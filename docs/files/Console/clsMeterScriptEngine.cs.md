# `Console/clsMeterScriptEngine.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** Scripting engine for user-defined meter faces/behaviors.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×8)
  - `Console/setup.cs` (calls ×4)
- Uses (outgoing references to other files):
  - `Console/MeterManager.cs` (references ×1)
- Most-referenced symbols from other files: `.BeginBatch()` (×4), `.EndBatch()` (×4), `.Stop()` (×1), `.UnregisterLed()` (×1), `.SetCondition()` (×1), `.ReadResult()` (×1)

## Outline

### Types

#### `MeterScriptEngine` (type, L53)

- `.Start()` — L139
- `.Stop()` — L172
- `.BeginBatch()` — L207
- `.EndBatch()` — L215
- `.RegisterLed()` — L245
- `.UnregisterLed()` — L285
- `.SetCondition()` — L325
- `.SetUpdateInterval()` — L382
- `.ReadResult()` — L395
- `.ReadError()` — L405
- `.ReadDiagnostic()` — L415
- `.SetLoopIntervalFloor()` — L425
- `.tick()` — L436
- `.build_globals_once()` — L527
- `.get_due_indices_nolock()` — L545
- `.recompute_loop_interval_nolock()` — L558
- `.schedule_compile_if_needed_nolock()` — L574
- `.compile_worker_entry()` — L581
- `.compile_one()` — L686
- `.warmup_roslyn_entry()` — L729

#### `BankVars` (type, L55)

_No extracted members._

#### `Snapshot` (type, L71)

_No extracted members._

#### `Globals` (type, L77)

_No extracted members._

#### `CompileResult` (type, L82)

_No extracted members._

#### `EvalResult` (type, L89)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsMeterScriptEngine.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
