# Developer Notes — Thetis for Hermes-Lite 2

## Build Process

### Prerequisites

| Tool | Version | Notes |
|------|---------|-------|
| Visual Studio | 2022 (v17.x) | Community or higher; install C++ and .NET desktop workloads |
| .NET Framework | 4.8 | Target runtime; included in VS workload |
| WiX Toolset | 3.14.x | Required for installer project; install from wixtoolset.org |
| NuGet | Any | Included with VS; restores Discord.Net and other packages |

### Solution File

```
Project Files/Source/Thetis_VS2026.sln
```

The solution name reflects the development year, not a VS version. Open it in Visual Studio 2022. Project files specify toolset `v145`; VS 2022 will prompt to retarget — accept the retarget or override the toolset manually.

### Build Configurations

The primary release configuration is **Release | x64**. The full matrix is:

| Configuration | Platform | Notes |
|--------------|----------|-------|
| Release | x64 | Production build — use this for releases |
| Release | x86 | 32-bit build |
| Release | Mixed Platforms | Builds the WiX installer automatically |
| Debug | x64 | Development/debugging |

> **Note:** The `Thetis-Installer` WiX project is excluded from the `Release|x64` solution-level build. Build it separately (see below) or switch to `Release|Mixed Platforms` to have it included automatically.

### Build Steps

**1. Restore NuGet packages**

```
nuget restore "Project Files/Source/Thetis_VS2026.sln"
```

**2. Build the main solution (Release | x64)**

```
msbuild "Project Files/Source/Thetis_VS2026.sln" \
    /p:Configuration=Release \
    /p:Platform=x64 \
    /p:PlatformToolset=v143 \
    /m /v:minimal
```

The `/p:PlatformToolset=v143` override is required when building outside VS 2022 (e.g., CI) because the project files declare `v145`.

**3. Build the installer separately**

```
msbuild "Project Files/Source/Thetis-Installer/Thetis-Installer.wixproj" \
    /p:Configuration=Release \
    /p:Platform=x64 \
    /p:VisualStudioVersion=17.0
```

Output: `Project Files/bin/Installers/Thetis-v<version>.x64.msi`

The installer AfterBuild target reads the version from the compiled `Thetis.exe` assembly and renames the MSI accordingly.

### Output Locations

| Artifact | Path |
|----------|------|
| Main executable | `Project Files/Bin/x64/Release/Thetis.exe` |
| DSP library | `Project Files/Bin/x64/Release/wdsp.dll` |
| Channel master | `Project Files/Bin/x64/Release/ChannelMaster.dll` |
| ASIO wrapper | `Project Files/Bin/x64/Release/cmASIO.dll` |
| MSI installer | `Project Files/bin/Installers/Thetis-v<version>.x64.msi` |

### CI/CD (GitHub Actions)

The workflow at `.github/workflows/build.yml` runs on every push to `master` and on `v*` tags.

- **Push to master:** builds the solution and installer, uploads both as workflow artifacts.
- **Push a `v*` tag:** additionally runs the `release` job, which publishes the MSI to GitHub Releases.

To make a release:

```bash
git tag v2.10.x.y
git push origin v2.10.x.y
```

---

## Source Code Description

### Project Map

```
Project Files/Source/
├── Console/            — Main application (C#, WinForms)
├── wdsp/               — DSP engine (C)
├── ChannelMaster/      — Audio I/O routing (C/C++)
├── cmASIO/             — ASIO audio driver wrapper (C++)
├── Midi2Cat/           — MIDI→CAT controller mapping (C#)
├── RawInput/           — Raw keyboard/mouse input (C#)
├── Thetis-Installer/   — Windows installer (WiX)
└── packages/           — NuGet package cache

Project Files/lib/
├── fftw_x64/           — FFTW3 DLLs (double, single, long-double precision)
├── NR_Algorithms_x64/  — rnnoise.dll, specbleach.dll + headers
├── portaudio-19.7.0/   — PortAudio source + build project
└── SharpDX/            — Managed DirectX wrappers (Direct2D, D3D11, DXGI)
```

---

### Console — Main Application

`Project Files/Source/Console/`

