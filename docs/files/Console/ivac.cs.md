# `Console/ivac.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** P/Invoke wrapper for ChannelMaster's VAC engine (`ivac.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `ivac` (type, L47)

- `.StartAudioIVAC()` — L53
- `.StopAudioIVAC()` — L56
- `.SetIVACstereo()` — L59
- `.SetIVACrun()` — L62
- `.SetIVACiqType()` — L65
- `.SetIVACvacRate()` — L77
- `.SetIVACvacSize()` — L92
- `.SetIVAChostAPIindex()` — L95
- `.SetIVACinputDEVindex()` — L98
- `.SetIVACoutputDEVindex()` — L101
- `.SetIVACnumChannels()` — L104
- `.SetIVACInLatency()` — L107
- `.SetIVACOutLatency()` — L110
- `.SetIVACPAInLatency()` — L113
- `.SetIVACPAOutLatency()` — L116
- `.SetIVACpreamp()` — L119
- `.SetIVACbypass()` — L122
- `.SetIVACRBReset()` — L125
- `.SetIVACvox()` — L128
- `.SetIVACrxscale()` — L131
- `.SetIVACcombine()` — L134
- `.SetIVACmon()` — L137
- `.SetIVACmonVol()` — L140
- `.SetIVACmox()` — L143
- `.getIVACdiags()` — L146
- `.forceIVACvar()` — L149
- `.resetIVACdiags()` — L152
- `.SetIVACFeedbackGain()` — L155
- `.SetIVACSlewTime()` — L158
- `.SetIVACPropRingMin()` — L162
- `.SetIVACPropRingMax()` — L165
- `.SetIVACFFRingMin()` — L168
- `.SetIVACFFRingMax()` — L171
- `.SetIVACFFAlpha()` — L174
- `.GetIVACControlFlag()` — L177
- `.SetIVACinitialVars()` — L180
- `.SetIVACswapIQout()` — L183
- `.SetIVACExclusiveOut()` — L186
- `.SetIVACExclusiveIn()` — L189

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ivac.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
