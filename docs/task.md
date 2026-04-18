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
- [ ] Flag `--base16`: `third_party/coreutils/src/basenc.c:L86`
- [ ] Flag `--base2lsbf`: `third_party/coreutils/src/basenc.c:L88`
- [ ] Flag `--base2msbf`: `third_party/coreutils/src/basenc.c:L87`
- [x] Flag `--base32`: `internal/commands/base32/base32.go`
- [ ] Flag `--base32hex`: `third_party/coreutils/src/basenc.c:L85`
- [ ] Flag `--base58`: `third_party/coreutils/src/basenc.c:L83`
- [x] Flag `--base64`: `internal/commands/base64/base64.go`
- [ ] Flag `--base64url`: `third_party/coreutils/src/basenc.c:L82`
- [ ] Flag `--z85`: `third_party/coreutils/src/basenc.c:L89`
- [x] Flag `-d`: `internal/commands/base32/base32.go`
- [x] Flag `-i`: `internal/commands/base32/base32.go`
- [x] Flag `-w`: `internal/commands/base32/base32.go`

### `basename`

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
- [ ] Flag `--z85`: `third_party/coreutils/src/basenc.c:L145`

### `bg`

- [ ] Upstream: `third_party/bash/builtins/fg_bg.def`
- [ ] Basic job management: Missing implementation
- [ ] Job specification support: `third_party/bash/builtins/fg_bg.def:L65`

### `bind`

- [ ] Upstream: `third_party/bash/builtins/bind.def`
- [ ] Keybinding management: Missing implementation
- [ ] Flag `-l`: `third_party/bash/builtins/bind.def:L139` (list)
- [ ] Flag `-v`: `third_party/bash/builtins/bind.def:L139` (list functions)
- [ ] Flag `-p`: `third_party/bash/builtins/bind.def:L139` (print status)
- [ ] Flag `-V`: `third_party/bash/builtins/bind.def:L139` (list variables)
- [ ] Flag `-P`: `third_party/bash/builtins/bind.def:L139` (print functions)
- [ ] Flag `-s`: `third_party/bash/builtins/bind.def:L139` (list macros)
- [ ] Flag `-S`: `third_party/bash/builtins/bind.def:L139` (print macros)
- [ ] Flag `-X`: `third_party/bash/builtins/bind.def:L139` (list keyseq bindings)
- [ ] Flag `-f=FILE`: `third_party/bash/builtins/bind.def:L139` (read from file)
- [ ] Flag `-q=FUNC`: `third_party/bash/builtins/bind.def:L139` (query keys for func)
- [ ] Flag `-u=FUNC`: `third_party/bash/builtins/bind.def:L139` (unbind func)
- [ ] Flag `-m=KEYMAP`: `third_party/bash/builtins/bind.def:L139` (keymap)
- [ ] Flag `-r=KEYSEQ`: `third_party/bash/builtins/bind.def:L139` (remove seq)
- [ ] Flag `-x=KEYSEQ:SHELLCMD`: `third_party/bash/builtins/bind.def:L139` (exec cmd)

### `break`

- [ ] Upstream: `third_party/bash/builtins/break.def`

### `builtin`

- [x] Upstream: `third_party/bash/builtins/builtin.def`
- [x] Basic execution: Implemented in `internal/commands/builtin/builtin.go`

### `caller`

- [ ] Upstream: `third_party/bash/builtins/caller.def`

### `cat`

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

- [ ] Upstream: `third_party/coreutils/src/chcon.c`
- [ ] Flag `-h`, `--no-dereference`: `third_party/coreutils/src/chcon.c:L125`
- [ ] Flag `-H`: `third_party/coreutils/src/chcon.c:L121`
- [ ] Flag `-L`: `third_party/coreutils/src/chcon.c:L131`
- [ ] Flag `-P`: `third_party/coreutils/src/chcon.c:L135`
- [ ] Flag `-R`, `--recursive`: `third_party/coreutils/src/chcon.c:L139`
- [ ] Flag `-u`, `--user=USER`: `third_party/coreutils/src/chcon.c:L147`
- [ ] Flag `-r`, `--role=ROLE`: `third_party/coreutils/src/chcon.c:L143`
- [ ] Flag `-t`, `--type=TYPE`: `third_party/coreutils/src/chcon.c:L151`
- [ ] Flag `-l`, `--range=RANGE`: `third_party/coreutils/src/chcon.c:L155`
- [ ] Flag `--reference=RFILE`: `third_party/coreutils/src/chcon.c:L131`

### `chgrp`

- [x] Upstream: `third_party/coreutils/src/chown-chgrp.c`
- [x] Inherits flags from `chown`: `--dereference`, `--no-dereference`, `--recursive`, `--from`, `--reference`, `-H`, `-L`, `-P`, `-c`, `-f`, `-v`
- [x] Basic group change: Implemented in `internal/commands/chown/chown.go`

### `chmod`

- [x] Basic mode change: Implemented in `internal/commands/chmod/chmod.go`
- [x] Numeric mode support: `internal/commands/chmod/chmod.go`
- [x] Symbolic mode support: `internal/commands/chmod/chmod.go`
- [x] Flag `--dereference`: `internal/commands/chmod/chmod.go` (Ignored)
- [x] Flag `--no-preserve-root`: `internal/commands/chmod/chmod.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chmod--chown))
- [x] Flag `--preserve-root`: `internal/commands/chmod/chmod.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chmod--chown))
- [x] Flag `--reference=RFILE`: `internal/commands/chmod/chmod.go`
- [x] Flag `-R`: `internal/commands/chmod/chmod.go`
- [x] Flag `-c`: `internal/commands/chmod/chmod.go`
- [x] Flag `-f`: `internal/commands/chmod/chmod.go`
- [x] Flag `-h`: `internal/commands/chmod/chmod.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`: `internal/commands/chmod/chmod.go`

### `chown`

- [x] Basic ownership change: Implemented in `internal/commands/chown/chown.go`
- [x] Flag `--dereference`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `--from`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `--from=CURRENT_OWNER:CURRENT_GROUP`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `--no-preserve-root`: `internal/commands/chown/chown.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chmod--chown))
- [x] Flag `--preserve-root`: `internal/commands/chown/chown.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#chmod--chown))
- [x] Flag `--reference=RFILE`: `internal/commands/chown/chown.go`
- [x] Flag `-H`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `-L`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `-P`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `-R`: `internal/commands/chown/chown.go`
- [x] Flag `-c`: `internal/commands/chown/chown.go`
- [x] Flag `-f`: `internal/commands/chown/chown.go`
- [x] Flag `-h`: `internal/commands/chown/chown.go` (Ignored)
- [x] Flag `-v`: `internal/commands/chown/chown.go`

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

- [ ] Upstream: `third_party/bash/builtins/complete.def`
- [ ] Inherits all `complete` flags: `third_party/bash/builtins/complete.def`

### `complete`

- [ ] Upstream: `third_party/bash/builtins/complete.def`
- [ ] Completion management: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/complete.def:L205` (alias)
- [ ] Flag `-b`: `third_party/bash/builtins/complete.def:L205` (builtin)
- [ ] Flag `-c`: `third_party/bash/builtins/complete.def:L205` (command)
- [ ] Flag `-d`: `third_party/bash/builtins/complete.def:L205` (directory)
- [ ] Flag `-e`: `third_party/bash/builtins/complete.def:L205` (export)
- [ ] Flag `-f`: `third_party/bash/builtins/complete.def:L204` (file)
- [ ] Flag `-g`: `third_party/bash/builtins/complete.def:L205` (group)
- [ ] Flag `-j`: `third_party/bash/builtins/complete.def:L205` (job)
- [ ] Flag `-k`: `third_party/bash/builtins/complete.def:L205` (keyword)
- [ ] Flag `-p`: `third_party/bash/builtins/complete.def:L205` (print)
- [ ] Flag `-r`: `third_party/bash/builtins/complete.def:L205` (remove)
- [ ] Flag `-s`: `third_party/bash/builtins/complete.def:L205` (service)
- [ ] Flag `-u`: `third_party/bash/builtins/complete.def:L205` (user)
- [ ] Flag `-v`: `third_party/bash/builtins/complete.def:L205` (variable)
- [ ] Flag `-o=OPT`: `third_party/bash/builtins/complete.def:L205` (option)
- [ ] Flag `-A=ACTION`: `third_party/bash/builtins/complete.def:L205` (action)
- [ ] Flag `-G=GLOB`: `third_party/bash/builtins/complete.def:L205` (glob)
- [ ] Flag `-W=WORDLIST`: `third_party/bash/builtins/complete.def:L205` (wordlist)
- [ ] Flag `-P=PREFIX`: `third_party/bash/builtins/complete.def:L205` (prefix)
- [ ] Flag `-S=SUFFIX`: `third_party/bash/builtins/complete.def:L205` (suffix)
- [ ] Flag `-X=FILTER`: `third_party/bash/builtins/complete.def:L205` (filter)
- [ ] Flag `-F=FUNC`: `third_party/bash/builtins/complete.def:L205` (function)
- [ ] Flag `-C=CMD`: `third_party/bash/builtins/complete.def:L205` (command)
- [ ] Flag `-E`: `third_party/bash/builtins/complete.def:L205` (empty)
- [ ] Flag `-I`: `third_party/bash/builtins/complete.def:L205` (initial)
- [ ] Flag `-D`: `third_party/bash/builtins/complete.def:L205` (default)

### `compopt`

- [ ] Upstream: `third_party/bash/builtins/complete.def`
- [ ] Flag `-o`, `--options`: `third_party/bash/builtins/complete.def:L323`
- [ ] Flag `-D`, `--default`: `third_party/bash/builtins/complete.def:L321`
- [ ] Flag `-E`, `--empty`: `third_party/bash/builtins/complete.def:L322`

### `continue`

- [ ] Upstream: `third_party/bash/builtins/break.def`
- [ ] Basic operation: Missing implementation

### `cp`

- [x] Basic copy: Implemented in `internal/commands/cp/cp.go`
- [ ] Flag `-a`, `--archive`: `third_party/coreutils/src/cp.c:L173`
- [ ] Flag `-b`, `--backup`: `third_party/coreutils/src/cp.c:L181`
- [ ] Flag `-d`: `third_party/coreutils/src/cp.c:L185` (implies -P --preserve=links)
- [x] Flag `-f`, `--force`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [ ] Flag `-H`: `third_party/coreutils/src/cp.c:L201`
- [x] Flag `-i`, `--interactive`: `internal/commands/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [ ] Flag `-l`, `--link`: `third_party/coreutils/src/cp.c:L209`
- [ ] Flag `-L`, `--dereference`: `third_party/coreutils/src/cp.c:L213`
- [x] Flag `-n`, `--no-clobber`: `internal/commands/cp/cp.go`
- [ ] Flag `-p`: `third_party/coreutils/src/cp.c:L234` (same as --preserve=mode,ownership,timestamps)
- [ ] Flag `-P`, `--no-dereference`: `third_party/coreutils/src/cp.c:L230`
- [x] Flag `-r`, `-R`, `--recursive`: `internal/commands/cp/cp.go`
- [ ] Flag `-s`, `--symbolic-link`: `third_party/coreutils/src/cp.c:L258`
- [x] Flag `-t`, `--target-directory`: `internal/commands/cp/cp.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/cp/cp.go`
- [x] Flag `-u`, `--update`: `internal/commands/cp/cp.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/cp/cp.go`
- [ ] Flag `-x`, `--one-file-system`: `third_party/coreutils/src/cp.c:L278`
- [ ] Flag `-Z`, `--context`: `third_party/coreutils/src/cp.c:L282`
- [ ] Flag `--attributes-only`: `third_party/coreutils/src/cp.c:L177`
- [ ] Flag `--preserve[=ATTR_LIST]`: `third_party/coreutils/src/cp.c:L242`
- [ ] Flag `--no-preserve=ATTR_LIST`: `third_party/coreutils/src/cp.c:L226`
- [ ] Flag `--parents`: `third_party/coreutils/src/cp.c:L238`
- [ ] Flag `--reflink[=WHEN]`: `third_party/coreutils/src/cp.c:L246`
- [ ] Flag `--sparse=WHEN`: `third_party/coreutils/src/cp.c:L254`
- [ ] Flag `--strip-trailing-slashes`: `third_party/coreutils/src/cat.c:L262`

