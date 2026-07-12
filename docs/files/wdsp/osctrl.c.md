# `wdsp/osctrl.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** TX overshoot control.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/TXA.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_osctrl()` (×1), `destroy_osctrl()` (×1), `flush_osctrl()` (×1), `xosctrl()` (×1), `setSamplerate_osctrl()` (×1), `setBuffers_osctrl()` (×1), `setSize_osctrl()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_osctrl()`** — L33 — `void calc_osctrl (OSCTRL a)`
  Called by: `create_osctrl()` (same file), `setSamplerate_osctrl()` (same file)
- **`decalc_osctrl()`** — L46 — `void decalc_osctrl (OSCTRL a)`
  Called by: `destroy_osctrl()` (same file), `setSamplerate_osctrl()` (same file)
- **`create_osctrl()`** — L52 — `OSCTRL create_osctrl ( int run, int size, double* inbuff, double* outbuff,`
  Constructor for the `osctrl` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_osctrl()`** — L72 — `void destroy_osctrl (OSCTRL a)`
  Destroys the `osctrl` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_osctrl()`** — L78 — `void flush_osctrl (OSCTRL a)`
  Flushes (zeroes) the `osctrl` block’s internal buffers/state.
  Called by: `setSize_osctrl()` (same file), `flush_txa()` (`wdsp/TXA.c`)
- **`xosctrl()`** — L84 — `void xosctrl (OSCTRL a)`
  Runs the `osctrl` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_osctrl()`** — L116 — `void setBuffers_osctrl (OSCTRL a, double* in, double* out)`
  Re-points the `osctrl` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_osctrl()`** — L122 — `void setSamplerate_osctrl (OSCTRL a, int rate)`
  Reconfigures the `osctrl` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_osctrl()`** — L129 — `void setSize_osctrl (OSCTRL a, int size)`
  Reconfigures the `osctrl` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAosctrlRun()`** — L141 — `PORT void SetTXAosctrlRun (int channel, int run)`
  Sets txaosctrl run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/osctrl.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
