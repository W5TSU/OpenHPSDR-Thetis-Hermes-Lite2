# `ChannelMaster/znobII.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/nobII.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRCVRNOBBuffsize()` (×1), `SetRCVRNOBSamplerate()` (×1)

## Outline

### Functions

- `GetRCVRNOBPointer()` — L29
- `SetRCVRNOBRun()` — L44
- `SetRCVRNOBMode()` — L50
- `SetRCVRNOBBuffsize()` — L56
- `SetRCVRNOBSamplerate()` — L62
- `SetRCVRNOBTau()` — L68
- `SetRCVRNOBHangtime()` — L74
- `SetRCVRNOBAdvtime()` — L80
- `SetRCVRNOBBacktau()` — L86
- `SetRCVRNOBThreshold()` — L92

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/znobII.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
