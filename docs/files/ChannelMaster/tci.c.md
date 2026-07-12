# `ChannelMaster/tci.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** TCI (Transceiver Control Interface) TCP server for SDC/logger integration at the audio layer.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `ChannelMaster/aamix.c` (calls ×14)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetTCIRxAudioRate()` (×1), `SetTCIRxAudioSize()` (×1), `SetTCITxMonitorRate()` (×1), `create_tci()` (×1), `destroy_tci()` (×1), `xtciOUT()` (×1)

## Outline

### Functions

- `get_tci_audio_mix_state()` — L65
- `apply_tci_audio_mix_state()` — L88
- `tci_audio_out()` — L118
- `create_tci_audio_mixer()` — L124
- `destroy_tci_audio_mixer()` — L162
- `create_tci()` — L179
- `destroy_tci()` — L192
- `xtciOUT()` — L200
- `SendpOutboundTCIRxAudio()` — L220
- `SetTCIRxAudioRate()` — L226
- `SetTCIRxAudioSize()` — L250
- `SetTCITxMonitorRate()` — L267
- `SetTCIRxAudioMox()` — L288
- `SetTCIRxAudioMon()` — L299
- `SetTCIRxAudioMonVol()` — L310

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/tci.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
