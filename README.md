# zohodesk-cli

A command-line interface for Zoho Desk API, designed for both human interaction and AI agent integration.

## Features

- **Multiple profiles**: Manage multiple Zoho Desk organizations
- **Multiple output formats**: JSON, YAML, and human-readable tables
- **TUI setup**: Interactive configuration wizard
- **Agent-friendly**: Designed for use by AI agents (Claude, GPT, etc.)
- **Self-Client OAuth**: No browser interaction required
- **Batch operations**: Accept JSON input via stdin for bulk operations

## Installation

### Binary Download (Recommended)

Download the latest release from [GitHub Releases](https://github.com/jrodriguezruibal/zohodesk-cli/releases):

```bash
# Linux (amd64)
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-linux-amd64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/

# Linux (arm64)
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-linux-arm64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/

# macOS (Apple Silicon)
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-darwin-arm64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/

# macOS (Intel)
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-darwin-amd64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/
```

### Homebrew

```bash
brew tap jrodriguezruibal/tap
brew install zohodesk-cli
```

### Go Install

```bash
go install github.com/jrodriguezruibal/zohodesk-cli@latest
```

### Docker

```bash
docker pull ghcr.io/jrodriguezruibal/zohodesk-cli:latest

# Run with mounted config
docker run --rm \
  -v ~/.config/zohodesk-cli:/root/.config/zohodesk-cli \
  ghcr.io/jrodriguezruibal/zohodesk-cli:latest tickets list
```

## Quick Start

### 1. Get Zoho API Credentials

1. Go to [Zoho API Console](https://api-console.zoho.com)
2. Create a new client with type **Self-Client**
3. Note your **Client ID** and **Client Secret**
4. Get your **Organization ID** from Zoho Desk Settings → Company → Organization

### 2. Initialize Configuration

```bash
# Interactive setup (recommended)
zohodesk-cli config init

# Or set environment variables for CI/CD
export ZOHO_CLIENT_ID="1000.xxxxx"
export ZOHO_CLIENT_SECRET="xxxxx"
export ZOHO_ORG_ID="12345678"
export ZOHO_REGION="com"  # Optional: com, eu, in, cn, au
```

###3. Use the CLI

```bash
# List tickets
zohodesk-cli tickets list

# List open tickets
zohodesk-cli tickets list --status Open

# Get ticket details
zohodesk-cli tickets get 123456789

# Create a ticket
zohodesk-cli tickets create \
  --subject "Login issue" \
  --description "User cannot login to portal" \
  --email user@example.com \
  --priority High

# Search tickets
zohodesk-cli tickets search --email user@example.com

# For AI agents (JSON output)
zohodesk-cli tickets list --output json
```

## Documentation

- [Installation Guide](./docs/INSTALLATION.md) - Detailed installation instructions
- [Configuration Guide](./docs/CONFIGURATION.md) - Configuration and profiles
- [Usage Examples](./docs/USAGE.md) - Common usage patterns
- [API Reference](./docs/API.md) - Complete command reference

## For AI Agents

This CLI is designed to be used by AI agents (Claude, GPT, etc.):

```bash
# All commands support JSON output
zohodesk-cli tickets list --output json

# Get ticket details as JSON
zohodesk-cli tickets get 123456789 --output json

# Create ticket and get JSON response
zohodesk-cli tickets create \
  --subject "Bug report" \
  --email user@example.com \
  --output json
```

### JSON Output Example

```json
[
  {
    "id": "123456789",
    "ticketNumber": "TKT-001",
    "subject": "Login issue",
    "status": "Open",
    "priority": "High",
    "createdTime": "2024-01-15T10:30:00Z",
    "contactEmail": "user@example.com"
  }
]
```

## Authentication

zohodesk-cli uses **Self-Client OAuth** which requires no browser interaction:

1. Access tokens are obtained automatically using client credentials
2. Tokens are cached locally in `~/.config/zohodesk-cli/tokens/`
3. Tokens are refreshed automatically when expired

## Regions

| Region | Base URL |
|--------|----------|
| com | https://desk.zoho.com |
| eu | https://desk.zoho.eu |
| in | https://desk.zoho.in |
| cn | https://desk.zoho.com.cn |
| au | https://desk.zoho.com.au |

## Development

```bash
# Clone the repository
git clone https://github.com/jrodriguezruibal/zohodesk-cli.git
cd zohodesk-cli

# Build
make build

# Install locally
make install

# Run tests
make test
```

## License

[MIT License](./LICENSE)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Support

- [GitHub Issues](https://github.com/jrodriguezruibal/zohodesk-cli/issues)
- [Zoho Desk API Documentation](https://www.zoho.com/desk/api.html)