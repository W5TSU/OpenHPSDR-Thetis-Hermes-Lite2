# `wdsp/gain.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

### Functions

- `create_gain()` — L29
- `destroy_gain()` — L44
- `flush_gain()` — L51
- `xgain()` — L57
- `setBuffers_gain()` — L80
- `setSamplerate_gain()` — L86
- `setSize_gain()` — L91
- `pSetTXOutputLevel()` — L102
- `pSetTXOutputLevelRun()` — L111
- `pSetTXOutputLevelSize()` — L119

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gain.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
