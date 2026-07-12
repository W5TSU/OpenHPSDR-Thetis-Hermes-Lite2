# `Console/RAForm.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Radio-astronomy data collection utility (niche feature retained from upstream).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `RAForm` (type, L55)

- `.Dispose()` — L79
- `.RArecordCheckBox_CheckedChanged()` — L125
- `.construct_header()` — L190
- `.write_line_to_file()` — L204
- `.numericUpDownTS1_ValueChanged()` — L212
- `.RA_timer_Tick()` — L218
- `.numericUpDownTS2_ValueChanged()` — L282
- `.picRAGraph_Paint()` — L288
- `.button_dBm_CheckedChanged()` — L648
- `.button_linear_CheckedChanged()` — L670
- `.picRAGraph_MouseMove()` — L676
- `.RAForm_MouseMove()` — L687
- `.button_readFile_Click()` — L693
- `.MyTryParse()` — L765
- `.button_writeFile_Click()` — L779
- `.manual_ymax_ValueChanged()` — L817
- `.manual_ymin_ValueChanged()` — L824
- `.manual_xmax_ValueChanged()` — L832
- `.manual_xmin_ValueChanged()` — L839
- `.StrToByteArray()` — L846
- `.RAForm_Load()` — L852

#### `Prompt` (type, L858)

- `.ShowDialog()` — L860

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/RAForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
