# `Console/display.cs`

**Functional area:** [5. Spectrum, waterfall, and panadapter display](../../CODE_OUTLINE.md#5-spectrum-waterfall-and-panadapter-display)

**Role:** The spectrum/waterfall renderer — SharpDX (Direct2D/D3D11) drawing of panadapter, waterfall, band edges, notches, cursors, and TX filter overlays.

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (calls ×11)
  - `Console/setup.cs` (calls ×2)
- Uses (outgoing references to other files):
  - `Console/enums.cs` (references ×14)
  - `Console/radio.cs` (references ×1, calls ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/Channel.cs` (references ×1)
  - `Console/SpotManager2.cs` (calls ×1)
  - `Console/clsBandStackManager.cs` (references ×1)
  - `Console/hiperftimer.cs` (references ×1)
- Most-referenced symbols from other files: `.DX2Adaptors()` (×2), `.PurgeBuffers()` (×1), `.RenderDX2D()` (×1), `.SetPendingWaterfallPixelRef()` (×1), `.ResetWaterfallTimers()` (×1), `.ShutdownDX2D()` (×1), `.SetDX2BackgoundImage()` (×1), `.SetupDelegates()` (×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Functions

- **`letItSnow()`** — L11750 — `private static void letItSnow()`
  Called by: `.RenderDX2D()` (same file)
- **`plotSanta()`** — L11789 — `private static void plotSanta()`
  Called by: `letItSnow()` (same file)
- **`SetSantaGif()`** — L11852 — `public static void SetSantaGif(System.Drawing.Image image)`
  Sets santa gif.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`santaCleanUp()`** — L11885 — `private static void santaCleanUp()`
  Called by: `.ShutdownDX2D()` (same file), `SetSantaGif()` (same file)

### Types

#### `WaterfallTimePosition` (type, L82)

_No extracted members._

#### `WaterfallTimeMode` (type, L91)

_No extracted members._

#### `Display` (type, L99)

- **`.SetupDelegates()`** — L767 — `public static void SetupDelegates()`
  Setups delegates.
  Called by: `.addDelegates()` (`Console/console.cs`)
- **`.RemoveDelegates()`** — L786 — `public static void RemoveDelegates()`
  Removes delegates.
  Called by: `.removeDelegates()` (`Console/console.cs`)
- **`.OnMinRXNotchWidthChanged()`** — L800 — `private static void OnMinRXNotchWidthChanged(int rx, double width)`
  Handles/raises the min rxnotch width changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnMinTXNotchWidthChanged()`** — L805 — `private static void OnMinTXNotchWidthChanged(double width)`
  Handles/raises the min txnotch width changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCTUNChanged()`** — L811 — `private static void OnCTUNChanged(int rx, bool oldCTUN, bool newCTUN, Band band)`
  Handles/raises the ctunchanged event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPowerChangeHander()`** — L818 — `private static void OnPowerChangeHander(bool oldPower, bool newPower)`
  Handles/raises the power change hander event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnBandChangeHandler()`** — L825 — `private static void OnBandChangeHandler(int rx, Band oldBand, Band newBand)`
  Handles/raises the band change handler event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.processBlobsActivePeakDisplayDelay()`** — L847 — `private static void processBlobsActivePeakDisplayDelay()`
  Called by: `.RenderDX2D()` (same file)
- **`.delayBlobsActivePeakDisplay()`** — L859 — `private static void delayBlobsActivePeakDisplay(int rx, bool blobs)`
  Called by: `.ResetSpectrumPeaks()` (same file), `.ResetBlobMaximums()` (same file)
- **`.resetPeaksAndNoise()`** — L879 — `private static void resetPeaksAndNoise(int rx)`
  Called by: `.clearBuffers()` (same file)
- **`.OnAttenuatorDataChanged()`** — L893 — `private static void OnAttenuatorDataChanged(int rx, int oldAtt, int newAtt)`
  Handles/raises the attenuator data changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnPreampModeChanged()`** — L900 — `private static void OnPreampModeChanged(int rx, PreampMode oldMode, PreampMode newMode)`
  Handles/raises the preamp mode changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnCentreFrequencyChanged()`** — L907 — `private static void OnCentreFrequencyChanged(int rx, double oldFreq, double newFreq, Band band, double offset)`
  Handles/raises the centre frequency changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.localMox()`** — L2731 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static bool localMox(int rx)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.DrawWaterfallDX2D()` (same file), `.handleNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.drawSpots()` (same file)
- **`.getWaterfallRxIndex()`** — L2784 — `private static int getWaterfallRxIndex(int rx)`
  Returns waterfall rx index.
  Called by: `.setCurrentWaterfallBand()` (same file), `.getWaterfallContextBandValue()` (same file)
- **`.setCurrentWaterfallBand()`** — L2789 — `private static void setCurrentWaterfallBand(int rx, Band band)`
  Sets current waterfall band.
  Called by: `.OnBandChangeHandler()` (same file)
- **`.getWaterfallContextBandValue()`** — L2796 — `private static int getWaterfallContextBandValue(int rx, bool isTxContext)`
  Returns waterfall context band value.
  Called by: `.getWaterfallCachedPreviousMinOrFloor()` (same file), `.updateWaterfallAgcCache()` (same file)
- **`.getWaterfallCachedPreviousMinOrFloor()`** — L2810 — `private static float getWaterfallCachedPreviousMinOrFloor(int rx, bool isTxContext)`
  Returns waterfall cached previous min or floor.
  Called by: `.OnBandChangeHandler()` (same file)
- **`.updateWaterfallAgcCache()`** — L2822 — `private static void updateWaterfallAgcCache(int rx, bool isTxContext, float previousMin)`
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.clearBuffers()`** — L2835 — `private static void clearBuffers(int W, int rx)`
  Called by: `.initDisplayArrays()` (same file), `.PurgeBuffers()` (same file)
- **`.initDisplayArrays()`** — L2878 — `private static void initDisplayArrays(int W, int H)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.changeAlpha()`** — L2938 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static Color changeAlpha(Color c, int A)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.dBToPixel()`** — L2943 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static float dBToPixel(float dB, int H, bool tx = false)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.dBToRX2Pixel()`** — L2955 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static float dBToRX2Pixel(float dB, int H, bool tx = false)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.updateSharePointsArray()`** — L2980 — `private static void updateSharePointsArray(int nW)`
  Called by: `.DrawHistogramDX2D()` (same file)
- **`.ResetWaterfallBmp()`** — L3073 — `private static void ResetWaterfallBmp()`
  Resets waterfall bmp.
  Called by: `.initDX2D()` (same file)
- **`.ResetWaterfallBmp2()`** — L3131 — `private static void ResetWaterfallBmp2()`
  Resets waterfall bmp2.
  Called by: `.initDX2D()` (same file)
- **`.ShutdownDX2D()`** — L3230 — `public static void ShutdownDX2D()`
  Called by: `.initDX2D()` (same file), `.ResetDX2DModeDescription()` (same file), `.RenderDX2D()` (same file), `.Console_Closing()` (`Console/console.cs`)
- **`.DX2Adaptors()`** — L3326 — `public static AdaptorInfo[] DX2Adaptors()`
  Called by: `.showHelpInfo()` (`Console/console.cs`), `.Main()` (`Console/console.cs`)
- **`.getGPUNameInUse()`** — L3386 — `private static string getGPUNameInUse()`
  Returns gpuname in use.
  Called by: `.initDX2D()` (same file)
- **`.initDX2D()`** — L3408 — `private static void initDX2D(DriverType driverType = DriverType.Hardware, AdaptorInfo adaptorInfo = null)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DXVersion()`** — L3615 — `public static int DXVersion()`
  Called by: `.miAbout_Click()` (`Console/console.cs`)
- **`.ResetDX2DModeDescription()`** — L3641 — `public static void ResetDX2DModeDescription()`
  Resets dx2 dmode description.
  Called by: `.udDisplayFPS_ValueChanged()` (`Console/setup.cs`)
- **`.resizeDX2D()`** — L3668 — `private static bool resizeDX2D(out string error)`
  Called by: `.ResetDX2DModeDescription()` (same file), `.RenderDX2D()` (same file)
- **`.setupAliasing()`** — L3781 — `private static void setupAliasing()`
  Called by: `.initDX2D()` (same file), `.resizeDX2D()` (same file)
- **`.pauseDisplay()`** — L3831 — `private static void pauseDisplay()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RenderDX2D()`** — L3882 — `public static void RenderDX2D()`
  Renders dx2 d.
  Called by: `.RunDisplay()` (`Console/console.cs`)
- **`.showFPSProfile()`** — L4303 — `private static void showFPSProfile()`
  Called by: `.RenderDX2D()` (same file)
- **`.calcFps()`** — L4365 — `private static void calcFps()`
  Called by: `.RenderDX2D()` (same file)
- **`.isOccupied()`** — L4435 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] static private int isOccupied(int rx, int nX)`
  Called by: `.processMaximums()` (same file)
- **`.processMaximums()`** — L4462 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] static private void processMaximums(int rx, float dbm, int nX, int nY)`
  Called by: `.DrawPanadapterDX2D()` (same file)
- **`.ResetSpectrumPeaks()`** — L4527 — `static public void ResetSpectrumPeaks(int rx)`
  Resets spectrum peaks.
  Called by: `.resetPeaksAndNoise()` (same file), `.OnCentreFrequencyChanged()` (same file), `.DrawPanadapterDX2D()` (same file)
- **`.ResetBlobMaximums()`** — L4542 — `static public void ResetBlobMaximums(int rx, bool bClear = false)`
  Resets blob maximums.
  Called by: `.resetPeaksAndNoise()` (same file), `.OnCentreFrequencyChanged()` (same file), `.DrawPanadapterDX2D()` (same file)
- **`.getFilterXPositions()`** — L4564 — `static private void getFilterXPositions(int rx, int W, bool local_mox, bool displayduplex, out int filter_left_x, out int filter_right_x)`
  Returns filter xpositions.
  Called by: `.DrawPanadapterDX2D()` (same file)
- **`.isRxDuplex()`** — L4618 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] static private bool isRxDuplex(int rx)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.DrawWaterfallDX2D()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.drawSpots()` (same file)
- **`.modifyDataForNotches()`** — L4733 — `private static void modifyDataForNotches(ref float[] data, int rx, bool bottom, bool local_mox, bool displayduplex, int W)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.DrawPanadapterDX2D()`** — L4976 — `unsafe static private bool DrawPanadapterDX2D(int nVerticalShift, int W, int H, int rx, bool bottom)`
  Draws panadapter dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.findImd()`** — L5731 — `private static int findImd(Maximums[] sorted, int imd, int pixel_jump, int offset, bool low, out int X)`
  Called by: `.DrawPanadapterDX2D()` (same file)
