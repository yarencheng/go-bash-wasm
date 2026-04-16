# Functional Parity Tracking

This document tracks the alignment of the Go Bash Simulator with upstream GNU implementations.

## Overview

Status codes:
- [x] : Fully implemented and verified.
- [ ] : Missing or incomplete.
- [-] : Deliberately skipped.

---

## Parity Matrix


### `alias`

- [ ] Upstream: `third_party/bash/builtins/alias.def`
- [ ] Basic management: Missing implementation
- [ ] Define/Display aliases: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/alias.def:L181`
- [ ] Flag `-p`: `third_party/bash/builtins/alias.def:L36`

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
- [ ] Flag `-a`: `third_party/coreutils/src/basename.c:L155`
- [ ] Flag `-s`: `third_party/coreutils/src/basename.c:L150`
- [ ] Flag `-z`: `third_party/coreutils/src/basename.c:L159`

### `basenc`

- [ ] Upstream: `third_party/coreutils/src/basenc.c`
- [ ] Flags to implement: -d, -i, -w

### `bind`

- [ ] Upstream: `third_party/bash/builtins/bind.def`
- [ ] Flags to implement: -P, -S, -V, -X, -f, -l, -lpvsPVS, -m, -p, -q, -r, -s, -u, -v, -x

### `break`

- [ ] Upstream: `third_party/bash/builtins/break.def`

### `builtin`

- [ ] Upstream: `third_party/bash/builtins/builtin.def`
- [ ] Basic execution: Missing implementation

### `caller`

- [ ] Upstream: `third_party/bash/builtins/caller.def`

### `cat`

- [ ] Basic output: Missing implementation
- [ ] Flag `-b`: `third_party/coreutils/src/cat.c:L103`
- [ ] Flag `-n`: `third_party/coreutils/src/cat.c:L112`
- [ ] Flag `-s`: `third_party/coreutils/src/cat.c:L115`
- [ ] Flag `-v`: `third_party/coreutils/src/cat.c:L127`

### `cd`

- [ ] Upstream: `third_party/bash/builtins/cd.def`
- [ ] Basic change directory: Missing implementation
- [ ] CDPATH support: `third_party/bash/builtins/cd.def:L84`
- [ ] Flag `-L`: `third_party/bash/builtins/cd.def:L94`
- [ ] Flag `-P`: `third_party/bash/builtins/cd.def:L96`
- [ ] Flags to implement: -e

### `chcon`

- [ ] Upstream: `third_party/coreutils/src/chcon.c`
- [ ] Flags to implement: -H, -L, -P, -R, -h, -l, -r, -t, -u, -v

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
- [ ] Flags to implement: -a, -b, -c, -l, -r, -s, -t, -w, -z

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
- [ ] Basic execution: Missing implementation
- [ ] Flag `-V`: `third_party/bash/builtins/command.def:L37`
- [ ] Flag `-p`: `third_party/bash/builtins/command.def:L33`
- [ ] Flag `-v`: `third_party/bash/builtins/command.def:L35`
- [ ] Flags to implement: -pVv

### `complete`

- [ ] Upstream: `third_party/bash/builtins/complete.def`
- [ ] Flags to implement: -A, -C, -D, -DEI, -E, -F, -G, -I, -P, -S, -V, -W, -X, -abcdefgjksuv, -abcdefgjkvu, -o, -p, -pr, -r

### `cp`

- [ ] Basic copy: Missing implementation
- [ ] Flag `-a`: `third_party/coreutils/src/cp.c:L173`
- [ ] Flag `-p`: `third_party/coreutils/src/cp.c:L234`
- [ ] Flag `-r`: `third_party/coreutils/src/cp.c:L250`

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
- [ ] Flag `-b`: `third_party/coreutils/src/cut.c:L143`
- [ ] Flag `-c`: `third_party/coreutils/src/cut.c:L147`
- [ ] Flag `-d`: `third_party/coreutils/src/cut.c:L151`
- [ ] Flag `-f`: `third_party/coreutils/src/cut.c:L155`
- [ ] Flags to implement: -n, -s, -z

### `date`

- [ ] Upstream: `third_party/coreutils/src/date.c`
- [ ] Basic output: Missing implementation
- [ ] Custom format `+FORMAT`: `third_party/coreutils/src/date.c:L607`
- [ ] Flag `-d`: `third_party/coreutils/src/date.c:L501`
- [ ] Flag `-u`: `third_party/coreutils/src/date.c:L561`
- [ ] Flags to implement: -I, -R, -f, -r, -s

### `dd`

- [ ] Upstream: `third_party/coreutils/src/dd.c`
- [ ] Basic copy: Missing implementation
- [ ] Flag `bs=BYTES`: `third_party/coreutils/src/dd.c:L536`
- [ ] Flag `conv=CONVS`: `third_party/coreutils/src/dd.c:L543`
- [ ] Flag `count=N`: `third_party/coreutils/src/dd.c:L546`
- [ ] Flag `if=FILE`: `third_party/coreutils/src/dd.c:L552` (usage)
- [ ] Flag `of=FILE`: `third_party/coreutils/src/dd.c:L561`
- [ ] Flags to implement: bs, cbs, conv, count, ibs, if, iflag, obs, of, oflag, seek, skip, status

### `declare`

- [ ] Upstream: `third_party/bash/builtins/declare.def`
- [ ] Attribute management (-i, -r, -x, -a, -A): Missing implementation
- [ ] Flag `-g`: `third_party/bash/builtins/declare.def:L320`
- [ ] Flag `-p`: `third_party/bash/builtins/declare.def:L306`
- [ ] Flags to implement: -A, -F, -I, -a, -f, -i, -l, -n, -r, -t, -u, -x

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

### `dircolors`

- [ ] Upstream: `third_party/coreutils/src/dircolors.c`
- [ ] Flags to implement: -b, -c, -p

### `dirname`

- [ ] Basic operation: Missing implementation
- [ ] Flag `-z`: `third_party/coreutils/src/dirname.c:L99`

### `du`

- [ ] Upstream: `third_party/coreutils/src/du.c`
- [ ] Basic du: Missing implementation
- [ ] Basic usage summary: Missing implementation
- [ ] Flag `--exclude=PATTERN`: `third_party/coreutils/src/du.c:L398`
- [ ] Flag `--files0-from=F`: `third_party/coreutils/src/du.c:L328`
- [ ] Flag `--inodes`: `third_party/coreutils/src/du.c:L341`
- [ ] Flag `--si`: `third_party/coreutils/src/du.c:L369`
- [ ] Flag `--time-style=STYLE`: `third_party/coreutils/src/du.c:L389`
- [ ] Flag `--time[=WORD]`: `third_party/coreutils/src/du.c:L382`
- [ ] Flag `-0`: `third_party/coreutils/src/du.c:L290`
- [ ] Flag `-A`: `third_party/coreutils/src/du.c:L298`
- [ ] Flag `-B`: `third_party/coreutils/src/du.c:L305`
- [ ] Flag `-D`: `third_party/coreutils/src/du.c:L318`
- [ ] Flag `-H`: `third_party/coreutils/src/du.c:L333`
- [ ] Flag `-L`: `third_party/coreutils/src/du.c:L349`
- [ ] Flag `-P`: `third_party/coreutils/src/du.c:L361`
- [ ] Flag `-S`: `third_party/coreutils/src/du.c:L365`
- [ ] Flag `-X`: `third_party/coreutils/src/du.c:L394`
- [ ] Flag `-a`: `third_party/coreutils/src/du.c:L294`
- [ ] Flag `-b`: `third_party/coreutils/src/du.c:L310`
- [ ] Flag `-c`: `third_party/coreutils/src/du.c:L314`
- [ ] Flag `-d`: `third_party/coreutils/src/du.c:L322`
- [ ] Flag `-h`: `third_party/coreutils/src/du.c:L337`
- [ ] Flag `-k`: `third_party/coreutils/src/du.c:L345`
- [ ] Flag `-l`: `third_party/coreutils/src/du.c:L353`
- [ ] Flag `-m`: `third_party/coreutils/src/du.c:L357`
- [ ] Flag `-s`: `third_party/coreutils/src/du.c:L373`
- [ ] Flag `-t`: `third_party/coreutils/src/du.c:L377`
- [ ] Flag `-x`: `third_party/coreutils/src/du.c:L402`

### `echo`

- [ ] Basic output: Missing implementation
- [ ] Flag `-E`: `third_party/bash/builtins/echo.def:L152`
- [ ] Flag `-e`: `third_party/bash/builtins/echo.def:L149`
- [ ] Flag `-n`: `third_party/bash/builtins/echo.def:L146`

### `enable`

- [ ] Upstream: `third_party/bash/builtins/enable.def`
- [ ] Flags to implement: -a, -d, -dnps, -f, -n, -p, -s

### `env`

- [ ] Upstream: `third_party/coreutils/src/env.c`
- [ ] Basic execution: Missing implementation
- [ ] Flag `--block-signal[=SIG]`: `third_party/coreutils/src/env.c:L148`
- [ ] Flag `--default-signal[=SIG]`: `third_party/coreutils/src/env.c:L152`
- [ ] Flag `--ignore-signal[=SIG]`: `third_party/coreutils/src/env.c:L156`
- [ ] Flag `--list-signal-handling`: `third_party/coreutils/src/env.c:L160`
- [ ] Flag `-0`: `third_party/coreutils/src/env.c:L131`
- [ ] Flag `-C`: `third_party/coreutils/src/env.c:L139`
- [ ] Flag `-S`: `third_party/coreutils/src/env.c:L143`
- [ ] Flag `-a`: `third_party/coreutils/src/env.c:L123`
- [ ] Flag `-i`: `third_party/coreutils/src/env.c:L127`
- [ ] Flag `-u`: `third_party/coreutils/src/env.c:L135`
- [ ] Flag `-v`: `third_party/coreutils/src/env.c:L164`

### `eval`

- [ ] Upstream: `third_party/bash/builtins/eval.def`
- [ ] Basic execution: Missing implementation

### `exec`

- [ ] Upstream: `third_party/bash/builtins/exec.def`
- [ ] Basic execution: Missing implementation
- [ ] Flag `-a name`: `third_party/bash/builtins/exec.def:L120`
- [ ] Flag `-c`: `third_party/bash/builtins/exec.def:L114`
- [ ] Flags to implement: -a, -cl

### `exit`

- [ ] Upstream: `third_party/bash/builtins/exit.def`
- [ ] Basic exit: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/exit.def:L25`

