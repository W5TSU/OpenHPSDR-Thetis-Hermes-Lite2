# `wdsp/fmd.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM/SAM (synchronous) and FM demodulators.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×11)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×13)
  - `wdsp/wcpAGC.c` (calls ×7)
  - `wdsp/iir.c` (calls ×6)
  - `wdsp/fcurve.c` (calls ×5)
  - `wdsp/fir.c` (calls ×5)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_fmd()` (×1), `destroy_fmd()` (×1), `flush_fmd()` (×1), `xfmd()` (×1), `setSamplerate_fmd()` (×1), `setBuffers_fmd()` (×1), `setSize_fmd()` (×1), `SetRXAFMNCaud()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_fmd()`** — L29 — `void calc_fmd (FMD a)`
  Called by: `create_fmd()` (same file), `setBuffers_fmd()` (same file), `setSamplerate_fmd()` (same file), `setSize_fmd()` (same file), `SetRXAFMLimGain()` (same file)
- **`decalc_fmd()`** — L75 — `void decalc_fmd (FMD a)`
  Called by: `destroy_fmd()` (same file), `setBuffers_fmd()` (same file), `setSamplerate_fmd()` (same file), `setSize_fmd()` (same file), `SetRXAFMLimGain()` (same file)
- **`create_fmd()`** — L81 — `FMD create_fmd( int run, int size, double* in, double* out, int rate, double deviation, double f_low, double f_high, double fmin, double fmax, double zeta, double omegaN, double ta`
  Constructor for the `fmd` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_fmd()`** — L122 — `void destroy_fmd (FMD a)`
  Destroys the `fmd` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_fmd()`** — L131 — `void flush_fmd (FMD a)`
  Flushes (zeroes) the `fmd` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xfmd()`** — L144 — `void xfmd (FMD a)`
  Runs the `fmd` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_fmd()`** — L190 — `void setBuffers_fmd (FMD a, double* in, double* out)`
  Re-points the `fmd` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_fmd()`** — L201 — `void setSamplerate_fmd (FMD a, int rate)`
  Reconfigures the `fmd` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_fmd()`** — L218 — `void setSize_fmd (FMD a, int size)`
  Reconfigures the `fmd` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAFMDeviation()`** — L245 — `PORT void SetRXAFMDeviation (int channel, double deviation)`
  Sets rxafmdeviation — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXACTCSSFreq()`** — L256 — `PORT void SetRXACTCSSFreq (int channel, double freq)`
  Sets rxactcssfreq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXACTCSSRun()`** — L267 — `PORT void SetRXACTCSSRun (int channel, int run)`
  Sets rxactcssrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAFMNCde()`** — L278 — `PORT void SetRXAFMNCde (int channel, int nc)`
  Sets rxafmncde — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`SetRXAFMMPde()`** — L295 — `PORT void SetRXAFMMPde (int channel, int mp)`
  Sets rxafmmpde — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)
- **`SetRXAFMNCaud()`** — L307 — `PORT void SetRXAFMNCaud (int channel, int nc)`
  Sets rxafmncaud — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetNC()` (`wdsp/RXA.c`)
- **`SetRXAFMMPaud()`** — L324 — `PORT void SetRXAFMMPaud (int channel, int mp)`
  Sets rxafmmpaud — API setter, typically called from the console via P/Invoke.
  Called by: `RXASetMP()` (`wdsp/RXA.c`)
- **`SetRXAFMLimRun()`** — L336 — `PORT void SetRXAFMLimRun (int channel, int run)`
  Sets rxafmlim run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAFMLimGain()`** — L349 — `PORT void SetRXAFMLimGain (int channel, double gaindB)`
  Sets rxafmlim gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAFMAFFilter()`** — L364 — `PORT void SetRXAFMAFFilter(int channel, double low, double high)`
  Sets rxafmaffilter — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmd.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
