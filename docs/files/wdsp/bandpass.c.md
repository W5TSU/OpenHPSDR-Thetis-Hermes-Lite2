# `wdsp/bandpass.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Main bandpass filter and the notched-bandpass (auto/manual notch database) filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×12)
  - `wdsp/TXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×19)
  - `wdsp/fir.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `flush_bandpass()` (×3), `create_bandpass()` (×2), `destroy_bandpass()` (×2), `xbandpass()` (×2), `setSamplerate_bandpass()` (×2), `setBuffers_bandpass()` (×2), `setSize_bandpass()` (×2), `setGain_bandpass()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_bps()`** — L35 — `void calc_bps (BPS a)`
  Called by: `create_bps()` (same file), `setBuffers_bps()` (same file), `setSamplerate_bps()` (same file), `setSize_bps()` (same file), `setFreqs_bps()` (same file)
- **`decalc_bps()`** — L47 — `void decalc_bps (BPS a)`
  Called by: `destroy_bps()` (same file), `setBuffers_bps()` (same file), `setSamplerate_bps()` (same file), `setSize_bps()` (same file), `setFreqs_bps()` (same file)
- **`create_bps()`** — L56 — `BPS create_bps (int run, int position, int size, double* in, double* out, double f_low, double f_high, int samplerate, int wintype, double gain)`
  Constructor for the `bps` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_bps()`** — L74 — `void destroy_bps (BPS a)`
  Destroys the `bps` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_bps()`** — L80 — `void flush_bps (BPS a)`
  Flushes (zeroes) the `bps` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xbps()`** — L85 — `void xbps (BPS a, int pos)`
  Runs the `bps` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_bps()`** — L107 — `void setBuffers_bps (BPS a, double* in, double* out)`
  Re-points the `bps` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_bps()`** — L115 — `void setSamplerate_bps (BPS a, int rate)`
  Reconfigures the `bps` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_bps()`** — L122 — `void setSize_bps (BPS a, int size)`
  Reconfigures the `bps` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setFreqs_bps()`** — L129 — `void setFreqs_bps (BPS a, double f_low, double f_high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_bandpass()`** — L284 — `BANDPASS create_bandpass (int run, int position, int size, int nc, int mp, double* in, double* out, double f_low, double f_high, int samplerate, int wintype, double gain)`
  Constructor for the `bandpass` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_bandpass()`** — L308 — `void destroy_bandpass (BANDPASS a)`
  Destroys the `bandpass` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_bandpass()`** — L314 — `void flush_bandpass (BANDPASS a)`
  Flushes (zeroes) the `bandpass` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `RXAbp1Set()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xbandpass()`** — L319 — `void xbandpass (BANDPASS a, int pos)`
  Runs the `bandpass` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_bandpass()`** — L327 — `void setBuffers_bandpass (BANDPASS a, double* in, double* out)`
  Re-points the `bandpass` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_bandpass()`** — L334 — `void setSamplerate_bandpass (BANDPASS a, int rate)`
  Reconfigures the `bandpass` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_bandpass()`** — L343 — `void setSize_bandpass (BANDPASS a, int size)`
  Reconfigures the `bandpass` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setGain_bandpass()`** — L355 — `void setGain_bandpass (BANDPASS a, double gain, int update)`
  Called by: `RXAbp1Check()` (`wdsp/RXA.c`)
- **`CalcBandpassFilter()`** — L364 — `void CalcBandpassFilter (BANDPASS a, double f_low, double f_high, double gain)`
  Called by: `TXASetupBPFilters()` (`wdsp/TXA.c`)
- **`SetRXABandpassRun()`** — L384 — `PORT void SetRXABandpassRun (int channel, int run)`
  Sets rxabandpass run — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXABandpassFreqs()`** — L392 — `PORT void SetRXABandpassFreqs (int channel, double f_low, double f_high)`
  Sets rxabandpass freqs — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetPassband()` (`wdsp/RXA.c`)
- **`SetRXABandpassWindow()`** — L411 — `PORT void SetRXABandpassWindow (int channel, int wintype)`
  Sets rxabandpass window — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXABandpassNC()`** — L429 — `PORT void SetRXABandpassNC (int channel, int nc)`
  Sets rxabandpass nc — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`SetRXABandpassMP()`** — L447 — `PORT void SetRXABandpassMP (int channel, int mp)`
  Sets rxabandpass mp — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)
- **`SetTXABandpassRun()`** — L465 — `PORT void SetTXABandpassRun (int channel, int run)`
  Sets txabandpass run — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXABandpassWindow()`** — L507 — `PORT void SetTXABandpassWindow (int channel, int wintype)`
  Sets txabandpass window — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXABandpassNC()`** — L538 — `PORT void SetTXABandpassNC (int channel, int nc)`
  Sets txabandpass nc — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetNC()` (`wdsp/TXA.c`)
- **`SetTXABandpassMP()`** — L572 — `PORT void SetTXABandpassMP (int channel, int mp)`
  Sets txabandpass mp — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetMP()` (`wdsp/TXA.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/bandpass.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
