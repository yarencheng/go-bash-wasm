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

### `stat`

- [ ] Basic output: Missing implementation
- [ ] Flag `-L` (dereference): `third_party/coreutils/src/stat.c:L1927` (main)
- [ ] Flag `-f` (file system): `third_party/coreutils/src/stat.c:L1932`
- [ ] Flag `-c` (format): `third_party/coreutils/src/stat.c:L1921`
- [ ] Flag `-t` (terse): `third_party/coreutils/src/stat.c:L1936`

### `dd`

- [ ] Basic copy: Missing implementation
- [ ] Operand `if=FILE`: `third_party/coreutils/src/dd.c:L552` (usage)
- [ ] Operand `of=FILE`: `third_party/coreutils/src/dd.c:L561`
- [ ] Operand `bs=BYTES`: `third_party/coreutils/src/dd.c:L536`
- [ ] Operand `count=N`: `third_party/coreutils/src/dd.c:L546`
- [ ] Operand `conv=CONVS`: `third_party/coreutils/src/dd.c:L543`

### `expr`

- [ ] Arithmetic (+, -, *, /, %): Missing implementation
- [ ] Comparison (=, !=, <, <=, >, >=): Missing implementation
- [ ] Logical (| , &): Missing implementation
- [ ] String ops (match, substr, index, length): Missing implementation

### `readlink`

- [ ] Basic output: Missing implementation
- [ ] Flag `-f` (canonicalize): `third_party/coreutils/src/readlink.c:L135`
- [ ] Flag `-n` (no newline): `third_party/coreutils/src/readlink.c:L141`

### `realpath`

- [ ] Basic output: Missing implementation
- [ ] Flag `-e` (existing): `third_party/coreutils/src/realpath.c:L220`
- [ ] Flag `-m` (missing): `third_party/coreutils/src/realpath.c:L224`
- [ ] Flag `--relative-to`: `third_party/coreutils/src/realpath.c:L246`

### `kill`

- [ ] Basic signaling: Missing implementation
- [ ] Flag `-s SIGNAL`: `third_party/coreutils/src/kill.c:L262` / `third_party/bash/builtins/kill.def:L129`
- [ ] Flag `-l` (list): `third_party/coreutils/src/kill.c:L277` / `third_party/bash/builtins/kill.def:L114`

### `timeout`

- [ ] Basic output: Missing implementation
- [ ] Flag `-k` (kill-after): `third_party/coreutils/src/timeout.c:L531`
- [ ] Flag `-s` (signal): `third_party/coreutils/src/timeout.c:L539`

### `exec` (builtin)

- [ ] Basic execution: Missing implementation
- [ ] Flag `-a name`: `third_party/bash/builtins/exec.def:L120`
- [ ] Flag `-c` (clean env): `third_party/bash/builtins/exec.def:L114`

### `eval` (builtin)

- [ ] Basic execution: Missing implementation

### `declare` / `typeset` / `local` / `export` / `readonly` (builtins)

- [ ] Attribute management (-i, -r, -x, -a, -A): Missing implementation
- [ ] Flag `-p` (print): `third_party/bash/builtins/declare.def:L306`
- [ ] Flag `-g` (global): `third_party/bash/builtins/declare.def:L320`

### `set` / `unset` (builtins)

- [ ] Option management (-e, -u, -x, -o): Missing implementation
- [ ] Positional parameters: `third_party/bash/builtins/set.def:L784`
- [ ] Flag `-f` (functions): `third_party/bash/builtins/set.def:L799` (unset)

<!-- Add new audits below this line -->










### `comm`

- [ ] Basic comparison: Missing implementation
- [ ] Flag `-1` (suppress column 1): `third_party/coreutils/src/comm.c:L467`
- [ ] Flag `-2` (suppress column 2): `third_party/coreutils/src/comm.c:L468`
- [ ] Flag `-3` (suppress column 3): `third_party/coreutils/src/comm.c:L469`
- [ ] Flag `--check-order`: `third_party/coreutils/src/comm.c:L473`
- [ ] Flag `--nocheck-order`: `third_party/coreutils/src/comm.c:L474`
- [ ] Flag `--output-delimiter`: `third_party/coreutils/src/comm.c:L477`
- [ ] Flag `--total`: `third_party/coreutils/src/comm.c:L478`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/comm.c:L480`

### `join`

- [ ] Basic join: Missing implementation
- [ ] Flag `-a FILENUM`: `third_party/coreutils/src/join.c:L210`
- [ ] Flag `-e STRING`: `third_party/coreutils/src/join.c:L215`
- [ ] Flag `-i` (`--ignore-case`): `third_party/coreutils/src/join.c:L220`
- [ ] Flag `-j FIELD`: `third_party/coreutils/src/join.c:L224`
- [ ] Flag `-o FORMAT`: `third_party/coreutils/src/join.c:L228`
- [ ] Flag `-t CHAR`: `third_party/coreutils/src/join.c:L232`
- [ ] Flag `-v FILENUM`: `third_party/coreutils/src/join.c:L236`
- [ ] Flag `-1 FIELD`: `third_party/coreutils/src/join.c:L240`
- [ ] Flag `-2 FIELD`: `third_party/coreutils/src/join.c:L244`
- [ ] Flag `--check-order`: `third_party/coreutils/src/join.c:L248`
- [ ] Flag `--nocheck-order`: `third_party/coreutils/src/join.c:L253`
- [ ] Flag `--header`: `third_party/coreutils/src/join.c:L257`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/join.c:L262`

### `paste`

