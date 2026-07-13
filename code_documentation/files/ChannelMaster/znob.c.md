# `ChannelMaster/znob.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/nob.c` (calls ×8)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRCVRANBBuffsize()` (×1), `SetRCVRANBSamplerate()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`GetRCVRANBPointer()`** — L29 — `ANB GetRCVRANBPointer (int stype, int id)`
  Returns rcvranbpointer — API getter, typically called from the console via P/Invoke.
  Called by: `SetRCVRANBRun()` (same file), `SetRCVRANBBuffsize()` (same file), `SetRCVRANBSamplerate()` (same file), `SetRCVRANBTau()` (same file), `SetRCVRANBHangtime()` (same file), `SetRCVRANBAdvtime()` (same file) — and 2 more
- **`SetRCVRANBRun()`** — L44 — `PORT void SetRCVRANBRun (int stype, int id, int run)`
  Sets rcvranbrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRANBBuffsize()`** — L50 — `PORT void SetRCVRANBBuffsize (int stype, int id, int size)`
  Sets rcvranbbuffsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetRCVRANBSamplerate()`** — L56 — `PORT void SetRCVRANBSamplerate (int stype, int id, int rate)`
  Sets rcvranbsamplerate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetRCVRANBTau()`** — L62 — `PORT void SetRCVRANBTau (int stype, int id, double tau)`
  Sets rcvranbtau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRANBHangtime()`** — L68 — `PORT void SetRCVRANBHangtime (int stype, int id, double time)`
  Sets rcvranbhangtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRANBAdvtime()`** — L74 — `PORT void SetRCVRANBAdvtime (int stype, int id, double time)`
  Sets rcvranbadvtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRANBBacktau()`** — L80 — `PORT void SetRCVRANBBacktau (int stype, int id, double tau)`
  Sets rcvranbbacktau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRANBThreshold()`** — L86 — `PORT void SetRCVRANBThreshold (int stype, int id, double thresh)`
  Sets rcvranbthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/znob.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
