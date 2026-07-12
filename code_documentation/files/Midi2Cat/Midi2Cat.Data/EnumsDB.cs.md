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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `EnumsDB` (type, L35)

- **`.AddControlTypes()`** — L53 — `private void AddControlTypes()`
  Adds control types.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddCatCmds()`** — L77 — `private void AddCatCmds()`
  Adds cat cmds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BindToDataSource()`** — L104 — `public void BindToDataSource(BindingSource source, string tableName)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCatCmdInUse()`** — L112 — `public void SetCatCmdInUse(CatCmd catCmd, bool inUse)`
  Sets cat cmd in use.
  Called by: `.MidiDeviceSetup_Load()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.mapControlToCommandGrid_CellValueChanged()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`), `.mappedCommandsGridView_CellClick()` (`Midi2Cat/Midi2Cat.IO/MidiDeviceSetup.cs`)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.Data/EnumsDB.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