### `csplit`

- [x] Upstream: `third_party/coreutils/src/csplit.c`
- [x] Basic split: Implemented in `internal/commands/csplit/csplit.go`
- [ ] Flag `--suppress-matched`: `third_party/coreutils/src/csplit.c:1435`
- [ ] Flag `-b`: `third_party/coreutils/src/csplit.c:1423`
- [ ] Flag `-f`: `third_party/coreutils/src/csplit.c:1427`
- [ ] Flag `-k`: `third_party/coreutils/src/csplit.c:1431`
- [ ] Flag `-n`: `third_party/coreutils/src/csplit.c:1439`
- [ ] Flag `-s`: `third_party/coreutils/src/csplit.c:1443`
- [ ] Flag `-z`: `third_party/coreutils/src/csplit.c:1447`

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
- [ ] Flag `-d`, `--date=STRING`: `third_party/coreutils/src/date.c:L501`
- [ ] Flag `-f`, `--file=DATEFILE`: `third_party/coreutils/src/date.c:L508`
- [ ] Flag `-I[FMT]`, `--iso-8601[=FMT]`: `third_party/coreutils/src/date.c:L513`
- [ ] Flag `-r`, `--reference=FILE`: `third_party/coreutils/src/date.c:L543`
- [ ] Flag `-R`, `--rfc-email`: `third_party/coreutils/src/date.c:L549`
- [ ] Flag `-s`, `--set=STRING`: `third_party/coreutils/src/date.c:L553`
- [ ] Flag `-u`, `--utc`, `--universal`: `third_party/coreutils/src/date.c:L561`
- [ ] Flag `--debug`: `third_party/coreutils/src/date.c:L497`

### `dd`

- [x] Upstream: `third_party/coreutils/src/dd.c`
- [x] Data copy: Implemented in `internal/commands/dd/dd.go`
- [x] Operand `bs=BYTES`: `internal/commands/dd/dd.go`
- [ ] Operand `cbs=BYTES`: `third_party/coreutils/src/dd.c:L539`
- [ ] Operand `conv=CONVS`: `third_party/coreutils/src/dd.c:L543`
- [x] Operand `count=N`: `internal/commands/dd/dd.go`
- [ ] Operand `ibs=BYTES`: `third_party/coreutils/src/dd.c:L549`
- [x] Operand `if=FILE`: `internal/commands/dd/dd.go`
- [ ] Operand `iflag=FLAGS`: `third_party/coreutils/src/dd.c:L555`
- [ ] Operand `obs=BYTES`: `third_party/coreutils/src/dd.c:L558`
- [x] Operand `of=FILE`: `internal/commands/dd/dd.go`
- [ ] Operand `oflag=FLAGS`: `third_party/coreutils/src/dd.c:L564`
- [x] Operand `seek=N`: `internal/commands/dd/dd.go`
- [x] Operand `skip=N`: `internal/commands/dd/dd.go`
- [ ] Operand `status=LEVEL`: `third_party/coreutils/src/dd.c:L573`
- [x] Operand `conv=notrunc`: `internal/commands/dd/dd.go`

### `declare`

- [x] Attribute management (-i, -r, -x, -a, -A): Implemented in `internal/commands/declare/declare.go`
- [x] Flag `-a`: `internal/commands/declare/declare.go`
- [x] Flag `-A`: `internal/commands/declare/declare.go`
- [x] Flag `-i`: `internal/commands/declare/declare.go`
- [x] Flag `-r`: `internal/commands/declare/declare.go`
- [x] Flag `-x`: `internal/commands/declare/declare.go`
- [ ] Flag `-l`: `third_party/bash/builtins/declare.def:L348` (lowercase)
- [ ] Flag `-u`: `third_party/bash/builtins/declare.def:L353` (uppercase)
- [ ] Flag `-n`: `third_party/bash/builtins/declare.def:L327` (nameref)
- [ ] Flag `-t`: `third_party/bash/builtins/declare.def:L333` (trace)
- [ ] Flag `-f`: `third_party/bash/builtins/declare.def:L313` (function)
- [ ] Flag `-F`: `third_party/bash/builtins/declare.def:L309` (function name only)
- [ ] Flag `-g`: `third_party/bash/builtins/declare.def:L320` (global)
- [x] Flag `-p`: `internal/commands/declare/declare.go`
- [ ] Flag `-I`: `third_party/bash/builtins/declare.def:L359` (inherit attributes)
- [x] Aliases: `typeset`

### `df`

- [x] Upstream: `third_party/coreutils/src/df.c`
- [x] Basic df: Implemented in `internal/commands/df/df.go`
- [x] Basic output: Implemented in `internal/commands/df/df.go`
- [ ] Flag `--no-sync`: `third_party/coreutils/src/df.c:L266`
- [ ] Flag `--output[=FIELD_LIST]`: `third_party/coreutils/src/df.c:L262`
- [ ] Flag `--sync`: `third_party/coreutils/src/df.c:L265`
- [ ] Flag `--total`: `third_party/coreutils/src/df.c:L267`
- [ ] Flag `-B`: `third_party/coreutils/src/df.c:L257`
- [x] Flag `-H`: `internal/commands/df/df.go`
- [ ] Flag `-P`: `third_party/coreutils/src/df.c:L263`
- [x] Flag `-T`: `internal/commands/df/df.go`
- [x] Flag `-a`: `internal/commands/df/df.go`
- [x] Flag `-h`: `internal/commands/df/df.go`
- [ ] Flag `-i`: `third_party/coreutils/src/df.c:L258`
- [x] Flag `-k`: `internal/commands/df/df.go`
- [ ] Flag `-l`: `third_party/coreutils/src/df.c:L261`
- [ ] Flag `-t`: `third_party/coreutils/src/df.c:L268`
- [ ] Flag `-x`: `third_party/coreutils/src/df.c:L269`

### `dir`

- [ ] Upstream: `third_party/coreutils/src/coreutils-dir.c`
- [ ] Inherits flags from `ls`

### `dircolors`

- [ ] Upstream: `third_party/coreutils/src/dircolors.c`
- [ ] Output configuration: Missing implementation
- [ ] Flag `-b`, `--sh`, `--bourne-shell`: `third_party/coreutils/src/dircolors.c:L158`
- [ ] Flag `-c`, `--csh`, `--c-shell`: `third_party/coreutils/src/dircolors.c:L162`
- [ ] Flag `-p`, `--print-database`: `third_party/coreutils/src/dircolors.c:L166`
- [ ] Flag `--print-ls-colors`: `third_party/coreutils/src/dircolors.c:L170`

### `dirname`

- [x] Upstream: `third_party/coreutils/src/dirname.c`
- [x] Basic operation: Implemented in `internal/commands/dirname/dirname.go`
- [x] Flag `-z`, `--zero`: `internal/commands/dirname/dirname.go`

### `dirs`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic listing: Implemented in `internal/commands/dirs/dirs.go`
- [x] Flag `-c`: `internal/commands/dirs/dirs.go` (clear stack)
- [ ] Flag `-l`: `third_party/bash/builtins/pushd.def:L380` (long listing)
- [x] Flag `-p`: `internal/commands/dirs/dirs.go` (print with one line per entry)
- [x] Flag `-v`: `internal/commands/dirs/dirs.go` (verbose)

