# `Console/Invoke/comboboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/frmAddCustomRadio.Designer.cs` (references ×1)
  - `Console/frmAddCustomRadio.cs` (references ×1)
  - `Console/frmIPv4Picker.Designer.cs` (references ×1)
  - `Console/frmMacroButtonConfig.Designer.cs` (references ×1)
  - `Console/frmSerialPortPicker.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
  - `Console/ucGradientDefault.Designer.cs` (references ×1)
  - …and 1 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `ComboBoxTS` (type, L33)

- `.BeginUpdate()` — L629
- `.BringToFront()` — L636
- `.Contains()` — L643
- `.CreateControl()` — L655
- `.CreateGraphics()` — L662
- `.Dispose()` — L686
- `.DoDragDrop()` — L693
- `.EndUpdate()` — L705
- `.Equals()` — L712
- `.FindForm()` — L724
- `.FindString()` — L736
- `.FindStringExact()` — L760
- `.Focus()` — L784
- `.GetChildAtPoint()` — L796
- `.GetContainerControl()` — L808
- `.GetHashCode()` — L820
- `.GetItemHeight()` — L832
- `.GetItemText()` — L844
- `.GetLifetimeService()` — L856
- `.GetNextControl()` — L868
- `.GetType()` — L880
- `.Hide()` — L892
- `.InitializeLifetimeService()` — L899
- `.Invalidate()` — L911
- `.PerformLayout()` — L973
- `.PointToClient()` — L991
- `.PointToScreen()` — L1003
- `.PreProcessMessage()` — L1015
- `.RectangleToClient()` — L1027
- `.RectangleToScreen()` — L1039
- `.Refresh()` — L1051
- `.ResetBackColor()` — L1058
- `.ResetBindings()` — L1065
- `.ResetCursor()` — L1072
- `.ResetFont()` — L1079
- `.ResetForeColor()` — L1086
- `.ResetImeMode()` — L1093
- `.ResetRightToLeft()` — L1100
- `.ResetText()` — L1107
- `.ResumeLayout()` — L1114
- `.Scale()` — L1132
- `.Select()` — L1143
- `.SelectAll()` — L1161
- `.SelectNextControl()` — L1168
- `.SendToBack()` — L1182
- `.SetBounds()` — L1189
- `.Show()` — L1212
- `.SuspendLayout()` — L1219
- `.ToString()` — L1226
- `.Update()` — L1238

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/comboboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
