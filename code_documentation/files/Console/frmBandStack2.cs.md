# `Console/frmBandStack2.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Per-band frequency stack (last-used frequencies per band) and its popup window.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmFilterManager.Designer.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/clsBandStackManager.cs` (calls ×4, references ×3)
  - `Console/common.cs` (calls ×4)
  - `Console/Andromeda/Andromeda.cs` (references ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmBandStack2` (type, L56)

- **`.InitForm()`** — L94 — `public void InitForm(Console console)`
  Inits form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMox()`** — L118 — `private void OnMox(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveDelegates()`** — L122 — `public void RemoveDelegates()`
  Removes delegates.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitBandStackFilter()`** — L129 — `public void InitBandStackFilter(BandStackFilter bsf, bool select = true)`
  Inits band stack filter.
  Called by: `.Show()` (same file)
- **`.UpdateSelected()`** — L180 — `public void UpdateSelected()`
  Updates selected.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupSelectedButtons()`** — L206 — `private void setupSelectedButtons()`
  Called by: `.InitBandStackFilter()` (same file), `.UpdateSelected()` (same file)
- **`.setupRadioButtons()`** — L233 — `private void setupRadioButtons(bool bCheck = true)`
  Called by: `.InitBandStackFilter()` (same file), `.radioLastUsedEntry_CheckedChanged()` (same file), `.radioSpecific_CheckedChanged()` (same file), `.radioLastUsed_CheckedChanged()` (same file)
- **`.btnOptions_Click()`** — L267 — `private void btnOptions_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOptions` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radioLastUsedEntry_CheckedChanged()`** — L281 — `private void radioLastUsedEntry_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radioLastUsedEntry` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radioSpecific_CheckedChanged()`** — L290 — `private void radioSpecific_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radioSpecific` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radioLastUsed_CheckedChanged()`** — L299 — `private void radioLastUsed_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radioLastUsed` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSetSpecific_Click()`** — L308 — `private void btnSetSpecific_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSetSpecific` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnLockSelected_Click()`** — L322 — `private void btnLockSelected_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnLockSelected` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnDeleteSelected_Click()`** — L340 — `private void btnDeleteSelected_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnDeleteSelected` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.bandStackListBox_SelectedIndexChanged()`** — L351 — `private void bandStackListBox_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `bandStackListBox` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnAddStackEntry_Click()`** — L374 — `private void btnAddStackEntry_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnAddStackEntry` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlwaysOnTop_CheckedChanged()`** — L379 — `private void chkAlwaysOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlwaysOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.HideClose()`** — L384 — `public void HideClose()`
  Hides close.
  Called by: `.frmBandStack2_FormClosing()` (same file)
- **`.frmBandStack2_FormClosing()`** — L396 — `private void frmBandStack2_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmBandStack2` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Show()`** — L402 — `public new void Show(bool is_popup = false, Point? popup_location = null, bool on_top = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Store()`** — L424 — `public void Store()`
  Called by: `.HideClose()` (same file)
- **`.btnUpdateEntry_Click()`** — L430 — `private void btnUpdateEntry_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnUpdateEntry` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkIgnoreDuplicates_CheckedChanged()`** — L441 — `private void chkIgnoreDuplicates_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkIgnoreDuplicates` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHideOnSelect_CheckedChanged()`** — L446 — `private void chkHideOnSelect_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHideOnSelect` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkShowInSpectrum_CheckedChanged()`** — L451 — `private void chkShowInSpectrum_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkShowInSpectrum` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.frmBandStack2_LocationChanged()`** — L463 — `private void frmBandStack2_LocationChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `BandStackListBox` (type, L471)

- **`.AddItem()`** — L537 — `public int AddItem(BandStackEntry bse)`
  Adds item.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearItems()`** — L547 — `public void ClearItems()`
  Clears items.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFontChanged()`** — L555 — `protected override void OnFontChanged(EventArgs e)`
  Handles/raises the font changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSelectedIndexChanged()`** — L565 — `protected override void OnSelectedIndexChanged(EventArgs e)`
  Handles/raises the selected index changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDrawItem()`** — L571 — `protected override void OnDrawItem(DrawItemEventArgs e)`
  Handles/raises the draw item event.
  Called by: `.OnPaint()` (same file)
- **`.OnMouseEnter()`** — L642 — `protected override void OnMouseEnter(EventArgs e)`
  Handles/raises the mouse enter event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseLeave()`** — L646 — `protected override void OnMouseLeave(EventArgs e)`
  Handles/raises the mouse leave event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L653 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L669 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmBandStack2.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
