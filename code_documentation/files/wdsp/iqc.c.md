# `wdsp/iqc.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** PureSignal calibration calculation and the I/Q correction applied to TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/calcc.c` (calls ×10)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetTXAiqcStart()` (×2), `SetTXAiqcSwap()` (×2), `create_iqc()` (×1), `destroy_iqc()` (×1), `flush_iqc()` (×1), `xiqc()` (×1), `setSamplerate_iqc()` (×1), `setBuffers_iqc()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`size_iqc()`** — L29 — `void size_iqc (IQC a)`
  Called by: `calc_iqc()` (same file), `SetPSIntsAndSpi()` (`wdsp/calcc.c`)
- **`desize_iqc()`** — L46 — `void desize_iqc (IQC a)`
  Called by: `decalc_iqc()` (same file), `SetPSIntsAndSpi()` (`wdsp/calcc.c`)
- **`calc_iqc()`** — L59 — `void calc_iqc (IQC a)`
  Called by: `create_iqc()` (same file), `setSamplerate_iqc()` (same file)
- **`decalc_iqc()`** — L80 — `void decalc_iqc (IQC a)`
  Called by: `destroy_iqc()` (same file), `setSamplerate_iqc()` (same file)
- **`create_iqc()`** — L87 — `IQC create_iqc (int run, int size, double* in, double* out, double rate, int ints, double tup, int spi)`
  Constructor for the `iqc` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_iqc()`** — L102 — `void destroy_iqc (IQC a)`
  Destroys the `iqc` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_iqc()`** — L108 — `void flush_iqc (IQC a)`
  Flushes (zeroes) the `iqc` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xiqc()`** — L122 — `void xiqc (IQC a)`
  Runs the `iqc` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_iqc()`** — L205 — `void setBuffers_iqc (IQC a, double* in, double* out)`
  Re-points the `iqc` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_iqc()`** — L211 — `void setSamplerate_iqc (IQC a, int rate)`
  Reconfigures the `iqc` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_iqc()`** — L218 — `void setSize_iqc (IQC a, int size)`
  Reconfigures the `iqc` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`GetTXAiqcValues()`** — L229 — `PORT void GetTXAiqcValues (int channel, double* cm, double* cc, double* cs)`
  Returns txaiqc values — API getter, typically called from the console via P/Invoke.
  Called by: `PSSaveCorrection()` (`wdsp/calcc.c`)
- **`SetTXAiqcValues()`** — L241 — `PORT void SetTXAiqcValues (int channel, double* cm, double* cc, double* cs)`
  Sets txaiqc values — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAiqcSwap()`** — L255 — `PORT void SetTXAiqcSwap (int channel, double* cm, double* cc, double* cs)`
  Sets txaiqc swap — API setter, typically called from the console via P/Invoke.
  Called by: `doPSCalcCorrection()` (`wdsp/calcc.c`), `PSRestoreCorrection()` (`wdsp/calcc.c`)
- **`SetTXAiqcStart()`** — L271 — `PORT void SetTXAiqcStart (int channel, double* cm, double* cc, double* cs)`
  Sets txaiqc start — API setter, typically called from the console via P/Invoke.
  Called by: `doPSCalcCorrection()` (`wdsp/calcc.c`), `PSRestoreCorrection()` (`wdsp/calcc.c`)
- **`SetTXAiqcEnd()`** — L288 — `PORT void SetTXAiqcEnd (int channel)`
  Sets txaiqc end — API setter, typically called from the console via P/Invoke.
  Called by: `doPSTurnoff()` (`wdsp/calcc.c`)
- **`GetTXAiqcDogCount()`** — L301 — `void GetTXAiqcDogCount (int channel, int* count)`
  Returns txaiqc dog count — API getter, typically called from the console via P/Invoke.
  Called by: `pscc()` (`wdsp/calcc.c`)
- **`SetTXAiqcDogCount()`** — L309 — `void SetTXAiqcDogCount (int channel, int count)`
  Sets txaiqc dog count — API setter, typically called from the console via P/Invoke.
  Called by: `pscc()` (`wdsp/calcc.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iqc.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
