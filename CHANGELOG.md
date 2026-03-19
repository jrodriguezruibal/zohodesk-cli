# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.9.1] - 2026-03-19

### Fixed
- **Authentication**: Fixed Self-Client OAuth flow for Zoho Desk API
  - Use `redirect_uri=self` for Self-Client authentication
  - Add `access_type=offline` to obtain refresh tokens
  - Remove client_credentials fallback (not supported by Zoho Desk)
  - Automatic token refresh using refresh tokens
  - Clear error message when authentication is required
- **API Response Parsing**: Handle responses with or without `{data: ...}` wrapper
  - Zoho Desk API returns data directly in some endpoints
  - Fixed Tickets Get/Create, Comments Add responses
- **Flag Conflicts**: Changed `--context` shorthand from `-c` to `-C` to avoid conflict with `--config`
- **Scopes**: Updated to valid Zoho Desk API scopes only
- **JSON Field Types**: Changed viewCount and likeCount in Article model from int to string

### Added
- `--contact-id` flag for ticket creation (contacts are required in Zoho Desk API)
- `unmarshalData` helper function to handle both API response formats
- `SearchByEmail` method in ContactsService

### Changed
- Ticket creation now requires `--department` and `--contact-id` (or `--email`)
- Updated documentation with correct authentication instructions
- Authentication is now one-time setup with automatic token refresh

### Removed
- **Time Tracking commands** (`time list/add/delete/report`): API endpoint not available in Zoho Desk
- **Categories commands** (`categories list/get`): API returns null data

## [0.9.0] - 2026-03-18

### Added
- **Tags Commands**: Manage tags on tickets
  - `tickets tags list <ticket-id>` - List tags on a ticket
  - `tickets tags add <ticket-id> --tag <tag>` - Add tags to a ticket
  - `tickets tags remove <ticket-id> <tag>` - Remove a tag from a ticket
- **Ticket Merge**: Merge duplicate tickets
  - `tickets merge <source-id> <target-id>` - Merge source into target
- **Ticket Follow/Unfollow**: Follow ticket updates
  - `tickets follow <ticket-id>` - Start following a ticket
  - `tickets unfollow <ticket-id>` - Stop following a ticket
- **Custom Fields Support**: Update ticket custom fields
  - `tickets update <id> --custom-fields '{"field": "value'}'`
- Tags field in Ticket model
- CustomFields field in Ticket and TicketUpdateRequest models
- New models: Tag, TagListResponse, TicketTagsRequest, TicketMergeRequest, FollowRequest
- New service: TagsService

### Changed
- Updated documentation (USAGE.md, API.md, README.md)

## [0.8.0] - 2026-03-18

### Added
- **Products Commands**: View products in your helpdesk
  - `products list` - List all products
  - `products get <id>` - Get product details
- **Accounts Commands**: Manage customer accounts
  - `accounts list` - List all accounts
  - `accounts get <id>` - Get account details
  - `accounts create` - Create a new account
  - `accounts update <id>` - Update an account
  - `accounts delete <id>` - Delete an account
- New models: Product, Account, AccountCreateRequest, AccountUpdateRequest
- New services: ProductsService, AccountsService
- Output formatters for products and accounts

### Changed
- Updated documentation (USAGE.md, API.md, README.md)

## [0.7.0] - 2026-03-18

### Added
- **Reports Commands**: Generate statistics and analytics
  - `reports tickets` - Ticket statistics (by status, priority, department)
  - `reports agents` - Agent performance statistics
  - `reports sla` - SLA compliance statistics
- New models: TicketStats, AgentStats, SLAStats
- New service: ReportsService
- Output formatters for reports (tables with statistics)
- Date range filtering for all reports

### Changed
- Updated documentation (USAGE.md, API.md, README.md)

## [0.6.0] - 2026-03-18

### Added
- **Tasks Commands**: Manage tasks associated with tickets
  - `tasks list` - List all tasks or filter by ticket
  - `tasks get <id>` - Get task details
  - `tasks create` - Create a new task
  - `tasks update <id>` - Update task details
  - `tasks complete <id>` - Mark task as completed
  - `tasks delete <id>` - Delete a task
- New models: Task, TaskCreateRequest, TaskUpdateRequest
- New service: TasksService
- Output formatters for tasks

### Changed
- Updated documentation (USAGE.md, API.md, README.md)

## [0.5.0] - 2026-03-18

### Added
- **Knowledge Base Commands**: Manage knowledge base articles and categories
  - `articles list` - List all articles
  - `articles get <id>` - Get article details
  - `articles search <query>` - Search articles by keyword
  - `articles create` - Create a new article
  - `articles update <id>` - Update an article
  - `articles delete <id>` - Delete an article
  - `categories list` - List all categories
  - `categories get <id>` - Get category details
