# `wdsp/emnr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Spectral noise reduction "NR2" (MMSE-based).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/FDnoiseIQ.h` (imports ×1)
  - `wdsp/calculus.h` (imports ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/zetaHat.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `create_emnr()` (×1), `destroy_emnr()` (×1), `flush_emnr()` (×1), `xemnr()` (×1), `setSamplerate_emnr()` (×1), `setBuffers_emnr()` (×1), `setSize_emnr()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`bessI0()`** — L44 — `double bessI0 (double x)`
  Called by: `calc_gain()` (same file)
- **`bessI1()`** — L82 — `double bessI1 (double x)`
  Called by: `calc_gain()` (same file)
- **`e1xb()`** — L128 — `double e1xb (double x)`
  Called by: `calc_gain()` (same file)
- **`calc_window()`** — L168 — `void calc_window (EMNR a)`
  Called by: `calc_emnr()` (same file)
- **`interpM()`** — L189 — `void interpM (double* res, double x, int nvals, double* xvals, double* yvals)`
  Called by: `calc_emnr()` (same file)
- **`readZetaHat()`** — L207 — `int readZetaHat(const char* zeta_file, int* rows, int* cols, double* gmin, double* gmax, double* ximin, double* ximax, double* zetaHat, int* zetaValid)`
  Called by: `calc_emnr()` (same file)
- **`CwriteZetaHat()`** — L247 — `void CwriteZetaHat(const char* cfile, int zetaHat_rows, int zetaHat_cols, double zetaHat_gmin, double zetaHat_gmax, double zetaHat_ximin, double zetaHat_ximax, double* zetaHat, int`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`calc_emnr()`** — L300 — `void calc_emnr(EMNR a)`
  Called by: `create_emnr()` (same file), `setSamplerate_emnr()` (same file), `setSize_emnr()` (same file)
- **`decalc_emnr()`** — L576 — `void decalc_emnr(EMNR a)`
  Called by: `destroy_emnr()` (same file), `setSamplerate_emnr()` (same file), `setSize_emnr()` (same file)
- **`create_emnr()`** — L637 — `EMNR create_emnr (int run, int position, int size, double* in, double* out, int fsize, int ovrlp, int rate, int wintype, double gain, int gain_method, int npe_method, int ae_run)`
  Constructor for the `emnr` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`flush_emnr()`** — L659 — `void flush_emnr (EMNR a)`
  Flushes (zeroes) the `emnr` block’s internal buffers/state.
  Called by: `SetRXAEMNRPosition()` (same file), `flush_rxa()` (`wdsp/RXA.c`)
- **`destroy_emnr()`** — L674 — `void destroy_emnr (EMNR a)`
  Destroys the `emnr` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`LambdaD()`** — L680 — `void LambdaD(EMNR a)`
  Called by: `calc_gain()` (same file)
- **`LambdaDs()`** — L805 — `void LambdaDs (EMNR a)`
  Called by: `calc_gain()` (same file)
- **`LambdaDl()`** — L820 — `void LambdaDl (EMNR a)`
  Called by: `calc_gain()` (same file)
- **`aepf()`** — L849 — `void aepf(EMNR a)`
  Called by: `calc_gain()` (same file)
- **`post2_calc_w()`** — L898 — `void post2_calc_w(EMNR a)`
  Called by: `calc_emnr()` (same file), `SetRXAEMNRpost2Taper()` (same file)
- **`post2()`** — L909 — `void post2(EMNR a)`
  Called by: `xemnr()` (same file)
- **`SetRXAEMNRpost2Run()`** — L951 — `PORT void SetRXAEMNRpost2Run(int channel, int run)`
  Sets rxaemnrpost2 run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRpost2Factor()`** — L959 — `PORT void SetRXAEMNRpost2Factor(int channel, double factor)`
  Sets rxaemnrpost2 factor — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRpost2Nlevel()`** — L967 — `PORT void SetRXAEMNRpost2Nlevel(int channel, double nlevel)`
  Sets rxaemnrpost2 nlevel — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRpost2Taper()`** — L975 — `PORT void SetRXAEMNRpost2Taper(int channel, int taper)`
  Sets rxaemnrpost2 taper — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRpost2Rate()`** — L985 — `PORT void SetRXAEMNRpost2Rate(int channel, double tc)`
  Sets rxaemnrpost2 rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`getKey()`** — L1001 — `double getKey(double* type, double gamma, double xi)`
  Called by: `calc_gain()` (same file)
- **`getZeta()`** — L1047 — `int getZeta( EMNR a, double gamma, double eps, double* zeta)`
  Called by: `calc_gain()` (same file)
- **`calc_gain()`** — L1069 — `void calc_gain (EMNR a)`
  Called by: `xemnr()` (same file)
- **`xemnr()`** — L1200 — `void xemnr (EMNR a, int pos)`
  Runs the `emnr` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_emnr()`** — L1256 — `void setBuffers_emnr (EMNR a, double* in, double* out)`
  Re-points the `emnr` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_emnr()`** — L1262 — `void setSamplerate_emnr (EMNR a, int rate)`
  Reconfigures the `emnr` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_emnr()`** — L1269 — `void setSize_emnr (EMNR a, int size)`
  Reconfigures the `emnr` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXAEMNRRun()`** — L1282 — `PORT void SetRXAEMNRRun (int channel, int run)`
  Sets rxaemnrrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRgainMethod()`** — L1299 — `PORT void SetRXAEMNRgainMethod (int channel, int method)`
  Sets rxaemnrgain method — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRnpeMethod()`** — L1307 — `PORT void SetRXAEMNRnpeMethod (int channel, int method)`
  Sets rxaemnrnpe method — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRaeRun()`** — L1315 — `PORT void SetRXAEMNRaeRun (int channel, int run)`
  Sets rxaemnrae run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRPosition()`** — L1323 — `PORT void SetRXAEMNRPosition (int channel, int position)`
  Sets rxaemnrposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRaeZetaThresh()`** — L1335 — `PORT void SetRXAEMNRaeZetaThresh (int channel, double zetathresh)`
  Sets rxaemnrae zeta thresh — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAEMNRaePsi()`** — L1343 — `PORT void SetRXAEMNRaePsi (int channel, double psi)`
  Sets rxaemnrae psi — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRXAEMNRtrainZetaThresh()`** — L1351 — `PORT void SetRXAEMNRtrainZetaThresh(int channel, double thresh)`
  Sets rxaemnrtrain zeta thresh — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXAEMNRtrainT2()`** — L1359 — `PORT void SetRXAEMNRtrainT2(int channel, double t2)`
  Sets rxaemnrtrain t2 — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/emnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
