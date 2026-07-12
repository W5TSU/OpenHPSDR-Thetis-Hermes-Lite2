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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `HardwareSpecific` (type, L59)

- **`.GetDefaultVoltCalibration()`** — L288 — `public static (float, float) GetDefaultVoltCalibration()`
  Returns default volt calibration.
  Called by: `.btnAmpDefault_Click()` (`Console/setup.cs`)
- **`.StringModelToEnum()`** — L352 — `public static HPSDRModel StringModelToEnum(string sModel)`
  enums/string conversion
  Called by: `.DBWritten()` (`Console/clsDBMan.cs`), `.createNewDB()` (`Console/clsDBMan.cs`), `.ImportAsAvailable()` (`Console/clsDBMan.cs`), `.MakeBackupAvailable()` (`Console/clsDBMan.cs`), `.getModelFromDB()` (`Console/setup.cs`), `.comboRadioModel_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.EnumModelToString()`** — L392 — `public static string EnumModelToString(HPSDRModel model)`
  Called by: `.InitAvailableDBs()` (`Console/frmDBMan.cs`)
- **`.RXMeterCalbrationOffsetDefaults()`** — L437 — `public static float RXMeterCalbrationOffsetDefaults(HPSDRModel model)`
  Called by: `.ResetLevelCalibration()` (`Console/console.cs`)
- **`.RXDisplayCalbrationOffsetDefauls()`** — L454 — `public static float RXDisplayCalbrationOffsetDefauls(HPSDRModel model)`
  Called by: `.ResetLevelCalibration()` (`Console/console.cs`)
- **`.DefaultPAGainsForBands()`** — L501 — `public static float[] DefaultPAGainsForBands(HPSDRModel model)`
  Called by: `.ResetGainDefaultsForModel()` (`Console/setup.cs`)
- **`.HasSteppedAttenuation()`** — L843 — `public static bool HasSteppedAttenuation(int rx)`
  Called by: `.updateAttenuationInfo()` (`Console/setup.cs`), `.chkHermesStepAttenuator_CheckedChanged()` (`Console/setup.cs`), `.chkRX2StepAtt_CheckedChanged()` (`Console/setup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsHardwareSpecific.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