### `expand`

- [ ] Upstream: `third_party/coreutils/src/expand.c`
- [ ] Basic conversion: Missing implementation
- [ ] Flag `-i`: `third_party/coreutils/src/expand.c:L78`
- [ ] Flag `-t`: `third_party/coreutils/src/expand.c:L82`

### `export`

- [ ] Upstream: `third_party/bash/builtins/setattr.def`
- [ ] Flags to implement: -f, -n, -p

### `expr`

- [ ] Upstream: `third_party/coreutils/src/expr.c`
- [ ] Arithmetic (+, -, *, /, %): Missing implementation
- [ ] Comparison (=, !=, <, <=, >, >=): Missing implementation
- [ ] Logical (| , &): Missing implementation
- [ ] String ops (match, substr, index, length): Missing implementation
- [ ] Flags to implement: specific

### `factor`

- [ ] Upstream: `third_party/coreutils/src/factor.c`
- [ ] Flags to implement: -h

### `fc`

- [ ] Upstream: `third_party/bash/builtins/fc.def`
- [ ] Basic editing/re-execution: Missing implementation
- [ ] Flag `-e ENAME`: `third_party/bash/builtins/fc.def:L232`
- [ ] Flag `-l`: `third_party/bash/builtins/fc.def:L220`
- [ ] Flag `-n`: `third_party/bash/builtins/fc.def:L216`
- [ ] Flag `-r`: `third_party/bash/builtins/fc.def:L224`
- [ ] Flag `-s`: `third_party/bash/builtins/fc.def:L228`
- [ ] Flags to implement: -e

