# `wdsp/cfcomp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** TX speech compressor and continuous frequency compressor.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `create_cfcomp()` (×1), `destroy_cfcomp()` (×1), `flush_cfcomp()` (×1), `xcfcomp()` (×1), `setSamplerate_cfcomp()` (×1), `setBuffers_cfcomp()` (×1), `setSize_cfcomp()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_cfcwindow()`** — L51 — `void calc_cfcwindow (CFCOMP a)`
  Called by: `calc_cfcomp()` (same file)
- **`fCOMPcompare()`** — L98 — `int fCOMPcompare (const void * a, const void * b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_comp()`** — L108 — `void calc_comp(CFCOMP a)`
  Called by: `calc_cfcomp()` (same file), `SetTXACFCOMPprofile()` (same file)
- **`calc_cfcomp()`** — L358 — `void calc_cfcomp(CFCOMP a)`
  Called by: `create_cfcomp()` (same file), `setSamplerate_cfcomp()` (same file), `setSize_cfcomp()` (same file)
- **`decalc_cfcomp()`** — L420 — `void decalc_cfcomp(CFCOMP a)`
  Called by: `destroy_cfcomp()` (same file), `setSamplerate_cfcomp()` (same file), `setSize_cfcomp()` (same file)
- **`create_cfcomp()`** — L449 — `CFCOMP create_cfcomp (int run, int position, int peq_run, int size, double* in, double* out, int fsize, int ovrlp, int rate, int wintype, int comp_method, int nfreqs, double precom`
  Constructor for the `cfcomp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`flush_cfcomp()`** — L482 — `void flush_cfcomp (CFCOMP a)`
  Flushes (zeroes) the `cfcomp` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`destroy_cfcomp()`** — L499 — `void destroy_cfcomp (CFCOMP a)`
  Destroys the `cfcomp` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`calc_mask()`** — L511 — `void calc_mask (CFCOMP a)`
  Called by: `xcfcomp()` (same file)
- **`xcfcomp()`** — L554 — `void xcfcomp (CFCOMP a, int pos)`
  Runs the `cfcomp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_cfcomp()`** — L607 — `void setBuffers_cfcomp (CFCOMP a, double* in, double* out)`
  Re-points the `cfcomp` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_cfcomp()`** — L613 — `void setSamplerate_cfcomp (CFCOMP a, int rate)`
  Reconfigures the `cfcomp` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_cfcomp()`** — L620 — `void setSize_cfcomp (CFCOMP a, int size)`
  Reconfigures the `cfcomp` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXACFCOMPRun()`** — L633 — `PORT void SetTXACFCOMPRun (int channel, int run)`
  Sets txacfcomprun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFCOMPPosition()`** — L645 — `PORT void SetTXACFCOMPPosition (int channel, int pos)`
  Sets txacfcompposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFCOMPprofile()`** — L657 — `PORT void SetTXACFCOMPprofile (int channel, int nfreqs, double* F, double* G, double *E, double *Qg, double *Qe)`
  Sets txacfcompprofile — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFCOMPPrecomp()`** — L702 — `PORT void SetTXACFCOMPPrecomp (int channel, double precomp)`
  Sets txacfcompprecomp — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFCOMPPeqRun()`** — L719 — `PORT void SetTXACFCOMPPeqRun (int channel, int run)`
  Sets txacfcomppeq run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACFCOMPPrePeq()`** — L731 — `PORT void SetTXACFCOMPPrePeq (int channel, double prepeq)`
  Sets txacfcomppre peq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetTXACFCOMPDisplayCompression()`** — L741 — `PORT void GetTXACFCOMPDisplayCompression (int channel, double* comp_values, int* ready)`
  Returns txacfcompdisplay compression — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/cfcomp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
