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

- [ ] Upstream: `third_party/bash/builtins/colon.def`
- [ ] Basic operation: Missing implementation

### `alias`

- [ ] Upstream: `third_party/bash/builtins/alias.def`
- [ ] Basic management: Missing implementation
- [ ] Flag `-p`: `third_party/bash/builtins/alias.def:L79` (print)

### `arch`

- [ ] Upstream: `third_party/coreutils/src/coreutils-arch.c`
- [ ] Inherits flags from `uname`

### `base32`

- [ ] Basic encoding/decoding: Missing implementation
- [ ] Flag `--base16`: `third_party/coreutils/src/basenc.c:L86`
- [ ] Flag `--base2lsbf`: `third_party/coreutils/src/basenc.c:L88`
- [ ] Flag `--base2msbf`: `third_party/coreutils/src/basenc.c:L87`
- [ ] Flag `--base32`: `third_party/coreutils/src/basenc.c:L84`
- [ ] Flag `--base32hex`: `third_party/coreutils/src/basenc.c:L85`
- [ ] Flag `--base58`: `third_party/coreutils/src/basenc.c:L83`
- [ ] Flag `--base64`: `third_party/coreutils/src/basenc.c:L81`
- [ ] Flag `--base64url`: `third_party/coreutils/src/basenc.c:L82`
- [ ] Flag `--z85`: `third_party/coreutils/src/basenc.c:L89`
- [ ] Flag `-d`: `third_party/coreutils/src/basenc.c:L77`
- [ ] Flag `-i`: `third_party/coreutils/src/basenc.c:L79`
- [ ] Flag `-w`: `third_party/coreutils/src/basenc.c:L78`

### `basename`

- [ ] Basic operation: Missing implementation
- [ ] Flag `-a`, `--multiple`: `third_party/coreutils/src/basename.c:L155`
- [ ] Flag `-s`, `--suffix=SUFFIX`: `third_party/coreutils/src/basename.c:L150`
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/basename.c:L159`

### `basenc`

- [ ] Upstream: `third_party/coreutils/src/basenc.c`
- [ ] Basic encoding/decoding: Missing implementation
- [ ] Flag `-d`, `--decode`: `third_party/coreutils/src/basenc.c:L125`
- [ ] Flag `-i`, `--ignore-garbage`: `third_party/coreutils/src/basenc.c:L129`
- [ ] Flag `-w`, `--wrap=COLS`: `third_party/coreutils/src/basenc.c:L133`
- [ ] Flag `--base16`: `third_party/coreutils/src/basenc.c:L142`
- [ ] Flag `--base32`: `third_party/coreutils/src/basenc.c:L141`
- [ ] Flag `--base32hex`: `third_party/coreutils/src/basenc.c:L143`
- [ ] Flag `--base64`: `third_party/coreutils/src/basenc.c:L141`
- [ ] Flag `--base64url`: `third_party/coreutils/src/basenc.c:L141`
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

- [ ] Upstream: `third_party/bash/builtins/builtin.def`
- [ ] Basic execution: Missing implementation

### `caller`

- [ ] Upstream: `third_party/bash/builtins/caller.def`

### `cat`

- [ ] Basic output: Missing implementation
- [ ] Flag `-A`, `--show-all`: `third_party/coreutils/src/cat.c:L100`
- [ ] Flag `-b`, `--number-nonblank`: `third_party/coreutils/src/cat.c:L103`
- [ ] Flag `-e`: `third_party/coreutils/src/cat.c:L105` (implies -vE)
- [ ] Flag `-E`, `--show-ends`: `third_party/coreutils/src/cat.c:L108`
- [ ] Flag `-n`, `--number`: `third_party/coreutils/src/cat.c:L112`
- [ ] Flag `-s`, `--squeeze-blank`: `third_party/coreutils/src/cat.c:L115`
- [ ] Flag `-t`: `third_party/coreutils/src/cat.c:L118` (implies -vT)
- [ ] Flag `-T`, `--show-tabs`: `third_party/coreutils/src/cat.c:L121`
- [ ] Flag `-u`: `third_party/coreutils/src/cat.c:L124` (ignored)
- [ ] Flag `-v`, `--show-nonprinting`: `third_party/coreutils/src/cat.c:L127`

### `cd`

- [ ] Upstream: `third_party/bash/builtins/cd.def`
- [ ] Basic change directory: Missing implementation
- [ ] CDPATH support: `third_party/bash/builtins/cd.def:L84`
- [ ] Flag `-e`: `third_party/bash/builtins/cd.def:L98` (exit status if -P cannot be satisfied)
- [ ] Flag `-L`: `third_party/bash/builtins/cd.def:L94`
- [ ] Flag `-P`: `third_party/bash/builtins/cd.def:L96`

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

- [ ] Upstream: `third_party/coreutils/src/chown-chgrp.c`
- [ ] Inherits flags from `chown`: `--dereference`, `--no-dereference`, `--recursive`, `--from`, `--reference`, `-H`, `-L`, `-P`, `-c`, `-f`, `-v`
- [ ] Basic group change: Missing implementation

### `chmod`

- [ ] Basic mode change: Missing implementation
- [ ] Numeric mode support: `third_party/coreutils/src/chmod.c:L415`
- [ ] Symbolic mode support: `third_party/coreutils/src/chmod.c:L414`
- [ ] Flag `--dereference`: `third_party/coreutils/src/chmod.c:L437`
- [ ] Flag `--no-preserve-root`: `third_party/coreutils/src/chmod.c:L446`
- [ ] Flag `--preserve-root`: `third_party/coreutils/src/chmod.c:L450`
- [ ] Flag `--reference=RFILE`: `third_party/coreutils/src/chmod.c:L454`
- [ ] Flag `-R`: `third_party/coreutils/src/chmod.c:L459`
- [ ] Flag `-c`: `third_party/coreutils/src/chmod.c:L425`
- [ ] Flag `-f`: `third_party/coreutils/src/chmod.c:L429`
- [ ] Flag `-h`: `third_party/coreutils/src/chmod.c:L442`
- [ ] Flag `-v`: `third_party/coreutils/src/chmod.c:L433`

### `chown`

- [ ] Basic ownership change: Missing implementation
- [ ] Flag `--dereference`: `third_party/coreutils/src/chown.c:L111`
- [ ] Flag `--from`: `third_party/coreutils/src/chown.c:L121`
- [ ] Flag `--from=CURRENT_OWNER:CURRENT_GROUP`: `third_party/coreutils/src/chown.c:L121`
- [ ] Flag `--no-preserve-root`: `third_party/coreutils/src/chown.c:L128`
- [ ] Flag `--preserve-root`: `third_party/coreutils/src/chown.c:L131`
- [ ] Flag `--reference=RFILE`: `third_party/coreutils/src/chown.c:L134`
- [ ] Flag `-H`: `third_party/coreutils/src/chown.c:L143`
- [ ] Flag `-L`: `third_party/coreutils/src/chown.c:L147`
- [ ] Flag `-P`: `third_party/coreutils/src/chown.c:L151`
- [ ] Flag `-R`: `third_party/coreutils/src/chown.c:L139`
- [ ] Flag `-c`: `third_party/coreutils/src/chown.c:L99`
- [ ] Flag `-f`: `third_party/coreutils/src/chown.c:L103`
- [ ] Flag `-h`: `third_party/coreutils/src/chown.c:L116`
- [ ] Flag `-v`: `third_party/coreutils/src/chown.c:L107`

### `chroot`

- [ ] Upstream: `third_party/coreutils/src/chroot.c`

### `cksum`

- [ ] Upstream: `third_party/coreutils/src/cksum.c`
- [ ] Basic CRC-32: Missing implementation
- [ ] Flag `-a`, `--algorithm=ALGO`: `third_party/coreutils/src/cksum.c:L186`
- [ ] Flag `-c`, `--check`: `third_party/coreutils/src/cksum.c:L148`
- [ ] Flag `-l`, `--length=BITS`: `third_party/coreutils/src/cksum.c:L181`
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/cksum.c:L158`
- [ ] Flag `--base64`: `third_party/coreutils/src/cksum.c:L187`
- [ ] Flag `--raw`: `third_party/coreutils/src/cksum.c:L189`
- [ ] Flag `--tag`: `third_party/coreutils/src/cksum.c:L157`
- [ ] Flag `--untagged`: `third_party/coreutils/src/cksum.c:L190`
- [ ] Flag `--ignore-missing`: `third_party/coreutils/src/cksum.c:L149`
- [ ] Flag `--quiet`: `third_party/coreutils/src/cksum.c:L150`
- [ ] Flag `--status`: `third_party/coreutils/src/cksum.c:L151`
- [ ] Flag `--strict`: `third_party/coreutils/src/cksum.c:L154`
- [ ] Flag `-w`, `--warn`: `third_party/coreutils/src/cksum.c:L153`

### `comm`

- [ ] Upstream: `third_party/coreutils/src/comm.c`
- [ ] Basic comparison: Missing implementation
- [ ] Flag `--check-order`: `third_party/coreutils/src/comm.c:L473`
- [ ] Flag `--nocheck-order`: `third_party/coreutils/src/comm.c:L474`
- [ ] Flag `--output-delimiter`: `third_party/coreutils/src/comm.c:L477`
- [ ] Flag `--total`: `third_party/coreutils/src/comm.c:L478`
- [ ] Flag `-1`: `third_party/coreutils/src/comm.c:L467`
- [ ] Flag `-2`: `third_party/coreutils/src/comm.c:L468`
- [ ] Flag `-3`: `third_party/coreutils/src/comm.c:L469`
- [ ] Flag `-z`: `third_party/coreutils/src/comm.c:L480`

