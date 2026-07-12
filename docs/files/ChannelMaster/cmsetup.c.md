# `ChannelMaster/cmsetup.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** System-wide setup: instantiates buffers, mixers, VAC, analyzers per radio model.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×30)
  - `ChannelMaster/pipe.c` (calls ×7)
  - `ChannelMaster/analyzers.c` (calls ×3)
  - `ChannelMaster/sync.c` (calls ×3)
  - `ChannelMaster/networkproto1.c` (calls ×2)
  - `ChannelMaster/ivac.c` (calls ×1)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
  - `ChannelMaster/pipe.c` (calls ×2)
  - `ChannelMaster/sync.c` (calls ×2)
  - `ChannelMaster/cmaster.h` (imports ×1)
  - `ChannelMaster/cmsetup.h` (imports ×1)
- Most-referenced symbols from other files: `inid()` (×15), `chid()` (×10), `getbuffsize()` (×8), `txid()` (×4), `rxid()` (×3), `stype()` (×3), `sp0id()` (×2), `mixinid()` (×2)

## Outline

### Functions

- `SetRadioStructure()` — L32
- `set_cmdefault_rates()` — L59
- `CreateRadio()` — L88
- `DestroyRadio()` — L96
- `getbuffsize()` — L105
- `getInputRate()` — L113
- `getChannelOutputRate()` — L119
- `rxid()` — L147
- `txid()` — L152
- `sp0id()` — L157
- `stype()` — L162
- `chid()` — L175
- `inid()` — L194
- `mixinid()` — L216

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmsetup.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
