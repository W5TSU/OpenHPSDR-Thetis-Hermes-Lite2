# `wdsp/utilities.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/iir.c` (calls ×14)
  - `wdsp/eq.c` (calls ×10)
  - `wdsp/fir.c` (calls ×9)
  - `wdsp/firmin.c` (calls ×7)
  - `ChannelMaster/aamix.c` (calls ×5)
  - `wdsp/RXA.c` (calls ×5)
  - `wdsp/TXA.c` (calls ×5)
  - `wdsp/dexp.c` (calls ×5)
  - `wdsp/rmatch.c` (calls ×5)
  - `wdsp/cfcomp.c` (calls ×4)
  - `wdsp/nbp.c` (calls ×4)
  - `wdsp/rnnr.c` (calls ×4)
  - …and 62 more files
- Uses (outgoing references to other files):
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `malloc0()` (×180)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`malloc0()`** — L36 — `PORT void *malloc0 (int size)`
  Called by: `analyze_bandpass_filter()` (same file), `WriteAudioWDSP()` (same file), `WriteScaledAudioFile()` (same file), `WriteScaledAudio()` (same file), `model_bandpass()` (same file), `create_bfcu()` (same file) — and 181 more
- **`NewCriticalSection()`** — L47 — `PORT void *NewCriticalSection()`
  Called from C# via P/Invoke — declared/wrapped in `Console/win32.cs`.
- **`DestroyCriticalSection()`** — L55 — `PORT void DestroyCriticalSection (LPCRITICAL_SECTION cs_ptr)`
  Called from C# via P/Invoke — declared/wrapped in `Console/win32.cs`.
- **`print_impulse()`** — L68 — `void print_impulse (const char* filename, int N, double* impulse, int rtype, int pr_mode)`
  Called by: `analyze_bandpass_filter()` (same file)
- **`analyze_bandpass_filter()`** — L90 — `PORT void analyze_bandpass_filter (int N, double f_low, double f_high, double samplerate, int wintype, int rtype, double scale)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_peak_val()`** — L103 — `void print_peak_val (const char* filename, int N, double* buff, double thresh)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_peak_env()`** — L123 — `void print_peak_env (const char* filename, int N, double* buff, double thresh)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_peak_env_f2()`** — L147 — `void print_peak_env_f2 (const char* filename, int N, float* Ibuff, float* Qbuff)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_iqc_values()`** — L166 — `void print_iqc_values (const char* filename, int state, double env_in, double I, double Q, double ym, double yc, double ys, double thresh)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_buffer_parameters()`** — L186 — `PORT void print_buffer_parameters (const char* filename, int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_meter()`** — L217 — `void print_meter (const char* filename, double* meter, int enum_av, int enum_pk, int enum_gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_message()`** — L231 — `void print_message (const char* filename, const char* message, int p0, int p1, int p2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_window_gain()`** — L243 — `void print_window_gain (const char* filename, int wintype, double inv_coherent_gain, double inherent_power_gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_deviation()`** — L281 — `void print_deviation (const char* filename, double dpmax, double rate)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CalccPrintSamples()`** — L293 — `void __cdecl CalccPrintSamples (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`doCalccPrintSamples()`** — L317 — `void doCalccPrintSamples(int channel)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`print_anb_parms()`** — L322 — `void print_anb_parms (const char* filename, ANB a)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`WriteAudioFile()`** — L346 — `void WriteAudioFile(void* arg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`WriteAudioWDSP()`** — L372 — `void WriteAudioWDSP (double seconds, int rate, int size, double* indata, int mode, double gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`WriteScaledAudioFile()`** — L425 — `void WriteScaledAudioFile (void* arg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`WriteScaledAudio()`** — L461 — `void WriteScaledAudio ( double seconds, int rate, int size, double* indata )`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`model_bandpass()`** — L506 — `double* model_bandpass(int nc, double f_low, double f_high, double rate, int wtype, int points)`
  Called by: `create_bfcu()` (same file)
- **`print_bandpass_response()`** — L536 — `void print_bandpass_response (const char* filename, int points, double* response)`
  Called by: `test_bfcu()` (same file)
- **`create_bfcu()`** — L551 — `PORT int create_bfcu(int id, int min_size, int max_size, double rate, double corner, int points)`
  Constructor for the `bfcu` block: allocates its state/buffers and computes initial coefficients.
  Called by: `test_bfcu()` (same file)
- **`destroy_bfcu()`** — L588 — `PORT void destroy_bfcu(int id)`
  Destroys the `bfcu` block, freeing its allocated buffers.
  Called by: `test_bfcu()` (same file)
- **`getFilterCorners()`** — L603 — `PORT void getFilterCorners(int id, int* lower_index, int* upper_index)`
  Called by: `test_bfcu()` (same file)
- **`getFilterCurve()`** — L612 — `PORT void getFilterCurve(int id, int size, int w_type, int index_low, int index_high, double* segment)`
  Called by: `test_bfcu()` (same file)
- **`test_bfcu()`** — L625 — `void test_bfcu()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/utilities.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
