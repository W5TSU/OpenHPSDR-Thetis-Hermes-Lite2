# `ChannelMaster/txgain.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** VOX detection and TX gain staging.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×4)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_txgain()` (×1), `destroy_txgain()` (×1), `xtxgain()` (×1), `SetTXGainSize()` (×1), `SetAmpProtectADCValue()` (×1)

## Outline

### Functions

- `create_txgain()` — L29
- `destroy_txgain()` — L58
- `xtxgain()` — L65
- `SetTXGainSize()` — L111
- `SetTXFixedGainRun()` — L116
- `SetTXFixedGain()` — L126
- `SetAmpProtectADCValue()` — L137
- `GetAndResetAmpProtect()` — L146
- `SetAmpProtectRun()` — L153
- `SetADCSupply()` — L163

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/txgain.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
