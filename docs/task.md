# Functional Parity Tracking

This document tracks the alignment of the Go Bash Simulator with upstream GNU implementations.

## Overview

Status codes:
- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

---

## Parity Matrix

### `# functional parity tracking`

- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

### `:`

- [x] Upstream: `third_party/bash/builtins/colon.def`
- [x] Basic operation: Implemented in `internal/commands/colon/colon.go`

### `alias`

- [x] Basic management: Implemented in `internal/commands/alias/alias.go`
- [x] Flag `-p`: `third_party/bash/builtins/alias.def:L79` (print)

### `arch`

- [x] Upstream: `third_party/coreutils/src/coreutils-arch.c`
- [x] Inherits flags from `uname`

### `base32`

- [x] Basic encoding/decoding: Implemented in `internal/commands/base32/base32.go`
- [-] Flag `--base16`: Use `basenc` instead
- [-] Flag `--base2lsbf`: Use `basenc` instead
- [-] Flag `--base2msbf`: Use `basenc` instead
- [x] Flag `--base32`: `internal/commands/base32/base32.go` (default)
- [-] Flag `--base32hex`: Use `basenc` instead
- [-] Flag `--base58`: Use `basenc` instead
- [x] Flag `--base64`: `internal/commands/base64/base64.go`
- [x] Flag `--base64url`: `internal/commands/base32/base32.go`
- [x] Flag `--z85`: `internal/commands/base32/base32.go`
- [x] Flag `-d`: `internal/commands/base32/base32.go`
- [x] Flag `-i`: `internal/commands/base32/base32.go`
- [x] Flag `-w`: `internal/commands/base32/base32.go`

### `base64`

- [x] Upstream: `third_party/coreutils/src/base64.c`
- [x] Basic encoding/decoding: Implemented in `internal/commands/base64/base64.go`
- [x] Flag `-d`, `--decode`: `internal/commands/base64/base64.go`
- [x] Flag `-i`, `--ignore-garbage`: `internal/commands/base64/base64.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`, `--wrap=COLS`: `internal/commands/base64/base64.go`

### `basename`

- [x] Upstream: `third_party/coreutils/src/basename.c`
- [x] Basic operation: Implemented in `internal/commands/basename/basename.go`
- [x] Flag `-a`, `--multiple`: `internal/commands/basename/basename.go`
- [x] Flag `-s`, `--suffix=SUFFIX`: `internal/commands/basename/basename.go`
- [x] Flag `-z`, `--zero`: `internal/commands/basename/basename.go`

### `basenc`

- [x] Upstream: `third_party/coreutils/src/basenc.c`
- [x] Basic encoding/decoding: Implemented in `internal/commands/basenc/basenc.go`
- [x] Flag `-d`, `--decode`: `internal/commands/basenc/basenc.go`
- [x] Flag `-i`, `--ignore-garbage`: `internal/commands/basenc/basenc.go`
- [x] Flag `-w`, `--wrap=COLS`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base16`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base32`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base32hex`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base64`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base64url`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base2lsbf`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base2msbf`: `internal/commands/basenc/basenc.go`
- [x] Flag `--base58`: `internal/commands/basenc/basenc.go`
- [x] Flag `--z85`: `internal/commands/basenc/basenc.go`

### `bash`

- [x] Upstream: `third_party/bash/shell.c`, `third_party/bash/version.c`
- [x] Flag `--version`: `third_party/bash/shell.c:L483`, `third_party/bash/version.c:L88`

### `bg`

- [x] Upstream: `third_party/bash/builtins/fg_bg.def`
- [x] Basic job management: Implemented in `internal/commands/bg/bg.go`
- [x] Job specification support: `internal/commands/bg/bg.go`

### `bind`

- [x] Upstream: `third_party/bash/builtins/bind.def`
- [x] Keybinding management: Implemented in `internal/commands/bind/bind.go`
- [x] Flag `-l`: `internal/commands/bind/bind.go` (list)
- [x] Flag `-v`: `internal/commands/bind/bind.go` (list functions)
- [x] Flag `-p`: `internal/commands/bind/bind.go` (print status)
- [x] Flag `-V`: `internal/commands/bind/bind.go` (list variables)
- [x] Flag `-P`: `internal/commands/bind/bind.go` (print functions)
- [x] Flag `-s`: `internal/commands/bind/bind.go` (list macros)
- [x] Flag `-S`: `internal/commands/bind/bind.go` (print macros)
- [x] Flag `-X`: `internal/commands/bind/bind.go` (list keyseq bindings)
- [x] Flag `-f=FILE`: `internal/commands/bind/bind.go` (read from file)
- [x] Flag `-q=FUNC`: `internal/commands/bind/bind.go` (query keys for func)
- [x] Flag `-u=FUNC`: `internal/commands/bind/bind.go` (unbind func)
- [x] Flag `-m=KEYMAP`: `internal/commands/bind/bind.go` (keymap)
- [x] Flag `-r=KEYSEQ`: `internal/commands/bind/bind.go` (remove seq)
- [x] Flag `-x=KEYSEQ:SHELLCMD`: `internal/commands/bind/bind.go` (exec cmd)

### `break`

- [x] Upstream: `third_party/bash/builtins/break.def`
- [x] Basic operation: Implemented in `internal/commands/break/break.go`

### `builtin`

- [x] Upstream: `third_party/bash/builtins/builtin.def`
- [x] Basic execution: Implemented in `internal/commands/builtin/builtin.go`

### `caller`

- [x] Upstream: `third_party/bash/builtins/caller.def`
- [x] Basic operation: Implemented in `internal/commands/caller/caller.go`

### `cat`

- [x] Upstream: `third_party/coreutils/src/cat.c`
- [x] Basic output: Implemented in `internal/commands/cat/cat.go`
- [x] Flag `-A`, `--show-all`: `internal/commands/cat/cat.go`
- [x] Flag `-b`, `--number-nonblank`: `internal/commands/cat/cat.go`
- [x] Flag `-e`: `internal/commands/cat/cat.go` (implies -vE)
- [x] Flag `-E`, `--show-ends`: `internal/commands/cat/cat.go`
- [x] Flag `-n`, `--number`: `internal/commands/cat/cat.go`
- [x] Flag `-s`, `--squeeze-blank`: `internal/commands/cat/cat.go`
- [x] Flag `-t`: `internal/commands/cat/cat.go` (implies -vT)
- [x] Flag `-T`, `--show-tabs`: `internal/commands/cat/cat.go`
- [x] Flag `-u`: `internal/commands/cat/cat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#cat))
- [x] Flag `-v`, `--show-nonprinting`: `internal/commands/cat/cat.go`

### `cd`

- [x] Upstream: `third_party/bash/builtins/cd.def`
- [x] Basic change directory: Implemented in `internal/commands/cd/cd.go`
- [x] CDPATH support: `internal/commands/cd/cd.go`
- [x] Flag `-e`: `internal/commands/cd/cd.go` (exit status if -P cannot be satisfied)
- [x] Flag `-L`: `internal/commands/cd/cd.go`
- [x] Flag `-P`: `internal/commands/cd/cd.go`

### `chcon`

- [x] Upstream: `third_party/coreutils/src/chcon.c`
- [x] Basic operation: `internal/commands/chcon/chcon.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chcon-runcon))
- [x] Flag `-h`, `--no-dereference`: `internal/commands/chcon/chcon.go`
- [x] Flag `-H`: `internal/commands/chcon/chcon.go`
- [x] Flag `-L`: `internal/commands/chcon/chcon.go`
- [x] Flag `-P`: `internal/commands/chcon/chcon.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/chcon/chcon.go`
- [x] Flag `-u`, `--user=USER`: `internal/commands/chcon/chcon.go`
- [x] Flag `-r`, `--role=ROLE`: `internal/commands/chcon/chcon.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/chcon/chcon.go`
- [x] Flag `-l`, `--range=RANGE`: `internal/commands/chcon/chcon.go`
- [x] Flag `--reference=RFILE`: `internal/commands/chcon/chcon.go`

### `chgrp`

- [x] Upstream: `third_party/coreutils/src/chgrp.c`
- [x] Basic operation: Implemented in `internal/commands/chown/chown.go` (via `NewChgrp`)
- [x] Flag `-c`, `--changes`: `internal/commands/chown/chown.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/chown/chown.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/chown/chown.go`
- [x] Flag `--dereference`: `internal/commands/chown/chown.go`
- [x] Flag `-h`, `--no-dereference`: `internal/commands/chown/chown.go`
- [x] Flag `--reference=RFILE`: `internal/commands/chown/chown.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/chown/chown.go`
- [x] Flag `-H`, `-L`, `-P`: `internal/commands/chown/chown.go` (Stub)

### `chmod`

- [x] Upstream: `third_party/coreutils/src/chmod.c`
- [x] Basic operation: Implemented in `internal/commands/chmod/chmod.go`
- [x] Flag `-c`, `--changes`: `internal/commands/chmod/chmod.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/chmod/chmod.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/chmod/chmod.go`
- [x] Flag `--reference=RFILE`: `internal/commands/chmod/chmod.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/chmod/chmod.go`
- [x] Symbolic modes (u+x, g-w, etc.): Implemented in `internal/commands/chmod/chmod.go`

### `chown`

- [x] Upstream: `third_party/coreutils/src/chown.c`
- [x] Basic operation: Implemented in `internal/commands/chown/chown.go`
- [x] Flag `-c`, `--changes`: `internal/commands/chown/chown.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/chown/chown.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/chown/chown.go`
- [x] Flag `--dereference`: `internal/commands/chown/chown.go`
- [x] Flag `-h`, `--no-dereference`: `internal/commands/chown/chown.go`
- [x] Flag `--from=CURRENT_OWNER:CURRENT_GROUP`: `internal/commands/chown/chown.go` (Stub)
- [x] Flag `--reference=RFILE`: `internal/commands/chown/chown.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/chown/chown.go`
- [x] Flag `-H`, `-L`, `-P`: `internal/commands/chown/chown.go` (Stub)

### `chroot`

- [x] Upstream: `third_party/coreutils/src/chroot.c`
- [x] Basic operation: Implemented in `internal/commands/chroot/chroot.go`

### `cksum`

- [x] Upstream: `third_party/coreutils/src/cksum.c`
- [x] Basic CRC-32: Implemented in `internal/commands/cksum/cksum.go`
- [x] Flag `-a`, `--algorithm=ALGO`: `internal/commands/cksum/cksum.go`
- [x] Flag `-c`, `--check`: `internal/commands/cksum/cksum.go`
- [x] Flag `-l`, `--length=BITS`: `internal/commands/cksum/cksum.go`
- [x] Flag `-z`, `--zero`: `internal/commands/cksum/cksum.go`
- [x] Flag `--base64`: `internal/commands/cksum/cksum.go`
- [x] Flag `--raw`: `internal/commands/cksum/cksum.go`
- [x] Flag `--tag`: `internal/commands/cksum/cksum.go`
- [x] Flag `--untagged`: `internal/commands/cksum/cksum.go`
- [x] Flag `--ignore-missing`: `internal/commands/cksum/cksum.go`
- [x] Flag `--quiet`: `internal/commands/cksum/cksum.go`
- [x] Flag `--status`: `internal/commands/cksum/cksum.go`
- [x] Flag `--strict`: `internal/commands/cksum/cksum.go`
- [x] Flag `-w`, `--warn`: `internal/commands/cksum/cksum.go`
 
### `clear`

- [x] Basic operation: Implemented in `internal/commands/clear/clear.go`
 
### `comm`

- [x] Upstream: `third_party/coreutils/src/comm.c`
- [x] Basic comparison: Implemented in `internal/commands/comm/comm.go`
- [x] Flag `--check-order`: `internal/commands/comm/comm.go`
- [x] Flag `--nocheck-order`: `internal/commands/comm/comm.go`
- [x] Flag `--output-delimiter`: `internal/commands/comm/comm.go`
- [x] Flag `--total`: `internal/commands/comm/comm.go`
- [x] Flag `-1`: `internal/commands/comm/comm.go`
- [x] Flag `-2`: `internal/commands/comm/comm.go`
- [x] Flag `-3`: `internal/commands/comm/comm.go`
- [x] Flag `-z`: `internal/commands/comm/comm.go`

### `command`

- [x] Upstream: `third_party/bash/builtins/command.def`
- [x] Execution override: Implemented in `internal/commands/command/command.go`
- [x] Flag `-p`: `internal/commands/command/command.go`
- [x] Flag `-v`: `internal/commands/command/command.go`
- [x] Flag `-V`: `internal/commands/command/command.go`

### `compgen`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Inherits all `complete` flags: Implemented in `internal/commands/compgen/compgen.go`

### `complete`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Completion management: Implemented in `internal/commands/complete/complete.go`
- [x] Flag `-a`: `internal/commands/complete/complete.go` (alias)
- [x] Flag `-b`: `internal/commands/complete/complete.go` (builtin)
- [x] Flag `-c`: `internal/commands/complete/complete.go` (command)
- [x] Flag `-d`: `internal/commands/complete/complete.go` (directory)
- [x] Flag `-e`: `internal/commands/complete/complete.go` (export)
- [x] Flag `-f`: `internal/commands/complete/complete.go` (file)
- [x] Flag `-g`: `internal/commands/complete/complete.go` (group)
- [x] Flag `-j`: `internal/commands/complete/complete.go` (job)
- [x] Flag `-k`: `internal/commands/complete/complete.go` (keyword)
- [x] Flag `-p`: `internal/commands/complete/complete.go` (print)
- [x] Flag `-r`: `internal/commands/complete/complete.go` (remove)
- [x] Flag `-s`: `internal/commands/complete/complete.go` (service)
- [x] Flag `-u`: `internal/commands/complete/complete.go` (user)
- [x] Flag `-v`: `internal/commands/complete/complete.go` (variable)
- [x] Flag `-o=OPT`: `internal/commands/complete/complete.go`
- [x] Flag `-A=ACTION`: `internal/commands/complete/complete.go`
- [x] Flag `-G=GLOB`: `internal/commands/complete/complete.go`
- [x] Flag `-W=WORDLIST`: `internal/commands/complete/complete.go`
- [x] Flag `-P=PREFIX`: `internal/commands/complete/complete.go`
- [x] Flag `-S=SUFFIX`: `internal/commands/complete/complete.go`
- [x] Flag `-X=FILTER`: `internal/commands/complete/complete.go`
- [x] Flag `-F=FUNC`: `internal/commands/complete/complete.go`
- [x] Flag `-C=CMD`: `internal/commands/complete/complete.go`
- [x] Flag `-E`: `internal/commands/complete/complete.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-I`: `internal/commands/complete/complete.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-D`: `internal/commands/complete/complete.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `compopt`

