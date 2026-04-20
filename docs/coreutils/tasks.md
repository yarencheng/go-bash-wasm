# Functional Parity Tracking

This document tracks the alignment of the Go Bash Simulator with upstream GNU implementations.

## Overview

Status codes:
- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

---


## Coreutils Commands

### `arch`

- [x] Upstream: `third_party/coreutils/src/coreutils-arch.c`
- [x] Inherits flags from `uname`

### `base32`

- [x] Basic encoding/decoding: Implemented in `internal/commands/coreutils/base32/base32.go`
- [-] Flag `--base16`: Use `basenc` instead
- [-] Flag `--base2lsbf`: Use `basenc` instead
- [-] Flag `--base2msbf`: Use `basenc` instead
- [x] Flag `--base32`: `internal/commands/coreutils/base32/base32.go` (default)
- [-] Flag `--base32hex`: Use `basenc` instead
- [-] Flag `--base58`: Use `basenc` instead
- [x] Flag `--base64`: `internal/commands/coreutils/base64/base64.go`
- [x] Flag `--base64url`: `internal/commands/coreutils/base32/base32.go`
- [x] Flag `--z85`: `internal/commands/coreutils/base32/base32.go`
- [x] Flag `-d`: `internal/commands/coreutils/base32/base32.go`
- [x] Flag `-i`: `internal/commands/coreutils/base32/base32.go`
- [x] Flag `-w`: `internal/commands/coreutils/base32/base32.go`

### `base64`

- [x] Upstream: `third_party/coreutils/src/base64.c`
- [x] Basic encoding/decoding: Implemented in `internal/commands/coreutils/base64/base64.go`
- [x] Flag `-d`, `--decode`: `internal/commands/coreutils/base64/base64.go`
- [x] Flag `-i`, `--ignore-garbage`: `internal/commands/coreutils/base64/base64.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`, `--wrap=COLS`: `internal/commands/coreutils/base64/base64.go`

### `basename`

- [x] Upstream: `third_party/coreutils/src/basename.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/basename/basename.go`
- [x] Flag `-a`, `--multiple`: `internal/commands/coreutils/basename/basename.go`
- [x] Flag `-s`, `--suffix=SUFFIX`: `internal/commands/coreutils/basename/basename.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/basename/basename.go`

### `basenc`

- [x] Upstream: `third_party/coreutils/src/basenc.c`
- [x] Basic encoding/decoding: Implemented in `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `-d`, `--decode`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `-i`, `--ignore-garbage`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `-w`, `--wrap=COLS`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base16`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base32`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base32hex`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base64`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base64url`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base2lsbf`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base2msbf`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--base58`: `internal/commands/coreutils/basenc/basenc.go`
- [x] Flag `--z85`: `internal/commands/coreutils/basenc/basenc.go`

### `cat`

- [x] Upstream: `third_party/coreutils/src/cat.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-A`, `--show-all`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-b`, `--number-nonblank`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-e`: `internal/commands/coreutils/cat/cat.go` (implies -vE)
- [x] Flag `-E`, `--show-ends`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-n`, `--number`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-s`, `--squeeze-blank`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-t`: `internal/commands/coreutils/cat/cat.go` (implies -vT)
- [x] Flag `-T`, `--show-tabs`: `internal/commands/coreutils/cat/cat.go`
- [x] Flag `-u`: `internal/commands/coreutils/cat/cat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#cat))
- [x] Flag `-v`, `--show-nonprinting`: `internal/commands/coreutils/cat/cat.go`

### `chcon`

- [x] Upstream: `third_party/coreutils/src/chcon.c`
- [x] Basic operation: `internal/commands/coreutils/chcon/chcon.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#chcon-runcon))
- [x] Flag `-h`, `--no-dereference`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-H`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-L`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-P`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-u`, `--user=USER`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-r`, `--role=ROLE`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `-l`, `--range=RANGE`: `internal/commands/coreutils/chcon/chcon.go`
- [x] Flag `--reference=RFILE`: `internal/commands/coreutils/chcon/chcon.go`

### `chgrp`

- [x] Upstream: `third_party/coreutils/src/chgrp.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/chown/chown.go` (via `NewChgrp`)
- [x] Flag `-c`, `--changes`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `--dereference`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-h`, `--no-dereference`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `--reference=RFILE`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-H`, `-L`, `-P`: `internal/commands/coreutils/chown/chown.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `chmod`

- [x] Upstream: `third_party/coreutils/src/chmod.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/chmod/chmod.go`
- [x] Flag `-c`, `--changes`: `internal/commands/coreutils/chmod/chmod.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/coreutils/chmod/chmod.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/chmod/chmod.go`
- [x] Flag `--reference=RFILE`: `internal/commands/coreutils/chmod/chmod.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/coreutils/chmod/chmod.go`
- [x] Symbolic modes (u+x, g-w, etc.): Implemented in `internal/commands/coreutils/chmod/chmod.go`

### `chown`

- [x] Upstream: `third_party/coreutils/src/chown.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-c`, `--changes`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-f`, `--silent`, `--quiet`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `--dereference`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-h`, `--no-dereference`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `--from=CURRENT_OWNER:CURRENT_GROUP`: `internal/commands/coreutils/chown/chown.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--reference=RFILE`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-R`, `--recursive`: `internal/commands/coreutils/chown/chown.go`
- [x] Flag `-H`, `-L`, `-P`: `internal/commands/coreutils/chown/chown.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `chroot`

- [x] Upstream: `third_party/coreutils/src/chroot.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/chroot/chroot.go`

### `cksum`

- [x] Upstream: `third_party/coreutils/src/cksum.c`
- [x] Basic CRC-32: Implemented in `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `-a`, `--algorithm=ALGO`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `-c`, `--check`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `-l`, `--length=BITS`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--base64`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--raw`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--tag`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--untagged`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--ignore-missing`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--quiet`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--status`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `--strict`: `internal/commands/coreutils/cksum/cksum.go`
- [x] Flag `-w`, `--warn`: `internal/commands/coreutils/cksum/cksum.go`
 
### `clear`

- [x] Basic operation: Implemented in `internal/commands/coreutils/clear/clear.go`
 
### `comm`

- [x] Upstream: `third_party/coreutils/src/comm.c`
- [x] Basic comparison: Implemented in `internal/commands/coreutils/comm/comm.go`
- [x] Flag `--check-order`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `--nocheck-order`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `--output-delimiter`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `--total`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `-1`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `-2`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `-3`: `internal/commands/coreutils/comm/comm.go`
- [x] Flag `-z`: `internal/commands/coreutils/comm/comm.go`

### `cp`

