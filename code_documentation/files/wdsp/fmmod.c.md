# `wdsp/fmmod.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM and FM modulators for TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/fir.c` (calls ×6)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_fmmod()` (×1), `destroy_fmmod()` (×1), `flush_fmmod()` (×1), `xfmmod()` (×1), `setSamplerate_fmmod()` (×1), `setBuffers_fmmod()` (×1), `setSize_fmmod()` (×1), `SetTXAFMNC()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_fmmod()`** — L29 — `void calc_fmmod (FMMOD a)`
  Called by: `create_fmmod()` (same file), `setBuffers_fmmod()` (same file), `setSamplerate_fmmod()` (same file), `setSize_fmmod()` (same file)
- **`create_fmmod()`** — L42 — `FMMOD create_fmmod (int run, int size, double* in, double* out, int rate, double dev, double f_low, double f_high, int ctcss_run, double ctcss_level, double ctcss_freq, int bp_run,`
  Constructor for the `fmmod` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_fmmod()`** — L68 — `void destroy_fmmod (FMMOD a)`
  Destroys the `fmmod` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_fmmod()`** — L74 — `void flush_fmmod (FMMOD a)`
  Flushes (zeroes) the `fmmod` block’s internal buffers/state.
  Called by: `flush_txa()` (`wdsp/TXA.c`)
- **`xfmmod()`** — L80 — `void xfmmod (FMMOD a)`
  Runs the `fmmod` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_fmmod()`** — L112 — `void setBuffers_fmmod (FMMOD a, double* in, double* out)`
  Re-points the `fmmod` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_fmmod()`** — L120 — `void setSamplerate_fmmod (FMMOD a, int rate)`
  Reconfigures the `fmmod` block for a new sample rate.
  Called by: `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_fmmod()`** — L130 — `void setSize_fmmod (FMMOD a, int size)`
  Reconfigures the `fmmod` block for a new buffer size.
  Called by: `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`SetTXAFMDeviation()`** — L147 — `PORT void SetTXAFMDeviation (int channel, double deviation)`
  Sets txafmdeviation — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACTCSSFreq()`** — L166 — `PORT void SetTXACTCSSFreq (int channel, double freq)`
  Sets txactcssfreq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXACTCSSRun()`** — L178 — `PORT void SetTXACTCSSRun (int channel, int run)`
  Sets txactcssrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAFMNC()`** — L186 — `PORT void SetTXAFMNC (int channel, int nc)`
  Sets txafmnc — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetNC()` (`wdsp/TXA.c`)
- **`SetTXAFMMP()`** — L203 — `PORT void SetTXAFMMP (int channel, int mp)`
  Sets txafmmp — API setter, typically called from the console via P/Invoke.
  Called by: `TXASetMP()` (`wdsp/TXA.c`)
- **`SetTXAFMAFFreqs()`** — L215 — `PORT void SetTXAFMAFFreqs (int channel, double low, double high)`
  Sets txafmaffreqs — API setter, typically called from the console via P/Invoke.
  Called by: `SetTXAFMAFFilter()` (`wdsp/TXA.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fmmod.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
