# `Console/CAT/SIOListenerII.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Serial-port wrapper and the per-port listener threads (CAT, PTT, keyer ports).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×7, calls ×3)
- Uses (outgoing references to other files):
  - `Console/CAT/SDRSerialPortII.cs` (calls ×36, references ×7)
  - `Console/CAT/CATParser.cs` (calls ×8, references ×7)
  - `Console/Andromeda/Andromeda.cs` (references ×7)
  - `Console/CAT/SerialRxEvent.cs` (references ×7)
  - `Console/CAT/JustinIO.cs` (calls ×1)
- Most-referenced symbols from other files: `.disableCAT2()` (×1), `.disableCAT3()` (×1), `.disableCAT4()` (×1)

## Outline

### Functions

- `ParseString()` — L199

### Types

#### `SIOListenerII` (type, L31)

- `.enableCAT()` — L67
- `.disableCAT()` — L149
- `.Initialize()` — L186
- `.console_Closing()` — L256
- `.console_Activated()` — L264
- `.SerialRXEventHandler()` — L274

#### `SIO2ListenerII` (type, L324)

- `.enableCAT2()` — L361
- `.disableCAT2()` — L435
- `.Initialize()` — L472
- `.console_Closing()` — L485
- `.console_Activated()` — L493
- `.SerialRX2EventHandler()` — L503

#### `SIO3ListenerII` (type, L556)

- `.enableCAT3()` — L593
- `.disableCAT3()` — L667
- `.Initialize()` — L704
- `.console_Closing()` — L717
- `.console_Activated()` — L725
- `.SerialRX3EventHandler()` — L735

#### `SIO4ListenerII` (type, L784)

- `.enableCAT4()` — L820
- `.disableCAT4()` — L894
- `.Initialize()` — L931
- `.console_Closing()` — L944
- `.console_Activated()` — L952
- `.SerialRX4EventHandler()` — L962

#### `SIO5ListenerII` (type, L1013)

- `.enableCAT5()` — L1050
- `.disableCAT5()` — L1067
- `.Initialize()` — L1104
- `.console_Closing()` — L1117
- `.console_Activated()` — L1125
- `.SerialRX5EventHandler()` — L1135

#### `SIO6ListenerII` (type, L1187)

- `.enableCAT6()` — L1224
- `.disableCAT6()` — L1241
- `.Initialize()` — L1278
- `.console_Closing()` — L1291
- `.console_Activated()` — L1299
- `.SerialRX6EventHandler()` — L1309

#### `SIO7ListenerII` (type, L1361)

- `.enableCAT7()` — L1398
- `.disableCAT7()` — L1415
- `.Initialize()` — L1452
- `.console_Closing()` — L1465
- `.console_Activated()` — L1473
- `.SerialRX7EventHandler()` — L1483

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SIOListenerII.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
