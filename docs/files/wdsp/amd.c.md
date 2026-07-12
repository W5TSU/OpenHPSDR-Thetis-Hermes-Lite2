# `wdsp/amd.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM/SAM (synchronous) and FM demodulators.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_amd()` (×1), `destroy_amd()` (×1), `flush_amd()` (×1), `xamd()` (×1), `setSamplerate_amd()` (×1), `setBuffers_amd()` (×1), `setSize_amd()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_amd()`** — L29 — `AMD create_amd ( int run, int buff_size, double *in_buff,`
  Constructor for the `amd` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_amd()`** — L67 — `void destroy_amd(AMD a)`
  Destroys the `amd` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`init_amd()`** — L72 — `void init_amd(AMD a)`
  Called by: `create_amd()` (same file), `setSamplerate_amd()` (same file)
- **`flush_amd()`** — L109 — `void flush_amd (AMD a)`
  Flushes (zeroes) the `amd` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xamd()`** — L115 — `void xamd (AMD a)`
  Runs the `amd` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_amd()`** — L241 — `void setBuffers_amd (AMD a, double* in, double* out)`
  Re-points the `amd` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_amd()`** — L247 — `void setSamplerate_amd (AMD a, int rate)`
  Reconfigures the `amd` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_amd()`** — L253 — `void setSize_amd (AMD a, int size)`
  Reconfigures the `amd` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAAMDRun()`** — L264 — `PORT void SetRXAAMDRun(int channel, int run)`
  Sets rxaamdrun — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAAMDSBMode()`** — L281 — `PORT void SetRXAAMDSBMode(int channel, int sbmode)`
  Sets rxaamdsbmode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAMDFadeLevel()`** — L289 — `PORT void SetRXAAMDFadeLevel(int channel, int levelfade)`
  Sets rxaamdfade level — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/amd.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
