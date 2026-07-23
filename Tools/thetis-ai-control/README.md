# thetisctl

A Go CLI for remotely controlling a running [Thetis](../../README.md) SDR
instance over its existing network protocols: CAT-over-TCP (frequency, mode,
filters, VFO, RIT/XIT, split, AGC, attenuator/preamp, band) and
TCI-over-WebSocket (the same controls plus RX/TX audio streaming, CW keying,
and transmit). No Thetis-side code changes are required — both servers
already exist in Thetis, they're just usually off by default.

This file is a plain command reference for a human running `thetisctl`
directly. For the AI-agent workflow and the full TX safety protocol, see
[`.claude/skills/thetis-control/SKILL.md`](../../.claude/skills/thetis-control/SKILL.md).

## Enabling the servers in Thetis

Open **Setup** in the Thetis instance you want to control and turn on:

- **TCP/IP CAT Server** — for Tier 1 commands (`thetisctl cat ...`). Default
  bind `127.0.0.1:13013`. If `thetisctl` runs on a different machine, rebind
  to the host's LAN IP (e.g. `192.168.1.50:13013`), not `127.0.0.1`.
- **TCI Server** — for Tier 2 commands (`thetisctl tci ...`). Default bind
  `127.0.0.1:50001`, same rebind note.

Neither server has authentication — anyone who can reach the bound
address/port can issue any command, including keying the transmitter. Only
bind these on a trusted LAN; never expose them to the internet.

## Build

```bash
cd Tools/thetis-ai-control
go build -o thetisctl ./cmd/thetisctl
go vet ./...
go test ./...
```

Pure Go, no cgo, no external dependencies — builds anywhere Go runs.

## Global flags

Both `cat` and `tci` take:

| Flag | Default | Meaning |
|---|---|---|
| `--host <ip>` | *(required)* | Thetis's address — never assumed local |
| `--port <n>` | `13013` (cat) / `50001` (tci) | Server port |
| `--timeout <duration>` | `3s` (cat) / `5s` (tci) | Network read/write timeout |

## CAT commands — control only (`thetisctl cat ...`)

