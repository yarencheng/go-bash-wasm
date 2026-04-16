# go-bash-wasm

**go-bash-wasm** is a high-fidelity, clean-room simulator of the GNU Bash shell and Coreutils, written in Go and optimized for WebAssembly (WASM). It brings the power of a standard UNIX environment to sandboxed ecosystems like browsers, edge computing, and secure server-side runtimes.

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
- `cmd/go-bash-wasm/main.go`: Entry point for native execution.
- `cmd/go-bash-wasm/main_js.go`: Entry point for JS/WASM execution (browser).
- `index.html`: Browser frontend using xterm.js.
- `internal/shell`: REPL and command execution logic with abstracted line reading.
- `internal/commands`: Registry and high-parity implementation of core utilities.

## ⚙️ Building and Running

### Prerequisites
- **Go 1.25+**
- **Docker** (Recommended) for containerized builds and web hosting.

### Run Locally (Native)
To start the interactive shell locally on your host machine:
```bash
go run ./cmd/go-bash-wasm/
```

### Build for WebAssembly (Browser)
To compile the project to a WASM binary for browser usage:
```bash
GOOS=js GOARCH=wasm go build -o main.wasm ./cmd/go-bash-wasm/
```

### Build & Run via Docker (Nginx)
To build the WASM binary and host it with an interactive terminal via Nginx:
```bash
docker build -t go-bash-wasm .
docker run -it --rm -p 8080:80 go-bash-wasm
```
Access the terminal at `http://localhost:8080`.

## 🧪 Testing

The project follows TDD (Test-Driven Development) to ensure 100% behavioral parity with upstream tools.
```bash
go test -v ./...
```

---
*Developed by the go-bash-wasm team. Aiming for 100% functional parity with GNU tools.*