# `wdsp/rnnr.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** RNNoise neural-network noise reduction "NR3" (uses `lib/NR_Algorithms_x64`).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×6)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/RXA.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_rnnr()` (×1), `destroy_rnnr()` (×1), `xrnnr()` (×1), `setSamplerate_rnnr()` (×1), `setBuffers_rnnr()` (×1), `setSize_rnnr()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`db_to_lin()`** — L52 — `static inline float db_to_lin(float db)`
  Called by: `rnnr_agc_init()` (same file), `xrnnr()` (same file)
- **`lin_to_db()`** — L53 — `static inline float lin_to_db(float lin)`
  Called by: `rnnr_agc_init()` (same file), `xrnnr()` (same file), `SetRXARNNRUseDefaultGain()` (same file)
- **`agc_alpha_ms()`** — L65 — `static float agc_alpha_ms(float ms, float frame_hz)`
  Called by: `rnnr_agc_init()` (same file)
- **`rnnr_agc_init()`** — L71 — `void rnnr_agc_init(RNNR a)`
  Called by: `setSamplerate_rnnr()` (same file), `create_rnnr()` (same file)
- **`frame_rms()`** — L89 — `static float frame_rms(const float* x, int n)`
  Called by: `xrnnr()` (same file)
- **`ring_buffer_init()`** — L105 — `static void ring_buffer_init(rnnr_ring_buffer* rb, int capacity)`
  ringbuffer
  Called by: `create_rnnr()` (same file)
- **`ring_buffer_free()`** — L114 — `static void ring_buffer_free(rnnr_ring_buffer* rb)`
  Called by: `destroy_rnnr()` (same file)
- **`ring_buffer_put()`** — L122 — `static void ring_buffer_put(rnnr_ring_buffer* rb, float v)`
  Called by: `xrnnr()` (same file)
- **`ring_buffer_get_bulk()`** — L132 — `static int ring_buffer_get_bulk(rnnr_ring_buffer* rb, float* dest, int n)`
  Called by: `xrnnr()` (same file)
- **`ring_buffer_resize()`** — L144 — `static void ring_buffer_resize(rnnr_ring_buffer* rb, int new_capacity)`
  Called by: `setSize_rnnr()` (same file)
- **`SetRXARNNRRun()`** — L161 — `PORT void SetRXARNNRRun (int channel, int run)`
  Sets rxarnnrrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`setSize_rnnr()`** — L178 — `void setSize_rnnr(RNNR a, int size)`
  Reconfigures the `rnnr` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setBuffers_rnnr()`** — L189 — `void setBuffers_rnnr(RNNR a, double* in, double* out)`
  Re-points the `rnnr` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_rnnr()`** — L195 — `void setSamplerate_rnnr(RNNR a, int rate)`
  Reconfigures the `rnnr` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`create_rnnr()`** — L201 — `RNNR create_rnnr(int run, int position, int size, double* in, double* out, int rate)`
  Constructor for the `rnnr` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`xrnnr()`** — L238 — `void xrnnr(RNNR a, int pos)`
  Runs the `rnnr` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`destroy_rnnr()`** — L315 — `void destroy_rnnr(RNNR a)`
  Destroys the `rnnr` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`RNNRloadModel()`** — L348 — `PORT void RNNRloadModel(const char* file_path)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXARNNRPosition()`** — L387 — `PORT void SetRXARNNRPosition(int channel, int position)`
  Sets rxarnnrposition — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXARNNRUseDefaultGain()`** — L396 — `PORT void SetRXARNNRUseDefaultGain(int channel, int use_default_gain)`
  Sets rxarnnruse default gain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/rnnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
