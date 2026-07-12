# `ChannelMaster/aamix.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Audio mixers (monitor mix, multi-RX audio combination) with per-input gain and slew.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×15)
  - `ChannelMaster/tci.c` (calls ×14)
  - `ChannelMaster/ivac.c` (calls ×12)
- Uses (outgoing references to other files):
  - `wdsp/resample.c` (calls ×8)
  - `wdsp/utilities.c` (calls ×5)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `create_aamix()` (×7), `destroy_aamix()` (×6), `SetAAudioMixState()` (×5), `SetAAudioMixWhat()` (×4), `SetAAudioStreamRate()` (×4), `xMixAudio()` (×3), `SetAAudioMixVol()` (×3), `SetAAudioRingInsize()` (×2)

## Outline

### Functions

- `mix_main()` — L32
- `start_mixthread()` — L51
- `create_aaslew()` — L69
- `destroy_aaslew()` — L109
- `flush_aaslew()` — L117
- `create_aamix()` — L128
- `destroy_aamix()` — L209
- `xMixAudio()` — L237
- `upslew()` — L280
- `downslew()` — L345
- `xaamix()` — L423
- `flush_mix_ring()` — L461
- `close_mixer()` — L471
- `open_mixer()` — L493
- `SetAAudioMixOutputPointer()` — L513
- `SetAAudioMixState()` — L521
- `SetAAudioMixStates()` — L554
- `SetAAudioMixWhat()` — L584
- `SetAAudioMixVolume()` — L596
- `SetAAudioMixVol()` — L610
- `SetAAudioRingInsize()` — L622
- `SetAAudioRingOutsize()` — L644
- `SetAAudioOutRate()` — L656
- `SetAAudioStreamRate()` — L683

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/aamix.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
