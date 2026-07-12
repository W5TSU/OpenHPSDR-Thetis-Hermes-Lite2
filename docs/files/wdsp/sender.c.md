# `wdsp/sender.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sends DSP data (spectrum, audio taps) back toward the console.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/analyzer.c` (calls ×1)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_sender()` (×1), `destroy_sender()` (×1), `flush_sender()` (×1), `xsender()` (×1), `setSamplerate_sender()` (×1), `setBuffers_sender()` (×1), `setSize_sender()` (×1)

## Outline

### Functions

- `calc_sender()` — L29
- `decalc_sender()` — L34
- `create_sender()` — L39
- `destroy_sender()` — L55
- `flush_sender()` — L61
- `xsender()` — L66
- `setBuffers_sender()` — L88
- `setSamplerate_sender()` — L93
- `setSize_sender()` — L98
- `SetRXASpectrum()` — L111

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/sender.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
