# `ChannelMaster/cmaster.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Stream lifecycle and top-level orchestration: creates channels, starts/stops audio and network streams.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmsetup.c` (calls ×2)
  - `ChannelMaster/netInterface.c` (calls ×2)
  - `ChannelMaster/cmbuffs.c` (calls ×1)
- Uses (outgoing references to other files):
  - `ChannelMaster/cmsetup.c` (calls ×30)
  - `ChannelMaster/aamix.c` (calls ×15)
  - `wdsp/channel.c` (calls ×8)
  - `ChannelMaster/ivac.c` (calls ×7)
  - `wdsp/analyzer.c` (calls ×5)
  - `ChannelMaster/ilv.c` (calls ×5)
  - `ChannelMaster/sidetone.c` (calls ×5)
  - `wdsp/dexp.c` (calls ×5)
  - `wdsp/eer.c` (calls ×5)
  - `ChannelMaster/txgain.c` (calls ×4)
  - `ChannelMaster/cmasio.c` (calls ×3)
  - `wdsp/nob.c` (calls ×3)
  - …and 12 more files
- Most-referenced symbols from other files: `create_cmaster()` (×1), `destroy_cmaster()` (×1), `xcmaster()` (×1), `SendpOutboundRx()` (×1), `SendpOutboundTx()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_rcvr()`** — L33 — `void create_rcvr()`
  standard receiver
  Called by: `create_cmaster()` (same file)
- **`destroy_rcvr()`** — L96 — `void destroy_rcvr()`
  Destroys the `rcvr` block, freeing its allocated buffers.
  Called by: `destroy_cmaster()` (same file)
- **`create_xmtr()`** — L112 — `void create_xmtr()`
  standard transmitter
  Called by: `create_cmaster()` (same file)
- **`destroy_xmtr()`** — L255 — `void destroy_xmtr()`
  Destroys the `xmtr` block, freeing its allocated buffers.
  Called by: `destroy_cmaster()` (same file)
- **`create_cmaster()`** — L273 — `void create_cmaster()`
  Constructor for the `cmaster` block: allocates its state/buffers and computes initial coefficients.
  Called by: `CreateRadio()` (`ChannelMaster/cmsetup.c`)
- **`destroy_cmaster()`** — L322 — `void destroy_cmaster()`
  Destroys the `cmaster` block, freeing its allocated buffers.
  Called by: `DestroyRadio()` (`ChannelMaster/cmsetup.c`)
- **`xcmaster()`** — L339 — `PORT void xcmaster (int stream)`
  Runs the `cmaster` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `cm_main()` (`ChannelMaster/cmbuffs.c`)
- **`SendpOutboundRx()`** — L407 — `PORT void SendpOutboundRx (void (*Outbound)(int id, int nsamples, double* buff))`
  Called by: `create_rnet()` (`ChannelMaster/netInterface.c`)
- **`SendpOutboundTx()`** — L414 — `PORT void SendpOutboundTx(void (*Outbound)(int id, int nsamples, double* buff))`
  Called by: `create_rnet()` (`ChannelMaster/netInterface.c`)
- **`SendpOutboundTCIRxIQ()`** — L422 — `PORT void SendpOutboundTCIRxIQ (void (*Outbound)(int id, int nsamples, double* buff))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SendpInboundTCITxAudio()`** — L428 — `PORT void SendpInboundTCITxAudio (void (*Inbound)(int nsamples, double* buff))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRXTCIRun()`** — L434 — `PORT void SetRXTCIRun (int active)`
  Sets rxtcirun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTXTCIAudioRun()`** — L440 — `PORT void SetTXTCIAudioRun (int txid, int active)`
  Sets txtciaudio run — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetRunPanadapter()`** — L447 — `PORT void SetRunPanadapter (int id, int run)`
  Sets run panadapter — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetXcmInrate()`** — L453 — `PORT void SetXcmInrate (int in_id, int rate)`
  Sets xcm inrate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetCMAudioOutrate()`** — L510 — `PORT void SetCMAudioOutrate (int in_id, int rate)`
  Sets cmaudio outrate — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRcvrChannelOutrate()`** — L522 — `PORT void SetRcvrChannelOutrate (int rcvr_id, int rate, int state)`
  Sets rcvr channel outrate — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetXmtrChannelOutrate()`** — L549 — `PORT void SetXmtrChannelOutrate (int xmtr_id, int rate, int state)`
  Sets xmtr channel outrate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAntiVOXSourceStates()`** — L583 — `PORT void SetAntiVOXSourceStates (int txid, int streams, int states)`
  Sets anti voxsource states — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetAntiVOXSourceWhat()`** — L590 — `PORT void SetAntiVOXSourceWhat (int txid, int stream, int state)`
  Sets anti voxsource what — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/cmaster.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