- New models: Article, ArticleCreateRequest, ArticleUpdateRequest, Category
- New services: ArticlesService, CategoriesService
- Output formatters for articles and categories

### Changed
- Updated documentation (USAGE.md, API.md, README.md)

## [0.4.0] - 2026-03-18

### Added
- **Attachments Commands**: Manage file attachments on tickets
  - `attachments list <ticket-id>` - List attachments for a ticket
  - `attachments upload <ticket-id> <file>` - Upload a file attachment
  - `attachments download <attachment-id>` - Download an attachment
  - `attachments delete <attachment-id>` - Delete an attachment
- Multipart file upload support in API client
- File size formatting in output (KB, MB, GB)
- Update documentation (USAGE.md, API.md, README.md)

### Changed
- Extended Attachment model with DownloadURL field

## [0.3.0] - 2026-03-18

### Added
- **Time Tracking Commands**: Manage time entries on tickets
  - `time list <ticket-id>` - List time entries for a ticket
  - `time add <ticket-id>` - Add time entry with duration (e.g., 1h30m, 2h, 45m)
  - `time delete <time-id>` - Delete a time entry
  - `time report` - Generate time report with filters
- Support for duration parsing (1h30m, 2h, 45m formats)
- New models: TimeEntry, TimeEntryCreateRequest, TimeReport
- New service: TimeService for time tracking API
- Output formatters for time entries and reports

### Changed
- Updated documentation with time tracking examples
- Added time tracking to feature list in README

## [0.2.2] - 2026-03-18

### Added
- **Departments Commands**: Manage Zoho Desk departments
  - `departments list` - List all departments
  - `departments get <id>` - Get department details
- **Agents Commands**: Manage Zoho Desk agents
  - `agents list` - List all agents
  - `agents get <id>` - Get agent details
- **Ticket Context Enrichment**: Get tickets with full context
  - `tickets get <id> --context` - Show ticket with department, assignee, and SLA info
- **Ticket Assignment**: Assign tickets to agents and departments
  - `tickets assign <id> --agent <agent-id>` - Assign ticket to agent
  - `tickets assign <id> --department <dept-id>` - Assign ticket to department
  - `tickets assign <id> --agent <agent-id> --department <dept-id>` - Assign to both
- New models: `SLA`, `Activity`, `TicketContext`
- New services: `DepartmentsService`, `AgentsService`
- Output formatters for departments and agents (JSON, YAML, Table)

### Changed
- Moved `Agent`, `Department` models to `pkg/models/resources.go`
- Enhanced models with additional fields (createdTime, modifiedTime, phone, mobile, photoURL)

## [0.2.1] - 2026-03-18

### Added
- Unit tests for batch operations (`internal/batch/processor_test.go`)
- Unit tests for data models (`pkg/models/comments_test.go`)
- Test coverage for Comments, Attachments, Agents, Departments
- Test coverage for BatchResult and BatchResponse JSON marshaling

### Changed
- Improved test coverage for batch operations error handling
- Added tests for edge cases in batch processing

## [0.2.0] - 2026-03-18

### Added
- **Batch Operations**: Create, update, and close multiple tickets from JSON input
  - `tickets create --batch` - Create multiple tickets from stdin or file
  - `tickets update --batch` - Update multiple tickets from stdin or file
  - `tickets close --batch` - Close multiple tickets from stdin or file
- **Comments/Replies Support**: Manage ticket comments
  - `comments list <ticket-id>` - List all comments on a ticket
  - `comments reply <ticket-id> --message` - Add a public reply
  - `comments note <ticket-id> --message` - Add a private note
- Rate limiting for batch operations (100ms delay between requests)
- New models: `Comment`, `Attachment`, `BatchRequest`, `BatchResponse`, `Agent`, `Department`
- New services: `CommentsService`, `BatchProcessor`
- `GetRaw` method in API client for streaming responses

### Documentation
- Updated README with batch operations and comments usage
- Updated docs/USAGE.md with examples for batch operations and comments
- Added JSON format examples for batch operations

## [0.1.2] - 2026-03-18

### Fixed
- Removed Homebrew brews from goreleaser (manual formula publication)
- Fixed changelog filter in goreleaser

## [0.1.1] - 2026-03-18

### Fixed
- Initial release with Docker builds disabled
- Fixed goreleaser configuration

## [0.1.0] - 2026-03-18

### Added
- Initial release
- Tickets commands: list, get, create, update, close, search
- Contacts commands: list, get, search
- Configuration management with multi-profile support
- TUI setup wizard for interactive configuration
- Multiple output formats: JSON, YAML, Table
- Self-Client OAuth authentication (no browser interaction)
- Token caching with automatic refresh
- Support for all Zoho regions (com, eu, in, cn, au)
- Cross-platform binaries (Linux, macOS, Windows) for amd64 and arm64
- Debian/RPM packages
- GitHub Actions CI/CD with goreleaser