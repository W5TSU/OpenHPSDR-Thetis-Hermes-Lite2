# `wdsp/calcc.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** PureSignal calibration calculation and the I/Q correction applied to TX.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×2)
  - `ChannelMaster/sync.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/iqc.c` (calls ×10)
  - `wdsp/delay.c` (calls ×8)
  - `wdsp/lmath.c` (calls ×3)
  - `wdsp/utilities.c` (calls ×2)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `pscc()` (×1), `create_calcc()` (×1), `destroy_calcc()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`size_calcc()`** — L30 — `void size_calcc (CALCC a)`
  Called by: `create_calcc()` (same file), `SetPSIntsAndSpi()` (same file)
- **`desize_calcc()`** — L83 — `void desize_calcc (CALCC a)`
  Called by: `destroy_calcc()` (same file), `SetPSIntsAndSpi()` (same file)
- **`create_calcc()`** — L119 — `CALCC create_calcc (int channel, int runcal, int size, int rate, int ints, int spi, double hw_scale, double moxdelay, double loopdelay, double ptol, int mox, int solidmox, int pin,`
  Constructor for the `calcc` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_txa()` (`wdsp/TXA.c`)
- **`destroy_calcc()`** — L205 — `void destroy_calcc (CALCC a)`
  Destroys the `calcc` block, freeing its allocated buffers.
  Called by: `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_calcc()`** — L241 — `void flush_calcc (CALCC a)`
  Flushes (zeroes) the `calcc` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`scheck()`** — L247 — `void scheck(CALCC a)`
  Called by: `calc()` (same file)
- **`rxscheck()`** — L294 — `void rxscheck (int rints, double* tvec, double* coef, int* info)`
  Called by: `calc()` (same file)
- **`calc()`** — L324 — `void calc (CALCC a)`
  Called by: `doPSCalcCorrection()` (same file)
- **`doPSCalcCorrection()`** — L485 — `void __cdecl doPSCalcCorrection (void *arg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`doPSTurnoff()`** — L509 — `void __cdecl doPSTurnoff (void *arg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`PSSaveCorrection()`** — L539 — `void __cdecl PSSaveCorrection (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`PSRestoreCorrection()`** — L572 — `void __cdecl PSRestoreCorrection(void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pscc()`** — L616 — `PORT void pscc (int channel, int size, double* tx, double* rx)`
  Called by: `psccF()` (same file), `InboundBlock()` (`ChannelMaster/sync.c`)
- **`psccF()`** — L839 — `PORT void psccF (int channel, int size, float *Itxbuff, float *Qtxbuff, float *Irxbuff, float *Qrxbuff, int mox, int solidmox)`
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`PSSaveCorr()`** — L859 — `PORT void PSSaveCorr (int channel, char* filename)`
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`PSRestoreCorr()`** — L871 — `PORT void PSRestoreCorr (int channel, char* filename)`
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSRunCal()`** — L890 — `PORT void SetPSRunCal (int channel, int run)`
  Sets psrun cal — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSMox()`** — L900 — `PORT void SetPSMox (int channel, int mox)`
  Sets psmox — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`GetPSInfo()`** — L913 — `PORT void GetPSInfo (int channel, int *info)`
  Returns psinfo — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSReset()`** — L923 — `PORT void SetPSReset (int channel, int reset)`
  Sets psreset — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSMancal()`** — L933 — `PORT void SetPSMancal (int channel, int mancal)`
  Sets psmancal — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSAutomode()`** — L941 — `PORT void SetPSAutomode (int channel, int automode)`
  Sets psautomode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSTurnon()`** — L949 — `PORT void SetPSTurnon (int channel, int turnon)`
  Sets psturnon — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSControl()`** — L957 — `PORT void SetPSControl (int channel, int reset, int mancal, int automode, int turnon)`
  Sets pscontrol — API setter, typically called from the console via P/Invoke.
  Called by: `SetPSIntsAndSpi()` (same file)
- **`SetPSLoopDelay()`** — L970 — `PORT void SetPSLoopDelay (int channel, double delay)`
  Sets psloop delay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSMoxDelay()`** — L981 — `PORT void SetPSMoxDelay (int channel, double delay)`
  Sets psmox delay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSTXDelay()`** — L992 — `PORT double SetPSTXDelay (int channel, double delay)`
  Sets pstxdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSHWPeak()`** — L1015 — `PORT void SetPSHWPeak (int channel, double peak)`
  Sets pshwpeak — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`GetPSHWPeak()`** — L1025 — `PORT void GetPSHWPeak (int channel, double* peak)`
  Returns pshwpeak — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`GetPSMaxTX()`** — L1033 — `PORT void GetPSMaxTX (int channel, double* maxtx)`
  Returns psmax tx — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSPtol()`** — L1041 — `PORT void SetPSPtol (int channel, double ptol)`
  Sets psptol — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`GetPSDisp()`** — L1049 — `PORT void GetPSDisp (int channel, double* x, double* ym, double* yc, double* ys, double* cm, double* cc, double* cs)`
  Returns psdisp — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSFeedbackRate()`** — L1064 — `PORT void SetPSFeedbackRate (int channel, int rate)`
  Sets psfeedback rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSPinMode()`** — L1093 — `PORT void SetPSPinMode (int channel, int pin)`
  Sets pspin mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSMapMode()`** — L1101 — `PORT void SetPSMapMode (int channel, int map)`
  Sets psmap mode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`SetPSStabilize()`** — L1109 — `PORT void SetPSStabilize (int channel, int stbl)`
  Sets psstabilize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.
- **`ForceShutDown()`** — L1117 — `void ForceShutDown (CALCC a, IQC b, int timeout)`
  Called by: `SetPSIntsAndSpi()` (same file)
- **`SetPSIntsAndSpi()`** — L1131 — `PORT void SetPSIntsAndSpi (int channel, int ints, int spi)`
  Sets psints and spi — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/PSForm.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/calcc.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
