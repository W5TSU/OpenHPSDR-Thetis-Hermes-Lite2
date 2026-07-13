# `Console/CAT/SDRSerialPortII.cs`

**Functional area:** [10. CAT control and external program interfaces](../../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Serial-port wrapper and the per-port listener threads (CAT, PTT, keyer ports).

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/SIOListenerII.cs` (calls ×36, references ×7)
  - `Console/setup.cs` (calls ×10)
  - `Console/CAT/SerialPortPTT.cs` (calls ×7, references ×1)
  - `Console/CW/CWInput.cs` (calls ×4, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Destroy()` (×15), `.put()` (×8), `.Create()` (×8), `.setCommParms()` (×7), `.StringToParity()` (×5), `.StringToStopBits()` (×5), `.Open()` (×2), `.Close()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `SDRSerialPort` (type, L31)

- **`.Open()`** — L50 — `public void Open()`
  Called by: `.SetPrimaryInput()` (`Console/CW/CWInput.cs`), `.SetSecondaryInput()` (`Console/CW/CWInput.cs`)
- **`.Close()`** — L55 — `public void Close()`
  Called by: `.SetPrimaryInput()` (`Console/CW/CWInput.cs`), `.SetSecondaryInput()` (`Console/CW/CWInput.cs`)
- **`.StringToParity()`** — L60 — `public static Parity StringToParity(string s)`
  Called by: `.initCATandPTTprops()` (`Console/setup.cs`), `.comboCATparity_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT2parity_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT3parity_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT4parity_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.StringToStopBits()`** — L70 — `public static StopBits StringToStopBits(string s)`
  Called by: `.initCATandPTTprops()` (`Console/setup.cs`), `.comboCATstopbits_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT2stopbits_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT3stopbits_SelectedIndexChanged()` (`Console/setup.cs`), `.comboCAT4stopbits_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.setCommParms()`** — L102 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L112 — `public uint put(string s)`
  Called by: `ParseString()` (`Console/CAT/SIOListenerII.cs`), `.SerialRXEventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L119 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`), `.Init()` (`Console/CAT/SerialPortPTT.cs`)
- **`.Destroy()`** — L137 — `public void Destroy()`
  Called by: `.disableCAT()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`), `.Destroy()` (`Console/CAT/SerialPortPTT.cs`)
- **`.isCTS()`** — L150 — `public bool isCTS()`
  Called by: `.isPTT()` (`Console/CAT/SerialPortPTT.cs`), `.isCTS()` (`Console/CAT/SerialPortPTT.cs`)
- **`.isDSR()`** — L156 — `public bool isDSR()`
  Called by: `.isPTT()` (`Console/CAT/SerialPortPTT.cs`), `.isDSR()` (`Console/CAT/SerialPortPTT.cs`)
- **`.isRI()`** — L162 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L168 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L174 — `public void setDTR(bool v)`
  Sets dtr.
  Called by: `.setDTR()` (`Console/CAT/SerialPortPTT.cs`)
- **`.SerialErrorReceived()`** — L180 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L234 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L339 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort2` (type, L346)

- **`.Open()`** — L365 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L370 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L375 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L385 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L418 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT2()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L428 — `public uint put(string s)`
  Called by: `.SerialRX2EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L435 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L453 — `public void Destroy()`
  Called by: `.disableCAT2()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L466 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L472 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L478 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L484 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L490 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L496 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L550 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L620 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort3` (type, L627)

- **`.Open()`** — L646 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L651 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L656 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L666 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L699 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT3()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L709 — `public uint put(string s)`
  Called by: `.SerialRX3EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L716 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L734 — `public void Destroy()`
  Called by: `.disableCAT3()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L747 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L753 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L759 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L765 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L771 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L777 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L831 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L901 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort4` (type, L908)

- **`.Open()`** — L927 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L932 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L937 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L947 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L980 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT4()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L990 — `public uint put(string s)`
  Called by: `.SerialRX4EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L997 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L1015 — `public void Destroy()`
  Called by: `.disableCAT4()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L1028 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L1034 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L1040 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L1046 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L1052 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L1058 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L1112 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L1182 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort5` (type, L1190)

- **`.Open()`** — L1209 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L1214 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L1219 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L1229 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L1262 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT5()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L1272 — `public uint put(string s)`
  Called by: `.SerialRX5EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L1279 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L1297 — `public void Destroy()`
  Called by: `.disableCAT5()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L1310 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L1316 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L1322 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L1328 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L1334 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L1340 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L1394 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L1464 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort6` (type, L1472)

- **`.Open()`** — L1491 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L1496 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L1501 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L1511 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L1544 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT6()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L1554 — `public uint put(string s)`
  Called by: `.SerialRX6EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L1561 — `public int Create()`
  Called by: `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L1579 — `public void Destroy()`
  Called by: `.disableCAT6()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L1592 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L1598 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L1604 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L1610 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L1616 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L1622 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L1676 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L1746 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SDRSerialPort7` (type, L1754)

- **`.Open()`** — L1773 — `public void Open()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Close()`** — L1778 — `public void Close()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToParity()`** — L1783 — `public static Parity StringToParity(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToStopBits()`** — L1793 — `public static StopBits StringToStopBits(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setCommParms()`** — L1826 — `public void setCommParms(int baudrate, Parity p, int databits, StopBits stop)`
  set the comm parms ... can only be done if port is not open -- silently fails if port is open (fixme -- add some error checking)
  Called by: `.enableCAT7()` (`Console/CAT/SIOListenerII.cs`)
- **`.put()`** — L1836 — `public uint put(string s)`
  Called by: `.SerialRX7EventHandler()` (`Console/CAT/SIOListenerII.cs`)
- **`.Create()`** — L1843 — `public int Create()`
  Called by: `.Create()` (same file), `.Initialize()` (`Console/CAT/SIOListenerII.cs`)
- **`.Destroy()`** — L1861 — `public void Destroy()`
  Called by: `.disableCAT7()` (`Console/CAT/SIOListenerII.cs`), `.console_Closing()` (`Console/CAT/SIOListenerII.cs`)
- **`.isCTS()`** — L1874 — `public bool isCTS()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isDSR()`** — L1880 — `public bool isDSR()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRI()`** — L1886 — `public bool isRI()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isRLSD()`** — L1892 — `public bool isRLSD()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setDTR()`** — L1898 — `public void setDTR(bool v)`
  Sets dtr.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialErrorReceived()`** — L1904 — `void SerialErrorReceived(object source, SerialErrorReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialPinChanged()`** — L1958 — `void SerialPinChanged(object source, SerialPinChangedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialReceivedData()`** — L2028 — `void SerialReceivedData(object source, SerialDataReceivedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SDRSerialPortII.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
