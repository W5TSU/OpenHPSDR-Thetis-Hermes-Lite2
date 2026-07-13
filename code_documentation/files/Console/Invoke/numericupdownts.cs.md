# `Console/Invoke/numericupdownts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - `Console/frmCFCConfig.Designer.cs` (references ×1)
  - `Console/frmSeqLog.Designer.cs` (references ×1)
  - `Console/Memory/MemoryForm.Designer.cs` (references ×1)
  - `Console/PSForm.designer.cs` (references ×1)
  - `Console/RAForm.Designer.cs` (references ×1)
  - `Console/rxa.Designer.cs` (references ×1)
  - `Console/rxaControls.Designer.cs` (references ×1)
  - `Console/setup.designer.cs` (references ×1)
  - …and 1 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `NumericUpDownTS` (type, L33)

- **`.BringToFront()`** — L640 — `public new void BringToFront()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L647 — `public new bool Contains(Control ctl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateControl()`** — L659 — `public new void CreateControl()`
  Creates control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateGraphics()`** — L666 — `public new Graphics CreateGraphics()`
  Creates graphics.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L690 — `public new virtual void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoDragDrop()`** — L697 — `public new DragDropEffects DoDragDrop(object data, DragDropEffects allowedEffects)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DownButton()`** — L709 — `public override void DownButton()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Equals()`** — L716 — `public new virtual object Equals(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindForm()`** — L728 — `public new Form FindForm()`
  Finds form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Focus()`** — L740 — `public new bool Focus()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetChildAtPoint()`** — L752 — `public new Control GetChildAtPoint(System.Drawing.Point pt)`
  Returns child at point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetContainerControl()`** — L764 — `public new IContainerControl GetContainerControl()`
  Returns container control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L776 — `public new virtual int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLifetimeService()`** — L788 — `public new virtual object GetLifetimeService()`
  Returns lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNextControl()`** — L800 — `public new Control GetNextControl(Control ctl, bool forward)`
  Returns next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetType()`** — L812 — `public new Type GetType()`
  Returns type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Hide()`** — L824 — `public new void Hide()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeLifetimeService()`** — L831 — `public new virtual object InitializeLifetimeService()`
  Initializes lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Invalidate()`** — L843 — `public new void Invalidate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformLayout()`** — L903 — `public new void PerformLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToClient()`** — L921 — `public new System.Drawing.Point PointToClient(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToScreen()`** — L933 — `public new System.Drawing.Point PointToScreen(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreProcessMessage()`** — L945 — `public new virtual bool PreProcessMessage(ref Message msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToClient()`** — L957 — `public new Rectangle RectangleToClient(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToScreen()`** — L969 — `public new Rectangle RectangleToScreen(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Refresh()`** — L981 — `public new virtual void Refresh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBackColor()`** — L988 — `public new virtual void ResetBackColor()`
  Resets back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBindings()`** — L995 — `public new void ResetBindings()`
  Resets bindings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetCursor()`** — L1002 — `public new virtual void ResetCursor()`
  Resets cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetFont()`** — L1009 — `public new virtual void ResetFont()`
  Resets font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetForeColor()`** — L1016 — `public new virtual void ResetForeColor()`
  Resets fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetImeMode()`** — L1023 — `public new void ResetImeMode()`
  Resets ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetRightToLeft()`** — L1030 — `public new virtual void ResetRightToLeft()`
  Resets right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetText()`** — L1037 — `public new virtual void ResetText()`
  Resets text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResumeLayout()`** — L1044 — `public new void ResumeLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Scale()`** — L1062 — `public new void Scale(SizeF ratio)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ScrollControlIntoView()`** — L1073 — `public new void ScrollControlIntoView(Control activeControl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Select()`** — L1082 — `public new void Select()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectNextControl()`** — L1098 — `public new bool SelectNextControl(Control ctl, bool forward, bool tabStopOnly, bool nested, bool wrap)`
  Selects next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendToBack()`** — L1112 — `public new void SendToBack()`
  Sends to back.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAutoScrollMargin()`** — L1119 — `public new void SetAutoScrollMargin(int x, int y)`
  Sets auto scroll margin.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBounds()`** — L1128 — `public new void SetBounds(int x, int y, int width, int height)`
  Sets bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L1151 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendLayout()`** — L1158 — `public new void SuspendLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L1165 — `public new virtual string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpButton()`** — L1177 — `public override void UpButton()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L1184 — `public new void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Validate()`** — L1191 — `public new bool Validate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseWheel()`** — L1212 — `protected override void OnMouseWheel(MouseEventArgs e)`
  Handles/raises the mouse wheel event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/numericupdownts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