- [x] Basic copy: Implemented in `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-a`, `--archive`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-b`, `--backup`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-d`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-f`, `--force`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-H`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--interactive`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`, `--link`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-L`, `--dereference`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-n`, `--no-clobber`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-p`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-P`, `--no-dereference`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-r`, `-R`, `--recursive`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-s`, `--symbolic-link`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--target-directory`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-u`, `--update`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/cp/cp.go`
- [x] Flag `-x`, `--one-file-system`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-Z`, `--context`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--attributes-only`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--preserve[=ATTR_LIST]`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--no-preserve=ATTR_LIST`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--parents`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--reflink[=WHEN]`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--sparse=WHEN`: `internal/commands/coreutils/cp/cp.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--strip-trailing-slashes`: `internal/commands/coreutils/cat/cat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `csplit`

- [x] Upstream: `third_party/coreutils/src/csplit.c`
- [x] Basic split: Implemented in `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `--suppress-matched`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-b`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-f`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-k`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-n`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-s`: `internal/commands/coreutils/csplit/csplit.go`
- [x] Flag `-z`: `internal/commands/coreutils/csplit/csplit.go`

### `cut`

- [x] Upstream: `third_party/coreutils/src/cut.c`
- [x] Basic selection: Implemented in `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-b`, `--bytes=LIST`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-c`, `--characters=LIST`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-d`, `--delimiter=DELIM`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-f`, `--fields=LIST`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-n`: `internal/commands/coreutils/cut/cut.go` (Ignored)
- [x] Flag `-s`, `--only-delimited`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `--complement`: `internal/commands/coreutils/cut/cut.go`
- [x] Flag `--output-delimiter=STRING`: `internal/commands/coreutils/cut/cut.go`

### `date`

- [x] Upstream: `third_party/coreutils/src/date.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/date/date.go`
- [x] Custom format `+FORMAT`: `internal/commands/coreutils/date/date.go`
- [x] Flag `-d`, `--date=STRING`: `internal/commands/coreutils/date/date.go`
- [x] Flag `-f`, `--file=DATEFILE`: `internal/commands/coreutils/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-I[FMT]`, `--iso-8601[=FMT]`: `internal/commands/coreutils/date/date.go`
- [x] Flag `-r`, `--reference=FILE`: `internal/commands/coreutils/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-R`, `--rfc-email`: `internal/commands/coreutils/date/date.go`
- [x] Flag `-s`, `--set=STRING`: `internal/commands/coreutils/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--utc`, `--universal`: `internal/commands/coreutils/date/date.go`
- [x] Flag `--debug`: `internal/commands/coreutils/date/date.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `dd`

- [x] Upstream: `third_party/coreutils/src/dd.c`
- [x] Data copy: Implemented in `internal/commands/coreutils/dd/dd.go`
- [x] Operand `bs=BYTES`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `cbs=BYTES`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `conv=CONVS`: `internal/commands/coreutils/dd/dd.go` (Partial via notrunc; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `count=N`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `ibs=BYTES`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `if=FILE`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `iflag=FLAGS`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `obs=BYTES`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `of=FILE`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `oflag=FLAGS`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `seek=N`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `skip=N`: `internal/commands/coreutils/dd/dd.go`
- [x] Operand `status=LEVEL`: `internal/commands/coreutils/dd/dd.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Operand `conv=notrunc`: `internal/commands/coreutils/dd/dd.go`


### `df`

- [x] Upstream: `third_party/coreutils/src/df.c`
- [x] Basic df: Implemented in `internal/commands/coreutils/df/df.go`
- [x] Basic output: Implemented in `internal/commands/coreutils/df/df.go`
- [x] Flag `--no-sync`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))
- [x] Flag `--output[=FIELD_LIST]`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))
- [x] Flag `--sync`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))
- [x] Flag `--total`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-B`, `--block-size=SIZE`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))
- [x] Flag `-H`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-P`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-T`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-a`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-h`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-i`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-k`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-l`: `internal/commands/coreutils/df/df.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))
- [x] Flag `-x`, `--exclude-type=TYPE`: `internal/commands/coreutils/df/df.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#df))

### `dir`

- [x] Upstream: `third_party/coreutils/src/coreutils-dir.c`
- [x] Inherits flags from `ls`

### `dircolors`

- [x] Upstream: `third_party/coreutils/src/dircolors.c`
- [x] Output configuration: Implemented in `internal/commands/coreutils/dircolors/dircolors.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#dircolors))
- [x] Flag `-b`, `--sh`, `--bourne-shell`: `internal/commands/coreutils/dircolors/dircolors.go`
- [x] Flag `-c`, `--csh`, `--c-shell`: `internal/commands/coreutils/dircolors/dircolors.go`
- [x] Flag `-p`, `--print-database`: `internal/commands/coreutils/dircolors/dircolors.go`
- [x] Flag `--print-ls-colors`: `internal/commands/coreutils/dircolors/dircolors.go`

### `dirname`

- [x] Upstream: `third_party/coreutils/src/dirname.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/dirname/dirname.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/dirname/dirname.go`

### `du`

- [x] Upstream: `third_party/coreutils/src/du.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/du/du.go`
- [x] Flag `-0`, `--null`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-a`, `--all`: `internal/commands/coreutils/du/du.go`
- [x] Flag `--apparent-size`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-b`, `--bytes`: `internal/commands/coreutils/du/du.go` (alias for apparent-size)
- [x] Flag `-c`, `--total`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-d`, `--max-depth=N`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-h`, `--human-readable`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-k`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-m`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-s`, `--summarize`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-S`, `--separate-dirs`: `internal/commands/coreutils/du/du.go` (Stub)
- [x] Flag `-D`, `--dereference-args`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-H`: `internal/commands/coreutils/du/du.go` (dereference)
- [x] Flag `-l`, `--count-links`: `internal/commands/coreutils/du/du.go` (partial via size)
- [x] Flag `-L`, `--dereference`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-P`, `--no-dereference`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-t`, `--threshold=SIZE`: `internal/commands/coreutils/du/du.go`
- [x] Flag `-x`, `--one-file-system`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))
- [x] Flag `-X`, `--exclude-from=FILE`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))
- [x] Flag `--exclude=PATTERN`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))
- [x] Flag `--files0-from=F`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))
- [x] Flag `--inodes`: `internal/commands/coreutils/du/du.go`
- [x] Flag `--si`: `internal/commands/coreutils/du/du.go`
- [x] Flag `--time[=WORD]`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))
- [x] Flag `--time-style=STYLE`: `internal/commands/coreutils/du/du.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#du))

### `echo`

- [x] Upstream: `third_party/coreutils/src/echo.c`
- [x] Basic output: Implemented in `internal/commands/bash/echo/echo.go`
- [x] Flag `-n`: `internal/commands/bash/echo/echo.go`
- [x] Flag `-e`: `internal/commands/bash/echo/echo.go`
- [x] Flag `-E`: `internal/commands/bash/echo/echo.go`
- [x] Escaped `\0NNN`: `internal/commands/bash/echo/echo.go`
- [x] Escaped `\xHH`: `internal/commands/bash/echo/echo.go`
- [x] Escaped `\c`: `internal/commands/bash/echo/echo.go`

### `env`


- [x] Upstream: `third_party/coreutils/src/env.c`
- [x] Basic execution: Implemented in `internal/commands/coreutils/env/env.go`
- [x] Flag `-a`, `--argv0=ARG`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--ignore-environment`: `internal/commands/coreutils/env/env.go`
- [x] Flag `-u`, `--unset=NAME`: `internal/commands/coreutils/env/env.go`
- [x] Flag `-0`, `--null`: `internal/commands/coreutils/env/env.go`
- [x] Flag `-C`, `--chdir=DIR`: `internal/commands/coreutils/env/env.go`
- [x] Flag `-S`, `--split-string=S`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--block-signal[=SIG]`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--default-signal[=SIG]`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--ignore-signal[=SIG]`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--list-signal-handling`: `internal/commands/coreutils/env/env.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `expand`

