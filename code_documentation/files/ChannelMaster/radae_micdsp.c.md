# `ChannelMaster/radae_micdsp.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Mic-path DSP helpers (biquad EQ, RNNoise, EBU R128 loudness normalisation, AGC) prepared for a future RADE V1 TX path — not yet wired to `xradae_tx` as of this branch.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/radae.c` (calls ×11)
- Uses (outgoing references to other files):
  - `ChannelMaster/radae_micdsp.h` (imports ×1)
- Most-referenced symbols from other files: `radae_micdsp_create()` (×1), `radae_micdsp_destroy()` (×1), `radae_micdsp_set_rnnoise_enabled()` (×1), `radae_micdsp_set_agc_enabled()` (×1), `radae_micdsp_set_agc_target_lufs()` (×1), `radae_micdsp_set_eq_enabled()` (×1), `radae_micdsp_set_eq_bass()` (×1), `radae_micdsp_set_eq_mid()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`biquad_reset()`** — L44 — `static void biquad_reset(biquad_t* b)`
  Called by: `radae_micdsp_reset()` (same file)
- **`biquad_clear()`** — L49 — `static void biquad_clear(biquad_t* b)`
  Called by: `biquad_design_peaking()` (same file), `biquad_design_lowshelf()` (same file), `biquad_design_highshelf()` (same file), `radae_micdsp_create()` (same file)
- **`biquad_process()`** — L60 — `static __forceinline float biquad_process(biquad_t* b, float in)`
  Process one sample through a Direct-Form-II Transposed biquad. The bypass branch is sample-cheap and lets the chain be wired unconditionally without a per-sample if. */
  Called by: `eq_step_block()` (same file)
- **`biquad_design_peaking()`** — L73 — `static void biquad_design_peaking(biquad_t* b, double sr, double f0, double gain_db, double q)`
  peakingEQ at f0 / Q with gain dB. Used for the mid band. */
  Called by: `eq_rebuild_if_dirty()` (same file)
- **`biquad_design_lowshelf()`** — L93 — `static void biquad_design_lowshelf(biquad_t* b, double sr, double f0, double gain_db)`
  lowshelf at f0 with gain dB, slope S=1.0 (matches SoX `bass`). */
  Called by: `eq_rebuild_if_dirty()` (same file)
- **`biquad_design_highshelf()`** — L116 — `static void biquad_design_highshelf(biquad_t* b, double sr, double f0, double gain_db)`
  highshelf at f0 with gain dB, slope S=1.0 (matches SoX `treble`). */
  Called by: `eq_rebuild_if_dirty()` (same file)
- **`biquad_design_gain()`** — L140 — `static void biquad_design_gain(biquad_t* b, double gain_db)`
  Linear gain biquad (b0 = gain, all other = 0). Trivially stable; Direct-Form path through biquad_process produces y = b0 * x. */
  Called by: `eq_rebuild_if_dirty()` (same file)
- **`agc_supported_rate()`** — L218 — `static int agc_supported_rate(int sr)`
  Called by: `agc_init_engines()` (same file)
- **`agc_init_engines()`** — L223 — `static void agc_init_engines(void)`
  Called by: `radae_micdsp_create()` (same file)
- **`agc_soft_reset()`** — L281 — `static void agc_soft_reset(void)`
  Soft reset -- mirrors freedv-gui AgcStep::reset() exactly: zero the smoothed-gain accumulators and clear the input FIFO, but leave the existing ebur128_state and WebRtcAgc instances allocated so their internal histories (loudness ring, limiter envelope) survive the reset. Our output FIFO is also…
  Called by: `radae_micdsp_reset()` (same file), `radae_micdsp_set_agc_enabled()` (same file)
- **`radae_micdsp_create()`** — L289 — `void radae_micdsp_create(int sample_rate)`
  Called by: `create_radae()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_destroy()`** — L314 — `void radae_micdsp_destroy(void)`
  Called by: `destroy_radae()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_reset()`** — L322 — `void radae_micdsp_reset(void)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`radae_micdsp_set_rnnoise_enabled()`** — L350 — `void radae_micdsp_set_rnnoise_enabled(int e)`
  Called by: `SetRadaeMicRNNoiseEnabled()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_agc_enabled()`** — L356 — `void radae_micdsp_set_agc_enabled(int e)`
  Called by: `SetRadaeMicAGCEnabled()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_agc_target_lufs()`** — L368 — `void radae_micdsp_set_agc_target_lufs(double t)`
  Called by: `SetRadaeMicAGCTargetLufs()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_eq_enabled()`** — L375 — `void radae_micdsp_set_eq_enabled(int e)`
  Called by: `SetRadaeMicEQEnabled()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_eq_bass()`** — L381 — `void radae_micdsp_set_eq_bass(double f, double g)`
  Called by: `SetRadaeMicEQBass()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_eq_mid()`** — L388 — `void radae_micdsp_set_eq_mid(double f, double g, double q)`
  Called by: `SetRadaeMicEQMid()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_eq_treble()`** — L396 — `void radae_micdsp_set_eq_treble(double f, double g)`
  Called by: `SetRadaeMicEQTreble()` (`ChannelMaster/radae.c`)
- **`radae_micdsp_set_eq_vol()`** — L403 — `void radae_micdsp_set_eq_vol(double db)`
  Called by: `SetRadaeMicEQVol()` (`ChannelMaster/radae.c`)
- **`eq_rebuild_if_dirty()`** — L413 — `static void eq_rebuild_if_dirty(void)`
  Called by: `radae_micdsp_process()` (same file)
- **`float_to_s16()`** — L464 — `static __forceinline int16_t float_to_s16(float f)`
  Called by: `agc_step_block()` (same file)
- **`s16_to_float()`** — L472 — `static __forceinline float s16_to_float(int16_t v)`
  Called by: `agc_step_block()` (same file)
- **`agc_run_one_block()`** — L477 — `static int agc_run_one_block(int16_t* in_block, int16_t* out_block, int n)`
  Called by: `agc_step_block()` (same file)
- **`agc_step_block()`** — L546 — `static int agc_step_block(float* buf, int n)`
  Returns the number of float samples actually written to buf. May be less than n during startup priming (AGC drains in 10 ms blocks, so up to ~10 ms of input is buffered before any output appears). Matches freedv-gui AgcStep::execute() semantics: the caller absorbs the variable rate. Do NOT…
  Called by: `radae_micdsp_process()` (same file)
- **`eq_step_block()`** — L606 — `static void eq_step_block(float* buf, int n)`
  Called by: `radae_micdsp_process()` (same file)
- **`rnnoise_step_block()`** — L632 — `static int rnnoise_step_block(float* buf, int n)`
  Called by: `radae_micdsp_process()` (same file)
- **`radae_micdsp_process()`** — L694 — `int radae_micdsp_process(float* buf, int n_in)`
  Called by: `xradae_tx()` (`ChannelMaster/radae.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/radae_micdsp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
