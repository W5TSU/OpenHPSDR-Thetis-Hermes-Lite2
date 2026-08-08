---
name: thetis-control
description: Remotely control a running Thetis SDR instance (frequency, mode, filters, VFO, RIT/XIT, split, AGC, attenuator/preamp, band — Tier 1; RX/TX audio and actually keying the transmitter — Tier 2) over its existing CAT-over-TCP and TCI-over-WebSocket network protocols, via the thetisctl Go CLI. Use when asked to "control Thetis", "set frequency on Thetis", "change mode on the radio", "key the HL2 transmitter", "stream TCI audio", "transmit over Thetis", or similar.
---

# Controlling Thetis remotely

Thetis (this repo) already exposes two network protocols that let a client
drive a **separate, already-running** Thetis instance — nothing in Thetis
itself needs to change. This skill wraps both in `thetisctl`, a Go CLI at
`Tools/thetis-ai-control/`.

Facts verified 2026-07-23 against source — recheck if stale:

- **CAT-over-TCP** (`Project Files/Source/Console/CAT/TCPIPcatServer.cs`,
  default `127.0.0.1:13013`) — Kenwood-style ASCII protocol. Covers Tier 1
  control state plus bare `TX;`/`RX;` PTT.
- **TCI-over-WebSocket** (`Project Files/Source/Console/TCIServer.cs`,
  default `127.0.0.1:50001`) — ExpertSDR3-compatible. Text frames mirror the
  same controls as CAT; binary frames carry RX/TX audio. Sending
  TX_AUDIO_STREAM frames while PTT is active is drained by
  `Console/cmaster.cs`'s `TCITxThreadProc` straight into the DSP TX chain —
  **this genuinely keys and modulates the transmitter**, not a simulation.
- **Neither protocol has authentication.** Any TCP client that can reach the
  bound address can issue any command, including PTT. Only bind these
  servers to a trusted LAN address — never expose them to the internet.

## Prerequisites

In the Thetis instance you want to control, open **Setup** and enable:

- **CAT server**: the TCP/IP CAT Server checkbox (`chkTCPIPCatServerListening`
  in code). Default bind is `127.0.0.1:13013` — if `thetisctl` runs on a
  different machine, rebind to the host's LAN IP (e.g. `192.168.1.50:13013`),
  not `127.0.0.1`.
- **TCI server**: the TCI Server checkbox (`chkTCIServerListening`), default
  bind `127.0.0.1:50001` — same rebind caveat.

Confirm you can reach the chosen address from wherever `thetisctl` will run
(same LAN, no firewall blocking the port) before troubleshooting the tool
itself.

## Build

```bash
cd Tools/thetis-ai-control
go build -o thetisctl ./cmd/thetisctl
go vet ./...
go test ./...
```

Pure Go, no cgo, no external dependencies.

## Tier 1 — control-only usage (CAT)

| Command | Effect |
|---|---|
| `cat --host <ip> freq get A\|B` / `freq set A\|B <hz>` | VFO A/B frequency |
| `cat --host <ip> mode get` / `mode set <USB\|LSB\|CW\|CWL\|FM\|AM\|DIGU\|DIGL>` | Demod mode |
| `cat --host <ip> rit on\|off\|get` / `xit on\|off\|get` | RIT/XIT enable |
| `cat --host <ip> split on\|off\|get` | VFO split |
| `cat --host <ip> agc get` / `agc set <FIXED\|LONG\|SLOW\|MEDIUM\|FAST\|CUSTOM>` | AGC mode |
| `cat --host <ip> atten get` / `atten set <0-31>` | RX1 step attenuator (dB) |
| `cat --host <ip> preamp set <0-9>` | RX1 preamp level |
| `cat --host <ip> band get` / `band set <name>` | Band (160-2, GEN, WWV, V0-V13) |
| `cat --host <ip> power get` / `power on\|off` | Start/stop Thetis's radio engine (software power, not mains) |
| `cat --host <ip> quickplay get` / `quickplay off` | Quick Play: read state / stop (always safe) |
| `cat --host <ip> quickplay on` | **TX-capable** — see the safety protocol below before using it |
| `cat --host <ip> quickrec get` / `quickrec on\|off` | Quick Rec: record RX audio to that same fixed file |
| `cat --host <ip> freedv get` / `freedv on\|off` | Enable/disable FreeDV RX decode (RX1 only) |
| `cat --host <ip> tciserver get` / `tciserver on\|off` | Enable/disable Thetis's TCI server — works even when TCI itself is unreachable (CAT doesn't depend on it), so it can bootstrap TCI back on after a restart left the checkbox unchecked |
| `cat --host <ip> freedv status` | Read FreeDV sync + SNR (read-only) |
| `cat --host <ip> status` | Combined ID + frequency/mode/RIT/XIT/split/TX status |
| `cat --host <ip> version` | Software version string, including the git short SHA the running build was made from (`ZZZV`) — check this before trusting that a remote instance is running the build you expect |
| `cat --host <ip> query <CODE>` | Raw passthrough for any CAT command not wrapped above (e.g. `query ZZZV`) |

