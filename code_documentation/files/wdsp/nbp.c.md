# `wdsp/nbp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Main bandpass filter and the notched-bandpass (auto/manual notch database) filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×12)
  - `wdsp/snb.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×17)
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/snb.c` (calls ×3)
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/fir.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_nbp()` (×2), `destroy_nbp()` (×2), `flush_nbp()` (×2), `xnbp()` (×2), `create_notchdb()` (×1), `destroy_notchdb()` (×1), `setSamplerate_nbp()` (×1), `setBuffers_nbp()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_notchdb()`** — L35 — `NOTCHDB create_notchdb (int master_run, int maxnotches)`
  Constructor for the `notchdb` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_notchdb()`** — L49 — `void destroy_notchdb (NOTCHDB b)`
  Destroys the `notchdb` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`fir_mbandpass()`** — L64 — `double* fir_mbandpass (int N, int nbp, double* flow, double* fhigh, double rate, double scale, int wintype)`
  Called by: `calc_nbp_lightweight()` (same file), `calc_nbp_impulse()` (same file)
- **`min_notch_width()`** — L82 — `double min_notch_width (NBP a)`
  Called by: `calc_nbp_lightweight()` (same file), `calc_nbp_impulse()` (same file), `RXANBPGetMinNotchWidth()` (same file)
- **`make_nbp()`** — L100 — `int make_nbp (int nn, int* active, double* center, double* width, double* nlow, double* nhigh, double minwidth, int autoincr, double flow, double fhigh, double* bplow, double* bphi`
  Called by: `calc_nbp_lightweight()` (same file), `calc_nbp_impulse()` (same file)
- **`calc_nbp_lightweight()`** — L184 — `void calc_nbp_lightweight (NBP a)`
  Called by: `UpdateNBPFiltersLightWeight()` (same file)
- **`calc_nbp_impulse()`** — L217 — `void calc_nbp_impulse (NBP a)`
  Called by: `create_nbp()` (same file), `setSamplerate_nbp()` (same file), `setSize_nbp()` (same file), `setNc_nbp()` (same file), `UpdateNBPFilters()` (same file), `RXANBPSetNotchesRun()` (same file) — and 4 more
- **`create_nbp()`** — L244 — `NBP create_nbp(int run, int fnfrun, int position, int size, int nc, int mp, double* in, double* out, double flow, double fhigh, int rate, int wintype, double gain, int autoincr, in`
  Constructor for the `nbp` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `calc_bpsnba()` (`wdsp/snb.c`)
- **`destroy_nbp()`** — L273 — `void destroy_nbp (NBP a)`
  Destroys the `nbp` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `decalc_bpsnba()` (`wdsp/snb.c`)
- **`flush_nbp()`** — L281 — `void flush_nbp (NBP a)`
  Flushes (zeroes) the `nbp` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `flush_bpsnba()` (`wdsp/snb.c`)
- **`xnbp()`** — L286 — `void xnbp (NBP a, int pos)`
  Runs the `nbp` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xbpsnbaout()` (`wdsp/snb.c`)
- **`setBuffers_nbp()`** — L294 — `void setBuffers_nbp (NBP a, double* in, double* out)`
  Re-points the `nbp` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_nbp()`** — L301 — `void setSamplerate_nbp (NBP a, int rate)`
  Reconfigures the `nbp` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_nbp()`** — L309 — `void setSize_nbp (NBP a, int size)`
  Reconfigures the `nbp` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setNc_nbp()`** — L319 — `void setNc_nbp (NBP a)`
  Called by: `RXANBPSetNC()` (same file), `RXABPSNBASetNC()` (`wdsp/snb.c`)
- **`setMp_nbp()`** — L326 — `void setMp_nbp (NBP a)`
  Called by: `RXANBPSetMP()` (same file), `RXABPSNBASetMP()` (`wdsp/snb.c`)
- **`UpdateNBPFiltersLightWeight()`** — L339 — `void UpdateNBPFiltersLightWeight (int channel)`
  Called by: `RXANBPSetTuneFrequency()` (same file), `RXANBPSetShiftFrequency()` (same file)
- **`UpdateNBPFilters()`** — L345 — `void UpdateNBPFilters(int channel)`
  Called by: `RXANBPAddNotch()` (same file), `RXANBPDeleteNotch()` (same file), `RXANBPEditNotch()` (same file)
- **`RXANBPAddNotch()`** — L361 — `PORT int RXANBPAddNotch (int channel, int notch, double fcenter, double fwidth, int active)`
  RXA chain operation — nbpadd notch; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPGetNotch()`** — L392 — `PORT int RXANBPGetNotch (int channel, int notch, double* fcenter, double* fwidth, int* active)`
  RXA chain operation — nbpget notch; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPDeleteNotch()`** — L417 — `PORT int RXANBPDeleteNotch (int channel, int notch)`
  RXA chain operation — nbpdelete notch; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPEditNotch()`** — L443 — `PORT int RXANBPEditNotch (int channel, int notch, double fcenter, double fwidth, int active)`
  RXA chain operation — nbpedit notch; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPGetNumNotches()`** — L464 — `PORT void RXANBPGetNumNotches (int channel, int* nnotches)`
  RXA chain operation — nbpget num notches; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetTuneFrequency()`** — L474 — `PORT void RXANBPSetTuneFrequency (int channel, double tunefreq)`
  RXA chain operation — nbpset tune frequency; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetShiftFrequency()`** — L486 — `PORT void RXANBPSetShiftFrequency (int channel, double shift)`
  RXA chain operation — nbpset shift frequency; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetNotchesRun()`** — L498 — `PORT void RXANBPSetNotchesRun (int channel, int run)`
  RXA chain operation — nbpset notches run; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetRun()`** — L520 — `PORT void RXANBPSetRun (int channel, int run)`
  RXA chain operation — nbpset run; part of the receive/transmit chain API.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`RXANBPSetFreqs()`** — L530 — `PORT void RXANBPSetFreqs (int channel, double flow, double fhigh)`
  RXA chain operation — nbpset freqs; part of the receive/transmit chain API.
  Called by: `RXASetPassband()` (`wdsp/RXA.c`)
- **`RXANBPSetWindow()`** — L545 — `PORT void RXANBPSetWindow (int channel, int wintype)`
  RXA chain operation — nbpset window; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetNC()`** — L566 — `PORT void RXANBPSetNC (int channel, int nc)`
  RXA chain operation — nbpset nc; part of the receive/transmit chain API.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`RXANBPSetMP()`** — L581 — `PORT void RXANBPSetMP (int channel, int mp)`
  RXA chain operation — nbpset mp; part of the receive/transmit chain API.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)
- **`RXANBPGetMinNotchWidth()`** — L593 — `PORT void RXANBPGetMinNotchWidth (int channel, double* minwidth)`
  RXA chain operation — nbpget min notch width; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXANBPSetAutoIncrease()`** — L603 — `PORT void RXANBPSetAutoIncrease (int channel, int autoincr)`
  RXA chain operation — nbpset auto increase; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nbp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
