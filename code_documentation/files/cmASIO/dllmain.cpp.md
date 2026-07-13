# `cmASIO/dllmain.cpp`

**Functional area:** [9. Audio devices, VAC, and ASIO](../../CODE_OUTLINE.md#9-audio-devices-vac-and-asio)

**Role:** The cmASIO DLL: thin host wrapper around the Steinberg ASIO SDK giving ChannelMaster direct ASIO driver access. The bundled `asiosdk_2.3.3.../` tree is the vendored Steinberg SDK.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files):
  - `cmASIO/pch.h` (imports ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`DllMain()`** — L4 — `BOOL APIENTRY DllMain( HMODULE hModule, DWORD ul_reason_for_call, LPVOID lpReserved )`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/cmASIO/dllmain.cpp`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
