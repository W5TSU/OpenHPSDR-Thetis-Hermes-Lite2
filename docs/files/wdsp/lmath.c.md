# `wdsp/lmath.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/calcc.c` (calls ×3)
  - `wdsp/snb.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_builder()` (×1), `destroy_builder()` (×1), `xbuilder()` (×1), `trI()` (×1), `asolve()` (×1), `median()` (×1)

## Outline

### Functions

- `dR()` — L29
- `trI()` — L51
- `asolve()` — L93
- `median()` — L127
- `create_builder()` — L186
- `destroy_builder()` — L226
- `flush_builder()` — L266
- `fcompare()` — L303
- `decomp()` — L313
- `dsolve()` — L372
- `cull()` — L394
- `xbuilder()` — L411

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/lmath.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
