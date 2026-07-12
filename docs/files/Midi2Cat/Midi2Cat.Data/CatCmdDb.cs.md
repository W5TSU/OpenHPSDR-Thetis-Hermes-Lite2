# `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (calls ×3, references ×1)
  - `Midi2Cat/Midi2Cat.Data/Database.cs` (calls ×2, references ×1)
  - `Midi2Cat/Midi2Cat.Data/EnumsDB.cs` (calls ×1, references ×1)
  - `Console/Midi2CatCommands.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/CatCommandAttribute.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/ControllerMapping.cs` (references ×1)
  - `Midi2Cat/MidiMessageManager.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.Data/CatCommandAttribute.cs` (references ×1)
- Most-referenced symbols from other files: `.Get()` (×6)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2Cat.Data` (namespace, L41)

_No extracted members._

#### `CatCmdDb` (type, L43)

- **`.Get()`** — L45 — `public static CatCommandAttribute Get(CatCmd Id)`
  Called by: `.GetMappings()` (`Midi2Cat/Midi2Cat.Data/Database.cs`), `.PopulateMapping()` (`Midi2Cat/Midi2Cat.Data/Database.cs`), `.AddCatCmds()` (`Midi2Cat/Midi2Cat.Data/EnumsDB.cs`), `.OnMidiInput()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.LoadNonDataBoundControls()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.mapControlToCommandGrid_CellEndEdit()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)

#### `CatCmd` (type, L56)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/CatCmdDb.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