- **`.processNoiseFloor()`** — L5872 — `private static void processNoiseFloor(int rx, int averageCount, float averageSum, int width, bool waterfall)`
  _bNoiseFloorAlreadyCalculatedRX2 = true; } }
  Called by: `.DrawPanadapterDX2D()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.resetWaterfallTimeOverlay()`** — L5922 — `private static void resetWaterfallTimeOverlay(int rx)`
  Called by: `.ResetWaterfallBmp()` (same file), `.ResetWaterfallBmp2()` (same file), `.resizeWaterfallTimeOverlay()` (same file), `.ResetWaterfallTimers()` (same file)
- **`.resizeWaterfallTimeOverlay()`** — L5936 — `private static void resizeWaterfallTimeOverlay(int rx, int waterHeight, int preservedRows)`
  Called by: `.ResetWaterfallBmp()` (same file), `.ResetWaterfallBmp2()` (same file)
- **`.getWaterfallLineIntervalMs()`** — L5976 — `private static double getWaterfallLineIntervalMs(int rx)`
  Returns waterfall line interval ms.
  Called by: `.recordWaterfallAdvance()` (same file)
- **`.recordWaterfallAdvance()`** — L5989 — `private static void recordWaterfallAdvance(int rx, int waterHeight)`
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.toDisplayTicks()`** — L6060 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static long toDisplayTicks(long utcTicks)`
  Called by: `.recordWaterfallAdvance()` (same file), `.drawWaterfallTimeOverlay()` (same file)
