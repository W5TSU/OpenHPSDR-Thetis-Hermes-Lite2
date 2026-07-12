# `Console/frmSeqLog.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** Sequence-error log window — shows dropped/out-of-order UDP packet statistics.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Dumpcap.cs` (calls ×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmSeqLog` (type, L54)

- **`.InitAndShow()`** — L74 — `public void InitAndShow()`
  Inits and show.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWireSharkPath()`** — L81 — `public void SetWireSharkPath(string sPath)`
  Sets wire shark path.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnClear_Click()`** — L86 — `private void btnClear_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnClear` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.LogString()`** — L98 — `public void LogString(string s)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmSeqLog_FormClosing()`** — L110 — `private void frmSeqLog_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmSeqLog` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCopyToClipboard_Click()`** — L125 — `private void btnCopyToClipboard_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCopyToClipboard` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCopyImageToClipboard_Click()`** — L130 — `private void btnCopyImageToClipboard_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCopyImageToClipboard` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnSetWireSharkFolder_Click()`** — L139 — `private void btnSetWireSharkFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnSetWireSharkFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setupControlsDumpCap()`** — L152 — `private void setupControlsDumpCap(string sPath)`
  Called by: `.SetWireSharkPath()` (same file), `.btnSetWireSharkFolder_Click()` (same file)
- **`.udInterface_ValueChanged()`** — L172 — `private void udInterface_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `udInterface` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkKillOnNegativeOnly_CheckedChanged()`** — L177 — `private void chkKillOnNegativeOnly_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKillOnNegativeOnly` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkDumpCapEnabled_CheckedChanged()`** — L182 — `private void chkDumpCapEnabled_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkDumpCapEnabled` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkClearRingBufferFolderOnRestart_CheckedChanged()`** — L187 — `private void chkClearRingBufferFolderOnRestart_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkClearRingBufferFolderOnRestart` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnShowDumpCapFolder_Click()`** — L192 — `private void btnShowDumpCapFolder_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnShowDumpCapFolder` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkStatusBarWarningNegativeOnly_CheckedChanged()`** — L197 — `private void chkStatusBarWarningNegativeOnly_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkStatusBarWarningNegativeOnly` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmSeqLog.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
