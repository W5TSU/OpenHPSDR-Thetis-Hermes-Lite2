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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2CatDatabase` (type, L38)

- **`.SaveChanges()`** — L75 — `public void SaveChanges(string MidiDeviceName)`
  Saves changes.
  Called by: `.Exit()` (same file), `.GetTable()` (same file), `.LoadMapping()` (same file), `.AddFromImport()` (same file), `.RemoveSavedMapping()` (same file), `.RenameSavedMapping()` (same file) — and 1 more
- **`.Exit()`** — L95 — `public void Exit()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTable()`** — L101 — `private DataTable GetTable(string MidiDeviceName, DataSet overrideDS = null)`
  Returns table.
  Called by: `.AddRow()` (same file), `.UpdateRow()` (same file), `.GetRow()` (same file), `.DeleteRow()` (same file), `.GetMappings()` (same file), `.GetMapping()` (same file) — and 8 more
- **`.AddRow()`** — L152 — `public void AddRow(string MidiDeviceName, DataRow row)`
  Adds row.
  Called by: `.UpdateOrAdd()` (same file)
- **`.UpdateRow()`** — L165 — `private bool UpdateRow(string MidiDeviceName, ControllerMapping mapping)`
  Updates row.
  Called by: `.UpdateOrAdd()` (same file)
- **`.GetRow()`** — L177 — `private ControllerMapping GetRow(string MidiDeviceName, int Idx)`
  Returns row.
  Called by: `.GetMappings()` (same file)
- **`.DeleteRow()`** — L188 — `public void DeleteRow(string MidiDeviceName, ControllerMapping mapping)`
  Deletes row.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateOrAdd()`** — L208 — `public void UpdateOrAdd(string MidiDeviceName, ControllerMapping mapping)`
  Updates or add.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMappings()`** — L216 — `public List<ControllerMapping> GetMappings(string MidiDeviceName, MappingFilter filter)`
  Returns mappings.
  Called by: `.IsDeviceSetup()` (same file)
- **`.GetMapping()`** — L240 — `public ControllerMapping GetMapping(string MidiDeviceName, int MidiControlId)`
  Returns mapping.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetReverseMapping()`** — L252 — `public ControllerMapping GetReverseMapping(string MidiDeviceName, CatCmd cmd)`
  Returns reverse mapping.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopulateMapping()`** — L269 — `private ControllerMapping PopulateMapping(DataRow dr)`
  Called by: `.GetRow()` (same file), `.GetMapping()` (same file), `.GetReverseMapping()` (same file)
- **`.PopulateRow()`** — L285 — `public DataRow PopulateRow(DataRow dr, ControllerMapping mapping)`
  Called by: `.AddRow()` (same file), `.UpdateRow()` (same file)
- **`.BindToDataSource()`** — L299 — `public void BindToDataSource(BindingSource source, string MidiDeviceName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsDeviceSetup()`** — L306 — `public bool IsDeviceSetup(string MidiDeviceName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveMapping()`** — L320 — `public void RemoveMapping(string MidiDeviceName, int catCmd)`
  Removes mapping.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ConvertFromDBVal()`** — L338 — `public T ConvertFromDBVal<T>(object obj)`
  Converts from dbval.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SaveMappingAs()`** — L350 — `public bool SaveMappingAs(string midiDeviceName, string Name, bool replace)`
  Saves mapping as.
  Called by: `.SaveChanges()` (same file)
- **`.GetSavedMappings()`** — L367 — `public string[] GetSavedMappings()`
  Returns saved mappings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadMapping()`** — L384 — `public bool LoadMapping(string midiDeviceName, string Name)`
  Loads mapping.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExportMappings()`** — L402 — `public bool ExportMappings(string fileName, string[] mappings)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ImportMappings()`** — L431 — `public bool ImportMappings(string fileName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetImportedMappings()`** — L446 — `public string[] GetImportedMappings()`
  Returns imported mappings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddFromImport()`** — L457 — `public bool AddFromImport( string[] mappings)`
  Adds from import.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveSavedMapping()`** — L497 — `public bool RemoveSavedMapping(string midiDeviceName, string Name)`
  Removes saved mapping.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RenameSavedMapping()`** — L523 — `public bool RenameSavedMapping(string midiDeviceName, string OldName, string NewName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetSettingTable()`** — L550 — `private DataTable GetSettingTable()`
  Returns setting table.
  Called by: `.SetSetting()` (same file), `.GetSetting()` (same file)
- **`.SetSetting()`** — L572 — `private void SetSetting(string name, string value)`
  Sets setting.
  Called by: `.SaveMappingAs()` (same file), `.LoadMapping()` (same file), `.RemoveSavedMapping()` (same file), `.RenameSavedMapping()` (same file)
- **`.GetSetting()`** — L593 — `private DataRow GetSetting(string name)`
  Returns setting.
  Called by: `.GetStringSetting()` (same file)
- **`.GetStringSetting()`** — L599 — `private string GetStringSetting(string name, string defaultValue)`
  Returns string setting.
  Called by: `.GetLoadedMappingName()` (same file)
- **`.GetPrefixFromMidiDeviceName()`** — L612 — `private string GetPrefixFromMidiDeviceName(string midiDeviceName)`
  Returns prefix from midi device name.
  Called by: `.SaveMappingAs()` (same file), `.LoadMapping()` (same file), `.RemoveSavedMapping()` (same file), `.RenameSavedMapping()` (same file), `.GetLoadedMappingName()` (same file)
- **`.GetLoadedMappingName()`** — L623 — `public string GetLoadedMappingName(string midiDeviceName)`
  Returns loaded mapping name.
  Called by: `.SaveChanges()` (same file), `.RemoveSavedMapping()` (same file), `.RenameSavedMapping()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/Database.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
