# `Console/ucQuickRecall.cs`

**Functional area:** [15. Memories, band stacks, and the database](../../CODE_OUTLINE.md#15-memories-band-stacks-and-the-database)

**Role:** Quick recall (recent frequencies) list.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×7)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/frmQuickRecallPopupList.Designer.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucQuickRecall` (type, L49)

- **`.resizeAndReposition()`** — L142 — `private void resizeAndReposition()`
  Called by: `.ucQuickRecall_Resize()` (same file)
- **`.OnPopupClosed()`** — L153 — `private void OnPopupClosed(object sender, ToolStripDropDownClosedEventArgs e)`
  Handles/raises the popup closed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMoxChanged()`** — L158 — `private void OnMoxChanged(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChanged()`** — L174 — `private void OnModeChanged(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOAChange()`** — L192 — `private void OnVFOAChange(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double oldCentreF, doubl`
  Handles/raises the vfoachange event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.buttonClicked()`** — L204 — `private void buttonClicked(EventArgs e)`
  Called by: `.btnPrevious_Click()` (same file), `.btnNext_Click()` (same file), `.btnList_Click()` (same file)
- **`.formatFrequencyToString()`** — L210 — `private string formatFrequencyToString(double f)`
  Called by: `.OnVFOATick()` (same file)
- **`.findExistingInVFOA()`** — L221 — `private int findExistingInVFOA(double f)`
  Called by: `.OnVFOAModeTick()` (same file), `.OnVFOATick()` (same file)
- **`.addVFOAEntry()`** — L230 — `private void addVFOAEntry(QuickInfo qi)`
  Called by: `.OnVFOATick()` (same file)
- **`.OnBackgroundColourPingerTick()`** — L238 — `private void OnBackgroundColourPingerTick(Object sender, EventArgs e)`
  Handles/raises the background colour pinger tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOAModeTick()`** — L244 — `private void OnVFOAModeTick(Object sender, EventArgs e)`
  Handles/raises the vfoamode tick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOATick()`** — L267 — `private void OnVFOATick(Object sender, EventArgs e)`
  Handles/raises the vfoatick event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDispose()`** — L299 — `private void OnDispose(object sender, EventArgs e)`
  Handles/raises the dispose event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEntrySelected()`** — L310 — `private void OnEntrySelected(int index)`
  Handles/raises the entry selected event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.btnPrevious_Click()`** — L318 — `private void btnPrevious_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnPrevious` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnNext_Click()`** — L329 — `private void btnNext_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnNext` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.selectVFOAEntry()`** — L340 — `private void selectVFOAEntry(int index)`
  Called by: `.OnEntrySelected()` (same file), `.btnPrevious_Click()` (same file), `.btnNext_Click()` (same file)
- **`.btnList_Click()`** — L348 — `private void btnList_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnList` is clicked.
  Called by: `.lblFlashColour_Click()` (same file)
- **`.buildAndShowPopup()`** — L361 — `private void buildAndShowPopup()`
  Called by: `.OnVFOATick()` (same file), `.btnList_Click()` (same file)
- **`.ucQuickRecall_BackColorChanged()`** — L378 — `private void ucQuickRecall_BackColorChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setButtonBackColour()`** — L383 — `private void setButtonBackColour()`
  Sets button back colour.
  Called by: `.ucQuickRecall_BackColorChanged()` (same file)
- **`.ucQuickRecall_Resize()`** — L390 — `private void ucQuickRecall_Resize(object sender, EventArgs e)`
  WinForms event handler: runs when `ucQuickRecall` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.lblFlashColour_Click()`** — L395 — `private void lblFlashColour_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `lblFlashColour` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `QuickInfo` (type, L64)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucQuickRecall.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
