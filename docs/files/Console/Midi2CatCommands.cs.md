# `Console/Midi2CatCommands.cs`

**Functional area:** [12. MIDI control (Midi2Cat)](../../CODE_OUTLINE.md#12-midi-control-midi2cat)

**Role:** The bridge: exposes console operations (tune, volume, filters, PTT…) as commands a MIDI control can bind to (256-edge god node).

## How this file is used

- Used by (incoming references from other files):
  - `Console/console.cs` (references ×1)
- Uses (outgoing references to other files):
  - `Console/CAT/CATCommands.cs` (calls ×273, references ×1)
  - `Midi2Cat/Midi2Cat.IO/MidiDevice.cs` (references ×241, calls ×12)
  - `Midi2Cat/Midi2Cat.Data/Enums.cs` (references ×69)
  - `Midi2Cat/MidiMessageManager.cs` (calls ×3, references ×1)
  - `Console/Andromeda/Andromeda.cs` (references ×1)
  - `Console/CAT/CATParser.cs` (references ×1)
  - `Midi2Cat/Midi2Cat.Data/CatCmdDb.cs` (references ×1)

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `Midi2CatCommands` (type, L54)

- **`.OpenMidi2Cat()`** — L73 — `public void OpenMidi2Cat()`
  Opens midi2 cat.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CloseMidi2Cat()`** — L78 — `public void CloseMidi2Cat()`
  Closes midi2 cat.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SendUpdateToMidi()`** — L97 — `public void SendUpdateToMidi(CatCmd cmd, double pct)`
  Sends update to midi.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MultiRxOnOff()`** — L109 — `public CmdState MultiRxOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1ModeNext()`** — L134 — `public void Rx1ModeNext(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1ModePrev()`** — L147 — `public void Rx1ModePrev(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1FilterWider()`** — L159 — `public void Rx1FilterWider(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1FilterNarrower()`** — L172 — `public void Rx1FilterNarrower(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VfoAtoB()`** — L185 — `public void VfoAtoB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VfoBtoA()`** — L201 — `public void VfoBtoA(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VfoSwap()`** — L216 — `public void VfoSwap(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.XIT()`** — L233 — `public void XIT(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RIT()`** — L252 — `public void RIT(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IsBehringerCMD()`** — L280 — `bool IsBehringerCMD(MidiDevice device)`
  Called by: `.RIT_inc()` (same file), `.XIT_inc()` (same file), `.AGCLevel()` (same file), `.RX2AGCLevel()` (same file), `.FilterHigh()` (same file), `.FilterLow()` (same file) — and 2 more
- **`.RIT_inc()`** — L286 — `public void RIT_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.XIT_inc()`** — L322 — `public void XIT_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RIT_clear()`** — L377 — `public void RIT_clear(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.XIT_clear()`** — L388 — `public void XIT_clear(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuningStepUp()`** — L399 — `public void TuningStepUp(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TuningStepDown()`** — L410 — `public void TuningStepDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VolumeVfoA()`** — L423 — `public void VolumeVfoA(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VolumeVfoA_inc()`** — L441 — `public void VolumeVfoA_inc(int msg, MidiDevice device)`
  -W2PA Incremental volume control for Behringer PL-1 or similar knobs as wheels. Also added an item for Wheel in CatCmdDb.cs
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VolumeVfoB()`** — L465 — `public void VolumeVfoB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VolumeVfoB_inc()`** — L485 — `public void VolumeVfoB_inc(int msg, MidiDevice device)`
  -W2PA Incremental volume control for Behringer PL-1 or similar knobs as wheels. Also added an item for Wheel in CatCmdDb.cs
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2Volume()`** — L509 — `public void RX2Volume(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2Pan()`** — L526 — `public void RX2Pan(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterBandwidth()`** — L544 — `public void FilterBandwidth(int msg, MidiDevice device)`
  case 52: // Pitch DeckA - FilterBandwidth
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterShift()`** — L565 — `public void FilterShift(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RatioMainSubRx()`** — L585 — `public void RatioMainSubRx(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AutoNotchOnOff()`** — L605 — `public CmdState AutoNotchOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1NoiseBlanker1OnOff()`** — L629 — `public CmdState Rx1NoiseBlanker1OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseBlanker1OnOff()`** — L654 — `public CmdState Rx2NoiseBlanker1OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx1Noiseblanker2OnOff()`** — L684 — `public CmdState Rx1Noiseblanker2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2Noiseblanker2OnOff()`** — L714 — `public CmdState Rx2Noiseblanker2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LockVFOOnOff()`** — L745 — `public CmdState LockVFOOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LockVFOAOnOff()`** — L768 — `public CmdState LockVFOAOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.LockVFOBOnOff()`** — L791 — `public CmdState LockVFOBOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RitOnOff()`** — L814 — `public CmdState RitOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.XitOnOff()`** — L837 — `public CmdState XitOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetAFGain()`** — L860 — `public void SetAFGain(int msg, MidiDevice device)`
  Sets afgain.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityFormOpen()`** — L868 — `public CmdState DiversityFormOpen(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityEnable()`** — L891 — `public CmdState DiversityEnable(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityPhase()`** — L914 — `public void DiversityPhase(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityGain()`** — L944 — `public void DiversityGain(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversityReference()`** — L974 — `public CmdState DiversityReference(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DiversitySource()`** — L997 — `public CmdState DiversitySource(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StringToFreq()`** — L1037 — `public int StringToFreq(string s)`
  Called by: `.ChangeFreqVfoA()` (same file), `.ChangeFreqVfoB()` (same file)
- **`.MidiMessagesPerTuneStepUp()`** — L1129 — `public void MidiMessagesPerTuneStepUp(int msg, MidiDevice device)`
  -W2PA This increments or decrements the number of MIDI messages that cause a single tune step increment It is useful when using coarse increments, such as 100kHz, and wanting more wheel rotation for each one so that tuning isn't so critical.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiMessagesPerTuneStepDown()`** — L1133 — `public void MidiMessagesPerTuneStepDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MidiMessagesPerTuneStepToggle()`** — L1137 — `public CmdState MidiMessagesPerTuneStepToggle(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ProcessStdMIDIWheelAsVFO()`** — L1173 — `private void ProcessStdMIDIWheelAsVFO(int direction, int step, bool round_to_step_size, long freq, int mode,`
  [2.10.3.9]MW0LGE refactor for speed, as other implemation was just a complete mess also using invoke and begininvoke with func/actions instead of helper functions. I am not interested in changing the code for the Behringer P1/micro etc.
  Called by: `.ChangeFreqVfoA()` (same file), `.ChangeFreqVfoB()` (same file)
- **`.ProcessBehringerMainWheelAsVFO()`** — L1377 — `private void ProcessBehringerMainWheelAsVFO(int direction, int step, bool RoundToStepSize, long freq, int mode, string vfo, string deviceName)`
  -W2PA Routine to implement variable speed tuning using the Behringer CMD PL-1 (and others) MIDI controller main wheel
  Called by: `.ChangeFreqVfoA()` (same file), `.ChangeFreqVfoB()` (same file)
- **`.ChangeFreqVfoA()`** — L1577 — `public void ChangeFreqVfoA(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SnapTune()`** — L1748 — `public long SnapTune(long freq, int step, int num_steps, bool RoundToStepSize)`
  Called by: `.ProcessStdMIDIWheelAsVFO()` (same file), `.ProcessBehringerMainWheelAsVFO()` (same file)
- **`.ChangeFreqVfoB()`** — L1887 — `public void ChangeFreqVfoB(int msg, MidiDevice device)`
  -W2PA Modified to select Behringer PL-1, Micro, or original code
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BinauralOnOff()`** — L1963 — `public CmdState BinauralOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MuteOnOff()`** — L1986 — `public CmdState MuteOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SpurReductionOnOff()`** — L2009 — `public CmdState SpurReductionOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseReduction4Amount()`** — L2032 — `public void NoiseReduction4Amount(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseReductionOnOff()`** — L2049 — `public CmdState NoiseReductionOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseReduction2OnOff()`** — L2072 — `public CmdState NoiseReduction2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseReduction3OnOff()`** — L2094 — `public CmdState NoiseReduction3OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.NoiseReduction4OnOff()`** — L2116 — `public CmdState NoiseReduction4OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseReduction4Amount()`** — L2138 — `public void Rx2NoiseReduction4Amount(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseReductionOnOff()`** — L2155 — `public CmdState Rx2NoiseReductionOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseReduction2OnOff()`** — L2177 — `public CmdState Rx2NoiseReduction2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseReduction3OnOff()`** — L2199 — `public CmdState Rx2NoiseReduction3OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2NoiseReduction4OnOff()`** — L2221 — `public CmdState Rx2NoiseReduction4OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2PreAmpOnOff()`** — L2243 — `public CmdState Rx2PreAmpOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VfoSyncOnOff()`** — L2273 — `public CmdState VfoSyncOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SplitOnOff()`** — L2296 — `public CmdState SplitOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MOXOnOff()`** — L2319 — `public CmdState MOXOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VOXOnOff()`** — L2342 — `public CmdState VOXOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CompanderOnOff()`** — L2365 — `public CmdState CompanderOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StereoDiversityOnOff()`** — L2388 — `public CmdState StereoDiversityOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DEXPOnOff()`** — L2411 — `public CmdState DEXPOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2OnOff()`** — L2434 — `public CmdState RX2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.StartOnOff()`** — L2463 — `public CmdState StartOnOff(int msg, MidiDevice device)`
  Starts on off.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TunerOnOff()`** — L2494 — `public CmdState TunerOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TunOnOff()`** — L2524 — `public CmdState TunOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TwoToneOnOff()`** — L2554 — `public CmdState TwoToneOnOff(int msg, MidiDevice device)`
  MW0LGE_21g
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TunerBypassOnOff()`** — L2583 — `public CmdState TunerBypassOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZeroBeatPress()`** — L2614 — `public void ZeroBeatPress(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandUp()`** — L2625 — `public void BandUp(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.BandDown()`** — L2644 — `public void BandDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2BandUp()`** — L2663 — `public void Rx2BandUp(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2BandDown()`** — L2682 — `public void Rx2BandDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreAmpSettingsKnob()`** — L2701 — `public void PreAmpSettingsKnob(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWBreakIn()`** — L2825 — `public CmdState CWBreakIn(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWQSK()`** — L2848 — `public CmdState CWQSK(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWSpeed()`** — L2871 — `public void CWSpeed(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWSpeed_inc()`** — L2893 — `public void CWSpeed_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFFreq()`** — L2917 — `public void APFFreq(int msg, MidiDevice device)`
  -W2PA Added knob/slider control of APF Tune
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFBandwidth()`** — L2943 — `public void APFBandwidth(int msg, MidiDevice device)`
  -W2PA Added knob/slider control of APF Bandwidth
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFGain()`** — L2961 — `public void APFGain(int msg, MidiDevice device)`
  -W2PA Added knob/slider control of APF Gain
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AGCLevel()`** — L2978 — `public void AGCLevel(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AGCLevel_inc()`** — L3009 — `public void AGCLevel_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AGCLevel()`** — L3046 — `public void RX2AGCLevel(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AGCLevel_inc()`** — L3077 — `public void RX2AGCLevel_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MicGain()`** — L3115 — `public void MicGain(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SquelchControl()`** — L3143 — `public void SquelchControl(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CPDRLevel()`** — L3160 — `public void CPDRLevel(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VOXGain()`** — L3203 — `public void VOXGain(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DEXPThreshold()`** — L3220 — `public void DEXPThreshold(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXAFMonitor()`** — L3245 — `public void TXAFMonitor(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DriveLevel()`** — L3262 — `public void DriveLevel(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DriveLevel_inc()`** — L3279 — `public void DriveLevel_inc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RXEQOnOff()`** — L3334 — `public CmdState RXEQOnOff(int msg, MidiDevice device)`
  commands.ZZPC(drive.ToString("000")); return; } catch { return; } }
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXEQOnOff()`** — L3357 — `public CmdState TXEQOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SquelchOnOff()`** — L3380 — `public CmdState SquelchOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AGCModeKnob()`** — L3410 — `public void AGCModeKnob(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AGCModeUp()`** — L3450 — `public void AGCModeUp(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AGCModeDown()`** — L3475 — `public void AGCModeDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PreampFlex5000()`** — L3500 — `public void PreampFlex5000(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayAverage()`** — L3531 — `public void DisplayAverage(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayPeak()`** — L3560 — `public void DisplayPeak(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayTxFilter()`** — L3589 — `public void DisplayTxFilter(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VACOnOff()`** — L3618 — `public CmdState VACOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VAC2OnOff()`** — L3648 — `public CmdState VAC2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IQtoVAC()`** — L3678 — `public void IQtoVAC(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.IQtoVACRX2()`** — L3707 — `public void IQtoVACRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayModePrev()`** — L3739 — `public void DisplayModePrev(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.DisplayModeNext()`** — L3771 — `public void DisplayModeNext(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomDec()`** — L3796 — `public void ZoomDec(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomInc()`** — L3840 — `public void ZoomInc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomSliderInc()`** — L3890 — `public void ZoomSliderInc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PanSliderInc()`** — L3918 — `public void PanSliderInc(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PanSlider()`** — L3945 — `public void PanSlider(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SpectralNoiseBlankerOnOff()`** — L3963 — `public CmdState SpectralNoiseBlankerOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SpectralNoiseBlankerRx2OnOff()`** — L3986 — `public CmdState SpectralNoiseBlankerRx2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickModeSave()`** — L4009 — `public void QuickModeSave(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro1()`** — L4028 — `public void CWXMacro1(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro2()`** — L4047 — `public void CWXMacro2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro3()`** — L4066 — `public void CWXMacro3(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro4()`** — L4085 — `public void CWXMacro4(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro5()`** — L4104 — `public void CWXMacro5(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro6()`** — L4123 — `public void CWXMacro6(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro7()`** — L4142 — `public void CWXMacro7(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro8()`** — L4161 — `public void CWXMacro8(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXMacro9()`** — L4180 — `public void CWXMacro9(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXStop()`** — L4199 — `public void CWXStop(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MONOnOff()`** — L4218 — `public CmdState MONOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PanCenter()`** — L4241 — `public void PanCenter(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickModeRestore()`** — L4260 — `public void QuickModeRestore(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomSliderFix()`** — L4279 — `public void ZoomSliderFix(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterHigh()`** — L4296 — `public void FilterHigh(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.FilterLow()`** — L4394 — `public void FilterLow(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VACGainRX()`** — L4493 — `public void VACGainRX(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VACGainTX()`** — L4511 — `public void VACGainTX(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VAC2GainRX()`** — L4530 — `public void VAC2GainRX(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.VAC2GainTX()`** — L4548 — `public void VAC2GainTX(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CTunOnOff()`** — L4590 — `public CmdState CTunOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ESCFormOnOff()`** — L4613 — `public CmdState ESCFormOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallLowLimit()`** — L4636 — `public void WaterfallLowLimit(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.WaterfallHighLimit()`** — L4655 — `public void WaterfallHighLimit(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MuteRX2OnOff()`** — L4674 — `public CmdState MuteRX2OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band160m()`** — L4697 — `public void Band160m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band80m()`** — L4716 — `public void Band80m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band60m()`** — L4735 — `public void Band60m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band40m()`** — L4754 — `public void Band40m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band30m()`** — L4773 — `public void Band30m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band20m()`** — L4792 — `public void Band20m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band17m()`** — L4812 — `public void Band17m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band15m()`** — L4831 — `public void Band15m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band12m()`** — L4850 — `public void Band12m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band10m()`** — L4869 — `public void Band10m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band6m()`** — L4888 — `public void Band6m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band2m()`** — L4907 — `public void Band2m(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band160mRX2()`** — L4926 — `public void Band160mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band80mRX2()`** — L4945 — `public void Band80mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band60mRX2()`** — L4964 — `public void Band60mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band40mRX2()`** — L4984 — `public void Band40mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band30mRX2()`** — L5003 — `public void Band30mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band20mRX2()`** — L5022 — `public void Band20mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band17mRX2()`** — L5042 — `public void Band17mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band15mRX2()`** — L5061 — `public void Band15mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band12mRX2()`** — L5080 — `public void Band12mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band10mRX2()`** — L5099 — `public void Band10mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band6mRX2()`** — L5118 — `public void Band6mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Band2mRX2()`** — L5137 — `public void Band2mRX2(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeSSB()`** — L5156 — `public void ModeSSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeLSB()`** — L5192 — `public void ModeLSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeUSB()`** — L5210 — `public void ModeUSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeDSB()`** — L5229 — `public void ModeDSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeCW()`** — L5248 — `public void ModeCW(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeCWL()`** — L5267 — `public void ModeCWL(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeCWU()`** — L5286 — `public void ModeCWU(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeFM()`** — L5305 — `public void ModeFM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeAM()`** — L5324 — `public void ModeAM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeDIGU()`** — L5343 — `public void ModeDIGU(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeSPEC()`** — L5363 — `public void ModeSPEC(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeDIGL()`** — L5382 — `public void ModeDIGL(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeSAM()`** — L5401 — `public void ModeSAM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ModeDRM()`** — L5421 — `public void ModeDRM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveVFOADown100Khz()`** — L5440 — `public void MoveVFOADown100Khz(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveVFOAUp100Khz()`** — L5459 — `public void MoveVFOAUp100Khz(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APF_OnOff()`** — L5478 — `public CmdState APF_OnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToggleVFOWheel()`** — L5501 — `public void ToggleVFOWheel(int msg, MidiDevice device)`
  Toggles vfowheel.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2ModeNext()`** — L5528 — `public void Rx2ModeNext(int msg, MidiDevice device)`
  DH1KLM_21g block of additions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2ModePrev()`** — L5540 — `public void Rx2ModePrev(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2FilterWider()`** — L5552 — `public void Rx2FilterWider(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.Rx2FilterNarrower()`** — L5564 — `public void Rx2FilterNarrower(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AutoNotchOnOff()`** — L5576 — `public CmdState RX2AutoNotchOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ToggleTX()`** — L5599 — `public CmdState ToggleTX(int msg, MidiDevice device)`
  Toggles tx.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TUNPowerLevel()`** — L5622 — `public void TUNPowerLevel(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AGCModeKnob()`** — L5639 — `public void RX2AGCModeKnob(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AGCModeUp()`** — L5679 — `public void RX2AGCModeUp(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AGCModeDown()`** — L5704 — `public void RX2AGCModeDown(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2CTunOnOff()`** — L5729 — `public CmdState RX2CTunOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.PSOnOff()`** — L5752 — `public CmdState PSOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeSSB()`** — L5775 — `public void RX2ModeSSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeLSB()`** — L5811 — `public void RX2ModeLSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeUSB()`** — L5829 — `public void RX2ModeUSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeDSB()`** — L5848 — `public void RX2ModeDSB(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeCW()`** — L5867 — `public void RX2ModeCW(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeCWL()`** — L5886 — `public void RX2ModeCWL(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeCWU()`** — L5905 — `public void RX2ModeCWU(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeFM()`** — L5924 — `public void RX2ModeFM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeAM()`** — L5943 — `public void RX2ModeAM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeDIGU()`** — L5962 — `public void RX2ModeDIGU(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeSPEC()`** — L5982 — `public void RX2ModeSPEC(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeDIGL()`** — L6001 — `public void RX2ModeDIGL(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeSAM()`** — L6020 — `public void RX2ModeSAM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2ModeDRM()`** — L6040 — `public void RX2ModeDRM(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveVFOBDown100Khz()`** — L6059 — `public void MoveVFOBDown100Khz(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.MoveVFOBUp100Khz()`** — L6078 — `public void MoveVFOBUp100Khz(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CloseConsole()`** — L6097 — `public void CloseConsole(int msg, MidiDevice device)`
  Closes console.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2SquelchOnOff()`** — L6116 — `public CmdState RX2SquelchOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2SquelchControl()`** — L6146 — `public void RX2SquelchControl(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXFilterHigh()`** — L6165 — `public void TXFilterHigh(int msg, MidiDevice device)`
  end DH1KLM_21h start block of additions
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.TXFilterLow()`** — L6262 — `public void TXFilterLow(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ExternalPAOnOff()`** — L6362 — `public CmdState ExternalPAOnOff(int msg, MidiDevice device)`
  MW0LGE_21j
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXKey()`** — L6385 — `public void CWXKey(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CWXPTT()`** — L6399 — `public void CWXPTT(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomToBandRecall()`** — L6414 — `public void ZoomToBandRecall(int msg, MidiDevice device)`
  MW0LGE_21k9d
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.ZoomToBandStore()`** — L6432 — `public void ZoomToBandStore(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX1AutoAGC()`** — L6451 — `public CmdState RX1AutoAGC(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.RX2AutoAGC()`** — L6473 — `public CmdState RX2AutoAGC(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SwapVFOWheels()`** — L6495 — `public CmdState SwapVFOWheels(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickSplitOnOff()`** — L6518 — `public CmdState QuickSplitOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickSplitOnOffandSplitOnOff()`** — L6540 — `public CmdState QuickSplitOnOffandSplitOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickPlayOnOff()`** — L6574 — `public CmdState QuickPlayOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.QuickRecOnOff()`** — L6603 — `public CmdState QuickRecOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.AudioAmpOnOff()`** — L6632 — `public CmdState AudioAmpOnOff(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFType_doublepole()`** — L6654 — `public void APFType_doublepole(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFType_matched()`** — L6668 — `public void APFType_matched(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFType_gaussian()`** — L6682 — `public void APFType_gaussian(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.APFType_biquad()`** — L6696 — `public void APFType_biquad(int msg, MidiDevice device)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Midi2CatCommands.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python docs/tools/gen_file_docs.py`._
