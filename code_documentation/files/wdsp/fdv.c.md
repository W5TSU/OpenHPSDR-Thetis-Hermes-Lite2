# `wdsp/fdv.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** FreeDV 700E RX decode block. Sits post-AGC in the RXA chain; resamples to/from the modem's 8 kHz rate, normalises blocks into `libcodec2`'s 16-bit domain via a smoothed AGC, drives `freedv_rx()` per `freedv_nin()`-sized block, and passes raw modem audio through until synced/primed so the signal stays audible for tuning. RADE V1's equivalent decode block lives in ChannelMaster (`radae.c`, §8) rather than here, since it uses a separate native library (`rade_c`) instead of `libcodec2`.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/resample.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_fdv()` (×1), `destroy_fdv()` (×1), `flush_fdv()` (×1), `xfdv()` (×1), `setSamplerate_fdv()` (×1), `setBuffers_fdv()` (×1), `setSize_fdv()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fdv_rb_init()`** — L74 — `static void fdv_rb_init(fdv_ring_buffer* rb, int capacity)`
  ringbuffer (same scheme as rnnr.c)
  Called by: `fdv_alloc_streams()` (same file)
- **`fdv_rb_free()`** — L83 — `static void fdv_rb_free(fdv_ring_buffer* rb)`
  Called by: `fdv_free_streams()` (same file)
- **`fdv_rb_clear()`** — L91 — `static void fdv_rb_clear(fdv_ring_buffer* rb)`
  Called by: `fdv_reset()` (same file)
- **`fdv_rb_put()`** — L96 — `static void fdv_rb_put(fdv_ring_buffer* rb, float v)`
  Called by: `xfdv()` (same file)
- **`fdv_rb_get_bulk()`** — L106 — `static int fdv_rb_get_bulk(fdv_ring_buffer* rb, float* dest, int n)`
  Called by: `xfdv()` (same file)
- **`fdv_alloc_streams()`** — L120 — `static void fdv_alloc_streams(FDV a)`
  resamplers and rings depend on size/rate; built here so the rate and size setters can simply tear down and rebuild around the freedv handle
  Called by: `create_fdv()` (same file), `setSize_fdv()` (same file), `setSamplerate_fdv()` (same file)
- **`fdv_free_streams()`** — L139 — `static void fdv_free_streams(FDV a)`
  Called by: `destroy_fdv()` (same file), `setSize_fdv()` (same file), `setSamplerate_fdv()` (same file)
- **`fdv_reset()`** — L153 — `static void fdv_reset(FDV a)`
  Called by: `flush_fdv()` (same file)
- **`create_fdv()`** — L165 — `FDV create_fdv(int run, int size, double* in, double* out, int rate)`
  Constructor for the `fdv` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_fdv()`** — L193 — `void destroy_fdv(FDV a)`
  Destroys the `fdv` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`flush_fdv()`** — L207 — `void flush_fdv(FDV a)`
  Flushes (zeroes) the `fdv` block’s internal buffers/state.
  Called by: `SetRXAFDVRun()` (same file), `flush_rxa()` (`wdsp/RXA.c`)
- **`setBuffers_fdv()`** — L214 — `void setBuffers_fdv(FDV a, double* in, double* out)`
  Re-points the `fdv` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSize_fdv()`** — L220 — `void setSize_fdv(FDV a, int size)`
  Reconfigures the `fdv` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`)
- **`setSamplerate_fdv()`** — L229 — `void setSamplerate_fdv(FDV a, int rate)`
  Reconfigures the `fdv` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`)
- **`fdv_block_rms()`** — L238 — `static float fdv_block_rms(const float* x, int n)`
  Called by: `xfdv()` (same file)
- **`xfdv()`** — L246 — `void xfdv(FDV a)`
  Runs the `fdv` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`)
- **`SetRXAFDVRun()`** — L496 — `PORT void SetRXAFDVRun(int channel, int run)`
  Sets rxafdvrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAFDVSync()`** — L509 — `PORT int GetRXAFDVSync(int channel)`
  Returns rxafdvsync — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRXAFDVSnr()`** — L515 — `PORT double GetRXAFDVSnr(int channel)`
  Returns rxafdvsnr — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`ResetRXAFDVDebug()`** — L526 — `PORT void ResetRXAFDVDebug(void)`
  W5TSU: DEBUG - temporary diagnostic dump control, remove before merge. Call at the start of a Quick-Play test session so fdv_debug.txt/ fdv_debug_audio.raw/fdv_debug_resamp.raw capture that run instead of staying silent because an earlier run in the same process already used up the counters.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/fdv.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
