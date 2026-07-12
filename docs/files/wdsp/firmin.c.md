# `wdsp/firmin.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/eq.c` (calls ×22)
  - `wdsp/bandpass.c` (calls ×19)
  - `wdsp/nbp.c` (calls ×17)
  - `wdsp/fmd.c` (calls ×13)
  - `wdsp/gaussian.c` (calls ×13)
  - `wdsp/doublepole.c` (calls ×12)
  - `wdsp/fmmod.c` (calls ×12)
  - `wdsp/matchedCW.c` (calls ×12)
  - `wdsp/emph.c` (calls ×11)
  - `wdsp/fmsq.c` (calls ×7)
  - `wdsp/cfir.c` (calls ×4)
  - `wdsp/dexp.c` (calls ×4)
  - …and 3 more files
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×7)
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setImpulse_fircore()` (×49), `setNc_fircore()` (×17), `create_fircore()` (×14), `destroy_fircore()` (×14), `xfircore()` (×13), `flush_fircore()` (×12), `setBuffers_fircore()` (×10), `setMp_fircore()` (×10)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_firmin()`** — L35 — `void calc_firmin (FIRMIN a)`
  Called by: `create_firmin()` (same file), `setSamplerate_firmin()` (same file), `setFreqs_firmin()` (same file)
