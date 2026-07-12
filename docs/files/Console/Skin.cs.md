# `Console/Skin.cs`

**Functional area:** [2. Settings and configuration](../../CODE_OUTLINE.md#2-settings-and-configuration)

**Role:** UI skin loading and application (SkiaSharp-backed image skins for console controls).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×3)
  - `Console/Andromeda/Andromeda.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/PrettyTrackBar.cs` (references ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×2)
  - `Console/ucQuickRecall.Designer.cs` (references ×1)
  - `Console/ucInfoBar.Designer.cs` (references ×1)
- Most-referenced symbols from other files: `.Restore()` (×3), `.SetConsole()` (×1), `.Save()` (×1)

## Outline

### Types

#### `Skin` (type, L63)

- `.SetConsole()` — L94
- `.Save()` — L105
- `.Restore()` — L134
- `.ReadImages()` — L399
- `.SaveForm()` — L486
- `.RestoreForm()` — L498
- `.SaveGroupBox()` — L539
- `.RestoreGroupBox()` — L551
- `.SavePanel()` — L592
- `.RestorePanel()` — L602
- `.SaveButton()` — L637
- `.RestoreButton()` — L654
- `.SetupQuickRecallImages()` — L710
- `.setupButtonHandlers()` — L764
- `.SetupButtonImages()` — L787
- `.Button_StateChanged()` — L858
- `.Button_MouseEnter()` — L881
- `.Button_MouseDown()` — L891
- `.Button_MouseUp()` — L901
- `.SetButtonImageState()` — L906
- `.SaveCheckBox()` — L918
- `.RestoreCheckBox()` — L937
- `.SetupCheckBoxImages()` — L1000
- `.setupCheckBoxHandlers()` — L1052
- `.SetupInfoBar()` — L1070
- `.CheckBox_StateChanged()` — L1162
- `.CheckBox_MouseEnter()` — L1187
- `.SetCheckBoxImageState()` — L1198
- `.SaveComboBox()` — L1210
- `.RestoreComboBox()` — L1224
- `.SaveLabel()` — L1265
- `.RestoreLabel()` — L1280
- `.SaveNumericUpDown()` — L1324
- `.RestoreNumericUpDown()` — L1338
- `.SavePictureBox()` — L1379
- `.RestorePictureBox()` — L1391
- `.SaveRadioButton()` — L1426
- `.RestoreRadioButton()` — L1445
- `.SetupRadioButtonImages()` — L1508
- `.RadioButton_StateChanged()` — L1576
- `.RadioButton_MouseEnter()` — L1603
- `.SetRadioButtonImageState()` — L1614
- `.SaveTextBox()` — L1626
- `.RestoreTextBox()` — L1642
- `.SavePrettyTrackBar()` — L1689
- `.RestorePrettyTrackBar()` — L1700
- `.SetupPrettyTrackBarImages()` — L1728
- `.SaveSize()` — L1768
- `.RestoreSize()` — L1776
- `.SaveFont()` — L1794
- `.RestoreFont()` — L1805
- `.SaveLocation()` — L1840
- `.RestoreLocation()` — L1848
- `.SaveFlatAppearance()` — L1866
- `.StringToColor()` — L1874
- `.SetBackgroundImage()` — L1882
- `.loadImage()` — L1921
- `.getImageFromFilePath()` — L1952
- `.computeHashFromImage()` — L1959
- `.resizeImage()` — L1995

#### `ImageState` (type, L67)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Skin.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
