.PHONY: test coverage lint fmt clean day

test:
	go test -v -race ./...

coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

coverage-cli:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	golangci-lint run --timeout 5m

fmt:
	go fmt ./...
	goimports -w .

fmt-check:
	@if [ -n "$$(go fmt ./...)" ]; then \
		echo "Code is not formatted"; \
		exit 1; \
	fi

vet:
	go vet ./...

clean:
	rm -f coverage.out coverage.html
	go clean -testcache

tidy:
	go mod tidy

# Create a new day directory with template files
day:
	@if [ -z "$(DAY)" ]; then \
		echo "Usage: make day DAY=01"; \
		exit 1; \
	fi
	@mkdir -p day$(DAY)
	@echo "package day$(DAY)\n\nfunc Solve(input string) (int, int) {\n\tpart1 := 0\n\tpart2 := 0\n\treturn part1, part2\n}" > day$(DAY)/solve.go
	@echo "package day$(DAY)\n\nimport (\n\t\"testing\"\n\t\"github.com/stretchr/testify/assert\"\n)\n\nfunc TestSolve(t *testing.T) {\n\tinput := \"\"\n\tpart1, part2 := Solve(input)\n\tassert.Equal(t, 0, part1)\n\tassert.Equal(t, 0, part2)\n}" > day$(DAY)/solve_test.go
	@echo "Created day$(DAY)"

all-checks: fmt-check vet lint test coverage-cli

help:
	@echo "Available targets:"
	@echo "  test           - Run all tests"
	@echo "  coverage       - Generate HTML coverage report"
	@echo "  coverage-cli   - Show coverage in CLI"
	@echo "  lint           - Run golangci-lint"
	@echo "  fmt            - Format and organize imports"
	@echo "  fmt-check      - Check if code is formatted"
	@echo "  vet            - Run go vet"
	@echo "  tidy           - Tidy go.mod"
	@echo "  day DAY=XX     - Create new day XX directory with templates"
	@echo "  all-checks     - Run all checks (fmt, vet, lint, test, coverage)"
	@echo "  clean          - Clean test cache and coverage files"
