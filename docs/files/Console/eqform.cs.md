# `Console/eqform.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX/TX graphic and parametric equalizer forms (backed by wdsp `eq.c`).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files):
  - `Console/Invoke/panelts.cs` (references ×1, calls ×1)
  - `Console/common.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/Invoke/trackbarts.cs` (references ×1)
  - `Console/ucParametricEq.cs` (references ×1)
- Most-referenced symbols from other files: `.DSPOptionsChanged()` (×1), `.Show()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `EQForm` (type, L54)

- **`.Dispose()`** — L232 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L251 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HighlightTXProfileSaveItems()`** — L2056 — `public void HighlightTXProfileSaveItems(bool bHighlight)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EQForm_Closing()`** — L2376 — `private void EQForm_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `EQForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRXEQ_Scroll()`** — L2385 — `private void tbRXEQ_Scroll(object sender, System.EventArgs e)`
  WinForms event handler: runs when `tbRXEQ` is scrolled.
  Called by: `.btnRXEQReset_Click()` (same file), `.rad3Band_CheckedChanged()` (same file), `.rad10Band_CheckedChanged()` (same file), `.chkLegacyEQ_CheckedChanged()` (same file)
- **`.setDBtip()`** — L2407 — `private void setDBtip(object sender)`
  Sets dbtip.
  Called by: `.tbRXEQ_Scroll()` (same file), `.setTXEQProfile()` (same file)
- **`.picRXEQ_Paint()`** — L2504 — `private void picRXEQ_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRXEQ` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRXEQEnabled_CheckedChanged()`** — L2548 — `private void chkRXEQEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkRXEQEnabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.enableRxEq()`** — L2555 — `private void enableRxEq(bool enable)`
  Called by: `.chkRXEQEnabled_CheckedChanged()` (same file), `.chkParaEQ_enabled_CheckedChanged()` (same file), `.UpdateEQEnabled()` (same file)
- **`.chkTXEQEnabled_CheckedChanged()`** — L2568 — `private void chkTXEQEnabled_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkTXEQEnabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.enableTxEq()`** — L2574 — `private void enableTxEq(bool enable)`
  Called by: `.chkTXEQEnabled_CheckedChanged()` (same file), `.chkParaEQ_enabled_CheckedChanged()` (same file), `.UpdateEQEnabled()` (same file)
- **`.btnRXEQReset_Click()`** — L2580 — `private void btnRXEQReset_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `btnRXEQReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rad3Band_CheckedChanged()`** — L2622 — `private void rad3Band_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `rad3Band` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.rad10Band_CheckedChanged()`** — L2694 — `private void rad10Band_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `rad10Band` checked state changes.
  Called by: `.chkLegacyEQ_CheckedChanged()` (same file)
