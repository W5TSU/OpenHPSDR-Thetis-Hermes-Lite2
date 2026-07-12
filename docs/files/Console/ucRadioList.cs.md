# `Console/ucRadioList.cs`

**Functional area:** [3. HPSDR network protocol and radio discovery](../../CODE_OUTLINE.md#3-hpsdr-network-protocol-and-radio-discovery)

**Role:** UI for picking among discovered radios and defining custom/static radio addresses.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `Console/HPSDR/clsRadioDiscovery.cs` (references ×12)

## Outline

### Types

#### `ucRadioList` (type, L57)

- `.DoesRadioExist()` — L391
- `.RadioConnected()` — L406
- `.RadioDisconnected()` — L429
- `.DisconnectAll()` — L468
- `.PLLLocked()` — L498
- `.UpdateSelectedDetails()` — L528
- `.AddRadio()` — L546
- `.RemoveRadio()` — L663
- `.ClearRadios()` — L695
- `.MakeRadioVisible()` — L715
- `.SaveToJson()` — L744
- `.LoadFromJson()` — L783
- `.IsInputKey()` — L903
- `.OnKeyDown()` — L910
- `.OnMouseWheel()` — L951
- `.OnMouseMove()` — L967
- `.OnMouseLeave()` — L983
- `.OnMouseDown()` — L995
- `.OnMouseUp()` — L1018
- `.OnFontChanged()` — L1050
- `.OnPaint()` — L1061
- `.drawRow()` — L1117
- `.drawRadioGlyph()` — L1238
- `.drawTrashGlyph()` — L1258
- `.buildLine1()` — L1280
- `.buildLine2()` — L1328
- `.buildNicLine3()` — L1387
- `.buildNicLine4()` — L1410
- `.hitTest()` — L1457
- `.setSelectedByIndex()` — L1489
- `.selectedIndex()` — L1508
- `.getSelectedItem()` — L1523
- `.indexOfKey()` — L1538
- `.getViewportRect()` — L1553
- `.normalRowHeight()` — L1565
- `.compactRowHeight()` — L1573
- `.autoRowIsCompact()` — L1581
- `.rowHeightForIndex()` — L1598
- `.contentHeight()` — L1608
- `.rowTopForIndex()` — L1620
- `.indexFromContentY()` — L1634
- `.scale()` — L1650
- `.updateScroll()` — L1667
- `.clampScrollValue()` — L1696
- `.setScrollValue()` — L1711
- `.scroll_ValueChanged()` — L1730
- `.ucRadioList_SizeChanged()` — L1735
- `.raiseSelectedChanged()` — L1742
- `.raiseListChanged()` — L1748
- `.ensureGuidString()` — L1754
- `.buildKey()` — L1768
- `.isLegacyMacOnlyKey()` — L1791
- `.isAutoKey()` — L1802
- `.ensureAutoEntry()` — L1807
- `.resetAutoItem()` — L1836
- `.fillItemFromInfo()` — L1886
- `.buildVersionText()` — L1960
- `.safe()` — L1992
- `.enc()` — L1998
- `.dec()` — L2005

#### `RowItem` (type, L64)

_No extracted members._

#### `PersistModel` (type, L115)

_No extracted members._

#### `PersistRow` (type, L122)

_No extracted members._

#### `HitTestResult` (type, L144)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucRadioList.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
