# `ChannelMaster/ilv.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/zeer.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_ilv()` (×1), `destroy_ilv()` (×1), `xilv()` (×1), `SetILVOutputPointer()` (×1), `pSetILVInsize()` (×1), `pSetILVRun()` (×1)

## Outline

### Functions

- `create_ilv()` — L31
- `destroy_ilv()` — L51
- `xilv()` — L57
- `SetILVOutputPointer()` — L97
- `SetILVRun()` — L103
- `SetILVWhat()` — L113
- `SetILVInsize()` — L123
- `SetILVOutboundId()` — L130
- `pSetILVRun()` — L137
- `pSetILVInsize()` — L145

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ilv.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
