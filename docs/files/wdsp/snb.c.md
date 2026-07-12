# `wdsp/snb.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Spectral noise blanker.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×19)
  - `wdsp/nbp.c` (calls ×3)
- Uses (outgoing references to other files):
  - `wdsp/nbp.c` (calls ×7)
  - `wdsp/resample.c` (calls ×6)
  - `wdsp/RXA.c` (calls ×4)
  - `wdsp/utilities.c` (calls ×4)
  - `wdsp/lmath.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/firmin.c` (calls ×1)
- Most-referenced symbols from other files: `recalc_bpsnba_filter()` (×4), `create_bpsnba()` (×1), `create_snba()` (×1), `destroy_bpsnba()` (×1), `destroy_snba()` (×1), `flush_bpsnba()` (×1), `flush_snba()` (×1), `xbpsnbain()` (×1)

## Outline

### Functions

- `calc_snba()` — L31
- `create_snba()` — L68
- `decalc_snba()` — L121
- `destroy_snba()` — L131
- `flush_snba()` — L161
- `setBuffers_snba()` — L187
- `setSamplerate_snba()` — L195
- `setSize_snba()` — L202
- `ATAc0()` — L209
- `multA1TA2()` — L218
- `multXKE()` — L241
- `multAv()` — L254
- `xHat()` — L265
- `invf()` — L306
- `det()` — L324
- `scanFrame()` — L404
- `execFrame()` — L492
- `xsnba()` — L539
- `SetRXASNBARun()` — L579
- `SetRXASNBAovrlp()` — L598
- `SetRXASNBAasize()` — L607
- `SetRXASNBAnpasses()` — L614
- `SetRXASNBAk1()` — L621
- `SetRXASNBAk2()` — L628
- `SetRXASNBAbridge()` — L635
- `SetRXASNBApresamps()` — L642
- `SetRXASNBApostsamps()` — L649
- `SetRXASNBApmultmin()` — L656
- `SetRXASNBAOutputBandwidth()` — L663
- `calc_bpsnba()` — L714
- `create_bpsnba()` — L736
- `decalc_bpsnba()` — L763
- `destroy_bpsnba()` — L769
- `flush_bpsnba()` — L775
- `setBuffers_bpsnba()` — L781
- `setSamplerate_bpsnba()` — L789
- `setSize_bpsnba()` — L796
- `xbpsnbain()` — L803
- `xbpsnbaout()` — L809
- `recalc_bpsnba_filter()` — L815
- `RXABPSNBASetNC()` — L837
- `RXABPSNBASetMP()` — L852

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/snb.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
