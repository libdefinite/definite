# AGENT.md

## Devcontainer

Devcontainer setup in `.devcontainer/devcontainer.json`

1. Go baseimage
2. Features to enable nodejs, python, `just`, buf, tmux
3. Markdown linting using `markdown-cli`. Markdown formating with `prettier`
4. `lefthook install` to setup git hooks
5. docker in docker for running test-container for testing

## Justfile

This project uses `just` as the task runner (see `Justfile`). Run `just` to see all available options. Tasks cover testing, linting, formatting, proto generation, and building.

## Architecture

**Definite** is an infrastructure orchestration system with a dual-plane, multi-server architecture.

### Entry point

`def` is the name of executable. Entry point for executable is `cmd/def/main.go`. Cobra used for cli. Two top-level subcommands:

- `node` — starts the backend servers (data plane + control plane)
- `ctl` — client-side tools (web console and definite management from commandline). CTL also has sub commands implemented in `internal/ctl/cmd.go`. Subcommands are as follows
  - `console`- to start web console

When definite runs in either

1. _ctl mode_ like kubectl
2. _node mode_ like like kube apiserver (node mode is a unified process for apiserver, kubecontroller, etcd and kubenode etc.)

---

### Node mode

`cmd/def/node.go` starts two gRPC servers concurrently:

Implementation in `internal/node`.

#### GRPC

GRPC handlers implemented in `internal/node/handler`

- Grpc used throughout for node to node communication and ctl to node communication
- _Data bus_ is grpc server which takes care of ctl to node communication. Proto file in `proto/data`
- _Control bus_ is grpc server which takes care of node to node communication. Proto files in `proto/control`
- ConnectRPC used for implementing GRPC servers

#### Proto and generated code

Protos live in `proto/`, organized by plane and version (`control/v1/`, `data/v1/`). Generated code goes to `gen/` (excluded from linting). Never edit files in `gen/` directly — run `just proto` instead.

Code generation config: `buf.gen.yaml` (uses `buf.build/protocolbuffers/go` + `buf.build/connectrpc/go`).

---

### CTL mode

implemented in `internal/ctl`

#### CTL Console

`internal/ctl/console` — implements web console.

- **web console** run using `def ctl console`
- **Templ** for type-safe, compile-time-checked HTML templates. All templates to be placed in `internal/ctl/console/templates` folder
- **Tailwind CSS v4** (generated from `internal/ctl/console/css/global.css`)
- **HTMX** + **Alpine.js** bundled locally via Vite (not CDN)

### Frontend build

Templ templates (`*.templ`) must be compiled to `*_templ.go` before building. `just build` and `just run` handle this automatically. If you edit `.templ` files, run `just templ` or `just format` to regenerate.

#### JS bundle

Entry point: `internal/ctl/console/js/main.js`. Vite config in `vite.config.js` bundles htmx and Alpine.js as IIFE to `internal/ctl/console/static/output.js`. Run `just js` to rebuild. Use JSDoc for type annotations — no TypeScript.

#### Tailwind CSS output

Tailwind CSS source in `internal/ctl/console/css/global.css`, scans `*.templ` files and outputs to `internal/ctl/console/static/output.css`.

### Linting

- Go linting configuration available in .golangci.yaml

---

## GitHub Actions

- `.github/workflows/VERIFY.yaml` — runs on PRs and pushes to main: lint job + coverage job (uploaded to codecov.io)

## Testing

- test-container library to be used for testing