### `disown`

- [ ] Upstream: `third_party/bash/builtins/jobs.def`
- [ ] Flag `-a`: `third_party/bash/builtins/jobs.def:L196` (all jobs)
- [ ] Flag `-h`: `third_party/bash/builtins/jobs.def:L199` (mark to not receive SIGHUP)
- [ ] Flag `-r`: `third_party/bash/builtins/jobs.def:L202` (running jobs only)

### `du`

- [x] Upstream: `third_party/coreutils/src/du.c`
- [x] Basic du: Implemented in `internal/commands/du/du.go`
- [x] Basic usage summary: Implemented in `internal/commands/du/du.go`
- [ ] Flag `-a`, `--all`: `third_party/coreutils/src/du.c:L294`
- [ ] Flag `-A`, `--apparent-size`: `third_party/coreutils/src/du.c:L298`
- [ ] Flag `-b`, `--bytes`: `third_party/coreutils/src/du.c:L310`
- [ ] Flag `-c`, `--total`: `third_party/coreutils/src/du.c:L314`
- [ ] Flag `-d`, `--max-depth=N`: `third_party/coreutils/src/du.c:L322`
- [ ] Flag `-D`, `--dereference-args`: `third_party/coreutils/src/du.c:L318`
- [x] Flag `-h`, `--human-readable`: `internal/commands/du/du.go`
- [ ] Flag `-H`: `third_party/coreutils/src/du.c:L333` (same as --dereference-args)
- [ ] Flag `-k`: `third_party/coreutils/src/du.c:L345` (1K blocks)
- [ ] Flag `-l`, `--count-links`: `third_party/coreutils/src/du.c:L353`
- [ ] Flag `-L`, `--dereference`: `third_party/coreutils/src/du.c:L349`
- [ ] Flag `-m`: `third_party/coreutils/src/du.c:L357` (1M blocks)
- [ ] Flag `-P`, `--no-dereference`: `third_party/coreutils/src/du.c:L361`
- [x] Flag `-s`, `--summarize`: `internal/commands/du/du.go`
- [ ] Flag `-S`, `--separate-dirs`: `third_party/coreutils/src/du.c:L365`
- [ ] Flag `-t`, `--threshold=SIZE`: `third_party/coreutils/src/du.c:L377`
- [ ] Flag `-x`, `--one-file-system`: `third_party/coreutils/src/du.c:L402`
- [ ] Flag `-X`, `--exclude-from=FILE`: `third_party/coreutils/src/du.c:L394`
- [ ] Flag `-0`, `--null`: `third_party/coreutils/src/du.c:L290`
- [ ] Flag `--exclude=PATTERN`: `third_party/coreutils/src/du.c:L398`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/du.c:L328`
- [ ] Flag `--inodes`: `third_party/coreutils/src/du.c:L341`
- [ ] Flag `--si`: `third_party/coreutils/src/du.c:L369` (1000 instead of 1024)
- [ ] Flag `--time[=WORD]`: `third_party/coreutils/src/du.c:L382`
- [ ] Flag `--time-style=STYLE`: `third_party/coreutils/src/du.c:L389`

### `echo`

- [x] Basic output: Implemented in `internal/commands/echo/echo.go`
- [x] Flag `-E`: `internal/commands/echo/echo.go`
- [x] Flag `-e`: `internal/commands/echo/echo.go`
- [x] Flag `-n`: `internal/commands/echo/echo.go`

### `enable`

- [ ] Upstream: `third_party/bash/builtins/enable.def`
- [ ] Flag `-a`: `third_party/bash/builtins/enable.def:L157` (display all)
- [ ] Flag `-d`: `third_party/bash/builtins/enable.def:L160` (delete loaded)
- [ ] Flag `-n`: `third_party/bash/builtins/enable.def:L163` (disable)
- [ ] Flag `-p`: `third_party/bash/builtins/enable.def:L166` (print status)
- [ ] Flag `-s`: `third_party/bash/builtins/enable.def:L169` (POSIX special only)
- [ ] Flag `-f filename`: `third_party/bash/builtins/enable.def:L172` (load from dynamic file)

### `env`

- [x] Upstream: `third_party/coreutils/src/env.c`
- [x] Basic execution: Implemented in `internal/commands/env/env.go`
- [ ] Flag `-a`, `--argv0=ARG`: `third_party/coreutils/src/env.c:L123`
- [x] Flag `-i`, `--ignore-environment`: `internal/commands/env/env.go`
- [x] Flag `-u`, `--unset=NAME`: `internal/commands/env/env.go`
- [x] Flag `-0`, `--null`: `internal/commands/env/env.go`
- [x] Flag `-C`, `--chdir=DIR`: `internal/commands/env/env.go`
- [ ] Flag `-S`, `--split-string=S`: `third_party/coreutils/src/env.c:L143`
- [ ] Flag `-v`, `--debug`: `third_party/coreutils/src/env.c:L164`
- [ ] Flag `--block-signal[=SIG]`: `third_party/coreutils/src/env.c:L148`
- [ ] Flag `--default-signal[=SIG]`: `third_party/coreutils/src/env.c:L152`
- [ ] Flag `--ignore-signal[=SIG]`: `third_party/coreutils/src/env.c:L156`
- [ ] Flag `--list-signal-handling`: `third_party/coreutils/src/env.c:L160`

### `eval`

- [x] Upstream: `third_party/bash/builtins/eval.def`
- [x] Basic execution: Implemented in `internal/commands/eval/eval.go`

### `exec`

- [x] Upstream: `third_party/bash/builtins/exec.def`
- [x] Basic execution: Implemented in `internal/commands/exec/exec.go`
- [ ] Flag `-l`: `third_party/bash/builtins/exec.def:L117` (login shell)
- [ ] Flag `-a name`: `third_party/bash/builtins/exec.def:L120`
- [ ] Flag `-c`: `third_party/bash/builtins/exec.def:L114`

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
- [ ] Flag `-f`: `third_party/bash/builtins/setattr.def:L142` (functions)
- [ ] Flag `-n`: `third_party/bash/builtins/setattr.def:L142` (remove export attribute)
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

### `false`

- [x] Upstream: `third_party/bash/builtins/colon.def`, `third_party/coreutils/src/false.c`
- [x] Basic operation: Implemented in `internal/commands/boolcmd/bool.go`

### `fc`

- [ ] Upstream: `third_party/bash/builtins/fc.def`
- [ ] Basic editing/re-execution: Missing implementation
- [ ] Flag `-e ENAME`: `third_party/bash/builtins/fc.def:L232` (editor name)
- [ ] Flag `-l`: `third_party/bash/builtins/fc.def:L220` (list mode)
- [ ] Flag `-n`: `third_party/bash/builtins/fc.def:L216` (no numbers)
- [ ] Flag `-r`: `third_party/bash/builtins/fc.def:L224` (reverse order)
- [ ] Flag `-s`: `third_party/bash/builtins/fc.def:L228` (re-execute)

### `fg`

- [ ] Upstream: `third_party/bash/builtins/fg_bg.def`

### `fmt`

- [x] Upstream: `third_party/coreutils/src/fmt.c`
- [x] Paragraph formatting: Implemented in `internal/commands/fmt/fmt.go`
- [ ] Flag `-c`, `--crown-margin`: `third_party/coreutils/src/fmt.c:L185`
- [ ] Flag `-p`, `--prefix=STRING`: `third_party/coreutils/src/fmt.c:L189`
- [x] Flag `-s`, `--split-only`: `internal/commands/fmt/fmt.go`
- [ ] Flag `-t`, `--tagged-paragraph`: `third_party/coreutils/src/fmt.c:L198`
- [ ] Flag `-u`, `--uniform-spacing`: `third_party/coreutils/src/fmt.c:L202`
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/fmt/fmt.go`
- [ ] Flag `-g`, `--goal=WIDTH`: `third_party/coreutils/src/fmt.c:L212`
- [ ] Flag `-WIDTH`: `third_party/coreutils/src/fmt.c:L178`

### `fold`