- [x] Upstream: `third_party/bash/builtins/complete.def`
- [x] Flag `-o`, `--options`: `internal/commands/compopt/compopt.go`
- [x] Flag `-D`, `--default`: `internal/commands/compopt/compopt.go`
- [x] Flag `-E`, `--empty`: `internal/commands/compopt/compopt.go`

### `continue`

- [x] Upstream: `third_party/bash/builtins/break.def`
- [x] Basic operation: Implemented in `internal/commands/continue/continue.go`

### `cp`

- [x] Basic copy: Implemented in `internal/commands/cp/cp.go`
- [x] Flag `-a`, `--archive`: `internal/commands/cp/cp.go`
- [x] Flag `-b`, `--backup`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-d`: `internal/commands/cp/cp.go`
- [x] Flag `-f`, `--force`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-H`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--interactive`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`, `--link`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-L`, `--dereference`: `internal/commands/cp/cp.go`
- [x] Flag `-n`, `--no-clobber`: `internal/commands/cp/cp.go`
- [x] Flag `-p`: `internal/commands/cp/cp.go`
- [x] Flag `-P`, `--no-dereference`: `internal/commands/cp/cp.go`
- [x] Flag `-r`, `-R`, `--recursive`: `internal/commands/cp/cp.go`
- [x] Flag `-s`, `--symbolic-link`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--target-directory`: `internal/commands/cp/cp.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/cp/cp.go`
- [x] Flag `-u`, `--update`: `internal/commands/cp/cp.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/cp/cp.go`
- [x] Flag `-x`, `--one-file-system`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-Z`, `--context`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--attributes-only`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--preserve[=ATTR_LIST]`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--no-preserve=ATTR_LIST`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--parents`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--reflink[=WHEN]`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--sparse=WHEN`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--strip-trailing-slashes`: `internal/commands/cat/cat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `csplit`

- [x] Upstream: `third_party/coreutils/src/csplit.c`
- [x] Basic split: Implemented in `internal/commands/csplit/csplit.go`
- [x] Flag `--suppress-matched`: `internal/commands/csplit/csplit.go`
- [x] Flag `-b`: `internal/commands/csplit/csplit.go`
- [x] Flag `-f`: `internal/commands/csplit/csplit.go`
- [x] Flag `-k`: `internal/commands/csplit/csplit.go`
- [x] Flag `-n`: `internal/commands/csplit/csplit.go`
- [x] Flag `-s`: `internal/commands/csplit/csplit.go`
- [x] Flag `-z`: `internal/commands/csplit/csplit.go`

### `cut`

- [x] Upstream: `third_party/coreutils/src/cut.c`
- [x] Basic selection: Implemented in `internal/commands/cut/cut.go`
- [x] Flag `-b`, `--bytes=LIST`: `internal/commands/cut/cut.go`
- [x] Flag `-c`, `--characters=LIST`: `internal/commands/cut/cut.go`
- [x] Flag `-d`, `--delimiter=DELIM`: `internal/commands/cut/cut.go`
- [x] Flag `-f`, `--fields=LIST`: `internal/commands/cut/cut.go`
- [x] Flag `-n`: `internal/commands/cut/cut.go` (Ignored)
- [x] Flag `-s`, `--only-delimited`: `internal/commands/cut/cut.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/cut/cut.go`
- [x] Flag `--complement`: `internal/commands/cut/cut.go`
- [x] Flag `--output-delimiter=STRING`: `internal/commands/cut/cut.go`

### `date`

- [x] Upstream: `third_party/coreutils/src/date.c`
- [x] Basic output: Implemented in `internal/commands/date/date.go`
- [x] Custom format `+FORMAT`: `internal/commands/date/date.go`
- [x] Flag `-d`, `--date=STRING`: `internal/commands/date/date.go`
- [x] Flag `-f`, `--file=DATEFILE`: `internal/commands/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-I[FMT]`, `--iso-8601[=FMT]`: `internal/commands/date/date.go`
- [x] Flag `-r`, `--reference=FILE`: `internal/commands/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-R`, `--rfc-email`: `internal/commands/date/date.go`
- [x] Flag `-s`, `--set=STRING`: `internal/commands/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--utc`, `--universal`: `internal/commands/date/date.go`
- [x] Flag `--debug`: `internal/commands/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `dd`