### `fg`

- [ ] Upstream: `third_party/bash/builtins/fg_bg.def`

### `fmt`

- [ ] Upstream: `third_party/coreutils/src/fmt.c`
- [ ] Basic formatting: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/fmt.c:L278`
- [ ] Flag `-g`: `third_party/coreutils/src/fmt.c:L305`
- [ ] Flag `-p`: `third_party/coreutils/src/fmt.c:L282`
- [ ] Flag `-s`: `third_party/coreutils/src/fmt.c:L287`
- [ ] Flag `-t`: `third_party/coreutils/src/fmt.c:L291`
- [ ] Flag `-u`: `third_party/coreutils/src/fmt.c:L295`
- [ ] Flag `-w`: `third_party/coreutils/src/fmt.c:L299`
- [ ] Flags to implement: -WIDTH

### `fold`

- [ ] Upstream: `third_party/coreutils/src/fold.c`
- [ ] Basic folding: Missing implementation
- [ ] Flag `-b`: `third_party/coreutils/src/fold.c:L88`
- [ ] Flag `-c`: `third_party/coreutils/src/fold.c:L92`
- [ ] Flag `-s`: `third_party/coreutils/src/fold.c:L96`
- [ ] Flag `-w`: `third_party/coreutils/src/fold.c:L100`

### `getlimits`

- [ ] Upstream: `third_party/coreutils/src/getlimits.c`
- [ ] Flags to implement: specific

### `getopts`

- [ ] Basic parsing: Missing implementation
- [ ] Silent mode support (`:`): `third_party/bash/builtins/getopts.def:L180`

### `groups`

- [ ] Upstream: `third_party/coreutils/src/groups.c`
- [ ] Basic listing: Missing implementation
- [ ] Multiple users support: `third_party/coreutils/src/groups.c:L120`
- [ ] Flags to implement: specific

### `hash`

- [ ] Upstream: `third_party/bash/builtins/hash.def`
- [ ] Basic hashing: Missing implementation
- [ ] Flag `-d`: `third_party/bash/builtins/hash.def:L32`
- [ ] Flag `-l`: `third_party/bash/builtins/hash.def:L33`
- [ ] Flag `-p pathname`: `third_party/bash/builtins/hash.def:L34`
- [ ] Flag `-r`: `third_party/bash/builtins/hash.def:L35`
- [ ] Flag `-t`: `third_party/bash/builtins/hash.def:L36`
- [ ] Flags to implement: -dt, -lr, -p

### `head`

- [ ] Basic output: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/head.c:L121`
- [ ] Flag `-n`: `third_party/coreutils/src/head.c:L126`
- [ ] Flag `-q`: `third_party/coreutils/src/head.c:L131`

