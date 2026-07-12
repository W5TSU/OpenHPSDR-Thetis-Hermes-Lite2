# `Midi2Cat/Midi2Cat.IO/WinMM.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** MIDI device open/close and message receive/send over the Windows Multimedia (winmm) API.

## How this file is used

- Used by (incoming references from other files):
  - `Midi2Cat/Midi2Cat.IO/MidiDevice.cs` (calls ×13)
  - `Midi2Cat/Midi2Cat.IO/MidiDevices.cs` (calls ×2)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.MidiInGetErrorText()` (×2), `.MidiInOpen()` (×1), `.MidiInStart()` (×1), `.MidiOutOpen()` (×1), `.MidiInClose()` (×1), `.MidiInReset()` (×1), `.MidiInStop()` (×1), `.MidiOutClose()` (×1)

## Outline

### Types

#### `WinMM` (type, L34)

- `.MidiInGetNumDevs()` — L58
- `.MidiInGetDevCaps()` — L61
- `.MidiInOpen()` — L64
- `.MidiInClose()` — L67
- `.MidiInReset()` — L70
- `.MidiInStart()` — L73
- `.MidiInStop()` — L76
- `.MidiInAddBuffer()` — L79
- `.MidiInPrepareHeader()` — L82
- `.MidiInUnprepareHeader()` — L85
- `.MidiInGetErrorText()` — L88
- `.MidiOutGetNumDevs()` — L107
- `.MidiOutGetDevCaps()` — L110
- `.MidiOutOpen()` — L113
- `.MidiOutClose()` — L116
- `.MidiOutShortMessage()` — L119
- `.MidiOutLongMessage()` — L122
- `.MidiOutPrepareHeader()` — L125
- `.MidiOutUnprepareHeader()` — L128

#### `MidiHeader` (type, L40)

_No extracted members._

#### `MIDIINCAPS` (type, L93)

_No extracted members._

#### `MIDIOUTCAPS` (type, L131)

_No extracted members._

#### `MidiInOpenFlags` (type, L145)

_No extracted members._

#### `MidiOutOpenFlags` (type, L155)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.IO/WinMM.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
