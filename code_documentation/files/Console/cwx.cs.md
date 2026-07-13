# `Console/cwx.cs`

**Functional area:** [11. CW keying](../../CODE_OUTLINE.md#11-cw-keying)

**Role:** CWX memory keyer window — canned messages, beacon loops, and keyboard CW.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×8, references ×1)
- Uses (outgoing references to other files):
  - `Console/ringbuffer.cs` (calls ×5, references ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Invoke/buttonts.cs` (references ×1)
  - `Console/Invoke/checkboxts.cs` (references ×1)
  - `Console/Invoke/comboboxts.cs` (references ×1)
  - `Console/Invoke/labelts.cs` (references ×1)
  - `Console/Invoke/numericupdownts.cs` (references ×1)
  - `Console/Invoke/textboxts.cs` (references ×1)
  - `Console/common.cs` (calls ×1)
- Most-referenced symbols from other files: `.StopEverything()` (×3), `.CWXStop()` (×1), `.Show()` (×1), `.KeyAction()` (×1), `.PressFNkey()` (×1), `.StopAction()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CWX` (type, L73)

- **`.timeGetDevCaps()`** — L213 — `[DllImport("winmm.dll")] private static extern int timeGetDevCaps(ref TimerCaps caps, int sizeOfTimerCaps)`
  Gets timer capabilities.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.timeSetEvent()`** — L218 — `[DllImport("winmm.dll")] private static extern int timeSetEvent(int delay, int resolution, TimeProc proc, int user, int mode)`
  Creates and starts the timer.
  Called by: `.setup_timer()` (same file)
- **`.timeKillEvent()`** — L223 — `[DllImport("winmm.dll")] private static extern int timeKillEvent(int id)`
  Stops and destroys the timer.
  Called by: `.setup_timer()` (same file), `.Dispose()` (same file), `.CWX_Closing()` (same file)
- **`.setup_timer()`** — L261 — `private void setup_timer()`
  Called by: `.Show()` (same file), `.udWPM_ValueChanged()` (same file)
- **`.setptt()`** — L282 — `private void setptt(bool state)`
  Called by: `.quitshut()` (same file), `.KeyAction()` (same file), `.process_element()` (same file)
- **`.setkey()`** — L305 — `private void setkey(bool state)`
  Called by: `.quitshut()` (same file), `.KeyAction()` (same file), `.process_element()` (same file)
- **`.quitshut()`** — L317 — `private void quitshut()`
  Called by: `.AbortSending()` (same file), `.KeyAction()` (same file), `.CWX_Closing()` (same file), `.editit()` (same file), `.process_element()` (same file), `.CWX_FormClosing()` (same file)
- **`.clear_fifo()`** — L328 — `private void clear_fifo()`
  Called by: `.quitshut()` (same file), `.loadmsg()` (same file)
- **`.push_fifo()`** — L336 — `private void push_fifo(byte data)`
  Called by: `.process_element()` (same file), `.queue_start()` (same file), `.loadchar()` (same file), `.loadmsg()` (same file)
- **`.pop_fifo()`** — L347 — `private byte pop_fifo()`
  Called by: `.process_element()` (same file)
- **`.clear_fifo2()`** — L368 — `private void clear_fifo2()`
  Called by: `.quitshut()` (same file), `.keyboardFifo()` (same file)
- **`.push_fifo2()`** — L376 — `private void push_fifo2(byte data)`
  Called by: `.keyboardDisplay()` (same file)
- **`.pop_fifo2()`** — L388 — `private byte pop_fifo2()`
  Called by: `.keyboardFifo()` (same file)
- **`.wpmrate()`** — L411 — `private int wpmrate()`
  Called by: `.setup_timer()` (same file)
- **`.help()`** — L415 — `private void help()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.notesButton_Click()`** — L460 — `private void notesButton_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `notesButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.build_mbits2()`** — L473 — `private void build_mbits2()`
  Called by: `.insert_and_reload()` (same file)
- **`.load_alpha()`** — L533 — `private void load_alpha()`
  Called by: `.insert_and_reload()` (same file)
- **`.RemoteMessage()`** — L657 — `public string RemoteMessage(byte[] msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendBufferMessage()`** — L672 — `private void SendBufferMessage()`
  Sends buffer message.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXStop()`** — L732 — `public void CWXStop()`
  Called by: `.Console_KeyDown()` (`Console/console.cs`)
- **`.AbortSending()`** — L740 — `public void AbortSending()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.startThreads()`** — L833 — `private void startThreads()`
  Called by: `.Show()` (same file)
- **`.Dispose()`** — L860 — `protected override void Dispose(bool disposing)`
  Clean up any resources being used.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeComponent()`** — L889 — `private void InitializeComponent()`
  Required method for Designer support - do not modify the contents of this method with the code editor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.expandButton_Click()`** — L1589 — `private void expandButton_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `expandButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.keyboardButton_Leave()`** — L1610 — `private void keyboardButton_Leave(object sender, System.EventArgs e)`
  WinForms event handler: runs when `keyboardButton` is left.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.keyboardButton_Enter()`** — L1619 — `private void keyboardButton_Enter(object sender, System.EventArgs e)`
  WinForms event handler: runs when `keyboardButton` is entered.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CWX_KeyUp_1()`** — L1629 — `private void CWX_KeyUp_1(object sender, System.Windows.Forms.KeyEventArgs e)`
  this guy checks for the release of the Alt key
  Called by: `.onGlobalKeyUp()` (same file)
- **`.CWX_KeyDown_1()`** — L1642 — `private void CWX_KeyDown_1(object sender, System.Windows.Forms.KeyEventArgs e)`
  the Esc, F1, F2, and Alt 1 thru Alt 9 are handled anywhere on the form the Alt key press is detected here and altkey set to true
  Called by: `.onGlobalKeyDown()` (same file)
- **`.keyboardButton_KeyPress()`** — L1689 — `private void keyboardButton_KeyPress(object sender, System.Windows.Forms.KeyPressEventArgs e)`
  WinForms event handler: runs when `keyboardButton` receives a key press.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.KeyAction()`** — L1696 — `public void KeyAction()`
  Called by: `.keyButton_Click()` (same file), `.DoOtherButtonAction()` (`Console/console.cs`)
- **`.keyButton_Click()`** — L1729 — `private void keyButton_Click(object sender, System.EventArgs e)`
  process the 'Key' button which start transmitter with key down
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.checkPTT()`** — L1733 — `private bool checkPTT(bool bShowWarning = true)`
  Called by: `.KeyAction()` (same file), `.process_element()` (same file)
- **`.CWX_Load()`** — L1749 — `private void CWX_Load(object sender, System.EventArgs e)`
  WinForms event handler: runs when `CWX` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CWX_Closing()`** — L1758 — `private void CWX_Closing(object sender, System.ComponentModel.CancelEventArgs e)`
  WinForms event handler: runs when `CWX` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Show()`** — L1788 — `public new void Show()`
  Called by: `.cWXToolStripMenuItem_Click()` (`Console/console.cs`)
- **`.TimerPeriodicEventCallback()`** — L1809 — `private void TimerPeriodicEventCallback(int id, int msg, int user, int param1, int param2)`
  Callback method called by the Win32 multimedia timer when a timer periodic event occurs.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PressFNkey()`** — L1814 — `public void PressFNkey(int fn_number)`
  Called by: `.DoOtherButtonAction()` (`Console/console.cs`)
- **`.s1_Click()`** — L1820 — `private void s1_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s1` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s2_Click()`** — L1826 — `private void s2_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s2` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s3_Click()`** — L1832 — `private void s3_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s3` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s4_Click()`** — L1838 — `private void s4_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s4` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s5_Click()`** — L1844 — `private void s5_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s5` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s6_Click()`** — L1850 — `private void s6_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s6` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s7_Click()`** — L1856 — `private void s7_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s7` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s8_Click()`** — L1861 — `private void s8_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s8` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s9_Click()`** — L1866 — `private void s9_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `s9` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.cbMorse_MouseDown()`** — L1871 — `private void cbMorse_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `cbMorse` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s1_MouseDown()`** — L1877 — `private void s1_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s1` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s2_MouseDown()`** — L1881 — `private void s2_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s2` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s3_MouseDown()`** — L1885 — `private void s3_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s3` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s4_MouseDown()`** — L1889 — `private void s4_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s4` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s5_MouseDown()`** — L1893 — `private void s5_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s5` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s6_MouseDown()`** — L1897 — `private void s6_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s6` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s7_MouseDown()`** — L1901 — `private void s7_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s7` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s8_MouseDown()`** — L1905 — `private void s8_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s8` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.s9_MouseDown()`** — L1909 — `private void s9_MouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `s9` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.StopAction()`** — L1914 — `public void StopAction()`
  Stops action.
  Called by: `.stopButton_Click()` (same file), `.DoOtherButtonAction()` (`Console/console.cs`)
- **`.stopActionCore()`** — L1919 — `private void stopActionCore()`
  Called by: `.StopAction()` (same file), `.StopEverything()` (same file)
- **`.stopButton_Click()`** — L1927 — `private void stopButton_Click(object sender, System.EventArgs e)`
  stop button clicked
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udWPM_ValueChanged()`** — L1931 — `private void udWPM_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udWPM` value changes.
  Called by: `.udWPM_LostFocus()` (same file)
- **`.udWPM_LostFocus()`** — L1939 — `private void udWPM_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udWPM` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDelay_ValueChanged()`** — L1944 — `private void udDelay_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDelay` value changes.
  Called by: `.udDelay_LostFocus()` (same file)
- **`.udDelay_LostFocus()`** — L1949 — `private void udDelay_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDelay` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udDrop_ValueChanged()`** — L1954 — `private void udDrop_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udDrop` value changes.
  Called by: `.udDrop_LostFocus()` (same file)
- **`.udDrop_LostFocus()`** — L1958 — `private void udDrop_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udDrop` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.udPtt_ValueChanged()`** — L1963 — `private void udPtt_ValueChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `udPtt` value changes.
  Called by: `.udPtt_LostFocus()` (same file)
- **`.udPtt_LostFocus()`** — L1971 — `private void udPtt_LostFocus(object sender, EventArgs e)`
  WinForms event handler: runs when `udPtt` loses focus.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CWX_MouseMove()`** — L1976 — `private void CWX_MouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  WinForms event handler: runs when `CWX` receives mouse movement.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.CWX_Paint()`** — L1981 — `private void CWX_Paint(object sender, System.Windows.Forms.PaintEventArgs e)`
  WinForms event handler: runs when `CWX` repaints.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkPause_CheckedChanged()`** — L1986 — `private void chkPause_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkPause` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkAlwaysOnTop_CheckedChanged()`** — L1991 — `private void chkAlwaysOnTop_CheckedChanged(object sender, System.EventArgs e)`
  WinForms event handler: runs when `chkAlwaysOnTop` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clear_keys()`** — L2011 — `private void clear_keys()`
  Called by: `.clear_show()` (same file)
- **`.show_keys()`** — L2025 — `private void show_keys(Graphics formGraphics = null)`
  Called by: `.CWX_Paint()` (same file), `.clear_show()` (same file), `.keyboardDisplay()` (same file), `.process_key()` (same file), `.backspace()` (same file), `.msg2keys()` (same file)
- **`.clearButton_Click()`** — L2089 — `private void clearButton_Click(object sender, System.EventArgs e)`
  WinForms event handler: runs when `clearButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.clear_show()`** — L2093 — `private void clear_show()`
  Called by: `.AbortSending()` (same file), `.CWX_KeyDown_1()` (same file), `.stopActionCore()` (same file), `.clearButton_Click()` (same file), `.editit()` (same file), `.CWX_FormClosing()` (same file)
- **`.editit()`** — L2106 — `private void editit()`
  Called by: `.cbMorse_MouseDown()` (same file)
- **`.insert_and_reload()`** — L2140 — `private void insert_and_reload(string s)`
  Called by: `.editit()` (same file)
- **`.write_a2m2()`** — L2171 — `private void write_a2m2()`
  Called by: `.insert_and_reload()` (same file)
- **`.process_element()`** — L2199 — `private void process_element()`
  Called by: `.TimerPeriodicEventCallback()` (same file)
- **`.keyboardFifo()`** — L2286 — `private void keyboardFifo()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.keyboardDisplay()`** — L2324 — `private void keyboardDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.queue_start()`** — L2361 — `private void queue_start(int qmsg)`
  Called by: `.CWX_KeyDown_1()` (same file), `.PressFNkey()` (same file), `.s1_Click()` (same file), `.s2_Click()` (same file), `.s3_Click()` (same file), `.s4_Click()` (same file) — and 5 more
- **`.loadchar()`** — L2396 — `private void loadchar(char cc)`
  Called by: `.RemoteMessage()` (same file), `.SendBufferMessage()` (same file), `.keyboardFifo()` (same file)
- **`.loadmsg()`** — L2429 — `private void loadmsg(string t)`
  Called by: `.process_element()` (same file), `.queue_start()` (same file)
- **`.process_key()`** — L2505 — `private void process_key(char key)`
  Called by: `.keyboardButton_KeyPress()` (same file)
- **`.CbMorse_SelectedIndexChanged()`** — L2524 — `private void CbMorse_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `CbMorse` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.insert_key()`** — L2529 — `private void insert_key(char key)`
  Called by: `.process_key()` (same file), `.msg2keys()` (same file)
- **`.CWX_FormClosing()`** — L2549 — `private void CWX_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `CWX` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.backspace()`** — L2564 — `private void backspace()`
  Called by: `.process_key()` (same file)
- **`.msg2keys()`** — L2580 — `private void msg2keys(int nmsg)`
  Called by: `.CWX_KeyDown_1()` (same file), `.s1_MouseDown()` (same file), `.s2_MouseDown()` (same file), `.s3_MouseDown()` (same file), `.s4_MouseDown()` (same file), `.s5_MouseDown()` (same file) — and 4 more
- **`.StopEverything()`** — L2617 — `public void StopEverything(bool bPowerState = false)`
  Stops everything.
  Called by: `.chkPower_CheckedChanged()` (`Console/console.cs`), `.Console_Closing()` (`Console/console.cs`), `.chkMOX_CheckedChanged2()` (`Console/console.cs`)
- **`.onGlobalKeyDown()`** — L2633 — `private void onGlobalKeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onGlobalKeyUp()`** — L2642 — `private void onGlobalKeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `TimerMode` (type, L197)

_No extracted members._

#### `TimerCaps` (type, L204)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cwx.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
