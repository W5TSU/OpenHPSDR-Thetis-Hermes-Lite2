# `Console/Invoke/trackbarts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmNotchPopup.Designer.cs` (references ×1)
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `TrackBarTS` (type, L33)

- **`.BringToFront()`** — L497 — `public new void BringToFront()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L504 — `public new bool Contains(Control ctl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateControl()`** — L516 — `public new void CreateControl()`
  Creates control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateGraphics()`** — L523 — `public new Graphics CreateGraphics()`
  Creates graphics.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L547 — `public new virtual void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoDragDrop()`** — L554 — `public new DragDropEffects DoDragDrop(object data, DragDropEffects allowedEffects)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Equals()`** — L566 — `public new virtual object Equals(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindForm()`** — L578 — `public new Form FindForm()`
  Finds form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Focus()`** — L590 — `public new bool Focus()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetChildAtPoint()`** — L602 — `public new Control GetChildAtPoint(System.Drawing.Point pt)`
  Returns child at point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetContainerControl()`** — L614 — `public new IContainerControl GetContainerControl()`
  Returns container control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L626 — `public new virtual int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLifetimeService()`** — L638 — `public new virtual object GetLifetimeService()`
  Returns lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNextControl()`** — L650 — `public new Control GetNextControl(Control ctl, bool forward)`
  Returns next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetType()`** — L662 — `public new Type GetType()`
  Returns type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Hide()`** — L674 — `public new void Hide()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeLifetimeService()`** — L681 — `public new virtual object InitializeLifetimeService()`
  Initializes lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Invalidate()`** — L693 — `public new void Invalidate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformLayout()`** — L753 — `public new void PerformLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToClient()`** — L771 — `public new System.Drawing.Point PointToClient(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToScreen()`** — L783 — `public new System.Drawing.Point PointToScreen(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreProcessMessage()`** — L795 — `public new virtual bool PreProcessMessage(ref Message msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToClient()`** — L807 — `public new Rectangle RectangleToClient(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToScreen()`** — L819 — `public new Rectangle RectangleToScreen(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Refresh()`** — L831 — `public new virtual void Refresh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBackColor()`** — L838 — `public new virtual void ResetBackColor()`
  Resets back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBindings()`** — L845 — `public new void ResetBindings()`
  Resets bindings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetCursor()`** — L852 — `public new virtual void ResetCursor()`
  Resets cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetFont()`** — L859 — `public new virtual void ResetFont()`
  Resets font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetForeColor()`** — L866 — `public new virtual void ResetForeColor()`
  Resets fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetImeMode()`** — L873 — `public new void ResetImeMode()`
  Resets ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetRightToLeft()`** — L880 — `public new virtual void ResetRightToLeft()`
  Resets right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetText()`** — L887 — `public new virtual void ResetText()`
  Resets text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResumeLayout()`** — L894 — `public new void ResumeLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Scale()`** — L912 — `public new void Scale(SizeF ratio)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Select()`** — L923 — `public new void Select()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectNextControl()`** — L930 — `public new bool SelectNextControl(Control ctl, bool forward, bool tabStopOnly, bool nested, bool wrap)`
  Selects next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendToBack()`** — L944 — `public new void SendToBack()`
  Sends to back.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBounds()`** — L951 — `public new void SetBounds(int x, int y, int width, int height)`
  Sets bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRange()`** — L974 — `public new void SetRange(int minValue, int maxValue)`
  Sets range.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L985 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendLayout()`** — L992 — `public new void SuspendLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L999 — `public new virtual string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L1011 — `public new void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/trackbarts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
