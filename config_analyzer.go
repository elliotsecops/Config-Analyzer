package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Severity string

const (
	Low    Severity = "Low"
	Medium Severity = "Medium"
	High   Severity = "High"
)

type ConfigPattern struct {
	pattern  string
	severity Severity
}

type ConfigCheck struct {
	filename string
	patterns []ConfigPattern
}

type Finding struct {
	File     string   `json:"file"`
	Line     int      `json:"line"`
	Pattern  string   `json:"pattern"`
	Content  string   `json:"content"`
	Severity Severity `json:"severity"`
}

var configChecks = []ConfigCheck{
	{
		filename: "apache2.conf",
		patterns: []ConfigPattern{
			{"ServerTokens OS", Medium},
			{"ServerSignature On", Medium},
			{"TraceEnable On", High},
			{"AllowOverride All", Medium},
			{"Options All", Medium},
		},
	},
	{
		filename: "nginx.conf",
		patterns: []ConfigPattern{
			{"server_tokens on", Medium},
			{"autoindex on", Low},
			{"ssl_protocols TLSv1 TLSv1.1", High},
		},
	},
	{
		filename: "sshd_config",
		patterns: []ConfigPattern{
			{"PermitRootLogin yes", High},
			{"PasswordAuthentication yes", Medium},
			{"X11Forwarding yes", Low},
			{"PermitEmptyPasswords yes", High},
		},
	},
	{
		filename: "my.cnf",
		patterns: []ConfigPattern{
			{"skip-networking", Low},
			{"bind-address = 0.0.0.0", Medium},
			{"local-infile=1", Medium},
		},
	},
}

type flagArray []string

func (i *flagArray) String() string {
	return strings.Join(*i, ", ")
}

func (i *flagArray) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	directories flagArray
	jsonOutput  bool
	ignoreFile  string
)

func main() {
	flag.Var(&directories, "dir", "Directories to scan (can be specified multiple times)")
	flag.BoolVar(&jsonOutput, "json", false, "Output results in JSON format")
	flag.StringVar(&ignoreFile, "ignore", "", "File containing patterns to ignore")
	flag.Parse()

	if len(directories) == 0 {
		directories = []string{"/etc/apache2", "/etc/nginx", "/etc/ssh", "/etc/mysql"}
	}

	var ignorePatterns []string
	if ignoreFile != "" {
		ignorePatterns = loadIgnorePatterns(ignoreFile)
	}

	var findings []Finding
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, dir := range directories {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			dirFindings := scanDirectory(dir, ignorePatterns)
			mu.Lock()
			findings = append(findings, dirFindings...)
			mu.Unlock()
		}(dir)
	}

	wg.Wait()

	if jsonOutput {
		outputJSON(findings)
	} else {
		outputText(findings)
	}
}

func loadIgnorePatterns(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening ignore file %s: %v\n", filename, err)
		return nil
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading ignore file %s: %v\n", filename, err)
	}

	return patterns
}

func scanDirectory(dir string, ignorePatterns []string) []Finding {
	var findings []Finding

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for _, check := range configChecks {
				if strings.HasSuffix(path, check.filename) {
					fmt.Printf("Scanning file: %s\n", path)
					findings = append(findings, scanFile(path, check.patterns, ignorePatterns)...)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory %s: %v\n", dir, err)
	}

	return findings
}

func scanFile(filepath string, patterns []ConfigPattern, ignorePatterns []string) []Finding {
	var findings []Finding

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filepath, err)
		return findings
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		for _, pattern := range patterns {
			if strings.Contains(line, pattern.pattern) && !containsPattern(ignorePatterns, line) {
				findings = append(findings, Finding{
					File:     filepath,
					Line:     lineNum,
					Pattern:  pattern.pattern,
					Content:  line,
					Severity: pattern.severity,
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filepath, err)
	}

	return findings
}

func containsPattern(patterns []string, line string) bool {
	for _, pattern := range patterns {
		if strings.Contains(line, pattern) {
			return true
		}
	}
	return false
}

func outputText(findings []Finding) {
	for _, finding := range findings {
		fmt.Printf("[%s] %s:%d - %s\n", finding.Severity, finding.File, finding.Line, finding.Pattern)
		fmt.Printf("  %s\n", finding.Content)
	}
}

func outputJSON(findings []Finding) {
	json, err := json.MarshalIndent(findings, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		return
	}
	fmt.Println(string(json))
}