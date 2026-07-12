# `wdsp/div.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Diversity combiner (mixes two receivers with adjustable gain/phase).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/sync.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_divEXT()` (×1), `destroy_divEXT()` (×1), `xdivEXT()` (×1)

## Outline

### Functions

- `create_div()` — L31
- `destroy_div()` — L50
- `flush_div()` — L62
- `xdiv()` — L67
- `create_divEXT()` — L107
- `destroy_divEXT()` — L113
- `flush_divEXT()` — L119
- `xdivEXT()` — L125
- `SetEXTDIVRun()` — L137
- `SetEXTDIVBuffsize()` — L147
- `SetEXTDIVNr()` — L157
- `SetEXTDIVOutput()` — L168
- `SetEXTDIVRotate()` — L179
- `xdivEXTF()` — L194

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/div.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