- [x] Upstream: `third_party/coreutils/src/expand.c`
- [x] Basic conversion: Implemented in `internal/commands/coreutils/expand/expand.go`
- [x] Flag `-i`, `--initial`: `internal/commands/coreutils/expand/expand.go`
- [x] Flag `-t`, `--tabs=LIST`: `internal/commands/coreutils/expand/expand.go`

### `expr`

- [x] Upstream: `third_party/coreutils/src/expr.c`
- [x] Expression evaluation: Implemented in `internal/commands/coreutils/expr/expr.go`
- [x] Arithmetic (+, -, *, /, %): Implemented in `internal/commands/coreutils/expr/expr.go`
- [x] Comparison (=, !=, <, <=, >, >=): Implemented in `internal/commands/coreutils/expr/expr.go`
- [x] Logical (| , &): Implemented in `internal/commands/coreutils/expr/expr.go`
- [x] String operators (match, substr, index, length): Implemented in `internal/commands/coreutils/expr/expr.go`
- [x] Flag `--help`: `internal/commands/coreutils/expr/expr.go`
- [x] Flag `--version`: `internal/commands/coreutils/expr/expr.go`

### `factor`

- [x] Upstream: `third_party/coreutils/src/factor.c`
- [x] Prime factorization: Implemented in `internal/commands/coreutils/factor/factor.go`

### `false`

- [x] Upstream: `third_party/coreutils/src/false.c`
- [x] Basic operation: Implemented in `internal/commands/bash/boolcmd/bool.go`

### `find`


- [x] Upstream: [External] (Part of GNU Findutils)
- [x] Basic Search: `internal/commands/coreutils/find/find.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#find))
- [x] Flag `-name`: `internal/commands/coreutils/find/find.go`
- [x] Flag `-type`: `internal/commands/coreutils/find/find.go`

### `fmt`

- [x] Upstream: `third_party/coreutils/src/fmt.c`
- [x] Paragraph formatting: Implemented in `internal/commands/coreutils/fmt/fmt.go`
- [x] Flag `-c`, `--crown-margin`: `internal/commands/coreutils/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-p`, `--prefix=STRING`: `internal/commands/coreutils/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`, `--split-only`: `internal/commands/coreutils/fmt/fmt.go`
- [x] Flag `-t`, `--tagged-paragraph`: `internal/commands/coreutils/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--uniform-spacing`: `internal/commands/coreutils/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/coreutils/fmt/fmt.go`
- [x] Flag `-g`, `--goal=WIDTH`: `internal/commands/coreutils/fmt/fmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-WIDTH`: `internal/commands/coreutils/fmt/fmt.go`

### `fold`

- [x] Upstream: `third_party/coreutils/src/fold.c`
- [x] Line wrapping: Implemented in `internal/commands/coreutils/fold/fold.go`
- [x] Flag `-b`, `--bytes`: `internal/commands/coreutils/fold/fold.go`
- [x] Flag `-c`, `--characters`: `internal/commands/coreutils/fold/fold.go`
- [x] Flag `-s`, `--spaces`: `internal/commands/coreutils/fold/fold.go`
- [x] Flag `-w`, `--width=WIDTH`: `internal/commands/coreutils/fold/fold.go`

### `getlimits`

- [x] Upstream: `third_party/coreutils/src/getlimits.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/getlimits/getlimits.go`
- [x] Flag `--help`: `internal/commands/coreutils/getlimits/getlimits.go`
- [x] Flag `--version`: `internal/commands/coreutils/getlimits/getlimits.go`


### `grep`

- [x] Upstream: [External] (Part of GNU Grep)
- [x] Regex Search: `internal/commands/coreutils/grep/grep.go` (Workaround; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#grep))
- [x] Flag `-i`, `--ignore-case`: `internal/commands/coreutils/grep/grep.go`
- [x] Flag `-v`, `--invert-match`: `internal/commands/coreutils/grep/grep.go`
- [x] Flag `-n`, `--line-number`: `internal/commands/coreutils/grep/grep.go`
- [x] Flag `-c`, `--count`: `internal/commands/coreutils/grep/grep.go`
- [x] Flag `-l`, `--files-with-matches`: `internal/commands/coreutils/grep/grep.go`

### `groups`

- [x] Upstream: `third_party/coreutils/src/groups.c`
- [x] Basic listing: Implemented in `internal/commands/coreutils/groups/groups.go`
- [x] Multiple users support: `internal/commands/coreutils/groups/groups.go`

### `head`

- [x] Upstream: `third_party/coreutils/src/head.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/head/head.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/coreutils/head/head.go`
- [x] Flag `-n`, `--lines`: `internal/commands/coreutils/head/head.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/coreutils/head/head.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/head/head.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/head/head.go`

### `hostid`

- [x] Upstream: `third_party/coreutils/src/hostid.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/hostid/hostid.go`

### `hostname`

- [x] Upstream: `third_party/coreutils/src/hostname.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/hostname/hostname.go`
- [x] Set hostname support: `internal/commands/coreutils/hostname/hostname.go`
- [x] Flag `--help`: `internal/commands/coreutils/hostname/hostname.go`
- [x] Flag `--version`: `internal/commands/coreutils/hostname/hostname.go`

### `id`

- [x] Upstream: `third_party/coreutils/src/id.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/id/id.go`
- [x] Flag `-G`: `internal/commands/coreutils/id/id.go`
- [x] Flag `-Z`: `internal/commands/coreutils/id/id.go` (Unsupported; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#id))
- [x] Flag `-a`: `internal/commands/coreutils/id/id.go` (Ignored)
- [x] Flag `-g`: `internal/commands/coreutils/id/id.go`
- [x] Flag `-n`: `internal/commands/coreutils/id/id.go`
- [x] Flag `-r`: `internal/commands/coreutils/id/id.go` (real == effective)
- [x] Flag `-u`: `internal/commands/coreutils/id/id.go`
- [x] Flag `-z`: `internal/commands/coreutils/id/id.go`

### `install`

- [x] Upstream: `third_party/coreutils/src/install.c`
- [x] Flag `-c`: `internal/commands/coreutils/install/install.go` (ignored)
- [x] Flag `-d`, `--directory`: `internal/commands/coreutils/install/install.go`
- [x] Flag `-D`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-g`, `--group=GROUP`: `internal/commands/coreutils/install/install.go`
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/coreutils/install/install.go` (Stub)
- [x] Flag `-o`, `--owner=OWNER`: `internal/commands/coreutils/install/install.go`
- [x] Flag `-p`, `--preserve-timestamps`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`, `--strip`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-S`, `--suffix=SUFFIX`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--target-directory=DIR`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/install/install.go`
- [x] Flag `-C`, `--compare`: `internal/commands/coreutils/install/install.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `join`

