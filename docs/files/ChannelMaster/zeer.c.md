# `ChannelMaster/zeer.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/eer.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/ilv.c` (calls ×1)

## Outline

### Functions

- `SetEERRun()` — L29
- `SetEERAMIQ()` — L38
- `SetEERMgain()` — L45
- `SetEERPgain()` — L52
- `SetEERRunDelays()` — L59
- `SetEERMdelay()` — L66
- `SetEERPdelay()` — L73
- `SetEERSize()` — L80
- `SetEERSamplerate()` — L87

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/zeer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
