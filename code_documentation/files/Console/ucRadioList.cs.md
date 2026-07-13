# `Console/ucRadioList.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** UI for picking among discovered radios and defining custom/static radio addresses.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/HPSDR/clsRadioDiscovery.cs` (references ×12)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `ucRadioList` (type, L57)

- **`.DoesRadioExist()`** — L391 — `public bool DoesRadioExist(string radioKey)`
  Called by: `.AddRadio()` (same file), `.LoadFromJson()` (same file), `.ensureAutoEntry()` (same file)
- **`.RadioConnected()`** — L406 — `public void RadioConnected(string radioKey)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RadioDisconnected()`** — L429 — `public void RadioDisconnected(string radioKey)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisconnectAll()`** — L468 — `public void DisconnectAll()`
  Disconnects all.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PLLLocked()`** — L498 — `public void PLLLocked(string radioKey, bool locked)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateSelectedDetails()`** — L528 — `public void UpdateSelectedDetails(NicRadioScanResult nic, RadioInfo radio)`
  Updates selected details.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddRadio()`** — L546 — `public string AddRadio(NicRadioScanResult nic, RadioInfo radio)`
  Adds radio.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveRadio()`** — L663 — `public bool RemoveRadio(string radioKey)`
  Removes radio.
  Called by: `.OnKeyDown()` (same file), `.OnMouseUp()` (same file)
- **`.ClearRadios()`** — L695 — `public void ClearRadios()`
  Clears radios.
  Called by: `.LoadFromJson()` (same file)
- **`.MakeRadioVisible()`** — L715 — `public bool MakeRadioVisible(string radioKey)`
  Called by: `.setSelectedByIndex()` (same file)
- **`.SaveToJson()`** — L744 — `public string SaveToJson()`
  Saves to json.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadFromJson()`** — L783 — `public void LoadFromJson(string json)`
  Loads from json.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsInputKey()`** — L903 — `protected override bool IsInputKey(Keys keyData)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnKeyDown()`** — L910 — `protected override void OnKeyDown(KeyEventArgs e)`
  Handles/raises the key down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseWheel()`** — L951 — `protected override void OnMouseWheel(MouseEventArgs e)`
  Handles/raises the mouse wheel event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L967 — `protected override void OnMouseMove(MouseEventArgs e)`
  Handles/raises the mouse move event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseLeave()`** — L983 — `protected override void OnMouseLeave(EventArgs e)`
  Handles/raises the mouse leave event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L995 — `protected override void OnMouseDown(MouseEventArgs e)`
  Handles/raises the mouse down event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseUp()`** — L1018 — `protected override void OnMouseUp(MouseEventArgs e)`
  Handles/raises the mouse up event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFontChanged()`** — L1050 — `protected override void OnFontChanged(EventArgs e)`
  Handles/raises the font changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPaint()`** — L1061 — `protected override void OnPaint(PaintEventArgs e)`
  Handles/raises the paint event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawRow()`** — L1117 — `private void drawRow(Graphics g, Rectangle rowRect, RowItem item, bool selected, bool hovered, bool hoverTrash, bool canRemove, bool compact)`
  Called by: `.OnPaint()` (same file)
- **`.drawRadioGlyph()`** — L1238 — `private void drawRadioGlyph(Graphics g, Rectangle rect, bool selected)`
  Called by: `.drawRow()` (same file)
- **`.drawTrashGlyph()`** — L1258 — `private void drawTrashGlyph(Graphics g, Rectangle rect, bool hot)`
  Called by: `.drawRow()` (same file)
- **`.buildLine1()`** — L1280 — `private string buildLine1(RowItem item)`
  Called by: `.drawRow()` (same file)
- **`.buildLine2()`** — L1328 — `private string buildLine2(RowItem item)`
  Called by: `.drawRow()` (same file)
- **`.buildNicLine3()`** — L1387 — `private string buildNicLine3(RowItem item)`
  Called by: `.drawRow()` (same file)
- **`.buildNicLine4()`** — L1410 — `private string buildNicLine4(RowItem item)`
  Called by: `.drawRow()` (same file)
- **`.hitTest()`** — L1457 — `private HitTestResult hitTest(Point p)`
  Called by: `.OnMouseMove()` (same file), `.OnMouseDown()` (same file), `.OnMouseUp()` (same file)
- **`.setSelectedByIndex()`** — L1489 — `private void setSelectedByIndex(int idx, bool ensureVisible)`
  Sets selected by index.
  Called by: `.OnKeyDown()` (same file), `.OnMouseDown()` (same file)
- **`.selectedIndex()`** — L1508 — `private int selectedIndex()`
  Called by: `.OnKeyDown()` (same file)
- **`.getSelectedItem()`** — L1523 — `private RowItem getSelectedItem()`
  Returns selected item.
  Called by: `.UpdateSelectedDetails()` (same file)
- **`.indexOfKey()`** — L1538 — `private int indexOfKey(string key)`
  Called by: `.RadioConnected()` (same file), `.RadioDisconnected()` (same file), `.PLLLocked()` (same file), `.RemoveRadio()` (same file), `.MakeRadioVisible()` (same file)
- **`.getViewportRect()`** — L1553 — `private Rectangle getViewportRect()`
  Returns viewport rect.
  Called by: `.MakeRadioVisible()` (same file), `.OnPaint()` (same file), `.hitTest()` (same file), `.updateScroll()` (same file)
- **`.normalRowHeight()`** — L1565 — `private int normalRowHeight()`
  Called by: `.OnMouseWheel()` (same file), `.rowHeightForIndex()` (same file), `.updateScroll()` (same file)
- **`.compactRowHeight()`** — L1573 — `private int compactRowHeight()`
  Called by: `.rowHeightForIndex()` (same file)
- **`.autoRowIsCompact()`** — L1581 — `private bool autoRowIsCompact(RowItem item)`
  Called by: `.OnPaint()` (same file), `.rowHeightForIndex()` (same file)
- **`.rowHeightForIndex()`** — L1598 — `private int rowHeightForIndex(int idx)`
  Called by: `.MakeRadioVisible()` (same file), `.OnPaint()` (same file), `.hitTest()` (same file), `.contentHeight()` (same file), `.rowTopForIndex()` (same file), `.indexFromContentY()` (same file)
- **`.contentHeight()`** — L1608 — `private int contentHeight()`
  Called by: `.updateScroll()` (same file)
- **`.rowTopForIndex()`** — L1620 — `private int rowTopForIndex(int idx)`
  Called by: `.MakeRadioVisible()` (same file), `.OnPaint()` (same file), `.hitTest()` (same file)
- **`.indexFromContentY()`** — L1634 — `private int indexFromContentY(int y)`
  Called by: `.OnPaint()` (same file), `.hitTest()` (same file)
- **`.scale()`** — L1650 — `private int scale(int px)`
  Called by: `.drawRow()` (same file), `.hitTest()` (same file), `.normalRowHeight()` (same file), `.compactRowHeight()` (same file)
- **`.updateScroll()`** — L1667 — `private void updateScroll()`
  Called by: `.RadioDisconnected()` (same file), `.DisconnectAll()` (same file), `.UpdateSelectedDetails()` (same file), `.AddRadio()` (same file), `.RemoveRadio()` (same file), `.ClearRadios()` (same file) — and 3 more
- **`.clampScrollValue()`** — L1696 — `private void clampScrollValue()`
  Called by: `.RadioDisconnected()` (same file), `.DisconnectAll()` (same file), `.UpdateSelectedDetails()` (same file), `.RemoveRadio()` (same file), `.LoadFromJson()` (same file), `.OnFontChanged()` (same file) — and 2 more
- **`.setScrollValue()`** — L1711 — `private void setScrollValue(int value)`
  Sets scroll value.
  Called by: `.ClearRadios()` (same file), `.MakeRadioVisible()` (same file), `.OnMouseWheel()` (same file)
- **`.scroll_ValueChanged()`** — L1730 — `private void scroll_ValueChanged(object sender, EventArgs e)`
  WinForms event handler: runs when `scroll` value changes.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.ucRadioList_SizeChanged()`** — L1735 — `private void ucRadioList_SizeChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.raiseSelectedChanged()`** — L1742 — `private void raiseSelectedChanged()`
  Called by: `.UpdateSelectedDetails()` (same file), `.RemoveRadio()` (same file), `.ClearRadios()` (same file), `.LoadFromJson()` (same file), `.setSelectedByIndex()` (same file)
