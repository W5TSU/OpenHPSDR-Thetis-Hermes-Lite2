# `wdsp/utilities.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/iir.c` (calls ×14)
  - `wdsp/eq.c` (calls ×10)
  - `wdsp/fir.c` (calls ×9)
  - `wdsp/firmin.c` (calls ×7)
  - `ChannelMaster/aamix.c` (calls ×5)
  - `wdsp/RXA.c` (calls ×5)
  - `wdsp/TXA.c` (calls ×5)
  - `wdsp/dexp.c` (calls ×5)
  - `wdsp/rmatch.c` (calls ×5)
  - `wdsp/cfcomp.c` (calls ×4)
  - `wdsp/nbp.c` (calls ×4)
  - `wdsp/rnnr.c` (calls ×4)
  - …and 62 more files
- Uses (outgoing references to other files):
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `malloc0()` (×180)

## Outline

### Functions

- `malloc0()` — L36
- `NewCriticalSection()` — L47
- `DestroyCriticalSection()` — L55
- `print_impulse()` — L68
- `analyze_bandpass_filter()` — L90
- `print_peak_val()` — L103
- `print_peak_env()` — L123
- `print_peak_env_f2()` — L147
- `print_iqc_values()` — L166
- `print_buffer_parameters()` — L186
- `print_meter()` — L217
- `print_message()` — L231
- `print_window_gain()` — L243
- `print_deviation()` — L281
- `CalccPrintSamples()` — L293
- `doCalccPrintSamples()` — L317
- `print_anb_parms()` — L322
- `WriteAudioFile()` — L346
- `WriteAudioWDSP()` — L372
- `WriteScaledAudioFile()` — L425
- `WriteScaledAudio()` — L461
- `model_bandpass()` — L506
- `print_bandpass_response()` — L536
- `create_bfcu()` — L551
- `destroy_bfcu()` — L588
- `getFilterCorners()` — L603
- `getFilterCurve()` — L612
- `test_bfcu()` — L625

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/utilities.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
