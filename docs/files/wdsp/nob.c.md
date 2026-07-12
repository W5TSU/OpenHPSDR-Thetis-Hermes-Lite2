# `wdsp/nob.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Impulse noise blankers (NB and NB2).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/znob.c` (calls ×8)
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_anb()` (×2), `destroy_anb()` (×2), `xanb()` (×2), `pSetRCVRANBRun()` (×1), `pSetRCVRANBBuffsize()` (×1), `pSetRCVRANBSamplerate()` (×1), `pSetRCVRANBTau()` (×1), `pSetRCVRANBHangtime()` (×1)

## Outline

### Functions

- `initBlanker()` — L33
- `create_anb()` — L54
- `destroy_anb()` — L89
- `flush_anb()` — L99
- `xanb()` — L107
- `setBuffers_anb()` — L189
- `setSamplerate_anb()` — L195
- `setSize_anb()` — L201
- `pSetRCVRANBRun()` — L232
- `pSetRCVRANBBuffsize()` — L240
- `pSetRCVRANBSamplerate()` — L248
- `pSetRCVRANBTau()` — L257
- `pSetRCVRANBHangtime()` — L266
- `pSetRCVRANBAdvtime()` — L275
- `pSetRCVRANBBacktau()` — L284
- `pSetRCVRANBThreshold()` — L293
- `create_anbEXT()` — L310
- `destroy_anbEXT()` — L326
- `flush_anbEXT()` — L332
- `xanbEXT()` — L338
- `SetEXTANBRun()` — L347
- `SetEXTANBBuffsize()` — L356
- `SetEXTANBSamplerate()` — L365
- `SetEXTANBTau()` — L375
- `SetEXTANBHangtime()` — L385
- `SetEXTANBAdvtime()` — L395
- `SetEXTANBBacktau()` — L405
- `SetEXTANBThreshold()` — L415
- `xanbEXTF()` — L430

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nob.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