The WinForms application (~180+ source files) that presents the SDR user interface and coordinates all subsystems.

#### Application Entry & Main Window

| File | Description |
|------|-------------|
| `console.cs` / `.Designer.cs` | Main application window; the central hub connecting all subsystems |
| `splash.cs` | Startup splash screen |
| `setup.cs` / `.designer.cs` | Setup/preferences dialog covering radio, audio, display, and CAT configuration |
| `AssemblyInfo.cs` | Version and assembly metadata |
| `app.config` | .NET application configuration |

#### Radio & Hardware Control

| File | Description |
|------|-------------|
| `radio.cs` | High-level radio operations (PTT, mode, band, VFO) |
| `common.cs` | Shared utility functions and constants |
| `clsHardwareSpecific.cs` | Per-hardware adjustments (power limits, band plans) |
| `xvtr.cs` | Transverter frequency offset configuration |
| `hiperftimer.cs` | High-resolution timer used for timing-critical radio operations |
| `Firewall.cs` | Adds/removes Windows Firewall rules at startup |

#### HPSDR Network Protocol

`Console/HPSDR/`

| File | Description |
|------|-------------|
| `NetworkIO.cs` / `NetworkIOImports.cs` | UDP packet I/O implementing the OpenHPSDR protocol |
| `clsRadioDiscovery.cs` | Discovers HPSDR-compatible radios on the LAN |
| `specHPSDR.cs` | HPSDR-specific spectrum data handling |
| `Alex.cs` | Alex antenna switch/LPF board interface |
| `Penny.cs` | Penny I/O board interface (band voltages, keying) |
| `IoBoardHl2.cs` | Hermes-Lite 2 specific I/O board extensions |

#### Display & Spectrum

| File | Description |
|------|-------------|
| `display.cs` | Core spectrum/waterfall rendering engine |
| `PanDisplay.cs` | Panadapter display control |
| `wbDisplay.cs` | Wideband (multi-band overview) display |
| `ucBandwidthView.cs` | Bandwidth visualization widget |
| `ucMeter.cs` | S-meter, power, and ALC meter rendering |
| `ucInfoBar.cs` | Information bar showing frequency, mode, and status |
| `ucVARGrapher.cs` | Voice-activated relay level graphing |
| `ucUnderOverFlowWarningViewer.cs` | Buffer underflow/overflow visual indicator |
| `TuneStep.cs` | Frequency tuning step selector |
| `Path_Illustrator.cs` | Signal-path diagram drawing |
| `Skin.cs` / `clsThetisSkinService.cs` | Skin/theme loading and application |

#### Audio & DSP Control

| File | Description |
|------|-------------|
| `audio.cs` | Audio device enumeration, selection, and stream management |
| `BasicAudio.cs` | Low-level audio utility functions |
| `dsp.cs` | Configures and controls the wdsp processing chain |
| `filter.cs` | Filter selection and management |
| `FilterForm.cs` / `frmFilterManager.cs` | Filter editor UI |
| `eqform.cs` | Parametric equalizer UI |
| `portaudio.cs` | .NET P/Invoke bindings to PortAudio |
| `ringbuffer.cs` | Lock-free ring buffer for audio streaming between threads |
| `clsAudioRecordPlayback.cs` | Audio file recording and playback |

#### CAT (Computer-Aided Transceiver) Protocol

`Console/CAT/`

| File | Description |
|------|-------------|
| `CATCommands.cs` | Implements the full Kenwood-compatible CAT command set |
| `CATParser.cs` | Parses incoming CAT byte streams |
| `CATStructs.xml` | XML definition of all CAT command structures |
| `CATTester.cs` | Built-in CAT protocol tester/debugger |
| `TCPIPcatServer.cs` | TCP server exposing CAT over Ethernet |
| `SIOListenerII.cs` | Serial port CAT listener |
| `SDRSerialPortII.cs` | Serial port abstraction for CAT and PTT |
| `SerialPortPTT.cs` | PTT keying via serial port RTS/DTR |
| `UsbBCDCable.cs` | USB CDC cable support for CAT |

#### CW (Morse Code)

