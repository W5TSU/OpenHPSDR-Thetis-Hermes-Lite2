# `wdsp/eer.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Envelope elimination and restoration (polar) TX processing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/zeer.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×5)
- Uses (outgoing references to other files):
  - `wdsp/delay.c` (calls ×17)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `pSetEERSamplerate()` (×2), `pSetEERSize()` (×2), `create_eer()` (×1), `destroy_eer()` (×1), `xeer()` (×1), `pSetEERRun()` (×1), `pSetEERAMIQ()` (×1), `pSetEERMgain()` (×1)

## Outline

### Functions

- `create_eer()` — L29
- `destroy_eer()` — L69
- `flush_eer()` — L78
- `xeer()` — L85
- `create_eerEXT()` — L134
- `destroy_eerEXT()` — L140
- `flush_eerEXT()` — L146
- `SetEERRun()` — L152
- `SetEERAMIQ()` — L161
- `SetEERMgain()` — L170
- `SetEERPgain()` — L179
- `SetEERRunDelays()` — L188
- `SetEERMdelay()` — L199
- `SetEERPdelay()` — L209
- `SetEERSize()` — L219
- `SetEERSamplerate()` — L230
- `pSetEERRun()` — L263
- `pSetEERAMIQ()` — L271
- `pSetEERMgain()` — L279
- `pSetEERPgain()` — L287
- `pSetEERRunDelays()` — L295
- `pSetEERMdelay()` — L305
- `pSetEERPdelay()` — L314
- `pSetEERSize()` — L323
- `pSetEERSamplerate()` — L333
- `xeerEXTF()` — L366

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/eer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
