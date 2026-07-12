# `wdsp/rmatch.c`

**Functional area:** [7. wdsp — the DSP engine](../../CODE_OUTLINE.md#7-wdsp--the-dsp-engine)

**Role:** Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks.

## How this file is used

- Used by (incoming references from other files):
  - `ChannelMaster/ivac.c` (calls ×19)
  - `ChannelMaster/cmasio.c` (calls ×9)
- Uses (outgoing references to other files):
  - `wdsp/utilities.c` (calls ×5)
  - `cmASIO/asiosdk_2.3.3_2019-06-14/common/combase.h` (calls ×3)
  - `wdsp/varsamp.c` (calls ×3)
  - `wdsp/comm.h` (imports ×1)
- Most-referenced symbols from other files: `xrmatchIN()` (×5), `xrmatchOUT()` (×4), `forceRMatchVar()` (×3), `create_rmatchV()` (×2), `destroy_rmatchV()` (×2), `getRMatchDiags()` (×2), `resetRMatchDiags()` (×2), `setRMatchFeedbackGain()` (×1)

## Outline

### Functions

- `create_mav()` — L29
- `destroy_mav()` — L43
- `flush_mav()` — L49
- `xmav()` — L56
- `create_aamav()` — L71
- `destroy_aamav()` — L86
- `flush_aamav()` — L92
- `xaamav()` — L101
- `calc_rmatch()` — L128
- `decalc_rmatch()` — L174
- `create_rmatch()` — L187
- `destroy_rmatch()` — L241
- `reset_rmatch()` — L247
- `control()` — L256
- `blend()` — L275
- `upslew()` — L285
- `xrmatchIN()` — L300
- `dslew()` — L364
- `xrmatchOUT()` — L427
- `getRMatchDiags()` — L469
- `resetRMatchDiags()` — L482
- `forceRMatchVar()` — L490
- `create_rmatchV()` — L500
- `destroy_rmatchV()` — L529
- `setRMatchInsize()` — L536
- `setRMatchOutsize()` — L548
- `setRMatchNomInrate()` — L560
- `setRMatchNomOutrate()` — L572
- `setRMatchRingsize()` — L584
- `setRMatchFeedbackGain()` — L596
- `setRMatchSlewTime()` — L606
- `setRMatchSlewTime1()` — L618
- `setRMatchPropRingMin()` — L641
- `setRMatchPropRingMax()` — L653
- `setRMatchFFRingMin()` — L665
- `setRMatchFFRingMax()` — L677
- `setRMatchFFAlpha()` — L689
- `getControlFlag()` — L699
- `create_rmatchLegacyV()` — L710

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/wdsp/rmatch.c`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