- **`.chooseWaterfallLabelIntervalMs()`** — L6068 — `private static long chooseWaterfallLabelIntervalMs(double msPerLine, double minLabelSpacing)`
  Called by: `.recordWaterfallAdvance()` (same file)
- **`.drawWaterfallTimeOverlay()`** — L6091 — `private static void drawWaterfallTimeOverlay(int nVerticalShift, int W, int H, int rx)`
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.SetPendingWaterfallPixelRef()`** — L6167 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public static void SetPendingWaterfallPixelRef(int rx, double pixel_ref)`
  Sets pending waterfall pixel ref.
  Called by: `.RunDisplay()` (`Console/console.cs`)
- **`.resetWaterfallBitmapAlignment()`** — L6174 — `private static void resetWaterfallBitmapAlignment(int rx)`
  Called by: `.ResetWaterfallBmp()` (same file), `.ResetWaterfallBmp2()` (same file), `.PurgeBuffers()` (same file)
- **`.getWaterfallSpanHz()`** — L6184 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static double getWaterfallSpanHz(int rx)`
  Returns waterfall span hz.
  Called by: `.prepareWaterfallBitmapShift()` (same file)
- **`.isWaterfallNoiseFloorCompensationEnabled()`** — L6189 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static bool isWaterfallNoiseFloorCompensationEnabled(int rx)`
  Called by: `.useWaterfallNoiseFloorCompensation()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.useWaterfallNoiseFloorCompensation()`** — L6196 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static bool useWaterfallNoiseFloorCompensation(int rx)`
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.getWaterfallNoiseFloorCompensationTarget()`** — L6204 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static float getWaterfallNoiseFloorCompensationTarget(int rx)`
  Returns waterfall noise floor compensation target.
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.prepareWaterfallBitmapShift()`** — L6212 — `private static int prepareWaterfallBitmapShift(int rx, int width, double centerMHz, out bool clearBitmap)`
  Called by: `.DrawWaterfallDX2D()` (same file)
