# A History of Thetis

### From the SDR-1000 and the HPSDR group to OpenHPSDR, Thetis, and the Hermes-Lite 2 forks

This document traces the twenty-plus-year lineage of the software in this repository — from the
first open-source amateur SDR through the HPSDR/TAPR hardware project, the PowerSDR console
family, and Thetis, down to the Hermes-Lite 2 forks of which this repo is one. Dates before ~2015
are drawn from project archives and community documentation; where sources give only approximate
dates, that is noted.

---

## 1. The roots: FlexRadio's SDR-1000 and PowerSDR (2003–2005)

The story begins outside HPSDR. In 2003 **Gerald Youngblood, K5SDR** (then AC5OG) of FlexRadio
Systems shipped the **SDR-1000**, grown out of his QEX article series "A Software-Defined Radio
for the Masses" — the first open-source SDR transceiver sold to hams. Its console software,
**PowerSDR**, was released under the GNU GPL, with DSP provided by the **DttSP** library written
by **Frank Brickle, AB2KT** and **Bob McGwier, N4HY**, working alongside FlexRadio's
**Eric Wachsmann, KE5DTO**.

That GPL release is the single most consequential licensing decision in this history: every
console in the Thetis family tree descends from it, and FlexRadio, Brickle, McGwier, Wigley, and
Samphire are still credited in this repository's source headers today.

## 2. The HPSDR group forms (2005–2006)

In 2005, three independent efforts converged:

- **Phil Covington, N8VB** started a "High Performance SDR" project around an FPGA motherboard
  with a USB 2.0 interface;
- **Phil Harman, VK6APH** (later VK6PH) and **Bill Tracey, KD5TFD** were building a sound-card
  replacement for the SDR-1000.

The groups merged in early 2006 and **HPSDR — High Performance Software Defined Radio** — was
born, with hpsdr.org (later openhpsdr.org) as its home. The concept was modular: a passive
backplane into which experimenters plug the boards they need. It was open source for software and
open (mixed licensing) for hardware, developed publicly on mailing-list reflectors by volunteers
worldwide.

### TAPR's role

**TAPR** (Tucson Amateur Packet Radio, tapr.org) — the nonprofit that had shepherded amateur
packet radio hardware since the early 1980s — became HPSDR's manufacturing and distribution
partner in 2006, kitting and selling the boards at cost so experimenters didn't have to source
their own PCBs. In March 2008 TAPR formally funded the Mercury receiver development. TAPR remains
the institutional steward today: the official OpenHPSDR firmware, protocol documentation, and
software mirrors live in the **github.com/TAPR** organization (OpenHPSDR-Firmware,
OpenHPSDR-PowerSDR, OpenHPSDR-Thetis).

## 3. The hardware era: Atlas to Hermes (2006–2012)

The board family, in rough order of appearance, with their lead designers:

| Board | Role | People / notes |
|-------|------|----------------|
| **Atlas** | Passive backplane | Phil Covington N8VB |
| **Ozy** | USB 2.0 PC interface | Phil Covington N8VB (replacing the earlier Xylo experiment) |
| **Janus** | A/D–D/A audio companion (SDR-1000 sound-card replacement) | Harman VK6APH / Tracey KD5TFD lineage |
| **Mercury** (2008) | Direct-sampling DDC receiver, 0–55 MHz | Funded by TAPR in March 2008; Phil Harman VK6APH FPGA work |
| **Penelope / PennyLane** | 0.5 W DUC exciter/transmitter | — |
| **Excalibur** | Frequency reference | — |
| **Alex** | RX/TX bandpass filter set | Named for Alexiares; controlled by the consoles to this day (`Console/HPSDR/Alex.cs`) |
| **Metis** (~2011) | Ethernet interface replacing USB | Basis of "Protocol 1" Ethernet framing |
| **Hermes** (2011–2012) | Single-board transceiver: Mercury + Metis + Penelope integrated | Concept by **Kevin Wheatley, M0KHZ** with the OpenHPSDR group |

Hermes is the pivotal design: a complete 0.5 W DDC/DUC transceiver on one board. Almost
everything that follows — the Apache Labs ANAN line and the Hermes-Lite — descends from it.

### The protocols

Communication between radio and PC settled into two openHPSDR standards, both largely authored by
**Phil Harman, VK6PH**:

- **Protocol 1** ("USB protocol", later carried over Metis Ethernet) — the framing this
  repository's `ChannelMaster/networkproto1.c` still implements, and the protocol the
  Hermes-Lite 2 speaks;
- **Protocol 2** (~2015 onward) — a cleaner, higher-performance Ethernet protocol for the newer
  ANAN radios, and the original reason Thetis exists.

## 4. The software family branches (2007–2014)

Multiple consoles grew up around the HPSDR hardware — the first big branching of the tree:

- **PowerSDR (HPSDR ports)** — FlexRadio's GPL console adapted to HPSDR hardware in three
  parallel ports: a simple one by **Bill Tracey, KD5TFD**; a featureful multi-receiver port
  (**PowerSDR mRX**) by **Doug Wigley, W5WC**; and a diversity-reception port for dual Mercury
  boards by **Joe Martin, K5SO**. W5WC's mRX line is Thetis's direct ancestor.
- **KISS Konsole** — Phil Harman VK6APH's deliberately simple C# console ("Keep It Simple,
  Stupid"), built on **Phil Covington N8VB's** SharpDSP, intended as a learning platform.
- **ghpsdr / ghpsdr3** — **John Melton, G0ORX/N6LYT**'s Linux console, which grew a
  server/dspserver/client architecture (early remote operation).
- **cuSDR** — **Hermann von Hasseln, DL3HVH**'s GPU-accelerated console; later continued as
  cudaSDR by **Rick Koch, N1GP** (whose G1 work appears in official Thetis to this day).
- **KE9NS PowerSDR** — **Darrin Kohn, KE9NS**'s long-running PowerSDR fork for the FlexRadio
  1500/3000/5000 side of the family (a sibling branch, not an HPSDR one).

### WDSP replaces DttSP

Around 2013 **Warren Pratt, NR0V** wrote **WDSP**, a complete modern replacement for the aging
DttSP engine — the `wdsp/` library in this repository, still copyright NR0V. With it came
**PureSignal** (NR0V's adaptive TX predistortion) and **PowerSDR mRX PS**, the W5WC/NR0V console
that carried the OpenHPSDR flag through the mid-2010s. NR0V also wrote **ChannelMaster**
(`ChannelMaster/` here), the audio/network routing engine both PowerSDR mRX PS and Thetis use.

## 5. Apache Labs and the ANAN line (2012–present)

**Apache Labs** (Gurgaon, India — founded by **Abhishek "Abhi" Prakash**) commercialized the
OpenHPSDR designs in cooperation with the project: the **ANAN-10/100** packaged the Hermes board
with 10 W/100 W PAs; the **ANAN-100D** used **Angelia** (dual-ADC Hermes descendant); the
**ANAN-200D** used **Orion**; later the 7000DLE/8000DLE and the current **G2 / G2 Ultra**
(Saturn) generation followed. Development stayed in the open — firmware in TAPR's GitHub, consoles
by the community — making ANAN the primary hardware target for Thetis.

## 6. Thetis (2016–present)

**Thetis** (in Greek myth, mother of Achilles — the HPSDR naming tradition of gods and heroes
continues) was created by **Doug Wigley, W5WC** as the next-generation console for **Protocol 2**,
rebuilt from PowerSDR mRX PS around NR0V's WDSP and ChannelMaster with a new display pipeline.
First public releases appeared around 2017–2018 under the TAPR/OpenHPSDR umbrella; Protocol 1
support was added later, letting Thetis serve both radio generations.

Around 2019–2020, stewardship passed to **Richard Samphire, MW0LGE** (GitHub **ramdor/Thetis**),
who drove six years of intense development — the modern SharpDX display, the multimeter system
(`MeterManager.cs`), NR3/NR4 neural and spectral noise reduction (RNNoise and libspecbleach — see
[NR3.md](NR3.md)), TCI, Andromeda/G2 front-panel support, CFC, skins, and the 2.9.x → 2.10.x
series. In **April 2026**, MW0LGE stepped back (commit "So long, and thanks for all the fish",
v2.10.3.14; a final v2.10.3.15 and an N1MM fix followed through May 2026). In **July 2026** the
ramdor repository was marked active again under new stewardship, with remote-operation features
announced as the development focus.

## 7. The Hermes-Lite: HPSDR for everyone (2014–present)

In parallel, **Steve Haynal, KF7O** asked how cheap an HPSDR-compatible transceiver could get.
His answer — the **Hermes-Lite** (github.com/softerhardware/Hermes-Lite2) — replaced the costly
FPGA+converter architecture with the **AD9866**, a mass-produced cable-modem front-end chip
(12-bit ADC/DAC, LNA) paired with a small FPGA speaking HPSDR **Protocol 1**:

- **Hermes-Lite 1.x** — proof of concept, developed from 2014, first builds ~2015–2016;
- **Hermes-Lite 2.x** (~2018 onward) — the refined 5 W, 0–38.4 MHz transceiver produced in
  community group buys via Makerfabs; fully open hardware, gateware, and software.

Because the HL2 speaks Protocol 1, every console in the family — Thetis, piHPSDR, linHPSDR,
Quisk, SparkSDR, deskHPSDR — can drive it. But the official Thetis targets ANAN hardware first,
which is exactly the gap the forks below fill.

## 8. The fork map

```
FlexRadio PowerSDR + DttSP (K5SDR, KE5DTO, AB2KT, N4HY)          2003–
    │  GPL release
    ├── FlexRadio PowerSDR v2.x ── KE9NS PowerSDR (Flex 3000/5000 fork)
    │
    └── HPSDR ports (2007–)
          ├─ KD5TFD simple port          ├─ K5SO diversity port
          └─ PowerSDR mRX (W5WC) ──► PowerSDR mRX PS (W5WC + NR0V wdsp/PureSignal)
                                            │
                                            ▼
                              Thetis (W5WC, ~2016–2018; Protocol 2)
                                            │
                     stewardship → ramdor/Thetis (MW0LGE, ~2019–2026;
                                    2026– new maintainers, remote-op focus)
                                            │   TAPR/OpenHPSDR-Thetis = official mirror
              ┌─────────────────────────────┼──────────────────────────┐
              ▼                             ▼                          ▼
   mi0bot/OpenHPSDR-Thetis        ON7OFF/Thetis                 (other private forks)
   (MI0BOT, Hermes-Lite 2         (2025; hand-merge of mi0bot
    adaptation: IoBoardHl2,        HL2 code onto ramdor
    HL2 TX latency/PTT hang,       2.10.3.11; inactive)
    I2C, "Radio Model" code)
              │
              ▼
   W5TSU/OpenHPSDR-Thetis-Hermes-Lite2   ◄── this repository
   (merges the best of the HL2 ecosystem: mi0bot HL2 line +
    official ramdor fixes; v2.10.3.17 released July 2026)

 Parallel non-Windows lineage (shares wdsp, not PowerSDR code):
   ghpsdr/ghpsdr3 → piHPSDR (G0ORX; Raspberry Pi) → linHPSDR (G0ORX) → deskHPSDR (DL1BZ)
   KISS Konsole (VK6APH) · cuSDR (DL3HVH) → cudaSDR (N1GP) · Quisk · SparkSDR
```

## 9. This repository's place in the story

This fork — **W5TSU/OpenHPSDR-Thetis-Hermes-Lite2**, maintained by **Mark Grennan, W5TSU** —
exists to combine the best parts of the Thetis ecosystem for Hermes-Lite 2 users:

- the **mi0bot** HL2 adaptation (**MI0BOT**, whose "Radio Model" code and register-level HL2 I/O
  board support — `Console/HPSDR/IoBoardHl2.cs` — make Thetis a first-class HL2 console);
- fixes tracked from the **official ramdor line** (e.g., the N1MM CW-shift fix cherry-picked as
  official Thetis `8071b543` into v2.10.3.17, July 2026);
- with an eye on **ON7OFF**'s Android-remote work and MW0LGE's remote-access direction.

Twenty-three years after the SDR-1000, the chain is unbroken: DttSP begat WDSP; PowerSDR begat
mRX begat Thetis; Atlas-and-Ozy begat Hermes begat both the ANAN flagships and the $300
Hermes-Lite 2 — and the source headers in this repository still name everyone from FlexRadio
Systems to Richard Samphire. It remains what it was in 2005: amateurs building world-class radio
in the open.

---

## Sources

- [OpenHPSDR — Wikipedia](https://en.wikipedia.org/wiki/OpenHPSDR)
- [openhpsdr.org](https://openhpsdr.org/) — project site, wiki (KISS Konsole, Ghpsdr/Ghpsdr3 pages)
- [History of HPSDR Mercury and QuickSilver — Phil Covington N8VB](http://pcovington.blogspot.com/2007/10/history-of-hpsdr-mercury-and-quick.html)
- [TAPR GitHub organization](https://github.com/TAPR) — OpenHPSDR-Firmware, OpenHPSDR-PowerSDR, OpenHPSDR-Thetis
- [Introducing the FLEX-5000A (TAPR DCC 2007, K5SDR/N4HY)](http://www.tapr.net/meetings/DCC_2007/DCC2007-FLEX5000A-K5SDR-N4HY.pdf) — PowerSDR/DttSP credits
- [Thetis User Manual](https://apache-labs.com/public/storage/download_file/1756364911_1020_Thetis-manual_1.0.pdf) — PowerSDR→Thetis lineage, W5WC, Protocol 2
- [ramdor/Thetis releases](https://github.com/ramdor/Thetis/releases) — MW0LGE era and handover notes
- [Hermes-Lite 2 — softerhardware GitHub](https://github.com/softerhardware/Hermes-Lite2) and [Makerfabs](https://www.makerfabs.com/hermes-lite-2.html)
- [The Hermes-Lite SDR — SWLing Post](https://swling.com/blog/2019/05/the-hermes-lite-sdr-an-open-source-hf-transceiver-based-on-a-broadband-modem-chip/)
- [ANAN SDR Transceivers overview (AB4OJ)](https://www.qsl.net/ab4oj/nsprog/anan_sdr.pdf) — Hermes/Angelia/Orion → ANAN mapping, M0KHZ credit
- [ARRL Surfin': Negotiating the HPSDR Highway](https://www.arrl.org/news/surfin-negotiating-the-hpsdr-highway)
- This repository's git history and source headers (FlexRadio, W5WC, MW0LGE, NR0V attributions)
