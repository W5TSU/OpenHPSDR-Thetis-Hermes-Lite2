# `Midi2Cat/Midi2Cat.Data/EnumsDB.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Persistence and object model for controller-to-command mappings.

## How this file is used

- Used by (incoming references from other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs` (calls ×1, references ×1)
- Most-referenced symbols from other files: `.SetCatCmdInUse()` (×3)

## Outline

### Types

#### `EnumsDB` (type, L35)

- `.AddControlTypes()` — L53
- `.AddCatCmds()` — L77
- `.BindToDataSource()` — L104
- `.SetCatCmdInUse()` — L112

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/EnumsDB.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
