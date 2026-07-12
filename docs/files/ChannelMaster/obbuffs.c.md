# `ChannelMaster/obbuffs.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×2)
  - `ChannelMaster/cmasio.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/obbuffs.h` (imports ×1)
- Most-referenced symbols from other files: `OutBound()` (×1), `destroy_obbuffs()` (×1), `create_obbuffs()` (×1)

## Outline

### Functions

- `start_obthread()` — L37
- `create_obbuffs()` — L43
- `destroy_obbuffs()` — L68
- `flush_obbuffs()` — L88
- `OutBound()` — L98
- `obdata()` — L132
- `ob_main()` — L153
- `SetOBRingOutsize()` — L175

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/obbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