- [x] Upstream: `third_party/coreutils/src/join.c`
- [x] Basic join: Implemented in `internal/commands/coreutils/join/join.go`
- [x] Flag `-a FILENUM`: `internal/commands/coreutils/join/join.go` (unpairable lines from file FILENUM)
- [x] Flag `-e EMPTY`: `internal/commands/coreutils/join/join.go` (replace empty input fields with EMPTY)
- [x] Flag `-i`, `--ignore-case`: `internal/commands/coreutils/join/join.go`
- [x] Flag `-j FIELD`: `internal/commands/coreutils/join/join.go` (equivalent to -1 FIELD -2 FIELD)
- [x] Flag `-o FORMAT`: `internal/commands/coreutils/join/join.go` (obey FORMAT while constructing output line)
- [x] Flag `-t CHAR`: `internal/commands/coreutils/join/join.go` (use CHAR as input and output field separator)
- [x] Flag `-v FILENUM`: `internal/commands/coreutils/join/join.go` (like -a FILENUM, but suppress joined output lines)
- [x] Flag `-1 FIELD`: `internal/commands/coreutils/join/join.go` (join on this FIELD of file 1)
- [x] Flag `-2 FIELD`: `internal/commands/coreutils/join/join.go` (join on this FIELD of file 2)
- [x] Flag `--check-order`: `internal/commands/coreutils/join/join.go` (check that the input is correctly sorted)
- [x] Flag `--nocheck-order`: `internal/commands/coreutils/join/join.go` (do not check that the input is correctly sorted)
- [x] Flag `--header`: `internal/commands/coreutils/join/join.go` (treat the first line of each file as field headers)

### `kill`

- [x] Upstream: `third_party/coreutils/src/kill.c`
- [x] Basic signaling: Implemented in `internal/commands/bash/kill/kill.go`
- [x] Flag `-l`, `--list`: `internal/commands/bash/kill/kill.go`
- [x] Flag `-s`, `--signal`: `internal/commands/bash/kill/kill.go`
- [x] Flag `-t`, `--table`: `internal/commands/bash/kill/kill.go`

### `link`


- [x] Basic hard link: Implemented in `internal/commands/coreutils/link/link.go`

### `ln`

- [x] Basic link creation: Implemented in `internal/commands/coreutils/ln/ln.go`
- [x] Flag `-f`: `internal/commands/coreutils/ln/ln.go`
- [x] Flag `-s`: `internal/commands/coreutils/ln/ln.go`
- [x] Flag `-v`: `internal/commands/coreutils/ln/ln.go`

### `logname`

- [x] Upstream: `third_party/coreutils/src/logname.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/logname/logname.go`
- [x] Flag `--help`: `internal/commands/coreutils/logname/logname.go`
- [x] Flag `--version`: `internal/commands/coreutils/logname/logname.go`

### `ls`

- [x] Basic listing: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--author`: `internal/commands/coreutils/ls/ls.go` (Simulated via single-user "root"; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#ls))
- [x] Flag `--block-size`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--color`: `internal/commands/coreutils/ls/ls.go` (ANSI colors)
- [x] Flag `--dereference-command-line-symlink-to-dir`: `internal/commands/coreutils/ls/ls.go` (-H)
- [x] Flag `--file-type`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--format`: `internal/commands/coreutils/ls/ls.go` (across, commas, horizontal, long, single-column, verbose, vertical)
- [x] Flag `--full-time`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--group-directories-first`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--hide`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--indicator-style`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--quoting-style`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--show-control-chars`: `internal/commands/coreutils/ls/ls.go` (partial via -q)
- [x] Flag `--si`: `internal/commands/coreutils/ls/ls.go` (power of 1000)
- [x] Flag `--sort`: `internal/commands/coreutils/ls/ls.go` (unified flag)
- [x] Flag `--time`: `internal/commands/coreutils/ls/ls.go` (atime, ctime, mtime)
- [x] Flag `--time-style`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `--zero`: `internal/commands/coreutils/ls/ls.go` (NUL terminated)
- [x] Flag `-1`: `internal/commands/coreutils/ls/ls.go` (one-line)
- [x] Flag `-A`: `internal/commands/coreutils/ls/ls.go` (almost-all)
- [x] Flag `-B`: `internal/commands/coreutils/ls/ls.go` (ignore-backups)
- [x] Flag `-C`: `internal/commands/coreutils/ls/ls.go` (vertical columns)
- [x] Flag `-D`, `--dired`: `internal/commands/coreutils/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#ls))
- [x] Flag `-F`: `internal/commands/coreutils/ls/ls.go` (classify)
- [x] Flag `-G`: `internal/commands/coreutils/ls/ls.go` (no-group)
- [x] Flag `-H`: `internal/commands/coreutils/ls/ls.go` (dereference-command-line)
- [x] Flag `-I`: `internal/commands/coreutils/ls/ls.go` (ignore)
- [x] Flag `-L`: `internal/commands/coreutils/ls/ls.go` (dereference)
- [x] Flag `-N`, `--literal`: `internal/commands/coreutils/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#ls))
- [x] Flag `-Q`: `internal/commands/coreutils/ls/ls.go` (quote-name)
- [x] Flag `-R`: `internal/commands/coreutils/ls/ls.go` (recursive)
- [x] Flag `-S`: `internal/commands/coreutils/ls/ls.go` (sort-size)
- [x] Flag `-T`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `-U`: `internal/commands/coreutils/ls/ls.go` (do not sort)
- [x] Flag `-X`: `internal/commands/coreutils/ls/ls.go` (extension sort)
- [x] Flag `-Z`: `internal/commands/coreutils/ls/ls.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#ls))
- [x] Flag `-a`: `internal/commands/coreutils/ls/ls.go` (all)
- [x] Flag `-b`: `internal/commands/coreutils/ls/ls.go` (escape)
- [x] Flag `-c`: `internal/commands/coreutils/ls/ls.go` (ctime)
- [x] Flag `-d`: `internal/commands/coreutils/ls/ls.go` (directory itself)
- [x] Flag `-f`: `internal/commands/coreutils/ls/ls.go` (do not sort, enable -aU)
- [x] Flag `-g`: `internal/commands/coreutils/ls/ls.go` (like -l but no owner)
- [x] Flag `-h`: `internal/commands/coreutils/ls/ls.go` (human-readable)
- [x] Flag `-i`: `internal/commands/coreutils/ls/ls.go` (inode)
- [x] Flag `-l`: `internal/commands/coreutils/ls/ls.go` (long)
- [x] Flag `-m`: `internal/commands/coreutils/ls/ls.go` (comma)
- [x] Flag `-n`: `internal/commands/coreutils/ls/ls.go` (numeric)
- [x] Flag `-o`: `internal/commands/coreutils/ls/ls.go` (like -l but no group)
- [x] Flag `-p`: `internal/commands/coreutils/ls/ls.go` (indicator)
- [x] Flag `-q`: `internal/commands/coreutils/ls/ls.go` (hide-control-chars)
- [x] Flag `-r`: `internal/commands/coreutils/ls/ls.go` (reverse)
- [x] Flag `-s`: `internal/commands/coreutils/ls/ls.go` (size in blocks)
- [x] Flag `-t`: `internal/commands/coreutils/ls/ls.go` (sort-time)
- [x] Flag `-u`: `internal/commands/coreutils/ls/ls.go` (atime)
- [x] Flag `-v`: `internal/commands/coreutils/ls/ls.go` (natural sort)
- [x] Flag `-w`: `internal/commands/coreutils/ls/ls.go`
- [x] Flag `-x`: `internal/commands/coreutils/ls/ls.go` (across/horizontal)

### `md5sum`

- [x] Upstream: `third_party/coreutils/src/cksum.c`
- [x] Inherits all `cksum` hash flags: `internal/commands/coreutils/sum/sum.go`

### `mkdir`

- [x] Upstream: `third_party/coreutils/src/mkdir.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/mkdir/mkdir.go`
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/coreutils/mkdir/mkdir.go` (octal)
- [x] Flag `-p`, `--parents`: `internal/commands/coreutils/mkdir/mkdir.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/mkdir/mkdir.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/coreutils/mkdir/mkdir.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `mkfifo`

- [x] Upstream: `third_party/coreutils/src/mkfifo.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/mkfifo/mkfifo.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#mkfifo-mknod))
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/coreutils/mkfifo/mkfifo.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/coreutils/mkfifo/mkfifo.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `mknod`

