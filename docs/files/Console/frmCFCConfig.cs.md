# `Console/frmCFCConfig.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Continuous Frequency Compressor (CFC) TX processing configuration (backed by wdsp `cfcomp.c`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmCFCConfig` (type, L53)

- **`.radCFC_bands_CheckedChanged()`** — L108 — `private void radCFC_bands_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radCFC_bands` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCFC_low_ValueChanged()`** — L120 — `private void udCFC_low_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCFC_low` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udCFC_high_ValueChanged()`** — L131 — `private void udCFC_high_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCFC_high` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_f_ValueChanged()`** — L142 — `private void nudCFC_f_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_f` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_precomp_ValueChanged()`** — L152 — `private void nudCFC_precomp_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_precomp` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_c_ValueChanged()`** — L159 — `private void nudCFC_c_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_c` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_posteqgain_ValueChanged()`** — L169 — `private void nudCFC_posteqgain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_posteqgain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_gain_ValueChanged()`** — L176 — `private void nudCFC_gain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_gain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_q_ValueChanged()`** — L186 — `private void nudCFC_q_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_q` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_cq_ValueChanged()`** — L196 — `private void nudCFC_cq_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_cq` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucCFC_comp_GlobalGainChanged()`** — L206 — `private void ucCFC_comp_GlobalGainChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_comp_PointDataChanged()`** — L218 — `private void ucCFC_comp_PointDataChanged(object sender, ucParametricEq.EqPointDataChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_comp_PointsChanged()`** — L234 — `private void ucCFC_comp_PointsChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_comp_PointSelected()`** — L240 — `private void ucCFC_comp_PointSelected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_comp_PointUnselected()`** — L248 — `private void ucCFC_comp_PointUnselected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_eq_GlobalGainChanged()`** — L257 — `private void ucCFC_eq_GlobalGainChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_eq_PointDataChanged()`** — L269 — `private void ucCFC_eq_PointDataChanged(object sender, ucParametricEq.EqPointDataChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_eq_PointsChanged()`** — L285 — `private void ucCFC_eq_PointsChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_eq_PointSelected()`** — L291 — `private void ucCFC_eq_PointSelected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucCFC_eq_PointUnselected()`** — L299 — `private void ucCFC_eq_PointUnselected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateSelected()`** — L316 — `private void updateSelected(ucParametricEq.EqPointSelectionChangedEventArgs e)`
  Called by: `.ucCFC_comp_PointSelected()` (same file), `.ucCFC_comp_PointUnselected()` (same file), `.ucCFC_eq_PointSelected()` (same file), `.ucCFC_eq_PointUnselected()` (same file)
- **`.setCFCProfile()`** — L333 — `private void setCFCProfile(int index = -1, bool just_text = false)`
  Sets cfcprofile.
  Called by: `.ucCFC_comp_GlobalGainChanged()` (same file), `.ucCFC_comp_PointDataChanged()` (same file), `.ucCFC_comp_PointsChanged()` (same file), `.ucCFC_eq_GlobalGainChanged()` (same file), `.ucCFC_eq_PointDataChanged()` (same file), `.ucCFC_eq_PointsChanged()` (same file) — and 2 more
- **`.timerTick()`** — L394 — `private void timerTick(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmCFCConfig_VisibleChanged()`** — L433 — `private void frmCFCConfig_VisibleChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTimer()`** — L437 — `private void setTimer()`
  Sets timer.
  Called by: `.frmCFCConfig_VisibleChanged()` (same file)
- **`.btnResetComp_Click()`** — L451 — `private void btnResetComp_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetComp` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnResetEQ_Click()`** — L458 — `private void btnResetEQ_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnResetEQ` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudCFC_selected_band_ValueChanged()`** — L465 — `private void nudCFC_selected_band_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudCFC_selected_band` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.frmCFCConfig_FormClosing()`** — L477 — `private void frmCFCConfig_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmCFCConfig` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCFC_UseQFactors_CheckedChanged()`** — L484 — `private void chkCFC_UseQFactors_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCFC_UseQFactors` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.HighlightTXProfileSaveItems()`** — L593 — `public void HighlightTXProfileSaveItems(bool bHighlight)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblOGGuide_LinkClicked()`** — L598 — `private void lblOGGuide_LinkClicked(object sender, LinkLabelLinkClickedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkLogScale_CheckedChanged()`** — L603 — `private void chkLogScale_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLogScale` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmCFCConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
