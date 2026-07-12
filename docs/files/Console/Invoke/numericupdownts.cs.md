# `Console/Invoke/numericupdownts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/Memory/MemoryForm.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - `Console/RAForm.Designer.cs` (references ×1)
  - `Console/rxa.Designer.cs` (references ×1)
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
  - …and 1 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `NumericUpDownTS` (type, L33)

- `.BringToFront()` — L640
- `.Contains()` — L647
- `.CreateControl()` — L659
- `.CreateGraphics()` — L666
- `.Dispose()` — L690
- `.DoDragDrop()` — L697
- `.DownButton()` — L709
- `.Equals()` — L716
- `.FindForm()` — L728
- `.Focus()` — L740
- `.GetChildAtPoint()` — L752
- `.GetContainerControl()` — L764
- `.GetHashCode()` — L776
- `.GetLifetimeService()` — L788
- `.GetNextControl()` — L800
- `.GetType()` — L812
- `.Hide()` — L824
- `.InitializeLifetimeService()` — L831
- `.Invalidate()` — L843
- `.PerformLayout()` — L903
- `.PointToClient()` — L921
- `.PointToScreen()` — L933
- `.PreProcessMessage()` — L945
- `.RectangleToClient()` — L957
- `.RectangleToScreen()` — L969
- `.Refresh()` — L981
- `.ResetBackColor()` — L988
- `.ResetBindings()` — L995
- `.ResetCursor()` — L1002
- `.ResetFont()` — L1009
- `.ResetForeColor()` — L1016
- `.ResetImeMode()` — L1023
- `.ResetRightToLeft()` — L1030
- `.ResetText()` — L1037
- `.ResumeLayout()` — L1044
- `.Scale()` — L1062
- `.ScrollControlIntoView()` — L1073
- `.Select()` — L1082
- `.SelectNextControl()` — L1098
- `.SendToBack()` — L1112
- `.SetAutoScrollMargin()` — L1119
- `.SetBounds()` — L1128
- `.Show()` — L1151
- `.SuspendLayout()` — L1158
- `.ToString()` — L1165
- `.UpButton()` — L1177
- `.Update()` — L1184
- `.Validate()` — L1191
- `.OnMouseWheel()` — L1212

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/numericupdownts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
