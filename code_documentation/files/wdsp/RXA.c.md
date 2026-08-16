# `wdsp/RXA.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Define the complete receive and transmit DSP graphs — every block below is instantiated and ordered here.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/main.c` (calls ×8)
  - `wdsp/snb.c` (calls ×4)
  - `wdsp/amd.c` (calls ×2)
  - `wdsp/anf.c` (calls ×2)
  - `wdsp/anr.c` (calls ×2)
  - `wdsp/emnr.c` (calls ×2)
  - `wdsp/rnnr.c` (calls ×2)
  - `wdsp/sbnr.c` (calls ×2)
  - `wdsp/nbp.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/snb.c` (calls ×19)
  - `wdsp/resample.c` (calls ×15)
  - `wdsp/iir.c` (calls ×14)
  - `wdsp/bandpass.c` (calls ×12)
  - `wdsp/nbp.c` (calls ×12)
  - `wdsp/fmd.c` (calls ×11)
  - `wdsp/shift.c` (calls ×11)
  - `wdsp/fmsq.c` (calls ×10)
  - `wdsp/eq.c` (calls ×9)
  - `wdsp/amsq.c` (calls ×8)
  - `wdsp/amd.c` (calls ×7)
  - `wdsp/anf.c` (calls ×7)
  - …and 21 more files
- Most-referenced symbols from other files: `RXAbp1Check()` (×7), `RXAbp1Set()` (×7), `RXAbpsnbaCheck()` (×2), `RXAbpsnbaSet()` (×2), `create_rxa()` (×1), `destroy_rxa()` (×1), `flush_rxa()` (×1), `xrxa()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_rxa()`** — L31 — `void create_rxa (int channel)`
  Constructor for the `rxa` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_main()` (`wdsp/main.c`)
- **`destroy_rxa()`** — L568 — `void destroy_rxa (int channel)`
  Destroys the `rxa` block, freeing its allocated buffers.
  Called by: `destroy_main()` (`wdsp/main.c`)
- **`flush_rxa()`** — L610 — `void flush_rxa (int channel)`
  Flushes (zeroes) the `rxa` block’s internal buffers/state.
  Called by: `flush_main()` (`wdsp/main.c`)
- **`xrxa()`** — L648 — `void xrxa (int channel)`
  Runs the `rxa` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `wdspmain()` (`wdsp/main.c`)
- **`setInputSamplerate_rxa()`** — L696 — `void setInputSamplerate_rxa (int channel)`
  Called by: `setInputSamplerate_main()` (`wdsp/main.c`)
- **`setOutputSamplerate_rxa()`** — L712 — `void setOutputSamplerate_rxa (int channel)`
  Called by: `setOutputSamplerate_main()` (`wdsp/main.c`)
- **`setDSPSamplerate_rxa()`** — L723 — `void setDSPSamplerate_rxa (int channel)`
  Called by: `setDSPSamplerate_main()` (`wdsp/main.c`)
- **`setDSPBuffsize_rxa()`** — L775 — `void setDSPBuffsize_rxa (int channel)`
  Called by: `setDSPBuffsize_main()` (`wdsp/main.c`)
- **`SetRXAMode()`** — L862 — `PORT void SetRXAMode (int channel, int mode)`
  Sets rxamode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXAResCheck()`** — L906 — `void RXAResCheck (int channel)`
  RXA chain operation — res check; part of the receive/transmit chain API.
  Called by: `create_rxa()` (same file), `setInputSamplerate_rxa()` (same file), `setOutputSamplerate_rxa()` (same file), `setDSPSamplerate_rxa()` (same file)
- **`RXAbp1Check()`** — L917 — `void RXAbp1Check (int channel, int amd_run, int snba_run, int emnr_run, int anf_run, int anr_run, int rnnr_run, int sbnr_run)`
  Called by: `SetRXAMode()` (same file), `SetRXAAMDRun()` (`wdsp/amd.c`), `SetRXAANFRun()` (`wdsp/anf.c`), `SetRXAANRRun()` (`wdsp/anr.c`), `SetRXAEMNRRun()` (`wdsp/emnr.c`), `SetRXARNNRRun()` (`wdsp/rnnr.c`) — and 2 more
- **`RXAbp1Set()`** — L935 — `void RXAbp1Set (int channel)`
  Called by: `SetRXAMode()` (same file), `SetRXAAMDRun()` (`wdsp/amd.c`), `SetRXAANFRun()` (`wdsp/anf.c`), `SetRXAANRRun()` (`wdsp/anr.c`), `SetRXAEMNRRun()` (`wdsp/emnr.c`), `SetRXARNNRRun()` (`wdsp/rnnr.c`) — and 2 more
- **`RXAbpsnbaCheck()`** — L951 — `void RXAbpsnbaCheck (int channel, int mode, int notch_run)`
  Called by: `SetRXAMode()` (same file), `RXANBPSetNotchesRun()` (`wdsp/nbp.c`), `SetRXASNBARun()` (`wdsp/snb.c`)
- **`RXAbpsnbaSet()`** — L1005 — `void RXAbpsnbaSet (int channel)`
  Called by: `SetRXAMode()` (same file), `RXANBPSetNotchesRun()` (`wdsp/nbp.c`), `SetRXASNBARun()` (`wdsp/snb.c`)
- **`RXASetPassband()`** — L1048 — `PORT void RXASetPassband (int channel, double f_low, double f_high)`
  RXA chain operation — set passband; part of the receive/transmit chain API.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`RXASetNC()`** — L1056 — `PORT void RXASetNC (int channel, int nc)`
  RXA chain operation — set nc; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`RXASetMP()`** — L1070 — `PORT void RXASetMP (int channel, int mp)`
  RXA chain operation — set mp; part of the receive/transmit chain API.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/RXA.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
