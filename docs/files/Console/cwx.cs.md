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

### Types

#### `CWX` (type, L73)

- `.timeGetDevCaps()` — L213
- `.timeSetEvent()` — L218
- `.timeKillEvent()` — L223
- `.setup_timer()` — L261
- `.setptt()` — L282
- `.setkey()` — L305
- `.quitshut()` — L317
- `.clear_fifo()` — L328
- `.push_fifo()` — L336
- `.pop_fifo()` — L347
- `.clear_fifo2()` — L368
- `.push_fifo2()` — L376
- `.pop_fifo2()` — L388
- `.wpmrate()` — L411
- `.help()` — L415
- `.notesButton_Click()` — L460
- `.build_mbits2()` — L473
- `.load_alpha()` — L533
- `.RemoteMessage()` — L657
- `.SendBufferMessage()` — L672
- `.CWXStop()` — L732
- `.AbortSending()` — L740
- `.startThreads()` — L833
- `.Dispose()` — L860
- `.InitializeComponent()` — L889
- `.expandButton_Click()` — L1589
- `.keyboardButton_Leave()` — L1610
- `.keyboardButton_Enter()` — L1619
- `.CWX_KeyUp_1()` — L1629
- `.CWX_KeyDown_1()` — L1642
- `.keyboardButton_KeyPress()` — L1689
- `.KeyAction()` — L1696
- `.keyButton_Click()` — L1729
- `.checkPTT()` — L1733
- `.CWX_Load()` — L1749
- `.CWX_Closing()` — L1758
- `.Show()` — L1788
- `.TimerPeriodicEventCallback()` — L1809
- `.PressFNkey()` — L1814
- `.s1_Click()` — L1820
- `.s2_Click()` — L1826
- `.s3_Click()` — L1832
- `.s4_Click()` — L1838
- `.s5_Click()` — L1844
- `.s6_Click()` — L1850
- `.s7_Click()` — L1856
- `.s8_Click()` — L1861
- `.s9_Click()` — L1866
- `.cbMorse_MouseDown()` — L1871
- `.s1_MouseDown()` — L1877
- `.s2_MouseDown()` — L1881
- `.s3_MouseDown()` — L1885
- `.s4_MouseDown()` — L1889
- `.s5_MouseDown()` — L1893
- `.s6_MouseDown()` — L1897
- `.s7_MouseDown()` — L1901
- `.s8_MouseDown()` — L1905
- `.s9_MouseDown()` — L1909
- `.StopAction()` — L1914
- `.stopActionCore()` — L1919
- `.stopButton_Click()` — L1927
- `.udWPM_ValueChanged()` — L1931
- `.udWPM_LostFocus()` — L1939
- `.udDelay_ValueChanged()` — L1944
- `.udDelay_LostFocus()` — L1949
- `.udDrop_ValueChanged()` — L1954
- `.udDrop_LostFocus()` — L1958
- `.udPtt_ValueChanged()` — L1963
- `.udPtt_LostFocus()` — L1971
- `.CWX_MouseMove()` — L1976
- `.CWX_Paint()` — L1981
- `.chkPause_CheckedChanged()` — L1986
- `.chkAlwaysOnTop_CheckedChanged()` — L1991
- `.clear_keys()` — L2011
- `.show_keys()` — L2025
- `.clearButton_Click()` — L2089
- `.clear_show()` — L2093
- `.editit()` — L2106
- `.insert_and_reload()` — L2140
- `.write_a2m2()` — L2171
- `.process_element()` — L2199
- `.keyboardFifo()` — L2286
- `.keyboardDisplay()` — L2324
- `.queue_start()` — L2361
- `.loadchar()` — L2396
- `.loadmsg()` — L2429
- `.process_key()` — L2505
- `.CbMorse_SelectedIndexChanged()` — L2524
- `.insert_key()` — L2529
- `.CWX_FormClosing()` — L2549
- `.backspace()` — L2564
- `.msg2keys()` — L2580
- `.StopEverything()` — L2617
- `.onGlobalKeyDown()` — L2633
- `.onGlobalKeyUp()` — L2642

#### `TimerMode` (type, L197)

_No extracted members._

#### `TimerCaps` (type, L204)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/cwx.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
