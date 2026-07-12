# `wdsp/fcurve.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/emph.c` (calls ×6)
  - `wdsp/fmd.c` (calls ×5)
- Uses (outgoing references to other files):
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/impulse_cache.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `fc_impulse()` (×10), `fc_mults()` (×1)

## Outline

### Functions

- `fc_impulse()` — L29
- `fc_mults()` — L185

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fcurve.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
