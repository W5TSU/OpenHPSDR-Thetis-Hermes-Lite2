# `ChannelMaster/analyzers.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Attaches wdsp spectrum analyzers to ChannelMaster streams (RX/TX displays).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×3)
  - `wdsp/analyzer.c` (calls ×2)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `wdsp/utilities.c` (calls ×1)
  - `wdsp/siphon.c` (calls ×1)
- Most-referenced symbols from other files: `create_analyzer_alloc()` (×1), `destroy_analyzer_alloc()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_analyzer_alloc()`** — L32 — `void* create_analyzer_alloc (int m_analyzers, int base_disp)`
  create the analyzer allocator; called from cmaster.c
  Called by: `create_cmaster()` (`ChannelMaster/cmaster.c`)
- **`destroy_analyzer_alloc()`** — L66 — `int destroy_analyzer_alloc ()`
  destroy the analyzer allocator; called from cmaster.c
  Called by: `destroy_cmaster()` (`ChannelMaster/cmaster.c`)
- **`tx_analyzers()`** — L87 — `void tx_analyzers ()`
  Called by: `alloc_analyzer()` (same file), `free_analyzer()` (same file), `run_analyzer()` (same file)
- **`alloc_analyzer()`** — L121 — `PORT int alloc_analyzer (int stype, int id, int max_fft_size)`
  Call from console to allocate a new analyzer. The 'disp' number (needed for SetAnalyzer(...), GetPixels(...), etc.), is returned. If no analyzer is available, the return value is negative.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`free_analyzer()`** — L157 — `PORT int free_analyzer (int disp)`
  Call from console to free an analyzer that is no longer in use.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`run_analyzer()`** — L170 — `PORT void run_analyzer(int disp, int run)`
  Call from console to turn OFF/ON one of the additional analyzers.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/analyzers.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
