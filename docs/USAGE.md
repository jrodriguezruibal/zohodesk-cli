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

# Ticket with full context (department, assignee, SLA)
zohodesk-cli tickets get 123456789 --context

# JSON output
zohodesk-cli tickets get 123456789 --output json
```

### Assign Ticket

```bash
# Assign to an agent
zohodesk-cli tickets assign 123456789 --agent 987654321

# Assign to a department
zohodesk-cli tickets assign 123456789 --department 456789

# Assign to both agent and department
zohodesk-cli tickets assign 123456789 \
  --agent 987654321 \
  --department 456789

# JSON output
zohodesk-cli tickets assign 123456789 --agent 987654321 --output json
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

## Departments

### List Departments

```bash
# List all departments
zohodesk-cli departments list

# JSON output
zohodesk-cli departments list --output json
```

### Get Department

```bash
# Department details
zohodesk-cli departments get 123456789

# JSON output
zohodesk-cli departments get 123456789 --output json
```

## Agents

### List Agents

```bash
# List all agents
zohodesk-cli agents list

# JSON output
zohodesk-cli agents list --output json
```

### Get Agent

```bash
# Agent details
zohodesk-cli agents get 123456789

# JSON output
zohodesk-cli agents get 123456789 --output json
```

## Time Tracking

### List Time Entries

```bash
# List all time entries for a ticket
zohodesk-cli time list 123456789

# JSON output
zohodesk-cli time list 123456789 --output json
```

### Add Time Entry

```bash
# Add 1 hour and 30 minutes
zohodesk-cli time add 123456789 --duration 1h30m --description "Investigated the issue"

# Add 2 hours
zohodesk-cli time add 123456789 --duration 2h --description "Fixed the bug"

# Add 45 minutes
zohodesk-cli time add 123456789 --duration 45m --description "Testing"

# Add time for specific agent
zohodesk-cli time add 123456789 --duration 1h --agent 987654321 --description "Review"

# Add time with execution date
zohodesk-cli time add 123456789 --duration 2h --executed 2024-01-15 --description "Work done yesterday"

# JSON output
zohodesk-cli time add 123456789 --duration 1h --description "Investigating" --output json
```

### Delete Time Entry

```bash
# Delete a time entry
zohodesk-cli time delete 987654321
```

### Time Report

```bash
# Get time report
zohodesk-cli time report

# Filter by agent
zohodesk-cli time report --agent 123456789

# Filter by date range
zohodesk-cli time report --from 2024-01-01 --to 2024-01-31

# JSON output
zohodesk-cli time report --output json
```

## Attachments

### List Attachments

```bash
# List all attachments for a ticket
zohodesk-cli attachments list 123456789

# JSON output
zohodesk-cli attachments list 123456789 --output json
```

### Upload Attachment

```bash
# Upload a file to a ticket
zohodesk-cli attachments upload 123456789 /path/to/file.pdf

# JSON output
zohodesk-cli attachments upload 123456789 /path/to/file.pdf --output json
```

### Download Attachment

```bash
# Download attachment to current directory
zohodesk-cli attachments download 987654321

# Download to specific path
zohodesk-cli attachments download 987654321 --output /path/to/save/file.pdf
```

### Delete Attachment

```bash
# Delete an attachment
zohodesk-cli attachments delete 987654321
```

## Knowledge Base

### List Articles

```bash
# List all articles
zohodesk-cli articles list

# Limit results
zohodesk-cli articles list --limit 20

# JSON output
zohodesk-cli articles list --output json
```

### Get Article

```bash
# Article details
zohodesk-cli articles get 123456789

# JSON output
zohodesk-cli articles get 123456789 --output json
```

### Search Articles

```bash
# Search articles by keyword
zohodesk-cli articles search "password reset"

# JSON output
zohodesk-cli articles search "password reset" --output json
```

### Create Article

```bash
# Create a new article
zohodesk-cli articles create \
  --title "How to reset password" \
  --content "Step by step guide..." \
  --summary "Password reset instructions" \
  --category 123456789

# JSON output
zohodesk-cli articles create \
  --title "FAQ" \
  --content "Common questions..." \
  --output json
```

### Update Article

```bash
# Update article
zohodesk-cli articles update 123456789 \
  --title "Updated title" \
  --status "Published"

# JSON output
zohodesk-cli articles update 123456789 --status "Draft" --output json
```

### Delete Article

```bash
# Delete an article
zohodesk-cli articles delete 123456789
```

### Categories

```bash
# List all categories
zohodesk-cli categories list

# Get category details
zohodesk-cli categories get 123456789

# JSON output
zohodesk-cli categories list --output json
```

