# HPSDR Network Protocols 1 and 2 — How Thetis Talks to the Radio

Thetis controls every supported radio over UDP/IP using one of two openHPSDR
network protocols. This document explains both protocols, where each is
implemented in this repository, how the protocol relates to the hardware, and
what that means for the Hermes-Lite 2.

Naming: the code and UI use several aliases for the same two things —

| Protocol | Also called | `RadioProtocol` enum (`enums.cs:437`) |
|----------|-------------|----------------------------------------|
| Protocol 1 | "P1", "USB over UDP", "Metis protocol" | `USB = 0` |
| Protocol 2 | "P2", "openHPSDR Ethernet protocol" | `ETH = 1` |

The "USB" name is historical: Protocol 1 wraps the 512-byte USB frames of the
original Atlas/Ozy USB radios in UDP packets, so the same firmware framing
survived the move to Ethernet (the Metis board).

---

## 1. Protocol 1 — everything through one port

Implemented in `Project Files/Source/ChannelMaster/networkproto1.c` (data
path) and `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs` (discovery).

- **Transport**: all traffic runs between the host and a single radio UDP
  port (1024 by default, `base_outbound_port` in `network.h:64`).
- **Framing**: each UDP payload is 1032 bytes — an 8-byte header
  (`0xEF 0xFE 0x01`, endpoint, sequence) plus **two 512-byte "USB frames"**.
  Every USB frame starts with the sync pattern `0x7F 0x7F 0x7F` followed by
  **5 command & control (C&C) bytes** `C0–C4`, then multiplexed payload
  (`WriteMainLoop`, `networkproto1.c:593`). There are no separate streams:
  receiver IQ, mic audio, speaker audio, TX IQ and all control bits share
  these frames.
- **C&C**: `C0` selects which register the remaining 4 bytes address; the host
  cycles through registers continuously (frequency, attenuation, filters…),
  and the radio returns status (PTT/dot/dash, ADC overflow, power, version)
  in its own C&C fields.
- **Discovery**: host broadcasts `0xEF 0xFE 0x02`; the radio replies with its
  MAC, firmware code version (byte 9) and **board ID (byte 10)** —
  parsed in `clsRadioDiscovery.cs:1159` and mapped by `mapP1DeviceType()`
  (`clsRadioDiscovery.cs:1236`).
- **Sample rates**: 48 / 96 / 192 kHz, plus 384 kHz for radios whose gateware
  supports it — in this fork the Hermes-Lite 2 and Red Pitaya
  (`setup.cs:852`, `InitAudioTab`). One rate applies to all receivers.

## 2. Protocol 2 — a port per function

Implemented in `Project Files/Source/ChannelMaster/network.c`. Instead of
multiplexing, every function gets its own UDP port and its own packet format,
which allows more receivers, independent per-DDC sample rates up to 1536 kHz,
and 24-bit samples.

Port map as implemented (offsets from the base port, normally 1024/1025 —
`p2_custom_port_base`, `network.c:108`):

**Host → radio** (command builders, `network.c`):

| Port | Function | Code |
|------|----------|------|
| 1024 | General packet (rates, options, every ~250 ms) | `CmdGeneral()` `network.c:812` |
| 1025 | DDC (receiver) configuration | `CmdRx()` `network.c:1057` |
| 1026 | DUC (transmitter) configuration | `CmdTx()` `network.c:1172` |
| 1027 | High-priority C&C (frequency, PTT, attenuation…) | `CmdHighPriority()` `network.c:904` |
| 1028+ | DUC TX IQ / speaker audio data | `sendOutbound()` `network.c:1241` |

**Radio → host** (dispatch by source port, `ReadUDPFrame()` `network.c:472`
and `ReadThreadMainLoop()` `network.c:636`):

| Port | Function |
|------|----------|
| 1025 | High-priority status: PTT/dot/dash, PLL lock, ADC overload, exciter/FWD/REV power, supply volts, user ADC & digital inputs |
| 1026 | Mic samples (16-bit, 48 ksps) |
| 1027–1034 | Wideband raw ADC snapshots (full-spectrum display) |
| 1035–1041 | DDC0–DDC6 IQ data (24-bit) |

