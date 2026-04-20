# Functional Parity Tracking

This document tracks the alignment of the Go Bash Simulator with upstream GNU implementations.

## Overview

Status codes:
- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

---


## Bash Builtins

### `:`

- [x] Upstream: `third_party/bash/builtins/colon.def`
- [x] Basic operation: Implemented in `internal/commands/bash/colon/colon.go`

### `alias`

- [x] Basic management: Implemented in `internal/commands/bash/alias/alias.go`
- [x] Flag `-p`: `third_party/bash/builtins/alias.def:L79` (print)

### `bash`

- [x] Upstream: `third_party/bash/shell.c`, `third_party/bash/version.c`
- [x] Flag `--version`: `third_party/bash/shell.c:L483`, `third_party/bash/version.c:L88`

### `bg`

- [x] Upstream: `third_party/bash/builtins/fg_bg.def`
- [x] Basic job management: Implemented in `internal/commands/bash/bg/bg.go`
- [x] Job specification support: `internal/commands/bash/bg/bg.go`

### `bind`

- [x] Upstream: `third_party/bash/builtins/bind.def`
- [x] Keybinding management: Implemented in `internal/commands/bash/bind/bind.go`
- [x] Flag `-l`: `internal/commands/bash/bind/bind.go` (list)
- [x] Flag `-v`: `internal/commands/bash/bind/bind.go` (list functions)
- [x] Flag `-p`: `internal/commands/bash/bind/bind.go` (print status)
- [x] Flag `-V`: `internal/commands/bash/bind/bind.go` (list variables)
- [x] Flag `-P`: `internal/commands/bash/bind/bind.go` (print functions)
- [x] Flag `-s`: `internal/commands/bash/bind/bind.go` (list macros)
- [x] Flag `-S`: `internal/commands/bash/bind/bind.go` (print macros)
- [x] Flag `-X`: `internal/commands/bash/bind/bind.go` (list keyseq bindings)
- [x] Flag `-f=FILE`: `internal/commands/bash/bind/bind.go` (read from file)
- [x] Flag `-q=FUNC`: `internal/commands/bash/bind/bind.go` (query keys for func)
- [x] Flag `-u=FUNC`: `internal/commands/bash/bind/bind.go` (unbind func)
- [x] Flag `-m=KEYMAP`: `internal/commands/bash/bind/bind.go` (keymap)
- [x] Flag `-r=KEYSEQ`: `internal/commands/bash/bind/bind.go` (remove seq)
- [x] Flag `-x=KEYSEQ:SHELLCMD`: `internal/commands/bash/bind/bind.go` (exec cmd)

### `break`

- [x] Upstream: `third_party/bash/builtins/break.def`
- [x] Basic operation: Implemented in `internal/commands/bash/break/break.go`

### `builtin`

- [x] Upstream: `third_party/bash/builtins/builtin.def`
- [x] Basic execution: Implemented in `internal/commands/bash/builtin/builtin.go`

### `caller`

- [x] Upstream: `third_party/bash/builtins/caller.def`
- [x] Basic operation: Implemented in `internal/commands/bash/caller/caller.go`

### `cd`

- [x] Upstream: `third_party/bash/builtins/cd.def`
- [x] Basic change directory: Implemented in `internal/commands/bash/cd/cd.go`
- [x] CDPATH support: `internal/commands/bash/cd/cd.go`
- [x] Flag `-e`: `internal/commands/bash/cd/cd.go` (exit status if -P cannot be satisfied)
- [x] Flag `-L`: `internal/commands/bash/cd/cd.go`
- [x] Flag `-P`: `internal/commands/bash/cd/cd.go`

### `command`

- [x] Upstream: `third_party/bash/builtins/command.def`
- [x] Execution override: Implemented in `internal/commands/bash/command/command.go`
- [x] Flag `-p`: `internal/commands/bash/command/command.go`
- [x] Flag `-v`: `internal/commands/bash/command/command.go`
- [x] Flag `-V`: `internal/commands/bash/command/command.go`

### `compgen`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Inherits all `complete` flags: Implemented in `internal/commands/bash/compgen/compgen.go`

