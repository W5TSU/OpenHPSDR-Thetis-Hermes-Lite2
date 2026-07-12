# `wdsp/ammod.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM and FM modulators for TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_ammod()` (×1), `destroy_ammod()` (×1), `flush_ammod()` (×1), `xammod()` (×1), `setSamplerate_ammod()` (×1), `setBuffers_ammod()` (×1), `setSize_ammod()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_ammod()`** — L29 — `AMMOD create_ammod (int run, int mode, int size, double* in, double* out, double c_level)`
  Constructor for the `ammod` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_ammod()`** — L43 — `void destroy_ammod (AMMOD a)`
  Destroys the `ammod` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_ammod()`** — L48 — `void flush_ammod (AMMOD a)`
  Flushes (zeroes) the `ammod` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xammod()`** — L53 — `void xammod (AMMOD a)`
  Runs the `ammod` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_ammod()`** — L81 — `void setBuffers_ammod (AMMOD a, double* in, double* out)`
  Re-points the `ammod` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_ammod()`** — L87 — `void setSamplerate_ammod (AMMOD a, int rate)`
  Reconfigures the `ammod` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_ammod()`** — L92 — `void setSize_ammod (AMMOD a, int size)`
  Reconfigures the `ammod` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAAMCarrierLevel()`** — L103 — `PORT void SetTXAAMCarrierLevel (int channel, double c_level)`
  Sets txaamcarrier level — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/ammod.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
