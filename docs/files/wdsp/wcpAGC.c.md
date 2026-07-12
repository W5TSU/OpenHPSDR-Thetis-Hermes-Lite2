# `wdsp/wcpAGC.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** The WDSP AGC (receive gain control and TX leveler).

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
  - `wdsp/fmd.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_wcpagc()` (×3), `destroy_wcpagc()` (×3), `flush_wcpagc()` (×3), `xwcpagc()` (×3), `setSamplerate_wcpagc()` (×3), `setBuffers_wcpagc()` (×3), `setSize_wcpagc()` (×3)

## Outline

### Functions

- `calc_wcpagc()` — L35
- `decalc_wcpagc()` — L54
- `create_wcpagc()` — L60
- `loadWcpAGC()` — L115
- `destroy_wcpagc()` — L148
- `flush_wcpagc()` — L154
- `xwcpagc()` — L161
- `setBuffers_wcpagc()` — L348
- `setSamplerate_wcpagc()` — L354
- `setSize_wcpagc()` — L361
- `SetRXAAGCMode()` — L374
- `SetRXAAGCAttack()` — L417
- `SetRXAAGCDecay()` — L426
- `SetRXAAGCHang()` — L435
- `GetRXAAGCHangLevel()` — L444
- `SetRXAAGCHangLevel()` — L453
- `GetRXAAGCHangThreshold()` — L472
- `SetRXAAGCHangThreshold()` — L481
- `GetRXAAGCThresh()` — L491
- `SetRXAAGCThresh()` — L503
- `GetRXAAGCTop()` — L517
- `SetRXAAGCTop()` — L526
- `SetRXAAGCSlope()` — L536
- `SetRXAAGCFixed()` — L545
- `SetRXAAGCMaxInputLevel()` — L554
- `SetTXAALCSt()` — L569
- `SetTXAALCAttack()` — L577
- `SetTXAALCDecay()` — L585
- `SetTXAALCHang()` — L594
- `SetTXAALCMaxGain()` — L603
- `SetTXALevelerSt()` — L612
- `SetTXALevelerAttack()` — L620
- `SetTXALevelerDecay()` — L629
- `SetTXALevelerHang()` — L638
- `SetTXALevelerTop()` — L647

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/wcpAGC.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
