# `wdsp/resample.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×15)
  - `wdsp/TXA.c` (calls ×15)
  - `ChannelMaster/aamix.c` (calls ×8)
  - `wdsp/snb.c` (calls ×6)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/fir.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `setBuffers_resample()` (×8), `create_resample()` (×6), `destroy_resample()` (×6), `setSize_resample()` (×6), `xresample()` (×4), `flush_resample()` (×4), `setInRate_resample()` (×4), `setOutRate_resample()` (×4)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_resample()`** — L35 — `void calc_resample (RESAMPLE a)`
  Called by: `create_resample()` (same file), `setInRate_resample()` (same file), `setOutRate_resample()` (same file), `setFCLow_resample()` (same file), `setBandwidth_resample()` (same file)
- **`decalc_resample()`** — L80 — `void decalc_resample (RESAMPLE a)`
  Called by: `destroy_resample()` (same file), `setInRate_resample()` (same file), `setOutRate_resample()` (same file), `setFCLow_resample()` (same file), `setBandwidth_resample()` (same file)
- **`create_resample()`** — L86 — `PORT RESAMPLE create_resample ( int run, int size, double* in, double* out, int in_rate, int out_rate, double fc, int ncoef, double gain)`
  Constructor for the `resample` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_resampleV()` (same file), `create_aamix()` (`ChannelMaster/aamix.c`), `SetAAudioOutRate()` (`ChannelMaster/aamix.c`), `SetAAudioStreamRate()` (`ChannelMaster/aamix.c`), `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`) — and 1 more
- **`destroy_resample()`** — L105 — `PORT void destroy_resample (RESAMPLE a)`
  Destroys the `resample` block, freeing its allocated buffers.
  Called by: `destroy_resampleV()` (same file), `destroy_aamix()` (`ChannelMaster/aamix.c`), `SetAAudioOutRate()` (`ChannelMaster/aamix.c`), `SetAAudioStreamRate()` (`ChannelMaster/aamix.c`), `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`) — and 1 more
- **`flush_resample()`** — L112 — `PORT void flush_resample (RESAMPLE a)`
  Flushes (zeroes) the `resample` block’s internal buffers/state.
  Called by: `setSize_resample()` (same file), `flush_mix_ring()` (`ChannelMaster/aamix.c`), `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`), `flush_snba()` (`wdsp/snb.c`)
- **`xresample()`** — L120 — `PORT int xresample(RESAMPLE a)`
  Runs the `resample` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xresampleV()` (same file), `xMixAudio()` (`ChannelMaster/aamix.c`), `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`), `xsnba()` (`wdsp/snb.c`)
- **`setBuffers_resample()`** — L166 — `void setBuffers_resample(RESAMPLE a, double* in, double* out)`
  Re-points the `resample` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`), `setOutputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setInputSamplerate_txa()` (`wdsp/TXA.c`), `setOutputSamplerate_txa()` (`wdsp/TXA.c`) — and 2 more
- **`setSize_resample()`** — L172 — `void setSize_resample(RESAMPLE a, int size)`
  Reconfigures the `resample` block for a new buffer size.
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setInputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setInRate_resample()`** — L178 — `void setInRate_resample(RESAMPLE a, int rate)`
  Called by: `setInputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setInputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setOutRate_resample()`** — L185 — `void setOutRate_resample(RESAMPLE a, int rate)`
  Called by: `setOutputSamplerate_rxa()` (`wdsp/RXA.c`), `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setOutputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setFCLow_resample()`** — L192 — `void setFCLow_resample (RESAMPLE a, double fc_low)`
  Called by: `calc_snba()` (`wdsp/snb.c`)
- **`setBandwidth_resample()`** — L202 — `void setBandwidth_resample (RESAMPLE a, double fc_low, double fc_high)`
  Called by: `SetRXASNBAOutputBandwidth()` (`wdsp/snb.c`)
- **`create_resampleV()`** — L215 — `PORT void* create_resampleV (int in_rate, int out_rate)`
  Constructor for the `resampleV` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xresampleV()`** — L221 — `PORT void xresampleV (double* input, double* output, int numsamps, int* outsamps, void* ptr)`
  Runs the `resampleV` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_resampleV()`** — L231 — `PORT void destroy_resampleV (void* ptr)`
  Destroys the `resampleV` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_resampleF()`** — L243 — `RESAMPLEF create_resampleF ( int run, int size, float* in, float* out, int in_rate, int out_rate)`
  Constructor for the `resampleF` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_resampleFV()` (same file)
- **`destroy_resampleF()`** — L289 — `void destroy_resampleF (RESAMPLEF a)`
  Destroys the `resampleF` block, freeing its allocated buffers.
  Called by: `destroy_resampleFV()` (same file)
- **`flush_resampleF()`** — L296 — `void flush_resampleF (RESAMPLEF a)`
  Flushes (zeroes) the `resampleF` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xresampleF()`** — L303 — `int xresampleF (RESAMPLEF a)`
  Runs the `resampleF` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xresampleFV()` (same file)
- **`create_resampleFV()`** — L341 — `PORT void* create_resampleFV (int in_rate, int out_rate)`
  Constructor for the `resampleFV` block: allocates its state/buffers and computes initial coefficients.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`xresampleFV()`** — L347 — `PORT void xresampleFV (float* input, float* output, int numsamps, int* outsamps, void* ptr)`
  Runs the `resampleFV` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`destroy_resampleFV()`** — L357 — `PORT void destroy_resampleFV (void* ptr)`
  Destroys the `resampleFV` block, freeing its allocated buffers.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/resample.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
