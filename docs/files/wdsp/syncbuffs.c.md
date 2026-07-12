# `wdsp/syncbuffs.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)

## Outline

### Functions

- `start_syncbthread()` — L29
- `create_syncbuffs()` — L35
- `destroy_syncbuffs()` — L65
- `flush_syncbuffs()` — L86
- `Syncbound()` — L97
- `syncbdata()` — L132
- `syncb_main()` — L162
- `SetSYNCBRingOutsize()` — L175
- `create_dumfilt()` — L199
- `destroy_dumfilt()` — L214
- `flush_dumfilt()` — L220
- `xdumfilt()` — L227

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/syncbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
