# API Reference

## Global Commands

### `zohodesk-cli version`

Show version information.

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

### `zohodesk-cli help`

Show help for any command.

```bash
zohodesk-cli help
zohodesk-cli tickets --help
zohodesk-cli tickets list --help
```

## Configuration Commands

### `zohodesk-cli config init`

Initialize configuration with interactive setup wizard.

```bash
zohodesk-cli config init
```

**Flags:**

| Flag | Description |
|------|-------------|
| None | Interactive TUI setup |

### `zohodesk-cli config set`

Set configuration values for a profile.

```bash
zohodesk-cli config set [flags]
```

**Flags:**

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--profile` | `-p` | Profile name | No (default: default) |
| `--client-id` | `-c` | Zoho Client ID | No |
| `--client-secret` | `-s` | Zoho Client Secret | No |
| `--org-id` | `-o` | Zoho Organization ID | No |
| `--region` | `-r` | Zoho region | No (default: com) |

### `zohodesk-cli config list`

List all configured profiles.

```bash
zohodesk-cli config list
```

### `zohodesk-cli config delete`

Delete a configuration profile.

```bash
zohodesk-cli config delete --profile <name>
```

**Flags:**

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--profile` | `-p` | Profile name to delete | Yes |

### `zohodesk-cli config use`

Set the default profile.

```bash
zohodesk-cli config use --profile <name>
```

**Flags:**

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--profile` | `-p` | Profile name | Yes |

### `zohodesk-cli config clear-tokens`

Clear cached authentication tokens.

```bash
# Clear tokens for current profile
zohodesk-cli config clear-tokens

# Clear tokens for specific profile
zohodesk-cli config clear-tokens --profile work

# Clear all tokens
zohodesk-cli config clear-tokens --profile all
```

## Ticket Commands

### `zohodesk-cli tickets list`

List tickets with optional filters.

```bash
zohodesk-cli tickets list [flags]
```

**Flags:**

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--status` | `-s` | Filter by status | all |
| `--priority` | `-P` | Filter by priority | all |
| `--limit` | `-l` | Maximum results | 50 |
| `--offset` | | Pagination offset |0 |
| `--output` | `-o` | Output format | table |

**Status Values:** `Open`, `Closed`, `On Hold`, `In Progress`

**Priority Values:** `Low`, `Medium`, `High`

### `zohodesk-cli tickets get`

Get ticket details.

```bash
zohodesk-cli tickets get <ticket-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `ticket-id` | Ticket ID | Yes |

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--full` | `-f` | Include threads and contact |
| `--output` | `-o` | Output format |

### `zohodesk-cli tickets create`

Create a new ticket.

```bash
zohodesk-cli tickets create [flags]
```

**Flags:**

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--subject` | `-s` | Ticket subject | Yes |
| `--description` | `-d` | Ticket description | No |
| `--email` | `-e` | Contact email | Yes |
| `--priority` | `-P` | Priority | No (default: Medium) |
| `--department` | `-D` | Department ID | No |
| `--output` | `-o` | Output format | No |

### `zohodesk-cli tickets update`

Update a ticket.

```bash
zohodesk-cli tickets update <ticket-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `ticket-id` | Ticket ID | Yes |

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--status` | `-s` | New status |
| `--priority` | `-P` | New priority |
| `--resolution` | `-r` | Resolution message |
| `--output` | `-o` | Output format |

### `zohodesk-cli tickets close`

Close a ticket.

```bash
zohodesk-cli tickets close <ticket-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `ticket-id` | Ticket ID | Yes |

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--resolution` | `-r` | Resolution message |

### `zohodesk-cli tickets search`

Search tickets.

```bash
zohodesk-cli tickets search [flags]
```

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--email` | `-e` | Search by contact email |
| `--query` | `-q` | Search by keyword |
| `--status` | `-s` | Filter by status |
| `--from` | `-f` | From date (YYYY-MM-DD) |
| `--to` | `-t` | To date (YYYY-MM-DD) |
| `--output` | `-o` | Output format |

## Contact Commands

### `zohodesk-cli contacts list`

List contacts.

```bash
zohodesk-cli contacts list [flags]
```

**Flags:**

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--limit` | `-l` | Maximum results | 50 |
| `--output` | `-o` | Output format | table |

### `zohodesk-cli contacts get`

Get contact details.

```bash
zohodesk-cli contacts get <contact-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `contact-id` | Contact ID | Yes |

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output format |

### `zohodesk-cli contacts search`

Search contacts.

```bash
zohodesk-cli contacts search [flags]
```

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
--email` | `-e` | Search by email |
| `--name` | `-n` | Search by name |
| `--output` | `-o` | Output format |

## Time Tracking Commands

### `zohodesk-cli time list`

List time entries for a ticket.

```bash
zohodesk-cli time list <ticket-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `ticket-id` | Ticket ID | Yes |

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output format |

### `zohodesk-cli time add`

Add a time entry to a ticket.

```bash
zohodesk-cli time add <ticket-id> [flags]
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `ticket-id` | Ticket ID | Yes |

**Flags:**

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--duration` | `-d` | Time duration (e.g., 1h30m, 2h, 45m) | Yes |
| `--description` | `-D` | Description of work performed | No |
| `--agent` | `-a` | Agent ID (default: current user) | No |
| `--executed` | `-e` | Execution date (YYYY-MM-DD) | No |
| `--output` | `-o` | Output format | No |

### `zohodesk-cli time delete`

Delete a time entry.

```bash
zohodesk-cli time delete <time-entry-id>
```

**Arguments:**

| Argument | Description | Required |
|----------|-------------|----------|
| `time-entry-id` | Time Entry ID | Yes |

### `zohodesk-cli time report`

Generate a time report.

```bash
zohodesk-cli time report [flags]
```

**Flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--agent` | `-a` | Filter by agent ID |
| `--from` | `-f` | From date (YYYY-MM-DD) |
| `--to` | `-t` | To date (YYYY-MM-DD) |
| `--output` | `-o` | Output format |

## Output Formats

### JSON

```bash
zohodesk-cli tickets list --output json
```

### YAML

```bash
zohodesk-cli tickets list --output yaml
```

### Table (default)

```bash
zohodesk-cli tickets list --output table
zohodesk-cli tickets list
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ZOHO_CLIENT_ID` | Zoho Client ID |
| `ZOHO_CLIENT_SECRET` | Zoho Client Secret |
| `ZOHO_ORG_ID` | Zoho Organization ID |
| `ZOHO_REGION` | Zoho region (com, eu, in, cn, au) |
| `ZOHO_PROFILE` | Default profile name |

## Exit Codes

| Code | Description |
|------|-------------|
|0 | Success |
|1 | General error |
|2 | Configuration error |
|3 | API error |