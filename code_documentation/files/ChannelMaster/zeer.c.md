# `ChannelMaster/zeer.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `wdsp/eer.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/ilv.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`SetEERRun()`** — L29 — `PORT void SetEERRun (int id, int run)`
  Sets eerrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERAMIQ()`** — L38 — `PORT void SetEERAMIQ (int id, int amiq)`
  Sets eeramiq — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERMgain()`** — L45 — `PORT void SetEERMgain (int id, double gain)`
  Sets eermgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERPgain()`** — L52 — `PORT void SetEERPgain (int id, double gain)`
  Sets eerpgain — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERRunDelays()`** — L59 — `PORT void SetEERRunDelays (int id, int run)`
  Sets eerrun delays — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERMdelay()`** — L66 — `PORT void SetEERMdelay (int id, double delay)`
  Sets eermdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERPdelay()`** — L73 — `PORT void SetEERPdelay (int id, double delay)`
  Sets eerpdelay — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`, `Console/dsp.cs`.
- **`SetEERSize()`** — L80 — `PORT void SetEERSize (int id, int size)`
  Sets eersize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetEERSamplerate()`** — L87 — `PORT void SetEERSamplerate (int id, int rate)`
  Sets eersamplerate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/zeer.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
