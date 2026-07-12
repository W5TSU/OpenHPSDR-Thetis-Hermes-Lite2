# `Console/ucMeter.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The meter user control and the floating multi-meter display window.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `Axis` (type, L49)

_No extracted members._

#### `ucMeter` (type, L60)

- `.HandleTouchDown()` — L170
- `.HandleTouchMove()` — L190
- `.HandleTouchUp()` — L209
- `.addDelegates()` — L262
- `.RemoveDelegates()` — L268
- `.OnMoxChangeHandler()` — L274
- `.pnlBar_MouseDown()` — L281
- `.Repaint()` — L295
- `.pnlBar_MouseLeave()` — L304
- `.pnlBar_MouseUp()` — L309
- `.pnlBar_MouseMove()` — L317
- `.showToolTip()` — L376
- `.hideToolTip()` — L384
- `.roundToNearestTen()` — L390
- `.pbGrab_MouseDown()` — L400
- `.pbGrab_MouseUp()` — L409
- `.forceResize()` — L416
- `.ChangeHeight()` — L451
- `.pbGrab_MouseMove()` — L489
- `.resize()` — L520
- `.storeLocation()` — L567
- `.RestoreLocation()` — L573
- `.setTopBarButtons()` — L605
- `.setTitle()` — L625
- `.getFirstLineOrWholeString()` — L631
- `.setupBorder()` — L640
- `.btnFloat_Click()` — L644
- `.pbGrab_MouseEnter()` — L649
- `.pbGrab_MouseLeave()` — L656
- `.mouseLeave()` — L664
- `.btnFloat_MouseLeave()` — L672
- `.lblRX_MouseDown()` — L687
- `.lblRX_MouseUp()` — L702
- `.lblRX_MouseMove()` — L710
- `.lblRX_MouseLeave()` — L764
- `.ucMeter_LocationChanged()` — L769
- `.ucMeter_SizeChanged()` — L777
- `.btnAxis_Click()` — L912
- `.setAxisButton()` — L936
- `.setPinOnTopButton()` — L969
- `.btnPin_Click()` — L974
- `.setTopMost()` — L980
- `.ToString()` — L1012
- `.TryParse()` — L1039
- `.btnAxis_MouseUp()` — L1161
- `.btnAxis_MouseLeave()` — L1166
- `.btnPin_MouseLeave()` — L1171
- `.btnSettings_Click()` — L1176
- `.btnSettings_MouseLeave()` — L1181
- `.uiComponentMouseLeave()` — L1186
- `.ucMeter_MouseLeave()` — L1192
- `.pnlContainer_MouseMove()` — L1198
- `.pnlContainer_MouseLeave()` — L1231

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucMeter.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
