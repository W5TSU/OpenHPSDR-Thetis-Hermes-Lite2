# `wdsp/cblock.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/ssql.c` (calls ×4)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_cbl()` (×2), `destroy_cbl()` (×2), `flush_cbl()` (×2), `xcbl()` (×2), `setSamplerate_cbl()` (×1), `setBuffers_cbl()` (×1), `setSize_cbl()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_cbl()`** — L29 — `void calc_cbl (CBL a)`
  Called by: `create_cbl()` (same file), `setSamplerate_cbl()` (same file)
- **`create_cbl()`** — L38 — `CBL create_cbl ( int run, int buff_size, double *in_buff,`
  Constructor for the `cbl` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `calc_ssql()` (`wdsp/ssql.c`)
- **`destroy_cbl()`** — L63 — `void destroy_cbl(CBL a)`
  Destroys the `cbl` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `decalc_ssql()` (`wdsp/ssql.c`)
- **`flush_cbl()`** — L68 — `void flush_cbl (CBL a)`
  Flushes (zeroes) the `cbl` block’s internal buffers/state.
  Called by: `setSize_cbl()` (same file), `flush_rxa()` (`wdsp/RXA.c`), `flush_ssql()` (`wdsp/ssql.c`)
- **`xcbl()`** — L76 — `void xcbl (CBL a, int position)`
  Runs the `cbl` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xssql()` (`wdsp/ssql.c`)
- **`setBuffers_cbl()`** — L99 — `void setBuffers_cbl (CBL a, double* in, double* out)`
  Re-points the `cbl` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_cbl()`** — L105 — `void setSamplerate_cbl (CBL a, int rate)`
  Reconfigures the `cbl` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_cbl()`** — L111 — `void setSize_cbl (CBL a, int size)`
  Reconfigures the `cbl` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXACBLRun()`** — L123 — `PORT void SetRXACBLRun(int channel, int setit)`
  Sets rxacblrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXACBLPosition()`** — L131 — `PORT void SetRXACBLPosition(int channel, int position)`
  Sets rxacblposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cblock.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
