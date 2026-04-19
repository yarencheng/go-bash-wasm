[![Go Report Card](https://goreportcard.com/badge/github.com/yarencheng/go-bash-wasm)](https://goreportcard.com/report/github.com/yarencheng/go-bash-wasm)
[![Go Reference](https://pkg.go.dev/badge/github.com/yarencheng/go-bash-wasm.svg)](https://pkg.go.dev/github.com/yarencheng/go-bash-wasm)
![Go Version](https://img.shields.io/github/go-mod/go-version/yarencheng/go-bash-wasm)
[![Build Status](https://github.com/yarencheng/go-bash-wasm/actions/workflows/go.yml/badge.svg)](https://github.com/yarencheng/go-bash-wasm/actions)
[![codecov](https://codecov.io/gh/yarencheng/go-bash-wasm/branch/main/graph/badge.svg)](https://codecov.io/gh/yarencheng/go-bash-wasm)
![GitHub License](https://img.shields.io/github/license/yarencheng/go-bash-wasm)

# go-bash-wasm

**go-bash-wasm** is a Go implementation of GNU Bash and Coreutils for WebAssembly, featuring a fully isolated in-memory filesystem. It enables running a shell environment in browsers and other sandboxed environments.

**Demo**: [https://bash.devops-playground.dev/](https://bash.devops-playground.dev/)

## Key Features

- **GNU Parity**: Targets Bash 5.3 and Coreutils 9.10 (e.g., `ls` flag support).
- **WASM Support**: Compiles to `js/wasm` and `wasip1/wasm`. Integrated with **xterm.js** for browser terminals.
- **In-Memory VFS**: Uses `afero` for total host isolation. Zero disk I/O.
- **Observability**: Structured logging via `zerolog` with browser console support.

## Architecture

- `cmd/go-bash-wasm/`: Entry points for native (`main.go`) and JS/WASM (`main_js.go`).
- `internal/shell`: Execution engine using `mvdan.cc/sh/syntax`.
- `internal/commands`: Coreutils implementation.
- `ui`: Svelte 5 frontend with xterm.js components.
- `Dockerfile`: Multi-stage builds for browser (Nginx) and CLI (Wasmtime).

## Prerequisites

- **Go 1.26+**
- **Node.js 20+**
- **Docker**

## Local Development

### 1. Native CLI
```bash
go run ./cmd/go-bash-wasm/
```

### 2. Browser UI (Svelte + WASM)
1. **Compile WASM**:
   ```bash
   GOOS=js GOARCH=wasm go build -o ui/static/main.wasm ./cmd/go-bash-wasm/
   cp $(go env GOROOT)/lib/wasm/wasm_exec.js ui/static/
   ```
2. **Run Svelte App**:
   ```bash
   cd ui && npm install && npm run dev
   ```
Accessible at `http://localhost:5173`.

## Docker Deployment

### 1. Browser Terminal (Nginx)
```bash
docker build -t go-bash-ui -f ui.Dockerfile .
docker run -it --rm -p 8080:80 go-bash-ui
```

### 2. Native CLI (Wasmtime)
```bash
docker build -t go-bash-cli -f Dockerfile .
docker run -it --rm go-bash-cli
```

## Testing

### 1. Go Backend
```bash
go test -v ./...
```

### 2. UI Frontend
```bash
cd ui && npm run test
```

### 3. Full Validation
```bash
docker build -f ui.Dockerfile .
```

---
*Developed by the go-bash-wasm team. Aiming for 100% functional parity with GNU tools.*