---
description: Comprehensive workflow guide for AI agents to audit upstream parity and maintain docs/task.md
---

// turbo-all

This workflow instructs the AI agent on how to systematically review and identify missing parity features directly from the upstream `bash` and `coreutils` repositories, maintaining an alignment checklist in `docs/task.md`.

## Goal
Achieve high-fidelity functional parity with GNU Bash and Coreutils by identifying implementation gaps and tracking them systematically.

## Phase 1: Context Discovery
Examine the master tracking file located at `docs/task.md`.
- If the file does not exist, initialize it with a basic structure.
- Review existing entries for the target command to avoid redundant research.

## Phase 2: Upstream Deep-Dive
Investigate the original implementation in the upstream repositories:
- **Bash builtins & logic:** `third_party/bash/*.c` (e.g., `alias.c`, `builtins/*.def`)
- **Coreutils binaries:** `third_party/coreutils/src/*.c` (e.g., `ls.c`, `cp.c`)

**Action:** Use `grep_search` to scan for:
- `long_options`: To find all available flags.
- `getopt` / `getopt_long`: To find flag parsing logic.
- `usage (`: To find the help text which often lists all supported features.
- Core logic functions (e.g., `do_ls`, `pwd_builtin`).

## Phase 3: Local Implementation Audit
Search the local codebase (`internal/`, `pkg/`, `cmd/`) for the Go implementation of the target command.
- Identify which flags are already handled.
- Check for existing tests to confirm implementation status.

## Phase 4: Gap Mapping & task.md Update
Log every identified flag and feature in a strict hierarchical checkbox list in `docs/task.md`.

**Review Checkbox Syntax:**
- `[x]` : Feature is fully implemented and verified. **Requirement:** Link to the local Go implementation.
- `[ ]` : Feature is incomplete or missing. **Requirement:** Link to the specific file and line range in `third_party/*` where this is handled upstream.
- `[-]` : Deliberately skipped. **Requirement:** State a brief rationale (e.g., *[Not Simulation-friendly] Relies on hardware IO*).

### Example Output for `docs/task.md`

```markdown
## Parity: `pwd`

- [x] Basic path reporting: `internal/commands/pwd.go`
- [x] Flag `-L` (logical path): `internal/commands/pwd.go`
- [ ] Flag `-P` (physical path): `third_party/coreutils/src/pwd.c:L120-145`
- [-] Flag `--help`: Handled by the shell's global help dispatcher.
```

## Phase 5: Actionable Planning
After finalizing the audit, propose the next steps (e.g., "Implement `-P` flag for `pwd`") and suggest a TDD approach.