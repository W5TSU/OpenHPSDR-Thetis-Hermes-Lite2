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
| `cat --host <ip> status` | Combined ID + frequency/mode/RIT/XIT/split/TX status |

```bash
./thetisctl cat --host 192.168.1.50 freq set A 14074000
./thetisctl cat --host 192.168.1.50 mode set USB
./thetisctl cat --host 192.168.1.50 status
```

`ptt on|off` also exists on the CAT channel but is TX-capable — see the
safety protocol below before using it.

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
| `tci --host <ip> cw send <rx> "<text>" --speed <wpm> --mode <cw\|cwu\|cwl>` | Key CW text via Thetis's own macro keyer |
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
same safety protocol applies.

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

`cat ptt`, `tci tune`, `tci ptt`, `tci cw send`, and `tci tx-audio send` can
key the transmitter. Since neither network protocol has authentication,
`thetisctl` enforces:

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

## Verification / reporting

After any control change, re-query the state (`cat status`, or the relevant
`tci query`) and report the *confirmed* value back to the user — don't just
assume a set command succeeded. Both CAT sets and most TCI sets are
fire-and-forget with no reply, so a follow-up read is the only way to know
Thetis actually applied the change.
