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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `MeterScriptEngine` (type, L53)

- **`.Start()`** — L139 — `public static void Start(Func<Snapshot> variable_provider_banked, int default_interval_ms, int bank_count)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Stop()`** — L172 — `public static void Stop()`
  Called by: `.Shutdown()` (`Console/MeterManager.cs`)
- **`.BeginBatch()`** — L207 — `public static void BeginBatch()`
  Called by: `.RestoreSettings()` (`Console/MeterManager.cs`), `.removeMeterItem()` (`Console/MeterManager.cs`), `.btnContainer_load_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.EndBatch()`** — L215 — `public static void EndBatch()`
  Called by: `.RestoreSettings()` (`Console/MeterManager.cs`), `.removeMeterItem()` (`Console/MeterManager.cs`), `.btnContainer_load_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.RegisterLed()`** — L245 — `public static int RegisterLed()`
  Registers led.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UnregisterLed()`** — L285 — `public static void UnregisterLed(int index)`
  Unregisters led.
  Called by: `.Removing()` (`Console/MeterManager.cs`)
- **`.SetCondition()`** — L325 — `public static bool SetCondition(int index, string condition)`
  _diagnostics[index] = string.Empty; _errors[index] = false; _needs_compile[index] = true; _needs_recompile = true; if (_batch_depth == 0) schedule_compile_if_needed_nolock(); } return true; }
  Called by: `.onTimerElapsedCondition()` (`Console/MeterManager.cs`)
- **`.SetUpdateInterval()`** — L382 — `public static void SetUpdateInterval(int index, int milliseconds)`
  Sets update interval.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadResult()`** — L395 — `public static bool ReadResult(int index)`
  Reads result.
  Called by: `.Update()` (`Console/MeterManager.cs`)
- **`.ReadError()`** — L405 — `public static bool ReadError(int index)`
  Reads error.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadDiagnostic()`** — L415 — `public static string ReadDiagnostic(int index)`
  Reads diagnostic.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLoopIntervalFloor()`** — L425 — `public static void SetLoopIntervalFloor(int milliseconds)`
  Sets loop interval floor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tick()`** — L436 — `private static void tick()`
  Called by: `.Start()` (same file)
- **`.build_globals_once()`** — L527 — `private static Globals build_globals_once()`
  Called by: `.tick()` (same file)
- **`.get_due_indices_nolock()`** — L545 — `private static List<int> get_due_indices_nolock()`
  Returns due indices nolock.
  Called by: `.tick()` (same file)
- **`.recompute_loop_interval_nolock()`** — L558 — `private static void recompute_loop_interval_nolock()`
  Called by: `.EndBatch()` (same file), `.RegisterLed()` (same file), `.UnregisterLed()` (same file), `.SetUpdateInterval()` (same file)
- **`.schedule_compile_if_needed_nolock()`** — L574 — `private static void schedule_compile_if_needed_nolock()`
  Called by: `.EndBatch()` (same file), `.SetCondition()` (same file), `.tick()` (same file)
- **`.compile_worker_entry()`** — L581 — `private static void compile_worker_entry()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.compile_one()`** — L686 — `private static CompileResult compile_one(string expr_trimmed)`
  Called by: `.compile_worker_entry()` (same file)
- **`.warmup_roslyn_entry()`** — L729 — `private static void warmup_roslyn_entry()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

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
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsMeterScriptEngine.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