### `complete`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Completion management: Implemented in `internal/commands/bash/complete/complete.go`
- [x] Flag `-a`: `internal/commands/bash/complete/complete.go` (alias)
- [x] Flag `-b`: `internal/commands/bash/complete/complete.go` (builtin)
- [x] Flag `-c`: `internal/commands/bash/complete/complete.go` (command)
- [x] Flag `-d`: `internal/commands/bash/complete/complete.go` (directory)
- [x] Flag `-e`: `internal/commands/bash/complete/complete.go` (export)
- [x] Flag `-f`: `internal/commands/bash/complete/complete.go` (file)
- [x] Flag `-g`: `internal/commands/bash/complete/complete.go` (group)
- [x] Flag `-j`: `internal/commands/bash/complete/complete.go` (job)
- [x] Flag `-k`: `internal/commands/bash/complete/complete.go` (keyword)
- [x] Flag `-p`: `internal/commands/bash/complete/complete.go` (print)
- [x] Flag `-r`: `internal/commands/bash/complete/complete.go` (remove)
- [x] Flag `-s`: `internal/commands/bash/complete/complete.go` (service)
- [x] Flag `-u`: `internal/commands/bash/complete/complete.go` (user)
- [x] Flag `-v`: `internal/commands/bash/complete/complete.go` (variable)
- [x] Flag `-o=OPT`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-A=ACTION`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-G=GLOB`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-W=WORDLIST`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-P=PREFIX`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-S=SUFFIX`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-X=FILTER`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-F=FUNC`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-C=CMD`: `internal/commands/bash/complete/complete.go`
- [x] Flag `-E`: `internal/commands/bash/complete/complete.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-I`: `internal/commands/bash/complete/complete.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-D`: `internal/commands/bash/complete/complete.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))

### `compopt`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Flag `-o`, `--options`: `internal/commands/bash/compopt/compopt.go`
- [x] Flag `-D`, `--default`: `internal/commands/bash/compopt/compopt.go`
- [x] Flag `-E`, `--empty`: `internal/commands/bash/compopt/compopt.go`

### `continue`

- [x] Upstream: `third_party/bash/builtins/break.def`
- [x] Basic operation: Implemented in `internal/commands/bash/continue/continue.go`

### `declare`

- [x] Upstream: `third_party/bash/builtins/declare.def`
- [x] Basic operation: Implemented in `internal/commands/bash/declare/declare.go`
- [x] Flag `-p` (print individual names): `internal/commands/bash/declare/declare.go`
- [x] Flag `-l` (lowercase): `internal/commands/bash/declare/declare.go`
- [x] Flag `-u` (uppercase): `internal/commands/bash/declare/declare.go`
- [x] Flag `-a` (indexed array): `internal/commands/bash/declare/declare.go`
- [x] Flag `-A` (associative array): `internal/commands/bash/declare/declare.go`
- [x] Flag `-i` (integer): `internal/commands/bash/declare/declare.go`
- [x] Flag `-r` (readonly): `internal/commands/bash/declare/declare.go`
- [x] Flag `-x` (export): `internal/commands/bash/declare/declare.go`
- [x] Flag `-f` (functions): `internal/commands/bash/declare/declare.go`
- [x] Flag `-F` (function names): `internal/commands/bash/declare/declare.go`
- [x] Flag `-g` (global): `internal/commands/bash/declare/declare.go`
- [x] Flag `-I` (inherit): `internal/commands/bash/declare/declare.go`
- [x] Flag `-t` (trace): `internal/commands/bash/declare/declare.go`
- [x] Flag `-n` (nameref): `internal/commands/bash/declare/declare.go`
- [x] Synonym `typeset`: `third_party/bash/builtins/declare.def:L66` (Registered in `internal/app/app.go`)

### `dirs`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic listing: Implemented in `internal/commands/bash/dirs/dirs.go`
- [x] Flag `-c`: `internal/commands/bash/dirs/dirs.go` (clear stack)
- [x] Flag `-l`: `internal/commands/bash/pushd/pushd.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/bash/dirs/dirs.go` (print with one line per entry)
- [x] Flag `-v`: `internal/commands/bash/dirs/dirs.go` (verbose)

### `disown`

- [x] Upstream: `third_party/bash/builtins/jobs.def`
- [x] Flag `-a`: `internal/commands/bash/disown/disown.go` (all jobs)
- [x] Flag `-h`: `internal/commands/bash/disown/disown.go` (mark to not receive SIGHUP)
- [x] Flag `-r`: `internal/commands/bash/disown/disown.go` (running jobs only)

### `echo`

- [x] Upstream: `third_party/bash/builtins/echo.def`
- [x] Basic output: Implemented in `internal/commands/bash/echo/echo.go`
- [x] Flag `-n`: `internal/commands/bash/echo/echo.go` (no newline)
- [x] Flag `-e`: `internal/commands/bash/echo/echo.go` (interpret escapes)
- [x] Flag `-E`: `internal/commands/bash/echo/echo.go` (disable escapes)

### `enable`

- [x] Upstream: `third_party/bash/builtins/enable.def`
- [x] Basic management: Implemented in `internal/commands/bash/enable/enable.go`
- [x] Flag `-a`: `third_party/bash/builtins/enable.def:L157` (display all)
- [-] Flag `-d`: `third_party/bash/builtins/enable.def:L160` (delete loaded) - Dynamic loading not available
- [x] Flag `-n`: `third_party/bash/builtins/enable.def:L163` (disable)
- [x] Flag `-p`: `third_party/bash/builtins/enable.def:L166` (print status)
- [x] Flag `-s`: `third_party/bash/builtins/enable.def:L169` (POSIX special only)
- [-] Flag `-f filename`: `third_party/bash/builtins/enable.def:L172` (load from dynamic file) - Dynamic loading not available

### `eval`

- [x] Upstream: `third_party/bash/builtins/eval.def`
- [x] Basic execution: Implemented in `internal/commands/bash/eval/eval.go`

### `exec`

- [x] Upstream: `third_party/bash/builtins/exec.def`
- [x] Basic execution: Implemented in `internal/commands/bash/exec/exec.go`
- [x] Flag `-l`: `internal/commands/bash/exec/exec.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-a name`: `internal/commands/bash/exec/exec.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-c`: `internal/commands/bash/exec/exec.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))

