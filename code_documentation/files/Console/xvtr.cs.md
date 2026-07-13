# `Console/xvtr.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Transverter band setup (frequency offsets, power limits per transverter band).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×40, references ×1)
  - `Console/Andromeda/Andromeda.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/comboboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.TranslateFreq()` (×15), `.XVTRFreq()` (×7), `.GetBegin()` (×4), `.GetEnd()` (×4), `.GetEnabled()` (×3), `.GetRXOnly()` (×3), `.GetDisablePA()` (×2), `.GetRXGain()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `XVTRForm` (type, L51)

- **`.Dispose()`** — L334 — `protected override void Dispose( bool disposing )`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L353 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.XVTRFreq()`** — L5335 — `public int XVTRFreq(double freq)`
  Returns an index that indicates which band the frequency is in.
  Called by: `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOABand_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`), `.UpdateVFOASub()` (`Console/console.cs`), `.ChangeNotchCentreFrequency()` (`Console/console.cs`), `.AddNotch()` (`Console/console.cs`) — and 1 more
- **`.TranslateFreq()`** — L5356 — `public double TranslateFreq(double freq)`
  Returns a translated frequency based on the xvtr data. Takes into account the LO Offset and correction.
  Called by: `.AntennaBandFromFreq()` (`Console/Andromeda/Andromeda.cs`), `.SetAriesAlexMode()` (`Console/Andromeda/Andromeda.cs`), `.RX1BandForVFOB()` (`Console/console.cs`), `.GetTransverterTranslatedRXBand()` (`Console/console.cs`), `.GetTransverterTranslatedTXBand()` (`Console/console.cs`), `.VFOASubUpdate()` (`Console/console.cs`) — and 9 more
- **`.SetupControlArrays()`** — L5374 — `private void SetupControlArrays()`
  Setups control arrays.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAntenna()`** — L5581 — `public int GetRXAntenna(int index)`
  Returns rxantenna.
  Called by: `.modifyXVTRantenna()` (`Console/console.cs`)
- **`.GetEnabled()`** — L5588 — `public bool GetEnabled(int index)`
  Returns enabled.
  Called by: `.EnableAllBands()` (`Console/console.cs`), `.CATGetXVTRBandNames()` (`Console/console.cs`), `.CATRX2BandUpDown()` (`Console/console.cs`)
- **`.GetBegin()`** — L5593 — `public float GetBegin(int index)`
  Returns begin.
  Called by: `.ChangeNotchCentreFrequency()` (`Console/console.cs`), `.AddNotch()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`), `.pnlDisplay_MouseMove()` (`Console/console.cs`)
- **`.GetEnd()`** — L5598 — `public float GetEnd(int index)`
  Returns end.
  Called by: `.ChangeNotchCentreFrequency()` (`Console/console.cs`), `.AddNotch()` (`Console/console.cs`), `.preBandSelect()` (`Console/console.cs`), `.pnlDisplay_MouseMove()` (`Console/console.cs`)
- **`.GetPower()`** — L5603 — `public int GetPower(int index)`
  Returns power.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPower()`** — L5608 — `public void SetPower(int index, int pwr)`
  Sets power.
  Called by: `.SetPowerUsingTargetDBM()` (`Console/console.cs`)
- **`.GetRXOnly()`** — L5613 — `public bool GetRXOnly(int index)`
  Returns rxonly.
  Called by: `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOABand_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.SetRXOnly()`** — L5618 — `public void SetRXOnly(int index, bool b)`
  Sets rxonly.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXGain()`** — L5623 — `public float GetRXGain(int index)`
  Returns rxgain.
  Called by: `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.GetXVTRRF()`** — L5635 — `public bool GetXVTRRF(int index)`
  Returns xvtrrf.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDisablePA()`** — L5640 — `public bool GetDisablePA(int index)`
  Returns disable pa.
  Called by: `.txtVFOAFreq_LostFocus()` (`Console/console.cs`), `.txtVFOBFreq_LostFocus()` (`Console/console.cs`)
- **`.XVTRForm_Closing()`** — L5650 — `private void XVTRForm_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `XVTRForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable0_CheckedChanged()`** — L5659 — `private void chkEnable0_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable0` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable1_CheckedChanged()`** — L5677 — `private void chkEnable1_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable2_CheckedChanged()`** — L5695 — `private void chkEnable2_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable3_CheckedChanged()`** — L5714 — `private void chkEnable3_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable3` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable4_CheckedChanged()`** — L5733 — `private void chkEnable4_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable4` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable5_CheckedChanged()`** — L5752 — `private void chkEnable5_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable5` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable6_CheckedChanged()`** — L5771 — `private void chkEnable6_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable6` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable7_CheckedChanged()`** — L5790 — `private void chkEnable7_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable7` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable8_CheckedChanged()`** — L5809 — `private void chkEnable8_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable8` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable9_CheckedChanged()`** — L5828 — `private void chkEnable9_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable9` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable10_CheckedChanged()`** — L5847 — `private void chkEnable10_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable10` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable11_CheckedChanged()`** — L5866 — `private void chkEnable11_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable11` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable12_CheckedChanged()`** — L5885 — `private void chkEnable12_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable12` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable13_CheckedChanged()`** — L5904 — `private void chkEnable13_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable13` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable14_CheckedChanged()`** — L5923 — `private void chkEnable14_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable14` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnable15_CheckedChanged()`** — L5940 — `private void chkEnable15_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable15` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtButtonText_TextChanged()`** — L5959 — `private void txtButtonText_TextChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `txtButtonText` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkUseXVTRTUNPWR_CheckedChanged()`** — L5967 — `private void chkUseXVTRTUNPWR_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkUseXVTRTUNPWR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkXVTRRF_CheckedChanged()`** — L5972 — `private void chkXVTRRF_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkXVTRRF` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udRXGain_ValueChanged()`** — L5999 — `private void udRXGain_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udRXGain` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlexTRRelay_CheckedChanged()`** — L6015 — `private void chkAlexTRRelay_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlexTRRelay` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.XVTRForm_Paint()`** — L6033 — `private void XVTRForm_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `XVTRForm` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udPower_ValueChanged()`** — L6038 — `private void udPower_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udPower` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboAnt0_SelectedIndexChanged()`** — L6044 — `private void comboAnt0_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboAnt0` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.XVTRForm_Activated()`** — L6056 — `private void XVTRForm_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/xvtr.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