### `help`

- [ ] Upstream: `third_party/bash/builtins/help.def`
- [ ] Basic discovery: Missing implementation
- [ ] Flag `-d`: `third_party/bash/builtins/help.def:L105`
- [ ] Flag `-m`: `third_party/bash/builtins/help.def:L108`
- [ ] Flag `-s`: `third_party/bash/builtins/help.def:L111`
- [ ] Flags to implement: -dms

### `history`

- [ ] Upstream: `third_party/bash/builtins/history.def`
- [ ] Basic management: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/history.def:L126`
- [ ] Flag `-c`: `third_party/bash/builtins/history.def:L129`
- [ ] Flag `-d offset`: `third_party/bash/builtins/history.def:L145`
- [ ] Flag `-n`: `third_party/bash/builtins/history.def:L132`
- [ ] Flag `-p`: `third_party/bash/builtins/history.def:L148`
- [ ] Flag `-r`: `third_party/bash/builtins/history.def:L135`
- [ ] Flag `-s`: `third_party/bash/builtins/history.def:L141`
- [ ] Flag `-w`: `third_party/bash/builtins/history.def:L138`
- [ ] Flags to implement: -d, offset

### `hostid`

- [ ] Upstream: `third_party/coreutils/src/hostid.c`
- [ ] Flags to implement: specific

### `hostname`

- [ ] Upstream: `third_party/coreutils/src/hostname.c`
- [ ] Basic output: Missing implementation
- [ ] Set hostname support: `third_party/coreutils/src/hostname.c:L95`
- [ ] Flags to implement: specific

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
- [ ] Flags to implement: -C, -D, -S, -T, -Z, -b, -c, -d, -g, -m, -o, -p, -s, -t, -v

### `jobs`

- [ ] Upstream: `third_party/bash/builtins/jobs.def`
- [ ] Flags to implement: -a, -h, -l, -lnprs, -n, -p, -r, -s, -x

### `join`

- [ ] Upstream: `third_party/coreutils/src/join.c`
- [ ] Basic join: Missing implementation
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
- [ ] Flags to implement: -1, -2, -a, -e, -j, -o, -t, -v, CHAR, FIELD, FILENUM, FORMAT, STRING

### `kill`

- [ ] Upstream: `third_party/bash/builtins/kill.def`
- [ ] Basic signaling: Missing implementation
- [ ] Flag `-l`: `third_party/coreutils/src/kill.c:L277` / `third_party/bash/builtins/kill.def:L114`
- [ ] Flag `-s SIGNAL`: `third_party/coreutils/src/kill.c:L262` / `third_party/bash/builtins/kill.def:L129`
- [ ] Flags to implement: -n, -s

### `let`

- [ ] Upstream: `third_party/bash/builtins/let.def`
- [ ] Flags to implement: specific

### `link`

- [ ] Basic hard link: Missing implementation (exactly 2 args required)

### `ln`

- [ ] Basic link creation: Missing implementation
- [ ] Flag `-f`: `third_party/coreutils/src/ln.c:L553`
- [ ] Flag `-s`: `third_party/coreutils/src/ln.c:L574`
- [ ] Flag `-v`: `third_party/coreutils/src/ln.c:L595`

### `local`

- [ ] Upstream: `third_party/bash/builtins/declare.def`
- [ ] Flags to implement: as, declare`)

### `logname`

- [ ] Upstream: `third_party/coreutils/src/logname.c`
- [ ] Flags to implement: specific

### `ls`

