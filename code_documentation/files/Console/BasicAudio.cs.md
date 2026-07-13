# `Console/BasicAudio.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** Simple WAV playback (beeps/announcements) outside the DSP path.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Play()` (×1), `.Stop()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `BasicAudio` (type, L50)

- **`.player_LocationChanged()`** — L84 — `private void player_LocationChanged(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.player_LoadCompleted()`** — L88 — `private void player_LoadCompleted(object sender, AsyncCompletedEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LoadSound()`** — L103 — `public void LoadSound(string sFile)`
  Loads sound.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Play()`** — L127 — `public void Play()`
  Called by: `.updateQSOTimerStatusbar()` (`Console/console.cs`)
- **`.Stop()`** — L143 — `public void Stop()`
  Called by: `.QSOTimerReset()` (`Console/console.cs`)
- **`.playSound()`** — L153 — `private void playSound()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/BasicAudio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
