# GitHub Copilot Instructions for Go Development

## Project Overview
This is a Go project that follows standard Go conventions and best practices.

## Code Style and Standards

### Go Formatting
- Use `gofmt` for formatting
- Follow effective Go naming conventions (camelCase for unexported, PascalCase for exported)
- Use meaningful variable and function names
- Keep functions small and focused

### Package Organization
- One package per directory
- Use lowercase package names without underscores

### Error Handling
- Always handle errors explicitly
- Use `fmt.Errorf` for error wrapping
- Return errors as the last return value

### Testing
- Write tests for all public functions
- Use table-driven tests when testing multiple scenarios
- Place tests in `*_test.go` files
- Include benchmarks for performance-critical code

### Documentation
- Write godoc comments for all exported functions, types, and constants
- Start comments with the name of the item being documented

### Dependencies
- Use Go modules (`go.mod`)
- Prefer standard library over external dependencies
- Pin dependency versions in production code
- Run `go mod tidy` regularly

### Security
- Validate all inputs
- Use context for timeouts and cancellation
- Avoid hardcoded secrets (use environment variables)
- Use secure random number generation

### Performance
- Prefer value types over pointers when possible
- Use `strings.Builder` for string concatenation
- Consider memory allocations in hot paths
- Use profiling tools (`go tool pprof`) for optimization

## Project Structure
```
src/
├── cli/                # Main applications
├── internal/           # Private application code
├── pkg/                # Public library code
├── api/                # API definitions
├── web/                # Web application assets
├── configs/            # Configuration files
├── deployments/        # Deployment configurations
├── test/               # Additional test files
├── docs/               # Documentation
├── scripts/            # Build and deployment scripts
└── vendor/             # Vendored dependencies (if used)
```

## Common Patterns

### CLI Applications
- Use `flag` package for command-line parsing
- Implement proper exit codes
- Provide help text and usage examples

### HTTP Services
- Use standard `net/http`
- Implement proper middleware for logging, auth, etc.
- Handle graceful shutdown
- Use context for request scoping

### Database Access
- Use `database/sql` with appropriate drivers
- Implement connection pooling
- Use prepared statements
- Handle transactions properly

### Configuration
- Use environment variables for runtime config
- Support configuration files (YAML, JSON, TOML)
- Validate configuration on startup

## Development Environment
- This project runs in a Debian 12 dev container
- Use `$BROWSER <url>` to open web pages in host browser
- Available tools: apt, git, curl, wget, ssh, gpg, etc.

## Build and Run
- Use `go build` for building
- Use `go run` for development
- Use `make` targets if Makefile exists
- Follow semantic versioning for releases

## Examples to Follow
When suggesting code, prefer patterns similar to:
- Standard library examples
- Go by Example patterns
- Effective Go guidelines
- Go Code Review Comments