- [ ] Basic listing: Missing implementation
- [ ] Color output (`--color`): `third_party/coreutils/src/ls.c`
- [ ] Flag `-R`: `third_party/coreutils/src/ls.c`
- [ ] Flag `-a`: `third_party/coreutils/src/ls.c:L41`
- [ ] Flag `-h`: `third_party/coreutils/src/ls.c:L47`
- [ ] Flag `-l`: `third_party/coreutils/src/ls.c`

### `mapfile`

- [ ] Upstream: `third_party/bash/builtins/mapfile.def`
- [ ] Flags to implement: -C, -O, -c, -d, -n, -s, -t, -u

### `md5sum`

- [ ] Upstream: `third_party/coreutils/src/cksum.c`
- [ ] Flags to implement: as, cksum`)

### `mkdir`

- [ ] Basic creation: Missing implementation
- [ ] Flag `-m`: `third_party/coreutils/src/mkdir.c:L65`
- [ ] Flag `-p`: `third_party/coreutils/src/mkdir.c:L69`
- [ ] Flag `-v`: `third_party/coreutils/src/mkdir.c:L74`

### `mkfifo`

- [ ] Upstream: `third_party/coreutils/src/mkfifo.c`
- [ ] Flags to implement: -Z, -m

### `mknod`

- [ ] Upstream: `third_party/coreutils/src/mknod.c`
- [ ] Flags to implement: -Z, -m

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
- [ ] Flag `-f`: `third_party/coreutils/src/mv.c:L282`
- [ ] Flag `-i`: `third_party/coreutils/src/mv.c:L286`
- [ ] Flag `-n`: `third_party/coreutils/src/mv.c:L290`

### `nice`

- [ ] Upstream: `third_party/coreutils/src/nice.c`
- [ ] Flags to implement: -N, -n

### `nl`

- [ ] Upstream: `third_party/coreutils/src/nl.c`
- [ ] Basic numbering: Missing implementation
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
- [ ] Flags to implement: CC, FORMAT, NUMBER, STRING, STYLE

### `nohup`

- [ ] Upstream: `third_party/coreutils/src/nohup.c`
- [ ] Flags to implement: specific

### `nproc`

- [ ] Upstream: `third_party/coreutils/src/nproc.c`

### `numfmt`

- [ ] Upstream: `third_party/coreutils/src/numfmt.c`
- [ ] Flags to implement: -d, -z

### `od`

- [ ] Upstream: `third_party/coreutils/src/od.c`
- [ ] Flags to implement: -A, -N, -S, -a, -b, -c, -d, -f, -i, -j, -l, -o, -s, -t, -v, -w, -x

### `paste`

- [ ] Upstream: `third_party/coreutils/src/paste.c`
- [ ] Basic paste: Missing implementation
- [ ] Flag `-d`: `third_party/coreutils/src/paste.c:L468`
- [ ] Flag `-s`: `third_party/coreutils/src/paste.c:L473`
- [ ] Flag `-z`: `third_party/coreutils/src/paste.c:L479`

### `pathchk`

- [ ] Upstream: `third_party/coreutils/src/pathchk.c`
- [ ] Flags to implement: -P, -p

### `pinky`

- [ ] Upstream: `third_party/coreutils/src/pinky.c`
- [ ] Flags to implement: -b, -f, -h, -i, -l, -p, -q, -s, -w

### `pr`

- [ ] Upstream: `third_party/coreutils/src/pr.c`
- [ ] Flags to implement: -COLS, -D, -F, -J, -N, -S, -T, -W, -a, -c, -d, -e, -f, -h, -i, -l, -m, -n, -o, -r, -s, -t, -v, -w

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
- [ ] Flags to implement: -A, -F, -G, -M, -O, -R, -S, -T, -W, -b, -f, -g, -i, -o, -r, -t, -w

### `pushd`

- [ ] Upstream: `third_party/bash/builtins/pushd.def`
- [ ] Flags to implement: -c, -clpv, -l, -n, -p, -v

### `pwd`

- [ ] Upstream: `third_party/bash/builtins/cd.def`
- [ ] Basic path reporting: Missing implementation
- [-] Flag `--help`: Handled by the shell's global help dispatcher.
- [ ] Flag `-L`: `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L291-316`
- [ ] Flag `-P`: `third_party/bash/builtins/cd.def:L435-513` / `third_party/coreutils/src/pwd.c:L371-383`

### `read`

- [ ] Upstream: `third_party/bash/builtins/read.def`
- [ ] Basic input: Missing implementation
- [ ] Flag `-a`: `third_party/bash/builtins/read.def:L39`
- [ ] Flag `-p`: `third_party/bash/builtins/read.def:L53`
- [ ] Flag `-r`: `third_party/bash/builtins/read.def:L55`
- [ ] Flags to implement: -N, -d, -ers, -i, -n, -t, -u

### `readlink`

- [ ] Upstream: `third_party/coreutils/src/readlink.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-f`: `third_party/coreutils/src/readlink.c:L135`
- [ ] Flag `-n`: `third_party/coreutils/src/readlink.c:L141`
- [ ] Flags to implement: -e, -m, -q, -s, -v, -z

### `readonly`

- [ ] Upstream: `third_party/bash/builtins/setattr.def`
- [ ] Flags to implement: -A, -a, -f, -p

### `realpath`

- [ ] Upstream: `third_party/coreutils/src/realpath.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `--relative-to`: `third_party/coreutils/src/realpath.c:L246`
- [ ] Flag `-e`: `third_party/coreutils/src/realpath.c:L220`
- [ ] Flag `-m`: `third_party/coreutils/src/realpath.c:L224`
- [ ] Flags to implement: -E, -L, -P, -q, -s, -z

