# `Console/Invoke/groupboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/BandButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/FilterButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmFilterManager.Designer.cs` (references ×1)
  - `Console/frmMacroButtonConfig.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - `Console/RAForm.Designer.cs` (references ×1)
  - …and 2 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `GroupBoxTS` (type, L33)

- `.BringToFront()` — L453
- `.Contains()` — L460
- `.CreateControl()` — L472
- `.CreateGraphics()` — L479
- `.Dispose()` — L503
- `.DoDragDrop()` — L510
- `.Equals()` — L522
- `.FindForm()` — L534
- `.Focus()` — L546
- `.GetChildAtPoint()` — L558
- `.GetContainerControl()` — L570
- `.GetHashCode()` — L582
- `.GetLifetimeService()` — L594
- `.GetNextControl()` — L606
- `.GetType()` — L618
- `.Hide()` — L630
- `.InitializeLifetimeService()` — L637
- `.Invalidate()` — L649
- `.PerformLayout()` — L711
- `.PointToClient()` — L729
- `.PointToScreen()` — L741
- `.PreProcessMessage()` — L753
- `.RectangleToClient()` — L765
- `.RectangleToScreen()` — L777
- `.Refresh()` — L789
- `.ResetBackColor()` — L796
- `.ResetBindings()` — L803
- `.ResetCursor()` — L810
- `.ResetFont()` — L817
- `.ResetForeColor()` — L824
- `.ResetImeMode()` — L831
- `.ResetRightToLeft()` — L838
- `.ResetText()` — L845
- `.ResumeLayout()` — L852
- `.Scale()` — L870
- `.Select()` — L881
- `.SelectNextControl()` — L888
- `.SendToBack()` — L902
- `.SetBounds()` — L909
- `.Show()` — L932
- `.SuspendLayout()` — L939
- `.ToString()` — L946
- `.Update()` — L958

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/groupboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
