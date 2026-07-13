# `Console/enums.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Shared enumerations (bands, modes, meter types, display modes, etc.) used across the whole console.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×106)
  - `Console/MeterManager.cs` (references ×83)
  - `Console/TCIServer.cs` (references ×37)
  - `Console/setup.cs` (references ×35)
  - `Console/clsBandStackManager.cs` (references ×24)
  - `Console/display.cs` (references ×14)
  - `Console/filter.cs` (references ×10)
  - `Console/HPSDR/Alex.cs` (references ×9)
  - `Console/HPSDR/Penny.cs` (references ×9)
  - `Console/PanDisplay.cs` (references ×8)
  - `Console/wbDisplay.cs` (references ×8)
  - `Console/CAT/CATCommands.cs` (references ×7)
  - …and 17 more files
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `FocusMasterMode` (type, L48)

_No extracted members._

#### `FWCAnt` (type, L56)

_No extracted members._

#### `ColorScheme` (type, L68)

_No extracted members._

#### `MultiMeterDisplayMode` (type, L81)

_No extracted members._

#### `MultiMeterMeasureMode` (type, L87)

_No extracted members._

#### `FilterWidthMode` (type, L96)

_No extracted members._

#### `DisplayEngine` (type, L103)

_No extracted members._

#### `HPSDRModel` (type, L109)

_No extracted members._

#### `DisplayMode` (type, L134)

_No extracted members._

#### `AGCMode` (type, L152)

_No extracted members._

#### `MeterRXMode` (type, L164)

_No extracted members._

#### `MeterTXMode` (type, L177)

_No extracted members._

#### `KeyerLine` (type, L198)

_No extracted members._

#### `FRSRegion` (type, L205)

_No extracted members._

#### `PreampMode` (type, L236)

_No extracted members._

#### `DSPMode` (type, L253)

_No extracted members._

#### `Band` (type, L273)

_No extracted members._

#### `Filter` (type, L328)

_No extracted members._

#### `PTTMode` (type, L347)

_No extracted members._

#### `DisplayLabelAlignment` (type, L362)

_No extracted members._

#### `ClickTuneMode` (type, L373)

_No extracted members._

#### `FMTXMode` (type, L381)

_No extracted members._

#### `HPSDRHW` (type, L389)

_No extracted members._

#### `DSPFilterType` (type, L404)

_No extracted members._

#### `DisplayRegion` (type, L410)

_No extracted members._

#### `BreakIn` (type, L430)

_No extracted members._

#### `RadioProtocol` (type, L437)

_No extracted members._

#### `TXPinActions` (type, L445)

_No extracted members._

#### `DrivePowerSource` (type, L458)

_No extracted members._

#### `SquelchState` (type, L465)

_No extracted members._

#### `StatusBarIconGroup` (type, L473)

_No extracted members._

#### `VFOSYNCinit` (type, L483)

_No extracted members._

#### `PAstatusIndicatorState` (type, L490)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/enums.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