- **`.clearWaterfallBitmapRegion()`** — L6265 — `private static void clearWaterfallBitmapRegion(SharpDX.Direct2D1.Bitmap bitmap, int x, int y, int width, int height)`
  Called by: `.ResetWaterfallBmp()` (same file), `.ResetWaterfallBmp2()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.ResetWaterfallTimers()`** — L6288 — `public static void ResetWaterfallTimers()`
  Resets waterfall timers.
  Called by: `.comboDisplayMode_SelectedIndexChanged()` (`Console/console.cs`)
- **`.OnWaterfallRXGradientChanged()`** — L6308 — `private static void OnWaterfallRXGradientChanged(int rx, Color[] colours)`
  Handles/raises the waterfall rxgradient changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.OnWaterfallTXGradientChanged()`** — L6340 — `private static void OnWaterfallTXGradientChanged(Color[] colours)`
  Handles/raises the waterfall txgradient changed event.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DrawWaterfallDX2D()`** — L6357 — `unsafe static private bool DrawWaterfallDX2D(int nVerticalShift, int W, int H, int rx, bool bottom)`
  Draws waterfall dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.convertColour()`** — L7717 — `private static Color4 convertColour(Color c)`
  Called by: `.convertBrush()` (same file), `.buildLinearGradientBrush()` (same file), `.buildLinearGradientBrushTX()` (same file), `.buildDX2Resources()` (same file)
- **`.convertBrush()`** — L7721 — `private static SharpDX.Direct2D1.SolidColorBrush convertBrush(SolidBrush b)`
  Called by: `.buildDX2Resources()` (same file), `.drawChannelBarDX2D()` (same file)
- **`.SetDX2BackgoundImage()`** — L7726 — `public static void SetDX2BackgoundImage(System.Drawing.Image image)`
  Sets dx2 backgound image.
  Called by: `.initDX2D()` (same file), `.setBackground()` (`Console/console.cs`)
- **`.SDXBitmapFromSysBitmap()`** — L7759 — `private static SharpDX.Direct2D1.Bitmap SDXBitmapFromSysBitmap(RenderTarget rt, System.Drawing.Bitmap bitmap)`
  Called by: `.SetDX2BackgoundImage()` (same file), `SetSantaGif()` (same file), `.getSpotFlagBitmap()` (same file)
- **`.buildLinearGradientBrush()`** — L7945 — `private static void buildLinearGradientBrush(int top, int bottom, int rx)`
  Called by: `.RenderDX2D()` (same file)
- **`.buildLinearGradientBrushTX()`** — L8032 — `private static void buildLinearGradientBrushTX(int top, int bottom, int rx)`
  Called by: `.RenderDX2D()` (same file)
- **`.releaseDX2Resources()`** — L8111 — `private static void releaseDX2Resources()`
  Called by: `.ShutdownDX2D()` (same file), `.buildDX2Resources()` (same file)
- **`.buildDX2Resources()`** — L8270 — `private static void buildDX2Resources()`
  Called by: `.initDX2D()` (same file)
- **`.releaseFonts()`** — L8374 — `private static void releaseFonts()`
  Called by: `.ShutdownDX2D()` (same file), `.buildFontsDX2D()` (same file)
- **`.buildFontsDX2D()`** — L8404 — `private static void buildFontsDX2D()`
  Called by: `.initDX2D()` (same file)
- **`.clearBackgroundDX2D()`** — L8427 — `static void clearBackgroundDX2D(int rx, int W, int H, bool bottom)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.drawLineDX2D()`** — L8457 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawLineDX2D(SharpDX.Direct2D1.Brush b, float x1, float y1, float x2, float y2, float strokeWidth = 1f)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.drawWaterfallTimeOverlay()` (same file), `.drawChannelBarDX2D()` (same file), `.handleNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.DrawSpectrumDX2D()` (same file) — and 7 more
- **`.drawFillRectangleDX2D()`** — L8467 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawFillRectangleDX2D(SharpDX.Direct2D1.Brush b, float x, float y, float w, float h)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.drawChannelBarDX2D()` (same file), `.handleNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.DrawSpectrumGridDX2D()` (same file), `.DrawPhaseDX2D()` (same file) — and 4 more
- **`.drawRectangleDX2D()`** — L8473 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawRectangleDX2D(SharpDX.Direct2D1.Brush b, float x, float y, float w, float h)`
  Called by: `.drawSpots()` (same file)
