# `ChannelMaster/router.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** The signal routing matrix — which input feeds which DSP channel and which output.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/networkproto1.c` (calls ×3)
  - `ChannelMaster/cmaster.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmbuffs.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `ChannelMaster/sync.c` (calls ×1)
- Most-referenced symbols from other files: `xrouter()` (×4), `create_router()` (×1), `destroy_router()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_router()`** — L32 — `void* create_router( int id )`
  Constructor for the `router` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_cmaster()` (`ChannelMaster/cmaster.c`)
- **`destroy_router()`** — L43 — `void destroy_router(void* ptr, int id)`
  Destroys the `router` block, freeing its allocated buffers.
  Called by: `destroy_cmaster()` (`ChannelMaster/cmaster.c`)
- **`flush_router()`** — L52 — `void flush_router(ROUTER a)`
  Flushes (zeroes) the `router` block’s internal buffers/state.
  Called by: `LoadRouterAll()` (same file)
- **`xrouter()`** — L70 — `PORT void xrouter(void* ptr, int id, int source, int nsamples, double* data)`
  Runs the `router` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `ReadThreadMainLoop()` (`ChannelMaster/network.c`), `twist()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`)
- **`LoadRouterAll()`** — L110 — `PORT void LoadRouterAll( void* ptr, int id, int sources,`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`LoadRouterControlBit()`** — L145 — `PORT void LoadRouterControlBit(void* ptr, int id, int var_number, int bit)`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/router.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
