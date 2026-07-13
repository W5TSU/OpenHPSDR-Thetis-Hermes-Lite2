# `wdsp/nob.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Impulse noise blankers (NB and NB2).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/znob.c` (calls ×8)
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_anb()` (×2), `destroy_anb()` (×2), `xanb()` (×2), `pSetRCVRANBRun()` (×1), `pSetRCVRANBBuffsize()` (×1), `pSetRCVRANBSamplerate()` (×1), `pSetRCVRANBTau()` (×1), `pSetRCVRANBHangtime()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`initBlanker()`** — L33 — `void initBlanker(ANB a)`
  Called by: `create_anb()` (same file), `flush_anb()` (same file), `setSamplerate_anb()` (same file), `setSize_anb()` (same file), `pSetRCVRANBSamplerate()` (same file), `pSetRCVRANBTau()` (same file) — and 8 more
- **`create_anb()`** — L54 — `PORT ANB create_anb ( int run, int buffsize, double* in,`
  Constructor for the `anb` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_anbEXT()` (same file), `create_rcvr()` (`ChannelMaster/cmaster.c`), `create_spc0()` (`ChannelMaster/pipe.c`)
- **`destroy_anb()`** — L89 — `PORT void destroy_anb (ANB a)`
  Destroys the `anb` block, freeing its allocated buffers.
  Called by: `destroy_anbEXT()` (same file), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_spc0()` (`ChannelMaster/pipe.c`)
- **`flush_anb()`** — L99 — `PORT void flush_anb (ANB a)`
  Flushes (zeroes) the `anb` block’s internal buffers/state.
  Called by: `flush_anbEXT()` (same file)
- **`xanb()`** — L107 — `PORT void xanb (ANB a)`
  Runs the `anb` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xanbEXT()` (same file), `xanbEXTF()` (same file), `xcmaster()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`setBuffers_anb()`** — L189 — `void setBuffers_anb (ANB a, double* in, double* out)`
  Re-points the `anb` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_anb()`** — L195 — `void setSamplerate_anb (ANB a, int rate)`
  Reconfigures the `anb` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_anb()`** — L201 — `void setSize_anb (ANB a, int size)`
  Reconfigures the `anb` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetRCVRANBRun()`** — L232 — `PORT void pSetRCVRANBRun (ANB a, int run)`
  Called by: `SetRCVRANBRun()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBBuffsize()`** — L240 — `PORT void pSetRCVRANBBuffsize (ANB a, int size)`
  Called by: `SetRCVRANBBuffsize()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBSamplerate()`** — L248 — `PORT void pSetRCVRANBSamplerate (ANB a, int rate)`
  Called by: `SetRCVRANBSamplerate()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBTau()`** — L257 — `PORT void pSetRCVRANBTau (ANB a, double tau)`
  Called by: `SetRCVRANBTau()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBHangtime()`** — L266 — `PORT void pSetRCVRANBHangtime (ANB a, double time)`
  Called by: `SetRCVRANBHangtime()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBAdvtime()`** — L275 — `PORT void pSetRCVRANBAdvtime (ANB a, double time)`
  Called by: `SetRCVRANBAdvtime()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBBacktau()`** — L284 — `PORT void pSetRCVRANBBacktau (ANB a, double tau)`
  Called by: `SetRCVRANBBacktau()` (`ChannelMaster/znob.c`)
- **`pSetRCVRANBThreshold()`** — L293 — `PORT void pSetRCVRANBThreshold (ANB a, double thresh)`
  Called by: `SetRCVRANBThreshold()` (`ChannelMaster/znob.c`)
- **`create_anbEXT()`** — L310 — `PORT void create_anbEXT ( int id, int run, int buffsize,`
  Constructor for the `anbEXT` block: allocates its state/buffers and computes initial coefficients.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`destroy_anbEXT()`** — L326 — `PORT void destroy_anbEXT (int id)`
  Destroys the `anbEXT` block, freeing its allocated buffers.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`flush_anbEXT()`** — L332 — `PORT void flush_anbEXT (int id)`
  Flushes (zeroes) the `anbEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xanbEXT()`** — L338 — `PORT void xanbEXT (int id, double* in, double* out)`
  Runs the `anbEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEXTANBRun()`** — L347 — `PORT void SetEXTANBRun (int id, int run)`
  Sets extanbrun — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEXTANBBuffsize()`** — L356 — `PORT void SetEXTANBBuffsize (int id, int size)`
  Sets extanbbuffsize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBSamplerate()`** — L365 — `PORT void SetEXTANBSamplerate (int id, int rate)`
  Sets extanbsamplerate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBTau()`** — L375 — `PORT void SetEXTANBTau (int id, double tau)`
  Sets extanbtau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBHangtime()`** — L385 — `PORT void SetEXTANBHangtime (int id, double time)`
  Sets extanbhangtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBAdvtime()`** — L395 — `PORT void SetEXTANBAdvtime (int id, double time)`
  Sets extanbadvtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBBacktau()`** — L405 — `PORT void SetEXTANBBacktau (int id, double tau)`
  Sets extanbbacktau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`SetEXTANBThreshold()`** — L415 — `PORT void SetEXTANBThreshold (int id, double thresh)`
  Sets extanbthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.
- **`xanbEXTF()`** — L430 — `PORT void xanbEXTF (int id, float *I, float *Q)`
  Runs the `anbEXTF` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called from C# via P/Invoke — declared/wrapped in `Console/HPSDR/specHPSDR.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nob.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
