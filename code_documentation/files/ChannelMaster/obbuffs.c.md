# `ChannelMaster/obbuffs.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/netInterface.c` (calls ×2)
  - `ChannelMaster/cmasio.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/network.c` (calls ×1)
  - `ChannelMaster/obbuffs.h` (imports ×1)
- Most-referenced symbols from other files: `OutBound()` (×1), `destroy_obbuffs()` (×1), `create_obbuffs()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`start_obthread()`** — L37 — `void start_obthread (int id)`
  Called by: `create_obbuffs()` (same file), `SetOBRingOutsize()` (same file)
- **`create_obbuffs()`** — L43 — `void create_obbuffs (int id, int accept, int max_insize, int outsize)`
  Constructor for the `obbuffs` block: allocates its state/buffers and computes initial coefficients.
  Called by: `UpdateRadioProtocolSampleSize()` (`ChannelMaster/netInterface.c`)
- **`destroy_obbuffs()`** — L68 — `void destroy_obbuffs (int id)`
  Destroys the `obbuffs` block, freeing its allocated buffers.
  Called by: `StopAudio()` (`ChannelMaster/netInterface.c`)
- **`flush_obbuffs()`** — L88 — `void flush_obbuffs (int id)`
  Flushes (zeroes) the `obbuffs` block’s internal buffers/state.
  Called by: `SetOBRingOutsize()` (same file)
- **`OutBound()`** — L98 — `PORT void OutBound (int id, int nsamples, double* in)`
  Called by: `asioOUT()` (`ChannelMaster/cmasio.c`)
- **`obdata()`** — L132 — `void obdata (int id, double* out)`
  Called by: `ob_main()` (same file)
- **`ob_main()`** — L153 — `void ob_main (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetOBRingOutsize()`** — L175 — `void SetOBRingOutsize (int id, int size)`
  Sets obring outsize — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/obbuffs.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
