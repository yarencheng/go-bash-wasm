# Functional Gap Map

This document tracks known functional gaps, intentional deviations, and implemented workarounds in the Go Bash Simulator. While `docs/task.md` serves as the primary progress tracker, this file provides detailed technical context for features that deviate from standard GNU/Bash behavior.

## Status Definitions

- `[x]` **Workaround Applied**: The feature is addressed using a simulator-specific solution rather than a 1:1 port.
- `[-]` **Deliberately Unsupported**: The feature is explicitly excluded (e.g., hardware-dependent, security-risk, or WASM-incompatible).
- `[ ]` **Unresolved / Decision Pending**: A gap has been identified, but the implementation strategy or priority remains to be decided.

---

## Gap Repository

### `cat`

- `[x]` Flag `-u` (Ignored): `internal/commands/coreutils/cat/cat.go`
  > Rationale: Standard GNU behavior uses `-u` for unbuffered I/O. In this simulator, Go's `io.Copy` and the virtual filesystem layer handle buffering at the stream level, making this flag redundant.

### `chmod` / `chown`

- `[-]` Flag `--preserve-root` (Sandbox Context): `internal/commands/coreutils/chmod/chmod.go`
  > Rationale: The WebAssembly sandbox operates on a virtual filesystem (Afero). "Root" preservation flags are omitted as the risk of accidental host-system corruption is mitigated by the sandbox itself.
- `[-]` Flag `-h`, `--no-dereference` (Simulated FS): `internal/commands/coreutils/chown/chown.go`
  > Rationale: Current `Afero` memory-mapped filesystem implementations handle symlinks via logical resolution; raw symlink attribute mutation is restricted.

### `cut`

- `[x]` Flag `-n` (Ignored): `internal/commands/coreutils/cut/cut.go:L35`
  > Rationale: Historic `cut` versions used `-n` to prevent splitting multi-byte characters. Modern Go string handling and the simulator's UTF-8 focus make this manual check unnecessary.

### `id`

- `[-]` Flag `-Z` (No SELinux): `internal/commands/coreutils/id/id.go:L29`
  > Rationale: SELinux context reporting relies on Linux Security Modules (LSM) which are not present in the WebAssembly runtime environment.
- `[x]` Flag `-a` (Ignored): `internal/commands/coreutils/id/id.go:L30`
  > Rationale: Included only for backward compatibility with older versions of `id`.

### `install`

- `[x]` Flags `-c`, `-m`, `-o`, `-g` (Ignored/Partial): `internal/commands/coreutils/install/install.go:L38`
  > Rationale: While ownership and mode flags are parsed, their effects are limited within the Afero MemMapFs sandbox. The `-c` flag (copy) is the default behavior and is strictly for compatibility.

- `[x]` Flag `--author` (Single User Simulation): `internal/commands/coreutils/ls/ls.go:L507`
  > Rationale: The simulator currently models a single-user environment. While the flag is not explicitly parsed, the "owner" and "group" fields are hardcoded to "root", effectively simulating a single-author system.
- `[x]` Flag `--color` (ANSI Simulation): `internal/commands/coreutils/ls/ls.go:L550`
  > Rationale: Instead of relying on `dircolors` databases, the simulator uses a hardcoded ANSI color-mapping logic tailored for web-based terminal themes.
- `[x]` Flags `-Z`, `-D`, `-N` (Ignored): `internal/commands/coreutils/ls/ls.go:L142-L144`
  > Rationale: Security contexts (SELinux), Emacs dired mode support, and raw literal printing are parsed for compatibility but perform no operations within the simulator's environment.

### `mktemp`

- `[x]` Flag `-q` (Ignored): `internal/commands/coreutils/mktemp/mktemp.go:L28`
  > Rationale: The simulator environment is designed to be failure-tolerant for educational purposes. Detailed diagnostics are preferred over silent failures in this context.

### `stat`

- `[x]` Flag `-f` (Ignored): `internal/commands/coreutils/stat/stat.go`
  > Rationale: Virtual filesystem (Afero MemMapFs) status information is static or unavailable. Standard file status reporting is prioritized over filesystem metadata.

### `df`

- `[x]` Flags `--sync`, `--no-sync`, `--output`, `-B`, `-t`, `-x` (Ignored): `internal/commands/coreutils/df/df.go:L42-L47`
  > Rationale: 
  > - **Sync**: The memory-mapped filesystem is always "synced" as there is no underlying physical disk or delayed write buffer.
  > - **Output/Type**: Advanced formatting and filesystem type filtering are restricted to a simplified static set of simulated mounts.

