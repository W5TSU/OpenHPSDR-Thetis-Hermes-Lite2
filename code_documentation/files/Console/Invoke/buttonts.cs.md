# `Console/Invoke/buttonts.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI.

## How this file is used

- Used by (incoming references from other files):
  - `Console/Andromeda/BandButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/FilterButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeButtonsPopup.cs` (references ×1)
  - `Console/Andromeda/ModeDependentSettingsForm.cs` (references ×1)
  - `Console/Andromeda/SliderSettingsForm.cs` (references ×1)
  - `Console/Andromeda/displaysettingsform.cs` (references ×1)
  - `Console/Andromeda/vfosettingspopup.cs` (references ×1)
  - `Console/ColorButton.cs` (inherits ×1)
  - `Console/DiversityForm.cs` (references ×1)
  - `Console/console.Designer.cs` (references ×1)
  - `Console/cwx.cs` (references ×1)
  - `Console/eqform.cs` (references ×1)
  - …and 24 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `System.Windows.Forms` (namespace, L31)

_No extracted members._

#### `ButtonTS` (type, L33)

- **`.BringToFront()`** — L546 — `public new void BringToFront()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Contains()`** — L553 — `public new bool Contains(Control ctl)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateControl()`** — L565 — `public new void CreateControl()`
  Creates control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CreateGraphics()`** — L572 — `public new Graphics CreateGraphics()`
  Creates graphics.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Dispose()`** — L596 — `public new virtual void Dispose()`
  Releases the object’s resources.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DoDragDrop()`** — L603 — `public new DragDropEffects DoDragDrop(object data, DragDropEffects allowedEffects)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Equals()`** — L615 — `public new virtual object Equals(object obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FindForm()`** — L627 — `public new Form FindForm()`
  Finds form.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Focus()`** — L639 — `public new bool Focus()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetChildAtPoint()`** — L651 — `public new Control GetChildAtPoint(System.Drawing.Point pt)`
  Returns child at point.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetContainerControl()`** — L663 — `public new IContainerControl GetContainerControl()`
  Returns container control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L675 — `public new virtual int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetLifetimeService()`** — L687 — `public new virtual object GetLifetimeService()`
  Returns lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNextControl()`** — L699 — `public new Control GetNextControl(Control ctl, bool forward)`
  Returns next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetType()`** — L711 — `public new Type GetType()`
  Returns type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Hide()`** — L723 — `public new void Hide()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeLifetimeService()`** — L730 — `public new virtual object InitializeLifetimeService()`
  Initializes lifetime service.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Invalidate()`** — L742 — `public new void Invalidate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NotifyDefault()`** — L802 — `public new virtual void NotifyDefault(bool val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformClick()`** — L811 — `public new virtual void PerformClick()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PerformLayout()`** — L818 — `public new void PerformLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToClient()`** — L836 — `public new System.Drawing.Point PointToClient(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PointToScreen()`** — L848 — `public new System.Drawing.Point PointToScreen(System.Drawing.Point p)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreProcessMessage()`** — L860 — `public new virtual bool PreProcessMessage(ref Message msg)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToClient()`** — L872 — `public new Rectangle RectangleToClient(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RectangleToScreen()`** — L884 — `public new Rectangle RectangleToScreen(Rectangle r)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Refresh()`** — L896 — `public new virtual void Refresh()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBackColor()`** — L903 — `public new virtual void ResetBackColor()`
  Resets back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetBindings()`** — L910 — `public new void ResetBindings()`
  Resets bindings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetCursor()`** — L917 — `public new virtual void ResetCursor()`
  Resets cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetFont()`** — L924 — `public new virtual void ResetFont()`
  Resets font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetForeColor()`** — L931 — `public new virtual void ResetForeColor()`
  Resets fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetImeMode()`** — L938 — `public new void ResetImeMode()`
  Resets ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetRightToLeft()`** — L945 — `public new virtual void ResetRightToLeft()`
  Resets right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetText()`** — L952 — `public new virtual void ResetText()`
  Resets text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResumeLayout()`** — L959 — `public new void ResumeLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Scale()`** — L977 — `public new void Scale(SizeF ratio)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Select()`** — L988 — `public new void Select()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SelectNextControl()`** — L995 — `public new bool SelectNextControl(Control ctl, bool forward, bool tabStopOnly, bool nested, bool wrap)`
  Selects next control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendToBack()`** — L1009 — `public new void SendToBack()`
  Sends to back.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBounds()`** — L1016 — `public new void SetBounds(int x, int y, int width, int height)`
  Sets bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Show()`** — L1039 — `public new void Show()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SuspendLayout()`** — L1046 — `public new void SuspendLayout()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L1053 — `public new virtual string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L1065 — `public new void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/buttonts.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
