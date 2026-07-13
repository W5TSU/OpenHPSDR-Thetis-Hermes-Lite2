# `ChannelMaster/tci.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** TCI (Transceiver Control Interface) TCP server for SDC/logger integration at the audio layer.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×3)
  - `ChannelMaster/pipe.c` (calls ×3)
- Uses (outgoing references to other files):
  - `ChannelMaster/aamix.c` (calls ×14)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `SetTCIRxAudioRate()` (×1), `SetTCIRxAudioSize()` (×1), `SetTCITxMonitorRate()` (×1), `create_tci()` (×1), `destroy_tci()` (×1), `xtciOUT()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`get_tci_audio_mix_state()`** — L65 — `static void get_tci_audio_mix_state(int rx, long* active, long* what)`
  Called by: `apply_tci_audio_mix_state()` (same file), `create_tci_audio_mixer()` (same file)
- **`apply_tci_audio_mix_state()`** — L88 — `static void apply_tci_audio_mix_state(int rx)`
  Called by: `SetTCIRxAudioRate()` (same file), `SetTCIRxAudioSize()` (same file), `SetTCITxMonitorRate()` (same file), `SetTCIRxAudioMox()` (same file), `SetTCIRxAudioMon()` (same file)
- **`tci_audio_out()`** — L118 — `static void tci_audio_out(int id, int nsamples, double* buff)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_tci_audio_mixer()`** — L124 — `static void create_tci_audio_mixer(int rx)`
  Constructor for the `tci_audio_mixer` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_tci()` (same file)
- **`destroy_tci_audio_mixer()`** — L162 — `static void destroy_tci_audio_mixer(int rx)`
  Destroys the `tci_audio_mixer` block, freeing its allocated buffers.
  Called by: `destroy_tci()` (same file)
- **`create_tci()`** — L179 — `void create_tci()`
  Constructor for the `tci` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_pipe()` (`ChannelMaster/pipe.c`)
- **`destroy_tci()`** — L192 — `void destroy_tci()`
  Destroys the `tci` block, freeing its allocated buffers.
  Called by: `destroy_pipe()` (`ChannelMaster/pipe.c`)
- **`xtciOUT()`** — L200 — `void xtciOUT(int id, int stream, double* data)`
  Runs the `tciOUT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`SendpOutboundTCIRxAudio()`** — L220 — `PORT void SendpOutboundTCIRxAudio(void (*Outbound)(int id, int nsamples, double* buff))`
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTCIRxAudioRate()`** — L226 — `void SetTCIRxAudioRate(int id, int rate)`
  Sets tcirx audio rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetTCIRxAudioSize()`** — L250 — `void SetTCIRxAudioSize(int id, int size)`
  Sets tcirx audio size — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetTCITxMonitorRate()`** — L267 — `void SetTCITxMonitorRate(int id, int rate)`
  Sets tcitx monitor rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetTCIRxAudioMox()`** — L288 — `PORT void SetTCIRxAudioMox(int id, int mox)`
  Sets tcirx audio mox — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTCIRxAudioMon()`** — L299 — `PORT void SetTCIRxAudioMon(int id, int mon)`
  Sets tcirx audio mon — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.
- **`SetTCIRxAudioMonVol()`** — L310 — `PORT void SetTCIRxAudioMonVol(int id, double vol)`
  Sets tcirx audio mon vol — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/cmaster.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/tci.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