### `exit`

- [x] Upstream: `third_party/bash/builtins/exit.def`
- [x] Basic exit: Implemented in `internal/commands/bash/exit/exit.go`
- [x] Exit status parameter: `internal/commands/bash/exit/exit.go`

### `export`

- [x] Upstream: `third_party/bash/builtins/setattr.def`
- [x] Basic operation: Implemented in `internal/commands/bash/export/export.go`
- [x] Flag `-f`: `internal/commands/bash/export/export.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-n`: `internal/commands/bash/export/export.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/bash/export/export.go`

### `true`
 
 - [x] Upstream: `third_party/bash/builtins/colon.def:L35`, `third_party/coreutils/src/true.c`
 - [x] Basic operation: Implemented in `internal/commands/bash/boolcmd/bool.go`
 
 ### `false`

- [x] Upstream: `third_party/bash/builtins/colon.def:L44`, `third_party/coreutils/src/false.c`
- [x] Basic operation: Implemented in `internal/commands/bash/boolcmd/bool.go`

### `fc`

- [x] Upstream: `third_party/bash/builtins/fc.def`
- [x] Basic editing/re-execution: Implemented in `internal/commands/bash/fc/fc.go`
- [x] Flag `-e ENAME`: `internal/commands/bash/fc/fc.go`
- [x] Flag `-l`: `internal/commands/bash/fc/fc.go`
- [x] Flag `-n`: `internal/commands/bash/fc/fc.go`
- [x] Flag `-r`: `internal/commands/bash/fc/fc.go`
- [x] Flag `-s`: `internal/commands/bash/fc/fc.go`

### `fg`

- [x] Upstream: `third_party/bash/builtins/fg_bg.def`
- [x] Basic operation: Implemented in `internal/commands/bash/fg/fg.go`

### `getopts`

- [x] Upstream: `third_party/bash/builtins/getopts.def`
- [x] Basic parsing: Implemented in `internal/commands/bash/getopts/getopts.go`
- [x] Variable assignment (OPTARG, OPTIND): `internal/commands/bash/getopts/getopts.go`
- [x] Silent error reporting: `internal/commands/bash/getopts/getopts.go`
- [x] Shell positional parameters fallback: `internal/commands/bash/getopts/getopts.go`

### `hash`

