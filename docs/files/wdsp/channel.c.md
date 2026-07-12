# `wdsp/channel.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Channel object lifecycle (create/destroy/run) and DLL entry points.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×8)
  - `wdsp/RXA.c` (calls ×1)
  - `wdsp/TXA.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/main.c` (calls ×10)
  - `wdsp/iobuffs.c` (calls ×9)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `OpenChannel()` (×2), `CloseChannel()` (×2), `SetOutputSamplerate()` (×2), `SetChannelState()` (×2), `SetInputBuffsize()` (×1), `SetInputSamplerate()` (×1)

## Outline

### Functions

- `start_thread()` — L31
- `pre_main_build()` — L37
- `post_main_build()` — L60
- `build_channel()` — L68
- `OpenChannel()` — L75
- `pre_main_destroy()` — L103
- `post_main_destroy()` — L113
- `CloseChannel()` — L120
- `flushChannel()` — L128
- `SetType()` — L156
- `SetInputBuffsize()` — L167
- `SetDSPBuffsize()` — L180
- `SetInputSamplerate()` — L196
- `SetDSPSamplerate()` — L210
- `SetOutputSamplerate()` — L226
- `SetAllRates()` — L240
- `SetChannelState()` — L258
- `SetChannelTDelayUp()` — L299
- `SetChannelTSlewUp()` — L311
- `SetChannelTDelayDown()` — L323
- `SetChannelTSlewDown()` — L335

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/channel.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
