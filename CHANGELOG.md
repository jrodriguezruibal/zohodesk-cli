# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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