- [ ] Basic paste: Missing implementation
- [ ] Flag `-d` (`--delimiters`): `third_party/coreutils/src/paste.c:L468`
- [ ] Flag `-s` (`--serial`): `third_party/coreutils/src/paste.c:L473`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/paste.c:L479`

### `split`

- [ ] Basic split: Missing implementation
- [ ] Flag `-a` (`--suffix-length`): `third_party/coreutils/src/split.c:L238`
- [ ] Flag `-b` (`--bytes`): `third_party/coreutils/src/split.c:L246`
- [ ] Flag `-l` (`--lines`): `third_party/coreutils/src/split.c:L278`
- [ ] Flag `-C` (`--line-bytes`): `third_party/coreutils/src/split.c:L250`
- [ ] Flag `-n` (`--number`): `third_party/coreutils/src/split.c:L282`
- [ ] Flag `-d` (`--numeric-suffixes`): `third_party/coreutils/src/split.c:L254`
- [ ] Flag `-x` (`--hex-suffixes`): `third_party/coreutils/src/split.c:L262`
- [ ] Flag `-e` (`--elide-empty-files`): `third_party/coreutils/src/split.c:L270`
- [ ] Flag `-t` (`--separator`): `third_party/coreutils/src/split.c:L286`
- [ ] Flag `-u` (`--unbuffered`): `third_party/coreutils/src/split.c:L291`
- [ ] Flag `--filter`: `third_party/coreutils/src/split.c:L274`
- [ ] Flag `--verbose`: `third_party/coreutils/src/split.c:L295`

### `csplit`

- [ ] Basic split: Missing implementation
- [ ] Flag `-b` (`--suffix-format`): `third_party/coreutils/src/csplit.c:1423`
- [ ] Flag `-f` (`--prefix`): `third_party/coreutils/src/csplit.c:1427`
- [ ] Flag `-k` (`--keep-files`): `third_party/coreutils/src/csplit.c:1431`
- [ ] Flag `--suppress-matched`: `third_party/coreutils/src/csplit.c:1435`
- [ ] Flag `-n` (`--digits`): `third_party/coreutils/src/csplit.c:1439`
- [ ] Flag `-s` (`--quiet`): `third_party/coreutils/src/csplit.c:1443`
- [ ] Flag `-z` (`--elide-empty-files`): `third_party/coreutils/src/csplit.c:1447`

### `mktemp`

- [ ] Basic creation: Missing implementation
- [ ] Flag `-d` (`--directory`): `third_party/coreutils/src/mktemp.c:L76`
- [ ] Flag `-u` (`--dry-run`): `third_party/coreutils/src/mktemp.c:L80`
- [ ] Flag `-q` (`--quiet`): `third_party/coreutils/src/mktemp.c:L84`
- [ ] Flag `--suffix`: `third_party/coreutils/src/mktemp.c:L88`
- [ ] Flag `-p` (`--tmpdir`): `third_party/coreutils/src/mktemp.c:L93`
- [ ] Flag `-t` (deprecated): `third_party/coreutils/src/mktemp.c:101`

### `truncate`

- [ ] Basic truncation: Missing implementation
- [ ] Flag `-c` (`--no-create`): `third_party/coreutils/src/truncate.c:L82`
- [ ] Flag `-o` (`--io-blocks`): `third_party/coreutils/src/truncate.c:L85`
- [ ] Flag `-r` (`--reference`): `third_party/coreutils/src/truncate.c:L88`
- [ ] Flag `-s` (`--size`): `third_party/coreutils/src/truncate.c:L91`

### `source` / `.` (builtin)

- [ ] Basic sourcing: Missing implementation
- [ ] Flag `-p path`: `third_party/bash/builtins/source.def:L126`

### `return` (builtin)

- [ ] Basic return: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/return.def:L61`

### `shift` (builtin)

- [ ] Basic shift: Missing implementation
- [ ] Shifting n parameters: `third_party/bash/builtins/shift.def:L64`

### `trap` (builtin)

- [ ] Basic trapping: Missing implementation
- [ ] Flag `-l` (list): `third_party/bash/builtins/trap.def:L125`
- [ ] Flag `-p` (print): `third_party/bash/builtins/trap.def:L128`
- [ ] Flag `-P` (print action only): `third_party/bash/builtins/trap.def:L131`

### `umask` (builtin)

- [ ] Basic mask management: Missing implementation
- [ ] Flag `-p` (reusable input): `third_party/bash/builtins/umask.def:L91`
- [ ] Flag `-S` (symbolic): `third_party/bash/builtins/umask.def:L88`


### `sort`

- [ ] Basic sort: Missing implementation
- [ ] Ordering flags (`-b`, `-i`, `-d`, `-f`, `-g`, `-h`, `-n`, `-M`, `-R`, `-V`, `-r`): `third_party/coreutils/src/sort.c:L437-490`
- [ ] Flag `-c` / `-C` (`--check`): `third_party/coreutils/src/sort.c:L501-507`
- [ ] Flag `-k` (`--key`): `third_party/coreutils/src/sort.c:L524`
- [ ] Flag `-m` (`--merge`): `third_party/coreutils/src/sort.c:L528`
- [ ] Flag `-o` (`--output`): `third_party/coreutils/src/sort.c:L532`
- [ ] Flag `-u` (`--unique`): `third_party/coreutils/src/sort.c:L557`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/sort.c:L562`
- [ ] Flag `-S` (`--buffer-size`): `third_party/coreutils/src/sort.c:L540`
- [ ] Flag `-t` (`--field-separator`): `third_party/coreutils/src/sort.c:L544`
- [ ] Flag `-T` (`--temporary-directory`): `third_party/coreutils/src/sort.c:L548`
- [ ] Flag `--parallel`: `third_party/coreutils/src/sort.c:L553`

### `uniq`

- [ ] Basic filtering: Missing implementation
- [ ] Flag `-c` (`--count`): `third_party/coreutils/src/uniq.c:172`
- [ ] Flag `-d` (`--repeated`): `third_party/coreutils/src/uniq.c:176`
- [ ] Flag `-D`: `third_party/coreutils/src/uniq.c:180`
- [ ] Flag `-u` (`--unique`): `third_party/coreutils/src/uniq.c:206`
- [ ] Flag `-i` (`--ignore-case`: `third_party/coreutils/src/uniq.c:198`
- [ ] Flag `-f` (`--skip-fields`): `third_party/coreutils/src/uniq.c:189`
- [ ] Flag `-s` (`--skip-chars`): `third_party/coreutils/src/uniq.c:202`
- [ ] Flag `-w` (`--check-chars`): `third_party/coreutils/src/uniq.c:214`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/uniq.c:210`

### `expand`

- [ ] Basic conversion: Missing implementation
- [ ] Flag `-i` (`--initial`): `third_party/coreutils/src/expand.c:L78`
- [ ] Flag `-t` (`--tabs`): `third_party/coreutils/src/expand.c:L82`

### `unexpand`

- [ ] Basic conversion: Missing implementation
- [ ] Flag `-a` (`--all`): `third_party/coreutils/src/unexpand.c:L87`
- [ ] Flag `--first-only`: `third_party/coreutils/src/unexpand.c:L91`
- [ ] Flag `-t` (`--tabs`): `third_party/coreutils/src/unexpand.c:L95`

### `fmt`

- [ ] Basic formatting: Missing implementation
- [ ] Flag `-w` (`--width`): `third_party/coreutils/src/fmt.c:L299`
- [ ] Flag `-g` (`--goal`): `third_party/coreutils/src/fmt.c:L305`
- [ ] Flag `-c` (`--crown-margin`): `third_party/coreutils/src/fmt.c:L278`
- [ ] Flag `-t` (`--tagged-paragraph`): `third_party/coreutils/src/fmt.c:L291`
- [ ] Flag `-u` (`--uniform-spacing`): `third_party/coreutils/src/fmt.c:L295`
- [ ] Flag `-p` (`--prefix`): `third_party/coreutils/src/fmt.c:L282`
- [ ] Flag `-s` (`--split-only`): `third_party/coreutils/src/fmt.c:L287`

### `fold`

