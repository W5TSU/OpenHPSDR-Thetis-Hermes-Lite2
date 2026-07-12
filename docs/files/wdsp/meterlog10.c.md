# `wdsp/meterlog10.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/analyzer.c` (calls ×3)
  - `wdsp/cfcomp.c` (calls ×1)
  - `wdsp/emnr.c` (calls ×1)
  - `wdsp/meter.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `mlog10()` (×7)

## Outline

### Functions

- `mlog10()` — L547

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/meterlog10.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
