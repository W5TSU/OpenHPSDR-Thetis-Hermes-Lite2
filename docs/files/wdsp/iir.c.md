# `wdsp/iir.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** IIR biquad sections (notches, peaking filters) and double-pole building blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×14)
  - `wdsp/apfshadow.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×7)
  - `wdsp/fmd.c` (calls ×6)
  - `wdsp/ssql.c` (calls ×4)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×14)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRXABiQuadBandwidth()` (×2), `SetRXABiQuadFreq()` (×2), `SetRXABiQuadGain()` (×2), `SetRXABiQuadRun()` (×2), `create_mpeak()` (×1), `create_speak()` (×1), `destroy_mpeak()` (×1), `destroy_speak()` (×1)

## Outline

### Functions

- `calc_snotch()` — L35
- `create_snotch()` — L50
- `destroy_snotch()` — L65
- `flush_snotch()` — L71
- `xsnotch()` — L76
- `setBuffers_snotch()` — L97
- `setSamplerate_snotch()` — L103
- `setSize_snotch()` — L109
- `SetSNCTCSSFreq()` — L121
- `SetSNCTCSSRun()` — L129
- `calc_speak()` — L143
- `create_speak()` — L218
- `destroy_speak()` — L242
- `flush_speak()` — L254
- `xspeak()` — L264
- `setBuffers_speak()` — L297
- `setSamplerate_speak()` — L303
- `setSize_speak()` — L309
- `SetRXABiQuadRun()` — L321
- `SetRXABiQuadFreq()` — L330
- `SetRXABiQuadBandwidth()` — L340
- `SetRXABiQuadGain()` — L350
- `calc_mpeak()` — L366
- `decalc_mpeak()` — L386
- `create_mpeak()` — L395
- `destroy_mpeak()` — L419
- `flush_mpeak()` — L431
- `xmpeak()` — L438
- `setBuffers_mpeak()` — L461
- `setSamplerate_mpeak()` — L469
- `setSize_mpeak()` — L476
- `SetRXAmpeakRun()` — L489
- `SetRXAmpeakNpeaks()` — L498
- `SetRXAmpeakFilEnable()` — L507
- `SetRXAmpeakFilFreq()` — L516
- `SetRXAmpeakFilBw()` — L527
- `SetRXAmpeakFilGain()` — L538
- `calc_phrot()` — L556
- `create_phrot()` — L569
- `decalc_phrot()` — L585
- `destroy_phrot()` — L593
- `flush_phrot()` — L600
- `xphrot()` — L608
- `setBuffers_phrot()` — L639
- `setSamplerate_phrot()` — L645
- `setSize_phrot()` — L652
- `SetTXAPHROTRun()` — L664
- `SetTXAPHROTCorner()` — L674
- `SetTXAPHROTNstages()` — L685
- `SetTXAPHROTReverse()` — L696
- `calc_bqlp()` — L711
- `create_bqlp()` — L726
- `destroy_bqlp()` — L749
- `flush_bqlp()` — L761
- `xbqlp()` — L771
- `setBuffers_bqlp()` — L804
- `setSamplerate_bqlp()` — L810
- `setSize_bqlp()` — L816
- `calc_dbqlp()` — L828
- `create_dbqlp()` — L843
- `destroy_dbqlp()` — L866
- `flush_dbqlp()` — L878
- `xdbqlp()` — L887
- `setBuffers_dbqlp()` — L917
- `setSamplerate_dbqlp()` — L923
- `setSize_dbqlp()` — L929
- `calc_bqbp()` — L942
- `create_bqbp()` — L961
- `destroy_bqbp()` — L984
- `flush_bqbp()` — L996
- `xbqbp()` — L1006
- `setBuffers_bqbp()` — L1039
- `setSamplerate_bqbp()` — L1045
- `setSize_bqbp()` — L1051
- `calc_dbqbp()` — L1063
- `create_dbqbp()` — L1082
- `destroy_dbqbp()` — L1105
- `flush_dbqbp()` — L1117
- `xdbqbp()` — L1126
- `setBuffers_dbqbp()` — L1156
- `setSamplerate_dbqbp()` — L1162
- `setSize_dbqbp()` — L1168
- `calc_sphp()` — L1180
- `create_sphp()` — L1193
- `decalc_sphp()` — L1208
- `destroy_sphp()` — L1216
- `flush_sphp()` — L1223
- `xsphp()` — L1231
- `setBuffers_sphp()` — L1260
- `setSamplerate_sphp()` — L1266
- `setSize_sphp()` — L1273
- `calc_dsphp()` — L1285
- `create_dsphp()` — L1298
- `decalc_dsphp()` — L1313
- `destroy_dsphp()` — L1321
- `flush_dsphp()` — L1328
- `xdsphp()` — L1336
- `setBuffers_dsphp()` — L1362
- `setSamplerate_dsphp()` — L1368
- `setSize_dsphp()` — L1375

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iir.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
