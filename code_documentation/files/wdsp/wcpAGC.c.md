# `wdsp/wcpAGC.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** The WDSP AGC (receive gain control and TX leveler).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
  - `wdsp/fmd.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_wcpagc()` (×3), `destroy_wcpagc()` (×3), `flush_wcpagc()` (×3), `xwcpagc()` (×3), `setSamplerate_wcpagc()` (×3), `setBuffers_wcpagc()` (×3), `setSize_wcpagc()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_wcpagc()`** — L35 — `void calc_wcpagc (WCPAGC a)`
  Called by: `create_wcpagc()` (same file), `setSamplerate_wcpagc()` (same file), `setSize_wcpagc()` (same file)
- **`decalc_wcpagc()`** — L54 — `void decalc_wcpagc (WCPAGC a)`
  Called by: `destroy_wcpagc()` (same file), `setSamplerate_wcpagc()` (same file), `setSize_wcpagc()` (same file)
- **`create_wcpagc()`** — L60 — `WCPAGC create_wcpagc ( int run, int mode, int pmode, double* in, double* out,`
  Constructor for the `wcpagc` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`), `calc_fmd()` (`wdsp/fmd.c`)
- **`loadWcpAGC()`** — L115 — `void loadWcpAGC (WCPAGC a)`
  Called by: `calc_wcpagc()` (same file), `SetRXAAGCMode()` (same file), `SetRXAAGCAttack()` (same file), `SetRXAAGCDecay()` (same file), `SetRXAAGCHang()` (same file), `SetRXAAGCHangLevel()` (same file) — and 14 more
- **`destroy_wcpagc()`** — L148 — `void destroy_wcpagc (WCPAGC a)`
  Destroys the `wcpagc` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`), `decalc_fmd()` (`wdsp/fmd.c`)
- **`flush_wcpagc()`** — L154 — `void flush_wcpagc (WCPAGC a)`
  Flushes (zeroes) the `wcpagc` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`), `flush_fmd()` (`wdsp/fmd.c`)
- **`xwcpagc()`** — L161 — `void xwcpagc (WCPAGC a)`
  Runs the `wcpagc` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`), `xfmd()` (`wdsp/fmd.c`)
- **`setBuffers_wcpagc()`** — L348 — `void setBuffers_wcpagc (WCPAGC a, double* in, double* out)`
  Re-points the `wcpagc` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`), `setBuffers_fmd()` (`wdsp/fmd.c`)
- **`setSamplerate_wcpagc()`** — L354 — `void setSamplerate_wcpagc (WCPAGC a, int rate)`
  Reconfigures the `wcpagc` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`), `setSamplerate_fmd()` (`wdsp/fmd.c`)
- **`setSize_wcpagc()`** — L361 — `void setSize_wcpagc (WCPAGC a, int size)`
  Reconfigures the `wcpagc` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`), `setSize_fmd()` (`wdsp/fmd.c`)
- **`SetRXAAGCMode()`** — L374 — `PORT void SetRXAAGCMode (int channel, int mode)`
  Sets rxaagcmode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCAttack()`** — L417 — `PORT void SetRXAAGCAttack (int channel, int attack)`
  Sets rxaagcattack — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCDecay()`** — L426 — `PORT void SetRXAAGCDecay (int channel, int decay)`
  Sets rxaagcdecay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCHang()`** — L435 — `PORT void SetRXAAGCHang (int channel, int hang)`
  Sets rxaagchang — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAAGCHangLevel()`** — L444 — `PORT void GetRXAAGCHangLevel(int channel, double *hangLevel)`
  Returns rxaagchang level — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCHangLevel()`** — L453 — `PORT void SetRXAAGCHangLevel(int channel, double hangLevel)`
  Sets rxaagchang level — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAAGCHangThreshold()`** — L472 — `PORT void GetRXAAGCHangThreshold(int channel, int *hangthreshold)`
  Returns rxaagchang threshold — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCHangThreshold()`** — L481 — `PORT void SetRXAAGCHangThreshold (int channel, int hangthreshold)`
  Sets rxaagchang threshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAAGCThresh()`** — L491 — `PORT void GetRXAAGCThresh(int channel, double *thresh, double size, double rate)`
  Returns rxaagcthresh — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCThresh()`** — L503 — `PORT void SetRXAAGCThresh(int channel, double thresh, double size, double rate)`
  Sets rxaagcthresh — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAAGCTop()`** — L517 — `PORT void GetRXAAGCTop(int channel, double *max_agc)`
  Returns rxaagctop — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCTop()`** — L526 — `PORT void SetRXAAGCTop (int channel, double max_agc)`
  Sets rxaagctop — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCSlope()`** — L536 — `PORT void SetRXAAGCSlope (int channel, int slope)`
  Sets rxaagcslope — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCFixed()`** — L545 — `PORT void SetRXAAGCFixed (int channel, double fixed_agc)`
  Sets rxaagcfixed — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAAGCMaxInputLevel()`** — L554 — `PORT void SetRXAAGCMaxInputLevel (int channel, double level)`
  Sets rxaagcmax input level — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAALCSt()`** — L569 — `PORT void SetTXAALCSt (int channel, int state)`
  Sets txaalcst — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAALCAttack()`** — L577 — `PORT void SetTXAALCAttack (int channel, int attack)`
  Sets txaalcattack — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAALCDecay()`** — L585 — `PORT void SetTXAALCDecay (int channel, int decay)`
  Sets txaalcdecay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAALCHang()`** — L594 — `PORT void SetTXAALCHang (int channel, int hang)`
  Sets txaalchang — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAALCMaxGain()`** — L603 — `PORT void SetTXAALCMaxGain (int channel, double maxgain)`
  Sets txaalcmax gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXALevelerSt()`** — L612 — `PORT void SetTXALevelerSt (int channel, int state)`
  Sets txaleveler st — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXALevelerAttack()`** — L620 — `PORT void SetTXALevelerAttack (int channel, int attack)`
  Sets txaleveler attack — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXALevelerDecay()`** — L629 — `PORT void SetTXALevelerDecay (int channel, int decay)`
  Sets txaleveler decay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXALevelerHang()`** — L638 — `PORT void SetTXALevelerHang (int channel, int hang)`
  Sets txaleveler hang — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXALevelerTop()`** — L647 — `PORT void SetTXALevelerTop (int channel, double maxgain)`
  Sets txaleveler top — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/wcpAGC.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
