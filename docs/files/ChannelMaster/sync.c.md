# `ChannelMaster/sync.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/router.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×3)
  - `wdsp/div.c` (calls ×3)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/calcc.c` (calls ×1)
- Most-referenced symbols from other files: `create_sync()` (×1), `destroy_sync()` (×1), `InboundBlock()` (×1)

## Outline

### Functions

- `create_sync()` — L32
- `destroy_sync()` — L38
- `InboundBlock()` — L44
- `SetPSTxIdx()` — L69
- `SetPSRxIdx()` — L75

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/sync.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
