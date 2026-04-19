[![Go Report Card](https://goreportcard.com/badge/github.com/yarencheng/go-bash-wasm)](https://goreportcard.com/report/github.com/yarencheng/go-bash-wasm)
[![Go Reference](https://pkg.go.dev/badge/github.com/yarencheng/go-bash-wasm.svg)](https://pkg.go.dev/github.com/yarencheng/go-bash-wasm)
![Go Version](https://img.shields.io/github/go-mod/go-version/yarencheng/go-bash-wasm)
[![Build Status](https://github.com/yarencheng/go-bash-wasm/actions/workflows/go.yml/badge.svg)](https://github.com/yarencheng/go-bash-wasm/actions)
[![codecov](https://codecov.io/gh/yarencheng/go-bash-wasm/branch/main/graph/badge.svg)](https://codecov.io/gh/yarencheng/go-bash-wasm)
![GitHub License](https://img.shields.io/github/license/yarencheng/go-bash-wasm)

# go-bash-wasm

**go-bash-wasm** is a Go implementation of GNU Bash and Coreutils for WebAssembly, featuring a fully isolated in-memory filesystem. It enables running a shell environment in browsers and other sandboxed environments.

## 🚀 Key Features

- **Strict Upstream Parity**: 
  - Tracks **GNU Bash 5.3** for shell logic and syntax.
  - Targets **GNU Coreutils 9.10** for utility behavior (e.g., `ls` with support for nearly all standard flags).
- **WebAssembly Browser & Native**:
  - Compiled with `GOOS=js GOARCH=wasm` for in-browser execution.
  - Interactive **xterm.js** terminal integration with full stdin/stdout piping.
  - Platform-agnostic input handling ensures compatibility with standard Go and JS/WASM environments.
- **Pure In-Memory Filesystem (VFS)**:
  - Uses `afero` for a fully detached, in-memory filesystem hierarchy.
  - Zero Disk I/O: Enforces absolute host isolation.
- **Structured Observability**:
  - Integrated with `zerolog` for high-performance, structured logging (native) and clean browser console output.

## 🛠 Architecture

The project follows a clean, modular architecture:
- `cmd/go-bash-wasm/`: Entry points for execution.
  - `main.go`: Native Go CLI entry point.
  - `main_js.go`: WebAssembly entry point using `syscall/js`.
- `internal/`: Core shell and command implementations.
  - `internal/shell`: REPL and command execution logic.
  - `internal/commands`: High-parity utilities (ls, cat, grep, etc.).
- `ui/`: Modern Svelte 5 frontend with xterm.js integration.
- `ui.Dockerfile`: Production-ready multi-stage build for the browser environment (includes `wasm-opt`).
- `Dockerfile`: CLI-focused environment using the `wasip1/wasm` target and Wasmtime.

## ⚙️ Building and Running

### Prerequisites
- **Go 1.26+**
- **Node.js 20+** (for local UI development)
- **Docker** (Recommended) for clean, isolated builds.

### 🏠 Local Development

#### 1. Native CLI Shell
Run the simulator directly on your host machine using the native Go runtime:
```bash
go run ./cmd/go-bash-wasm/
```

#### 2. Browser UI (Svelte + WASM)
To develop the frontend locally:
1. **Compile WASM**:
   ```bash
   GOOS=js GOARCH=wasm go build -o ui/static/main.wasm ./cmd/go-bash-wasm/
   cp $(go env GOROOT)/lib/wasm/wasm_exec.js ui/static/
   ```
2. **Run Svelte App**:
   ```bash
   cd ui
   npm install
   npm run dev
   ```
Access at `http://localhost:5173`.

### 🐳 Docker Deployment

#### 1. Browser Terminal (Svelte + Nginx)
Built for the web, including WASM optimizations via `binaryen`.
```bash
# Build and run (mapped to port 8080)
docker build -t go-bash-ui -f ui.Dockerfile .
docker run -it --rm -p 8080:80 go-bash-ui
```
*Supports `OPTIMIZE=fast` (default) or `OPTIMIZE=small` build args.*

#### 2. Native CLI (Wasmtime)
Runs the shell in a secure `wasip1` container.
```bash
docker build -t go-bash-cli -f Dockerfile .
docker run -it --rm go-bash-cli
```

## 🧪 Testing

We ensure 100% behavioral parity through rigorous testing in both backend and frontend.

### 1. Go Backend Tests
Runs all unit tests for the shell and coreutils implementations:
```bash
go test -v ./...
```

### 2. UI Frontend Tests
Runs Svelte component and logic tests:
```bash
cd ui
npm run test
```

### 3. Full Docker Validation
You can run a full build/test cycle inside Docker to verify environment parity:
```bash
# This triggers both go tests and npm tests as part of the build
docker build -f ui.Dockerfile .
```

---
*Developed by the go-bash-wasm team. Aiming for 100% functional parity with GNU tools.*