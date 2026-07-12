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

### Types

#### `CATQueueBatching` (namespace, L48)

_No extracted members._

#### `MessageQueueSystem` (type, L50)

- `.createBatch()` — L137
- `.sendBatch()` — L142
- `.isBusy()` — L156
- `.getPending()` — L163
- `.isEmpty()` — L170
- `.stopAndClearQueue()` — L177
- `.stopAll()` — L199
- `.startQueue()` — L208
- `.ensureQueueRunning()` — L226
- `.workerLoop()` — L231
- `.drainQueue()` — L277
- `.throwIfDisposed()` — L284
- `.Dispose()` — L289

#### `MessageBatch` (type, L55)

- `.add()` — L66
- `.isEmpty()` — L76
- `.snapshot()` — L84

#### `QueuedItem` (type, L93)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/clsCATMessageQueue.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
