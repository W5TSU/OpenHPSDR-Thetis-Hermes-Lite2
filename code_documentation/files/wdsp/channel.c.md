# `wdsp/channel.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Channel object lifecycle (create/destroy/run) and DLL entry points.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×8)
  - `wdsp/RXA.c` (calls ×1)
  - `wdsp/TXA.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/main.c` (calls ×10)
  - `wdsp/iobuffs.c` (calls ×9)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `OpenChannel()` (×2), `CloseChannel()` (×2), `SetOutputSamplerate()` (×2), `SetChannelState()` (×2), `SetInputBuffsize()` (×1), `SetInputSamplerate()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`start_thread()`** — L31 — `void start_thread (int channel)`
  Called by: `post_main_build()` (same file)
- **`pre_main_build()`** — L37 — `void pre_main_build (int channel)`
  Called by: `build_channel()` (same file), `SetInputBuffsize()` (same file), `SetDSPBuffsize()` (same file), `SetInputSamplerate()` (same file), `SetDSPSamplerate()` (same file), `SetOutputSamplerate()` (same file) — and 1 more
- **`post_main_build()`** — L60 — `void post_main_build (int channel)`
  Called by: `build_channel()` (same file), `SetInputBuffsize()` (same file), `SetDSPBuffsize()` (same file), `SetInputSamplerate()` (same file), `SetDSPSamplerate()` (same file), `SetOutputSamplerate()` (same file) — and 1 more
- **`build_channel()`** — L68 — `void build_channel (int channel)`
  Called by: `OpenChannel()` (same file), `SetType()` (same file)
- **`OpenChannel()`** — L75 — `PORT void OpenChannel (int channel, int in_size, int dsp_size, int input_samplerate, int dsp_rate, int output_samplerate, int type, int state, double tdelayup, double tslewup, doub`
  Called by: `create_rcvr()` (`ChannelMaster/cmaster.c`), `create_xmtr()` (`ChannelMaster/cmaster.c`)
- **`pre_main_destroy()`** — L103 — `void pre_main_destroy (int channel)`
  Called by: `CloseChannel()` (same file), `SetInputBuffsize()` (same file), `SetDSPBuffsize()` (same file), `SetInputSamplerate()` (same file), `SetDSPSamplerate()` (same file), `SetOutputSamplerate()` (same file) — and 1 more
- **`post_main_destroy()`** — L113 — `void post_main_destroy (int channel)`
  Called by: `CloseChannel()` (same file), `SetInputBuffsize()` (same file), `SetDSPBuffsize()` (same file), `SetInputSamplerate()` (same file), `SetDSPSamplerate()` (same file), `SetOutputSamplerate()` (same file) — and 1 more
- **`CloseChannel()`** — L120 — `PORT void CloseChannel (int channel)`
  Called by: `SetType()` (same file), `destroy_rcvr()` (`ChannelMaster/cmaster.c`), `destroy_xmtr()` (`ChannelMaster/cmaster.c`)
- **`flushChannel()`** — L128 — `void flushChannel (void* p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetType()`** — L156 — `PORT void SetType (int channel, int type)`
  Sets type — API setter, typically called from the console via P/Invoke.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`SetInputBuffsize()`** — L167 — `PORT void SetInputBuffsize (int channel, int in_size)`
  Sets input buffsize — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetDSPBuffsize()`** — L180 — `PORT void SetDSPBuffsize (int channel, int dsp_size)`
  Sets dspbuffsize — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetInputSamplerate()`** — L196 — `PORT void SetInputSamplerate (int channel, int in_rate)`
  Sets input samplerate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetDSPSamplerate()`** — L210 — `PORT void SetDSPSamplerate (int channel, int dsp_rate)`
  Sets dspsamplerate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetOutputSamplerate()`** — L226 — `PORT void SetOutputSamplerate (int channel, int out_rate)`
  Sets output samplerate — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`), `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetAllRates()`** — L240 — `PORT void SetAllRates (int channel, int in_rate, int dsp_rate, int out_rate)`
  Sets all rates — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetChannelState()`** — L258 — `PORT int SetChannelState (int channel, int state, int dmode)`
  Sets channel state — API setter, typically called from the console via P/Invoke.
  Called by: `SetDSPBuffsize()` (same file), `SetDSPSamplerate()` (same file), `RXASetNC()` (`wdsp/RXA.c`), `TXASetNC()` (`wdsp/TXA.c`)
- **`SetChannelTDelayUp()`** — L299 — `PORT void SetChannelTDelayUp (int channel, double time)`
  Sets channel tdelay up — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetChannelTSlewUp()`** — L311 — `PORT void SetChannelTSlewUp (int channel, double time)`
  Sets channel tslew up — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetChannelTDelayDown()`** — L323 — `PORT void SetChannelTDelayDown (int channel, double time)`
  Sets channel tdelay down — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.
- **`SetChannelTSlewDown()`** — L335 — `PORT void SetChannelTSlewDown (int channel, double time)`
  Sets channel tslew down — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/dsp.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/channel.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