| File | Description |
|------|-------------|
| `cwkeyer.cs` | CW keyer timing and iambic keyer logic |
| `cwedit.cs` | CW message editor |
| `cwx.cs` | CW memory keyer (macros, playback) |
| `CW/CWInput.cs` | CW paddle/key input handling |

#### Receiver Channels

| File | Description |
|------|-------------|
| `Channel.cs` | Receiver channel state and parameters |
| `rxa.cs` | RX-A channel controls (gain, filter, mode per receiver) |
| `rxaControls.cs` | UI controls bound to RX channel parameters |
| `AmpView.cs` | Amplifier/attenuator status display |

#### Memory & Band Stack

| File | Description |
|------|-------------|
| `clsBandStackManager.cs` | Manages per-band frequency/mode memory stacks |
| `frmBandStack2.cs` | Band stack browser and editor UI |
| `Memory/MemoryForm.cs` | General memory channel browser |
| `Memory/MemoryRecord.cs` | Memory channel data structure |
| `Memory/DXMemRecord.cs` | DX spot memory record |
| `database.cs` / `clsDBMan.cs` | SQLite-backed settings and memory persistence |

#### Andromeda UI Framework

`Console/Andromeda/` — An alternative touch-friendly control surface that mirrors the main console.

| File | Description |
|------|-------------|
| `Andromeda.cs` | Main Andromeda surface controller |
| `AndromedaEditForm.cs` | Layout editor for the Andromeda surface |
| `BandButtonsPopup.cs` | Band selection popup panel |
| `FilterButtonsPopup.cs` | Filter selection popup |
| `ModeButtonsPopup.cs` | Modulation mode popup |
| `vfosettingspopup.cs` | VFO configuration popup |
| `displaysettingsform.cs` | Display settings accessible from Andromeda |

#### Thread-Safe UI Controls

`Console/Invoke/` — Wrappers making standard WinForms controls safe to update from non-UI threads.

`invoke.cs`, `buttonts.cs`, `checkboxts.cs`, `comboboxts.cs`, `labelts.cs`, `numericupdownts.cs`, `radiobuttonts.cs`, `textboxts.cs`, `trackbarts.cs`

#### External Integrations

| File | Description |
|------|-------------|
| `N1MM.cs` | N1MM+ contest logger UDP interface |
| `SpotManager2.cs` | DX cluster spot management and display |
| `TCIServer.cs` | TCI (Thetis Control Interface) TCP server for external SDR apps |
| `clsDiscord.cs` | Discord rich-presence and webhook integration |
| `clsMeterScriptEngine.cs` | Scripting engine for custom meter layouts |
| `MusicPlayback.cs` | Background music playback for demonstrations |
| `Midi2CatCommands.cs` | MIDI-to-CAT command binding definitions (used with Midi2Cat project) |
| `NetworkThrottle.cs` | Limits outgoing network bandwidth for constrained links |

#### Utilities

| File | Description |
|------|-------------|
| `enums.cs` | All global enum definitions |
| `keyboard.cs` | Keyboard shortcut handling |
| `clsDPISafeTools.cs` | DPI-aware layout helpers |
| `clsCountryData.cs` / `clsFlagAtlas.cs` | Country identification and flag rendering for DX spots |
| `clsSingleInstance.cs` | Prevents multiple simultaneous instances |
| `clsSpectrumProcessor.cs` | Spectrum data post-processing (peak hold, averaging) |
| `Versions.cs` / `VersionInfo.cs` | Runtime version checking and display |
| `GlobalMouseHandler.cs` | Application-wide mouse wheel interception |
| `clsTouchHandler.cs` | Touchscreen gesture support |

---

### wdsp — DSP Signal Processing Engine

`Project Files/Source/wdsp/`

A pure-C library (~143 files) that implements all digital signal processing. It is loaded as `wdsp.dll` and called from the Console via P/Invoke. FFTW3 provides the underlying FFT.

#### Core Signal Chains

| File | Description |
|------|-------------|
| `RXA.c` / `RXA.h` | Receive signal processing chain (builds the RX DSP graph) |
| `TXA.c` / `TXA.h` | Transmit signal processing chain |
| `channel.c` / `channel.h` | Multi-channel management and routing |
| `main.c` / `main.h` | Library initialization, DLL exports |

