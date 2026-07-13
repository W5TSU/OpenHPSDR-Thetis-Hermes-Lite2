# `ChannelMaster/aamix.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Audio mixers (monitor mix, multi-RX audio combination) with per-input gain and slew.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×15)
  - `ChannelMaster/tci.c` (calls ×14)
  - `ChannelMaster/ivac.c` (calls ×12)
- Uses (outgoing references to other files):
  - `wdsp/resample.c` (calls ×8)
  - `wdsp/utilities.c` (calls ×5)
  - `ChannelMaster/cmcomm.h` (imports ×1)
- Most-referenced symbols from other files: `create_aamix()` (×7), `destroy_aamix()` (×6), `SetAAudioMixState()` (×5), `SetAAudioMixWhat()` (×4), `SetAAudioStreamRate()` (×4), `xMixAudio()` (×3), `SetAAudioMixVol()` (×3), `SetAAudioRingInsize()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`mix_main()`** — L32 — `void mix_main (void *pargs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`start_mixthread()`** — L51 — `void start_mixthread (AAMIX a)`
  Called by: `create_aamix()` (same file), `open_mixer()` (same file)
- **`create_aaslew()`** — L69 — `void create_aaslew (AAMIX a)`
  Constructor for the `aaslew` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_aamix()` (same file)
- **`destroy_aaslew()`** — L109 — `void destroy_aaslew (AAMIX a)`
  Destroys the `aaslew` block, freeing its allocated buffers.
  Called by: `destroy_aamix()` (same file)
- **`flush_aaslew()`** — L117 — `void flush_aaslew (AAMIX a)`
  Flushes (zeroes) the `aaslew` block’s internal buffers/state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`create_aamix()`** — L128 — `void* create_aamix ( int id, int outbound_id, int ringinsize, int outsize,`
  Constructor for the `aamix` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_xmtr()` (`ChannelMaster/cmaster.c`), `create_cmaster()` (`ChannelMaster/cmaster.c`), `create_ivac()` (`ChannelMaster/ivac.c`), `SetIVACaudioRate()` (`ChannelMaster/ivac.c`), `SetIVACtxmonRate()` (`ChannelMaster/ivac.c`), `SetIVACaudioSize()` (`ChannelMaster/ivac.c`) — and 1 more
- **`destroy_aamix()`** — L209 — `void destroy_aamix (void* ptr, int id)`
  Destroys the `aamix` block, freeing its allocated buffers.
  Called by: `destroy_xmtr()` (`ChannelMaster/cmaster.c`), `destroy_cmaster()` (`ChannelMaster/cmaster.c`), `SetIVACaudioRate()` (`ChannelMaster/ivac.c`), `SetIVACtxmonRate()` (`ChannelMaster/ivac.c`), `SetIVACaudioSize()` (`ChannelMaster/ivac.c`), `destroy_tci_audio_mixer()` (`ChannelMaster/tci.c`)
- **`xMixAudio()`** — L237 — `void xMixAudio (void* ptr, int id, int stream, double* data)`
  loads data from a buffer into an audio mixer ring
  Called by: `xcmaster()` (`ChannelMaster/cmaster.c`), `xvacOUT()` (`ChannelMaster/ivac.c`), `xtciOUT()` (`ChannelMaster/tci.c`)
- **`upslew()`** — L280 — `void upslew (AAMIX a)`
  Called by: `xaamix()` (same file)
- **`downslew()`** — L345 — `void downslew (AAMIX a)`
  Called by: `xaamix()` (same file)
- **`xaamix()`** — L423 — `void xaamix (AAMIX a)`
  pulls data from audio rings and mixes with output
  Called by: `mix_main()` (same file)
- **`flush_mix_ring()`** — L461 — `void flush_mix_ring (AAMIX a, int stream)`
  Flushes (zeroes) the `mix_ring` block’s internal buffers/state.
  Called by: `close_mixer()` (same file)
- **`close_mixer()`** — L471 — `void close_mixer (AAMIX a)`
  Called by: `SetAAudioMixState()` (same file), `SetAAudioMixStates()` (same file), `SetAAudioRingInsize()` (same file), `SetAAudioRingOutsize()` (same file), `SetAAudioOutRate()` (same file)
- **`open_mixer()`** — L493 — `void open_mixer (AAMIX a)`
  Called by: `SetAAudioMixState()` (same file), `SetAAudioMixStates()` (same file), `SetAAudioRingInsize()` (same file), `SetAAudioRingOutsize()` (same file), `SetAAudioOutRate()` (same file)
- **`SetAAudioMixOutputPointer()`** — L513 — `void SetAAudioMixOutputPointer(void* ptr, int id, void (*Outbound)(int id, int nsamples, double* buff))`
  Sets aaudio mix output pointer — API setter, typically called from the console via P/Invoke.
  Called by: `SendpOutboundRx()` (`ChannelMaster/cmaster.c`)
- **`SetAAudioMixState()`** — L521 — `PORT void SetAAudioMixState (void* ptr, int id, int stream, int state)`
  Sets aaudio mix state — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`), `apply_tci_audio_mix_state()` (`ChannelMaster/tci.c`), `SetTCIRxAudioRate()` (`ChannelMaster/tci.c`), `SetTCITxMonitorRate()` (`ChannelMaster/tci.c`)
- **`SetAAudioMixStates()`** — L554 — `PORT void SetAAudioMixStates (void* ptr, int id, int streams, int states)`
  SetAAudioMixStates() is an alternative to SetAAudioMixState() that can be used to set multiple mix states with only a single call. 'streams' has one bit per mix state that you want to set and 'states' has one bit specifying the state of each stream that you want to set. For example, if you want…
  Called by: `SetAntiVOXSourceStates()` (`ChannelMaster/cmaster.c`)
- **`SetAAudioMixWhat()`** — L584 — `PORT void SetAAudioMixWhat (void* ptr, int id, int stream, int state)`
  Sets aaudio mix what — API setter, typically called from the console via P/Invoke.
  Called by: `SetAntiVOXSourceWhat()` (`ChannelMaster/cmaster.c`), `SetIVACmox()` (`ChannelMaster/ivac.c`), `SetIVACmon()` (`ChannelMaster/ivac.c`), `apply_tci_audio_mix_state()` (`ChannelMaster/tci.c`)
- **`SetAAudioMixVolume()`** — L596 — `PORT void SetAAudioMixVolume (void* ptr, int id, double volume)`
  Sets aaudio mix volume — API setter, typically called from the console via P/Invoke.
  Called by: `SetIVACrxscale()` (`ChannelMaster/ivac.c`)
- **`SetAAudioMixVol()`** — L610 — `PORT void SetAAudioMixVol (void* ptr, int id, int stream, double vol)`
  Sets aaudio mix vol — API setter, typically called from the console via P/Invoke.
  Called by: `SetIVACmonVol()` (`ChannelMaster/ivac.c`), `create_tci_audio_mixer()` (`ChannelMaster/tci.c`), `SetTCIRxAudioMonVol()` (`ChannelMaster/tci.c`)
- **`SetAAudioRingInsize()`** — L622 — `void SetAAudioRingInsize (void* ptr, int id, int size)`
  Sets aaudio ring insize — API setter, typically called from the console via P/Invoke.
  Called by: `SetCMAudioOutrate()` (`ChannelMaster/cmaster.c`), `SetTCIRxAudioSize()` (`ChannelMaster/tci.c`)
- **`SetAAudioRingOutsize()`** — L644 — `void SetAAudioRingOutsize (void* ptr, int id, int size)`
  Sets aaudio ring outsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetCMAudioOutrate()` (`ChannelMaster/cmaster.c`), `SetTCIRxAudioSize()` (`ChannelMaster/tci.c`)
- **`SetAAudioOutRate()`** — L656 — `void SetAAudioOutRate (void* ptr, int id, int rate)`
  Sets aaudio out rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetCMAudioOutrate()` (`ChannelMaster/cmaster.c`), `SetTCIRxAudioRate()` (`ChannelMaster/tci.c`)
- **`SetAAudioStreamRate()`** — L683 — `void SetAAudioStreamRate (void* ptr, int id, int mixinid, int rate)`
  Sets aaudio stream rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetTCIRxAudioRate()` (`ChannelMaster/tci.c`), `SetTCITxMonitorRate()` (`ChannelMaster/tci.c`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/aamix.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