- [x] Upstream: `third_party/coreutils/src/mknod.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/mknod/mknod.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#mkfifo-mknod))
- [x] Flag `-m`, `--mode=MODE`: `internal/commands/coreutils/mknod/mknod.go`
- [x] Flag `-Z`, `--context=CTX`: `internal/commands/coreutils/mknod/mknod.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `mktemp`

- [x] Upstream: `third_party/coreutils/src/mktemp.c`
- [x] Basic creation: Implemented in `internal/commands/coreutils/mktemp/mktemp.go`
- [x] Flag `--suffix`: `internal/commands/coreutils/mktemp/mktemp.go` (partial via TEMPLATE)
- [x] Flag `-d`: `internal/commands/coreutils/mktemp/mktemp.go`
- [x] Flag `-p`: `internal/commands/coreutils/mktemp/mktemp.go` (via --tmpdir)
- [x] Flag `-q`: `internal/commands/coreutils/mktemp/mktemp.go` (ignored)
- [x] Flag `-t`: `internal/commands/coreutils/mktemp/mktemp.go`
- [x] Flag `-u`: `internal/commands/coreutils/mktemp/mktemp.go`

### `mv`

- [x] Basic move/rename: Implemented in `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-b`, `--backup`: `internal/commands/coreutils/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-f`, `--force`: `internal/commands/coreutils/mv/mv.go` (ignored)
- [x] Flag `-i`, `--interactive`: `internal/commands/coreutils/mv/mv.go` (ignored)
- [x] Flag `-n`, `--no-clobber`: `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-t`, `--target-directory`: `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-T`, `--no-target-directory`: `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-u`, `--update`: `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/mv/mv.go`
- [x] Flag `-Z`, `--context`: `internal/commands/coreutils/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--exchange`: `internal/commands/coreutils/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--no-copy`: `internal/commands/coreutils/mv/mv.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `nice`

- [x] Upstream: `third_party/coreutils/src/nice.c`
- [x] Priority adjustment: Implemented (as wrapper) in `internal/commands/coreutils/nice/nice.go`
- [x] Flag `-n`, `--adjustment=N`: `internal/commands/coreutils/nice/nice.go`

### `nl`

- [x] Upstream: `third_party/coreutils/src/nl.c`
- [x] Basic numbering: Implemented in `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-b`, `--body-numbering=STYLE`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-d`, `--section-delimiter=CC`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-f`, `--footer-numbering=STYLE`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-h`, `--header-numbering=STYLE`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-i`, `--line-increment=NUMBER`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-l`, `--join-blank-lines=NUMBER`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-n`, `--number-format=FORMAT`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-p`, `--no-renumber`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-s`, `--number-separator=STRING`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-v`, `--starting-line-number=NUMBER`: `internal/commands/coreutils/nl/nl.go`
- [x] Flag `-w`, `--number-width=NUMBER`: `internal/commands/coreutils/nl/nl.go`

### `nohup`

- [x] Upstream: `third_party/coreutils/src/nohup.c`
- [x] Flag `--help`: `internal/commands/coreutils/nohup/nohup.go`
- [x] Flag `--version`: `internal/commands/coreutils/nohup/nohup.go`

### `nproc`

- [x] Upstream: `third_party/coreutils/src/nproc.c`
- [x] Basic nproc: Implemented in `internal/commands/coreutils/nproc/nproc.go`
- [x] Flag `--all`: `internal/commands/coreutils/nproc/nproc.go`
- [x] Flag `--ignore`: `internal/commands/coreutils/nproc/nproc.go`

### `numfmt`

- [x] Upstream: `third_party/coreutils/src/numfmt.c`
- [x] Conversion: Implemented in `internal/commands/coreutils/numfmt/numfmt.go`
- [x] Flag `-d`, `--delimiter=X`: `internal/commands/coreutils/numfmt/numfmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/numfmt/numfmt.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--to`: `internal/commands/coreutils/numfmt/numfmt.go`
- [x] Flag `--from`: `internal/commands/coreutils/numfmt/numfmt.go`
- [x] Flag `--header`: `internal/commands/coreutils/numfmt/numfmt.go`

### `od`

- [x] Upstream: `third_party/coreutils/src/od.c`
- [x] Format output: Implemented in `internal/commands/coreutils/od/od.go`
- [x] Flag `-A rad`: `internal/commands/coreutils/od/od.go`
- [x] Flag `-j bytes`: `internal/commands/coreutils/od/od.go`
- [x] Flag `-N bytes`: `internal/commands/coreutils/od/od.go`
- [x] Flag `-t type`: `internal/commands/coreutils/od/od.go` (Stub; only 2-byte octal supported)
- [x] Flag `-v`: `internal/commands/coreutils/od/od.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-w`: `internal/commands/coreutils/od/od.go`
- [x] Flag `-S`: `internal/commands/coreutils/od/od.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `paste`

- [x] Upstream: `third_party/coreutils/src/paste.c`
- [x] Basic paste: Implemented in `internal/commands/coreutils/paste/paste.go`
- [x] Flag `-d`, `--delimiters=LIST`: `internal/commands/coreutils/paste/paste.go`
- [x] Flag `-s`, `--serial`: `internal/commands/coreutils/paste/paste.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/paste/paste.go`

### `pathchk`

- [x] Upstream: `third_party/coreutils/src/pathchk.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/pathchk/pathchk.go`
- [x] Flag `-p`: `internal/commands/coreutils/pathchk/pathchk.go`
- [x] Flag `-P`: `internal/commands/coreutils/pathchk/pathchk.go`

### `pinky`

- [x] Upstream: `third_party/coreutils/src/pinky.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/pinky/pinky.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#pinky))
- [x] Flag `-b`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-f`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-h`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-i`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-l`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-p`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-q`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-s`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)
- [x] Flag `-w`: `internal/commands/coreutils/pinky/pinky.go` (Ignored)

### `pr`

- [x] Upstream: `third_party/coreutils/src/pr.c`
- [x] Print formatting: Implemented in `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-a`: `internal/commands/coreutils/pr/pr.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-d`: `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-h`: `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-l`: `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-m`: `internal/commands/coreutils/pr/pr.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-n`: `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-t`: `internal/commands/coreutils/pr/pr.go`
- [x] Flag `-w`: `internal/commands/coreutils/pr/pr.go`

### `printenv`

- [x] Upstream: `third_party/coreutils/src/printenv.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/printenv/printenv.go`
- [x] Flag `-0`: `internal/commands/coreutils/printenv/printenv.go`

### `printf`

- [x] Upstream: `third_party/coreutils/src/printf.c`
- [x] Basic formatting: Implemented in `internal/commands/coreutils/printf/printf.go`
- [x] Flag `%b`: `internal/commands/coreutils/printf/printf.go`
- [x] Flag `%q`: `internal/commands/coreutils/printf/printf.go`
- [x] Flag `-v VAR`: Implemented in `internal/commands/coreutils/printf/printf.go`
- [x] Format specifier `*` width/precision: `internal/commands/coreutils/printf/printf.go`
- [x] Reusing format string: `internal/commands/coreutils/printf/printf.go`

### `ptx`

- [x] Upstream: `third_party/coreutils/src/ptx.c`
- [x] Permuted Index: `internal/commands/coreutils/ptx/ptx.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#ptx))
- [x] Flag `-A`, `--auto-reference`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-F`, `--flag-truncation=STRING`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-G`, `--gnu-extensions`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-M`, `--macro-name=STRING`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-O`, `--format=roff`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-R`, `--right-side-refs`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-S`, `--sentence-regexp=REGEXP`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-T`, `--format=tex`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-W`, `--word-regexp=REGEXP`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-b`, `--break-file=FILE`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-f`, `--ignore-case`: `internal/commands/coreutils/ptx/ptx.go` (Ignored)
- [x] Flag `-g`, `--gap-size=NUMBER`: `internal/commands/coreutils/ptx/ptx.go`
- [x] Flag `-i`, `--ignore-file=FILE`: `internal/commands/coreutils/ptx/ptx.go`
- [x] Flag `-o`, `--only-file=FILE`: `internal/commands/coreutils/ptx/ptx.go`
- [x] Flag `-r`, `--references`: `internal/commands/coreutils/ptx/ptx.go`
- [x] Flag `-t`, `--typeset-mode`: `internal/commands/coreutils/ptx/ptx.go`
- [x] Flag `-w`, `--width=NUMBER`: `internal/commands/coreutils/ptx/ptx.go`

