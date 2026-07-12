# `ChannelMaster/cmsetup.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** System-wide setup: instantiates buffers, mixers, VAC, analyzers per radio model.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×30)
  - `ChannelMaster/pipe.c` (calls ×7)
  - `ChannelMaster/analyzers.c` (calls ×3)
  - `ChannelMaster/sync.c` (calls ×3)
  - `ChannelMaster/networkproto1.c` (calls ×2)
  - `ChannelMaster/ivac.c` (calls ×1)
  - `ChannelMaster/network.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmaster.c` (calls ×2)
  - `ChannelMaster/pipe.c` (calls ×2)
  - `ChannelMaster/sync.c` (calls ×2)
  - `ChannelMaster/cmaster.h` (imports ×1)
  - `ChannelMaster/cmsetup.h` (imports ×1)
- Most-referenced symbols from other files: `inid()` (×15), `chid()` (×10), `getbuffsize()` (×8), `txid()` (×4), `rxid()` (×3), `stype()` (×3), `sp0id()` (×2), `mixinid()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`SetRadioStructure()`** — L32 — `PORT void SetRadioStructure ( int cmSTREAM, int cmRCVR, int cmXMTR,`
  set radio structure, call this first these parameters are used by create_cmaster() to determine units to create & buffer sizes
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`set_cmdefault_rates()`** — L59 — `PORT void set_cmdefault_rates ( int* xcm_inrates, int aud_outrate, int* rcvr_ch_outrates,`
  set default sample rates, call this before 'create'
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CreateRadio()`** — L88 — `PORT void CreateRadio()`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`DestroyRadio()`** — L96 — `PORT void DestroyRadio()`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`getbuffsize()`** — L105 — `PORT int getbuffsize (int rate)`
  buffer sizes are a function of sample rate to yield constant latency
  Called by: `set_cmdefault_rates()` (same file), `create_rcvr()` (`ChannelMaster/cmaster.c`), `create_xmtr()` (`ChannelMaster/cmaster.c`), `create_cmaster()` (`ChannelMaster/cmaster.c`), `SetXcmInrate()` (`ChannelMaster/cmaster.c`), `SetCMAudioOutrate()` (`ChannelMaster/cmaster.c`) — and 3 more
- **`getInputRate()`** — L113 — `PORT int getInputRate (int stype, int id)`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`getChannelOutputRate()`** — L119 — `PORT int getChannelOutputRate (int stype, int id)`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`rxid()`** — L147 — `int rxid (int stream)`
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`), `SetXcmInrate()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`txid()`** — L152 — `int txid (int stream)`
  Called by: `tx_analyzers()` (`ChannelMaster/analyzers.c`), `xcmaster()` (`ChannelMaster/cmaster.c`), `SetXcmInrate()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`sp0id()`** — L157 — `int sp0id (int stream)`
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`stype()`** — L162 — `int stype (int stream)`
  Called by: `chid()` (same file), `mixinid()` (same file), `xcmaster()` (`ChannelMaster/cmaster.c`), `SetXcmInrate()` (`ChannelMaster/cmaster.c`), `xpipe()` (`ChannelMaster/pipe.c`)
- **`chid()`** — L175 — `PORT int chid (int stream, int subrx)`
  Called by: `tx_analyzers()` (`ChannelMaster/analyzers.c`), `create_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `create_xmtr()` (`ChannelMaster/cmaster.c`), `destroy_xmtr()` (`ChannelMaster/cmaster.c`), `xcmaster()` (`ChannelMaster/cmaster.c`) — and 4 more
- **`inid()`** — L194 — `PORT int inid (int stype, int id)`
  Called by: `getInputRate()` (same file), `alloc_analyzer()` (`ChannelMaster/analyzers.c`), `create_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `create_xmtr()` (`ChannelMaster/cmaster.c`), `destroy_xmtr()` (`ChannelMaster/cmaster.c`) — and 10 more
- **`mixinid()`** — L216 — `int mixinid (int stream, int subrx)`
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmsetup.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