- **`.drawElipseDX2D()`** — L8479 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawElipseDX2D(SharpDX.Direct2D1.Brush b, float xMiddle, float yMiddle, float w, float h)`
  Called by: `.DrawPhaseGridDX2D()` (same file)
- **`.drawFillElipseDX2D()`** — L8485 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawFillElipseDX2D(SharpDX.Direct2D1.Brush b, float xMiddle, float yMiddle, float w, float h)`
  Called by: `.drawSpots()` (same file)
- **`.drawStringDX2D()`** — L8503 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] private static void drawStringDX2D(string s, SharpDX.DirectWrite.TextFormat tf, SharpDX.Direct2D1.Brush b, float x, float y)`
  Called by: `.RenderDX2D()` (same file), `.DrawPanadapterDX2D()` (same file), `.drawWaterfallTimeOverlay()` (same file), `.handleNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.DrawCursorInfoPanel()` (same file) — and 3 more
- **`.drawFilterOverlayDX2D()`** — L8509 — `private static void drawFilterOverlayDX2D(SharpDX.Direct2D1.Brush brush, int filter_left_x, int filter_right_x, int W, int H, int rx, int top, bool bottom, int nVerticalShfit)`
  Called by: `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.drawChannelBarDX2D()`** — L8520 — `private static void drawChannelBarDX2D(Channel chan, int left, int right, int top, int height, Color c, Color h)`
  Called by: `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.measureStringDX2D()`** — L8542 — `private static System.Drawing.SizeF measureStringDX2D(string s, SharpDX.DirectWrite.TextFormat tf, bool cacheStringLength = false)`
  Called by: `.recordWaterfallAdvance()` (same file), `.drawWaterfallTimeOverlay()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.DrawCursorInfoPanel()` (same file), `.DrawSpectrumGridDX2D()` (same file), `.drawSpots()` (same file) — and 1 more
- **`.getCWSideToneShift()`** — L8601 — `private static int getCWSideToneShift(int rx, DSPMode forceMode = DSPMode.FIRST)`
  Returns cwside tone shift.
  Called by: `.modifyDataForNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file), `.drawSpots()` (same file)
