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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_txa()`** — L31 — `void create_txa (int channel)`
  Constructor for the `txa` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_main()` (`wdsp/main.c`)
- **`destroy_txa()`** — L481 — `void destroy_txa (int channel)`
  Destroys the `txa` block, freeing its allocated buffers.
  Called by: `destroy_main()` (`wdsp/main.c`)
- **`flush_txa()`** — L520 — `void flush_txa (int channel)`
  Flushes (zeroes) the `txa` block’s internal buffers/state.
  Called by: `flush_main()` (`wdsp/main.c`)
- **`xtxa()`** — L557 — `void xtxa (int channel)`
  Runs the `txa` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `wdspmain()` (`wdsp/main.c`)
- **`setInputSamplerate_txa()`** — L594 — `void setInputSamplerate_txa (int channel)`
  Called by: `setInputSamplerate_main()` (`wdsp/main.c`)
- **`setOutputSamplerate_txa()`** — L606 — `void setOutputSamplerate_txa (int channel)`
  Called by: `setOutputSamplerate_main()` (`wdsp/main.c`)
- **`setDSPSamplerate_txa()`** — L623 — `void setDSPSamplerate_txa (int channel)`
  Called by: `setDSPSamplerate_main()` (`wdsp/main.c`)
- **`setDSPBuffsize_txa()`** — L671 — `void setDSPBuffsize_txa (int channel)`
  Called by: `setDSPBuffsize_main()` (`wdsp/main.c`)
- **`SetTXAMode()`** — L752 — `PORT void SetTXAMode (int channel, int mode)`
  Sets txamode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXABandpassFreqs()`** — L791 — `PORT void SetTXABandpassFreqs (int channel, double f_low, double f_high)`
  Sets txabandpass freqs — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXAResCheck()`** — L809 — `void TXAResCheck (int channel)`
  TXA chain operation — res check; part of the receive/transmit chain API.
  Called by: `create_txa()` (same file), `setInputSamplerate_txa()` (same file), `setOutputSamplerate_txa()` (same file), `setDSPSamplerate_txa()` (same file)
- **`TXAUslewCheck()`** — L819 — `int TXAUslewCheck (int channel)`
  TXA chain operation — uslew check; part of the receive/transmit chain API.
  Called by: `xuslew()` (`wdsp/slew.c`)
- **`TXASetupBPFilters()`** — L827 — `void TXASetupBPFilters (int channel)`
  TXA chain operation — setup bpfilters; part of the receive/transmit chain API.
  Called by: `SetTXAMode()` (same file), `SetTXABandpassFreqs()` (same file), `SetTXACompressorRun()` (`wdsp/compress.c`), `SetTXAosctrlRun()` (`wdsp/osctrl.c`)
- **`TXASetNC()`** — L909 — `PORT void TXASetNC (int channel, int nc)`
  TXA chain operation — set nc; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`TXASetMP()`** — L921 — `PORT void TXASetMP (int channel, int mp)`
  TXA chain operation — set mp; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetTXAFMAFFilter()`** — L930 — `PORT void SetTXAFMAFFilter (int channel, double low, double high)`
  Sets txafmaffilter — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/TXA.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
