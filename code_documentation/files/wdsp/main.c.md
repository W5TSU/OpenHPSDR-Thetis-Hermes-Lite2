# `wdsp/main.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Channel object lifecycle (create/destroy/run) and DLL entry points.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/channel.c` (calls ×10)
- Uses (outgoing references to other files):
  - `wdsp/RXA.c` (calls ×8)
  - `wdsp/TXA.c` (calls ×8)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/iobuffs.c` (calls ×1)
- Most-referenced symbols from other files: `setInputSamplerate_main()` (×2), `setDSPSamplerate_main()` (×2), `setOutputSamplerate_main()` (×2), `create_main()` (×1), `destroy_main()` (×1), `flush_main()` (×1), `setDSPBuffsize_main()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`wdspmain()`** — L29 — `void wdspmain (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_main()`** — L63 — `void create_main (int channel)`
  Constructor for the `main` block: allocates its state/buffers and computes initial coefficients.
  Called by: `build_channel()` (`wdsp/channel.c`)
- **`destroy_main()`** — L79 — `void destroy_main (int channel)`
  Destroys the `main` block, freeing its allocated buffers.
  Called by: `CloseChannel()` (`wdsp/channel.c`)
- **`flush_main()`** — L95 — `void flush_main (int channel)`
  Flushes (zeroes) the `main` block’s internal buffers/state.
  Called by: `flushChannel()` (`wdsp/channel.c`)
- **`setInputSamplerate_main()`** — L111 — `void setInputSamplerate_main (int channel)`
  Called by: `SetInputSamplerate()` (`wdsp/channel.c`), `SetAllRates()` (`wdsp/channel.c`)
- **`setOutputSamplerate_main()`** — L127 — `void setOutputSamplerate_main (int channel)`
  Called by: `SetOutputSamplerate()` (`wdsp/channel.c`), `SetAllRates()` (`wdsp/channel.c`)
- **`setDSPSamplerate_main()`** — L143 — `void setDSPSamplerate_main (int channel)`
  Called by: `SetDSPSamplerate()` (`wdsp/channel.c`), `SetAllRates()` (`wdsp/channel.c`)
- **`setDSPBuffsize_main()`** — L159 — `void setDSPBuffsize_main (int channel)`
  Called by: `SetDSPBuffsize()` (`wdsp/channel.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/main.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
