# Installation Guide

## Requirements

- Zoho Desk account with API access
- Self-Client OAuth credentials (see Configuration Guide)

## Installation Methods

### 1. Binary Download (Recommended)

Download the latest release for your platform:

#### Linux (amd64)

```bash
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-linux-amd64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/
sudo chmod +x /usr/local/bin/zohodesk-cli
```

#### Linux (arm64)

```bash
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-linux-arm64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/
sudo chmod +x /usr/local/bin/zohodesk-cli
```

#### macOS (Apple Silicon - M1/M2/M3)

```bash
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-darwin-arm64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/
sudo chmod +x /usr/local/bin/zohodesk-cli
```

#### macOS (Intel)

```bash
curl -sL https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-darwin-amd64.tar.gz | tar xz
sudo mv zohodesk-cli /usr/local/bin/
sudo chmod +x /usr/local/bin/zohodesk-cli
```

#### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://github.com/jrodriguezruibal/zohodesk-cli/releases/latest/download/zohodesk-cli-windows-amd64.zip -OutFile zohodesk-cli.zip
Expand-Archive zohodesk-cli.zip
Move-Item zohodesk-cli/zohodesk-cli.exe C:\Windows\
```

### 2. Homebrew (macOS/Linux)

```bash
brew tap jrodriguezruibal/tap
brew install zohodesk-cli
```

###3. Go Install

```bash
go install github.com/jrodriguezruibal/zohodesk-cli@latest

# Binary will be at ~/go/bin/zohodesk-cli
# Add to PATH:
export PATH=$PATH:~/go/bin
```

### 4. Docker

```bash
# Pull the image
docker pull ghcr.io/jrodriguezruibal/zohodesk-cli:latest

# Create alias for convenience
alias zohodesk-cli='docker run --rm -v ~/.config/zohodesk-cli:/root/.config/zohodesk-cli ghcr.io/jrodriguezruibal/zohodesk-cli:latest'

# Use normally
zohodesk-cli tickets list
```

### 5. Package Managers

#### Debian/Ubuntu (.deb)

```bash
sudo dpkg -i zohodesk-cli_*_linux_amd64.deb
```

#### RHEL/CentOS/Fedora (.rpm)

```bash
sudo rpm -i zohodesk-cli_*_linux_amd64.rpm
```

#### Arch Linux (AUR)

```bash
yay -s zohodesk-cli-bin
```

## Verify Installation

```bash
zohodesk-cli version
```

Output:

```
zohodesk-cli v0.1.0
  Build time: 2024-01-15T10:30:00Z
  Go version: go1.21.5
  OS/Arch:linux/amd64
```

## Next Steps

- [Configuration Guide](./CONFIGURATION.md) - Set up your Zoho credentials
- [Usage Examples](./USAGE.md) - Learn how to use the CLI