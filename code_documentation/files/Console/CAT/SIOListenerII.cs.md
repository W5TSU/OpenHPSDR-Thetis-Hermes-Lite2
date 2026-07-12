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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`ParseString()`** — L199 — `private void ParseString(byte[] rxdata, uint count)`
  segment incoming string into CAT commands ... handle leftovers from when we read a parial
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

### Types

#### `SIOListenerII` (type, L31)

- **`.enableCAT()`** — L67 — `public void enableCAT()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT()`** — L149 — `public void disableCAT()`
  typically called when the end user has disabled CAT control through a UI element ... this closes the serial port and neutralized the listeners we have in place
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialize()`** — L186 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L256 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L264 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRXEventHandler()`** — L274 — `private void SerialRXEventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO2ListenerII` (type, L324)

- **`.enableCAT2()`** — L361 — `public void enableCAT2()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT2()`** — L435 — `public void disableCAT2()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.Initialize()`** — L472 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L485 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L493 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX2EventHandler()`** — L503 — `private void SerialRX2EventHandler(object sender, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO3ListenerII` (type, L556)

- **`.enableCAT3()`** — L593 — `public void enableCAT3()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT3()`** — L667 — `public void disableCAT3()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.Initialize()`** — L704 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L717 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L725 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX3EventHandler()`** — L735 — `private void SerialRX3EventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO4ListenerII` (type, L784)

- **`.enableCAT4()`** — L820 — `public void enableCAT4()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT4()`** — L894 — `public void disableCAT4()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.Initialize()`** — L931 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L944 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L952 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX4EventHandler()`** — L962 — `private void SerialRX4EventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO5ListenerII` (type, L1013)

- **`.enableCAT5()`** — L1050 — `public void enableCAT5()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT5()`** — L1067 — `public void disableCAT5()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialize()`** — L1104 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L1117 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L1125 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX5EventHandler()`** — L1135 — `private void SerialRX5EventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO6ListenerII` (type, L1187)

- **`.enableCAT6()`** — L1224 — `public void enableCAT6()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT6()`** — L1241 — `public void disableCAT6()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialize()`** — L1278 — `private void Initialize()`
  Called when the console is activated for the first time.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.console_Closing()`** — L1291 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L1299 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX6EventHandler()`** — L1309 — `private void SerialRX6EventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SIO7ListenerII` (type, L1361)

- **`.enableCAT7()`** — L1398 — `public void enableCAT7()`
  Called by: `.console_Activated()` (same file)
- **`.disableCAT7()`** — L1415 — `public void disableCAT7()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialize()`** — L1452 — `private void Initialize()`
  Called when the console is activated for the first time.
  Called by: `.enableCAT()` (same file), `.enableCAT2()` (same file), `.enableCAT3()` (same file), `.enableCAT4()` (same file), `.enableCAT5()` (same file), `.enableCAT6()` (same file) — and 1 more
- **`.console_Closing()`** — L1465 — `private void console_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `console` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.console_Activated()`** — L1473 — `private void console_Activated(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SerialRX7EventHandler()`** — L1483 — `private void SerialRX7EventHandler(object source, SerialRXEvent e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/CAT/SIOListenerII.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
