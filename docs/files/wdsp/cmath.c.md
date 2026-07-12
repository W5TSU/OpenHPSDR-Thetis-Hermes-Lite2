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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`mag()`** — L29 — `double mag(double* value)`
  function to calculate the magnitude of a complex value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`cadd()`** — L35 — `void cadd(double* a, double* b, double* sum)`
  function to perform a Complex Add, a+b; it returns a complex value, 'sum'
  Called by: `cpar()` (same file)
- **`csub()`** — L42 — `void csub(double* a, double* b, double* diff)`
  function to perform a Complex Subtract, a-b; it returns a complex value, 'diff'
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`cmult()`** — L49 — `void cmult(double* a, double* b, double* product)`
  function to perform a Complex Multiply, a*b; it returns a complex value, 'product'
  Called by: `cpar()` (same file), `build_doublepole_1sided()` (`wdsp/doublepole.c`)
- **`cdiv()`** — L56 — `void cdiv(double* a, double* b, double* quotient)`
  function to perform a Complex Divide, a/b; it returns a complex value, 'quotient'
  Called by: `cpar()` (same file)
- **`cpar()`** — L64 — `void cpar(double* Z1, double* Z2, double* Zpar)`
  function to calculate complex Z (series equivalent) of two parallel elements
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`cser_to_par()`** — L73 — `void cser_to_par(double* Z1, double* ZR, double* ZX)`
  function to convert a complex Z to parallel R and X values
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cmath.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
