# `ChannelMaster/ivac.c`

**Functional area:** [8. ChannelMaster — audio and network routing](../../CODE_OUTLINE.md#8-channelmaster--audio-and-network-routing)

**Role:** Virtual Audio Cable engine — PortAudio streams with variable-ratio resampling between Thetis and other PC apps.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×7)
  - `ChannelMaster/pipe.c` (calls ×4)
  - `ChannelMaster/cmasio.c` (calls ×1)
- Uses (outgoing references to other files):
  - `wdsp/rmatch.c` (calls ×19)
  - `ChannelMaster/aamix.c` (calls ×12)
  - `wdsp/utilities.c` (calls ×2)
  - `ChannelMaster/cmcomm.h` (imports ×1)
  - `ChannelMaster/cmsetup.c` (calls ×1)
- Most-referenced symbols from other files: `combinebuff()` (×1), `SetIVACiqSizeAndRate()` (×1), `SetIVACmicRate()` (×1), `SetIVACmicSize()` (×1), `SetIVACaudioRate()` (×1), `SetIVACaudioSize()` (×1), `SetIVACtxmonRate()` (×1), `SetIVACtxmonSize()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`create_resamps()`** — L33 — `void create_resamps(IVAC a)`
  Constructor for the `resamps` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_ivac()` (same file), `SetIVACiqType()` (same file), `SetIVACvacRate()` (same file), `SetIVACmicRate()` (same file), `SetIVACaudioRate()` (same file), `SetIVACvacSize()` (same file) — and 6 more
- **`create_ivac()`** — L55 — `PORT void create_ivac( int id, int run, int iq_type, int stereo,`
  Constructor for the `ivac` block: allocates its state/buffers and computes initial coefficients.
  Called by: `create_pipe()` (`ChannelMaster/pipe.c`)
- **`destroy_resamps()`** — L111 — `void destroy_resamps(IVAC a)`
  Destroys the `resamps` block, freeing its allocated buffers.
  Called by: `destroy_ivac()` (same file), `SetIVACiqType()` (same file), `SetIVACvacRate()` (same file), `SetIVACmicRate()` (same file), `SetIVACaudioRate()` (same file), `SetIVACvacSize()` (same file) — and 6 more
- **`destroy_ivac()`** — L121 — `PORT void destroy_ivac(int id)`
  Destroys the `ivac` block, freeing its allocated buffers.
  Called by: `destroy_pipe()` (`ChannelMaster/pipe.c`)
- **`xvacIN()`** — L129 — `PORT void xvacIN(int id, double* in_tx, int bypass)`
  Runs the `vacIN` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`xvacOUT()`** — L145 — `PORT void xvacOUT(int id, int stream, double* data)`
  Runs the `vacOUT` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  Called by: `xpipe()` (`ChannelMaster/pipe.c`)
