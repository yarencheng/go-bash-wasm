---
description: Comprehensive workflow guide for AI agents to audit upstream parity and maintain docs/task.md
---

// turbo-all

# FakeBash Parity Review Workflow

This workflow instructs the AI agent on how to systematically review and identify missing parity features directly from the upstream `bash` and `coreutils` repositories, maintaining an alignment checklist in `docs/task.md`.

## Step 1: Research Upstream Specifications
When asked to review functional parity for a specific shell command or builtin, you must first investigate its original implementation in the upstream repositories:
- **For Bash builtins & logic:** Search within `third_party/bash/`
- **For Coreutils binaries:** Search within `third_party/coreutils/`

**Action:** Use specific tools like `grep_search` to scan the upstream C source code. Look specifically for flag definitions (e.g., `getopt_long` usages) to discover all supported properties.

## Step 2: Prepare the Tracking Document
Examine the master tracking file located at `docs/task.md`.
- If the file does not exist, initialize it.
- Never delete existing items unrelated to your current review scope.

## Step 3: Map the Parity Items
Log every identified flag and feature in a strict hierarchical checkbox list under the targeted command's heading. Ensure the logic path mapping is included for items needing work.

**Review Checkbox Syntax:**
- `[x]` : Feature is fully implemented in the Go simulator.
- `[ ]` : Feature is incomplete or missing. **Requirement:** Include a specific reference path/file to the logic in `third_party/*` where this is handled upstream.
- `[-]` : Deliberately skipped or impossible to implement. **Requirement:** State a brief rationale (e.g., *[Not Simulation-friendly] Relies on pure hardware disk I/O*).

**Example Output:**
```markdown
## Parity of `ls`

- [x] Basic directory listing
- [ ] Flag `-a` (all files): `third_party/coreutils/src/ls.c` (see `decode_switches`)
- [ ] Flag `-l` (long format): `third_party/coreutils/src/ls.c`
- [-] Flag `--color`: [Not Simulation-friendly] Terminal color rendering is handled by the unified host console runner.
```
