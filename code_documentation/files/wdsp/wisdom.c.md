# `wdsp/wisdom.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FFTW wisdom generation/caching and FIR impulse-response caching for fast startup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`wisdom_get_status()`** — L32 — `PORT char* wisdom_get_status()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`WDSPwisdom()`** — L38 — `PORT int WDSPwisdom (char* directory)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/wisdom.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
