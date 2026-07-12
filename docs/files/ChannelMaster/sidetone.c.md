# `ChannelMaster/sidetone.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** CW sidetone generation.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/netInterface.c` (calls ×5)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `keySidetone()` (×2), `create_sidetone()` (×1), `destroy_sidetone()` (×1), `xsidetone()` (×1), `setSidetoneRate()` (×1), `setSidetoneSize()` (×1), `SetSidetoneWPM()` (×1), `SetSidetoneRun()` (×1)

## Outline

### Functions

- `calc_tone1()` — L31
- `calc_rising_edge()` — L41
- `decalc_rising_edge()` — L60
- `calc_falling_edge()` — L65
- `decalc_falling_edge()` — L84
- `calc_wpm_times()` — L89
- `calc_sidetone()` — L98
- `decalc_sidetone()` — L108
- `create_sidetone()` — L114
- `destroy_sidetone()` — L157
- `osc_init()` — L172
- `osc()` — L178
- `xsidetone()` — L193
- `setSidetoneRate()` — L290
- `setSidetoneSize()` — L300
- `SetSidetoneSelectKey()` — L308
- `keySidetone()` — L317
- `makedotSidetone()` — L327
- `makedashSidetone()` — L336
- `SetCWtxIQpolarity()` — L345
- `SetSidetoneVolume()` — L354
- `SetCWtxVolume()` — L363
- `SetSidetoneWPM()` — L373
- `SetSidetoneRun()` — L383
- `SetCWtxRun()` — L392
- `SetSidetonePitch()` — L401
- `SetSidetoneEdgetype()` — L411
- `SetSidetoneEdgelength()` — L424

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/sidetone.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
