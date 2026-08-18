# FreeDV in Thetis HL2 — User Guide

> **Status: prototype, `FreeDV` branch, decode is usable — transmit is not.** Both
> modes' RX/decode side work as described below. RADE V1 TX exists and, as of
> 2026-08-18, has been keyed twice over real RF as controlled mechanics tests
> (`thetisctl` + `--confirm-tx`, not through the normal UI), the second one
> carrying a valid station-ID (end-of-over callsign) burst — but it is **not ready
> for normal operation**: nothing confirms the transmitted signal actually decodes
> anywhere, since there's no second receiver at this station yet to check. **700E
> TX** also now exists (a from-scratch encoder, separate effort from RADE V1) and
> has had one real-RF verification test confirming clean output, but it has no
> MOX/PTT arbiter at all yet — even further from usable than RADE V1 TX. Both TX
> encoders are disarmed by default. Treat TX as an active development item, not a
> feature — see "Known issues / recent changes" below. This is not part of a
> released Thetis build; it's available on side-loaded
> `Thetis-Test` builds from the `FreeDV` development branch. See
> `Documentation/FreeDV-Plan.md` for build/implementation details and progress notes.

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
noticeably more natural-sounding speech than Codec2 at a comparable bit rate — now
confirmed audible on this build, see the note below on how long sync needs to run
first. Its "sync"/SNR figures come from pilot correlation and model confidence rather
than an LDPC parity-check pass, so they don't map onto 700E's numbers directly.

> ℹ️ **RADE V1 needs a sustained signal before audio engages.** As of 2026-08-16, RX
> sync and decoded audio both work end-to-end — first confirmed intelligible RADE V1
> decode through Thetis — but only after a couple of minutes of continuous sync;
> a quick or weak signal may show sync with no audible speech yet (that's expected
> behavior for this vocoder, not a bug — its own reference decoder does the same).
> A separate, now-fixed bug (`1425318d`) had VAC (Virtual Audio Cable) users hearing
> nothing regardless of sync duration, because VAC was tapping the audio stream
> *before* RADE V1's decode ran; if you're on the native speaker/monitor path rather
> than VAC, that specific bug never applied to you. See `FreeDV-Plan.md` Stage C for
> the full trace, including a self-correction along the way (an early read of the
> decoded audio as "broken" turned out to be normal — this vocoder's raw output really
> is mostly silence per short frame; only a proper full-length comparison against
> freedv-gui's own reference decoder caught that).

| | FreeDV 700E | RADE V1 |
|---|---|---|
| Speech engine | Codec2 vocoder, 700 bit/s | Neural autoencoder, continuous latent, trained jointly with the channel |
| Modem/FEC | Coherent OFDM/PSK + LDPC | Learned constellation, own pilot/sync scheme |
| Sound | Compressed/robotic, clearly a vocoder | More natural ("AI voice"-ish) — confirmed audible, needs sustained sync first |
| Sync behavior | Locks/drops fairly quickly; passthrough audio while unsynced | Reproduces reliably on a controlled signal; audio follows after ~2 min of sustained sync |
| Underlying library | `libcodec2.dll` | `radae_c` (`rade.lib`) — separate native library |
| Maturity here | RX prototype working, tested against bench + real-RF signals | Sync + decoded audio both confirmed working (HackRF positive control); real off-air confirmation still open |

Bottom line: **700E** is the mature, predictable mode — reach for it first when
checking that a signal is even present, or working weaker/shorter signals where a
clean sync/no-sync behavior is what you want. **RADE V1** now genuinely works and
sounds better once it's decoding, but give it real time — don't judge it on a signal
that's only been synced for a few seconds.

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

- **RADE V1 decoded audio (fixed 2026-08-16).** Confirmed working end-to-end for the
  first time: sync, decode, and audible speech, given a couple of minutes of
  sustained signal. The real bug (`1425318d`) was specific to VAC (Virtual Audio
  Cable) users — VAC tapped the audio stream *before* RADE V1's decode ran, so VAC
  output was always raw/undecoded regardless of sync state, the same class of gap
  the `1c185f14` local-speaker fix had already closed for the native speaker path a
  day earlier. Native speaker/monitor listeners were never affected by this specific
  bug. Separately (not a bug, just a characteristic worth knowing): RADE V1's raw
  decoder output is mostly silence within any short window — this vocoder needs a
  sustained sync run before real speech content shows up, confirmed against
  freedv-gui's own reference decoder behaving the same way.
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
- **RADE V1 TX encoder wired, loopback-confirmed (2026-08-17/18).** The encoder
  (mic conditioning → LPCNet → `rade_tx`) is hooked into Thetis's TX audio path
  and, via a dedicated no-RF loopback test mode, produced its first audible,
  intelligible synthesized speech end-to-end. See `FreeDV-Plan.md`'s 2026-08-17/18
  entry for the full trace.
- **MOX/PTT wired; first real over-the-air transmission (2026-08-18) — mechanics
  only, not ready for use.** Real PTT keying is now wired (`d4486c84`): key-down
  arms the encoder, key-up holds real PTT open up to 2s (hard-capped) so the
  end-of-over burst can flush before the radio unkeys, normal SSB/CW/FM PTT release
  is unaffected. Tested with a real 6-second over-the-air transmission — clean
  key/unkey, no stuck PTT, Thetis stayed responsive.
- **Callsign ID wired, second on-air test (2026-08-18).** `SetRadaeEooCallsign` is
  now reachable via a new CAT field, `ZZDJ` (fixed 15-char, `KY`-style). A second
  real 6-second over-the-air transmission confirmed the end-of-over burst now
  carries a real, operator-set callsign rather than the earlier uninitialized bit
  pattern — that specific gap is closed. **What's still open:** no second receiver
  at this station confirms the transmitted signal actually decodes anywhere —
  everything tested so far proves the TX chain keys/unkeys cleanly and IDs
  correctly, not that the signal itself is received/decoded elsewhere. TX encoder
  is disarmed by default after testing. Don't treat this as an operator-usable
  transmit mode yet — see `FreeDV-Plan.md`'s 2026-08-18 entries.
- **700E TX encoder built, wired, verified via real PTT test (2026-08-18) —
  further from usable than RADE V1 TX.** A separate, from-scratch encoder (700E
  had no TX code to reuse, unlike RADE V1). CAT-armable via a test-only field,
  `ZZEF`. One real 4-second silent PTT test confirmed the encoder produces clean,
  correctly-scaled modem audio (no clipping, no NaN, spectrum matches expected
  audio-passband OFDM occupancy) — a real structural checkpoint. But **there is no
  MOX/PTT arbiter for this mode at all yet** — nothing manages the end of a
  transmission properly — so this is meaningfully less ready than RADE V1 TX, not
  a parallel feature at the same stage. Disarmed by default. See `FreeDV-Plan.md`
  Stage B for the full trace.

See `Documentation/FreeDV-Plan.md` for the full dated history, evidence, and any
issues still open.

## See also

- `Documentation/FreeDV-Plan.md` — implementation details, architecture, and dated
  progress/testing notes for this feature.
- `Documentation/Thetis_VB-Audio_config.md` §7 — using the external `freedv-gui` app
  with Thetis instead of native decode.
- <https://freedv.org> and <https://freedv.org/radio-autoencoder/> — upstream FreeDV
  and RADE project pages.
