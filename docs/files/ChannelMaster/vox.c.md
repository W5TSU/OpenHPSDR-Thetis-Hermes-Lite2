# `ChannelMaster/vox.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** VOX detection and TX gain staging.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

### Functions

- `create_vox()` — L29
- `destroy_vox()` — L45
- `flush_vox()` — L51
- `xvox()` — L58
- `create_voxEXT()` — L112
- `destroy_voxEXT()` — L117
- `flush_voxEXT()` — L122
- `xvoxEXT()` — L127
- `SendCBPushVox()` — L140
- `SetTXAVoxRun()` — L146
- `SetTXAVoxSize()` — L155
- `SetTXAVoxThresh()` — L164
- `GetTXAVoxPeak()` — L173

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/vox.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
