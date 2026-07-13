# `wdsp/apfshadow.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Matched CW filtering and audio peaking filter support.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/RXA.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/iir.c` (calls ×8)
  - `wdsp/doublepole.c` (calls ×7)
  - `wdsp/gaussian.c` (calls ×7)
  - `wdsp/matchedCW.c` (calls ×7)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_apfshadow()` (×1), `destroy_apfshadow()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_apfshadow()`** — L28 — `APFSHADOW create_apfshadow (int selection, int run, double f_center, double bandwidth, double gain)`
  Constructor for the `apfshadow` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`)
- **`destroy_apfshadow()`** — L39 — `void destroy_apfshadow (APFSHADOW a)`
  Destroys the `apfshadow` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`)
- **`SetRXASPCWSelection()`** — L44 — `PORT void SetRXASPCWSelection (int channel, int selection)`
  Sets rxaspcwselection — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASPCWRun()`** — L92 — `PORT void SetRXASPCWRun (int channel, int run)`
  Sets rxaspcwrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASPCWFreq()`** — L116 — `PORT void SetRXASPCWFreq (int channel, double f_center)`
  Sets rxaspcwfreq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASPCWBandwidth()`** — L140 — `PORT void SetRXASPCWBandwidth (int channel, double bandwidth)`
  Sets rxaspcwbandwidth — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRXASPCWGain()`** — L164 — `PORT void SetRXASPCWGain (int channel, double gain)`
  Sets rxaspcwgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/apfshadow.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
