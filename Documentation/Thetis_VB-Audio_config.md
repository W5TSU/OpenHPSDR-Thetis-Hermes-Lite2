# Thetis + VB-Audio Voicemeeter Station Audio Configuration (W5TSU)

TX/RX audio setup for Thetis (Hermes-Lite 2 build) using a **Focusrite Scarlett Solo**
interface, an **Audio-Technica AT2020** condenser microphone, and **VB-Audio Voicemeeter**
as the routing mixer.

Why this chain exists: the **Hermes-Lite 2 has no microphone input** — TX audio must reach
the radio from the PC through Thetis's VAC (Virtual Audio Cable) path.

```
AT2020 ──XLR──► Scarlett Solo ──► Voicemeeter (mic strip → B1) ──► Thetis VAC1 in ──► HL2 TX
HL2 RX ──► Thetis VAC1 out ──► Voicemeeter VAIO strip ──► A1 ──► Scarlett headphones
```

---

## 1. Windows and Scarlett Solo groundwork

- Set **48000 Hz** on every device in the chain: Scarlett input and output (Windows Sound →
  Device properties), all Voicemeeter virtual devices, Voicemeeter's internal rate, and
  Thetis VAC1. One rate end-to-end avoids resampling crackle.
- Untick **"Allow applications to take exclusive control"** on every device Voicemeeter uses.
- Scarlett: mic on the XLR input, **INST off**, Air mode to taste.

## 2. AT2020 specifics

- **48V phantom power ON** (button on the Scarlett) — the AT2020 is silent without it.
  Switch phantom on *after* plugging the XLR in, off *before* unplugging.
- It's a **side-address** mic: speak into the front (logo side), not the top.
- Working distance **4–8 inches** with a pop filter. Closer = more level and more bass
  (proximity effect), which the TX EQ then trims.
- Output is modest (~ −37 dB sensitivity): expect the Solo's gain at **2–3 o'clock** for
  close speech. Peaks green/amber on the gain halo, never red. If you run out of knob,
  move closer instead of boosting downstream.
- Condensers hear everything: monitor RX on **headphones**, not speakers, while
  transmitting; use Thetis DEXP to gate room noise between words.

## 3. Voicemeeter

Standard Voicemeeter suffices for one mic + one radio; **Banana** adds a second virtual
pair (B2) for digital-mode apps later.

| Setting | Value |
|---------|-------|
| Hardware Input 1 | Scarlett Solo (**WDM** flavor) |
| Mic strip routing | **B1** on (A1 off unless you want local self-monitor) |
| A1 hardware out | Scarlett Solo output (headphones) |
| VAIO strip (Thetis RX) | routes to **A1** |
| Internal sample rate / buffer | 48000 / 512 |

## 4. Thetis (Setup → Audio → VAC1)

| Setting | Value |
|---------|-------|
| Enable VAC1 | ✔ (and the **VAC1** button on the console front panel) |
| Driver | MME or Windows WASAPI |
| Input | `Voicemeeter Out B1` (the mic) |
| Output | `Voicemeeter Input` (RX audio into Voicemeeter) |
| Sample rate | 48000 |
| Buffer | 1024 (512 if clean, 2048 if audio gaps) |

## 5. Level-setting order (once, in this order)

1. Scarlett hardware gain — voice peaks green/amber
2. Voicemeeter mic strip fader — 0 dB
3. Thetis VAC1 TX gain — Mic/ALC meter just touches 0 on peaks
4. Only then enable processing (DEXP → EQ → CFC/compander), re-checking levels after EQ

## 6. TX EQ profiles

Open with the **EQ** button on the console. The parametric EQ (default mode) offers up to
10 bands with frequency/gain/width; "Legacy EQ" gives the classic 10 sliders. In the wdsp
TX chain the EQ runs **before** the leveler and CFC, so: set EQ first with CFC off, then
add compression. Settings save into the active **TX profile** — keep one profile per style.

### Profile "Rag Chew" — natural, full audio

| Item | Setting |
|------|---------|
| TX filter | 150–2900 Hz |
| EQ 200 Hz | −3 dB (tame proximity boom) |
| EQ 1 kHz | 0 dB (reference) |
| EQ 2 kHz | +2 dB (gentle presence) |
| EQ preamp | 0 dB |
| DEXP | on, threshold just above room noise |
| CFC / compander | off or very light |

### Profile "DX / Contest" — punch and intelligibility

| Item | Setting |
|------|---------|
| TX filter | 250–2900 Hz (hard low cut) |
| EQ below 250 Hz | −6 to −10 dB |
| EQ 600 Hz | −2 dB (de-mud) |
| EQ 2–2.5 kHz | +4 dB (articulation) |
| DEXP | on |
| CFC | on — per-band gains **flat** (tone already shaped by EQ), compression only |
| Leveler | on |

### Verifying by ear

MON (with headphone latency and coloring) is only a rough guide. Record yourself off a
**WebSDR/KiwiSDR** or have a local station record you, then iterate the curve. RX EQ:
leave near flat; +2–3 dB around 1.5–2 kHz helps pull voices from band noise.

## 7. Gotchas

- **Spacebar PTT vs VAC**: Thetis can bypass/mute the VAC mic path when keying with the
  space bar (the "allow space bypass" behavior). No TX audio when space-keying → check
  that Setup option first.
- Digital modes later: keep voice on VAC1/B1; give WSJT-X etc. **VAC2 ↔ B2**
  (Voicemeeter Banana) so they never fight the mic routing.
- If the chain is ever reduced to *mic → radio only*, Thetis VAC1 can take the Scarlett
  directly and Voicemeeter can be dropped — fewer moving parts, less latency.
