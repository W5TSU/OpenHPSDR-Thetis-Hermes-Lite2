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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Skin` (type, L63)

- **`.SetConsole()`** — L94 — `public static void SetConsole(Console objConsole)`
  Sets console.
  Called by: `.AfterConstructor()` (`Console/setup.cs`)
- **`.Save()`** — L105 — `public static void Save(string name, string p, Form f)`
  Saves a forms appearance including properties of the form and its controls to xml.
  Called by: `.btnSkinExport_Click()` (`Console/setup.cs`)
- **`.Restore()`** — L134 — `public static bool Restore(string name, string p, Form f)`
  Restores a forms appearance including properties of the form and its controls from xml.
  Called by: `.ExecuteButtonAction()` (`Console/Andromeda/Andromeda.cs`), `.UpdateAndromedaSkins()` (`Console/Andromeda/Andromeda.cs`), `.comboAppSkin_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ReadImages()`** — L399 — `private static void ReadImages(Control c)`
  Reads images.
  Called by: `.Restore()` (same file)
- **`.SaveForm()`** — L486 — `private static void SaveForm(Form ctrl, XmlTextWriter writer)`
  Saves form.
  Called by: `.Save()` (same file)
- **`.RestoreForm()`** — L498 — `private static void RestoreForm(Form ctrl, XmlDocument doc)`
  Restores form.
  Called by: `.Restore()` (same file)
- **`.SaveGroupBox()`** — L539 — `private static void SaveGroupBox(GroupBox ctrl, XmlTextWriter writer)`
  Saves group box.
  Called by: `.Save()` (same file)
- **`.RestoreGroupBox()`** — L551 — `private static void RestoreGroupBox(GroupBox ctrl, XmlDocument doc)`
  Restores group box.
  Called by: `.Restore()` (same file)
- **`.SavePanel()`** — L592 — `private static void SavePanel(Panel ctrl, XmlTextWriter writer)`
  Saves panel.
  Called by: `.Save()` (same file)
- **`.RestorePanel()`** — L602 — `private static void RestorePanel(Panel ctrl, XmlDocument doc)`
  Restores panel.
  Called by: `.Restore()` (same file)
- **`.SaveButton()`** — L637 — `private static void SaveButton(Button ctrl, XmlTextWriter writer)`
  Saves button.
  Called by: `.Save()` (same file)
- **`.RestoreButton()`** — L654 — `private static void RestoreButton(Button ctrl, XmlDocument doc)`
  Restores button.
  Called by: `.Restore()` (same file)
- **`.SetupQuickRecallImages()`** — L710 — `private static void SetupQuickRecallImages(ucQuickRecall ctrl)`
  Setups quick recall images.
  Called by: `.ReadImages()` (same file)
- **`.setupButtonHandlers()`** — L764 — `private static void setupButtonHandlers(Button ctrl)`
  Called by: `.SetupQuickRecallImages()` (same file), `.SetupButtonImages()` (same file)
- **`.SetupButtonImages()`** — L787 — `private static void SetupButtonImages(Button ctrl)`
  Setups button images.
  Called by: `.ReadImages()` (same file)
- **`.Button_StateChanged()`** — L858 — `private static void Button_StateChanged(object sender, EventArgs e)`
  Called by: `.SetupQuickRecallImages()` (same file), `.SetupButtonImages()` (same file), `.Button_MouseUp()` (same file)
- **`.Button_MouseEnter()`** — L881 — `private static void Button_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `Button` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Button_MouseDown()`** — L891 — `private static void Button_MouseDown(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `Button` receives a mouse-down.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.Button_MouseUp()`** — L901 — `private static void Button_MouseUp(object sender, MouseEventArgs e)`
  WinForms event handler: runs when `Button` receives a mouse-up.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetButtonImageState()`** — L906 — `private static void SetButtonImageState(Button ctrl, ImageState state)`
  Sets button image state.
  Called by: `.Button_StateChanged()` (same file), `.Button_MouseEnter()` (same file), `.Button_MouseDown()` (same file)
- **`.SaveCheckBox()`** — L918 — `private static void SaveCheckBox(CheckBox ctrl, XmlTextWriter writer)`
  Saves check box.
  Called by: `.Save()` (same file)
- **`.RestoreCheckBox()`** — L937 — `private static void RestoreCheckBox(CheckBox ctrl, XmlDocument doc)`
  Restores check box.
  Called by: `.Restore()` (same file)
- **`.SetupCheckBoxImages()`** — L1000 — `private static void SetupCheckBoxImages(CheckBox ctrl)`
  Setups check box images.
  Called by: `.ReadImages()` (same file)
- **`.setupCheckBoxHandlers()`** — L1052 — `private static void setupCheckBoxHandlers(CheckBox ctrl)`
  Called by: `.SetupCheckBoxImages()` (same file), `.SetupInfoBar()` (same file)
- **`.SetupInfoBar()`** — L1070 — `private static void SetupInfoBar(ucInfoBar ctrl)`
  Setups info bar.
  Called by: `.ReadImages()` (same file)
- **`.CheckBox_StateChanged()`** — L1162 — `private static void CheckBox_StateChanged(object sender, EventArgs e)`
  Checks box state changed.
  Called by: `.SetupCheckBoxImages()` (same file), `.SetupInfoBar()` (same file)
- **`.CheckBox_MouseEnter()`** — L1187 — `private static void CheckBox_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `CheckBox` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetCheckBoxImageState()`** — L1198 — `private static void SetCheckBoxImageState(CheckBox ctrl, ImageState state)`
  Sets check box image state.
  Called by: `.CheckBox_StateChanged()` (same file), `.CheckBox_MouseEnter()` (same file)
- **`.SaveComboBox()`** — L1210 — `private static void SaveComboBox(ComboBox ctrl, XmlTextWriter writer)`
  Saves combo box.
  Called by: `.Save()` (same file)
- **`.RestoreComboBox()`** — L1224 — `private static void RestoreComboBox(ComboBox ctrl, XmlDocument doc)`
  Restores combo box.
  Called by: `.Restore()` (same file)
- **`.SaveLabel()`** — L1265 — `private static void SaveLabel(Label ctrl, XmlTextWriter writer)`
  Saves label.
  Called by: `.Save()` (same file)
- **`.RestoreLabel()`** — L1280 — `private static void RestoreLabel(Label ctrl, XmlDocument doc)`
  Restores label.
  Called by: `.Restore()` (same file)
- **`.SaveNumericUpDown()`** — L1324 — `private static void SaveNumericUpDown(NumericUpDown ctrl, XmlTextWriter writer)`
  Saves numeric up down.
  Called by: `.Save()` (same file)
- **`.RestoreNumericUpDown()`** — L1338 — `private static void RestoreNumericUpDown(NumericUpDown ctrl, XmlDocument doc)`
  Restores numeric up down.
  Called by: `.Restore()` (same file)
- **`.SavePictureBox()`** — L1379 — `private static void SavePictureBox(PictureBox ctrl, XmlTextWriter writer)`
  Saves picture box.
  Called by: `.Save()` (same file)
- **`.RestorePictureBox()`** — L1391 — `private static void RestorePictureBox(PictureBox ctrl, XmlDocument doc)`
  Restores picture box.
  Called by: `.Restore()` (same file)
- **`.SaveRadioButton()`** — L1426 — `private static void SaveRadioButton(RadioButton ctrl, XmlTextWriter writer)`
  Saves radio button.
  Called by: `.Save()` (same file)
- **`.RestoreRadioButton()`** — L1445 — `private static void RestoreRadioButton(RadioButton ctrl, XmlDocument doc)`
  Restores radio button.
  Called by: `.Restore()` (same file)
- **`.SetupRadioButtonImages()`** — L1508 — `private static void SetupRadioButtonImages(RadioButton ctrl)`
  Setups radio button images.
  Called by: `.ReadImages()` (same file)
- **`.RadioButton_StateChanged()`** — L1576 — `private static void RadioButton_StateChanged(object sender, EventArgs e)`
  Called by: `.SetupRadioButtonImages()` (same file)
- **`.RadioButton_MouseEnter()`** — L1603 — `private static void RadioButton_MouseEnter(object sender, EventArgs e)`
  WinForms event handler: runs when `RadioButton` is entered by the mouse.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.SetRadioButtonImageState()`** — L1614 — `private static void SetRadioButtonImageState(RadioButton ctrl, ImageState state)`
  Sets radio button image state.
  Called by: `.RadioButton_StateChanged()` (same file), `.RadioButton_MouseEnter()` (same file)
- **`.SaveTextBox()`** — L1626 — `private static void SaveTextBox(TextBox ctrl, XmlTextWriter writer)`
  Saves text box.
  Called by: `.Save()` (same file)
- **`.RestoreTextBox()`** — L1642 — `private static void RestoreTextBox(TextBox ctrl, XmlDocument doc)`
  Restores text box.
  Called by: `.Restore()` (same file)
- **`.SavePrettyTrackBar()`** — L1689 — `private static void SavePrettyTrackBar(PrettyTrackBar ctrl, XmlTextWriter writer)`
  Saves pretty track bar.
  Called by: `.Save()` (same file)
- **`.RestorePrettyTrackBar()`** — L1700 — `private static void RestorePrettyTrackBar(PrettyTrackBar ctrl, XmlDocument doc)`
  Restores pretty track bar.
  Called by: `.Restore()` (same file)
- **`.SetupPrettyTrackBarImages()`** — L1728 — `private static void SetupPrettyTrackBarImages(PrettyTrackBar ctrl)`
  Setups pretty track bar images.
  Called by: `.ReadImages()` (same file)
- **`.SaveSize()`** — L1768 — `private static void SaveSize(Size s, XmlTextWriter writer)`
  Saves size.
  Called by: `.SaveForm()` (same file), `.SaveGroupBox()` (same file), `.SavePanel()` (same file), `.SaveButton()` (same file), `.SaveCheckBox()` (same file), `.SaveComboBox()` (same file) — and 6 more
- **`.RestoreSize()`** — L1776 — `private static Size RestoreSize(XmlNode node)`
  Restores size.
  Called by: `.RestoreForm()` (same file), `.RestoreGroupBox()` (same file), `.RestorePanel()` (same file), `.RestoreButton()` (same file), `.RestoreCheckBox()` (same file), `.RestoreComboBox()` (same file) — and 6 more
- **`.SaveFont()`** — L1794 — `private static void SaveFont(Font f, XmlTextWriter writer)`
  Saves font.
  Called by: `.SaveForm()` (same file), `.SaveGroupBox()` (same file), `.SaveButton()` (same file), `.SaveCheckBox()` (same file), `.SaveComboBox()` (same file), `.SaveLabel()` (same file) — and 3 more
- **`.RestoreFont()`** — L1805 — `private static Font RestoreFont(XmlNode node)`
  Restores font.
  Called by: `.RestoreForm()` (same file), `.RestoreGroupBox()` (same file), `.RestoreButton()` (same file), `.RestoreCheckBox()` (same file), `.RestoreComboBox()` (same file), `.RestoreLabel()` (same file) — and 3 more
- **`.SaveLocation()`** — L1840 — `private static void SaveLocation(Point p, XmlTextWriter writer)`
  Saves location.
  Called by: `.SaveGroupBox()` (same file), `.SavePanel()` (same file), `.SaveButton()` (same file), `.SaveCheckBox()` (same file), `.SaveComboBox()` (same file), `.SaveLabel()` (same file) — and 5 more
- **`.RestoreLocation()`** — L1848 — `private static Point RestoreLocation(XmlNode node)`
  Restores location.
  Called by: `.RestoreGroupBox()` (same file), `.RestorePanel()` (same file), `.RestoreButton()` (same file), `.RestoreCheckBox()` (same file), `.RestoreComboBox()` (same file), `.RestoreLabel()` (same file) — and 5 more
- **`.SaveFlatAppearance()`** — L1866 — `private static void SaveFlatAppearance(FlatButtonAppearance fa, XmlTextWriter writer)`
  Saves flat appearance.
  Called by: `.SaveButton()` (same file), `.SaveCheckBox()` (same file), `.SaveRadioButton()` (same file)
- **`.StringToColor()`** — L1874 — `private static Color StringToColor(string s)`
  Called by: `.RestoreForm()` (same file), `.RestoreGroupBox()` (same file), `.RestorePanel()` (same file), `.RestoreButton()` (same file), `.RestoreCheckBox()` (same file), `.RestoreComboBox()` (same file) — and 6 more
- **`.SetBackgroundImage()`** — L1882 — `private static void SetBackgroundImage(Control c)`
  Sets background image.
  Called by: `.ReadImages()` (same file)
- **`.loadImage()`** — L1921 — `private static Image loadImage(string path)`
  Called by: `.Restore()` (same file), `.SetupQuickRecallImages()` (same file), `.SetupButtonImages()` (same file), `.SetupCheckBoxImages()` (same file), `.SetupInfoBar()` (same file), `.SetupRadioButtonImages()` (same file) — and 2 more
- **`.getImageFromFilePath()`** — L1952 — `private static Image getImageFromFilePath(string path)`
  Returns image from file path.
  Called by: `.SetupButtonImages()` (same file), `.SetupCheckBoxImages()` (same file), `.SetupRadioButtonImages()` (same file)
- **`.computeHashFromImage()`** — L1959 — `private static string computeHashFromImage(Image image)`
  Called by: `.loadImage()` (same file)
- **`.resizeImage()`** — L1995 — `private static Image resizeImage(Image image, Control c)`
  Called by: `.SetupButtonImages()` (same file), `.SetupCheckBoxImages()` (same file), `.SetupRadioButtonImages()` (same file)

#### `ImageState` (type, L67)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Skin.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
