# `Console/Invoke/buttonts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/BandButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/FilterButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeDependentSettingsForm.cs` (references ×1)
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/Andromeda/vfosettingspopup.cs` (references ×1)
  - `Console/ColorButton.cs` (inherits ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - …and 24 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `System.Windows.Forms` (namespace, L31)

_No extracted members._

#### `ButtonTS` (type, L33)

- `.BringToFront()` — L546
- `.Contains()` — L553
- `.CreateControl()` — L565
- `.CreateGraphics()` — L572
- `.Dispose()` — L596
- `.DoDragDrop()` — L603
- `.Equals()` — L615
- `.FindForm()` — L627
- `.Focus()` — L639
- `.GetChildAtPoint()` — L651
- `.GetContainerControl()` — L663
- `.GetHashCode()` — L675
- `.GetLifetimeService()` — L687
- `.GetNextControl()` — L699
- `.GetType()` — L711
- `.Hide()` — L723
- `.InitializeLifetimeService()` — L730
- `.Invalidate()` — L742
- `.NotifyDefault()` — L802
- `.PerformClick()` — L811
- `.PerformLayout()` — L818
- `.PointToClient()` — L836
- `.PointToScreen()` — L848
- `.PreProcessMessage()` — L860
- `.RectangleToClient()` — L872
- `.RectangleToScreen()` — L884
- `.Refresh()` — L896
- `.ResetBackColor()` — L903
- `.ResetBindings()` — L910
- `.ResetCursor()` — L917
- `.ResetFont()` — L924
- `.ResetForeColor()` — L931
- `.ResetImeMode()` — L938
- `.ResetRightToLeft()` — L945
- `.ResetText()` — L952
- `.ResumeLayout()` — L959
- `.Scale()` — L977
- `.Select()` — L988
- `.SelectNextControl()` — L995
- `.SendToBack()` — L1009
- `.SetBounds()` — L1016
- `.Show()` — L1039
- `.SuspendLayout()` — L1046
- `.ToString()` — L1053
- `.Update()` — L1065

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/buttonts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
