# `Midi2Cat/Midi2CatSetupForm.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** Mapping editor UI: wiggle a control, pick a function.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2CatSetupForm` (type, L38)

- **`.Midi2CatSetupForm_Load()`** — L53 — `private void Midi2CatSetupForm_Load(object sender, EventArgs e)`
  WinForms event handler: runs when `Midi2CatSetupForm` loads.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.LoadSetup()`** — L60 — `private void LoadSetup()`
  Loads setup.
  Called by: `.Midi2CatSetupForm_Load()` (same file)
- **`.Midi2CatSetupForm_FormClosing()`** — L76 — `private void Midi2CatSetupForm_FormClosing(object sender, FormClosingEventArgs e)`
  WinForms event handler: runs when `Midi2CatSetupForm` is closing.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.saveButton_Click()`** — L85 — `private void saveButton_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `saveButton` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.startTimer_Tick()`** — L90 — `private void startTimer_Tick(object sender, EventArgs e)`
  WinForms event handler: runs when `startTimer` timer fires.
  Called by: WinForms event wiring at runtime (no static call sites).

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Midi2Cat/Midi2CatSetupForm.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
