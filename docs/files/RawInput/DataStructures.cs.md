# `RawInput/DataStructures.cs`

**Functional area:** [18. Raw keyboard/mouse input (RawInput)](../../CODE_OUTLINE.md#18-raw-keyboardmouse-input-rawinput)

**Role:** Win32 interop declarations and device-name registry lookup.

## How this file is used

- Used by (incoming references from other files):
  - `RawInput/Win32.cs` (references ×3, calls ×1)
  - `RawInput/RawKeyboard.cs` (references ×1)
  - `RawInput/RawMouse.cs` (references ×1)
- Uses (outgoing references to other files):
  - `RawInput/Enumerations.cs` (references ×4)
- Most-referenced symbols from other files: `.ToString()` (×1)

## Outline

### Types

#### `RawInput_dll` (namespace, L4)

_No extracted members._

#### `DeviceInfo` (type, L8)

- `.ToString()` — L23

#### `DeviceInfoMouse` (type, L29)

- `.ToString()` — L37

#### `DeviceInfoKeyboard` (type, L43)

- `.ToString()` — L52

#### `DeviceInfoHid` (type, L64)

- `.ToString()` — L72

#### `BroadcastDeviceInterface` (type, L78)

_No extracted members._

#### `Rawinputdevicelist` (type, L91)

_No extracted members._

#### `RawData` (type, L98)

_No extracted members._

#### `InputData` (type, L109)

_No extracted members._

#### `Rawinputheader` (type, L116)

- `.ToString()` — L124

#### `Rawhid` (type, L130)

- `.ToString()` — L137

#### `Rawmouse` (type, L143)

_No extracted members._

#### `Rawkeyboard` (type, L164)

- `.ToString()` — L174

#### `RawInputDevice` (type, L181)

- `.ToString()` — L189

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/DataStructures.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
