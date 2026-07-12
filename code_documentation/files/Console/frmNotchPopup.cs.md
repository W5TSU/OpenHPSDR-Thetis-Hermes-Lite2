# `Console/frmNotchPopup.cs`

**Functional area:** [6. DSP control from the console](../../CODE_OUTLINE.md#6-dsp-control-from-the-console)

**Role:** Manual notch filter add/edit popup (backed by wdsp `nbp`).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/radio.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmNotchPopup` (type, L48)

- **`.Show()`** — L74 — `public void Show(MNotch notch, int minWidth, int maxWidth, bool top, int notch_index = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FrmNotchPopup_Deactivate()`** — L117 — `private void FrmNotchPopup_Deactivate(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BtnDelete_Click()`** — L125 — `private void BtnDelete_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `BtnDelete` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setBW()`** — L160 — `private void setBW(int width)`
  Sets bw.
  Called by: `.Btn25_Click()` (same file), `.Btn50_Click()` (same file), `.Btn100_Click()` (same file), `.Btn200_Click()` (same file)
- **`.Btn25_Click()`** — L167 — `private void Btn25_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `Btn25` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Btn50_Click()`** — L172 — `private void Btn50_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `Btn50` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Btn100_Click()`** — L177 — `private void Btn100_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `Btn100` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Btn200_Click()`** — L182 — `private void Btn200_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `Btn200` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.TrkWidth_Scroll()`** — L187 — `private void TrkWidth_Scroll(object sender, EventArgs e)`
  WinForms event handler: runs when `TrkWidth` is scrolled.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.setText()`** — L194 — `private void setText(int v)`
  Sets text.
  Called by: `.Show()` (same file), `.setBW()` (same file), `.TrkWidth_Scroll()` (same file)
- **`.ChkActive_CheckedChanged()`** — L199 — `private void ChkActive_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `ChkActive` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmNotchPopup.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
