# `wdsp/slew.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/TXA.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_uslew()` (×1), `destroy_uslew()` (×1), `flush_uslew()` (×1), `xuslew()` (×1), `setSamplerate_uslew()` (×1), `setBuffers_uslew()` (×1), `setSize_uslew()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_uslew()`** — L37 — `void calc_uslew (USLEW a)`
  Called by: `create_uslew()` (same file), `setSamplerate_uslew()` (same file), `SetTXAuSlewTime()` (same file)
- **`decalc_uslew()`** — L57 — `void decalc_uslew (USLEW a)`
  Called by: `destroy_uslew()` (same file), `setSamplerate_uslew()` (same file), `SetTXAuSlewTime()` (same file)
- **`create_uslew()`** — L62 — `USLEW create_uslew (int channel, volatile long *ch_upslew, int size, double* in, double* out, double rate, double tdelay, double tupslew)`
  Constructor for the `uslew` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_uslew()`** — L77 — `void destroy_uslew (USLEW a)`
  Destroys the `uslew` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_uslew()`** — L83 — `void flush_uslew (USLEW a)`
  Flushes (zeroes) the `uslew` block’s internal buffers/state.
  Called by: `setSize_uslew()` (same file), `flush_txa()` (`wdsp/TXA.c`)
- **`xuslew()`** — L90 — `void xuslew (USLEW a)`
  Runs the `uslew` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_uslew()`** — L157 — `void setBuffers_uslew (USLEW a, double* in, double* out)`
  Re-points the `uslew` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_uslew()`** — L163 — `void setSamplerate_uslew (USLEW a, int rate)`
  Reconfigures the `uslew` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_uslew()`** — L170 — `void setSize_uslew (USLEW a, int size)`
  Reconfigures the `uslew` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAuSlewTime()`** — L182 — `PORT void SetTXAuSlewTime(int channel, double time)`
  Sets txau slew time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/slew.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