- **`xvac_out()`** — L165 — `void xvac_out(int id, int nsamples, double* buff)`
  Runs the `vac_out` block on one buffer of samples — the per-buffer processing entry called from the owning chain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`CallbackIVAC()`** — L196 — `int CallbackIVAC(const void* input, void* output, unsigned long frameCount, const PaStreamCallbackTimeInfo* timeInfo, PaStreamCallbackFlags statusFlags,`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`StartAudioIVAC()`** — L265 — `PORT int StartAudioIVAC(int id)`
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACRBReset()`** — L361 — `PORT void SetIVACRBReset(int id, int reset)`
  Sets ivacrbreset — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`StopAudioIVAC()`** — L367 — `PORT void StopAudioIVAC(int id)`
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACrun()`** — L373 — `PORT void SetIVACrun(int id, int run)`
  Sets ivacrun — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACiqType()`** — L379 — `PORT void SetIVACiqType(int id, int type)`
  Sets ivaciq type — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACstereo()`** — L392 — `PORT void SetIVACstereo(int id, int stereo)`
  Sets ivacstereo — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACvacRate()`** — L398 — `PORT void SetIVACvacRate(int id, int rate)`
  Sets ivacvac rate — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACmicRate()`** — L411 — `PORT void SetIVACmicRate(int id, int rate)`
  Sets ivacmic rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACaudioRate()`** — L424 — `PORT void SetIVACaudioRate(int id, int rate)`
  Sets ivacaudio rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACtxmonRate()`** — L442 — `void SetIVACtxmonRate(int id, int rate)`
  Sets ivactxmon rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACvacSize()`** — L458 — `PORT void SetIVACvacSize(int id, int size)`
  Sets ivacvac size — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACmicSize()`** — L471 — `PORT void SetIVACmicSize(int id, int size)`
  Sets ivacmic size — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACiqSizeAndRate()`** — L484 — `PORT void SetIVACiqSizeAndRate(int id, int size, int rate)`
  Sets ivaciq size and rate — API setter, typically called from the console via P/Invoke.
  Called by: `SetXcmInrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACaudioSize()`** — L501 — `PORT void SetIVACaudioSize(int id, int size)`
  Sets ivacaudio size — API setter, typically called from the console via P/Invoke.
  Called by: `SetRcvrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVACtxmonSize()`** — L516 — `void SetIVACtxmonSize(int id, int size)`
  Sets ivactxmon size — API setter, typically called from the console via P/Invoke.
  Called by: `SetXmtrChannelOutrate()` (`ChannelMaster/cmaster.c`)
- **`SetIVAChostAPIindex()`** — L522 — `PORT void SetIVAChostAPIindex(int id, int index)`
  Sets ivachost apiindex — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACinputDEVindex()`** — L528 — `PORT void SetIVACinputDEVindex(int id, int index)`
  Sets ivacinput devindex — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACoutputDEVindex()`** — L534 — `PORT void SetIVACoutputDEVindex(int id, int index)`
  Sets ivacoutput devindex — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACnumChannels()`** — L540 — `PORT void SetIVACnumChannels(int id, int n)`
  Sets ivacnum channels — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACInLatency()`** — L546 — `PORT void SetIVACInLatency(int id, double lat, int reset)`
  Sets ivacin latency — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACOutLatency()`** — L559 — `PORT void SetIVACOutLatency(int id, double lat, int reset)`
  Sets ivacout latency — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACPAInLatency()`** — L572 — `PORT void SetIVACPAInLatency(int id, double lat, int reset)`
  Sets ivacpain latency — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACPAOutLatency()`** — L582 — `PORT void SetIVACPAOutLatency(int id, double lat, int reset)`
  Sets ivacpaout latency — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACvox()`** — L592 — `PORT void SetIVACvox(int id, int vox)`
  Sets ivacvox — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACmox()`** — L598 — `PORT void SetIVACmox(int id, int mox)`
  Sets ivacmox — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACmon()`** — L630 — `PORT void SetIVACmon(int id, int mon)`
  Sets ivacmon — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACmonVol()`** — L662 — `PORT void SetIVACmonVol(int id, double vol)`
  Sets ivacmon vol — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACpreamp()`** — L669 — `PORT void SetIVACpreamp(int id, double preamp)`
  Sets ivacpreamp — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACrxscale()`** — L675 — `PORT void SetIVACrxscale(int id, double scale)`
  Sets ivacrxscale — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACbypass()`** — L682 — `PORT void SetIVACbypass(int id, int bypass)`
  Sets ivacbypass — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACcombine()`** — L688 — `PORT void SetIVACcombine(int id, int combine)`
  Sets ivaccombine — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`combinebuff()`** — L694 — `void combinebuff(int n, double* a, double* combined)`
  Called by: `xvacIN()` (same file), `asioIN()` (`ChannelMaster/cmasio.c`)
- **`scalebuff()`** — L701 — `void scalebuff(int size, double* in, double scale, double* out)`
  Called by: `xvacIN()` (same file)
- **`getIVACdiags()`** — L708 — `PORT void getIVACdiags (int id, int type, int* underflows, int* overflows, double* var, int* ringsize, int* nring)`
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`forceIVACvar()`** — L722 — `PORT void forceIVACvar (int id, int type, int force, double fvar)`
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`resetIVACdiags()`** — L742 — `PORT void resetIVACdiags(int id, int type)`
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACFeedbackGain()`** — L757 — `PORT void SetIVACFeedbackGain(int id, int type, double feedback_gain)`
  MW0LGE_21h
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACSlewTime()`** — L770 — `PORT void SetIVACSlewTime(int id, int type, double slew_time)`
  Sets ivacslew time — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACPropRingMin()`** — L785 — `PORT void SetIVACPropRingMin(int id, int type, int prop_min)`
  MW0LGE_21j
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACPropRingMax()`** — L798 — `PORT void SetIVACPropRingMax(int id, int type, int prop_max)`
  Sets ivacprop ring max — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACFFRingMin()`** — L811 — `PORT void SetIVACFFRingMin(int id, int type, int ff_ringmin)`
  Sets ivacffring min — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACFFRingMax()`** — L824 — `PORT void SetIVACFFRingMax(int id, int type, int ff_ringmax)`
  Sets ivacffring max — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACFFAlpha()`** — L837 — `PORT void SetIVACFFAlpha(int id, int type, double ff_alpha)`
  Sets ivacffalpha — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`GetIVACControlFlag()`** — L850 — `PORT void GetIVACControlFlag(int id, int type, int* control_flag)`
  Returns ivaccontrol flag — API getter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACinitialVars()`** — L863 — `PORT void SetIVACinitialVars(int id, double INvar, double OUTvar)`
  Sets ivacinitial vars — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACswapIQout()`** — L888 — `PORT void SetIVACswapIQout(int id, int swap)`
  Sets ivacswap iqout — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACExclusiveOut()`** — L895 — `PORT void SetIVACExclusiveOut(int id, int exclusive_out)`
  Sets ivacexclusive out — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.
- **`SetIVACExclusiveIn()`** — L902 — `PORT void SetIVACExclusiveIn(int id, int exclusive_in)`
  Sets ivacexclusive in — API setter, typically called from the console via P/Invoke.
  Called from C# via P/Invoke — declared/wrapped in `Console/ivac.cs`.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ivac.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
