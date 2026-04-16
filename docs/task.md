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

### `head`

- [ ] Basic output: Missing implementation
- [ ] Flag `-n` (lines): `third_party/coreutils/src/head.c:L126`
- [ ] Flag `-c` (bytes): `third_party/coreutils/src/head.c:L121`
- [ ] Flag `-q` (quiet): `third_party/coreutils/src/head.c:L131`

### `tail`

- [ ] Basic output: Missing implementation
- [ ] Flag `-n` (lines): `third_party/coreutils/src/tail.c:L314`
- [ ] Flag `-c` (bytes): `third_party/coreutils/src/tail.c:L296`
- [ ] Flag `-f` (follow): `third_party/coreutils/src/tail.c:L305`

### `wc`

- [ ] Basic counts: Missing implementation
- [ ] Flag `-l` (lines): `third_party/coreutils/src/wc.c:L196`
- [ ] Flag `-w` (words): `third_party/coreutils/src/wc.c:L214`
- [ ] Flag `-c` (bytes): `third_party/coreutils/src/wc.c:L188`
- [ ] Flag `-m` (chars): `third_party/coreutils/src/wc.c:L192`

### `chmod`

- [ ] Basic mode change: Missing implementation
- [ ] Flag `-R` (recursive): `third_party/coreutils/src/chmod.c:L459`
- [ ] Numeric mode support: `third_party/coreutils/src/chmod.c:L415`
- [ ] Symbolic mode support: `third_party/coreutils/src/chmod.c:L414`

### `chown`

- [ ] Basic ownership change: Missing implementation
- [ ] Flag `-R` (recursive): `third_party/coreutils/src/chown.c:L141`
- [ ] Flag `--from`: `third_party/coreutils/src/chown.c:L121`

### `ln`

- [ ] Basic link creation: Missing implementation
- [ ] Flag `-s` (symbolic): `third_party/coreutils/src/ln.c:L574`
- [ ] Flag `-f` (force): `third_party/coreutils/src/ln.c:L553`
- [ ] Flag `-v` (verbose): `third_party/coreutils/src/ln.c:L595`

### `touch`

- [ ] Basic touch: Missing implementation
- [ ] Flag `-a` (access time): `third_party/coreutils/src/touch.c:L299`
- [ ] Flag `-m` (mod time): `third_party/coreutils/src/touch.c:L318`
- [ ] Flag `-c` (no create): `third_party/coreutils/src/touch.c:L303`

### `rmdir`

- [ ] Basic rmdir: Missing implementation
- [ ] Flag `-p` (parents): `third_party/coreutils/src/rmdir.c:L215`
- [ ] Flag `-v` (verbose): `third_party/coreutils/src/rmdir.c:L221`

### `du`

- [ ] Basic du: Missing implementation
- [ ] Flag `-s` (summarize): `third_party/coreutils/src/du.c:L373` (usage) / `L787` (main)
- [ ] Flag `-h` (human-readable): `third_party/coreutils/src/du.c:L337`
- [ ] Flag `-a` (all files): `third_party/coreutils/src/du.c:L294`

### `df`

- [ ] Basic df: Missing implementation
- [ ] Flag `-h` (human-readable): `third_party/coreutils/src/df.c:L1644`
- [ ] Flag `-a` (all files): `third_party/coreutils/src/df.c:L1625`

### `printf`

- [ ] Basic formatting: Missing implementation
- [ ] Flag `-v VAR` (assign to variable): `third_party/bash/builtins/printf.def:L301`
- [ ] Format `%b` (expand escapes): `third_party/bash/builtins/printf.def:L558`
- [ ] Format `%q` (shell quote): `third_party/bash/builtins/printf.def:L672`

### `test` / `[`

- [ ] Unary operators (-e, -f, -d, etc.): Missing implementation
- [ ] String operators (=, !=, -z, -n): Missing implementation
- [ ] Numeric operators (-eq, -ne, etc.): Missing implementation
- [ ] Logical operators (!, -a, -o): Missing implementation

### `sleep`

- [ ] Basic sleep: Missing implementation
- [ ] Multiple arguments (sum): `third_party/coreutils/src/sleep.c:L135`
- [ ] Suffixes (s, m, h, d): `third_party/coreutils/src/sleep.c:L65`

<!-- Add new audits below this line -->