### `command`

- [ ] Upstream: `third_party/bash/builtins/command.def`
- [ ] Execution override: Missing implementation
- [ ] Flag `-p`: `third_party/bash/builtins/command.def:L75` (default PATH)
- [ ] Flag `-v`: `third_party/bash/builtins/command.def:L75` (identify command)
- [ ] Flag `-V`: `third_party/bash/builtins/command.def:L75` (verbose identify)

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

- [ ] Basic copy: Missing implementation
- [ ] Flag `-a`, `--archive`: `third_party/coreutils/src/cp.c:L173`
- [ ] Flag `-b`, `--backup`: `third_party/coreutils/src/cp.c:L181`
- [ ] Flag `-d`: `third_party/coreutils/src/cp.c:L185` (implies -P --preserve=links)
- [ ] Flag `-f`, `--force`: `third_party/coreutils/src/cp.c:L193`
- [ ] Flag `-H`: `third_party/coreutils/src/cp.c:L201`
- [ ] Flag `-i`, `--interactive`: `third_party/coreutils/src/cp.c:L205`
- [ ] Flag `-l`, `--link`: `third_party/coreutils/src/cp.c:L209`
- [ ] Flag `-L`, `--dereference`: `third_party/coreutils/src/cp.c:L213`
- [ ] Flag `-n`, `--no-clobber`: `third_party/coreutils/src/cp.c:L217`
- [ ] Flag `-p`: `third_party/coreutils/src/cp.c:L234` (same as --preserve=mode,ownership,timestamps)
- [ ] Flag `-P`, `--no-dereference`: `third_party/coreutils/src/cp.c:L230`
- [ ] Flag `-r`, `-R`, `--recursive`: `third_party/coreutils/src/cp.c:L250`
- [ ] Flag `-s`, `--symbolic-link`: `third_party/coreutils/src/cp.c:L258`
- [ ] Flag `-t`, `--target-directory`: `third_party/coreutils/src/cp.c:L262`
- [ ] Flag `-T`, `--no-target-directory`: `third_party/coreutils/src/cp.c:L266`
- [ ] Flag `-u`, `--update`: `third_party/coreutils/src/cp.c:L270`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/cp.c:L274`
- [ ] Flag `-x`, `--one-file-system`: `third_party/coreutils/src/cp.c:L278`
- [ ] Flag `-Z`, `--context`: `third_party/coreutils/src/cp.c:L282`
- [ ] Flag `--attributes-only`: `third_party/coreutils/src/cp.c:L177`
- [ ] Flag `--preserve[=ATTR_LIST]`: `third_party/coreutils/src/cp.c:L242`
- [ ] Flag `--no-preserve=ATTR_LIST`: `third_party/coreutils/src/cp.c:L226`
- [ ] Flag `--parents`: `third_party/coreutils/src/cp.c:L238`
- [ ] Flag `--reflink[=WHEN]`: `third_party/coreutils/src/cp.c:L246`
- [ ] Flag `--sparse=WHEN`: `third_party/coreutils/src/cp.c:L254`
- [ ] Flag `--strip-trailing-slashes`: `third_party/coreutils/src/cp.c:L262`

### `csplit`

- [ ] Upstream: `third_party/coreutils/src/csplit.c`
- [ ] Basic split: Missing implementation
- [ ] Flag `--suppress-matched`: `third_party/coreutils/src/csplit.c:1435`
- [ ] Flag `-b`: `third_party/coreutils/src/csplit.c:1423`
- [ ] Flag `-f`: `third_party/coreutils/src/csplit.c:1427`
- [ ] Flag `-k`: `third_party/coreutils/src/csplit.c:1431`
- [ ] Flag `-n`: `third_party/coreutils/src/csplit.c:1439`
- [ ] Flag `-s`: `third_party/coreutils/src/csplit.c:1443`
- [ ] Flag `-z`: `third_party/coreutils/src/csplit.c:1447`

### `cut`

- [ ] Upstream: `third_party/coreutils/src/cut.c`
- [ ] Basic selection: Missing implementation
- [ ] Flag `-b`, `--bytes=LIST`: `third_party/coreutils/src/cut.c:L143`
- [ ] Flag `-c`, `--characters=LIST`: `third_party/coreutils/src/cut.c:L147`
- [ ] Flag `-d`, `--delimiter=DELIM`: `third_party/coreutils/src/cut.c:L151`
- [ ] Flag `-f`, `--fields=LIST`: `third_party/coreutils/src/cut.c:L155`
- [ ] Flag `-n`: `third_party/coreutils/src/cut.c:L159` (ignored)
- [ ] Flag `-s`, `--only-delimited`: `third_party/coreutils/src/cut.c:L163`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/cut.c:L175`
- [ ] Flag `--complement`: `third_party/coreutils/src/cut.c:L167`
- [ ] Flag `--output-delimiter=STRING`: `third_party/coreutils/src/cut.c:L171`

### `date`

- [ ] Upstream: `third_party/coreutils/src/date.c`
- [ ] Basic output: Missing implementation
- [ ] Custom format `+FORMAT`: `third_party/coreutils/src/date.c:L607`
- [ ] Flag `-d`, `--date=STRING`: `third_party/coreutils/src/date.c:L501`
- [ ] Flag `-f`, `--file=DATEFILE`: `third_party/coreutils/src/date.c:L508`
- [ ] Flag `-I[FMT]`, `--iso-8601[=FMT]`: `third_party/coreutils/src/date.c:L513`
- [ ] Flag `-r`, `--reference=FILE`: `third_party/coreutils/src/date.c:L543`
- [ ] Flag `-R`, `--rfc-email`: `third_party/coreutils/src/date.c:L549`
- [ ] Flag `-s`, `--set=STRING`: `third_party/coreutils/src/date.c:L553`
- [ ] Flag `-u`, `--utc`, `--universal`: `third_party/coreutils/src/date.c:L561`
- [ ] Flag `--debug`: `third_party/coreutils/src/date.c:L497`

### `dd`

- [ ] Upstream: `third_party/coreutils/src/dd.c`
- [ ] Data copy: Missing implementation
- [ ] Operand `bs=BYTES`: `third_party/coreutils/src/dd.c:L536`
- [ ] Operand `cbs=BYTES`: `third_party/coreutils/src/dd.c:L539`
- [ ] Operand `conv=CONVS`: `third_party/coreutils/src/dd.c:L543`
- [ ] Operand `count=N`: `third_party/coreutils/src/dd.c:L546`
- [ ] Operand `ibs=BYTES`: `third_party/coreutils/src/dd.c:L549`
- [ ] Operand `if=FILE`: `third_party/coreutils/src/dd.c:L552`
- [ ] Operand `iflag=FLAGS`: `third_party/coreutils/src/dd.c:L555`
- [ ] Operand `obs=BYTES`: `third_party/coreutils/src/dd.c:L558`
- [ ] Operand `of=FILE`: `third_party/coreutils/src/dd.c:L561`
- [ ] Operand `oflag=FLAGS`: `third_party/coreutils/src/dd.c:L564`
- [ ] Operand `seek=N`: `third_party/coreutils/src/dd.c:L567`
- [ ] Operand `skip=N`: `third_party/coreutils/src/dd.c:L570`
- [ ] Operand `status=LEVEL`: `third_party/coreutils/src/dd.c:L573`

### `declare`

- [ ] Upstream: `third_party/bash/builtins/declare.def`
- [ ] Attribute management (-i, -r, -x, -a, -A): Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/declare.def:L282` (indexed array)
- [ ] Flag `-A`: `third_party/bash/builtins/declare.def:L294` (associative array)
- [ ] Flag `-i`: `third_party/bash/builtins/declare.def:L324` (integer)
- [ ] Flag `-r`: `third_party/bash/builtins/declare.def:L330` (readonly)
- [ ] Flag `-x`: `third_party/bash/builtins/declare.def:L336` (export)
- [ ] Flag `-l`: `third_party/bash/builtins/declare.def:L348` (lowercase)
- [ ] Flag `-u`: `third_party/bash/builtins/declare.def:L353` (uppercase)
- [ ] Flag `-n`: `third_party/bash/builtins/declare.def:L327` (nameref)
- [ ] Flag `-t`: `third_party/bash/builtins/declare.def:L333` (trace)
- [ ] Flag `-f`: `third_party/bash/builtins/declare.def:L313` (function)
- [ ] Flag `-F`: `third_party/bash/builtins/declare.def:L309` (function name only)
- [ ] Flag `-g`: `third_party/bash/builtins/declare.def:L320` (global)
- [ ] Flag `-p`: `third_party/bash/builtins/declare.def:L306` (print)
- [ ] Flag `-I`: `third_party/bash/builtins/declare.def:L359` (inherit attributes)
- [ ] Aliases: `typeset`

### `df`

- [ ] Upstream: `third_party/coreutils/src/df.c`
- [ ] Basic df: Missing implementation
- [ ] Basic output: Missing implementation
- [ ] Flag `--no-sync`: `third_party/coreutils/src/df.c:L266`
- [ ] Flag `--output[=FIELD_LIST]`: `third_party/coreutils/src/df.c:L262`
- [ ] Flag `--sync`: `third_party/coreutils/src/df.c:L265`
- [ ] Flag `--total`: `third_party/coreutils/src/df.c:L267`
- [ ] Flag `-B`: `third_party/coreutils/src/df.c:L257`
- [ ] Flag `-H`: `third_party/coreutils/src/df.c:L260`
- [ ] Flag `-P`: `third_party/coreutils/src/df.c:L263`
- [ ] Flag `-T`: `third_party/coreutils/src/df.c:L264`
- [ ] Flag `-a`: `third_party/coreutils/src/df.c:L256`
- [ ] Flag `-h`: `third_party/coreutils/src/df.c:L259`
- [ ] Flag `-i`: `third_party/coreutils/src/df.c:L258`
- [ ] Flag `-k`: `third_party/coreutils/src/df.c:L1307`
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

- [ ] Upstream: `third_party/coreutils/src/dirname.c`
- [ ] Basic operation: Missing implementation
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/dirname.c:L76`

