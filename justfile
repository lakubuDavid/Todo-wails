# List all available commands
default:
    just --list

# Run Wails in development mode with specific options
run:
    wails dev -skipbindings -noreload

# Run Wails in development mode
dev:
    wails dev

# Build the Wails project with optional arguments
build +args="":
    wails build {{args}}

# Generate code using sqlc and Wails
generate:
    sqlc generate
    wails generate module

# Alias for dbmate with optional arguments
db +args="":
    dbmate {{args}}

# Clean the build directory
clean:
    rm -rf build

# Lint the codebase
lint:
    golangci-lint run ./...

# Test the project
test:
    go test ./...

# Add a brief description for each command
@description:
    default: "List all available commands"
    run: "Run Wails in development mode with specific options"
    dev: "Run Wails in development mode"
    build: "Build the Wails project with optional arguments"
    generate: "Generate code using sqlc and Wails"
    db: "Alias for dbmate with optional arguments"
    clean: "Clean the build directory"
    lint: "Lint the codebase"
    test: "Test the project"

