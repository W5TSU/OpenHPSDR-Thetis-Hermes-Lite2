# First-party source tree

## Purpose

The six subprojects composing Thetis for the Hermes-Lite 2: `Console/` (C# WinForms UI and
radio control), `wdsp/` (C DSP library), `ChannelMaster/` (C/C++ audio routing and HPSDR
network protocols), `cmASIO/` (ASIO wrapper), `Midi2Cat/`, `RawInput/`.

## Ownership

- This doc owns code conventions for everything under `Project Files/Source`
- `Thetis-Installer/` (WiX) and `packages/` are build plumbing owned by the root doc

## Local Contracts

- This is a fork periodically synced from mi0bot/OpenHPSDR-Thetis: keep diffs against upstream
  minimal, and tag fork-specific changes with a callsign comment (`// MI0BOT:`, `// W5TSU:`)
  as the existing code does
- `*.designer.cs` files are Visual Studio-generated; when editing without VS, mirror the
  designer's exact output patterns (declaration + init block + Controls.Add + field) and never
  reformat surrounding lines
- Version is bumped only in `Console/AssemblyInfo.cs` and only during the release step
  (`.claude/skills/thetis-merge/SKILL.md`); the MSI version derives from it via WiX bind
- `HPSDRModel`/`HPSDRHW` enums (`Console/enums.cs`): add members before `LAST`, never reorder —
  the int values are persisted in user databases
- Setup/console controls persist to the settings database keyed by control name; renaming a
  control silently orphans users' saved settings
- HL2-specific behavior concentrates in `Console/HPSDR/IoBoardHl2.cs`, the `HERMESLITE` /
  `HPSDRHW.HermesLite` switch cases, and `ChannelMaster/networkproto1.c` `_HL2` loops

## Work Guidance

- The wdsp and ChannelMaster C code is called from C# via P/Invoke — signature changes must be
  mirrored on both sides
- UI updates from DSP/audio threads must go through the thread-safe `*TS` wrappers in
  `Console/Invoke/`

## Verification

- No local build exists on Linux; compile-check via GitHub Actions:
  `gh workflow run build.yml --ref <branch> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`, then
  `gh run watch <id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 --exit-status` (~4 min)
- Runtime behavior needs the user's on-air test with the HL2 hardware — ask rather than assume

## Child DOX Index

(none)
