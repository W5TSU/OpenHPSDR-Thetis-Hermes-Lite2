# `wdsp/shift.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×11)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `setBuffers_shift()` (×3), `setSize_shift()` (×3), `create_shift()` (×1), `destroy_shift()` (×1), `flush_shift()` (×1), `xshift()` (×1), `setSamplerate_shift()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_shift()`** — L29 — `void calc_shift (SHIFT a)`
  Called by: `create_shift()` (same file), `setSamplerate_shift()` (same file), `SetRXAShiftFreq()` (same file)
- **`create_shift()`** — L36 — `SHIFT create_shift (int run, int size, double* in, double* out, int rate, double fshift)`
  Constructor for the `shift` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_shift()`** — L50 — `void destroy_shift (SHIFT a)`
  Destroys the `shift` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_shift()`** — L55 — `void flush_shift (SHIFT a)`
  Flushes (zeroes) the `shift` block’s internal buffers/state.
  Called by: `setSize_shift()` (same file), `flush_rxa()` (`wdsp/RXA.c`)
- **`xshift()`** — L60 — `void xshift (SHIFT a)`
  Runs the `shift` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_shift()`** — L87 — `void setBuffers_shift(SHIFT a, double* in, double* out)`
  Re-points the `shift` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_shift()`** — L93 — `void setSamplerate_shift (SHIFT a, int rate)`
  Reconfigures the `shift` block for a new sample rate.
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_shift()`** — L100 — `void setSize_shift (SHIFT a, int size)`
  Reconfigures the `shift` block for a new buffer size.
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAShiftRun()`** — L112 — `PORT void SetRXAShiftRun (int channel, int run)`
  Sets rxashift run — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAShiftFreq()`** — L120 — `PORT void SetRXAShiftFreq (int channel, double fshift)`
  Sets rxashift freq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/shift.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
