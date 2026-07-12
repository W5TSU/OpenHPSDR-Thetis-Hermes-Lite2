# `Console/MeterManager.cs`

**Functional area:** [14. Metering](../../CODE_OUTLINE.md#14-metering)

**Role:** The metering subsystem (~30k lines): collects signal/power/SWR/voltage readings, renders configurable meter faces (analog, bar, LED) via DirectX, and manages multiple meter containers.

## How this file is used

- Used by (incoming references from other files):
  - `Console/setup.cs` (calls ×68, references ×5)
  - `Console/console.cs` (calls ×20, references ×1)
  - `Console/TCIServer.cs` (references ×5, calls ×2)
  - `Console/ucSignalSelect.cs` (references ×4)
  - `Console/CAT/TCPIPcatServer.cs` (calls ×2)
  - `Console/dsp.cs` (references ×2)
  - `Console/frmVariablePicker.cs` (calls ×2)
  - `Console/CAT/CATCommands.cs` (calls ×1)
  - `Console/clsMeterScriptEngine.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×83)
  - `Console/ucOtherButtonsOptionsGrid.cs` (calls ×11, references ×8)
  - `Console/common.cs` (calls ×16)
  - `Console/clsCATMessageQueue.cs` (calls ×7, references ×1)
  - `Console/clsBandStackManager.cs` (calls ×8)
  - `Console/clsMeterScriptEngine.cs` (calls ×8)
  - `Console/ucMeter.Designer.cs` (references ×7)
  - `Console/clsCatAtonic.cs` (references ×4, calls ×3)
  - `Console/Andromeda/Andromeda.cs` (references ×6)
  - `Console/HPSDR/specHPSDR.cs` (calls ×5, references ×1)
  - `Console/clsImgeFetcher.cs` (references ×1, calls ×1)
  - `Console/clsDiscord.cs` (calls ×2)
  - …and 9 more files
- Most-referenced symbols from other files: `.RequiresUpdate()` (×9), `.FilterItemFrequencies()` (×3), `.GetContainerNotes()` (×3), `.QuickestUpdateInterval()` (×3), `.RemoveMMIO()` (×3), `.Start()` (×2), `.Stop()` (×2), `.MeterName()` (×2)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Reading` (type, L78)

_No extracted members._

#### `MeterType` (type, L178)

_No extracted members._

#### `BandGroups` (type, L234)

_No extracted members._

#### `Globals` (type, L241)

_No extracted members._

#### `MeterManager` (type, L245)

- **`.SetLedVariable()`** — L404 — `internal static void SetLedVariable(int rx, string variable, object value)`
  Sets led variable.
  Called by: `.add_readings()` (same file)
- **`.provide_variables()`** — L410 — `private static MeterScriptEngine.Snapshot provide_variables()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CatVariables()`** — L441 — `public static List<string> CatVariables()`
  Called by: `.Init()` (`Console/frmVariablePicker.cs`)
- **`.OnCatQmessage()`** — L454 — `private static void OnCatQmessage(int queue_index, Guid guid, ScriptCommand sc)`
  cat queue
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.guard_holds()`** — L485 — `private static bool guard_holds(ScriptCommand c)`
  Called by: `.OnCatQmessage()` (same file)
- **`.eval_now()`** — L509 — `private static bool eval_now(string id, int macro, int button_index, string name)`
  Called by: `.guard_holds()` (same file)
- **`.normalise_var()`** — L552 — `private static string normalise_var(string name)`
  Called by: `.OnCatQmessage()` (same file)
- **`.OnCatQstate()`** — L559 — `private static void OnCatQstate(int queue_index, bool running)`
  Handles/raises the cat qstate event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.rebuildAllLedReadings()`** — L565 — `private static void rebuildAllLedReadings()`
  Called by: `.OnCatQmessage()` (same file)
- **`.GetLedFrom4Char()`** — L576 — `public static clsLed GetLedFrom4Char(string fourchar)`
  Returns led from4 char.
  Called by: `.Update()` (same file)
- **`.GetVoiceRecordPlayFrom4Char()`** — L592 — `public static clsVoiceRecordPlay GetVoiceRecordPlayFrom4Char(string fourchar)`
  Returns voice record play from4 char.
  Called by: `.ZZJQ()` (`Console/CAT/CATCommands.cs`)
- **`.GetOtherButtonsFromID()`** — L608 — `public static clsOtherButtons GetOtherButtonsFromID(string id)`
  Returns other buttons from id.
  Called by: `.eval_now()` (same file)
- **`.GetWebImageIDsFrom4Char()`** — L629 — `public static (string, string) GetWebImageIDsFrom4Char(string fourchar)`
  Returns web image ids from4 char.
  Called by: `.btnWebImage_goto_next_Click()` (`Console/setup.cs`)
- **`.IsWebImageBackgroundShown()`** — L647 — `public static bool IsWebImageBackgroundShown()`
  Called by: `.Update()` (same file)
- **`.FilterItemFrequencies()`** — L662 — `public static void FilterItemFrequencies(FilterItemSnapFrequencies setting_group, string settings)`
  Called by: `.txtFilter_sideband_frequencies_TextChanged()` (`Console/setup.cs`), `.txtFilter_cw_frequencies_TextChanged()` (`Console/setup.cs`), `.txtFilter_other_frequencies_TextChanged()` (`Console/setup.cs`)
- **`.GetFilterItemFrequencies()`** — L715 — `public static float[] GetFilterItemFrequencies(FilterItemSnapFrequencies setting_group)`
  Returns filter item frequencies.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadingsCustom()`** — L2010 — `public static CustomReadings ReadingsCustom(int rx)`
  Called by: `.Update()` (same file), `.parseText()` (same file), `.onTimerElapsedCondition()` (same file), `.add_readings()` (same file), `.expand_placeholders()` (same file), `.ZeroOut()` (same file) — and 1 more
- **`.ZeroReading()`** — L2015 — `public static void ZeroReading(out float value, int rx, Reading reading)`
  zero reading
  Called by: `.ZeroOut()` (same file)
- **`.GetMeterTXRXType()`** — L2144 — `public static int GetMeterTXRXType(MeterType meter)`
  Returns meter txrxtype.
  Called by: `.lstMetersInUse_DrawItem()` (`Console/setup.cs`)
- **`.MeterName()`** — L2206 — `public static string MeterName(MeterType meter)`
  Called by: `.ToString()` (`Console/setup.cs`), `.findIndexForInsertOfSpecialItem()` (`Console/setup.cs`)
- **`.ReadingName()`** — L2261 — `public static string ReadingName(Reading reading)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ReadingUnits()`** — L2322 — `public static string ReadingUnits(Reading reading)`
  Called by: `.updateReadingText()` (same file), `.renderHBar()` (same file), `.renderSignalTextDisplay()` (same file)
- **`.UpdateMeters()`** — L2376 — `private static void UpdateMeters()`
  Updates meters.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.normaliseWaveRecordPath()`** — L2423 — `private static string normaliseWaveRecordPath(string path)`
  Called by: `.waveRecordPathsEqual()` (same file), `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file), `.sanitiseStoredPaths()` (same file)
- **`.waveRecordPathsEqual()`** — L2437 — `private static bool waveRecordPathsEqual(string left, string right)`
  Called by: `.AddFiles()` (same file), `.IsPlaying()` (same file), `.MouseUp()` (same file), `.containsFilePathLocked()` (same file), `.pathArraysEqual()` (same file), `.handlePlay()` (same file) — and 1 more
- **`.loadImages()`** — L2445 — `private static void loadImages(bool init = false)`
  images, used by dxrenderer
  Called by: `.RefreshAllImages()` (same file), `.loadDXSkinImages()` (same file), `.RunRendererDisplay()` (same file)
- **`.loadImage()`** — L2495 — `private static void loadImage(string sFilePath, bool isSkinImage)`
  Called by: `.loadImages()` (same file)
- **`.loadResouceImages()`** — L2519 — `private static void loadResouceImages()`
  Called by: `.RunRendererDisplay()` (same file)
- **`.clearAllCachedImageData()`** — L2574 — `private static void clearAllCachedImageData(bool bOnlySkins = false)`
  Called by: `.RefreshAllImages()` (same file), `.loadDXSkinImages()` (same file)
- **`.removeImageCacheData()`** — L2594 — `private static bool removeImageCacheData(string sKey, bool bOnlySkins = false)`
  Called by: `.loadImages()` (same file), `.clearAllCachedImageData()` (same file)
- **`.addBitmap()`** — L2619 — `private static bool addBitmap(string sKey, System.Drawing.Bitmap image)`
  Called by: `.loadImage()` (same file), `.loadResouceImages()` (same file)
- **`.GetBitmap()`** — L2631 — `internal static System.Drawing.Bitmap GetBitmap(string sKey)`
  Returns bitmap.
  Called by: `.convertImageToDX()` (same file)
- **`.ContainsBitmap()`** — L2641 — `internal static bool ContainsBitmap(string sKey)`
  Called by: `.loadImage()` (same file), `.loadResouceImages()` (same file), `.renderImage()` (same file)
- **`.AddStreamData()`** — L2649 — `internal static void AddStreamData(string sId, MemoryStream tempStream)`
  Adds stream data.
  Called by: `.bitmapFromSystemBitmap()` (same file)
- **`.GetStreamData()`** — L2659 — `internal static MemoryStream GetStreamData(string sKey)`
  Returns stream data.
  Called by: `.bitmapFromSystemBitmap()` (same file)
- **`.ContainsStreamData()`** — L2669 — `internal static bool ContainsStreamData(string sKey)`
  Called by: `.bitmapFromSystemBitmap()` (same file)
- **`.RemoveStreamData()`** — L2677 — `internal static void RemoveStreamData(string sKey)`
  Removes stream data.
  Called by: `.Removing()` (same file), `.renderWebImage()` (same file)
- **`.ContainerBorder()`** — L2695 — `public static void ContainerBorder(string sId, bool border)`
  Called by: `.chkContainerBorder_CheckedChanged()` (`Console/setup.cs`)
- **`.NoTitle()`** — L2705 — `public static void NoTitle(string sId, bool noTitle)`
  Called by: `.chkContainerNoTitle_CheckedChanged()` (`Console/setup.cs`)
- **`.AutoContainerHeight()`** — L2716 — `public static void AutoContainerHeight(string sId, bool auto_height)`
  Called by: `.chkMultiMeter_auto_container_height_CheckedChanged()` (`Console/setup.cs`)
- **`.SetContainerHeight()`** — L2727 — `public static void SetContainerHeight(string sId, int height)`
  Sets container height.
  Called by: `.dxRender()` (same file)
- **`.ContainerMinimises()`** — L2745 — `public static void ContainerMinimises(string sId, bool minimises)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`), `.chkContainerMinimises_CheckedChanged()` (`Console/setup.cs`)
- **`.RefreshContainerVisible()`** — L2766 — `public static void RefreshContainerVisible(string id)`
  Refreshes container visible.
  Called by: `.ucOtherButtonsOptionsGrid_buttons_MacroSetupClicked()` (`Console/setup.cs`)
- **`.containerVisible()`** — L2771 — `private static void containerVisible(string id, bool visible)`
  Called by: `.RefreshContainerVisible()` (same file), `.ContainerHidesWhenRXNotUsed()` (same file), `.enableContainer()` (same file), `.returnMeterFromFloating()` (same file), `.setMeterFloating()` (same file), `.OnRX2EnabledPreChanged()` (same file) — and 1 more
- **`.ContainerHidesWhenRXNotUsed()`** — L2784 — `public static void ContainerHidesWhenRXNotUsed(string sId, bool hides)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`), `.chkContainer_hidewhennotused_CheckedChanged()` (`Console/setup.cs`)
- **`.LockContainer()`** — L2827 — `public static void LockContainer(string sId, bool locked)`
  Called by: `.chkLockContainer_CheckedChanged()` (`Console/setup.cs`)
- **`.SetContainerRX()`** — L2837 — `public static void SetContainerRX(string sId, int rx)`
  Sets container rx.
  Called by: `.radContainer_rx1_data_CheckedChanged()` (`Console/setup.cs`), `.radContainer_rx2_data_CheckedChanged()` (`Console/setup.cs`)
- **`.GetContainerRX()`** — L2870 — `public static int GetContainerRX(string sId)`
  Returns container rx.
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.updateHiddenMacroInfo()`** — L2882 — `private static void updateHiddenMacroInfo(ucMeter ucM)`
  Called by: `.HiddenByMacro()` (same file), `.enableContainer()` (same file), `.AddMeterContainer()` (same file)
- **`.HiddenByMacro()`** — L2894 — `public static void HiddenByMacro(string sId, bool hide, string sId_copy_size = null)`
  Called by: `.handleMacroButtonPress()` (same file)
- **`.ShowContainerOnRX()`** — L2946 — `public static void ShowContainerOnRX(string sId, bool visible)`
  Shows container on rx.
  Called by: `.chkContainerShowRX_CheckedChanged()` (`Console/setup.cs`)
- **`.ShowContainerOnTX()`** — L2965 — `public static void ShowContainerOnTX(string sId, bool visible)`
  Shows container on tx.
  Called by: `.chkContainerShowTX_CheckedChanged()` (`Console/setup.cs`)
- **`.enableContainer()`** — L2984 — `public static void enableContainer(string sId, bool enabled, bool undo_hidden_by_macro = false)`
  Called by: `.HiddenByMacro()` (same file), `.ShowContainerOnRX()` (same file), `.ShowContainerOnTX()` (same file), `.OnMox()` (same file), `.RecoverContainer()` (same file)
- **`.ContainerHasBorder()`** — L3071 — `public static bool ContainerHasBorder(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerNoTitleBar()`** — L3082 — `public static bool ContainerNoTitleBar(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerAutoHeight()`** — L3093 — `public static bool ContainerAutoHeight(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerLocked()`** — L3104 — `public static bool ContainerLocked(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerIsHidden()`** — L3126 — `public static bool ContainerIsHidden(string sId)`
  Called by: `.RefreshContainerVisible()` (same file), `.setupButtons()` (same file), `.handleMacroButtonPress()` (same file)
- **`.ContainerShowOnRX()`** — L3162 — `public static bool ContainerShowOnRX(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerShowOnTX()`** — L3173 — `public static bool ContainerShowOnTX(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.ContainerNotes()`** — L3195 — `public static void ContainerNotes(string sId, string notes)`
  Called by: `.txtContainerNotes_TextChanged()` (`Console/setup.cs`)
- **`.GetContainerNotes()`** — L3207 — `public static string GetContainerNotes(string sId)`
  Returns container notes.
  Called by: `.containerNameFromId()` (`Console/setup.cs`), `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`), `.txtContainerNotes_TextChanged()` (`Console/setup.cs`)
- **`.ContainerBackgroundColour()`** — L3218 — `public static void ContainerBackgroundColour(string sId, System.Drawing.Color c)`
  Called by: `.clrbtnContainerBackground_Changed()` (`Console/setup.cs`)
- **`.GetContainerBackgroundColour()`** — L3232 — `public static System.Drawing.Color GetContainerBackgroundColour(string sId)`
  Returns container background colour.
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.HighlightContainer()`** — L3243 — `public static void HighlightContainer(string sId)`
  Called by: `.comboContainerSelect_SelectedIndexChanged()` (`Console/setup.cs`), `.chkContainerHighlight_CheckedChanged()` (`Console/setup.cs`)
- **`.DisposeImageData()`** — L3257 — `public static void DisposeImageData()`
  Called by: `.Shutdown()` (same file)
- **`.SetAntennaAuxText()`** — L3274 — `public static void SetAntennaAuxText(string n1, string n2, string n3)`
  Sets antenna aux text.
  Called by: `.comboRadioModel_SelectedIndexChanged()` (`Console/setup.cs`)
- **`.initAntennaArrays()`** — L3284 — `private static void initAntennaArrays()`
  Called by: `.Init()` (same file)
- **`.Init()`** — L3336 — `public static void Init(Console c, Display.AdaptorInfo adaptor = null)`
  Called by: `.InitConsole()` (`Console/console.cs`)
- **`.addRenderer()`** — L3362 — `private static void addRenderer(string sId, int rx, Panel target, clsMeter meter, System.Drawing.Color backColour)`
  Called by: `.AddMeterContainer()` (same file)
- **`.UpdateS9()`** — L3369 — `public static void UpdateS9(double s9freq)`
  Updates s9.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RefreshAllImages()`** — L3378 — `public static void RefreshAllImages()`
  Refreshes all images.
  Called by: `.fileDownloadHandler()` (`Console/setup.cs`)
- **`.loadDXSkinImages()`** — L3393 — `private static void loadDXSkinImages()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunRendererDisplay()`** — L3408 — `public static void RunRendererDisplay(string sId, bool init = false)`
  Called by: `.RunAllRendererDisplays()` (same file), `.AddMeterContainer()` (same file), `.btnContainer_load_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.RunAllRendererDisplays()`** — L3419 — `public static void RunAllRendererDisplays()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVsync()`** — L3428 — `public static void SetVsync(bool vsync)`
  Sets vsync.
  Called by: `.chkVSyncDX_CheckedChanged()` (`Console/setup.cs`)
- **`.removeRenderer()`** — L3453 — `private static void removeRenderer(string sId)`
  Called by: `.RemoveMeterContainer()` (same file)