### `return`

- [ ] Upstream: `third_party/bash/builtins/return.def`
- [ ] Basic return: Missing implementation
- [ ] Exit status parameter: `third_party/bash/builtins/return.def:L61`

### `rm`

- [ ] Upstream: `third_party/coreutils/src/rm.c`
- [ ] Basic removal: Missing implementation
- [ ] Flag `-f`: `third_party/coreutils/src/rm.c:L137`
- [ ] Flag `-i`: `third_party/coreutils/src/rm.c:L142`
- [ ] Flag `-r`: `third_party/coreutils/src/rm.c:L172`
- [ ] Flags to implement: -I, -d, -v

### `rmdir`

- [ ] Upstream: `third_party/coreutils/src/rmdir.c`
- [ ] Basic rmdir: Missing implementation
- [ ] Basic removal: Missing implementation
- [ ] Flag `--ignore-fail-on-non-empty`: `third_party/coreutils/src/rmdir.c:L178`
- [ ] Flag `-p`: `third_party/coreutils/src/rmdir.c:L182`
- [ ] Flag `-v`: `third_party/coreutils/src/rmdir.c:L187`

### `runcon`

- [ ] Upstream: `third_party/coreutils/src/runcon.c`
- [ ] Flags to implement: -c, -l, -r, -t, -u

### `seq`

- [ ] Upstream: `third_party/coreutils/src/seq.c`
- [ ] Basic sequence: Missing implementation
- [ ] Flag `-s`: `third_party/coreutils/src/seq.c:L596`
- [ ] Flag `-w`: `third_party/coreutils/src/seq.c:L600`
- [ ] Flags to implement: -f

### `set`

- [ ] Upstream: `third_party/bash/builtins/set.def`
- [ ] Option management (-e, -u, -x, -o): Missing implementation
- [ ] Positional parameters: `third_party/bash/builtins/set.def:L784`
- [ ] Flag `-f`: `third_party/bash/builtins/set.def:L799` (unset)
- [ ] Flags to implement: -abefhkmnptuvxBCEHPT, -n, -o, -v

### `shift`

- [ ] Upstream: `third_party/bash/builtins/shift.def`
- [ ] Basic shift: Missing implementation
- [ ] Shifting n parameters: `third_party/bash/builtins/shift.def:L64`

### `shopt`

- [ ] Upstream: `third_party/bash/builtins/shopt.def`
- [ ] Basic option management: Missing implementation
- [ ] Flag `-o`: `third_party/bash/builtins/shopt.def:L316`
- [ ] Flag `-p`: `third_party/bash/builtins/shopt.def:L319`
- [ ] Flag `-q`: `third_party/bash/builtins/shopt.def:L313`
- [ ] Flag `-s`: `third_party/bash/builtins/shopt.def:L307`
- [ ] Flag `-u`: `third_party/bash/builtins/shopt.def:L310`

### `shred`

- [ ] Upstream: `third_party/coreutils/src/shred.c`
- [ ] Flags to implement: -f, -n, -s, -u, -v, -x, -z

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
- [ ] Flags to implement: specific

### `sort`

