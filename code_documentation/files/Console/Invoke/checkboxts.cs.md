# `Console/Invoke/checkboxts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/frmInfoBarPopup.cs` (references ×2)
  - `Console/setup.cs` (references ×1, calls ×1)
  - `Console/ucInfoBar.cs` (references ×2)
  - `Console/AmpView.Designer.cs` (references ×1)
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmBandStack2.Designer.cs` (references ×1)
  - `Console/frmBandwidth.Designer.cs` (references ×1)
  - …and 18 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.GetType()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CheckBoxTS` (type, L33)

- **`.BringToFront()`** — L586 — `public new void BringToFront()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L593 — `public new bool Contains(Control ctl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateControl()`** — L605 — `public new void CreateControl()`
  Creates control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateGraphics()`** — L612 — `public new Graphics CreateGraphics()`
  Creates graphics.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L636 — `public new virtual void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoDragDrop()`** — L643 — `public new DragDropEffects DoDragDrop(object data, DragDropEffects allowedEffects)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Equals()`** — L655 — `public new virtual object Equals(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindForm()`** — L667 — `public new Form FindForm()`
  Finds form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Focus()`** — L679 — `public new bool Focus()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetChildAtPoint()`** — L691 — `public new Control GetChildAtPoint(System.Drawing.Point pt)`
  Returns child at point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetContainerControl()`** — L703 — `public new IContainerControl GetContainerControl()`
  Returns container control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L715 — `public new virtual int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLifetimeService()`** — L727 — `public new virtual object GetLifetimeService()`
  Returns lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNextControl()`** — L739 — `public new Control GetNextControl(Control ctl, bool forward)`
  Returns next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetType()`** — L751 — `public new Type GetType()`
  Returns type.
  Called by: `.getOptions()` (`Console/setup.cs`)
- **`.Hide()`** — L763 — `public new void Hide()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeLifetimeService()`** — L770 — `public new virtual object InitializeLifetimeService()`
  Initializes lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Invalidate()`** — L782 — `public new void Invalidate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformLayout()`** — L844 — `public new void PerformLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToClient()`** — L862 — `public new System.Drawing.Point PointToClient(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToScreen()`** — L874 — `public new System.Drawing.Point PointToScreen(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreProcessMessage()`** — L886 — `public new virtual bool PreProcessMessage(ref Message msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToClient()`** — L898 — `public new Rectangle RectangleToClient(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToScreen()`** — L910 — `public new Rectangle RectangleToScreen(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Refresh()`** — L922 — `public new virtual void Refresh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBackColor()`** — L929 — `public new virtual void ResetBackColor()`
  Resets back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBindings()`** — L936 — `public new void ResetBindings()`
  Resets bindings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetCursor()`** — L943 — `public new virtual void ResetCursor()`
  Resets cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetFont()`** — L950 — `public new virtual void ResetFont()`
  Resets font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetForeColor()`** — L957 — `public new virtual void ResetForeColor()`
  Resets fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetImeMode()`** — L964 — `public new void ResetImeMode()`
  Resets ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetRightToLeft()`** — L971 — `public new virtual void ResetRightToLeft()`
  Resets right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetText()`** — L978 — `public new virtual void ResetText()`
  Resets text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResumeLayout()`** — L985 — `public new void ResumeLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Scale()`** — L1003 — `public new void Scale(SizeF ratio)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Select()`** — L1014 — `public new void Select()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectNextControl()`** — L1021 — `public new bool SelectNextControl(Control ctl, bool forward, bool tabStopOnly, bool nested, bool wrap)`
  Selects next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendToBack()`** — L1035 — `public new void SendToBack()`
  Sends to back.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBounds()`** — L1042 — `public new void SetBounds(int x, int y, int width, int height)`
  Sets bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L1065 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendLayout()`** — L1072 — `public new void SuspendLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L1079 — `public new virtual string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L1091 — `public new void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/checkboxts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
