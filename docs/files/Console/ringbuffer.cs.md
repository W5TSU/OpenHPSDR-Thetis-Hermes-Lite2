# `Console/ringbuffer.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** Managed ring buffer used by audio record/playback.

## How this file is used

- Used by (incoming references from other files):
  - `Console/clsAudioRecordPlayback.cs` (calls ×7, references ×2)
  - `Console/cwx.cs` (calls ×5, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.ReadSpace()` (×3), `.Read()` (×2), `.Write()` (×2), `.Reset()` (×2), `.WritePtr()` (×1), `.WriteSpace()` (×1), `.ReadPtr()` (×1)

## Outline

### Types

#### `RingBufferFloat` (type, L38)

- `.npoof2()` — L65
- `.nblock2()` — L78
- `.ReadSpace()` — L87
- `.WriteSpace()` — L98
- `.Read()` — L112
- `.ReadPtr()` — L148
- `.Write()` — L185
- `.WritePtr()` — L222
- `.Reset()` — L256
- `.Clear()` — L266
- `.Restart()` — L277
- `.Peek()` — L283
- `.ReadAdvance()` — L339
- `.WriteAdvance()` — L344

#### `RingBufferByte` (type, L352)

- `.npoof2()` — L379
- `.nblock2()` — L392
- `.ReadSpace()` — L401
- `.WriteSpace()` — L412
- `.Read()` — L426
- `.ReadPtr()` — L463
- `.Write()` — L500
- `.WritePtr()` — L537
- `.Reset()` — L571
- `.Clear()` — L581
- `.Restart()` — L592
- `.Peek()` — L598
- `.ReadAdvance()` — L654
- `.WriteAdvance()` — L659

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ringbuffer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
