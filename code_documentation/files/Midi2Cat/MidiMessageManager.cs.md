# `Midi2Cat/MidiMessageManager.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Routes incoming MIDI messages to their mapped commands.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Midi2CatCommands.cs` (calls ×3, references ×1)
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDevice.cs` (calls ×5, references ×4)
  - `Midi2Cat/Midi2Cat.Data/ControllerMapping.cs` (references ×2)
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (references ×2)
  - `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/ControllerBinding.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/Database.cs` (references ×1)
- Most-referenced symbols from other files: `.Open()` (×1), `.Close()` (×1), `.SendUpdateToMidi()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2Cat` (namespace, L37)

_No extracted members._

#### `MidiMessageManager` (type, L39)

- **`.Open()`** — L55 — `public void Open()`
  Called by: `.OpenMidi2Cat()` (`Console/Midi2CatCommands.cs`)
- **`.Close()`** — L77 — `public void Close()`
  Called by: `.Open()` (same file), `.CloseMidi2Cat()` (`Console/Midi2CatCommands.cs`)
- **`.PL1Device()`** — L87 — `public MidiDevice PL1Device()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PL1Index()`** — L102 — `public int PL1Index()`
  Called by: `.SendUpdateToMidi()` (same file)
- **`.SendUpdateToMidi()`** — L118 — `public void SendUpdateToMidi(CatCmd cmd, double pct)`
  Sends update to midi.
  Called by: `.SendUpdateToMidi()` (`Console/Midi2CatCommands.cs`)
- **`.PL1InitialButtonLights()`** — L139 — `public void PL1InitialButtonLights(MidiDevice device)`
  Called by: `.InitDevice()` (same file)
- **`.MicroInitialButtonLights()`** — L197 — `public void MicroInitialButtonLights(MidiDevice device)`
  Called by: `.InitDevice()` (same file)
- **`.InitDevice()`** — L206 — `void InitDevice(string deviceName,List<ControllerMapping> mappings, int Idx)`
  Inits device.
  Called by: `.Open()` (same file)
- **`.BindMappingHandlers()`** — L224 — `Dictionary<int, MidMessageHandler> BindMappingHandlers(List<ControllerMapping> mappings)`
  Called by: `.InitDevice()` (same file)
- **`.onMidiDebugMsg()`** — L256 — `void onMidiDebugMsg(int Device, Direction direction, Status status, string msg1, string msg2)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMidiInput()`** — L261 — `void OnMidiInput(MidiDevice Device, int DeviceIdx, int ControlId, int Data, int Status, int Voice, int Channel)`
  Handles/raises the midi input event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/MidiMessageManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
