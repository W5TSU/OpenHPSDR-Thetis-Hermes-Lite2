# `wdsp/TXA.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Define the complete receive and transmit DSP graphs — every block below is instantiated and ordered here.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/main.c` (calls ×8)
  - `wdsp/slew.c` (calls ×1)
  - `wdsp/compress.c` (calls ×1)
  - `wdsp/osctrl.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/resample.c` (calls ×15)
  - `wdsp/meter.c` (calls ×12)
  - `wdsp/bandpass.c` (calls ×10)
  - `wdsp/emph.c` (calls ×10)
  - `wdsp/fmmod.c` (calls ×10)
  - `wdsp/cfir.c` (calls ×9)
  - `wdsp/eq.c` (calls ×9)
  - `wdsp/amsq.c` (calls ×8)
  - `wdsp/ammod.c` (calls ×7)
  - `wdsp/cfcomp.c` (calls ×7)
  - `wdsp/compress.c` (calls ×7)
  - `wdsp/gen.c` (calls ×7)
  - …and 11 more files
- Most-referenced symbols from other files: `TXASetupBPFilters()` (×2), `create_txa()` (×1), `destroy_txa()` (×1), `flush_txa()` (×1), `xtxa()` (×1), `setInputSamplerate_txa()` (×1), `setOutputSamplerate_txa()` (×1), `setDSPSamplerate_txa()` (×1)

## Outline

### Functions

- `create_txa()` — L31
- `destroy_txa()` — L481
- `flush_txa()` — L520
- `xtxa()` — L557
- `setInputSamplerate_txa()` — L594
- `setOutputSamplerate_txa()` — L606
- `setDSPSamplerate_txa()` — L623
- `setDSPBuffsize_txa()` — L671
- `SetTXAMode()` — L752
- `SetTXABandpassFreqs()` — L791
- `TXAResCheck()` — L809
- `TXAUslewCheck()` — L819
- `TXASetupBPFilters()` — L827
- `TXASetNC()` — L909
- `TXASetMP()` — L921
- `SetTXAFMAFFilter()` — L930

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/TXA.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
