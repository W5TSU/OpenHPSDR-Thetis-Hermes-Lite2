# `wdsp/nobII.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Impulse noise blankers (NB and NB2).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/znobII.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_nob()` (×2), `destroy_nob()` (×2), `xnob()` (×2), `pSetRCVRNOBRun()` (×1), `pSetRCVRNOBMode()` (×1), `pSetRCVRNOBBuffsize()` (×1), `pSetRCVRNOBSamplerate()` (×1), `pSetRCVRNOBTau()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`init_nob()`** — L36 — `void init_nob (NOB a)`
  Called by: `create_nob()` (same file), `setSamplerate_nob()` (same file), `pSetRCVRNOBSamplerate()` (same file), `pSetRCVRNOBTau()` (same file), `pSetRCVRNOBHangtime()` (same file), `pSetRCVRNOBAdvtime()` (same file) — and 6 more
- **`create_nob()`** — L63 — `PORT NOB create_nob ( int run, int buffsize, double* in,`
  Constructor for the `nob` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_nobEXT()` (same file), `create_rcvr()` (`ChannelMaster/cmaster.c`), `create_spc0()` (`ChannelMaster/pipe.c`)
- **`destroy_nob()`** — L126 — `PORT void destroy_nob (NOB a)`
  Destroys the `nob` block, freeing its allocated buffers.
  Called by: `destroy_nobEXT()` (same file), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_spc0()` (`ChannelMaster/pipe.c`)
- **`flush_nob()`** — L140 — `PORT void flush_nob (NOB a)`
  Flushes (zeroes) the `nob` block’s internal buffers/state.
  Called by: `init_nob()` (same file), `setSize_nob()` (same file), `flush_nobEXT()` (same file)
- **`xnob()`** — L157 — `PORT void xnob (NOB a)`
  Runs the `nob` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xnobEXT()` (same file), `xnobEXTF()` (same file), `xcmaster()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`setBuffers_nob()`** — L496 — `void setBuffers_nob (NOB a, double* in, double* out)`
  Re-points the `nob` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_nob()`** — L502 — `void setSamplerate_nob (NOB a, int rate)`
  Reconfigures the `nob` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_nob()`** — L508 — `void setSize_nob (NOB a, int size)`
  Reconfigures the `nob` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetRCVRNOBRun()`** — L520 — `PORT void pSetRCVRNOBRun (NOB a, int run)`
  Called by: `SetRCVRNOBRun()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBMode()`** — L528 — `PORT void pSetRCVRNOBMode (NOB a, int mode)`
  Called by: `SetRCVRNOBMode()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBBuffsize()`** — L536 — `PORT void pSetRCVRNOBBuffsize (NOB a, int size)`
  Called by: `SetRCVRNOBBuffsize()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBSamplerate()`** — L544 — `PORT void pSetRCVRNOBSamplerate (NOB a, int rate)`
  Called by: `SetRCVRNOBSamplerate()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBTau()`** — L553 — `PORT void pSetRCVRNOBTau (NOB a, double tau)`
  Called by: `SetRCVRNOBTau()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBHangtime()`** — L563 — `PORT void pSetRCVRNOBHangtime (NOB a, double time)`
  Called by: `SetRCVRNOBHangtime()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBAdvtime()`** — L572 — `PORT void pSetRCVRNOBAdvtime (NOB a, double time)`
  Called by: `SetRCVRNOBAdvtime()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBBacktau()`** — L581 — `PORT void pSetRCVRNOBBacktau (NOB a, double tau)`
  Called by: `SetRCVRNOBBacktau()` (`ChannelMaster/znobII.c`)
- **`pSetRCVRNOBThreshold()`** — L590 — `PORT void pSetRCVRNOBThreshold (NOB a, double thresh)`
  Called by: `SetRCVRNOBThreshold()` (`ChannelMaster/znobII.c`)
- **`create_nobEXT()`** — L607 — `PORT void create_nobEXT ( int id, int run, int mode,`
  Constructor for the `nobEXT` block: allocates its state/buffers and computes initial coefficients.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`destroy_nobEXT()`** — L627 — `PORT void destroy_nobEXT (int id)`
  Destroys the `nobEXT` block, freeing its allocated buffers.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`flush_nobEXT()`** — L633 — `PORT void flush_nobEXT (int id)`
  Flushes (zeroes) the `nobEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xnobEXT()`** — L639 — `PORT void xnobEXT (int id, double* in, double* out)`
  Runs the `nobEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEXTNOBRun()`** — L648 — `PORT void SetEXTNOBRun (int id, int run)`
  Sets extnobrun — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEXTNOBMode()`** — L657 — `PORT void SetEXTNOBMode (int id, int mode)`
  Sets extnobmode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBBuffsize()`** — L666 — `PORT void SetEXTNOBBuffsize (int id, int size)`
  Sets extnobbuffsize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBSamplerate()`** — L675 — `PORT void SetEXTNOBSamplerate (int id, int rate)`
  Sets extnobsamplerate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBTau()`** — L685 — `PORT void SetEXTNOBTau (int id, double tau)`
  Sets extnobtau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBHangtime()`** — L696 — `PORT void SetEXTNOBHangtime (int id, double time)`
  Sets extnobhangtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBAdvtime()`** — L706 — `PORT void SetEXTNOBAdvtime (int id, double time)`
  Sets extnobadvtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBBacktau()`** — L716 — `PORT void SetEXTNOBBacktau (int id, double tau)`
  Sets extnobbacktau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTNOBThreshold()`** — L726 — `PORT void SetEXTNOBThreshold (int id, double thresh)`
  Sets extnobthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`xnobEXTF()`** — L741 — `PORT void xnobEXTF (int id, float *I, float *Q)`
  Runs the `nobEXTF` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nobII.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