- **`.handleNotches()`** — L8644 — `private static List<clsNotchCoords> handleNotches(int rx, bool bottom, int cwSideToneShift, int Low, int High, int nVerticalShift, int top, int width, int W, int H, bool bDraw)`
  Called by: `.modifyDataForNotches()` (same file), `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.fastPow10Shifted()`** — L8772 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] static unsafe float fastPow10Shifted(float dBdiv10)`
  return ret; }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.fastPow10Raw()`** — L8794 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] static unsafe float fastPow10Raw(float dB)`
  return ret; }
  Called by: `.DrawPanadapterDX2D()` (same file), `.processNoiseFloor()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.drawPanadapterAndWaterfallGridDX2D()`** — L8802 — `private static int drawPanadapterAndWaterfallGridDX2D(int nVerticalShift, int W, int H, int rx, bool bottom, out long left_edge, out long right_edge, bool bIsWaterfall = false)`
  Called by: `.DrawPanadapterDX2D()` (same file), `.DrawWaterfallDX2D()` (same file)
- **`.AppendCursorInfoLines()`** — L10014 — `private static void AppendCursorInfoLines(List<string> lines, string text)`
  Called by: `.DrawCursorInfo()` (same file)
- **`.DrawCursorInfoPanel()`** — L10027 — `private static void DrawCursorInfoPanel(List<string> lines, int W, float yPos)`
  Draws cursor info panel.
  Called by: `.DrawCursorInfo()` (same file)
- **`.DrawCursorInfo()`** — L10072 — `private static void DrawCursorInfo(int W)`
  Draws cursor info.
  Called by: `.RenderDX2D()` (same file)
- **`.DrawSpectrumDX2D()`** — L10088 — `unsafe static private bool DrawSpectrumDX2D(int rx, int W, int H, bool bottom)`
  Draws spectrum dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.DrawSpectrumGridDX2D()`** — L10233 — `private static void DrawSpectrumGridDX2D(int W, int H, bool bottom)`
  Draws spectrum grid dx2 d.
  Called by: `.DrawSpectrumDX2D()` (same file), `.DrawHistogramDX2D()` (same file)
- **`.DrawScopeDX2D()`** — L10613 — `unsafe private static bool DrawScopeDX2D(int W, int H, bool bottom)`
  Draws scope dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.DrawScope2DX2D()`** — L10702 — `unsafe private static bool DrawScope2DX2D(int W, int H, bool bottom)`
  Draws scope2 dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.lerp()`** — L10784 — `private static float lerp(float first, float second, float by)`
  Called by: `.lerpPointF()` (same file)
- **`.lerpPointF()`** — L10788 — `private static PointF lerpPointF(PointF first, PointF second, float by)`
  Called by: `.DrawPhaseDX2D()` (same file)
- **`.DrawPhaseDX2D()`** — L10796 — `unsafe private static bool DrawPhaseDX2D(int W, int H, bool bottom)`
  Draws phase dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.DrawPhaseGridDX2D()`** — L10894 — `private static void DrawPhaseGridDX2D(int W, int H, bool bottom)`
  Draws phase grid dx2 d.
  Called by: `.DrawPhaseDX2D()` (same file), `.DrawPhase2DX2D()` (same file)
- **`.DrawPhase2DX2D()`** — L10906 — `unsafe private static void DrawPhase2DX2D(int W, int H, bool bottom)`
  Draws phase2 dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.DrawHistogramDX2D()`** — L10963 — `unsafe static private bool DrawHistogramDX2D(int rx, int W, int H)`
  Draws histogram dx2 d.
  Called by: `.RenderDX2D()` (same file)
- **`.getSpotLayer()`** — L11159 — `private static int getSpotLayer(int rx, int leftX)`
  Returns spot layer.
  Called by: `.drawSpots()` (same file)
- **`.updateLayer()`** — L11192 — `private static void updateLayer(int rx, int layer, int rightX)`
  Called by: `.drawSpots()` (same file)
- **`.getCallsignString()`** — L11212 — `private static string getCallsignString(SpotManager2.smSpot spot)`
  Returns callsign string.
  Called by: `.drawSpots()` (same file)
