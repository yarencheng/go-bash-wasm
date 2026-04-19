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

### `cut`

- `[x]` Flag `-n` (Ignored): `internal/commands/cut/cut.go:L35`
  > Rationale: Historic `cut` versions used `-n` to prevent splitting multi-byte characters. Modern Go string handling and the simulator's UTF-8 focus make this manual check unnecessary.

### `id`

- `[-]` Flag `-Z` (No SELinux): `internal/commands/id/id.go:L29`
  > Rationale: SELinux context reporting relies on Linux Security Modules (LSM) which are not present in the WebAssembly runtime environment.
- `[x]` Flag `-a` (Ignored): `internal/commands/id/id.go:L30`
  > Rationale: Included only for backward compatibility with older versions of `id`.

### `install`

- `[x]` Flags `-c`, `-m`, `-o`, `-g` (Ignored/Partial): `internal/commands/install/install.go:L38`
  > Rationale: While ownership and mode flags are parsed, their effects are limited within the Afero MemMapFs sandbox. The `-c` flag (copy) is the default behavior and is strictly for compatibility.

### `ls`

- `[x]` Flag `--author` (Single User Simulation): `internal/commands/ls/ls.go:L472`
  > Rationale: The simulator currently models a single-user environment. The "author" field is hardcoded to the simulated user identity.
- `[x]` Flag `--color` (ANSI Simulation): `internal/commands/ls/ls.go:L515`
  > Rationale: Instead of relying on `dircolors` databases, the simulator uses a hardcoded ANSI color-mapping logic tailored for web-based terminal themes.

### `mktemp`

- `[x]` Flag `-q` (Ignored): `internal/commands/mktemp/mktemp.go:L28`
  > Rationale: The simulator environment is designed to be failure-tolerant for educational purposes. Detailed diagnostics are preferred over silent failures in this context.

### `stat`

- `[x]` Flag `-f` (Ignored): `internal/commands/stat/stat.go:L28`
  > Rationale: Virtual filesystem (Afero MemMapFs) status information is static or unavailable. Standard file status reporting is prioritized over filesystem metadata.

### System Management (`mount`, `umount`, `su`)

- `[-]` Commands (WASM/Sandbox Limitation): `docs/coreutils/tasks.md`
  > Rationale:
  > - **mount/umount**: The WebAssembly runtime lacks the necessary syscalls for block-device management and mounting.
  > - **su**: The simulator operates as a single-user environment. Privilege escalation is restricted to maintain sandbox integrity and simplify state management.

### Commonly Ignored Flags

Across multiple commands (`cp`, `mv`, `rm`, `chmod`, `chown`, `realpath`), several flags are implemented but intentionally perform no operation:

- `[x]` **Interactive/Force Flags** (`-i`, `-f`): `internal/commands/`
  > Rationale: The simulator is designed for non-interactive automation and educational purposes. Destructive actions are currently allowed without prompts to streamline the user experience in the browser.
- `[x]` **Path/Symlink Logic** (`-L`, `-P`, `-H`): `internal/commands/`
  > Rationale: The virtual filesystem (Afero) primarily handles logical path resolution. Physical vs. Logical path distinctions are ignored where the underlying implementation treats them as equivalent in the sandbox.
- `[x]` **Filesystem Hints** (`-f`, `--file-system`, `--dereference`): `internal/commands/`
  > Rationale: Metadata like filesystem type or raw mount information is not available or static in the WASM memory-mapped filesystem.

### `chcon` / `runcon`
  
- `[-]` SELinux Contexts (WASM Limitation): `internal/commands/chcon/chcon.go`, `internal/commands/runcon/runcon.go`
  > Rationale: SELinux security contexts are not supported by the browser's WebAssembly runtime or the virtual filesystem. These commands return errors or exit with status 1 when context manipulation is attempted.

### `dircolors`

- `[x]` Basic Output (Stub): `internal/commands/dircolors/dircolors.go`
  > Rationale: The simulator uses hardcoded ANSI color mappings in the shell and `ls` implementation. `dircolors` is provided as a stub to prevent script failures but does not yet influence the environment's color database.

### `find`

- `[x]` Basic Search (Workaround): `internal/commands/find/find.go`
  > Rationale: The simulator implements a limited subset of `find` functionality (supporting `-name` and `-type`). Advanced predicates like `-mtime`, `-exec`, or `-size` are currently pending.

### `grep`

- `[x]` Regex Support (Workaround): `internal/commands/grep/grep.go`
  > Rationale: Uses Go's `regexp` package which implements RE2 syntax. Some GNU-specific extensions or backreferences may not behave identically to the original `grep`.

### `mkfifo` / `mknod`

- `[x]` Device Creation (Stub): `internal/commands/mkfifo/mkfifo.go`, `internal/commands/mknod/mknod.go`
  > Rationale: The Afero MemMapFs does not support real FIFOs or device nodes. These commands create normal files with the appropriate mode bits as placeholders to allow scripts to proceed without error.

### `pinky`

- `[x]` User Information (Stub): `internal/commands/pinky/pinky.go`
  > Rationale: Reports static simulated information for the single user in the environment. Real finger/pinky protocols or multiple user tracking are not supported.

### `shred`

- `[x]` Data Erasure (Workaround): `internal/commands/shred/shred.go`
  > Rationale: In a memory-mapped virtual filesystem, hardware-level secure deletion is not possible. `shred` performs basic buffer overwriting (using fixed patterns) to simulate the command's behavior without providing actual physical security.

### `shuf`

- `[x]` Flag `--random-source` (Ignored): `internal/commands/shuf/shuf.go`
  > Rationale: Seeding for shuffle operations is handled by the WebAssembly environment's source of randomness (via `time.Now().UnixNano()`) rather than external polling files.

### `stdbuf`

- `[x]` Stream Buffering (Stub): `internal/commands/stdbuf/stdbuf.go`
  > Rationale: GNU `stdbuf` relies on `LD_PRELOAD` to intercept library calls, which is not possible in the WebAssembly runtime. Standard I/O streams are managed by Go's runtime and the simulator's shell logic.

### `stty`

- `[x]` TTY Configuration (Stub): `internal/commands/stty/stty.go`
  > Rationale: The simulator's terminal is a web frontend component (Xterm.js) and does not provide a raw Unix TTY device that can be configured via `ioctl` as required by `stty`.

### `ptx`

- `[x]` Permuted Index (Stub): `internal/commands/ptx/ptx.go`
  > Rationale: Currently implements a simplified output that mirrors input lines, rather than performing full permuted index generation.

---

*Last Updated: 2026-04-19* (Current Date)