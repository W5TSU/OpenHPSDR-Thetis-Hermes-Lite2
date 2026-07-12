# `Console/Invoke/panelts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/eqform.cs` (references ×1, calls ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/rxa.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
  - `Console/ucTunestepOptionsGrid.Designer.cs` (references ×1)
  - `Console/wideband.Designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.BringToFront()` (×1)

## Outline

### Types

#### `PanelTS` (type, L33)

- `.BringToFront()` — L509
- `.Contains()` — L516
- `.CreateControl()` — L528
- `.Dispose()` — L559
- `.DoDragDrop()` — L566
- `.Equals()` — L578
- `.FindForm()` — L590
- `.Focus()` — L602
- `.GetChildAtPoint()` — L614
- `.GetContainerControl()` — L626
- `.GetHashCode()` — L638
- `.GetLifetimeService()` — L650
- `.GetNextControl()` — L662
- `.GetType()` — L674
- `.Hide()` — L686
- `.InitializeLifetimeService()` — L693
- `.Invalidate()` — L705
- `.PerformLayout()` — L765
- `.PointToClient()` — L783
- `.PointToScreen()` — L795
- `.PreProcessMessage()` — L807
- `.RectangleToClient()` — L819
- `.RectangleToScreen()` — L831
- `.Refresh()` — L843
- `.ResetBackColor()` — L850
- `.ResetBindings()` — L857
- `.ResetCursor()` — L864
- `.ResetFont()` — L871
- `.ResetForeColor()` — L878
- `.ResetImeMode()` — L885
- `.ResetRightToLeft()` — L892
- `.ResetText()` — L899
- `.ResumeLayout()` — L906
- `.Scale()` — L924
- `.ScrollControlIntoView()` — L935
- `.Select()` — L946
- `.SelectNextControl()` — L953
- `.SendToBack()` — L967
- `.SetAutoScrollMargin()` — L974
- `.SetBounds()` — L985
- `.Show()` — L1008
- `.SuspendLayout()` — L1015
- `.ToString()` — L1022
- `.Update()` — L1034

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/panelts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
