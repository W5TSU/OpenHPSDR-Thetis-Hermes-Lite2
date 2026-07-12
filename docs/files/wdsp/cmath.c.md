# `wdsp/cmath.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/doublepole.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `cmult()` (×1)

## Outline

### Functions

- `mag()` — L29
- `cadd()` — L35
- `csub()` — L42
- `cmult()` — L49
- `cdiv()` — L56
- `cpar()` — L64
- `cser_to_par()` — L73

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cmath.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
