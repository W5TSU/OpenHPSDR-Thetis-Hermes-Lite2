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

- `.GetRelayValues()` — L168
- `.GetRelayValue()` — L173
- `.SetRelay()` — L193
- `.SetRelays()` — L226
- `.CloseDevice()` — L242

#### `UsbBCDDevices` (type, L258)

- `.OpenDevice()` — L365
- `.SetRelay()` — L371
- `.SetRelays()` — L376
- `.SetBCDbyBand()` — L381
- `.GetRelay()` — L464
- `.GetRelays()` — L469
- `.CloseDevice()` — L474
- `.ResetDevice()` — L480
- `.GetDeviceCount()` — L493
- `.PopulateDeviceList()` — L505

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/UsbBCDCable.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
