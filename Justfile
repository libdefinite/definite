default:
    just --list

# Generate Go code from proto files
proto: proto-clean
    buf generate proto

# Clean generated proto files
proto-clean:
    rm -rf gen/

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
templ: templ-clean
    go tool templ generate

# Clean generated templ files
templ-clean:
    find . -name "*_templ.go" | xargs rm -f

# Build JS bundle (pass minify=false to skip minification)
js minify="false":
    npx vite build {{ if minify == "false" { "--no-minify" } else { "" } }}

# Build Tailwind CSS (pass minify=true to minify)
css minify="false":
    npx @tailwindcss/cli \
        -i ./internal/ctl/console/css/global.css \
        -o ./internal/ctl/console/static/output.css \
        {{ if minify == "true" { "--minify" } else { "" } }}

# Run go tests
test *ARGS:
    go test {{ ARGS }} ./...

# Generate code coverage report
coverage:
    go test -coverprofile=coverage.out ./...
    grep -v -e "_templ.go" -e "^github.com/libdefinite/definite/gen/" coverage.out > coverage.filtered.out
    go tool cover -func=coverage.filtered.out
    go tool cover -html=coverage.filtered.out -o coverage.html

# Run console with hot reload (proto, js, css changes do not trigger restart — run those tasks manually in a separate terminal)
dev *ARGS: proto js css
    air -- ctl console {{ ARGS }}

# Build CLI binary
build: proto templ (js "true") (css "true")
    go build -ldflags="-s -w" -o bin/def ./cmd/def

# Run (pass args with: just run -- --flag value)
run *ARGS: proto templ js css
    go run ./cmd/def {{ ARGS }}