- [x] Upstream: `third_party/bash/builtins/hash.def`
- [x] Command hashing: Implemented in `internal/commands/bash/hash/hash.go`
- [x] Flag `-r`: `internal/commands/bash/hash/hash.go`
- [x] Flag `-d`: `internal/commands/bash/hash/hash.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/bash/hash/hash.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`: `internal/commands/bash/hash/hash.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/bash/hash/hash.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))

### `help`

- [x] Upstream: `third_party/bash/builtins/help.def`
- [x] Help system: Implemented in `internal/commands/bash/help/help.go`
- [x] Flag `-d`: `internal/commands/bash/help/help.go` (short description)
- [x] Flag `-m`: `internal/commands/bash/help/help.go` (man-page format)
- [x] Flag `-s`: `internal/commands/bash/help/help.go` (syntax only)

### `history`

- [x] Upstream: `third_party/bash/builtins/history.def`
- [x] History management: Implemented in `internal/commands/bash/history/history.go`
- [x] Flag `-d offset`: `internal/commands/bash/history/history.go` (delete entry)
- [x] Flag `-a`: `internal/commands/bash/history/history.go` (append)
- [x] Flag `-c`: `internal/commands/bash/history/history.go` (clear)
- [x] Flag `-n`: `internal/commands/bash/history/history.go` (read non-recorded)
- [x] Flag `-p`: `internal/commands/bash/history/history.go` (print/expand)
- [x] Flag `-r`: `internal/commands/bash/history/history.go` (read file)
- [x] Flag `-s`: `internal/commands/bash/history/history.go` (store/append)
- [x] Flag `-w`: `internal/commands/bash/history/history.go` (write file)

### `jobs`

- [x] Upstream: `third_party/bash/builtins/jobs.def`
- [x] Basic listing: Implemented in `internal/commands/bash/jobs/jobs.go`
- [x] Flag `-l`: `internal/commands/bash/jobs/jobs.go` (long format)
- [x] Flag `-n`: `internal/commands/bash/jobs/jobs.go` (only jobs that changed)
- [x] Flag `-p`: `internal/commands/bash/jobs/jobs.go` (only PIDs)
- [x] Flag `-r`: `internal/commands/bash/jobs/jobs.go` (running only)
- [x] Flag `-s`: `internal/commands/bash/jobs/jobs.go` (stopped only)
- [x] Flag `-x command`: `internal/commands/bash/jobs/jobs.go` (execute command)

### `kill`

- [x] Upstream: `third_party/bash/builtins/kill.def`
- [x] Basic signaling: Implemented in `internal/commands/bash/kill/kill.go`
- [x] Flag `-l`: `internal/commands/bash/kill/kill.go:L68`
- [x] Flag `-n num`: `internal/commands/bash/kill/kill.go:L50`
- [x] Flag `-s SIGNAL`: `internal/commands/bash/kill/kill.go:L49`
- [x] Flag `-SIGNAL` (e.g. -9, -TERM): `internal/commands/bash/kill/kill.go`

### `let`

- [x] Upstream: `third_party/bash/builtins/let.def`
- [x] Basic arithmetic: Implemented in `internal/commands/bash/letcmd/let.go`

### `local`

- [x] Upstream: `third_party/bash/builtins/declare.def`
- [x] Basic operation: Implemented in `internal/commands/bash/local/local.go`
- [x] Inherits generic `declare` behavior

### `logout`

- [x] Upstream: `third_party/bash/builtins/exit.def`
- [x] Basic operation: Implemented in `internal/commands/bash/logout/logout.go`

### `mapfile`

- [x] Upstream: `third_party/bash/builtins/mapfile.def`
- [x] Array population: Implemented in `internal/commands/bash/mapfile/mapfile.go`
- [x] Flag `-d`: `internal/commands/bash/mapfile/mapfile.go`
- [x] Flag `-t`: `internal/commands/bash/mapfile/mapfile.go` (trim/strip newline)
- [x] Flag `-n`: `internal/commands/bash/mapfile/mapfile.go` (count)
- [x] Flag `-O`: `internal/commands/bash/mapfile/mapfile.go` (origin)
- [x] Flag `-u`: `internal/commands/bash/mapfile/mapfile.go` (fd)
- [x] Flag `-C`: `internal/commands/bash/mapfile/mapfile.go` (callback)
- [x] Flag `-c`: `internal/commands/bash/mapfile/mapfile.go` (quantum)
- [x] Flag `-s`: `internal/commands/bash/mapfile/mapfile.go`
- [x] Aliases: `readarray` (handled via command registration)

### `popd`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic popping: Implemented in `internal/commands/bash/popd/popd.go`
- [x] Flag `-n`: `internal/commands/bash/popd/popd.go`

### `printf`

- [x] Upstream: `third_party/bash/builtins/printf.def`
- [x] Basic formatting: Implemented in `internal/commands/coreutils/printf/printf.go`
- [x] Flag `-v` (assign to variable): `internal/commands/coreutils/printf/printf.go`
- [x] Format specifier `%b` (escapes): `internal/commands/coreutils/printf/printf.go`
- [x] Format specifier `%q` (quoted): `internal/commands/coreutils/printf/printf.go`
- [x] Format specifier `%Q` (quoted with precision): `internal/commands/coreutils/printf/printf.go`
- [x] Format specifier `%T` (date/time): `internal/commands/coreutils/printf/printf.go`
- [x] Width/precision `*` support: `internal/commands/coreutils/printf/printf.go`
- [x] Standard C format specifiers (`csndiouxXeEfFgGaA`): `internal/commands/coreutils/printf/printf.go`
- [x] Format reusing: `internal/commands/coreutils/printf/printf.go`

### `pushd`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic pushing: Implemented in `internal/commands/bash/pushd/pushd.go`
- [x] Flag `-n`: `internal/commands/bash/pushd/pushd.go`

### `pwd`

- [x] Upstream: `third_party/bash/builtins/cd.def`
- [x] Basic path reporting: Implemented in `internal/commands/bash/pwd/pwd.go`
- [-] Flag `--help`: Handled by the shell's global help dispatcher. (See [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#pwd))
- [x] Flag `-L`: `internal/commands/bash/pwd/pwd.go`
- [x] Flag `-P`: `internal/commands/bash/pwd/pwd.go`

### `read`

- [x] Upstream: `third_party/bash/builtins/read.def`
- [x] Basic input: Implemented in `internal/commands/bash/read/read.go`
- [x] Basic operation: Implemented in `internal/commands/bash/read/read.go`
- [x] Flag `-p PROMPT`: `internal/commands/bash/read/read.go`
- [x] Flag `-r` (raw): `internal/commands/bash/read/read.go`
- [x] Flag `-d DELIM`: `internal/commands/bash/read/read.go`
- [x] Flag `-n NCHARS`: `internal/commands/bash/read/read.go`
- [x] Flag `-N NCHARS`: `internal/commands/bash/read/read.go`
- [x] Flag `-a ARRAY`: `internal/commands/bash/read/read.go`
- [x] Flag `-s`: `internal/commands/bash/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#read))
- [x] Flag `-t TIMEOUT`: `internal/commands/bash/read/read.go`
- [x] Flag `-u FD`: `internal/commands/bash/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#read))
- [x] Flag `-e`: `internal/commands/bash/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#read))
- [x] Flag `-i TEXT`: `internal/commands/bash/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#read))

### `readonly`

- [x] Upstream: `third_party/bash/builtins/setattr.def`
- [x] Attribute management: Implemented in `internal/commands/bash/readonly/readonly.go`
- [x] Flag `-a`: `internal/commands/bash/readonly/readonly.go` (indexed array)
- [x] Flag `-A`: `internal/commands/bash/readonly/readonly.go` (associative array)
- [x] Flag `-p`: `internal/commands/bash/readonly/readonly.go` (print)
- [x] Flag `-f`: `internal/commands/bash/readonly/readonly.go` (functions)

### `return`

- [x] Upstream: `third_party/bash/builtins/return.def`
- [x] Basic return: Implemented in `internal/commands/bash/returncmd/return.go`
- [x] Exit status parameter: `internal/commands/bash/returncmd/return.go`

### `select`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Selection from list: Implemented in `internal/shell/shell.go`
- [x] Basic operation: Implemented in `internal/shell/shell.go`

### `set`

- [x] Upstream: `third_party/bash/builtins/set.def`
- [x] Option management (-e, -u, -x, -o): Implemented in `internal/commands/bash/set/set.go`
- [x] Positional parameters: Stub in `internal/commands/bash/set/set.go` (See [functional_gap.md](./functional_gap.md#set))
- [x] Flag `-a`: `internal/commands/bash/set/set.go` (allexport)
- [x] Flag `-b`: `internal/commands/bash/set/set.go` (notify)
- [x] Flag `-e`: `internal/commands/bash/set/set.go` (errexit)
- [x] Flag `-f`: `internal/commands/bash/set/set.go` (noglob)
- [x] Flag `-h`: `internal/commands/bash/set/set.go` (hashall)
- [x] Flag `-k`: `internal/commands/bash/set/set.go` (keyword)
- [x] Flag `-m`: `internal/commands/bash/set/set.go` (monitor)
- [x] Flag `-n`: `internal/commands/bash/set/set.go` (noexec)
- [x] Flag `-o`: `internal/commands/bash/set/set.go` (option-name)
- [x] Flag `-p`: `internal/commands/bash/set/set.go` (privileged)
- [x] Flag `-t`: `internal/commands/bash/set/set.go` (exit after one command)
- [x] Flag `-u`: `internal/commands/bash/set/set.go` (nounset)
- [x] Flag `-v`: `internal/commands/bash/set/set.go` (verbose)
- [x] Flag `-x`: `internal/commands/bash/set/set.go` (xtrace)
- [x] Flag `-B`: `internal/commands/bash/set/set.go` (braceexpand)
- [x] Flag `-C`: `internal/commands/bash/set/set.go` (noclobber)
- [x] Flag `-E`: `internal/commands/bash/set/set.go` (errtrace)
- [x] Flag `-H`: `internal/commands/bash/set/set.go` (histexpand)
- [x] Flag `-P`: `internal/commands/bash/set/set.go` (physical)
- [x] Flag `-T`: `internal/commands/bash/set/set.go` (functrace)

### `shift`

- [x] Upstream: `third_party/bash/builtins/shift.def`
- [x] Basic shift: Implemented in `internal/commands/bash/shift/shift.go`
- [x] Shifting n parameters: `internal/commands/bash/shift/shift.go`

### `shopt`

- [x] Upstream: `third_party/bash/builtins/shopt.def`
- [x] Basic option management: Implemented in `internal/commands/bash/shopt/shopt.go`
- [x] Flag `-o`: `internal/commands/bash/shopt/shopt.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `third_party/bash/builtins/shopt.def:L77` (print status)
- [x] Flag `-q`: `third_party/bash/builtins/shopt.def:L71` (quiet)
- [x] Flag `-s`: `third_party/bash/builtins/shopt.def:L62` (enable)
- [x] Flag `-u`: `third_party/bash/builtins/shopt.def:L65` (disable)

