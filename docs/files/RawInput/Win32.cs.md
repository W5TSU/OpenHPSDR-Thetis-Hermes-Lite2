# `RawInput/Win32.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Win32 interop declarations and device-name registry lookup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `RawInput/DataStructures.cs` (references ×3, calls ×1)
  - `RawInput/Enumerations.cs` (references ×3)
  - `RawInput/KeyPressEvent.cs` (calls ×1)

## Outline

### Types

#### `Win32` (type, L9)

- `.LoWord()` — L11
- `.HiWord()` — L16
- `.LowWord()` — L21
- `.HighWord()` — L26
- `.BuildWParam()` — L31
- `.GetRawInputData()` — L77
- `.GetRawInputDeviceInfo()` — L83
- `.GetRawInputDeviceList()` — L90
- `.RegisterRawInputDevices()` — L93
- `.RegisterDeviceNotification()` — L96
- `.UnregisterDeviceNotification()` — L99
- `.DeviceAudit()` — L102
- `.GetDeviceType()` — L181
- `.GetDeviceDescription()` — L195

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/Win32.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