### `dirs`

- [ ] Upstream: `third_party/bash/builtins/pushd.def`
- [ ] Flag `-c`: `third_party/bash/builtins/pushd.def:L377` (clear stack)
- [ ] Flag `-l`: `third_party/bash/builtins/pushd.def:L380` (long listing)
- [ ] Flag `-p`: `third_party/bash/builtins/pushd.def:L387` (print with one line per entry)
- [ ] Flag `-v`: `third_party/bash/builtins/pushd.def:L390` (verbose)

### `disown`

- [ ] Upstream: `third_party/bash/builtins/jobs.def`
- [ ] Flag `-a`: `third_party/bash/builtins/jobs.def:L196` (all jobs)
- [ ] Flag `-h`: `third_party/bash/builtins/jobs.def:L199` (mark to not receive SIGHUP)
- [ ] Flag `-r`: `third_party/bash/builtins/jobs.def:L202` (running jobs only)

### `du`

- [ ] Upstream: `third_party/coreutils/src/du.c`
- [ ] Basic du: Missing implementation
- [ ] Basic usage summary: Missing implementation
- [ ] Flag `-a`, `--all`: `third_party/coreutils/src/du.c:L294`
- [ ] Flag `-A`, `--apparent-size`: `third_party/coreutils/src/du.c:L298`
- [ ] Flag `-b`, `--bytes`: `third_party/coreutils/src/du.c:L310`
- [ ] Flag `-c`, `--total`: `third_party/coreutils/src/du.c:L314`
- [ ] Flag `-d`, `--max-depth=N`: `third_party/coreutils/src/du.c:L322`
- [ ] Flag `-D`, `--dereference-args`: `third_party/coreutils/src/du.c:L318`
- [ ] Flag `-h`, `--human-readable`: `third_party/coreutils/src/du.c:L337`
- [ ] Flag `-H`: `third_party/coreutils/src/du.c:L333` (same as --dereference-args)
- [ ] Flag `-k`: `third_party/coreutils/src/du.c:L345` (1K blocks)
- [ ] Flag `-l`, `--count-links`: `third_party/coreutils/src/du.c:L353`
- [ ] Flag `-L`, `--dereference`: `third_party/coreutils/src/du.c:L349`
- [ ] Flag `-m`: `third_party/coreutils/src/du.c:L357` (1M blocks)
- [ ] Flag `-P`, `--no-dereference`: `third_party/coreutils/src/du.c:L361`
- [ ] Flag `-s`, `--summarize`: `third_party/coreutils/src/du.c:L373`
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

- [ ] Basic output: Missing implementation
- [ ] Flag `-E`: `third_party/bash/builtins/echo.def:L152`
- [ ] Flag `-e`: `third_party/bash/builtins/echo.def:L149`
- [ ] Flag `-n`: `third_party/bash/builtins/echo.def:L146`

### `enable`

- [ ] Upstream: `third_party/bash/builtins/enable.def`
- [ ] Flag `-a`: `third_party/bash/builtins/enable.def:L157` (display all)
- [ ] Flag `-d`: `third_party/bash/builtins/enable.def:L160` (delete loaded)
- [ ] Flag `-n`: `third_party/bash/builtins/enable.def:L163` (disable)
- [ ] Flag `-p`: `third_party/bash/builtins/enable.def:L166` (print status)
- [ ] Flag `-s`: `third_party/bash/builtins/enable.def:L169` (POSIX special only)
- [ ] Flag `-f filename`: `third_party/bash/builtins/enable.def:L172` (load from dynamic file)

### `env`

- [ ] Upstream: `third_party/coreutils/src/env.c`
- [ ] Basic execution: Missing implementation
- [ ] Flag `-a`, `--argv0=ARG`: `third_party/coreutils/src/env.c:L123`
- [ ] Flag `-i`, `--ignore-environment`: `third_party/coreutils/src/env.c:L127`
- [ ] Flag `-u`, `--unset=NAME`: `third_party/coreutils/src/env.c:L135`
- [ ] Flag `-0`, `--null`: `third_party/coreutils/src/env.c:L131`
- [ ] Flag `-C`, `--chdir=DIR`: `third_party/coreutils/src/env.c:L139`
- [ ] Flag `-S`, `--split-string=S`: `third_party/coreutils/src/env.c:L143`
- [ ] Flag `-v`, `--debug`: `third_party/coreutils/src/env.c:L164`
- [ ] Flag `--block-signal[=SIG]`: `third_party/coreutils/src/env.c:L148`
- [ ] Flag `--default-signal[=SIG]`: `third_party/coreutils/src/env.c:L152`
- [ ] Flag `--ignore-signal[=SIG]`: `third_party/coreutils/src/env.c:L156`
- [ ] Flag `--list-signal-handling`: `third_party/coreutils/src/env.c:L160`

### `eval`

- [ ] Upstream: `third_party/bash/builtins/eval.def`
- [ ] Basic execution: Missing implementation

### `exec`

- [ ] Upstream: `third_party/bash/builtins/exec.def`
- [ ] Basic execution: Missing implementation
- [ ] Flag `-l`: `third_party/bash/builtins/exec.def:L117` (login shell)
- [ ] Flag `-a name`: `third_party/bash/builtins/exec.def:L120`
- [ ] Flag `-c`: `third_party/bash/builtins/exec.def:L114`

### `exit`