- **`.GetMeterIDsFromSaveData()`** — L3462 — `public static List<string> GetMeterIDsFromSaveData(Dictionary<string, string> data, bool include_ig_ids = false)`
  Returns meter ids from save data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Shutdown()`** — L3524 — `public static void Shutdown()`
  Called by: `.Console_Closing()` (`Console/console.cs`)
- **`.addDelegates()`** — L3568 — `private static void addDelegates()`
  Called by: `.Init()` (same file)
- **`.removeDelegates()`** — L3679 — `private static void removeDelegates()`
  Called by: `.Shutdown()` (same file)
- **`.OnContainerHiddenByMacro()`** — L3795 — `private static void OnContainerHiddenByMacro(string id, bool hidden)`
  Handles/raises the container hidden by macro event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnContainerEnabled()`** — L3807 — `private static void OnContainerEnabled(string id, bool enabled)`
  Handles/raises the container enabled event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCWXShown()`** — L3819 — `public static void OnCWXShown(bool shown)`
  Handles/raises the cwxshown event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCWPitchChanged()`** — L3831 — `public static void OnCWPitchChanged(int old_pitch, int new_pitch, bool show_cwzero)`
  Handles/raises the cwpitch changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTNFChanged()`** — L3844 — `public static void OnTNFChanged(bool old_tnf, bool new_tnf)`
  Handles/raises the tnfchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRXSpecGridMinMaxChanged()`** — L3856 — `private static void OnRXSpecGridMinMaxChanged(int rx, int min, int max)`
  Handles/raises the rxspec grid min max changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXSpecGridMinMaxChanged()`** — L3869 — `private static void OnTXSpecGridMinMaxChanged(int min, int max)`
  Handles/raises the txspec grid min max changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnWaterfallRXGradientChanged()`** — L3883 — `private static void OnWaterfallRXGradientChanged(int rx, System.Drawing.Color[] colours)`
  Handles/raises the waterfall rxgradient changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnWaterfallTXGradientChanged()`** — L3895 — `private static void OnWaterfallTXGradientChanged(System.Drawing.Color[] colours)`
  Handles/raises the waterfall txgradient changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRXWaterfallMinMaxChanged()`** — L3907 — `private static void OnRXWaterfallMinMaxChanged(int rx, int min, int max)`
  Handles/raises the rxwaterfall min max changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXWaterfallMinMaxChanged()`** — L3920 — `private static void OnTXWaterfallMinMaxChanged(int min, int max)`
  Handles/raises the txwaterfall min max changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPAProfileChanged()`** — L3934 — `private static void OnPAProfileChanged(string old_profile, string new_profile)`
  Handles/raises the paprofile changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXProfileChanged()`** — L3948 — `private static void OnTXProfileChanged(string old_profile, string new_profile)`
  Handles/raises the txprofile changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXFiltersChanged()`** — L3962 — `private static void OnTXFiltersChanged(int low, int high)`
  Handles/raises the txfilters changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTuneStepIndexChanged()`** — L3977 — `private static void OnTuneStepIndexChanged(int rx, int old_index, int new_index)`
  Handles/raises the tune step index changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOSyncChanged()`** — L3990 — `private static void OnVFOSyncChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfosync changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOaLockChanged()`** — L4002 — `private static void OnVFOaLockChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfoa lock changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFObLockChanged()`** — L4014 — `private static void OnVFObLockChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the vfob lock changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateAntennaMeters()`** — L4027 — `private static void updateAntennaMeters()`
  Called by: `.SetAntennaAuxText()` (same file), `.OnAntennaRXChanged()` (same file), `.OnAntennaTXChanged()` (same file), `.OnAntennaAuxChanged()` (same file), `.OnAntennaDoNotTX()` (same file), `.OnAntennaRxTx()` (same file)
- **`.OnTXFrequencyChanged()`** — L4038 — `private static void OnTXFrequencyChanged(double old_frequency, double new_frequency, Band old_band, Band new_band, bool rx2_enabled, bool tx_vfob, double centre_freq)`
  Handles/raises the txfrequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAntennaRXChanged()`** — L4066 — `private static void OnAntennaRXChanged(Band b, int antenna, bool old_state, bool new_state)`
  Handles/raises the antenna rxchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAntennaTXChanged()`** — L4078 — `private static void OnAntennaTXChanged(Band b, int antenna, bool old_state, bool new_state)`
  Handles/raises the antenna txchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAntennaAuxChanged()`** — L4090 — `private static void OnAntennaAuxChanged(Band b, int antenna, bool old_state, bool new_state, string button_text)`
  Handles/raises the antenna aux changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAntennaDoNotTX()`** — L4102 — `private static void OnAntennaDoNotTX(int antenna, bool old_state, bool new_state)`
  Handles/raises the antenna do not tx event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAntennaRxTx()`** — L4113 — `private static void OnAntennaRxTx(bool old_state, bool new_state)`
  Handles/raises the antenna rx tx event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNoTX()`** — L4130 — `public static bool GetNoTX(int antenna)`
  Returns no tx.
  Called by: `.setupButtons()` (same file)
- **`.GetRXAntState()`** — L4138 — `public static bool GetRXAntState(Band b, int antenna)`
  Returns rxant state.
  Called by: `.setupButtons()` (same file)
- **`.GetTXAntState()`** — L4147 — `public static bool GetTXAntState(Band b, int antenna)`
  Returns txant state.
  Called by: `.setupButtons()` (same file)
- **`.GetRXAuxState()`** — L4156 — `public static bool GetRXAuxState(Band b, int antenna)`
  Returns rxaux state.
  Called by: `.setupButtons()` (same file)
- **`.GetRXAuxName()`** — L4165 — `public static string GetRXAuxName(int antenna)`
  Returns rxaux name.
  Called by: `.setupButtons()` (same file)
- **`.GetTXEnabled()`** — L4173 — `public static bool GetTXEnabled(int antenna)`
  Returns txenabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetRXAnt()`** — L4181 — `public static int GetRXAnt(Band b, bool xvtr)`
  Returns rxant.
  Called by: `.setupButtons()` (same file)
- **`.GetTXAnt()`** — L4210 — `public static int GetTXAnt(Band b)`
  Returns txant.
  Called by: `.setupButtons()` (same file)
- **`.OnSplitChanged()`** — L4220 — `private static void OnSplitChanged(int rx, bool oldSplit, bool newSplit)`
  Handles/raises the split changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMinimumNotchWidthChangedRX()`** — L4236 — `private static void OnMinimumNotchWidthChangedRX(int rx, double width)`
  Handles/raises the minimum notch width changed rx event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMinimumNotchWidthChangedTX()`** — L4245 — `private static void OnMinimumNotchWidthChangedTX(double width)`
  Handles/raises the minimum notch width changed tx event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterEdgesChanged()`** — L4254 — `private static void OnFilterEdgesChanged(int rx, Filter newFilter, Band band, int low, int high, string sName, int max_width, int max_shift)`
  Handles/raises the filter edges changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterChanged()`** — L4258 — `private static void OnFilterChanged(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high, string sName)`
  Handles/raises the filter changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateFilterInfo()`** — L4262 — `private static void updateFilterInfo(int rx, Filter oldFilter, Filter newFilter, Band band, int low, int high, string sName, int max_width = -1, int max_shift = -1)`
  Called by: `.OnFilterEdgesChanged()` (same file), `.OnFilterChanged()` (same file)
