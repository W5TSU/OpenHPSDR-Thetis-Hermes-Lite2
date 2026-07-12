# `Midi2Cat/Midi2Cat.IO/MidiDevices.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** MIDI device open/close and message receive/send over the Windows Multimedia (winmm) API.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Midi2Cat/Midi2Cat.IO/WinMM.cs` (calls ×2)

## Outline

### Types

#### `MidiDevices` (type, L34)

- `.MidiInGetName()` — L80
- `.MidiOutGetName()` — L89

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.IO/MidiDevices.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
