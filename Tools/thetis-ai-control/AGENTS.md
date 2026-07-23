# thetisctl (Thetis AI-control CLI)

## Purpose

`thetisctl` is a standalone Go CLI that gives an AI agent (or any script)
remote control of a running Thetis instance over its existing CAT-over-TCP
and TCI-over-WebSocket network protocols — control state, RX/TX audio, and
transmit. Companion to `.claude/skills/thetis-control/SKILL.md`, which is the
canonical usage and safety-protocol document; keep this file and that skill
consistent when either changes.

## Ownership

This subtree only (`Tools/thetis-ai-control/`). Not part of the VS solution,
not built by `.github/workflows/build.yml` (the Windows Thetis build), and
has no dependency on `Project Files/`. It has its own CI job (see Verification).

## Local Contracts

- `gofmt -l .` must be empty; `go vet ./...` must be clean.
- No cgo, no live local audio device I/O (mic/speaker) — audio is file
  (WAV) and stdin/stdout PCM only, so the tool stays pure Go and builds
  anywhere. This is a deliberate scope decision, not a placeholder for later.
- Any code path that can key the transmitter (CAT `TX;`, TCI `trx:...,true`,
  TCI `tune:...,true`) must route through `internal/safety`'s dry-run +
  `--confirm-tx` literal-match gate — never key directly.
- Wire formats in `internal/cat` and `internal/tci` were confirmed by reading
  `Project Files/Source/Console/CAT/{CATStructs.xml,CATCommands.cs,CATParser.cs}`
  and `Project Files/Source/Console/TCIServer.cs` directly (comments on each
  typed helper cite line numbers). Re-verify against those files — not just
  each other — before changing a wire format, since Thetis is periodically
  synced from upstream and command dispatch can move.

## Work Guidance

- Dry-run-by-default with an explicit, hard-to-fat-finger `--confirm-tx`
  phrase for anything TX-capable is a durable, user-stated safety
  requirement (neither Thetis network protocol has authentication, and TCI
  TX audio genuinely transmits RF) — do not relax it for convenience.

## Verification

```bash
cd Tools/thetis-ai-control
gofmt -l .            # must be empty
go vet ./...
go test ./...
go build -o thetisctl ./cmd/thetisctl
```

CI: `.github/workflows/thetisctl.yml` runs the above on push/PR touching this
directory, independent of the Windows-only `build.yml`.

Manual end-to-end checklist against a real Thetis+HL2 instance (see
`.claude/skills/thetis-control/SKILL.md` for the full protocol): CAT
freq/mode set with visual confirmation in Thetis's UI; TCI dry-run TX
confirming Thetis's PTT/MOX indicator stays unkeyed; a real low-power TX test
only with the operator present and explicit in-session go-ahead.

## Child DOX Index

None.