- [ ] Upstream: `third_party/bash/builtins/exit.def`
- [ ] Basic exit: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/exit.def:L25`

### `expand`

- [ ] Upstream: `third_party/coreutils/src/expand.c`
- [ ] Basic conversion: Missing implementation
- [ ] Flag `-i`, `--initial`: `third_party/coreutils/src/expand.c:L78`
- [ ] Flag `-t`, `--tabs=LIST`: `third_party/coreutils/src/expand.c:L82`

### `export`

- [ ] Upstream: `third_party/bash/builtins/setattr.def`
- [ ] Flag `-f`: `third_party/bash/builtins/setattr.def:L142` (functions)
- [ ] Flag `-n`: `third_party/bash/builtins/setattr.def:L142` (remove export attribute)
- [ ] Flag `-p`: `third_party/bash/builtins/setattr.def:L142` (print)

### `expr`

- [ ] Upstream: `third_party/coreutils/src/expr.c`
- [ ] Expression evaluation: Missing implementation
- [ ] Arithmetic (+, -, *, /, %): Missing implementation
- [ ] Comparison (=, !=, <, <=, >, >=): Missing implementation
- [ ] Logical (| , &): Missing implementation
- [ ] String operators (match, substr, index, length): `third_party/coreutils/src/expr.c:L186`
- [ ] Flag `--help`: `third_party/coreutils/src/expr.c:L157`
- [ ] Flag `--version`: `third_party/coreutils/src/expr.c:L160`

### `factor`

- [ ] Upstream: `third_party/coreutils/src/factor.c`
- [ ] Prime factorization: Missing implementation
- [ ] Flag `-h`, `--exponents`: `third_party/coreutils/src/factor.c:L2480`

### `false`

- [ ] Upstream: `third_party/bash/builtins/colon.def`, `third_party/coreutils/src/false.c`
- [ ] Basic operation: Missing implementation

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

- [ ] Upstream: `third_party/coreutils/src/fmt.c`
- [ ] Paragraph formatting: Missing implementation
- [ ] Flag `-c`, `--crown-margin`: `third_party/coreutils/src/fmt.c:L185`
- [ ] Flag `-p`, `--prefix=STRING`: `third_party/coreutils/src/fmt.c:L189`
- [ ] Flag `-s`, `--split-only`: `third_party/coreutils/src/fmt.c:L194`
- [ ] Flag `-t`, `--tagged-paragraph`: `third_party/coreutils/src/fmt.c:L198`
- [ ] Flag `-u`, `--uniform-spacing`: `third_party/coreutils/src/fmt.c:L202`
- [ ] Flag `-w`, `--width=WIDTH`: `third_party/coreutils/src/fmt.c:L206`
- [ ] Flag `-g`, `--goal=WIDTH`: `third_party/coreutils/src/fmt.c:L212`
- [ ] Flag `-WIDTH`: `third_party/coreutils/src/fmt.c:L178`

### `fold`

- [ ] Upstream: `third_party/coreutils/src/fold.c`
- [ ] Line wrapping: Missing implementation
- [ ] Flag `-b`, `--bytes`: `third_party/coreutils/src/fold.c:72`
- [ ] Flag `-c`, `--characters`: `third_party/coreutils/src/fold.c:76`
- [ ] Flag `-s`, `--spaces`: `third_party/coreutils/src/fold.c:80`
- [ ] Flag `-w`, `--width=WIDTH`: `third_party/coreutils/src/fold.c:84`
- [ ] Flag `-s`: `third_party/coreutils/src/fold.c:L96`
- [ ] Flag `-w`: `third_party/coreutils/src/fold.c:L100`

### `getlimits`

- [ ] Upstream: `third_party/coreutils/src/getlimits.c`
- [ ] Flag `--help`: `third_party/coreutils/src/getlimits.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/getlimits.c:L52`

### `getopts`

- [ ] Basic parsing: Missing implementation
- [ ] Silent mode support (`:`): `third_party/bash/builtins/getopts.def:L180`

### `groups`

- [ ] Upstream: `third_party/coreutils/src/groups.c`
- [ ] Basic listing: Missing implementation
- [ ] Multiple users support: `third_party/coreutils/src/groups.c:L120`

### `hash`

- [ ] Upstream: `third_party/bash/builtins/hash.def`
- [ ] Command hashing: Missing implementation
- [ ] Flag `-r`: `third_party/bash/builtins/hash.def:L124` (forget all)
- [ ] Flag `-d`: `third_party/bash/builtins/hash.def:L125` (forget name)
- [ ] Flag `-p`: `third_party/bash/builtins/hash.def:L126` (use path)
- [ ] Flag `-t`: `third_party/bash/builtins/hash.def:L127` (list name)
- [ ] Flag `-l`: `third_party/bash/builtins/hash.def:L128` (output format)

### `head`

- [ ] Basic output: Missing implementation
- [ ] Flag `-c`, `--bytes`: `third_party/coreutils/src/head.c:L121`
- [ ] Flag `-n`, `--lines`: `third_party/coreutils/src/head.c:L126`
- [ ] Flag `-q`, `--quiet`, `--silent`: `third_party/coreutils/src/head.c:L131`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/head.c:L135`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/head.c:L139`

### `help`

- [ ] Upstream: `third_party/bash/builtins/help.def`
- [ ] Basic discovery: Missing implementation
- [ ] Flag `-d`: `third_party/bash/builtins/help.def:L105` (short description)
- [ ] Flag `-m`: `third_party/bash/builtins/help.def:L108` (man-page format)
- [ ] Flag `-s`: `third_party/bash/builtins/help.def:L111` (syntax only)

### `history`

- [ ] Upstream: `third_party/bash/builtins/history.def`
- [ ] Basic management: Missing implementation
- [ ] Flag `-d offset`: `third_party/bash/builtins/history.def:L145` (delete entry)
- [ ] Flag `-a`: `third_party/bash/builtins/history.def:L126` (append)
- [ ] Flag `-c`: `third_party/bash/builtins/history.def:L129` (clear)
- [ ] Flag `-n`: `third_party/bash/builtins/history.def:L132` (read non-recorded)
- [ ] Flag `-p`: `third_party/bash/builtins/history.def:L148` (print/expand)
- [ ] Flag `-r`: `third_party/bash/builtins/history.def:L135` (read file)
- [ ] Flag `-s`: `third_party/bash/builtins/history.def:L141` (store/append)
- [ ] Flag `-w`: `third_party/bash/builtins/history.def:L138` (write file)

### `hostid`

- [ ] Upstream: `third_party/coreutils/src/hostid.c`
- [ ] Flag `--help`: `third_party/coreutils/src/hostid.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/hostid.c:L52`

### `hostname`

- [ ] Upstream: `third_party/coreutils/src/hostname.c`
- [ ] Basic output: Missing implementation
- [ ] Set hostname support: `third_party/coreutils/src/hostname.c:L95`
- [ ] Flag `--help`: `third_party/coreutils/src/hostname.c:L65`
- [ ] Flag `--version`: `third_party/coreutils/src/hostname.c:L65`

### `id`

- [ ] Upstream: `third_party/coreutils/src/id.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-G`: `third_party/coreutils/src/id.c:L113`
- [ ] Flag `-Z`: `third_party/coreutils/src/id.c:L105`
- [ ] Flag `-a`: `third_party/coreutils/src/id.c:L101`
- [ ] Flag `-g`: `third_party/coreutils/src/id.c:L109`
- [ ] Flag `-n`: `third_party/coreutils/src/id.c:L117`
- [ ] Flag `-r`: `third_party/coreutils/src/id.c:L121`
- [ ] Flag `-u`: `third_party/coreutils/src/id.c:L125`
- [ ] Flag `-z`: `third_party/coreutils/src/id.c:L129`

### `install`

- [ ] Upstream: `third_party/coreutils/src/install.c`
- [ ] Flag `-c`: `third_party/coreutils/src/install.c:L129` (ignored)
- [ ] Flag `-d`, `--directory`: `third_party/coreutils/src/install.c:L132`
- [ ] Flag `-D`: `third_party/coreutils/src/install.c:L135`
- [ ] Flag `-g`, `--group=GROUP`: `third_party/coreutils/src/install.c:L138`
- [ ] Flag `-m`, `--mode=MODE`: `third_party/coreutils/src/install.c:L141`
- [ ] Flag `-o`, `--owner=OWNER`: `third_party/coreutils/src/install.c:L144`
- [ ] Flag `-p`, `--preserve-timestamps`: `third_party/coreutils/src/install.c:L147`
- [ ] Flag `-s`, `--strip`: `third_party/coreutils/src/install.c:L150`
- [ ] Flag `-S`, `--suffix=SUFFIX`: `third_party/coreutils/src/install.c:L153`
- [ ] Flag `-t`, `--target-directory=DIR`: `third_party/coreutils/src/install.c:L156`
- [ ] Flag `-T`, `--no-target-directory`: `third_party/coreutils/src/install.c:L159`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/install.c:L162`
- [ ] Flag `-C`, `--compare`: `third_party/coreutils/src/install.c:L165`

### `jobs`

- [ ] Upstream: `third_party/bash/builtins/jobs.def`
- [ ] Flag `-l`: `third_party/bash/builtins/jobs.def:L94` (long format)
- [ ] Flag `-n`: `third_party/bash/builtins/jobs.def:L97` (only jobs that changed)
- [ ] Flag `-p`: `third_party/bash/builtins/jobs.def:L100` (only PIDs)
- [ ] Flag `-r`: `third_party/bash/builtins/jobs.def:L103` (running only)
- [ ] Flag `-s`: `third_party/bash/builtins/jobs.def:L106` (stopped only)
- [ ] Flag `-x command`: `third_party/bash/builtins/jobs.def:L109` (execute command)

### `join`

- [ ] Upstream: `third_party/coreutils/src/join.c`
- [ ] Basic join: Missing implementation
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

- [ ] Upstream: `third_party/bash/builtins/kill.def`
- [ ] Basic signaling: Missing implementation
- [ ] Flag `-n num`: `third_party/bash/builtins/kill.def:L130`
- [ ] Flag `-l`: `third_party/coreutils/src/kill.c:L277` / `third_party/bash/builtins/kill.def:L114`
- [ ] Flag `-s SIGNAL`: `third_party/coreutils/src/kill.c:L262` / `third_party/bash/builtins/kill.def:L129`

### `let`

- [ ] Upstream: `third_party/bash/builtins/let.def`
- [ ] Flag `--help`: `third_party/bash/builtins/let.def:L52` (hypothetical, usually handled as builtin)

### `link`

- [ ] Basic hard link: Missing implementation (exactly 2 args required)

### `ln`

- [ ] Basic link creation: Missing implementation
- [ ] Flag `-f`: `third_party/coreutils/src/ln.c:L553`
- [ ] Flag `-s`: `third_party/coreutils/src/ln.c:L574`
- [ ] Flag `-v`: `third_party/coreutils/src/ln.c:L595`

### `local`

- [ ] Upstream: `third_party/bash/builtins/declare.def`
- [ ] Inherits all `declare` attributes: `third_party/bash/builtins/declare.def`

### `logname`

- [ ] Upstream: `third_party/coreutils/src/logname.c`
- [ ] Flag `--help`: `third_party/coreutils/src/logname.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/logname.c:L44`

### `logout`

- [ ] Upstream: `third_party/bash/builtins/exit.def`
- [ ] Basic operation: Missing implementation

### `ls`