- **`.OnMultiRxChanged()`** — L4311 — `private static void OnMultiRxChanged(int rx, bool newState, bool oldState, double vfoASubFrequency, Band b, bool bRx2Enabled)`
  Handles/raises the multi rx changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTXBandChanged()`** — L4326 — `private static void OnTXBandChanged(Band oldBand, Band newBand, double tx_frequency)`
  Handles/raises the txband changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOTXChanged()`** — L4338 — `private static void OnVFOTXChanged(bool vfoB, bool oldState, bool newState)`
  Handles/raises the vfotxchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChangeHandler()`** — L4352 — `private static void OnModeChangeHandler(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVHFDetailsChanged()`** — L4377 — `private static void OnVHFDetailsChanged(int idx, bool old_state, bool new_state, string old_text, string new_text)`
  Handles/raises the vhfdetails changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBandPanelChanged()`** — L4391 — `private static void OnBandPanelChanged(int rx, bool gen, bool hf, bool vhf)`
  Handles/raises the band panel changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.bandChange()`** — L4405 — `private static void bandChange(int rx, Band old_band, Band new_band, bool update_button_boxes = false, bool update_band = true)`
  Called by: `.OnBandChange()` (same file), `.OnVFOA()` (same file), `.OnVFOB()` (same file)
- **`.OnBandChange()`** — L4476 — `private static void OnBandChange(int rx, Band oldBand, Band newBand)`
  Handles/raises the band change event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPreBandChange()`** — L4480 — `private static void OnPreBandChange(int rx, Band currentBand)`
  Handles/raises the pre band change event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTransverterIndexChanged()`** — L4484 — `private static void OnTransverterIndexChanged(int oldIndex, int newIndex)`
  Handles/raises the transverter index changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAlexPresentChanged()`** — L4488 — `private static void OnAlexPresentChanged(bool oldSetting, bool newSetting)`
  Handles/raises the alex present changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPAPresentChanged()`** — L4492 — `private static void OnPAPresentChanged(bool oldSetting, bool newSetting)`
  Handles/raises the papresent changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnApolloPresentChanged()`** — L4496 — `private static void OnApolloPresentChanged(bool oldSetting, bool newSetting)`
  Handles/raises the apollo present changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCurrentModelChanged()`** — L4500 — `private static void OnCurrentModelChanged(HPSDRModel oldModel, HPSDRModel newModel)`
  Handles/raises the current model changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOA()`** — L4504 — `private static void OnVFOA(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double oldCentreF, doub`
  Handles/raises the vfoa event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOB()`** — L4545 — `private static void OnVFOB(Band oldBand, Band newBand, DSPMode oldMode, DSPMode newMode, Filter oldFilter, Filter newFilter, double oldFreq, double newFreq, double oldCentreF, doub`
  Handles/raises the vfob event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVFOASub()`** — L4578 — `private static void OnVFOASub(Band oldBand, Band newBand, DSPMode newMode, Filter newFilter, double oldFreq, double newFreq, double newCentreF, bool newCTUN, int newZoomSlider, dou`
  Handles/raises the vfoasub event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDuplexChanged()`** — L4597 — `private static void OnDuplexChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the duplex changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPSAChanged()`** — L4614 — `private static void OnPSAChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the psachanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnQuickPlayChanged()`** — L4627 — `private static void OnQuickPlayChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the quick play changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnANFChanged()`** — L4640 — `private static void OnANFChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the anfchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSNBChanged()`** — L4652 — `private static void OnSNBChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the snbchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAVGChanged()`** — L4664 — `private static void OnAVGChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the avgchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPeakChanged()`** — L4676 — `private static void OnPeakChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the peak changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L4688 — `private static void OnCTUNChanged(int rx, bool old_state, bool new_state, Band b)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnVACEnabledChanged()`** — L4700 — `private static void OnVACEnabledChanged(int vac, bool old_state, bool new_state)`
  Handles/raises the vacenabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMuteChanged()`** — L4720 — `private static void OnMuteChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the mute changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBINChanged()`** — L4732 — `private static void OnBINChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the binchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPanSwapChanged()`** — L4744 — `private static void OnPanSwapChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the pan swap changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDisplayModeChanged()`** — L4757 — `private static void OnDisplayModeChanged(int rx, DisplayMode old_mode, DisplayMode new_mode)`
  Handles/raises the display mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAGCModeChanged()`** — L4769 — `private static void OnAGCModeChanged(int rx, AGCMode old_mode, AGCMode new_mode)`
  Handles/raises the agcmode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAGCAutoModeChanged()`** — L4781 — `private static void OnAGCAutoModeChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the agcauto mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSqlChanged()`** — L4793 — `private static void OnSqlChanged(int rx, SquelchState old_state, SquelchState new_state)`
  Handles/raises the sql changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnGeneralSettingsChanged()`** — L4812 — `private static void OnGeneralSettingsChanged(int rx, OtherButtonId setting, bool old_state, bool new_state, Dictionary<OtherButtonId, bool> settings)`
  Handles/raises the general settings changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnXPAChanged()`** — L4825 — `private static void OnXPAChanged(bool in_use, bool old_state, bool new_state)`
  Handles/raises the xpachanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnQuickRecordChanged()`** — L4839 — `private static void OnQuickRecordChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the quick record changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnWaveRecordChanged()`** — L4852 — `private static void OnWaveRecordChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the wave record changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTwoToneChanged()`** — L4865 — `private static void OnTwoToneChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the two tone changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTuneChanged()`** — L4877 — `private static void OnTuneChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the tune changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnNRChanged()`** — L4889 — `private static void OnNRChanged(int rx, int old_nr, int new_nr)`
  Handles/raises the nrchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnNBChanged()`** — L4901 — `private static void OnNBChanged(int rx, int old_nb, int new_nb)`
  Handles/raises the nbchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMONChanged()`** — L4913 — `private static void OnMONChanged(bool old_state, bool new_state)`
  Handles/raises the monchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.initAllConsoleData()`** — L4926 — `private static void initAllConsoleData()`
  Called by: `.FinishSetupAndDisplay()` (same file), `.OnRX2EnabledChanged()` (same file)
- **`.initConsoleData()`** — L4946 — `private static void initConsoleData(clsMeter m)`
  Called by: `.SetContainerRX()` (same file), `.initAllConsoleData()` (same file), `.AddMeterContainer()` (same file), `.FinishSetupAndDisplay()` (same file), `.AddMeter()` (same file)
- **`.getFilterName()`** — L5167 — `private static string getFilterName(int rx)`
  Returns filter name.
  Called by: `.initConsoleData()` (same file), `.setupButtons()` (same file)
- **`.OnPower()`** — L5223 — `private static void OnPower(bool oldPower, bool newPower)`
  Handles/raises the power event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPreMox()`** — L5242 — `private static void OnPreMox(int rx, bool oldMox, bool newMox)`
  Handles/raises the pre mox event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMox()`** — L5252 — `private static void OnMox(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnEQChanged()`** — L5314 — `private static void OnEQChanged(bool oldState, bool newState)`
  Handles/raises the eqchanged event.
  Called by: `.AddMeter()` (same file)
- **`.OnLevelerChanged()`** — L5327 — `private static void OnLevelerChanged(bool oldState, bool newState)`
  Handles/raises the leveler changed event.
  Called by: `.AddMeter()` (same file)
- **`.OnCFCChanged()`** — L5341 — `private static void OnCFCChanged(bool oldState, bool newState)`
  Handles/raises the cfcchanged event.
  Called by: `.AddMeter()` (same file)
- **`.OnCompandChanged()`** — L5355 — `private static void OnCompandChanged(bool oldState, bool newState)`
  Handles/raises the compand changed event.
  Called by: `.AddMeter()` (same file)
- **`.OnQuickSplitChanged()`** — L5368 — `private static void OnQuickSplitChanged(bool oldState, bool newState)`
  Handles/raises the quick split changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dbmOffsetForAboveS9Frequency()`** — L5419 — `private static float dbmOffsetForAboveS9Frequency(int rx)`
  Called by: `.getPerc()` (same file)
- **`.IsAboveS9Frequency()`** — L5426 — `private static bool IsAboveS9Frequency(int rx)`
  Called by: `.ZeroReading()` (same file), `.dbmOffsetForAboveS9Frequency()` (same file), `.addSMeterBar()` (same file), `.AddMagicEye()` (same file), `.AddAnanMM()` (same file), `.renderHBar()` (same file) — and 1 more
- **`.getReading()`** — L5447 — `private static float getReading(int rx, Reading rt, bool bUseReading = false)`
  Returns reading.
  Called by: `.TakeReading()` (same file), `.Update()` (same file)
- **`.setReading()`** — L5456 — `private static void setReading(int rx, Reading rt, ref Dictionary<Reading, float> readings, bool bChangeOverride = false)`
  Sets reading.
  Called by: `.OnMeterReadings()` (same file)
- **`.setReadingForced()`** — L5461 — `private static void setReadingForced(int rx, Reading rt, float reading)`
  Sets reading forced.
  Called by: `.ZeroReading()` (same file), `.ZeroOut()` (same file)
- **`.OnMeterReadings()`** — L5468 — `private static void OnMeterReadings(int rx, bool mox, ref Dictionary<Reading, float> readings)`
  Handles/raises the meter readings event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RequiresUpdate()`** — L5543 — `public static bool RequiresUpdate(int rx, Reading rt)`
  Called by: `.computeRefPower()` (`Console/console.cs`), `.computeAlexFwdPower()` (`Console/console.cs`), `.computeExciterPower()` (`Console/console.cs`), `.computeOrionMkIIExciterPower()` (`Console/console.cs`), `.computeOrionExciterPower()` (`Console/console.cs`), `.computeANANExciterPower()` (`Console/console.cs`) — and 3 more
- **`.QuickestUpdateInterval()`** — L5552 — `public static int QuickestUpdateInterval(int rx, bool mox)`
  meters
  Called by: `.chkPower_CheckedChanged()` (`Console/console.cs`), `.MultiMeter2UpdateRX1()` (`Console/console.cs`), `.MultiMeter2UpdateRX2()` (`Console/console.cs`)
- **`.KeycodeInUse()`** — L5569 — `public static bool KeycodeInUse(Keys keycode)`
  Called by: `.onGlobalKeyDown()` (`Console/setup.cs`)
- **`.AbortAllVoiceRecordRepeatPlaybacks()`** — L5587 — `public static void AbortAllVoiceRecordRepeatPlaybacks()`
  Called by: `.onGlobalKeyDown()` (`Console/setup.cs`)
- **`.MeterFromId()`** — L5598 — `public static clsMeter MeterFromId(string sId)`
  Called by: `.RestoreSettings()` (same file), `.ContainerToString()` (same file), `.ContainerFromString()` (same file), `.StoreSettings2()` (same file), `.meterFromSelectedContainer()` (`Console/setup.cs`)
- **`.MeterExists()`** — L5608 — `public static bool MeterExists(string sId)`
  Called by: `.RestoreSettings()` (same file), `.ContainerFromString()` (same file)
- **`.onGlobalKeyDown()`** — L5616 — `private static void onGlobalKeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onGlobalKeyUp()`** — L5627 — `private static void onGlobalKeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddMeterContainer()`** — L5638 — `public static void AddMeterContainer(ucMeter ucM, bool bFromRestore = false)`
  Adds meter container.
  Called by: `.RestoreSettings()` (same file), `.ContainerFromString()` (same file), `.btnAddRX1Container_Click()` (`Console/setup.cs`), `.btnAddRX2Container_Click()` (`Console/setup.cs`)
- **`.IsOnTop()`** — L5699 — `public static bool IsOnTop(string sId)`
  Called by: `.MouseUp()` (same file), `.PopBandStack()` (same file)
- **`.FinishSetupAndDisplay()`** — L5712 — `public static void FinishSetupAndDisplay()`
  Called by: `.btnContainer_load_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.BringToFront()`** — L5774 — `public static void BringToFront()`
  Called by: `.updateAttNudsCombos()` (`Console/console.cs`)
- **`.ucMeter_SettingsClicked()`** — L5809 — `private static void ucMeter_SettingsClicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucMeter_FloatingDockedClicked()`** — L5817 — `private static void ucMeter_FloatingDockedClicked(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ucMeter_FloatingDockedMoved()`** — L5831 — `private static void ucMeter_FloatingDockedMoved(object sender, EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPositionOfDockedMeters()`** — L5837 — `public static void SetPositionOfDockedMeters()`
  Sets position of docked meters.
  Called by: `.ResizeConsole()` (`Console/console.cs`)
- **`.setPoisitionOfDockedMeter()`** — L5847 — `private static void setPoisitionOfDockedMeter(ucMeter m)`
  Sets poisition of docked meter.
  Called by: `.SetPositionOfDockedMeters()` (same file), `.returnMeterFromFloating()` (same file)
- **`.returnMeterFromFloating()`** — L5892 — `private static void returnMeterFromFloating(ucMeter m, frmMeterDisplay frm)`
  Called by: `.ContainerHidesWhenRXNotUsed()` (same file), `.enableContainer()` (same file), `.AddMeterContainer()` (same file), `.FinishSetupAndDisplay()` (same file), `.ucMeter_FloatingDockedClicked()` (same file), `.OnRX2EnabledChanged()` (same file)
- **`.setMeterFloating()`** — L5919 — `private static void setMeterFloating(ucMeter m, frmMeterDisplay frm)`
  Sets meter floating.
  Called by: `.ContainerHidesWhenRXNotUsed()` (same file), `.enableContainer()` (same file), `.AddMeterContainer()` (same file), `.FinishSetupAndDisplay()` (same file), `.ucMeter_FloatingDockedClicked()` (same file), `.OnRX2EnabledChanged()` (same file)
- **`.OnRX2EnabledChanged()`** — L5944 — `private static void OnRX2EnabledChanged(bool enabled)`
  Handles/raises the rx2 enabled changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnRX2EnabledPreChanged()`** — L5983 — `private static void OnRX2EnabledPreChanged(bool enabled)`
  Handles/raises the rx2 enabled pre changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.zeroAllMeters()`** — L6026 — `private static void zeroAllMeters()`
  Called by: `.UpdateS9()` (same file), `.FinishSetupAndDisplay()` (same file)
- **`.RestoreSettings()`** — L6037 — `public static bool RestoreSettings(ref Dictionary<string, string> settings)`
  Restores settings.
  Called by: `.getOptions()` (`Console/setup.cs`)
- **`.GetFormGuidList()`** — L6131 — `public static List<string> GetFormGuidList()`
  Returns form guid list.
  Called by: `.SaveOptions()` (`Console/setup.cs`)
- **`.ContainerToString()`** — L6147 — `public static string ContainerToString(string id)`
  Called by: `.btnContainer_save_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.ContainerFromString()`** — L6217 — `public static ucMeter ContainerFromString(string data64, List<string> web_images = null)`
  Called by: `.btnContainer_load_Click()` (`Console/setup.cs`), `.btnContainer_dupe_Click()` (`Console/setup.cs`)
- **`.StoreSettings2()`** — L6416 — `public static bool StoreSettings2(ref Dictionary<string, string> a)`
  Called by: `.SaveOptions()` (`Console/setup.cs`)
- **`.RecoverContainer()`** — L6539 — `public static void RecoverContainer(string sId)`
  Called by: `.btnRecoverContainer_Click()` (`Console/setup.cs`)
- **`.RemoveMeterContainer()`** — L6558 — `public static void RemoveMeterContainer(string sId)`
  Removes meter container.
  Called by: `.ContainerFromString()` (same file), `.btnContainerDelete_Click()` (`Console/setup.cs`), `.btnContainer_load_Click()` (`Console/setup.cs`)

#### `FilterItemSnapFrequencies` (type, L248)

_No extracted members._

#### `CustomReadings` (type, L732)

- **`.setupReadings()`** — L750 — `private void setupReadings()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.formatNumber()`** — L839 — `private string formatNumber(double number)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.formatElapsedTimeCompact()`** — L849 — `private string formatElapsedTimeCompact(long elapsedSeconds)`
  Called by: `.GetReading()` (same file)
- **`.formatElapsedTime()`** — L870 — `private string formatElapsedTime(long elapsedSeconds)`
  Called by: `.GetReading()` (same file)
- **`.GetPlaceholders()`** — L899 — `public List<string> GetPlaceholders(string text)`
  Returns placeholders.
  Called by: `.onTimerElapsedCondition()` (same file)
- **`.TakeReading()`** — L923 — `public void TakeReading(Reading reading)`
  Called by: `.Update()` (same file), `.add_readings()` (same file)
- **`.IsCustomString()`** — L930 — `public bool IsCustomString(string custom)`
  Called by: `.GetReading()` (same file), `.add_readings()` (same file), `.expand_placeholders()` (same file), `.ZeroOut()` (same file)
- **`.GetReading()`** — L999 — `public object GetReading(string reading, clsMeter owningMeter)`
  Returns reading.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getIntPart()`** — L1249 — `private static int getIntPart(string input, string prefix)`
  Returns int part.
  Called by: `.GetReading()` (same file)
- **`.returnTuneStep()`** — L1267 — `public void returnTuneStep(string key)`
  Called by: `.GetReading()` (same file)
- **`.returnPAProfile()`** — L1278 — `public void returnPAProfile(string key)`
  Called by: `.GetReading()` (same file)
- **`.UpdateReadings()`** — L1289 — `public void UpdateReadings(string text)`
  Updates readings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addReading()`** — L1387 — `private void addReading(Reading reading, string text)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addReadingText()`** — L1391 — `private void addReadingText(string reading, string text)`
  Called by: `.UpdateReadings()` (same file)
- **`.GetAvailableReadings()`** — L1395 — `public List<string> GetAvailableReadings()`
  Returns available readings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsIGSettings` (type, L1433)

- **`.SetSetting()`** — L1510 — `public void SetSetting(string setting, object value)`
  Sets setting.
  Called by: `.TryParse()` (same file), `.SetMMIOGuid()` (same file), `.SetMMIOVariable()` (same file)
- **`.GetSetting()`** — L1524 — `public object GetSetting(string setting, Type type)`
  Returns setting.
  Called by: `.GetMMIOGuid()` (same file), `.GetMMIOVariable()` (same file)
- **`.ToString2()`** — L1573 — `public string ToString2()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse2()`** — L1630 — `public bool TryParse2(string str)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse()`** — L1734 — `public bool TryParse(string str)`
  return sRet; }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMMIOGuid()`** — L1972 — `public Guid GetMMIOGuid(int index)`
  Returns mmioguid.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMMIOGuid()`** — L1976 — `public void SetMMIOGuid(int index, Guid g)`
  Sets mmioguid.
  Called by: `.TryParse()` (same file)
- **`.GetMMIOVariable()`** — L1980 — `public string GetMMIOVariable(int index)`
  Returns mmiovariable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMMIOVariable()`** — L1984 — `public void SetMMIOVariable(int index, string variable)`
  Sets mmiovariable.
  Called by: `.TryParse()` (same file)

#### `GeneralOtherButtonSettings` (type, L4805)

_No extracted members._

#### `clsMeterItem` (type, L6625)

- **`.PrepareCalibration()`** — L6815 — `public void PrepareCalibration()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddPerc()`** — L6883 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public void AddPerc(clsPercCache pc)`
  [2.10.30.9]MW0LGE this perc cache code totally refactored, and only caches to 2 decimal precision for the dB value, and is keyed on the int version of that
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetPerc()`** — L6899 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public clsPercCache GetPerc(float value)`
  Returns perc.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HasPerc()`** — L6908 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool HasPerc(float value)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L7049 — `public virtual void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearHistory()`** — L7072 — `public virtual void ClearHistory()`
  Clears history.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.History()`** — L7075 — `public virtual void History(out float minHistory, out float maxHistory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToString()`** — L7124 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse()`** — L7128 — `public virtual bool TryParse(string val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleIncrement()`** — L7140 — `public virtual void HandleIncrement()`
  Handles increment.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleDecrement()`** — L7143 — `public virtual void HandleDecrement()`
  Handles decrement.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L7146 — `public virtual bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseClick()`** — L7152 — `public virtual void MouseClick(MouseEventArgs e)`
  mouse
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L7156 — `public virtual void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L7160 — `public virtual void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyDown()`** — L7198 — `public virtual void KeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyUp()`** — L7202 — `public virtual void KeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheel()`** — L7206 — `public virtual void MouseWheel(int number_of_moves)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Removing()`** — L7210 — `public virtual void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandPanelsChanged()`** — L7214 — `public virtual void BandPanelsChanged(bool gen, bool hf, bool vhf)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandChanged()`** — L7218 — `public virtual void BandChanged(Band oldBand, Band newBand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialise()`** — L7222 — `public virtual void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeChanged()`** — L7226 — `public virtual void ModeChanged(DSPMode oldMode, DSPMode newMode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuneStepIndexChanged()`** — L7230 — `public virtual void TuneStepIndexChanged(int old_index, int new_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterChanged()`** — L7234 — `public virtual void FilterChanged(Filter f, string name, int low, int high, bool vfoA, bool vfoB, int max_width, int max_shift)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXFilterChanged()`** — L7238 — `public virtual void TXFilterChanged(int low, int high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PAProfileChanged()`** — L7242 — `public virtual void PAProfileChanged(string name)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXProfileChanged()`** — L7246 — `public virtual void TXProfileChanged(string name)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXSpectrumGridMin()`** — L7250 — `public virtual void SetRXSpectrumGridMin(int min)`
  Sets rxspectrum grid min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXSpectrumGridMax()`** — L7254 — `public virtual void SetRXSpectrumGridMax(int min)`
  Sets rxspectrum grid max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXSpectrumGridMin()`** — L7258 — `public virtual void SetTXSpectrumGridMin(int min)`
  Sets txspectrum grid min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXSpectrumGridMax()`** — L7262 — `public virtual void SetTXSpectrumGridMax(int min)`
  Sets txspectrum grid max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXWaterfallMin()`** — L7266 — `public virtual void SetRXWaterfallMin(int min)`
  Sets rxwaterfall min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXWaterfallMax()`** — L7270 — `public virtual void SetRXWaterfallMax(int min)`
  Sets rxwaterfall max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXWaterfallMin()`** — L7274 — `public virtual void SetTXWaterfallMin(int min)`
  Sets txwaterfall min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXWaterfallMax()`** — L7278 — `public virtual void SetTXWaterfallMax(int min)`
  Sets txwaterfall max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallRXGradient()`** — L7282 — `public virtual void WaterfallRXGradient(System.Drawing.Color[] colours)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallTXGradient()`** — L7286 — `public virtual void WaterfallTXGradient(System.Drawing.Color[] colours)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ContainerEnabled()`** — L7340 — `public virtual void ContainerEnabled(string id, bool enabled)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ContainerHiddenByMacro()`** — L7344 — `public virtual void ContainerHiddenByMacro(string id, bool hidden)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `MeterItemType` (type, L6627)

_No extracted members._

#### `clsPercCache` (type, L6666)

_No extracted members._

#### `clsItemGroup` (type, L7521)

- **`.ToString()`** — L7543 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse()`** — L7565 — `public override bool TryParse(string str)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsClickBox` (type, L7597)

_No extracted members._

#### `clsSolidColour` (type, L7657)

_No extracted members._

#### `clsFadeCover` (type, L7691)

_No extracted members._

#### `clsFilterButtonBox` (type, L7700)

- **`.Initialise()`** — L7714 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterChanged()`** — L7731 — `public override void FilterChanged(Filter f, string name, int low, int high, bool vfoA, bool vfoB, int max_width, int max_shift)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitFilterButtons()`** — L7751 — `public void InitFilterButtons()`
  Inits filter buttons.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L7768 — `private void setupButtons()`
  SetText(1, index, new_name); }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L7861 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L7910 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L7923 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setFilter()`** — L7956 — `private void setFilter(Filter f)`
  Sets filter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsTunestepButtons` (type, L8025)

- **`.setupButtons()`** — L8059 — `private void setupButtons()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L8166 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L8215 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L8228 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuneStepIndexChanged()`** — L8245 — `public override void TuneStepIndexChanged(int old_index, int new_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsOtherButtons` (type, L8251)

- **`.Removing()`** — L8332 — `public override void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onPlayingChanged()`** — L8340 — `private void onPlayingChanged(bool playing, string id, string filename, bool isWdsp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onRecordingChanged()`** — L8356 — `private void onRecordingChanged(bool recording, string id, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnContainerVisible()`** — L8404 — `private void OnContainerVisible(string id, bool visible)`
  Handles/raises the container visible event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCatState()`** — L8456 — `private void OnCatState(int queue, Guid id, ScriptCommand sc, string cat_result)`
  Handles/raises the cat state event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMacroSettings()`** — L8812 — `public override OtherButtonMacroSettings GetMacroSettings(int macro)`
  Returns macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMacroSettings()`** — L8820 — `public override void SetMacroSettings(int macro, OtherButtonMacroSettings settings)`
  Sets macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ContainerHiddenByMacro()`** — L8850 — `public override void ContainerHiddenByMacro(string id, bool hidden)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateOn()`** — L8880 — `private void updateOn(OtherButtonId id, bool val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Initialise()`** — L8886 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVisibleBits()`** — L8890 — `public override int GetVisibleBits(int bit_group)`
  Returns visible bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVisibleBits()`** — L8895 — `public override void SetVisibleBits(int bit_group, int bit_field)`
  Sets visible bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.isVisible()`** — L8902 — `private bool isVisible(int bit_group, int bit)`
  Called by: `.setupButtons()` (same file)
- **`.try_index_from_group_bit()`** — L8908 — `private bool try_index_from_group_bit(int bit_group, int bit, out int index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L8971 — `private void setupButtons(bool init = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L9273 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L9322 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L9346 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.handleMacroButtonPress()`** — L9428 — `private void handleMacroButtonPress(int macro, int index)`
  Called by: `.MouseUp()` (same file)
- **`.MoveButton()`** — L9546 — `public override void MoveButton(int button_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L9551 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsAntennaButtonBox` (type, L9622)

- **`.setupButtons()`** — L9669 — `private void setupButtons()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.formatAux()`** — L9826 — `private string formatAux(string name)`
  Called by: `.setupButtons()` (same file)
- **`.AntennasChanged()`** — L9834 — `public void AntennasChanged(Band rx1_band, Band tx_band, double vfoa_freq, double tx_freq, int rxtx_swap)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L9899 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L9948 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L9961 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.toggleTxRxAnt()`** — L9986 — `private void toggleTxRxAnt()`
  Called by: `.MouseUp()` (same file)
- **`.setRXAntenna()`** — L9995 — `private void setRXAntenna(int antenna, Band b)`
  Sets rxantenna.
  Called by: `.MouseUp()` (same file)
- **`.setAuxAntenna()`** — L10005 — `private void setAuxAntenna(int antenna, Band b, bool byp, bool ext1, bool xvtr)`
  Sets aux antenna.
  Called by: `.MouseUp()` (same file)
- **`.setTXAntenna()`** — L10015 — `private void setTXAntenna(int antenna, Band b)`
  Sets txantenna.
  Called by: `.MouseUp()` (same file)

#### `clsModeButtonBox` (type, L10071)

- **`.Initialise()`** — L10084 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeChanged()`** — L10097 — `public override void ModeChanged(DSPMode oldMode, DSPMode newMode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L10109 — `private void setupButtons()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L10183 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L10232 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L10245 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.abortForLockedVFO()`** — L10262 — `private bool abortForLockedVFO()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setMode()`** — L10275 — `private void setMode(DSPMode m)`
  Sets mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsVoiceRecordPlay` (type, L10342)

- **`.Removing()`** — L10492 — `public override void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RecordToSlot()`** — L10509 — `public void RecordToSlot(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PlayFromSlot()`** — L10519 — `public void PlayFromSlot(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyUp()`** — L10597 — `public override void KeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyDown()`** — L10641 — `public override void KeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SlotDuration()`** — L10703 — `public float SlotDuration(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onJsonWritten()`** — L10708 — `private void onJsonWritten(string unique_id, clsAudioRecordPlayback.RecordingJsonModel json)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onPlayingChanged()`** — L10728 — `private void onPlayingChanged(bool playing, string id, string filename, bool isWdsp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onRecordingChanged()`** — L10765 — `private void onRecordingChanged(bool recording, string id, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupRepeatTimer()`** — L10798 — `private void setupRepeatTimer(bool setup, int slot, int delay)`
  Called by: `.onPlayingChanged()` (same file), `.clearRunningRepeat()` (same file)
- **`.enableAllSlots()`** — L10844 — `private void enableAllSlots(bool enabled)`
  Called by: `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file)
- **`.AbortVoiceRecordRepeatPlayback()`** — L10851 — `public void AbortVoiceRecordRepeatPlayback()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsPlaying()`** — L10865 — `public bool IsPlaying(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsRecording()`** — L10871 — `public bool IsRecording(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AtivateTime()`** — L10877 — `public DateTime AtivateTime(int slot)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetSlotLocked()`** — L10883 — `public void SetSlotLocked(int slot, bool locked)`
  Sets slot locked.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetSlotLocked()`** — L10889 — `public bool GetSlotLocked(int slot)`
  Returns slot locked.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.GetKeybind()`** — L10895 — `public Keys GetKeybind(int slot)`
  Returns keybind.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetKeybind()`** — L10900 — `public void SetKeybind(int slot, Keys keybind)`
  Sets keybind.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetUsesKeybind()`** — L10905 — `public bool GetUsesKeybind(int slot)`
  Returns uses keybind.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetUsesKeybind()`** — L10910 — `public void SetUsesKeybind(int slot, bool useskeybind)`
  Sets uses keybind.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.UsesKeybind()`** — L10927 — `public bool UsesKeybind(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCanRepeat()`** — L10940 — `public void SetCanRepeat(int slot, bool canrepeat)`
  Sets can repeat.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetCanRepeat()`** — L10948 — `public bool GetCanRepeat(int slot)`
  Returns can repeat.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetRepeatDelay()`** — L10954 — `public void SetRepeatDelay(int slot, int repeatdelay)`
  Sets repeat delay.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetRepeatDelay()`** — L10961 — `public int GetRepeatDelay(int slot)`
  Returns repeat delay.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.GetRepeatEnabled()`** — L10967 — `public bool GetRepeatEnabled(int slot)`
  Returns repeat enabled.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetRepeatEnabled()`** — L10973 — `public void SetRepeatEnabled(int slot, bool enabled)`
  Sets repeat enabled.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetDelayElapsed()`** — L10979 — `public float GetDelayElapsed(int slot)`
  Returns delay elapsed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetGainAdjust()`** — L10988 — `public double GetGainAdjust(int slot)`
  Returns gain adjust.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetGainAdjust()`** — L10994 — `public void SetGainAdjust(int slot, double gain_adjust)`
  Sets gain adjust.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetIgnorePlayTempChanges()`** — L11000 — `public bool GetIgnorePlayTempChanges(int slot)`
  Returns ignore play temp changes.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetIgnorePlayTempChanges()`** — L11006 — `public void SetIgnorePlayTempChanges(int slot, bool ignore)`
  Sets ignore play temp changes.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.GetIgnoreRecordTempChanges()`** — L11012 — `public bool GetIgnoreRecordTempChanges(int slot)`
  Returns ignore record temp changes.
  Called by: `.GetSettingsForMeterGroup()` (same file)
- **`.SetIgnoreRecordTempChanges()`** — L11018 — `public void SetIgnoreRecordTempChanges(int slot, bool ignore)`
  Sets ignore record temp changes.
  Called by: `.ApplySettingsForMeterGroup()` (same file)
- **`.Initialise()`** — L11040 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L11069 — `private void setupButtons()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveButton()`** — L11207 — `public override void MoveButton(int button_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L11217 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L11266 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L11295 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clearRunningRepeat()`** — L11379 — `private void clearRunningRepeat(int slot, bool force = false)`
  Called by: `.Removing()` (same file), `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file), `.AbortVoiceRecordRepeatPlayback()` (same file), `.SetCanRepeat()` (same file), `.SetRepeatDelay()` (same file) — and 2 more
- **`.handleClicked()`** — L11386 — `private void handleClicked(int slot, bool right_click, bool long_hold, bool from_repeat_timer = false, bool from_keybind = false)`
  Called by: `.RecordToSlot()` (same file), `.PlayFromSlot()` (same file), `.KeyUp()` (same file), `.setupRepeatTimer()` (same file), `.MouseUp()` (same file)

#### `clsWaveRecord` (type, L11602)

- **`.Initialise()`** — L11754 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Removing()`** — L11759 — `public override void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onPlayingChanged()`** — L11782 — `private void onPlayingChanged(bool playing, string id, string filename, bool isWdsp)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.onRecordingChanged()`** — L11798 — `private void onRecordingChanged(bool recording, string id, string filename)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CanAcceptFiles()`** — L12073 — `public bool CanAcceptFiles(string[] filePaths)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddFiles()`** — L12078 — `public void AddFiles(string[] filePaths)`
  Adds files.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsPlaying()`** — L12107 — `public bool IsPlaying(string filePath)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRenderLayout()`** — L12115 — `public void SetRenderLayout( SharpDX.RectangleF contentRect, SharpDX.RectangleF scrollTrackRect, SharpDX.RectangleF scrollThumbRect, float rowPitch,`
  Sets render layout.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateDragFromMouse()`** — L12137 — `public void UpdateDragFromMouse()`
  Updates drag from mouse.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HitTest()`** — L12143 — `public WaveRecordHitRegion HitTest(PointF point)`
  Called by: `.MouseDown()` (same file), `.MouseUp()` (same file)
- **`.Update()`** — L12157 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheel()`** — L12162 — `public override void MouseWheel(int number_of_moves)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L12189 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L12236 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RefreshEntries()`** — L12285 — `public void RefreshEntries()`
  Refreshes entries.
  Called by: `.Initialise()` (same file), `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file), `.AddFiles()` (same file), `.handlePlay()` (same file), `.handleDelete()` (same file) — and 1 more
- **`.buildJsonDataDisplay()`** — L12347 — `private WaveRecordJsonDataDisplay buildJsonDataDisplay(string wavPath)`
  Called by: `.RefreshEntries()` (same file)
- **`.sanitiseJsonDataValue()`** — L12370 — `private static string sanitiseJsonDataValue(string value, bool frequency)`
  Called by: `.buildJsonDataDisplay()` (same file), `.formatJsonDataUtc()` (same file)
- **`.formatJsonDataDuration()`** — L12392 — `private static string formatJsonDataDuration(double? durationSeconds)`
  Called by: `.buildJsonDataDisplay()` (same file)
- **`.formatJsonDataAudio()`** — L12401 — `private static string formatJsonDataAudio(short bitDepth, int sampleRate, short channels)`
  Called by: `.buildJsonDataDisplay()` (same file)
- **`.formatJsonDataUtc()`** — L12412 — `private static string formatJsonDataUtc(string value)`
  Called by: `.buildJsonDataDisplay()` (same file)
- **`.containsFilePathLocked()`** — L12426 — `private bool containsFilePathLocked(string filePath)`
  Called by: `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file)
- **`.sanitiseStoredPaths()`** — L12432 — `private static string[] sanitiseStoredPaths(IEnumerable<string> filePaths)`
  Called by: `.CanAcceptFiles()` (same file), `.AddFiles()` (same file)
- **`.pathArraysEqual()`** — L12452 — `private static bool pathArraysEqual(string[] left, string[] right)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.shortArraysEqual()`** — L12466 — `private static bool shortArraysEqual(short[] left, short[] right)`
  Called by: `.RefreshEntries()` (same file), `.moveEntry()` (same file)
- **`.sanitiseOrderMap()`** — L12482 — `private static short[] sanitiseOrderMap(IEnumerable<short> rawOrderMap, int itemCount)`
  Called by: `.AddFiles()` (same file), `.RefreshEntries()` (same file), `.removeFromOrderMap()` (same file), `.moveEntry()` (same file)
- **`.removeFromOrderMap()`** — L12520 — `private static short[] removeFromOrderMap(short[] orderMap, int removedIndex, int oldCount)`
  Called by: `.handleDelete()` (same file)
- **`.handleHit()`** — L12547 — `private void handleHit(WaveRecordHitRegion hit)`
  Called by: `.MouseUp()` (same file)
- **`.handlePlay()`** — L12566 — `private void handlePlay(string filePath)`
  Called by: `.handleHit()` (same file)
- **`.handleDelete()`** — L12614 — `private void handleDelete(string filePath)`
  Called by: `.handleHit()` (same file)
- **`.moveEntry()`** — L12660 — `private void moveEntry(int fromDisplayIndex, int toDisplayIndex, bool swapOnly)`
  Called by: `.MouseUp()` (same file)
- **`.adjustScroll()`** — L12707 — `private void adjustScroll(float delta)`
  Called by: `.MouseWheel()` (same file)
- **`.updateReorderDrag()`** — L12718 — `private void updateReorderDrag()`
  Called by: `.UpdateDragFromMouse()` (same file)
- **`.updateScrollDrag()`** — L12738 — `private void updateScrollDrag()`
  Called by: `.UpdateDragFromMouse()` (same file), `.Update()` (same file)
- **`.scrollTrackToLocked()`** — L12755 — `private void scrollTrackToLocked(float mouseY)`
  Called by: `.MouseDown()` (same file)
- **`.rowIndexFromPointLocked()`** — L12772 — `private int rowIndexFromPointLocked(PointF point, bool clampToValid)`
  Called by: `.MouseDown()` (same file), `.updateReorderDrag()` (same file)
- **`.clampScrollLocked()`** — L12790 — `private void clampScrollLocked()`
  Called by: `.SetRenderLayout()` (same file), `.RefreshEntries()` (same file), `.adjustScroll()` (same file), `.updateScrollDrag()` (same file), `.scrollTrackToLocked()` (same file)

#### `WaveRecordHitType` (type, L11604)

_No extracted members._

#### `WaveRecordEntry` (type, L11614)

_No extracted members._

#### `WaveRecordJsonDataDisplay` (type, L11622)

_No extracted members._

#### `WaveRecordHitRegion` (type, L11645)

_No extracted members._

#### `clsBandButtonBox` (type, L12802)

- **`.Initialise()`** — L12817 — `public override void Initialise()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandPanelsChanged()`** — L12852 — `public override void BandPanelsChanged(bool gen, bool hf, bool vhf)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VHFUpdate()`** — L12861 — `public void VHFUpdate(int index, bool enabled, string text)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandChanged()`** — L12876 — `public override void BandChanged(Band oldBand, Band newBand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L12930 — `private void setupButtons()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupClick()`** — L13084 — `private void setupClick(bool setup)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L13133 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L13146 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setBand()`** — L13224 — `private void setBand(Band b)`
  Sets band.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.abortForLockedVFO()`** — L13244 — `private bool abortForLockedVFO()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsDiscordButtonBox` (type, L13303)

- **`.OnReady()`** — L13324 — `private void OnReady()`
  Handles/raises the ready event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDisconnected()`** — L13334 — `private void OnDisconnected()`
  Handles/raises the disconnected event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupButtons()`** — L13345 — `private void setupButtons()`
  Called by: `.Initialise()` (same file), `.InitFilterButtons()` (same file), `.TuneStepIndexChanged()` (same file), `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file), `.SetVisibleBits()` (same file) — and 2 more
- **`.formatNumber()`** — L13413 — `private string formatNumber(double number)`
  Called by: `.GetReading()` (same file), `.MouseUp()` (same file)
- **`.setupClick()`** — L13427 — `private void setupClick(bool setup)`
  Called by: `.MouseDown()` (same file), `.MouseUp()` (same file)
- **`.MouseDown()`** — L13476 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L13489 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.sendMsg()`** — L13573 — `private void sendMsg(string msg)`
  Called by: `.MouseUp()` (same file)

#### `clsButtonBox` (type, L13627)

- **`.setupArrays()`** — L13727 — `private void setupArrays()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ResetButtons()`** — L13797 — `public void ResetButtons()`
  Resets buttons.
  Called by: `.setupArrays()` (same file)
- **`.GetVisibleBits()`** — L13893 — `public virtual int GetVisibleBits(int bit_group)`
  Returns visible bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVisibleBits()`** — L13894 — `public virtual void SetVisibleBits(int bit_group, int bit_field)`
  Sets visible bits.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetMacroSettings()`** — L13895 — `public virtual void SetMacroSettings(int macro, OtherButtonMacroSettings settings)`
  Sets macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMacroSettings()`** — L13896 — `public virtual OtherButtonMacroSettings GetMacroSettings(int macro)`
  Returns macro settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetFillColour()`** — L13922 — `public void SetFillColour(int bank, int button, System.Drawing.Color colour)`
  Sets fill colour.
  Called by: `.setupButtons()` (same file)
- **`.GetFillColour()`** — L13927 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetFillColour(int bank, int button)`
  Returns fill colour.
  Called by: `.setupButtons()` (same file)
- **`.SetHoverColour()`** — L13933 — `public void SetHoverColour(int bank, int button, System.Drawing.Color colour)`
  Sets hover colour.
  Called by: `.setupButtons()` (same file)
- **`.GetHoverColour()`** — L13938 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetHoverColour(int bank, int button)`
  Returns hover colour.
  Called by: `.setupButtons()` (same file)
- **`.SetClickColour()`** — L13944 — `public void SetClickColour(int bank, int button, System.Drawing.Color colour)`
  Sets click colour.
  Called by: `.setupButtons()` (same file)
- **`.GetClickColour()`** — L13949 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetClickColour(int bank, int button)`
  Returns click colour.
  Called by: `.setupButtons()` (same file)
- **`.SetBorderColour()`** — L13955 — `public void SetBorderColour(int bank, int button, System.Drawing.Color colour)`
  Sets border colour.
  Called by: `.setupButtons()` (same file)
- **`.GetBorderColour()`** — L13960 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetBorderColour(int bank, int button)`
  Returns border colour.
  Called by: `.setupButtons()` (same file)
- **`.SetUseOffColour()`** — L13966 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public void SetUseOffColour(int bank, int button, bool use)`
  Sets use off colour.
  Called by: `.setupButtons()` (same file)
- **`.GetUseOffColour()`** — L13972 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetUseOffColour(int bank, int button)`
  Returns use off colour.
  Called by: `.setupButtons()` (same file)
- **`.SetUseIndicator()`** — L13978 — `public void SetUseIndicator(int bank, int button, bool use)`
  Sets use indicator.
  Called by: `.setupButtons()` (same file)
- **`.GetUseIndicator()`** — L13983 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetUseIndicator(int bank, int button)`
  Returns use indicator.
  Called by: `.setupButtons()` (same file)
- **`.SetOn()`** — L13989 — `public void SetOn(int bank, int button, bool on)`
  Sets on.
  Called by: `.FilterChanged()` (same file), `.setupButtons()` (same file), `.OnContainerVisible()` (same file), `.OnCatState()` (same file), `.ContainerHiddenByMacro()` (same file), `.updateOn()` (same file) — and 4 more
- **`.GetOn()`** — L13994 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetOn(int bank, int button)`
  Returns on.
  Called by: `.setupButtons()` (same file), `.MouseUp()` (same file), `.handleMacroButtonPress()` (same file)
- **`.SetIndicatorWidth()`** — L14000 — `public void SetIndicatorWidth(int bank, int button, float width)`
  Sets indicator width.
  Called by: `.setupButtons()` (same file)
- **`.GetIndicatorWidth()`** — L14005 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public float GetIndicatorWidth(int bank, int button)`
  Returns indicator width.
  Called by: `.setupButtons()` (same file)
- **`.SetOnColour()`** — L14011 — `public void SetOnColour(int bank, int button, System.Drawing.Color colour)`
  Sets on colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file)
- **`.GetOnColour()`** — L14016 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetOnColour(int bank, int button)`
  Returns on colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file)
- **`.SetOffColour()`** — L14022 — `public void SetOffColour(int bank, int button, System.Drawing.Color colour)`
  Sets off colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file)
- **`.GetOffColour()`** — L14027 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetOffColour(int bank, int button)`
  Returns off colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file)
- **`.SetFontFamily()`** — L14033 — `public void SetFontFamily(int bank, int button, string font_family)`
  Sets font family.
  Called by: `.setupButtons()` (same file)
- **`.GetFontFamily()`** — L14038 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public string GetFontFamily(int bank, int button)`
  Returns font family.
  Called by: `.setupButtons()` (same file)
- **`.SetFontStyle()`** — L14044 — `public void SetFontStyle(int bank, int button, FontStyle font_style)`
  Sets font style.
  Called by: `.setupButtons()` (same file)
- **`.GetFontStyle()`** — L14049 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public FontStyle GetFontStyle(int bank, int button)`
  Returns font style.
  Called by: `.setupButtons()` (same file)
- **`.SetFontSize()`** — L14055 — `public void SetFontSize(int bank, int button, float size)`
  Sets font size.
  Called by: `.setupButtons()` (same file)
- **`.GetFontSize()`** — L14060 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public float GetFontSize(int bank, int button)`
  Returns font size.
  Called by: `.setupButtons()` (same file)
- **`.SetFontColour()`** — L14066 — `public void SetFontColour(int bank, int button, System.Drawing.Color colour)`
  Sets font colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file), `.VHFUpdate()` (same file), `.OnReady()` (same file), `.OnDisconnected()` (same file)
- **`.GetFontColour()`** — L14071 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public System.Drawing.Color GetFontColour(int bank, int button)`
  Returns font colour.
  Called by: `.setupButtons()` (same file), `.Update()` (same file), `.OnReady()` (same file), `.OnDisconnected()` (same file)
- **`.SetText()`** — L14077 — `public void SetText(int bank, int button, string text)`
  Sets text.
  Called by: `.FilterChanged()` (same file), `.setupButtons()` (same file), `.OnCatState()` (same file), `.ContainerHiddenByMacro()` (same file), `.handleMacroButtonPress()` (same file), `.Update()` (same file) — and 2 more
- **`.GetText()`** — L14084 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public string GetText(int bank, int button)`
  Returns text.
  Called by: `.setupButtons()` (same file), `.OnCatState()` (same file), `.ContainerHiddenByMacro()` (same file), `.handleMacroButtonPress()` (same file), `.Update()` (same file), `.GetSettingsForMeterGroup()` (same file)
- **`.GetTextUpdateChanged()`** — L14090 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public string GetTextUpdateChanged(int bank, int button)`
  Returns text update changed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetTextChanged()`** — L14099 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetTextChanged(int bank, int button)`
  Returns text changed.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIconOn()`** — L14105 — `public void SetIconOn(int bank, int button, string text)`
  Sets icon on.
  Called by: `.setupButtons()` (same file)
- **`.GetIconOn()`** — L14110 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public string GetIconOn(int bank, int button)`
  Returns icon on.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetIconOff()`** — L14116 — `public void SetIconOff(int bank, int button, string text)`
  Sets icon off.
  Called by: `.setupButtons()` (same file)
- **`.GetIconOff()`** — L14121 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public string GetIconOff(int bank, int button)`
  Returns icon off.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUseIcon()`** — L14127 — `public void SetUseIcon(int bank, int button, bool icon)`
  Sets use icon.
  Called by: `.setupButtons()` (same file)
- **`.GetUseIcon()`** — L14132 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetUseIcon(int bank, int button)`
  Returns use icon.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetEnabled()`** — L14138 — `public void SetEnabled(int bank, int button, bool enabled)`
  Sets enabled.
  Called by: `.setupButtons()` (same file), `.onPlayingChanged()` (same file), `.onRecordingChanged()` (same file), `.Update()` (same file), `.enableAllSlots()` (same file), `.VHFUpdate()` (same file) — and 2 more
- **`.GetEnabled()`** — L14143 — `public bool GetEnabled(int bank, int button)`
  Returns enabled.
  Called by: `.setupButtons()` (same file), `.MouseDown()` (same file), `.MouseUp()` (same file), `.Update()` (same file), `.KeyUp()` (same file), `.KeyDown()` (same file) — and 1 more
- **`.SetVisible()`** — L14148 — `public void SetVisible(int bank, int button, bool enabled)`
  Sets visible.
  Called by: `.setupButtons()` (same file)
- **`.GetVisible()`** — L14164 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool GetVisible(int bank, int button)`
  Returns visible.
  Called by: `.setupButtons()` (same file)
- **`.SetIndicatorType()`** — L14170 — `public void SetIndicatorType(int bank, int button, IndicatorType type)`
  Sets indicator type.
  Called by: `.setupButtons()` (same file)
- **`.GetIndicatorType()`** — L14175 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public IndicatorType GetIndicatorType(int bank, int button)`
  Returns indicator type.
  Called by: `.setupButtons()` (same file)
- **`.MoveButton()`** — L14191 — `public virtual void MoveButton(int button_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `IndicatorType` (type, L13629)

_No extracted members._

#### `clsVfoDisplay` (type, L14201)

- **`.getVfo()`** — L14397 — `private double getVfo()`
  Returns vfo.
  Called by: `.KeyDown()` (same file)
- **`.setVfo()`** — L14430 — `private void setVfo(double value)`
  Sets vfo.
  Called by: `.KeyDown()` (same file)
- **`.abortForLockedVFO()`** — L14462 — `private bool abortForLockedVFO()`
  Called by: `.setMode()` (same file), `.setBand()` (same file), `.setVfo()` (same file), `.adjustVfo()` (same file), `.MouseUp()` (same file)
- **`.adjustVfo()`** — L14495 — `private void adjustVfo(double adjustment)`
  Called by: `.MouseWheel()` (same file), `.MouseUp()` (same file)
- **`.KeyDown()`** — L14527 — `public override void KeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.findOnePosition()`** — L14560 — `private int findOnePosition(double number)`
  Called by: `.KeyDown()` (same file)
- **`.MouseWheel()`** — L14572 — `public override void MouseWheel(int number_of_moves)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopBandStack()`** — L14584 — `public void PopBandStack()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PopFilterMenu()`** — L14599 — `public void PopFilterMenu()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L14613 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L14631 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setTuneStep()`** — L14809 — `private bool setTuneStep(bool vfoB)`
  Sets tune step.
  Called by: `.MouseUp()` (same file)
- **`.setLock()`** — L14879 — `private void setLock(bool vfoB)`
  Sets lock.
  Called by: `.MouseUp()` (same file)
- **`.setVfoSync()`** — L14897 — `private void setVfoSync()`
  Sets vfo sync.
  Called by: `.MouseUp()` (same file)
- **`.setTX()`** — L14905 — `private void setTX(bool vfoB)`
  Sets tx.
  Called by: `.MouseUp()` (same file)
- **`.toggleSplit()`** — L14933 — `private void toggleSplit(MouseEventArgs e)`
  Called by: `.MouseUp()` (same file)
- **`.setFilter()`** — L14948 — `private bool setFilter(bool vfoB)`
  Sets filter.
  Called by: `.MouseUp()` (same file)
- **`.setMode()`** — L14991 — `private bool setMode(bool vfoB)`
  Sets mode.
  Called by: `.MouseUp()` (same file)
- **`.setBand()`** — L15030 — `private bool setBand(bool vfoB)`
  Sets band.
  Called by: `.MouseUp()` (same file)

#### `renderState` (type, L14203)

_No extracted members._

#### `buttonState` (type, L14211)

_No extracted members._

#### `DSPModeForModeDisplay` (type, L14232)

_No extracted members._

#### `VFODisplayMode` (type, L14247)

_No extracted members._

#### `clsClock` (type, L15395)

_No extracted members._

#### `clsWebImage` (type, L15485)

- **`.OnWebImageRemoved()`** — L15571 — `private void OnWebImageRemoved(object sender, string fourchar)`
  Handles/raises the web image removed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnShowWebImageBackground()`** — L15578 — `private void OnShowWebImageBackground(object sender, string fourchar)`
  Handles/raises the show web image background event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.timerCallback()`** — L15644 — `private void timerCallback(object state)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L15667 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearBackgroundFourChar()`** — L15737 — `public void ClearBackgroundFourChar()`
  Clears background four char.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Removing()`** — L15767 — `public override void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnState()`** — L15919 — `private void OnState(object sender, ImageFetcher.StateEventArgs e)`
  Handles/raises the state event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnImage()`** — L15940 — `private void OnImage(object sender, Guid guid)`
  Handles/raises the image event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsImage` (type, L16025)

_No extracted members._

#### `clsScaleItem` (type, L16105)

- **`.ToString()`** — L16191 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsNeedleScalePwrItem` (type, L16208)

- **`.ToString()`** — L16311 — `public override string ToString()`
  Returns the string representation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsRotatorItem` (type, L16362)

- **`.SendRotatorMessage()`** — L16475 — `public void SendRotatorMessage(bool dragging_rotator_ele, float dragging_rotator_degrees, bool stop)`
  Sends rotator message.
  Called by: `.MouseUp()` (same file), `.MouseDown()` (same file)
- **`.Update()`** — L16564 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  public bool ShowElevation { get { return _show_elevation; } set { _show_elevation = value; } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L16696 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L16704 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseDown()`** — L16711 — `public override void MouseDown(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `RotatorMode` (type, L16364)

_No extracted members._

#### `clsDialDisplay` (type, L16719)

- **`.undoTuneStep()`** — L17099 — `private void undoTuneStep()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L17116 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L17134 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheel()`** — L17157 — `public override void MouseWheel(int number_of_moves)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsMagicEyeItem` (type, L17175)

- **`.Update()`** — L17208 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearHistory()`** — L17319 — `public override void ClearHistory()`
  Clears history.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.History()`** — L17328 — `public override void History(out float minHistory, out float maxHistory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L17351 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsDataOut` (type, L17367)

- **`.Update()`** — L17376 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsSpacerItem` (type, L17436)

_No extracted members._

#### `clsHistoryItem` (type, L17469)

- **`.ZeroOut()`** — L17594 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.addReading()`** — L17755 — `private void addReading(int axis, float value)`
  private float test = -140; private bool testadd = true;
  Called by: `.UpdateReadings()` (same file), `.Update()` (same file)
- **`.Update()`** — L18096 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `HistoryData` (type, L17471)

_No extracted members._

#### `clsFilterItem` (type, L18172)

- **`.Initialise()`** — L18366 — `public override void Initialise()`
  Called by: `.UpdateItems()` (same file)
- **`.WaterfallRXGradient()`** — L18498 — `public override void WaterfallRXGradient(System.Drawing.Color[] colours)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallTXGradient()`** — L18509 — `public override void WaterfallTXGradient(System.Drawing.Color[] colours)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.findNearestVGrid()`** — L18586 — `public int findNearestVGrid(int hz)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Removing()`** — L18609 — `public override void Removing()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXProfileChanged()`** — L18796 — `public override void TXProfileChanged(string name)`
  public override void PAProfileChanged(string name) { _pa_profile = name; }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXFilterChanged()`** — L18800 — `public override void TXFilterChanged(int low, int high)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterChanged()`** — L18819 — `public override void FilterChanged(Filter f, string name, int low, int high, bool vfoA, bool vfoB, int max_width, int max_shift)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseUp()`** — L18858 — `public override void MouseUp(MouseEventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MouseWheel()`** — L19013 — `public override void MouseWheel(int number_of_moves)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AdjustNotch()`** — L19351 — `public void AdjustNotch(float delta)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Shift()`** — L19378 — `public void Shift(float shift_hz)`
  Called by: `.MouseUp()` (same file), `.MouseWheel()` (same file)
- **`.buildSpectrumGreyScale()`** — L19560 — `private void buildSpectrumGreyScale(bool do_low, bool do_high)`
  Called by: `.Initialise()` (same file), `.TXFilterChanged()` (same file), `.FilterChanged()` (same file), `.Update()` (same file)
- **`.Update()`** — L19616 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXSpectrumGridMin()`** — L19931 — `public override void SetRXSpectrumGridMin(int min)`
  Sets rxspectrum grid min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXSpectrumGridMax()`** — L19939 — `public override void SetRXSpectrumGridMax(int max)`
  Sets rxspectrum grid max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXSpectrumGridMin()`** — L19947 — `public override void SetTXSpectrumGridMin(int min)`
  Sets txspectrum grid min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXSpectrumGridMax()`** — L19955 — `public override void SetTXSpectrumGridMax(int max)`
  Sets txspectrum grid max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXWaterfallMin()`** — L19964 — `public override void SetRXWaterfallMin(int min)`
  Sets rxwaterfall min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRXWaterfallMax()`** — L19972 — `public override void SetRXWaterfallMax(int max)`
  Sets rxwaterfall max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXWaterfallMin()`** — L19980 — `public override void SetTXWaterfallMin(int min)`
  Sets txwaterfall min.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTXWaterfallMax()`** — L19988 — `public override void SetTXWaterfallMax(int max)`
  Sets txwaterfall max.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `FIWaterfallPalette` (type, L18174)

_No extracted members._

#### `FIDisplayMode` (type, L18185)

_No extracted members._

#### `clsTextOverlay` (type, L20066)

- **`.Update()`** — L20529 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.parseText()`** — L20587 — `private string parseText(int rx, int text_line)`
  Called by: `.Update()` (same file)
- **`.ZeroOut()`** — L20718 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsLed` (type, L20768)

- **`.Removing()`** — L20935 — `public override void Removing()`
  Called by: `.removeMeterItem()` (same file)
- **`.onTimerElapsedCondition()`** — L20942 — `private void onTimerElapsedCondition()`
  Called by: `.RebuildCondition()` (same file)
- **`.UpdateReadings()`** — L20961 — `public void UpdateReadings()`
  Updates readings.
  Called by: `.add_readings()` (same file)
- **`.add_readings()`** — L20965 — `private void add_readings(bool update_custom = false)`
  Called by: `.onTimerElapsedCondition()` (same file), `.Update()` (same file)
- **`.expand_placeholders()`** — L21007 — `private string expand_placeholders(string expr, int rx)`
  Called by: `.onTimerElapsedCondition()` (same file)
- **`.RebuildCondition()`** — L21066 — `public void RebuildCondition()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Update()`** — L21177 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.stopMox()`** — L21206 — `private void stopMox()`
  Called by: `.Update()` (same file)
- **`.ZeroOut()`** — L21213 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Led_Shape` (type, L20770)

_No extracted members._

#### `Led_Style` (type, L20776)

_No extracted members._

#### `clsBarItem` (type, L21237)

- **`.Update()`** — L21323 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearHistory()`** — L21449 — `public override void ClearHistory()`
  Clears history.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.History()`** — L21458 — `public override void History(out float minHistory, out float maxHistory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleIncrement()`** — L21539 — `public override void HandleIncrement()`
  Handles increment.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleDecrement()`** — L21557 — `public override void HandleDecrement()`
  Handles decrement.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L21584 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `Units` (type, L21239)

_No extracted members._

#### `BarStyle` (type, L21246)

_No extracted members._

#### `clsSignalText` (type, L21606)

- **`.Update()`** — L21677 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearHistory()`** — L21763 — `public override void ClearHistory()`
  Clears history.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.History()`** — L21772 — `public override void History(out float minHistory, out float maxHistory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleIncrement()`** — L21836 — `public override void HandleIncrement()`
  Handles increment.
  Called by: `.incrementMeterItem()` (same file)
- **`.HandleDecrement()`** — L21852 — `public override void HandleDecrement()`
  Handles decrement.
  Called by: `.decrementMeterItem()` (same file)
- **`.ZeroOut()`** — L21870 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsNeedleItem` (type, L21900)

- **`.Update()`** — L21974 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearHistory()`** — L22131 — `public override void ClearHistory()`
  Clears history.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.History()`** — L22140 — `public override void History(out float minHistory, out float maxHistory)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L22185 — `public override bool ZeroOut(ref Dictionary<Reading, float> values, int rx)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `NeedlePlacement` (type, L21902)

_No extracted members._

#### `NeedleStyle` (type, L21909)

_No extracted members._

#### `NeedleDirection` (type, L21915)

_No extracted members._

#### `clsText` (type, L22251)

- **`.updateReadingText()`** — L22274 — `private void updateReadingText(float reading)`
  Called by: `.Update()` (same file)
- **`.Update()`** — L22293 — `public override void Update(int rx, ref List<Reading> readingsUsed, Dictionary<Reading, object> all_list_item_readings = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `clsMeter` (type, L22359)

- **`.addMeterItem()`** — L22620 — `private void addMeterItem(clsMeterItem mi)`
  Called by: `.addSMeterBar()` (same file), `.AddADCMaxMag()` (same file), `.AddSMeterBarText()` (same file), `.AddADCBar()` (same file), `.AddPBSNRBar()` (same file), `.AddAGCGainBar()` (same file) — and 37 more
- **`.MeterVariablesReadingString()`** — L22628 — `public string MeterVariablesReadingString(MeterType meter, int variable_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MeterVariables()`** — L22684 — `public int MeterVariables(MeterType meter)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddMeter()`** — L22740 — `public void AddMeter(MeterType meter, clsItemGroup restoreIg = null)`
  Adds meter.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AddSMeterBarSignal()`** — L22822 — `public string AddSMeterBarSignal(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds smeter bar signal.
  Called by: `.AddMeter()` (same file)
- **`.AddSMeterBarSignalAvg()`** — L22826 — `public string AddSMeterBarSignalAvg(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds smeter bar signal avg.
  Called by: `.AddMeter()` (same file)
- **`.AddSMeterBarMaxBin()`** — L22830 — `public string AddSMeterBarMaxBin(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds smeter bar max bin.
  Called by: `.AddMeter()` (same file)
- **`.getFadeCover()`** — L22834 — `private clsFadeCover getFadeCover(string sId, bool ignore_empty_check = false)`
  Returns fade cover.
  Called by: `.addSMeterBar()` (same file), `.AddADCMaxMag()` (same file), `.AddSMeterBarText()` (same file), `.AddADCBar()` (same file), `.AddPBSNRBar()` (same file), `.AddAGCGainBar()` (same file) — and 35 more
- **`.addSMeterBar()`** — L22846 — `private string addSMeterBar(int nMSupdate, Reading reading, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Called by: `.AddSMeterBarSignal()` (same file), `.AddSMeterBarSignalAvg()` (same file), `.AddSMeterBarMaxBin()` (same file)
- **`.AddADCMaxMag()`** — L22940 — `public string AddADCMaxMag(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds adcmax mag.
  Called by: `.AddMeter()` (same file)
- **`.AddSMeterBarText()`** — L23001 — `private string AddSMeterBarText(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds smeter bar text.
  Called by: `.AddMeter()` (same file)
- **`.AddADCBar()`** — L23063 — `public string AddADCBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds adcbar.
  Called by: `.AddMeter()` (same file)
- **`.AddPBSNRBar()`** — L23151 — `public string AddPBSNRBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds pbsnrbar.
  Called by: `.AddMeter()` (same file)
- **`.AddAGCGainBar()`** — L23222 — `public string AddAGCGainBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds agcgain bar.
  Called by: `.AddMeter()` (same file)
- **`.AddAGCBar()`** — L23284 — `public string AddAGCBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds agcbar.
  Called by: `.AddMeter()` (same file)
- **`.AddCustomBar()`** — L23374 — `public string AddCustomBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds custom bar.
  Called by: `.AddMeter()` (same file)
- **`.AddRotator()`** — L23466 — `public string AddRotator(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds rotator.
  Called by: `.AddMeter()` (same file)
- **`.AddMagicEye()`** — L23572 — `public string AddMagicEye(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds magic eye.
  Called by: `.AddMeter()` (same file)
- **`.AddDial()`** — L23619 — `public string AddDial(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds dial.
  Called by: `.AddMeter()` (same file)
- **`.AddSpacer()`** — L23647 — `public string AddSpacer(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds spacer.
  Called by: `.AddMeter()` (same file)
- **`.AddWebImage()`** — L23676 — `public string AddWebImage(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds web image.
  Called by: `.AddMeter()` (same file)
- **`.AddDataOut()`** — L23704 — `public string AddDataOut(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds data out.
  Called by: `.AddMeter()` (same file)
- **`.AddTextOverlay()`** — L23728 — `public string AddTextOverlay(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds text overlay.
  Called by: `.AddMeter()` (same file)
- **`.AddLed()`** — L23756 — `public string AddLed(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds led.
  Called by: `.AddMeter()` (same file)
- **`.AddAnanMM()`** — L23784 — `public string AddAnanMM(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds anan mm.
  Called by: `.AddMeter()` (same file)
- **`.AddCrossNeedle()`** — L24140 — `public string AddCrossNeedle(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds cross needle.
  Called by: `.AddMeter()` (same file)
- **`.AddMicBar()`** — L24327 — `public string AddMicBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds mic bar.
  Called by: `.AddMeter()` (same file)
- **`.AddEQBar()`** — L24414 — `public string AddEQBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds eqbar.
  Called by: `.AddMeter()` (same file)
- **`.AddLevelerBar()`** — L24502 — `public string AddLevelerBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds leveler bar.
  Called by: `.AddMeter()` (same file)
- **`.AddLevelerGainBar()`** — L24588 — `public string AddLevelerGainBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds leveler gain bar.
  Called by: `.AddMeter()` (same file)
- **`.AddALCBar()`** — L24649 — `public string AddALCBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds alcbar.
  Called by: `.AddMeter()` (same file)
- **`.AddALCGainBar()`** — L24735 — `public string AddALCGainBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds alcgain bar.
  Called by: `.AddMeter()` (same file)
- **`.AddALCGroupBar()`** — L24796 — `public string AddALCGroupBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds alcgroup bar.
  Called by: `.AddMeter()` (same file)
- **`.AddCFCBar()`** — L24857 — `public string AddCFCBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds cfcbar.
  Called by: `.AddMeter()` (same file)
- **`.AddCFCGainBar()`** — L24943 — `public string AddCFCGainBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds cfcgain bar.
  Called by: `.AddMeter()` (same file)
- **`.AddCompBar()`** — L25004 — `public string AddCompBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds comp bar.
  Called by: `.AddMeter()` (same file)
- **`.AddPWRBar()`** — L25177 — `public string AddPWRBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  return cb.ID; }
  Called by: `.AddMeter()` (same file), `.AddREVPWRBar()` (same file)
- **`.AddREVPWRBar()`** — L25181 — `public string AddREVPWRBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds revpwrbar.
  Called by: `.AddMeter()` (same file)
- **`.AddSWRBar()`** — L25313 — `public string AddSWRBar(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds swrbar.
  Called by: `.AddMeter()` (same file)
- **`.AddBandButtons()`** — L25376 — `public string AddBandButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds band buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddDiscordButtons()`** — L25412 — `public string AddDiscordButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds discord buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddModeButtons()`** — L25448 — `public string AddModeButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds mode buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddVoiceRecordPlay()`** — L25484 — `public string AddVoiceRecordPlay(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds voice record play.
  Called by: `.AddMeter()` (same file)
- **`.AddWaveRecord()`** — L25520 — `public string AddWaveRecord(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds wave record.
  Called by: `.AddMeter()` (same file)
- **`.AddFilterButtons()`** — L25549 — `public string AddFilterButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds filter buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddHistory()`** — L25585 — `public string AddHistory(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds history.
  Called by: `.AddMeter()` (same file)
- **`.AddAntennaButtons()`** — L25614 — `public string AddAntennaButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds antenna buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddTunestepButtons()`** — L25650 — `public string AddTunestepButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds tunestep buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddOtherButtons()`** — L25686 — `public string AddOtherButtons(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds other buttons.
  Called by: `.AddMeter()` (same file)
- **`.AddFilterDisplay()`** — L25722 — `public string AddFilterDisplay(int nMSupdate, float fTop, out float fBottom, float fSize, clsItemGroup restoreIg = null)`
  Adds filter display.
  Called by: `.AddMeter()` (same file)
- **`.AddVFODisplay()`** — L25751 — `public string AddVFODisplay(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds vfodisplay.
  Called by: `.AddMeter()` (same file)
- **`.AddClock()`** — L25815 — `public string AddClock(int nMSupdate, float fTop, out float fBottom, clsItemGroup restoreIg = null)`
  Adds clock.
  Called by: `.AddMeter()` (same file)
- **`.GetBandGroupFromBand()`** — L25864 — `public BandGroups GetBandGroupFromBand(Band b)`
  Returns band group from band.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetBandPanel()`** — L25917 — `public void SetBandPanel(Console c, int rx, bool gen, bool hf, bool vhf)`
  Sets band panel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetWebImageIDsFrom4Char()`** — L25947 — `public (string, string) GetWebImageIDsFrom4Char(string fourchar)`
  Returns web image ids from4 char.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsWebImageBackgroundShown()`** — L25966 — `public bool IsWebImageBackgroundShown()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateFilterDetails()`** — L25983 — `public void UpdateFilterDetails(Filter newFilter, string name, int low, int high, bool vfoA, bool vfoB, int max_width, int max_size)`
  Updates filter details.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PAProfileChanged()`** — L25997 — `public void PAProfileChanged(string profile)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXProfileChanged()`** — L26258 — `public void TXProfileChanged(string profile)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateTXFilterDetails()`** — L26272 — `public void UpdateTXFilterDetails(int low, int high)`
  Updates txfilter details.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.InitFilterButtons()`** — L26286 — `public void InitFilterButtons()`
  Inits filter buttons.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeChanged()`** — L26314 — `public void ModeChanged(DSPMode oldMode, DSPMode newMode)`
  mi.FilterNameChanged(f, new_name); } } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuneStepIndexChanged()`** — L26328 — `public void TuneStepIndexChanged(int old_index, int new_index)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandChanged()`** — L26342 — `public void BandChanged(Band oldBand, Band newBand)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AntennasChanged()`** — L26356 — `public void AntennasChanged(Band rx1_band, Band tx_band, double vfoa_freq, double tx_freq, int rxtx_swap = 0)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandPanelsChanged()`** — L26370 — `public void BandPanelsChanged(bool gen, bool hf, bool vhf)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VHFDetailsChanged()`** — L26384 — `public void VHFDetailsChanged(int idx, bool enabled, string text)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroOut()`** — L26398 — `public void ZeroOut(bool bRxReadings, bool bTxReadings)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RebuildLedReadings()`** — L26421 — `public void RebuildLedReadings()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RebuildLedConditions()`** — L26436 — `public void RebuildLedConditions()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyDown()`** — L26452 — `public void KeyDown(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.KeyUp()`** — L26463 — `public void KeyUp(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TryParse()`** — L26474 — `public bool TryParse(string str)`
  Called by: `.FilterItemFrequencies()` (same file), `.getIntPart()` (same file), `.TryParse()` (same file), `.ContainerFromString()` (same file), `.onJsonWritten()` (same file), `.sanitiseJsonDataValue()` (same file) — and 4 more
- **`.numberOfMeterGroups()`** — L26517 — `private int numberOfMeterGroups()`
  Called by: `.addSMeterBar()` (same file), `.AddADCMaxMag()` (same file), `.AddSMeterBarText()` (same file), `.AddADCBar()` (same file), `.AddPBSNRBar()` (same file), `.AddAGCGainBar()` (same file) — and 37 more
- **`.removeMeterItem()`** — L26541 — `private void removeMeterItem(string sId, bool bRebuild = false)`
  Called by: `.RemoveMeterType()` (same file), `.RemoveAllMeterTypes()` (same file)
- **`.RemoveMeterType()`** — L26570 — `public void RemoveMeterType(MeterType mt, int order, bool bRebuild = false)`
  Removes meter type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetMeterItem()`** — L26613 — `public clsMeterItem GetMeterItem(MeterType mt, int order, clsMeterItem.MeterItemType mit)`
  Returns meter item.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MetersVoiceItemsUseThisKeycode()`** — L26649 — `public bool MetersVoiceItemsUseThisKeycode(Keys keycode)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AbortAllVoiceRecordRepeatPlaybacks()`** — L26671 — `public void AbortAllVoiceRecordRepeatPlaybacks()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MeterHasLockedVoiceRecords()`** — L26691 — `public bool MeterHasLockedVoiceRecords()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveAllMeterTypes()`** — L26713 — `public void RemoveAllMeterTypes(bool bRebuild = false)`
  Removes all meter types.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HasMeterType()`** — L26746 — `public bool HasMeterType(MeterType mt)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisableMeterType()`** — L26773 — `public void DisableMeterType(MeterType mt, bool bDisable)`
  Disables meter type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MeterGroupID()`** — L26814 — `public string MeterGroupID(MeterType mt = MeterType.NONE, int order = -1)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Find4Chars()`** — L26844 — `public void Find4Chars(ref Dictionary<string, string> fourchars, clsIGSettings igs)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ApplySettingsForMeterGroup()`** — L26908 — `public void ApplySettingsForMeterGroup(MeterType mt, clsIGSettings igs, List<string> webimages = null, int order = -1, bool from_setup_form = false, bool ignore_led_condition = fal`
  Applys settings for meter group.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetSettingsForMeterGroup()`** — L28355 — `public clsIGSettings GetSettingsForMeterGroup(MeterType mt, int order = -1)`
  Returns settings for meter group.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetOrderForMeterType()`** — L29384 — `public List<int> GetOrderForMeterType(MeterType mt)`
  Returns order for meter type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetOrderForMeterType()`** — L29414 — `public void SetOrderForMeterType(MeterType mt, int nOrder, bool bRebuild, bool bUp, int order = -1)`
  Sets order for meter type.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateAlways()`** — L29500 — `public bool UpdateAlways()`
  Updates always.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.itemFromID()`** — L29531 — `internal clsMeterItem itemFromID(string sId)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getBounds()`** — L29541 — `internal System.Drawing.RectangleF getBounds(string sId)`
  Returns bounds.
  Called by: `.getFadeCover()` (same file), `.ApplySettingsForMeterGroup()` (same file)
- **`.LedIndicatorFromFourChar()`** — L29567 — `internal clsLed LedIndicatorFromFourChar(string fourchar)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VoiceRecordPlayFromFourChar()`** — L29585 — `internal clsVoiceRecordPlay VoiceRecordPlayFromFourChar(string fourchar)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.itemsFromID()`** — L29603 — `internal Dictionary<string, clsMeterItem> itemsFromID(string sId, bool bIncludeTheParent = true, bool bOnlyChildren = false)`
  Called by: `.removeMeterItem()` (same file), `.GetMeterItem()` (same file), `.DisableMeterType()` (same file), `.Find4Chars()` (same file), `.ApplySettingsForMeterGroup()` (same file), `.GetSettingsForMeterGroup()` (same file) — and 2 more
- **`.hasReading()`** — L29620 — `private bool hasReading(Reading reading)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetBottom()`** — L29636 — `public float GetBottom()`
  Returns bottom.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateItems()`** — L29672 — `public void UpdateItems()`
  if (meterItems.Count == 1) { return (clsItemGroup)meterItems.First().Value; } return null; } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getMeterGroups()`** — L29729 — `internal Dictionary<string, clsItemGroup> getMeterGroups()`
  Returns meter groups.
  Called by: `.Rebuild()` (same file)
- **`.UpdateIntervals()`** — L29756 — `public void UpdateIntervals()`
  Updates intervals.
  Called by: `.ApplySettingsForMeterGroup()` (same file), `.Rebuild()` (same file)
- **`.Rebuild()`** — L29761 — `public void Rebuild()`
  Called by: `.removeMeterItem()` (same file), `.RemoveMeterType()` (same file), `.RemoveAllMeterTypes()` (same file), `.ApplySettingsForMeterGroup()` (same file), `.SetOrderForMeterType()` (same file)
- **`.setupSortedZOrder()`** — L29795 — `private void setupSortedZOrder()`
  Called by: `.Rebuild()` (same file)
- **`.MouseUp()`** — L29814 — `internal void MouseUp(System.Windows.Forms.MouseEventArgs e, clsMeter m, clsClickBox cb)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateVfoABandText()`** — L29976 — `private void updateVfoABandText(object _)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updateVfoBBandText()`** — L29996 — `private void updateVfoBBandText(object _)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UpdateBandText()`** — L30240 — `public void UpdateBandText(bool is_vfoA)`
  Updates band text.
  Called by: `.updateVfoABandText()` (same file), `.updateVfoBBandText()` (same file)
- **`.ContainerEnabled()`** — L30625 — `public void ContainerEnabled(string id, bool enabled)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ContainerHiddenByMacro()`** — L30636 — `public void ContainerHiddenByMacro(string id, bool hidden)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickestUpdateInterval()`** — L30913 — `public int QuickestUpdateInterval(bool mox, bool console_only_refresh)`
  Called by: `.QuickestUpdateInterval()` (same file), `.UpdateIntervals()` (same file)
- **`.addUpdateReading()`** — L31017 — `private void addUpdateReading(ref Dictionary<Reading, object> all_readings, Reading reading, object value)`
  Called by: `.Update()` (same file)
- **`.Update()`** — L31024 — `public void Update(ref List<Reading> readingsUsed)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.clearB()`** — L31134 — `private string clearB(string b)`
  Called by: `.Update()` (same file)
- **`.DelayForUpdate()`** — L31139 — `public int DelayForUpdate()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.incrementDisplayGroup()`** — L31165 — `private void incrementDisplayGroup()`
  Called by: `.MouseUp()` (same file)
- **`.decrementDisplayGroup()`** — L31170 — `private void decrementDisplayGroup()`
  Called by: `.MouseUp()` (same file)
- **`.incrementMeterItem()`** — L31175 — `private void incrementMeterItem(clsClickBox cb)`
  Called by: `.MouseUp()` (same file)
- **`.decrementMeterItem()`** — L31181 — `private void decrementMeterItem(clsClickBox cb)`
  Called by: `.MouseUp()` (same file)
- **`.AddDisplayGroup()`** — L31201 — `public void AddDisplayGroup(string sName)`
  Adds display group.
  Called by: `.AddAnanMM()` (same file)
- **`.RemoveDisplayGroup()`** — L31208 — `public void RemoveDisplayGroup(string sName)`
  Removes display group.
  Called by: `.RemoveMeterType()` (same file)
- **`.ToString()`** — L31225 — `public override string ToString()`
  Returns the string representation.
  Called by: `.GetReading()` (same file), `.ContainerFromString()` (same file), `.StoreSettings2()` (same file), `.ToString()` (same file), `.handleClicked()` (same file), `.handlePlay()` (same file) — and 29 more

#### `clsReading` (type, L31244)

_No extracted members._

#### `clsReadings` (type, L31250)

- **`.GetReading()`** — L31263 — `public float GetReading(Reading rt, bool useReading = false)`
  Returns reading.
  Called by: `.getReading()` (same file), `.parseText()` (same file), `.add_readings()` (same file)
- **`.SetReading()`** — L31269 — `public void SetReading(Reading rt, float value)`
  Sets reading.
  Called by: `.setReading()` (same file), `.setReadingForced()` (same file)
- **`.RequiresUpdate()`** — L31275 — `public bool RequiresUpdate(Reading rt)`
  Called by: `.setReading()` (same file), `.RequiresUpdate()` (same file)
- **`.UseReading()`** — L31280 — `public void UseReading(Reading rt)`
  Called by: `.GetReading()` (same file), `.UpdateMeters()` (same file)

#### `DXRenderer` (type, L31290)

- **`.HandleTouchDown()`** — L31454 — `private void HandleTouchDown(int x, int y)`
  Handles touch down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleTouchUp()`** — L31460 — `private void HandleTouchUp(int x, int y)`
  Handles touch up.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.HandleTouchMove()`** — L31466 — `private void HandleTouchMove(int x, int y)`
  Handles touch move.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RunDisplay()`** — L31483 — `public void RunDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveAnySkinImages()`** — L31530 — `internal void RemoveAnySkinImages()`
  Removes any skin images.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.convertImageToDX()`** — L31545 — `private void convertImageToDX(string sID)`
  Called by: `.drawWaveRecordIcon()` (same file), `.renderButtonBox()` (same file), `.renderVfoDisplay()` (same file), `.renderImage()` (same file)
- **`.getMaxSamples()`** — L31573 — `private int getMaxSamples()`
  Returns max samples.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dxInit()`** — L31586 — `private void dxInit(DriverType driverType = DriverType.Hardware, Display.AdaptorInfo adaptorInfo = null)`
  Called by: `.RunDisplay()` (same file)
- **`.setupFilterWaterfallBitmap()`** — L31809 — `private void setupFilterWaterfallBitmap()`
  Called by: `.renderFilterDisplay()` (same file)
- **`.resetDX2DModeDescription()`** — L31845 — `private void resetDX2DModeDescription(int fps)`
  Called by: `.dxRender()` (same file)
- **`.ShutdownDX()`** — L31870 — `public void ShutdownDX(bool bFromRenderThread = false)`
  Called by: `.Shutdown()` (same file), `.dxInit()` (same file), `.dxRender()` (same file)
- **`.RemoveAllDXImages()`** — L31984 — `internal void RemoveAllDXImages()`
  Removes all dximages.
  Called by: `.ShutdownDX()` (same file)
- **`.RemoveDXImage()`** — L32004 — `internal void RemoveDXImage(string sKey)`
  Removes dximage.
  Called by: `.RemoveAnySkinImages()` (same file)
- **`.dxRender()`** — L32036 — `private void dxRender()`
  { double late = _dElapsedFrameStart - (_fLastTime + 1000); if (late > 2000 || late < 0) late = 0; // ignore if too late _nFps = _nFrameCount; _nFrameCount = 0; _fLastTime = _dElapsedFrameStart - late; } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.setupAliasing()`** — L32193 — `private void setupAliasing()`
  Called by: `.dxInit()` (same file), `.resizeDX()` (same file)
- **`.releaseDXResources()`** — L32245 — `private void releaseDXResources()`
  Called by: `.ShutdownDX()` (same file)
- **`.buildDXFonts()`** — L32252 — `private void buildDXFonts()`
  Called by: `.dxInit()` (same file)
- **`.releaseDXFonts()`** — L32258 — `private void releaseDXFonts()`
  Called by: `.ShutdownDX()` (same file)
- **`.clearAllDynamicBrushes()`** — L32274 — `private void clearAllDynamicBrushes()`
  Called by: `.releaseDXResources()` (same file)
- **`.clearAllDynamicTextFormats()`** — L32288 — `private void clearAllDynamicTextFormats()`
  Called by: `.releaseDXResources()` (same file)
- **`.getDXBrushForColour()`** — L32305 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private SharpDX.Direct2D1.Brush getDXBrushForColour(System.Drawing.Color c, int replaceAlpha = -1)`
  Returns dxbrush for colour.
  Called by: `.dxRender()` (same file), `.renderNeedleScale()` (same file), `.renderScale()` (same file), `.generalScale()` (same file), `.renderGroup()` (same file), `.renderLed()` (same file) — and 24 more
- **`.convertColour()`** — L32333 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private SharpDX.Color4 convertColour(System.Drawing.Color c)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.getDXTextFormatForFont()`** — L32338 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private SharpDX.DirectWrite.TextFormat getDXTextFormatForFont(string sFontFamily, float emSize, FontStyle style, bool bAlignRight`
  Returns dxtext format for font.
  Called by: `.renderNeedleScale()` (same file), `.renderScale()` (same file), `.generalScale()` (same file), `.renderGroup()` (same file), `.renderHBar()` (same file), `.plotText()` (same file) — and 1 more
- **`.resizeDX()`** — L32374 — `private bool resizeDX(out string error)`
  Called by: `.dxRender()` (same file)
- **`.target_Resize()`** — L32429 — `private void target_Resize(object sender, System.EventArgs e)`
  WinForms event handler: runs when `target` is resized.
  Called by: WinForms event wiring at runtime (no static call sites).
- **`.target_VisibleChanged()`** — L32436 — `private void target_VisibleChanged(object sender, System.EventArgs e)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseEnter()`** — L32442 — `private void OnMouseEnter(object sender, System.EventArgs e)`
  Handles/raises the mouse enter event.
  Called by: `.HandleTouchDown()` (same file)
- **`.OnMouseLeave()`** — L32462 — `private void OnMouseLeave(object sender, System.EventArgs e)`
  Handles/raises the mouse leave event.
  Called by: `.HandleTouchUp()` (same file)
- **`.OnMouseCaptureChanged()`** — L32483 — `private void OnMouseCaptureChanged(object sender, System.EventArgs e)`
  Handles/raises the mouse capture changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseMove()`** — L32504 — `private void OnMouseMove(object sender, System.Windows.Forms.MouseEventArgs e)`
  Handles/raises the mouse move event.
  Called by: `.HandleTouchMove()` (same file)
- **`.OnMouseWheel()`** — L32554 — `private void OnMouseWheel(object sender, System.Windows.Forms.MouseEventArgs e)`
  Handles/raises the mouse wheel event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseClick()`** — L32604 — `private void OnMouseClick(object sender, MouseEventArgs e)`
  Handles/raises the mouse click event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMouseDown()`** — L32658 — `private void OnMouseDown(object sender, System.Windows.Forms.MouseEventArgs e)`
  Handles/raises the mouse down event.
  Called by: `.HandleTouchDown()` (same file)
- **`.OnMouseUp()`** — L32712 — `private void OnMouseUp(object sender, System.Windows.Forms.MouseEventArgs e)`
  Handles/raises the mouse up event.
  Called by: `.HandleTouchUp()` (same file)
- **`.OnDragEnter()`** — L32787 — `private void OnDragEnter(object sender, DragEventArgs e)`
  Handles/raises the drag enter event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDragOver()`** — L32800 — `private void OnDragOver(object sender, DragEventArgs e)`
  Handles/raises the drag over event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnDragDrop()`** — L32813 — `private void OnDragDrop(object sender, DragEventArgs e)`
  Handles/raises the drag drop event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.tryGetWaveRecordDropTarget()`** — L32824 — `private clsWaveRecord tryGetWaveRecordDropTarget(object sender, System.Drawing.Point screenPoint)`
  Called by: `.OnDragEnter()` (same file), `.OnDragOver()` (same file), `.OnDragDrop()` (same file)
- **`.drawMeters()`** — L32869 — `private int drawMeters(out int height)`
  Called by: `.dxRender()` (same file)
- **`.measureString()`** — L33023 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private SizeF measureString(string sText, string sFontFamily, FontStyle style, float emSize, bool ignore_caching = false)`
  Called by: `.renderNeedleScale()` (same file), `.renderScale()` (same file), `.generalScale()` (same file), `.renderGroup()` (same file), `.renderTextOverlay()` (same file), `.renderDialDisplay()` (same file) — and 4 more
- **`.fade()`** — L33153 — `private int fade(clsMeterItem mi, clsMeter m)`
  return size; } [2.10.1.0] MW0LGE
  Called by: `.renderFadeCover()` (same file)
- **`.renderNeedleScale()`** — L33179 — `private void renderNeedleScale(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.tidyPower()`** — L33385 — `private string tidyPower(float fPower)`
  Called by: `.renderNeedleScale()` (same file), `.renderScale()` (same file)
- **`.renderScale()`** — L33393 — `private void renderScale(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.generalScale()`** — L33872 — `private void generalScale(float x,float y,float w,float h,clsScaleItem scale, int lowLongTicks, int highLongTicks, int lowStartNumber, int highEndNumber, int lowIncrement, int high`
  Called by: `.renderScale()` (same file)
- **`.renderGroup()`** — L33959 — `private void renderGroup(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.slits()`** — L33977 — `private void slits(Vector2 centre, float radiusX, float radiusY, float w, float h, SharpDX.Direct2D1.Brush closedSectionBrush)`
  Called by: `.renderEye()` (same file)
- **`.renderLed()`** — L34010 — `private void renderLed(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m, bool draw_led)`
  Called by: `.drawMeters()` (same file)
- **`.renderTextOverlay()`** — L34135 — `private bool renderTextOverlay(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m, bool render_text)`
  Called by: `.drawMeters()` (same file)
- **`.buildWaterfall()`** — L34289 — `private (byte[], byte[], bool) buildWaterfall(clsFilterItem filter)`
  Called by: `.renderFilterDisplay()` (same file)
- **`.renderDialDisplay()`** — L35284 — `private void renderDialDisplay(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderFilterDisplay()`** — L35466 — `private void renderFilterDisplay(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.hzToPixels()`** — L36272 — `private float hzToPixels(double hz, float pixel_span, int hz_span)`
  Called by: `.renderFilterDisplay()` (same file)
- **`.pixelsToHz()`** — L36277 — `private double pixelsToHz(float pixels, float pixel_span, int hz_span)`
  Called by: `.renderFilterDisplay()` (same file)
- **`.isMouseNearLine()`** — L36282 — `public bool isMouseNearLine(RawVector2 point1, RawVector2 point2, PointF mousePosition, float proximityThreshold = 3.0f)`
  Called by: `.renderFilterDisplay()` (same file)
- **`.pointToSegmentDistance()`** — L36288 — `private float pointToSegmentDistance(RawVector2 lineStart, RawVector2 lineEnd, PointF point)`
  Called by: `.renderFilterDisplay()` (same file), `.isMouseNearLine()` (same file)
- **`.distanceBetweenPoints()`** — L36310 — `private float distanceBetweenPoints(RawVector2 p1, PointF p2)`
  Called by: `.renderFilterDisplay()` (same file), `.pointToSegmentDistance()` (same file)
- **`.renderHistory()`** — L36315 — `private void renderHistory(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderSpacer()`** — L36544 — `private void renderSpacer(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderWebImage()`** — L36578 — `private void renderWebImage(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.convertDegreesToCardinal()`** — L36671 — `private string convertDegreesToCardinal(float degrees)`
  Called by: `.renderRotator()` (same file)
- **`.renderRotator()`** — L36704 — `private void renderRotator(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.calculateDistance()`** — L37227 — `private float calculateDistance(PointF point1, PointF point2)`
  Called by: `.renderRotator()` (same file)
- **`.renderEye()`** — L37233 — `private void renderEye(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderText()`** — L37357 — `private void renderText(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderHBarMarkersOnly()`** — L37388 — `private void renderHBarMarkersOnly(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderHBar()`** — L37424 — `private clsMeterItem renderHBar(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderSolidColour()`** — L37813 — `private void renderSolidColour(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  SharpDX.RectangleF txtrect = new SharpDX.RectangleF(x, y + (h * 0.2f), w, h); _renderTarget.DrawText(sText, getDXTextFormatForFont(cbi.FontFamily, newSize, cbi.FntStyle), txtrect, getDXBrushForColour(cbi.FontColour)); } }
  Called by: `.drawMeters()` (same file)
- **`.renderFadeCover()`** — L37826 — `private void renderFadeCover(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.getParts()`** — L37847 — `private void getParts(double vfoFreq, out string MHz, out string kHz, out string hz)`
  Returns parts.
  Called by: `.renderVfoDisplay()` (same file)
- **`.plotText()`** — L37857 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private (float, float) plotText(string sText, float x, float y, float containerWidth, float fTextSize, System.Drawing.Color c, in`
  Called by: `.renderTextOverlay()` (same file), `.renderDialDisplay()` (same file), `.renderFilterDisplay()` (same file), `.renderHistory()` (same file), `.renderWebImage()` (same file), `.renderRotator()` (same file) — and 10 more
- **`.highlightBox()`** — L38052 — `private void highlightBox(float x, float y, float w, float h, SharpDX.RectangleF rect, clsVfoDisplay vfo, int bx, int by, float gap, clsMeter m, float shift)`
  return (szTextSize.Width, szTextSize.Height); }
  Called by: `.drawTuneStep()` (same file), `.drawBand()` (same file), `.drawMode()` (same file), `.drawFilter()` (same file)
- **`.drawTuneStep()`** — L38063 — `private clsVfoDisplay.buttonState drawTuneStep(float x, float y, float w, float h, SharpDX.RectangleF rect, clsVfoDisplay vfo, clsMeter m, float shift, float x_multy)`
  Called by: `.renderVfoDisplay()` (same file)
- **`.drawBand()`** — L38170 — `private clsVfoDisplay.buttonState drawBand(float x, float y, float w, float h, SharpDX.RectangleF rect, clsVfoDisplay vfo, clsMeter m, float shift, float x_multy)`
  Called by: `.renderVfoDisplay()` (same file)
- **`.drawMode()`** — L38343 — `private clsVfoDisplay.buttonState drawMode(float x, float y, float w, float h, SharpDX.RectangleF rect, clsVfoDisplay vfo, clsMeter m, float shift, float x_multy)`
  Called by: `.renderVfoDisplay()` (same file)
- **`.drawFilter()`** — L38438 — `private clsVfoDisplay.buttonState drawFilter(float x, float y, float w, float h, SharpDX.RectangleF rect, clsVfoDisplay vfo, clsMeter m, float shift, float x_multy)`
  Called by: `.renderVfoDisplay()` (same file)
- **`.shrinkRectangle()`** — L38552 — `private void shrinkRectangle(SharpDX.RectangleF original, float ratio, ref SharpDX.RectangleF shrunk, float absolute = 0f)`
  Called by: `.renderButtonBox()` (same file)
- **`.drawRoundedRectangle()`** — L38570 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private void drawRoundedRectangle(RoundedRectangle rr, SharpDX.Direct2D1.Brush b, float stroke, bool centred = false)`
  Called by: `.renderWaveRecord()` (same file), `.drawWaveRecordButton()` (same file), `.renderButtonBox()` (same file)
- **`.fillRoundedRectangle()`** — L38579 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private void fillRoundedRectangle(RoundedRectangle rr, SharpDX.Direct2D1.Brush b, bool centred = false)`
  Called by: `.renderWaveRecord()` (same file), `.drawWaveRecordButton()` (same file), `.renderButtonBox()` (same file)
- **`.drawSafeLine()`** — L38587 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private void drawSafeLine(RawVector2 start, RawVector2 end, SharpDX.Direct2D1.Brush b, float width)`
  Called by: `.renderButtonBox()` (same file)
- **`.dimColour()`** — L38600 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private System.Drawing.Color dimColour(System.Drawing.Color colour, bool dim, float amount = 0.35f)`
  Called by: `.renderButtonBox()` (same file)
- **`.renderWaveRecord()`** — L38611 — `private void renderWaveRecord(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.drawWaveRecordButton()`** — L38948 — `private void drawWaveRecordButton(SharpDX.RectangleF buttonRect, bool hovered, System.Drawing.Color fillColour, System.Drawing.Color borderColour, System.Drawing.Color hoverColour,`
  Called by: `.renderWaveRecord()` (same file)
- **`.drawWaveRecordIcon()`** — L38962 — `private void drawWaveRecordIcon(string icon, SharpDX.RectangleF buttonRect, System.Drawing.Color iconColour, int fade)`
  Called by: `.renderWaveRecord()` (same file)
- **`.clipRect()`** — L38990 — `private SharpDX.RectangleF clipRect(SharpDX.RectangleF value, SharpDX.RectangleF clip)`
  Called by: `.renderWaveRecord()` (same file)
- **`.rectEmpty()`** — L39000 — `private bool rectEmpty(SharpDX.RectangleF value)`
  Called by: `.renderWaveRecord()` (same file)
- **`.samePath()`** — L39004 — `private bool samePath(string left, string right)`
  Called by: `.renderWaveRecord()` (same file)
- **`.renderButtonBox()`** — L39009 — `private void renderButtonBox(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.adjustTextColourForContrast()`** — L39820 — `private System.Drawing.Color adjustTextColourForContrast(System.Drawing.Color textColor, System.Drawing.Color backgroundColor)`
  Called by: `.renderButtonBox()` (same file)
- **`.calculateContrastRatio()`** — L39849 — `private double calculateContrastRatio(double luminance1, double luminance2)`
  Called by: `.adjustTextColourForContrast()` (same file)
- **`.renderVfoDisplay()`** — L39857 — `private void renderVfoDisplay(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderClock()`** — L40474 — `private void renderClock(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderSignalTextDisplay()`** — L40538 — `private void renderSignalTextDisplay(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderImage()`** — L40629 — `private void renderImage(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.renderNeedle()`** — L40740 — `private void renderNeedle(SharpDX.RectangleF rect, clsMeterItem mi, clsMeter m)`
  Called by: `.drawMeters()` (same file)
- **`.degToRad()`** — L41033 — `private double degToRad(float deg)`
  Called by: `.renderNeedleScale()` (same file), `.renderEye()` (same file), `.renderNeedle()` (same file)
- **`.radToDeg()`** — L41037 — `private double radToDeg(float rad)`
  Called by: `.renderNeedle()` (same file)
- **`.getPerc()`** — L41041 — `private void getPerc(clsMeterItem mi, float rawValue, out float percX, out float percY, out PointF min, out PointF max)`
  Returns perc.
  Called by: `.renderNeedleScale()` (same file), `.renderEye()` (same file), `.renderHBarMarkersOnly()` (same file), `.renderHBar()` (same file), `.renderNeedle()` (same file)
- **`.bitmapFromSystemBitmap()`** — L41249 — `private SharpDX.Direct2D1.Bitmap bitmapFromSystemBitmap(RenderTarget rt, System.Drawing.Bitmap bitmap, string sId)`
  pc = new clsMeterItem.clsPercCache() { _value = value, _percX = percX, _percY = percY, _min = min, _max = max }; mi.AddPerc(pc); } }
  Called by: `.convertImageToDX()` (same file), `.renderWebImage()` (same file)

#### `MultiMeterIO` (type, L41321)

- **`.StartListeningUDP()`** — L42777 — `public static bool StartListeningUDP(clsMMIO mmio)`
  Starts listening udp.
  Called by: `.StartConnection()` (same file)
- **`.StartListeningTCPIP()`** — L42805 — `public static bool StartListeningTCPIP(clsMMIO mmio)`
  Starts listening tcpip.
  Called by: `.StartConnection()` (same file)
- **`.StartTcpClient()`** — L42831 — `public static bool StartTcpClient(clsMMIO mmio)`
  Starts tcp client.
  Called by: `.StartConnection()` (same file)
- **`.StartSerialPort()`** — L42859 — `public static bool StartSerialPort(clsMMIO mmio)`
  Starts serial port.
  Called by: `.StartConnection()` (same file)
- **`.StopConnection()`** — L42887 — `public static void StopConnection(Guid guid)`
  Stops connection.
  Called by: `.StopConnection()` (same file), `.RestoreSaveData()` (same file), `.AddMMIO()` (same file), `.RemoveMMIO()` (same file), `.chkMMIO_network_enabled_CheckedChanged()` (`Console/setup.cs`)
- **`.StopConnections()`** — L42912 — `public static void StopConnections()`
  Stops connections.
  Called by: `.Shutdown()` (same file), `.RestoreSaveData2()` (same file), `.RestoreSaveData()` (same file)
- **`.AlreadyConfigured()`** — L42938 — `public static bool AlreadyConfigured(string ip, int port, MMIOType type)`
  Called by: `.addEditConnector()` (`Console/setup.cs`)
- **`.GetSaveData()`** — L42951 — `public static string GetSaveData()`
  Returns save data.
  Called by: `.SaveOptions()` (`Console/setup.cs`)
- **`.RestoreSaveData2()`** — L43006 — `public static bool RestoreSaveData2(string data)`
  Restores save data2.
  Called by: `.getOptions()` (`Console/setup.cs`)
- **`.RestoreSaveData()`** — L43036 — `public static bool RestoreSaveData(string data)`
  Restores save data.
  Called by: `.getOptions()` (`Console/setup.cs`)
- **`.AddMMIO()`** — L43163 — `public static bool AddMMIO(clsMMIO mmio)`
  Adds mmio.
  Called by: `.showSerialPortPicker()` (`Console/setup.cs`), `.addEditConnector()` (`Console/setup.cs`)
- **`.RemoveMMIO()`** — L43176 — `public static bool RemoveMMIO(Guid guid)`
  Removes mmio.
  Called by: `.showSerialPortPicker()` (`Console/setup.cs`), `.addEditConnector()` (`Console/setup.cs`), `.btnMMIO_network_delete_Click()` (`Console/setup.cs`)
- **`.SendDataMMIO()`** — L43186 — `public static void SendDataMMIO(Guid guid, string data)`
  Sends data mmio.
  Called by: `.handleMacroButtonPress()` (same file), `.Update()` (same file)
- **`.GuidfromFourChar()`** — L43194 — `public static Guid GuidfromFourChar(string fourChar)`
  Called by: `.handleMacroButtonPress()` (same file), `.updateMeterType()` (`Console/setup.cs`)
- **`.MultiMeterIO_ListenerRunning()`** — L43203 — `private static void MultiMeterIO_ListenerRunning(Guid guid, MMIOType type, bool running)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiMeterIO_TransmittedData()`** — L43209 — `private static void MultiMeterIO_TransmittedData(Guid guid)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsValidXml()`** — L43212 — `[DebuggerHidden] public static bool IsValidXml(string xmlString)`
  Called by: `.MultiMeterIO_ReceivedDataString()` (same file)
- **`.MultiMeterIO_ReceivedDataString()`** — L43238 — `private static void MultiMeterIO_ReceivedDataString(Guid guid, string dataString)`
  [DebuggerHidden]
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.parseJsonToken()`** — L43301 — `public static void parseJsonToken(JToken token, string currentPath, Dictionary<string, string> keyValuePairs)`
  Called by: `.MultiMeterIO_ReceivedDataString()` (same file)
- **`.parseXMLElement()`** — L43336 — `private static void parseXMLElement(XElement element, string currentPath, Dictionary<string, string> keyValuePairs)`
  Called by: `.MultiMeterIO_ReceivedDataString()` (same file)

#### `MMIODirection` (type, L41324)

_No extracted members._

#### `MMIOFormat` (type, L41331)

_No extracted members._

#### `MMIOType` (type, L41339)

_No extracted members._

#### `MMIOTerminator` (type, L41348)

_No extracted members._

#### `clsMMIO` (type, L41358)

- **`.OnDeserialized()`** — L41398 — `[OnDeserialized] private void OnDeserialized(StreamingContext context)`
  Handles/raises the deserialized event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.init()`** — L41405 — `private void init()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.refreshUdpEndpoint()`** — L41538 — `private void refreshUdpEndpoint()`
  Called by: `.StartConnection()` (same file)
- **`.EnqueueOutbound()`** — L41606 — `public void EnqueueOutbound(string data)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DequeueOutbound()`** — L41616 — `public string DequeueOutbound()`
  Called by: `.listen()` (same file), `.Connect()` (same file)
- **`.StartConnection()`** — L41644 — `public bool StartConnection()`
  Starts connection.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StopConnection()`** — L41665 — `public void StopConnection()`
  Stops connection.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetVariable()`** — L41669 — `public bool SetVariable(string key, object value)`
  Sets variable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetVariable()`** — L41683 — `public object GetVariable(string key, string precision_format = "")`
  Returns variable.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DetermineType()`** — L41707 — `public Type DetermineType(string value)`
  Called by: `.GetVariable()` (same file)
- **`.ConvertToType()`** — L41757 — `public object ConvertToType(string value, Type type)`
  [DebuggerHidden]
  Called by: `.GetVariable()` (same file)
- **`.Variables()`** — L41791 — `public ConcurrentDictionary<string, object> Variables()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VariableValueType()`** — L41795 — `public string VariableValueType(object obj, string float_precision = "")`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RemoveVariable()`** — L41826 — `public void RemoveVariable(string key)`
  IFormatter formatter = new BinaryFormatter(); using (Stream stream = new MemoryStream()) { formatter.Serialize(stream, obj); stream.Seek(0, SeekOrigin.Begin); return (T)formatter.Deserialize(stream); } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `TcpListener` (type, L41831)

- **`.Start()`** — L41855 — `public void Start()`
  Called by: `.listen()` (same file), `.StartServer()` (`Console/CAT/TCPIPcatServer.cs`), `.StartServer()` (`Console/TCIServer.cs`)
- **`.Stop()`** — L41864 — `public void Stop()`
  Called by: `.listen()` (same file), `.StopConnection()` (same file), `.StopServer()` (`Console/CAT/TCPIPcatServer.cs`), `.StopServer()` (`Console/TCIServer.cs`)
- **`.listen()`** — L41889 — `private void listen()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `TcpClientHandler` (type, L42092)

- **`.Start()`** — L42116 — `public void Start()`
  Called by: `.StartTcpClient()` (same file)
- **`.Stop()`** — L42125 — `public void Stop()`
  Called by: `.StopConnection()` (same file)
- **`.Connect()`** — L42148 — `private void Connect()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `UdpListener` (type, L42335)

- **`.Start()`** — L42358 — `public void Start()`
  Called by: `.StartListeningUDP()` (same file), `.StartListeningTCPIP()` (same file)
- **`.Stop()`** — L42367 — `public void Stop()`
  Called by: `.StopConnection()` (same file)
- **`.listen()`** — L42375 — `private void listen()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `SerialPortHandler` (type, L42521)

- **`.Start()`** — L42552 — `public void Start()`
  Called by: `.StartSerialPort()` (same file)
- **`.Stop()`** — L42561 — `public void Stop()`
  Called by: `.StopConnection()` (same file), `.StopConnections()` (same file)
- **`.Connect()`** — L42572 — `private void Connect()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetAvailableComPorts()`** — L42738 — `public static List<string> GetAvailableComPorts()`
  Returns available com ports.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `MiniSpec` (type, L43359)

- **`.Init()`** — L43417 — `public static void Init(Console console)`
  Called by: `.InitConsole()` (`Console/console.cs`)
- **`.UpdateRXFilterCharacteristics()`** — L43469 — `public static void UpdateRXFilterCharacteristics(DSPMode mode, double[] segments, int index_low, int index_upper, double corner_freq)`
  Updates rxfilter characteristics.
  Called by: `.BuildFilterCharacteristics()` (`Console/console.cs`)
- **`.UpdateTXFilterCharacteristics()`** — L43491 — `public static void UpdateTXFilterCharacteristics(double[] segments, int index_low, int index_upper, double corner_freq)`
  Updates txfilter characteristics.
  Called by: `.BuildFilterCharacteristics()` (`Console/console.cs`)
- **`.GetRXCharacteristic()`** — L43519 — `public static FilterCharacteristics GetRXCharacteristic(DSPMode mode)`
  Returns rxcharacteristic.
  Called by: `.renderFilterDisplay()` (same file)
- **`.GetTXCharacteristic()`** — L43527 — `public static FilterCharacteristics GetTXCharacteristic()`
  Returns txcharacteristic.
  Called by: `.renderFilterDisplay()` (same file)
- **`.UsingAFilter()`** — L43534 — `public static bool UsingAFilter(int id, bool sub_receiver = false)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.UsingFilter()`** — L43540 — `public static void UsingFilter(int id, bool sub_receiver = false)`
  Called by: `.Initialise()` (same file)
- **`.StopUsingFilter()`** — L43564 — `public static void StopUsingFilter(int id, bool sub_receiver = false)`
  Stops using filter.
  Called by: `.Removing()` (same file)
- **`.OnSpectrumDetailsChanged()`** — L43584 — `private static void OnSpectrumDetailsChanged(int rx)`
  Handles/raises the spectrum details changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnAVGChanged()`** — L43592 — `private static void OnAVGChanged(int rx, bool old_state, bool new_state)`
  Handles/raises the avgchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnFilterEdgesChanged()`** — L43600 — `private static void OnFilterEdgesChanged(int rx, Filter filter, Band band, int low, int high, string sName, int max_width, int max_shift)`
  Handles/raises the filter edges changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMinNotchWidth()`** — L43610 — `private static void OnMinNotchWidth(int rx, double width)`
  Handles/raises the min notch width event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnTNFChanged()`** — L43618 — `private static void OnTNFChanged(bool old_tnf, bool new_tnf)`
  Handles/raises the tnfchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMox()`** — L43659 — `private static void OnMox(int rx, bool oldMox, bool newMox)`
  Handles/raises the mox event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Add()`** — L43667 — `public static void Add(int rx, int id, bool sub_receiver = false)`
  Called by: `.addSMeterBar()` (same file), `.AddADCMaxMag()` (same file), `.AddADCBar()` (same file), `.AddPBSNRBar()` (same file), `.AddAGCGainBar()` (same file), `.AddAGCBar()` (same file) — and 17 more
- **`.GetMiniRX()`** — L43693 — `public static clsMiniSpec GetMiniRX(int id, bool sub_receiver = false)`
  Returns mini rx.
  Called by: `.OnTXFrequencyChanged()` (same file), `.OnVFOA()` (same file), `.OnVFOB()` (same file), `.OnVFOASub()` (same file), `.OnRX2EnabledPreChanged()` (same file), `.Update()` (same file)
- **`.shutdownRX()`** — L43714 — `private static void shutdownRX(int key)`
  Called by: `.ShutdownAllRX()` (same file)
- **`.ShutdownAllRX()`** — L43721 — `public static void ShutdownAllRX()`
  Called by: `.Shutdown()` (same file)
- **`.OnPowerChanged()`** — L43748 — `private static void OnPowerChanged(bool old_state, bool new_state)`
  Handles/raises the power changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L43760 — `private static void OnCTUNChanged(int rx, bool oldCTUN, bool newCTUN, Band band)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnModeChanged()`** — L43768 — `private static void OnModeChanged(int rx, DSPMode oldMode, DSPMode newMode, Band oldBand, Band newBand)`
  Handles/raises the mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCWPitchChanged()`** — L43776 — `private static void OnCWPitchChanged(int old_pitch, int new_pitch, bool show_cwzero)`
  Handles/raises the cwpitch changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnSpectrumSettingsChanged()`** — L43784 — `private static void OnSpectrumSettingsChanged(int rx)`
  Handles/raises the spectrum settings changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnHWSampleRateChanged()`** — L43792 — `private static void OnHWSampleRateChanged(int rx, int old_rate, int new_rate)`
  Handles/raises the hwsample rate changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCentreFrequency()`** — L43800 — `private static void OnCentreFrequency(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnNotchChanged()`** — L43811 — `private static void OnNotchChanged(int notch_index, double old_bw, double new_bw, bool active, double old_centre_freq, double new_centre_freq, bool added, bool removed)`
  Handles/raises the notch changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetNotches()`** — L43866 — `public static List<Notch> GetNotches(double centre_hz, int half_bandwidth)`
  Returns notches.
  Called by: `.renderFilterDisplay()` (same file), `.runDisplay()` (same file)
- **`.GetNotch()`** — L43878 — `public static Notch GetNotch(int notch_index)`
  Returns notch.
  Called by: `.MouseWheel()` (same file), `.AdjustNotch()` (same file)

#### `FilterCharacteristics` (type, L43367)

_No extracted members._

#### `Notch` (type, L43394)

_No extracted members._

#### `clsMiniSpec` (type, L43886)

- **`.setupSpecDetails()`** — L44118 — `private void setupSpecDetails()`
  Called by: `.UpdateSpecSettings()` (same file)
- **`.resetBuffers()`** — L44150 — `private void resetBuffers()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ClearData()`** — L44162 — `public void ClearData()`
  Clears data.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.runDisplay()`** — L44173 — `private void runDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.attenuateData()`** — L44251 — `private void attenuateData(int center_index, float attenuation, int span_in_pixels)`
  Called by: `.runDisplay()` (same file)
- **`.UpdateSpecSettings()`** — L44265 — `public void UpdateSpecSettings()`
  Updates spec settings.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Shutdown()`** — L44335 — `public void Shutdown()`
  Called by: `.Add()` (same file), `.shutdownRX()` (same file)
- **`.rateLimitSetPan()`** — L44385 — `public void rateLimitSetPan()`
  rate limit setPan() to 1/8 frame rate
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.updatePan()`** — L44406 — `private void updatePan(object _)`
  Called by: `.rateLimitSetPan()` (same file)
- **`.setPan()`** — L44429 — `private void setPan()`
  Sets pan.
  Called by: `.UpdateSpecSettings()` (same file), `.updatePan()` (same file)
- **`.zoom()`** — L44473 — `private void zoom()`
  Called by: `.UpdateSpecSettings()` (same file)

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/MeterManager.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
