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

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `RingBufferFloat` (type, L38)

- **`.npoof2()`** — L65 — `public int npoof2(int n)`
  returns the power of 2 that is equal/larger than n
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.nblock2()`** — L78 — `public int nblock2(int n)`
  returns the next power of 2 larger/equal to n
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadSpace()`** — L87 — `public int ReadSpace()`
  Get the number of elements available to be read from the ringbuffer.
  Called by: `.ProcessRecordBuffers()` (`Console/clsAudioRecordPlayback.cs`), `.GetPlayBuffer()` (`Console/clsAudioRecordPlayback.cs`)
- **`.WriteSpace()`** — L98 — `public int WriteSpace()`
  Get the number of elements that will fit into the ringbuffer.
  Called by: `.ProcessBuffers()` (`Console/clsAudioRecordPlayback.cs`)
- **`.Read()`** — L112 — `public int Read(float[] dest, int cnt)`
  Reads data out of the ringbuffer into the dest array.
  Called by: `.WriteBuffer()` (`Console/clsAudioRecordPlayback.cs`)
- **`.ReadPtr()`** — L148 — `public int ReadPtr(float* dest, int cnt)`
  Read elements out of the ringbuffer into the array pointed to by dest.
  Called by: `.GetPlayBuffer()` (`Console/clsAudioRecordPlayback.cs`)
- **`.Write()`** — L185 — `public int Write(float[] src, int cnt)`
  Writes from the src array into the ringbuffer.
  Called by: `.ReadBuffer()` (`Console/clsAudioRecordPlayback.cs`)
- **`.WritePtr()`** — L222 — `public int WritePtr(float* src, int cnt)`
  Writes from the array pointed to by src into the ringbuffer.
  Called by: `.AddWriteBuffer()` (`Console/clsAudioRecordPlayback.cs`)
- **`.Reset()`** — L256 — `public void Reset()`
  Resets the ringbuffer pointers (will be empty afterwards).
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Clear()`** — L266 — `public void Clear(int nfloats)`
  Zero the data in the buffer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Restart()`** — L277 — `public void Restart(int nfloats)`
  Reset the pointers and zero the actual data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Peek()`** — L283 — `public int Peek(float[] dest, int cnt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadAdvance()`** — L339 — `public void ReadAdvance(int cnt)`
  Reads advance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WriteAdvance()`** — L344 — `public void WriteAdvance(int cnt)`
  Writes advance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RingBufferByte` (type, L352)

- **`.npoof2()`** — L379 — `public int npoof2(int n)`
  returns the power of 2 that is equal/larger than n
  Called by: `.nblock2()` (same file)
- **`.nblock2()`** — L392 — `public int nblock2(int n)`
  returns the next power of 2 larger/equal to n
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadSpace()`** — L401 — `public int ReadSpace()`
  Get the number of elements available to be read from the ringbuffer.
  Called by: `.Read()` (same file), `.ReadPtr()` (same file), `.Peek()` (same file), `.SendBufferMessage()` (`Console/cwx.cs`)
- **`.WriteSpace()`** — L412 — `public int WriteSpace()`
  Get the number of elements that will fit into the ringbuffer.
  Called by: `.Write()` (same file), `.WritePtr()` (same file)
- **`.Read()`** — L426 — `public int Read(byte[] dest, int cnt)`
  Reads data out of the ringbuffer into the dest array.
  Called by: `.SendBufferMessage()` (`Console/cwx.cs`)
- **`.ReadPtr()`** — L463 — `public int ReadPtr(byte* dest, int cnt)`
  Read elements out of the ringbuffer into the array pointed to by dest.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Write()`** — L500 — `public int Write(byte[] src, int cnt)`
  Writes from the src array into the ringbuffer.
  Called by: `.Clear()` (same file), `.RemoteMessage()` (`Console/cwx.cs`)
- **`.WritePtr()`** — L537 — `public int WritePtr(byte* src, int cnt)`
  Writes from the array pointed to by src into the ringbuffer.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Reset()`** — L571 — `public void Reset()`
  Resets the ringbuffer pointers (will be empty afterwards).
  Called by: `.Restart()` (same file), `.CWXStop()` (`Console/cwx.cs`), `.AbortSending()` (`Console/cwx.cs`)
- **`.Clear()`** — L581 — `public void Clear(int nbytes)`
  Zero the data in the buffer.
  Called by: `.Restart()` (same file)
- **`.Restart()`** — L592 — `public void Restart(int nbytes)`
  Reset the pointers and zero the actual data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Peek()`** — L598 — `public int Peek(byte[] dest, int cnt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadAdvance()`** — L654 — `public void ReadAdvance(int cnt)`
  Reads advance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WriteAdvance()`** — L659 — `public void WriteAdvance(int cnt)`
  Writes advance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/ringbuffer.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
