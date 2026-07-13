# `wdsp/patchpanel.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_panel()` (×2), `destroy_panel()` (×2), `flush_panel()` (×2), `xpanel()` (×2), `setSamplerate_panel()` (×2), `setBuffers_panel()` (×2), `setSize_panel()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_panel()`** — L29 — `PANEL create_panel (int channel, int run, int size, double* in, double* out, double gain1, double gain2I, double gain2Q, int inselect, int copy)`
  Constructor for the `panel` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_panel()`** — L45 — `void destroy_panel (PANEL a)`
  Destroys the `panel` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_panel()`** — L50 — `void flush_panel (PANEL a)`
  Flushes (zeroes) the `panel` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xpanel()`** — L55 — `void xpanel (PANEL a)`
  Runs the `panel` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_panel()`** — L103 — `void setBuffers_panel (PANEL a, double* in, double* out)`
  Re-points the `panel` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_panel()`** — L109 — `void setSamplerate_panel (PANEL a, int rate)`
  Reconfigures the `panel` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_panel()`** — L114 — `void setSize_panel (PANEL a, int size)`
  Reconfigures the `panel` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetRXAPanelRun()`** — L125 — `PORT void SetRXAPanelRun (int channel, int run)`
  Sets rxapanel run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPanelSelect()`** — L133 — `PORT void SetRXAPanelSelect (int channel, int select)`
  Sets rxapanel select — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPanelGain1()`** — L141 — `PORT void SetRXAPanelGain1 (int channel, double gain)`
  Sets rxapanel gain1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPanelGain2()`** — L149 — `PORT void SetRXAPanelGain2 (int channel, double gainI, double gainQ)`
  Sets rxapanel gain2 — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAPanelPan()`** — L158 — `PORT void SetRXAPanelPan (int channel, double pan)`
  Sets rxapanel pan — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAPanelCopy()`** — L178 — `PORT void SetRXAPanelCopy (int channel, int copy)`
  Sets rxapanel copy — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAPanelBinaural()`** — L186 — `PORT void SetRXAPanelBinaural (int channel, int bin)`
  Sets rxapanel binaural — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPanelRun()`** — L200 — `PORT void SetTXAPanelRun (int channel, int run)`
  Sets txapanel run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPanelGain1()`** — L208 — `PORT void SetTXAPanelGain1 (int channel, double gain)`
  Sets txapanel gain1 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAPanelSelect()`** — L217 — `PORT void SetTXAPanelSelect (int channel, int select)`
  Sets txapanel select — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/patchpanel.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
