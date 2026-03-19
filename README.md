# zohodesk-cli

A powerful command-line interface for Zoho Desk API, designed for both human interaction and AI agent integration.

## Overview

zohodesk-cli provides complete access to Zoho Desk's API through an intuitive command-line interface. Whether you're a support agent managing tickets, a developer automating workflows, or an AI agent integrating with Zoho Desk, this tool makes it simple.

**Key Features:**
- 🔐 **Simple Authentication** - One-time OAuth setup with automatic token refresh
- 📋 **Full Ticket Management** - Create, update, close, merge, assign tickets
- 💬 **Comments & Threads** - Add replies and private notes to tickets
- 🏷️ **Tags** - Organize tickets with custom tags
- 👥 **Contacts & Agents** - Manage contacts and support agents
- 📁 **Attachments** - Upload and download files
- 📝 **Knowledge Base** - Manage articles
- ✅ **Tasks** - Task management with priorities and due dates
- 🔄 **Batch Operations** - Process multiple tickets at once
- 📊 **Multiple Output Formats** - Human-readable tables, JSON, YAML

## Installation

### Download Binary (Recommended)

Download the latest release for your platform from [GitHub Releases](https://github.com/jrodriguezruibal/zohodesk-cli/releases):

```bash
# Extract and install
tar xzf zohodesk-cli-*.tar.gz
sudo mv zohodesk-cli /usr/local/bin/

# Verify installation
zohodesk-cli version
```

### Build from Source

```bash
git clone https://github.com/jrodriguezruibal/zohodesk-cli.git
cd zohodesk-cli
go build -o zohodesk-cli .
```

## Quick Start

### Step 1: Create Zoho API Credentials

1. Go to [Zoho API Console](https://api-console.zoho.com)
2. Click **"Add Client"** → Select **"Self-Client"**
3. Copy your **Client ID** and **Client Secret**
4. Get your **Organization ID** from Zoho Desk → Settings → Organization

### Step 2: Initialize Configuration

```bash
zohodesk-cli config init
```

Enter your Client ID, Client Secret, Organization ID, and region when prompted.

### Step 3: Authenticate (One-Time Setup)

Generate an authorization code in Zoho API Console:

1. In Zoho API Console, select your Self-Client
2. Go to **"Generate Code"** tab
3. Enter these scopes:
   ```
   Desk.tickets.ALL,Desk.contacts.READ,Desk.basic.READ,Desk.tasks.ALL,Desk.articles.READ,Desk.articles.CREATE,Desk.articles.UPDATE,Desk.articles.DELETE,Desk.search.READ,Desk.settings.READ
   ```
4. Set duration to 10 minutes
5. Click **"Create"**
6. Copy the generated code

Then authenticate:

```bash
zohodesk-cli config auth -a YOUR_AUTHORIZATION_CODE
```

**You're all set!** Tokens are automatically refreshed when they expire.

## Usage Examples

### Working with Tickets

```bash
# List tickets
zohodesk-cli tickets list
zohodesk-cli tickets list --status Open --priority High
zohodesk-cli tickets list --limit 50 --output json

# Create a ticket
zohodesk-cli tickets create \
  --subject "Cannot login to portal" \
  --description "User reports login issues" \
  --department DEPT_ID \
  --contact-id CONTACT_ID \
  --priority High

# Get ticket details
zohodesk-cli tickets get TICKET_ID
zohodesk-cli tickets get TICKET_ID --context  # includes related data

# Update a ticket
zohodesk-cli tickets update TICKET_ID --status "In Progress" --priority High

# Close a ticket
zohodesk-cli tickets close TICKET_ID --resolution "Issue resolved"

# Add a comment
zohodesk-cli comments note TICKET_ID --message "Investigating the issue"
zohodesk-cli comments reply TICKET_ID --message "Thank you for your patience"
```

### Working with Contacts

```bash
# List contacts
zohodesk-cli contacts list

# Search contacts
zohodesk-cli contacts search --email user@example.com
```

### Working with Tags

```bash
# Add tags to a ticket
zohodesk-cli tags add TICKET_ID --tag urgent --tag bug

# List tags on a ticket
zohodesk-cli tags list TICKET_ID
```

### Working with Agents

```bash
# List agents
zohodesk-cli agents list

# Assign ticket to agent
zohodesk-cli tickets assign TICKET_ID --agent AGENT_ID
```

### Working with Articles (Knowledge Base)

```bash
# List articles
zohodesk-cli articles list

# Search articles
zohodesk-cli articles search "password reset"
```

### Working with Tasks

```bash
# List tasks
zohodesk-cli tasks list

# Create a task
zohodesk-cli tasks create --title "Follow up with customer" --priority High
```

### Working with Attachments

```bash
# List attachments on a ticket
zohodesk-cli attachments list TICKET_ID

# Upload a file
zohodesk-cli attachments upload TICKET_ID /path/to/file.pdf

# Download an attachment
zohodesk-cli attachments download ATTACHMENT_ID --output /path/to/save.pdf
```

### Batch Operations

Process multiple tickets at once using JSON input:

```bash
# Create multiple tickets
echo '[
  {"subject": "Issue 1", "departmentId": "123", "contactId": "456"},
  {"subject": "Issue 2", "departmentId": "123", "contactId": "789"}
]' | zohodesk-cli tickets create --batch

# Close multiple tickets
echo '["TICKET_ID_1", "TICKET_ID_2"]' | zohodesk-cli tickets close --batch
```

### Output Formats

All commands support multiple output formats:

```bash
# Human-readable table (default)
zohodesk-cli tickets list --limit 10

# JSON (for scripting and AI agents)
zohodesk-cli tickets list --output json

# YAML
zohodesk-cli tickets list --output yaml
```

## Command Reference

| Command | Description |
|---------|-------------|
| `tickets` | Manage tickets (create, list, get, update, close, merge, assign) |
| `comments` | Add notes and replies to tickets |
| `tags` | Manage ticket tags |
| `contacts` | List and search contacts |
| `agents` | List support agents |
| `articles` | Knowledge base articles |
| `tasks` | Task management |
| `attachments` | File attachments |
| `products` | List products |
| `accounts` | Customer accounts |
| `config` | Manage configuration |

Use `zohodesk-cli [command] --help` for detailed usage information.

## Configuration

### Profiles

Manage multiple Zoho Desk organizations:

```bash
# Create new profile
zohodesk-cli config set --profile production \
  --client-id "PROD_CLIENT_ID" \
  --client-secret "PROD_SECRET" \
  --org-id "PROD_ORG_ID" \
  --region com

# Switch profiles
zohodesk-cli config use production

# List profiles
zohodesk-cli config list
```

### Regions

| Region | Flag | Zoho URL |
|--------|------|----------|
| United States | `--region com` | desk.zoho.com |
| Europe | `--region eu` | desk.zoho.eu |
| India | `--region in` | desk.zoho.in |
| China | `--region cn` | desk.zoho.com.cn |
| Australia | `--region au` | desk.zoho.com.au |

## For AI Agents

Designed for AI agent integration (Claude, GPT, etc.):

- All commands output structured JSON with `--output json`
- Clear error messages for debugging
- Batch operations for bulk processing
- No interactive prompts blocking execution
- Full ticket context with `--context` flag

Example context for AI agents:

```text
The zohodesk-cli tool provides complete access to Zoho Desk API. 

To create a ticket, you need:
1. A department ID (find in Zoho Desk settings or use a known ID)
2. A contact ID (search with `zohodesk-cli contacts search --email EMAIL`)

Example workflow:
1. Search for contact: zohodesk-cli contacts search --email user@example.com --output json
2. Extract contactId from response
3. Create ticket with required IDs: zohodesk-cli tickets create --subject "Issue" --department DEPT_ID --contact-id CONTACT_ID
```

## Troubleshooting

### Authentication Issues

**Error: "no valid authentication token found"**

Run authentication:
```bash
zohodesk-cli config auth -a YOUR_AUTHORIZATION_CODE
```

**Error: "invalid_code"**

The authorization code has expired (valid for 10 minutes). Generate a new code in Zoho API Console.

### Common Issues

**Ticket creation requires department and contact:**

```bash
# Find or search contact
zohodesk-cli contacts search --email user@example.com --output json

# Create ticket with IDs (department ID can be found in Zoho Desk settings)
zohodesk-cli tickets create \
  --subject "Issue" \
  --department DEPT_ID \
  --contact-id CONTACT_ID
```

**Empty responses:** Some fields may return empty from the Zoho Desk API (e.g., task titles, account names). This is expected behavior from Zoho.

## Contributing

Contributions are welcome! Please read our contributing guidelines before submitting PRs.

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `go test ./...`
5. Submit a pull request

## License

[MIT License](LICENSE)

## Support

- **Issues:** [GitHub Issues](https://github.com/jrodriguezruibal/zohodesk-cli/issues)
- **Zoho Desk API:** [Official Documentation](https://www.zoho.com/desk/api.html)