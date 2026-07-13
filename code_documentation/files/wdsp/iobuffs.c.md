# `wdsp/iobuffs.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Sample buffering between the audio callback world and DSP blocks.

## How this file is used

- Used by (incoming references from other files):
  - `wdsp/channel.c` (calls ×9)
  - `ChannelMaster/cmaster.c` (calls ×1)
  - `wdsp/main.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×2)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `flush_slews()` (×2), `create_slews()` (×2), `destroy_slews()` (×2), `fexchange0()` (×1), `create_iobuffs()` (×1), `destroy_iobuffs()` (×1), `flush_iobuffs()` (×1), `dexchange()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_slews()`** — L47 — `void create_slews (IOB a)`
  Constructor for the `slews` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_iobuffs()` (same file), `SetChannelTSlewUp()` (`wdsp/channel.c`), `SetChannelTSlewDown()` (`wdsp/channel.c`)
- **`destroy_slews()`** — L82 — `void destroy_slews(IOB a)`
  Destroys the `slews` block, freeing its allocated buffers.
  Called by: `destroy_iobuffs()` (same file), `SetChannelTSlewUp()` (`wdsp/channel.c`), `SetChannelTSlewDown()` (`wdsp/channel.c`)
- **`flush_slews()`** — L88 — `void flush_slews (IOB a)`
  Flushes (zeroes) the `slews` block’s internal buffers/state.
  Called by: `flush_iobuffs()` (same file), `SetChannelTDelayUp()` (`wdsp/channel.c`), `SetChannelTDelayDown()` (`wdsp/channel.c`)
- **`upslew0()`** — L98 — `void upslew0 (IOB a, double* pin)`
  Called by: `fexchange0()` (same file)
- **`upslew2()`** — L162 — `void upslew2 (IOB a, INREAL* pIin, INREAL* pQin)`
  Called by: `fexchange2()` (same file)
- **`downslew0()`** — L226 — `void downslew0 (IOB a, double* pout)`
  Called by: `fexchange0()` (same file)
- **`downslew2()`** — L302 — `void downslew2 (IOB a, OUTREAL* pIout, OUTREAL* pQout)`
  Called by: `fexchange2()` (same file)
- **`create_iobuffs()`** — L384 — `void create_iobuffs (int channel)`
  Constructor for the `iobuffs` block: allocates its state/buffers and computes initial coefficients.
  Called by: `pre_main_build()` (`wdsp/channel.c`)
- **`destroy_iobuffs()`** — L425 — `void destroy_iobuffs (int channel)`
  Destroys the `iobuffs` block, freeing its allocated buffers.
  Called by: `post_main_destroy()` (`wdsp/channel.c`)
- **`flush_iobuffs()`** — L443 — `void flush_iobuffs (int channel)`
  Flushes (zeroes) the `iobuffs` block’s internal buffers/state.
  Called by: `flushChannel()` (`wdsp/channel.c`)
- **`fexchange0()`** — L464 — `PORT void fexchange0 (int channel, double* in, double* out, int* error)`
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`fexchange2()`** — L518 — `PORT void fexchange2 (int channel, INREAL *Iin, INREAL *Qin, OUTREAL *Iout, OUTREAL *Qout, int* error)`
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`dexchange()`** — L583 — `void dexchange (int channel, double* in, double* out)`
  Called by: `wdspmain()` (`wdsp/main.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/iobuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
