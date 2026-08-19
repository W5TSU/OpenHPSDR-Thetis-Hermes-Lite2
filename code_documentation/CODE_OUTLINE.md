# Thetis (Hermes-Lite 2) — Code Outline

Thetis is a Windows SDR console for OpenHPSDR-protocol radios; this fork adapts it for the
**Hermes-Lite 2 (HL2)**. The program is a four-layer stack: the **Console** (C#, WinForms) provides
all user interaction and radio control, and calls down via P/Invoke into **wdsp.dll** (C — all
signal processing), **ChannelMaster.dll** (C/C++ — audio and network I/O routing), and
**cmASIO.dll** (C++ — ASIO driver access). All paths below are relative to `Project Files/Source/`.

This outline was generated with the help of a [graphify](https://graphify.net) knowledge graph of
the source tree (15,653 nodes / 36,901 edges / 461 communities). See
[Exploring further](#exploring-the-code-with-the-knowledge-graph) at the end.

**Per-file documentation:** every file listed below has its own page under
[`code_documentation/files/`](files/README.md) with an outline of its classes, methods, and functions — each
with line number, signature, a short description, and how it is called — plus a graph-derived
summary of how the file is used (callers and callees). Regenerate those pages with
`python code_documentation/tools/gen_file_docs.py` after rebuilding the graph.

---

## Contents

Each entry links to its section below, plus a component-relationship diagram (`diagrams/`)
covering the files in that section — generated with the `diagram-design` plugin, architecture
type. Diagrams are grouped/simplified views for orientation, not a substitute for the file-role
tables or the knowledge graph.

1. [Application core and main window](#1-application-core-and-main-window) — [diagram](diagrams/01-application-core.svg)
2. [Settings and configuration](#2-settings-and-configuration) — [diagram](diagrams/02-settings-config.svg)
3. [HPSDR network protocol and radio discovery](#3-hpsdr-network-protocol-and-radio-discovery) — [diagram](diagrams/03-hpsdr-network.svg)
4. [Hermes-Lite 2 specific hardware control](#4-hermes-lite-2-specific-hardware-control) — [diagram](diagrams/04-hl2-hardware.svg)
5. [Spectrum, waterfall, and panadapter display](#5-spectrum-waterfall-and-panadapter-display) — [diagram](diagrams/05-spectrum-display.svg)
6. [DSP control from the console](#6-dsp-control-from-the-console) — [diagram](diagrams/06-dsp-control.svg)
7. [wdsp — the DSP engine](#7-wdsp--the-dsp-engine) — [diagram](diagrams/07-wdsp-engine.svg)
8. [ChannelMaster — audio and network routing](#8-channelmaster--audio-and-network-routing) — [diagram](diagrams/08-channelmaster.svg)
9. [Audio devices, VAC, and ASIO](#9-audio-devices-vac-and-asio) — [diagram](diagrams/09-audio-vac-asio.svg)
10. [CAT control and external program interfaces](#10-cat-control-and-external-program-interfaces) — [diagram](diagrams/10-cat-control.svg)
11. [CW keying](#11-cw-keying) — [diagram](diagrams/11-cw-keying.svg)
12. [MIDI control (Midi2Cat)](#12-midi-control-midi2cat) — [diagram](diagrams/12-midi-control.svg)
13. [Andromeda control surface](#13-andromeda-control-surface) — [diagram](diagrams/13-andromeda.svg)
14. [Metering](#14-metering) — [diagram](diagrams/14-metering.svg)
15. [Memories, band stacks, and the database](#15-memories-band-stacks-and-the-database) — [diagram](diagrams/15-memories-bandstacks.svg)
16. [DX spots and cluster display](#16-dx-spots-and-cluster-display) — [diagram](diagrams/16-dx-spots.svg)
17. [Thread-safe UI plumbing and shared controls](#17-thread-safe-ui-plumbing-and-shared-controls) — [diagram](diagrams/17-threadsafe-ui.svg)
18. [Raw keyboard/mouse input (RawInput)](#18-raw-keyboardmouse-input-rawinput) — [diagram](diagrams/18-rawinput.svg)

---

## 1. Application core and main window

The heart of the C# application: the main window that owns every subsystem, plus startup,
version, and OS-integration helpers.

| File | Role |
|------|------|
| `Console/console.cs` | The main window and central hub (~50k lines). Owns VFOs, band/mode/filter state, PTT/MOX sequencing, menus, and wires every other subsystem together. The graph's second-biggest god node (1,285 edges). |
| `Console/radio.cs` | Radio/receiver object model — bands, modes, filter presets per mode, and per-RX DSP state that the console manipulates. |
| `Console/enums.cs` | Shared enumerations (bands, modes, meter types, display modes, etc.) used across the whole console. |
| `Console/common.cs` | Grab-bag of static helpers: string/number formatting, control lookup, debugging aids, exception reporting. |
| `Console/keyboard.cs` | Static key-state helpers over `GetKeyState`/`GetAsyncKeyState` P/Invoke; the async variant backs the spacebar PTT-hold release detection (shortcut handling itself lives in `console.cs` `Console_KeyDown`). |
| `Console/titlebar.cs` | Custom title bar text/version display. |
| `Console/splash.cs`, `Console/progress.cs`, `Console/ucProgress.cs` | Startup splash screen and progress reporting during initialization. |
| `Console/VersionInfo.cs`, `Console/Versions.cs`, `Console/frmReleaseNotes.cs`, `Console/frmAbout.cs` | Version identification, release notes, and About box. |
| `Console/clsSingleInstance.cs` | Enforces a single running instance of Thetis. |
| `Console/win32.cs` | Win32 P/Invoke declarations (window messages, power management, multimedia timers). |
| `Console/hiperftimer.cs` | High-resolution performance timer used for timing-critical UI work. |
| `Console/clsDPISafeTools.cs` | High-DPI/monitor-scaling safety helpers for WinForms layout. |
| `Console/GlobalMouseHandler.cs`, `Console/clsTouchHandler.cs` | Application-wide mouse message filtering and touch-screen gesture support. |
| `Console/TimeOutTimerManager.cs` | Transmit time-out timer (limits continuous TX time). |
| `Console/frmLog.cs`, `Console/clsProgressLog.cs` | Diagnostic/status logging windows. |
| `Console/Firewall.cs` | Adds Windows Firewall rules so radio UDP traffic is not blocked. |
| `Console/Dumpcap.cs` | Drives Wireshark's `dumpcap` to capture radio network traffic for debugging. |
| `Console/clsDiscord.cs` | Discord rich-presence integration. |
| `Console/InputBox.cs`, `Console/frmFinder.cs` | Simple text-input dialog; searchable "find a setting" helper. |

## 2. Settings and configuration

Persistent user configuration: the giant Setup dialog, its database backing, and skinning.

| File | Role |
|------|------|
| `Console/setup.cs` | The Setup dialog (~90k lines) — every configurable option (radio, audio, display, DSP, CAT, VAC…) lives here. The graph's biggest god node (2,234 edges) and, per its cohesion scores, the prime refactoring candidate. |
| `Console/database.cs` | Settings persistence — reads/writes all option and control state to the database (XML/SQLite), including import/merge on upgrade. |
| `Console/clsDBMan.cs`, `Console/frmDBMan.cs` | Database manager: multiple named databases, backup/restore/switch between them. |
| `Console/Skin.cs`, `Console/clsThetisSkinService.cs` | UI skin loading and application (SkiaSharp-backed image skins for console controls). |
| `Console/clsHardwareSpecific.cs` | Per-hardware-model capability flags and defaults (which options apply to which radio model, incl. HL2). |
| `Console/clsCMASIOConfig.cs` | Configuration UI/state for the cmASIO ASIO driver connection. |
| `Console/xvtr.cs` | Transverter band setup (frequency offsets, power limits per transverter band). |
| `Console/TuneStep.cs`, `Console/ucTunestepOptionsGrid.cs` | Tuning-step definitions and their configuration grid. |
| `Console/clsLegacyItemController.cs` | Maps legacy/renamed UI items to their current equivalents so old databases and skins keep working. |

## 3. HPSDR network protocol and radio discovery

Finding the radio on the network and exchanging Protocol-1 UDP traffic with it. The C# side
configures and supervises; the actual packet pumps live in ChannelMaster (§8).

| File | Role |
|------|------|
| `Console/HPSDR/NetworkIO.cs` | High-level radio session control: init/start/stop, VFO frequency-to-phase-word conversion, sample rate, and control-register updates to the radio. |
| `Console/HPSDR/NetworkIOImports.cs` | The P/Invoke surface into ChannelMaster's network code — every `extern` for radio I/O (128 imports). |
| `Console/HPSDR/clsRadioDiscovery.cs` | UDP broadcast discovery of HPSDR radios on all NICs; produces the discovered-radio list. |
| `Console/clsDiscoveredRadioPicker.cs`, `Console/frmAddCustomRadio.cs`, `Console/ucRadioList.cs` | UI for picking among discovered radios and defining custom/static radio addresses. |
| `Console/HPSDR/specHPSDR.cs` | Configures the wdsp spectrum analyzer instances for HPSDR data streams. |
| `Console/HPSDR/Alex.cs` | Alex RF filter board control (antenna and band-filter relay selection). Retained from upstream; antenna switching from the console is disabled in this HL2 fork. |
| `Console/HPSDR/Penny.cs` | PennyLane/Penelope open-collector output and mic-gain control by band. |
| `Console/NetworkThrottle.cs` | Network send-rate throttling to smooth UDP bursts. |
| `Console/frmSeqLog.cs` | Sequence-error log window — shows dropped/out-of-order UDP packet statistics. |

## 4. Hermes-Lite 2 specific hardware control

The HL2 additions this fork exists for.

| File | Role |
|------|------|
| `Console/HPSDR/IoBoardHl2.cs` | Register-level control of the HL2 I/O board: antenna tuner (ATU) commands and status, TX frequency bytes sent to the board, fault detection, and control-register read/write over the HL2's i2c-style interface. |

Supporting pieces elsewhere: `clsHardwareSpecific.cs` (§2) carries HL2 model quirks;
`NetworkIO.cs`/`NetworkIOImports.cs` (§3) carry the HL2's gateware control bits inside Protocol-1
command words.

## 5. Spectrum, waterfall, and panadapter display

Everything drawn on the panadapter: FFT data collection, rendering, and wideband views.

| File | Role |
|------|------|
| `Console/display.cs` | The spectrum/waterfall renderer — SharpDX (Direct2D/D3D11) drawing of panadapter, waterfall, band edges, notches, cursors, and TX filter overlays. |
| `Console/clsSpectrumProcessor.cs` | Pulls pixel-ready FFT data from the wdsp analyzer and post-processes it (averaging, peak/blend detection) for the display. |
| `Console/PanDisplay.cs` | Panadapter display user control hosting the render surface, mouse tuning, and zoom/pan interaction. |
| `Console/wbDisplay.cs`, `Console/wideband.cs` | Wideband (full 0–61 MHz) spectrum display and its data acquisition from the radio's wideband sample stream. |
| `Console/Path_Illustrator.cs` | Interactive block diagram of the whole signal path (what's enabled where, RX/TX routing). |
| `Console/N1MM.cs` | Streams spectrum display data over UDP to the N1MM+ logger's spectrum window. |
| `Console/ucVARGrapher.cs` | Small graphing control (used for VAC variable-rate resampler diagnostics). |
| `Console/ucUnderOverFlowWarningViewer.cs` | Audio buffer underflow/overflow warning indicator. |

## 6. DSP control from the console

C# classes that configure the wdsp processing chains — the "knobs" side of the DSP.

| File | Role |
|------|------|
| `Console/dsp.cs` | Central DSP settings hub: creates wdsp RX/TX channels and pushes every DSP parameter (NR, NB, AGC, filters, TX processing) down via P/Invoke. |
| `Console/rxa.cs`, `Console/rxaControls.cs` | Typed wrappers for wdsp RXA (receiver chain) settings and the UI controls bound to them. |
| `Console/filter.cs`, `Console/FilterForm.cs`, `Console/frmFilterManager.cs` | RX filter preset model per mode, the filter-edit form, and the filter-set manager. |
| `Console/frmBandwidth.cs`, `Console/ucBandwidthView.cs` | Variable-bandwidth adjustment popup and its graphical bandwidth view. |
| `Console/frmNotchPopup.cs` | Manual notch filter add/edit popup (backed by wdsp `nbp`). |
| `Console/eqform.cs`, `Console/ucParametricEq.cs` | RX/TX graphic and parametric equalizer forms (backed by wdsp `eq.c`). |
| `Console/frmCFCConfig.cs` | Continuous Frequency Compressor (CFC) TX processing configuration (backed by wdsp `cfcomp.c`). |
| `Console/PSForm.cs`, `Console/AmpView.cs` | PureSignal TX linearization control panel and the amplifier gain/phase view (backed by wdsp `calcc.c`/`iqc.c`). |
| `Console/DiversityForm.cs` | Two-receiver diversity reception control (phase/gain mixing of RX1/RX2). |
| `Console/cmaster.cs` | P/Invoke wrapper for ChannelMaster: channel setup, audio mixer (`aamix`), radio protocol selection, and stream control from C#. |

## 7. wdsp — the DSP engine

Pure C library (`wdsp/`); every audio sample and spectrum pixel passes through it. Created by
Warren Pratt NR0V. The console talks to it through `dsp.cs`/`specHPSDR.cs`.

**Chain assembly and infrastructure**

| File | Role |
|------|------|
| `RXA.c` / `TXA.c` | Define the complete receive and transmit DSP graphs — every block below is instantiated and ordered here. |
| `channel.c`, `main.c` | Channel object lifecycle (create/destroy/run) and DLL entry points. |
| `iobuffs.c`, `syncbuffs.c`, `cblock.c` | Sample buffering between the audio callback world and DSP blocks. |
| `utilities.c`, `cmath.c`, `lmath.c`, `calculus.c`, `gaussian.c`, `fcurve.c`, `meterlog10.c`, `zetaHat.c` | Shared math: aligned allocation (`malloc0`, a god node with 189 edges), complex math, interpolation, statistics. |
| `wisdom.c`, `impulse_cache.c` | FFTW wisdom generation/caching and FIR impulse-response caching for fast startup. |
| `patchpanel.c`, `gain.c`, `shift.c`, `slew.c`, `delay.c` | Signal patch/routing points, gain staging, frequency shift, envelope slewing, delays. |
| `meter.c` | Signal level metering taps feeding the console's meters. |
| `version.c` | Library version export. |

**Filtering**

| File | Role |
|------|------|
| `fir.c`, `firmin.c`, `cfir.c`, `icfir.c` | FIR filter design and fast-convolution (overlap-save) filtering, including CIC-compensation filters. |
| `iir.c`, `doublepole.c` | IIR biquad sections (notches, peaking filters) and double-pole building blocks. |
| `bandpass.c`, `nbp.c` | Main bandpass filter and the notched-bandpass (auto/manual notch database) filter. |
| `eq.c` | Graphic/parametric equalizer. |
| `emph.c` | FM pre-/de-emphasis. |
| `resample.c`, `varsamp.c`, `rmatch.c` | Fixed and variable-ratio resamplers, and the adaptive rate-matcher that reconciles independent sample clocks. |

**Demodulation, modulation, and squelch**

| File | Role |
|------|------|
| `amd.c`, `fmd.c` | AM/SAM (synchronous) and FM demodulators. |
| `ammod.c`, `fmmod.c` | AM and FM modulators for TX. |
| `eer.c` | Envelope elimination and restoration (polar) TX processing. |
| `amsq.c`, `fmsq.c`, `ssql.c` | AM squelch, FM squelch, and syllabic (voice-detecting) squelch. |

**Digital voice (FreeDV)** — HL2 fork addition, `FreeDV` branch

| File | Role |
|------|------|
| `fdv.c` | FreeDV 700E RX decode block. Sits post-AGC in the RXA chain; resamples to/from the modem's 8 kHz rate, normalises blocks into `libcodec2`'s 16-bit domain via a smoothed AGC, drives `freedv_rx()` per `freedv_nin()`-sized block, and passes raw modem audio through until synced/primed so the signal stays audible for tuning. RADE V1's equivalent decode block lives in ChannelMaster (`radae.c`, §8) rather than here, since it uses a separate native library (`rade_c`) instead of `libcodec2`. |

**Noise reduction and blanking**

| File | Role |
|------|------|
| `anr.c`, `anf.c` | Legacy LMS adaptive noise reduction (NR) and automatic notch filter. |
| `emnr.c` | Spectral noise reduction "NR2" (MMSE-based). |
| `rnnr.c` | RNNoise neural-network noise reduction "NR3" (uses `lib/NR_Algorithms_x64`). |
| `sbnr.c` | libspecbleach spectral noise reduction "NR4". |
| `nob.c`, `nobII.c` | Impulse noise blankers (NB and NB2). |
| `snb.c` | Spectral noise blanker. |
| `FDnoiseIQ.c` | Frequency-domain I/Q noise processing. |

**Levels, AGC, and TX processing**

| File | Role |
|------|------|
| `wcpAGC.c` | The WDSP AGC (receive gain control and TX leveler). |
| `compress.c`, `cfcomp.c` | TX speech compressor and continuous frequency compressor. |
| `dexp.c` | Downward expander / noise gate with VOX tie-in. |
| `osctrl.c` | TX overshoot control. |
| `siphon.c` | Taps TX samples out of the chain (e.g., for the TX display). |

**PureSignal and diversity**

| File | Role |
|------|------|
| `calcc.c`, `iqc.c` | PureSignal calibration calculation and the I/Q correction applied to TX. |
| `div.c` | Diversity combiner (mixes two receivers with adjustable gain/phase). |

**Spectrum analyzer and misc**

| File | Role |
|------|------|
| `analyzer.c` | The multi-instance FFT spectrum analyzer behind every panadapter/waterfall. |
| `sender.c` | Sends DSP data (spectrum, audio taps) back toward the console. |
| `gen.c` | Signal generators (tone, two-tone, noise, sweep) for testing and tune. |
| `matchedCW.c`, `apfshadow.c` | Matched CW filtering and audio peaking filter support. |

## 8. ChannelMaster — audio and network routing

C/C++ DLL (`ChannelMaster/`) that owns real-time I/O: it moves samples between the radio (UDP),
wdsp, the sound card, and virtual audio cables. The graph confirmed it calls directly into wdsp's
resamplers and allocators (`aamix.c → resample.c/utilities.c`).

| File | Role |
|------|------|
| `cmaster.c` | Stream lifecycle and top-level orchestration: creates channels, starts/stops audio and network streams. |
| `cmsetup.c` | System-wide setup: instantiates buffers, mixers, VAC, analyzers per radio model. |
| `router.c` | The signal routing matrix — which input feeds which DSP channel and which output. |
| `network.c`, `networkproto1.c`, `netInterface.c` | HPSDR Protocol-1 UDP implementation: socket setup, packet build/parse, EP2/EP4/EP6 endpoint handling, sequence tracking. |
| `aamix.c`, `amix.c` | Audio mixers (monitor mix, multi-RX audio combination) with per-input gain and slew. |
| `ivac.c` | Virtual Audio Cable engine — PortAudio streams with variable-ratio resampling between Thetis and other PC apps. |
| `cmasio.c` | Bridge to the cmASIO DLL for direct ASIO device I/O. |
| `tci.c` | TCI (Transceiver Control Interface) TCP server for SDC/logger integration at the audio layer. |
| `sidetone.c` | CW sidetone generation. |
| `vox.c`, `dexp` hooks, `txgain.c` | VOX detection and TX gain staging. |
| `analyzers.c` | Attaches wdsp spectrum analyzers to ChannelMaster streams (RX/TX displays). |
| `cmbuffs.c`, `obbuffs.c`, `ring.c`, `pipe.c`, `sync.c`, `ilv.c` | Ring buffers, output buffers, thread synchronization, and sample interleaving plumbing. |
| `bandwidth_monitor.c`, `nanotime.c` | Network bandwidth statistics and high-resolution timestamps. |
| `pro.c`, `zeer.c`, `znob.c`, `znobII.c` | Auxiliary DSP experiments retained from upstream (protocol processing, zero-delay EER, noise blanker variants). |
| `cmUtilities.c`, `version.c` | Shared helpers and version export. |
| `radae.c` | RADE V1 neural-mode RX decode block (HL2 fork addition, `FreeDV` branch). RX-only: hooked into `pipe.c`'s `xpipe()` hot path, drives the `rade_c`/`rade.lib` native decoder plus `lpcnet`/`fargan` speech synthesis, gated by `RXRadaeEnabled` (CAT `ZZDW`/`ZZDZ`, Setup → DSP → FreeDV tab). |
| `radae_micdsp.c` | Mic-path DSP helpers (biquad EQ, RNNoise, EBU R128 loudness normalisation, AGC) prepared for a future RADE V1 TX path — not yet wired to `xradae_tx` as of this branch. |
| `r8brain_wrap.cpp` | C-callable wrapper around the vendored r8brain-free-src `CDSPResampler24`, used by `radae.c` for its own internal sample-rate conversion. |

## 9. Audio devices, VAC, and ASIO

Sound-card management on the C# side, plus the native ASIO wrapper.

| File | Role |
|------|------|
| `Console/audio.cs` | Audio device enumeration and configuration: sample rates, buffer sizes, VAC on/off, device selection; drives ChannelMaster accordingly. |
| `Console/portaudio.cs` | P/Invoke wrapper for PortAudio device/host-API enumeration. |
| `Console/ivac.cs` | P/Invoke wrapper for ChannelMaster's VAC engine (`ivac.c`). |
| `Console/BasicAudio.cs` | Simple WAV playback (beeps/announcements) outside the DSP path. |
| `Console/clsAudioRecordPlayback.cs` | RX/TX audio and I/Q recording and playback (wave capture of what you hear/transmit). |
| `Console/ringbuffer.cs` | Managed ring buffer used by audio record/playback. |
| `cmASIO/dllmain.cpp`, `cmASIO/hostsample.cpp`, `cmASIO/version.cpp` | The cmASIO DLL: thin host wrapper around the Steinberg ASIO SDK giving ChannelMaster direct ASIO driver access. The bundled `asiosdk_2.3.3.../` tree is the vendored Steinberg SDK. |

## 10. CAT control and external program interfaces

Kenwood-style CAT over serial and TCP, plus modern TCI — how loggers and digimode apps control
Thetis.

| File | Role |
|------|------|
| `Console/CAT/CATCommands.cs` | Implements the CAT command set (ZZxx extended + Kenwood TS-2000 subset) — 399-edge god node touching most console state. |
| `Console/CAT/CATParser.cs` | Tokenizes/validates incoming CAT strings and dispatches to `CATCommands`. |
| `Console/CAT/SDRSerialPortII.cs`, `Console/CAT/SIOListenerII.cs` | Serial-port wrapper and the per-port listener threads (CAT, PTT, keyer ports). |
| `Console/CAT/TCPIPcatServer.cs` | CAT over TCP/IP server (multiple client connections). |
| `Console/CAT/SerialPortPTT.cs`, `Console/CAT/UsbBCDCable.cs` | PTT via serial control lines; band-decoder output over a USB BCD cable. |
| `Console/CAT/JustinIO.cs` | Low-level Win32 serial I/O used beneath the serial classes. |
| `Console/CAT/SerialRxEvent.cs`, `Console/CAT/CATTester.cs` | Serial receive event plumbing and an interactive CAT test window. |
| `Console/clsCATMessageQueue.cs`, `Console/clsCatAtonic.cs` | Queued/asynchronous CAT message handling and scripted ("atomic") CAT command sequences. |
| `Console/TCIServer.cs` | TCI WebSocket server (protocol used by SDC, LogHX, etc.): exposes VFOs, modes, spots, and audio to TCI clients. |

## 11. CW keying

| File | Role |
|------|------|
| `Console/cwkeyer.cs` | The CW keyer: iambic keying, key/paddle input sources (serial, radio, MIDI), break-in timing, and keying the radio. |
| `Console/cwx.cs` | CWX memory keyer window — canned messages, beacon loops, and keyboard CW. |
| `Console/cwedit.cs` | Editor for CWX stored messages. |
| `Console/CW/CWInput.cs` | Abstraction over CW key input sources (which line/device is dot, dash, PTT). |

(Signal-side CW: sidetone in `ChannelMaster/sidetone.c`, matched CW filter in `wdsp/matchedCW.c`.)

## 12. MIDI control (Midi2Cat)

Separate C# assembly (`Midi2Cat/`) mapping MIDI controllers to console functions, bridged into the
console by `Console/Midi2CatCommands.cs`.

| File | Role |
|------|------|
| `Console/Midi2CatCommands.cs` | The bridge: exposes console operations (tune, volume, filters, PTT…) as commands a MIDI control can bind to (256-edge god node). |
| `Midi2Cat/Midi2Cat.IO/MidiDevice.cs`, `MidiDevices.cs`, `WinMM.cs` | MIDI device open/close and message receive/send over the Windows Multimedia (winmm) API. |
| `Midi2Cat/MidiMessageManager.cs` | Routes incoming MIDI messages to their mapped commands. |
| `Midi2Cat/Midi2Cat.Data/*.cs` (`Database.cs`, `ControllerMapping.cs`, `ControllerBinding.cs`, `CatCmdDb.cs`, `MappedCommands.cs`, `Enums.cs`) | Persistence and object model for controller-to-command mappings. |
| `Midi2Cat/Midi2CatSetupForm.cs`, `Midi2Cat.IO/MidiDeviceSetup.cs`, dialogs (`PickDialog`, `LoadDialog`, `SaveAsDialog`, `OrganiseDialog`) | Mapping editor UI: wiggle a control, pick a function. |
| `Midi2Cat/Helpers/ControlHelpers.cs`, `Midi2Cat.Data/MidiDiag.cs` | UI helpers and MIDI diagnostics. |

## 13. Andromeda control surface

Touch/hardware front-panel support (Apache Labs Andromeda and similar G2 panels) mirroring main
console functions.

| File | Role |
|------|------|
| `Console/Andromeda/Andromeda.cs` | Core handler: decodes Andromeda encoder/button CAT messages and applies them to the console; manages button-bar text feedback. |
| `Console/Andromeda/AndromedaEditForm.cs` | Editor for assigning functions to Andromeda encoders and buttons. |
| `Console/Andromeda/BandButtonsPopup.cs`, `FilterButtonsPopup.cs`, `ModeButtonsPopup.cs` | On-screen popups for band/filter/mode selection from the panel. |
| `Console/Andromeda/vfosettingspopup.cs`, `displaysettingsform.cs`, `ModeDependentSettingsForm.cs`, `SliderSettingsForm.cs` | Panel-oriented quick-settings popups (VFO, display, per-mode, slider assignments). |
| `Console/frmMacroButtonConfig.cs`, `Console/ucOtherButtonsOptionsGrid.cs` | User-programmable macro buttons and their configuration grid. |

## 14. Metering

| File | Role |
|------|------|
| `Console/MeterManager.cs` | The metering subsystem (~30k lines): collects signal/power/SWR/voltage readings, renders configurable meter faces (analog, bar, LED) via DirectX, and manages multiple meter containers. |
| `Console/ucMeter.cs`, `Console/frmMeterDisplay.cs` | The meter user control and the floating multi-meter display window. |
| `Console/clsMeterScriptEngine.cs` | Scripting engine for user-defined meter faces/behaviors. |
| `Console/ucOCLedStrip.cs`, `Console/ucSignalSelect.cs`, `Console/ucLGPicker.cs`, `Console/ucGradientDefault.cs` | Meter-related picker controls (open-collector LED strip, signal source, linear-gradient color pickers). |

## 15. Memories, band stacks, and the database

| File | Role |
|------|------|
| `Console/Memory/MemoryForm.cs`, `MemoryList.cs`, `MemoryRecord.cs` | Memory channel list UI and its record/list model (frequency, mode, filter, tones per memory). |
| `Console/Memory/DXMemList.cs`, `DXMemRecord.cs` | Memory list variant used for DX cluster spot storage. |
| `Console/clsBandStackManager.cs`, `Console/frmBandStack2.cs` | Per-band frequency stack (last-used frequencies per band) and its popup window. |
| `Console/frmQuickRecallPopupList.cs`, `Console/ucQuickRecall.cs` | Quick recall (recent frequencies) list. |
| `Console/Channel.cs` | Simple channel object used by scanning/memory features. |
| `Console/SortableBindingList.cs` | Sortable list base used by memory grids. |

(`database.cs`/`clsDBMan.cs` in §2 provide the storage these features persist into.)

## 16. DX spots and cluster display

| File | Role |
|------|------|
| `Console/SpotManager2.cs` | DX spot store and on-panadapter spot rendering (callsigns on the spectrum); fed by TCI and cluster sources. |
| `Console/clsCountryData.cs` | DXCC country/prefix lookup for spot flag/bearing data. |
| `Console/clsFlagAtlas.cs`, `Console/clsImgeFetcher.cs` | Country flag atlas and web image fetching (e.g., QRZ pictures). |

## 17. Thread-safe UI plumbing and shared controls

DSP/audio/network threads must update WinForms controls safely; these wrappers marshal to the UI
thread. `ts` suffix = thread-safe.

| File | Role |
|------|------|
| `Console/Invoke/invoke.cs` | Core invoke helpers — run any control update on the UI thread. |
| `Console/Invoke/*ts.cs` (`buttonts`, `checkboxts`, `comboboxts`, `groupboxts`, `labelts`, `numericupdownts`, `panelts`, `radiobuttonts`, `textboxts`, `trackbarts`) | Thread-safe subclasses of the standard WinForms controls, used everywhere in the console/setup UI. |
| `Console/PrettyTrackBar.cs`, `Console/ColorButton.cs` | Custom-drawn slider and color-picker button used across forms. |
| `Console/ucInfoBar.cs`, `Console/frmInfoBarPopup.cs` | The info bar (status/warning strip) and its popup. |
| `Console/frmIPv4Picker.cs`, `Console/frmSerialPortPicker.cs`, `Console/frmVariablePicker.cs` | Small shared picker dialogs. |
| `Console/ShutdownForm.cs` | Orderly-shutdown progress window. |
| `Console/RAForm.cs` | Radio-astronomy data collection utility (niche feature retained from upstream). |

## 18. Raw keyboard/mouse input (RawInput)

Separate assembly (`RawInput/`) using the Win32 Raw Input API so specific devices (e.g., a
dedicated USB keypad) can drive the radio even when Thetis isn't focused.

| File | Role |
|------|------|
| `RawInput/RawInput.cs` | Entry point: registers for raw input and dispatches device events. |
| `RawInput/RawKeyboard.cs`, `RawInput/RawMouse.cs` | Per-device keyboard and mouse message decoding. |
| `RawInput/KeyMapper.cs`, `KeyPressEvent.cs`, `MouseEvent.cs`, `RawInputEventArg.cs` | Key mapping and event argument types. |
| `RawInput/Win32.cs`, `DataStructures.cs`, `Enumerations.cs`, `RegistryAccess.cs`, `PreMessageFilter.cs` | Win32 interop declarations and device-name registry lookup. |

---

## Exploring the code with the knowledge graph

The graphify outputs live in `graphify-out/` at the repository root:

- **`graph.html`** — interactive graph, open in a browser (aggregated to 434 community nodes).
- **`GRAPH_REPORT.md`** — full audit: god nodes, per-community listings, cohesion scores.
- **`graph.json`** — raw graph data (15,492 nodes / 36,536 edges), GraphRAG-ready.

Useful commands (run from the repository root):

```bash
graphify query "How does the console send TX frequency to the HL2 I/O board?"
graphify explain "IoBoardHl2"
graphify path "CATCommands" "NetworkIO"
graphify update "Project Files/Source"   # incremental rebuild after code changes
```

Notable graph findings:

- **God nodes** (most-connected abstractions): `Setup` (2,234 edges), `Console` (1,285),
  `CATCommands` (399), `TCPIPtciSocketListener` (346), `MidiDevice` (277), `WDSP` (261).
- **Cross-layer coupling**: ChannelMaster links directly against wdsp internals
  (`aamix.c` calls `resample.c`'s and `utilities.c`'s functions; `ivac.c` calls `rmatch.c`) —
  the two native DLLs are not independent.
- **No import cycles** were detected between modules.
