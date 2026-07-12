# `ChannelMaster/znob.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/nob.c` (calls ×8)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRCVRANBBuffsize()` (×1), `SetRCVRANBSamplerate()` (×1)

## Outline

### Functions

- `GetRCVRANBPointer()` — L29
- `SetRCVRANBRun()` — L44
- `SetRCVRANBBuffsize()` — L50
- `SetRCVRANBSamplerate()` — L56
- `SetRCVRANBTau()` — L62
- `SetRCVRANBHangtime()` — L68
- `SetRCVRANBAdvtime()` — L74
- `SetRCVRANBBacktau()` — L80
- `SetRCVRANBThreshold()` — L86

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/znob.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
