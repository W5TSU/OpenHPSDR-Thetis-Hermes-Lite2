# `Midi2Cat/Midi2Cat.Data/Database.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.IO/OrganiseDialog.cs` (references ×1)
  - `Midi2Cat/MidiMessageManager.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.Data/ControllerMapping.cs` (references ×10)
  - `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs` (calls ×2, references ×1)
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (references ×1, calls ×1)

## Outline

### Types

#### `Midi2CatDatabase` (type, L38)

- `.SaveChanges()` — L75
- `.Exit()` — L95
- `.GetTable()` — L101
- `.AddRow()` — L152
- `.UpdateRow()` — L165
- `.GetRow()` — L177
- `.DeleteRow()` — L188
- `.UpdateOrAdd()` — L208
- `.GetMappings()` — L216
- `.GetMapping()` — L240
- `.GetReverseMapping()` — L252
- `.PopulateMapping()` — L269
- `.PopulateRow()` — L285
- `.BindToDataSource()` — L299
- `.IsDeviceSetup()` — L306
- `.RemoveMapping()` — L320
- `.ConvertFromDBVal()` — L338
- `.SaveMappingAs()` — L350
- `.GetSavedMappings()` — L367
- `.LoadMapping()` — L384
- `.ExportMappings()` — L402
- `.ImportMappings()` — L431
- `.GetImportedMappings()` — L446
- `.AddFromImport()` — L457
- `.RemoveSavedMapping()` — L497
- `.RenameSavedMapping()` — L523
- `.GetSettingTable()` — L550
- `.SetSetting()` — L572
- `.GetSetting()` — L593
- `.GetStringSetting()` — L599
- `.GetPrefixFromMidiDeviceName()` — L612
- `.GetLoadedMappingName()` — L623

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/Database.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
