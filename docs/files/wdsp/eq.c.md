# `wdsp/eq.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Graphic/parametric equalizer.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×9)
  - `wdsp/TXA.c` (calls ×9)
  - `wdsp/fmsq.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×22)
  - `wdsp/utilities.c` (calls ×10)
  - `wdsp/fir.c` (calls ×3)
  - `wdsp/impulse_cache.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_eqp()` (×2), `destroy_eqp()` (×2), `flush_eqp()` (×2), `xeqp()` (×2), `setSamplerate_eqp()` (×2), `setBuffers_eqp()` (×2), `setSize_eqp()` (×2), `eq_impulse()` (×2)

## Outline

### Functions

- `fEQcompare()` — L54
- `fEQcompare3()` — L64
- `eq_impulse()` — L74
- `create_eqp()` — L508
- `destroy_eqp()` — L534
- `flush_eqp()` — L540
- `xeqp()` — L545
- `setBuffers_eqp()` — L553
- `setSamplerate_eqp()` — L560
- `setSize_eqp()` — L569
- `SetRXAEQRun()` — L585
- `SetRXAEQNC()` — L593
- `SetRXAEQMP()` — L610
- `SetRXAEQProfile()` — L622
- `SetRXAEQCtfmode()` — L651
- `SetRXAEQWintype()` — L663
- `SetRXAGrphEQ()` — L675
- `SetRXAGrphEQ10()` — L703
- `SetTXAEQRun()` — L742
- `SetTXAEQNC()` — L750
- `SetTXAEQMP()` — L767
- `SetTXAEQProfile()` — L779
- `SetTXAEQCtfmode()` — L807
- `SetTXAEQWintype()` — L819
- `SetTXAGrphEQ()` — L831
- `SetTXAGrphEQ10()` — L859
- `eq_mults()` — L898
- `calc_eq()` — L906
- `decalc_eq()` — L916
- `create_eq()` — L925
- `destroy_eq()` — L944
- `flush_eq()` — L953
- `xeq()` — L958
- `setBuffers_eq()` — L980
- `setSamplerate_eq()` — L988
- `setSize_eq()` — L995

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/eq.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
