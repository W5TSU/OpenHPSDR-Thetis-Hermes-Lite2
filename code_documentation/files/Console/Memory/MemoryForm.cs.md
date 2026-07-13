# `Console/Memory/MemoryForm.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Memory channel list UI and its record/list model (frequency, mode, filter, tones per memory).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×7)
  - `Console/Andromeda/Andromeda.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `MemoryForm` (type, L45)

- **`.dataGridView1_DragEnter()`** — L290 — `private void dataGridView1_DragEnter(object sender, DragEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dataGridView1_DragDrop()`** — L328 — `private void dataGridView1_DragDrop(object sender, DragEventArgs e)`
  ke9ns add ONCE DRAGENTER VALIDATES YOUR URL, YOU RELEASE YOUR MOUSE OVER THE WINDOW AND COMMENT FIELD IS UPDATED WITH THE URL
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dataGridView1_CellMouseDown()`** — L339 — `private void dataGridView1_CellMouseDown(object sender, DataGridViewCellMouseEventArgs e)`
  ke9ns add USED TO OPEN WEB BROWSER if comment field has URL
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dataGridView1_CellClick()`** — L376 — `private void dataGridView1_CellClick(object sender, DataGridViewCellEventArgs e)`
  ke9ns add COMES HERE AFTER YOU CLICK ON A FIELD BOX TO DETERMINE WHICH ROW YOU ARE WORKING IN
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dataGridView1_CellValidating()`** — L392 — `void dataGridView1_CellValidating(object sender, DataGridViewCellValidatingEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MemoryRecordAdd_DragEnter()`** — L427 — `private void MemoryRecordAdd_DragEnter(object sender, DragEventArgs e)`
  ke9ns add YOU DRAG YOUR URL ALONG WITH YOUR MOUSE OVER THE ADD BUTTON
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MemoryRecordAdd_DragDrop()`** — L467 — `private void MemoryRecordAdd_DragDrop(object sender, DragEventArgs e)`
  Ke9ns add YOUR URL (after being VERIFIED) YOU LET GO THE LEFT MOUSE BUTTON TO DROP ONTO THE ADD BUTTON
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MemoryRecordAdd_Click()`** — L502 — `public void MemoryRecordAdd_Click(object sender, EventArgs e)`
  Add a new Memory entry based on the current console settings.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMemoryRecordCopy_Click()`** — L548 — `private void btnMemoryRecordCopy_Click(object sender, EventArgs e)`
  Copy an existing row into a new one.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnMemoryRecordDelete_Click()`** — L564 — `private void btnMemoryRecordDelete_Click(object sender, EventArgs e)`
  Delete the current row (after confirmation).
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSelect_Click()`** — L605 — `private void btnSelect_Click(object sender, EventArgs e)`
  Makes the selected row active -- sends it to console
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MemoryForm_FormClosing()`** — L642 — `private void MemoryForm_FormClosing(object sender, FormClosingEventArgs e)`
  Don't actually close the form, just hide it and save the position/size.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ReadURL()`** — L675 — `private string ReadURL(IDataObject data)`
  ke9ns look for URL or file
  Called by: `.dataGridView1_DragEnter()` (same file), `.MemoryRecordAdd_DragEnter()` (same file)
- **`.DoesDragDropDataContainUrl1()`** — L718 — `private static bool DoesDragDropDataContainUrl1(IDataObject data, string urlDataFormatName)`
  Called by: `.ReadURL()` (same file)
- **`.textBox1_TextChanged()`** — L728 — `private void textBox1_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `textBox1` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.MemoryForm_Load()`** — L733 — `private void MemoryForm_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `MemoryForm` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleDurationTime_ValueChanged()`** — L743 — `private void ScheduleDurationTime_ValueChanged(object sender, EventArgs e)`
  ke9ns add Schedule duration of recording if enabled
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleRepeat_CheckedChanged()`** — L762 — `private void ScheduleRepeat_CheckedChanged(object sender, EventArgs e)`
  ke9ns add Schedule weekly ON/OFF
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleRepeatm_CheckedChanged()`** — L792 — `private void ScheduleRepeatm_CheckedChanged(object sender, EventArgs e)`
  ke9ns add Schedule monthly ON/OFF
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleRecord_CheckedChanged()`** — L825 — `private void ScheduleRecord_CheckedChanged(object sender, EventArgs e)`
  ke9ns add Schedule Record ON/OFF
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleOn_CheckedChanged()`** — L841 — `private void ScheduleOn_CheckedChanged(object sender, EventArgs e)`
  ke9ns add Schedule ON/OFF (NOT USED AT THIS TIME)
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleStartDate_ValueChanged()`** — L869 — `private void ScheduleStartDate_ValueChanged(object sender, EventArgs e)`
  ke9ns add DATE for datetime for schedule
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ScheduleUpdate()`** — L930 — `public void ScheduleUpdate()`
  ke9ns add update the boxes at the bottom of the memory screen
  Called by: `.dataGridView1_CellMouseDown()` (same file), `.dataGridView1_CellClick()` (same file), `.dataGridView1_CellValidating()` (same file), `.MemoryRecordAdd_DragEnter()` (same file), `.MemoryRecordAdd_DragDrop()` (same file), `.MemoryRecordAdd_Click()` (same file)
- **`.chkAlwaysOnTop_CheckedChanged()`** — L985 — `private void chkAlwaysOnTop_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkAlwaysOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SCHEDULER()`** — L1010 — `private void SCHEDULER()`
  ke9ns add Thread routine (checks the scheduler)
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.buttonTS1_Click()`** — L1387 — `private void buttonTS1_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `buttonTS1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.WaveToMP3()`** — L1422 — `public static void WaveToMP3(string waveFileName, string mp3FileName, int bitRate = 128)`
  ke9ns add NOT USED AT THE MOMENT
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TOMP3()`** — L1432 — `public void TOMP3()`
  ke9ns add Thread
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ConvertWavStreamToMp3File()`** — L1471 — `public static void ConvertWavStreamToMp3File(ref MemoryStream ms, string savetofilename)`
  MemoryStream ms = new MemoryStream(); ke9ns add
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `AutoClosingMessageBox` (type, L1493)

- **`.Show()`** — L1506 — `public static void Show(string text, string caption, int timeout)`
  Called by: `.SCHEDULER()` (same file)
- **`.OnTimerElapsed()`** — L1510 — `void OnTimerElapsed(object state)`
  Handles/raises the timer elapsed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindWindow()`** — L1518 — `[System.Runtime.InteropServices.DllImport("user32.dll", SetLastError = true)] static extern IntPtr FindWindow(string lpClassName, string lpWindowName)`
  Finds window.
  Called by: `.OnTimerElapsed()` (same file)
- **`.SendMessage()`** — L1520 — `[System.Runtime.InteropServices.DllImport("user32.dll", CharSet = System.Runtime.InteropServices.CharSet.Auto)] static extern IntPtr SendMessage(IntPtr hWnd, UInt32 Msg, IntPtr wPa`
  Sends message.
  Called by: `.OnTimerElapsed()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Memory/MemoryForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
