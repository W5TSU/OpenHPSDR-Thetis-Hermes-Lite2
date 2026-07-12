# `ChannelMaster/pro.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/networkproto1.c` (calls ×5)
- Uses (outgoing references to other files):
  - `ChannelMaster/pro.h` (imports ×1)
- Most-referenced symbols from other files: `destroy_pro()` (×3), `create_pro()` (×1), `xpro()` (×1)

## Outline

### Functions

- `create_pro()` — L29
- `destroy_pro()` — L55
- `xpro()` — L67

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/pro.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