```bash
./thetisctl cat --host 192.168.1.50 freq set A 14074000
./thetisctl cat --host 192.168.1.50 mode set USB
./thetisctl cat --host 192.168.1.50 status
```

`ptt on|off` also exists on the CAT channel but is TX-capable — see the
safety protocol below before using it.

`power on|off` starts/stops Thetis's *software* radio engine (the main Power
button — `console.PowerOn`) — it does NOT toggle mains/PoE power to the
physical HL2 board; if the board itself has no power at all, this cannot
bring it up. It is not TX-capable (it can't key the transmitter), but it is a
bigger action than the other Tier 1 controls — it starts/stops the actual
hardware connection and DSP audio engine — so treat it with the same "ask
before doing it" judgment as any other state-changing remote action, even
though it doesn't need `--confirm-tx`. Powering on can take a few seconds;
pass a longer `--timeout` if `power on`'s readback confirmation times out.

`quickrec on|off` and `freedv on|off|status` are not TX-capable — Quick Rec
records RX audio to a fixed file, and `freedv` controls/reads Thetis's
FreeDV RX decode block (`fdv.c`); neither touches the transmitter
(`RecordToFileFromWDSP` was checked and confirmed to never touch
`_console.MOX`).

**`quickplay on` is TX-capable — do not treat it as safe.** It was
originally documented here (and believed, based on that documentation) to
be RX-only: it injects a fixed audio file as RX I/Q *before* the antenna
input. In practice it calls Thetis's `PlayFileViaWDSP`
(`Console/clsAudioRecordPlayback.cs`), a function *shared* with a genuine
TX-audio-preview feature, which contains
`if (!_console.MOX && MoxOnPlayback) _console.MOX = true;` — and
`MoxOnPlayback` **defaults to `true`** in this codebase (Setup →
Recording's "MOX on Playback" checkbox). This was discovered by live
testing 2026-08-04 *after* several sessions of calling `quickplay on`
believing it was RF-free — it went through `catToggle` exactly like
`quickrec`, with no `--confirm-tx` gate at all. It's since been fixed in
`thetisctl` (`cat_cmd.go`'s `catQuickPlay`) to require `--confirm-tx` and
auto-stop after `--hold` (default 15s), same as `ptt`/`tune`. Follow the
full safety protocol below before using it — same as `ptt`. If you need it
to be genuinely RX-only, ask the operator to confirm "MOX on Playback" is
unchecked in Thetis's Setup → Recording tab first; `thetisctl` cannot read
or change that setting remotely, only the resulting MOX state:

```bash
./thetisctl cat --host 192.168.1.50 freedv on
./thetisctl cat --host 192.168.1.50 quickplay on
[dry-run] would send: quickplay on
WARNING: this may key MOX for real. ...
Pass --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO to proceed.

# after the operator explicitly confirms, in this conversation, for this
# specific test:
./thetisctl cat --host 192.168.1.50 quickplay on \
    --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO --hold 15s &
sleep 3                                             # let the decoder attempt sync
./thetisctl cat --host 192.168.1.50 freedv status   # "SYNC  SNR 15.3 dB" or "no sync" — objective, no listening required
wait                                                # quickplay auto-stops after --hold and confirms it
```

`tci rx-audio capture` (below) can still be used alongside this to actually
listen to what came out — useful for judging decode *quality*, not just
sync — but `freedv status` alone is enough to script a pass/fail loop when
iterating on `fdv.c`. If Quick Play's target file doesn't exist or playback
otherwise fails, Thetis shows a local error dialog and reverts the
checkbox — `quickplay on` itself won't report an error for that (the CAT
set is fire-and-forget); use `quickplay get` afterward to confirm it's
actually still running.

## Tier 2 — audio and transmit usage (TCI)

`rx` selects the receiver: `0` = RX1, `1` = RX2.

| Command | Effect |
|---|---|
| `tci --host <ip> vfo <rx> <chan 0\|1> <hz>` | VFO A(0)/B(1) frequency |
| `tci --host <ip> modulation <rx> <lsb\|usb\|dsb\|am\|sam\|fm\|cw\|cwl\|cwu\|digl\|digu>` | Demod mode |
| `tci --host <ip> split\|rit\|xit <rx> on\|off` | Split/RIT/XIT enable |
| `tci --host <ip> rit-offset\|xit-offset <rx> <hz>` | RIT/XIT offset |
| `tci --host <ip> filter <rx> <lowHz> <highHz>` | RX filter passband |
| `tci --host <ip> atten <rx> <dB>` / `preamp <rx> <dB<=0>` | Step attenuator / preamp gain |
| `tci --host <ip> agc <rx> <mode>` / `agc-gain <rx> <-20..120>` | AGC |
| `tci --host <ip> drive <rx> <0-100>` | TX drive power |
| `tci --host <ip> power on\|off` | Start/stop Thetis's radio engine (software power, not mains); waits for confirmation |
| `tci --host <ip> cw send <rx> "<text>" --speed <wpm> --mode <cw\|cwu\|cwl>` | Key CW text via Thetis's own macro keyer |
| `tci --host <ip> freedv-scan [--dwell 6s] [--out-dir <dir>]` | RX-only: tune RX1 through the FreeDV calling frequencies, record + report peak/RMS per band |
| `tci --host <ip> query <cmd> [args...]` | Raw passthrough for anything not listed above |

```bash
./thetisctl tci --host 192.168.1.50 vfo 0 0 14074000
./thetisctl tci --host 192.168.1.50 modulation 0 usb
./thetisctl tci --host 192.168.1.50 rx-audio capture 0 --duration 5s --out rx.wav
```

RX audio capture:

```bash
./thetisctl tci --host 192.168.1.50 rx-audio capture 0 --duration 10s --out capture.wav
./thetisctl tci --host 192.168.1.50 rx-audio stream 0 --duration 10s > raw.pcm   # float32 LE PCM
```

FreeDV calling-frequency scan — RX-only, tunes RX1 through each frequency in
`freeDVCallingFrequencies` (`cmd/thetisctl/tci_cmd.go`; the table is also
saved to project memory, `freedv-calling-frequencies`), records a WAV per
band, and restores the original frequency/mode when done:

```bash
./thetisctl tci --host 192.168.1.50 freedv-scan --out-dir /tmp/freedv-scan
```

**This does not identify FreeDV.** Peak/RMS is only a prioritization hint
for which captures are worth listening to — telling a real FreeDV signal
apart from a mistuned SSB voice transmission or plain band noise from
spectral shape alone was tried and found unreliable in practice (a captured
signal that looked "flat and broadband, no speech pauses" turned out on
listening to be a mistuned voice transmission, not FreeDV). Report the
files and their peak/RMS back to the user; let them (or real FreeDV
software) make the actual identification.

TX audio — **always run without `--confirm-tx` first**:

```bash
$ ./thetisctl tci --host 192.168.1.50 tx-audio send 0 --file cq.wav
[dry-run] would send: trx:0,true,tci; then stream 4.2s of TX audio from cq.wav
(48000 Hz, 1 ch, peak 0.812) as int16 frames; then trx:0,false,tci;
Pass --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO to actually transmit this audio.

$ ./thetisctl tci --host 192.168.1.50 tx-audio send 0 --file cq.wav \
    --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO --max-duration 10s
```

`tci tune <rx> on|off` and `tci ptt <rx> on|off` are also TX-capable (the
latter with `--audio` to declare this connection as the TX audio source) —
same safety protocol applies. `tune` is a bare, unmodulated carrier — the
highest-nuisance thing this tool can transmit — so it's hard-capped at 5
seconds total on-time regardless of `--hold`, non-configurable.

CW — Thetis's own macro keyer sends the text and manages PTT itself, so
`--confirm-tx` here authorizes the whole message, not per-character:

```bash
$ ./thetisctl tci --host 192.168.1.50 cw send 0 "CQ CQ DE W5TSU" --speed 5 --mode cwu
[dry-run] would send: modulation:0,cwu; cw_macros_speed:5; cw_macros:0,CQ CQ DE W5TSU;
Pass --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO to actually transmit this message.

$ ./thetisctl tci --host 192.168.1.50 cw send 0 "CQ CQ DE W5TSU" --speed 5 --mode cwu \
    --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO --max-duration 60s
```

`cw send` sets the target receiver's mode explicitly (`--mode`, default `cw`
which lets Thetis auto-pick CWL/CWU) before keying — don't assume the
receiver is already in a CW mode.

## Safety protocol — read before any TX-capable command

`cat ptt`, `cat quickplay on`, `tci tune`, `tci ptt`, `tci cw send`, and `tci
tx-audio send` can key the transmitter. Since neither network protocol has
authentication, `thetisctl` enforces:

1. **Dry-run by default.** Without `--confirm-tx`, these commands print
   exactly what they would send and do nothing TX-capable. Always run the
   dry-run first and show the output to the user.
2. **Never pass `--confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO` without
   asking the human operator in the *current* conversation and getting an
   explicit go-ahead.** A prior general "ok to test the radio" earlier in the
   session is not sufficient — ask again for the specific frequency, mode,
   content, and duration each time.
3. Prefer the smallest `--max-duration`/`--hold` and lowest drive that
   accomplishes the task. `tx-audio send` always hard-caps at
   `--max-duration` (default 10s) and auto-unkeys on completion, error, or
   Ctrl-C.
4. If anything about frequency, band, mode, or power is ambiguous or could
   put the transmission out of the amateur band or outside what the operator
   is licensed for, stop and ask rather than guessing.
5. **Check `split` before assuming TX frequency matches VFO A.** If split is
   enabled, the radio transmits on VFO B's frequency, not VFO A's — even
   though CAT `status`/`freq get` and everything you've been setting on
   VFO A stays exactly where you left it. This produced a real incident: a
   test session transmitted on a completely different frequency/band than
   intended, undetected until the operator noticed, because split had been
   on since before the session started and nothing in routine status output
   flagged the mismatch. Cross-check TCI's `tx_frequency` (or `query
   tx_frequency`) against VFO A before any real transmission if split status
   is unknown or was set by someone/something else.
6. **Never trust a fire-and-forget unkey.** Sending an unkey command and
   closing the connection immediately afterward can silently drop it — see
   the gotcha below. Every TX-capable command's unkey path (`tune off`,
   `ptt off`, auto-unkey-after-`--hold`, `tx-audio`'s completion/interrupt
   unkey, `cw send`'s stop-on-error) verifies the radio actually unkeyed
   before returning, retrying if not yet confirmed — any new TX-capable
   command must do the same (`confirmTCIUnkeyed`/`confirmCATUnkeyed` in
   `cmd/thetisctl/{tci,cat}_cmd.go`), never a bare fire-and-forget send.

## Extending the command set

Wire formats for the commands above were confirmed directly against source,
not just the protocol's own advertised spec — verify against the same files
before adding more:

- `Project Files/Source/Console/CAT/CATStructs.xml` — canonical CAT
  parameter-width table (note: several classic Kenwood codes like `RT`, `XT`,
  `RA`, `PA` are disabled stubs in this codebase; the working equivalents are
  `ZZRT`, `ZZXS`, `ZZRX`, `ZZPA` — check `CATParser.cs`/`CATCommands.cs`
  dispatch before trusting the XML alone).
- `Project Files/Source/Console/TCIServer.cs` — TCI command handlers (search
  `handle<Name>Message`) and the binary stream frame builder
  (`buildStreamPayload`, `encodeSamples`/`decodeSamples`).
- `Documentation/Radio/Thetis-CAT-Command-Reference-Guide-V3.pdf` — official
  CAT command reference.

Gotchas found by testing against a real Thetis instance, not obvious from
reading the protocol alone:

- **CAT sends an unsolicited connect banner.** If "Send Welcome" is on,
  Thetis writes `#Thetis TCP/IP Cat - <version>#;` right after a client
  connects, before any reply to the first command — a naive client can
  misread it as the answer to its first query. `internal/cat/client.go`'s
  `Query` skips non-matching replies for this reason; do the same in any new
  CAT client code.
- **`cw_macros_empty` is CW-Terminal-mode-only, not "message finished."**
  Both places TCIServer.cs fires it are gated on `isTerminalEnabledLocked`
  (lines ~8547, ~8852-8853) — it never fires for a plain `cw_macros` send, no
  matter how long you wait. `cw send` instead polls the bare `trx:<rx>;`
  query (1-arg = get, `handleTrxMessage` TCIServer.cs:3690-3693) and watches
  live PTT/MOX go true then false.
- **The classic Kenwood `PS` (power) command is a disabled stub; `ZZPS` is
  the real, active one.** Same pattern as `RT`/`XT`/`RA`/`PA` above. On TCI,
  the equivalent is the bare `start;`/`stop;` command (not `power:...;`) —
  it's also what the server broadcasts unsolicited to every connected client
  whenever the state changes via any source (`PowerChange` →
  `sendStart`/`sendStop`, TCIServer.cs:1500-1504, 1911-1917), so a client
  sending `start;` must wait for that broadcast echo rather than assume the
  send alone means it worked (see `SetPower`'s doc comment).
- **Not every fully-implemented CAT command is actually reachable — check
  `CATParser.cs`'s dispatch switch, not just `CATCommands.cs`.** `ZZQA`
  (Quick Play) and `ZZQB` (Quick Rec) had complete, correct implementations
  in `CATCommands.cs` that were simply never wired into `CATParser.cs`'s
  switch and had no `CATStructs.xml` entry — unlike the `RT`/`XT`/`RA`/`PA`/`PS`
  stubs above, this wasn't even commented out as deliberately disabled, just
  orphaned. No CAT client, `thetisctl` or otherwise, could reach them until
  wired in (2026-07-30). If a command you expect to work doesn't, check all
  three places (`CATCommands.cs` implementation, `CATParser.cs` dispatch,
  `CATStructs.xml` param widths) before assuming it's a `thetisctl` bug.
- **TCI's "send initial state on connect" burst can shadow real replies.**
  If enabled (`chkTCIsendInitialStateOnConnect`), Thetis dumps ~100+
  unsolicited status frames (one per control) immediately after connect,
  terminated by a bare `ready;` frame. A client that sends a "get" query
  right after connecting and just grabs the first reply matching that
  command name can easily grab a **stale** value from this burst instead of
  the genuine reply — this produced a full page of false "set didn't work"
  failures in the live test suite until fixed by draining the burst (through
  `ready;`) before issuing any query. Confirmed via cross-check: the
  underlying `Set` calls were working correctly the whole time (verified
  independently over CAT while the TCI connection stayed open).
- **`rx_preamp_att_ex` (TCI) / `ZZPA` (CAT) preamp values are quantized, not
  continuous.** Server-side they resolve through `PreampMode` — discrete
  steps only (0, -10, -20, -30, -40, -50 dB, plus SA-prefixed variants,
  `Project Files/Source/Console/enums.cs:236-251`) — not an arbitrary
  integer. Sending an in-between value (e.g. -1) gets silently snapped to
  the nearest valid step; don't treat that as a bug.
- **CAT's step-attenuator getter (`ZZRX`) can hang indefinitely on a live
  radio — but only the getter.** Reproduced repeatedly against a real,
  actively-receiving HL2, independent of Thetis's software power state
  (confirmed on both). `SetAttenuatorDB` was separately confirmed to return
  instantly every time (it's fire-and-forget over CAT, never waits for a
  reply) — the hang is specific to `Query("ZZRX")` waiting on a reply that
  never comes. At the same time, the equivalent TCI value
  (`rx_step_att_ex`) was observed changing on its own with no client
  touching it, suggesting an automatic overload-protection feature reacting
  to real signal conditions and possibly saturating the console's UI-thread
  `Invoke` queue that CAT's getter blocks on (an `Invoke`-based getter blocks
  the caller; a `Send`-based setter does not). Unresolved — if you hit this,
  don't assume it's `thetisctl`'s bug; cross-check via TCI first.
- **`tci query`'s raw passthrough can return the wrong reply if used right
  after connecting.** Caught by the live test suite: `query vfo 0 0` right
  after dialing returned `protocol: [ExpertSDR3 2.0]` — the first line of
  the initial-state burst above — instead of a `vfo:...` reply. Fixed in
  `tciQuery` (`cmd/thetisctl/tci_cmd.go`) to loop until a reply's command
  name matches what was sent, not just take the first frame; see that
  function's doc comment for the residual ambiguity it can't fully resolve
  (matching by command name only, not by leading arguments, since `query`
  accepts arbitrary commands whose argument shape it doesn't know).
- **Sending a TCI command and closing the connection immediately afterward
  can silently drop the command.** This was a real, live safety incident,
  not a theoretical concern: every TX-capable command originally sent its
  unkey (`tune:rx,false;`, `trx:rx,false;`, etc.) and then let the CLI's
  `defer conn.Close()` fire right after — and against a real radio, this
  dropped the unkey more than once, leaving TUNE keyed indefinitely with no
  time bound until a human operator noticed and intervened manually via
  Thetis's own UI. Confirmed by direct testing: the identical unkey command
  sent over a connection kept open a couple of seconds afterward worked
  reliably every time; immediate-close did not. Root cause not fully
  pinned down (client-side write/close race vs. a server-side frame
  ordering issue when a data frame is immediately followed by a close
  frame) — the fix doesn't depend on knowing which: every unkey now goes
  through `confirmTCIUnkeyed`/`confirmCATUnkeyed`
  (`cmd/thetisctl/{tci,cat}_cmd.go`), which sends, then verifies via a query
  that the state actually changed, retrying until confirmed or a timeout is
  hit (in which case it returns an error — it never silently claims
  success). Any new TX-capable command must use this pattern; a bare
  fire-and-forget send followed by closing the connection is not safe for
  anything that unkeys the transmitter.

## Verification / reporting

After any control change, re-query the state (`cat status`, or the relevant
`tci query`) and report the *confirmed* value back to the user — don't just
assume a set command succeeded. Both CAT sets and most TCI sets are
fire-and-forget with no reply, so a follow-up read is the only way to know
Thetis actually applied the change.

## Live test suite

Three files, all build-tag `live` (excluded from normal `go test ./...` and
CI), together covering every remote function this tool exposes:

- `internal/cat/live_test.go` — every exported CAT client function.
- `internal/tci/live_test.go` — every exported TCI client function.
- `cmd/thetisctl/live_test.go` — CLI-layer code the two above bypass by
  calling library functions directly: `rx-audio capture`/`stream` (WAV file
  I/O, stdout PCM streaming) and `query` (raw passthrough), plus a dry run
  (never `--confirm-tx`) of every TX-capable command.

```bash
THETIS_HOST=192.168.2.12 go test -tags=live ./internal/cat/... -v
THETIS_HOST=192.168.2.12 go test -tags=live ./internal/tci/... -v
THETIS_HOST=192.168.2.12 go test -tags=live ./cmd/thetisctl/... -v
```

(`THETIS_CAT_PORT`/`THETIS_TCI_PORT` and `THETIS_LIVE_TIMEOUT` override the
defaults if needed.) Every settable function round-trips: read the current
value, change it, verify, then restore the original via `t.Cleanup` — never
assume a specific starting state. Two exceptions, both documented in their
test's doc comment: `SetBand` is read-only in the test (can retune the VFO
to a stored per-band frequency — bigger disruption than the other reversible
toggles); the CAT attenuator's `Get` is allowed to `t.Skip` on its known hang
(see the gotcha above) rather than fail, with a separate test confirming
`Set` alone doesn't hang.

**None of these three files ever transmits for real.** TX-capable functions
are only ever called in their safe form (`SetPTT`/`SetTrx`/`SetTune` with
`false`; the CLI dry-run path for `ptt`/`tune`/`cw send`/`tx-audio send`,
never with `--confirm-tx`) — an unattended test run can't provide the
per-invocation human confirmation real TX requires.

A fourth file, `cmd/thetisctl/txlive_test.go`, exists **only** for a human
operator to run deliberately when they want real end-to-end TX coverage —
it is never something an agent should run on its own. Skips unconditionally
unless a *second*, independent env var is set to the exact confirm phrase:

```bash
THETIS_HOST=192.168.2.12 THETIS_LIVE_ALLOW_TX=I-UNDERSTAND-THIS-KEYS-THE-RADIO \
    go test -tags=live ./cmd/thetisctl/... -run TestLiveTX -v
```

If an agent is ever asked to "run all the live tests" or similar, that
instruction covers the first three files, not this one — running
`txlive_test.go` requires the same explicit, per-invocation, in-conversation
go-ahead as passing `--confirm-tx` directly, because that's exactly what it
does.
