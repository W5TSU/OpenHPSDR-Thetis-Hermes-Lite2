# `Console/win32.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Win32 P/Invoke declarations (window messages, power management, multimedia timers).

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Win32` (type, L48)

- **`.memcpy()`** — L52 — `[DllImport("msvcrt.dll", EntryPoint = "memcpy", CallingConvention = CallingConvention.Cdecl)] public static extern void memcpy(void* destptr, void* srcptr, int n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.EnterCriticalSection()`** — L55 — `[DllImport("kernel32.dll", EntryPoint = "EnterCriticalSection")] public static extern void EnterCriticalSection(void* cs_ptr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LeaveCriticalSection()`** — L58 — `[DllImport("kernel32.dll", EntryPoint = "LeaveCriticalSection")] public static extern void LeaveCriticalSection(void* cs_ptr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeCriticalSection()`** — L61 — `[DllImport("kernel32.dll", EntryPoint = "InitializeCriticalSection")] public static extern void InitializeCriticalSection(void* cs_ptr)`
  Initializes critical section.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitializeCriticalSectionAndSpinCount()`** — L64 — `[DllImport("kernel32.dll", EntryPoint = "InitializeCriticalSectionAndSpinCount")] public static extern int InitializeCriticalSectionAndSpinCount(void* cs_ptr, uint spincount)`
  Initializes critical section and spin count.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DeleteCriticalSection()`** — L67 — `[DllImport("kernel32.dll", EntryPoint = "DeleteCriticalSection")] public static extern byte DeleteCriticalSection(void* cs_ptr)`
  Deletes critical section.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetCurrentThread()`** — L70 — `[DllImport("kernel32.dll")] public static extern IntPtr GetCurrentThread()`
  Returns current thread.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetThreadAffinityMask()`** — L73 — `[DllImport("kernel32.dll", SetLastError = true)] public static extern int SetThreadAffinityMask(IntPtr hThread, IntPtr dwThreadAffinityMask)`
  Sets thread affinity mask.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NewCriticalSection()`** — L76 — `[DllImport("wdsp.dll", EntryPoint = "NewCriticalSection", CallingConvention = CallingConvention.Cdecl)] public static extern void* NewCriticalSection()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DestroyCriticalSection()`** — L79 — `[DllImport("wdsp.dll", EntryPoint = "DestroyCriticalSection", CallingConvention = CallingConvention.Cdecl)] public static extern void DestroyCriticalSection(void* cs_ptr)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.memset()`** — L82 — `[DllImport("msvcrt.dll", EntryPoint = "memset", CallingConvention = CallingConvention.Cdecl)] public static extern void memset(void* addr, byte val, int n)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetWindowPos()`** — L85 — `[DllImport("user32.dll", EntryPoint = "SetWindowPos")] public static extern int SetWindowPos(int hwnd, int hWndInsertAfter, int x, int y, int cx, int cy, int wFlags)`
  Sets window pos.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.keybd_event()`** — L88 — `[DllImport("user32.dll", EntryPoint = "keybd_event", CharSet = CharSet.Auto, ExactSpelling = true)] public static extern void keybd_event(byte vk, byte scan, int flags, int extrain`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AllocConsole()`** — L91 — `[DllImport("kernel32.dll", SetLastError = true)] public static extern bool AllocConsole()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FreeConsole()`** — L94 — `[DllImport("kernel32.dll", SetLastError = true)] public static extern bool FreeConsole()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AttachConsole()`** — L97 — `[DllImport("kernel32.dll", SetLastError = true)] public static extern bool AttachConsole(int dwProcessId)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetWindowTextW()`** — L100 — `[DllImport("user32.dll", CharSet = CharSet.Auto, ExactSpelling = true)] public static extern int GetWindowTextW(IntPtr hwnd, System.Text.StringBuilder lpString, int maxcount)`
  Returns window text w.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetWindowTextLengthW()`** — L103 — `[DllImport("user32.dll", CharSet = CharSet.Auto, ExactSpelling = true)] public static extern int GetWindowTextLengthW(IntPtr hwnd)`
  Returns window text length w.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsWindow()`** — L106 — `[DllImport("user32.dll", CharSet = CharSet.Auto, ExactSpelling = true)] public static extern bool IsWindow(IntPtr hWnd)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetForegroundWindow()`** — L109 — `[DllImport("user32.dll", SetLastError = true)] public static extern IntPtr GetForegroundWindow()`
  Returns foreground window.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetWindowThreadProcessId()`** — L112 — `[DllImport("user32.dll", SetLastError = true)] public static extern uint GetWindowThreadProcessId(IntPtr hWnd, out int lpdwProcessId)`
  Returns window thread process id.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddFontMemResourceEx()`** — L115 — `[DllImport("gdi32.dll")] public static extern IntPtr AddFontMemResourceEx(byte[] pbFont, int cbFont, IntPtr pdv, out uint pcFonts)`
  Adds font mem resource ex.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ShowWindowAsync()`** — L118 — `[DllImport("user32.dll")] [return: MarshalAs(UnmanagedType.Bool)] public static extern bool ShowWindowAsync(HandleRef hWnd, int nCmdShow)`
  Shows window async.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetForegroundWindow()`** — L122 — `[DllImport("user32.dll")] [return: MarshalAs(UnmanagedType.Bool)] public static extern bool SetForegroundWindow(IntPtr hWnd)`
  Sets foreground window.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WSAStartup()`** — L126 — `[DllImport("ws2_32.dll", CharSet = CharSet.Auto, SetLastError = true)] public static extern Int32 WSAStartup(Int16 wVersionRequested, out WSAData wsaData)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WSACleanup()`** — L129 — `[DllImport("ws2_32.dll", CharSet = CharSet.Auto, SetLastError = true)] public static extern Int32 WSACleanup()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TimeBeginPeriod()`** — L151 — `[System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Interoperability", "CA1401:PInvokesShouldNotBeVisible"), System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Secu`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TimeEndPeriod()`** — L158 — `[System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Interoperability", "CA1401:PInvokesShouldNotBeVisible"), System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Secu`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `WSAData` (type, L132)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/win32.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
