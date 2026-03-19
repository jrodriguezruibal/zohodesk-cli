# zohodesk-cli

A command-line interface for Zoho Desk API, designed for both human interaction and AI agent integration.

## Features

- **Multiple profiles**: Manage multiple Zoho Desk organizations
- **Multiple output formats**: JSON, YAML, and human-readable tables
- **TUI setup**: Interactive configuration wizard
- **Agent-friendly**: Designed for use by AI agents (Claude, GPT, etc.)
- **Self-Client OAuth**: No browser interaction required
- **Batch operations**: Create, update, or close multiple tickets from JSON input
- **Comments/Replies**: Add public replies and private notes to tickets
- **Ticket assignment**: Assign tickets to agents and departments
- **Context enrichment**: Get tickets with full context (department, assignee, SLA)
- **Departments & Agents**: List and manage departments and agents
- **Time tracking**: Log time on tickets and generate reports
- **Attachments**: Upload, download, and manage file attachments
- **Knowledge Base**: Manage articles and categories
- **Tasks**: Manage tasks with due dates, priorities, and statuses
- **Reports**: Ticket statistics, agent performance, and SLA compliance
- **Products & Accounts**: Manage products and customer accounts
- **Rate limiting**: Built-in rate limiting for batch operations

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
```

### 3. Authenticate (Self-Client OAuth) - One-Time Setup

Zoho Desk uses OAuth2 with Self-Client for API access. You need to generate an authorization code **once**. After that, tokens are refreshed automatically.

**Step 1: Create Self-Client in Zoho API Console**

1. Go to [Zoho API Console](https://api-console.zoho.com)
2. Click **"Add Client"** → Select **"Self-Client"**
3. Note your **Client ID** and **Client Secret**
4. Get your **Organization ID** from Zoho Desk → Settings → Organization

**Step 2: Generate Authorization Code**

1. In Zoho API Console, select your Self-Client
2. Go to **"Generate Code"** tab
3. Enter the following scopes:

```
Desk.tickets.ALL,Desk.contacts.READ,Desk.basic.READ,Desk.tasks.ALL,Desk.articles.READ,Desk.articles.CREATE,Desk.articles.UPDATE,Desk.articles.DELETE,Desk.search.READ,Desk.settings.READ
```

4. Set **Code Expiry Duration** to 10 minutes (or more)
5. Click **"Create"**
6. Copy the generated authorization code

**Step 3: Authenticate**

```bash
zohodesk-cli config auth -a YOUR_AUTHORIZATION_CODE
```

**That's it!** After the initial authentication:
- Access tokens are automatically refreshed when they expire
- You only need to re-authenticate if the refresh token expires (rare)
- Tokens are stored locally in `~/.config/zohodesk-cli/tokens/`

### 4. Use the CLI

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

## Batch Operations

Create, update, or close multiple tickets from JSON input:

```bash
# Create multiple tickets from JSON
cat tickets.json | zohodesk-cli tickets create --batch

# Create from file
zohodesk-cli tickets create --batch --file tickets.json

# Update multiple tickets
cat updates.json | zohodesk-cli tickets update --batch

# Close multiple tickets
echo '["123", "456", "789"]' | zohodesk-cli tickets close --batch
```

### Batch JSON Format

**Create tickets:**
```json
[
  {"subject": "Ticket 1", "email": "user1@example.com", "priority": "High"},
  {"subject": "Ticket 2", "email": "user2@example.com", "priority": "Medium"}
]
```

**Update tickets:**
```json
[
  {"id": "123", "status": "Closed"},
  {"id": "456", "priority": "High"}
]
```

**Close tickets:**
```json
["123", "456", "789"]
```

## Comments and Replies

Add public replies or private notes to tickets:

```bash
# List comments on a ticket
zohodesk-cli comments list <ticket-id>

# Add a public reply
zohodesk-cli comments reply <ticket-id> --message "Thank you for your report"

# Add a private note
zohodesk-cli comments note <ticket-id> --message "Internal note for the team"
```

## Ticket Assignment

Assign tickets to agents and departments:

```bash
# Assign to an agent
zohodesk-cli tickets assign <ticket-id> --agent <agent-id>

# Assign to a department
zohodesk-cli tickets assign <ticket-id> --department <dept-id>

