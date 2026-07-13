# `wdsp/eer.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Envelope elimination and restoration (polar) TX processing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/zeer.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×5)
- Uses (outgoing references to other files):
  - `wdsp/delay.c` (calls ×17)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `pSetEERSamplerate()` (×2), `pSetEERSize()` (×2), `create_eer()` (×1), `destroy_eer()` (×1), `xeer()` (×1), `pSetEERRun()` (×1), `pSetEERAMIQ()` (×1), `pSetEERMgain()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_eer()`** — L29 — `PORT EER create_eer (int run, int size, double* in, double* out, double* outM, int rate, double mgain, double pgain, int rundelays, double mdelay, double pdelay, int amiq)`
  Constructor for the `eer` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_eerEXT()` (same file), `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`destroy_eer()`** — L69 — `PORT void destroy_eer (EER a)`
  Destroys the `eer` block, freeing its allocated buffers.
  Called by: `destroy_eerEXT()` (same file), `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`flush_eer()`** — L78 — `PORT void flush_eer (EER a)`
  Flushes (zeroes) the `eer` block’s internal buffers/state.
  Called by: `flush_eerEXT()` (same file)
- **`xeer()`** — L85 — `PORT void xeer (EER a)`
  Runs the `eer` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xeerEXTF()` (same file), `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`create_eerEXT()`** — L134 — `PORT void create_eerEXT (int id, int run, int size, int rate, double mgain, double pgain, int rundelays, double mdelay, double pdelay, int amiq)`
  Constructor for the `eerEXT` block: allocates its state/buffers and computes initial coefficients.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`destroy_eerEXT()`** — L140 — `PORT void destroy_eerEXT (int id)`
  Destroys the `eerEXT` block, freeing its allocated buffers.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`flush_eerEXT()`** — L146 — `PORT void flush_eerEXT (int id)`
  Flushes (zeroes) the `eerEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetEERRun()`** — L152 — `PORT void SetEERRun (int id, int run)`
  Sets eerrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERAMIQ()`** — L161 — `PORT void SetEERAMIQ (int id, int amiq)`
  Sets eeramiq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERMgain()`** — L170 — `PORT void SetEERMgain (int id, double gain)`
  Sets eermgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERPgain()`** — L179 — `PORT void SetEERPgain (int id, double gain)`
  Sets eerpgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERRunDelays()`** — L188 — `PORT void SetEERRunDelays (int id, int run)`
  Sets eerrun delays — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERMdelay()`** — L199 — `PORT void SetEERMdelay (int id, double delay)`
  Sets eermdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERPdelay()`** — L209 — `PORT void SetEERPdelay (int id, double delay)`
  Sets eerpdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERSize()`** — L219 — `PORT void SetEERSize (int id, int size)`
  Sets eersize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetEERSamplerate()`** — L230 — `PORT void SetEERSamplerate (int id, int rate)`
  Sets eersamplerate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`pSetEERRun()`** — L263 — `PORT void pSetEERRun (EER a, int run)`
  Called by: `SetEERRun()` (`ChannelMaster/zeer.c`)
- **`pSetEERAMIQ()`** — L271 — `PORT void pSetEERAMIQ (EER a, int amiq)`
  Called by: `SetEERAMIQ()` (`ChannelMaster/zeer.c`)
- **`pSetEERMgain()`** — L279 — `PORT void pSetEERMgain (EER a, double gain)`
  Called by: `SetEERMgain()` (`ChannelMaster/zeer.c`)
- **`pSetEERPgain()`** — L287 — `PORT void pSetEERPgain (EER a, double gain)`
  Called by: `SetEERPgain()` (`ChannelMaster/zeer.c`)
- **`pSetEERRunDelays()`** — L295 — `PORT void pSetEERRunDelays (EER a, int run)`
  Called by: `SetEERRunDelays()` (`ChannelMaster/zeer.c`)
- **`pSetEERMdelay()`** — L305 — `PORT void pSetEERMdelay (EER a, double delay)`
  Called by: `SetEERMdelay()` (`ChannelMaster/zeer.c`)
- **`pSetEERPdelay()`** — L314 — `PORT void pSetEERPdelay (EER a, double delay)`
  Called by: `SetEERPdelay()` (`ChannelMaster/zeer.c`)
- **`pSetEERSize()`** — L323 — `PORT void pSetEERSize (EER a, int size)`
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetEERSize()` (`ChannelMaster/zeer.c`)
- **`pSetEERSamplerate()`** — L333 — `PORT void pSetEERSamplerate (EER a, int rate)`
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetEERSamplerate()` (`ChannelMaster/zeer.c`)
- **`xeerEXTF()`** — L366 — `PORT void xeerEXTF (int id, float* inI, float* inQ, float* outI, float* outQ, float* outMI, float* outMQ, int mox, int size)`
  Runs the `eerEXTF` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/eer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
