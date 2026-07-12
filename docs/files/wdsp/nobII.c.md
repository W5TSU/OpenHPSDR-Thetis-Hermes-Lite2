# `wdsp/nobII.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Impulse noise blankers (NB and NB2).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/znobII.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_nob()` (×2), `destroy_nob()` (×2), `xnob()` (×2), `pSetRCVRNOBRun()` (×1), `pSetRCVRNOBMode()` (×1), `pSetRCVRNOBBuffsize()` (×1), `pSetRCVRNOBSamplerate()` (×1), `pSetRCVRNOBTau()` (×1)

## Outline

### Functions

- `init_nob()` — L36
- `create_nob()` — L63
- `destroy_nob()` — L126
- `flush_nob()` — L140
- `xnob()` — L157
- `setBuffers_nob()` — L496
- `setSamplerate_nob()` — L502
- `setSize_nob()` — L508
- `pSetRCVRNOBRun()` — L520
- `pSetRCVRNOBMode()` — L528
- `pSetRCVRNOBBuffsize()` — L536
- `pSetRCVRNOBSamplerate()` — L544
- `pSetRCVRNOBTau()` — L553
- `pSetRCVRNOBHangtime()` — L563
- `pSetRCVRNOBAdvtime()` — L572
- `pSetRCVRNOBBacktau()` — L581
- `pSetRCVRNOBThreshold()` — L590
- `create_nobEXT()` — L607
- `destroy_nobEXT()` — L627
- `flush_nobEXT()` — L633
- `xnobEXT()` — L639
- `SetEXTNOBRun()` — L648
- `SetEXTNOBMode()` — L657
- `SetEXTNOBBuffsize()` — L666
- `SetEXTNOBSamplerate()` — L675
- `SetEXTNOBTau()` — L685
- `SetEXTNOBHangtime()` — L696
- `SetEXTNOBAdvtime()` — L706
- `SetEXTNOBBacktau()` — L716
- `SetEXTNOBThreshold()` — L726
- `xnobEXTF()` — L741

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/nobII.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
