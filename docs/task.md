# Functional Parity Tracking

This document tracks the alignment of the Go Bash Simulator with upstream GNU implementations.

## Overview
Status codes:
- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

---

## Parity Matrix

### `pwd`

- [ ] Basic path reporting: Missing implementation
- [ ] Flag `-L` (logical path): `third_party/bash/builtins/cd.def:L435`
- [ ] Flag `-P` (physical path): `third_party/bash/builtins/cd.def:L435`
- [-] Flag `--help`: Handled by the shell's global help dispatcher.

<!-- Add new audits below this line -->
