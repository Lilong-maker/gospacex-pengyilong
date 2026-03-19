# AGENTS.md

Guidelines for AI coding agents working in this Go repository.

## Build / Test / Lint Commands

```bash
# Build
go build ./...                    # All packages
go build ./path/to/package        # Specific package

# Test
go test ./...                     # All tests
go test -v ./...                  # Verbose
go test -v ./pkg -run TestName    # Single test
go test -cover ./...              # Coverage
go test -race ./...               # Race detection

# Format & Lint
go fmt ./...                      # Format code
go vet ./...                      # Static analysis
go mod tidy                       # Tidy dependencies
staticcheck ./...                 # Advanced linting (install: go install honnef.co/go/tools/cmd/staticcheck@latest)
```

## Project Structure

```
gospacex-pengyilong/
├── cmd/appname/         # Main applications
├── internal/            # Private code
├── pkg/                 # Public libraries
├── api/                 # API definitions
├── configs/             # Configuration files
├── go.mod               # Module definition
└── AGENTS.md            # This file
```

## Code Style Guidelines

### Imports

Group in three sections (stdlib, third-party, local):

```go
import (
    "context"
    "fmt"
    
    "github.com/some/pkg"
    
    "gospacex-pengyilong/internal/x"
)
```

### Naming Conventions

| Type | Convention | Example |
|------|------------|---------|
| Package | lowercase, no underscores | `package mypkg` |
| Exported | PascalCase | `MyStruct`, `MyFunc()` |
| Unexported | camelCase | `myStruct`, `myFunc()` |
| Interface | ends with -er | `Reader`, `Writer` |
| Exported const | PascalCase | `MaxRetries` |
| Unexported const | camelCase | `defaultTimeout` |
| Acronyms | ALL UPPERCASE | `HTTPServer`, `URL`, `ID` |

### Error Handling

```go
// Always check errors
result, err := fn()
if err != nil {
    return nil, fmt.Errorf("context: %w", err)
}

// Sentinel errors
var ErrNotFound = errors.New("not found")
if errors.Is(err, ErrNotFound) { }

// Custom errors
var customErr *MyError
if errors.As(err, &customErr) { }
```

### Context

```go
// First param, named ctx
func DoWork(ctx context.Context, arg string) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
}
// Never store context in structs
```

### Types

```go
type Config struct {
    Host    *string // optional (pointer)
    Port    int     // required
    Enabled bool    // required (zero value ok)
}

// Small interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Testing

```go
package mypkg_test  // Black-box testing

func TestFunction(t *testing.T) {
    t.Run("case name", func(t *testing.T) {
        got, err := Function()
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}

// Table-driven tests
func TestProcess(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid", "in", "out", false},
        {"invalid", "bad", "", true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Process(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Documentation

```go
// Package mypkg provides ...
package mypkg

// Function does X. Returns Y on success, error on failure.
//
// Example:
//   result, err := Function(ctx, input)
func Function(ctx context.Context, input string) (Result, error) { }

// Explain WHY, not WHAT. Use complete sentences.
```

### Common Patterns

```go
// Functional options
type Option func(*Config)
func WithTimeout(d time.Duration) Option {
    return func(c *Config) { c.Timeout = d }
}

// Defer cleanup
f, err := os.Open("file")
if err != nil { return err }
defer f.Close()
```

## Key Principles

1. **Simplicity** > cleverness
2. **Explicit** > implicit
3. **Check errors** immediately
4. **Test** critical paths
5. **Document** exports and complex logic
6. **Run `go fmt`** before committing
7. **Atomic commits** with clear messages