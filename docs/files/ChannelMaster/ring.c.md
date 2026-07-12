# `ChannelMaster/ring.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)

## Outline

### Functions

- `ringbuffer_create()` — L36
- `ringbuffer_free()` — L62
- `ringbuffer_reset_size()` — L69
- `ringbuffer_reset()` — L80
- `ringbuffer_clear()` — L89
- `ringbuffer_restart()` — L99
- `ringbuffer_read_space()` — L106
- `ringbuffer_write_space()` — L125
- `ringbuffer_write()` — L144
- `ringbuffer_read()` — L181

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ring.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
