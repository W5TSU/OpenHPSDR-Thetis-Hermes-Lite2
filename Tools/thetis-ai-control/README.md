# thetisctl

A Go CLI for remotely controlling a running [Thetis](../../README.md) SDR
instance over its existing network protocols: CAT-over-TCP (frequency, mode,
filters, VFO, RIT/XIT, split, AGC, attenuator/preamp, band) and
TCI-over-WebSocket (the same controls plus RX/TX audio streaming and
transmit). No Thetis-side code changes are required — both servers already
exist in Thetis, they're just usually off by default.

Full usage, the TX safety protocol, and the intended AI-agent workflow are
documented in [`.claude/skills/thetis-control/SKILL.md`](../../.claude/skills/thetis-control/SKILL.md).
This file only covers building and running the binary directly.

## Build

```bash
cd Tools/thetis-ai-control
go build -o thetisctl ./cmd/thetisctl
go vet ./...
go test ./...
```

Pure Go, no cgo, no external dependencies — builds anywhere Go runs.

## Quick usage

```bash
./thetisctl help

# Tier 1: control only, over CAT
./thetisctl cat --host 192.168.1.50 freq set A 14074000
./thetisctl cat --host 192.168.1.50 status

# Tier 2: control + audio + transmit, over TCI
./thetisctl tci --host 192.168.1.50 vfo 0 0 14074000
./thetisctl tci --host 192.168.1.50 rx-audio capture 0 --duration 5s --out rx.wav
```

Every TX-capable command (`cat ptt`, `tci tune`, `tci ptt`, `tci tx-audio
send`) defaults to a dry run. See the skill doc before using
`--confirm-tx`.
