# `Console/frmMacroButtonConfig.cs`

**Functional area:** [13. Andromeda control surface](../../CODE_OUTLINE.md#13-andromeda-control-surface)

**Role:** User-programmable macro buttons and their configuration grid.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
  - `Console/clsCatAtonic.cs` (references ×1, calls ×1)
  - `Console/ucOtherButtonsOptionsGrid.cs` (references ×2)
  - `Console/CAT/CATTester.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `frmMacroButtonConfig` (type, L48)

- **`.InitAndShow()`** — L84 — `public DialogResult InitAndShow(OtherButtonMacroSettings settings, Dictionary<string, string> containers, ref OtherButtonMacroSettings working_set, Console c)`
  Inits and show.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.txtON_TextChanged()`** — L229 — `private void txtON_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtON` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtOFF_TextChanged()`** — L235 — `private void txtOFF_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtOFF` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtNotes_TextChanged()`** — L241 — `private void txtNotes_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtNotes` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkClosesParent_CheckedChanged()`** — L247 — `private void chkClosesParent_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkClosesParent` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.getIndexFromName()`** — L253 — `private int getIndexFromName(object sender)`
  Returns index from name.
  Called by: `.chkClosesContainer_n_CheckedChanged()` (same file), `.comboOpenContainer_n_SelectedIndexChanged()` (same file), `.chkOpensContainer_n_CheckedChanged()` (same file), `.comboCloseContainer_n_SelectedIndexChanged()` (same file), `.chkUseParentCoodsForOpen_n_CheckedChanged()` (same file), `.chkSendMssageViaMMIO_n_CheckedChanged()` (same file) — and 4 more
- **`.chkClosesContainer_n_CheckedChanged()`** — L272 — `private void chkClosesContainer_n_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkClosesContainer_n` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboOpenContainer_n_SelectedIndexChanged()`** — L281 — `private void comboOpenContainer_n_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboOpenContainer_n` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkOpensContainer_n_CheckedChanged()`** — L291 — `private void chkOpensContainer_n_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkOpensContainer_n` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboCloseContainer_n_SelectedIndexChanged()`** — L300 — `private void comboCloseContainer_n_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboCloseContainer_n` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.updateUseParent()`** — L309 — `private void updateUseParent(int idx)`
  Called by: `.InitAndShow()` (same file), `.chkClosesContainer_n_CheckedChanged()` (same file), `.comboOpenContainer_n_SelectedIndexChanged()` (same file), `.chkOpensContainer_n_CheckedChanged()` (same file), `.comboCloseContainer_n_SelectedIndexChanged()` (same file)
- **`.chkUseParentCoodsForOpen_n_CheckedChanged()`** — L353 — `private void chkUseParentCoodsForOpen_n_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkUseParentCoodsForOpen_n` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkSendMssageViaMMIO_n_CheckedChanged()`** — L360 — `private void chkSendMssageViaMMIO_n_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkSendMssageViaMMIO_n` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_4char_n_TextChanged()`** — L367 — `private void txtMMIO_4char_n_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_4char_n` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_message_n_TextChanged()`** — L374 — `private void txtMMIO_message_n_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_message_n` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtMMIO_message_n_off_TextChanged()`** — L380 — `private void txtMMIO_message_n_off_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtMMIO_message_n_off` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.radButtonState_n_CheckedChanged()`** — L386 — `private void radButtonState_n_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `radButtonState_n` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtButtonState_led_4char_TextChanged()`** — L423 — `private void txtButtonState_led_4char_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtButtonState_led_4char` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.comboButtonState_container_visibility_SelectedIndexChanged()`** — L429 — `private void comboButtonState_container_visibility_SelectedIndexChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `comboButtonState_container_visibility` selection changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtButtonState_cat_on_reply_TextChanged()`** — L437 — `private void txtButtonState_cat_on_reply_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtButtonState_cat_on_reply` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.txtCatMacro_TextChanged()`** — L443 — `private void txtCatMacro_TextChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `txtCatMacro` text changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkCatSend_CheckedChanged()`** — L472 — `private void chkCatSend_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkCatSend` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCancel_Click()`** — L479 — `private void btnCancel_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCancel` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnOK_Click()`** — L484 — `private void btnOK_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnOK` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.btnCatTest_Click()`** — L489 — `private void btnCatTest_Click(object sender, EventArgs e)`
  WinForms event handler: runs when `btnCatTest` is clicked.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.chkRunStateCommandOnVisible_CheckedChanged()`** — L500 — `private void chkRunStateCommandOnVisible_CheckedChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `chkRunStateCommandOnVisible` checked state changes.
  Called by: WinForms event wiring at runtime (no static call sites).

#### `clsContainerComboboxItem` (type, L50)

- **`.ToString()`** — L56 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmMacroButtonConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
