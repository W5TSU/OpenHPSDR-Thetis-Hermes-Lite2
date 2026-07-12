# `ChannelMaster/analyzers.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Attaches wdsp spectrum analyzers to ChannelMaster streams (RX/TX displays).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×3)
  - `wdsp/analyzer.c` (calls ×2)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Most-referenced symbols from other files: `create_analyzer_alloc()` (×1), `destroy_analyzer_alloc()` (×1)

## Outline

### Functions

- `create_analyzer_alloc()` — L32
- `destroy_analyzer_alloc()` — L66
- `tx_analyzers()` — L87
- `alloc_analyzer()` — L121
- `free_analyzer()` — L157
- `run_analyzer()` — L170

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/analyzers.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
