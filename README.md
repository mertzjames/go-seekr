# go-seekr

A simple Go-based tool for scanning repositories and files for leaked secrets and potentially vulnerable environment variables.

## Overview

go-seekr is a security scanner that helps identify potentially leaked secrets and sensitive environment variables in your codebase. It uses pattern matching to detect over 400 common secret patterns including API keys, tokens, passwords, and other sensitive variables that could pose security risks if exposed.

The detection patterns are based on the comprehensive list from the [LinPEAS script](https://github.com/peass-ng/PEASS-ng/blob/master/linPEAS/builder/linpeas_parts/variables/pwd_in_variables.sh), covering a wide range of services including:

- Cloud providers (AWS, Google Cloud, Azure, DigitalOcean)
- CI/CD platforms (GitHub, GitLab, Travis CI, Jenkins)
- Communication services (Twilio, Slack, Discord)
- Database credentials
- API keys and tokens
- Docker and container secrets
- And many more...

## Features

- **Recursive Directory Scanning**: Scans all files in a specified directory and its subdirectories
- **Comprehensive Pattern Detection**: Detects over 400 types of potentially sensitive variables
- **Line Number Reporting**: Shows exact line numbers where potential secrets are found
- **Simple CLI Interface**: Easy-to-use command-line tool
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

2. Build the application:

```bash
go build -o seekr seekr.go
```

### Using Pre-built Binaries

Pre-built binaries are available in the `bin/` directory:

- `seekr-amd64-linux` - Linux AMD64
- `seekr-arm64-linux` - Linux ARM64

## Usage

### Basic Usage

Scan the current directory:

```bash
./seekr
```

Scan a specific directory:

```bash
./seekr -path /path/to/your/project
```

### Command-line Options

- `-path` : Specifies the path to scan (default: current directory ".")

### Example Output

```text
[INFO]  Scanning for potentially vulnerable variables in files under: /path/to/project
[INFO]  Scanning file:  /path/to/project/config.py
[VULN]    Found the following potential vulnerable variables:
[VULN]      - 005: AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
[VULN]      - 006: AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
[VULN]  Potentially vulnerable variables found in /path/to/project/config.py
[INFO]  Scanning file:  /path/to/project/deploy.sh
[INFO]    No vulnerable variables found.
[INFO]  Scan completed successfully.
```

## What It Detects

The tool scans for environment variables and configuration values that commonly contain sensitive information, including but not limited to:

- **AWS Credentials**: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
- **GitHub Tokens**: `GITHUB_TOKEN`, `GITHUB_API_KEY`
- **API Keys**: Various service API keys (Stripe, Twilio, SendGrid, etc.)
- **Database Credentials**: `DATABASE_PASSWORD`, `DB_USER`, `MYSQL_PASSWORD`
- **OAuth Secrets**: `CLIENT_SECRET`, `OAUTH_TOKEN`
- **SSH Keys**: `SSH_PRIVATE_KEY`, `id_rsa`
- **Docker Secrets**: `DOCKER_PASSWORD`, `DOCKER_TOKEN`
- **CI/CD Variables**: Travis, Jenkins, and other CI platform secrets

## Use Cases

- **Pre-commit Hooks**: Integrate into your Git workflow to catch secrets before they're committed
- **Security Audits**: Regular scanning of codebases for exposed credentials
- **CI/CD Pipeline Integration**: Automated secret detection in build processes
- **Code Reviews**: Manual verification during code review processes
- **Compliance**: Help meet security compliance requirements

## Important Notes

- This tool uses pattern matching and may produce false positives
- Always review the results manually to confirm actual secrets
- Consider using `.gitignore` or similar mechanisms to exclude files with legitimate environment variables
- This is a detection tool - it doesn't automatically remediate found secrets

## Future Enhancements

- Custom detection pattern support (see TODO in code)
- Configuration file support
- Integration with secret management services
- Additional output formats (JSON, CSV)
- Whitelist/blacklist functionality

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## License

See the [LICENSE](LICENSE) file for license information.