- [ ] Basic listing: Missing implementation
- [ ] Color output (`--color`): `third_party/coreutils/src/ls.c:L215`
- [ ] Flag `--author`: `third_party/coreutils/src/ls.c:L157`
- [ ] Flag `--block-size`: `third_party/coreutils/src/ls.c:L165`
- [ ] Flag `--color`: `third_party/coreutils/src/ls.c:L215`
- [ ] Flag `--dereference-command-line-symlink-to-dir`: `third_party/coreutils/src/ls.c:L274`
- [ ] Flag `--file-type`: `third_party/coreutils/src/ls.c:L237`
- [ ] Flag `--format`: `third_party/coreutils/src/ls.c:L241`
- [ ] Flag `--full-time`: `third_party/coreutils/src/ls.c:L246`
- [ ] Flag `--group-directories-first`: `third_party/coreutils/src/ls.c:L254`
- [ ] Flag `--hide`: `third_party/coreutils/src/ls.c:L285`
- [ ] Flag `--hyperlink`: `third_party/coreutils/src/ls.c:L289` (note: usage lists it)
- [ ] Flag `--indicator-style`: `third_party/coreutils/src/ls.c:L289`
- [ ] Flag `--quoting-style`: `third_party/coreutils/src/ls.c:L353`
- [ ] Flag `--show-control-chars`: `third_party/coreutils/src/ls.c:L345`
- [ ] Flag `--si`: `third_party/coreutils/src/ls.c:L266`
- [ ] Flag `--sort`: `third_party/coreutils/src/ls.c:L373`
- [ ] Flag `--time`: `third_party/coreutils/src/ls.c:L388`
- [ ] Flag `--time-style`: `third_party/coreutils/src/ls.c:L397`
- [ ] Flag `--zero`: `third_party/coreutils/src/ls.c:L435` (mapped to ZERO_OPTION)
- [ ] Flag `-1`: `third_party/coreutils/src/ls.c:L443`
- [ ] Flag `-A`: `third_party/coreutils/src/ls.c:L45`
- [ ] Flag `-B`: `third_party/coreutils/src/ls.c:L170`
- [ ] Flag `-C`: `third_party/coreutils/src/ls.c:L181`
- [ ] Flag `-D`: `third_party/coreutils/src/ls.c:L224`
- [ ] Flag `-F`: `third_party/coreutils/src/ls.c:L232`
- [ ] Flag `-G`: `third_party/coreutils/src/ls.c:L258`
- [ ] Flag `-H`: `third_party/coreutils/src/ls.c:L270`
- [ ] Flag `-I`: `third_party/coreutils/src/ls.c:L309`
- [ ] Flag `-L`: `third_party/coreutils/src/ls.c:L317`
- [ ] Flag `-N`: `third_party/coreutils/src/ls.c:L329`
- [ ] Flag `-Q`: `third_party/coreutils/src/ls.c:L349`
- [ ] Flag `-R`: `third_party/coreutils/src/ls.c:L361`
- [ ] Flag `-S`: `third_party/coreutils/src/ls.c:L369`
- [ ] Flag `-T`: `third_party/coreutils/src/ls.c:L411`
- [ ] Flag `-U`: `third_party/coreutils/src/ls.c:L419`
- [ ] Flag `-X`: `third_party/coreutils/src/ls.c:L435`
- [ ] Flag `-Z`: `third_party/coreutils/src/ls.c:L439`
- [ ] Flag `-a`: `third_party/coreutils/src/ls.c:L41`
- [ ] Flag `-b`: `third_party/coreutils/src/ls.c:L161`
- [ ] Flag `-c`: `third_party/coreutils/src/ls.c:L174`
- [ ] Flag `-d`: `third_party/coreutils/src/ls.c:L220`
- [ ] Flag `-f`: `third_party/coreutils/src/ls.c:L228`
- [ ] Flag `-g`: `third_party/coreutils/src/ls.c:L250`
- [ ] Flag `-h`: `third_party/coreutils/src/ls.c:L262`
- [ ] Flag `-i`: `third_party/coreutils/src/ls.c:L305`
- [ ] Flag `-k`: `third_party/coreutils/src/ls.c:L313`
- [ ] Flag `-l`: `third_party/coreutils/src/ls.c`
- [ ] Flag `-m`: `third_party/coreutils/src/ls.c:L321`
- [ ] Flag `-n`: `third_party/coreutils/src/ls.c:L325`
- [ ] Flag `-o`: `third_party/coreutils/src/ls.c:L333`
- [ ] Flag `-p`: `third_party/coreutils/src/ls.c:L337`
- [ ] Flag `-q`: `third_party/coreutils/src/ls.c:L341`
- [ ] Flag `-r`: `third_party/coreutils/src/ls.c:L357`
- [ ] Flag `-s`: `third_party/coreutils/src/ls.c:L365`
- [ ] Flag `-t`: `third_party/coreutils/src/ls.c:L384`
- [ ] Flag `-u`: `third_party/coreutils/src/ls.c:L415`
- [ ] Flag `-v`: `third_party/coreutils/src/ls.c:L423`
- [ ] Flag `-w`: `third_party/coreutils/src/ls.c:L427`
- [ ] Flag `-x`: `third_party/coreutils/src/ls.c:L431`

### `mapfile`

- [ ] Upstream: `third_party/bash/builtins/mapfile.def`
- [ ] Array population: Missing implementation
- [ ] Flag `-d`: `third_party/bash/builtins/mapfile.def:L238` (delimiter)
- [ ] Flag `-n`: `third_party/bash/builtins/mapfile.def:L238` (count)
- [ ] Flag `-O`: `third_party/bash/builtins/mapfile.def:L238` (origin)
- [ ] Flag `-t`: `third_party/bash/builtins/mapfile.def:L238` (strip newline)
- [ ] Flag `-u`: `third_party/bash/builtins/mapfile.def:L238` (fd)
- [ ] Flag `-C`: `third_party/bash/builtins/mapfile.def:L238` (callback)
- [ ] Flag `-c`: `third_party/bash/builtins/mapfile.def:L238` (quantum)
- [ ] Flag `-s`: `third_party/bash/builtins/mapfile.def:L76` (array is same)
- [ ] Aliases: `readarray`

### `md5sum`

- [ ] Upstream: `third_party/coreutils/src/cksum.c`
- [ ] Inherits all `cksum` hash flags: `third_party/coreutils/src/cksum.c`

### `mkdir`

- [ ] Upstream: `third_party/coreutils/src/mkdir.c`
- [ ] Flag `-m`, `--mode=MODE`: `third_party/coreutils/src/mkdir.c:L65`
- [ ] Flag `-p`, `--parents`: `third_party/coreutils/src/mkdir.c:L69`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/mkdir.c:L74`
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

- [ ] Upstream: `third_party/coreutils/src/mktemp.c`
- [ ] Basic creation: Missing implementation
- [ ] Flag `--suffix`: `third_party/coreutils/src/mktemp.c:L88`
- [ ] Flag `-d`: `third_party/coreutils/src/mktemp.c:L76`
- [ ] Flag `-p`: `third_party/coreutils/src/mktemp.c:L93`
- [ ] Flag `-q`: `third_party/coreutils/src/mktemp.c:L84`
- [ ] Flag `-t`: `third_party/coreutils/src/mktemp.c:101`
- [ ] Flag `-u`: `third_party/coreutils/src/mktemp.c:L80`

### `mv`

- [ ] Basic move/rename: Missing implementation
- [ ] Flag `-b`, `--backup`: `third_party/coreutils/src/mv.c:L278`
- [ ] Flag `-f`, `--force`: `third_party/coreutils/src/mv.c:L282`
- [ ] Flag `-i`, `--interactive`: `third_party/coreutils/src/mv.c:L286`
- [ ] Flag `-n`, `--no-clobber`: `third_party/coreutils/src/mv.c:L290`
- [ ] Flag `-t`, `--target-directory`: `third_party/coreutils/src/mv.c:L294`
- [ ] Flag `-T`, `--no-target-directory`: `third_party/coreutils/src/mv.c:L298`
- [ ] Flag `-u`, `--update`: `third_party/coreutils/src/mv.c:L302`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/mv.c:L306`
- [ ] Flag `-Z`, `--context`: `third_party/coreutils/src/mv.c:L310`
- [ ] Flag `--exchange`: `third_party/coreutils/src/mv.c:L314`
- [ ] Flag `--no-copy`: `third_party/coreutils/src/mv.c:L318`

### `nice`

- [ ] Upstream: `third_party/coreutils/src/nice.c`
- [ ] Priority adjustment: Missing implementation
- [ ] Flag `-n`, `--adjustment=N`: `third_party/coreutils/src/nice.c:L160`

### `nl`

- [ ] Upstream: `third_party/coreutils/src/nl.c`
- [ ] Basic numbering: Missing implementation
- [ ] Flag `-b STYLE`: `third_party/coreutils/src/nl.c:L129`
- [ ] Flag `-d CC`: `third_party/coreutils/src/nl.c:L132`
- [ ] Flag `-f STYLE`: `third_party/coreutils/src/nl.c:L135`
- [ ] Flag `-h STYLE`: `third_party/coreutils/src/nl.c:L138`
- [ ] Flag `-i NUMBER`: `third_party/coreutils/src/nl.c:L141`
- [ ] Flag `-l NUMBER`: `third_party/coreutils/src/nl.c:L144`
- [ ] Flag `-n FORMAT`: `third_party/coreutils/src/nl.c:L147`
- [ ] Flag `-p`: `third_party/coreutils/src/nl.c:L150`
- [ ] Flag `-s STRING`: `third_party/coreutils/src/nl.c:L153`
- [ ] Flag `-v NUMBER`: `third_party/coreutils/src/nl.c:L156`
- [ ] Flag `-w NUMBER`: `third_party/coreutils/src/nl.c:L159`
- [ ] Flag `-b`: `third_party/coreutils/src/nl.c:L153`
- [ ] Flag `-d`: `third_party/coreutils/src/nl.c:L162`
- [ ] Flag `-f`: `third_party/coreutils/src/nl.c:L154`
- [ ] Flag `-h`: `third_party/coreutils/src/nl.c:L152`
- [ ] Flag `-i`: `third_party/coreutils/src/nl.c:L156`
- [ ] Flag `-l`: `third_party/coreutils/src/nl.c:L158`
- [ ] Flag `-n`: `third_party/coreutils/src/nl.c:L161`
- [ ] Flag `-p`: `third_party/coreutils/src/nl.c:L157`
- [ ] Flag `-s`: `third_party/coreutils/src/nl.c:L159`
- [ ] Flag `-v`: `third_party/coreutils/src/nl.c:L155`
- [ ] Flag `-w`: `third_party/coreutils/src/nl.c:L160`

