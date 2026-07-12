# `Console/ucInfoBar.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** The info bar (status/warning strip) and its popup.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/Andromeda/Andromeda.cs` (references ×2)
  - `Console/Invoke/checkboxts.cs` (references ×2)
  - `Console/frmInfoBarPopup.Designer.cs` (references ×2)

## Outline

### Types

#### `ucInfoBar` (type, L52)

- `.SendMessage()` — L55
- `.actionString()` — L321
- `.OnActionClicked_Button1()` — L326
- `.OnActionClicked_Button2()` — L346
- `.doAction()` — L366
- `.addPopup()` — L394
- `.OnPopupClosed()` — L417
- `.ShutDown()` — L420
- `.onWarning()` — L443
- `.onTick()` — L453
- `.LateInit()` — L519
- `.OnMoxChangeHandler()` — L554
- `.setPSboolsToFalse()` — L562
- `.chkButton1_CheckedChanged()` — L568
- `.chkButton2_CheckedChanged()` — L578
- `.UpdateButtonState()` — L588
- `.Left1()` — L657
- `.Left2()` — L673
- `.Left3()` — L689
- `.Right1()` — L705
- `.Right2()` — L721
- `.Right3()` — L737
- `.SetToolTipLeft()` — L753
- `.SetToolTipRight()` — L777
- `.PSInfo()` — L808
- `.updatePSDisplay()` — L839
- `.Warning()` — L911
- `.InfoBar_Resize()` — L934
- `.InfoBar_Click()` — L946
- `.flip()` — L950
- `.updateLabels()` — L958
- `.chkButton1_MouseDown()` — L989
- `.chkButton2_MouseDown()` — L1013
- `.IsRightButton()` — L1037
- `.lblFB_MouseDown()` — L1042
- `.setToolTips()` — L1081
- `.replaceMainButton()` — L1098
- `.GetPopupButton()` — L1115
- `.lblSplitter_MouseDown()` — L1130
- `.lblSplitter_MouseEnter()` — L1138
- `.lblSplitter_MouseHover()` — L1146
- `.lblSplitter_MouseLeave()` — L1151
- `.lblSplitter_MouseMove()` — L1157
- `.repositionControls()` — L1187
- `.lblSplitter_MouseUp()` — L1261

#### `ActionTypes` (type, L59)

_No extracted members._

#### `InfoBarAction` (type, L73)

_No extracted members._

#### `ActionState` (type, L137)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucInfoBar.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
