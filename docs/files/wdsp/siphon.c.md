# `wdsp/siphon.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Taps TX samples out of the chain (e.g., for the TX display).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
  - `ChannelMaster/pipe.c` (calls ×3)
  - `ChannelMaster/analyzers.c` (calls ×1)
  - `ChannelMaster/cmaster.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/analyzer.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/meterlog10.c` (calls ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_siphon()` (×2), `destroy_siphon()` (×2), `flush_siphon()` (×2), `xsiphon()` (×2), `setSamplerate_siphon()` (×2), `setBuffers_siphon()` (×2), `setSize_siphon()` (×2), `TXASetSipAllocDisps()` (×1)

## Outline

### Functions

- `build_window()` — L29
- `create_siphon()` — L53
- `destroy_siphon()` — L80
- `flush_siphon()` — L93
- `xsiphon()` — L101
- `setBuffers_siphon()` — L139
- `setSamplerate_siphon()` — L144
- `setSize_siphon()` — L149
- `suck()` — L155
- `sip_spectrum()` — L172
- `RXAGetaSipF()` — L189
- `RXAGetaSipF1()` — L204
- `TXASetSipPosition()` — L226
- `TXASetSipMode()` — L235
- `TXASetSipDisplay()` — L244
- `TXAGetaSipF()` — L253
- `TXAGetaSipF1()` — L268
- `TXASetSipSpecmode()` — L284
- `TXAGetSpecF1()` — L294
- `TXASetSipAllocDisps()` — L321
- `create_siphonEXT()` — L346
- `destroy_siphonEXT()` — L352
- `flush_siphonEXT()` — L358
- `xsiphonEXT()` — L364
- `GetaSipF1EXT()` — L372
- `SetSiphonInsize()` — L388

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/siphon.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
