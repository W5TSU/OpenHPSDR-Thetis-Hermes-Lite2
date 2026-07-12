# `Console/Memory/MemoryForm.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Memory channel list UI and its record/list model (frequency, mode, filter, tones per memory).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×7)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

### Types

#### `MemoryForm` (type, L45)

- `.dataGridView1_DragEnter()` — L290
- `.dataGridView1_DragDrop()` — L328
- `.dataGridView1_CellMouseDown()` — L339
- `.dataGridView1_CellClick()` — L376
- `.dataGridView1_CellValidating()` — L392
- `.MemoryRecordAdd_DragEnter()` — L427
- `.MemoryRecordAdd_DragDrop()` — L467
- `.MemoryRecordAdd_Click()` — L502
- `.btnMemoryRecordCopy_Click()` — L548
- `.btnMemoryRecordDelete_Click()` — L564
- `.btnSelect_Click()` — L605
- `.MemoryForm_FormClosing()` — L642
- `.ReadURL()` — L675
- `.DoesDragDropDataContainUrl1()` — L718
- `.textBox1_TextChanged()` — L728
- `.MemoryForm_Load()` — L733
- `.ScheduleDurationTime_ValueChanged()` — L743
- `.ScheduleRepeat_CheckedChanged()` — L762
- `.ScheduleRepeatm_CheckedChanged()` — L792
- `.ScheduleRecord_CheckedChanged()` — L825
- `.ScheduleOn_CheckedChanged()` — L841
- `.ScheduleStartDate_ValueChanged()` — L869
- `.ScheduleUpdate()` — L930
- `.chkAlwaysOnTop_CheckedChanged()` — L985
- `.SCHEDULER()` — L1010
- `.buttonTS1_Click()` — L1387
- `.WaveToMP3()` — L1422
- `.TOMP3()` — L1432
- `.ConvertWavStreamToMp3File()` — L1471

#### `AutoClosingMessageBox` (type, L1493)

- `.Show()` — L1506
- `.OnTimerElapsed()` — L1510
- `.FindWindow()` — L1518
- `.SendMessage()` — L1520

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/MemoryForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
