# `wdsp/delay.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/eer.c` (calls ×17)
  - `wdsp/calcc.c` (calls ×8)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_delay()` (×5), `destroy_delay()` (×5), `SetDelayValue()` (×5), `SetDelayBuffs()` (×4), `flush_delay()` (×2), `xdelay()` (×2), `SetDelayRun()` (×2)

## Outline

### Functions

- `create_delay()` — L29
- `destroy_delay()` — L57
- `flush_delay()` — L65
- `xdelay()` — L71
- `SetDelayRun()` — L107
- `SetDelayValue()` — L114
- `SetDelayBuffs()` — L128

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/delay.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
