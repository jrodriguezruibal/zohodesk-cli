# Usage Examples

## Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--config` | `-c` | Config file path |
| `--profile` | `-p` | Profile to use |
| `--output` | `-o` | Output format: `json`, `yaml`, `table` |
| `--help` | `-h` | Show help |

## Tickets

### List Tickets

```bash
# List all tickets (default: 50)
zohodesk-cli tickets list

# List more tickets
zohodesk-cli tickets list --limit 100

# Filter by status
zohodesk-cli tickets list --status Open
zohodesk-cli tickets list --status Closed
zohodesk-cli tickets list --status "On Hold"

# Filter by priority
zohodesk-cli tickets list --priority High
zohodesk-cli tickets list --priority Medium
zohodesk-cli tickets list --priority Low

# Combine filters
zohodesk-cli tickets list --status Open --priority High --limit 10

# JSON output
zohodesk-cli tickets list --output json

# YAML output
zohodesk-cli tickets list --output yaml
```

### Get Ticket Details

```bash
# Basic ticket info
zohodesk-cli tickets get 123456789

# Full ticket with threads and contact
zohodesk-cli tickets get 123456789 --full

# JSON output
zohodesk-cli tickets get 123456789 --output json
```

### Create Ticket

```bash
# Minimum required fields
zohodesk-cli tickets create \
  --subject "Login issue" \
  --email user@example.com

# With description
zohodesk-cli tickets create \
  --subject "Bug report" \
  --description "User cannot login after password change" \
  --email user@example.com

# With priority
zohodesk-cli tickets create \
  --subject "Critical bug" \
  --description "System is down" \
  --email admin@example.com \
  --priority High

# JSON output
zohodesk-cli tickets create \
  --subject "Issue" \
  --email user@example.com \
  --output json
```

### Update Ticket

```bash
# Change status
zohodesk-cli tickets update 123456789 --status Closed
zohodesk-cli tickets update 123456789 --status "On Hold"

# Change priority
zohodesk-cli tickets update 123456789 --priority High

# Multiple fields
zohodesk-cli tickets update 123456789 \
  --status "On Hold" \
  --priority High
```

### Close Ticket

```bash
# Close without resolution
zohodesk-cli tickets close 123456789

# Close with resolution
zohodesk-cli tickets close 123456789 \
  --resolution "Issue resolved by clearing cache"
```

### Search Tickets

```bash
# Search by email
zohodesk-cli tickets search --email user@example.com

# Search by status
zohodesk-cli tickets search --status Open

# Combine filters
zohodesk-cli tickets search \
  --email user@example.com \
  --status Open

# JSON output
zohodesk-cli tickets search --email user@example.com --output json
```

## Batch Operations

### Create Multiple Tickets

```bash
# From stdin
cat tickets.json | zohodesk-cli tickets create --batch

# From file
zohodesk-cli tickets create --batch --file tickets.json
```

**JSON format:**
```json
[
  {"subject": "Ticket 1", "email": "user1@example.com", "priority": "High"},
  {"subject": "Ticket 2", "email": "user2@example.com", "priority": "Medium"}
]
```

### Update Multiple Tickets

```bash
# From stdin
cat updates.json | zohodesk-cli tickets update --batch

# From file
zohodesk-cli tickets update --batch --file updates.json
```

**JSON format:**
```json
[
  {"id": "123", "status": "Closed"},
  {"id": "456", "status": "On Hold", "priority": "High"}
]
```

### Close Multiple Tickets

```bash
# From stdin
echo '["123", "456", "789"]' | zohodesk-cli tickets close --batch

# From file
zohodesk-cli tickets close --batch --file ticket_ids.json
```

**JSON format:**
```json
["123456789", "987654321", "111222333"]
```

### Batch Output

```json
{
  "total": 3,
  "succeeded": 2,
  "failed": 1,
  "results": [
    {"index": 0, "status": "success", "data": {...}},
    {"index": 1, "status": "success", "data": {...}},
    {"index": 2, "status": "failed", "error": "API error: ..."}
  ]
}
```

