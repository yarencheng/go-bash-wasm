# Functional Gap Map

This document tracks known functional gaps, intentional deviations, and implemented workarounds in the Go Bash Simulator. While `docs/task.md` serves as the primary progress tracker, this file provides detailed technical context for features that deviate from standard GNU/Bash behavior.

## Status Definitions

- `[x]` **Workaround Applied**: The feature is addressed using a simulator-specific solution rather than a 1:1 port.
- `[-]` **Deliberately Unsupported**: The feature is explicitly excluded (e.g., hardware-dependent, security-risk, or WASM-incompatible).
- `[ ]` **Unresolved / Decision Pending**: A gap has been identified, but the implementation strategy or priority remains to be decided.

---

## Gap Repository

### Commonly Ignored Flags

Across multiple Bash builtins, several flags are implemented in the parser but intentionally perform no operation to maintain script compatibility without introducing unnecessary complexity to the simulation:

- `[x]` **Readline Contexts** (`-E`, `-I`, `-D` in `complete`): Handled by frontend lookup logic.
- `[x]` **Physical vs. Logical Paths** (`-L`, `-P` in `pwd`, `cd`, `dirs`): The virtual filesystem primarily uses logical resolution.
- `[x]` **Process/Environment Hints** (`-l`, `-a`, `-c` in `exec`, `export`): Not applicable or not yet implemented in the simulator's process model.

---

### `complete`

- `[x]` Flags `-E`, `-I`, `-D` (Ignored): `internal/commands/bash/complete/complete.go`
  > Rationale: These flags control specific Readline completion contexts (empty lines, initial positions, default fallback). The current completion engine in the WASM frontend implements a simplified lookup logic that does not yet differentiate between these granular contexts.

### `dirs`

- `[x]` Flag `-l` (Ignored): `internal/commands/bash/dirs/dirs.go`
  > Rationale: The simulator's directory stack always uses logical path resolution. Toggling between physical and logical listing is redundant in this environment.

### `enable`

- `[-]` Flags `-d`, `-f` (Unsupported): `internal/commands/bash/enable/enable.go`
  > Rationale: GNU Bash uses these flags for dynamic loading of builtins from shared objects. The WebAssembly runtime and the simulator's static command registration model do not support dynamic object loading or symbol resolution at runtime.

### `exec`

- `[x]` Flags `-l`, `-a`, `-c` (Ignored): `internal/commands/bash/exec/exec.go`
  > Rationale: 
  > - `-l` (Login): The simulator doesn't distinguish between login and non-login shells in its process model.
  > - `-a` (Zeroth Argument): Command execution is handled via high-level Go interfaces; raw `execve` style argument manipulation for the zeroth index is not exposed.
  > - `-c` (Clean Environment): The simulator environment is persistent for the session; clearing it for a single command would require complex state-snapshotting not yet implemented.

### `export`

- `[x]` Flags `-f`, `-n` (Ignored): `internal/commands/bash/export/export.go`
  > Rationale: 
  > - `-f` (Functions): Function exporting is managed globally by the simulator's environment.
  > - `-n` (Remove): Export status is currently tied to variable presence in the `EnvVars` map; explicit removal of the export attribute while keeping the variable value is pending.

### `hash`

- `[x]` Flags `-d`, `-p`, `-t`, `-l` (Ignored): `internal/commands/bash/hash/hash.go`
  > Rationale: Hashing in the simulator is managed by the shell's command resolution cache. Manual manipulation of this cache via flags is currently stubbed to prevent script errors, while basic `-r` (reset) is supported.

### `pwd`
  
- `[x]` Flag `--help` (Global Dispatcher): `docs/bash/tasks.md`
  > Rationale: Built-in help is intercepted by the shell's help dispatcher to ensure a consistent instructional experience across all simulated commands.

### `read`

- `[x]` Flags `-s`, `-u`, `-e`, `-i` (Stubs): `internal/commands/bash/read/read.go`
  > Rationale: 
  > - `-s` (Silent): Interactive terminal echoing is managed by the frontend (Xterm.js).
  > - `-u` (FD): Browser WASM runtime has limited support for arbitrary file descriptor redirection beyond standard streams.
  > - `-e` / `-i` (Readline): The simulator uses standard Go I/O rather than a full `readline` library for input.

### `set`

- `[x]` Positional Parameters (Stub): `internal/commands/bash/set/set.go`
  > Rationale: Setting positional parameters via `set -- arg1 arg2` is partially implemented in the parser but the full update of the shell's positional parameter state is still being refined in the execution engine.

### `shopt`

- `[x]` Flag `-o` (Ignored): `internal/commands/bash/shopt/shopt.go`
  > Rationale: `set -o` is the primary mechanism for managing shell options in the simulator. `shopt -o` is provided as a stub for compatibility but redirects users to the `set` builtin logic.

### `suspend`

- `[-]` Process Suspension (Unsupported): `internal/commands/bash/suspend/suspend.go`
  > Rationale: WebAssembly processes in the browser cannot be suspended/resumed by the shell in the same way as OS-level processes.

### `times`

- `[x]` Basic Output (Simulation): `internal/commands/bash/times/times.go`
  > Rationale: Accurate process accounting for child processes is not supported by the WebAssembly environment. The shell's own user and system times are simulated based on the time elapsed since startup.

### `trap`

- `[x]` Flag `-P` (Ignored): `internal/commands/bash/trap/trap.go`
  > Rationale: The simulator uses a simplified signal-to-trap mapping. Advanced POSIX-specific trap listing formats are simplified to a standard output that covers most common use cases.

### `ulimit`

- `[x]` Resource Management (Simulation): `internal/commands/bash/ulimit/ulimit.go`
  > Rationale: Setting real resource limits (e.g., stack size, file descriptors) is not possible within the browser's WebAssembly sandbox. All limits are reported as static simulated values, and setting new limits is disabled.
- `[x]` Flags `-S`, `-H` (Ignored): `internal/commands/bash/ulimit/ulimit.go`
  > Rationale: The browser environment does not distinguish between "soft" and "hard" limits. All reported values are simulated constants.

### `unset`

- `[x]` Flag `-n` (Ignored): `internal/commands/bash/unset/unset.go`
  > Rationale: Nameref support (`unset -n`) is currently restricted as the simulator's environment storage model treats all variables as direct references.

### `wait`

- `[x]` Flags `-f`, `-n` (Ignored): `internal/commands/bash/wait/wait.go`
  > Rationale: Job control in the simulator is currently basic. Synchronous waiting is implemented, but advanced polling flags are ignored to simplify the state-machine.

### `coproc`

- `[-]` Asynchronous Coprocesses (Unsupported): `docs/bash/tasks.md`
  > Rationale: The simulator's execution model handles foreground and background processes via Go goroutines, but the complex bi-directional pipe management required for `coproc` is not supported in the current architecture.

### Mail Notification

- `[-]` MAILCHECK, MAILPATH (Unsupported): `docs/bash/tasks.md`
  > Rationale: The browser-based simulator does not implement a mail delivery system or mailbox monitoring. These variables are ignored.

*Last Updated: 2026-04-19*