# `wdsp/ssql.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** AM squelch, FM squelch, and syllabic (voice-detecting) squelch.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/cblock.c` (calls ×4)
  - `wdsp/iir.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_ssql()` (×1), `destroy_ssql()` (×1), `flush_ssql()` (×1), `xssql()` (×1), `setSamplerate_ssql()` (×1), `setBuffers_ssql()` (×1), `setSize_ssql()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_ftov()`** — L35 — `FTOV create_ftov (int run, int size, int rate, int rsize, double fmax, double* in, double* out)`
  Constructor for the `ftov` block: allocates its state/buffers and computes initial coefficients.
  Called by: `calc_ssql()` (same file)
- **`destroy_ftov()`** — L56 — `void destroy_ftov (FTOV a)`
  Destroys the `ftov` block, freeing its allocated buffers.
  Called by: `decalc_ssql()` (same file)
- **`flush_ftov()`** — L62 — `void flush_ftov (FTOV a)`
  Flushes (zeroes) the `ftov` block’s internal buffers/state.
  Called by: `flush_ssql()` (same file)
- **`xftov()`** — L70 — `void xftov (FTOV a)`
  Runs the `ftov` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xssql()` (same file)
- **`compute_ssql_slews()`** — L113 — `void compute_ssql_slews(SSQL a)`
  Called by: `calc_ssql()` (same file)
- **`calc_ssql()`** — L133 — `void calc_ssql (SSQL a)`
  Called by: `create_ssql()` (same file), `setBuffers_ssql()` (same file), `setSamplerate_ssql()` (same file), `setSize_ssql()` (same file)
- **`decalc_ssql()`** — L162 — `void decalc_ssql (SSQL a)`
  Called by: `destroy_ssql()` (same file), `setBuffers_ssql()` (same file), `setSamplerate_ssql()` (same file), `setSize_ssql()` (same file)
- **`create_ssql()`** — L177 — `SSQL create_ssql (int run, int size, double* in, double* out, int rate, double tup, double tdown, double muted_gain, double tau_mute, double tau_unmute, double wthresh, double tr_t`
  Constructor for the `ssql` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_ssql()`** — L202 — `void destroy_ssql (SSQL a)`
  Destroys the `ssql` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_ssql()`** — L208 — `void flush_ssql (SSQL a)`
  Flushes (zeroes) the `ssql` block’s internal buffers/state.
  Called by: `flush_rxa()` (`wdsp/RXA.c`)
- **`xssql()`** — L230 — `void xssql (SSQL a)`
  Runs the `ssql` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`setBuffers_ssql()`** — L302 — `void setBuffers_ssql (SSQL a, double* in, double* out)`
  Re-points the `ssql` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_ssql()`** — L310 — `void setSamplerate_ssql (SSQL a, int rate)`
  Reconfigures the `ssql` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`setSize_ssql()`** — L317 — `void setSize_ssql (SSQL a, int size)`
  Reconfigures the `ssql` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`SetRXASSQLRun()`** — L330 — `PORT void SetRXASSQLRun (int channel, int run)`
  Sets rxassqlrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASSQLThreshold()`** — L338 — `PORT void SetRXASSQLThreshold (int channel, double threshold)`
  Sets rxassqlthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASSQLTauMute()`** — L348 — `PORT void SetRXASSQLTauMute (int channel, double tau_mute)`
  Sets rxassqltau mute — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASSQLTauUnMute()`** — L360 — `PORT void SetRXASSQLTauUnMute (int channel, double tau_unmute)`
  Sets rxassqltau un mute — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/ssql.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
