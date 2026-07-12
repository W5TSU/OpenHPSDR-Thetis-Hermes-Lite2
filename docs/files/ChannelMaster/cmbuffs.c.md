# `ChannelMaster/cmbuffs.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/networkproto1.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/router.c` (calls ×1)
  - `ChannelMaster/sync.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmaster.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `Inbound()` (×5), `create_cmbuffs()` (×1), `destroy_cmbuffs()` (×1), `SetCMRingOutsize()` (×1)

## Outline

### Functions

- `start_cmthread()` — L29
- `create_cmbuffs()` — L35
- `destroy_cmbuffs()` — L60
- `flush_cmbuffs()` — L78
- `Inbound()` — L88
- `cmdata()` — L123
- `cm_main()` — L151
- `SetCMRingOutsize()` — L170

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
