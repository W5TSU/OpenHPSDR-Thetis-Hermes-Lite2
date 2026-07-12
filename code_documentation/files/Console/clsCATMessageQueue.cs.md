# `Console/clsCATMessageQueue.cs`

**Functional area:** [10. CAT control and external program interfaces](../../CODE_OUTLINE.md#10-cat-control-and-external-program-interfaces)

**Role:** Queued/asynchronous CAT message handling and scripted ("atomic") CAT command sequences.

## How this file is used

- Used by (incoming references from other files):
  - `Console/MeterManager.cs` (calls ×7, references ×1)
- Uses (outgoing references to other files):
  - `Console/clsCatAtonic.cs` (references ×2)
- Most-referenced symbols from other files: `.add()` (×2), `.createBatch()` (×2), `.sendBatch()` (×2), `.stopAll()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `CATQueueBatching` (namespace, L48)

_No extracted members._

#### `MessageQueueSystem` (type, L50)

- **`.createBatch()`** — L137 — `public MessageBatch createBatch()`
  Called by: `.OnContainerVisible()` (`Console/MeterManager.cs`), `.handleMacroButtonPress()` (`Console/MeterManager.cs`)
- **`.sendBatch()`** — L142 — `public void sendBatch(int queue_index, MessageBatch batch)`
  Called by: `.OnContainerVisible()` (`Console/MeterManager.cs`), `.handleMacroButtonPress()` (`Console/MeterManager.cs`)
- **`.isBusy()`** — L156 — `public bool isBusy(int queue_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getPending()`** — L163 — `public int getPending(int queue_index)`
  Returns pending.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isEmpty()`** — L170 — `public bool isEmpty(int queue_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.stopAndClearQueue()`** — L177 — `public void stopAndClearQueue(int queue_index)`
  Called by: `.stopAll()` (same file)
- **`.stopAll()`** — L199 — `public void stopAll()`
  Called by: `.Shutdown()` (`Console/MeterManager.cs`)
- **`.startQueue()`** — L208 — `private void startQueue(int index)`
  Called by: `.ensureQueueRunning()` (same file)
- **`.ensureQueueRunning()`** — L226 — `private void ensureQueueRunning(int index)`
  Called by: `.sendBatch()` (same file)
- **`.workerLoop()`** — L231 — `private void workerLoop(int index, CancellationToken token)`
  Called by: `.startQueue()` (same file)
- **`.drainQueue()`** — L277 — `private void drainQueue(int index)`
  Called by: `.stopAndClearQueue()` (same file)
- **`.throwIfDisposed()`** — L284 — `private void throwIfDisposed()`
  Called by: `.sendBatch()` (same file), `.isBusy()` (same file), `.getPending()` (same file), `.isEmpty()` (same file), `.stopAndClearQueue()` (same file), `.stopAll()` (same file)
- **`.Dispose()`** — L289 — `public void Dispose()`
  Releases the object’s resources.
  Called by: `.startQueue()` (same file)

#### `MessageBatch` (type, L55)

- **`.add()`** — L66 — `public Guid add(ScriptCommand cmd)`
  Called by: `.OnContainerVisible()` (`Console/MeterManager.cs`), `.handleMacroButtonPress()` (`Console/MeterManager.cs`)
- **`.isEmpty()`** — L76 — `internal bool isEmpty()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.snapshot()`** — L84 — `internal QueuedItem[] snapshot()`
  Called by: `.sendBatch()` (same file)

#### `QueuedItem` (type, L93)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCATMessageQueue.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
