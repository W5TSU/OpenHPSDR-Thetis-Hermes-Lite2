# `Console/DiversityForm.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Two-receiver diversity reception control (phase/gain mixing of RX1/RX2).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/groupboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/radiobuttonts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `DiversityForm` (type, L59)

- **`.applyControlStyles()`** — L265 — `private void applyControlStyles(Control parent, bool dark_mode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L359 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L377 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerializeObjectToString()`** — L1458 — `public string SerializeObjectToString<T>(T obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DeserializeStringToObject()`** — L1470 — `public T DeserializeStringToObject<T>(string str)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initMemories()`** — L1481 — `private void initMemories()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.picRadar_Paint()`** — L1489 — `private void picRadar_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `picRadar` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getControlHandlePoint()`** — L1603 — `private Point getControlHandlePoint()`
  Returns control handle point.
  Called by: `.picRadar_Paint()` (same file)
- **`.PolarToXY()`** — L1629 — `private Point PolarToXY(double r, double angle)`
  Called by: `.picRadar_Paint()` (same file), `.getControlHandlePoint()` (same file), `.getMemoryIndexAtPoint()` (same file)
- **`.picRadar_MouseMove()`** — L1636 — `private void picRadar_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picRadar` receives mouse movement.
  Called by: `.picRadar_MouseDown()` (same file)
