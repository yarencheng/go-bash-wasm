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
- [ ] Flag `-L` (logical path): `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L291-316`
- [ ] Flag `-P` (physical path): `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L371-383`
- [-] Flag `--help`: Handled by the shell's global help dispatcher.

### `ls`

- [ ] Basic listing: Missing implementation
- [ ] Flag `-a` (`--all`): `third_party/coreutils/src/ls.c:L41`
- [ ] Flag `-l` (long format): `third_party/coreutils/src/ls.c`
- [ ] Flag `-h` (`--human-readable`): `third_party/coreutils/src/ls.c:L47`
- [ ] Flag `-R` (`--recursive`): `third_party/coreutils/src/ls.c`
- [ ] Color output (`--color`): `third_party/coreutils/src/ls.c`

### `echo`

- [ ] Basic output: Missing implementation
- [ ] Flag `-n` (no newline): `third_party/bash/builtins/echo.def:L146`
- [ ] Flag `-e` (enable backslash escapes): `third_party/bash/builtins/echo.def:L149`
- [ ] Flag `-E` (disable backslash escapes): `third_party/bash/builtins/echo.def:L152`

### `cat`

- [ ] Basic output: Missing implementation
- [ ] Flag `-n` (number lines): `third_party/coreutils/src/cat.c:L112`
- [ ] Flag `-b` (number non-blank): `third_party/coreutils/src/cat.c:L103`
- [ ] Flag `-s` (squeeze blank): `third_party/coreutils/src/cat.c:L115`
- [ ] Flag `-v` (show non-printing): `third_party/coreutils/src/cat.c:L127`

### `mkdir`

- [ ] Basic creation: Missing implementation
- [ ] Flag `-p` (parents): `third_party/coreutils/src/mkdir.c:L69`
- [ ] Flag `-m` (mode): `third_party/coreutils/src/mkdir.c:L65`
- [ ] Flag `-v` (verbose): `third_party/coreutils/src/mkdir.c:L74`

### `rm`

- [ ] Basic removal: Missing implementation
- [ ] Flag `-f` (force): `third_party/coreutils/src/rm.c:L137`
- [ ] Flag `-r` / `-R` (recursive): `third_party/coreutils/src/rm.c:L172`
- [ ] Flag `-i` (interactive): `third_party/coreutils/src/rm.c:L142`

### `cp`

- [ ] Basic copy: Missing implementation
- [ ] Flag `-r` / `-R` (recursive): `third_party/coreutils/src/cp.c:L250`
- [ ] Flag `-p` (preserve): `third_party/coreutils/src/cp.c:L234`
- [ ] Flag `-a` (archive): `third_party/coreutils/src/cp.c:L173`

### `mv`

- [ ] Basic move/rename: Missing implementation
- [ ] Flag `-f` (force): `third_party/coreutils/src/mv.c:L282`
- [ ] Flag `-i` (interactive): `third_party/coreutils/src/mv.c:L286`
- [ ] Flag `-n` (no-clobber): `third_party/coreutils/src/mv.c:L290`

### `cd` (builtin)

- [ ] Basic change directory: Missing implementation
- [ ] Flag `-L` (logical path): `third_party/bash/builtins/cd.def:L94`
- [ ] Flag `-P` (physical path): `third_party/bash/builtins/cd.def:L96`
- [ ] CDPATH support: `third_party/bash/builtins/cd.def:L84`

### `read` (builtin)

- [ ] Basic input: Missing implementation
- [ ] Flag `-r` (raw mode): `third_party/bash/builtins/read.def:L55`
- [ ] Flag `-p` (prompt): `third_party/bash/builtins/read.def:L53`
- [ ] Flag `-a` (array): `third_party/bash/builtins/read.def:L39`

### `exit` (builtin)

- [ ] Basic exit: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/exit.def:L25`

<!-- Add new audits below this line -->


