# `wdsp/gen.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal generators (tone, two-tone, noise, sweep) for testing and tune.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×7)
  - `wdsp/TXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `create_gen()` (×2), `destroy_gen()` (×2), `flush_gen()` (×2), `xgen()` (×2), `setSamplerate_gen()` (×2), `setBuffers_gen()` (×2), `setSize_gen()` (×2)

## Outline

### Functions

- `calc_tone()` — L29
- `calc_tt()` — L37
- `calc_sweep()` — L49
- `calc_sawtooth()` — L57
- `calc_triangle()` — L64
- `calc_pulse()` — L73
- `calc_ttpulse()` — L98
- `calc_gen()` — L127
- `decalc_gen()` — L138
- `create_gen()` — L144
- `destroy_gen()` — L195
- `flush_gen()` — L201
- `xgen()` — L215
- `setBuffers_gen()` — L530
- `setSamplerate_gen()` — L536
- `setSize_gen()` — L543
- `SetRXAPreGenRun()` — L558
- `SetRXAPreGenMode()` — L566
- `SetRXAPreGenToneMag()` — L574
- `SetRXAPreGenToneFreq()` — L582
- `SetRXAPreGenNoiseMag()` — L591
- `SetRXAPreGenSweepMag()` — L599
- `SetRXAPreGenSweepFreq()` — L607
- `SetRXAPreGenSweepRate()` — L617
- `SetTXAPreGenRun()` — L635
- `SetTXAPreGenMode()` — L643
- `SetTXAPreGenToneMag()` — L651
- `SetTXAPreGenToneFreq()` — L659
- `SetTXAPreGenNoiseMag()` — L668
- `SetTXAPreGenSweepMag()` — L676
- `SetTXAPreGenSweepFreq()` — L684
- `SetTXAPreGenSweepRate()` — L694
- `SetTXAPreGenSawtoothMag()` — L703
- `SetTXAPreGenSawtoothFreq()` — L711
- `SetTXAPreGenTriangleMag()` — L720
- `SetTXAPreGenTriangleFreq()` — L728
- `SetTXAPreGenPulseMag()` — L737
- `SetTXAPreGenPulseFreq()` — L745
- `SetTXAPreGenPulseDutyCycle()` — L754
- `SetTXAPreGenPulseToneFreq()` — L763
- `SetTXAPreGenPulseTransition()` — L772
- `SetTXAPostGenRun()` — L783
- `SetTXAPostGenMode()` — L791
- `SetTXAPostGenToneMag()` — L799
- `SetTXAPostGenToneFreq()` — L807
- `SetTXAPostGenTTMag()` — L816
- `SetTXAPostGenTTFreq()` — L825
- `SetTXAPostGenSweepMag()` — L835
- `SetTXAPostGenSweepFreq()` — L843
- `SetTXAPostGenSweepRate()` — L853
- `SetTXAPostGenPulseMag()` — L862
- `SetTXAPostGenPulseFreq()` — L870
- `SetTXAPostGenPulseDutyCycle()` — L879
- `SetTXAPostGenPulseToneFreq()` — L888
- `SetTXAPostGenPulseTransition()` — L897
- `SetTXAPostGenPulseIQout()` — L906
- `SetTXAPostGenTTPulseMag()` — L914
- `SetTXAPostGenTTPulseFreq()` — L925
- `SetTXAPostGenTTPulseDutyCycle()` — L934
- `SetTXAPostGenTTPulseToneFreq()` — L943
- `SetTXAPostGenTTPulseTransition()` — L954
- `SetTXAPostGenTTPulseIQout()` — L963

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/gen.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