## Tasks

### List Tasks

```bash
# List all tasks
zohodesk-cli tasks list

# List tasks for a specific ticket
zohodesk-cli tasks list --ticket 123456789

# JSON output
zohodesk-cli tasks list --output json
```

### Get Task

```bash
# Task details
zohodesk-cli tasks get 123456789

# JSON output
zohodesk-cli tasks get 123456789 --output json
```

### Create Task

```bash
# Create a task
zohodesk-cli tasks create --title "Review code changes" --description "Review PR #42"

# Create task with due date and priority
zohodesk-cli tasks create \
  --title "Complete feature" \
  --priority "High" \
  --due 2024-01-15

# Create task for a ticket
zohodesk-cli tasks create \
  --title "Follow up with customer" \
  --ticket 123456789

# JSON output
zohodesk-cli tasks create --title "New task" --output json
```

### Update Task

```bash
# Update task
zohodesk-cli tasks update 123456789 --status "In Progress" --priority "High"

# JSON output
zohodesk-cli tasks update 123456789 --title "Updated title" --output json
```

### Complete Task

```bash
# Mark task as completed
zohodesk-cli tasks complete 123456789
```

### Delete Task

```bash
# Delete a task
zohodesk-cli tasks delete 123456789
```

## Reports

### Ticket Statistics

```bash
# Get ticket statistics
zohodesk-cli reports tickets

# Filter by date range
zohodesk-cli reports tickets --from 2024-01-01 --to 2024-01-31

# Filter by department
zohodesk-cli reports tickets --department 123456789

# JSON output
zohodesk-cli reports tickets --output json
```

### Agent Statistics

```bash
# Get agent performance statistics
zohodesk-cli reports agents

# Filter by date range
zohodesk-cli reports agents --from 2024-01-01 --to 2024-01-31

# Filter by specific agent
zohodesk-cli reports agents --agent 123456789

# JSON output
zohodesk-cli reports agents --output json
```

### SLA Statistics

```bash
# Get SLA compliance statistics
zohodesk-cli reports sla

# Filter by date range
zohodesk-cli reports sla --from 2024-01-01 --to 2024-01-31

# Filter by department
zohodesk-cli reports sla --department 123456789

# JSON output
zohodesk-cli reports sla --output json
```

## Products

### List Products

```bash
# List all products
zohodesk-cli products list

# JSON output
zohodesk-cli products list --output json
```

### Get Product

```bash
# Product details
zohodesk-cli products get 123456789

# JSON output
zohodesk-cli products get 123456789 --output json
```

## Accounts

### List Accounts

```bash
# List all accounts
zohodesk-cli accounts list

# JSON output
zohodesk-cli accounts list --output json
```

### Get Account

```bash
# Account details
zohodesk-cli accounts get 123456789

# JSON output
zohodesk-cli accounts get 123456789 --output json
```

### Create Account

```bash
# Create account
zohodesk-cli accounts create \
  --name "Acme Corp" \
  --email "contact@acme.com" \
  --phone "+1234567890"

# With additional fields
zohodesk-cli accounts create \
  --name "Acme Corp" \
  --type "Customer" \
  --industry "Technology"

# JSON output
zohodesk-cli accounts create --name "New Account" --output json
```

### Update Account

```bash
# Update account
zohodesk-cli accounts update 123456789 --name "Updated Name"

# JSON output
zohodesk-cli accounts update 123456789 --type "Partner" --output json
```

### Delete Account

```bash
# Delete account
zohodesk-cli accounts delete 123456789
```

## Tags

### List Tags

```bash
# List tags on a ticket
zohodesk-cli tickets tags list 123456789

# JSON output
zohodesk-cli tickets tags list 123456789 --output json
```

### Add Tags

```bash
# Add multiple tags
zohodesk-cli tickets tags add 123456789 --tag urgent --tag bug --tag customer

# JSON output
zohodesk-cli tickets tags add 123456789 --tag urgent --output json
```

### Remove Tag

```bash
# Remove a specific tag
zohodesk-cli tickets tags remove 123456789 urgent
```

## Ticket Operations

### Merge Tickets

```bash
# Merge source ticket into target ticket
zohodesk-cli tickets merge 123456789 987654321
```

### Follow/Unfollow Tickets

```bash
# Follow a ticket (receive notifications)
zohodesk-cli tickets follow 123456789

# Stop following a ticket
zohodesk-cli tickets unfollow 123456789
```

### Update with Custom Fields

```bash
# Update ticket with custom fields
zohodesk-cli tickets update 123456789 --custom-fields '{"priority_level": "high", "customer_type": "enterprise"}'
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