# Assign to both
zohodesk-cli tickets assign <ticket-id> --agent <agent-id> --department <dept-id>
```

## Ticket Context

Get tickets with full context (department, assignee, SLA):

```bash
# Get ticket with enriched context
zohodesk-cli tickets get <ticket-id> --context --output json
```

## Departments

Manage Zoho Desk departments:

```bash
# List all departments
zohodesk-cli departments list

# Get department details
zohodesk-cli departments get <dept-id>
```

## Agents

Manage Zoho Desk agents:

```bash
# List all agents
zohodesk-cli agents list

# Get agent details
zohodesk-cli agents get <agent-id>
```

## Time Tracking

Log time spent on tickets:

```bash
# List time entries for a ticket
zohodesk-cli time list <ticket-id>

# Add time entry (1 hour 30 minutes)
zohodesk-cli time add <ticket-id> --duration 1h30m --description "Investigating issue"

# Add time for another agent
zohodesk-cli time add <ticket-id> --duration 2h --agent <agent-id> --description "Code review"

# View time report
zohodesk-cli time report

# Filter report by agent
zohodesk-cli time report --agent <agent-id>

# Filter by date range
zohodesk-cli time report --from 2024-01-01 --to 2024-01-31
```

## Attachments

Manage file attachments on tickets:

```bash
# List attachments for a ticket
zohodesk-cli attachments list <ticket-id>

# Upload a file
zohodesk-cli attachments upload <ticket-id> /path/to/file.pdf

# Download an attachment
zohodesk-cli attachments download <attachment-id> --output /path/to/save.pdf

# Delete an attachment
zohodesk-cli attachments delete <attachment-id>
```

## Knowledge Base

Manage knowledge base articles and categories:

```bash
# List articles
zohodesk-cli articles list

# Search articles
zohodesk-cli articles search "password reset"

# Create article
zohodesk-cli articles create --title "FAQ" --content "Common questions..."

# Update article
zohodesk-cli articles update <article-id> --status "Published"

# List categories
zohodesk-cli categories list

# Get category
zohodesk-cli categories get <category-id>
```

## Tasks

Manage tasks associated with tickets:

```bash
# List all tasks
zohodesk-cli tasks list

# List tasks for a ticket
zohodesk-cli tasks list --ticket <ticket-id>

# Create task
zohodesk-cli tasks create --title "Follow up" --priority "High" --due 2024-01-15

# Create task for a ticket
zohodesk-cli tasks create --title "Review" --ticket <ticket-id>

# Update task
zohodesk-cli tasks update <task-id> --status "In Progress"

# Complete task
zohodesk-cli tasks complete <task-id>

# Delete task
zohodesk-cli tasks delete <task-id>
```

## Reports

Generate statistics and analytics:

```bash
# Ticket statistics
zohodesk-cli reports tickets

# Ticket stats by date range
zohodesk-cli reports tickets --from 2024-01-01 --to 2024-01-31

# Agent performance
zohodesk-cli reports agents

# SLA compliance
zohodesk-cli reports sla

# SLA by department
zohodesk-cli reports sla --department <dept-id>
```

## Products

View products in your helpdesk:

```bash
# List products
zohodesk-cli products list

# Get product details
zohodesk-cli products get <product-id>
```

## Accounts

Manage customer accounts:

```bash
# List accounts
zohodesk-cli accounts list

# Get account details
zohodesk-cli accounts get <account-id>

# Create account
zohodesk-cli accounts create --name "Acme Corp" --email "contact@acme.com"

# Update account
zohodesk-cli accounts update <account-id> --type "Customer"

# Delete account
zohodesk-cli accounts delete <account-id>
```

## Tags

Manage tags on tickets:

```bash
# List tags on a ticket
zohodesk-cli tickets tags list <ticket-id>

# Add tags
zohodesk-cli tickets tags add <ticket-id> --tag urgent --tag bug

# Remove tag
zohodesk-cli tickets tags remove <ticket-id> urgent
```

## Ticket Operations

```bash
# Merge tickets
zohodesk-cli tickets merge <source-id> <target-id>

# Follow/unfollow tickets
zohodesk-cli tickets follow <ticket-id>
zohodesk-cli tickets unfollow <ticket-id>

# Update with custom fields
zohodesk-cli tickets update <ticket-id> --custom-fields '{"field": "value"}'
```

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