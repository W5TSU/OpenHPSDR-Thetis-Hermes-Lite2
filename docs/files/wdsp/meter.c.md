# `wdsp/meter.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Signal level metering taps feeding the console's meters.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/TXA.c` (calls ×12)
  - `wdsp/RXA.c` (calls ×7)
- Uses (outgoing references to other files):
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/meterlog10.c` (calls ×1)
- Most-referenced symbols from other files: `setBuffers_meter()` (×4), `setSize_meter()` (×4), `setSamplerate_meter()` (×3), `create_meter()` (×2), `destroy_meter()` (×2), `flush_meter()` (×2), `xmeter()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`calc_meter()`** — L29 — `void calc_meter (METER a)`
  Called by: `create_meter()` (same file), `setSamplerate_meter()` (same file)
- **`create_meter()`** — L36 — `METER create_meter (int run, int* prun, int size, double* buff, int rate, double tau_av, double tau_decay, double* result, CRITICAL_SECTION** pmtupdate, int enum_av, int enum_pk, i`
  Constructor for the `meter` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_rxa()` (`wdsp/RXA.c`), `create_txa()` (`wdsp/TXA.c`)
- **`destroy_meter()`** — L59 — `void destroy_meter (METER a)`
  Destroys the `meter` block, freeing its allocated buffers.
  Called by: `destroy_rxa()` (`wdsp/RXA.c`), `destroy_txa()` (`wdsp/TXA.c`)
- **`flush_meter()`** — L65 — `void flush_meter (METER a)`
  Flushes (zeroes) the `meter` block’s internal buffers/state.
  Called by: `calc_meter()` (same file), `setSize_meter()` (same file), `flush_rxa()` (`wdsp/RXA.c`), `flush_txa()` (`wdsp/TXA.c`)
- **`xmeter()`** — L75 — `void xmeter (METER a)`
  Runs the `meter` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xrxa()` (`wdsp/RXA.c`), `xtxa()` (`wdsp/TXA.c`)
- **`setBuffers_meter()`** — L110 — `void setBuffers_meter (METER a, double* in)`
  Re-points the `meter` block’s input/output buffers (called when the channel’s buffers change).
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setOutputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`setSamplerate_meter()`** — L115 — `void setSamplerate_meter (METER a, int rate)`
  Reconfigures the `meter` block for a new sample rate.
  Called by: `setDSPSamplerate_rxa()` (`wdsp/RXA.c`), `setOutputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`)
- **`setSize_meter()`** — L121 — `void setSize_meter (METER a, int size)`
  Reconfigures the `meter` block for a new buffer size.
  Called by: `setDSPBuffsize_rxa()` (`wdsp/RXA.c`), `setOutputSamplerate_txa()` (`wdsp/TXA.c`), `setDSPSamplerate_txa()` (`wdsp/TXA.c`), `setDSPBuffsize_txa()` (`wdsp/TXA.c`)
- **`GetRXAMeter()`** — L133 — `PORT double GetRXAMeter (int channel, int mt)`
  Returns rxameter — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetTXAMeter()`** — L150 — `PORT double GetTXAMeter (int channel, int mt)`
  Returns txameter — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/meter.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