- **`.drawSpots()`** — L11230 — `public static void drawSpots(int rx, int nVerticalShift, int W, bool bottom)`
  Called by: `.RenderDX2D()` (same file)
- **`.clearAllDynamicBrushes()`** — L11553 — `private static void clearAllDynamicBrushes()`
  Called by: `.releaseDX2Resources()` (same file)
- **`.getDXBrushForColour()`** — L11570 — `private static SharpDX.Direct2D1.Brush getDXBrushForColour(Color c, int replaceAlpha = -1)`
  Returns dxbrush for colour.
  Called by: `.drawWaterfallTimeOverlay()` (same file), `.buildDX2Resources()` (same file), `.DrawCursorInfoPanel()` (same file), `.drawSpots()` (same file), `letItSnow()` (same file)
- **`.PurgeBuffers()`** — L11630 — `public static void PurgeBuffers()`
  Called by: `.OnPowerChangeHander()` (same file), `.RunDisplay()` (`Console/console.cs`)
- **`.EnumDisplaySettings()`** — L11935 — `[DllImport("user32.dll", CharSet = CharSet.Auto)] private static extern bool EnumDisplaySettings(string deviceName, int modeNum, ref DEVMODE devMode)`
  Called by: `.GetCurrentMonitorRefreshRate()` (same file)
- **`.GetCurrentMonitorRefreshRate()`** — L11938 — `public static int GetCurrentMonitorRefreshRate(Form form)`
  Returns current monitor refresh rate.
  Called by: `.btnGetMonitorHz_Click()` (`Console/setup.cs`)
- **`.getSpotTagSize()`** — L11952 — `private static SizeF getSpotTagSize(string displayText, System.Drawing.Image flagImage)`
  Returns spot tag size.
  Called by: `.drawSpots()` (same file)
- **`.drawSpotTagContent()`** — L11967 — `private static void drawSpotTagContent(SpotManager2.smSpot spot, string displayText, SharpDX.Direct2D1.Brush textBrush, Rectangle bounds)`
  Called by: `.drawSpots()` (same file)
- **`.getSpotFlagBitmap()`** — L12000 — `private static SharpDX.Direct2D1.Bitmap getSpotFlagBitmap(System.Drawing.Image flagImage)`
  Returns spot flag bitmap.
  Called by: `.drawSpotTagContent()` (same file)
- **`.getSpotFlagRenderSize()`** — L12039 — `private static void getSpotFlagRenderSize(System.Drawing.Image flagImage, int targetHeight, out int width, out int height)`
  Returns spot flag render size.
  Called by: `.getSpotTagSize()` (same file), `.drawSpotTagContent()` (same file)
- **`.clearSpotFlagBitmapCache()`** — L12051 — `private static void clearSpotFlagBitmapCache()`
  Called by: `.releaseDX2Resources()` (same file)

#### `BandEdgeRegionCacheDX2D` (type, L103)

- **`.Update()`** — L153 — `public void Update(FRSRegion region)`
  Called by: `.drawPanadapterAndWaterfallGridDX2D()` (same file)
- **`.Contains()`** — L203 — `[MethodImpl(MethodImplOptions.AggressiveInlining)] public bool Contains(int frequencyHz)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `WaterfallAgcCacheKey` (type, L2752)

- **`.Equals()`** — L2763 — `public bool Equals(WaterfallAgcCacheKey other)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.GetHashCode()`** — L2773 — `public override int GetHashCode()`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `WaterfallAgcCacheEntry` (type, L2779)

_No extracted members._

#### `AdaptorInfo` (type, L3313)

_No extracted members._

#### `Maximums` (type, L4427)

_No extracted members._

#### `clsNotchCoords` (type, L8627)

_No extracted members._

#### `SnowFlake` (type, L11662)

- **`.Update()`** — L11685 — `public void Update()`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

#### `DEVMODE` (type, L11903)

_No extracted members._

#### `ReferenceEqualityComparer` (type, L12085)

- **`.Equals()`** — L12089 — `public bool Equals(T x, T y)`
  Called by: `.Equals()` (same file), `.drawSpots()` (same file)
- **`.GetHashCode()`** — L12094 — `public int GetHashCode(T obj)`
  Returns hash code.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/display.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
