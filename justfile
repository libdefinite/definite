default:
    @just --list

# Generate Go code from proto files
proto:
    buf generate proto

# Lint proto files
proto-lint:
    buf lint proto

# Tidy dependencies
tidy:
    go mod tidy

# Format code (requires: go install golang.org/x/tools/cmd/goimports@latest)
format:
    goimports -w -local github.com/libdefinite/definite $(find . -name "*.go" -not -path "./gen/*")

# Lint code (requires: golangci-lint)
lint:
    golangci-lint run

# Build CLI binary
build-cli:
    go build -o bin/definite ./cmd/cli

# Build server binary
build-server:
    go build -o bin/server ./cmd/server

# Run server
serve:
    go run ./cmd/server

# Run CLI (pass args with: just run -- --flag value)
run *ARGS:
    go run ./cmd/cli {{ ARGS }}