### `nohup`

- [ ] Upstream: `third_party/coreutils/src/nohup.c`
- [ ] Flag `--help`: `third_party/coreutils/src/nohup.c:L59`
- [ ] Flag `--version`: `third_party/coreutils/src/nohup.c:L59`

### `nproc`

- [ ] Upstream: `third_party/coreutils/src/nproc.c`

### `numfmt`

- [ ] Upstream: `third_party/coreutils/src/numfmt.c`
- [ ] Flag `-d`, `--delimiter=X`: `third_party/coreutils/src/numfmt.c:L1022`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/numfmt.c:L1053`

### `od`

- [ ] Upstream: `third_party/coreutils/src/od.c`
- [ ] Format output: Missing implementation
- [ ] Flag `-A rad`: `third_party/coreutils/src/od.c:L316` (address radix)
- [ ] Flag `-j bytes`: `third_party/coreutils/src/od.c:L315` (skip bytes)
- [ ] Flag `-N bytes`: `third_party/coreutils/src/od.c:L317` (read bytes)
- [ ] Flag `-t type`: `third_party/coreutils/src/od.c:L318` (format spec)
- [ ] Flag `-v`: `third_party/coreutils/src/od.c:L319` (output duplicates)
- [ ] Flag `-w`: `third_party/coreutils/src/od.c:L322` (width)
- [ ] Flag `-S`: `third_party/coreutils/src/od.c:L320` (strings)

### `paste`

- [ ] Upstream: `third_party/coreutils/src/paste.c`
- [ ] Basic paste: Missing implementation
- [ ] Flag `-d`, `--delimiters=LIST`: `third_party/coreutils/src/paste.c:L468`
- [ ] Flag `-s`, `--serial`: `third_party/coreutils/src/paste.c:L473`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/paste.c:L479`

### `pathchk`

- [ ] Upstream: `third_party/coreutils/src/pathchk.c`
- [ ] Flag `-p`: `third_party/coreutils/src/pathchk.c:L360`
- [ ] Flag `-P`: `third_party/coreutils/src/pathchk.c:L363`

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