- **Discovery**: host sends `0x00 0x00 0x00 0x00 0x02`; the reply carries the
  board ID (byte 11), the **protocol version supported** (byte 12 — stored as
  `Protocol2Supported`, `clsRadioDiscovery.cs:1215`), firmware version and a
  beta/sub-version byte.
- **Sample rates**: 48/96/192/384/768/1536 kHz (`setup.cs:858`).

## 3. How Thetis picks a protocol

There is no manual protocol switch — **the radio tells Thetis what it speaks**:

1. Discovery (`clsRadioDiscovery.cs`) broadcasts *both* discovery formats
   (`RadioDiscoveryProtocolMode.Auto`) and tags each reply `P1` or `P2` by its
   header.
2. When you power on against a selected radio, `NetworkIO.InitJanusAudio()`
   re-verifies the radio (`NetworkIO.cs:72`) and passes `protocol` (0 or 1)
   to the native side: `nativeInitMetis(..., protocol, model_id, ...)`
   (`network.c:88`).
3. `CurrentRadioProtocol` (`NetworkIO.cs:17`) then steers everything
   protocol-dependent in the C# layer: available sample rates
   (`InitAudioTab`), PureSignal defaults (`HardwareSpecific.PSDefaultPeak`),
   audio amplifier availability, firmware minimum-version checks.
4. On the native side the read thread runs the matching loop: P1 →
   `MetisReadThreadMainLoop()` / `MetisReadThreadMainLoop_HL2()`
   (`networkproto1.c:245`), P2 → `ReadThreadMainLoop()` (`network.c:636`).

A radio model whose board runs *either* protocol is therefore handled
automatically: the protocol follows whatever firmware the radio has loaded.

This fork also extends P1/P2 addressing for **remote (WAN) use**: MI0BOT added
support for non-standard port bases so several HL2s behind one router can be
reached on different forwarded ports (`network.c:86`, custom radios in the
radio list keep their own discovery port).

## 4. Protocol vs hardware

The protocol is a property of the **FPGA gateware/firmware loaded in the
radio**, not of Thetis and not (usually) of the board itself:

| `HPSDRHW` board (`enums.cs:389`) | Radios | Protocol |
|------------------|--------|----------|
| Atlas (0) | Original Atlas/Metis stack | P1 only |
| Hermes (1) | Hermes, ANAN-10/100 | P1 **or** P2 — separate firmware images exist for each; you choose the protocol by flashing the corresponding firmware |
| HermesII (2) | ANAN-10E/100B | P1 or P2 (firmware choice) |
| Angelia (3) | ANAN-100D | P1 or P2 (firmware choice) |
| Orion (4) | ANAN-200D | P1 or P2 (firmware choice) |
| OrionMKII (5) | ANAN-7000DLE/8000DLE, Anvelina-Pro3, Red Pitaya | P1 or P2 (firmware choice) |
| **HermesLite (6)** | **Hermes-Lite 2** | **P1 only** (with HL2 extensions, below) |
| Saturn (10) | ANAN-G2 / G2-1K | P2 (native; version shown as `fpga=… p2app=…` in the radio list, `ucRadioList.cs:1966`) |
| HermesC10 (20) | ANAN-G2E | P2 |

Because ANAN-class boards can be flashed either way, Thetis keeps both stacks
alive and lets discovery decide per session. The board ID byte is the same in
both protocols' discovery replies, which is how a P1-flashed and a P2-flashed
ANAN-100D both show up as `Angelia`.

## 5. The Hermes-Lite 2: Protocol 1, extended

