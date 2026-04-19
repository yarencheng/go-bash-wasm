---
description: Workflow for documenting technical deviations, workarounds, and sandbox limitations in docs/coreutils/functional_gap.md
---

// turbo-all

This workflow guides the AI agent on how to capture and document technical rationales for Coreutils features that cannot be implemented 1:1 with upstream behavior, ensuring all deviations are justified and tracked in `docs/coreutils/functional_gap.md`.

## Goal
Provide a transparent record of the simulator's technical limitations and the strategic decisions behind specific Coreutils workarounds.

## Phase 1: Audit Alignment
Before starting, review `docs/coreutils/tasks.md` for the target command.
- Identify flags/features marked as `[-]` (skipped) or `[x]` (implemented but potentially via workaround).
- If a flag in the task list has a brief note like *[Not Simulation-friendly]*, it likely requires a detailed entry here.

## Phase 2: Technical Root Cause Analysis
Investigate why the Coreutils utility deviates from upstream:
- **Sandbox Barriers**: Does it require `syscalls` unavailable in WASM (e.g., `chown`, `mknod`, `ptrace`)?
- **Hardware/OS Dependency**: Does it rely on `/dev/` nodes or specific Linux kernel features?
- **Security Decisions**: Is the feature excluded to prevent sandbox escapes?

## Phase 3: Workaround Documentation
If a workaround exists:
- Locate the internal implementation in `internal/commands/`.
- Verify the specific logic that replaces the upstream behavior.

## Phase 4: functional_gap.md Update
Update the `docs/coreutils/functional_gap.md` file using the established status codes:

- `[x]` **Workaround**: For features implemented using simulator-specific logic. 
  - **Requirement**: Path to local Go code + Rationale.
- `[-]` **Unsupported**: For features explicitly excluded.
  - **Requirement**: Path to upstream C code + Detailed rationale (e.g., WASM limitation).
- `[ ]` **Pending**: For identified gaps without a clear resolution path.

### Example Entry

```markdown
### `ls`

- `[x]` Flag `--author` (Single User Simulation): `internal/commands/ls/ls.go:L472`
  > Rationale: The simulator currently models a single-user environment. The "author" field is hardcoded to the simulated user identity.
```

## Phase 5: Verification
Ensure that every entry in the gap map has a corresponding reference (even if just a placeholder) in `docs/coreutils/tasks.md` to maintain a single source of truth for progress.
