# `wdsp/fmsq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×7)
  - `wdsp/eq.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setBuffers_fmsq()` (×2), `create_fmsq()` (×1), `destroy_fmsq()` (×1), `flush_fmsq()` (×1), `xfmsq()` (×1), `setSamplerate_fmsq()` (×1), `setSize_fmsq()` (×1), `SetRXAFMSQNC()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_fmsq()`** — L29 — `void calc_fmsq (FMSQ a)`
  Called by: `create_fmsq()` (same file), `setSamplerate_fmsq()` (same file), `setSize_fmsq()` (same file)
- **`decalc_fmsq()`** — L80 — `void decalc_fmsq (FMSQ a)`
  Called by: `destroy_fmsq()` (same file), `setSamplerate_fmsq()` (same file), `setSize_fmsq()` (same file)
- **`create_fmsq()`** — L88 — `FMSQ create_fmsq (int run, int size, double* insig, double* outsig, double* trigger, int rate, double fc, double* pllpole, double tdelay, double avtau, double longtau, double tup, `
  Constructor for the `fmsq` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_fmsq()`** — L116 — `void destroy_fmsq (FMSQ a)`
  Destroys the `fmsq` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_fmsq()`** — L122 — `void flush_fmsq (FMSQ a)`
  Flushes (zeroes) the `fmsq` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xfmsq()`** — L141 — `void xfmsq (FMSQ a)`
  Runs the `fmsq` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_fmsq()`** — L207 — `void setBuffers_fmsq (FMSQ a, double* in, double* out, double* trig)`
  Re-points the `fmsq` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_fmsq()`** — L215 — `void setSamplerate_fmsq (FMSQ a, int rate)`
  Reconfigures the `fmsq` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_fmsq()`** — L222 — `void setSize_fmsq (FMSQ a, int size)`
  Reconfigures the `fmsq` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAFMSQRun()`** — L235 — `PORT void SetRXAFMSQRun (int channel, int run)`
  Sets rxafmsqrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAFMSQThreshold()`** — L243 — `PORT void SetRXAFMSQThreshold (int channel, double threshold)`
  Sets rxafmsqthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAFMSQNC()`** — L252 — `PORT void SetRXAFMSQNC (int channel, int nc)`
  Sets rxafmsqnc — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`SetRXAFMSQMP()`** — L269 — `PORT void SetRXAFMSQMP (int channel, int mp)`
  Sets rxafmsqmp — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmsq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
