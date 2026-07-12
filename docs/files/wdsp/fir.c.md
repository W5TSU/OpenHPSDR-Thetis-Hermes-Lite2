# `wdsp/fir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/bandpass.c` (calls ×12)
  - `wdsp/fmmod.c` (calls ×6)
  - `wdsp/fmd.c` (calls ×5)
  - `wdsp/eq.c` (calls ×3)
  - `wdsp/fcurve.c` (calls ×3)
  - `wdsp/firmin.c` (calls ×3)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/nbp.c` (calls ×2)
  - `wdsp/resample.c` (calls ×2)
  - `wdsp/cfir.c` (calls ×1)
  - `wdsp/delay.c` (calls ×1)
  - `wdsp/dexp.c` (calls ×1)
  - …and 3 more files
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×9)
  - `wdsp/impulse_cache.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `fir_bandpass()` (×33), `fir_fsamp()` (×4), `fftcv_mults()` (×3), `fir_fsamp_odd()` (×2), `mp_imp()` (×2), `analytic()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fftcv_mults()`** — L29 — `double* fftcv_mults (int NM, double* c_impulse)`
  Called by: `calc_bps()` (`wdsp/bandpass.c`), `eq_mults()` (`wdsp/eq.c`), `fc_mults()` (`wdsp/fcurve.c`)
- **`get_fsamp_window()`** — L44 — `double* get_fsamp_window(int N, int wintype)`
  Called by: `fir_fsamp_odd()` (same file), `fir_fsamp()` (same file)
- **`fir_fsamp_odd()`** — L83 — `double* fir_fsamp_odd (int N, double* A, int rtype, double scale, int wintype)`
  Called by: `eq_impulse()` (`wdsp/eq.c`), `fc_impulse()` (`wdsp/fcurve.c`)
- **`fir_fsamp()`** — L127 — `double* fir_fsamp (int N, double* A, int rtype, double scale, int wintype)`
  Called by: `cfir_impulse()` (`wdsp/cfir.c`), `eq_impulse()` (`wdsp/eq.c`), `fc_impulse()` (`wdsp/fcurve.c`), `icfir_impulse()` (`wdsp/icfir.c`)
- **`fir_bandpass()`** — L187 — `double* fir_bandpass (int N, double f_low, double f_high, double samplerate, int wintype, int rtype, double scale)`
  Called by: `calc_bps()` (`wdsp/bandpass.c`), `create_bandpass()` (`wdsp/bandpass.c`), `setSamplerate_bandpass()` (`wdsp/bandpass.c`), `setSize_bandpass()` (`wdsp/bandpass.c`), `setGain_bandpass()` (`wdsp/bandpass.c`), `CalcBandpassFilter()` (`wdsp/bandpass.c`) — and 27 more
- **`fir_read()`** — L288 — `double *fir_read (int N, const char *filename, int rtype, double scale)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`analytic()`** — L330 — `void analytic (int N, double* in, double* out)`
  Called by: `mp_imp()` (same file), `build_doublepole_1sided()` (`wdsp/doublepole.c`)
- **`mp_imp()`** — L357 — `void mp_imp (int N, double* fir, double* mpfir, int pfactor, int polarity)`
  Called by: `calc_fircore()` (`wdsp/firmin.c`), `analyze_bandpass_filter()` (`wdsp/utilities.c`)
- **`zff_impulse()`** — L442 — `double* zff_impulse(int nc, double scale)`
  impulse response of a zero frequency filter comprising a cascade of two resonators, each followed by a detrending filter
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
