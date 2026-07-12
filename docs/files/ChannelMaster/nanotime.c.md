# `ChannelMaster/nanotime.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Network bandwidth statistics and high-resolution timestamps.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/nanotimer.h` (imports ×1)

## Outline

### Functions

- `getPerfTicks()` — L32
- `getPerfFreq()` — L46
- `perfTicksToNanos()` — L57
- `updateHLA()` — L70
- `initHLA()` — L82
- `printHLA()` — L90
- `printHLANano()` — L109

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/nanotime.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
