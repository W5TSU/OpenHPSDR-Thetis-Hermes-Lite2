# `Console/ucMeter.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The meter user control and the floating multi-meter display window.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Axis` (type, L49)

_No extracted members._

#### `ucMeter` (type, L60)

- **`.HandleTouchDown()`** — L170 — `private void HandleTouchDown(int x, int y)`
  Handles touch down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleTouchMove()`** — L190 — `private void HandleTouchMove(int x, int y)`
  Handles touch move.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleTouchUp()`** — L209 — `private void HandleTouchUp(int x, int y)`
  Handles touch up.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addDelegates()`** — L262 — `private void addDelegates()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveDelegates()`** — L268 — `public void RemoveDelegates()`
  Removes delegates.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChangeHandler()`** — L274 — `private void OnMoxChangeHandler(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pnlBar_MouseDown()`** — L281 — `private void pnlBar_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlBar` receives a mouse-down.
  Called by: `.HandleTouchDown()` (same file)
- **`.Repaint()`** — L295 — `public void Repaint()`
  Called by: `.pnlBar_MouseMove()` (same file), `.resize()` (same file), `.RestoreLocation()` (same file), `.lblRX_MouseMove()` (same file)
- **`.pnlBar_MouseLeave()`** — L304 — `private void pnlBar_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlBar` is left by the mouse.
  Called by: `.HandleTouchUp()` (same file)
- **`.pnlBar_MouseUp()`** — L309 — `private void pnlBar_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlBar` receives a mouse-up.
  Called by: `.HandleTouchUp()` (same file)
- **`.pnlBar_MouseMove()`** — L317 — `private void pnlBar_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlBar` receives mouse movement.
  Called by: `.HandleTouchMove()` (same file)
- **`.showToolTip()`** — L376 — `private void showToolTip(string msg, Control window, bool is_resize = false)`
  Called by: `.pnlBar_MouseMove()` (same file), `.pbGrab_MouseMove()` (same file), `.lblRX_MouseMove()` (same file)
- **`.hideToolTip()`** — L384 — `private void hideToolTip()`
  Called by: `.pnlBar_MouseUp()` (same file), `.pbGrab_MouseUp()` (same file), `.lblRX_MouseUp()` (same file)
- **`.roundToNearestTen()`** — L390 — `private int roundToNearestTen(int number)`
  Called by: `.pnlBar_MouseMove()` (same file), `.pbGrab_MouseMove()` (same file), `.lblRX_MouseMove()` (same file)
- **`.pbGrab_MouseDown()`** — L400 — `private void pbGrab_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pbGrab` receives a mouse-down.
  Called by: `.HandleTouchDown()` (same file)
- **`.pbGrab_MouseUp()`** — L409 — `private void pbGrab_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pbGrab` receives a mouse-up.
  Called by: `.HandleTouchUp()` (same file)
- **`.forceResize()`** — L416 — `private void forceResize(bool shrink = false)`
  Called by: `.pbGrab_MouseUp()` (same file), `.ChangeHeight()` (same file)
- **`.ChangeHeight()`** — L451 — `public void ChangeHeight(int height)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.pbGrab_MouseMove()`** — L489 — `private void pbGrab_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pbGrab` receives mouse movement.
  Called by: `.HandleTouchMove()` (same file)
- **`.resize()`** — L520 — `private void resize(int x, int y, bool shrink = false)`
  Called by: `.forceResize()` (same file), `.pbGrab_MouseMove()` (same file)
- **`.storeLocation()`** — L567 — `private void storeLocation()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RestoreLocation()`** — L573 — `[Browsable(false), EditorBrowsable(EditorBrowsableState.Never)] public void RestoreLocation()`
  Restores location.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTopBarButtons()`** — L605 — `private void setTopBarButtons()`
  Sets top bar buttons.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTitle()`** — L625 — `private void setTitle()`
  Sets title.
  Called by: `.OnMoxChangeHandler()` (same file)
- **`.getFirstLineOrWholeString()`** — L631 — `private string getFirstLineOrWholeString(string input)`
  Returns first line or whole string.
  Called by: `.setTitle()` (same file)
