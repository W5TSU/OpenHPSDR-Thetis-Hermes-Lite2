# `ChannelMaster/sync.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/router.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×3)
  - `wdsp/div.c` (calls ×3)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/calcc.c` (calls ×1)
- Most-referenced symbols from other files: `create_sync()` (×1), `destroy_sync()` (×1), `InboundBlock()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_sync()`** — L32 — `void create_sync()`
  Constructor for the `sync` block: allocates its state/buffers and computes initial coefficients.
  Called by: `CreateRadio()` (`ChannelMaster/cmsetup.c`)
- **`destroy_sync()`** — L38 — `void destroy_sync()`
  Destroys the `sync` block, freeing its allocated buffers.
  Called by: `DestroyRadio()` (`ChannelMaster/cmsetup.c`)
- **`InboundBlock()`** — L44 — `PORT void InboundBlock (int id, int nsamples, double** data)`
  Called by: `xrouter()` (`ChannelMaster/router.c`)
- **`SetPSTxIdx()`** — L69 — `PORT void SetPSTxIdx (int id, int idx)`
  Sets pstx idx — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetPSRxIdx()`** — L75 — `PORT void SetPSRxIdx (int id, int idx)`
  Sets psrx idx — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/sync.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
