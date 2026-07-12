# `Console/frmCFCConfig.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Continuous Frequency Compressor (CFC) TX processing configuration (backed by wdsp `cfcomp.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)

## Outline

### Types

#### `frmCFCConfig` (type, L53)

- `.radCFC_bands_CheckedChanged()` — L108
- `.udCFC_low_ValueChanged()` — L120
- `.udCFC_high_ValueChanged()` — L131
- `.nudCFC_f_ValueChanged()` — L142
- `.nudCFC_precomp_ValueChanged()` — L152
- `.nudCFC_c_ValueChanged()` — L159
- `.nudCFC_posteqgain_ValueChanged()` — L169
- `.nudCFC_gain_ValueChanged()` — L176
- `.nudCFC_q_ValueChanged()` — L186
- `.nudCFC_cq_ValueChanged()` — L196
- `.ucCFC_comp_GlobalGainChanged()` — L206
- `.ucCFC_comp_PointDataChanged()` — L218
- `.ucCFC_comp_PointsChanged()` — L234
- `.ucCFC_comp_PointSelected()` — L240
- `.ucCFC_comp_PointUnselected()` — L248
- `.ucCFC_eq_GlobalGainChanged()` — L257
- `.ucCFC_eq_PointDataChanged()` — L269
- `.ucCFC_eq_PointsChanged()` — L285
- `.ucCFC_eq_PointSelected()` — L291
- `.ucCFC_eq_PointUnselected()` — L299
- `.updateSelected()` — L316
- `.setCFCProfile()` — L333
- `.timerTick()` — L394
- `.frmCFCConfig_VisibleChanged()` — L433
- `.setTimer()` — L437
- `.btnResetComp_Click()` — L451
- `.btnResetEQ_Click()` — L458
- `.nudCFC_selected_band_ValueChanged()` — L465
- `.frmCFCConfig_FormClosing()` — L477
- `.chkCFC_UseQFactors_CheckedChanged()` — L484
- `.HighlightTXProfileSaveItems()` — L593
- `.lblOGGuide_LinkClicked()` — L598
- `.chkLogScale_CheckedChanged()` — L603

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmCFCConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
