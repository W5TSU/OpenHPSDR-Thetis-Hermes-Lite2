# `Console/filter.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** RX filter preset model per mode, the filter-edit form, and the filter-set manager.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×2)
  - `Console/FilterForm.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×10)

## Outline

### Types

#### `FilterPreset` (type, L47)

- `.SetLow()` — L60
- `.SetHigh()` — L65
- `.SetName()` — L70
- `.SetFilter()` — L75
- `.GetLow()` — L82
- `.GetHigh()` — L87
- `.GetBW()` — L92
- `.GetName()` — L97
- `.ToString()` — L109

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/filter.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
