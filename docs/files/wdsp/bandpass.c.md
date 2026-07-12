# `wdsp/bandpass.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Main bandpass filter and the notched-bandpass (auto/manual notch database) filter.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×12)
  - `wdsp/TXA.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/firmin.c` (calls ×19)
  - `wdsp/fir.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `flush_bandpass()` (×3), `create_bandpass()` (×2), `destroy_bandpass()` (×2), `xbandpass()` (×2), `setSamplerate_bandpass()` (×2), `setBuffers_bandpass()` (×2), `setSize_bandpass()` (×2), `setGain_bandpass()` (×1)

## Outline

### Functions

- `calc_bps()` — L35
- `decalc_bps()` — L47
- `create_bps()` — L56
- `destroy_bps()` — L74
- `flush_bps()` — L80
- `xbps()` — L85
- `setBuffers_bps()` — L107
- `setSamplerate_bps()` — L115
- `setSize_bps()` — L122
- `setFreqs_bps()` — L129
- `create_bandpass()` — L284
- `destroy_bandpass()` — L308
- `flush_bandpass()` — L314
- `xbandpass()` — L319
- `setBuffers_bandpass()` — L327
- `setSamplerate_bandpass()` — L334
- `setSize_bandpass()` — L343
- `setGain_bandpass()` — L355
- `CalcBandpassFilter()` — L364
- `SetRXABandpassRun()` — L384
- `SetRXABandpassFreqs()` — L392
- `SetRXABandpassWindow()` — L411
- `SetRXABandpassNC()` — L429
- `SetRXABandpassMP()` — L447
- `SetTXABandpassRun()` — L465
- `SetTXABandpassWindow()` — L507
- `SetTXABandpassNC()` — L538
- `SetTXABandpassMP()` — L572

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/bandpass.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
