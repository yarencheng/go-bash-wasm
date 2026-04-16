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

### `basename`

- [ ] Basic operation: Missing implementation
- [ ] Flag `-a` (multiple args): `third_party/coreutils/src/basename.c:L155`
- [ ] Flag `-s` (suffix): `third_party/coreutils/src/basename.c:L150`
- [ ] Flag `-z` (zero-terminated): `third_party/coreutils/src/basename.c:L159`

### `dirname`

- [ ] Basic operation: Missing implementation
- [ ] Flag `-z` (zero-terminated): `third_party/coreutils/src/dirname.c:L99`

### `env`

- [ ] Basic execution: Missing implementation
- [ ] Flag `-i` (ignore env): `third_party/coreutils/src/env.c:L790`
- [ ] Flag `-u` (unset): `third_party/coreutils/src/env.c:L793`
- [ ] Flag `-C` (chdir): `third_party/coreutils/src/env.c:L811`
- [ ] Flag `-S` (split-string): `third_party/coreutils/src/env.c:L814`

### `printenv`

- [ ] Basic output: Missing implementation
- [ ] Flag `-0` (null-terminated): `third_party/coreutils/src/printenv.c:L100`

### `uname`

- [ ] Basic output: Missing implementation
- [ ] Flag `-a` (all): `third_party/coreutils/src/uname.c:L233`
- [ ] Flag `-s` (kernel name): `third_party/coreutils/src/uname.c:L237`
- [ ] Flag `-m` (machine): `third_party/coreutils/src/uname.c:L253`

### `id`

- [ ] Basic output: Missing implementation
- [ ] Flag `-u` (user): `third_party/coreutils/src/id.c:L191`
- [ ] Flag `-g` (group): `third_party/coreutils/src/id.c:L182`
- [ ] Flag `-G` (groups): `third_party/coreutils/src/id.c:L197`
- [ ] Flag `-n` (name): `third_party/coreutils/src/id.c:L185`

### `whoami`

- [ ] Basic output: Missing implementation

### `sort`

- [ ] Basic sorting: Missing implementation
- [ ] Flag `-n` (numeric): `third_party/coreutils/src/sort.c:L464`
- [ ] Flag `-r` (reverse): `third_party/coreutils/src/sort.c:L478`
- [ ] Flag `-u` (unique): `third_party/coreutils/src/sort.c:L557`
- [ ] Flag `-k` (key): `third_party/coreutils/src/sort.c:L524`
- [ ] Flag `-t` (separator): `third_party/coreutils/src/sort.c:L544`

### `uniq`

- [ ] Basic filtering: Missing implementation
- [ ] Flag `-c` (count): `third_party/coreutils/src/uniq.c:L172`
- [ ] Flag `-d` (repeated): `third_party/coreutils/src/uniq.c:L176`
- [ ] Flag `-u` (unique): `third_party/coreutils/src/uniq.c:L207`
- [ ] Flag `-i` (ignore case): `third_party/coreutils/src/uniq.c:L198`

### `cut`

- [ ] Basic selection: Missing implementation
- [ ] Flag `-f` (fields): `third_party/coreutils/src/cut.c:L155`
- [ ] Flag `-d` (delimiter): `third_party/coreutils/src/cut.c:L151`
- [ ] Flag `-b` (bytes): `third_party/coreutils/src/cut.c:L143`
- [ ] Flag `-c` (chars): `third_party/coreutils/src/cut.c:L147`

### `tee`

- [ ] Basic copy: Missing implementation
- [ ] Flag `-a` (append): `third_party/coreutils/src/tee.c:L93`
- [ ] Flag `-i` (ignore interrupts): `third_party/coreutils/src/tee.c:L97`

### `alias` / `unalias`

- [ ] Basic management: Missing implementation
- [ ] Flag `-p` (alias): `third_party/bash/builtins/alias.def:L79`
- [ ] Flag `-a` (unalias): `third_party/bash/builtins/alias.def:L181`

### `type`

- [ ] Basic lookup: Missing implementation
- [ ] Flag `-a` (all): `third_party/bash/builtins/type.def:L150`
- [ ] Flag `-p` (path): `third_party/bash/builtins/type.def:L156`
- [ ] Flag `-t` (type): `third_party/bash/builtins/type.def:L160`
- [ ] Flag `-P` (force path): `third_party/bash/builtins/type.def:L164`

### `date`

- [ ] Basic output: Missing implementation
- [ ] Flag `-d` (date string): `third_party/coreutils/src/date.c:L501`
- [ ] Flag `-u` (UTC): `third_party/coreutils/src/date.c:L561`
- [ ] Custom format `+FORMAT`: `third_party/coreutils/src/date.c:L607`

### `seq`

- [ ] Basic sequence: Missing implementation
- [ ] Flag `-s` (separator): `third_party/coreutils/src/seq.c:L596`
- [ ] Flag `-w` (equal width): `third_party/coreutils/src/seq.c:L600`

### `tr`

- [ ] Basic translation: Missing implementation
- [ ] Flag `-d` (delete): `third_party/coreutils/src/tr.c:L300`
- [ ] Flag `-s` (squeeze): `third_party/coreutils/src/tr.c:L304`
- [ ] Flag `-c` (complement): `third_party/coreutils/src/tr.c:L296`

### `yes`

- [ ] Basic operation: Missing implementation

### `true` / `false`

- [ ] Basic operation: Missing implementation

<!-- Add new audits below this line -->









