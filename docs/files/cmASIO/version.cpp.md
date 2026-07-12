# `cmASIO/version.cpp`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** The cmASIO DLL: thin host wrapper around the Steinberg ASIO SDK giving ChannelMaster direct ASIO driver access. The bundled `asiosdk_2.3.3.../` tree is the vendored Steinberg SDK.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `cmASIO/version.h` (imports ×1)

## Outline

### Functions

- `GetCMasioVersion()` — L10

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/cmASIO/version.cpp`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
