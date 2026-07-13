# `ChannelMaster/ilv.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
  - `ChannelMaster/zeer.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `create_ilv()` (×1), `destroy_ilv()` (×1), `xilv()` (×1), `SetILVOutputPointer()` (×1), `pSetILVInsize()` (×1), `pSetILVRun()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_ilv()`** — L31 — `ILV create_ilv ( int run, int outbound_id, int insize, int ninputs,`
  Constructor for the `ilv` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`destroy_ilv()`** — L51 — `void destroy_ilv (ILV a)`
  Destroys the `ilv` block, freeing its allocated buffers.
  Called by: `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`xilv()`** — L57 — `void xilv (ILV a, double** data)`
  Runs the `ilv` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`)
- **`SetILVOutputPointer()`** — L97 — `void SetILVOutputPointer (int xmtr_id, void(*Outbound)(int id, int nsamples, double* buff))`
  Sets ilvoutput pointer — API setter, typically called from the console via P/Invoke.
  Called by: `SendpOutboundTx()` (`ChannelMaster/cmaster.c`)
- **`SetILVRun()`** — L103 — `PORT void SetILVRun (int xmtr_id, int run)`
  Sets ilvrun — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetILVWhat()`** — L113 — `PORT void SetILVWhat(int xmtr_id, int stream, int state)`
  Sets ilvwhat — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetILVInsize()`** — L123 — `PORT void SetILVInsize(int xmtr_id, int size)`
  Sets ilvinsize — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetILVOutboundId()`** — L130 — `PORT void SetILVOutboundId(int xmtr_id, int obid)`
  Sets ilvoutbound id — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`pSetILVRun()`** — L137 — `void pSetILVRun(ILV a, int run)`
  Called by: `SetEERRun()` (`ChannelMaster/zeer.c`)
- **`pSetILVInsize()`** — L145 — `void pSetILVInsize(ILV a, int size)`
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ilv.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
