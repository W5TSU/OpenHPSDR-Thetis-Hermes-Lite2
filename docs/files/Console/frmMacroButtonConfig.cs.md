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

### Types

#### `frmMacroButtonConfig` (type, L48)

- `.InitAndShow()` — L84
- `.txtON_TextChanged()` — L229
- `.txtOFF_TextChanged()` — L235
- `.txtNotes_TextChanged()` — L241
- `.chkClosesParent_CheckedChanged()` — L247
- `.getIndexFromName()` — L253
- `.chkClosesContainer_n_CheckedChanged()` — L272
- `.comboOpenContainer_n_SelectedIndexChanged()` — L281
- `.chkOpensContainer_n_CheckedChanged()` — L291
- `.comboCloseContainer_n_SelectedIndexChanged()` — L300
- `.updateUseParent()` — L309
- `.chkUseParentCoodsForOpen_n_CheckedChanged()` — L353
- `.chkSendMssageViaMMIO_n_CheckedChanged()` — L360
- `.txtMMIO_4char_n_TextChanged()` — L367
- `.txtMMIO_message_n_TextChanged()` — L374
- `.txtMMIO_message_n_off_TextChanged()` — L380
- `.radButtonState_n_CheckedChanged()` — L386
- `.txtButtonState_led_4char_TextChanged()` — L423
- `.comboButtonState_container_visibility_SelectedIndexChanged()` — L429
- `.txtButtonState_cat_on_reply_TextChanged()` — L437
- `.txtCatMacro_TextChanged()` — L443
- `.chkCatSend_CheckedChanged()` — L472
- `.btnCancel_Click()` — L479
- `.btnOK_Click()` — L484
- `.btnCatTest_Click()` — L489
- `.chkRunStateCommandOnVisible_CheckedChanged()` — L500

#### `clsContainerComboboxItem` (type, L50)

- `.ToString()` — L56

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/frmMacroButtonConfig.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
