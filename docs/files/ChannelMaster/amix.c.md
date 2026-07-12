# `ChannelMaster/amix.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Audio mixers (monitor mix, multi-RX audio combination) with per-input gain and slew.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

### Functions

- `create_amix()` — L29
- `destroy_amix()` — L51
- `xamix()` — L57
- `create_amixEXT()` — L101
- `destroy_amixEXT()` — L106
- `xamixEXT()` — L111
- `SetAudioMixWhat()` — L121
- `SetAudioMixSize()` — L144
- `SetAudioMixVolume()` — L153
- `SetAudioMixVol()` — L168

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/amix.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