#### Buffering & I/O

| File | Description |
|------|-------------|
| `iobuffs.c` / `iobuffs.h` | Input/output buffer management between audio and DSP |
| `syncbuffs.c` / `syncbuffs.h` | Synchronized buffer hand-off |
| `cblock.c` / `cblock.h` | Fixed-size processing block management |

#### Filtering

| File | Description |
|------|-------------|
| `fir.c` / `fir.h` | FIR (Finite Impulse Response) filter core |
| `cfir.c` / `cfir.h` | Complex FIR filter |
| `icfir.c` / `icfir.h` | Interpolating complex FIR |
| `iir.c` / `iir.h` | IIR (Infinite Impulse Response) biquad filters |
| `bandpass.c` / `bandpass.h` | Overlap-save bandpass filter |
| `nbp.c` / `nbp.h` | Narrow bandpass (used for IF passband) |
| `firmin.c` / `firmin.h` | Minimum-phase FIR design |
| `eq.c` / `eq.h` | Graphic and parametric equalizer |
| `emph.c` / `emph.h` | Pre/de-emphasis for FM and speech |
| `doublepole.c` / `doublepole.h` | Two-pole resonator |
| `fcurve.c` / `fcurve.h` | Frequency response curve shaping |

#### Modulation & Demodulation

| File | Description |
|------|-------------|
| `amd.c` / `amd.h` | AM (synchronous) demodulator |
| `ammod.c` / `ammod.h` | AM modulator |
| `fmd.c` / `fmd.h` | FM demodulator |
| `fmmod.c` / `fmmod.h` | FM modulator |
| `matchedCW.c` / `matchedCW.h` | Matched filter for CW signal detection |

#### Noise Reduction & Blanking

| File | Description |
|------|-------------|
| `anr.c` / `anr.h` | Adaptive Noise Reduction (LMS-based) |
| `anf.c` / `anf.h` | Adaptive Notch Filter (removes heterodynes) |
| `emnr.c` / `emnr.h` | Spectral-subtraction noise reduction |
| `rnnr.c` / `rnnr.h` | RNNoise neural-network noise reduction (uses `lib/NR_Algorithms_x64/rnnoise.dll`) |
| `snb.c` / `snb.h` | Spectral Noise Blanking |
| `sbnr.c` / `sbnr.h` | Spectral Blanking + Noise Reduction |
| `nob.c` / `nob.h` | Noise blanker (pulse/impulse suppression) |
| `nobII.c` / `nobII.h` | Advanced noise blanker with improved detection |

#### Squelch & Gating

| File | Description |
|------|-------------|
| `amsq.c` / `amsq.h` | AM carrier squelch |
| `fmsq.c` / `fmsq.h` | FM sub-tone squelch |
| `ssql.c` / `ssql.h` | Spectral squelch (signal-to-noise based) |
| `vox.c` / `vox.h` | Voice Operated eXchange gate |

#### Level & Gain Control

| File | Description |
|------|-------------|
| `gain.c` / `gain.h` | Linear gain block |
| `compress.c` / `compress.h` | Speech compressor |
| `cfcomp.c` / `cfcomp.h` | Frequency-domain compressor |
| `meter.c` / `meter.h` | Signal level metering |
| `meterlog10.c` / `meterlog10.h` | Log-scale (dBm/dBFS) metering |

#### Frequency & Phase

| File | Description |
|------|-------------|
| `shift.c` / `shift.h` | Frequency shift (complex multiplier) |
| `resample.c` / `resample.h` | Sample rate conversion |
| `varsamp.c` / `varsamp.h` | Variable-ratio resampler |
| `delay.c` / `delay.h` | Fixed delay line |
| `slew.c` / `slew.h` | Slew rate limiter (anti-click) |
| `iqc.c` / `iqc.h` | IQ balance and phase/amplitude correction |
| `osctrl.c` / `osctrl.h` | Oscillator control |

#### Transmit Processing

