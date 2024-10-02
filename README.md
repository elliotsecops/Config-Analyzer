# Configuration Security Analyzer

## Overview

The Configuration Security Analyzer is a minimalist Go script designed to scan configuration files of common services such as Apache, Nginx, SSH, and MySQL for insecure or non-recommended configurations. This tool performs the following tasks:

1. **Scan Specific Directories**: Scans specified directories for known configuration files.
2. **Check for Insecure Patterns**: Scans each configuration file for insecure configuration patterns.
3. **Compare Against Best Practices**: Compares the configurations against a list of security best practices.

## Features

- **Severity Levels**: Each pattern has an associated severity level (Low, Medium, High).
- **Command-Line Flags**: Users can specify directories to scan and choose the output format (text or JSON).
- **Ignore Patterns**: Users can specify patterns to ignore during the scan.
- **Parallel Scanning**: Uses goroutines to scan directories and files concurrently, improving performance.
- **Minimalist Design**: All tasks are executed within a single script, maintaining simplicity and ease of use.

## Installation

### Prerequisites

- Go (version 1.16 or higher)

### Steps

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/elliotsecops/config_analyzer.git
   cd config_analyzer
   ```

2. **Initialize the Go Module**:
   ```sh
   go mod init config_analyzer
   ```

3. **Build the Script**:
   ```sh
   go build -o config_analyzer
   ```

4. **Run the Script**:
   ```sh
   ./config_analyzer -dir /etc/apache2 -dir /etc/nginx -json -ignore ignore_patterns.txt
   ```

## Usage

### Command-Line Flags

- **-dir**: Directories to scan (can be specified multiple times).
- **-json**: Output results in JSON format.
- **-ignore**: File containing patterns to ignore.

### Example Commands

1. **Scan Specific Directories and Output in JSON**:
   ```sh
   ./config_analyzer -dir /etc/apache2 -dir /etc/nginx -json -ignore ignore_patterns.txt
   ```

2. **Scan Default Directories and Output in Text**:
   ```sh
   ./config_analyzer
   ```

### Example `ignore_patterns.txt` Content

```
ServerTokens OS
ServerSignature On
```

## How It Works

### Severity Levels

Each pattern checked by the script is associated with a severity level (Low, Medium, High). This helps in prioritizing the findings based on their potential impact.

### Configuration Checks

The script defines a list of configuration checks for each service. Each check includes the filename of the configuration file and a list of patterns to look for. These patterns are associated with their respective severity levels.

### Main Function

The main function initializes the directories to be scanned and loads any ignore patterns specified by the user. It then uses goroutines to scan each directory concurrently, collecting findings in a thread-safe manner.

### Load Ignore Patterns

The script can load patterns to ignore from a file specified by the user. This allows users to exclude known false positives or intentionally insecure configurations.

### Scan Directory

The `scanDirectory` function walks through the specified directory and its subdirectories, looking for files that match the target configuration filenames. When a matching file is found, it calls the `scanFile` function to analyze its contents.

### Scan File

The `scanFile` function reads the specified configuration file line by line and checks for the presence of insecure configurations defined in the patterns. If a match is found and it is not in the ignore list, it records the finding.

### Output

The script supports two output formats:

- **Text Output**: Human-readable text format.
- **JSON Output**: Machine-parsable JSON format.

## Troubleshooting

1. **Verify Directory Existence**:
   - Ensure that the `/etc/apache2` and `/etc/nginx` directories exist on your system.
   - If these directories do not exist, you can either create them or adjust the directories you are scanning.

2. **Check Permissions**:
   - Ensure you have the necessary permissions to read the specified directories.

3. **Run the Script from the Correct Directory**:
   - Ensure you are in the directory where `config_analyzer.go` and `ignore_patterns.txt` are located.


This Configuration Security Analyzer is a powerful tool for identifying insecure configurations in common services. Its minimalist design and support for parallel scanning make it efficient and easy to use. By following the steps outlined in this README, you can effectively use this tool to enhance the security of your configurations.