The HL2 gateware implements **openHPSDR Protocol 1 only** (board ID `0x06`).
The official protocol documentation is explicit that HL2 remains compatible
with "a core subset of the openHPSDR protocol" so standard P1 software works,
and it lists **no Protocol 2 support or plans**
(<https://github.com/softerhardware/Hermes-Lite2/wiki/Protocol>).

The HL2 *extends* P1, and this fork implements those extensions with dedicated
code paths (`// MI0BOT` commented):

- **Dedicated P1 loops**: `MetisReadThreadMainLoop_HL2()`
  (`networkproto1.c:427`) and `WriteMainLoop_HL2()` (`networkproto1.c:874`)
  replace the generic loops when `HPSDRModel == HERMESLITE`
  (`networkproto1.c:257`).
- **ACK / request-response C&C**: `C0` bit 7 marks a specific
  request-response exchange instead of the classic free-running register
  cycle.
- **I²C tunnelling**: C&C addresses `0x3C`/`0x3D` carry read/write
  transactions on the HL2's two I²C buses (`WriteMainLoop_HL2`,
  `network.h:118` queue structure) — this is how the N2ADR filter board,
  PA bias, and the HL2 I/O board (`IoBoardHl2.cs`) are controlled.
- **Extended discovery reply**: EEPROM-configured fixed IP (bytes 13–16),
  EEPROM config bytes (11–12), gateware beta version (byte 21) — parsed at
  `clsRadioDiscovery.cs:1172` and shown as `major.minor.beta` in the radio
  picker (`ucRadioList.cs:1974`).
- **384 kHz** maximum sample rate under P1 (`setup.cs:854`).

### Can it be updated?

Two different questions hide here:

**Can the HL2 be switched to Protocol 2?** No. Protocol support lives in the
FPGA gateware, and no official Protocol 2 gateware exists for the HL2 (the
project documents no plans for one). Unlike ANAN boards, there is no P2 image
to flash. If a future community gateware implemented P2 with the standard
discovery reply, this Thetis would in principle drive it via its existing P2
stack — but today HL2 = P1.

**Can the HL2's gateware be updated?** Yes, easily, and this is routine:

- **Over Ethernet (recommended)** using an `.rbf` gateware file. Tools that
  can flash the HL2 in-place: **SparkSDR** (right-click the discovered radio
  while disconnected → program), **Quisk** (config → radio → hardware →
  "Program from RBF file…"), or the `hermeslite` **Python module**
  (`update_gateware_github()`). Programming takes seconds and the radio
  restarts itself.
- The HL2 flash holds **two gateware slots**: slot 1 is a factory fallback
  image, slot 2 the application image. A failed update falls back to the
  factory image, so Ethernet updates are low-risk. A JTAG programmer
  (Altera/Intel USB Blaster) or Raspberry Pi can recover/program directly if
  ever needed.
- Use the `.rbf` matching the board build (e.g. `hl2b5up_main.rbf` for
  build 5+). See
  <https://github.com/softerhardware/hermes-lite2/wiki/Updating-Gateware>.
- **Thetis cannot flash gateware** — this codebase contains no programming
  support; it only *reads* the gateware version from discovery and warns when
  a minimum version isn't met (`NetworkIO.FWVersionsChecked`).

## 6. Quick file map

| Concern | File |
|---------|------|
| Protocol/board enums | `Project Files/Source/Console/enums.cs` (`RadioProtocol`, `HPSDRHW`) |
| Discovery (both protocols) | `Project Files/Source/Console/HPSDR/clsRadioDiscovery.cs` |
| Connect + protocol handoff to native | `Project Files/Source/Console/HPSDR/NetworkIO.cs` |
| Socket setup, **Protocol 2** command/data engine | `Project Files/Source/ChannelMaster/network.c`, `network.h` |
| **Protocol 1** engine incl. HL2 loops & I²C | `Project Files/Source/ChannelMaster/networkproto1.c` |
| Per-model DDC stream/router layout | `Project Files/Source/Console/cmaster.cs` |
| Per-protocol sample-rate lists | `Project Files/Source/Console/setup.cs` (`InitAudioTab`) |
| HL2 I/O board over I²C | `Project Files/Source/Console/HPSDR/IoBoardHl2.cs` |

## References

- [HL2 protocol definition (official wiki)](https://github.com/softerhardware/Hermes-Lite2/wiki/Protocol)
- [HL2 gateware update procedures (official wiki)](https://github.com/softerhardware/hermes-lite2/wiki/Updating-Gateware)
- openHPSDR protocol specifications: <https://github.com/TAPR/OpenHPSDR-SVN>
  (Protocol 1 "USB" and Protocol 2 "Ethernet" documents)
