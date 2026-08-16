# FreeDV in Thetis HL2 — User Guide

> **Status: prototype, `FreeDV` branch, RX-only.** Both modes below are decode-only —
> there is no FreeDV/RADE transmit path yet. This is not part of a released Thetis
> build; it's available on side-loaded `Thetis-Test` builds from the `FreeDV`
> development branch. See `Documentation/FreeDV-Plan.md` for build/implementation
> details and progress notes.

## What is FreeDV?

[FreeDV](https://freedv.org) is an open-source digital voice protocol for HF radio.
It packs a compressed speech codec and a modem into a normal SSB-width channel, so it
runs over the same rigs and antennas as regular voice — no dedicated digital hardware
needed. This build adds native FreeDV decode directly into Thetis's RX chain: no
external FreeDV app, no virtual audio cable, no second computer.

Two decode modes are available, from two different generations of the FreeDV project.
They sound different, behave differently at low SNR, and are implemented by different
libraries — the difference is worth understanding before picking one.

## FreeDV 700E vs. RADE V1

**FreeDV 700E** is the classical, all-DSP mode. Speech is compressed by the Codec2
vocoder to 700 bits/second, protected with LDPC forward error correction, and carried
on a coherent OFDM/PSK modem — the "E" revision is the most fading-resistant of the
700-series. Everything here is conventional, deterministic signal processing (matched
filters, coherent carrier/timing recovery, parity-check decoding), the same lineage
digital voice has used for years. It sounds compressed/synthetic — intelligible, but
clearly a vocoder — and it either syncs or it doesn't: sync drops cleanly under fast
fading or low SNR, and Thetis passes the raw modem audio through until it locks again.

**RADE V1** ("Radio Autoencoder") is FreeDV's newer machine-learning mode. Instead of
a hand-designed vocoder plus a separate modem, RADE trains a neural encoder and
decoder *end-to-end*, together with a model of the radio channel itself — there's no
discrete "700 bits/second" bitstream in the middle, just a continuous learned
representation optimized jointly with how HF actually distorts a signal. The payoff is
noticeably more natural-sounding speech than Codec2 at a comparable bit rate. The
tradeoff, at least as currently implemented, is a different and slower sync/acquisition
behavior — RADE V1 has needed much longer observation windows in on-air testing here to
lock reliably (see `FreeDV-Plan.md` for specifics), and its "sync"/SNR figures come
from pilot correlation and model confidence rather than an LDPC parity-check pass, so
they don't map onto 700E's numbers directly.

| | FreeDV 700E | RADE V1 |
|---|---|---|
| Speech engine | Codec2 vocoder, 700 bit/s | Neural autoencoder, continuous latent, trained jointly with the channel |
| Modem/FEC | Coherent OFDM/PSK + LDPC | Learned constellation, own pilot/sync scheme |
| Sound | Compressed/robotic, clearly a vocoder | More natural ("AI voice"-ish) |
| Sync behavior | Locks/drops fairly quickly; passthrough audio while unsynced | Needs a longer observation window to lock; different sync semantics |
| Underlying library | `libcodec2.dll` | `radae_c` (`rade.lib`) — separate native library |
| Maturity here | RX prototype working, tested against bench + real-RF signals | RX-only wiring in place; first real-RF decode confirmed, sync duration still being characterized |

Bottom line: **700E** is the mature, predictable mode — reach for it first when
checking that a signal is even present, or working weaker signals where a clean
sync/no-sync behavior is what you want. **RADE V1** is the newer, better-sounding mode
worth trying once you have a strong, sustained signal to give it time to lock.

## Using it in Thetis

Both modes live in **Setup → DSP → FreeDV** tab, as two separate group boxes:

- **"FreeDV (prototype)"** — **Decode FreeDV 700E (RX1)** checkbox, with a live sync/SNR
  status label underneath.
- **"RADE V1 (prototype)"** — **Decode RADE V1 (RX1)** checkbox, with its own status
  label.

Only one should normally be enabled at a time, on RX1.

**Setup, both modes:**
1. RX1 mode: **DIGU** (USB-side digital voice convention).
2. Filter: roughly 3 kHz wide, centered normally.
3. Turn **off** NR, NR2, ANF, NB, and squelch — these DSP stages hurt a modem/codec
   signal rather than helping it.
4. Tune to a FreeDV or RADE V1 calling frequency (see the calling-frequency table
   referenced from the `thetisctl freedv-scan` tool, or use
   `thetisctl freedv-reporter watch` to auto-tune to live activity reported on
   [FreeDV Reporter](https://qso.freedv.org)).
5. Check the matching decode box. Audio passes through unmodified until sync is
   achieved, then switches to decoded speech. For 700E, decoded volume is tuned to
   roughly match passthrough loudness (see "Known issues / recent changes" below) —
   expect a similar level, not a sudden drop, when sync engages.

**Remote/CAT status:** the 700E sync/SNR state is also exposed over CAT as `ZZFD`
(run flag) / `ZZFS` (sync/SNR query), and RADE V1 status via `ZZDW`/`ZZDZ`, if you're
scripting or monitoring remotely (e.g. via `thetisctl`).

## Known issues / recent changes

- **700E decoded volume (fixed 2026-08-16).** Earlier builds dropped noticeably
  quieter than passthrough audio the moment sync engaged (~28 dB quieter, RMS) —
  a jarring "why did it just go quiet" effect. Fixed by raising the decoder's
  internal speech gain; live-verified numbers: decoded speech now runs
  **-30.6 dBFS RMS / -6.0 dBFS peak** against passthrough's **-27.1 dBFS RMS /
  -20.2 dBFS peak** — RMS loudness is now essentially matched (within ~3.5 dB),
  and peaks are actually louder than passthrough's. No clipping observed. This
  was verified against one test signal, not yet broadly on-air across varying
  signal strengths — if 700E decode sounds unexpectedly loud or quiet on real
  traffic, that's worth reporting rather than assuming it's expected.
- **RX-only.** Neither mode transmits. There is no FreeDV/RADE keying path in
  this build yet.

See `Documentation/FreeDV-Plan.md` for the full dated history, evidence, and any
issues still open.

## See also

- `Documentation/FreeDV-Plan.md` — implementation details, architecture, and dated
  progress/testing notes for this feature.
- `Documentation/Thetis_VB-Audio_config.md` §7 — using the external `freedv-gui` app
  with Thetis instead of native decode.
- <https://freedv.org> and <https://freedv.org/radio-autoencoder/> — upstream FreeDV
  and RADE project pages.