- [ ] Upstream: `third_party/bash/builtins/pushd.def`
- [ ] Flag `-n`: `third_party/bash/builtins/pushd.def:L165` (don't rotate)

### `pr`

- [ ] Upstream: `third_party/coreutils/src/pr.c`
- [ ] Print formatting: Missing implementation
- [ ] Flag `-a`: `third_party/coreutils/src/pr.c:L316` (across)
- [ ] Flag `-d`: `third_party/coreutils/src/pr.c:L318` (double space)
- [ ] Flag `-h`: `third_party/coreutils/src/pr.c:L322` (header)
- [ ] Flag `-l`: `third_party/coreutils/src/pr.c:L325` (length)
- [ ] Flag `-m`: `third_party/coreutils/src/pr.c:L326` (merge)
- [ ] Flag `-n`: `third_party/coreutils/src/pr.c:L327` (number)
- [ ] Flag `-t`: `third_party/coreutils/src/pr.c:L333` (omit header)
- [ ] Flag `-w`: `third_party/coreutils/src/pr.c:L336` (width)

### `printenv`

- [ ] Upstream: `third_party/coreutils/src/printenv.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-0`: `third_party/coreutils/src/printenv.c:L70`

### `printf`

- [ ] Basic formatting: Missing implementation
- [ ] Flag `%b`: `third_party/bash/builtins/printf.def:L558`
- [ ] Flag `%q`: `third_party/bash/builtins/printf.def:L672`
- [ ] Flag `-v VAR`: `third_party/bash/builtins/printf.def:L301`

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

- [ ] Upstream: `third_party/bash/builtins/pushd.def`
- [ ] Flag `-n`: `third_party/bash/builtins/pushd.def:L129` (don't change directory)

### `pwd`

- [ ] Upstream: `third_party/bash/builtins/cd.def`
- [ ] Basic path reporting: Missing implementation
- [-] Flag `--help`: Handled by the shell's global help dispatcher.
- [ ] Flag `-L`: `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L291-316`
- [ ] Flag `-P`: `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L371-383`

### `read`

- [ ] Upstream: `third_party/bash/builtins/read.def`
- [ ] Basic input: Missing implementation
- [ ] Flag `-a`, `--array`: `third_party/bash/builtins/read.def:L39`
- [ ] Flag `-d`, `--delimiter`: `third_party/bash/builtins/read.def:L41`
- [ ] Flag `-e`: `third_party/bash/builtins/read.def:L43` (use Readline)
- [ ] Flag `-i`, `--initial-text`: `third_party/bash/builtins/read.def:L49`
- [ ] Flag `-n`, `--nchars`: `third_party/bash/builtins/read.def:L50`
- [ ] Flag `-N`, `--Nchars`: `third_party/bash/builtins/read.def:L51`
- [ ] Flag `-p`, `--prompt`: `third_party/bash/builtins/read.def:L53`
- [ ] Flag `-r`: `third_party/bash/builtins/read.def:L55` (raw mode)
- [ ] Flag `-s`, `--silent`: `third_party/bash/builtins/read.def:L56`
- [ ] Flag `-t`, `--timeout`: `third_party/bash/builtins/read.def:L57`
- [ ] Flag `-u`, `--fd`: `third_party/bash/builtins/read.def:L61`

### `readlink`

- [ ] Upstream: `third_party/coreutils/src/readlink.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-e`, `--canonicalize-existing`: `third_party/coreutils/src/readlink.c:L44`
- [ ] Flag `-m`, `--canonicalize-missing`: `third_party/coreutils/src/readlink.c:L45`
- [ ] Flag `-q`, `--quiet`: `third_party/coreutils/src/readlink.c:L46`
- [ ] Flag `-s`, `--silent`: `third_party/coreutils/src/readlink.c:L47`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/readlink.c:L48`
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/readlink.c:L49`
- [ ] Flag `-f`: `third_party/coreutils/src/readlink.c:L135`
- [ ] Flag `-n`: `third_party/coreutils/src/readlink.c:L141`

### `readonly`

- [ ] Upstream: `third_party/bash/builtins/setattr.def`
- [ ] Attribute management: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/setattr.def:L181` (indexed array)
- [ ] Flag `-A`: `third_party/bash/builtins/setattr.def:L181` (associative array)
- [ ] Flag `-p`: `third_party/bash/builtins/setattr.def:L181` (print)
- [ ] Flag `-f`: `third_party/bash/builtins/setattr.def:L181` (functions)

### `realpath`

- [ ] Upstream: `third_party/coreutils/src/realpath.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-E`, `--canonicalize-existing`: `third_party/coreutils/src/realpath.c:L44`
- [ ] Flag `-L`, `--logical`: `third_party/coreutils/src/realpath.c:L45`
- [ ] Flag `-P`, `--physical`: `third_party/coreutils/src/realpath.c:L46`
- [ ] Flag `-q`, `--quiet`: `third_party/coreutils/src/realpath.c:L47`
- [ ] Flag `-s`, `--strip`: `third_party/coreutils/src/realpath.c:L48`
- [ ] Flag `-z`, `--zero`: `third_party/coreutils/src/realpath.c:L49`
- [ ] Flag `--relative-to`: `third_party/coreutils/src/realpath.c:L246`
- [ ] Flag `-e`: `third_party/coreutils/src/realpath.c:L220`
- [ ] Flag `-m`: `third_party/coreutils/src/realpath.c:L224`

### `return`

- [ ] Upstream: `third_party/bash/builtins/return.def`
- [ ] Basic return: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/return.def:L61`

### `rm`

- [ ] Upstream: `third_party/coreutils/src/rm.c`
- [ ] Basic removal: Missing implementation
- [ ] Flag `-I`: `third_party/coreutils/src/rm.c:L157` (prompt once)
- [ ] Flag `-d`, `--dir`: `third_party/coreutils/src/rm.c:L162` (remove empty directories)
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/rm.c:L167`
- [ ] Flag `-f`: `third_party/coreutils/src/rm.c:L137`
- [ ] Flag `-i`: `third_party/coreutils/src/rm.c:L142`
- [ ] Flag `-r`: `third_party/coreutils/src/rm.c:L172`

### `rmdir`

- [ ] Upstream: `third_party/coreutils/src/rmdir.c`
- [ ] Basic rmdir: Missing implementation
- [ ] Basic removal: Missing implementation
- [ ] Flag `--ignore-fail-on-non-empty`: `third_party/coreutils/src/rmdir.c:L178`
- [ ] Flag `-p`: `third_party/coreutils/src/rmdir.c:L182`
- [ ] Flag `-v`: `third_party/coreutils/src/rmdir.c:L187`

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

- [ ] Upstream: `third_party/coreutils/src/seq.c`
- [ ] Basic sequence: Missing implementation
- [ ] Flag `-f`, `--format=FORMAT`: `third_party/coreutils/src/seq.c:L592`
- [ ] Flag `-s`, `--separator=STRING`: `third_party/coreutils/src/seq.c:L596`
- [ ] Flag `-w`, `--equal-width`: `third_party/coreutils/src/seq.c:L600`
- [ ] Flag `-s`: `third_party/coreutils/src/seq.c:L596`
- [ ] Flag `-w`: `third_party/coreutils/src/seq.c:L600`

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

### `sha1sum`

- [ ] Upstream: `third_party/coreutils/src/coreutils-sha1sum.c`
- [ ] Inherits flags from `cksum`

### `sha224sum`

- [ ] Upstream: `third_party/coreutils/src/coreutils-sha224sum.c`
- [ ] Inherits flags from `cksum`

### `sha256sum`

- [ ] Upstream: `third_party/coreutils/src/coreutils-sha256sum.c`
- [ ] Inherits flags from `cksum`

### `sha384sum`

- [ ] Upstream: `third_party/coreutils/src/coreutils-sha384sum.c`
- [ ] Inherits flags from `cksum`

### `sha512sum`

- [ ] Upstream: `third_party/coreutils/src/coreutils-sha512sum.c`
- [ ] Inherits flags from `cksum`

### `shift`

- [ ] Upstream: `third_party/bash/builtins/shift.def`
- [ ] Basic shift: Missing implementation
- [ ] Shifting n parameters: `third_party/bash/builtins/shift.def:L64`

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

- [ ] Upstream: `third_party/coreutils/src/shuf.c`
- [ ] Basic shuffling: Missing implementation
- [ ] Flag `--random-source=FILE`: `third_party/coreutils/src/shuf.c:L111`
- [ ] Flag `-e`: `third_party/coreutils/src/shuf.c:L107`
- [ ] Flag `-i`: `third_party/coreutils/src/shuf.c:L108`
- [ ] Flag `-n`: `third_party/coreutils/src/shuf.c:L109`
- [ ] Flag `-o`: `third_party/coreutils/src/shuf.c:L110`
- [ ] Flag `-r`: `third_party/coreutils/src/shuf.c:L112`
- [ ] Flag `-z`: `third_party/coreutils/src/shuf.c:L113`

### `sleep`

- [ ] Upstream: `third_party/coreutils/src/sleep.c`
- [ ] Basic sleep: Missing implementation
- [ ] Multiple arguments (sum): `third_party/coreutils/src/sleep.c:L135`
- [ ] Suffixes (s, m, h, d): `third_party/coreutils/src/sleep.c:L65`
- [ ] Flag `--help`: `third_party/coreutils/src/sleep.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/sleep.c:L44`

### `sort`

- [ ] Upstream: `third_party/coreutils/src/sort.c`
- [ ] Basic sorting: Missing implementation
- [ ] Ordering flags (`-b`, `-i`, `-d`, `-f`, `-g`, `-h`, `-n`, `-M`, `-R`, `-V`, `-r`): `third_party/coreutils/src/sort.c:L437-490`
- [ ] Flag `-b`, `--ignore-leading-blanks`: `third_party/coreutils/src/sort.c:L437`
- [ ] Flag `-c`, `-C`, `--check`: `third_party/coreutils/src/sort.c:L441`
- [ ] Flag `-d`, `--dictionary-order`: `third_party/coreutils/src/sort.c:L451`
- [ ] Flag `-f`, `--ignore-case`: `third_party/coreutils/src/sort.c:L455`
- [ ] Flag `-g`, `--general-numeric-sort`: `third_party/coreutils/src/sort.c:L459`
- [ ] Flag `-h`, `--human-numeric-sort`: `third_party/coreutils/src/sort.c:L463`
- [ ] Flag `-i`, `--ignore-nonprinting`: `third_party/coreutils/src/sort.c:L467`
- [ ] Flag `-k`, `--key=KEYDEF`: `third_party/coreutils/src/sort.c:L473`
- [ ] Flag `-m`, `--merge`: `third_party/coreutils/src/sort.c:L477`
- [ ] Flag `-M`, `--month-sort`: `third_party/coreutils/src/sort.c:L481`
- [ ] Flag `-n`, `--numeric-sort`: `third_party/coreutils/src/sort.c:L485`
- [ ] Flag `-o`, `--output=FILE`: `third_party/coreutils/src/sort.c:L489`
- [ ] Flag `-r`, `--reverse`: `third_party/coreutils/src/sort.c:L493`
- [ ] Flag `-s`, `--stable`: `third_party/coreutils/src/sort.c:L497`
- [ ] Flag `-S`, `--buffer-size=SIZE`: `third_party/coreutils/src/sort.c:L501`
- [ ] Flag `-t`, `--field-separator=SEP`: `third_party/coreutils/src/sort.c:L505`
- [ ] Flag `-T`, `--temporary-directory=DIR`: `third_party/coreutils/src/sort.c:L509`
- [ ] Flag `-u`, `--unique`: `third_party/coreutils/src/sort.c:L513`
- [ ] Flag `-V`, `--version-sort`: `third_party/coreutils/src/sort.c:L517`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/sort.c:L521`
- [ ] Flag `--parallel=N`: `third_party/coreutils/src/sort.c:L525`
- [ ] Flag `--random-sort` (`-R`): `third_party/coreutils/src/sort.c:L529`
- [ ] Flag `--debug`: `third_party/coreutils/src/sort.c:L533`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/sort.c:L537`

### `source`

- [ ] Upstream: `third_party/bash/builtins/source.def`
- [ ] Basic sourcing: Missing implementation
- [ ] Aliases: `.`
- [ ] Flag `-p path`: `third_party/bash/builtins/source.def:L126`

### `split`

- [ ] Basic split: Missing implementation
- [ ] Flag `--filter`: `third_party/coreutils/src/split.c:L274`
- [ ] Flag `--verbose`: `third_party/coreutils/src/split.c:L295`
- [ ] Flag `-C`: `third_party/coreutils/src/split.c:L250`
- [ ] Flag `-a`: `third_party/coreutils/src/split.c:L238`
- [ ] Flag `-b`: `third_party/coreutils/src/split.c:L246`
- [ ] Flag `-d`: `third_party/coreutils/src/split.c:L254`
- [ ] Flag `-e`: `third_party/coreutils/src/split.c:L270`
- [ ] Flag `-l`: `third_party/coreutils/src/split.c:L278`
- [ ] Flag `-n`: `third_party/coreutils/src/split.c:L282`
- [ ] Flag `-t`: `third_party/coreutils/src/split.c:L286`
- [ ] Flag `-u`: `third_party/coreutils/src/split.c:L291`
- [ ] Flag `-x`: `third_party/coreutils/src/split.c:L262`

### `stat`

- [ ] Upstream: `third_party/coreutils/src/stat.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-c`, `--format=FORMAT`: `third_party/coreutils/src/stat.c:L1921`
- [ ] Flag `-f`, `--file-system`: `third_party/coreutils/src/stat.c:L1932`
- [ ] Flag `-L`, `--dereference`: `third_party/coreutils/src/stat.c:L1927`
- [ ] Flag `-t`, `--terse`: `third_party/coreutils/src/stat.c:L1936`
- [ ] Flag `--printf=FORMAT`: `third_party/coreutils/src/stat.c:L1924`
- [ ] Flag `--cached={always,never,default}`: `third_party/coreutils/src/stat.c:L1917`

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

- [ ] Upstream: `third_party/coreutils/src/sync.c`
- [ ] Flush buffers: Missing implementation
- [ ] Flag `-d`, `--data`: `third_party/coreutils/src/sync.c:L129`
- [ ] Flag `-f`, `--file-system`: `third_party/coreutils/src/sync.c:L132`

### `tac`

- [ ] Upstream: `third_party/coreutils/src/tac.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-b`: `third_party/coreutils/src/tac.c:L103`
- [ ] Flag `-r`: `third_party/coreutils/src/tac.c:L104`
- [ ] Flag `-s`: `third_party/coreutils/src/tac.c:L105`

### `tail`

- [ ] Basic output: Missing implementation
- [ ] Flag `-c`, `--bytes`: `third_party/coreutils/src/tail.c:L296`
- [ ] Flag `-f`, `--follow[={name|descriptor}]`: `third_party/coreutils/src/tail.c:L305`
- [ ] Flag `-F`: `third_party/coreutils/src/tail.c:L309` (implies --follow=name --retry)
- [ ] Flag `-n`, `--lines`: `third_party/coreutils/src/tail.c:L314`
- [ ] Flag `-q`, `--quiet`, `--silent`: `third_party/coreutils/src/tail.c:L318`
- [ ] Flag `-s`, `--sleep-interval`: `third_party/coreutils/src/tail.c:L322`
- [ ] Flag `-v`, `--verbose`: `third_party/coreutils/src/tail.c:L326`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/tail.c:L330`
- [ ] Flag `--pid`: `third_party/coreutils/src/tail.c:L334`
- [ ] Flag `--retry`: `third_party/coreutils/src/tail.c:L338`
- [ ] Flag `--max-unchanged-stats`: `third_party/coreutils/src/tail.c:L342`

### `tee`

- [ ] Basic copy: Missing implementation
- [ ] Flag `-a`, `--append`: `third_party/coreutils/src/tee.c:L93`
- [ ] Flag `-i`, `--ignore-interrupts`: `third_party/coreutils/src/tee.c:L97`
- [ ] Flag `-p`, `--output-error[=MODE]`: `third_party/coreutils/src/tee.c:L101`

### `test`

- [ ] Unary operators (-e, -f, -d, etc.): Missing implementation
- [ ] String operators (=, !=, -z, -n): Missing implementation
- [ ] Numeric operators (-eq, -ne, etc.): Missing implementation
- [ ] Logical operators (!, -a, -o): Missing implementation
- [ ] Aliases: `[`

### `time`

- [ ] Upstream: `third_party/bash/builtins/reserved.def`
- [ ] Basic operation: Missing implementation

### `timeout`

- [ ] Upstream: `third_party/coreutils/src/timeout.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `--help`: `third_party/coreutils/src/timeout.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/timeout.c:L44`
- [ ] Flag `-k`: `third_party/coreutils/src/timeout.c:L531`
- [ ] Flag `-s`: `third_party/coreutils/src/timeout.c:L539`

### `times`

- [ ] Upstream: `third_party/bash/builtins/times.def`
- [ ] Basic output: Missing implementation

### `touch`

- [ ] Upstream: `third_party/coreutils/src/touch.c`
- [ ] Basic touch: Missing implementation
- [ ] Basic timestamp update: Missing implementation
- [ ] Flag `-t STAMP`: `third_party/coreutils/src/touch.c:L259` (explicit timestamp)
- [ ] Flag `-a`: `third_party/coreutils/src/touch.c:L230`
- [ ] Flag `-c`: `third_party/coreutils/src/touch.c:L234`
- [ ] Flag `-d`: `third_party/coreutils/src/touch.c:L238`
- [ ] Flag `-h`: `third_party/coreutils/src/touch.c:L246`
- [ ] Flag `-m`: `third_party/coreutils/src/touch.c:L251`
- [ ] Flag `-r`: `third_party/coreutils/src/touch.c:L255`
- [ ] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `third_party/coreutils/src/touch.c:L259`

### `tr`

- [ ] Upstream: `third_party/coreutils/src/tr.c`
- [ ] Basic translation: Missing implementation
- [ ] Flag `-c`, `-C`, `--complement`: `third_party/coreutils/src/tr.c:L296`
- [ ] Flag `-d`, `--delete`: `third_party/coreutils/src/tr.c:L300`
- [ ] Flag `-s`, `--squeeze-repeats`: `third_party/coreutils/src/tr.c:L304`
- [ ] Flag `-t`, `--truncate-set1`: `third_party/coreutils/src/tr.c:L310`

### `trap`

- [ ] Upstream: `third_party/bash/builtins/trap.def`
- [ ] Basic trapping: Missing implementation
- [ ] Flag `-P`: `third_party/bash/builtins/trap.def:L131`
- [ ] Flag `-l`: `third_party/bash/builtins/trap.def:L125`
- [ ] Flag `-p`: `third_party/bash/builtins/trap.def:L128`

### `true`

- [ ] Basic operation: Missing implementation

### `truncate`

- [ ] Upstream: `third_party/coreutils/src/truncate.c`
- [ ] Basic truncation: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/truncate.c:L82`
- [ ] Flag `-o`: `third_party/coreutils/src/truncate.c:L85`
- [ ] Flag `-r`: `third_party/coreutils/src/truncate.c:L88`
- [ ] Flag `-s`: `third_party/coreutils/src/truncate.c:L91`

### `tsort`

- [ ] Upstream: `third_party/coreutils/src/tsort.c`
- [ ] Flag `--help`: `third_party/coreutils/src/tsort.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/tsort.c:L52`

### `tty`

- [ ] Upstream: `third_party/coreutils/src/tty.c`
- [ ] TTY reporting: Missing implementation
- [ ] Flag `-s`, `--silent`, `--quiet`: `third_party/coreutils/src/tty.c:L71`

### `type`

- [ ] Upstream: `third_party/bash/builtins/type.def`
- [ ] Command identification: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/type.def:L129` (all occurrences)
- [ ] Flag `-p`: `third_party/bash/builtins/type.def:L130` (path only)
- [ ] Flag `-t`: `third_party/bash/builtins/type.def:L131` (type only)
- [ ] Flag `-f`: `third_party/bash/builtins/type.def:L132` (skip functions)
- [ ] Flag `-P`: `third_party/bash/builtins/type.def:L133` (force path search)

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

- [ ] Upstream: `third_party/bash/builtins/umask.def`
- [ ] Basic mask management: Missing implementation
- [ ] Flag `-S`: `third_party/bash/builtins/umask.def:L88`
- [ ] Flag `-p`: `third_party/bash/builtins/umask.def:L91`

### `unalias`

- [ ] Upstream: `third_party/bash/builtins/alias.def`
- [ ] Remove aliases: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/unalias.def:L165`

### `uname`

- [ ] Upstream: `third_party/coreutils/src/uname.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-a`: `third_party/coreutils/src/uname.c:L123`
- [ ] Flag `-i`: `third_party/coreutils/src/uname.c:L145`
- [ ] Flag `-m`: `third_party/coreutils/src/uname.c:L139`
- [ ] Flag `-n`: `third_party/coreutils/src/uname.c:L130`
- [ ] Flag `-o`: `third_party/coreutils/src/uname.c:L148`
- [ ] Flag `-p`: `third_party/coreutils/src/uname.c:L142`
- [ ] Flag `-r`: `third_party/coreutils/src/uname.c:L133`
- [ ] Flag `-s`: `third_party/coreutils/src/uname.c:L127`
- [ ] Flag `-v`: `third_party/coreutils/src/uname.c:L136`

### `unexpand`

- [ ] Upstream: `third_party/coreutils/src/unexpand.c`
- [ ] Basic conversion: Missing implementation
- [ ] Flag `-a`, `--all`: `third_party/coreutils/src/unexpand.c:L87`
- [ ] Flag `-t`, `--tabs=LIST`: `third_party/coreutils/src/unexpand.c:L95`
- [ ] Flag `--first-only`: `third_party/coreutils/src/unexpand.c:L91`

### `uniq`

- [ ] Upstream: `third_party/coreutils/src/uniq.c`
- [ ] Basic filtering: Missing implementation
- [ ] Flag `-c`, `--count`: `third_party/coreutils/src/uniq.c:L172`
- [ ] Flag `-d`, `--repeated`: `third_party/coreutils/src/uniq.c:L176`
- [ ] Flag `-D`, `--all-repeated[=METHOD]`: `third_party/coreutils/src/uniq.c:180`
- [ ] Flag `-f`, `--skip-fields=N`: `third_party/coreutils/src/uniq.c:189`
- [ ] Flag `-i`, `--ignore-case`: `third_party/coreutils/src/uniq.c:198`
- [ ] Flag `-s`, `--skip-chars=N`: `third_party/coreutils/src/uniq.c:202`
- [ ] Flag `-u`, `--unique`: `third_party/coreutils/src/uniq.c:206`
- [ ] Flag `-w`, `--check-chars=N`: `third_party/coreutils/src/uniq.c:214`
- [ ] Flag `-z`, `--zero-terminated`: `third_party/coreutils/src/uniq.c:210`
- [ ] Flag `--group[=METHOD]`: `third_party/coreutils/src/uniq.c:184`

### `unlink`

- [ ] Upstream: `third_party/coreutils/src/unlink.c`
- [ ] Basic removal: Missing implementation (exactly 1 arg required)
- [ ] Flag `--help`: `third_party/coreutils/src/unlink.c:L52`
- [ ] Flag `--version`: `third_party/coreutils/src/unlink.c:L52`

### `unset`

- [ ] Upstream: `third_party/bash/builtins/set.def`
- [ ] Flag `-f`: `third_party/bash/builtins/set.def:L643` (functions)
- [ ] Flag `-v`: `third_party/bash/builtins/set.def:L637` (variables)
- [ ] Flag `-n`: `third_party/bash/builtins/set.def:L640` (nameref)

### `uptime`

- [ ] Upstream: `third_party/coreutils/src/uptime.c`
- [ ] Flag `-p`, `--pretty`: `third_party/coreutils/src/uptime.c:L130`
- [ ] Flag `-s`, `--since`: `third_party/coreutils/src/uptime.c:L135`

### `users`

- [ ] Upstream: `third_party/coreutils/src/users.c`
- [ ] Flag `--help`: `third_party/coreutils/src/users.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/users.c:L44`

### `vdir`

- [ ] Upstream: `third_party/coreutils/src/ls.c`

### `wait`

- [ ] Upstream: `third_party/bash/builtins/wait.def`
- [ ] Basic waiting: Missing implementation
- [ ] Optional: jobspec or process ID
- [ ] Flag `-f`: `third_party/bash/builtins/wait.def:L134`
- [ ] Flag `-n`: `third_party/bash/builtins/wait.def:L131`
- [ ] Flag `-p var`: `third_party/bash/builtins/wait.def:L137`

### `wc`

- [ ] Basic counts: Missing implementation
- [ ] Flag `-c`, `--bytes`: `third_party/coreutils/src/wc.c:L188`
- [ ] Flag `-m`, `--chars`: `third_party/coreutils/src/wc.c:L192`
- [ ] Flag `-l`, `--lines`: `third_party/coreutils/src/wc.c:L196`
- [ ] Flag `-w`, `--words`: `third_party/coreutils/src/wc.c:L214`
- [ ] Flag `-L`, `--max-line-length`: `third_party/coreutils/src/wc.c:L218`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/wc.c:L222`
- [ ] Flag `--total={auto,always,only,never}`: `third_party/coreutils/src/wc.c:L226`

### `who`

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

- [ ] Upstream: `third_party/coreutils/src/whoami.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `--help`: `third_party/coreutils/src/whoami.c:L44`
- [ ] Flag `--version`: `third_party/coreutils/src/whoami.c:L44`

### `yes`

- [ ] Upstream: `third_party/coreutils/src/yes.c`
- [ ] Basic operation: Missing implementation
- [ ] Basic repetition: Missing implementation
- [ ] Flag `--help`: `third_party/coreutils/src/yes.c:L48`
- [ ] Flag `--version`: `third_party/coreutils/src/yes.c:L48`

