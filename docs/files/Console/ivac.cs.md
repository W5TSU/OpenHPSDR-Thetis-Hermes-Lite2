# `Console/ivac.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** P/Invoke wrapper for ChannelMaster's VAC engine (`ivac.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ivac` (type, L47)

- **`.StartAudioIVAC()`** — L53 — `[DllImport("ChannelMaster.dll", EntryPoint = "StartAudioIVAC", CallingConvention = CallingConvention.Cdecl)] public static extern int StartAudioIVAC(int id)`
  Starts audio ivac.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopAudioIVAC()`** — L56 — `[DllImport("ChannelMaster.dll", EntryPoint = "StopAudioIVAC", CallingConvention = CallingConvention.Cdecl)] public static extern void StopAudioIVAC(int id)`
  Stops audio ivac.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACstereo()`** — L59 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACstereo", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACstereo(int id, int stereo)`
  Sets ivacstereo.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACrun()`** — L62 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACrun", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACrun(int id, int run)`
  Sets ivacrun.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACiqType()`** — L65 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACiqType", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACiqType(int id, int type)`
  Sets ivaciq type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACvacRate()`** — L77 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACvacRate", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACvacRate(int id, int rate)`
  Sets ivacvac rate.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACvacSize()`** — L92 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACvacSize", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACvacSize(int id, int size)`
  Sets ivacvac size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVAChostAPIindex()`** — L95 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVAChostAPIindex", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVAChostAPIindex(int id, int index)`
  Sets ivachost apiindex.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACinputDEVindex()`** — L98 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACinputDEVindex", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACinputDEVindex(int id, int index)`
  Sets ivacinput devindex.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACoutputDEVindex()`** — L101 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACoutputDEVindex", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACoutputDEVindex(int id, int inde`
  Sets ivacoutput devindex.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACnumChannels()`** — L104 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACnumChannels", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACnumChannels(int id, int n)`
  Sets ivacnum channels.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACInLatency()`** — L107 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACInLatency", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACInLatency(int id, double lat, int re`
  Sets ivacin latency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACOutLatency()`** — L110 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACOutLatency", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACOutLatency(int id, double lat, int `
  Sets ivacout latency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACPAInLatency()`** — L113 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACPAInLatency", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACPAInLatency(int id, double lat, in`
  Sets ivacpain latency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACPAOutLatency()`** — L116 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACPAOutLatency", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACPAOutLatency(int id, double lat, `
  Sets ivacpaout latency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACpreamp()`** — L119 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACpreamp", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACpreamp(int id, double preamp)`
  Sets ivacpreamp.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACbypass()`** — L122 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACbypass", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACbypass(int id, int enabled)`
  Sets ivacbypass.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACRBReset()`** — L125 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACRBReset", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACRBReset(int id, int enabled)`
  Sets ivacrbreset.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACvox()`** — L128 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACvox", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACvox(int id, int enabled)`
  Sets ivacvox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACrxscale()`** — L131 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACrxscale", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACrxscale(int id, double scale)`
  Sets ivacrxscale.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACcombine()`** — L134 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACcombine", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACcombine(int id, int combine)`
  Sets ivaccombine.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACmon()`** — L137 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACmon", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACmon(int id, int mon)`
  Sets ivacmon.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACmonVol()`** — L140 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACmonVol", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACmonVol(int id, double vol)`
  Sets ivacmon vol.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACmox()`** — L143 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACmox", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACmox(int id, int mox)`
  Sets ivacmox.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getIVACdiags()`** — L146 — `[DllImport("ChannelMaster.dll", EntryPoint = "getIVACdiags", CallingConvention = CallingConvention.Cdecl)] public static extern void getIVACdiags(int id, int type, int* underflows,`
  Returns ivacdiags.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.forceIVACvar()`** — L149 — `[DllImport("ChannelMaster.dll", EntryPoint = "forceIVACvar", CallingConvention = CallingConvention.Cdecl)] public static extern void forceIVACvar(int id, int type, bool force, doub`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.resetIVACdiags()`** — L152 — `[DllImport("ChannelMaster.dll", EntryPoint = "resetIVACdiags", CallingConvention = CallingConvention.Cdecl)] public static extern void resetIVACdiags(int id, int type)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACFeedbackGain()`** — L155 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACFeedbackGain", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACFeedbackGain(int id, int type, do`
  Sets ivacfeedback gain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACSlewTime()`** — L158 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACSlewTime", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACSlewTime(int id, int type, double sle`
  Sets ivacslew time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACPropRingMin()`** — L162 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACPropRingMin", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACPropRingMin(int id, int type, int `
  Sets ivacprop ring min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACPropRingMax()`** — L165 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACPropRingMax", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACPropRingMax(int id, int type, int `
  Sets ivacprop ring max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACFFRingMin()`** — L168 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACFFRingMin", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACFFRingMin(int id, int type, int ff_r`
  Sets ivacffring min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACFFRingMax()`** — L171 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACFFRingMax", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACFFRingMax(int id, int type, int ff_r`
  Sets ivacffring max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACFFAlpha()`** — L174 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACFFAlpha", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACFFAlpha(int id, int type, double ff_al`
  Sets ivacffalpha.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetIVACControlFlag()`** — L177 — `[DllImport("ChannelMaster.dll", EntryPoint = "GetIVACControlFlag", CallingConvention = CallingConvention.Cdecl)] public static extern void GetIVACControlFlag(int id, int type, int*`
  Returns ivaccontrol flag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACinitialVars()`** — L180 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACinitialVars", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACinitialVars(int id, double INvar, `
  Sets ivacinitial vars.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACswapIQout()`** — L183 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACswapIQout", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACswapIQout(int id, int swap)`
  Sets ivacswap iqout.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACExclusiveOut()`** — L186 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACExclusiveOut", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACExclusiveOut(int id, int exclusiv`
  Sets ivacexclusive out.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIVACExclusiveIn()`** — L189 — `[DllImport("ChannelMaster.dll", EntryPoint = "SetIVACExclusiveIn", CallingConvention = CallingConvention.Cdecl)] public static extern void SetIVACExclusiveIn(int id, int exclusive_`
  Sets ivacexclusive in.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ivac.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
