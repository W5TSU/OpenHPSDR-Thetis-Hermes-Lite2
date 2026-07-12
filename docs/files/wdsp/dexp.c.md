# `wdsp/dexp.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Downward expander / noise gate with VOX tie-in.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/cmaster.c` (calls ×5)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×5)
  - `wdsp/firmin.c` (calls ×4)
  - `wdsp/comm.h` (imports ×1)
  - `wdsp/fir.c` (calls ×1)
- Most-referenced symbols from other files: `create_dexp()` (×1), `destroy_dexp()` (×1), `xdexp()` (×1), `SetDEXPRate()` (×1), `SetDEXPSize()` (×1)

## Outline

### Functions

- `calc_delring()` — L31
- `decalc_delring()` — L45
- `flush_delring()` — L51
- `xdelring()` — L58
- `calc_slews()` — L91
- `calc_buffs()` — L111
- `decalc_buffs()` — L118
- `calc_dexp()` — L125
- `decalc_dexp()` — L149
- `calc_filter()` — L156
- `decalc_filter()` — L168
- `calc_antivox()` — L174
- `decalc_antivox()` — L181
- `create_dexp()` — L186
- `destroy_dexp()` — L229
- `flush_dexp()` — L241
- `xdexp()` — L265
- `SendCBPushDexpVox()` — L398
- `SetDEXPRun()` — L406
- `SetDEXPSize()` — L416
- `SetDEXPIOBuffers()` — L435
- `SetDEXPRate()` — L450
- `SetDEXPDetectorTau()` — L465
- `SetDEXPAttackTime()` — L478
- `SetDEXPReleaseTime()` — L491
- `SetDEXPHoldTime()` — L504
- `SetDEXPExpansionRatio()` — L517
- `SetDEXPHysteresisRatio()` — L530
- `SetDEXPAttackThreshold()` — L543
- `SetDEXPFilterTaps()` — L555
- `SetDEXPWindowType()` — L567
- `SetDEXPLowCut()` — L581
- `SetDEXPHighCut()` — L593
- `SetDEXPRunSideChannelFilter()` — L605
- `SetDEXPRunVox()` — L615
- `SetDEXPRunAudioDelay()` — L625
- `SetDEXPAudioDelay()` — L635
- `GetDEXPPeakSignal()` — L647
- `SetAntiVOXRun()` — L656
- `SetAntiVOXSize()` — L665
- `SetAntiVOXRate()` — L676
- `SetAntiVOXGain()` — L687
- `SetAntiVOXDetectorTau()` — L696
- `SendAntiVOXData()` — L707

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/dexp.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