- [x] Upstream: `third_party/coreutils/src/fold.c`
- [x] Line wrapping: Implemented in `internal/commands/fold/fold.go`
- [x] Flag `-b`, `--bytes`: `internal/commands/fold/fold.go`
- [x] Flag `-c`, `--characters`: `internal/commands/fold/fold.go`
- [x] Flag `-s`, `--spaces`: `internal/commands/fold/fold.go`
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/fold/fold.go`

### `getlimits`

- [ ] Upstream: `third_party/coreutils/src/getlimits.c`
- [ ] Flag `--help`: `third_party/coreutils/src/getlimits.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/getlimits.c:L52`

### `getopts`

- [x] Basic parsing: Implemented in `internal/commands/getopts/getopts.go`
- [x] Silent mode support (`:`): `internal/commands/getopts/getopts.go`

### `groups`

- [x] Upstream: `third_party/coreutils/src/groups.c`
- [x] Basic listing: Implemented in `internal/commands/groups/groups.go`
- [x] Multiple users support: `internal/commands/groups/groups.go`

### `hash`

- [x] Upstream: `third_party/bash/builtins/hash.def`
- [x] Command hashing: Implemented in `internal/commands/hash/hash.go`
- [x] Flag `-r`: `internal/commands/hash/hash.go`
- [ ] Flag `-d`: `third_party/bash/builtins/hash.def:L125` (forget name)
- [ ] Flag `-p`: `third_party/bash/builtins/hash.def:L126` (use path)
- [ ] Flag `-t`: `third_party/bash/builtins/hash.def:L127` (list name)
- [ ] Flag `-l`: `third_party/bash/builtins/hash.def:L128` (output format)

### `head`

- [x] Basic output: Implemented in `internal/commands/head/head.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/head/head.go`
- [x] Flag `-n`, `--lines`: `internal/commands/head/head.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/head/head.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/head/head.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/head/head.go`

### `help`

- [x] Upstream: `third_party/bash/builtins/help.def`
- [x] Help system: Implemented in `internal/commands/help/help.go`
- [ ] Flag `-d`: `third_party/bash/builtins/help.def:L105` (short description)
- [ ] Flag `-m`: `third_party/bash/builtins/help.def:L108` (man-page format)
- [ ] Flag `-s`: `third_party/bash/builtins/help.def:L111` (syntax only)

### `history`

- [x] Upstream: `third_party/bash/builtins/history.def`
- [x] History management: Implemented in `internal/commands/history/history.go`
- [x] Flag `-d offset`: `internal/commands/history/history.go` (delete entry)
- [x] Flag `-a`: `internal/commands/history/history.go` (append)
- [x] Flag `-c`: `internal/commands/history/history.go` (clear)
- [x] Flag `-n`: `internal/commands/history/history.go` (read non-recorded)
- [ ] Flag `-p`: `third_party/bash/builtins/history.def:L148` (print/expand)
- [x] Flag `-r`: `internal/commands/history/history.go` (read file)
- [ ] Flag `-s`: `third_party/bash/builtins/history.def:L141` (store/append)
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
- [ ] Flag `-D`: `third_party/coreutils/src/install.c:L135`
- [x] Flag `-g`, `--group=GROUP`: `internal/commands/install/install.go`
- [ ] Flag `-m`, `--mode=MODE`: `third_party/coreutils/src/install.c:L141`
- [x] Flag `-o`, `--owner=OWNER`: `internal/commands/install/install.go`
- [ ] Flag `-p`, `--preserve-timestamps`: `third_party/coreutils/src/install.c:L147`
- [ ] Flag `-s`, `--strip`: `third_party/coreutils/src/install.c:L150`
- [ ] Flag `-S`, `--suffix=SUFFIX`: `third_party/coreutils/src/install.c:L153`
- [ ] Flag `-t`, `--target-directory=DIR`: `third_party/coreutils/src/install.c:L156`
- [ ] Flag `-T`, `--no-target-directory`: `third_party/coreutils/src/install.c:L159`
- [x] Flag `-v`, `--verbose`: `internal/commands/install/install.go`
- [ ] Flag `-C`, `--compare`: `third_party/coreutils/src/install.c:L165`

### `jobs`

- [x] Upstream: `third_party/bash/builtins/jobs.def`
- [x] Basic listing: Implemented in `internal/commands/jobs/jobs.go`
- [x] Flag `-l`: `internal/commands/jobs/jobs.go` (long format)
- [ ] Flag `-n`: `third_party/bash/builtins/jobs.def:L97` (only jobs that changed)
- [x] Flag `-p`: `internal/commands/jobs/jobs.go` (only PIDs)
- [x] Flag `-r`: `internal/commands/jobs/jobs.go` (running only)
- [x] Flag `-s`: `internal/commands/jobs/jobs.go` (stopped only)
- [ ] Flag `-x command`: `third_party/bash/builtins/jobs.def:L109` (execute command)

### `join`

- [x] Upstream: `third_party/coreutils/src/join.c`
- [x] Basic join: Implemented in `internal/commands/join/join.go`
- [ ] Flag `-1 FIELD`: `third_party/coreutils/src/join.c:L186`
- [ ] Flag `-2 FIELD`: `third_party/coreutils/src/join.c:L187`
- [ ] Flag `-a FILENUM`: `third_party/coreutils/src/join.c:L188`
- [ ] Flag `-e STRING`: `third_party/coreutils/src/join.c:L189`
- [ ] Flag `-i`, `--ignore-case`: `third_party/coreutils/src/join.c:L190`
- [ ] Flag `-j FIELD`: `third_party/coreutils/src/join.c:L191`
- [ ] Flag `-o FORMAT`: `third_party/coreutils/src/join.c:L192`
- [ ] Flag `-t CHAR`: `third_party/coreutils/src/join.c:L193`
- [ ] Flag `-v FILENUM`: `third_party/coreutils/src/join.c:L194`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/join.c:L195`
- [ ] Flag `--check-order`: `third_party/coreutils/src/join.c:L196`
- [ ] Flag `--nocheck-order`: `third_party/coreutils/src/join.c:L197`
- [ ] Flag `--header`: `third_party/coreutils/src/join.c:L198`
- [ ] Flag `--check-order`: `third_party/coreutils/src/join.c:L248`
- [ ] Flag `--header`: `third_party/coreutils/src/join.c:L257`
- [ ] Flag `--nocheck-order`: `third_party/coreutils/src/join.c:L253`
- [ ] Flag `-1 FIELD`: `third_party/coreutils/src/join.c:L240`
- [ ] Flag `-2 FIELD`: `third_party/coreutils/src/join.c:L244`
- [ ] Flag `-a FILENUM`: `third_party/coreutils/src/join.c:L210`
- [ ] Flag `-e STRING`: `third_party/coreutils/src/join.c:L215`
- [ ] Flag `-i`: `third_party/coreutils/src/join.c:L220`
- [ ] Flag `-j FIELD`: `third_party/coreutils/src/join.c:L224`
- [ ] Flag `-o FORMAT`: `third_party/coreutils/src/join.c:L228`
- [ ] Flag `-t CHAR`: `third_party/coreutils/src/join.c:L232`
- [ ] Flag `-v FILENUM`: `third_party/coreutils/src/join.c:L236`
- [ ] Flag `-z`: `third_party/coreutils/src/join.c:L262`

### `kill`

- [x] Upstream: `third_party/bash/builtins/kill.def`
- [x] Basic signaling: Implemented in `internal/commands/kill/kill.go`
- [x] Flag `-l`: `internal/commands/kill/kill.go`
- [ ] Flag `-n num`: `third_party/bash/builtins/kill.def:L130`
- [ ] Flag `-l`: `third_party/coreutils/src/kill.c:L277` / `third_party/bash/builtins/kill.def:L114`
- [ ] Flag `-s SIGNAL`: `third_party/coreutils/src/kill.c:L262` / `third_party/bash/builtins/kill.def:L129`

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

- [ ] Upstream: `third_party/bash/builtins/declare.def`
- [ ] Inherits all `declare` attributes: `third_party/bash/builtins/declare.def`

### `logname`

- [x] Upstream: `third_party/coreutils/src/logname.c`
- [x] Basic operation: Implemented in `internal/commands/logname/logname.go`
- [x] Flag `--help`: `internal/commands/logname/logname.go`
- [x] Flag `--version`: `internal/commands/logname/logname.go`

### `logout`

- [x] Upstream: `third_party/bash/builtins/exit.def`
- [x] Basic operation: Implemented in `internal/commands/logout/logout.go`

#### `ls`

- [x] Basic listing: `internal/commands/ls/ls.go`
- [ ] Color output (`--color`): `third_party/coreutils/src/ls.c:L215`
- [x] Flag `--author`: `internal/commands/ls/ls.go` (partial via info)
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
- [ ] Flag `-C`: `third_party/coreutils/src/ls.c:L181`
- [ ] Flag `-D`: `third_party/coreutils/src/ls.c:L224`
- [x] Flag `-F`: `internal/commands/ls/ls.go` (classify)
- [x] Flag `-G`: `internal/commands/ls/ls.go` (no-group)
- [x] Flag `-H`: `internal/commands/ls/ls.go` (dereference-command-line)
- [x] Flag `-I`: `internal/commands/ls/ls.go` (ignore)
- [x] Flag `-L`: `internal/commands/ls/ls.go` (dereference)
- [ ] Flag `-N`: `third_party/coreutils/src/ls.c:L329`
- [x] Flag `-Q`: `internal/commands/ls/ls.go` (quote-name)
- [x] Flag `-R`: `internal/commands/ls/ls.go` (recursive)
- [x] Flag `-S`: `internal/commands/ls/ls.go` (sort-size)
- [ ] Flag `-T`: `third_party/coreutils/src/ls.c:L411`
- [x] Flag `-U`: `internal/commands/ls/ls.go` (do not sort)
- [x] Flag `-X`: `internal/commands/ls/ls.go` (extension sort)
- [ ] Flag `-Z`: `third_party/coreutils/src/ls.c:L439`
- [x] Flag `-a`: `internal/commands/ls/ls.go` (all)
- [x] Flag `-b`: `internal/commands/ls/ls.go` (escape)
- [x] Flag `-c`: `internal/commands/ls/ls.go` (ctime)
- [x] Flag `-d`: `internal/commands/ls/ls.go` (directory itself)
- [x] Flag `-f`: `internal/commands/ls/ls.go` (do not sort, enable -aU)
- [x] Flag `-g`: `internal/commands/ls/ls.go` (like -l but no owner)
- [x] Flag `-h`: `internal/commands/ls/ls.go` (human-readable)
- [x] Flag `-i`: `internal/commands/ls/ls.go` (inode)
- [x] Flag `-k`: `internal/commands/ls/ls.go` (kibibytes)
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
- [ ] Flag `-w`: `third_party/coreutils/src/ls.c:L427`
- [x] Flag `-x`: `internal/commands/ls/ls.go` (across/horizontal)

### `mapfile`

- [x] Upstream: `third_party/bash/builtins/mapfile.def`
- [x] Array population: Implemented in `internal/commands/mapfile/mapfile.go`
- [ ] Flag `-d`: `third_party/bash/builtins/mapfile.def:L238` (delimiter)
- [x] Flag `-t`: `internal/commands/mapfile/mapfile.go` (trim)
- [ ] Flag `-n`: `third_party/bash/builtins/mapfile.def:L238` (count)
- [ ] Flag `-O`: `third_party/bash/builtins/mapfile.def:L238` (origin)
- [ ] Flag `-t`: `third_party/bash/builtins/mapfile.def:L238` (strip newline)
- [ ] Flag `-u`: `third_party/bash/builtins/mapfile.def:L238` (fd)
- [ ] Flag `-C`: `third_party/bash/builtins/mapfile.def:L238` (callback)
- [ ] Flag `-c`: `third_party/bash/builtins/mapfile.def:L238` (quantum)
- [ ] Flag `-s`: `third_party/bash/builtins/mapfile.def:L76` (array is same)
- [ ] Aliases: `readarray`