### `du`

- `[x]` Flags `-x`, `-X`, `--exclude`, `--files0-from`, `--time` (Ignored): `internal/commands/coreutils/du/du.go:L56-L61`
  > Rationale: 
  > - **Exclusions**: Advanced pattern-based exclusion and external file-list reading are omitted to keep the recursive walker implementation focused on core usage calculation.
  > - **Times**: File modification times in the sandbox are handled by standard `os.FileInfo` but advanced formatting/styles are not yet linked to `du`.

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
- `[x]` **Chown/Chgrp Stubs** (`-H`, `-L`, `-P`, `--from`): `internal/commands/coreutils/chown/chown.go`
  > Rationale: Symlink dereference behavior and owner-based filtering for `chown` are currently stubbed. The simulator prioritizes simple path-based ownership changes over complex POSIX link-traversal logic.

### `chcon` / `runcon`
  
- `[-]` SELinux Contexts (WASM Limitation): `internal/commands/coreutils/chcon/chcon.go`, `internal/commands/coreutils/runcon/runcon.go`
  > Rationale: SELinux security contexts are not supported by the browser's WebAssembly runtime or the virtual filesystem. These commands return errors or exit with status 1 when context manipulation is attempted.

### `dircolors`

- `[x]` Basic Output (Stub): `internal/commands/coreutils/dircolors/dircolors.go`
  > Rationale: The simulator uses hardcoded ANSI color mappings in the shell and `ls` implementation. `dircolors` is provided as a stub to prevent script failures but does not yet influence the environment's color database.

### `find`

- `[x]` Basic Search (Workaround): `internal/commands/coreutils/find/find.go`
  > Rationale: The simulator implements a limited subset of `find` functionality (supporting `-name` and `-type`). Advanced predicates like `-mtime`, `-exec`, or `-size` are currently pending.

### `grep`

- `[x]` Regex Support (Workaround): `internal/commands/coreutils/grep/grep.go`
  > Rationale: Uses Go's `regexp` package which implements RE2 syntax. Some GNU-specific extensions or backreferences may not behave identically to the original `grep`.

### `mkfifo` / `mknod`

- `[x]` Device Creation (Stub): `internal/commands/coreutils/mkfifo/mkfifo.go`, `internal/commands/coreutils/mknod/mknod.go`
  > Rationale: The Afero MemMapFs does not support real FIFOs or device nodes. These commands create normal files with the appropriate mode bits as placeholders to allow scripts to proceed without error.

### `pinky`

- `[x]` User Information (Stub): `internal/commands/coreutils/pinky/pinky.go`
  > Rationale: Reports static simulated information for the single user in the environment. Real finger/pinky protocols or multiple user tracking are not supported.

### `shred`

- `[x]` Data Erasure (Workaround): `internal/commands/coreutils/shred/shred.go`
  > Rationale: In a memory-mapped virtual filesystem, hardware-level secure deletion is not possible. `shred` performs basic buffer overwriting (using fixed patterns) to simulate the command's behavior without providing actual physical security.

### `shuf`

- `[x]` Flag `--random-source` (Ignored): `internal/commands/coreutils/shuf/shuf.go`
  > Rationale: Seeding for shuffle operations is handled by the WebAssembly environment's source of randomness (via `time.Now().UnixNano()`) rather than external polling files.

### `stdbuf`

- `[x]` Stream Buffering (Stub): `internal/commands/coreutils/stdbuf/stdbuf.go`
  > Rationale: GNU `stdbuf` relies on `LD_PRELOAD` to intercept library calls, which is not possible in the WebAssembly runtime. Standard I/O streams are managed by Go's runtime and the simulator's shell logic.

### `stty`

- `[x]` TTY Configuration (Stub): `internal/commands/coreutils/stty/stty.go`
  > Rationale: The simulator's terminal is a web frontend component (Xterm.js) and does not provide a raw Unix TTY device that can be configured via `ioctl` as required by `stty`.

### `ptx`

- `[x]` Permuted Index (Stub): `internal/commands/coreutils/ptx/ptx.go`
  > Rationale: Currently implements a simplified output that mirrors input lines, rather than performing full permuted index generation.

---

*Last Updated: 2026-04-19* (Current Date)