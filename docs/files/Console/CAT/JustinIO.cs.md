# `Console/CAT/JustinIO.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Low-level Win32 serial I/O used beneath the serial classes.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/SIOListenerII.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.dbgWriteLine()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ModifiedJustinIO` (namespace, L40)

_No extracted members._

#### `CommPort` (type, L41)

- **`.CreateFile()`** — L188 — `[DllImport("kernel32.dll")] private static extern int CreateFile( string lpFileName, uint dwDesiredAccess, int dwShareMode,`
  Creates file.
  Called by: `.Open()` (same file)
- **`.GetCommState()`** — L198 — `[DllImport("kernel32.dll")] private static extern bool GetCommState( int hFile, ref DCB lpDCB )`
  Returns comm state.
  Called by: `.Open()` (same file)
- **`.BuildCommDCB()`** — L203 — `[DllImport("kernel32.dll")] private static extern bool BuildCommDCB( string lpDef, ref DCB lpDCB )`
  Builds comm dcb.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCommState()`** — L208 — `[DllImport("kernel32.dll")] private static extern bool SetCommState( int hFile, ref DCB lpDCB )`
  Sets comm state.
  Called by: `.Open()` (same file)
- **`.GetCommTimeouts()`** — L213 — `[DllImport("kernel32.dll")] private static extern bool GetCommTimeouts( int hFile, ref COMMTIMEOUTS lpCommTimeouts )`
  Returns comm timeouts.
  Called by: `.Open()` (same file)
- **`.SetCommTimeouts()`** — L218 — `[DllImport("kernel32.dll")] private static extern bool SetCommTimeouts( int hFile, ref COMMTIMEOUTS lpCommTimeouts )`
  Sets comm timeouts.
  Called by: `.Open()` (same file)
- **`.ReadFile()`** — L223 — `[DllImport("kernel32.dll")] private static extern bool ReadFile( int hFile, byte[] lpBuffer, int nNumberOfBytesToRead,`
  Reads file.
  Called by: `.Read()` (same file)
- **`.WriteFile()`** — L231 — `[DllImport("kernel32.dll")] private static extern bool WriteFile( int hFile, byte[] lpBuffer, int nNumberOfBytesToWrite,`
  Writes file.
  Called by: `.Write()` (same file)
- **`.CloseHandle()`** — L239 — `[DllImport("kernel32.dll")] private static extern bool CloseHandle( int hObject )`
  Closes handle.
  Called by: `.Close()` (same file)
- **`.GetLastError()`** — L243 — `[DllImport("kernel32.dll")] private static extern uint GetLastError()`
  Returns last error.
  Called by: `.Open()` (same file)
- **`.GetCommModemStatus()`** — L246 — `[DllImport("kernel32.dll")] private static extern Boolean GetCommModemStatus ( int hFile, ref uint lpModemStat`
  Returns comm modem status.
  Called by: `.isModemBitOn()` (same file)
- **`.EscapeCommFunction()`** — L254 — `[DllImport("kernel32.dll")] private static extern Boolean EscapeCommFunction ( int hFile, uint dwFunc`
  Called by: `.setRTS()` (same file), `.setDTR()` (same file)
- **`.DeviceIoControl()`** — L271 — `[DllImport("kernel32.dll")] private static extern Boolean DeviceIoControl ( int hfile, UInt32 IOctlcode,`
  DWORD dwIoControlCode, LPVOID lpInBuffer, DWORD nInBufferSize, LPVOID lpOutBuffer, DWORD nOutBufferSize, LPDWORD lpBytesReturned, LPOVERLAPPED lpOverlapped );
  Called by: `.checkIfCharsAvail()` (same file)
- **`.dbgWriteLine()`** — L288 — `private void dbgWriteLine(string s)`
  Called by: `.Read()` (same file), `.Write()` (same file), `ParseString()` (`Console/CAT/SIOListenerII.cs`)
- **`.Open()`** — L308 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L373 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.bytesToString()`** — L384 — `private static string bytesToString(byte[] b, uint n)`
  convert an array of bytes to a string -- assuming bytes are ascii char fixme ... surely there's a libarary routine to do this!
  Called by: `.Read()` (same file), `.Write()` (same file)
- **`.checkIfCharsAvail()`** — L444 — `private uint checkIfCharsAvail()`
  returns count of chars available to be read ... sort of a hack as this is using an ioctl that is not in the documentation but does appear in the dd source this is needed when using the MixW CommEmulDrv virt serial ports becuse it does not respect the SetCommTimeout values, so a read with no data…
  Called by: `.Read()` (same file)
- **`.Read()`** — L467 — `public int Read(byte[] buf_to_fill, int max_size)`
  read chars from serial port, returns number of chars this routine will not block waiting for chars, if there are none avail it will return 0 immediately
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Write()`** — L505 — `public int Write(byte[] b, int count)`
  returns number of bytes written
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setRTS()`** — L528 — `public void setRTS(bool on)`
  routines to bit bang RTS and DTR
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L540 — `public void setDTR(bool on)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isModemBitOn()`** — L555 — `private bool isModemBitOn(uint which_bit)`
  Called by: `.isCTS()` (same file), `.isDSR()` (same file), `.isRI()` (same file), `.isRLSD()` (same file)
- **`.isCTS()`** — L572 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L577 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L582 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L586 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `DCB` (type, L79)

_No extracted members._

#### `SERIAL_STATUS` (type, L115)

_No extracted members._

#### `COMMTIMEOUTS` (type, L170)

_No extracted members._

#### `OVERLAPPED` (type, L179)

_No extracted members._

#### `HexCon` (type, L592)

- **`.ByteToString()`** — L594 — `public static string ByteToString(byte[] InBytes)`
  converter hex string to byte and byte to hex string
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToByte()`** — L601 — `public static byte[] StringToByte(string InString)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/JustinIO.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
