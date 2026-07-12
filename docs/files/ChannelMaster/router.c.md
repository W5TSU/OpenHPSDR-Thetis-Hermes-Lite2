# `ChannelMaster/router.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** The signal routing matrix — which input feeds which DSP channel and which output.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/networkproto1.c` (calls ×3)
  - `ChannelMaster/cmaster.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `ChannelMaster/sync.c` (calls ×1)
- Most-referenced symbols from other files: `xrouter()` (×4), `create_router()` (×1), `destroy_router()` (×1)

## Outline

### Functions

- `create_router()` — L32
- `destroy_router()` — L43
- `flush_router()` — L52
- `xrouter()` — L70
- `LoadRouterAll()` — L110
- `LoadRouterControlBit()` — L145

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/router.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
