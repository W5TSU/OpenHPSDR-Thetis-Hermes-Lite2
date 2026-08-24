# Radio Documentation

Vendor and third-party reference manuals for the hardware and software this fork runs on —
kept here as a local archive so they're available offline and versioned alongside the code that
depends on them. These are unmodified copies of external documents (Apache Labs, HPSDR project,
FlexRadio Systems, individual authors); none of them are generated or maintained by this repo.

| Document | Contents |
|----------|----------|
| [`Thetis manual_1.1.pdf`](Thetis%20manual_1.1.pdf) | The general **THETIS User Manual** (Rev 1.1, edited by Laurence Barker, G8NJJ, for the HPSDR project). Start here for Thetis itself: screen layout, the RX/TX signal processing chain, installation, voice/CW/digital-mode operation, the full console-control reference, and the "other forms" (bandstacks, memory, equalizer, PureSignal, diversity, DX spotting, radio astronomy, wideband display). |
| [`Thetis-CAT-Command-Reference-Guide-V3.pdf`](Thetis-CAT-Command-Reference-Guide-V3.pdf) | **Thetis and PowerSDR 3.x CAT Command Reference Guide** (developed by BobT, K5KDN; updated for 3.x by G8NJJ). Every standard Kenwood-style and `ZZxx` extended CAT command, organized by functional group (VFO, filters, noise rejection, metering, digital modes, TX audio, antennas, mixer control, and more). The canonical reference for this fork's own CAT work — see `.claude/skills/thetis-control/SKILL.md`. |
| [`Thetis Network Settings_0.2.pdf`](Thetis%20Network%20Settings_0.2.pdf) | **Thetis Network Settings** (MW0LGE, 2026, v0.2). A focused deep-dive on Setup → H/W Select's networking/discovery controls: discovery speed, NIC selection (all NICs vs. a specific one), subnet handling, any-radio vs. specific-radio targeting, and the discovery listen-port options (random vs. fixed, P2-only port negotiation). |
| [`WDSP Guide, Rev 1.23.pdf`](WDSP%20Guide%2C%20Rev%201.23.pdf) | **The WDSP Guide — Using WDSP for Software Developers** (Dr. Warren C. Pratt, NR0V). The developer-facing reference for the `wdsp` DSP library this fork's `Project Files/Source/wdsp/` vendors: the channel API, and a block-by-block walkthrough of the `RXA` receive chain and `TXA` transmit chain (filters, AGC, noise reduction, squelch, equalizer, CESSB, PureSignal predistortion, panadapter/display senders, and more). |
| [`cmASIO Guide.pdf`](cmASIO%20Guide.pdf) | Quickstart for **cmASIO**'s optional direct-ASIO audio routing feature — an advanced, opt-in path that lets ChannelMaster sink/source audio via a hardware ASIO device instead of the radio's own codec. Covers the registry-key setup (`ASIOdrivername`), the ring-buffer size and low-latency "lockMode" tuning knobs, and the hardware/firmware prerequisites (Protocol-2 only, a real hardware ASIO driver, a shared clock source). |
| [`Hermes Lite 2 Thetis Installation and 3rd Party Apps V8 05142024.pdf`](Hermes%20Lite%202%20Thetis%20Installation%20and%203rd%20Party%20Apps%20V8%2005142024.pdf) | **Hermes Lite 2 and Thetis Installations and Configurations** (N8SDR/MI0BOT, rev. 2024-05-14, v9), 70 pages. Walks through installing and configuring Thetis for the HL2 and remote-connecting to it, then gives worked examples for ~10 popular third-party digital-mode apps (WSJT-X, JTDX, MSHV, JS8Call, VarAC, FLDIGI) and logging programs (N3FJP ACLog, LogHX, Log4OM, N1MM) — plus a grab-bag of operating tips (SWL filter settings, CW keyer setup, SparkSDR firmware updates). |
| [`26-October-2017-ANAN-7000DLE-User-Guide.pdf`](26-October-2017-ANAN-7000DLE-User-Guide.pdf) | Apache Labs' **ANAN-7000DLE User Guide** (26 Oct 2017) — the original ANAN-7000DLE (100W HF/6M) hardware manual: specifications, network setup, PowerSDR mRX PS installation, SSB/CW operation, linear-amplifier and transverter connections, and PureSignal adaptive predistortion. Not HL2-specific, but the same PowerSDR/Thetis-family operating concepts apply. |
| [`7000DLE MKII User guide.pdf`](7000DLE%20MKII%20User%20guide.pdf) | Apache Labs' **ANAN-7000DLE MKII User Guide** (26 Feb 2019) — the MKII hardware revision of the guide above (adds the embedded i5/i7 front-panel quick-start, otherwise the same structure/scope). |
| [`HPSDRProgrammer_HPSDRBootloader-User-Guide.pdf`](HPSDRProgrammer_HPSDRBootloader-User-Guide.pdf) | Apache Labs' guide to **HPSDRProgrammer** (v2.0.4.10) and **HPSDRBootloader** (v2.0) — the two Windows tools for flashing/updating an HPSDR-protocol radio's firmware over Ethernet, and for assigning the radio a static IP address. Relevant any time HL2 gateware needs updating; see also `code_documentation/Protocols.md` for how this fork's own network layer talks to that firmware. |

## Where these fit alongside this repo's own docs

These are third-party references, not this fork's own documentation — for how the *code* actually
works, see `code_documentation/` (architecture, per-file reference) and `Documentation/` (this
fork's own guides: FreeDV/RADE, audio pathway, station audio config). Reach for a PDF here when the
question is about vendor hardware behavior, the CAT protocol as a standalone spec, or a third-party
app's own setup — reach for the rest of the repo's docs when the question is about what this
codebase does.