- **`.raiseListChanged()`** — L1748 — `private void raiseListChanged()`
  Called by: `.UpdateSelectedDetails()` (same file), `.AddRadio()` (same file), `.RemoveRadio()` (same file), `.ClearRadios()` (same file), `.LoadFromJson()` (same file)
- **`.ensureGuidString()`** — L1754 — `private string ensureGuidString(string guid)`
  Called by: `.AddRadio()` (same file), `.LoadFromJson()` (same file), `.buildKey()` (same file), `.fillItemFromInfo()` (same file)
- **`.buildKey()`** — L1768 — `private string buildKey(string mac, string guid, bool isCustom, IPAddress ip, int port, string nicId, RadioDiscoveryRadioProtocol proto)`
  Called by: `.AddRadio()` (same file), `.LoadFromJson()` (same file)
- **`.isLegacyMacOnlyKey()`** — L1791 — `private bool isLegacyMacOnlyKey(string key)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isAutoKey()`** — L1802 — `private bool isAutoKey(string key)`
  Called by: `.RadioDisconnected()` (same file), `.DisconnectAll()` (same file), `.AddRadio()` (same file), `.RemoveRadio()` (same file), `.SaveToJson()` (same file), `.LoadFromJson()` (same file) — and 7 more
- **`.ensureAutoEntry()`** — L1807 — `private void ensureAutoEntry(bool selectAutoIfMissingSelection)`
  Called by: `.AddRadio()` (same file), `.RemoveRadio()` (same file), `.ClearRadios()` (same file), `.LoadFromJson()` (same file)
- **`.resetAutoItem()`** — L1836 — `private void resetAutoItem(RowItem item)`
  Called by: `.RadioDisconnected()` (same file), `.DisconnectAll()` (same file), `.ensureAutoEntry()` (same file)
- **`.fillItemFromInfo()`** — L1886 — `private void fillItemFromInfo(RowItem item, NicRadioScanResult nic, RadioInfo radio)`
  Called by: `.UpdateSelectedDetails()` (same file)
- **`.buildVersionText()`** — L1960 — `private string buildVersionText(RadioInfo radio)`
  Called by: `.AddRadio()` (same file), `.fillItemFromInfo()` (same file)
- **`.safe()`** — L1992 — `private string safe(string s)`
  Called by: `.AddRadio()` (same file), `.buildLine1()` (same file), `.buildLine2()` (same file), `.buildNicLine3()` (same file), `.buildNicLine4()` (same file), `.ensureGuidString()` (same file) — and 6 more
- **`.enc()`** — L1998 — `private string enc(string s)`
  Called by: `.SaveToJson()` (same file)
- **`.dec()`** — L2005 — `private string dec(string s)`
  Called by: `.LoadFromJson()` (same file)

#### `RowItem` (type, L64)

_No extracted members._

#### `PersistModel` (type, L115)

_No extracted members._

#### `PersistRow` (type, L122)

_No extracted members._

#### `HitTestResult` (type, L144)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucRadioList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
