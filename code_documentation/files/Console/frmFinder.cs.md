# `Console/frmFinder.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Simple text-input dialog; searchable "find a setting" helper.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/common.cs` (calls ×3)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmFinder` (type, L55)

- **`.GatherSearchData()`** — L98 — `public void GatherSearchData(Form frm, ToolTip tt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GatherCATStructData()`** — L119 — `public void GatherCATStructData(string file_path)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.gatherCATStructSearchDataThread()`** — L148 — `private void gatherCATStructSearchDataThread(string file_path)`
  Called by: `.GatherCATStructData()` (same file)
- **`.gatherSearchDataThread()`** — L211 — `private void gatherSearchDataThread(Control frm, ToolTip tt)`
  Called by: `.GatherSearchData()` (same file)
- **`.stripPrefix()`** — L236 — `private static string stripPrefix(string name)`
  Called by: `.getControlList()` (same file)
- **`.getControlList()`** — L247 — `private void getControlList(Control root, ToolTip tt)`
  Returns control list.
  Called by: `.gatherSearchDataThread()` (same file)
- **`.txtSearch_TextChanged()`** — L316 — `private void txtSearch_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtSearch` text changes.
  Called by: `.chkFullDetails_CheckedChanged()` (same file), `.frmFinder_KeyDown()` (same file), `.chkHighlight_CheckedChanged()` (same file), `.chkKeywords_CheckedChanged()` (same file)
- **`.lstResults_SelectedIndexChanged()`** — L372 — `private void lstResults_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `lstResults` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lstResults_DrawItem()`** — L395 — `private void lstResults_DrawItem(object sender, DrawItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.highlight()`** — L468 — `private void highlight(string sSearchText, string sLineText, ListBox listBox, int xPos, int yPos, Graphics g)`
  Called by: `.lstResults_DrawItem()` (same file)
- **`.findSubstringOccurrences()`** — L492 — `private List<Tuple<int, int>> findSubstringOccurrences(string inputString, string searchString)`
  Called by: `.highlight()` (same file)
- **`.applyTint()`** — L516 — `private Color applyTint(Color baseColor, Color tintColor)`
  Called by: `.lstResults_DrawItem()` (same file)
- **`.frmFinder_FormClosing()`** — L525 — `private void frmFinder_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `frmFinder` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Show()`** — L536 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.lstResults_MeasureItem()`** — L544 — `private void lstResults_MeasureItem(object sender, MeasureItemEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.showControl()`** — L563 — `private void showControl(Control c)`
  Called by: `.lstResults_SelectedIndexChanged()` (same file)
- **`.selectRequiredTabs()`** — L582 — `private void selectRequiredTabs(Control parentControl, Control targetControl)`
  Called by: `.showControl()` (same file)
- **`.chkFullDetails_CheckedChanged()`** — L653 — `private void chkFullDetails_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkFullDetails` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ReadXmlFinderFile()`** — L665 — `public void ReadXmlFinderFile(string directoryPath)`
  Reads xml finder file.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WriteXmlFinderFile()`** — L693 — `public void WriteXmlFinderFile(string directoryPath)`
  Writes xml finder file.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.frmFinder_KeyDown()`** — L746 — `private void frmFinder_KeyDown(object sender, KeyEventArgs e)`
  WinForms event handler: runs when `frmFinder` receives a key-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkHighlight_CheckedChanged()`** — L777 — `private void chkHighlight_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkHighlight` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkKeywords_CheckedChanged()`** — L789 — `private void chkKeywords_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkKeywords` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `SearchData` (type, L57)

_No extracted members._

#### `CatStructEntry` (type, L143)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmFinder.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
