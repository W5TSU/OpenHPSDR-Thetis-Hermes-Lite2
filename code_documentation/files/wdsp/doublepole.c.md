# `wdsp/doublepole.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** IIR biquad sections (notches, peaking filters) and double-pole building blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/apfshadow.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/cmath.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `SetRXADoublepoleFreqs()` (×3), `SetRXADoublepoleGain()` (×2), `SetRXADoublepoleRun()` (×2), `create_doublepole()` (×1), `destroy_doublepole()` (×1), `flush_doublepole()` (×1), `xdoublepole()` (×1), `setSamplerate_doublepole()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_dpole_nc()`** — L30 — `static int calc_dpole_nc (double rate, double bandwidth)`
  Called by: `build_doublepole_1sided()` (same file), `build_doublepole_2sided()` (same file)
- **`build_doublepole_1sided()`** — L56 — `double* build_doublepole_1sided (int* nc, double rate, double fcenter, double bandwidth, double scale)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`build_doublepole_2sided()`** — L106 — `double* build_doublepole_2sided (int* nc, double rate, double fcenter, double bandwidth, double scale)`
  Called by: `create_doublepole()` (same file), `setSamplerate_doublepole()` (same file), `setSize_doublepole()` (same file), `setGain_doublepole()` (same file), `CalcDoublepoleFilter()` (same file)
- **`create_doublepole()`** — L146 — `DOUBLEPOLE create_doublepole (int run, int position, int size, double* in, double* out, double f_center, double bandwidth, int samplerate, double gain, int mode)`
  Constructor for the `doublepole` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_doublepole()`** — L168 — `void destroy_doublepole (DOUBLEPOLE a)`
  Destroys the `doublepole` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_doublepole()`** — L174 — `void flush_doublepole (DOUBLEPOLE a)`
  Flushes (zeroes) the `doublepole` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xdoublepole()`** — L179 — `void xdoublepole (DOUBLEPOLE a, int pos)`
  Runs the `doublepole` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_doublepole()`** — L200 — `void setBuffers_doublepole (DOUBLEPOLE a, double* in, double* out)`
  Re-points the `doublepole` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_doublepole()`** — L207 — `void setSamplerate_doublepole (DOUBLEPOLE a, int rate)`
  Reconfigures the `doublepole` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_doublepole()`** — L220 — `void setSize_doublepole (DOUBLEPOLE a, int size)`
  Reconfigures the `doublepole` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setGain_doublepole()`** — L232 — `void setGain_doublepole (DOUBLEPOLE a, double gain)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CalcDoublepoleFilter()`** — L242 — `void CalcDoublepoleFilter (DOUBLEPOLE a, double f_center, double bandwidth, double gain)`
  Called by: `SetRXADoublepoleFreqs()` (same file), `SetRXADoublepoleGain()` (same file)
- **`SetRXADoublepoleRun()`** — L267 — `PORT void SetRXADoublepoleRun (int channel, int run)`
  Sets rxadoublepole run — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWRun()` (`wdsp/apfshadow.c`)
- **`SetRXADoublepoleFreqs()`** — L276 — `PORT void SetRXADoublepoleFreqs (int channel, double f_center, double bandwidth)`
  Sets rxadoublepole freqs — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWFreq()` (`wdsp/apfshadow.c`), `SetRXASPCWBandwidth()` (`wdsp/apfshadow.c`)
- **`SetRXADoublepoleGain()`** — L285 — `PORT void SetRXADoublepoleGain (int channel, double gain)`
  Sets rxadoublepole gain — API setter, typically called from the console via P/Invoke.
  Called by: `SetRXASPCWSelection()` (`wdsp/apfshadow.c`), `SetRXASPCWGain()` (`wdsp/apfshadow.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/doublepole.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
