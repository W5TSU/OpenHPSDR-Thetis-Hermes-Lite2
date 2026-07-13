# `wdsp/div.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Diversity combiner (mixes two receivers with adjustable gain/phase).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/sync.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_divEXT()` (×1), `destroy_divEXT()` (×1), `xdivEXT()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_div()`** — L31 — `MDIV create_div (int run, int nr, int size, double **in, double *out)`
  Constructor for the `div` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_divEXT()` (same file)
- **`destroy_div()`** — L50 — `void destroy_div (MDIV a)`
  Destroys the `div` block, freeing its allocated buffers.
  Called by: `destroy_divEXT()` (same file)
- **`flush_div()`** — L62 — `void flush_div (MDIV a)`
  Flushes (zeroes) the `div` block’s internal buffers/state.
  Called by: `flush_divEXT()` (same file)
- **`xdiv()`** — L67 — `void xdiv (MDIV a)`
  Runs the `div` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xdivEXT()` (same file), `xdivEXTF()` (same file)
- **`create_divEXT()`** — L107 — `PORT void create_divEXT (int id, int run, int nr, int size)`
  Constructor for the `divEXT` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_sync()` (`ChannelMaster/sync.c`)
- **`destroy_divEXT()`** — L113 — `PORT void destroy_divEXT (int id)`
  Destroys the `divEXT` block, freeing its allocated buffers.
  Called by: `destroy_sync()` (`ChannelMaster/sync.c`)
- **`flush_divEXT()`** — L119 — `PORT void flush_divEXT (int id)`
  Flushes (zeroes) the `divEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xdivEXT()`** — L125 — `PORT void xdivEXT (int id, int nsamples, double **in, double *out)`
  Runs the `divEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `InboundBlock()` (`ChannelMaster/sync.c`)
- **`SetEXTDIVRun()`** — L137 — `PORT void SetEXTDIVRun (int id, int run)`
  0 - does nothing; 1 - operates
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetEXTDIVBuffsize()`** — L147 — `PORT void SetEXTDIVBuffsize (int id, int size)`
  size of data buffer in complex samples
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEXTDIVNr()`** — L157 — `PORT void SetEXTDIVNr (int id, int nr)`
  number of receivers being used for diversity
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetEXTDIVOutput()`** — L168 — `PORT void SetEXTDIVOutput (int id, int output)`
  number of which receiver to output if output==nr, mixing occurs
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetEXTDIVRotate()`** — L179 — `PORT void SetEXTDIVRotate (int id, int nr, double *Irotate, double *Qrotate)`
  I and Q "rotate" multipliers for each receiver can be set to 1.0 and 0.0 for "reference receiver"
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`xdivEXTF()`** — L194 — `PORT void xdivEXTF (int id, int size, float **input, float *Iout, float *Qout)`
  LEGACY INTERFACE - REMOVE *
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/div.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
