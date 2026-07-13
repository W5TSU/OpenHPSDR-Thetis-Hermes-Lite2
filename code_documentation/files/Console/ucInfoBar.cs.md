# `Console/ucInfoBar.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** The info bar (status/warning strip) and its popup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
  - `Console/Invoke/checkboxts.cs` (references ×2)
  - `Console/frmInfoBarPopup.Designer.cs` (references ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucInfoBar` (type, L52)

- **`.SendMessage()`** — L55 — `[DllImport("user32.dll")] private static extern int SendMessage(IntPtr hWnd, Int32 wMsg, bool wParam, Int32 lParam)`
  Sends message.
  Called by: `.lblSplitter_MouseMove()` (same file)
- **`.actionString()`** — L321 — `private string actionString(ActionTypes action)`
  Called by: `.UpdateButtonState()` (same file)
- **`.OnActionClicked_Button1()`** — L326 — `private void OnActionClicked_Button1(object sender, frmInfoBarPopup.PopupActionSelected e)`
  Handles/raises the action clicked button1 event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnActionClicked_Button2()`** — L346 — `private void OnActionClicked_Button2(object sender, frmInfoBarPopup.PopupActionSelected e)`
  Handles/raises the action clicked button2 event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.doAction()`** — L366 — `private void doAction(int button, ActionTypes action, bool bState, MouseButtons mouseButton)`
  Called by: `.OnActionClicked_Button1()` (same file), `.OnActionClicked_Button2()` (same file)
- **`.addPopup()`** — L394 — `private void addPopup(frmInfoBarPopup frm, ToolStripControlHost host, ToolStripDropDown dropDown)`
  Called by: `.LateInit()` (same file)
- **`.OnPopupClosed()`** — L417 — `private void OnPopupClosed(object sender, ToolStripDropDownClosedEventArgs e)`
  Handles/raises the popup closed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShutDown()`** — L420 — `public void ShutDown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onWarning()`** — L443 — `private void onWarning(object sender, System.Timers.ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onTick()`** — L453 — `private void onTick(object sender, System.Timers.ElapsedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LateInit()`** — L519 — `public void LateInit(Console c)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChangeHandler()`** — L554 — `private void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setPSboolsToFalse()`** — L562 — `private void setPSboolsToFalse()`
  Sets psbools to false.
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.chkButton1_CheckedChanged()`** — L568 — `private void chkButton1_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButton1` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButton2_CheckedChanged()`** — L578 — `private void chkButton2_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkButton2` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.UpdateButtonState()`** — L588 — `public void UpdateButtonState(ActionTypes action, bool bEnabled, bool bIncludePopup = true)`
  Updates button state.
  Called by: `.replaceMainButton()` (same file)
- **`.Left1()`** — L657 — `public void Left1(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Left2()`** — L673 — `public void Left2(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Left3()`** — L689 — `public void Left3(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Right1()`** — L705 — `public void Right1(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Right2()`** — L721 — `public void Right2(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Right3()`** — L737 — `public void Right3(int flipLayer, string value, int width = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetToolTipLeft()`** — L753 — `public void SetToolTipLeft(int flipLayer, int labelIndex, string text)`
  Sets tool tip left.
  Called by: `.updateLabels()` (same file)
- **`.SetToolTipRight()`** — L777 — `public void SetToolTipRight(int flipLayer, int labelIndex, string text)`
  Sets tool tip right.
  Called by: `.updateLabels()` (same file)
- **`.PSInfo()`** — L808 — `public void PSInfo(int level, bool bFeedbackLevelOk, bool bCorrectionsBeingApplied, bool bCalibrationAttemptsChanged, Color feedbackColour)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updatePSDisplay()`** — L839 — `private void updatePSDisplay()`
  Called by: `.OnMoxChangeHandler()` (same file), `.PSInfo()` (same file)
- **`.Warning()`** — L911 — `public void Warning(string msg, bool red_warning = false, int show_duration = 2000)`
  public void Warning(string msg, int nOverloadColourCount = -1, bool bExtendedShow = false)
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InfoBar_Resize()`** — L934 — `private void InfoBar_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `InfoBar` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.InfoBar_Click()`** — L946 — `private void InfoBar_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `InfoBar` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.flip()`** — L950 — `private void flip()`
  Called by: `.InfoBar_Click()` (same file)
- **`.updateLabels()`** — L958 — `private void updateLabels()`
  Called by: `.LateInit()` (same file), `.flip()` (same file)
- **`.chkButton1_MouseDown()`** — L989 — `private void chkButton1_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkButton1` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkButton2_MouseDown()`** — L1013 — `private void chkButton2_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `chkButton2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.IsRightButton()`** — L1037 — `private bool IsRightButton(MouseEventArgs e)`
  Called by: `.chkButton1_MouseDown()` (same file), `.chkButton2_MouseDown()` (same file)
- **`.lblFB_MouseDown()`** — L1042 — `private void lblFB_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblFB` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setToolTips()`** — L1081 — `private void setToolTips()`
  Sets tool tips.
  Called by: `.LateInit()` (same file)
- **`.replaceMainButton()`** — L1098 — `private void replaceMainButton(int button, ActionTypes action, bool bState, MouseButtons mouseButton)`
  Called by: `.OnActionClicked_Button1()` (same file), `.OnActionClicked_Button2()` (same file)
- **`.GetPopupButton()`** — L1115 — `public CheckBoxTS GetPopupButton(int infoBarButton, int index)`
  Returns popup button.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblSplitter_MouseDown()`** — L1130 — `private void lblSplitter_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblSplitter` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblSplitter_MouseEnter()`** — L1138 — `private void lblSplitter_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `lblSplitter` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblSplitter_MouseHover()`** — L1146 — `private void lblSplitter_MouseHover(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lblSplitter_MouseLeave()`** — L1151 — `private void lblSplitter_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `lblSplitter` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblSplitter_MouseMove()`** — L1157 — `private void lblSplitter_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblSplitter` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.repositionControls()`** — L1187 — `private void repositionControls()`
  Called by: `.Left1()` (same file), `.Left2()` (same file), `.Left3()` (same file), `.Right1()` (same file), `.Right2()` (same file), `.Right3()` (same file) — and 3 more
- **`.lblSplitter_MouseUp()`** — L1261 — `private void lblSplitter_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblSplitter` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `ActionTypes` (type, L59)

_No extracted members._

#### `InfoBarAction` (type, L73)

_No extracted members._

#### `ActionState` (type, L137)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucInfoBar.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