- [x] Upstream: `third_party/coreutils/src/dd.c`
- [x] Data copy: Implemented in `internal/commands/dd/dd.go`
- [x] Operand `bs=BYTES`: `internal/commands/dd/dd.go`
- [x] Operand `cbs=BYTES`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `conv=CONVS`: `internal/commands/dd/dd.go` (Partial via notrunc; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `count=N`: `internal/commands/dd/dd.go`
- [x] Operand `ibs=BYTES`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `if=FILE`: `internal/commands/dd/dd.go`
- [x] Operand `iflag=FLAGS`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `obs=BYTES`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `of=FILE`: `internal/commands/dd/dd.go`
- [x] Operand `oflag=FLAGS`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `seek=N`: `internal/commands/dd/dd.go`
- [x] Operand `skip=N`: `internal/commands/dd/dd.go`
- [x] Operand `status=LEVEL`: `internal/commands/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Operand `conv=notrunc`: `internal/commands/dd/dd.go`

### `declare`

- [x] Attribute management (-i, -r, -x, -a, -A): Implemented in `internal/commands/declare/declare.go`
- [x] Flag `-a`: `internal/commands/declare/declare.go`
- [x] Flag `-A`: `internal/commands/declare/declare.go`
- [x] Flag `-i`: `internal/commands/declare/declare.go`
- [x] Flag `-r`: `internal/commands/declare/declare.go`
- [x] Flag `-x`: `internal/commands/declare/declare.go`
- [x] Flag `-l`: `internal/commands/declare/declare.go` (lowercase)
- [x] Flag `-u`: `internal/commands/declare/declare.go` (uppercase)
- [x] Flag `-n`: `internal/commands/declare/declare.go` (nameref stub)
- [x] Flag `-t`: `internal/commands/declare/declare.go` (trace stub)
- [x] Flag `-f`: `internal/commands/declare/declare.go` (function stub)
- [x] Flag `-F`: `internal/commands/declare/declare.go` (function name stub)
- [x] Flag `-g`: `internal/commands/declare/declare.go` (global stub)
- [x] Flag `-p`: `internal/commands/declare/declare.go`
- [x] Flag `-I`: `internal/commands/declare/declare.go` (inherit stub)
- [x] Aliases: `typeset`

### `df`

- [x] Upstream: `third_party/coreutils/src/df.c`
- [x] Basic df: Implemented in `internal/commands/df/df.go`
- [x] Basic output: Implemented in `internal/commands/df/df.go`
- [x] Flag `--no-sync`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--output[=FIELD_LIST]`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--sync`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--total`: `internal/commands/df/df.go`
- [x] Flag `-B`, `--block-size=SIZE`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-H`: `internal/commands/df/df.go`
- [x] Flag `-P`: `internal/commands/df/df.go`
- [x] Flag `-T`: `internal/commands/df/df.go`
- [x] Flag `-a`: `internal/commands/df/df.go`
- [x] Flag `-h`: `internal/commands/df/df.go`
- [x] Flag `-i`: `internal/commands/df/df.go`
- [x] Flag `-k`: `internal/commands/df/df.go`
- [x] Flag `-l`: `internal/commands/df/df.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-x`, `--exclude-type=TYPE`: `internal/commands/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `dir`

- [x] Upstream: `third_party/coreutils/src/coreutils-dir.c`
- [x] Inherits flags from `ls`

### `dircolors`

- [x] Upstream: `third_party/coreutils/src/dircolors.c`
- [x] Output configuration: Implemented in `internal/commands/dircolors/dircolors.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#dircolors))
- [x] Flag `-b`, `--sh`, `--bourne-shell`: `internal/commands/dircolors/dircolors.go`
- [x] Flag `-c`, `--csh`, `--c-shell`: `internal/commands/dircolors/dircolors.go`
- [x] Flag `-p`, `--print-database`: `internal/commands/dircolors/dircolors.go`
- [x] Flag `--print-ls-colors`: `internal/commands/dircolors/dircolors.go`

### `dirname`

- [x] Upstream: `third_party/coreutils/src/dirname.c`
- [x] Basic operation: Implemented in `internal/commands/dirname/dirname.go`
- [x] Flag `-z`, `--zero`: `internal/commands/dirname/dirname.go`

### `dirs`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic listing: Implemented in `internal/commands/dirs/dirs.go`
- [x] Flag `-c`: `internal/commands/dirs/dirs.go` (clear stack)
- [x] Flag `-l`: `internal/commands/pushd/pushd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/dirs/dirs.go` (print with one line per entry)
- [x] Flag `-v`: `internal/commands/dirs/dirs.go` (verbose)

### `disown`

- [x] Upstream: `third_party/bash/builtins/jobs.def`
- [x] Flag `-a`: `internal/commands/disown/disown.go` (all jobs)
- [x] Flag `-h`: `internal/commands/disown/disown.go` (mark to not receive SIGHUP)
- [x] Flag `-r`: `internal/commands/disown/disown.go` (running jobs only)

### `du`

- [x] Upstream: `third_party/coreutils/src/du.c`
- [x] Basic operation: Implemented in `internal/commands/du/du.go`
- [x] Flag `-0`, `--null`: `internal/commands/du/du.go`
- [x] Flag `-a`, `--all`: `internal/commands/du/du.go`
- [x] Flag `--apparent-size`: `internal/commands/du/du.go`
- [x] Flag `-b`, `--bytes`: `internal/commands/du/du.go` (alias for apparent-size)
- [x] Flag `-c`, `--total`: `internal/commands/du/du.go`
- [x] Flag `-d`, `--max-depth=N`: `internal/commands/du/du.go`
- [x] Flag `-h`, `--human-readable`: `internal/commands/du/du.go`
- [x] Flag `-k`: `internal/commands/du/du.go`
- [x] Flag `-m`: `internal/commands/du/du.go`
- [x] Flag `-s`, `--summarize`: `internal/commands/du/du.go`
- [x] Flag `-S`, `--separate-dirs`: `internal/commands/du/du.go` (Stub)
- [x] Flag `-D`, `--dereference-args`: `internal/commands/du/du.go`
- [x] Flag `-H`: `internal/commands/du/du.go` (dereference)
- [x] Flag `-l`, `--count-links`: `internal/commands/du/du.go` (partial via size)
- [x] Flag `-L`, `--dereference`: `internal/commands/du/du.go`
- [x] Flag `-P`, `--no-dereference`: `internal/commands/du/du.go`
- [x] Flag `-t`, `--threshold=SIZE`: `internal/commands/du/du.go`
- [x] Flag `-x`, `--one-file-system`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-X`, `--exclude-from=FILE`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--exclude=PATTERN`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--files0-from=F`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--inodes`: `internal/commands/du/du.go`
- [x] Flag `--si`: `internal/commands/du/du.go`
- [x] Flag `--time[=WORD]`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--time-style=STYLE`: `internal/commands/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `echo`

- [x] Upstream: `third_party/bash/builtins/echo.def`
- [x] Basic output: Implemented in `internal/commands/echo/echo.go`
- [x] Flag `-n`: `internal/commands/echo/echo.go` (no newline)
- [x] Flag `-e`: `internal/commands/echo/echo.go` (interpret escapes)
- [x] Flag `-E`: `internal/commands/echo/echo.go` (disable escapes)

### `enable`

- [x] Upstream: `third_party/bash/builtins/enable.def`
- [x] Basic management: Implemented in `internal/commands/enable/enable.go`
- [x] Flag `-a`: `third_party/bash/builtins/enable.def:L157` (display all)
- [-] Flag `-d`: `third_party/bash/builtins/enable.def:L160` (delete loaded) - Dynamic loading not available
- [x] Flag `-n`: `third_party/bash/builtins/enable.def:L163` (disable)
- [x] Flag `-p`: `third_party/bash/builtins/enable.def:L166` (print status)
- [x] Flag `-s`: `third_party/bash/builtins/enable.def:L169` (POSIX special only)
- [-] Flag `-f filename`: `third_party/bash/builtins/enable.def:L172` (load from dynamic file) - Dynamic loading not available

### `env`

- [x] Upstream: `third_party/coreutils/src/env.c`
- [x] Basic execution: Implemented in `internal/commands/env/env.go`
- [x] Flag `-a`, `--argv0=ARG`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--ignore-environment`: `internal/commands/env/env.go`
- [x] Flag `-u`, `--unset=NAME`: `internal/commands/env/env.go`
- [x] Flag `-0`, `--null`: `internal/commands/env/env.go`
- [x] Flag `-C`, `--chdir=DIR`: `internal/commands/env/env.go`
- [x] Flag `-S`, `--split-string=S`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--block-signal[=SIG]`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--default-signal[=SIG]`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--ignore-signal[=SIG]`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--list-signal-handling`: `internal/commands/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `eval`

- [x] Upstream: `third_party/bash/builtins/eval.def`
- [x] Basic execution: Implemented in `internal/commands/eval/eval.go`

### `exec`

- [x] Upstream: `third_party/bash/builtins/exec.def`
- [x] Basic execution: Implemented in `internal/commands/exec/exec.go`
- [x] Flag `-l`: `internal/commands/exec/exec.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-a name`: `internal/commands/exec/exec.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-c`: `internal/commands/exec/exec.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `exit`

- [x] Upstream: `third_party/bash/builtins/exit.def`
- [x] Basic exit: Implemented in `internal/commands/exit/exit.go`
- [x] Exit status parameter: `internal/commands/exit/exit.go`

### `expand`

- [x] Upstream: `third_party/coreutils/src/expand.c`
- [x] Basic conversion: Implemented in `internal/commands/expand/expand.go`
- [x] Flag `-i`, `--initial`: `internal/commands/expand/expand.go`
- [x] Flag `-t`, `--tabs=LIST`: `internal/commands/expand/expand.go`

### `export`

- [x] Upstream: `third_party/bash/builtins/setattr.def`
- [x] Basic operation: Implemented in `internal/commands/export/export.go`
- [x] Flag `-f`: `internal/commands/export/export.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-n`: `internal/commands/export/export.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/export/export.go`

### `expr`

- [x] Upstream: `third_party/coreutils/src/expr.c`
- [x] Expression evaluation: Implemented in `internal/commands/expr/expr.go`
- [x] Arithmetic (+, -, *, /, %): Implemented in `internal/commands/expr/expr.go`
- [x] Comparison (=, !=, <, <=, >, >=): Implemented in `internal/commands/expr/expr.go`
- [x] Logical (| , &): Implemented in `internal/commands/expr/expr.go`
- [x] String operators (match, substr, index, length): Implemented in `internal/commands/expr/expr.go`
- [x] Flag `--help`: `internal/commands/expr/expr.go`
- [x] Flag `--version`: `internal/commands/expr/expr.go`

### `factor`

- [x] Upstream: `third_party/coreutils/src/factor.c`
- [x] Prime factorization: Implemented in `internal/commands/factor/factor.go`

### `find`

- [x] Upstream: `third_party/coreutils/src/find.c`
- [x] Basic Search: `internal/commands/find/find.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#find))
- [x] Flag `-name`: `internal/commands/find/find.go`
- [x] Flag `-type`: `internal/commands/find/find.go`

### `false`

- [x] Upstream: `third_party/bash/builtins/colon.def`, `third_party/coreutils/src/false.c`
- [x] Basic operation: Implemented in `internal/commands/boolcmd/bool.go`

### `fc`

- [x] Upstream: `third_party/bash/builtins/fc.def`
- [x] Basic editing/re-execution: Implemented in `internal/commands/fc/fc.go`
- [x] Flag `-e ENAME`: `internal/commands/fc/fc.go`
- [x] Flag `-l`: `internal/commands/fc/fc.go`
- [x] Flag `-n`: `internal/commands/fc/fc.go`
- [x] Flag `-r`: `internal/commands/fc/fc.go`
- [x] Flag `-s`: `internal/commands/fc/fc.go`

### `fg`

- [x] Upstream: `third_party/bash/builtins/fg_bg.def`
- [x] Basic operation: Implemented in `internal/commands/fg/fg.go`

### `fmt`

- [x] Upstream: `third_party/coreutils/src/fmt.c`
- [x] Paragraph formatting: Implemented in `internal/commands/fmt/fmt.go`
- [x] Flag `-c`, `--crown-margin`: `internal/commands/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`, `--prefix=STRING`: `internal/commands/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`, `--split-only`: `internal/commands/fmt/fmt.go`
- [x] Flag `-t`, `--tagged-paragraph`: `internal/commands/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--uniform-spacing`: `internal/commands/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/fmt/fmt.go`
- [x] Flag `-g`, `--goal=WIDTH`: `internal/commands/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-WIDTH`: `internal/commands/fmt/fmt.go`

### `fold`

- [x] Upstream: `third_party/coreutils/src/fold.c`
- [x] Line wrapping: Implemented in `internal/commands/fold/fold.go`
- [x] Flag `-b`, `--bytes`: `internal/commands/fold/fold.go`
- [x] Flag `-c`, `--characters`: `internal/commands/fold/fold.go`
- [x] Flag `-s`, `--spaces`: `internal/commands/fold/fold.go`
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/fold/fold.go`

### `getlimits`

- [x] Upstream: `third_party/coreutils/src/getlimits.c`
- [x] Basic operation: Implemented in `internal/commands/getlimits/getlimits.go`
- [x] Flag `--help`: `internal/commands/getlimits/getlimits.go`
- [x] Flag `--version`: `internal/commands/getlimits/getlimits.go`

### `getopts`

- [x] Basic parsing: Implemented in `internal/commands/getopts/getopts.go`
- [x] Silent mode support (`:`): `internal/commands/getopts/getopts.go`

### `grep`

- [x] Upstream: `third_party/coreutils/src/grep.c`
- [x] Regex Search: `internal/commands/grep/grep.go` (Workaround; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#grep))
- [x] Flag `-i`, `--ignore-case`: `internal/commands/grep/grep.go`
- [x] Flag `-v`, `--invert-match`: `internal/commands/grep/grep.go`
- [x] Flag `-n`, `--line-number`: `internal/commands/grep/grep.go`
- [x] Flag `-c`, `--count`: `internal/commands/grep/grep.go`
- [x] Flag `-l`, `--files-with-matches`: `internal/commands/grep/grep.go`

### `groups`

- [x] Upstream: `third_party/coreutils/src/groups.c`
- [x] Basic listing: Implemented in `internal/commands/groups/groups.go`
- [x] Multiple users support: `internal/commands/groups/groups.go`

### `hash`

- [x] Upstream: `third_party/bash/builtins/hash.def`
- [x] Command hashing: Implemented in `internal/commands/hash/hash.go`
- [x] Flag `-r`: `internal/commands/hash/hash.go`
- [x] Flag `-d`: `internal/commands/hash/hash.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `internal/commands/hash/hash.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`: `internal/commands/hash/hash.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/hash/hash.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `head`

- [x] Upstream: `third_party/coreutils/src/head.c`
- [x] Basic output: Implemented in `internal/commands/head/head.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/head/head.go`
- [x] Flag `-n`, `--lines`: `internal/commands/head/head.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/head/head.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/head/head.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/head/head.go`

### `help`

- [x] Upstream: `third_party/bash/builtins/help.def`
- [x] Help system: Implemented in `internal/commands/help/help.go`
- [x] Flag `-d`: `internal/commands/help/help.go` (short description)
- [x] Flag `-m`: `internal/commands/help/help.go` (man-page format)
- [x] Flag `-s`: `internal/commands/help/help.go` (syntax only)

### `history`

- [x] Upstream: `third_party/bash/builtins/history.def`
- [x] History management: Implemented in `internal/commands/history/history.go`
- [x] Flag `-d offset`: `internal/commands/history/history.go` (delete entry)
- [x] Flag `-a`: `internal/commands/history/history.go` (append)
- [x] Flag `-c`: `internal/commands/history/history.go` (clear)
- [x] Flag `-n`: `internal/commands/history/history.go` (read non-recorded)
- [x] Flag `-p`: `internal/commands/history/history.go` (print/expand)
- [x] Flag `-r`: `internal/commands/history/history.go` (read file)
- [x] Flag `-s`: `internal/commands/history/history.go` (store/append)
- [x] Flag `-w`: `internal/commands/history/history.go` (write file)

### `hostid`

- [x] Upstream: `third_party/coreutils/src/hostid.c`
- [x] Basic operation: Implemented in `internal/commands/hostid/hostid.go`

### `hostname`

- [x] Upstream: `third_party/coreutils/src/hostname.c`
- [x] Basic output: Implemented in `internal/commands/hostname/hostname.go`
- [x] Set hostname support: `internal/commands/hostname/hostname.go`
- [x] Flag `--help`: `internal/commands/hostname/hostname.go`
- [x] Flag `--version`: `internal/commands/hostname/hostname.go`

### `id`

- [x] Upstream: `third_party/coreutils/src/id.c`
- [x] Basic output: Implemented in `internal/commands/id/id.go`
- [x] Flag `-G`: `internal/commands/id/id.go`
- [x] Flag `-Z`: `internal/commands/id/id.go` (Unsupported; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#id))
- [x] Flag `-a`: `internal/commands/id/id.go` (Ignored)
- [x] Flag `-g`: `internal/commands/id/id.go`
- [x] Flag `-n`: `internal/commands/id/id.go`
- [x] Flag `-r`: `internal/commands/id/id.go` (real == effective)
- [x] Flag `-u`: `internal/commands/id/id.go`
- [x] Flag `-z`: `internal/commands/id/id.go`

### `install`

- [x] Upstream: `third_party/coreutils/src/install.c`
- [x] Flag `-c`: `internal/commands/install/install.go` (ignored)
- [x] Flag `-d`, `--directory`: `internal/commands/install/install.go`
- [x] Flag `-D`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-g`, `--group=GROUP`: `internal/commands/install/install.go`
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/install/install.go` (Stub)
- [x] Flag `-o`, `--owner=OWNER`: `internal/commands/install/install.go`
- [x] Flag `-p`, `--preserve-timestamps`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`, `--strip`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-S`, `--suffix=SUFFIX`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--target-directory=DIR`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/install/install.go`
- [x] Flag `-C`, `--compare`: `internal/commands/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `jobs`

- [x] Upstream: `third_party/bash/builtins/jobs.def`
- [x] Basic listing: Implemented in `internal/commands/jobs/jobs.go`
- [x] Flag `-l`: `internal/commands/jobs/jobs.go` (long format)
- [x] Flag `-n`: `internal/commands/jobs/jobs.go` (only jobs that changed)
- [x] Flag `-p`: `internal/commands/jobs/jobs.go` (only PIDs)
- [x] Flag `-r`: `internal/commands/jobs/jobs.go` (running only)
- [x] Flag `-s`: `internal/commands/jobs/jobs.go` (stopped only)
- [x] Flag `-x command`: `internal/commands/jobs/jobs.go` (execute command)

### `join`

- [x] Upstream: `third_party/coreutils/src/join.c`
- [x] Basic join: Implemented in `internal/commands/join/join.go`
- [x] Flag `-a FILENUM`: `internal/commands/join/join.go` (unpairable lines from file FILENUM)
- [x] Flag `-e EMPTY`: `internal/commands/join/join.go` (replace empty input fields with EMPTY)
- [x] Flag `-i`, `--ignore-case`: `internal/commands/join/join.go`
- [x] Flag `-j FIELD`: `internal/commands/join/join.go` (equivalent to -1 FIELD -2 FIELD)
- [x] Flag `-o FORMAT`: `internal/commands/join/join.go` (obey FORMAT while constructing output line)
- [x] Flag `-t CHAR`: `internal/commands/join/join.go` (use CHAR as input and output field separator)
- [x] Flag `-v FILENUM`: `internal/commands/join/join.go` (like -a FILENUM, but suppress joined output lines)
- [x] Flag `-1 FIELD`: `internal/commands/join/join.go` (join on this FIELD of file 1)
- [x] Flag `-2 FIELD`: `internal/commands/join/join.go` (join on this FIELD of file 2)
- [x] Flag `--check-order`: `internal/commands/join/join.go` (check that the input is correctly sorted)
- [x] Flag `--nocheck-order`: `internal/commands/join/join.go` (do not check that the input is correctly sorted)
- [x] Flag `--header`: `internal/commands/join/join.go` (treat the first line of each file as field headers)

### `kill`

- [x] Upstream: `third_party/bash/builtins/kill.def`
- [x] Basic signaling: Implemented in `internal/commands/kill/kill.go`
- [x] Flag `-l`: `internal/commands/kill/kill.go`
- [x] Flag `-n num`: `internal/commands/kill/kill.go`
- [x] Flag `-l`: `internal/commands/kill/kill.go`
- [x] Flag `-s SIGNAL`: `internal/commands/kill/kill.go`

### `let`

- [x] Upstream: `third_party/bash/builtins/let.def`
- [x] Basic arithmetic: Implemented in `internal/commands/letcmd/let.go`

### `link`

- [x] Basic hard link: Implemented in `internal/commands/link/link.go`

### `ln`

- [x] Basic link creation: Implemented in `internal/commands/ln/ln.go`
- [x] Flag `-f`: `internal/commands/ln/ln.go`
- [x] Flag `-s`: `internal/commands/ln/ln.go`
- [x] Flag `-v`: `internal/commands/ln/ln.go`

### `local`

- [x] Upstream: `third_party/bash/builtins/declare.def`
- [x] Basic operation: Implemented in `internal/commands/local/local.go`
- [x] Inherits generic `declare` behavior

### `logname`

- [x] Upstream: `third_party/coreutils/src/logname.c`
- [x] Basic operation: Implemented in `internal/commands/logname/logname.go`
- [x] Flag `--help`: `internal/commands/logname/logname.go`
- [x] Flag `--version`: `internal/commands/logname/logname.go`

### `logout`

- [x] Upstream: `third_party/bash/builtins/exit.def`
- [x] Basic operation: Implemented in `internal/commands/logout/logout.go`

### `ls`

- [x] Basic listing: `internal/commands/ls/ls.go`
- [x] Color output (`--color`): `internal/commands/ls/ls.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#ls))
- [x] Flag `--author`: `internal/commands/ls/ls.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#ls))
- [x] Flag `--block-size`: `internal/commands/ls/ls.go`
- [x] Flag `--color`: `internal/commands/ls/ls.go` (ANSI colors)
- [x] Flag `--dereference-command-line-symlink-to-dir`: `internal/commands/ls/ls.go` (-H)
- [x] Flag `--file-type`: `internal/commands/ls/ls.go`
- [x] Flag `--format`: `internal/commands/ls/ls.go` (across, commas, horizontal, long, single-column, verbose, vertical)
- [x] Flag `--full-time`: `internal/commands/ls/ls.go`
- [x] Flag `--group-directories-first`: `internal/commands/ls/ls.go`
- [x] Flag `--hide`: `internal/commands/ls/ls.go`
- [x] Flag `--indicator-style`: `internal/commands/ls/ls.go`
- [x] Flag `--quoting-style`: `internal/commands/ls/ls.go`
- [x] Flag `--show-control-chars`: `internal/commands/ls/ls.go` (partial via -q)
- [x] Flag `--si`: `internal/commands/ls/ls.go` (power of 1000)
- [x] Flag `--sort`: `internal/commands/ls/ls.go` (unified flag)
- [x] Flag `--time`: `internal/commands/ls/ls.go` (atime, ctime, mtime)
- [x] Flag `--time-style`: `internal/commands/ls/ls.go`
- [x] Flag `--zero`: `internal/commands/ls/ls.go` (NUL terminated)
- [x] Flag `-1`: `internal/commands/ls/ls.go` (one-line)
- [x] Flag `-A`: `internal/commands/ls/ls.go` (almost-all)
- [x] Flag `-B`: `internal/commands/ls/ls.go` (ignore-backups)
- [x] Flag `-C`: `internal/commands/ls/ls.go` (vertical columns)
- [x] Flag `-D`, `--dired`: `internal/commands/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-F`: `internal/commands/ls/ls.go` (classify)
- [x] Flag `-G`: `internal/commands/ls/ls.go` (no-group)
- [x] Flag `-H`: `internal/commands/ls/ls.go` (dereference-command-line)
- [x] Flag `-I`: `internal/commands/ls/ls.go` (ignore)
- [x] Flag `-L`: `internal/commands/ls/ls.go` (dereference)
- [x] Flag `-N`, `--literal`: `internal/commands/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-Q`: `internal/commands/ls/ls.go` (quote-name)
- [x] Flag `-R`: `internal/commands/ls/ls.go` (recursive)
- [x] Flag `-S`: `internal/commands/ls/ls.go` (sort-size)
- [x] Flag `-T`: `internal/commands/ls/ls.go`
- [x] Flag `-U`: `internal/commands/ls/ls.go` (do not sort)
- [x] Flag `-X`: `internal/commands/ls/ls.go` (extension sort)
- [x] Flag `-Z`: `internal/commands/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-a`: `internal/commands/ls/ls.go` (all)
- [x] Flag `-b`: `internal/commands/ls/ls.go` (escape)
- [x] Flag `-c`: `internal/commands/ls/ls.go` (ctime)
- [x] Flag `-d`: `internal/commands/ls/ls.go` (directory itself)
- [x] Flag `-f`: `internal/commands/ls/ls.go` (do not sort, enable -aU)
- [x] Flag `-g`: `internal/commands/ls/ls.go` (like -l but no owner)
- [x] Flag `-h`: `internal/commands/ls/ls.go` (human-readable)
- [x] Flag `-i`: `internal/commands/ls/ls.go` (inode)
- [x] Flag `-l`: `internal/commands/ls/ls.go` (long)
- [x] Flag `-m`: `internal/commands/ls/ls.go` (comma)
- [x] Flag `-n`: `internal/commands/ls/ls.go` (numeric)
- [x] Flag `-o`: `internal/commands/ls/ls.go` (like -l but no group)
- [x] Flag `-p`: `internal/commands/ls/ls.go` (indicator)
- [x] Flag `-q`: `internal/commands/ls/ls.go` (hide-control-chars)
- [x] Flag `-r`: `internal/commands/ls/ls.go` (reverse)
- [x] Flag `-s`: `internal/commands/ls/ls.go` (size in blocks)
- [x] Flag `-t`: `internal/commands/ls/ls.go` (sort-time)
- [x] Flag `-u`: `internal/commands/ls/ls.go` (atime)
- [x] Flag `-v`: `internal/commands/ls/ls.go` (natural sort)
- [x] Flag `-w`: `internal/commands/ls/ls.go`
- [x] Flag `-x`: `internal/commands/ls/ls.go` (across/horizontal)

### `mapfile`

- [x] Upstream: `third_party/bash/builtins/mapfile.def`
- [x] Array population: Implemented in `internal/commands/mapfile/mapfile.go`
- [x] Flag `-d`: `internal/commands/mapfile/mapfile.go`
- [x] Flag `-t`: `internal/commands/mapfile/mapfile.go` (trim/strip newline)
- [x] Flag `-n`: `internal/commands/mapfile/mapfile.go` (count)
- [x] Flag `-O`: `internal/commands/mapfile/mapfile.go` (origin)
- [x] Flag `-u`: `internal/commands/mapfile/mapfile.go` (fd)
- [x] Flag `-C`: `internal/commands/mapfile/mapfile.go` (callback)
- [x] Flag `-c`: `internal/commands/mapfile/mapfile.go` (quantum)
- [x] Flag `-s`: `internal/commands/mapfile/mapfile.go`
- [x] Aliases: `readarray` (handled via command registration)

### `md5sum`

- [x] Upstream: `third_party/coreutils/src/cksum.c`
- [x] Inherits all `cksum` hash flags: `internal/commands/sum/sum.go`

### `mkdir`

- [x] Upstream: `third_party/coreutils/src/mkdir.c`
- [x] Basic operation: Implemented in `internal/commands/mkdir/mkdir.go`
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/mkdir/mkdir.go` (octal)
- [x] Flag `-p`, `--parents`: `internal/commands/mkdir/mkdir.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/mkdir/mkdir.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/mkdir/mkdir.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `mkfifo`

- [x] Upstream: `third_party/coreutils/src/mkfifo.c`
- [x] Basic operation: Implemented in `internal/commands/mkfifo/mkfifo.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#mkfifo-mknod))
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/mkfifo/mkfifo.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/mkfifo/mkfifo.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `mknod`

- [x] Upstream: `third_party/coreutils/src/mknod.c`
- [x] Basic operation: Implemented in `internal/commands/mknod/mknod.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#mkfifo-mknod))
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/mknod/mknod.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/mknod/mknod.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `mktemp`

- [x] Upstream: `third_party/coreutils/src/mktemp.c`
- [x] Basic creation: Implemented in `internal/commands/mktemp/mktemp.go`
- [x] Flag `--suffix`: `internal/commands/mktemp/mktemp.go` (partial via TEMPLATE)
- [x] Flag `-d`: `internal/commands/mktemp/mktemp.go`
- [x] Flag `-p`: `internal/commands/mktemp/mktemp.go` (via --tmpdir)
- [x] Flag `-q`: `internal/commands/mktemp/mktemp.go` (ignored)
- [x] Flag `-t`: `internal/commands/mktemp/mktemp.go`
- [x] Flag `-u`: `internal/commands/mktemp/mktemp.go`

### `mv`

- [x] Basic move/rename: Implemented in `internal/commands/mv/mv.go`
- [x] Flag `-b`, `--backup`: `internal/commands/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-f`, `--force`: `internal/commands/mv/mv.go` (ignored)
- [x] Flag `-i`, `--interactive`: `internal/commands/mv/mv.go` (ignored)
- [x] Flag `-n`, `--no-clobber`: `internal/commands/mv/mv.go`
- [x] Flag `-t`, `--target-directory`: `internal/commands/mv/mv.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/mv/mv.go`
- [x] Flag `-u`, `--update`: `internal/commands/mv/mv.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/mv/mv.go`
- [x] Flag `-Z`, `--context`: `internal/commands/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--exchange`: `internal/commands/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--no-copy`: `internal/commands/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `nice`

- [x] Upstream: `third_party/coreutils/src/nice.c`
- [x] Priority adjustment: Implemented (as wrapper) in `internal/commands/nice/nice.go`
- [x] Flag `-n`, `--adjustment=N`: `internal/commands/nice/nice.go`

### `nl`

- [x] Upstream: `third_party/coreutils/src/nl.c`
- [x] Basic numbering: Implemented in `internal/commands/nl/nl.go`
- [x] Flag `-b`, `--body-numbering=STYLE`: `internal/commands/nl/nl.go`
- [x] Flag `-d`, `--section-delimiter=CC`: `internal/commands/nl/nl.go`
- [x] Flag `-f`, `--footer-numbering=STYLE`: `internal/commands/nl/nl.go`
- [x] Flag `-h`, `--header-numbering=STYLE`: `internal/commands/nl/nl.go`
- [x] Flag `-i`, `--line-increment=NUMBER`: `internal/commands/nl/nl.go`
- [x] Flag `-l`, `--join-blank-lines=NUMBER`: `internal/commands/nl/nl.go`
- [x] Flag `-n`, `--number-format=FORMAT`: `internal/commands/nl/nl.go`
- [x] Flag `-p`, `--no-renumber`: `internal/commands/nl/nl.go`
- [x] Flag `-s`, `--number-separator=STRING`: `internal/commands/nl/nl.go`
- [x] Flag `-v`, `--starting-line-number=NUMBER`: `internal/commands/nl/nl.go`
- [x] Flag `-w`, `--number-width=NUMBER`: `internal/commands/nl/nl.go`

### `nohup`

- [x] Upstream: `third_party/coreutils/src/nohup.c`
- [x] Flag `--help`: `internal/commands/nohup/nohup.go`
- [x] Flag `--version`: `internal/commands/nohup/nohup.go`

### `nproc`

- [x] Upstream: `third_party/coreutils/src/nproc.c`
- [x] Basic nproc: Implemented in `internal/commands/nproc/nproc.go`
- [x] Flag `--all`: `internal/commands/nproc/nproc.go`
- [x] Flag `--ignore`: `internal/commands/nproc/nproc.go`

### `numfmt`

- [x] Upstream: `third_party/coreutils/src/numfmt.c`
- [x] Conversion: Implemented in `internal/commands/numfmt/numfmt.go`
- [x] Flag `-d`, `--delimiter=X`: `internal/commands/numfmt/numfmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/numfmt/numfmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--to`: `internal/commands/numfmt/numfmt.go`
- [x] Flag `--from`: `internal/commands/numfmt/numfmt.go`
- [x] Flag `--header`: `internal/commands/numfmt/numfmt.go`

### `od`

- [x] Upstream: `third_party/coreutils/src/od.c`
- [x] Format output: Implemented in `internal/commands/od/od.go`
- [x] Flag `-A rad`: `internal/commands/od/od.go`
- [x] Flag `-j bytes`: `internal/commands/od/od.go`
- [x] Flag `-N bytes`: `internal/commands/od/od.go`
- [x] Flag `-t type`: `internal/commands/od/od.go` (Stub; only 2-byte octal supported)
- [x] Flag `-v`: `internal/commands/od/od.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`: `internal/commands/od/od.go`
- [x] Flag `-S`: `internal/commands/od/od.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `paste`

- [x] Upstream: `third_party/coreutils/src/paste.c`
- [x] Basic paste: Implemented in `internal/commands/paste/paste.go`
- [x] Flag `-d`, `--delimiters=LIST`: `internal/commands/paste/paste.go`
- [x] Flag `-s`, `--serial`: `internal/commands/paste/paste.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/paste/paste.go`

### `pathchk`

- [x] Upstream: `third_party/coreutils/src/pathchk.c`
- [x] Basic operation: Implemented in `internal/commands/pathchk/pathchk.go`
- [x] Flag `-p`: `internal/commands/pathchk/pathchk.go`
- [x] Flag `-P`: `internal/commands/pathchk/pathchk.go`

### `pinky`

- [x] Upstream: `third_party/coreutils/src/pinky.c`
- [x] Basic operation: Implemented in `internal/commands/pinky/pinky.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#pinky))
- [x] Flag `-b`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-f`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-h`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-i`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-l`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-p`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-q`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-s`: `internal/commands/pinky/pinky.go` (Ignored)
- [x] Flag `-w`: `internal/commands/pinky/pinky.go` (Ignored)

### `popd`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic popping: Implemented in `internal/commands/popd/popd.go`
- [x] Flag `-n`: `internal/commands/popd/popd.go`

### `pr`

- [x] Upstream: `third_party/coreutils/src/pr.c`
- [x] Print formatting: Implemented in `internal/commands/pr/pr.go`
- [x] Flag `-a`: `internal/commands/pr/pr.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-d`: `internal/commands/pr/pr.go`
- [x] Flag `-h`: `internal/commands/pr/pr.go`
- [x] Flag `-l`: `internal/commands/pr/pr.go`
- [x] Flag `-m`: `internal/commands/pr/pr.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-n`: `internal/commands/pr/pr.go`
- [x] Flag `-t`: `internal/commands/pr/pr.go`
- [x] Flag `-w`: `internal/commands/pr/pr.go`

### `printenv`

- [x] Upstream: `third_party/coreutils/src/printenv.c`
- [x] Basic output: Implemented in `internal/commands/printenv/printenv.go`
- [x] Flag `-0`: `internal/commands/printenv/printenv.go`

### `printf`

- [x] Basic formatting: Implemented in `internal/commands/printf/printf.go`
- [x] Flag `%b`: `internal/commands/printf/printf.go`
- [x] Flag `%q`: `internal/commands/printf/printf.go`
- [x] Flag `-v VAR`: Implemented in `internal/commands/printf/printf.go`

### `ptx`

- [x] Upstream: `third_party/coreutils/src/ptx.c`
- [x] Permuted Index: `internal/commands/ptx/ptx.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#ptx))
- [x] Flag `-A`, `--auto-reference`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-F`, `--flag-truncation=STRING`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-G`, `--gnu-extensions`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-M`, `--macro-name=STRING`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-O`, `--format=roff`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-R`, `--right-side-refs`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-S`, `--sentence-regexp=REGEXP`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-T`, `--format=tex`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-W`, `--word-regexp=REGEXP`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-b`, `--break-file=FILE`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-f`, `--ignore-case`: `internal/commands/ptx/ptx.go` (Ignored)
- [x] Flag `-g`, `--gap-size=NUMBER`: `internal/commands/ptx/ptx.go`
- [x] Flag `-i`, `--ignore-file=FILE`: `internal/commands/ptx/ptx.go`
- [x] Flag `-o`, `--only-file=FILE`: `internal/commands/ptx/ptx.go`
- [x] Flag `-r`, `--references`: `internal/commands/ptx/ptx.go`
- [x] Flag `-t`, `--typeset-mode`: `internal/commands/ptx/ptx.go`
- [x] Flag `-w`, `--width=NUMBER`: `internal/commands/ptx/ptx.go`

### `pushd`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic pushing: Implemented in `internal/commands/pushd/pushd.go`
- [x] Flag `-n`: `internal/commands/pushd/pushd.go`

### `pwd`

- [x] Upstream: `third_party/bash/builtins/cd.def`
- [x] Basic path reporting: Implemented in `internal/commands/pwd/pwd.go`
- [-] Flag `--help`: Handled by the shell's global help dispatcher. (See [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#pwd))
- [x] Flag `-L`: `internal/commands/pwd/pwd.go`
- [x] Flag `-P`: `internal/commands/pwd/pwd.go`

### `read`

- [x] Upstream: `third_party/bash/builtins/read.def`
- [x] Basic input: Implemented in `internal/commands/read/read.go`
- [x] Basic operation: Implemented in `internal/commands/read/read.go`
- [x] Flag `-p PROMPT`: `internal/commands/read/read.go`
- [x] Flag `-r` (raw): `internal/commands/read/read.go`
- [x] Flag `-d DELIM`: `internal/commands/read/read.go`
- [x] Flag `-n NCHARS`: `internal/commands/read/read.go`
- [x] Flag `-N NCHARS`: `internal/commands/read/read.go`
- [x] Flag `-a ARRAY`: `internal/commands/read/read.go`
- [x] Flag `-s`: `internal/commands/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#read))
- [x] Flag `-t TIMEOUT`: `internal/commands/read/read.go`
- [x] Flag `-u FD`: `internal/commands/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#read))
- [x] Flag `-e`: `internal/commands/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#read))
- [x] Flag `-i TEXT`: `internal/commands/read/read.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#read))

### `readlink`

- [x] Upstream: `third_party/coreutils/src/readlink.c`
- [x] Basic output: Implemented in `internal/commands/readlink/readlink.go`
- [x] Flag `-e`, `--canonicalize-existing`: `internal/commands/readlink/readlink.go`
- [x] Flag `-m`, `--canonicalize-missing`: `internal/commands/readlink/readlink.go`
- [x] Flag `-q`, `--quiet`: `internal/commands/readlink/readlink.go`
- [x] Flag `-s`, `--silent`: `internal/commands/readlink/readlink.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/readlink/readlink.go`
- [x] Flag `-z`, `--zero`: `internal/commands/readlink/readlink.go`
- [x] Flag `-f`: `internal/commands/readlink/readlink.go`
- [x] Flag `-n`: `internal/commands/readlink/readlink.go`

### `readonly`

- [x] Upstream: `third_party/bash/builtins/setattr.def`
- [x] Attribute management: Implemented in `internal/commands/readonly/readonly.go`
- [x] Flag `-a`: `internal/commands/readonly/readonly.go` (indexed array)
- [x] Flag `-A`: `internal/commands/readonly/readonly.go` (associative array)
- [x] Flag `-p`: `internal/commands/readonly/readonly.go` (print)
- [x] Flag `-f`: `internal/commands/readonly/readonly.go` (functions)

### `realpath`

- [x] Upstream: `third_party/coreutils/src/realpath.c`
- [x] Basic functionality: Implemented in `internal/commands/realpath/realpath.go`
- [x] Basic output: Implemented in `internal/commands/realpath/realpath.go`
- [x] Flag `-E`, `--canonicalize-existing`: `internal/commands/realpath/realpath.go`
- [x] Flag `-L`, `--logical`: `internal/commands/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-P`, `--physical`: `internal/commands/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-q`, `--quiet`: `internal/commands/realpath/realpath.go`
- [x] Flag `-s`, `--strip`: `internal/commands/realpath/realpath.go`
- [x] Flag `-z`, `--zero`: `internal/commands/realpath/realpath.go`
- [x] Flag `--relative-to`: `internal/commands/realpath/realpath.go`
- [x] Flag `--relative-base`: `internal/commands/realpath/realpath.go`
- [x] Flag `-e`: `internal/commands/realpath/realpath.go`
- [x] Flag `-m`: `internal/commands/realpath/realpath.go`

### `return`

- [x] Upstream: `third_party/bash/builtins/return.def`
- [x] Basic return: Implemented in `internal/commands/returncmd/return.go`
- [x] Exit status parameter: `internal/commands/returncmd/return.go`

### `rm`

- [x] Upstream: `third_party/coreutils/src/rm.c`
- [x] Basic removal: Implemented in `internal/commands/rm/rm.go`
- [x] Flag `-d`, `--dir`: `internal/commands/rm/rm.go`
- [x] Flag `-f`: `internal/commands/rm/rm.go`
- [x] Flag `-i`: `internal/commands/rm/rm.go`
- [x] Flag `-I`: `internal/commands/rm/rm.go`
- [x] Flag `-r`, `-R`, `--recursive`: `internal/commands/rm/rm.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/rm/rm.go`
- [x] Flag `--one-file-system`: `internal/commands/rm/rm.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `rmdir`

- [x] Upstream: `third_party/coreutils/src/rmdir.c`
- [x] Basic rmdir: Implemented in `internal/commands/rmdir/rmdir.go`
- [x] Basic removal: Implemented in `internal/commands/rmdir/rmdir.go`
- [x] Flag `--ignore-fail-on-non-empty`: `internal/commands/rmdir/rmdir.go`
- [x] Flag `-p`: `internal/commands/rmdir/rmdir.go`
- [x] Flag `-v`: `internal/commands/rmdir/rmdir.go`

### `runcon`

- [x] Upstream: `third_party/coreutils/src/runcon.c`
- [x] Basic operation: `internal/commands/runcon/runcon.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chcon-runcon))
- [x] Flag `-c`, `--compute`: `internal/commands/runcon/runcon.go`
- [x] Flag `-l`, `--user=USER`: `internal/commands/runcon/runcon.go`
- [x] Flag `-r`, `--role=ROLE`: `internal/commands/runcon/runcon.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/runcon/runcon.go`
- [x] Flag `-u`, `--user=USER`: `internal/commands/runcon/runcon.go`

### `select`

- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Basic operation: Implemented in `internal/shell/shell.go`

### `seq`

- [x] Upstream: `third_party/coreutils/src/seq.c`
- [x] Basic sequence: Implemented in `internal/commands/seq/seq.go`
- [x] Flag `-f`, `--format=FORMAT`: `internal/commands/seq/seq.go`
- [x] Flag `-s`, `--separator=STRING`: `internal/commands/seq/seq.go`
- [x] Flag `-w`, `--equal-width`: `internal/commands/seq/seq.go`

### `set`

- [ ] Upstream: `third_party/bash/builtins/set.def`
- [x] Option management (-e, -u, -x, -o): Implemented in `internal/commands/set/set.go`
- [x] Positional parameters: Stub in `internal/commands/set/set.go`
- [x] Flag `-a`: `internal/commands/set/set.go` (allexport)
- [x] Flag `-b`: `internal/commands/set/set.go` (notify)
- [x] Flag `-e`: `internal/commands/set/set.go` (errexit)
- [x] Flag `-f`: `internal/commands/set/set.go` (noglob)
- [x] Flag `-h`: `internal/commands/set/set.go` (hashall)
- [x] Flag `-k`: `internal/commands/set/set.go` (keyword)
- [x] Flag `-m`: `internal/commands/set/set.go` (monitor)
- [x] Flag `-n`: `internal/commands/set/set.go` (noexec)
- [x] Flag `-o`: `internal/commands/set/set.go` (option-name)
- [x] Flag `-p`: `internal/commands/set/set.go` (privileged)
- [x] Flag `-t`: `internal/commands/set/set.go` (exit after one command)
- [x] Flag `-u`: `internal/commands/set/set.go` (nounset)
- [x] Flag `-v`: `internal/commands/set/set.go` (verbose)
- [x] Flag `-x`: `internal/commands/set/set.go` (xtrace)
- [x] Flag `-B`: `internal/commands/set/set.go` (braceexpand)
- [x] Flag `-C`: `internal/commands/set/set.go` (noclobber)
- [x] Flag `-E`: `internal/commands/set/set.go` (errtrace)
- [x] Flag `-H`: `internal/commands/set/set.go` (histexpand)
- [x] Flag `-P`: `internal/commands/set/set.go` (physical)
- [x] Flag `-T`: `internal/commands/set/set.go` (functrace)

### `sha1sum`

### `sha1sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha1sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `sha224sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha224sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `sha256sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha256sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `sha384sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha384sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `sha512sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha512sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `shift`

- [x] Upstream: `third_party/bash/builtins/shift.def`
- [x] Basic shift: Implemented in `internal/commands/shift/shift.go`
- [x] Shifting n parameters: `internal/commands/shift/shift.go`

### `shopt`

- [x] Upstream: `third_party/bash/builtins/shopt.def`
- [x] Basic option management: Implemented in `internal/commands/shopt/shopt.go`
- [x] Flag `-o`: `internal/commands/shopt/shopt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`: `third_party/bash/builtins/shopt.def:L77` (print status)
- [x] Flag `-q`: `third_party/bash/builtins/shopt.def:L71` (quiet)
- [x] Flag `-s`: `third_party/bash/builtins/shopt.def:L62` (enable)
- [x] Flag `-u`: `third_party/bash/builtins/shopt.def:L65` (disable)

### `shred`

- [x] Upstream: `third_party/coreutils/src/shred.c`
- [x] Data erasure: Implemented in `internal/commands/shred/shred.go` (Workaround; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#shred))
- [x] Flag `-f`, `--force`: `internal/commands/shred/shred.go`
- [x] Flag `-n`, `--iterations=N`: `internal/commands/shred/shred.go`
- [x] Flag `-s`, `--size=N`: `internal/commands/shred/shred.go` (partial via iteration)
- [x] Flag `-u`, `--remove`: `internal/commands/shred/shred.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/shred/shred.go`
- [x] Flag `-x`, `--exact`: `internal/commands/shred/shred.go`
- [x] Flag `-z`, `--zero`: `internal/commands/shred/shred.go`

### `shuf`

- [x] Upstream: `third_party/coreutils/src/shuf.c`
- [x] Basic operation: Implemented in `internal/commands/shuf/shuf.go`
- [x] Flag `-e`, `--echo`: `internal/commands/shuf/shuf.go`
- [x] Flag `-i`, `--input-range=LO-HI`: `internal/commands/shuf/shuf.go`
- [x] Flag `-n`, `--head-count=COUNT`: `internal/commands/shuf/shuf.go`
- [x] Flag `-o`, `--output=FILE`: `internal/commands/shuf/shuf.go`
- [x] Flag `-r`, `--repeat`: `internal/commands/shuf/shuf.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/shuf/shuf.go`
- [x] Flag `--random-source=FILE`: `internal/commands/shuf/shuf.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#shuf))

### `sleep`

- [x] Upstream: `third_party/coreutils/src/sleep.c`
- [x] Basic sleep: Implemented in `internal/commands/sleep/sleep.go`
- [x] Multiple arguments (sum): `internal/commands/sleep/sleep.go`
- [x] Suffixes (s, m, h, d): `internal/commands/sleep/sleep.go`
- [x] Flag `--help`: `internal/commands/sleep/sleep.go`
- [x] Flag `--version`: `internal/commands/sleep/sleep.go`

### `sort`

- [x] Upstream: `third_party/coreutils/src/sort.c`
- [x] Basic sorting: Implemented in `internal/commands/sort/sort.go`
- [x] Ordering flags (`-b`, `-i`, `-d`, `-f`, `-g`, `-h`, `-n`, `-M`, `-R`, `-V`, `-r`): Implemented in `internal/commands/sort/sort.go`
- [x] Flag `-b`, `--ignore-leading-blanks`: `internal/commands/sort/sort.go`
- [x] Flag `-c`, `-C`, `--check`: `internal/commands/sort/sort.go`
- [x] Flag `-d`, `--dictionary-order`: `internal/commands/sort/sort.go`
- [x] Flag `-f`, `--ignore-case`: `internal/commands/sort/sort.go`
- [x] Flag `-g`, `--general-numeric-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-h`, `--human-numeric-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-i`, `--ignore-nonprinting`: `internal/commands/sort/sort.go`
- [x] Flag `-k`, `--key=KEYDEF`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-m`, `--merge`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-M`, `--month-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-n`, `--numeric-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-o`, `--output=FILE`: `internal/commands/sort/sort.go`
- [x] Flag `-r`, `--reverse`: `internal/commands/sort/sort.go`
- [x] Flag `-s`, `--stable`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-S`, `--buffer-size=SIZE`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--field-separator=SEP`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-T`, `--temporary-directory=DIR`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--unique`: `internal/commands/sort/sort.go`
- [x] Flag `-V`, `--version-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/sort/sort.go`
- [x] Flag `--parallel=N`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--random-sort` (`-R`): `internal/commands/sort/sort.go`
- [x] Flag `--debug`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--files0-from=F`: `internal/commands/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `source`

- [x] Upstream: `third_party/bash/builtins/source.def`
- [x] Basic sourcing: Implemented in `internal/commands/source/source.go`
- [x] Aliases: `.`

### `split`

- [x] Basic split: Implemented in `internal/commands/split/split.go`
- [x] Flag `--filter=COMMAND`: `internal/commands/split/split.go` (Stub)
- [x] Flag `--verbose`: `internal/commands/split/split.go`
- [x] Flag `-C`, `--line-bytes=SIZE`: `internal/commands/split/split.go` (Stub)
- [x] Flag `-a`: `internal/commands/split/split.go`
- [x] Flag `-b`: `internal/commands/split/split.go`
- [x] Flag `-d`: `internal/commands/split/split.go`
- [x] Flag `-e`, `--elide-empty-files`: `internal/commands/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/split/split.go`
- [x] Flag `-n`: `internal/commands/split/split.go`
- [x] Flag `-t`: `internal/commands/split/split.go`
- [x] Flag `-u`, `--unbuffered`: `internal/commands/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-x`, `--hex-suffixes`: `internal/commands/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `stat`

- [x] Upstream: `third_party/coreutils/src/stat.c`
- [x] Basic output: Implemented in `internal/commands/stat/stat.go`
- [x] Flag `-c`, `--format=FORMAT`: `internal/commands/stat/stat.go`
- [x] Flag `-f`, `--file-system`: `internal/commands/stat/stat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-L`, `--dereference`: `internal/commands/stat/stat.go`
- [x] Flag `-t`, `--terse`: `internal/commands/stat/stat.go`
- [x] Flag `--printf=FORMAT`: `internal/commands/stat/stat.go` (Stub)
- [x] Flag `--cached={always,never,default}`: `internal/commands/stat/stat.go` (Stub)

### `stdbuf`

- [x] Upstream: `third_party/coreutils/src/stdbuf.c`
- [x] Stream Buffering: `internal/commands/stdbuf/stdbuf.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#stdbuf))
- [x] Flag `-e`, `--error=MODE`: `internal/commands/stdbuf/stdbuf.go` (Ignored)
- [x] Flag `-i`, `--input=MODE`: `internal/commands/stdbuf/stdbuf.go` (Ignored)
- [x] Flag `-o`, `--output=MODE`: `internal/commands/stdbuf/stdbuf.go` (Ignored)

### `stty`

- [x] Upstream: `third_party/coreutils/src/stty.c`
- [x] TTY Configuration: `internal/commands/stty/stty.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#stty))
- [x] Flag `-F`, `--file=DEVICE`: `internal/commands/stty/stty.go` (Ignored)
- [x] Flag `-a`, `--all`: `internal/commands/stty/stty.go` (Partial output)
- [x] Flag `-g`, `--save`: `internal/commands/stty/stty.go` (Ignored)

### `sum`

- [x] Upstream: `third_party/coreutils/src/sum.c`
- [x] Basic checksum: Implemented in `internal/commands/sumlegacy/sum.go`
- [x] Flag `-r`: `internal/commands/sumlegacy/sum.go` (BSD algorithm)
- [x] Flag `-s`, `--sysv`: `internal/commands/sumlegacy/sum.go` (System V algorithm)

### `suspend`

- [x] Upstream: `third_party/bash/builtins/suspend.def`
- [x] Basic operation: `internal/commands/suspend/suspend.go` (Unsupported; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#suspend))
- [x] Flag `-f`: `internal/commands/suspend/suspend.go`

### `sync`

- [x] Upstream: `third_party/coreutils/src/sync.c`
- [x] Basic operation: Implemented in `internal/commands/sync/sync.go`
- [x] Flag `-d`, `--data`: `internal/commands/sync/sync.go` (no-op)
- [x] Flag `-f`, `--file-system`: `internal/commands/sync/sync.go` (no-op)

### `tac`

- [x] Upstream: `third_party/coreutils/src/tac.c`
- [x] Basic output: Implemented in `internal/commands/tac/tac.go`
- [x] Flag `-b`: `internal/commands/tac/tac.go`
- [x] Flag `-r`, `--regex`: `internal/commands/tac/tac.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`: `internal/commands/tac/tac.go`

### `tail`

- [x] Basic output: Implemented in `internal/commands/tail/tail.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/tail/tail.go`
- [x] Flag `-f`, `--follow[={name|descriptor}]`: `internal/commands/tail/tail.go`
- [x] Flag `-F`: `internal/commands/tail/tail.go` (partial)
- [x] Flag `-n`, `--lines`: `internal/commands/tail/tail.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/tail/tail.go`
- [x] Flag `-s`, `--sleep-interval`: `internal/commands/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/tail/tail.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/tail/tail.go`
- [x] Flag `--pid`: `internal/commands/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--retry`: `internal/commands/tail/tail.go` (Stub)
- [x] Flag `--max-unchanged-stats`: `internal/commands/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `tee`

- [x] Basic copy: Implemented in `internal/commands/tee/tee.go`
- [x] Flag `-a`, `--append`: `internal/commands/tee/tee.go`
- [x] Flag `-i`, `--ignore-interrupts`: `internal/commands/tee/tee.go`
- [x] Flag `-p`, `--output-error[=MODE]`: `internal/commands/tee/tee.go`

### `test`

- [x] Unary operators (-e, -f, -d, etc.): Implemented in `internal/commands/test/test.go`
- [x] String operators (=, !=, -z, -n): Implemented in `internal/commands/test/test.go`
- [x] Numeric operators (-eq, -ne, etc.): Implemented in `internal/commands/test/test.go`
- [x] Logical operators (!, -a, -o): Implemented in `internal/commands/test/test.go`
- [x] Aliases: `[`

### `time`

- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Basic operation: Implemented in `internal/commands/time/time.go`

### `timeout`

- [x] Upstream: `third_party/coreutils/src/timeout.c`
- [x] Basic output: Implemented in `internal/commands/timeout/timeout.go`
- [x] Flag `--help`: `internal/commands/timeout/timeout.go`
- [x] Flag `--version`: `internal/commands/timeout/timeout.go`
- [x] Flag `-k`: `internal/commands/timeout/timeout.go`
- [x] Flag `-s`: `internal/commands/timeout/timeout.go`

### `times`

- [x] Upstream: `third_party/bash/builtins/times.def`
- [x] Basic output: Implemented in `internal/commands/times/times.go`

### `touch`

- [x] Upstream: `third_party/coreutils/src/touch.c`
- [x] Basic touch: Implemented in `internal/commands/touch/touch.go`
- [x] Basic timestamp update: Implemented in `internal/commands/touch/touch.go`
- [x] Flag `-t STAMP`: `internal/commands/touch/touch.go` (explicit timestamp)
- [x] Flag `-a`: `internal/commands/touch/touch.go`
- [x] Flag `-c`: `internal/commands/touch/touch.go`
- [x] Flag `-d`: `internal/commands/touch/touch.go`
- [x] Flag `-h`: `internal/commands/touch/touch.go`
- [x] Flag `-m`: `internal/commands/touch/touch.go`
- [x] Flag `-r`: `internal/commands/touch/touch.go`
- [x] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `internal/commands/touch/touch.go`

### `tr`

- [x] Upstream: `third_party/coreutils/src/tr.c`
- [x] Basic translation: Implemented in `internal/commands/tr/tr.go`
- [x] Flag `-c`, `-C`, `--complement`: `internal/commands/tr/tr.go`
- [x] Flag `-d`, `--delete`: `internal/commands/tr/tr.go`
- [x] Flag `-s`, `--squeeze-repeats`: `internal/commands/tr/tr.go`
- [x] Flag `-t`, `--truncate-set1`: `internal/commands/tr/tr.go`

### `trap`

- [x] Upstream: `third_party/bash/builtins/trap.def`
- [x] Basic trapping: Implemented in `internal/commands/trap/trap.go`
- [x] Flag `-P`: `internal/commands/trap/trap.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/trap/trap.go`
- [x] Flag `-p`: `internal/commands/trap/trap.go`

### `true`

- [x] Basic operation: Implemented in `internal/commands/boolcmd/bool.go`

### `truncate`

- [x] Upstream: `third_party/coreutils/src/truncate.c`
- [x] Basic truncation: Implemented in `internal/commands/truncate/truncate.go`
- [x] Flag `-c`: `internal/commands/truncate/truncate.go`
- [x] Flag `-o`: `internal/commands/truncate/truncate.go` (ignored stub)
- [x] Flag `-r`: `internal/commands/truncate/truncate.go`
- [x] Flag `-s`: `internal/commands/truncate/truncate.go`

### `tsort`

- [x] Upstream: `third_party/coreutils/src/tsort.c`
- [x] Topological sort: Implemented in `internal/commands/tsort/tsort.go`
- [x] Flag `--help`: `internal/commands/tsort/tsort.go`
- [x] Flag `--version`: `internal/commands/tsort/tsort.go`

### `tty`

- [x] Upstream: `third_party/coreutils/src/tty.c`
- [x] TTY reporting: Implemented in `internal/commands/tty/tty.go`
- [x] Flag `-s`, `--silent`, `--quiet`: `internal/commands/tty/tty.go`

### `type`

- [x] Upstream: `third_party/bash/builtins/type.def`
- [x] Command identification: Implemented in `internal/commands/type/type.go`
- [x] Flag `-a`: `internal/commands/type/type.go` (all occurrences)
- [x] Flag `-p`: `internal/commands/type/type.go` (path only)
- [x] Flag `-t`: `internal/commands/type/type.go` (type only)
- [x] Flag `-f`: `internal/commands/type/type.go` (skip functions)
- [x] Flag `-P`: `internal/commands/type/type.go` (force path search)

### `ulimit`

- [x] Upstream: `third_party/bash/builtins/ulimit.def`
- [x] Resource management: Implemented in `internal/commands/ulimit/ulimit.go` (Simulation)
- [x] Flag `-a`: `third_party/bash/builtins/ulimit.def:L35` (all)
- [x] Flag `-c`: `third_party/bash/builtins/ulimit.def:L37` (core)
- [x] Flag `-d`: `third_party/bash/builtins/ulimit.def:L38` (data)
- [x] Flag `-e`: `third_party/bash/builtins/ulimit.def:L39` (priority)
- [x] Flag `-f`: `third_party/bash/builtins/ulimit.def:L40` (file size)
- [x] Flag `-n`: `third_party/bash/builtins/ulimit.def:L45` (opened files)
- [x] Flag `-u`: `third_party/bash/builtins/ulimit.def:L51` (user processes)
- [x] Flag `-S`: `internal/commands/ulimit/ulimit.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-H`: `internal/commands/ulimit/ulimit.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `umask`

- [x] Upstream: `third_party/bash/builtins/umask.def`
- [x] Basic mask management: Implemented in `internal/commands/umask/umask.go`
- [x] Flag `-S`: `internal/commands/umask/umask.go`
- [x] Flag `-p`: `internal/commands/umask/umask.go`

### `unalias`

- [x] Upstream: `third_party/bash/builtins/alias.def`
- [x] Remove aliases: Implemented in `internal/commands/unalias/unalias.go`
- [x] Flag `-a`: `internal/commands/unalias/unalias.go` (remove all)

### `uname`

- [x] Upstream: `third_party/coreutils/src/uname.c`
- [x] Basic output: Implemented in `internal/commands/uname/uname.go`
- [x] Flag `-a`: `internal/commands/uname/uname.go`
- [x] Flag `-i`: `internal/commands/uname/uname.go`
- [x] Flag `-m`: `internal/commands/uname/uname.go`
- [x] Flag `-n`: `internal/commands/uname/uname.go`
- [x] Flag `-o`: `internal/commands/uname/uname.go`
- [x] Flag `-p`: `internal/commands/uname/uname.go`
- [x] Flag `-r`: `internal/commands/uname/uname.go`
- [x] Flag `-s`: `internal/commands/uname/uname.go`
- [x] Flag `-v`: `internal/commands/uname/uname.go`

### `unexpand`

- [x] Upstream: `third_party/coreutils/src/unexpand.c`
- [x] Basic conversion: Implemented in `internal/commands/unexpand/unexpand.go`
- [x] Flag `-a`, `--all`: `internal/commands/unexpand/unexpand.go`
- [x] Flag `-t`, `--tabs=LIST`: `internal/commands/unexpand/unexpand.go`
- [x] Flag `--first-only`: `internal/commands/unexpand/unexpand.go`

### `uniq`

- [x] Upstream: `third_party/coreutils/src/uniq.c`
- [x] Basic filtering: Implemented in `internal/commands/uniq/uniq.go`
- [x] Flag `-c`, `--count`: `internal/commands/uniq/uniq.go`
- [x] Flag `-d`, `--repeated`: `internal/commands/uniq/uniq.go`
- [x] Flag `-D`, `--all-repeated[=METHOD]`: `internal/commands/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-f`, `--skip-fields=N`: `internal/commands/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--ignore-case`: `internal/commands/uniq/uniq.go`
- [x] Flag `-s`, `--skip-chars=N`: `internal/commands/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--unique`: `internal/commands/uniq/uniq.go`
- [x] Flag `-w`, `--check-chars=N`: `internal/commands/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/uniq/uniq.go`
- [x] Flag `--group[=METHOD]`: `internal/commands/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `unlink`

- [x] Basic removal: Implemented in `internal/commands/unlink/unlink.go` (exactly 1 arg required)
- [x] Flag `--help`: `internal/commands/unlink/unlink.go`
- [x] Flag `--version`: `internal/commands/unlink/unlink.go`

### `unset`

- [x] Upstream: `third_party/bash/builtins/set.def`
- [x] Attribute management: Implemented in `internal/commands/unset/unset.go`
- [x] Flag `-f`: `internal/commands/unset/unset.go` (functions)
- [x] Flag `-v`: `internal/commands/unset/unset.go` (variables)
- [x] Flag `-n`: `internal/commands/declare/declare.go` (nameref - Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))

### `uptime`

- [x] Upstream: `third_party/coreutils/src/uptime.c`
- [x] Basic output: Implemented in `internal/commands/uptime/uptime.go`
- [x] Flag `-p`, `--pretty`: `internal/commands/uptime/uptime.go`
- [x] Flag `-s`, `--since`: `internal/commands/uptime/uptime.go`

### `users`

- [x] Upstream: `third_party/coreutils/src/users.c`
- [x] Basic output: Implemented in `internal/commands/users/users.go`
- [x] Flag `--help`: `internal/commands/users/users.go`
- [x] Flag `--version`: `internal/commands/users/users.go`

### `vdir`

- [x] Upstream: `third_party/coreutils/src/ls.c`
- [x] Inherits flags from `ls`

### `wait`

- [x] Upstream: `third_party/bash/builtins/wait.def`
- [x] Basic waiting: Implemented in `internal/commands/wait/wait.go`
- [x] Optional: jobspec or process ID: `internal/commands/wait/wait.go`
- [x] Flag `-f`: `internal/commands/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#wait))
- [x] Flag `-n`: `internal/commands/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#wait))
- [x] Flag `-p var`: `internal/commands/wait/wait.go`

### `wc`

- [x] Basic counts: Implemented in `internal/commands/wc/wc.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/wc/wc.go`
- [x] Flag `-m`, `--chars`: `internal/commands/wc/wc.go`
- [x] Flag `-l`, `--lines`: `internal/commands/wc/wc.go`
- [x] Flag `-w`, `--words`: `internal/commands/wc/wc.go`
- [x] Flag `-L`, `--max-line-length`: `internal/commands/wc/wc.go`
- [x] Flag `--files0-from=F`: `internal/commands/wc/wc.go`
- [x] Flag `--total={auto,always,only,never}`: `internal/commands/wc/wc.go`

### `who`

- [x] Upstream: `third_party/coreutils/src/who.c`
- [x] Basic output: Implemented in `internal/commands/who/who.go`
- [x] Flag `-H`, `--heading`: `internal/commands/who/who.go`
- [x] Flag `-a`, `--all`: `internal/commands/who/who.go`
- [x] Flag `-b`, `--boot`: `internal/commands/who/who.go`
- [x] Flag `-d`, `--dead`: `internal/commands/who/who.go`
- [x] Flag `-l`, `--login`: `internal/commands/who/who.go`
- [x] Flag `-m`: `internal/commands/who/who.go` (current user)
- [x] Flag `-p`, `--process`: `internal/commands/who/who.go`
- [x] Flag `-q`, `--count`: `internal/commands/who/who.go`
- [x] Flag `-r`, `--runlevel`: `internal/commands/who/who.go`
- [x] Flag `-s`, `--short`: `internal/commands/who/who.go`
- [x] Flag `-t`, `--time`: `internal/commands/who/who.go`
- [x] Flag `-u`, `--users`: `internal/commands/who/who.go`

### `whoami`

- [x] Upstream: `third_party/coreutils/src/whoami.c`
- [x] Basic output: Implemented in `internal/commands/whoami/whoami.go`
- [x] Flag `--help`: `internal/commands/whoami/whoami.go`
- [x] Flag `--version`: `internal/commands/whoami/whoami.go`

### `yes`

- [x] Upstream: `third_party/coreutils/src/yes.c`
- [x] Basic operation: Implemented in `internal/commands/yes/yes.go`
- [x] Basic repetition: Implemented in `internal/commands/yes/yes.go`
- [x] Flag `--help`: `internal/commands/yes/yes.go`
- [x] Flag `--version`: `internal/commands/yes/yes.go`


## Shell Keywords & Grammar

### `!`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pipeline negation: Implemented in `internal/shell/shell.go`

### `[[`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Conditional expressions: Implemented in `internal/shell/shell.go`
- [x] Pattern matching (`==`, `!=`): Implemented in `internal/shell/shell.go`
- [x] Regex matching (`=~`): Implemented in `internal/shell/shell.go`
- [ ] Aliases: `]]`

### `((`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Arithmetic evaluation: Implemented in `internal/shell/shell.go`
- [ ] Aliases: `))`

### `{`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Command grouping: Implemented in `internal/shell/shell.go`
- [ ] Aliases: `}`

### `case`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Pattern-based branching: Implemented in `internal/shell/shell.go`

### `coproc`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Asynchronous coprocesses: Missing implementation

### `for`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] List-based iteration: Missing implementation
- [x] C-style arithmetic iteration (`for ((`): Implemented in `internal/shell/shell.go`

### `function`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Shell function definition: Missing implementation

### `if`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Conditional branching (if/then/elif/else/fi): Missing implementation

### `until`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Negative condition looping: Implemented in `internal/shell/shell.go`

### `while`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Positive condition looping: Missing implementation
- [x] Sequential list `;`: Implemented in `internal/shell/shell.go`

## Shell Variables

### `BASH_VERSION`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Version information string: Implemented in `internal/app/app.go`

### `CDPATH`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Search path for `cd` command: Implemented in `internal/commands/cd/cd.go`

### `GLOBIGNORE`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pattern-based pathname expansion ignore: Missing implementation

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
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] EOF handling for interactive shells: Missing implementation

