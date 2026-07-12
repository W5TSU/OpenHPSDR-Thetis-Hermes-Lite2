# `ChannelMaster/znobII.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `wdsp/nobII.c` (calls ×9)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetRCVRNOBBuffsize()` (×1), `SetRCVRNOBSamplerate()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`GetRCVRNOBPointer()`** — L29 — `NOB GetRCVRNOBPointer (int stype, int id)`
  Returns rcvrnobpointer — API getter, typically called from the console via P/Invoke.
  Called by: `SetRCVRNOBRun()` (same file), `SetRCVRNOBMode()` (same file), `SetRCVRNOBBuffsize()` (same file), `SetRCVRNOBSamplerate()` (same file), `SetRCVRNOBTau()` (same file), `SetRCVRNOBHangtime()` (same file) — and 3 more
- **`SetRCVRNOBRun()`** — L44 — `PORT void SetRCVRNOBRun (int stype, int id, int run)`
  Sets rcvrnobrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBMode()`** — L50 — `PORT void SetRCVRNOBMode (int stype, int id, int mode)`
  Sets rcvrnobmode — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBBuffsize()`** — L56 — `PORT void SetRCVRNOBBuffsize (int stype, int id, int size)`
  Sets rcvrnobbuffsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetRCVRNOBSamplerate()`** — L62 — `PORT void SetRCVRNOBSamplerate (int stype, int id, int rate)`
  Sets rcvrnobsamplerate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetRCVRNOBTau()`** — L68 — `PORT void SetRCVRNOBTau (int stype, int id, double tau)`
  Sets rcvrnobtau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBHangtime()`** — L74 — `PORT void SetRCVRNOBHangtime (int stype, int id, double time)`
  Sets rcvrnobhangtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBAdvtime()`** — L80 — `PORT void SetRCVRNOBAdvtime (int stype, int id, double time)`
  Sets rcvrnobadvtime — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBBacktau()`** — L86 — `PORT void SetRCVRNOBBacktau (int stype, int id, double tau)`
  Sets rcvrnobbacktau — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRCVRNOBThreshold()`** — L92 — `PORT void SetRCVRNOBThreshold (int stype, int id, double thresh)`
  Sets rcvrnobthreshold — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/znobII.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
