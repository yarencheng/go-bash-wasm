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

