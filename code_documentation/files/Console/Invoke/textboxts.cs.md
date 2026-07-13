# `Console/Invoke/textboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/CAT/CATTester.cs` (calls ×2, references ×1)
  - `Console/Andromeda/vfosettingspopup.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/frmAddCustomRadio.Designer.cs` (references ×1)
  - `Console/frmFilterManager.Designer.cs` (references ×1)
  - `Console/frmFinder.Designer.cs` (references ×1)
  - `Console/frmLog.Designer.cs` (references ×1)
  - `Console/frmMacroButtonConfig.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - …and 4 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Focus()` (×1), `.Clear()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TextBoxTS` (type, L33)

- **`.AppendText()`** — L652 — `public new void AppendText(string text)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BringToFront()`** — L661 — `public new void BringToFront()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Clear()`** — L668 — `public new void Clear()`
  Called by: `.ExecuteCommand()` (`Console/CAT/CATTester.cs`)
- **`.ClearUndo()`** — L675 — `public new void ClearUndo()`
  Clears undo.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L682 — `public new bool Contains(Control ctl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Copy()`** — L694 — `public new void Copy()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateControl()`** — L701 — `public new void CreateControl()`
  Creates control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateGraphics()`** — L708 — `public new Graphics CreateGraphics()`
  Creates graphics.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Cut()`** — L732 — `public new void Cut()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L739 — `public new virtual void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoDragDrop()`** — L746 — `public new DragDropEffects DoDragDrop(object data, DragDropEffects allowedEffects)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Equals()`** — L758 — `public new virtual object Equals(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindForm()`** — L770 — `public new Form FindForm()`
  Finds form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Focus()`** — L782 — `public new bool Focus()`
  Called by: `.Setup()` (`Console/CAT/CATTester.cs`)
- **`.GetChildAtPoint()`** — L794 — `public new Control GetChildAtPoint(System.Drawing.Point pt)`
  Returns child at point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetContainerControl()`** — L806 — `public new IContainerControl GetContainerControl()`
  Returns container control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L818 — `public new virtual int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLifetimeService()`** — L830 — `public new virtual object GetLifetimeService()`
  Returns lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNextControl()`** — L842 — `public new Control GetNextControl(Control ctl, bool forward)`
  Returns next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetType()`** — L854 — `public new Type GetType()`
  Returns type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Hide()`** — L866 — `public new void Hide()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeLifetimeService()`** — L873 — `public new virtual object InitializeLifetimeService()`
  Initializes lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Invalidate()`** — L885 — `public new void Invalidate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Paste()`** — L947 — `public new void Paste()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformLayout()`** — L954 — `public new void PerformLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToClient()`** — L972 — `public new System.Drawing.Point PointToClient(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToScreen()`** — L984 — `public new System.Drawing.Point PointToScreen(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreProcessMessage()`** — L996 — `public new virtual bool PreProcessMessage(ref Message msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToClient()`** — L1008 — `public new Rectangle RectangleToClient(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToScreen()`** — L1020 — `public new Rectangle RectangleToScreen(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Refresh()`** — L1032 — `public new virtual void Refresh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBackColor()`** — L1039 — `public new virtual void ResetBackColor()`
  Resets back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBindings()`** — L1046 — `public new void ResetBindings()`
  Resets bindings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetCursor()`** — L1053 — `public new virtual void ResetCursor()`
  Resets cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetFont()`** — L1060 — `public new virtual void ResetFont()`
  Resets font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetForeColor()`** — L1067 — `public new virtual void ResetForeColor()`
  Resets fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetImeMode()`** — L1074 — `public new void ResetImeMode()`
  Resets ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetRightToLeft()`** — L1081 — `public new virtual void ResetRightToLeft()`
  Resets right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetText()`** — L1088 — `public new virtual void ResetText()`
  Resets text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResumeLayout()`** — L1095 — `public new void ResumeLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Scale()`** — L1113 — `public new void Scale(SizeF ratio)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ScrollToCaret()`** — L1124 — `public new void ScrollToCaret()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Select()`** — L1131 — `public new void Select(int start, int length)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectAll()`** — L1149 — `public new void SelectAll()`
  Selects all.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectNextControl()`** — L1156 — `public new bool SelectNextControl(Control ctl, bool forward, bool tabStopOnly, bool nested, bool wrap)`
  Selects next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendToBack()`** — L1170 — `public new void SendToBack()`
  Sends to back.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBounds()`** — L1177 — `public new void SetBounds(int x, int y, int width, int height)`
  Sets bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L1200 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendLayout()`** — L1207 — `public new void SuspendLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L1214 — `public new virtual string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Undo()`** — L1226 — `public new void Undo()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L1233 — `public new void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/textboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
