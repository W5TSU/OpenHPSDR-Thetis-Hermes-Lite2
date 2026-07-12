# `wdsp/iobuffs.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/channel.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×1)
  - `wdsp/main.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `flush_slews()` (×2), `create_slews()` (×2), `destroy_slews()` (×2), `fexchange0()` (×1), `create_iobuffs()` (×1), `destroy_iobuffs()` (×1), `flush_iobuffs()` (×1), `dexchange()` (×1)

## Outline

### Functions

- `create_slews()` — L47
- `destroy_slews()` — L82
- `flush_slews()` — L88
- `upslew0()` — L98
- `upslew2()` — L162
- `downslew0()` — L226
- `downslew2()` — L302
- `create_iobuffs()` — L384
- `destroy_iobuffs()` — L425
- `flush_iobuffs()` — L443
- `fexchange0()` — L464
- `fexchange2()` — L518
- `dexchange()` — L583

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iobuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
