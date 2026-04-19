---
description: Workflow guide for AI agents to audit upstream parity and maintain docs/coreutils/tasks.md
---

// turbo-all

This workflow instructs the AI agent on how to systematically review and identify missing parity features directly from the upstream `coreutils` repository, maintaining alignment checklists in `docs/coreutils/tasks.md`.

## Goal
Achieve high-fidelity functional parity with GNU Coreutils by identifying implementation gaps and tracking them systematically.

## Phase 1: Context Discovery
Examine the tracking file located at `docs/coreutils/tasks.md`.
- Review existing entries for the target command to avoid redundant research.
- Ensure the command is a Coreutils utility.

## Phase 2: Upstream Deep-Dive
Investigate the original implementation in the upstream Coreutils repository:
- **Coreutils binaries:** `third_party/coreutils/src/*.c` (e.g., `ls.c`, `cp.c`)

**Action:** Use `grep_search` to scan for:
- `long_options`: To find all available flags.
- `getopt` / `getopt_long`: To find flag parsing logic.
- `usage (`: To find the help text which often lists all supported features.
- Core logic functions (e.g., `do_ls`).

## Phase 3: Local Implementation Audit
Search the local codebase (`internal/`, `pkg/`, `cmd/`) for the Go implementation of the target command.
- Identify which flags are already handled.
- Check for existing tests to confirm implementation status.

## Phase 4: Gap Mapping & tasks.md Update
Log every identified flag and feature in a strict hierarchical checkbox list in `docs/coreutils/tasks.md`.

**Review Checkbox Syntax:**
- `[x]` : Feature is fully implemented and verified. **Requirement:** Link to the local Go implementation.
- `[ ]` : Feature is incomplete or missing. **Requirement:** Link to the specific file and line range in `third_party/coreutils/src/*` where this is handled upstream.
- `[-]` : Deliberately skipped. **Requirement:** State a brief rationale (e.g., *[Not Simulation-friendly] Relies on hardware IO*).

### Example Output for `docs/coreutils/tasks.md`

```markdown
## Parity: `ls`

- [x] Basic directory listing: `internal/commands/ls.go`
- [x] Flag `-l` (long format): `internal/commands/ls.go`
- [ ] Flag `--author`: `third_party/coreutils/src/ls.c:L...`
- [-] Flag `--color`: Handled via specialized ANSI mapping.
```

## Phase 5: Actionable Planning
After finalizing the audit, propose the next steps (e.g., "Implement `--author` flag for `ls`") and suggest a TDD approach.
