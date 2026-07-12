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

### Functions

- `create_resamps()` — L33
- `create_ivac()` — L55
- `destroy_resamps()` — L111
- `destroy_ivac()` — L121
- `xvacIN()` — L129
- `xvacOUT()` — L145
- `xvac_out()` — L165
- `CallbackIVAC()` — L196
- `StartAudioIVAC()` — L265
- `SetIVACRBReset()` — L361
- `StopAudioIVAC()` — L367
- `SetIVACrun()` — L373
- `SetIVACiqType()` — L379
- `SetIVACstereo()` — L392
- `SetIVACvacRate()` — L398
- `SetIVACmicRate()` — L411
- `SetIVACaudioRate()` — L424
- `SetIVACtxmonRate()` — L442
- `SetIVACvacSize()` — L458
- `SetIVACmicSize()` — L471
- `SetIVACiqSizeAndRate()` — L484
- `SetIVACaudioSize()` — L501
- `SetIVACtxmonSize()` — L516
- `SetIVAChostAPIindex()` — L522
- `SetIVACinputDEVindex()` — L528
- `SetIVACoutputDEVindex()` — L534
- `SetIVACnumChannels()` — L540
- `SetIVACInLatency()` — L546
- `SetIVACOutLatency()` — L559
- `SetIVACPAInLatency()` — L572
- `SetIVACPAOutLatency()` — L582
- `SetIVACvox()` — L592
- `SetIVACmox()` — L598
- `SetIVACmon()` — L630
- `SetIVACmonVol()` — L662
- `SetIVACpreamp()` — L669
- `SetIVACrxscale()` — L675
- `SetIVACbypass()` — L682
- `SetIVACcombine()` — L688
- `combinebuff()` — L694
- `scalebuff()` — L701
- `getIVACdiags()` — L708
- `forceIVACvar()` — L722
- `resetIVACdiags()` — L742
- `SetIVACFeedbackGain()` — L757
- `SetIVACSlewTime()` — L770
- `SetIVACPropRingMin()` — L785
- `SetIVACPropRingMax()` — L798
- `SetIVACFFRingMin()` — L811
- `SetIVACFFRingMax()` — L824
- `SetIVACFFAlpha()` — L837
- `GetIVACControlFlag()` — L850
- `SetIVACinitialVars()` — L863
- `SetIVACswapIQout()` — L888
- `SetIVACExclusiveOut()` — L895
- `SetIVACExclusiveIn()` — L902

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/ChannelMaster/ivac.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
