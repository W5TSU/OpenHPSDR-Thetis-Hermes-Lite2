# `Console/BasicAudio.cs`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** Simple WAV playback (beeps/announcements) outside the DSP path.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×2, references ×1)
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Most-referenced symbols from other files: `.Play()` (×1), `.Stop()` (×1)

## Outline

### Types

#### `BasicAudio` (type, L50)

- `.player_LocationChanged()` — L84
- `.player_LoadCompleted()` — L88
- `.LoadSound()` — L103
- `.Play()` — L127
- `.Stop()` — L143
- `.playSound()` — L153

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/BasicAudio.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