## Comments

### List Comments

```bash
# List all comments on a ticket
zohodesk-cli comments list 123456789

# JSON output
zohodesk-cli comments list 123456789 --output json
```

### Add Public Reply

```bash
# Add a public reply (customer can see)
zohodesk-cli comments reply 123456789 \
  --message "Thank you for your report. We are investigating."

# JSON output
zohodesk-cli comments reply 123456789 \
  --message "Your ticket is being processed." \
  --output json
```

### Add Private Note

```bash
# Add a private note (only agents can see)
zohodesk-cli comments note 123456789 \
  --message "Customer mentioned they tried clearing cache"

# JSON output
zohodesk-cli comments note 123456789 \
  --message "Internal note here" \
  --output json
```

## Contacts

### List Contacts

```bash
# List all contacts (default: 50)
zohodesk-cli contacts list

# List more contacts
zohodesk-cli contacts list --limit 100

# JSON output
zohodesk-cli contacts list --output json
```

### Get Contact

```bash
# Contact details
zohodesk-cli contacts get 123456789

# JSON output
zohodesk-cli contacts get 123456789 --output json
```

### Search Contacts

```bash
# By email
zohodesk-cli contacts search --email user@example.com

# By name
zohodesk-cli contacts search --name "John"

# JSON output
zohodesk-cli contacts search --email user@example.com --output json
```

## Configuration

### Initialize

```bash
# Interactive setup
zohodesk-cli config init
```

### List Profiles

```bash
zohodesk-cli config list
```

### Set Profile

```bash
# Create or update profile
zohodesk-cli config set \
  --profile work \
  --client-id "1000.xxx" \
  --client-secret "xxx" \
  --org-id "12345678" \
  --region "eu"
```

### Switch Default Profile

```bash
zohodesk-cli config use --profile work
```

### Delete Profile

```bash
zohodesk-cli config delete --profile work
```

## Working with Profiles

### Use Specific Profile

```bash
# For single command
zohodesk-cli tickets list --profile work

# Set as default
zohodesk-cli config use --profile work
```

## For AI Agents

### JSON Output

All commands support JSON output for easy parsing:

```bash
# Tickets
zohodesk-cli tickets list --output json
zohodesk-cli tickets get 123456789 --output json

# Comments
zohodesk-cli comments list 123456789 --output json

# Contacts
zohodesk-cli contacts list --output json
zohodesk-cli contacts search --email user@example.com --output json
```

### Example JSON Responses

#### Tickets List

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

#### Single Ticket

```json
{
  "id": "123456789",
  "ticketNumber": "TKT-001",
  "subject": "Login issue",
  "description": "User cannot login to portal",
  "status": "Open",
  "priority": "High",
  "createdTime": "2024-01-15T10:30:00Z",
  "modifiedTime": "2024-01-15T11:00:00Z",
  "contactId": "987654321",
  "contactEmail": "user@example.com",
  "contactName": "John Doe"
}
```

#### Comments List

```json
[
  {
    "id": "987654321",
    "ticketId": "123456789",
    "content": "Thank you for your report.",
    "authorName": "Agent Smith",
    "authorEmail": "agent@example.com",
    "isPublic": true,
    "createdTime": "2024-01-15T14:30:00Z"
  }
]
```

### Scripting Examples

```bash
# Count open tickets
zohodesk-cli tickets list --status Open --output json | jq 'length'

# Extract ticket IDs
zohodesk-cli tickets list --output json | jq '.[].id'

# Get all high priority tickets
zohodesk-cli tickets list --priority High --output json | \
  jq '.[] | select(.status == "Open")'

# Create ticket from script
SUBJECT="Bug report"
EMAIL="user@example.com"
zohodesk-cli tickets create \
  --subject "$SUBJECT" \
  --email "$EMAIL" \
  --output json
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Configuration error |
| 3 | API error |