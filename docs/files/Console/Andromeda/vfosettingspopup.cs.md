# `Console/Andromeda/vfosettingspopup.cs`

**Functional area:** [13. Andromeda control surface](../../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)

## Outline

### Types

#### `VFOSettingsPopup` (type, L36)

- `.Dispose()` — L60
- `.InitializeComponent()` — L77
- `.VFOSettingsPopup_FormClosing()` — L170
- `.TextBoxTuneStep_MouseDown()` — L178
- `.ButtonMinus_Click()` — L185
- `.ButtonPlus_Click()` — L196
- `.VFOSettingsPopup_Load()` — L206
- `.ButtonClose_Click()` — L214

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Andromeda/vfosettingspopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
