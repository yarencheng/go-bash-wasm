# Functional Gap Map

This document tracks known functional gaps, intentional deviations, and implemented workarounds in the Go Bash Simulator. While `docs/task.md` serves as the primary progress tracker, this file provides detailed technical context for features that deviate from standard GNU/Bash behavior.

## Status Definitions

Use these codes to categorize gaps found during implementation or auditing:

- `[x]` **Workaround Applied**: The feature is addressed using a simulator-specific solution rather than a 1:1 port. Must include the local path and a brief technical rationale.
- `[-]` **Deliberately Unsupported**: The feature is explicitly excluded from the implementation plan (e.g., hardware-dependent, security-risk, or WASM-incompatible). Must include a rationale.
- `[ ]` **Unresolved / Decision Pending**: A gap has been identified, but the implementation strategy or priority remains to be decided.

---

## Instructions for AI Agents

When performing command audits or adding new features:
1. **Consistency**: Ensure any flags marked as `[-]` or `[x]` in `docs/task.md` are documented here if they require technical explanation.
2. **Identification**: If you discover a flag that relies on unavailable system calls (e.g., `ptrace`, `mknod`), mark it as `[-]` and explain why.
3. **Traceability**: Always link to the upstream line number in `third_party/` for every gap.
4. **Implementation Focus**: Use this document to justify shortcuts or deviations taken during development.

---

## Gap Repository

### Example: `my-command`

*Detailed analysis of deviations for `my-command` implementation.*

- `[x]` Flag `-a` (Simulated via Internal Buffer): `internal/commands/my_cmd/workaround.go:L45`
  > Rationale: Upstream `-a` requires raw device access; the simulator uses an in-memory buffer to mimic this behavior.
- `[-]` Flag `-b` (Sandbox Limitation): `third_party/coreutils/src/my_cmd.c:L120`
  > Rationale: Relies on `sys/mount.h` structures which are not available in the WebAssembly sandbox.
- `[ ]` Flag `-c` (Feasibility Under Review): `third_party/coreutils/src/my_cmd.c:L201`
  > Status: Investigating if `atob` can be used to handle this encoding gap.

---

<!-- Gaps will be appended below this line -->