| Command | Effect |
|---|---|
| `freq get A\|B` | Read VFO A or B frequency (Hz) |
| `freq set A\|B <hz>` | Set VFO A or B frequency, then reads it back to confirm |
| `mode get` | Read the demod mode |
| `mode set <name>` | Set mode: `USB`, `LSB`, `CW`, `CWL`, `FM`, `AM`, `DIGU`, `DIGL` |
| `rit on\|off\|get` | RIT enable/disable/query |
| `xit on\|off\|get` | XIT enable/disable/query |
| `split on\|off\|get` | VFO split enable/disable/query |
| `agc get` | Read AGC mode |
| `agc set <name>` | Set AGC: `FIXED`, `LONG`, `SLOW`, `MEDIUM`, `FAST`, `CUSTOM` |
| `atten get` | Read RX1 step attenuator (dB) |
| `atten set <0-31>` | Set RX1 step attenuator (dB) |
| `preamp set <0-9>` | Set RX1 preamp level (0=off, 1=on, 2-6=-10..-50dB, 7-9=SA -10..-30dB) |
| `band get` | Read current band |
| `band set <name>` | Set band: `160`,`80`,`60`,`40`,`30`,`20`,`17`,`15`,`12`,`10`,`6`,`2`,`GEN`,`WWV`,`V0`-`V13` |
| `status` | Rig ID + frequency/mode/RIT/XIT/split/TX in one call |
| `ptt on\|off` | **TX-capable** — see [Transmitting](#transmitting-tx-capable-commands) |

```bash
./thetisctl cat --host 192.168.1.50 freq set A 14074000
./thetisctl cat --host 192.168.1.50 mode set USB
./thetisctl cat --host 192.168.1.50 status
```

## TCI commands — control, audio, and transmit (`thetisctl tci ...`)

`rx` selects the receiver: `0` = RX1, `1` = RX2.

| Command | Effect |
|---|---|
| `vfo <rx> <chan 0\|1> <hz>` | Set VFO A(0)/B(1) frequency |
| `modulation <rx> <mode>` | Set mode: `lsb`,`usb`,`dsb`,`am`,`sam`,`fm`,`cw`,`cwl`,`cwu`,`digl`,`digu` |
| `split <rx> on\|off` | VFO split |
| `rit <rx> on\|off` | RIT enable |
| `xit <rx> on\|off` | XIT enable |
| `rit-offset <rx> <hz>` | RIT offset |
| `xit-offset <rx> <hz>` | XIT offset |
| `filter <rx> <lowHz> <highHz>` | RX filter passband edges |
| `atten <rx> <dB>` | Step attenuator (dB, ≥0) |
| `preamp <rx> <dB>` | Preamp gain, expressed as attenuation ≤0 (e.g. `-10`) |
| `agc <rx> <mode>` | AGC mode: `off`/`fixed`, `long`, `slow`, `medium`/`normal`, `fast`, `custom` |
| `agc-gain <rx> <-20..120>` | AGC gain/threshold |
| `drive <rx> <0-100>` | TX drive power |
| `rx-audio capture <rx> --duration <d> --out <file.wav>` | Record RX audio to a WAV file |
| `rx-audio stream <rx> --duration <d>` | Stream RX audio as raw float32 LE PCM to stdout |
| `tune <rx> on\|off` | **TX-capable** — key TUNE (bare carrier) |
| `ptt <rx> on\|off [--audio]` | **TX-capable** — key PTT (`--audio` marks this connection as the TX audio source) |
| `cw send <rx> "<text>" --speed <wpm> --mode <cw\|cwu\|cwl>` | **TX-capable** — key CW text via Thetis's own macro keyer |
| `tx-audio send <rx> --file <wav>` | **TX-capable** — stream a WAV file as TX audio |
| `query <cmd> [args...]` | Raw passthrough — send any TCI text command not listed above and print the reply |

```bash
./thetisctl tci --host 192.168.1.50 vfo 0 0 14074000
./thetisctl tci --host 192.168.1.50 modulation 0 usb
./thetisctl tci --host 192.168.1.50 rx-audio capture 0 --duration 10s --out capture.wav
./thetisctl tci --host 192.168.1.50 rx-audio stream 0 --duration 10s > raw.pcm
```

## Transmitting (TX-capable commands)

`cat ptt`, `tci tune`, `tci ptt`, `tci cw send`, and `tci tx-audio send` can
key the transmitter. **Every one of them defaults to a dry run** — without
`--confirm-tx`, they print exactly what they would send and do nothing
TX-capable:

```
$ ./thetisctl tci --host 192.168.1.50 cw send 0 "CQ CQ DE W5TSU" --speed 5 --mode cwu
[dry-run] would send: modulation:0,cwu; cw_macros_speed:5; cw_macros:0,CQ CQ DE W5TSU;
Pass --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO to actually transmit this message.
```

To actually transmit, add the exact phrase (not a bare flag — this is
deliberate, so nothing else can accidentally trigger it):

```bash
./thetisctl tci --host 192.168.1.50 cw send 0 "CQ CQ DE W5TSU" --speed 5 --mode cwu \
    --confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO --max-duration 60s
```

Other TX flags:

| Flag | Applies to | Meaning |
|---|---|---|
| `--confirm-tx=I-UNDERSTAND-THIS-KEYS-THE-RADIO` | all TX-capable commands | Required to key for real; anything else stays a dry run |
| `--hold <duration>` (default `3s`) | `cat ptt`, `tci tune`, `tci ptt` | Auto-unkeys after this long |
| `--max-duration <duration>` (default `10s` for tx-audio, `90s` for cw) | `tci tx-audio send`, `tci cw send` | Hard cap; truncates/stops and unkeys if exceeded |

Every TX-capable command also unkeys automatically on completion, error, or
Ctrl-C — the radio is never deliberately left keyed by a crashed or
interrupted invocation.

## Notes on real-world behavior

- **CAT connect banner.** If Thetis's "Send Welcome" option is on, connecting
  to the CAT port gets you an unsolicited `#Thetis TCP/IP Cat - <version>#;`
  line before any command reply. `thetisctl` already accounts for this.
- **CW completion isn't a "message finished" event.** Thetis's TCI protocol
  has no such event for a plain `cw_macros` send (`cw_macros_empty` only
  fires in CW Terminal mode, which `cw send` doesn't use). `cw send` instead
  polls live PTT state and reports done once it sees the radio key up and
  then release.
- **Switching to CW mode** can shift the displayed/tuned frequency by
  Thetis's CW pitch offset (commonly 600 Hz) — this is normal sidetone-offset
  behavior, not a bug in `thetisctl`.

## Extending thetisctl

Wire formats were confirmed by reading Thetis's own source, not just
protocol docs — see
[`SKILL.md`'s "Extending the command set"](../../.claude/skills/thetis-control/SKILL.md#extending-the-command-set)
for the exact files and gotchas to check before adding new commands.
