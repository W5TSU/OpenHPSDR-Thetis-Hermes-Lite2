# `Console/Invoke/radiobuttonts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×2)
  - `Console/Andromeda/BandButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/FilterButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeButtonsPopup.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/FilterForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmBandwidth.Designer.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
  - `Console/frmMacroButtonConfig.Designer.cs` (references ×1)
  - …and 5 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `RadioButtonTS` (type, L33)

- `.BringToFront()` — L552
- `.Contains()` — L559
- `.CreateControl()` — L571
- `.CreateGraphics()` — L578
- `.Dispose()` — L602
- `.DoDragDrop()` — L609
- `.Equals()` — L621
- `.FindForm()` — L633
- `.Focus()` — L645
- `.GetChildAtPoint()` — L657
- `.GetContainerControl()` — L669
- `.GetHashCode()` — L681
- `.GetLifetimeService()` — L693
- `.GetNextControl()` — L705
- `.GetType()` — L717
- `.Hide()` — L729
- `.InitializeLifetimeService()` — L736
- `.Invalidate()` — L748
- `.PerformClick()` — L810
- `.PerformLayout()` — L817
- `.PointToClient()` — L835
- `.PointToScreen()` — L847
- `.PreProcessMessage()` — L859
- `.RectangleToClient()` — L871
- `.RectangleToScreen()` — L883
- `.Refresh()` — L895
- `.ResetBackColor()` — L902
- `.ResetBindings()` — L909
- `.ResetCursor()` — L916
- `.ResetFont()` — L923
- `.ResetForeColor()` — L930
- `.ResetImeMode()` — L937
- `.ResetRightToLeft()` — L944
- `.ResetText()` — L951
- `.ResumeLayout()` — L958
- `.Scale()` — L976
- `.Select()` — L987
- `.SelectNextControl()` — L994
- `.SendToBack()` — L1008
- `.SetBounds()` — L1015
- `.Show()` — L1038
- `.SuspendLayout()` — L1045
- `.ToString()` — L1052
- `.Update()` — L1064

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/radiobuttonts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