### `MAILCHECK`, `MAILPATH`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Mail notification settings: Missing implementation

### `PATH`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Command search path: Initialized in `internal/app/app.go`

### `PROMPT_COMMAND`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pre-prompt execution hook: Missing implementation

### `PS1`, `PS2`, `PS3`, `PS4`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Interactive prompt formatting: Initialized in `internal/app/app.go`

### `PWD`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Logical current directory tracking: Implemented in `internal/commands/cd/cd.go`

### `SHELLOPTS`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] List of enabled shell options: Missing implementation

### `TERM`
- [x] Upstream: `third_party/bash/builtins/reserved.def`
- [x] Terminal environment identification: Initialized in `internal/app/app.go`

### `TIMEFORMAT`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Output format for `time` reserved word: Missing implementation

## Interactive Shell Features

- [ ] Interactive history navigation (Up/Down arrow keys)
- [ ] Command line editing (Backspace, Ctrl+L, etc.)
- [ ] Tab completion: Missing implementation

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
- [x] IFS-based splitting in `read`: Implemented in `internal/commands/read/read.go`

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
- [ ] Moving input `[n]<&digit-`: `third_party/bash/redir.c:L1117`
- [ ] Moving output `[n]>&digit-`: `third_party/bash/redir.c:L1118`

### Advanced Redirections
- [ ] Here-Documents `[n]<<[-]word`: `third_party/bash/redir.c:L1042`
- [x] Here-Strings `[n]<<<word`: Implemented in `internal/shell/shell.go`
- [ ] Process Substitution `<(list)`, `>(list)`: `third_party/bash/subst.c:L321`