- **`.SetTXProfile()`** — L2765 — `public void SetTXProfile()`
  Sets txprofile.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTXEQProfile()`** — L2777 — `private void setTXEQProfile(object sender, EventArgs e)`
  Sets txeqprofile.
  Called by: `.SetTXProfile()` (same file), `.chkLegacyEQ_CheckedChanged()` (same file)
- **`.chkLegacyEQ_CheckedChanged()`** — L2880 — `private void chkLegacyEQ_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLegacyEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUseQFactors_CheckedChanged()`** — L2931 — `private void chkUseQFactors_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUseQFactors` checked state changes.
  Called by: `.chkLegacyEQ_CheckedChanged()` (same file)
- **`.setupWDSPdataFromParaEQ()`** — L2950 — `private void setupWDSPdataFromParaEQ(bool is_rx)`
  Called by: `.SetTXProfile()` (same file), `.radParaEQ_CheckedChanged()` (same file), `.ucParametricEq1_GlobalGainChanged()` (same file), `.ucParametricEq1_PointsChanged()` (same file), `.nudParaEQ_preamp_ValueChanged()` (same file)
- **`.dspUpdateTimerTick()`** — L2977 — `private void dspUpdateTimerTick(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.sendRXDspUpdate()`** — L2998 — `private void sendRXDspUpdate()`
  Called by: `.dspUpdateTimerTick()` (same file)
- **`.sendTXDspUpdate()`** — L3041 — `private void sendTXDspUpdate()`
  Called by: `.dspUpdateTimerTick()` (same file)
- **`.btnParaEQReset_Click()`** — L3083 — `private void btnParaEQReset_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnParaEQReset` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkParaEQ_enabled_CheckedChanged()`** — L3090 — `private void chkParaEQ_enabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkParaEQ_enabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radParaEQ_RX_CheckedChanged()`** — L3110 — `private void radParaEQ_RX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radParaEQ_RX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radParaEQ_CheckedChanged()`** — L3121 — `private void radParaEQ_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radParaEQ` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucParametricEq1_GlobalGainChanged()`** — L3167 — `private void ucParametricEq1_GlobalGainChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucParametricEq1_PointDataChanged()`** — L3189 — `private void ucParametricEq1_PointDataChanged(object sender, ucParametricEq.EqPointDataChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucParametricEq1_PointsChanged()`** — L3197 — `private void ucParametricEq1_PointsChanged(object sender, ucParametricEq.EqDraggingEventArgs e)`
  Called by: `.chkUseQFactors_CheckedChanged()` (same file)
- **`.setParaEQData()`** — L3322 — `private void setParaEQData()`
  Sets para eqdata.
  Called by: `.radParaEQ_CheckedChanged()` (same file), `.radParaEQ_RXTX_CheckedChanged()` (same file)
- **`.radParaEQ_RXTX_CheckedChanged()`** — L3380 — `private void radParaEQ_RXTX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radParaEQ_RXTX` checked state changes.
  Called by: `.chkLegacyEQ_CheckedChanged()` (same file)
- **`.nudParaEQ_selected_band_ValueChanged()`** — L3409 — `private void nudParaEQ_selected_band_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_selected_band` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateBandData()`** — L3418 — `private void updateBandData()`
  Called by: `.chkUseQFactors_CheckedChanged()` (same file), `.ucParametricEq1_PointDataChanged()` (same file), `.nudParaEQ_selected_band_ValueChanged()` (same file), `.ucParametricEq1_PointSelected()` (same file), `.ucParametricEq1_PointUnselected()` (same file)
- **`.ucParametricEq1_PointSelected()`** — L3447 — `private void ucParametricEq1_PointSelected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucParametricEq1_PointUnselected()`** — L3455 — `private void ucParametricEq1_PointUnselected(object sender, ucParametricEq.EqPointSelectionChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nudParaEQ_f_ValueChanged()`** — L3463 — `private void nudParaEQ_f_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_f` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudParaEQ_gain_ValueChanged()`** — L3476 — `private void nudParaEQ_gain_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_gain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudParaEQ_q_ValueChanged()`** — L3489 — `private void nudParaEQ_q_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_q` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPanaEQ_live_CheckedChanged()`** — L3502 — `private void chkPanaEQ_live_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkPanaEQ_live` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.nudParaEQ_preamp_ValueChanged()`** — L3508 — `private void nudParaEQ_preamp_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_preamp` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DSPOptionsChanged()`** — L3527 — `public void DSPOptionsChanged()`
  Called by: `.radParaEQ_RXTX_CheckedChanged()` (same file), `.chkPanaEQ_live_CheckedChanged()` (same file), `.UpdateDSP()` (`Console/console.cs`)
- **`.nudParaEQ_low_ValueChanged()`** — L3539 — `private void nudParaEQ_low_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_low` value changes.
  Called by: `.setupLowHigh()` (same file)
- **`.nudParaEQ_high_ValueChanged()`** — L3559 — `private void nudParaEQ_high_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `nudParaEQ_high` value changes.
  Called by: `.setupLowHigh()` (same file)
- **`.setupLowHigh()`** — L3578 — `private void setupLowHigh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateEQEnabled()`** — L3583 — `private void UpdateEQEnabled(bool is_rx)`
  Updates eqenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupTimer()`** — L3604 — `private void setupTimer(bool run)`
  Called by: `.EQForm_Closing()` (same file), `.chkLegacyEQ_CheckedChanged()` (same file), `.EQForm_VisibleChanged()` (same file)
- **`.EQForm_VisibleChanged()`** — L3618 — `private void EQForm_VisibleChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L3626 — `public new void Show()`
  Called by: `.equalizerToolStripMenuItem_Click()` (`Console/console.cs`)
- **`.chkLogScale_CheckedChanged()`** — L3640 — `private void chkLogScale_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLogScale` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `ParaEQState` (type, L2827)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/eqform.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
