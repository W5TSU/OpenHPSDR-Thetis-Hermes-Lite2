# `ChannelMaster/cmbuffs.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/networkproto1.c` (calls ×2)
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/router.c` (calls ×1)
  - `ChannelMaster/sync.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmaster.c` (calls ×1)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
- Most-referenced symbols from other files: `Inbound()` (×5), `create_cmbuffs()` (×1), `destroy_cmbuffs()` (×1), `SetCMRingOutsize()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`start_cmthread()`** — L29 — `void start_cmthread (int id)`
  Called by: `create_cmbuffs()` (same file), `SetCMRingOutsize()` (same file)
- **`create_cmbuffs()`** — L35 — `void create_cmbuffs (int id, int accept, int max_insize, int max_outsize, int outsize)`
  Constructor for the `cmbuffs` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_cmaster()` (`ChannelMaster/cmaster.c`)
- **`destroy_cmbuffs()`** — L60 — `void destroy_cmbuffs (int id)`
  Destroys the `cmbuffs` block, freeing its allocated buffers.
  Called by: `destroy_cmaster()` (`ChannelMaster/cmaster.c`)
- **`flush_cmbuffs()`** — L78 — `void flush_cmbuffs (int id)`
  Flushes (zeroes) the `cmbuffs` block’s internal buffers/state.
  Called by: `SetCMRingOutsize()` (same file)
- **`Inbound()`** — L88 — `PORT void Inbound (int id, int nsamples, double* in)`
  Called by: `ReadThreadMainLoop()` (`ChannelMaster/network.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`), `xrouter()` (`ChannelMaster/router.c`), `InboundBlock()` (`ChannelMaster/sync.c`)
- **`cmdata()`** — L123 — `void cmdata (int id, double* out)`
  Called by: `cm_main()` (same file)
- **`cm_main()`** — L151 — `void cm_main (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetCMRingOutsize()`** — L170 — `void SetCMRingOutsize (int id, int size)`
  Sets cmring outsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