- [ ] Upstream: `third_party/coreutils/src/sort.c`
- [ ] Basic sorting: Missing implementation
- [ ] Basic sort: Missing implementation
- [ ] Ordering flags (`-b`, `-i`, `-d`, `-f`, `-g`, `-h`, `-n`, `-M`, `-R`, `-V`, `-r`): `third_party/coreutils/src/sort.c:L437-490`
- [ ] Flag `--parallel`: `third_party/coreutils/src/sort.c:L553`
- [ ] Flag `-S`: `third_party/coreutils/src/sort.c:L540`
- [ ] Flag `-T`: `third_party/coreutils/src/sort.c:L548`
- [ ] Flag `-c`: `third_party/coreutils/src/sort.c:L501-507`
- [ ] Flag `-k`: `third_party/coreutils/src/sort.c:L524`
- [ ] Flag `-m`: `third_party/coreutils/src/sort.c:L528`
- [ ] Flag `-n`: `third_party/coreutils/src/sort.c:L464`
- [ ] Flag `-o`: `third_party/coreutils/src/sort.c:L532`
- [ ] Flag `-r`: `third_party/coreutils/src/sort.c:L478`
- [ ] Flag `-t`: `third_party/coreutils/src/sort.c:L544`
- [ ] Flag `-u`: `third_party/coreutils/src/sort.c:L557`
- [ ] Flag `-z`: `third_party/coreutils/src/sort.c:L562`
- [ ] Flags to implement: -C, -M, -R, -V, -b, -d, -f, -g, -h, -i, -s

### `source`

- [ ] Upstream: `third_party/bash/builtins/source.def`
- [ ] Basic sourcing: Missing implementation
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
- [ ] Flag `-L`: `third_party/coreutils/src/stat.c:L1927` (main)
- [ ] Flag `-c`: `third_party/coreutils/src/stat.c:L1921`
- [ ] Flag `-f`: `third_party/coreutils/src/stat.c:L1932`
- [ ] Flag `-t`: `third_party/coreutils/src/stat.c:L1936`

### `stdbuf`

- [ ] Upstream: `third_party/coreutils/src/stdbuf.c`
- [ ] Flags to implement: -e, -i, -o

### `stty`

- [ ] Upstream: `third_party/coreutils/src/stty.c`
- [ ] Flags to implement: -F, -a, -g

### `sum`

- [ ] Upstream: `third_party/coreutils/src/sum.c`
- [ ] Flags to implement: -r, -s

### `suspend`

- [ ] Upstream: `third_party/bash/builtins/suspend.def`
- [ ] Flags to implement: -f

### `sync`

- [ ] Upstream: `third_party/coreutils/src/sync.c`
- [ ] Flags to implement: -d, -f

### `tac`

- [ ] Upstream: `third_party/coreutils/src/tac.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-b`: `third_party/coreutils/src/tac.c:L103`
- [ ] Flag `-r`: `third_party/coreutils/src/tac.c:L104`
- [ ] Flag `-s`: `third_party/coreutils/src/tac.c:L105`

### `tail`

- [ ] Basic output: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/tail.c:L296`
- [ ] Flag `-f`: `third_party/coreutils/src/tail.c:L305`
- [ ] Flag `-n`: `third_party/coreutils/src/tail.c:L314`

### `tee`

- [ ] Basic copy: Missing implementation
- [ ] Flag `-a`: `third_party/coreutils/src/tee.c:L93`
- [ ] Flag `-i`: `third_party/coreutils/src/tee.c:L97`

### `test`

- [ ] Unary operators (-e, -f, -d, etc.): Missing implementation
- [ ] String operators (=, !=, -z, -n): Missing implementation
- [ ] Numeric operators (-eq, -ne, etc.): Missing implementation
- [ ] Logical operators (!, -a, -o): Missing implementation

### `timeout`

- [ ] Upstream: `third_party/coreutils/src/timeout.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-k`: `third_party/coreutils/src/timeout.c:L531`
- [ ] Flag `-s`: `third_party/coreutils/src/timeout.c:L539`
- [ ] Flags to implement: -f, -p, -v

### `times`

- [ ] Upstream: `third_party/bash/builtins/times.def`
- [ ] Basic output: Missing implementation

### `touch`

- [ ] Upstream: `third_party/coreutils/src/touch.c`
- [ ] Basic touch: Missing implementation
- [ ] Basic timestamp update: Missing implementation
- [ ] Flag `-a`: `third_party/coreutils/src/touch.c:L230`
- [ ] Flag `-c`: `third_party/coreutils/src/touch.c:L234`
- [ ] Flag `-d`: `third_party/coreutils/src/touch.c:L238`
- [ ] Flag `-h`: `third_party/coreutils/src/touch.c:L246`
- [ ] Flag `-m`: `third_party/coreutils/src/touch.c:L251`
- [ ] Flag `-r`: `third_party/coreutils/src/touch.c:L255`
- [ ] Flag `-t [[CC]YY]MMDDhhmm[.ss]`: `third_party/coreutils/src/touch.c:L259`
- [ ] Flags to implement: -f, -t

### `tr`

- [ ] Upstream: `third_party/coreutils/src/tr.c`
- [ ] Basic translation: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/tr.c:L296`
- [ ] Flag `-d`: `third_party/coreutils/src/tr.c:L300`
- [ ] Flag `-s`: `third_party/coreutils/src/tr.c:L304`
- [ ] Flag `-t`: `third_party/coreutils/src/tr.c:L310`
- [ ] Flags to implement: -C

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
- [ ] Flags to implement: -w, a, but, it's, no-op), specific

