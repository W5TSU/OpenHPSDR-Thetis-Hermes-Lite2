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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `RawInput_dll` (namespace, L4)

_No extracted members._

#### `DeviceInfo` (type, L8)

- **`.ToString()`** — L23 — `public override string ToString()`
  Returns the string representation.
  Called by: `.DeviceAudit()` (`RawInput/Win32.cs`)

#### `DeviceInfoMouse` (type, L29)

- **`.ToString()`** — L37 — `public override string ToString()`
  ReSharper restore MemberCanBePrivate.Global
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `DeviceInfoKeyboard` (type, L43)

- **`.ToString()`** — L52 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `DeviceInfoHid` (type, L64)

- **`.ToString()`** — L72 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `BroadcastDeviceInterface` (type, L78)

_No extracted members._

#### `Rawinputdevicelist` (type, L91)

_No extracted members._

#### `RawData` (type, L98)

_No extracted members._

#### `InputData` (type, L109)

_No extracted members._

#### `Rawinputheader` (type, L116)

- **`.ToString()`** — L124 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Rawhid` (type, L130)

- **`.ToString()`** — L137 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Rawmouse` (type, L143)

_No extracted members._

#### `Rawkeyboard` (type, L164)

- **`.ToString()`** — L174 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RawInputDevice` (type, L181)

- **`.ToString()`** — L189 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/DataStructures.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
