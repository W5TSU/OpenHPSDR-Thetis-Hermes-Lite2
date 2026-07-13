# `wdsp/amsq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×8)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_amsq()` (×2), `destroy_amsq()` (×2), `flush_amsq()` (×2), `xamsq()` (×2), `xamsqcap()` (×2), `setSamplerate_amsq()` (×2), `setBuffers_amsq()` (×2), `setSize_amsq()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`compute_slews()`** — L29 — `void compute_slews(AMSQ a)`
  Called by: `calc_amsq()` (same file), `SetTXAAMSQMutedGain()` (same file)
- **`calc_amsq()`** — L49 — `void calc_amsq(AMSQ a)`
  Called by: `create_amsq()` (same file), `setSamplerate_amsq()` (same file), `setSize_amsq()` (same file)
- **`decalc_amsq()`** — L66 — `void decalc_amsq (AMSQ a)`
  Called by: `destroy_amsq()` (same file), `setSamplerate_amsq()` (same file), `setSize_amsq()` (same file)
- **`create_amsq()`** — L73 — `AMSQ create_amsq (int run, int size, double* in, double* out, double* trigger, int rate, double avtau, double tup, double tdown, double tail_thresh, double unmute_thresh, double mi`
  Constructor for the `amsq` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_amsq()`** — L95 — `void destroy_amsq (AMSQ a)`
  Destroys the `amsq` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_amsq()`** — L101 — `void flush_amsq (AMSQ a)`
  Flushes (zeroes) the `amsq` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xamsq()`** — L117 — `void xamsq (AMSQ a)`
  Runs the `amsq` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`xamsqcap()`** — L178 — `void xamsqcap (AMSQ a)`
  Runs the `amsqcap` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_amsq()`** — L183 — `void setBuffers_amsq (AMSQ a, double* in, double* out, double* trigger)`
  Re-points the `amsq` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_amsq()`** — L190 — `void setSamplerate_amsq (AMSQ a, int rate)`
  Reconfigures the `amsq` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_amsq()`** — L197 — `void setSize_amsq (AMSQ a, int size)`
  Reconfigures the `amsq` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetRXAAMSQRun()`** — L210 — `PORT void SetRXAAMSQRun (int channel, int run)`
  Sets rxaamsqrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAMSQThreshold()`** — L218 — `PORT void SetRXAAMSQThreshold (int channel, double threshold)`
  Sets rxaamsqthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAMSQMaxTail()`** — L228 — `PORT void SetRXAAMSQMaxTail (int channel, double tail)`
  Sets rxaamsqmax tail — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAAMSQRun()`** — L245 — `PORT void SetTXAAMSQRun (int channel, int run)`
  Sets txaamsqrun — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAAMSQMutedGain()`** — L253 — `PORT void SetTXAAMSQMutedGain (int channel, double dBlevel)`
  Sets txaamsqmuted gain — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAAMSQThreshold()`** — L264 — `PORT void SetTXAAMSQThreshold (int channel, double threshold)`
  Sets txaamsqthreshold — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/amsq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
