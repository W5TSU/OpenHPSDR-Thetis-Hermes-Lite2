# `Console/common.cs`

**Functional area:** [1. Application core and main window](../../CODE_OUTLINE.md#1-application-core-and-main-window)

**Role:** Grab-bag of static helpers: string/number formatting, control lookup, debugging aids, exception reporting.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×49)
  - `Console/console.cs` (calls ×23)
  - `Console/MeterManager.cs` (calls ×16)
  - `Console/Memory/MemoryForm.cs` (calls ×7)
  - `Console/frmAbout.cs` (calls ×6)
  - `Console/clsDBMan.cs` (calls ×5)
  - `Console/frmBandStack2.cs` (calls ×4)
  - `Console/AmpView.cs` (calls ×3)
  - `Console/SpotManager2.cs` (calls ×3)
  - `Console/clsAudioRecordPlayback.cs` (calls ×3)
  - `Console/frmCFCConfig.cs` (calls ×3)
  - `Console/frmFinder.cs` (calls ×3)
  - …and 25 more files
- Uses (outgoing references to other files):
  - `Console/hiperftimer.cs` (references ×1, calls ×1)
  - `Console/database.cs` (calls ×2)
  - `Console/clsDPISafeTools.cs` (calls ×1)
- Most-referenced symbols from other files: `.SaveForm()` (×28), `.OpenUri()` (×15), `.ForceFormOnScreen()` (×9), `.FindNextPowerOf2()` (×9), `.GetComPortNumber()` (×9), `.FindPreviousPowerOf2()` (×8), `.GetVerNum()` (×7), `.RestoreForm()` (×6)

## Outline

### Functions

- `.Add()` — L191
- `.Create()` — L207

### Types

#### `FloatExtensions` (type, L70)

- `.Clamp()` — L72

#### `StringExtensions` (type, L75)

- `.Truncate()` — L77
- `.Contains()` — L87
- `.Left()` — L101
- `.Right()` — L116
- `.ReplaceIgnoreTokenCase()` — L131

#### `ControlExtentions` (type, L155)

- `.GetFullName()` — L157

#### `TypeRenameBinder` (type, L182)

- `.BindToType()` — L197

#### `Common` (type, L214)

- `.HightlightControl()` — L230
- `.DwmGetWindowAttribute()` — L304
- `.DropShadowSize()` — L315
- `.ControlList()` — L339
- `.SaveForm()` — L359
- `.RestoreForm()` — L416
- `.ForceFormOnScreen()` — L510
- `.TabControlInsert()` — L522
- `.SortedComPorts()` — L547
- `.RevToString()` — L567
- `.SetLogPath()` — L576
- `.LogString()` — L580
- `.LogException()` — L602
- `.GetVerNum()` — L640
- `.GetFileVersion()` — L661
- `.GetRevision()` — L669
- `.setupVersions()` — L677
- `.IsAdministrator()` — L692
- `.DoubleBuffered()` — L729
- `.FiveDigitHash()` — L739
- `.ColourToString()` — L758
- `.ColourFromString()` — L762
- `.UVfromDBM()` — L781
- `.SMeterFromDBM()` — L786
- `.SMeterFromDBM_Spaceless()` — L834
- `.GetSMeterUnits()` — L881
- `.SMeterFromDBM2()` — L888
- `.DwmSetWindowAttribute()` — L939
- `.UseImmersiveDarkMode()` — L943
- `.IsWindows10OrGreater()` — L959
- `.DateTimeStringForFile()` — L965
- `.FadeIn()` — L988
- `.FadeOut()` — L1000
- `.CompareVersions()` — L1013
- `.tryParseVersionPart()` — L1035
- `.IsValidUri()` — L1044
- `.OpenUri()` — L1060
- `.FindNextPowerOf2()` — L1074
- `.FindPreviousPowerOf2()` — L1084
- `.IsIpv4Valid()` — L1094
- `.SerializeToBase64()` — L1119
- `.DeserializeFromBase64()` — L1132
- `.HasArg()` — L1156
- `.ArgParam()` — L1166
- `.GetLuminance()` — L1186
- `.rGBtoLin()` — L1196
- `.DoubleBufferAll()` — L1210
- `.IsValidFilename()` — L1219
- `.IsValidPath()` — L1234
- `.DebugPrintCallStack()` — L1249
- `.FourChar()` — L1261
- `.convertToFourChar()` — L1272
- `.CanCreateFile()` — L1288
- `.hasWritePermissionOnDir()` — L1330
- `.isFileWritable()` — L1345
- `.GetComPortNumber()` — L1359
- `.SetProcessPriorityBoost()` — L1373
- `.DisableForegroundPriorityBoost()` — L1375
- `.GetCpuName()` — L1388
- `.GetGpuNames()` — L1407
- `.GetTotalRam()` — L1437
- `.GetInstalledRam()` — L1460
- `.GetDpiForMonitor()` — L1487
- `.MonitorFromWindow()` — L1494
- `.GetScalingForWindow()` — L1499
- `.ProcessCPUUsage()` — L1537
- `.SetThreadExecutionState()` — L1570
- `.PreventSleep()` — L1578
- `.PreventScreenSaver()` — L1589
- `.ResumeSleep()` — L1600
- `.ResumeScreenSaver()` — L1620
- `.GenerateKeyBase64()` — L1658
- `.EncryptAndCombineIvToBase64()` — L1667
- `.DecryptFromCombinedIvBase64()` — L1704
- `.Compress_gzip()` — L1745
- `.Decompress_gzip()` — L1764
- `.GetDiskFreeSpaceExW()` — L1794
- `.TryGetDriveTotalAndFreeBytes()` — L1801
- `.getUncShareRoot()` — L1917

#### `HighlightData` (type, L220)

_No extracted members._

#### `RECT` (type, L306)

_No extracted members._

#### `MonitorDpiType` (type, L1480)

_No extracted members._

#### `ExecutionState` (type, L1562)

_No extracted members._

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/common.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
