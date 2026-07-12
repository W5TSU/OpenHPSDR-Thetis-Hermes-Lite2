# `Console/CAT/JustinIO.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Low-level Win32 serial I/O used beneath the serial classes.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/SIOListenerII.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.dbgWriteLine()` (×1)

## Outline

### Types

#### `ModifiedJustinIO` (namespace, L40)

_No extracted members._

#### `CommPort` (type, L41)

- `.CreateFile()` — L188
- `.GetCommState()` — L198
- `.BuildCommDCB()` — L203
- `.SetCommState()` — L208
- `.GetCommTimeouts()` — L213
- `.SetCommTimeouts()` — L218
- `.ReadFile()` — L223
- `.WriteFile()` — L231
- `.CloseHandle()` — L239
- `.GetLastError()` — L243
- `.GetCommModemStatus()` — L246
- `.EscapeCommFunction()` — L254
- `.DeviceIoControl()` — L271
- `.dbgWriteLine()` — L288
- `.Open()` — L308
- `.Close()` — L373
- `.bytesToString()` — L384
- `.checkIfCharsAvail()` — L444
- `.Read()` — L467
- `.Write()` — L505
- `.setRTS()` — L528
- `.setDTR()` — L540
- `.isModemBitOn()` — L555
- `.isCTS()` — L572
- `.isDSR()` — L577
- `.isRI()` — L582
- `.isRLSD()` — L586

#### `DCB` (type, L79)

_No extracted members._

#### `SERIAL_STATUS` (type, L115)

_No extracted members._

#### `COMMTIMEOUTS` (type, L170)

_No extracted members._

#### `OVERLAPPED` (type, L179)

_No extracted members._

#### `HexCon` (type, L592)

- `.ByteToString()` — L594
- `.StringToByte()` — L601

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/JustinIO.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