- [ ] Basic folding: Missing implementation
- [ ] Flag `-b` (`--bytes`): `third_party/coreutils/src/fold.c:L88`
- [ ] Flag `-c` (`--characters`): `third_party/coreutils/src/fold.c:L92`
- [ ] Flag `-s` (`--spaces`): `third_party/coreutils/src/fold.c:L96`
- [ ] Flag `-w` (`--width`): `third_party/coreutils/src/fold.c:L100`

### `alias` (builtin)

- [ ] Define/Display aliases: Missing implementation
- [ ] Flag `-p` (reusable format): `third_party/bash/builtins/alias.def:L36`

### `unalias` (builtin)

- [ ] Remove aliases: Missing implementation
- [ ] Flag `-a` (remove all): `third_party/bash/builtins/unalias.def:L165`

### `builtin` (builtin)

- [ ] Basic execution: Missing implementation

### `command` (builtin)

- [ ] Basic execution: Missing implementation
- [ ] Flag `-p` (default PATH): `third_party/bash/builtins/command.def:L33`
- [ ] Flag `-v` (identify command): `third_party/bash/builtins/command.def:L35`
- [ ] Flag `-V` (verbose description): `third_party/bash/builtins/command.def:L37`

### `type` (builtin)

- [ ] Basic identification: Missing implementation
- [ ] Flag `-a` (all): `third_party/bash/builtins/type.def:L32`
- [ ] Flag `-f` (no functions): `third_party/bash/builtins/type.def:L35`
- [ ] Flag `-p` (path only): `third_party/bash/builtins/type.def:L39`
- [ ] Flag `-P` (force path): `third_party/bash/builtins/type.def:L36`
- [ ] Flag `-t` (type): `third_party/bash/builtins/type.def:L41`


### `printenv`

- [ ] Basic output: Missing implementation
- [ ] Flag `-0` (`--null`): `third_party/coreutils/src/printenv.c:L70`

### `whoami`

- [ ] Basic output: Missing implementation

### `groups`

- [ ] Basic listing: Missing implementation
- [ ] Multiple users support: `third_party/coreutils/src/groups.c:L120`

### `id`

- [ ] Basic output: Missing implementation
- [ ] Flag `-a` (ignored): `third_party/coreutils/src/id.c:L101`
- [ ] Flag `-Z` (`--context`): `third_party/coreutils/src/id.c:L105`
- [ ] Flag `-g` (`--group`): `third_party/coreutils/src/id.c:L109`
- [ ] Flag `-G` (`--groups`): `third_party/coreutils/src/id.c:L113`
- [ ] Flag `-n` (`--name`): `third_party/coreutils/src/id.c:L117`
- [ ] Flag `-r` (`--real`): `third_party/coreutils/src/id.c:L121`
- [ ] Flag `-u` (`--user`): `third_party/coreutils/src/id.c:L125`
- [ ] Flag `-z` (`--zero`): `third_party/coreutils/src/id.c:L129`

### `tty`

- [ ] Basic output: Missing implementation
- [ ] Flag `-s` (`--silent`): `third_party/coreutils/src/tty.c:L71`

### `uname`

- [ ] Basic output: Missing implementation
- [ ] Flag `-a` (`--all`): `third_party/coreutils/src/uname.c:L123`
- [ ] Flag `-s` (`--kernel-name`): `third_party/coreutils/src/uname.c:L127`
- [ ] Flag `-n` (`--nodename`): `third_party/coreutils/src/uname.c:L130`
- [ ] Flag `-r` (`--kernel-release`): `third_party/coreutils/src/uname.c:L133`
- [ ] Flag `-v` (`--kernel-version`): `third_party/coreutils/src/uname.c:L136`
- [ ] Flag `-m` (`--machine`): `third_party/coreutils/src/uname.c:L139`
- [ ] Flag `-p` (`--processor`): `third_party/coreutils/src/uname.c:L142`
- [ ] Flag `-i` (`--hardware-platform`): `third_party/coreutils/src/uname.c:L145`
- [ ] Flag `-o` (`--operating-system`): `third_party/coreutils/src/uname.c:L148`

### `hostname`

- [ ] Basic output: Missing implementation
- [ ] Set hostname support: `third_party/coreutils/src/hostname.c:L95`

### `getopts` (builtin)

- [ ] Basic parsing: Missing implementation
- [ ] Silent mode support (`:`): `third_party/bash/builtins/getopts.def:L180`

### `hash` (builtin)

- [ ] Basic hashing: Missing implementation
- [ ] Flag `-d` (forget): `third_party/bash/builtins/hash.def:L32`
- [ ] Flag `-l` (reusable format): `third_party/bash/builtins/hash.def:L33`
- [ ] Flag `-p pathname` (set path): `third_party/bash/builtins/hash.def:L34`
- [ ] Flag `-r` (forget all): `third_party/bash/builtins/hash.def:L35`
- [ ] Flag `-t` (print location): `third_party/bash/builtins/hash.def:L36`

### `times` (builtin)

- [ ] Basic output: Missing implementation

### `wait` (builtin)

- [ ] Basic waiting: Missing implementation
- [ ] Flag `-n` (wait for any): `third_party/bash/builtins/wait.def:L131`
- [ ] Flag `-f` (force): `third_party/bash/builtins/wait.def:L134`
- [ ] Flag `-p var` (set pid var): `third_party/bash/builtins/wait.def:L137`


### `df`

- [ ] Basic output: Missing implementation
- [ ] Flag `-a` (`--all`): `third_party/coreutils/src/df.c:L256`
- [ ] Flag `-B` (`--block-size=SIZE`): `third_party/coreutils/src/df.c:L257`
- [ ] Flag `-h` (`--human-readable`): `third_party/coreutils/src/df.c:L259`
- [ ] Flag `-H` (`--si`): `third_party/coreutils/src/df.c:L260`
- [ ] Flag `-i` (`--inodes`): `third_party/coreutils/src/df.c:L258`
- [ ] Flag `-k` (1K blocks): `third_party/coreutils/src/df.c:L1307`
- [ ] Flag `-l` (`--local`): `third_party/coreutils/src/df.c:L261`
- [ ] Flag `--no-sync`: `third_party/coreutils/src/df.c:L266`
- [ ] Flag `--output[=FIELD_LIST]`: `third_party/coreutils/src/df.c:L262`
- [ ] Flag `-P` (`--portability`): `third_party/coreutils/src/df.c:L263`
- [ ] Flag `--sync`: `third_party/coreutils/src/df.c:L265`
- [ ] Flag `--total`: `third_party/coreutils/src/df.c:L267`
- [ ] Flag `-t` (`--type=TYPE`): `third_party/coreutils/src/df.c:L268`
- [ ] Flag `-T` (`--print-type`): `third_party/coreutils/src/df.c:L264`
- [ ] Flag `-x` (`--exclude-type=TYPE`): `third_party/coreutils/src/df.c:L269`

### `base32` / `base64` / `basenc`

