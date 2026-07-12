# `Console/frmSeqLog.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Sequence-error log window — shows dropped/out-of-order UDP packet statistics.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Dumpcap.cs` (calls ×2)

## Outline

### Types

#### `frmSeqLog` (type, L54)

- `.InitAndShow()` — L74
- `.SetWireSharkPath()` — L81
- `.btnClear_Click()` — L86
- `.LogString()` — L98
- `.frmSeqLog_FormClosing()` — L110
- `.btnCopyToClipboard_Click()` — L125
- `.btnCopyImageToClipboard_Click()` — L130
- `.btnSetWireSharkFolder_Click()` — L139
- `.setupControlsDumpCap()` — L152
- `.udInterface_ValueChanged()` — L172
- `.chkKillOnNegativeOnly_CheckedChanged()` — L177
- `.chkDumpCapEnabled_CheckedChanged()` — L182
- `.chkClearRingBufferFolderOnRestart_CheckedChanged()` — L187
- `.btnShowDumpCapFolder_Click()` — L192
- `.chkStatusBarWarningNegativeOnly_CheckedChanged()` — L197

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmSeqLog.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
