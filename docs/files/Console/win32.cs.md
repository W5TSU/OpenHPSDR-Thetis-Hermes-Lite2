# `Console/win32.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Win32 P/Invoke declarations (window messages, power management, multimedia timers).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

### Types

#### `Win32` (type, L48)

- `.memcpy()` — L52
- `.EnterCriticalSection()` — L55
- `.LeaveCriticalSection()` — L58
- `.InitializeCriticalSection()` — L61
- `.InitializeCriticalSectionAndSpinCount()` — L64
- `.DeleteCriticalSection()` — L67
- `.GetCurrentThread()` — L70
- `.SetThreadAffinityMask()` — L73
- `.NewCriticalSection()` — L76
- `.DestroyCriticalSection()` — L79
- `.memset()` — L82
- `.SetWindowPos()` — L85
- `.keybd_event()` — L88
- `.AllocConsole()` — L91
- `.FreeConsole()` — L94
- `.AttachConsole()` — L97
- `.GetWindowTextW()` — L100
- `.GetWindowTextLengthW()` — L103
- `.IsWindow()` — L106
- `.GetForegroundWindow()` — L109
- `.GetWindowThreadProcessId()` — L112
- `.AddFontMemResourceEx()` — L115
- `.ShowWindowAsync()` — L118
- `.SetForegroundWindow()` — L122
- `.WSAStartup()` — L126
- `.WSACleanup()` — L129
- `.TimeBeginPeriod()` — L151
- `.TimeEndPeriod()` — L158

#### `WSAData` (type, L132)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/win32.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
