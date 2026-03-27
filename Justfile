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

# Format and lint docs markdown files
docs:
    prettier --write "docs/**/*.md"
    markdownlint "docs/**/*.md"

# Generate templ files
templ:
    go tool templ generate

# Build Tailwind CSS (pass minify=true to minify)
css minify="false":
    npx @tailwindcss/cli \
        -i ./internal/ctl/web/static/input.css \
        -o ./internal/ctl/web/static/output.css \
        {{ if minify == "true" { "--minify" } else { "" } }}

# Build CLI binary
build: templ css
    go build -o bin/def ./cmd/def

# Run  (pass args with: just run -- --flag value)
run *ARGS: templ css
    go run ./cmd/def {{ ARGS }}