### `md5sum`

- [x] Upstream: `third_party/coreutils/src/cksum.c`
- [x] Inherits all `cksum` hash flags: `internal/commands/sum/sum.go`

### `mkdir`

- [x] Upstream: `third_party/coreutils/src/mkdir.c`
- [x] Basic operation: Implemented in `internal/commands/mkdir/mkdir.go`
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/mkdir/mkdir.go` (octal)
- [x] Flag `-p`, `--parents`: `internal/commands/mkdir/mkdir.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/mkdir/mkdir.go`
- [ ] Flag `-Z`, `--context=CTX`: `third_party/coreutils/src/mkdir.c:L78`

### `mkfifo`

- [ ] Upstream: `third_party/coreutils/src/mkfifo.c`
- [ ] Flag `-m`, `--mode=MODE`: `third_party/coreutils/src/mkfifo.c:L60`
- [ ] Flag `-Z`, `--context=CTX`: `third_party/coreutils/src/mkfifo.c:L64`

### `mknod`

- [ ] Upstream: `third_party/coreutils/src/mknod.c`
- [ ] Flag `-m`, `--mode=MODE`: `third_party/coreutils/src/mknod.c:L157`
- [ ] Flag `-Z`, `--context=CTX`: `third_party/coreutils/src/mknod.c:L161`

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
- [ ] Flag `-b`, `--backup`: `third_party/coreutils/src/mv.c:L278`
- [x] Flag `-f`, `--force`: `internal/commands/mv/mv.go` (ignored)
- [x] Flag `-i`, `--interactive`: `internal/commands/mv/mv.go` (ignored)
- [x] Flag `-n`, `--no-clobber`: `internal/commands/mv/mv.go`
- [x] Flag `-t`, `--target-directory`: `internal/commands/mv/mv.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/mv/mv.go`
- [x] Flag `-u`, `--update`: `internal/commands/mv/mv.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/mv/mv.go`
- [ ] Flag `-Z`, `--context`: `third_party/coreutils/src/mv.c:L310`
- [ ] Flag `--exchange`: `third_party/coreutils/src/mv.c:L314`
- [ ] Flag `--no-copy`: `third_party/coreutils/src/mv.c:L318`

### `nice`

- [x] Upstream: `third_party/coreutils/src/nice.c`
- [x] Priority adjustment: Implemented (as wrapper) in `internal/commands/nice/nice.go`
- [x] Flag `-n`, `--adjustment=N`: `internal/commands/nice/nice.go`

### `nl`

- [x] Upstream: `third_party/coreutils/src/nl.c`
- [x] Basic numbering: Implemented in `internal/commands/nl/nl.go`
- [x] Flag `-b`, `--body-numbering`: `internal/commands/nl/nl.go`
- [ ] Flag `-d`, `--section-delimiter`: `third_party/coreutils/src/nl.c:L132`
- [x] Flag `-f`, `--footer-numbering`: `internal/commands/nl/nl.go`
- [x] Flag `-h`, `--header-numbering`: `internal/commands/nl/nl.go`
- [x] Flag `-i`, `--line-increment`: `internal/commands/nl/nl.go`
- [ ] Flag `-l`, `--join-blank-lines`: `third_party/coreutils/src/nl.c:L144`
- [x] Flag `-n`, `--number-format`: `internal/commands/nl/nl.go`
- [ ] Flag `-p`, `--no-renumber`: `third_party/coreutils/src/nl.c:L150`
- [x] Flag `-s`, `--number-separator`: `internal/commands/nl/nl.go`
- [x] Flag `-v`, `--starting-line-number`: `internal/commands/nl/nl.go`
- [x] Flag `-w`, `--number-width`: `internal/commands/nl/nl.go`

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
- [ ] Flag `-d`, `--delimiter=X`: `third_party/coreutils/src/numfmt.c:L1022`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/numfmt.c:L1053`
- [x] Flag `--to`: `internal/commands/numfmt/numfmt.go`
- [x] Flag `--from`: `internal/commands/numfmt/numfmt.go`
- [x] Flag `--header`: `internal/commands/numfmt/numfmt.go`

### `od`

- [x] Upstream: `third_party/coreutils/src/od.c`
- [x] Format output: Implemented in `internal/commands/od/od.go`
- [x] Flag `-A rad`: `internal/commands/od/od.go`
- [x] Flag `-j bytes`: `internal/commands/od/od.go`
- [x] Flag `-N bytes`: `internal/commands/od/od.go`
- [ ] Flag `-t type`: `third_party/coreutils/src/od.c:L318` (only default 2-byte octal supported)
- [ ] Flag `-v`: `third_party/coreutils/src/od.c:L319`
- [x] Flag `-w`: `internal/commands/od/od.go`
- [ ] Flag `-S`: `third_party/coreutils/src/od.c:L320`

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

- [ ] Upstream: `third_party/coreutils/src/pinky.c`
- [ ] Flag `-b`: `third_party/coreutils/src/pinky.c:L467`
- [ ] Flag `-f`: `third_party/coreutils/src/pinky.c:L468`
- [ ] Flag `-h`: `third_party/coreutils/src/pinky.c:L469`
- [ ] Flag `-i`: `third_party/coreutils/src/pinky.c:L470`
- [ ] Flag `-l`: `third_party/coreutils/src/pinky.c:L471`
- [ ] Flag `-p`: `third_party/coreutils/src/pinky.c:L473`
- [ ] Flag `-q`: `third_party/coreutils/src/pinky.c:L472`
- [ ] Flag `-s`: `third_party/coreutils/src/pinky.c:L474`
- [ ] Flag `-w`: `third_party/coreutils/src/pinky.c:L475`

### `popd`

- [x] Upstream: `third_party/bash/builtins/pushd.def`
- [x] Basic popping: Implemented in `internal/commands/popd/popd.go`
- [x] Flag `-n`: `internal/commands/popd/popd.go`

### `pr`

- [x] Upstream: `third_party/coreutils/src/pr.c`
- [x] Print formatting: Implemented in `internal/commands/pr/pr.go`
- [ ] Flag `-a`: `third_party/coreutils/src/pr.c:L316` (multi-column not implemented)
- [x] Flag `-d`: `internal/commands/pr/pr.go`
- [x] Flag `-h`: `internal/commands/pr/pr.go`
- [x] Flag `-l`: `internal/commands/pr/pr.go`
- [ ] Flag `-m`: `third_party/coreutils/src/pr.c:L326`
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

- [ ] Upstream: `third_party/coreutils/src/ptx.c`
- [ ] Flag `-A`, `--auto-reference`: `third_party/coreutils/src/ptx.c:L1863`
- [ ] Flag `-F`, `--flag-truncation=STRING`: `third_party/coreutils/src/ptx.c:L1867`
- [ ] Flag `-G`, `--gnu-extensions`: `third_party/coreutils/src/ptx.c:L1871`
- [ ] Flag `-M`, `--macro-name=STRING`: `third_party/coreutils/src/ptx.c:L1875`
- [ ] Flag `-O`, `--format=roff`: `third_party/coreutils/src/ptx.c:L1879`
- [ ] Flag `-R`, `--right-side-refs`: `third_party/coreutils/src/ptx.c:L1883`
- [ ] Flag `-S`, `--sentence-regexp=REGEXP`: `third_party/coreutils/src/ptx.c:L1887`
- [ ] Flag `-T`, `--format=tex`: `third_party/coreutils/src/ptx.c:L1891`
- [ ] Flag `-W`, `--word-regexp=REGEXP`: `third_party/coreutils/src/ptx.c:L1895`
- [ ] Flag `-b`, `--break-file=FILE`: `third_party/coreutils/src/ptx.c:L1899`
- [ ] Flag `-f`, `--ignore-case`: `third_party/coreutils/src/ptx.c:L1903`
- [ ] Flag `-g`, `--gap-size=NUMBER`: `third_party/coreutils/src/ptx.c:L1907`
- [ ] Flag `-i`, `--ignore-file=FILE`: `third_party/coreutils/src/ptx.c:L1911`
- [ ] Flag `-o`, `--only-file=FILE`: `third_party/coreutils/src/ptx.c:L1915`
- [ ] Flag `-r`, `--references`: `third_party/coreutils/src/ptx.c:L1919`
- [ ] Flag `-t`, `--typeset-mode`: `third_party/coreutils/src/ptx.c:L1923`
- [ ] Flag `-w`, `--width=NUMBER`: `third_party/coreutils/src/ptx.c:L1927`

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
- [x] Flag `-a`, `--array`: Implemented in `internal/commands/read/read.go`
- [x] Flag `-d`, `--delimiter`: Implemented in `internal/commands/read/read.go`
- [ ] Flag `-e`: `third_party/bash/builtins/read.def:L43` (use Readline)
- [ ] Flag `-i`, `--initial-text`: `third_party/bash/builtins/read.def:L49`
- [x] Flag `-n`, `--nchars`: Implemented in `internal/commands/read/read.go`
- [x] Flag `-N`, `--Nchars`: Implemented in `internal/commands/read/read.go`
- [x] Flag `-p`, `--prompt`: `internal/commands/read/read.go`
- [x] Flag `-r`: Implemented in `internal/commands/read/read.go` (raw mode)
- [x] Flag `-s`, `--silent`: `internal/commands/read/read.go`
- [x] Flag `-t`, `--timeout`: `internal/commands/read/read.go`
- [ ] Flag `-u`, `--fd`: `third_party/bash/builtins/read.def:L61`

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

