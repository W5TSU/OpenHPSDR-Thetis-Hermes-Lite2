# `Console/ucLGPicker.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (references ×16, calls ×2)
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.InterpolateBetween()` (×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucLGPicker` (type, L54)

- **`.rebuildSortedColours()`** — L135 — `private void rebuildSortedColours()`
  Called by: `.LGPicker_MouseMove()` (same file), `.highlightGripper()` (same file), `.setColour()` (same file), `.RemoveSelectedGripper()` (same file), `.addGripper()` (same file), `.Clear()` (same file) — and 1 more
- **`.enableGripper()`** — L150 — `private void enableGripper(int index, bool bEnable)`
  Called by: `.RemoveSelectedGripper()` (same file), `.Clear()` (same file)
- **`.addColour()`** — L161 — `private void addColour(int index, float percpos, Color c, Dictionary<int,GradColours> lstColours, bool bEnable = true)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawTextCentre()`** — L180 — `private void drawTextCentre(Graphics g, string s, int x)`
  Called by: `.drawScales()` (same file)
- **`.drawScales()`** — L189 — `private void drawScales(Graphics g)`
  Called by: `.LGPicker_Paint()` (same file)
- **`.LGPicker_Paint()`** — L213 — `private void LGPicker_Paint(object sender, PaintEventArgs e)`
  WinForms event handler: runs when `LGPicker` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.findColourIndexGrip()`** — L279 — `private int findColourIndexGrip(int X, int Y)`
  Called by: `.LGPicker_MouseMove()` (same file), `.LGPicker_MouseDown()` (same file)
- **`.LGPicker_MouseMove()`** — L296 — `private void LGPicker_MouseMove(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `LGPicker` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.highlightGripper()`** — L364 — `private void highlightGripper(int index, bool bHightlight)`
  Called by: `.LGPicker_MouseDown()` (same file), `.HighlightFirstGripper()` (same file)
- **`.percFromPixels()`** — L378 — `private float percFromPixels(int X)`
  Called by: `.LGPicker_MouseMove()` (same file), `.LGPicker_MouseDown()` (same file)
- **`.LGPicker_MouseDown()`** — L387 — `private void LGPicker_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `LGPicker` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.LGPicker_MouseUp()`** — L447 — `private void LGPicker_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `LGPicker` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.indexOfClosestToLeft()`** — L469 — `private int indexOfClosestToLeft(float perc)`
  Called by: `.LGPicker_MouseMove()` (same file), `.LGPicker_MouseDown()` (same file), `.GetColourAtPercent()` (same file)
- **`.indexAtPerc()`** — L499 — `private int indexAtPerc(float perc)`
  Called by: `.GetColourAtPercent()` (same file), `.GetColourGradientDataForDBMRange()` (same file)
- **`.indexOfClosestToRight()`** — L513 — `private int indexOfClosestToRight(float perc)`
  Called by: `.LGPicker_MouseMove()` (same file), `.LGPicker_MouseDown()` (same file), `.GetColourAtPercent()` (same file), `.GetColourGradientDataForDBMRange()` (same file)
- **`.LGPicker_MouseLeave()`** — L544 — `private void LGPicker_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `LGPicker` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setColour()`** — L566 — `private void setColour(int index, Color c, bool bRefresh = false)`
  Sets colour.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveSelectedGripper()`** — L582 — `public void RemoveSelectedGripper(bool bRefresh = false)`
  Removes selected gripper.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addGripper()`** — L599 — `private int addGripper(float perc, Color colour, bool bRefresh = false)`
  Called by: `.LGPicker_MouseDown()` (same file)
- **`.findFreeDisabledGripper()`** — L625 — `private int findFreeDisabledGripper()`
  Called by: `.addGripper()` (same file)
- **`.HighlightFirstGripper()`** — L639 — `public void HighlightFirstGripper()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Clear()`** — L653 — `public void Clear()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnChanged()`** — L740 — `private void OnChanged(EventArgs e)`
  Handles/raises the changed event.
  Called by: `.LGPicker_MouseUp()` (same file), `.RemoveSelectedGripper()` (same file), `.addGripper()` (same file), `.Clear()` (same file), `.ApplyGlobalAlpha()` (same file)
- **`.OnGripperSelected()`** — L744 — `private void OnGripperSelected(Color c)`
  Handles/raises the gripper selected event.
  Called by: `.LGPicker_MouseDown()` (same file), `.LGPicker_MouseUp()` (same file), `.HighlightFirstGripper()` (same file)
- **`.OnGripperMouseEnter()`** — L748 — `private void OnGripperMouseEnter(int dbm, float percent)`
  Handles/raises the gripper mouse enter event.
  Called by: `.LGPicker_MouseMove()` (same file)
- **`.OnGripperMouseLeave()`** — L752 — `private void OnGripperMouseLeave(int dbm, float percent)`
  Handles/raises the gripper mouse leave event.
  Called by: `.LGPicker_MouseMove()` (same file), `.LGPicker_MouseLeave()` (same file)
- **`.OnGripperDBMChanged()`** — L756 — `private void OnGripperDBMChanged(int dbm, float percent)`
  Handles/raises the gripper dbmchanged event.
  Called by: `.LGPicker_MouseMove()` (same file)
- **`.LGPicker_EnabledChanged()`** — L760 — `private void LGPicker_EnabledChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetColourForDBM()`** — L764 — `public Color GetColourForDBM(float dbm)`
  Returns colour for dbm.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetColourAtPercent()`** — L770 — `public Color GetColourAtPercent(float perc)`
  Returns colour at percent.
  Called by: `.LGPicker_MouseDown()` (same file), `.GetColourForDBM()` (same file), `.GetColourGradientDataForDBMRange()` (same file)
- **`.GetPercForDBM()`** — L793 — `public float GetPercForDBM(float dbm)`
  Returns perc for dbm.
  Called by: `.GetColourForDBM()` (same file), `.GetColourGradientDataForDBMRange()` (same file)
- **`.addColourGradientData()`** — L804 — `private void addColourGradientData(List<ColourGradientData> lst, Color color, float perc)`
  Called by: `.GetColourGradientDataForDBMRange()` (same file)
- **`.GetColourGradientDataForDBMRange()`** — L813 — `public List<ColourGradientData> GetColourGradientDataForDBMRange(float low, float high)`
  making changes
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ApplyGlobalAlpha()`** — L891 — `public void ApplyGlobalAlpha(int A)`
  Applys global alpha.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `ColourGradientData` (type, L70)

_No extracted members._

#### `GradColours` (type, L75)

_No extracted members._

#### `ColourEventArgs` (type, L958)

_No extracted members._

#### `GripperEventArgs` (type, L963)

_No extracted members._

#### `ColorInterpolator` (type, L971)

- **`.InterpolateBetween()`** — L979 — `public static Color InterpolateBetween( Color endPoint1, Color endPoint2, double lambda)`
  Called by: `.GetColourAtPercent()` (same file), `.renderLed()` (`Console/MeterManager.cs`), `.OnRxTimerTick()` (`Console/setup.cs`), `.OnTxTimerTick()` (`Console/setup.cs`)
- **`.InterpolateComponent()`** — L999 — `static byte InterpolateComponent( Color endPoint1, Color endPoint2, double lambda, ComponentSelector selector)`
  Called by: `.InterpolateBetween()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucLGPicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
