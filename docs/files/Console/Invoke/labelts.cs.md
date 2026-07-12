# `Console/Invoke/labelts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/Andromeda/vfosettingspopup.cs` (references ×1)
  - `Console/CAT/CATTester.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmAbout.Designer.cs` (references ×1)
  - `Console/frmAddCustomRadio.Designer.cs` (references ×1)
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
  - …and 18 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `LabelTS` (type, L33)

- `.BringToFront()` — L519
- `.Contains()` — L526
- `.CreateControl()` — L538
- `.CreateGraphics()` — L545
- `.Dispose()` — L569
- `.DoDragDrop()` — L576
- `.Equals()` — L588
- `.FindForm()` — L600
- `.Focus()` — L612
- `.GetChildAtPoint()` — L624
- `.GetContainerControl()` — L636
- `.GetHashCode()` — L648
- `.GetLifetimeService()` — L660
- `.GetNextControl()` — L672
- `.GetType()` — L684
- `.Hide()` — L696
- `.InitializeLifetimeService()` — L703
- `.Invalidate()` — L715
- `.PerformLayout()` — L775
- `.PointToClient()` — L793
- `.PointToScreen()` — L805
- `.PreProcessMessage()` — L817
- `.RectangleToClient()` — L829
- `.RectangleToScreen()` — L841
- `.Refresh()` — L853
- `.ResetBackColor()` — L860
- `.ResetBindings()` — L867
- `.ResetCursor()` — L874
- `.ResetFont()` — L881
- `.ResetForeColor()` — L888
- `.ResetImeMode()` — L895
- `.ResetRightToLeft()` — L902
- `.ResetText()` — L909
- `.ResumeLayout()` — L916
- `.Scale()` — L934
- `.Select()` — L945
- `.SelectNextControl()` — L952
- `.SendToBack()` — L966
- `.SetBounds()` — L973
- `.Show()` — L996
- `.SuspendLayout()` — L1003
- `.ToString()` — L1010
- `.Update()` — L1022

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/labelts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
