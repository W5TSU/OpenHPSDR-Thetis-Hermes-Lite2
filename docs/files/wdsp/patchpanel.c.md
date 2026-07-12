# `wdsp/patchpanel.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_panel()` (×2), `destroy_panel()` (×2), `flush_panel()` (×2), `xpanel()` (×2), `setSamplerate_panel()` (×2), `setBuffers_panel()` (×2), `setSize_panel()` (×2)

## Outline

### Functions

- `create_panel()` — L29
- `destroy_panel()` — L45
- `flush_panel()` — L50
- `xpanel()` — L55
- `setBuffers_panel()` — L103
- `setSamplerate_panel()` — L109
- `setSize_panel()` — L114
- `SetRXAPanelRun()` — L125
- `SetRXAPanelSelect()` — L133
- `SetRXAPanelGain1()` — L141
- `SetRXAPanelGain2()` — L149
- `SetRXAPanelPan()` — L158
- `SetRXAPanelCopy()` — L178
- `SetRXAPanelBinaural()` — L186
- `SetTXAPanelRun()` — L200
- `SetTXAPanelGain1()` — L208
- `SetTXAPanelSelect()` — L217

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/patchpanel.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
