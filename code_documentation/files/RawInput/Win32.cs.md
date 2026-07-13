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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Win32` (type, L9)

- **`.LoWord()`** — L11 — `public static int LoWord(int dwValue)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HiWord()`** — L16 — `public static int HiWord(Int64 dwValue)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LowWord()`** — L21 — `public static ushort LowWord(uint val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HighWord()`** — L26 — `public static ushort HighWord(uint val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BuildWParam()`** — L31 — `public static uint BuildWParam(ushort low, ushort high)`
  Builds wparam.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRawInputData()`** — L77 — `[DllImport("User32.dll", SetLastError = true)] internal static extern int GetRawInputData(IntPtr hRawInput, DataCommand command, [Out] out InputData buffer, [In, Out] ref int size,`
  Returns raw input data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRawInputDeviceInfo()`** — L83 — `[DllImport("User32.dll", SetLastError = true)] internal static extern uint GetRawInputDeviceInfo(IntPtr hDevice, RawInputDeviceInfo command, IntPtr pData, ref uint size)`
  Returns raw input device info.
  Called by: `.DeviceAudit()` (same file)
- **`.GetRawInputDeviceList()`** — L90 — `[DllImport("User32.dll", SetLastError = true)] internal static extern uint GetRawInputDeviceList(IntPtr pRawInputDeviceList, ref uint numberDevices, uint size)`
  Returns raw input device list.
  Called by: `.DeviceAudit()` (same file)
- **`.RegisterRawInputDevices()`** — L93 — `[DllImport("User32.dll", SetLastError = true)] internal static extern bool RegisterRawInputDevices(RawInputDevice[] pRawInputDevice, uint numberDevices, uint size)`
  Registers raw input devices.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RegisterDeviceNotification()`** — L96 — `[DllImport("user32.dll", SetLastError = true)] internal static extern IntPtr RegisterDeviceNotification(IntPtr hRecipient, IntPtr notificationFilter, DeviceNotification flags)`
  Registers device notification.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UnregisterDeviceNotification()`** — L99 — `[DllImport("user32.dll", SetLastError = true)] internal static extern bool UnregisterDeviceNotification(IntPtr handle)`
  Unregisters device notification.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DeviceAudit()`** — L102 — `public static void DeviceAudit()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDeviceType()`** — L181 — `public static string GetDeviceType(uint device)`
  Returns device type.
  Called by: `.DeviceAudit()` (same file)
- **`.GetDeviceDescription()`** — L195 — `public static string GetDeviceDescription(string device)`
  Returns device description.
  Called by: `.DeviceAudit()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/RawInput/Win32.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
