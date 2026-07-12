# `wdsp/main.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Channel object lifecycle (create/destroy/run) and DLL entry points.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/channel.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×8)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/iobuffs.c` (calls ×1)
- Most-referenced symbols from other files: `setInputSamplerate_main()` (×2), `setDSPSamplerate_main()` (×2), `setOutputSamplerate_main()` (×2), `create_main()` (×1), `destroy_main()` (×1), `flush_main()` (×1), `setDSPBuffsize_main()` (×1)

## Outline

### Functions

- `wdspmain()` — L29
- `create_main()` — L63
- `destroy_main()` — L79
- `flush_main()` — L95
- `setInputSamplerate_main()` — L111
- `setOutputSamplerate_main()` — L127
- `setDSPSamplerate_main()` — L143
- `setDSPBuffsize_main()` — L159

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/main.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
