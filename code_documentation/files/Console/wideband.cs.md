# `Console/wideband.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Wideband (full 0–61 MHz) spectrum display and its data acquisition from the radio's wideband sample stream.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/database.cs` (calls ×2)
  - `Console/wbDisplay.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `wideband` (type, L14)

- **`.wideband_Resize()`** — L36 — `private void wideband_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `wideband` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.wideband_FormClosing()`** — L49 — `private void wideband_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `wideband` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.contextMenuStripWideBand_Opening()`** — L59 — `private void contextMenuStripWideBand_Opening(object sender, CancelEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToolStripMenuItem_MouseEnter()`** — L72 — `private void ToolStripMenuItem_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ToolStripMenuItem_MouseLeave()`** — L83 — `private void ToolStripMenuItem_MouseLeave(object sender, EventArgs e)`
  WinForms event handler: runs when `ToolStripMenuItem` is left by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ContextMenuStrip_Closing()`** — L95 — `private void ContextMenuStrip_Closing(object sender, ToolStripDropDownClosingEventArgs e)`
  WinForms event handler: runs when `ContextMenuStrip` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.wbAvgtoolStripMenuItem_Click()`** — L107 — `private void wbAvgtoolStripMenuItem_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `wbAvgtoolStripMenuItem` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.wbdisplay_Resize()`** — L112 — `private void wbdisplay_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `wbdisplay` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.wbUpdatetoolStripComboBox_SelectedIndexChanged()`** — L117 — `private void wbUpdatetoolStripComboBox_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `wbUpdatetoolStripComboBox` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.wbFrameSizetoolStripComboBox_SelectedIndexChanged()`** — L123 — `private void wbFrameSizetoolStripComboBox_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `wbFrameSizetoolStripComboBox` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SaveWideBand()`** — L129 — `public void SaveWideBand()`
  Saves wide band.
  Called by: `.wideband_FormClosing()` (same file)
- **`.GetWideBand()`** — L147 — `public void GetWideBand()`
  Returns wide band.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/wideband.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
