# `ChannelMaster/pro.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/networkproto1.c` (calls ×5)
- Uses (outgoing references to other files):
  - `ChannelMaster/pro.h` (imports ×1)
- Most-referenced symbols from other files: `destroy_pro()` (×3), `create_pro()` (×1), `xpro()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_pro()`** — L29 — `PRO create_pro ( int run, int psize, int npacks, int lpacks )`
  Constructor for the `pro` block: allocates its state/buffers and computes initial coefficients.
  Called by: `SendStartToMetis()` (`ChannelMaster/networkproto1.c`)
- **`destroy_pro()`** — L55 — `void destroy_pro ( PRO a )`
  Destroys the `pro` block, freeing its allocated buffers.
  Called by: `SendStopToMetis()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop()` (`ChannelMaster/networkproto1.c`), `MetisReadThreadMainLoop_HL2()` (`ChannelMaster/networkproto1.c`)
- **`xpro()`** — L67 — `void xpro (PRO a, unsigned int seqnum, char* buffer)`
  Runs the `pro` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `MetisReadDirect()` (`ChannelMaster/networkproto1.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/pro.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