- [ ] Upstream: `third_party/bash/builtins/setattr.def`
- [ ] Attribute management: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/setattr.def:L181` (indexed array)
- [ ] Flag `-A`: `third_party/bash/builtins/setattr.def:L181` (associative array)
- [ ] Flag `-p`: `third_party/bash/builtins/setattr.def:L181` (print)
- [ ] Flag `-f`: `third_party/bash/builtins/setattr.def:L181` (functions)

### `realpath`

- [ ] Upstream: `third_party/coreutils/src/realpath.c`
- [x] Basic output: Implemented in `internal/commands/realpath/realpath.go`
- [x] Flag `-E`, `--canonicalize-existing`: `internal/commands/realpath/realpath.go`
- [x] Flag `-L`, `--logical`: `internal/commands/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-P`, `--physical`: `internal/commands/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-q`, `--quiet`: `internal/commands/realpath/realpath.go`
- [x] Flag `-s`, `--strip`: `internal/commands/realpath/realpath.go`
- [x] Flag `-z`, `--zero`: `internal/commands/realpath/realpath.go`
- [ ] Flag `--relative-to`: `third_party/coreutils/src/realpath.c:L246`
- [x] Flag `-e`: `internal/commands/realpath/realpath.go`
- [x] Flag `-m`: `internal/commands/realpath/realpath.go`

### `return`

- [ ] Upstream: `third_party/bash/builtins/return.def`
- [ ] Basic return: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/return.def:L61`

### `rm`

- [x] Upstream: `third_party/coreutils/src/rm.c`
- [x] Basic removal: Implemented in `internal/commands/rm/rm.go`
- [ ] Flag `-I`: `third_party/coreutils/src/rm.c:L157` (prompt once)
- [x] Flag `-d`, `--dir`: `internal/commands/rm/rm.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/rm/rm.go`
- [x] Flag `-f`: `internal/commands/rm/rm.go`
- [x] Flag `-i`: `internal/commands/rm/rm.go`
- [x] Flag `-r`: `internal/commands/rm/rm.go`

### `rmdir`

- [x] Upstream: `third_party/coreutils/src/rmdir.c`
- [x] Basic rmdir: Implemented in `internal/commands/rmdir/rmdir.go`
- [x] Basic removal: Implemented in `internal/commands/rmdir/rmdir.go`
- [x] Flag `--ignore-fail-on-non-empty`: `internal/commands/rmdir/rmdir.go`
- [x] Flag `-p`: `internal/commands/rmdir/rmdir.go`
- [x] Flag `-v`: `internal/commands/rmdir/rmdir.go`

### `runcon`

- [ ] Upstream: `third_party/coreutils/src/runcon.c`
- [ ] Flag `-c`, `--compute`: `third_party/coreutils/src/runcon.c:L123`
- [ ] Flag `-l`, `--user=USER`: `third_party/coreutils/src/runcon.c:L127`
- [ ] Flag `-r`, `--role=ROLE`: `third_party/coreutils/src/runcon.c:L131`
- [ ] Flag `-t`, `--type=TYPE`: `third_party/coreutils/src/runcon.c:L135`
- [ ] Flag `-u`, `--user=USER`: `third_party/coreutils/src/runcon.c:L139`

### `select`

- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Basic operation: Missing implementation

### `seq`

- [x] Upstream: `third_party/coreutils/src/seq.c`
- [x] Basic sequence: Implemented in `internal/commands/seq/seq.go`
- [ ] Flag `-f`, `--format=FORMAT`: `third_party/coreutils/src/seq.c:L592`
- [x] Flag `-s`, `--separator=STRING`: `internal/commands/seq/seq.go`
- [x] Flag `-w`, `--equal-width`: `internal/commands/seq/seq.go`

### `set`

- [ ] Upstream: `third_party/bash/builtins/set.def`
- [ ] Option management (-e, -u, -x, -o): Missing implementation
- [ ] Positional parameters: `third_party/bash/builtins/set.def:L784`
- [ ] Flag `-a`: `third_party/bash/builtins/set.def:L843` (allexport)
- [ ] Flag `-b`: `third_party/bash/builtins/set.def:L843` (notify)
- [ ] Flag `-e`: `third_party/bash/builtins/set.def:L843` (errexit)
- [ ] Flag `-f`: `third_party/bash/builtins/set.def:L843` (noglob)
- [ ] Flag `-h`: `third_party/bash/builtins/set.def:L843` (hashall)
- [ ] Flag `-k`: `third_party/bash/builtins/set.def:L843` (keyword)
- [ ] Flag `-m`: `third_party/bash/builtins/set.def:L843` (monitor)
- [ ] Flag `-n`: `third_party/bash/builtins/set.def:L849` (noexec)
- [ ] Flag `-o`: `third_party/bash/builtins/set.def:L732` (option-name)
- [ ] Flag `-p`: `third_party/bash/builtins/set.def:L843` (privileged)
- [ ] Flag `-t`: `third_party/bash/builtins/set.def:L843` (exit after one command)
- [ ] Flag `-u`: `third_party/bash/builtins/set.def:L843` (nounset)
- [ ] Flag `-v`: `third_party/bash/builtins/set.def:L846` (verbose)
- [ ] Flag `-x`: `third_party/bash/builtins/set.def:L843` (xtrace)
- [ ] Flag `-B`: `third_party/bash/builtins/set.def:L843` (braceexpand)
- [ ] Flag `-C`: `third_party/bash/builtins/set.def:L843` (noclobber)
- [ ] Flag `-E`: `third_party/bash/builtins/set.def:L843` (errtrace)
- [ ] Flag `-H`: `third_party/bash/builtins/set.def:L843` (histexpand)
- [ ] Flag `-P`: `third_party/bash/builtins/set.def:L843` (physical)
- [ ] Flag `-T`: `third_party/bash/builtins/set.def:L843` (functrace)

- [x] Upstream: `third_party/coreutils/src/coreutils-sha1sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha224sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha256sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha384sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha512sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/sum/sum.go`

### `shift`

- [x] Upstream: `third_party/bash/builtins/shift.def`
- [x] Basic shift: Implemented in `internal/commands/shift/shift.go`
- [x] Shifting n parameters: `internal/commands/shift/shift.go`

### `shopt`

- [ ] Upstream: `third_party/bash/builtins/shopt.def`
- [ ] Basic option management: Missing implementation
- [ ] Flag `-o`: `third_party/bash/builtins/shopt.def:L65` (restrict to set -o)
- [ ] Flag `-p`: `third_party/bash/builtins/shopt.def:L77` (print status)
- [ ] Flag `-q`: `third_party/bash/builtins/shopt.def:L71` (quiet)
- [ ] Flag `-s`: `third_party/bash/builtins/shopt.def:L62` (enable)
- [ ] Flag `-u`: `third_party/bash/builtins/shopt.def:L65` (disable)

### `shred`

- [ ] Upstream: `third_party/coreutils/src/shred.c`
- [ ] Data erasure: Missing implementation
- [ ] Flag `-f`, `--force`: `third_party/coreutils/src/shred.c:L157`
- [ ] Flag `-n`, `--iterations=N`: `third_party/coreutils/src/shred.c:L158`
- [ ] Flag `-s`, `--size=N`: `third_party/coreutils/src/shred.c:L159`
- [ ] Flag `-u`, `--remove`: `third_party/coreutils/src/shred.c:L160`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/shred.c:L161`
- [ ] Flag `-x`, `--exact`: `third_party/coreutils/src/shred.c:L162`
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/shred.c:L163`

### `shuf`

