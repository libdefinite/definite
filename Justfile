default:
    @just --list

# Generate Go code from proto files
proto:
    buf generate proto

# Lint proto files
proto-lint:
    buf lint proto

# Format code
format:
    go tool goimports \
        -w \
        -local github.com/libdefinite/definite \
        $(find . -name "*.go" -not -path "./gen/*")
    go tool templ fmt .

# Lint code
lint:
    go tool golangci-lint run
    go tool templ fmt -fail .

# Format and lint markdown files
md-lint:
    prettier --write "**/*.md" --ignore-path .gitignore
    markdownlint "**/*.md" --ignore-path .gitignore

# Generate templ files
templ:
    go tool templ generate

# Build Tailwind CSS (pass minify=true to minify)
css minify="false":
    npx @tailwindcss/cli \
        -i ./internal/ctl/console/static/input.css \
        -o ./internal/ctl/console/static/output.css \
        {{ if minify == "true" { "--minify" } else { "" } }}

# Run tests
test *ARGS:
    go test {{ ARGS }} ./...

# Generate code coverage report
coverage:
    go test -coverprofile=coverage.out ./...
    grep -v -e "_templ.go" -e "^github.com/libdefinite/definite/gen/" coverage.out > coverage.filtered.out
    go tool cover -func=coverage.filtered.out
    go tool cover -html=coverage.filtered.out -o coverage.html

# Build CLI binary
build: templ css
    go build -o bin/def ./cmd/def

# Run  (pass args with: just run -- --flag value)
run *ARGS: templ css
    go run ./cmd/def {{ ARGS }}