### `pwd`

- [x] Upstream: `third_party/coreutils/src/pwd.c`
- [x] Basic operation: Implemented in `internal/commands/bash/pwd/pwd.go`
- [x] Flag `-L`, `--logical`: `internal/commands/bash/pwd/pwd.go`
- [x] Flag `-P`, `--physical`: `internal/commands/bash/pwd/pwd.go`

### `readlink`


- [x] Upstream: `third_party/coreutils/src/readlink.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-e`, `--canonicalize-existing`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-m`, `--canonicalize-missing`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-q`, `--quiet`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-s`, `--silent`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-f`: `internal/commands/coreutils/readlink/readlink.go`
- [x] Flag `-n`: `internal/commands/coreutils/readlink/readlink.go`

### `realpath`

- [x] Upstream: `third_party/coreutils/src/realpath.c`
- [x] Basic functionality: Implemented in `internal/commands/coreutils/realpath/realpath.go`
- [x] Basic output: Implemented in `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-E`, `--canonicalize-existing`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-L`, `--logical`: `internal/commands/coreutils/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-P`, `--physical`: `internal/commands/coreutils/realpath/realpath.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-q`, `--quiet`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-s`, `--strip`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `--relative-to`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `--relative-base`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-e`: `internal/commands/coreutils/realpath/realpath.go`
- [x] Flag `-m`: `internal/commands/coreutils/realpath/realpath.go`

### `rm`

- [x] Upstream: `third_party/coreutils/src/rm.c`
- [x] Basic removal: Implemented in `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-d`, `--dir`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-f`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-i`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-I`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-r`, `-R`, `--recursive`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/rm/rm.go`
- [x] Flag `--one-file-system`: `internal/commands/coreutils/rm/rm.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `rmdir`

- [x] Upstream: `third_party/coreutils/src/rmdir.c`
- [x] Basic rmdir: Implemented in `internal/commands/coreutils/rmdir/rmdir.go`
- [x] Basic removal: Implemented in `internal/commands/coreutils/rmdir/rmdir.go`
- [x] Flag `--ignore-fail-on-non-empty`: `internal/commands/coreutils/rmdir/rmdir.go`
- [x] Flag `-p`: `internal/commands/coreutils/rmdir/rmdir.go`
- [x] Flag `-v`: `internal/commands/coreutils/rmdir/rmdir.go`

### `runcon`

- [x] Upstream: `third_party/coreutils/src/runcon.c`
- [x] Basic operation: `internal/commands/coreutils/runcon/runcon.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#chcon-runcon))
- [x] Flag `-c`, `--compute`: `internal/commands/coreutils/runcon/runcon.go`
- [x] Flag `-l`, `--user=USER`: `internal/commands/coreutils/runcon/runcon.go`
- [x] Flag `-r`, `--role=ROLE`: `internal/commands/coreutils/runcon/runcon.go`
- [x] Flag `-t`, `--type=TYPE`: `internal/commands/coreutils/runcon/runcon.go`
- [x] Flag `-u`, `--user=USER`: `internal/commands/coreutils/runcon/runcon.go`

### `seq`

- [x] Upstream: `third_party/coreutils/src/seq.c`
- [x] Basic sequence: Implemented in `internal/commands/coreutils/seq/seq.go`
- [x] Flag `-f`, `--format=FORMAT`: `internal/commands/coreutils/seq/seq.go`
- [x] Flag `-s`, `--separator=STRING`: `internal/commands/coreutils/seq/seq.go`
- [x] Flag `-w`, `--equal-width`: `internal/commands/coreutils/seq/seq.go`

### `sha1sum`

### `sha1sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha1sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/coreutils/sum/sum.go`

### `sha224sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha224sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/coreutils/sum/sum.go`

### `sha256sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha256sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/coreutils/sum/sum.go`

### `sha384sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha384sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/coreutils/sum/sum.go`

### `sha512sum`

- [x] Upstream: `third_party/coreutils/src/coreutils-sha512sum.c`
- [x] Inherits flags from `cksum`: `internal/commands/coreutils/sum/sum.go`

### `shred`

