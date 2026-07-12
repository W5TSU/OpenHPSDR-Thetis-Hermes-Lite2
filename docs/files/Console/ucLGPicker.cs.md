# `Console/ucLGPicker.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers).

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (references ×16, calls ×2)
  - `Console/MeterManager.cs` (calls ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.InterpolateBetween()` (×3)

## Outline

### Types

#### `ucLGPicker` (type, L54)

- `.rebuildSortedColours()` — L135
- `.enableGripper()` — L150
- `.addColour()` — L161
- `.drawTextCentre()` — L180
- `.drawScales()` — L189
- `.LGPicker_Paint()` — L213
- `.findColourIndexGrip()` — L279
- `.LGPicker_MouseMove()` — L296
- `.highlightGripper()` — L364
- `.percFromPixels()` — L378
- `.LGPicker_MouseDown()` — L387
- `.LGPicker_MouseUp()` — L447
- `.indexOfClosestToLeft()` — L469
- `.indexAtPerc()` — L499
- `.indexOfClosestToRight()` — L513
- `.LGPicker_MouseLeave()` — L544
- `.setColour()` — L566
- `.RemoveSelectedGripper()` — L582
- `.addGripper()` — L599
- `.findFreeDisabledGripper()` — L625
- `.HighlightFirstGripper()` — L639
- `.Clear()` — L653
- `.OnChanged()` — L740
- `.OnGripperSelected()` — L744
- `.OnGripperMouseEnter()` — L748
- `.OnGripperMouseLeave()` — L752
- `.OnGripperDBMChanged()` — L756
- `.LGPicker_EnabledChanged()` — L760
- `.GetColourForDBM()` — L764
- `.GetColourAtPercent()` — L770
- `.GetPercForDBM()` — L793
- `.addColourGradientData()` — L804
- `.GetColourGradientDataForDBMRange()` — L813
- `.ApplyGlobalAlpha()` — L891

#### `ColourGradientData` (type, L70)

_No extracted members._

#### `GradColours` (type, L75)

_No extracted members._

#### `ColourEventArgs` (type, L958)

_No extracted members._

#### `GripperEventArgs` (type, L963)

_No extracted members._

#### `ColorInterpolator` (type, L971)

- `.InterpolateBetween()` — L979
- `.InterpolateComponent()` — L999

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ucLGPicker.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
