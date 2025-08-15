# go-seekr

A simple Go-based tool for scanning repositories and files for leaked secrets and potentially vulnerable environment variables.

## Overview

go-seekr is a security scanner that helps identify potentially leaked secrets and sensitive environment variables in your codebase. It uses pattern matching to detect over 400 common secret patterns including API keys, tokens, passwords, and other sensitive variables that could pose security risks if exposed.

The detection patterns are embedded directly in the application using a comprehensive list based on the [LinPEAS script](https://github.com/peass-ng/PEASS-ng/blob/master/linPEAS/builder/linpeas_parts/variables/pwd_in_variables.sh), covering a wide range of services including:

- Cloud providers (AWS, Google Cloud, Azure, DigitalOcean)
- CI/CD platforms (GitHub, GitLab, Travis CI, Jenkins)
- Communication services (Twilio, Slack, Discord)
- Database credentials
- API keys and tokens
- Docker and container secrets
- And many more...

## Features

- **Recursive Directory Scanning**: Scans all files in a specified directory and its subdirectories
- **Comprehensive Pattern Detection**: Detects over 400 types of potentially sensitive variables using embedded pattern database
- **Line Number Reporting**: Shows exact line numbers where potential secrets are found
- **Custom Variable Support**: Add your own variable names to scan for using the `--vars` flag or file-based input with `--vars_file`
- **Custom Regex Patterns**: Define custom regular expressions for advanced pattern matching via `--regex_str` or `--regex_file`
- **Flexible Input Methods**: Support both command-line arguments and file-based input for scalable custom pattern management
- **Default Pattern Override**: Option to ignore built-in patterns and use only custom definitions with `--ignore_default`
- **SSH Private Key Detection**: Automatically detects private SSH keys in PEM format
- **Language-Specific Scanning**: Filter scanning by programming language or file type
- **Binary File Support**: Optional scanning of binary files with text extraction
- **Simple CLI Interface**: Easy-to-use command-line tool with intuitive flags
- **Cross-platform**: Works on Linux, macOS, and Windows

## Installation

### Prerequisites

- Go 1.24.6 or later

### Building from Source

1. Clone the repository:

```bash
git clone https://github.com/mertzjames/go-seekr.git
cd go-seekr
```

1. Build the application:

```bash
go build -o seekr main.go
```

### Using Pre-built Binaries

Pre-built binaries are available in the [latest GitHub release](https://github.com/mertzjames/go-seekr/releases/latest):

**Linux:**

- `seekr-amd64-linux` - Linux AMD64
- `seekr-arm64-linux` - Linux ARM64

**macOS:**

- `seekr-amd64-darwin` - macOS AMD64 (Intel)
- `seekr-arm64-darwin` - macOS ARM64 (Apple Silicon)

**Windows:**

- `seekr-amd64-windows.exe` - Windows AMD64
- `seekr-arm64-windows.exe` - Windows ARM64

Download the appropriate binary for your platform from the release page and make it executable:

**Linux/macOS:**

```bash
chmod +x seekr-amd64-linux
./seekr-amd64-linux
```

**Windows:**

```cmd
seekr-amd64-windows.exe
```

## Usage

The application provides a simple command-line interface for scanning files and directories for secrets.

### Basic Usage

Scan the current directory for secrets supporting the most common programming language source files:

```bash
./seekr
```

Scan a specific directory:

```bash
./seekr --path /path/to/your/project
```

Scan only specific programming languages:

```bash
./seekr --path /path/to/your/project --language python,javascript
```

Scan all text based files:

```bash
./seekr --path /path/to/your/project --all_files
```

Scan all files including binary files:

```bash
./seekr --path /path/to/your/project --all_files --binary_check
```

Scan with custom variables:

```bash
./seekr --path /path/to/your/project --vars "MY_SECRET_VAR,CUSTOM_API_KEY"
```

Scan with custom regex pattern:

```bash
./seekr --path /path/to/your/project --regex_str "secret[_-]?key.*=.*"
```

Load custom variables from a file:

```bash
./seekr --path /path/to/your/project --vars_file /path/to/custom_vars.txt
```

Load custom regex patterns from a file:

```bash
./seekr --path /path/to/your/project --regex_file /path/to/custom_patterns.txt
```

Use only custom variables and regex patterns (ignore defaults):

```bash
./seekr --path /path/to/your/project --vars_file custom_vars.txt --ignore_default
```

### Command-line Options

- `-p, --path` : Specifies the path to scan (default: current directory ".")
- `-l, --language` : Programming language(s) to scan for secrets. Comma-separated list or "all" for all languages
- `-a, --all_files` : Include all files in the scan, regardless of file extension (overrides language flag)
- `-b, --binary_check` : Include binary files in the scan (only effective when used with --all_files)
- `-v, --vars` : Comma-separated list of additional variables to include in the scan
- `-r, --regex_str` : User-defined regular expression for matching custom/unsupported secrets
- `-V, --vars_file` : Path to a file containing additional variables to include in the scan, one per line
- `-R, --regex_file` : Path to a file containing user-defined regular expressions for matching custom/unsupported secrets, one per line
- `-i, --ignore_default` : Ignore the default set of vulnerable variables and only use user-defined variables and regex patterns

### File-based Input Formats

#### Variables File Format

When using the `--vars_file` flag, create a text file with one variable name per line. Lines starting with `#` are treated as comments and ignored:

```text
# Custom variables for my organization
MY_SECRET_API_KEY
CUSTOM_DATABASE_PASSWORD
INTERNAL_SERVICE_TOKEN
# Add more variables below
COMPANY_SPECIFIC_SECRET
```

#### Regex File Format

When using the `--regex_file` flag, create a text file with one regular expression per line. Lines starting with `#` are treated as comments and ignored:

```text
# Custom regex patterns
secret[_-]?key.*=.*
password[_-]?(hash|pwd).*=.*
# API key patterns
api[_-]?key[_-]?\w*.*=.*
```

**Note**: Regular expressions are automatically case-insensitive and wrapped with additional matching logic. You only need to provide the core pattern.

### ⚠️ Binary File Scanning Warning

**Important**: When using the `--binary_check` flag, be aware of the following potential issues:

- **Performance Impact**: Scanning binary files can significantly slow down the scanning process, especially for large files
- **System Instability**: Processing certain binary files may cause system instability or unexpected behavior
- **High Memory Usage**: Large binary files can consume substantial memory during text extraction
- **False Positives**: Binary files may contain byte sequences that match secret patterns but are not actual secrets
- **Limited Effectiveness**: The tool only scans for embedded text-based secrets within binary files

**Recommendation**: Use binary file scanning sparingly and only when necessary. Consider excluding large binary files or known safe binaries from your scan path.

### Supported Languages

The tool supports scanning files for the following programming languages. Use the exact language names shown below with the `--language` flag:

**Web Development:**

- `javascript` (.js, .mjs, .cjs)
- `typescript` (.ts, .tsx)
- `php` (.php, .phtml)

**Backend & Systems:**

- `python` (.py, .pyw, .pyi)
- `java` (.java, .class, .jar)
- `csharp` (.cs, .csx)
- `go` (.go)
- `ruby` (.rb, .rbw)
- `rust` (.rs)
- `c` (.c, .h)
- `cpp` (.cpp, .hpp, .cc, .h)

**Mobile Development:**

- `swift` (.swift)
- `kotlin` (.kt, .kts)
- `dart` (.dart)
- `objc` (.m, .h)

**Scripting & Shell:**

- `shell` (.sh, .bash, .zsh)
- `powershell` (.ps1, .psm1)
- `perl` (.pl, .pm)
- `lua` (.lua)

**Data & Configuration:**

- `sql` (.sql)
- `r` (.r, .R)
- `matlab` (.m)
- `yaml` (.yaml, .yml)
- `xml` (.xml)
- `json` (.json)

**Other Languages:**

- `scala` (.scala, .sc)

**Usage Examples:**

```bash
# Scan only Python files
./seekr --path . --language python

# Scan multiple languages
./seekr --path . --language python,javascript,go

# Scan all supported languages
./seekr --path . --language all
```

Use `--language all` or omit the language flag to scan all supported file types.

### Example Output

```text
[INFO]    Scanning directory: /path/to/project
[INFO]    Scanning file: /path/to/project/config.py
[VULN]    Found the following potential vulnerable variables:
[VULN]      - 005: AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
[VULN]      - 006: AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
[INFO]    Scanning file: /path/to/project/deploy.sh
[INFO]    No vulnerable variables found.
```

## What It Detects

The tool scans for environment variables and configuration values that commonly contain sensitive information, including but not limited to:

- **AWS Credentials**: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
- **GitHub Tokens**: `GITHUB_TOKEN`, `GITHUB_API_KEY`
- **API Keys**: Various service API keys (Stripe, Twilio, SendGrid, etc.)
- **Database Credentials**: `DATABASE_PASSWORD`, `DB_USER`, `MYSQL_PASSWORD`
- **OAuth Secrets**: `CLIENT_SECRET`, `OAUTH_TOKEN`
- **SSH Private Keys**: Automatically detects PEM-formatted private keys (`-----BEGIN PRIVATE KEY-----`)
- **Docker Secrets**: `DOCKER_PASSWORD`, `DOCKER_TOKEN`
- **CI/CD Variables**: Travis, Jenkins, and other CI platform secrets
- **Custom Variables**: User-defined variable names via `--vars` flag or `--vars_file` for file-based input
- **Custom Patterns**: User-defined regex patterns via `--regex_str` flag or `--regex_file` for file-based input

### Advanced Detection Features

- **Case-Insensitive Matching**: All pattern matching is performed case-insensitively
- **SSH Key Detection**: Recognizes various private key formats (RSA, DSA, ECDSA, etc.)
- **Custom Variable Lists**: Specify additional variables to scan for with comma-separated lists or file input
- **Regex Pattern Matching**: Define custom regular expressions for organization-specific secrets via command line or file input
- **Flexible Input Methods**: Support both command-line arguments and file-based input for custom variables and patterns
- **Default Pattern Override**: Use `--ignore_default` to scan only for custom patterns, ignoring the built-in pattern database

## Use Cases

- **Pre-commit Hooks**: Integrate into your Git workflow to catch secrets before they're committed
- **Security Audits**: Regular scanning of codebases for exposed credentials
- **CI/CD Pipeline Integration**: Automated secret detection in build processes with custom organization-specific patterns
- **Code Reviews**: Manual verification during code review processes
- **Compliance**: Help meet security compliance requirements
- **Organization-Specific Scanning**: Use file-based custom variables and patterns to detect company-specific sensitive data
- **Scalable Pattern Management**: Maintain and version-control custom detection patterns using file-based inputs

## Important Notes

- This tool uses pattern matching and may produce false positives
- Always review the results manually to confirm actual secrets
- Consider using `.gitignore` or similar mechanisms to exclude files with legitimate environment variables
- This is a detection tool - it doesn't automatically remediate found secrets

## Future Enhancements

- Configuration file support
- Integration with secret management services
- Additional output formats (JSON, CSV)
- Whitelist/blacklist functionality
- Performance optimizations for large repositories

## AI Disclosure

This README was generated with assistance from Claude (Anthropic's AI assistant). Additionally, other AI tools were used to assist in the creation of this tool. However, no code was directly copy-pasted from AI sources - all code was written and reviewed by human developers.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## License

See the [LICENSE](LICENSE) file for license information.
