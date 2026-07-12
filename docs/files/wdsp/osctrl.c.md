# `wdsp/osctrl.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** TX overshoot control.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/TXA.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_osctrl()` (×1), `destroy_osctrl()` (×1), `flush_osctrl()` (×1), `xosctrl()` (×1), `setSamplerate_osctrl()` (×1), `setBuffers_osctrl()` (×1), `setSize_osctrl()` (×1)

## Outline

### Functions

- `calc_osctrl()` — L33
- `decalc_osctrl()` — L46
- `create_osctrl()` — L52
- `destroy_osctrl()` — L72
- `flush_osctrl()` — L78
- `xosctrl()` — L84
- `setBuffers_osctrl()` — L116
- `setSamplerate_osctrl()` — L122
- `setSize_osctrl()` — L129
- `SetTXAosctrlRun()` — L141

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/osctrl.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
