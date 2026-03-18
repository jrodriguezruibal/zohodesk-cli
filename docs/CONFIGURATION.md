# Configuration Guide

## Prerequisites

### Zoho API Credentials

1. Go to [Zoho API Console](https://api-console.zoho.com)
2. Sign in with your Zoho account
3. Click **Add Client** → select **Self-Client**
4. Fill in:
   - **Client Name**: `zohodesk-cli` (or any name you prefer)
   - **Homepage URL**: `http://localhost` (not used for Self-Client)
   - **Authorized Redirect URI**: `http://localhost` (not used for Self-Client)
5. Click **Create**
6. Note your **Client ID** and **Client Secret**

### Get Organization ID

1. Open [Zoho Desk](https://desk.zoho.com)
2. Go to **Settings** (gear icon) → **Company** → **Organization**
3. Copy your **Organization ID**

## Configuration Methods

### 1. Interactive TUI (Recommended)

```bash
zohodesk-cli config init
```

This launches an interactive wizard that:
- Prompts for Client ID, Client Secret, Organization ID, Region
- Validates credentials against Zoho API
- Creates configuration files with correct permissions

### 2. Environment Variables

For CI/CD pipelines or containers:

```bash
export ZOHO_CLIENT_ID="1000.xxxxxxxxxxxxx"
export ZOHO_CLIENT_SECRET="xxxxxxxxxxxxxxxxx"
export ZOHO_ORG_ID="12345678"
export ZOHO_REGION="com"          # Optional: com, eu, in, cn, au
```

Add to your shell profile (`~/.bashrc`, `~/.zshrc`):

```bash
# ~/.bashrc
export ZOHO_CLIENT_ID="1000.xxxxxxxxxxxxx"
export ZOHO_CLIENT_SECRET="xxxxxxxxxxxxxxxxx"
export ZOHO_ORG_ID="12345678"
export ZOHO_REGION="com"
```

### 3. Configuration File

Create `~/.config/zohodesk-cli/config.yaml`:

```yaml
profiles:
  default:
    client_id: "1000.xxxxxxxxxxxxx"
    client_secret: "xxxxxxxxxxxxxxxxx"
    org_id: "12345678"
    region: "com"

  work:
    client_id: "1000.yyyyyyyyyyyyy"
    client_secret: "yyyyyyyyyyyyyyyyy"
    org_id: "87654321"
    region: "eu"

current_profile: "default"
```

**Important**: Set correct permissions:

```bash
chmod 600 ~/.config/zohodesk-cli/config.yaml
```

## Multi-Profile Management

### List Profiles

```bash
zohodesk-cli config list
```

Output:

```
Current profile: default

Profiles:
  * default (region: com)
    work (region: eu)
    personal (region: com)
```

### Add Profile

```bash
zohodesk-cli config set --profile work \
  --client-id "1000.yyy" \
  --client-secret "yyy" \
  --org-id "87654321" \
  --region "eu"
```

### Switch Default Profile

```bash
zohodesk-cli config use --profile work
```

### Use Profile for Single Command

```bash
zohodesk-cli tickets list --profile work
```

### Delete Profile

```bash
zohodesk-cli config delete --profile work
```

## Regions

zohodesk-cli supports all Zoho Desk regions:

| Region | Flag | API Endpoint |
|--------|------|--------------|
| United States | `--region com` | https://desk.zoho.com |
| Europe | `--region eu` | https://desk.zoho.eu |
| India | `--region in` | https://desk.zoho.in |
| China | `--region cn` | https://desk.zoho.com.cn |
| Australia | `--region au` | https://desk.zoho.com.au |

## Token Cache

Access tokens are cached in:

```
~/.config/zohodesk-cli/tokens/<profile-hash>.json
```

### Clear Token Cache

```bash
# Clear tokens for current profile
zohodesk-cli config clear-tokens

# Clear tokens for specific profile
zohodesk-cli config clear-tokens --profile work

# Clear all tokens
zohodesk-cli config clear-tokens --profile all
```

## Security Best Practices

1. **File Permissions**: Config files have `600` permissions (owner read/write only)

2. **Environment Variables**: Never commit `.env` files to version control

3. **Token Expiry**: Tokens expire and are automatically refreshed

4. **Secrets Rotation**: Periodically rotate your Client Secret in Zoho API Console

5. **CI/CD**: Use environment variables or secret management systems

## Troubleshooting

### Invalid Credentials

```
Error: auth failed (400): invalid_client
```

**Solution**: Verify your Client ID and Client Secret in the Zoho API Console.

### Invalid Organization ID

```
Error: profile 'default' not found
```

**Solution**: Check your Organization ID in Zoho Desk Settings.

### Token Expired

```
Error: API error (401): Invalid token
```

**Solution**: Clear token cache to force re-authentication:

```bash
zohodesk-cli config clear-tokens
```

### Region Not Supported

```
Error: invalid region 'xyz'
```

**Solution**: Use valid regions: `com`, `eu`, `in`, `cn`, `au`