- **`.setupBorder()`** — L640 — `private void setupBorder()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnFloat_Click()`** — L644 — `private void btnFloat_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFloat` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pbGrab_MouseEnter()`** — L649 — `private void pbGrab_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `pbGrab` is entered by the mouse.
  Called by: `.HandleTouchDown()` (same file)
- **`.pbGrab_MouseLeave()`** — L656 — `private void pbGrab_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pbGrab` is left by the mouse.
  Called by: `.HandleTouchUp()` (same file)
- **`.mouseLeave()`** — L664 — `private void mouseLeave()`
  Called by: `.pbGrab_MouseLeave()` (same file), `.uiComponentMouseLeave()` (same file), `.ucMeter_MouseLeave()` (same file), `.pnlContainer_MouseLeave()` (same file)
- **`.btnFloat_MouseLeave()`** — L672 — `private void btnFloat_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `btnFloat` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblRX_MouseDown()`** — L687 — `private void lblRX_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblRX` receives a mouse-down.
  Called by: `.HandleTouchDown()` (same file)
- **`.lblRX_MouseUp()`** — L702 — `private void lblRX_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblRX` receives a mouse-up.
  Called by: `.HandleTouchUp()` (same file)
- **`.lblRX_MouseMove()`** — L710 — `private void lblRX_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `lblRX` receives mouse movement.
  Called by: `.HandleTouchMove()` (same file)
- **`.lblRX_MouseLeave()`** — L764 — `private void lblRX_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `lblRX` is left by the mouse.
  Called by: `.HandleTouchUp()` (same file)
- **`.ucMeter_LocationChanged()`** — L769 — `private void ucMeter_LocationChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucMeter_SizeChanged()`** — L777 — `private void ucMeter_SizeChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnAxis_Click()`** — L912 — `private void btnAxis_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAxis` is clicked.
  Called by: `.btnAxis_MouseUp()` (same file)
- **`.setAxisButton()`** — L936 — `private void setAxisButton()`
  Sets axis button.
  Called by: `.setTopBarButtons()` (same file), `.btnAxis_Click()` (same file)
- **`.setPinOnTopButton()`** — L969 — `private void setPinOnTopButton()`
  Sets pin on top button.
  Called by: `.setTopBarButtons()` (same file), `.btnPin_Click()` (same file)
- **`.btnPin_Click()`** — L974 — `private void btnPin_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPin` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setTopMost()`** — L980 — `private void setTopMost()`
  Sets top most.
  Called by: `.btnPin_Click()` (same file)
- **`.ToString()`** — L1012 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse()`** — L1039 — `public bool TryParse(string str)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnAxis_MouseUp()`** — L1161 — `private void btnAxis_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `btnAxis` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAxis_MouseLeave()`** — L1166 — `private void btnAxis_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAxis` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnPin_MouseLeave()`** — L1171 — `private void btnPin_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPin` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSettings_Click()`** — L1176 — `private void btnSettings_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSettings` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSettings_MouseLeave()`** — L1181 — `private void btnSettings_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSettings` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.uiComponentMouseLeave()`** — L1186 — `private void uiComponentMouseLeave()`
  Called by: `.pnlBar_MouseLeave()` (same file), `.btnFloat_MouseLeave()` (same file), `.lblRX_MouseLeave()` (same file), `.btnAxis_MouseLeave()` (same file), `.btnPin_MouseLeave()` (same file), `.btnSettings_MouseLeave()` (same file)
- **`.ucMeter_MouseLeave()`** — L1192 — `private void ucMeter_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `ucMeter` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.pnlContainer_MouseMove()`** — L1198 — `private void pnlContainer_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `pnlContainer` receives mouse movement.
  Called by: `.HandleTouchDown()` (same file), `.HandleTouchMove()` (same file)
- **`.pnlContainer_MouseLeave()`** — L1231 — `private void pnlContainer_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `pnlContainer` is left by the mouse.
  Called by: `.HandleTouchUp()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucMeter.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