| File | Description |
|------|-------------|
| `eer.c` / `eer.h` | Envelope Elimination and Restoration (for linear amplifiers) |
| `dexp.c` / `dexp.h` | Downward expander / noise gate for TX |
| `sender.c` / `sender.h` | Transmit output stage |
| `patchpanel.c` / `patchpanel.h` | Internal signal routing matrix |

#### Analysis & Utilities

| File | Description |
|------|-------------|
| `analyzer.c` / `analyzer.h` | FFT-based spectrum analyzer |
| `siphon.c` / `siphon.h` | Signal tap for scope/analyzer feeds |
| `gen.c` / `gen.h` | Test tone / noise generator |
| `calculus.c` / `calculus.h` | Pre-computed mathematical tables (large) |
| `FDnoiseIQ.c` / `FDnoiseIQ.h` | Frequency-domain IQ noise model (large) |
| `wisdom.c` / `wisdom.h` | FFTW wisdom cache (FFT plan optimization) |
| `lmath.c` / `lmath.h` | Large-array math operations |
| `cmath.c` / `cmath.h` | Complex arithmetic helpers |
| `gaussian.c` / `gaussian.h` | Gaussian window and distribution functions |
| `utilities.c` / `utilities.h` | General-purpose utilities |
| `rmatch.c` / `rmatch.h` | Reactive impedance matching |
| `apfshadow.c` / `apfshadow.h` | All-pass filter shadow buffer |

---

### ChannelMaster — Audio I/O Router

`Project Files/Source/ChannelMaster/`

A C/C++ DLL that owns all audio device I/O and routes streams between the radio hardware (via PortAudio or ASIO), the wdsp DSP engine, and the network.

#### Core Routing

| File | Description |
|------|-------------|
| `cmaster.c` / `cmaster.h` | Top-level audio routing, stream lifecycle, and device management |
| `cmasio.c` / `cmasio.h` | ASIO device control via the cmASIO wrapper |
| `cmbuffs.c` / `cmbuffs.h` | Shared circular buffers between audio threads and DSP |
| `router.c` / `router.h` | Configurable signal routing matrix |
| `patchpanel.c` | Cross-point switching between inputs and outputs |

#### Network Protocol

| File | Description |
|------|-------------|
| `network.c` / `network.h` | UDP network I/O for HPSDR protocol |
| `networkproto1.c` | HPSDR Protocol 1 packet framing and parsing |
| `netInterface.c` | Network interface selection and management |
| `tci.c` / `tci.h` | TCI (Thetis Control Interface) TCP protocol server |
| `pipe.c` / `pipe.h` | Named pipe for inter-process communication |
| `ring.c` / `ring.h` | Lock-free ring buffer for network streaming |

#### Audio Processing

| File | Description |
|------|-------------|
| `aamix.c` / `aamix.h` | Advanced audio mixer (multiple sources to one output) |
| `amix.c` / `amix.h` | Basic audio mixer |
| `sidetone.c` / `sidetone.h` | CW sidetone generation |
| `ivac.c` / `ivac.h` | VAC (Virtual Audio Cable) interface |
| `ilv.c` / `ilv.h` | Sample interleaving/de-interleaving |
| `bandwidth_monitor.c` / `bandwidth_monitor.h` | Measures and reports network bandwidth usage |

#### Control & Timing

| File | Description |
|------|-------------|
| `txgain.c` / `txgain.h` | Transmit gain staging |
| `vox.c` / `vox.h` | VOX detection and keying |
| `sync.c` / `sync.h` | Thread synchronization primitives |
| `nanotime.c` / `nanotimer.h` | High-resolution nanosecond timer |
| `obbuffs.c` / `obbuffs.h` | Output block buffers |
| `cmsetup.c` / `cmsetup.h` | Initialization and configuration |
| `cmUtilities.c` / `cmUtilities.h` | Helper functions |
| `version.c` / `version.h` | DLL version reporting |

---

### cmASIO — ASIO Audio Driver Wrapper

`Project Files/Source/cmASIO/`

A thin C++ DLL wrapping the Steinberg ASIO SDK. ChannelMaster loads it at runtime to access ASIO-compatible audio hardware (e.g., RME, Focusrite, dedicated SDR audio devices).

