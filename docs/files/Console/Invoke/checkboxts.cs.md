# `Console/Invoke/checkboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmInfoBarPopup.cs` (references ×2)
  - `Console/setup.cs` (references ×1, calls ×1)
  - `Console/ucInfoBar.cs` (references ×2)
  - `Console/AmpView.Designer.cs` (references ×1)
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmBandwidth.Designer.cs` (references ×1)
  - …and 18 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetType()` (×1)

## Outline

### Types

#### `CheckBoxTS` (type, L33)

- `.BringToFront()` — L586
- `.Contains()` — L593
- `.CreateControl()` — L605
- `.CreateGraphics()` — L612
- `.Dispose()` — L636
- `.DoDragDrop()` — L643
- `.Equals()` — L655
- `.FindForm()` — L667
- `.Focus()` — L679
- `.GetChildAtPoint()` — L691
- `.GetContainerControl()` — L703
- `.GetHashCode()` — L715
- `.GetLifetimeService()` — L727
- `.GetNextControl()` — L739
- `.GetType()` — L751
- `.Hide()` — L763
- `.InitializeLifetimeService()` — L770
- `.Invalidate()` — L782
- `.PerformLayout()` — L844
- `.PointToClient()` — L862
- `.PointToScreen()` — L874
- `.PreProcessMessage()` — L886
- `.RectangleToClient()` — L898
- `.RectangleToScreen()` — L910
- `.Refresh()` — L922
- `.ResetBackColor()` — L929
- `.ResetBindings()` — L936
- `.ResetCursor()` — L943
- `.ResetFont()` — L950
- `.ResetForeColor()` — L957
- `.ResetImeMode()` — L964
- `.ResetRightToLeft()` — L971
- `.ResetText()` — L978
- `.ResumeLayout()` — L985
- `.Scale()` — L1003
- `.Select()` — L1014
- `.SelectNextControl()` — L1021
- `.SendToBack()` — L1035
- `.SetBounds()` — L1042
- `.Show()` — L1065
- `.SuspendLayout()` — L1072
- `.ToString()` — L1079
- `.Update()` — L1091

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/checkboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
