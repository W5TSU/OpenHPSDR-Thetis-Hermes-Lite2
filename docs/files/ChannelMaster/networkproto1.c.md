# `ChannelMaster/networkproto1.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×1)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/netInterface.c` (calls ×6)
  - `ChannelMaster/pro.c` (calls ×5)
  - `ChannelMaster/network.c` (calls ×3)
  - `ChannelMaster/router.c` (calls ×3)
  - `ChannelMaster/cmbuffs.c` (calls ×2)
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/bandwidth_monitor.c` (calls ×1)
  - `ChannelMaster/network.h` (imports ×1)
  - `ChannelMaster/pro.h` (imports ×1)
- Most-referenced symbols from other files: `SendStartToMetis()` (×1), `SendStopToMetis()` (×1)

## Outline

### Functions

- `SendStartToMetis()` — L38
- `SendStopToMetis()` — L76
- `ForceCandCFrames()` — L111
- `ForceCandCFrame()` — L139
- `MetisReadDirect()` — L146
- `MetisWriteFrame()` — L221
- `MetisReadThreadMain()` — L245
- `twist()` — L268
- `MetisReadThreadMainLoop()` — L281
- `MetisReadThreadMainLoop_HL2()` — L427
- `WriteMainLoop()` — L593
- `WriteMainLoop_HL2()` — L874
- `sendProtocol1Samples()` — L1209

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/networkproto1.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
