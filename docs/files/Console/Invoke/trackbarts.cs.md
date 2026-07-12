# `Console/Invoke/trackbarts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmNotchPopup.Designer.cs` (references ×1)
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `TrackBarTS` (type, L33)

- `.BringToFront()` — L497
- `.Contains()` — L504
- `.CreateControl()` — L516
- `.CreateGraphics()` — L523
- `.Dispose()` — L547
- `.DoDragDrop()` — L554
- `.Equals()` — L566
- `.FindForm()` — L578
- `.Focus()` — L590
- `.GetChildAtPoint()` — L602
- `.GetContainerControl()` — L614
- `.GetHashCode()` — L626
- `.GetLifetimeService()` — L638
- `.GetNextControl()` — L650
- `.GetType()` — L662
- `.Hide()` — L674
- `.InitializeLifetimeService()` — L681
- `.Invalidate()` — L693
- `.PerformLayout()` — L753
- `.PointToClient()` — L771
- `.PointToScreen()` — L783
- `.PreProcessMessage()` — L795
- `.RectangleToClient()` — L807
- `.RectangleToScreen()` — L819
- `.Refresh()` — L831
- `.ResetBackColor()` — L838
- `.ResetBindings()` — L845
- `.ResetCursor()` — L852
- `.ResetFont()` — L859
- `.ResetForeColor()` — L866
- `.ResetImeMode()` — L873
- `.ResetRightToLeft()` — L880
- `.ResetText()` — L887
- `.ResumeLayout()` — L894
- `.Scale()` — L912
- `.Select()` — L923
- `.SelectNextControl()` — L930
- `.SendToBack()` — L944
- `.SetBounds()` — L951
- `.SetRange()` — L974
- `.Show()` — L985
- `.SuspendLayout()` — L992
- `.ToString()` — L999
- `.Update()` — L1011

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/trackbarts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