## Globbing Patterns

### Standard Wildcards
- [x] Match any string `*`: Implemented in `internal/shell/shell.go`
- [x] Match any character `?`: Implemented in `internal/shell/shell.go`

### Character Classes
- [ ] Match set of characters `[...]`: `third_party/bash/lib/glob/smatch.c`
- [ ] Negative match set `[!...]`, `[^...]`: `third_party/bash/lib/glob/smatch.c`

### Extended Globbing (extglob)
- [ ] Option `?(list)` (zero or one): `third_party/bash/lib/glob/smatch.c`
- [ ] Option `*(list)` (zero or more): `third_party/bash/lib/glob/smatch.c`
- [ ] Option `+(list)` (one or more): `third_party/bash/lib/glob/smatch.c`
- [ ] Option `@(list)` (exactly one): `third_party/bash/lib/glob/smatch.c`
- [ ] Option `!(list)` (anything but): `third_party/bash/lib/glob/smatch.c`

## Execution Flow

### Pipelines
- [x] Basic pipe `|`: Implemented in `internal/shell/shell.go`
- [x] Combined stderr pipe `|&`: Implemented in `internal/shell/shell.go`

### Compound Commands & Lists
- [x] Sequential list `;`: Implemented in `internal/shell/shell.go`
- [ ] Background execution `&`: `third_party/bash/execute_cmd.c:L193`
- [ ] Logical AND `&&`: `third_party/bash/execute_cmd.c:L193`
- [ ] Logical OR `||`: `third_party/bash/execute_cmd.c:L193`
- [ ] Subshell execution `( list )`: `third_party/bash/execute_cmd.c:L185`