| File | Description |
|------|-------------|
| `hostsample.cpp` / `hostsample.h` | ASIO host implementation — device enumeration, buffer setup, callback |
| `dllmain.cpp` | DLL entry point |
| `pch.cpp` / `pch.h` | Pre-compiled header |
| `asiosdk_2.3.3_2019-06-14/` | Steinberg ASIO SDK 2.3.3 (included in-tree) |

---

### Midi2Cat — MIDI Controller Mapping

`Project Files/Source/Midi2Cat/`

A C# .NET library/application that intercepts MIDI controller messages and translates them to CAT commands sent to the Console.

| File | Description |
|------|-------------|
| `Midi2CatSetupForm.cs` | Configuration UI — assigns MIDI controls to CAT functions |
| `MidiMessageManager.cs` | MIDI device listener and event dispatcher |
| `Data/Database.cs` / `CatCmdDb.cs` | Persists MIDI-to-CAT mappings |
| `Data/MappedCommands.cs` | Runtime table of active MIDI→CAT bindings |
| `Data/ControllerMapping.cs` / `ControllerBinding.cs` | Data model for a single MIDI control binding |
| `Data/MidiDiag.cs` | MIDI message diagnostic/logging |
| `IO/MidiDevices.cs` / `MidiDevice.cs` | MIDI device enumeration via Windows Multimedia API |
| `IO/WinMM.cs` | P/Invoke bindings to `winmm.dll` (MIDI I/O) |
| `IO/MidiDeviceSetup.cs` | Per-device configuration dialog |

---

### RawInput — Raw Input Library

`Project Files/Source/RawInput/`

A C# library that captures keyboard and mouse input at the WM_INPUT level, bypassing standard WinForms focus rules. Used to allow hotkeys to work regardless of which control has focus.

| File | Description |
|------|-------------|
| `RawInput.cs` | Registers the raw input device and dispatches events |
| `RawKeyboard.cs` | Keyboard event decoder |
| `RawMouse.cs` | Mouse event decoder |
| `Win32.cs` | P/Invoke for `RegisterRawInputDevices` and `GetRawInputData` |
| `DataStructures.cs` | `RAWINPUT`, `RAWKEYBOARD`, `RAWMOUSE` struct definitions |
| `Enumerations.cs` | Virtual key codes and input flag enumerations |
| `KeyPressEvent.cs` / `MouseEvent.cs` | Typed event argument classes |
| `PreMessageFilter.cs` | WinForms message filter integration |
| `KeyMapper.cs` | Maps virtual key codes to readable key names |

---

### Thetis-Installer — Windows Installer

`Project Files/Source/Thetis-Installer/`

A WiX v3 project that packages the built binaries into a distributable MSI.

| File | Description |
|------|-------------|
| `Product.wxs` | Master installer definition — features, components, file list, registry entries |
| `Thetis_en-us.wxl` | English localization strings |
| `dotnet48.wxs` | .NET Framework 4.8 prerequisite check and bootstrapper |
| `Microsoft_VC14x_CRT_x64.msm` | Visual C++ runtime merge modules (VC142, VC143, VC145) included for all C++ DLLs |
| `MeterSkinInstaller.exe` | Bundled installer for meter skin package |
| `OpenHPSDR_Skins.exe` | Bundled installer for UI skins |

The AfterBuild target reads the `Thetis.exe` assembly version and renames the MSI to `Thetis-v<major>.<minor>.<build>.x64.msi`.

---

### External Libraries (Project Files/lib)

| Directory | Library | Used By |
|-----------|---------|---------|
| `fftw_x64/` | FFTW3 (double, single, long-double) | wdsp — all FFT operations |
| `NR_Algorithms_x64/` | rnnoise.dll, specbleach.dll | wdsp — `rnnr.c`, `sbnr.c` |
| `portaudio-19.7.0/` | PortAudio 19.7 | ChannelMaster — non-ASIO audio I/O |
| `SharpDX/` | SharpDX (Direct2D, D3D11, DXGI) | Console — hardware-accelerated display rendering |
| `SkiaSharp/` | SkiaSharp / libSkiaSharp | Console — 2D graphics and skin rendering |
