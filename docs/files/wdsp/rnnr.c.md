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

### Functions

- `db_to_lin()` — L52
- `lin_to_db()` — L53
- `agc_alpha_ms()` — L65
- `rnnr_agc_init()` — L71
- `frame_rms()` — L89
- `ring_buffer_init()` — L105
- `ring_buffer_free()` — L114
- `ring_buffer_put()` — L122
- `ring_buffer_get_bulk()` — L132
- `ring_buffer_resize()` — L144
- `SetRXARNNRRun()` — L161
- `setSize_rnnr()` — L178
- `setBuffers_rnnr()` — L189
- `setSamplerate_rnnr()` — L195
- `create_rnnr()` — L201
- `xrnnr()` — L238
- `destroy_rnnr()` — L315
- `RNNRloadModel()` — L348
- `SetRXARNNRPosition()` — L387
- `SetRXARNNRUseDefaultGain()` — L396

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/rnnr.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