### `tty`

- [ ] Upstream: `third_party/coreutils/src/tty.c`
- [ ] Basic output: Missing implementation
- [ ] Flag `-s`: `third_party/coreutils/src/tty.c:L71`

### `type`

- [ ] Upstream: `third_party/bash/builtins/type.def`
- [ ] Basic lookup: Missing implementation
- [ ] Basic identification: Missing implementation
- [ ] Flag `-P`: `third_party/bash/builtins/type.def:L36`
- [ ] Flag `-a`: `third_party/bash/builtins/type.def:L32`
- [ ] Flag `-f`: `third_party/bash/builtins/type.def:L35`
- [ ] Flag `-p`: `third_party/bash/builtins/type.def:L39`
- [ ] Flag `-t`: `third_party/bash/builtins/type.def:L41`

### `ulimit`

- [ ] Upstream: `third_party/bash/builtins/ulimit.def`
- [ ] Flags to implement: -H, -P, -R, -S, -SHabcdefiklmnpqrstuvxPRT, -T, -a, -b, -c, -d, -e, -f, -i, -k, -l, -m, -n, -p, -q, -r, -s, -t, -u, -v, -x

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
- [ ] Flag `--first-only`: `third_party/coreutils/src/unexpand.c:L91`
- [ ] Flag `-a`: `third_party/coreutils/src/unexpand.c:L87`
- [ ] Flag `-t`: `third_party/coreutils/src/unexpand.c:L95`

### `uniq`

- [ ] Upstream: `third_party/coreutils/src/uniq.c`
- [ ] Basic filtering: Missing implementation
- [ ] Flag `-D`: `third_party/coreutils/src/uniq.c:180`
- [ ] Flag `-c`: `third_party/coreutils/src/uniq.c:172`
- [ ] Flag `-d`: `third_party/coreutils/src/uniq.c:176`
- [ ] Flag `-f`: `third_party/coreutils/src/uniq.c:189`
- [ ] Flag `-i`: `third_party/coreutils/src/uniq.c:198`
- [ ] Flag `-s`: `third_party/coreutils/src/uniq.c:202`
- [ ] Flag `-u`: `third_party/coreutils/src/uniq.c:206`
- [ ] Flag `-w`: `third_party/coreutils/src/uniq.c:214`
- [ ] Flag `-z`: `third_party/coreutils/src/uniq.c:210`

### `unlink`

- [ ] Upstream: `third_party/coreutils/src/unlink.c`
- [ ] Basic removal: Missing implementation (exactly 1 arg required)
- [ ] Flags to implement: specific

### `uptime`

- [ ] Upstream: `third_party/coreutils/src/uptime.c`
- [ ] Flags to implement: -p, -s, specific

### `users`

- [ ] Upstream: `third_party/coreutils/src/users.c`
- [ ] Flags to implement: beyond, specific, standard

### `vdir`

- [ ] Upstream: `third_party/coreutils/src/ls.c`

### `wait`

- [ ] Upstream: `third_party/bash/builtins/wait.def`
- [ ] Basic waiting: Missing implementation
- [ ] Flag `-f`: `third_party/bash/builtins/wait.def:L134`
- [ ] Flag `-n`: `third_party/bash/builtins/wait.def:L131`
- [ ] Flag `-p var`: `third_party/bash/builtins/wait.def:L137`
- [ ] Flags to implement: -p, var

### `wc`

- [ ] Basic counts: Missing implementation
- [ ] Flag `-c`: `third_party/coreutils/src/wc.c:L188`
- [ ] Flag `-l`: `third_party/coreutils/src/wc.c:L196`
- [ ] Flag `-m`: `third_party/coreutils/src/wc.c:L192`
- [ ] Flag `-w`: `third_party/coreutils/src/wc.c:L214`

### `who`

- [ ] Upstream: `third_party/coreutils/src/who.c`
- [ ] Flags to implement: -H, -T, -a, -b, -d, -l, -m, -p, -q, -r, -s, -t, -u, -w, -y

### `whoami`

- [ ] Upstream: `third_party/coreutils/src/whoami.c`
- [ ] Basic output: Missing implementation
- [ ] Flags to implement: specific

### `yes`

- [ ] Upstream: `third_party/coreutils/src/yes.c`
- [ ] Basic operation: Missing implementation
- [ ] Basic repetition: Missing implementation
- [ ] Flags to implement: specific

