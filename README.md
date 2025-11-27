# Advent of Code 2025 - Go Learning Project

Learning Go through Advent of Code 2025 with Test-Driven Development (TDD).

## Setup

### Prerequisites
- Docker (for devcontainer)
- VS Code with "Dev Containers" extension

### Getting Started

1. **Open in devcontainer:**
   - Open workspace in VS Code
   - Click "Reopen in Container" (bottom-left corner)
   - Wait for container to build and initialize

2. **Verify setup:**
   ```bash
   make test           # Run all tests
   make coverage       # Generate coverage report
   make lint           # Run static analysis
   ```

## Project Structure

```
.
├── day01/              # Day 1 solution (example)
│   ├── solve.go        # Main solution code
│   └── solve_test.go   # TDD tests
├── day02/              # Day 2, etc...
├── utils/              # Shared utility functions
│   ├── input.go        # Input parsing functions
│   └── input_test.go   # Tests for utilities
├── Makefile            # Development commands
├── go.mod              # Go module definition
├── .devcontainer/      # Dev container configuration
└── .github/workflows/  # CI/CD pipelines
```

## Common Utility Functions

The `utils/input.go` provides helper functions for common AoC input patterns:

### Reading Input
```go
// Read file as lines
lines, err := utils.ReadLines("input.txt")

// Read string as lines
lines := utils.ReadLinesString(input)

// Read entire file as string
content, err := utils.ReadRawFile("input.txt")
```

### Grid/Matrix Operations
```go
// Convert lines to 2D character grid
grid := utils.Read2DGrid(lines)        // [][]rune
grid := utils.Read2DGridFromString(input)

// Access grid elements
if grid[row][col] == '#' { ... }
```

### Processing Lines
```go
// Parse integer per line
nums := utils.ParseIntLines(lines)

// Remove empty lines
lines := utils.FilterEmptyLines(lines)

// Split lines by delimiter
groups := utils.SplitOn(lines, "")  // blank line separator
```

## Workflow

### Creating a New Day

```bash
# Create day02 directory with templates
make day DAY=02

# This creates:
# - day02/solve.go (main solution)
# - day02/solve_test.go (TDD test file)
```

### TDD Workflow

1. **Write a failing test** in `dayXX/solve_test.go`
   ```go
   func TestSolvePart1(t *testing.T) {
       input := `your example`
       result := solvePart1(parseInput(input))
       assert.Equal(t, expectedValue, result)
   }
   ```

2. **Run tests** to see failure
   ```bash
   make test
   # or
   go test ./dayXX -v
   ```

3. **Implement code** in `dayXX/solve.go` to pass tests

4. **Run tests again** to verify pass
   ```bash
   make test
   ```

5. **Refactor** and repeat

### Make Commands

```bash
# Testing
make test                 # Run all tests
make coverage             # Generate HTML coverage report
make coverage-cli         # Show coverage in terminal

# Code Quality
make lint                 # Run golangci-lint
make fmt                  # Format code
make fmt-check            # Check formatting (CI)
make vet                  # Run go vet

# Utilities
make day DAY=XX           # Create day directory
make tidy                 # Tidy go.mod
make all-checks           # Run all checks
make clean                # Clean test cache

make help                 # Show all commands
```

## CI/CD Pipeline

GitHub Actions automatically runs on push/PR:

✅ **Unit Tests** - All tests with race detector
✅ **Code Coverage** - Must maintain 80%+ coverage
✅ **Linting** - golangci-lint analysis
✅ **Format Check** - Ensures `go fmt` compliance
✅ **Vet Check** - Runs `go vet` for correctness

Coverage reports uploaded to Codecov.

## Dev Container Features

Includes:
- **Go 1.23** latest
- **VS Code Extensions:**
  - Go language support
  - golangci-lint integration
  - Makefile support
- **Post-create tools:**
  - golangci-lint pre-installed
- **SSH forwarding** for git operations

## Tips for AoC

1. **Use grid parsing for spatial puzzles:**
   ```go
   grid := utils.Read2DGridFromString(input)
   for row := 0; row < len(grid); row++ {
       for col := 0; col < len(grid[row]); col++ {
           cell := grid[row][col]
       }
   }
   ```

2. **Split on blank lines for grouping:**
   ```go
   groups := utils.SplitOn(lines, "")
   ```

3. **Write tests with examples first** - AoC usually provides examples!

4. **Focus on one puzzle at a time** - Use day-focused packages

## Resources

- [Advent of Code 2025](https://adventofcode.com/2025)
- [Go Documentation](https://golang.org/doc)
- [Testify Package](https://github.com/stretchr/testify)