- **`.picRadar_MouseDown()`** — L1730 — `private void picRadar_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picRadar` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.picRadar_MouseUp()`** — L1757 — `private void picRadar_MouseUp(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `picRadar` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udR_ValueChanged()`** — L1777 — `private void udR_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udR` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udTheta_ValueChanged()`** — L1783 — `private void udTheta_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udTheta` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateDirection()`** — L1789 — `private void UpdateDirection()`
  Updates direction.
  Called by: `.UpdateDiversity()` (same file)
- **`.UpdateDiversity()`** — L1834 — `private unsafe void UpdateDiversity()`
  Updates diversity.
  Called by: `.udR_ValueChanged()` (same file), `.udTheta_ValueChanged()` (same file), `.radioButtonMerc1_CheckedChanged()` (same file), `.radioButtonMerc2_CheckedChanged()` (same file), `.udR2_ValueChanged()` (same file), `.udR1_ValueChanged()` (same file) — and 8 more
- **`.chkEnable_CheckedChanged()`** — L1898 — `private void chkEnable_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkEnable` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DiversityForm_Closing()`** — L1911 — `private void DiversityForm_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `DiversityForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShiftUp45_Click()`** — L1918 — `private void btnShiftUp45_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShiftUp45` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShift180_Click()`** — L1955 — `private void btnShift180_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShift180` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShiftDwn45_Click()`** — L1967 — `private void btnShiftDwn45_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShiftDwn45` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radioButtonMerc1_CheckedChanged()`** — L2004 — `private void radioButtonMerc1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radioButtonMerc1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radioButtonMerc2_CheckedChanged()`** — L2028 — `private void radioButtonMerc2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radioButtonMerc2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLockAngle_CheckedChanged()`** — L2048 — `private void chkLockAngle_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLockAngle` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkLockR_CheckedChanged()`** — L2054 — `private void chkLockR_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkLockR` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.groupBox_refMerc_Enter()`** — L2060 — `private void groupBox_refMerc_Enter(object sender, EventArgs e)`
  WinForms event handler: runs when `groupBox_refMerc` is entered.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.groupBox_udPhase_Enter()`** — L2064 — `private void groupBox_udPhase_Enter(object sender, EventArgs e)`
  WinForms event handler: runs when `groupBox_udPhase` is entered.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udR2_ValueChanged()`** — L2070 — `private void udR2_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udR2` value changes.
  Called by: `.radioButtonMerc1_CheckedChanged()` (same file), `.udGainMulti_ValueChanged()` (same file)
- **`.udR1_ValueChanged()`** — L2140 — `private void udR1_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udR1` value changes.
  Called by: `.radioButtonMerc2_CheckedChanged()` (same file), `.udGainMulti_ValueChanged()` (same file)
- **`.udCalib_ValueChanged()`** — L2341 — `private void udCalib_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udCalib` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAngle0_ValueChanged()`** — L2348 — `private void udAngle0_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udAngle0` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ConvertAngleToAngle0()`** — L2360 — `private double ConvertAngleToAngle0(double e)`
  Converts angle to angle0.
  Called by: `.picRadar_MouseMove()` (same file), `.stepAngle()` (same file)
- **`.ConvertAngle0ToAngle()`** — L2370 — `private double ConvertAngle0ToAngle(double e)`
  Converts angle0 to angle.
  Called by: `.udAngle0_ValueChanged()` (same file), `.udAntSpacing_ValueChanged_1()` (same file), `.chkCrossFire_CheckedChanged()` (same file)
- **`.panelDivControls_Enter()`** — L2385 — `private void panelDivControls_Enter(object sender, EventArgs e)`
  private void labelTS13_Click(object sender, EventArgs e) { }
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udAntSpacing_ValueChanged_1()`** — L2390 — `private void udAntSpacing_ValueChanged_1(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CalcVrms()`** — L2398 — `private double CalcVrms(double a, double b)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.chkCrossFire_CheckedChanged()`** — L2442 — `private void chkCrossFire_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCrossFire` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udFineNull_ValueChanged()`** — L2453 — `private void udFineNull_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udFineNull` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRxSource1_CheckedChanged()`** — L2534 — `private void radRxSource1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRxSource1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRxSource2_CheckedChanged()`** — L2548 — `private void radRxSource2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRxSource2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radRxSourceRx1Rx2_CheckedChanged()`** — L2562 — `private void radRxSourceRx1Rx2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radRxSourceRx1Rx2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkEnableDiversity_CheckedChanged()`** — L2576 — `private void chkEnableDiversity_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkEnableDiversity` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.FormEncoderEvent()`** — L2594 — `public void FormEncoderEvent()`
  method called by console encoder event. Provides option of auto-show and auto-hide if form was not shown, mark it as opened by an encoder event
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Callback()`** — L2613 — `private void Callback(object source, ElapsedEventArgs e)`
  callback function when 10 second timer expires; hide the form
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityForm_Load()`** — L2621 — `private void DiversityForm_Load(object sender, EventArgs e)`
  check if we want portrait or landscape format. If landscape change form size and panel positions
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.DiversityForm_Resize()`** — L2635 — `private void DiversityForm_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `DiversityForm` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udGainMulti_ValueChanged()`** — L2640 — `private void udGainMulti_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udGainMulti` value changes.
  Called by: `.btnMemory_Click()` (same file), `.recallMemory()` (same file)
- **`.chkAlwaysOnTop_CheckedChanged()`** — L2677 — `private void chkAlwaysOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlwaysOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkNoAttLink_CheckedChanged()`** — L2682 — `private void chkNoAttLink_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkNoAttLink` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkVFOSync_CheckedChanged()`** — L2687 — `private void chkVFOSync_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkVFOSync` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMemory_Click()`** — L2707 — `private void btnMemory_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnMemory` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.NormalizeAngle()`** — L2770 — `private double NormalizeAngle(double angle)`
  Called by: `.stepAngle()` (same file)
- **`.stepAngle()`** — L2785 — `private void stepAngle(double degrees)`
  Called by: `.btnShiftUp45_Click()` (same file), `.btnShift180_Click()` (same file), `.btnShiftDwn45_Click()` (same file), `.btnShiftUp10_Click()` (same file), `.btnShift90_Click()` (same file), `.btnShiftDown10_Click()` (same file)
- **`.btnShiftUp10_Click()`** — L2798 — `private void btnShiftUp10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShiftUp10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShift90_Click()`** — L2803 — `private void btnShift90_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShift90` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShiftDown10_Click()`** — L2808 — `private void btnShiftDown10_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShiftDown10` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udZoom_ValueChanged()`** — L2813 — `private void udZoom_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udZoom` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setNewNaming()`** — L2828 — `private void setNewNaming(bool new_naming)`
  Sets new naming.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getMemoryIndexAtPoint()`** — L2855 — `private int getMemoryIndexAtPoint(Point pt, int hit_radius_px)`
  memory click in radar
  Called by: `.updateHoverMemory()` (same file)
- **`.updateHoverMemory()`** — L2873 — `private void updateHoverMemory(Point pt)`
  Called by: `.picRadar_MouseMove()` (same file), `.picRadar_MouseDown()` (same file), `.picRadar_MouseUp()` (same file)
- **`.recallMemory()`** — L2885 — `private void recallMemory(int index)`
  Called by: `.picRadar_MouseUp()` (same file)
- **`.picRadar_MouseLeave()`** — L2929 — `private void picRadar_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `picRadar` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `memorySettings` (type, L1443)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/DiversityForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
