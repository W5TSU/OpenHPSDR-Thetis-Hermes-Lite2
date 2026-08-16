# `ChannelMaster/radae.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** RADE V1 neural-mode RX decode block (HL2 fork addition, `FreeDV` branch). RX-only: hooked into `pipe.c`'s `xpipe()` hot path, drives the `rade_c`/`rade.lib` native decoder plus `lpcnet`/`fargan` speech synthesis, gated by `RXRadaeEnabled` (CAT `ZZDW`/`ZZDZ`, Setup → DSP → FreeDV tab).

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/pipe.c` (calls ×4)
- Uses (outgoing references to other files):
  - `ChannelMaster/radae_micdsp.c` (calls ×11)
  - `wdsp/resample.c` (calls ×7)
  - `wdsp/rmatch.c` (calls ×5)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/cmsetup.c` (calls ×1)
  - `ChannelMaster/r8brain_wrap.h` (imports ×1)
  - `ChannelMaster/radae.h` (imports ×1)
  - `ChannelMaster/radae_micdsp.h` (imports ×1)
- Most-referenced symbols from other files: `create_radae()` (×1), `destroy_radae()` (×1), `GetRadaeRxEnabled()` (×1), `xradae_rx()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`fifo_push_check()`** — L305 — `static void fifo_push_check(float* buf, int* n, int cap, const float* src, int count, long* counter, const char* tag)`
  Called by: `xradae_rx()` (same file), `xradae_tx()` (same file)
- **`fifo_pop()`** — L329 — `static void fifo_pop(float* buf, int* n, float* dst, int count)`
  Called by: `xradae_rx()` (same file), `xradae_tx()` (same file)
- **`rebuild_rx_resamplers()`** — L341 — `static void rebuild_rx_resamplers(int rx, int new_outrate)`
  Called by: `xradae_rx()` (same file)
- **`rebuild_tx_resamplers()`** — L351 — `static void rebuild_tx_resamplers(int new_outrate, int outsize)`
  Called by: `xradae_tx()` (same file)
- **`on_radae_text_rx()`** — L386 — `static void on_radae_text_rx(rade_text_t rt, const char* txt_ptr, int length, void* state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`radae_compute_geom_maxima()`** — L422 — `static void radae_compute_geom_maxima(void)`
  Probe RADE V1 (the larger geometry in every dimension) once so that all geometry-dependent buffers can be sized at the max of both protocols. Sizing at the max lets a per-RX live protocol switch reopen a handle without re-allocating any buffer. Real handle geometry is folded in by the caller…
  Called by: `create_radae()` (same file)
- **`create_radae()`** — L443 — `void create_radae(void)`
  Constructor for the `radae` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_pipe()` (`ChannelMaster/pipe.c`)
- **`destroy_radae()`** — L611 — `void destroy_radae(void)`
  Destroys the `radae` block, freeing its allocated buffers.
  Called by: `destroy_pipe()` (`ChannelMaster/pipe.c`)
- **`radae_rx_valid()`** — L692 — `static int radae_rx_valid(int rx)`
  Called by: `SetRadaeRxEnabled()` (same file), `GetRadaeRxEnabled()` (same file), `GetRadaeSync()` (same file), `GetRadaeSnrDb()` (same file), `GetRadaeRxLevelDb()` (same file), `GetRadaeClip()` (same file) — and 13 more
- **`SetRadaeRxEnabled()`** — L694 — `PORT void SetRadaeRxEnabled(int rx, int enable)`
  Sets radae rx enabled — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetRadaeTxEnabled()`** — L707 — `PORT void SetRadaeTxEnabled(int enable)`
  Sets radae tx enabled — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeRxEnabled()`** — L712 — `PORT int GetRadaeRxEnabled(int rx)`
  Returns radae rx enabled — API getter, typically called from the console via P/Invoke.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`GetRadaeTxEnabled()`** — L718 — `PORT int GetRadaeTxEnabled(void)`
  Returns radae tx enabled — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeSync()`** — L723 — `PORT int GetRadaeSync(int rx)`
  Returns radae sync — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRadaeSnrDb()`** — L729 — `PORT int GetRadaeSnrDb(int rx)`
  Returns radae snr db — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`GetRadaeRxLevelDb()`** — L735 — `PORT int GetRadaeRxLevelDb(int rx)`
  Returns radae rx level db — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeClip()`** — L741 — `PORT int GetRadaeClip(int rx)`
  Returns radae clip — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeRemoteCallsign()`** — L750 — `PORT int GetRadaeRemoteCallsign(int rx, char* dst, int max)`
  Returns radae remote callsign — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeRemoteCallsignSeq()`** — L767 — `PORT int GetRadaeRemoteCallsignSeq(int rx)`
  Returns radae remote callsign seq — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeEooDecodePulse()`** — L773 — `PORT int GetRadaeEooDecodePulse(int rx)`
  Returns radae eoo decode pulse — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeTxMicLevelDb()`** — L782 — `PORT int GetRadaeTxMicLevelDb(void)`
  Returns radae tx mic level db — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeTxMicClip()`** — L787 — `PORT int GetRadaeTxMicClip(void)`
  Returns radae tx mic clip — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeFreqOffset()`** — L795 — `PORT float GetRadaeFreqOffset(int rx)`
  Returns radae freq offset — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeFreqOffset()`** — L801 — `PORT void SetRadaeFreqOffset(int rx, float hz)`
  Sets radae freq offset — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`RadaeNotifyEndOfOver()`** — L807 — `PORT void RadaeNotifyEndOfOver(void)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`RadaeNotifyBeginOver()`** — L813 — `PORT void RadaeNotifyBeginOver(void)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeEooFlushed()`** — L823 — `PORT int GetRadaeEooFlushed(void)`
  Returns radae eoo flushed — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeTxSilenceHold()`** — L828 — `PORT void SetRadaeTxSilenceHold(int on)`
  Sets radae tx silence hold — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeLoopbackEnabled()`** — L833 — `PORT void SetRadaeLoopbackEnabled(int rx, int enable)`
  Sets radae loopback enabled — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeLoopbackEnabled()`** — L857 — `PORT int GetRadaeLoopbackEnabled(int rx)`
  Returns radae loopback enabled — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeProtocolV2()`** — L867 — `PORT void SetRadaeProtocolV2(int rx, int on)`
  Sets radae protocol v2 — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeProtocolV2()`** — L947 — `PORT int GetRadaeProtocolV2(int rx)`
  Returns radae protocol v2 — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeTxRx()`** — L953 — `PORT void SetRadaeTxRx(int rx)`
  Sets radae tx rx — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeTxRx()`** — L959 — `PORT int GetRadaeTxRx(void)`
  Returns radae tx rx — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMoxState()`** — L964 — `PORT void SetRadaeMoxState(int mox)`
  Sets radae mox state — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicScale()`** — L980 — `PORT void SetRadaeMicScale(double scale)`
  Sets radae mic scale — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeRxScale()`** — L987 — `PORT void SetRadaeRxScale(int rx, double scale)`
  Sets radae rx scale — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeRxDialScale()`** — L995 — `PORT void SetRadaeRxDialScale(int rx, double scale)`
  Sets radae rx dial scale — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeRxAFGain()`** — L1003 — `PORT void SetRadaeRxAFGain(int rx, double gain)`
  Sets radae rx afgain — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`GetRadaeRxAFGain()`** — L1011 — `PORT float GetRadaeRxAFGain(int rx)`
  Returns radae rx afgain — API getter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicRNNoiseEnabled()`** — L1020 — `PORT void SetRadaeMicRNNoiseEnabled(int e)`
  Pre-encoder mic-conditioning chain (TX-side, single).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicAGCEnabled()`** — L1021 — `PORT void SetRadaeMicAGCEnabled(int e)`
  Sets radae mic agcenabled — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicAGCTargetLufs()`** — L1022 — `PORT void SetRadaeMicAGCTargetLufs(double t)`
  Sets radae mic agctarget lufs — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicEQEnabled()`** — L1023 — `PORT void SetRadaeMicEQEnabled(int e)`
  Sets radae mic eqenabled — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicEQBass()`** — L1024 — `PORT void SetRadaeMicEQBass(double f, double g)`
  Sets radae mic eqbass — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicEQMid()`** — L1025 — `PORT void SetRadaeMicEQMid(double f, double g, double q)`
  Sets radae mic eqmid — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicEQTreble()`** — L1026 — `PORT void SetRadaeMicEQTreble(double f, double g)`
  Sets radae mic eqtreble — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeMicEQVol()`** — L1027 — `PORT void SetRadaeMicEQVol(double db)`
  Sets radae mic eqvol — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeBypassMicDsp()`** — L1032 — `PORT void SetRadaeBypassMicDsp(int enable)`
  Diagnostic bypass flags (TX-side, single).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeBypassEncoderCore()`** — L1033 — `PORT void SetRadaeBypassEncoderCore(int enable)`
  Sets radae bypass encoder core — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeBypassRmatch()`** — L1034 — `PORT void SetRadaeBypassRmatch(int enable)`
  Sets radae bypass rmatch — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeBypassEncoder()`** — L1035 — `PORT void SetRadaeBypassEncoder(int enable)`
  Sets radae bypass encoder — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeBypassAll()`** — L1036 — `PORT void SetRadaeBypassAll(int enable)`
  Sets radae bypass all — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetRadaeEooCallsign()`** — L1038 — `PORT void SetRadaeEooCallsign(const char* callsign)`
  Sets radae eoo callsign — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`xradae_rx()`** — L1055 — `void xradae_rx(int rx, double* rbuff_io)`
  Runs the `radae_rx` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`xradae_tx()`** — L1288 — `void xradae_tx(double* mic_io)`
  Runs the `radae_tx` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/radae.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
