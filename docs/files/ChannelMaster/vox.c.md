# `ChannelMaster/vox.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** VOX detection and TX gain staging.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_vox()`** — L29 — `VOX create_vox (int id, int run, int size, double* in, int mode, double thresh)`
  Constructor for the `vox` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_voxEXT()` (same file)
- **`destroy_vox()`** — L45 — `void destroy_vox (VOX a)`
  Destroys the `vox` block, freeing its allocated buffers.
  Called by: `destroy_voxEXT()` (same file)
- **`flush_vox()`** — L51 — `void flush_vox (VOX a)`
  Flushes (zeroes) the `vox` block’s internal buffers/state.
  Called by: `flush_voxEXT()` (same file)
- **`xvox()`** — L58 — `void xvox (VOX a)`
  Runs the `vox` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xvoxEXT()` (same file)
- **`create_voxEXT()`** — L112 — `void create_voxEXT (int channel, int run, int size, double* in, int mode, double thresh)`
  Constructor for the `voxEXT` block: allocates its state/buffers and computes initial coefficients.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`destroy_voxEXT()`** — L117 — `void destroy_voxEXT (int channel)`
  Destroys the `voxEXT` block, freeing its allocated buffers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`flush_voxEXT()`** — L122 — `void flush_voxEXT (int channel)`
  Flushes (zeroes) the `voxEXT` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xvoxEXT()`** — L127 — `void xvoxEXT (int channel, double* in)`
  Runs the `voxEXT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SendCBPushVox()`** — L140 — `PORT void SendCBPushVox (int id, void (__stdcall *pushvox)(int id, int active))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTXAVoxRun()`** — L146 — `PORT void SetTXAVoxRun (int id, int run)`
  Sets txavox run — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAVoxSize()`** — L155 — `PORT void SetTXAVoxSize (int id, int size)`
  Sets txavox size — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetTXAVoxThresh()`** — L164 — `PORT void SetTXAVoxThresh (int id, double thresh)`
  Sets txavox thresh — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`GetTXAVoxPeak()`** — L173 — `PORT void GetTXAVoxPeak (int id, double* peak)`
  Returns txavox peak — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/vox.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
