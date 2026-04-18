# Functional Gap Map

This document tracks known functional gaps, intentional deviations, and implemented workarounds in the Go Bash Simulator. While `docs/task.md` serves as the primary progress tracker, this file provides detailed technical context for features that deviate from standard GNU/Bash behavior.

## Status Definitions

- `[x]` **Workaround Applied**: The feature is addressed using a simulator-specific solution rather than a 1:1 port.
- `[-]` **Deliberately Unsupported**: The feature is explicitly excluded (e.g., hardware-dependent, security-risk, or WASM-incompatible).
- `[ ]` **Unresolved / Decision Pending**: A gap has been identified, but the implementation strategy or priority remains to be decided.

---

## Gap Repository

### `cat`

- `[x]` Flag `-u` (Ignored): `internal/commands/cat/cat.go`
  > Rationale: Standard GNU behavior uses `-u` for unbuffered I/O. In this simulator, Go's `io.Copy` and the virtual filesystem layer handle buffering at the stream level, making this flag redundant.

### `chmod` / `chown`

- `[-]` Flag `--preserve-root` (Sandbox Context): `internal/commands/chmod/chmod.go`
  > Rationale: The WebAssembly sandbox operates on a virtual filesystem (Afero). "Root" preservation flags are omitted as the risk of accidental host-system corruption is mitigated by the sandbox itself.
- `[-]` Flag `-h`, `--no-dereference` (Simulated FS): `internal/commands/chown/chown.go`
  > Rationale: Current `Afero` memory-mapped filesystem implementations handle symlinks via logical resolution; raw symlink attribute mutation is restricted.

### `id`

- `[-]` Flag `-Z` (No SELinux): `internal/commands/id/id.go:L43`
  > Rationale: SELinux context reporting relies on Linux Security Modules (LSM) which are not present in the WebAssembly runtime environment.

### `ls`

- `[x]` Flag `--author` (Single User Simulation): `internal/commands/ls/ls.go:L472`
  > Rationale: The simulator currently models a single-user environment. The "author" field is hardcoded to the simulated user identity.
- `[x]` Flag `--color` (ANSI Simulation): `internal/commands/ls/ls.go:L515`
  > Rationale: Instead of relying on `dircolors` databases, the simulator uses a hardcoded ANSI color-mapping logic tailored for web-based terminal themes.

### `pwd`

- `[x]` Flag `--help` (Global Dispatcher): `docs/task.md:L1065`
  > Rationale: Built-in help is intercepted by the shell's help dispatcher to ensure a consistent instructional experience across all simulated commands.

### `wait`

- `[x]` Flags `-f`, `-n` (Ignored): `internal/commands/wait/wait.go`
  > Rationale: Job control in the simulator is currently basic. Synchronous waiting is implemented, but advanced polling flags are ignored to simplify the state-machine.

### Commonly Ignored Flags

Across multiple commands (`cp`, `mv`, `rm`, `chmod`, `chown`, `realpath`), several flags are implemented but intentionally perform no operation:

- `[x]` **Interactive/Force Flags** (`-i`, `-f`): `internal/commands/`
  > Rationale: The simulator is designed for non-interactive automation and educational purposes. Destructive actions are currently allowed without prompts to streamline the user experience in the browser.
- `[x]` **Path/Symlink Logic** (`-L`, `-P`, `-H`): `internal/commands/`
  > Rationale: The virtual filesystem (Afero) primarily handles logical path resolution. Physical vs. Logical path distinctions are ignored where the underlying implementation treats them as equivalent in the sandbox.
- `[x]` **Filesystem Hints** (`-f`, `--file-system`, `--dereference`): `internal/commands/`
  > Rationale: Metadata like filesystem type or raw mount information is not available or static in the WASM memory-mapped filesystem.

---

*Last Updated: 2026-04-18*