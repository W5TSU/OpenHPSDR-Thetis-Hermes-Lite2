# `ChannelMaster/pipe.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/cmaster.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×7)
  - `ChannelMaster/ivac.c` (calls ×4)
  - `ChannelMaster/radae.c` (calls ×4)
  - `wdsp/nob.c` (calls ×3)
  - `wdsp/nobII.c` (calls ×3)
  - `ChannelMaster/tci.c` (calls ×3)
  - `wdsp/siphon.c` (calls ×3)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/analyzer.c` (calls ×1)
- Most-referenced symbols from other files: `xpipe()` (×1), `create_pipe()` (×1), `destroy_pipe()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_spc0()`** — L32 — `void create_spc0()`
  Constructor for the `spc0` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_pipe()` (same file)
- **`destroy_spc0()`** — L69 — `void destroy_spc0()`
  Destroys the `spc0` block, freeing its allocated buffers.
  Called by: `destroy_pipe()` (same file)
- **`create_pipe()`** — L79 — `void create_pipe()`
  Constructor for the `pipe` block: allocates its state/buffers and computes initial coefficients.
  Called by: `CreateRadio()` (`ChannelMaster/cmsetup.c`)
- **`destroy_pipe()`** — L120 — `void destroy_pipe()`
  Destroys the `pipe` block, freeing its allocated buffers.
  Called by: `DestroyRadio()` (`ChannelMaster/cmsetup.c`)
- **`xplaywave()`** — L135 — `void xplaywave(int rx, int state, double* data)`
  Runs the `playwave` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (same file)
- **`xrecordwave()`** — L144 — `void xrecordwave(int rx, int state, int pos, double* data)`
  Runs the `recordwave` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (same file)
- **`xscope()`** — L153 — `void xscope(int rx, int state, double* data)`
  Runs the `scope` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (same file)
- **`xpipe()`** — L162 — `void xpipe (int stream, int pos, double** buffs)`
  Runs the `pipe` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`SendCBCreateScope()`** — L273 — `PORT void SendCBCreateScope (void (__stdcall *create_Scope)(int id))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendCBScope()`** — L279 — `PORT void SendCBScope (int id, void (__stdcall *xscope)(int state, double* data))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetScopeRun()`** — L285 — `PORT void SetScopeRun(int id, int run)`
  Sets scope run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendCBCreateWRecord()`** — L291 — `PORT void SendCBCreateWRecord (void (__stdcall *create_WaveRecord)(int id))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendCBWaveRecorder()`** — L297 — `PORT void SendCBWaveRecorder (int id, void (__stdcall *xrecordwave)(int state, int pos, double* data))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetWaveRecorderRun()`** — L303 — `PORT void SetWaveRecorderRun(int id, int run)`
  Sets wave recorder run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendCBCreateWPlay()`** — L309 — `PORT void SendCBCreateWPlay (void (__stdcall *create_WavePlay)(int id))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendCBWavePlayer()`** — L315 — `PORT void SendCBWavePlayer (int id, void (__stdcall *xplaywave)(int state, double* data))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetWavePlayerRun()`** — L321 — `PORT void SetWavePlayerRun(int id, int run)`
  Sets wave player run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTopPan3Run()`** — L327 — `PORT void SetTopPan3Run (int run)`
  Sets top pan3 run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTXVAC()`** — L333 — `PORT void SetTXVAC (int txid, int txvac)`
  Sets txvac — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/pipe.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