- **`create_firmin()`** — L44 — `FIRMIN create_firmin (int run, int position, int size, double* in, double* out, int nc, double f_low, double f_high, int samplerate, int wintype, double gain)`
  Constructor for the `firmin` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_firmin()`** — L63 — `void destroy_firmin (FIRMIN a)`
  Destroys the `firmin` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_firmin()`** — L70 — `void flush_firmin (FIRMIN a)`
  Flushes (zeroes) the `firmin` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xfirmin()`** — L76 — `void xfirmin (FIRMIN a, int pos)`
  Runs the `firmin` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_firmin()`** — L101 — `void setBuffers_firmin (FIRMIN a, double* in, double* out)`
  Re-points the `firmin` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_firmin()`** — L107 — `void setSamplerate_firmin (FIRMIN a, int rate)`
  Reconfigures the `firmin` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_firmin()`** — L113 — `void setSize_firmin (FIRMIN a, int size)`
  Reconfigures the `firmin` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setFreqs_firmin()`** — L118 — `void setFreqs_firmin (FIRMIN a, double f_low, double f_high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`plan_firopt()`** — L131 — `void plan_firopt (FIROPT a)`
  Called by: `create_firopt()` (same file), `setBuffers_firopt()` (same file), `setSize_firopt()` (same file)
- **`calc_firopt()`** — L155 — `void calc_firopt (FIROPT a)`
  Called by: `create_firopt()` (same file), `setBuffers_firopt()` (same file), `setSamplerate_firopt()` (same file), `setSize_firopt()` (same file), `setFreqs_firopt()` (same file)
- **`create_firopt()`** — L172 — `FIROPT create_firopt (int run, int position, int size, double* in, double* out, int nc, double f_low, double f_high, int samplerate, int wintype, double gain)`
  Constructor for the `firopt` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`deplan_firopt()`** — L192 — `void deplan_firopt (FIROPT a)`
  Called by: `destroy_firopt()` (same file), `setBuffers_firopt()` (same file), `setSize_firopt()` (same file)
- **`destroy_firopt()`** — L212 — `void destroy_firopt (FIROPT a)`
  Destroys the `firopt` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_firopt()`** — L218 — `void flush_firopt (FIROPT a)`
  Flushes (zeroes) the `firopt` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xfiropt()`** — L227 — `void xfiropt (FIROPT a, int pos)`
  Runs the `firopt` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setBuffers_firopt()`** — L253 — `void setBuffers_firopt (FIROPT a, double* in, double* out)`
  Re-points the `firopt` block’s input/output buffers (called when the channel’s buffers change).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSamplerate_firopt()`** — L262 — `void setSamplerate_firopt (FIROPT a, int rate)`
  Reconfigures the `firopt` block for a new sample rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setSize_firopt()`** — L268 — `void setSize_firopt (FIROPT a, int size)`
  Reconfigures the `firopt` block for a new buffer size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`setFreqs_firopt()`** — L276 — `void setFreqs_firopt (FIROPT a, double f_low, double f_high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`plan_fircore()`** — L290 — `void plan_fircore (FIRCORE a)`
  Called by: `create_fircore()` (same file), `setBuffers_fircore()` (same file), `setSize_fircore()` (same file), `setNc_fircore()` (same file)
- **`calc_fircore()`** — L322 — `void calc_fircore (FIRCORE a, int flip)`
  Called by: `create_fircore()` (same file), `setBuffers_fircore()` (same file), `setSize_fircore()` (same file), `setImpulse_fircore()` (same file), `setNc_fircore()` (same file), `setMp_fircore()` (same file)
- **`create_fircore()`** — L348 — `FIRCORE create_fircore (int size, double* in, double* out, int nc, int mp, double* impulse)`
  Constructor for the `fircore` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_bandpass()` (`wdsp/bandpass.c`), `calc_cfir()` (`wdsp/cfir.c`), `calc_filter()` (`wdsp/dexp.c`), `create_doublepole()` (`wdsp/doublepole.c`), `create_emphp()` (`wdsp/emph.c`), `create_eqp()` (`wdsp/eq.c`) — and 8 more
- **`deplan_fircore()`** — L365 — `void deplan_fircore (FIRCORE a)`
  Called by: `destroy_fircore()` (same file), `setBuffers_fircore()` (same file), `setSize_fircore()` (same file), `setNc_fircore()` (same file)
- **`destroy_fircore()`** — L391 — `void destroy_fircore (FIRCORE a)`
  Destroys the `fircore` block, freeing its allocated buffers.
  Called by: `destroy_bandpass()` (`wdsp/bandpass.c`), `decalc_cfir()` (`wdsp/cfir.c`), `decalc_filter()` (`wdsp/dexp.c`), `destroy_doublepole()` (`wdsp/doublepole.c`), `destroy_emphp()` (`wdsp/emph.c`), `destroy_eqp()` (`wdsp/eq.c`) — and 8 more
- **`flush_fircore()`** — L400 — `void flush_fircore (FIRCORE a)`
  Flushes (zeroes) the `fircore` block’s internal buffers/state.
  Called by: `flush_bandpass()` (`wdsp/bandpass.c`), `flush_cfir()` (`wdsp/cfir.c`), `flush_dexp()` (`wdsp/dexp.c`), `flush_doublepole()` (`wdsp/doublepole.c`), `flush_emphp()` (`wdsp/emph.c`), `flush_eqp()` (`wdsp/eq.c`) — and 6 more
- **`xfircore()`** — L409 — `void xfircore(FIRCORE a)`
  Runs the `fircore` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xbandpass()` (`wdsp/bandpass.c`), `xcfir()` (`wdsp/cfir.c`), `xdexp()` (`wdsp/dexp.c`), `xdoublepole()` (`wdsp/doublepole.c`), `xemphp()` (`wdsp/emph.c`), `xeqp()` (`wdsp/eq.c`) — and 7 more
- **`setBuffers_fircore()`** — L441 — `void setBuffers_fircore (FIRCORE a, double* in, double* out)`
  Re-points the `fircore` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setBuffers_bandpass()` (`wdsp/bandpass.c`), `setBuffers_doublepole()` (`wdsp/doublepole.c`), `setBuffers_emphp()` (`wdsp/emph.c`), `setBuffers_eqp()` (`wdsp/eq.c`), `setBuffers_fmd()` (`wdsp/fmd.c`), `setBuffers_fmmod()` (`wdsp/fmmod.c`) — and 4 more
- **`setSize_fircore()`** — L450 — `void setSize_fircore (FIRCORE a, int size)`
  Reconfigures the `fircore` block for a new buffer size.
  Called by: `setSize_bandpass()` (`wdsp/bandpass.c`), `setSize_doublepole()` (`wdsp/doublepole.c`), `setSize_emphp()` (`wdsp/emph.c`), `setSize_eqp()` (`wdsp/eq.c`), `setSize_fmmod()` (`wdsp/fmmod.c`), `setSize_gaussian()` (`wdsp/gaussian.c`) — and 2 more
- **`setImpulse_fircore()`** — L458 — `void setImpulse_fircore (FIRCORE a, double* impulse, int update)`
  Called by: `setSamplerate_bandpass()` (`wdsp/bandpass.c`), `setSize_bandpass()` (`wdsp/bandpass.c`), `setGain_bandpass()` (`wdsp/bandpass.c`), `CalcBandpassFilter()` (`wdsp/bandpass.c`), `SetRXABandpassFreqs()` (`wdsp/bandpass.c`), `SetRXABandpassWindow()` (`wdsp/bandpass.c`) — and 43 more
- **`setNc_fircore()`** — L464 — `void setNc_fircore (FIRCORE a, int nc, double* impulse)`
  Called by: `SetRXABandpassNC()` (`wdsp/bandpass.c`), `SetTXABandpassNC()` (`wdsp/bandpass.c`), `setSamplerate_doublepole()` (`wdsp/doublepole.c`), `CalcDoublepoleFilter()` (`wdsp/doublepole.c`), `SetTXAFMEmphNC()` (`wdsp/emph.c`), `SetRXAEQNC()` (`wdsp/eq.c`) — and 11 more
- **`setMp_fircore()`** — L478 — `void setMp_fircore (FIRCORE a, int mp)`
  Called by: `SetRXABandpassMP()` (`wdsp/bandpass.c`), `SetTXABandpassMP()` (`wdsp/bandpass.c`), `SetTXAFMEmphMP()` (`wdsp/emph.c`), `SetRXAEQMP()` (`wdsp/eq.c`), `SetTXAEQMP()` (`wdsp/eq.c`), `SetRXAFMMPde()` (`wdsp/fmd.c`) — and 4 more
- **`setUpdate_fircore()`** — L484 — `void setUpdate_fircore (FIRCORE a)`
  Called by: `RXAbp1Set()` (`wdsp/RXA.c`), `RXAbpsnbaSet()` (`wdsp/RXA.c`), `SetRXABandpassFreqs()` (`wdsp/bandpass.c`), `SetRXABandpassWindow()` (`wdsp/bandpass.c`), `SetTXAFMDeviation()` (`wdsp/fmmod.c`), `RXANBPSetNotchesRun()` (`wdsp/nbp.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/firmin.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