- [ ] Basic encoding/decoding: Missing implementation
- [ ] Flag `-d` (`--decode`): `third_party/coreutils/src/basenc.c:L77`
- [ ] Flag `-i` (`--ignore-garbage`): `third_party/coreutils/src/basenc.c:L79`
- [ ] Flag `-w` (`--wrap=COLS`): `third_party/coreutils/src/basenc.c:L78`
- [ ] Flag `--base64`: `third_party/coreutils/src/basenc.c:L81`
- [ ] Flag `--base64url`: `third_party/coreutils/src/basenc.c:L82`
- [ ] Flag `--base58`: `third_party/coreutils/src/basenc.c:L83`
- [ ] Flag `--base32`: `third_party/coreutils/src/basenc.c:L84`
- [ ] Flag `--base32hex`: `third_party/coreutils/src/basenc.c:L85`
- [ ] Flag `--base16`: `third_party/coreutils/src/basenc.c:L86`
- [ ] Flag `--base2msbf`: `third_party/coreutils/src/basenc.c:L87`
- [ ] Flag `--base2lsbf`: `third_party/coreutils/src/basenc.c:L88`
- [ ] Flag `--z85`: `third_party/coreutils/src/basenc.c:L89`

### `nl`

- [ ] Basic numbering: Missing implementation
- [ ] Flag `-b` (`--body-numbering=STYLE`): `third_party/coreutils/src/nl.c:L153`
- [ ] Flag `-d` (`--section-delimiter=CC`): `third_party/coreutils/src/nl.c:L162`
- [ ] Flag `-f` (`--footer-numbering=STYLE`): `third_party/coreutils/src/nl.c:L154`
- [ ] Flag `-h` (`--header-numbering=STYLE`): `third_party/coreutils/src/nl.c:L152`
- [ ] Flag `-i` (`--line-increment=NUMBER`): `third_party/coreutils/src/nl.c:L156`
- [ ] Flag `-l` (`--join-blank-lines=NUMBER`): `third_party/coreutils/src/nl.c:L158`
- [ ] Flag `-n` (`--number-format=FORMAT`): `third_party/coreutils/src/nl.c:L161`
- [ ] Flag `-p` (`--no-renumber`: `third_party/coreutils/src/nl.c:L157`
- [ ] Flag `-s` (`--number-separator=STRING`): `third_party/coreutils/src/nl.c:L159`
- [ ] Flag `-v` (`--starting-line-number=NUMBER`): `third_party/coreutils/src/nl.c:L155`
- [ ] Flag `-w` (`--number-width=NUMBER`): `third_party/coreutils/src/nl.c:L160`

### `shuf`

- [ ] Basic shuffling: Missing implementation
- [ ] Flag `-e` (`--echo`): `third_party/coreutils/src/shuf.c:L107`
- [ ] Flag `-i` (`--input-range=LO-HI`): `third_party/coreutils/src/shuf.c:L108`
- [ ] Flag `-n` (`--head-count=COUNT`): `third_party/coreutils/src/shuf.c:L109`
- [ ] Flag `-o` (`--output=FILE`): `third_party/coreutils/src/shuf.c:L110`
- [ ] Flag `--random-source=FILE`: `third_party/coreutils/src/shuf.c:L111`
- [ ] Flag `-r` (`--repeat`): `third_party/coreutils/src/shuf.c:L112`
- [ ] Flag `-z` (`--zero-terminated`): `third_party/coreutils/src/shuf.c:L113`

### `tac`

- [ ] Basic output: Missing implementation
- [ ] Flag `-b` (`--before`): `third_party/coreutils/src/tac.c:L103`
- [ ] Flag `-r` (`--regex`): `third_party/coreutils/src/tac.c:L104`
- [ ] Flag `-s` (`--separator=STRING`): `third_party/coreutils/src/tac.c:L105`

### `shopt` (builtin)

- [ ] Basic option management: Missing implementation
- [ ] Flag `-o` (restrict to set -o): `third_party/bash/builtins/shopt.def:L316`
- [ ] Flag `-p` (reusable output): `third_party/bash/builtins/shopt.def:L319`
- [ ] Flag `-q` (suppress output): `third_party/bash/builtins/shopt.def:L313`
- [ ] Flag `-s` (enable): `third_party/bash/builtins/shopt.def:L307`
- [ ] Flag `-u` (disable): `third_party/bash/builtins/shopt.def:L310`

### `help` (builtin)

- [ ] Basic discovery: Missing implementation
- [ ] Flag `-d` (short description): `third_party/bash/builtins/help.def:L105`
- [ ] Flag `-m` (manpage format): `third_party/bash/builtins/help.def:L108`
- [ ] Flag `-s` (usage synopsis): `third_party/bash/builtins/help.def:L111`

### `fc` (builtin)

- [ ] Basic editing/re-execution: Missing implementation
- [ ] Flag `-e ENAME` (select editor): `third_party/bash/builtins/fc.def:L232`
- [ ] Flag `-l` (list): `third_party/bash/builtins/fc.def:L220`
- [ ] Flag `-n` (omit numbers): `third_party/bash/builtins/fc.def:L216`
- [ ] Flag `-r` (reverse): `third_party/bash/builtins/fc.def:L224`
- [ ] Flag `-s` (re-execute): `third_party/bash/builtins/fc.def:L228`

### `history` (builtin)

- [ ] Basic management: Missing implementation
- [ ] Flag `-c` (clear): `third_party/bash/builtins/history.def:L129`
- [ ] Flag `-d offset` (delete): `third_party/bash/builtins/history.def:L145`
- [ ] Flag `-a` (append to file): `third_party/bash/builtins/history.def:L126`
- [ ] Flag `-n` (read new): `third_party/bash/builtins/history.def:L132`
- [ ] Flag `-r` (read file): `third_party/bash/builtins/history.def:L135`
- [ ] Flag `-w` (write file): `third_party/bash/builtins/history.def:L138`
- [ ] Flag `-p` (expansion): `third_party/bash/builtins/history.def:L148`
- [ ] Flag `-s` (append entry): `third_party/bash/builtins/history.def:L141`

### `chown` / `chgrp`

- [ ] Basic ownership change: Missing implementation
- [ ] Flag `-c` (`--changes`): `third_party/coreutils/src/chown.c:L99`
- [ ] Flag `-f` (`--silent`, `--quiet`): `third_party/coreutils/src/chown.c:L103`
- [ ] Flag `-v` (`--verbose`): `third_party/coreutils/src/chown.c:L107`
- [ ] Flag `--dereference`: `third_party/coreutils/src/chown.c:L111`
- [ ] Flag `-h` (`--no-dereference`): `third_party/coreutils/src/chown.c:L116`
- [ ] Flag `--from=CURRENT_OWNER:CURRENT_GROUP`: `third_party/coreutils/src/chown.c:L121`
- [ ] Flag `--no-preserve-root`: `third_party/coreutils/src/chown.c:L128`
- [ ] Flag `--preserve-root`: `third_party/coreutils/src/chown.c:L131`
- [ ] Flag `--reference=RFILE`: `third_party/coreutils/src/chown.c:L134`
- [ ] Flag `-R` (`--recursive`): `third_party/coreutils/src/chown.c:L139`
- [ ] Flag `-H`: `third_party/coreutils/src/chown.c:L143`
- [ ] Flag `-L`: `third_party/coreutils/src/chown.c:L147`
- [ ] Flag `-P`: `third_party/coreutils/src/chown.c:L151`

### `chmod`

- [ ] Basic mode change: Missing implementation
- [ ] Flag `-c` (`--changes`): `third_party/coreutils/src/chmod.c:L425`
- [ ] Flag `-f` (`--silent`, `--quiet`): `third_party/coreutils/src/chmod.c:L429`
- [ ] Flag `-v` (`--verbose`): `third_party/coreutils/src/chmod.c:L433`
- [ ] Flag `--dereference`: `third_party/coreutils/src/chmod.c:L437`
- [ ] Flag `-h` (`--no-dereference`): `third_party/coreutils/src/chmod.c:L442`
- [ ] Flag `--no-preserve-root`: `third_party/coreutils/src/chmod.c:L446`
- [ ] Flag `--preserve-root`: `third_party/coreutils/src/chmod.c:L450`
- [ ] Flag `--reference=RFILE`: `third_party/coreutils/src/chmod.c:L454`
- [ ] Flag `-R` (`--recursive`): `third_party/coreutils/src/chmod.c:L459`

### `touch`

- [ ] Basic timestamp update: Missing implementation
- [ ] Flag `-a` (access time only): `third_party/coreutils/src/touch.c:L230`
- [ ] Flag `-c` (`--no-create`): `third_party/coreutils/src/touch.c:L234`
- [ ] Flag `-d` (`--date=STRING`): `third_party/coreutils/src/touch.c:L238`
- [ ] Flag `-h` (`--no-dereference`): `third_party/coreutils/src/touch.c:L246`
- [ ] Flag `-m` (modification time only): `third_party/coreutils/src/touch.c:L251`
- [ ] Flag `-r` (`--reference=FILE`): `third_party/coreutils/src/touch.c:L255`
- [ ] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `third_party/coreutils/src/touch.c:L259`

### `du`

- [ ] Basic usage summary: Missing implementation
- [ ] Flag `-0` (`--null`): `third_party/coreutils/src/du.c:L290`
- [ ] Flag `-a` (`--all`): `third_party/coreutils/src/du.c:L294`
- [ ] Flag `-A` (`--apparent-size`): `third_party/coreutils/src/du.c:L298`
- [ ] Flag `-B` (`--block-size=SIZE`): `third_party/coreutils/src/du.c:L305`
- [ ] Flag `-b` (`--bytes`): `third_party/coreutils/src/du.c:L310`
- [ ] Flag `-c` (`--total`): `third_party/coreutils/src/du.c:L314`
- [ ] Flag `-D` (`--dereference-args`): `third_party/coreutils/src/du.c:L318`
- [ ] Flag `-d` (`--max-depth=N`): `third_party/coreutils/src/du.c:L322`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/du.c:L328`
- [ ] Flag `-H`: `third_party/coreutils/src/du.c:L333`
- [ ] Flag `-h` (`--human-readable`): `third_party/coreutils/src/du.c:L337`
- [ ] Flag `--inodes`: `third_party/coreutils/src/du.c:L341`
- [ ] Flag `-k`: `third_party/coreutils/src/du.c:L345`
- [ ] Flag `-L` (`--dereference`): `third_party/coreutils/src/du.c:L349`
- [ ] Flag `-l` (`--count-links`): `third_party/coreutils/src/du.c:L353`
- [ ] Flag `-m`: `third_party/coreutils/src/du.c:L357`
- [ ] Flag `-P` (`--no-dereference`): `third_party/coreutils/src/du.c:L361`
- [ ] Flag `-S` (`--separate-dirs`): `third_party/coreutils/src/du.c:L365`
- [ ] Flag `--si`: `third_party/coreutils/src/du.c:L369`
- [ ] Flag `-s` (`--summarize`): `third_party/coreutils/src/du.c:L373`
- [ ] Flag `-t` (`--threshold=SIZE`): `third_party/coreutils/src/du.c:L377`
- [ ] Flag `--time[=WORD]`: `third_party/coreutils/src/du.c:L382`
- [ ] Flag `--time-style=STYLE`: `third_party/coreutils/src/du.c:L389`
- [ ] Flag `-X` (`--exclude-from=FILE`): `third_party/coreutils/src/du.c:L394`
- [ ] Flag `--exclude=PATTERN`: `third_party/coreutils/src/du.c:L398`
- [ ] Flag `-x` (`--one-file-system`: `third_party/coreutils/src/du.c:L402`

### `tr`

- [ ] Basic translation: Missing implementation
- [ ] Flag `-c`, `-C`, `--complement`: `third_party/coreutils/src/tr.c:L296`
- [ ] Flag `-d`, `--delete`: `third_party/coreutils/src/tr.c:L300`
- [ ] Flag `-s`, `--squeeze-repeats`: `third_party/coreutils/src/tr.c:L304`
- [ ] Flag `-t`, `--truncate-set1`: `third_party/coreutils/src/tr.c:L310`

### `rmdir`

- [ ] Basic removal: Missing implementation
- [ ] Flag `--ignore-fail-on-non-empty`: `third_party/coreutils/src/rmdir.c:L178`
- [ ] Flag `-p`, `--parents`: `third_party/coreutils/src/rmdir.c:L182`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/rmdir.c:L187`

### `link`

- [ ] Basic hard link: Missing implementation (exactly 2 args required)

### `unlink`

- [ ] Basic removal: Missing implementation (exactly 1 arg required)

### `yes`

- [ ] Basic repetition: Missing implementation

### `env`

- [ ] Basic execution: Missing implementation
- [ ] Flag `-a`, `--argv0=ARG`: `third_party/coreutils/src/env.c:L123`
- [ ] Flag `-i`, `--ignore-environment`: `third_party/coreutils/src/env.c:L127`
- [ ] Flag `-0`, `--null`: `third_party/coreutils/src/env.c:L131`
- [ ] Flag `-u`, `--unset=NAME`: `third_party/coreutils/src/env.c:L135`
- [ ] Flag `-C`, `--chdir=DIR`: `third_party/coreutils/src/env.c:L139`
- [ ] Flag `-S`, `--split-string=S`: `third_party/coreutils/src/env.c:L143`
- [ ] Flag `--block-signal[=SIG]`: `third_party/coreutils/src/env.c:L148`
- [ ] Flag `--default-signal[=SIG]`: `third_party/coreutils/src/env.c:L152`
- [ ] Flag `--ignore-signal[=SIG]`: `third_party/coreutils/src/env.c:L156`
- [ ] Flag `--list-signal-handling`: `third_party/coreutils/src/env.c:L160`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/env.c:L164`
### Batch 17: Multi-Utility Expansion (od, who, uptime, users, pinky, shred, cksum/sum, mknod, mkfifo, nproc, hostid, logname, pathchk, tsort, vdir, chroot, nice, nohup, stdbuf, runcon, timeout, truncate, factor, numfmt, pr, ptx)
- [ ] Audit `od` (octal dump)
    - Upstream: `third_party/coreutils/src/od.c`
    - Flags: `-A`, `-j`, `-N`, `-S`, `-t`, `-v`, `-w`
- [ ] Audit `who` (logged in users)
    - Upstream: `third_party/coreutils/src/who.c`
    - Flags: `-a`, `-b`, `-d`, `-H`, `-l`, `-m`, `-p`, `-q`, `-r`, `-s`, `-t`, `-u`, `-w`, `-y`
- [ ] Audit `uptime` (system runtime)
    - Upstream: `third_party/coreutils/src/uptime.c`
    - Flags: `-p`, `-s`
- [ ] Audit `users` (list users)
    - Upstream: `third_party/coreutils/src/users.c`
    - Flags: None specific beyond standard.
- [ ] Audit `pinky` (user info)
    - Upstream: `third_party/coreutils/src/pinky.c`
    - Flags: `-l`, `-b`, `-h`, `-p`, `-s`, `-f`, `-w`, `-i`, `-q`
- [ ] Audit `shred` (secure delete)
    - Upstream: `third_party/coreutils/src/shred.c`
    - Flags: `-f`, `-n`, `--random-source`, `-s`, `-u`, `-v`, `-x`, `-z`
- [ ] Audit `cksum` / `sum` (checksum)
    - Upstream: `third_party/coreutils/src/cksum.c`
    - Flags: `-a`, `--base64`, `-c`, `-l`, `--raw`, `--tag`, `--untagged`, `-z`, `--ignore-missing`, `--quiet`, `--status`, `--strict`, `-w`, `--debug`
- [ ] Audit `mknod` (make special files)
    - Upstream: `third_party/coreutils/src/mknod.c`
    - Flags: `-m`, `-Z`, `--context`
- [ ] Audit `mkfifo` (make FIFOs)
    - Upstream: `third_party/coreutils/src/mkfifo.c`
    - Flags: `-m`, `-Z`, `--context`
- [ ] Audit `nproc` (CPU count)
    - Upstream: `third_party/coreutils/src/nproc.c`
    - Flags: `--all`, `--ignore=N`
- [ ] Audit `hostid` (host identifier)
    - Upstream: `third_party/coreutils/src/hostid.c`
    - Flags: None specific.
- [ ] Audit `logname` (login name)
    - Upstream: `third_party/coreutils/src/logname.c`
    - Flags: None specific.
- [ ] Audit `pathchk` (path validity)
    - Upstream: `third_party/coreutils/src/pathchk.c`
    - Flags: `-p`, `-P`, `--portability`
- [ ] Audit `tsort` (topo sort)
    - Upstream: `third_party/coreutils/src/tsort.c`
    - Flags: None (found `-w` but it's a no-op).
- [ ] Audit `vdir` (ls -lb)
    - Upstream: `third_party/coreutils/src/ls.c` (as vdir)
    - Parity: Alias of ls with specific defaults.
- [ ] Audit `chroot` (change root)
    - Upstream: `third_party/coreutils/src/chroot.c`
    - Flags: `--groups`, `--userspec`, `--skip-chdir`
- [ ] Audit `nice` (scheduling priority)
    - Upstream: `third_party/coreutils/src/nice.c`
    - Flags: `-n`, `-N` (legacy)
- [ ] Audit `nohup` (immune to hangups)
    - Upstream: `third_party/coreutils/src/nohup.c`
    - Flags: None specific.
- [ ] Audit `stdbuf` (buffer control)
    - Upstream: `third_party/coreutils/src/stdbuf.c`
    - Flags: `-i`, `-o`, `-e`
- [ ] Audit `runcon` (security context)
    - Upstream: `third_party/coreutils/src/runcon.c`
    - Flags: `-r`, `-t`, `-u`, `-l`, `-c`
- [ ] Audit `timeout` (timed run)
    - Upstream: `third_party/coreutils/src/timeout.c`
    - Flags: `-f`, `-k`, `-p`, `-s`, `-v`
- [ ] Audit `truncate` (resize file)
    - Upstream: `third_party/coreutils/src/truncate.c`
    - Flags: `-c`, `-o`, `-r`, `-s`
- [ ] Audit `factor` (prime factors)
    - Upstream: `third_party/coreutils/src/factor.c`
    - Flags: `-h`, `--exponents`
- [ ] Audit `numfmt` (format numbers)
    - Upstream: `third_party/coreutils/src/numfmt.c`
    - Flags: `--debug`, `-d`, `--field`, `--format`, `--from`, `--from-unit`, `--grouping`, `--header`, `--invalid`, `--padding`, `--round`, `--suffix`, `--unit-separator`, `--to`, `--to-unit`, `-z`
- [ ] Audit `pr` (paginate)
    - Upstream: `third_party/coreutils/src/pr.c`
    - Flags: `+N`, `-N`, `-a`, `-c`, `-d`, `-D`, `-e`, `-F`, `-h`, `-i`, `-J`, `-l`, `-m`, `-n`, `-N`, `-o`, `-r`, `-s`, `-S`, `-t`, `-T`, `-v`, `-w`, `-W`
- [ ] Audit `ptx` (permuted index)
    - Upstream: `third_party/coreutils/src/ptx.c`
    - Flags: `-A`, `-G`, `-F`, `-M`, `-O`, `-R`, `-S`, `-T`, `-W`, `-b`, `-f`, `-g`, `-i`, `-o`, `-r`, `-t`, `-w`
### Batch 18: System & Environment (stty, install, dircolors, sync, getlimits)
- [ ] Audit `stty` (terminal settings)
    - Upstream: `third_party/coreutils/src/stty.c`
    - Flags: `-a`, `-g`, `-F`
- [ ] Audit `install` (ginstall)
    - Upstream: `third_party/coreutils/src/install.c`
    - Flags: `--backup`, `-b`, `-c`, `-C`, `-d`, `-D`, `--debug`, `-g`, `-m`, `-o`, `-p`, `-s`, `--strip-program`, `-S`, `-t`, `-T`, `-v`, `--preserve-context`, `-Z`
- [ ] Audit `dircolors` (LS_COLORS setup)
    - Upstream: `third_party/coreutils/src/dircolors.c`
    - Flags: `-b`, `-c`, `-p`, `--print-ls-colors`
- [ ] Audit `sync` (sync disks)
    - Upstream: `third_party/coreutils/src/sync.c`
    - Flags: `-d`, `-f`
- [ ] Audit `getlimits` (platform limits)
    - Upstream: `third_party/coreutils/src/getlimits.c`
    - Flags: None specific.
### Batch 19: Security Context (chcon)
- [ ] Audit `chcon` (change context)
    - Upstream: `third_party/coreutils/src/chcon.c`
    - Flags: `-R`, `--dereference`, `-h`, `--no-preserve-root`, `--preserve-root`, `--reference`, `-u`, `-r`, `-t`, `-l`, `-v`, `-H`, `-L`, `-P`
### Batch 21: Bash Builtins (bind, job control, resource limits, etc.)
- [ ] Audit `bind` (readline bindings)
    - Upstream: `third_party/bash/builtins/bind.def`
    - Flags: `-m`, `-l`, `-P`, `-p`, `-S`, `-s`, `-V`, `-v`, `-q`, `-u`, `-r`, `-f`, `-x`, `-X`
- [ ] Audit `break` / `continue` (loop control)
    - Upstream: `third_party/bash/builtins/break.def`
    - Flags: `[n]`
- [ ] Audit `caller` (stack trace)
    - Upstream: `third_party/bash/builtins/caller.def`
    - Flags: `[expr]`
- [ ] Audit `complete` / `compgen` / `compopt` (programmable completion)
    - Upstream: `third_party/bash/builtins/complete.def`
    - Flags: `-abcdefgjksuv`, `-o`, `-A`, `-G`, `-W`, `-F`, `-C`, `-X`, `-P`, `-S`, `-V`, `-p`, `-r`, `-D`, `-E`, `-I`
- [ ] Audit `enable` (builtin control)
    - Upstream: `third_party/bash/builtins/enable.def`
    - Flags: `-a`, `-n`, `-p`, `-s`, `-f`, `-d`
- [ ] Audit `fg` / `bg` (job control)
    - Upstream: `third_party/bash/builtins/fg_bg.def`
    - Flags: `[job_spec]`
- [ ] Audit `jobs` / `disown` (job management)
    - Upstream: `third_party/bash/builtins/jobs.def`
    - Flags: `-l`, `-n`, `-p`, `-r`, `-s`, `-x` (jobs); `-h`, `-a`, `-r` (disown)
- [ ] Audit `let` (arithmetic)
    - Upstream: `third_party/bash/builtins/let.def`
    - Flags: None specific
- [ ] Audit `mapfile` / `readarray` (array input)
    - Upstream: `third_party/bash/builtins/mapfile.def`
    - Flags: `-d`, `-u`, `-n`, `-O`, `-t`, `-C`, `-c`, `-s`
- [ ] Audit `pushd` / `popd` / `dirs` (directory stack)
    - Upstream: `third_party/bash/builtins/pushd.def`
    - Flags: `-n` (pushd/popd); `-c`, `-l`, `-p`, `-v` (dirs)
- [ ] Audit `suspend` (shell suspension)
    - Upstream: `third_party/bash/builtins/suspend.def`
    - Flags: `-f`
- [ ] Audit `ulimit` (resource limits)
    - Upstream: `third_party/bash/builtins/ulimit.def`
    - Flags: `-S`, `-H`, `-a`, `-b`, `-c`, `-d`, `-e`, `-f`, `-i`, `-k`, `-l`, `-m`, `-n`, `-p`, `-q`, `-r`, `-s`, `-t`, `-u`, `-v`, `-x`, `-P`, `-R`, `-T`
### Batch 22: Bash Builtins (State & Navigation)
- [ ] Audit `export` (environment variables)
    - Upstream: `third_party/bash/builtins/setattr.def`
    - Flags: `-f`, `-n`, `-p`
- [ ] Audit `readonly` (constant variables)
    - Upstream: `third_party/bash/builtins/setattr.def`
    - Flags: `-a`, `-A`, `-f`, `-p`
- [ ] Audit `cd` (navigation)
    - Upstream: `third_party/bash/builtins/cd.def`
    - Flags: `-L`, `-P`, `-e`, `-@`
- [ ] Audit `pwd` (working directory)
    - Upstream: `third_party/bash/builtins/cd.def`
    - Flags: `-L`, `-P`
### Batch 23: Coreutils (basenc, checksums, env, etc.)
- [ ] Audit `basenc` (encoding)
    - Upstream: `third_party/coreutils/src/basenc.c`
    - Flags: `--base64`, `--base64url`, `--base58`, `--base32`, `--base32hex`, `--base16`, `--base2msbf`, `--base2lsbf`, `-d`, `--decode`, `-i`, `--ignore-garbage`, `-w`, `--wrap=COLS`, `--z85`
- [ ] Audit `cksum` (checksums)
    - Upstream: `third_party/coreutils/src/cksum.c`
    - Flags: `-r`, `-s`, `--sysv`, `-a`, `--algorithm=TYPE`, `--base64`, `-b`, `--binary`, `-c`, `--check`, `-l`, `--length=BITS`, `--raw`, `--tag`, `--untagged`, `-t`, `--text`, `-z`, `--zero`, `--ignore-missing`, `--quiet`, `--status`, `--strict`, `-w`, `--warn`, `--debug`
- [ ] Audit `comm` (comparison)
    - Upstream: `third_party/coreutils/src/comm.c`
    - Flags: `-1`, `-2`, `-3`, `--check-order`, `--nocheck-order`, `--output-delimiter=STR`, `--total`, `-z`, `--zero-terminated`
- [ ] Audit `csplit` (split by pattern)
    - Upstream: `third_party/coreutils/src/csplit.c`
    - Flags: `-b`, `--suffix-format=FORMAT`, `-f`, `--prefix=PREFIX`, `-k`, `--keep-files`, `--suppress-matched`, `-n`, `--digits=DIGITS`, `-s`, `--quiet`, `--silent`, `-z`, `--elide-empty-files`
- [ ] Audit `env` (environment)
    - Upstream: `third_party/coreutils/src/env.c`
    - Flags: `-a`, `--argv0=ARG`, `-i`, `--ignore-environment`, `-0`, `--null`, `-u`, `--unset=NAME`, `-C`, `--chdir=DIR`, `-S`, `--split-string=S`, `--block-signal[=SIG]`, `--default-signal[=SIG]`, `--ignore-signal[=SIG]`, `--list-signal-handling`, `-v`, `--debug`, `-`
- [ ] Audit `expand` (tabs to spaces)
    - Upstream: `third_party/coreutils/src/expand.c`
    - Flags: `-i`, `--initial`, `-t`, `--tabs=N`
### Batch 24: Coreutils (Text Processing & Math)
- [ ] Audit `expr` (expressions)
    - Upstream: `third_party/coreutils/src/expr.c`
    - Flags: None specific
- [ ] Audit `factor` (prime factors)
    - Upstream: `third_party/coreutils/src/factor.c`
    - Flags: `-h`, `--exponents`
- [ ] Audit `fmt` (paragraph formatting)
    - Upstream: `third_party/coreutils/src/fmt.c`
    - Flags: `-WIDTH`, `-c`, `--crown-margin`, `-p`, `--prefix=STRING`, `-s`, `--split-only`, `-t`, `--tagged-paragraph`, `-u`, `--uniform-spacing`, `-w`, `--width=WIDTH`, `-g`, `--goal=WIDTH`
- [ ] Audit `fold` (line wrapping)
    - Upstream: `third_party/coreutils/src/fold.c`
    - Flags: `-b`, `--bytes`, `-c`, `--characters`, `-s`, `--spaces`, `-w`, `--width=WIDTH`
- [ ] Audit `groups` (group memberships)
    - Upstream: `third_party/coreutils/src/groups.c`
    - Flags: None specific
### Batch 25: Coreutils (System Info & Joining)
- [ ] Audit `hostid` (host id)
    - Upstream: `third_party/coreutils/src/hostid.c`
    - Flags: None specific
- [ ] Audit `hostname` (host name)
    - Upstream: `third_party/coreutils/src/hostname.c`
    - Flags: None specific
- [ ] Audit `id` (user/group id)
    - Upstream: `third_party/coreutils/src/id.c`
    - Flags: `-a`, `-Z`, `--context`, `-g`, `--group`, `-G`, `--groups`, `-n`, `--name`, `-r`, `--real`, `-u`, `--user`, `-z`, `--zero`
- [ ] Audit `join` (join lines)
    - Upstream: `third_party/coreutils/src/join.c`
    - Flags: `-a FILENUM`, `-e STRING`, `-i`, `--ignore-case`, `-j FIELD`, `-o FORMAT`, `-t CHAR`, `-v FILENUM`, `-1 FIELD`, `-2 FIELD`, `--check-order`, `--nocheck-order`, `--header`, `-z`, `--zero-terminated`
- [ ] Audit `logname` (login name)
    - Upstream: `third_party/coreutils/src/logname.c`
    - Flags: None specific
### Batch 26: Coreutils (Process & Number Formatting)
- [ ] Audit `nice` (niceness)
    - Upstream: `third_party/coreutils/src/nice.c`
    - Flags: `-n`, `--adjustment=N`
- [ ] Audit `nl` (number lines)
    - Upstream: `third_party/coreutils/src/nl.c`
    - Flags: `-b STYLE`, `-d CC`, `-f STYLE`, `-h STYLE`, `-i NUMBER`, `-l NUMBER`, `-n FORMAT`, `-p`, `-s STRING`, `-v NUMBER`, `-w NUMBER`
- [ ] Audit `nohup` (hangup immune)
    - Upstream: `third_party/coreutils/src/nohup.c`
    - Flags: None specific
- [ ] Audit `nproc` (processor count)
    - Upstream: `third_party/coreutils/src/nproc.c`
    - Flags: `--all`, `--ignore=N`
- [ ] Audit `numfmt` (reformat numbers)
    - Upstream: `third_party/coreutils/src/numfmt.c`
    - Flags: `--debug`, `-d`, `--delimiter=X`, `--field=FIELDS`, `--format=FORMAT`, `--from=UNIT`, `--from-unit=N`, `--grouping`, `--header[=N]`, `--invalid=MODE`, `--padding=N`, `--round=METHOD`, `--suffix=SUFFIX`, `--unit-separator=SEP`, `--to=UNIT`, `--to-unit=N`, `-z`, `--zero-terminated`
- [ ] Audit `od` (octal dump)
    - Upstream: `third_party/coreutils/src/od.c`
    - Flags: `-A`, `--address-radix=RADIX`, `--endian`, `-j`, `--skip-bytes=BYTES`, `-N`, `--read-bytes=BYTES`, `-S`, `--strings`, `-t`, `--format=TYPE`, `-v`, `--output-duplicates`, `-w`, `--width`, `--traditional`, `-a`, `-b`, `-c`, `-d`, `-f`, `-i`, `-l`, `-o`, `-s`, `-x`
### Batch 27: Coreutils (Text Formatting & Info)
- [ ] Audit `paste` (merge lines)
    - Upstream: `third_party/coreutils/src/paste.c`
    - Flags: `-d`, `--delimiters=LIST`, `-s`, `--serial`, `-z`, `--zero-terminated`
- [ ] Audit `pathchk` (check filenames)
    - Upstream: `third_party/coreutils/src/pathchk.c`
    - Flags: `-p`, `-P`, `--portability`
- [ ] Audit `pinky` (user info)
    - Upstream: `third_party/coreutils/src/pinky.c`
    - Flags: `-l`, `-b`, `-h`, `-p`, `-s`, `-f`, `-w`, `-i`, `-q`, `--lookup`
- [ ] Audit `pr` (format for print)
    - Upstream: `third_party/coreutils/src/pr.c`
    - Flags: `+FIRST_PAGE[:LAST_PAGE]`, `--pages=FIRST_PAGE[:LAST_PAGE]`, `-COLS`, `--columns=COLS`, `-a`, `--across`, `-c`, `--show-control-chars`, `-d`, `--double-space`, `-D`, `--date-format=FORMAT`, `-e`, `--expand-tabs`, `-F`, `-f`, `--form-feed`, `-h`, `--header=HEADER`, `-i`, `--output-tabs`, `-J`, `--join-lines`, `-l`, `--length=PAGE_LENGTH`, `-m`, `--merge`, `-n`, `--number-lines`, `-N`, `--first-line-number=NUMBER`, `-o`, `--indent=MARGIN`, `-r`, `--no-file-warnings`, `-s`, `--separator`, `-S`, `--sep-string`, `-t`, `--omit-header`, `-T`, `--omit-pagination`, `-v`, `--show-nonprinting`, `-w`, `--width=PAGE_WIDTH`, `-W`, `--page-width=PAGE_WIDTH`
- [ ] Audit `printenv` (environment)
    - Upstream: `third_party/coreutils/src/printenv.c`
    - Flags: `-0`, `--null`
- [ ] Audit `ptx` (permuted index)
    - Upstream: `third_party/coreutils/src/ptx.c`
    - Flags: `-A`, `--auto-reference`, `-G`, `--traditional`, `-F`, `--flag-truncation=STRING`, `-M`, `--macro-name=STRING`, `-O`, `--format=roff`, `-R`, `--right-side-refs`, `-S`, `--sentence-regexp=REGEXP`, `-T`, `--format=tex`, `-W`, `--word-regexp=REGEXP`, `-b`, `--break-file=FILE`, `-f`, `--ignore-case`, `-g`, `--gap-size=NUMBER`, `-i`, `--ignore-file=FILE`, `-o`, `--only-file=FILE`, `-r`, `--references`, `-t`, `--typeset-mode`, `-w`, `--width=NUMBER`
