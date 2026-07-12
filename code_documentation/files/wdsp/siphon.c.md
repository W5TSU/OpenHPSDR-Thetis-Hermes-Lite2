# `wdsp/siphon.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Taps TX samples out of the chain (e.g., for the TX display).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
  - `ChannelMaster/pipe.c` (calls ×3)
  - `ChannelMaster/analyzers.c` (calls ×1)
  - `ChannelMaster/cmaster.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/analyzer.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_siphon()` (×2), `destroy_siphon()` (×2), `flush_siphon()` (×2), `xsiphon()` (×2), `setSamplerate_siphon()` (×2), `setBuffers_siphon()` (×2), `setSize_siphon()` (×2), `TXASetSipAllocDisps()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`build_window()`** — L29 — `void build_window (SIPHON a)`
  Called by: `create_siphon()` (same file)
- **`create_siphon()`** — L53 — `SIPHON create_siphon (int run, int position, int mode, int disp, int insize, double* in, int sipsize, int fftsize, int specmode)`
  Constructor for the `siphon` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_siphonEXT()` (same file), `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_siphon()`** — L80 — `void destroy_siphon (SIPHON a)`
  Destroys the `siphon` block, freeing its allocated buffers.
  Called by: `destroy_siphonEXT()` (same file), `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_siphon()`** — L93 — `void flush_siphon (SIPHON a)`
  Flushes (zeroes) the `siphon` block’s internal buffers/state.
  Called by: `setSamplerate_siphon()` (same file), `setSize_siphon()` (same file), `flush_siphonEXT()` (same file), `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xsiphon()`** — L101 — `void xsiphon (SIPHON a, int pos)`
  Runs the `siphon` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xsiphonEXT()` (same file), `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_siphon()`** — L139 — `void setBuffers_siphon (SIPHON a, double* in)`
  Re-points the `siphon` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_siphon()`** — L144 — `void setSamplerate_siphon (SIPHON a, int rate)`
  Reconfigures the `siphon` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_siphon()`** — L149 — `void setSize_siphon (SIPHON a, int size)`
  Reconfigures the `siphon` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`suck()`** — L155 — `void suck (SIPHON a)`
  Called by: `RXAGetaSipF()` (same file), `RXAGetaSipF1()` (same file), `TXAGetaSipF()` (same file), `TXAGetaSipF1()` (same file), `TXAGetSpecF1()` (same file), `GetaSipF1EXT()` (same file)
- **`sip_spectrum()`** — L172 — `void sip_spectrum (SIPHON a)`
  Called by: `TXAGetSpecF1()` (same file)
- **`RXAGetaSipF()`** — L189 — `PORT void RXAGetaSipF (int channel, float* out, int size)`
  RXA chain operation — geta sip f; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXAGetaSipF1()`** — L204 — `PORT void RXAGetaSipF1 (int channel, float* out, int size)`
  RXA chain operation — geta sip f1; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetSipPosition()`** — L226 — `PORT void TXASetSipPosition (int channel, int pos)`
  TXA chain operation — set sip position; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetSipMode()`** — L235 — `PORT void TXASetSipMode (int channel, int mode)`
  TXA chain operation — set sip mode; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetSipDisplay()`** — L244 — `PORT void TXASetSipDisplay (int channel, int disp)`
  TXA chain operation — set sip display; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXAGetaSipF()`** — L253 — `PORT void TXAGetaSipF (int channel, float* out, int size)`
  TXA chain operation — geta sip f; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXAGetaSipF1()`** — L268 — `PORT void TXAGetaSipF1 (int channel, float* out, int size)`
  TXA chain operation — geta sip f1; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetSipSpecmode()`** — L284 — `PORT void TXASetSipSpecmode (int channel, int mode)`
  TXA chain operation — set sip specmode; part of the receive/transmit chain API.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`TXAGetSpecF1()`** — L294 — `PORT void TXAGetSpecF1 (int channel, float* out)`
  TXA chain operation — get spec f1; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetSipAllocDisps()`** — L321 — `PORT void TXASetSipAllocDisps (int channel, int n_alloc_disps, int* alloc_run, int* alloc_disp)`
  TXA chain operation — set sip alloc disps; part of the receive/transmit chain API.
  Called by: `tx_analyzers()` (`ChannelMaster/analyzers.c`)
- **`create_siphonEXT()`** — L346 — `PORT void create_siphonEXT (int id, int run, int insize, int sipsize, int fftsize, int specmode)`
  Constructor for the `siphonEXT` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_pipe()` (`ChannelMaster/pipe.c`)
- **`destroy_siphonEXT()`** — L352 — `PORT void destroy_siphonEXT (int id)`
  Destroys the `siphonEXT` block, freeing its allocated buffers.
  Called by: `destroy_pipe()` (`ChannelMaster/pipe.c`)
- **`flush_siphonEXT()`** — L358 — `PORT void flush_siphonEXT (int id)`
  Flushes (zeroes) the `siphonEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xsiphonEXT()`** — L364 — `PORT void xsiphonEXT (int id, double* buff)`
  Runs the `siphonEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`GetaSipF1EXT()`** — L372 — `PORT void GetaSipF1EXT (int id, float* out, int size)`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetSiphonInsize()`** — L388 — `PORT void SetSiphonInsize (int id, int size)`
  Sets siphon insize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/siphon.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
