# Analizador de configuraciones de seguridad

Este Configuration Security Analyzer es un script minimalista hecho en Go diseñado para escanear archivos de configuración de servicios comunes como Apache, Nginx, SSH y MySQL en busca de configuraciones inseguras o no recomendadas. Esta herramienta realiza las siguientes tareas:

1. **Scan Specific Directories**: Escanea directorios especificados en busca de archivos de configuración conocidos.
2. **Check for Insecure Patterns**: Escanea cada archivo de configuración en busca de patrones de configuración inseguros.
3. **Compare Against Best Practices**: Compara las configuraciones con una lista de mejores prácticas de seguridad.

## Caracteristicas

- **Severity Levels**: Cada patrón tiene un nivel de severidad asociado (Low, Medium, High).
- **Command-Line Flags**: Los usuarios pueden especificar directorios para escanear y elegir el formato de salida (texto o JSON).
- **Ignore Patterns**: Los usuarios pueden especificar patrones para ignorar durante el escaneo.
- **Parallel Scanning**: Usa goroutines para escanear directorios y archivos de manera concurrente, mejorando el rendimiento.
- **Minimalist Design**: Todas las tareas se ejecutan dentro de un solo script, manteniendo la simplicidad y facilidad de uso.

## Instalaciónn

### Prerequisitos

- Go (version 1.16 or higher)

### Pasos

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

- **-dir**: Directorios a escanear (pueden especificarse múltiples veces).
- **-json**: Salida de resultados en formato JSON.
- **-ignore**: Archivo que contiene patrones a ignorar.

### Comandos de ejemplo

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

## ¿Cómo funciona?

### Niveles de severidad

Cada patrón verificado por el script está asociado con un nivel de severidad (Low, Medium, High). Esto ayuda a priorizar los hallazgos según su impacto potencial.

### Configuration Checks

El script define una lista de verificaciones de configuración para cada servicio. Cada verificación incluye el nombre del archivo de configuración y una lista de patrones a buscar. Estos patrones están asociados con sus respectivos niveles de severidad.

### Main Function

La función principal inicializa los directorios a escanear y carga cualquier patrón de ignorar especificado por el usuario. Luego usa goroutines para escanear cada directorio de manera concurrente, recolectando hallazgos de manera segura.

### Load Ignore Patterns

El script puede cargar patrones a ignorar desde un archivo especificado por el usuario. Esto permite a los usuarios excluir falsos positivos conocidos o configuraciones intencionalmente inseguras.

### Scan Directory

La función `scanDirectory` recorre el directorio especificado y sus subdirectorios, buscando archivos que coincidan con los nombres de archivos de configuración objetivo. Cuando se encuentra un archivo coincidente, llama a la función `scanFile` para analizar su contenido.

### Scan File

La función `scanFile` lee el archivo de configuración especificado línea por línea y verifica la presencia de configuraciones inseguras definidas en los patrones. Si se encuentra una coincidencia y no está en la lista de ignorar, registra el hallazgo.

### Output

El script soporta dos formatos de salida:

- **Text Output**: Formato de texto legible por humanos.
- **JSON Output**: Formato JSON analizable por máquinas.

## Troubleshooting

1. **Verify Directory Existence**:
   - Asegúrate de que los directorios `/etc/apache2` y `/etc/nginx` existan en tu sistema.
   - Si estos directorios no existen, puedes crearlos o ajustar los directorios que estás escaneando.

2. **Check Permissions**:
   - Asegúrate de tener los permisos necesarios para leer los directorios especificados.

3. **Run the Script from the Correct Directory**:
   - Asegúrate de estar en el directorio donde se encuentran `config_analyzer.go` y `ignore_patterns.txt`.

Este Configuration Security Analyzer es una herramienta poderosa para identificar configuraciones inseguras en servicios comunes. Su diseño minimalista y soporte para escaneo paralelo lo hacen eficiente y fácil de usar. Siguiendo los pasos descritos en este README, puedes utilizar esta herramienta de manera efectiva para mejorar la seguridad de tus configuraciones.

---

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
