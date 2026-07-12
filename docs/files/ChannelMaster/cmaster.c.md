# `ChannelMaster/cmaster.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Stream lifecycle and top-level orchestration: creates channels, starts/stops audio and network streams.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/netInterface.c` (calls ×2)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×30)
  - `ChannelMaster/aamix.c` (calls ×15)
  - `wdsp/channel.c` (calls ×8)
  - `ChannelMaster/ivac.c` (calls ×7)
  - `wdsp/analyzer.c` (calls ×5)
  - `ChannelMaster/ilv.c` (calls ×5)
  - `ChannelMaster/sidetone.c` (calls ×5)
  - `wdsp/dexp.c` (calls ×5)
  - `wdsp/eer.c` (calls ×5)
  - `ChannelMaster/txgain.c` (calls ×4)
  - `ChannelMaster/cmasio.c` (calls ×3)
  - `wdsp/nob.c` (calls ×3)
  - …and 12 more files
- Most-referenced symbols from other files: `create_cmaster()` (×1), `destroy_cmaster()` (×1), `xcmaster()` (×1), `SendpOutboundRx()` (×1), `SendpOutboundTx()` (×1)

## Outline

### Functions

- `create_rcvr()` — L33
- `destroy_rcvr()` — L96
- `create_xmtr()` — L112
- `destroy_xmtr()` — L255
- `create_cmaster()` — L273
- `destroy_cmaster()` — L322
- `xcmaster()` — L339
- `SendpOutboundRx()` — L407
- `SendpOutboundTx()` — L414
- `SendpOutboundTCIRxIQ()` — L422
- `SendpInboundTCITxAudio()` — L428
- `SetRXTCIRun()` — L434
- `SetTXTCIAudioRun()` — L440
- `SetRunPanadapter()` — L447
- `SetXcmInrate()` — L453
- `SetCMAudioOutrate()` — L510
- `SetRcvrChannelOutrate()` — L522
- `SetXmtrChannelOutrate()` — L549
- `SetAntiVOXSourceStates()` — L583
- `SetAntiVOXSourceWhat()` — L590

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmaster.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
