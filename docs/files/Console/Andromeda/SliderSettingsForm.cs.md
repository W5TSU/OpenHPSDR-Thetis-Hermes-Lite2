# `Console/Andromeda/SliderSettingsForm.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/trackbarts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SliderSettingsForm` (type, L35)

- **`.Dispose()`** — L114 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L131 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SliderSettingsForm_Closing()`** — L1240 — `private void SliderSettingsForm_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `SliderSettingsForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SliderSettingsForm_Activated()`** — L1251 — `private void SliderSettingsForm_Activated(object sender, EventArgs e)`
  copy initial settings from console controls when form shown
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tbRX1AF_Scroll()`** — L1309 — `private void tbRX1AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX2AF_Scroll()`** — L1314 — `private void tbRX2AF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2AF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbSubRXAF_Scroll()`** — L1319 — `private void tbSubRXAF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbSubRXAF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX1RF_Scroll()`** — L1324 — `private void tbRX1RF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1RF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX2RF_Scroll()`** — L1330 — `private void tbRX2RF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2RF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX1Sql_Scroll()`** — L1336 — `private void tbRX1Sql_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1Sql` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX2Sql_Scroll()`** — L1341 — `private void tbRX2Sql_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2Sql` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX1Pan_Scroll()`** — L1346 — `private void tbRX1Pan_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1Pan` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX2Pan_Scroll()`** — L1351 — `private void tbRX2Pan_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2Pan` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbSubRXPan_Scroll()`** — L1356 — `private void tbSubRXPan_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbSubRXPan` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbMasterAF_Scroll()`** — L1361 — `private void tbMasterAF_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbMasterAF` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbDrive_Scroll()`** — L1366 — `private void tbDrive_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbDrive` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSubRX_CheckedChanged()`** — L1371 — `private void chkSubRX_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSubRX` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnClose_Click()`** — L1379 — `private void btnClose_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClose` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SliderSettingsForm_FormClosing()`** — L1384 — `private void SliderSettingsForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `SliderSettingsForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX1Atten_Scroll()`** — L1389 — `private void tbRX1Atten_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX1Atten` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.tbRX2Atten_Scroll()`** — L1394 — `private void tbRX2Atten_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `tbRX2Atten` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX1Mute_CheckedChanged()`** — L1399 — `private void chkRX1Mute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX1Mute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRX2Mute_CheckedChanged()`** — L1404 — `private void chkRX2Mute_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRX2Mute` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX1VAC_CheckedChanged()`** — L1409 — `private void ChkRX1VAC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX1VAC` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ChkRX2VAC_CheckedChanged()`** — L1415 — `private void ChkRX2VAC_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkRX2VAC` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TbMicGain_Scroll()`** — L1421 — `private void TbMicGain_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbMicGain` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TbRX1VACRX_Scroll()`** — L1426 — `private void TbRX1VACRX_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbRX1VACRX` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TbRX1VACTX_Scroll()`** — L1432 — `private void TbRX1VACTX_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbRX1VACTX` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TbRX2VACRX_Scroll()`** — L1438 — `private void TbRX2VACRX_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbRX2VACRX` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TbRX2VACTX_Scroll()`** — L1444 — `private void TbRX2VACTX_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TbRX2VACTX` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.FormEncoderEvent()`** — L1453 — `public void FormEncoderEvent()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Callback()`** — L1472 — `private void Callback(object source, ElapsedEventArgs e)`
  callback function when 10 second timer expires; hide the form
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX1Sql_CheckStateChanged()`** — L1493 — `private void chkRX1Sql_CheckStateChanged(object sender, EventArgs e)`
  private void chkRX2Sql_CheckedChanged(object sender, EventArgs e) { if (chkRX2Sql.Checked == true) console.CATSquelch2 = 1; else console.CATSquelch2 = 0; }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkRX2Sql_CheckStateChanged()`** — L1509 — `private void chkRX2Sql_CheckStateChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateSQLButtons()`** — L1525 — `private void updateSQLButtons(int rx)`
  Called by: `.chkRX1Sql_CheckStateChanged()` (same file), `.chkRX2Sql_CheckStateChanged()` (same file)
- **`.getSQLinfoOnFormActivate()`** — L1542 — `private void getSQLinfoOnFormActivate()`
  Returns sqlinfo on form activate.
  Called by: `.SliderSettingsForm_Activated()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/SliderSettingsForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
