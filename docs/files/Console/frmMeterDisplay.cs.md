# `Console/frmMeterDisplay.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The meter user control and the floating multi-meter display window.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×2)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/ucMeter.Designer.cs` (references ×1)

## Outline

### Types

#### `frmMeterDisplay` (type, L49)

- `.OnRX2Enabled()` — L90
- `.OnWindowStateChanged()` — L114
- `.setTitle()` — L140
- `.frmMeterDisplay_FormClosing()` — L158
- `.TakeOwner()` — L168
- `.SetWindowPos()` — L180
- `.OnLoad()` — L184

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmMeterDisplay.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
