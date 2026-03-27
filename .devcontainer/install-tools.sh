#!/usr/bin/env bash
set -euo pipefail

# Install buf
echo "Installing buf..."
BUF_VERSION=$(curl -fsSL https://api.github.com/repos/bufbuild/buf/releases/latest | grep '"tag_name"' | sed 's/.*"v\([^"]*\)".*/\1/')
curl -fsSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-Linux-x86_64" -o /usr/local/bin/buf
chmod +x /usr/local/bin/buf
echo "buf ${BUF_VERSION} installed"

# Install just
echo "Installing just..."
curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin
echo "just $(just --version) installed"

# Install markdownlint-cli and prettier
echo "Installing markdownlint-cli and prettier..."
npm install -g markdownlint-cli prettier
echo "markdownlint $(markdownlint --version) installed"
echo "prettier $(prettier --version) installed"
