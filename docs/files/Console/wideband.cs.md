# `Console/wideband.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** Wideband (full 0–61 MHz) spectrum display and its data acquisition from the radio's wideband sample stream.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/database.cs` (calls ×2)
  - `Console/wbDisplay.cs` (references ×1)

## Outline

### Types

#### `wideband` (type, L14)

- `.wideband_Resize()` — L36
- `.wideband_FormClosing()` — L49
- `.contextMenuStripWideBand_Opening()` — L59
- `.ToolStripMenuItem_MouseEnter()` — L72
- `.ToolStripMenuItem_MouseLeave()` — L83
- `.ContextMenuStrip_Closing()` — L95
- `.wbAvgtoolStripMenuItem_Click()` — L107
- `.wbdisplay_Resize()` — L112
- `.wbUpdatetoolStripComboBox_SelectedIndexChanged()` — L117
- `.wbFrameSizetoolStripComboBox_SelectedIndexChanged()` — L123
- `.SaveWideBand()` — L129
- `.GetWideBand()` — L147

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/wideband.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
