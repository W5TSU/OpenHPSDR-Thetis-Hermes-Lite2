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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fc_impulse()`** — L29 — `double* fc_impulse (int nc, double f0, double f1, double g0, double g1, int curve, double samplerate, double scale, int ctfmode, int wintype)`
  Called by: `fc_mults()` (same file), `create_emphp()` (`wdsp/emph.c`), `setSamplerate_emphp()` (`wdsp/emph.c`), `setSize_emphp()` (`wdsp/emph.c`), `SetTXAFMEmphNC()` (`wdsp/emph.c`), `SetTXAFMPreEmphFreqs()` (`wdsp/emph.c`) — and 5 more
- **`fc_mults()`** — L185 — `double* fc_mults (int size, double f0, double f1, double g0, double g1, int curve, double samplerate, double scale, int ctfmode, int wintype)`
  generate mask for Overlap-Save Filter
  Called by: `calc_emph()` (`wdsp/emph.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fcurve.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
