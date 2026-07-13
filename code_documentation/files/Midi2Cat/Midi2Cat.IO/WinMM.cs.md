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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `WinMM` (type, L34)

- **`.MidiInGetNumDevs()`** — L58 — `[DllImport("winmm.dll", EntryPoint = "midiInGetNumDevs", CharSet = CharSet.Ansi)] public static extern int MidiInGetNumDevs()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiInGetDevCaps()`** — L61 — `[DllImport("winmm.dll", EntryPoint = "midiInGetDevCaps", CharSet = CharSet.Ansi)] public static extern int MidiInGetDevCaps(int uDeviceID, ref MIDIINCAPS caps, int cbMidiInCaps)`
  Called by: `.MidiInGetName()` (`Midi2Cat/Midi2Cat.IO/MidiDevices.cs`)
- **`.MidiInOpen()`** — L64 — `[DllImport("winmm.dll", EntryPoint = "midiInOpen", CharSet = CharSet.Ansi)] public static extern int MidiInOpen(out IntPtr lphMidiIn, uint uDeviceID, MidiInCallback dwCallback, Int`
  Called by: `.OpenMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiInClose()`** — L67 — `[DllImport("winmm.dll", EntryPoint = "midiInClose", CharSet = CharSet.Ansi)] public static extern int MidiInClose(IntPtr hMidiIn)`
  Called by: `.CloseMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiInReset()`** — L70 — `[DllImport("winmm.dll", EntryPoint = "midiInReset", CharSet = CharSet.Ansi)] public static extern int MidiInReset(IntPtr hMidiIn)`
  Called by: `.CloseMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiInStart()`** — L73 — `[DllImport("winmm.dll", EntryPoint = "midiInStart", CharSet = CharSet.Ansi)] public static extern int MidiInStart(IntPtr hMidiIn)`
  Called by: `.OpenMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiInStop()`** — L76 — `[DllImport("winmm.dll", EntryPoint = "midiInStop", CharSet = CharSet.Ansi)] public static extern int MidiInStop(IntPtr hMidiIn)`
  Called by: `.CloseMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiInAddBuffer()`** — L79 — `[DllImport("winmm.dll", EntryPoint = "midiInAddBuffer")] public static extern int MidiInAddBuffer(IntPtr hMidiIn, IntPtr headerPtr, int cbMidiInHdr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiInPrepareHeader()`** — L82 — `[DllImport("winmm.dll", EntryPoint = "midiInPrepareHeader")] public static extern int MidiInPrepareHeader(IntPtr hMidiIn, IntPtr headerPtr, int cbMidiInHdr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiInUnprepareHeader()`** — L85 — `[DllImport("winmm.dll", EntryPoint = "midiInUnprepareHeader")] public static extern int MidiInUnprepareHeader(int hMidiIn, IntPtr headerPtr, int cbMidiInHdr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiInGetErrorText()`** — L88 — `[DllImport("winmm.dll", EntryPoint = "midiInGetErrorText")] public static extern int MidiInGetErrorText(int wError, StringBuilder lpText, int cchText)`
  Called by: `.OpenMidiIn()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`), `.OpenMidiOut()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutGetNumDevs()`** — L107 — `[DllImport("winmm.dll", EntryPoint = "midiOutGetNumDevs")] public static extern int MidiOutGetNumDevs()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiOutGetDevCaps()`** — L110 — `[DllImport("winmm.dll", EntryPoint = "midiOutGetDevCaps")] public static extern int MidiOutGetDevCaps(int uDeviceID, ref MIDIOUTCAPS caps, int cbMidiOutCaps)`
  Called by: `.MidiOutGetName()` (`Midi2Cat/Midi2Cat.IO/MidiDevices.cs`)
- **`.MidiOutOpen()`** — L113 — `[DllImport("winmm.dll", EntryPoint = "midiOutOpen")] public static extern int MidiOutOpen(out IntPtr lphMidiOut, uint uDeviceID, IntPtr dwCallback, IntPtr dwInstance, MidiOutOpenFl`
  Called by: `.OpenMidiOut()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutClose()`** — L116 — `[DllImport("winmm.dll", EntryPoint = "midiOutClose")] public static extern int MidiOutClose(IntPtr hMidiOut)`
  Called by: `.CloseMidiOut()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutShortMessage()`** — L119 — `[DllImport("winmm.dll", EntryPoint = "midiOutShortMsg")] public static extern int MidiOutShortMessage(IntPtr hMidiOut, uint dwMsg)`
  Called by: `.SendMsg()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutLongMessage()`** — L122 — `[DllImport("winmm.dll", EntryPoint = "midiOutLongMsg")] public static extern int MidiOutLongMessage(int handle, IntPtr headerPtr, int sizeOfMidiHeader)`
  Called by: `.SendLongMessage()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutPrepareHeader()`** — L125 — `[DllImport("winmm.dll", EntryPoint = "midiOutPrepareHeader")] public static extern int MidiOutPrepareHeader(int handle, IntPtr headerPtr, int sizeOfMidiHeader)`
  Called by: `.SendLongMessage()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)
- **`.MidiOutUnprepareHeader()`** — L128 — `[DllImport("winmm.dll", EntryPoint = "midiOutUnprepareHeader")] public static extern int MidiOutUnprepareHeader(int handle, IntPtr headerPtr, int sizeOfMidiHeader)`
  Called by: `.SendLongMessage()` (`Midi2Cat/Midi2Cat.IO/MidiDevice.cs`)

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
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2Cat.IO/WinMM.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
