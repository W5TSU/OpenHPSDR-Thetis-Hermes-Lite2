# `Console/CAT/UsbBCDCable.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** PTT via serial control lines; band-decoder output over a USB BCD cable.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×6, references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×1)
- Most-referenced symbols from other files: `.SetBCDbyBand()` (×2), `.OpenDevice()` (×1), `.SetRelays()` (×1), `.GetRelays()` (×1), `.CloseDevice()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `UsbRelayDeviceNotFoundException` (type, L38)

_No extracted members._

#### `UsbRelayConfigurationException` (type, L46)

_No extracted members._

#### `UsbRelayInvalidRelayException` (type, L54)

_No extracted members._

#### `UsbRelayReadException` (type, L62)

_No extracted members._

#### `UsbRelayWriteException` (type, L70)

_No extracted members._

#### `UsbRelayStatusException` (type, L78)

_No extracted members._

#### `UsbBCDCable` (type, L96)

- **`.GetRelayValues()`** — L168 — `public byte GetRelayValues()`
  Get the relay values
  Called by: `.GetRelayValue()` (same file), `.SetRelay()` (same file), `.GetRelays()` (same file)
- **`.GetRelayValue()`** — L173 — `public bool GetRelayValue(int relay)`
  Returns relay value.
  Called by: `.GetRelay()` (same file)
- **`.SetRelay()`** — L193 — `public void SetRelay(int relay, bool value)`
  Set an individual relay value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRelays()`** — L226 — `public void SetRelays(byte values)`
  Set all the relay values, each relay is a bit
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CloseDevice()`** — L242 — `public void CloseDevice()`
  Closes device.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `UsbBCDDevices` (type, L258)

- **`.OpenDevice()`** — L365 — `public void OpenDevice(string serialNumber)`
  Opens device.
  Called by: `.comboUsbDevices_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.SetRelay()`** — L371 — `public void SetRelay(string serialNumber, int relay, bool value)`
  Sets relay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRelays()`** — L376 — `public void SetRelays(string serialNumber, byte values)`
  Sets relays.
  Called by: `.SetRelay()` (same file), `.SetBCDbyBand()` (same file), `.CloseUsbBcdDevice()` (`Console/setup.cs`)
- **`.SetBCDbyBand()`** — L381 — `public void SetBCDbyBand(string serialNumber, Band b)`
  Sets bcdby band.
  Called by: `.updateUsbBCDdevice()` (`Console/setup.cs`), `.comboUsbDevices_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.GetRelay()`** — L464 — `public bool GetRelay(string serialNumber, int relay)`
  Returns relay.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRelays()`** — L469 — `public byte GetRelays(string serialNumber)`
  Returns relays.
  Called by: `.comboUsbDevices_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.CloseDevice()`** — L474 — `public void CloseDevice(string serialNumber)`
  Closes device.
  Called by: `.CloseUsbBcdDevice()` (`Console/setup.cs`)
- **`.ResetDevice()`** — L480 — `public void ResetDevice()`
  Resets device.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetDeviceCount()`** — L493 — `private int GetDeviceCount()`
  Return the device count
  Called by: `.PopulateDeviceList()` (same file)
- **`.PopulateDeviceList()`** — L505 — `private void PopulateDeviceList()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/UsbBCDCable.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
