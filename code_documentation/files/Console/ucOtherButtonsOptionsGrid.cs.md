# `Console/ucOtherButtonsOptionsGrid.cs`

**Functional area:** [13. Andromeda control surface](../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** User-programmable macro buttons and their configuration grid.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×11, references ×8)
  - `Console/console.cs` (references ×6)
  - `Console/frmMacroButtonConfig.cs` (references ×2)
- Uses (outgoing references to other files):
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
- Most-referenced symbols from other files: `.BitFromID()` (×7), `.BitToID()` (×2), `.BitToIcon()` (×1), `.BitToText()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `OtherButtonId` (type, L54)

_No extracted members._

#### `OtherButtonMacroSettings` (type, L259)

- **`.deep_clone()`** — L306 — `private static T deep_clone<T>(T obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `OB_ButtonState` (type, L262)

_No extracted members._

#### `OtherButtonIdHelpers` (type, L521)

- **`.OtherButtonIDToText()`** — L776 — `public static string OtherButtonIDToText(OtherButtonId id)`
  Called by: `.BitToText()` (same file), `.initialise_checkboxes()` (same file)
- **`.OtherButtonIDToIconOn()`** — L781 — `public static string OtherButtonIDToIconOn(OtherButtonId id)`
  Called by: `.BitToIcon()` (same file)
- **`.OtherButtonIDToIconOff()`** — L786 — `public static string OtherButtonIDToIconOff(OtherButtonId id)`
  Called by: `.BitToIcon()` (same file)
- **`.BitToID()`** — L793 — `public static OtherButtonId BitToID(int bit_group, int bit_number)`
  Called by: `.BitToText()` (same file), `.BitToIcon()` (same file), `.setupButtons()` (`Console/MeterManager.cs`), `.MouseUp()` (`Console/MeterManager.cs`)
- **`.BitFromID()`** — L799 — `public static (int bit_group, int bit) BitFromID(OtherButtonId id)`
  Called by: `.onPlayingChanged()` (`Console/MeterManager.cs`), `.onRecordingChanged()` (`Console/MeterManager.cs`), `.OnContainerVisible()` (`Console/MeterManager.cs`), `.OnCatState()` (`Console/MeterManager.cs`), `.ContainerHiddenByMacro()` (`Console/MeterManager.cs`), `.updateOn()` (`Console/MeterManager.cs`) — and 1 more
- **`.BitToText()`** — L805 — `public static string BitToText(int bit_group, int bit_number)`
  Called by: `.setupButtons()` (`Console/MeterManager.cs`)
- **`.BitToIcon()`** — L810 — `public static (string, string) BitToIcon(int bit_group, int bit_number)`
  Called by: `.setupButtons()` (`Console/MeterManager.cs`)
- **`.OtherButtonIDToTooltip()`** — L815 — `public static string OtherButtonIDToTooltip(OtherButtonId id)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkImplemented()`** — L820 — `private static void checkImplemented()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `ucOtherButtonsOptionsGrid` (type, L842)

- **`.initialise_checkboxes()`** — L913 — `private void initialise_checkboxes()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.checkbox_checked_changed()`** — L1103 — `private void checkbox_checked_changed(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.button_clicked()`** — L1108 — `private void button_clicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetBitfield()`** — L1119 — `public int GetBitfield(int bit_group)`
  Returns bitfield.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBitfield()`** — L1131 — `public void SetBitfield(int bit_group, int value)`
  Sets bitfield.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCheckedCount()`** — L1147 — `public int GetCheckedCount(int bit_group)`
  Returns checked count.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMacroSettings()`** — L1159 — `public OtherButtonMacroSettings GetMacroSettings(int macro)`
  Returns macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMacroSettings()`** — L1165 — `public void SetMacroSettings(int macro, OtherButtonMacroSettings settings)`
  Sets macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `MacroButtonEventArgs` (type, L844)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucOtherButtonsOptionsGrid.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
