# `Console/Invoke/textboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/CATTester.cs` (calls ×2, references ×1)
  - `Console/Andromeda/vfosettingspopup.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/frmAddCustomRadio.Designer.cs` (references ×1)
  - `Console/frmFilterManager.Designer.cs` (references ×1)
  - `Console/frmFinder.Designer.cs` (references ×1)
  - `Console/frmLog.Designer.cs` (references ×1)
  - `Console/frmMacroButtonConfig.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - …and 4 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Focus()` (×1), `.Clear()` (×1)

## Outline

### Types

#### `TextBoxTS` (type, L33)

- `.AppendText()` — L652
- `.BringToFront()` — L661
- `.Clear()` — L668
- `.ClearUndo()` — L675
- `.Contains()` — L682
- `.Copy()` — L694
- `.CreateControl()` — L701
- `.CreateGraphics()` — L708
- `.Cut()` — L732
- `.Dispose()` — L739
- `.DoDragDrop()` — L746
- `.Equals()` — L758
- `.FindForm()` — L770
- `.Focus()` — L782
- `.GetChildAtPoint()` — L794
- `.GetContainerControl()` — L806
- `.GetHashCode()` — L818
- `.GetLifetimeService()` — L830
- `.GetNextControl()` — L842
- `.GetType()` — L854
- `.Hide()` — L866
- `.InitializeLifetimeService()` — L873
- `.Invalidate()` — L885
- `.Paste()` — L947
- `.PerformLayout()` — L954
- `.PointToClient()` — L972
- `.PointToScreen()` — L984
- `.PreProcessMessage()` — L996
- `.RectangleToClient()` — L1008
- `.RectangleToScreen()` — L1020
- `.Refresh()` — L1032
- `.ResetBackColor()` — L1039
- `.ResetBindings()` — L1046
- `.ResetCursor()` — L1053
- `.ResetFont()` — L1060
- `.ResetForeColor()` — L1067
- `.ResetImeMode()` — L1074
- `.ResetRightToLeft()` — L1081
- `.ResetText()` — L1088
- `.ResumeLayout()` — L1095
- `.Scale()` — L1113
- `.ScrollToCaret()` — L1124
- `.Select()` — L1131
- `.SelectAll()` — L1149
- `.SelectNextControl()` — L1156
- `.SendToBack()` — L1170
- `.SetBounds()` — L1177
- `.Show()` — L1200
- `.SuspendLayout()` — L1207
- `.ToString()` — L1214
- `.Undo()` — L1226
- `.Update()` — L1233

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/textboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