- [x] Upstream: `third_party/coreutils/src/shuf.c`
- [x] Basic shuffling: Implemented in `internal/commands/shuf/shuf.go`
- [ ] Flag `--random-source=FILE`: `third_party/coreutils/src/shuf.c:L111`
- [x] Flag `-e`: `internal/commands/shuf/shuf.go`
- [x] Flag `-i`: `internal/commands/shuf/shuf.go`
- [x] Flag `-n`: `internal/commands/shuf/shuf.go`
- [x] Flag `-o`: `internal/commands/shuf/shuf.go`
- [x] Flag `-r`: `internal/commands/shuf/shuf.go`
- [x] Flag `-z`: `internal/commands/shuf/shuf.go`

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
- [ ] Flag `-k`, `--key=KEYDEF`: `third_party/coreutils/src/sort.c:L473`
- [ ] Flag `-m`, `--merge`: `third_party/coreutils/src/sort.c:L477`
- [x] Flag `-M`, `--month-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-n`, `--numeric-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-o`, `--output=FILE`: `internal/commands/sort/sort.go`
- [x] Flag `-r`, `--reverse`: `internal/commands/sort/sort.go`
- [ ] Flag `-s`, `--stable`: `third_party/coreutils/src/sort.c:L497`
- [ ] Flag `-S`, `--buffer-size=SIZE`: `third_party/coreutils/src/sort.c:L501`
- [ ] Flag `-t`, `--field-separator=SEP`: `third_party/coreutils/src/sort.c:L505`
- [ ] Flag `-T`, `--temporary-directory=DIR`: `third_party/coreutils/src/sort.c:L509`
- [x] Flag `-u`, `--unique`: `internal/commands/sort/sort.go`
- [x] Flag `-V`, `--version-sort`: `internal/commands/sort/sort.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/sort/sort.go`
- [ ] Flag `--parallel=N`: `third_party/coreutils/src/sort.c:L525`
- [x] Flag `--random-sort` (`-R`): `internal/commands/sort/sort.go`
- [ ] Flag `--debug`: `third_party/coreutils/src/sort.c:L533`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/sort.c:L537`

### `source`

- [x] Upstream: `third_party/bash/builtins/source.def`
- [x] Basic sourcing: Implemented in `internal/commands/source/source.go`
- [x] Aliases: `.`

### `split`

- [x] Basic split: Implemented in `internal/commands/split/split.go`
- [ ] Flag `--filter`: `third_party/coreutils/src/split.c:L274`
- [x] Flag `--verbose`: `internal/commands/split/split.go`
- [ ] Flag `-C`: `third_party/coreutils/src/split.c:L250`
- [x] Flag `-a`: `internal/commands/split/split.go`
- [x] Flag `-b`: `internal/commands/split/split.go`
- [x] Flag `-d`: `internal/commands/split/split.go`
- [ ] Flag `-e`: `third_party/coreutils/src/split.c:L270`
- [x] Flag `-l`: `internal/commands/split/split.go`
- [ ] Flag `-n`: `third_party/coreutils/src/split.c:L282`
- [ ] Flag `-t`: `third_party/coreutils/src/split.c:L286`
- [ ] Flag `-u`: `third_party/coreutils/src/split.c:L291`
- [ ] Flag `-x`: `third_party/coreutils/src/split.c:L262`

### `stat`

- [x] Upstream: `third_party/coreutils/src/stat.c`
- [x] Basic output: Implemented in `internal/commands/stat/stat.go`
- [x] Flag `-c`, `--format=FORMAT`: `internal/commands/stat/stat.go`
- [x] Flag `-f`, `--file-system`: `internal/commands/stat/stat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-L`, `--dereference`: `internal/commands/stat/stat.go`
- [x] Flag `-t`, `--terse`: `internal/commands/stat/stat.go`
- [ ] Flag `--printf=FORMAT`: `third_party/coreutils/src/stat.c:1924`
- [ ] Flag `--cached={always,never,default}`: `third_party/coreutils/src/stat.c:1917`

### `stdbuf`

- [ ] Upstream: `third_party/coreutils/src/stdbuf.c`
- [ ] Flag `-e`, `--error=MODE`: `third_party/coreutils/src/stdbuf.c:L130`
- [ ] Flag `-i`, `--input=MODE`: `third_party/coreutils/src/stdbuf.c:L134`
- [ ] Flag `-o`, `--output=MODE`: `third_party/coreutils/src/stdbuf.c:L138`

### `stty`

- [ ] Upstream: `third_party/coreutils/src/stty.c`
- [ ] Flag `-F`, `--file=DEVICE`: `third_party/coreutils/src/stty.c:L1022`
- [ ] Flag `-a`, `--all`: `third_party/coreutils/src/stty.c:L1026`
- [ ] Flag `-g`, `--save`: `third_party/coreutils/src/stty.c:1030`

### `sum`

- [ ] Upstream: `third_party/coreutils/src/sum.c`
- [ ] Flag `-r`: `third_party/coreutils/src/sum.c:L142` (BSD algorithm)
- [ ] Flag `-s`, `--sysv`: `third_party/coreutils/src/sum.c:L146` (System V algorithm)

### `suspend`

- [ ] Upstream: `third_party/bash/builtins/suspend.def`
- [ ] Flag `-f`: `third_party/bash/builtins/suspend.def:L64` (force)

### `sync`

- [x] Upstream: `third_party/coreutils/src/sync.c`
- [x] Basic operation: Implemented in `internal/commands/sync/sync.go`
- [ ] Flag `-d`, `--data`: `third_party/coreutils/src/sync.c:L129`
- [ ] Flag `-f`, `--file-system`: `third_party/coreutils/src/sync.c:L132`

### `tac`

- [x] Upstream: `third_party/coreutils/src/tac.c`
- [x] Basic output: Implemented in `internal/commands/tac/tac.go`
- [ ] Flag `-b`: `third_party/coreutils/src/tac.c:L103`
- [ ] Flag `-r`: `third_party/coreutils/src/tac.c:L104`
- [ ] Flag `-s`: `third_party/coreutils/src/tac.c:L105`

### `tail`

- [x] Basic output: Implemented in `internal/commands/tail/tail.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/tail/tail.go`
- [ ] Flag `-f`, `--follow[={name|descriptor}]`: `third_party/coreutils/src/tail.c:L305`
- [ ] Flag `-F`: `third_party/coreutils/src/tail.c:L309` (implies --follow=name --retry)
- [x] Flag `-n`, `--lines`: `internal/commands/tail/tail.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/tail/tail.go`
- [ ] Flag `-s`, `--sleep-interval`: `third_party/coreutils/src/tail.c:L322`
- [x] Flag `-v`, `--verbose`: `internal/commands/tail/tail.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/tail/tail.go`
- [ ] Flag `--pid`: `third_party/coreutils/src/tail.c:L334`
- [ ] Flag `--retry`: `third_party/coreutils/src/tail.c:L338`
- [ ] Flag `--max-unchanged-stats`: `third_party/coreutils/src/tail.c:L342`

### `tee`

- [x] Basic copy: Implemented in `internal/commands/tee/tee.go`
- [x] Flag `-a`, `--append`: `internal/commands/tee/tee.go`
- [ ] Flag `-i`, `--ignore-interrupts`: `third_party/coreutils/src/tee.c:L97`
- [ ] Flag `-p`, `--output-error[=MODE]`: `third_party/coreutils/src/tee.c:L101`

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
- [ ] Flag `--help`: `third_party/coreutils/src/timeout.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/timeout.c:L44`
- [ ] Flag `-k`: `third_party/coreutils/src/timeout.c:L531`
- [ ] Flag `-s`: `third_party/coreutils/src/timeout.c:L539`

### `times`

- [ ] Upstream: `third_party/bash/builtins/times.def`
- [ ] Basic output: Missing implementation

### `touch`

- [x] Upstream: `third_party/coreutils/src/touch.c`
- [x] Basic touch: Implemented in `internal/commands/touch/touch.go`
- [x] Basic timestamp update: Implemented in `internal/commands/touch/touch.go`
- [ ] Flag `-t STAMP`: `third_party/coreutils/src/touch.c:L259` (explicit timestamp)
- [x] Flag `-a`: `internal/commands/touch/touch.go`
- [x] Flag `-c`: `internal/commands/touch/touch.go`
- [x] Flag `-d`: `internal/commands/touch/touch.go`
- [x] Flag `-h`: `internal/commands/touch/touch.go`
- [x] Flag `-m`: `internal/commands/touch/touch.go`
- [x] Flag `-r`: `internal/commands/touch/touch.go`
- [ ] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `third_party/coreutils/src/touch.c:L259`

### `tr`

- [x] Upstream: `third_party/coreutils/src/tr.c`
- [x] Basic translation: Implemented in `internal/commands/tr/tr.go`
- [x] Flag `-c`, `-C`, `--complement`: `internal/commands/tr/tr.go`
- [x] Flag `-d`, `--delete`: `internal/commands/tr/tr.go`
- [x] Flag `-s`, `--squeeze-repeats`: `internal/commands/tr/tr.go`
- [x] Flag `-t`, `--truncate-set1`: `internal/commands/tr/tr.go`

### `trap`

- [ ] Upstream: `third_party/bash/builtins/trap.def`
- [ ] Basic trapping: Missing implementation
- [ ] Flag `-P`: `third_party/bash/builtins/trap.def:L131`
- [ ] Flag `-l`: `third_party/bash/builtins/trap.def:L125`
- [ ] Flag `-p`: `third_party/bash/builtins/trap.def:L128`

### `true`

- [x] Basic operation: Implemented in `internal/commands/boolcmd/bool.go`

### `truncate`

- [x] Upstream: `third_party/coreutils/src/truncate.c`
- [x] Basic truncation: Implemented in `internal/commands/truncate/truncate.go`
- [x] Flag `-c`: `internal/commands/truncate/truncate.go`
- [ ] Flag `-o`: `third_party/coreutils/src/truncate.c:L85`
- [x] Flag `-r`: `internal/commands/truncate/truncate.go`
- [x] Flag `-s`: `internal/commands/truncate/truncate.go`

### `tsort`

- [x] Upstream: `third_party/coreutils/src/tsort.c`
- [x] Topological sort: Implemented in `internal/commands/tsort/tsort.go`
- [ ] Flag `--help`: `third_party/coreutils/src/tsort.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/tsort.c:L52`

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

- [ ] Upstream: `third_party/bash/builtins/ulimit.def`
- [ ] Resource management: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/ulimit.def:L35` (all)
- [ ] Flag `-c`: `third_party/bash/builtins/ulimit.def:L37` (core)
- [ ] Flag `-d`: `third_party/bash/builtins/ulimit.def:L38` (data)
- [ ] Flag `-e`: `third_party/bash/builtins/ulimit.def:L39` (priority)
- [ ] Flag `-f`: `third_party/bash/builtins/ulimit.def:L40` (file size)
- [ ] Flag `-n`: `third_party/bash/builtins/ulimit.def:L45` (opened files)
- [ ] Flag `-u`: `third_party/bash/builtins/ulimit.def:L51` (user processes)
- [ ] Flag `-S`: `third_party/bash/builtins/ulimit.def:L33` (soft)
- [ ] Flag `-H`: `third_party/bash/builtins/ulimit.def:L34` (hard)

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
- [ ] Flag `--first-only`: `third_party/coreutils/src/unexpand.c:L91`

### `uniq`

