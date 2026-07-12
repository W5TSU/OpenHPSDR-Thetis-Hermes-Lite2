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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`dR()`** — L29 — `void dR (int n, double* r, double* y, double* z)`
  Called by: `trI()` (same file)
- **`trI()`** — L51 — `void trI ( int n, double* r, double* B, double* y,`
  Called by: `xHat()` (`wdsp/snb.c`)
- **`asolve()`** — L93 — `void asolve(int xsize, int asize, double* x, double* a, double* r, double* z)`
  Called by: `execFrame()` (`wdsp/snb.c`)
- **`median()`** — L127 — `void median (int n, double* a, double* med)`
  Called by: `det()` (`wdsp/snb.c`)
- **`create_builder()`** — L186 — `BLDR create_builder(int points, int ints)`
  Constructor for the `builder` block: allocates its state/buffers and computes initial coefficients.
  Called by: `size_calcc()` (`wdsp/calcc.c`)
- **`destroy_builder()`** — L226 — `void destroy_builder(BLDR a)`
  Destroys the `builder` block, freeing its allocated buffers.
  Called by: `desize_calcc()` (`wdsp/calcc.c`)
- **`flush_builder()`** — L266 — `void flush_builder(BLDR a, int points, int ints)`
  Flushes (zeroes) the `builder` block’s internal buffers/state.
  Called by: `xbuilder()` (same file)
- **`fcompare()`** — L303 — `int fcompare(const void* a, const void* b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`decomp()`** — L313 — `void decomp(int n, double* a, int* piv, int* info, double* wrk)`
  Called by: `xbuilder()` (same file)
- **`dsolve()`** — L372 — `void dsolve(int n, double* a, int* piv, double* b, double* x)`
  Called by: `xbuilder()` (same file)
- **`cull()`** — L394 — `void cull(int* n, int ints, double* x, double* t, double ptol)`
  Called by: `xbuilder()` (same file)
- **`xbuilder()`** — L411 — `void xbuilder(BLDR a, int points, double* x, double* y, int ints, double* t, int* info, double* c, double ptol)`
  Runs the `builder` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `calc()` (`wdsp/calcc.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/lmath.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
