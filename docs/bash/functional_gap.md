# Functional Gap Map

This document tracks known functional gaps, intentional deviations, and implemented workarounds in the Go Bash Simulator. While `docs/task.md` serves as the primary progress tracker, this file provides detailed technical context for features that deviate from standard GNU/Bash behavior.

## Status Definitions

- `[x]` **Workaround Applied**: The feature is addressed using a simulator-specific solution rather than a 1:1 port.
- `[-]` **Deliberately Unsupported**: The feature is explicitly excluded (e.g., hardware-dependent, security-risk, or WASM-incompatible).
- `[ ]` **Unresolved / Decision Pending**: A gap has been identified, but the implementation strategy or priority remains to be decided.

---

## Gap Repository

### `pwd`
  
- `[x]` Flag `--help` (Global Dispatcher): `docs/bash/tasks.md`
  > Rationale: Built-in help is intercepted by the shell's help dispatcher to ensure a consistent instructional experience across all simulated commands.

### `read`

- `[x]` Flags `-s`, `-u`, `-e`, `-i` (Stubs): `internal/commands/read/read.go`
  > Rationale: 
  > - `-s` (Silent): Interactive terminal echoing is managed by the frontend (Xterm.js).
  > - `-u` (FD): Browser WASM runtime has limited support for arbitrary file descriptor redirection beyond standard streams.
  > - `-e` / `-i` (Readline): The simulator uses standard Go I/O rather than a full `readline` library for input.

### `times`

- `[x]` Basic Output (Simulation): `internal/commands/times/times.go`
  > Rationale: Accurate process accounting for child processes is not supported by the WebAssembly environment. The shell's own user and system times are simulated based on the time elapsed since startup.

### `ulimit`

- `[x]` Resource Management (Simulation): `internal/commands/ulimit/ulimit.go`
  > Rationale: Setting real resource limits (e.g., stack size, file descriptors) is not possible within the browser's WebAssembly sandbox. All limits are reported as static simulated values, and setting new limits is disabled.

### `wait`

- `[x]` Flags `-f`, `-n` (Ignored): `internal/commands/wait/wait.go:L1579`
  > Rationale: Job control in the simulator is currently basic. Synchronous waiting is implemented, but advanced polling flags are ignored to simplify the state-machine.

### `suspend`

- `[-]` Process Suspension (Unsupported): `internal/commands/suspend/suspend.go`
  > Rationale: WebAssembly processes in the browser cannot be suspended/resumed by the shell in the same way as OS-level processes.


*Last Updated: 2026-04-19* (Current Date)