## Signal & Trap Handling

### Core Signal Handling
- [ ] Trap initialization: `third_party/bash/trap.c:L154`
- [ ] Signal decoding (names/numbers): `third_party/bash/trap.c:L236`
- [ ] Pending trap execution: `third_party/bash/trap.c:L328`

### Subshell & Inheritance
- [ ] Signal inheritance rules: `third_party/bash/trap.c:L568`
- [ ] Trap reset in subshells: `third_party/bash/trap.c:L447`

## Advanced Shell Features

### Alias Expansion
- [ ] Initialization: `initialize_aliases` -> `third_party/bash/alias.c:L71`
- [ ] Expansion Logic (Recursive): `alias_expand` -> `third_party/bash/alias.c:L465`
- [ ] Tokenization for Aliases: `rd_token` -> `third_party/bash/alias.c:L425`
- [ ] Whitespace handling: `skipws` -> `third_party/bash/alias.c:L339`

### Array Support
- [ ] **Indexed Arrays**: Doubly-linked list implementation -> `third_party/bash/array.c`
    - [ ] `array_insert`: `third_party/bash/array.c:L516`
    - [ ] `array_reference`: `third_party/bash/array.c:L657`
    - [ ] Subrange expansion `${a[@]:s:n}`: `third_party/bash/array.c:L377`
- [ ] **Associative Arrays**: Hash table implementation -> `third_party/bash/assoc.c`
    - [ ] `assoc_insert`: `third_party/bash/assoc.c:L68`
    - [ ] `assoc_reference`: `third_party/bash/assoc.c:L120`

### Programmable Completion
- [ ] **Core Logic**: `gen_progcomp_completions` -> `third_party/bash/pcomplete.c:L127`
- [ ] **Builtin Integration**: `compgen`, `complete` logic -> `third_party/bash/pcomplete.c:L142`
- [ ] Item Generators (Aliases, Jobs, etc.): `third_party/bash/pcomplete.c:L155-178`