- [x] Upstream: `third_party/coreutils/src/uniq.c`
- [x] Basic filtering: Implemented in `internal/commands/uniq/uniq.go`
- [x] Flag `-c`, `--count`: `internal/commands/uniq/uniq.go`
- [x] Flag `-d`, `--repeated`: `internal/commands/uniq/uniq.go`
- [ ] Flag `-D`, `--all-repeated[=METHOD]`: `third_party/coreutils/src/uniq.c:180`
- [ ] Flag `-f`, `--skip-fields=N`: `third_party/coreutils/src/uniq.c:189`
- [x] Flag `-i`, `--ignore-case`: `internal/commands/uniq/uniq.go`
- [ ] Flag `-s`, `--skip-chars=N`: `third_party/coreutils/src/uniq.c:202`
- [x] Flag `-u`, `--unique`: `internal/commands/uniq/uniq.go`
- [ ] Flag `-w`, `--check-chars=N`: `third_party/coreutils/src/uniq.c:214`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/uniq/uniq.go`
- [ ] Flag `--group[=METHOD]`: `third_party/coreutils/src/uniq.c:184`

### `unlink`

- [x] Basic removal: Implemented in `internal/commands/unlink/unlink.go` (exactly 1 arg required)
- [x] Flag `--help`: `internal/commands/unlink/unlink.go`
- [x] Flag `--version`: `internal/commands/unlink/unlink.go`

### `unset`

- [x] Upstream: `third_party/bash/builtins/set.def`
- [x] Attribute management: Implemented in `internal/commands/unset/unset.go`
- [x] Flag `-f`: `internal/commands/unset/unset.go` (functions)
- [x] Flag `-v`: `internal/commands/unset/unset.go` (variables)
- [ ] Flag `-n`: `third_party/bash/builtins/set.def:L640` (nameref)

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

- [ ] Upstream: `third_party/coreutils/src/ls.c`

### `wait`

- [x] Upstream: `third_party/bash/builtins/wait.def`
- [x] Basic waiting: Implemented in `internal/commands/wait/wait.go`
- [x] Optional: jobspec or process ID: `internal/commands/wait/wait.go`
- [x] Flag `-f`: `internal/commands/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#wait))
- [x] Flag `-n`: `internal/commands/wait/wait.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/functional_gap.md#wait))
- [ ] Flag `-p var`: `third_party/bash/builtins/wait.def:L137`

### `wc`

- [x] Basic counts: Implemented in `internal/commands/wc/wc.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/wc/wc.go`
- [x] Flag `-m`, `--chars`: `internal/commands/wc/wc.go`
- [x] Flag `-l`, `--lines`: `internal/commands/wc/wc.go`
- [x] Flag `-w`, `--words`: `internal/commands/wc/wc.go`
- [x] Flag `-L`, `--max-line-length`: `internal/commands/wc/wc.go`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/wc.c:L222`
- [ ] Flag `--total={auto,always,only,never}`: `third_party/coreutils/src/wc.c:L226`

### `who`

- [x] Upstream: `third_party/coreutils/src/who.c`
- [x] Basic output: Implemented in `internal/commands/who/who.go`

- [ ] Upstream: `third_party/coreutils/src/who.c`
- [ ] Flag `-H`, `--heading`: `third_party/coreutils/src/who.c:L592`
- [ ] Flag `-a`, `--all`: `third_party/coreutils/src/who.c:L584`
- [ ] Flag `-b`, `--boot`: `third_party/coreutils/src/who.c:L588`
- [ ] Flag `-d`, `--dead`: `third_party/coreutils/src/who.c:L596`
- [ ] Flag `-l`, `--login`: `third_party/coreutils/src/who.c:L604`
- [ ] Flag `-m`: `third_party/coreutils/src/who.c:L608`
- [ ] Flag `-p`, `--process`: `third_party/coreutils/src/who.c:L616`
- [ ] Flag `-q`, `--count`: `third_party/coreutils/src/who.c:L620`
- [ ] Flag `-r`, `--runlevel`: `third_party/coreutils/src/who.c:L624`
- [ ] Flag `-s`, `--short`: `third_party/coreutils/src/who.c:L628`
- [ ] Flag `-t`, `--time`: `third_party/coreutils/src/who.c:L632`
- [ ] Flag `-u`, `--users`: `third_party/coreutils/src/who.c:L636`

### `whoami`

- [x] Upstream: `third_party/coreutils/src/whoami.c`
- [x] Basic output: Implemented in `internal/commands/whoami/whoami.go`
- [ ] Flag `--help`: `third_party/coreutils/src/whoami.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/whoami.c:L44`

### `yes`

- [x] Upstream: `third_party/coreutils/src/yes.c`
- [x] Basic operation: Implemented in `internal/commands/yes/yes.go`
- [ ] Basic repetition: Missing implementation
- [ ] Flag `--help`: `third_party/coreutils/src/yes.c:L48`
- [ ] Flag `--version`: `third_party/coreutils/src/yes.c:L48`


## Shell Keywords & Grammar

### `!`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pipeline negation: Missing implementation

### `[[`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Conditional expressions: Missing implementation
- [ ] Pattern matching (`==`, `!=`): Missing implementation
- [ ] Regex matching (`=~`): Missing implementation
- [ ] Aliases: `]]`

### `((`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Arithmetic evaluation: Missing implementation
- [ ] Aliases: `))`

### `{`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Command grouping: Missing implementation
- [ ] Aliases: `}`

### `case`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pattern-based branching: Missing implementation

### `coproc`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Asynchronous coprocesses: Missing implementation

### `for`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] List-based iteration: Missing implementation
- [ ] C-style arithmetic iteration (`for ((`): Missing implementation

### `function`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Shell function definition: Missing implementation

### `if`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Conditional branching (if/then/elif/else/fi): Missing implementation

### `until`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Negative condition looping: Missing implementation

### `while`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Positive condition looping: Missing implementation

## Shell Variables

### `BASH_VERSION`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Version information string: Missing implementation

### `CDPATH`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Search path for `cd` command: Missing implementation

### `GLOBIGNORE`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pattern-based pathname expansion ignore: Missing implementation

### `HISTFILE`, `HISTFILESIZE`, `HISTSIZE`, `HISTIGNORE`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] History management persistence: Missing implementation

### `HOME`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Logical login directory: Missing implementation

### `HOSTNAME`, `HOSTTYPE`, `MACHTYPE`, `OSTYPE`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] System identity metadata: Missing implementation

### `IGNOREEOF`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] EOF handling for interactive shells: Missing implementation

### `MAILCHECK`, `MAILPATH`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Mail notification settings: Missing implementation

### `PATH`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Command search path: Missing implementation

### `PROMPT_COMMAND`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Pre-prompt execution hook: Missing implementation

### `PS1`, `PS2`, `PS3`, `PS4`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Interactive prompt formatting: Missing implementation

### `PWD`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Logical current directory tracking: Missing implementation

### `SHELLOPTS`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] List of enabled shell options: Missing implementation

### `TERM`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Terminal environment identification: Missing implementation

### `TIMEFORMAT`
- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Output format for `time` reserved word: Missing implementation

## Interactive Shell Features

- [ ] Interactive history navigation (Up/Down arrow keys)
- [ ] Command line editing (Backspace, Ctrl+L, etc.)
- [ ] Tab completion: Missing implementation

## Shell Expansions

### Parameter Expansion
- [ ] Basic expansion `${var}`: Missing implementation
- [ ] Substring expansion `${var:offset:length}`: `third_party/bash/subst.c:L10170`
- [ ] Prefix removal `${var#pattern}`, `${var##pattern}`: `third_party/bash/subst.c:L10313`
- [ ] Suffix removal `${var%pattern}`, `${var%%pattern}`: `third_party/bash/subst.c:L10314`
- [ ] Substring replacement `${var/pattern/string}`: `third_party/bash/subst.c:L10205`
- [ ] Case modification `${var^}`, `${var^^}`, `${var,}`, `${var,,}`: `third_party/bash/subst.c:L10234`
- [ ] Default values `${var:-default}`, `${var:=default}`: `third_party/bash/subst.c:L10338-10341`
- [ ] Alternative/Error values `${var:?error}`, `${var:+alternative}`: `third_party/bash/subst.c:L10338-10341`

### Command Substitution
- [ ] Basic substitution $(command), `command`: `third_party/bash/subst.c:L11000`

### Arithmetic Expansion
- [ ] Basic expansion $( (expression) ): `third_party/bash/subst.c:L10825`

### Brace Expansion
- [ ] basic expansion {a,b,c}: `third_party/bash/braces.c`

### Tilde Expansion
- [ ] basic expansion ~, ~user: `third_party/bash/subst.c:L10740`

## Redirections

### Standard Redirections
- [ ] Input redirection `[n]<word`: `third_party/bash/redir.c:L897`
- [ ] Output redirection `[n]>word`: `third_party/bash/redir.c:L895`
- [ ] Append redirection `[n]>>word`: `third_party/bash/redir.c:L896`
- [ ] Force output `[n]>|word`: `third_party/bash/redir.c:L902`
- [ ] Combined stderr/stdout `&>word`, `&>>word`: `third_party/bash/redir.c:L899-900`

### File Descriptor Manipulation
- [ ] Duplicating input `[n]<&word`: `third_party/bash/redir.c:L1115`
- [ ] Duplicating output `[n]>&word`: `third_party/bash/redir.c:L1116`
- [ ] Moving input `[n]<&digit-`: `third_party/bash/redir.c:L1117`
- [ ] Moving output `[n]>&digit-`: `third_party/bash/redir.c:L1118`

### Advanced Redirections
- [ ] Here-Documents `[n]<<[-]word`: `third_party/bash/redir.c:L1042`
- [ ] Here-Strings `[n]<<<word`: `third_party/bash/redir.c:L1044`
- [ ] Process Substitution `<(list)`, `>(list)`: `third_party/bash/subst.c:L321`

## Globbing Patterns

### Standard Wildcards
- [ ] Match any string `*`: `third_party/bash/lib/glob/glob.c`
- [ ] Match any character `?`: `third_party/bash/lib/glob/glob.c`

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
- [ ] Basic pipe `|`: `third_party/bash/execute_cmd.c:L191`
- [ ] Combined stderr pipe `|&`: `third_party/bash/execute_cmd.c:L191`

### Compound Commands & Lists
- [ ] Sequential list `;`: `third_party/bash/execute_cmd.c:L193`
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
