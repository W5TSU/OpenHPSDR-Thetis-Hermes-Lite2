# `ChannelMaster/cmasio.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Bridge to the cmASIO DLL for direct ASIO device I/O.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/netInterface.c` (calls ×2)
- Uses (outgoing references to other files):
  - `cmASIO/hostsample.cpp` (calls ×9)
  - `wdsp/rmatch.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/obbuffs.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `ChannelMaster/ivac.c` (calls ×1)
  - `ChannelMaster/obbuffs.c` (calls ×1)
- Most-referenced symbols from other files: `create_cmasio()` (×1), `destroy_cmasio()` (×1), `asioIN()` (×1), `cm_asioStart()` (×1), `cm_asioStop()` (×1)

## Outline

### Functions

- `create_cmasio()` — L47
- `destroy_cmasio()` — L104
- `asioIN()` — L120
- `asioOUT()` — L137
- `CallbackASIO()` — L150
- `cm_asioStart()` — L337
- `cm_asioStop()` — L358
- `getCMAstate()` — L369
- `getCMAevents()` — L379
- `resetCMAevents()` — L396

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmasio.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
