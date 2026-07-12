# `Console/clsHardwareSpecific.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** Per-hardware-model capability flags and defaults (which options apply to which radio model, incl. HL2).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×7)
  - `Console/clsDBMan.cs` (calls ×4)
  - `Console/console.cs` (calls ×2)
  - `Console/frmDBMan.cs` (calls ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×7)
- Most-referenced symbols from other files: `.StringModelToEnum()` (×6), `.HasSteppedAttenuation()` (×3), `.GetDefaultVoltCalibration()` (×1), `.EnumModelToString()` (×1), `.RXMeterCalbrationOffsetDefaults()` (×1), `.RXDisplayCalbrationOffsetDefauls()` (×1), `.DefaultPAGainsForBands()` (×1)

## Outline

### Types

#### `HardwareSpecific` (type, L59)

- `.GetDefaultVoltCalibration()` — L288
- `.StringModelToEnum()` — L352
- `.EnumModelToString()` — L392
- `.RXMeterCalbrationOffsetDefaults()` — L437
- `.RXDisplayCalbrationOffsetDefauls()` — L454
- `.DefaultPAGainsForBands()` — L501
- `.HasSteppedAttenuation()` — L843

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsHardwareSpecific.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
