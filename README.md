# go-bash-wasm

**go-bash-wasm** is a high-fidelity, clean-room simulator of the GNU Bash shell and Coreutils, written in Go and optimized for WebAssembly (WASM). It brings the power of a standard UNIX environment to sandboxed ecosystems like browsers, edge computing, and secure server-side runtimes.

## 🚀 Key Features

- **Strict Upstream Parity**: 
  - Tracks **GNU Bash 5.3** for shell logic and syntax.
  - Targets **GNU Coreutils 9.10** for utility behavior (e.g., `ls` with support for nearly all standard flags).
- **WebAssembly Native (`wasip1`)**:
  - Compiled with `GOOS=wasip1 GOARCH=wasm`.
  - Platform-agnostic input handling ensures compatibility with WASM runtimes (Wasmtime, browser) while preserving advanced terminal features on native OSs.
- **Pure In-Memory Filesystem (VFS)**:
  - Uses `afero` for a fully detached, in-memory filesystem hierarchy.
  - Zero Disk I/O: Enforces absolute host isolation.
- **Structured Observability**:
  - Integrated with `zerolog` for high-performance, structured logging across all shell operations and command executions.

## 🛠 Architecture

The project follows a clean, modular architecture:
- `cmd/go-bash-wasm`: Minimal entry point for the simulator.
- `internal/app`: Central application bootstrap and lifecycle management.
- `internal/shell`: REPL and command execution logic with abstracted line reading.
- `internal/commands`: Registry and high-parity implementation of core utilities (starting with `ls`).

## ⚙️ Building and Running

### Prerequisites
- **Go 1.25+**
- (Optional) **Docker** for containerized WASM builds.
- (Optional) **Wasmtime** for running the compiled WASM binary.

### Run Locally (Native)
To start the interactive shell locally on your host machine:
```bash
go run ./cmd/go-bash-wasm/
```

### Build for WebAssembly
To compile the project to a WASM binary compatible with WASI Preview 1:
```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm ./cmd/go-bash-wasm/
```

### Build via Docker
A multi-stage `build.Dockerfile` is provided for automated WASM compilation and verification:
```bash
docker build -t go-bash-wasm -f build.Dockerfile .
```

## 🧪 Testing

The project follows TDD (Test-Driven Development) to ensure 100% behavioral parity with upstream tools.
```bash
go test -v ./...
```

---
*Developed by the go-bash-wasm team. Aiming for 100% functional parity with GNU tools.*