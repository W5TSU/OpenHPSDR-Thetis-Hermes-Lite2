# `wdsp/meter.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal level metering taps feeding the console's meters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×12)
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `setBuffers_meter()` (×4), `setSize_meter()` (×4), `setSamplerate_meter()` (×3), `create_meter()` (×2), `destroy_meter()` (×2), `flush_meter()` (×2), `xmeter()` (×2)

## Outline

### Functions

- `calc_meter()` — L29
- `create_meter()` — L36
- `destroy_meter()` — L59
- `flush_meter()` — L65
- `xmeter()` — L75
- `setBuffers_meter()` — L110
- `setSamplerate_meter()` — L115
- `setSize_meter()` — L121
- `GetRXAMeter()` — L133
- `GetTXAMeter()` — L150

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/meter.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