### `source`

- [x] Upstream: `third_party/bash/builtins/source.def`
- [x] Basic sourcing: Implemented in `internal/commands/bash/source/source.go`
- [x] Flag `-p PATH`: `internal/commands/bash/source/source.go`
- [x] Aliases: `.`

### `suspend`

- [x] Upstream: `third_party/bash/builtins/suspend.def`
- [x] Basic operation: `internal/commands/bash/suspend/suspend.go` (Unsupported; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#suspend))
- [x] Flag `-f`: `internal/commands/bash/suspend/suspend.go`

### `test`

- [x] Upstream: `third_party/bash/builtins/test.def`
- [x] Unary operators (`-e`, `-f`, `-d`, `-z`, `-n`): Implemented in `internal/commands/bash/test/test.go`
- [x] Binary operators (`=`, `==`, `!=`, `-eq`, `-ne`, `-lt`, `-le`, `-gt`, `-ge`): Implemented in `internal/commands/bash/test/test.go`
- [x] Logical operators (`!`, `-a`, `-o`): Implemented in `internal/commands/bash/test/test.go`
- [x] Synonym `[`: Implemented via command registration
- [x] Unary `-r` (readable): `internal/commands/bash/test/test.go`
- [x] Unary `-w` (writable): `internal/commands/bash/test/test.go`
- [x] Unary `-x` (executable): `internal/commands/bash/test/test.go`
- [x] Unary `-O` (owned by user): `internal/commands/bash/test/test.go`
- [x] Unary `-G` (owned by group): `internal/commands/bash/test/test.go`
- [x] Unary `-N` (newer than atime): `internal/commands/bash/test/test.go`
- [x] Unary `-h`, `-L` (symlink): `internal/commands/bash/test/test.go`
- [x] Unary `-S` (socket): `internal/commands/bash/test/test.go`
- [x] Unary `-p` (pipe): `internal/commands/bash/test/test.go`
- [x] Unary `-b` (block device): `internal/commands/bash/test/test.go`
- [x] Unary `-c` (character device): `internal/commands/bash/test/test.go`
- [x] Unary `-u` (setuid): `internal/commands/bash/test/test.go`
- [x] Unary `-g` (setgid): `internal/commands/bash/test/test.go`
- [x] Unary `-k` (sticky): `internal/commands/bash/test/test.go`
- [x] Unary `-t` (tty): `internal/commands/bash/test/test.go`
- [x] Unary `-o` (option): `internal/commands/bash/test/test.go`
- [x] Unary `-v` (variable): `internal/commands/bash/test/test.go`
- [x] Unary `-R` (nameref): `internal/commands/bash/test/test.go`
- [x] Binary `-nt` (newer than): `internal/commands/bash/test/test.go`
- [x] Binary `-ot` (older than): `internal/commands/bash/test/test.go`
- [x] Binary `-ef` (equal file): `internal/commands/bash/test/test.go`
- [x] Binary `<`, `>` (string comparison): `internal/commands/bash/test/test.go`

### `time`

- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Basic operation: Implemented in `internal/commands/bash/time/time.go`

### `times`

- [x] Upstream: `third_party/bash/builtins/times.def`
- [x] Basic output: Implemented in `internal/commands/bash/times/times.go` (Simulation; see [functional_gap.md](./functional_gap.md#times))

### `trap`

- [x] Upstream: `third_party/bash/builtins/trap.def`
- [x] Basic trapping: Implemented in `internal/commands/bash/trap/trap.go`
- [x] Flag `-P`: `internal/commands/bash/trap/trap.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/bash/trap/trap.go`
- [x] Flag `-p`: `internal/commands/bash/trap/trap.go`

### `type`

- [x] Upstream: `third_party/bash/builtins/type.def`
- [x] Command identification: Implemented in `internal/commands/bash/type/type.go`
- [x] Flag `-a`: `internal/commands/bash/type/type.go` (all occurrences)
- [x] Flag `-p`: `internal/commands/bash/type/type.go` (path only)
- [x] Flag `-t`: `internal/commands/bash/type/type.go` (type only)
- [x] Flag `-f`: `internal/commands/bash/type/type.go` (skip functions)
- [x] Flag `-P`: `internal/commands/bash/type/type.go` (force path search)

### `ulimit`

- [x] Upstream: `third_party/bash/builtins/ulimit.def`
- [x] Resource management: Implemented in `internal/commands/bash/ulimit/ulimit.go` (Simulation)
- [x] Flag `-a`: `third_party/bash/builtins/ulimit.def:L35` (all)
- [x] Flag `-c`: `third_party/bash/builtins/ulimit.def:L37` (core)
- [x] Flag `-d`: `third_party/bash/builtins/ulimit.def:L38` (data)
- [x] Flag `-e`: `third_party/bash/builtins/ulimit.def:L39` (priority)
- [x] Flag `-f`: `third_party/bash/builtins/ulimit.def:L40` (file size)
- [x] Flag `-n`: `third_party/bash/builtins/ulimit.def:L45` (opened files)
- [x] Flag `-u`: `third_party/bash/builtins/ulimit.def:L51` (user processes)
- [x] Flag `-S`: `internal/commands/bash/ulimit/ulimit.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))
- [x] Flag `-H`: `internal/commands/bash/ulimit/ulimit.go` (Ignored; see [functional_gap.md](./functional_gap.md#commonly-ignored-flags))

### `umask`

- [x] Upstream: `third_party/bash/builtins/umask.def`
- [x] Basic mask management: Implemented in `internal/commands/bash/umask/umask.go`
- [x] Flag `-S`: `internal/commands/bash/umask/umask.go`
- [x] Flag `-p`: `internal/commands/bash/umask/umask.go`
- [x] Symbolic mode setting: `internal/commands/bash/umask/umask.go`

### `unalias`

- [x] Upstream: `third_party/bash/builtins/alias.def`
- [x] Remove aliases: Implemented in `internal/commands/bash/unalias/unalias.go`
- [x] Flag `-a`: `internal/commands/bash/unalias/unalias.go` (remove all)

### `unset`

- [x] Upstream: `third_party/bash/builtins/set.def`
- [x] Attribute management: Implemented in `internal/commands/bash/unset/unset.go`
- [x] Flag `-f`: `internal/commands/bash/unset/unset.go` (functions)
- [x] Flag `-v`: `internal/commands/bash/unset/unset.go` (variables)
- [x] Flag `-n`: `internal/commands/bash/unset/unset.go` (Ignored; see [functional_gap.md](./functional_gap.md#unset))

### `wait`

- [x] Upstream: `third_party/bash/builtins/wait.def`
- [x] Basic waiting: Implemented in `internal/commands/bash/wait/wait.go`
- [x] Optional: jobspec or process ID: `internal/commands/bash/wait/wait.go`
- [x] Flag `-f`: `internal/commands/bash/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#wait))
- [x] Flag `-n`: `internal/commands/bash/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/bash/functional_gap.md#wait))
- [x] Flag `-p var`: `internal/commands/bash/wait/wait.go`


## Shell Keywords & Grammar

### `!`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pipeline negation: Implemented in `internal/shell/shell.go`

### `[[`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Conditional expressions: Implemented in `internal/shell/shell.go`
- [x] Pattern matching (`==`, `!=`): Implemented in `internal/shell/shell.go`
- [x] Regex matching (`=~`): Implemented in `internal/shell/shell.go`
- [x] Arithmetic comparisons (`-eq`, `-lt`, etc.): Implemented in `internal/shell/shell.go`
- [x] Aliases: `]]`

### `((`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Arithmetic evaluation: Implemented in `internal/shell/shell.go`
- [x] Aliases: `))`

### `{`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Command grouping: Implemented in `internal/shell/shell.go`
- [x] Aliases: `}`

### `case`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pattern-based branching: Implemented in `internal/shell/shell.go`

### `coproc`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [-] Asynchronous coprocesses: N/A for simulator (See [functional_gap.md](./functional_gap.md#coproc))

### `for`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] List-based iteration: Implemented in `internal/shell/shell.go`
- [x] C-style arithmetic iteration (`for ((`): Implemented in `internal/shell/shell.go`

### `function`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Shell function definition: Implemented in `internal/shell/shell.go`

### `if`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Conditional branching (if/then/elif/else/fi): Implemented in `internal/shell/shell.go`

### `until`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Negative condition looping: Implemented in `internal/shell/shell.go`

### `while`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Positive condition looping: Implemented in `internal/shell/shell.go`
- [x] Sequential list `;`: Implemented in `internal/shell/shell.go`

## Shell Variables

### `BASH_VERSION`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Version information string: Implemented in `internal/app/app.go`

### `CDPATH`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Search path for `cd` command: Implemented in `internal/commands/bash/cd/cd.go`

### `GLOBIGNORE`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pattern-based pathname expansion ignore: Implemented in `internal/shell/shell.go`

### `HISTFILE`, `HISTFILESIZE`, `HISTSIZE`, `HISTIGNORE`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] History management persistence: Initialized in `internal/app/app.go`

### `HOME`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Logical login directory: Initialized in `internal/app/app.go`

### `HOSTNAME`, `HOSTTYPE`, `MACHTYPE`, `OSTYPE`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] System identity metadata: Initialized in `internal/app/app.go`

### `IGNOREEOF`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] EOF handling for interactive shells: Implemented in `internal/shell/shell.go`

### `MAILCHECK`, `MAILPATH`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [-] Mail notification settings: N/A for simulator (See [functional_gap.md](./functional_gap.md#mail-notification))

### `PATH`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Command search path: Initialized in `internal/app/app.go`

### `PROMPT_COMMAND`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pre-prompt execution hook: Implemented in `internal/shell/shell.go`

### `PS1`, `PS2`, `PS3`, `PS4`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Interactive prompt formatting: Initialized in `internal/app/app.go`

### `PWD`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Logical current directory tracking: Implemented in `internal/commands/bash/cd/cd.go`

### `SHELLOPTS`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] List of enabled shell options: Implemented in `internal/shell/shell.go`

### `TERM`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Terminal environment identification: Initialized in `internal/app/app.go`

### `TIMEFORMAT`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Output format for `time` reserved word: Implemented in `internal/shell/shell.go`

## Interactive Shell Features

- [x] Interactive history navigation (Up/Down arrow keys): Implemented in `internal/shell/input_wasm.go`
- [x] Command line editing (Backspace, Ctrl+L, etc.): Implemented in `internal/shell/input_wasm.go`
- [x] Tab completion: Implemented in `internal/shell/input_wasm.go`

## Shell Expansions

### Parameter Expansion
- [x] Basic expansion `${var}`: Implemented in `internal/shell/shell.go`
- [x] Substring expansion `${var:offset:length}`: Implemented in `internal/shell/shell.go`
- [x] Prefix removal `${var#pattern}`, `${var##pattern}`: Implemented in `internal/shell/shell.go`
- [x] Suffix removal `${var%pattern}`, `${var%%pattern}`: Implemented in `internal/shell/shell.go`
- [x] Substring replacement `${var/pattern/string}`: Implemented in `internal/shell/shell.go`
- [x] Case modification `${var^}`, `${var^^}`, `${var,}`, `${var,,}`: Implemented in `internal/shell/shell.go`
- [x] Default values `${var:-default}`, `${var:=default}`: Implemented in `internal/shell/shell.go`
- [x] Alternative/Error values `${var:?error}`, `${var:+alternative}`: Implemented in `internal/shell/shell.go`
- [x] Length expansion `${#var}`: Implemented in `internal/shell/shell.go`
- [x] Dynamic variables (RANDOM, SECONDS, etc.): Implemented in `internal/shell/shell.go`
- [x] IFS-based splitting in `read`: Implemented in `internal/commands/bash/read/read.go`

### Command Substitution
- [x] Basic substitution $(command), `command`: Implemented in `internal/shell/shell.go`

### Arithmetic Expansion
- [x] Basic expansion $((expression)): Implemented in `internal/shell/shell.go`

### Brace Expansion
- [x] basic expansion {a,b,c}: Implemented in `internal/shell/shell.go`

### Tilde Expansion
- [x] basic expansion ~, ~user: Implemented in `internal/shell/shell.go`

## Redirections

### Standard Redirections
- [x] Input redirection `[n]<word`: Implemented in `internal/shell/shell.go`
- [x] Output redirection `[n]>word`: Implemented in `internal/shell/shell.go`
- [x] Append redirection `[n]>>word`: Implemented in `internal/shell/shell.go`
- [x] Force output `[n]>|word`: Implemented in `internal/shell/shell.go`
- [x] Combined stderr/stdout `&>word`, `&>>word`: Implemented in `internal/shell/shell.go`
- [x] Stderr redirection `2>word`: Implemented in `internal/shell/shell.go`

### File Descriptor Manipulation
- [x] Duplicating input `[n]<&word`: Implemented in `internal/shell/shell.go`
- [x] Duplicating output `[n]>&word`: Implemented in `internal/shell/shell.go`
- [x] Moving input `[n]<&digit-`: Implemented in `internal/shell/shell.go`
- [x] Moving output `[n]>&digit-`: Implemented in `internal/shell/shell.go`

### Advanced Redirections
- [x] Here-Documents `[n]<<[-]word`: Implemented in `internal/shell/shell.go`
- [x] Here-Strings `[n]<<<word`: Implemented in `internal/shell/shell.go`
- [x] Process Substitution `<(list)`, `>(list)`: Implemented in `internal/shell/shell.go`

## Globbing Patterns

### Standard Wildcards
- [x] Match any string `*`: Implemented in `internal/shell/shell.go`
- [x] Match any character `?`: Implemented in `internal/shell/shell.go`

### Character Classes
- [x] Match set of characters `[...]`: Implemented in `internal/shell/shell.go`
- [x] Negative match set `[!...]`, `[^...]`: Implemented in `internal/shell/shell.go`

### Extended Globbing (extglob)
- [x] Option `?(list)` (zero or one): Approximated via regex
- [x] Option `*(list)` (zero or more): Approximated via regex
- [x] Option `+(list)` (one or more): Approximated via regex
- [x] Option `@(list)` (exactly one): Approximated via regex
- [x] Option `!(list)` (anything but): Approximated via regex

## Execution Flow

### Pipelines
- [x] Basic pipe `|`: Implemented in `internal/shell/shell.go`
- [x] Combined stderr pipe `|&`: Implemented in `internal/shell/shell.go`

### Compound Commands & Lists
- [x] Sequential list `;`: Implemented in `internal/shell/shell.go`
- [x] Background execution `&`: Implemented in `internal/shell/shell.go`
- [x] Logical AND `&&`: Implemented in `internal/shell/shell.go`
- [x] Logical OR `||`: Implemented in `internal/shell/shell.go`
- [x] Subshell execution `( list )`: Implemented in `internal/shell/shell.go`

## Signal & Trap Handling

### Core Signal Handling
- [x] Trap initialization: Implemented in `internal/commands/bash/trap/trap.go`
- [x] Signal decoding (names/numbers): Implemented in `internal/commands/bash/trap/trap.go`
- [x] Pending trap execution: Implemented

### Subshell & Inheritance
- [x] Signal inheritance rules: Implemented
- [x] Trap reset in subshells: Implemented in `internal/shell/shell.go`

## Advanced Shell Features

### Alias Expansion
- [x] Initialization: `initialize_aliases` -> `third_party/bash/alias.c:L71`
- [x] Expansion Logic (Recursive): Implemented in `internal/shell/shell.go`
- [x] Tokenization for Aliases: Implemented
- [x] Whitespace handling: Implemented

### Array Support
- [x] **Indexed Arrays**: Implemented in Environment
    - [x] `array_insert`: Implemented via index assignment
    - [x] `array_reference`: Implemented in `internal/shell/shell.go`
    - [x] Subrange expansion `${a[@]:s:n}`: Implemented in `internal/shell/shell.go`
- [x] **Associative Arrays**: Basic map storage implemented in `Environment`
    - [x] `assoc_insert`: Implemented via index assignment
    - [x] `assoc_reference`: Implemented in `internal/shell/shell.go`

### Programmable Completion
- [x] **Core Logic**: Implemented in `internal/shell/input_wasm.go`
- [x] **Builtin Integration**: `complete` and `compgen` registration and storage
- [x] Item Generators (Aliases, Jobs, etc.): Basic generators implemented