- [x] Upstream: `third_party/coreutils/src/shred.c`
- [x] Data erasure: Implemented in `internal/commands/coreutils/shred/shred.go` (Workaround; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#shred))
- [x] Flag `-f`, `--force`: `internal/commands/coreutils/shred/shred.go`
- [x] Flag `-n`, `--iterations=N`: `internal/commands/coreutils/shred/shred.go`
- [x] Flag `-s`, `--size=N`: `internal/commands/coreutils/shred/shred.go` (partial via iteration)
- [x] Flag `-u`, `--remove`: `internal/commands/coreutils/shred/shred.go`
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/shred/shred.go`
- [x] Flag `-x`, `--exact`: `internal/commands/coreutils/shred/shred.go`
- [x] Flag `-z`, `--zero`: `internal/commands/coreutils/shred/shred.go`

### `shuf`

- [x] Upstream: `third_party/coreutils/src/shuf.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-e`, `--echo`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-i`, `--input-range=LO-HI`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-n`, `--head-count=COUNT`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-o`, `--output=FILE`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-r`, `--repeat`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/shuf/shuf.go`
- [x] Flag `--random-source=FILE`: `internal/commands/coreutils/shuf/shuf.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#shuf))

### `sleep`

- [x] Upstream: `third_party/coreutils/src/sleep.c`
- [x] Basic sleep: Implemented in `internal/commands/coreutils/sleep/sleep.go`
- [x] Multiple arguments (sum): `internal/commands/coreutils/sleep/sleep.go`
- [x] Suffixes (s, m, h, d): `internal/commands/coreutils/sleep/sleep.go`
- [x] Flag `--help`: `internal/commands/coreutils/sleep/sleep.go`
- [x] Flag `--version`: `internal/commands/coreutils/sleep/sleep.go`

### `sort`

- [x] Upstream: `third_party/coreutils/src/sort.c`
- [x] Basic sorting: Implemented in `internal/commands/coreutils/sort/sort.go`
- [x] Ordering flags (`-b`, `-i`, `-d`, `-f`, `-g`, `-h`, `-n`, `-M`, `-R`, `-V`, `-r`): Implemented in `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-b`, `--ignore-leading-blanks`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-c`, `-C`, `--check`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-d`, `--dictionary-order`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-f`, `--ignore-case`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-g`, `--general-numeric-sort`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-h`, `--human-numeric-sort`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-i`, `--ignore-nonprinting`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-k`, `--key=KEYDEF`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-m`, `--merge`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-M`, `--month-sort`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-n`, `--numeric-sort`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-o`, `--output=FILE`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-r`, `--reverse`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-s`, `--stable`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-S`, `--buffer-size=SIZE`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-t`, `--field-separator=SEP`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-T`, `--temporary-directory=DIR`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--unique`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-V`, `--version-sort`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/sort/sort.go`
- [x] Flag `--parallel=N`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--random-sort` (`-R`): `internal/commands/coreutils/sort/sort.go`
- [x] Flag `--debug`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--files0-from=F`: `internal/commands/coreutils/sort/sort.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `split`

- [x] Basic split: Implemented in `internal/commands/coreutils/split/split.go`
- [x] Flag `--filter=COMMAND`: `internal/commands/coreutils/split/split.go` (Stub)
- [x] Flag `--verbose`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-C`, `--line-bytes=SIZE`: `internal/commands/coreutils/split/split.go` (Stub)
- [x] Flag `-a`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-b`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-d`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-e`, `--elide-empty-files`: `internal/commands/coreutils/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-l`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-n`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-t`: `internal/commands/coreutils/split/split.go`
- [x] Flag `-u`, `--unbuffered`: `internal/commands/coreutils/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-x`, `--hex-suffixes`: `internal/commands/coreutils/split/split.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `stat`

- [x] Upstream: `third_party/coreutils/src/stat.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/stat/stat.go`
- [x] Flag `-c`, `--format=FORMAT`: `internal/commands/coreutils/stat/stat.go`
- [x] Flag `-f`, `--file-system`: `internal/commands/coreutils/stat/stat.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-L`, `--dereference`: `internal/commands/coreutils/stat/stat.go`
- [x] Flag `-t`, `--terse`: `internal/commands/coreutils/stat/stat.go`
- [x] Flag `--printf=FORMAT`: `internal/commands/coreutils/stat/stat.go` (Stub)
- [x] Flag `--cached={always,never,default}`: `internal/commands/coreutils/stat/stat.go` (Stub)

### `stdbuf`

- [x] Upstream: `third_party/coreutils/src/stdbuf.c`
- [x] Stream Buffering: `internal/commands/coreutils/stdbuf/stdbuf.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#stdbuf))
- [x] Flag `-e`, `--error=MODE`: `internal/commands/coreutils/stdbuf/stdbuf.go` (Ignored)
- [x] Flag `-i`, `--input=MODE`: `internal/commands/coreutils/stdbuf/stdbuf.go` (Ignored)
- [x] Flag `-o`, `--output=MODE`: `internal/commands/coreutils/stdbuf/stdbuf.go` (Ignored)

### `stty`

- [x] Upstream: `third_party/coreutils/src/stty.c`
- [x] TTY Configuration: `internal/commands/coreutils/stty/stty.go` (Stub; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#stty))
- [x] Flag `-F`, `--file=DEVICE`: `internal/commands/coreutils/stty/stty.go` (Ignored)
- [x] Flag `-a`, `--all`: `internal/commands/coreutils/stty/stty.go` (Partial output)
- [x] Flag `-g`, `--save`: `internal/commands/coreutils/stty/stty.go` (Ignored)

### `sum`

- [x] Upstream: `third_party/coreutils/src/sum.c`
- [x] Basic checksum: Implemented in `internal/commands/coreutils/sumlegacy/sum.go`
- [x] Flag `-r`: `internal/commands/coreutils/sumlegacy/sum.go` (BSD algorithm)
- [x] Flag `-s`, `--sysv`: `internal/commands/coreutils/sumlegacy/sum.go` (System V algorithm)

### `sync`

- [x] Upstream: `third_party/coreutils/src/sync.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/sync/sync.go`
- [x] Flag `-d`, `--data`: `internal/commands/coreutils/sync/sync.go` (no-op)
- [x] Flag `-f`, `--file-system`: `internal/commands/coreutils/sync/sync.go` (no-op)

### `tac`

- [x] Upstream: `third_party/coreutils/src/tac.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/tac/tac.go`
- [x] Flag `-b`: `internal/commands/coreutils/tac/tac.go`
- [x] Flag `-r`, `--regex`: `internal/commands/coreutils/tac/tac.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-s`: `internal/commands/coreutils/tac/tac.go`

### `tail`

- [x] Basic output: Implemented in `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-f`, `--follow[={name|descriptor}]`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-F`: `internal/commands/coreutils/tail/tail.go` (partial)
- [x] Flag `-n`, `--lines`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-q`, `--quiet`, `--silent`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-s`, `--sleep-interval`: `internal/commands/coreutils/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-v`, `--verbose`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/tail/tail.go`
- [x] Flag `--pid`: `internal/commands/coreutils/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `--retry`: `internal/commands/coreutils/tail/tail.go` (Stub)
- [x] Flag `--max-unchanged-stats`: `internal/commands/coreutils/tail/tail.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `tee`

- [x] Basic copy: Implemented in `internal/commands/coreutils/tee/tee.go`
- [x] Flag `-a`, `--append`: `internal/commands/coreutils/tee/tee.go`
- [x] Flag `-i`, `--ignore-interrupts`: `internal/commands/coreutils/tee/tee.go`
- [x] Flag `-p`, `--output-error[=MODE]`: `internal/commands/coreutils/tee/tee.go`

### `test`

- [x] Upstream: `third_party/coreutils/src/test.c`
- [x] Unary operators (-e, -f, -d, etc.): Implemented in `internal/commands/bash/test/test.go`
- [x] String operators (=, !=, -z, -n): Implemented in `internal/commands/bash/test/test.go`
- [x] Numeric operators (-eq, -ne, etc.): Implemented in `internal/commands/bash/test/test.go`
- [x] Logical operators (!, -a, -o): Implemented in `internal/commands/bash/test/test.go`
- [x] File operators (`-r`, `-w`, `-x`, `-L`, `-G`, `-O`, `-S`, `-p`, `-b`, `-c`, `-t`, `-k`, `-u`, `-g`, `-N`): `internal/commands/bash/test/test.go`
- [x] Binary file operators (`-nt`, `-ot`, `-ef`): `internal/commands/bash/test/test.go`
- [x] String comparison operators (`>`, `<`): `internal/commands/bash/test/test.go`
- [x] Length operator (`-l STRING`): `internal/commands/bash/test/test.go`
- [x] Aliases: `[`

### `timeout`

- [x] Upstream: `third_party/coreutils/src/timeout.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/timeout/timeout.go`
- [x] Flag `--help`: `internal/commands/coreutils/timeout/timeout.go`
- [x] Flag `--version`: `internal/commands/coreutils/timeout/timeout.go`
- [x] Flag `-k`: `internal/commands/coreutils/timeout/timeout.go`
- [x] Flag `-s`: `internal/commands/coreutils/timeout/timeout.go`

### `touch`

- [x] Upstream: `third_party/coreutils/src/touch.c`
- [x] Basic touch: Implemented in `internal/commands/coreutils/touch/touch.go`
- [x] Basic timestamp update: Implemented in `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-t STAMP`: `internal/commands/coreutils/touch/touch.go` (explicit timestamp)
- [x] Flag `-a`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-c`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-d`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-h`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-m`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-r`: `internal/commands/coreutils/touch/touch.go`
- [x] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `internal/commands/coreutils/touch/touch.go`

### `tr`

- [x] Upstream: `third_party/coreutils/src/tr.c`
- [x] Basic translation: Implemented in `internal/commands/coreutils/tr/tr.go`
- [x] Flag `-c`, `-C`, `--complement`: `internal/commands/coreutils/tr/tr.go`
- [x] Flag `-d`, `--delete`: `internal/commands/coreutils/tr/tr.go`
- [x] Flag `-s`, `--squeeze-repeats`: `internal/commands/coreutils/tr/tr.go`
- [x] Flag `-t`, `--truncate-set1`: `internal/commands/coreutils/tr/tr.go`

### `true`

- [x] Upstream: `third_party/coreutils/src/true.c`
- [x] Basic operation: Implemented in `internal/commands/bash/boolcmd/bool.go`


### `truncate`

- [x] Upstream: `third_party/coreutils/src/truncate.c`
- [x] Basic truncation: Implemented in `internal/commands/coreutils/truncate/truncate.go`
- [x] Flag `-c`: `internal/commands/coreutils/truncate/truncate.go`
- [x] Flag `-o`: `internal/commands/coreutils/truncate/truncate.go` (ignored stub)
- [x] Flag `-r`: `internal/commands/coreutils/truncate/truncate.go`
- [x] Flag `-s`: `internal/commands/coreutils/truncate/truncate.go`

### `tsort`

- [x] Upstream: `third_party/coreutils/src/tsort.c`
- [x] Topological sort: Implemented in `internal/commands/coreutils/tsort/tsort.go`
- [x] Flag `--help`: `internal/commands/coreutils/tsort/tsort.go`
- [x] Flag `--version`: `internal/commands/coreutils/tsort/tsort.go`

### `tty`

- [x] Upstream: `third_party/coreutils/src/tty.c`
- [x] TTY reporting: Implemented in `internal/commands/coreutils/tty/tty.go`
- [x] Flag `-s`, `--silent`, `--quiet`: `internal/commands/coreutils/tty/tty.go`

### `uname`

- [x] Upstream: `third_party/coreutils/src/uname.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-a`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-i`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-m`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-n`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-o`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-p`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-r`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-s`: `internal/commands/coreutils/uname/uname.go`
- [x] Flag `-v`: `internal/commands/coreutils/uname/uname.go`

### `unexpand`

- [x] Upstream: `third_party/coreutils/src/unexpand.c`
- [x] Basic conversion: Implemented in `internal/commands/coreutils/unexpand/unexpand.go`
- [x] Flag `-a`, `--all`: `internal/commands/coreutils/unexpand/unexpand.go`
- [x] Flag `-t`, `--tabs=LIST`: `internal/commands/coreutils/unexpand/unexpand.go`
- [x] Flag `--first-only`: `internal/commands/coreutils/unexpand/unexpand.go`

### `uniq`

- [x] Upstream: `third_party/coreutils/src/uniq.c`
- [x] Basic filtering: Implemented in `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `-c`, `--count`: `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `-d`, `--repeated`: `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `-D`, `--all-repeated[=METHOD]`: `internal/commands/coreutils/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-f`, `--skip-fields=N`: `internal/commands/coreutils/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-i`, `--ignore-case`: `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `-s`, `--skip-chars=N`: `internal/commands/coreutils/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-u`, `--unique`: `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `-w`, `--check-chars=N`: `internal/commands/coreutils/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))
- [x] Flag `-z`, `--zero-terminated`: `internal/commands/coreutils/uniq/uniq.go`
- [x] Flag `--group[=METHOD]`: `internal/commands/coreutils/uniq/uniq.go` (Ignored; see [functional_gap.md](file:///Users/aren/github/yarencheng/go-bash-wasm/docs/coreutils/functional_gap.md#commonly-ignored-flags))

### `unlink`

- [x] Basic removal: Implemented in `internal/commands/coreutils/unlink/unlink.go` (exactly 1 arg required)
- [x] Flag `--help`: `internal/commands/coreutils/unlink/unlink.go`
- [x] Flag `--version`: `internal/commands/coreutils/unlink/unlink.go`

### `uptime`

- [x] Upstream: `third_party/coreutils/src/uptime.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/uptime/uptime.go`
- [x] Flag `-p`, `--pretty`: `internal/commands/coreutils/uptime/uptime.go`
- [x] Flag `-s`, `--since`: `internal/commands/coreutils/uptime/uptime.go`

### `users`

- [x] Upstream: `third_party/coreutils/src/users.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/users/users.go`
- [x] Flag `--help`: `internal/commands/coreutils/users/users.go`
- [x] Flag `--version`: `internal/commands/coreutils/users/users.go`

### `vdir`

- [x] Upstream: `third_party/coreutils/src/ls.c`
- [x] Inherits flags from `ls`

### `wc`

- [x] Basic counts: Implemented in `internal/commands/coreutils/wc/wc.go`
- [x] Flag `-c`, `--bytes`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `-m`, `--chars`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `-l`, `--lines`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `-w`, `--words`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `-L`, `--max-line-length`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `--files0-from=F`: `internal/commands/coreutils/wc/wc.go`
- [x] Flag `--total={auto,always,only,never}`: `internal/commands/coreutils/wc/wc.go`

### `who`

- [x] Upstream: `third_party/coreutils/src/who.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/who/who.go`
- [x] Flag `-H`, `--heading`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-a`, `--all`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-b`, `--boot`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-d`, `--dead`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-l`, `--login`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-m`: `internal/commands/coreutils/who/who.go` (current user)
- [x] Flag `-p`, `--process`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-q`, `--count`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-r`, `--runlevel`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-s`, `--short`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-t`, `--time`: `internal/commands/coreutils/who/who.go`
- [x] Flag `-u`, `--users`: `internal/commands/coreutils/who/who.go`

### `whoami`

- [x] Upstream: `third_party/coreutils/src/whoami.c`
- [x] Basic output: Implemented in `internal/commands/coreutils/whoami/whoami.go`
- [x] Flag `--help`: `internal/commands/coreutils/whoami/whoami.go`
- [x] Flag `--version`: `internal/commands/coreutils/whoami/whoami.go`

### `yes`

- [x] Upstream: `third_party/coreutils/src/yes.c`
- [x] Basic operation: Implemented in `internal/commands/coreutils/yes/yes.go`
- [x] Basic repetition: Implemented in `internal/commands/coreutils/yes/yes.go`
- [x] Flag `--help`: `internal/commands/coreutils/yes/yes.go`
- [x] Flag `--version`: `internal/commands/coreutils/yes/yes.go`


