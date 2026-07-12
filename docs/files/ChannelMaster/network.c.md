# `ChannelMaster/network.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×64)
  - `ChannelMaster/networkproto1.c` (calls ×3)
  - `ChannelMaster/obbuffs.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/netInterface.c` (calls ×4)
  - `ChannelMaster/bandwidth_monitor.c` (calls ×2)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmsetup.c` (calls ×1)
  - `ChannelMaster/network.h` (imports ×1)
  - `ChannelMaster/networkproto1.c` (calls ×1)
  - `wdsp/analyzer.c` (calls ×1)
  - `ChannelMaster/router.c` (calls ×1)
  - `ChannelMaster/sidetone.c` (calls ×1)
  - `ChannelMaster/txgain.c` (calls ×1)
- Most-referenced symbols from other files: `CmdHighPriority()` (×26), `CmdTx()` (×23), `CmdGeneral()` (×7), `CmdRx()` (×4), `sendPacket()` (×3), `StopReadThread()` (×2), `StartReadThread()` (×1), `IOThreadStop()` (×1)

## Outline

### Functions

- `initWSA()` — L53
- `DeInitMetisSockets()` — L75
- `nativeInitMetis()` — L87
- `GetMetisIPAddr()` — L348
- `SendStart()` — L353
- `SendStop()` — L363
- `StartReadThread()` — L371
- `StopReadThread()` — L389
- `addSnapShot()` — L398
- `storeRXSeqDelta()` — L456
- `ReadUDPFrame()` — L472
- `ReadThreadMainLoop()` — L635
- `CmdGeneral()` — L812
- `CmdHighPriority()` — L904
- `CmdRx()` — L1056
- `CmdTx()` — L1172
- `sendOutbound()` — L1241
- `WriteUDPFrame()` — L1347
- `sendPacket()` — L1386
- `KeepAliveLoop()` — L1408
- `IOThreadStop()` — L1438
- `ReadThreadMain()` — L1478
- `KeepAliveMain()` — L1493

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/network.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
