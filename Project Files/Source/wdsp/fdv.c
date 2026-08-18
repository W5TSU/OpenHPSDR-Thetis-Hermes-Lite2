/*  fdv.c

This file is part of a program that implements a Software-Defined Radio.

Copyright (C) 2026 Mark Grennan W5TSU

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

FreeDV digital voice RX decode (see fdv.h). Modeled on rnnr.c: a
post-AGC RXA block with float ring buffers to bridge the dsp buffer
size and the modem's variable freedv_nin() block size. Uses libcodec2
(LGPL 2.1, dynamically linked).
*/

#define _CRT_SECURE_NO_WARNINGS
#include "comm.h"
#include "codec2/freedv_api.h"
#include "codec2/modem_stats.h"

// freedv_rx() wants 16-bit-range samples; normalise each nin block to
// this RMS (in short counts) with a smoothed gain so post-AGC level
// changes don't starve or clip the modem
#define FDV_TARGET_RMS_DB   (72.0f)     // ~4000 counts
#define FDV_GAIN_MIN_DB     (-20.0f)
#define FDV_GAIN_MAX_DB     (140.0f)
#define FDV_GAIN_SMOOTH     (0.3f)
#define FDV_SHORT_CEIL      (32000.0f)
// W5TSU: decoded speech level into midbuff. Was 0.30f -- measured live
// against hl2winbox (2026-08-16, FreeDV-Plan.md Phase 3 item 7) at ~28.5 dB
// RMS / ~11 dB peak quieter than the raw modem passthrough audio an operator
// hears while tuning, a jarring drop right when sync engages. Raised to
// 0.75f (~+8 dB): ceiling stays at FDV_SPEECH_GAIN itself (speech_out is
// int16-range, so max amplitude here is FDV_SPEECH_GAIN * 32767/32768), i.e.
// ~-2.5 dBFS worst case -- still well short of clipping, no limiter added.
// Re-measure after this lands; 0.75f is a reasoned first pass, not a
// re-verified final value.
#define FDV_SPEECH_GAIN     (0.75f)

// W5TSU: TX-side mic normalisation into the short domain for freedv_tx(),
// mirroring the RX target above but tuned for a mic-level speech signal
// rather than a received/attenuated RF one -- ~8000 counts, a reasoned
// first pass (not yet measured against real hardware; the RX target above
// only reached its current value after a live measurement pass of its own).
#define FDV_TX_TARGET_RMS_DB (78.0f)     // ~8000 counts
// W5TSU: modem output level into midbuff, matching the RX side's
// FDV_SPEECH_GAIN comment style -- ceiling is this value itself (mod_out is
// int16-range), so max amplitude here is FDV_TX_MODEM_GAIN * 32767/32768.
// Unmeasured first pass; revisit once there's a way to check the resulting
// on-air level (TX-side meter or a captured loopback, matching how RADE V1
// TX's own RADE_TX_SCALE_* constants got tuned via measured modem-RMS ratio).
#define FDV_TX_MODEM_GAIN    (0.5f)

// W5TSU: DEBUG - temporary diagnostic dump state, remove before merge.
// Process-lifetime counters, not per-channel: reset via ResetRXAFDVDebug()
// so each Quick-Play session starts a fresh capture instead of exhausting
// the cap on the first run and silently going quiet on every run after.
static int fdv_dbg_count = 0;
static int fdv_dbg_audio_count = 0;
static int fdv_dbg_resamp_count = 0;
static int fdv_dbg_nin_count = 0; // W5TSU: DEBUG - traces freedv_nin()/ring state on every xfdv() call, to catch the modem-block loop silently never running; remove before merge

// W5TSU: DEBUG - temporary TX-side dump counter, reset via ResetTXAFDVDebug().
// See xfdvtx()'s fdvtx_debug_modem.raw dump.
static int fdvtx_dbg_count = 0;

// W5TSU: DEBUG - scratch buffer for freedv_get_modem_extended_stats(), which
// exposes the OFDM demod's own internal sync/timing/freq-offset estimates
// (sync_metric, foff, rx_timing, clock_offset) instead of just the coarse
// sync flag freedv_get_modem_stats() gives. Wide margin between "sync never
// even gets close" (sync_metric pegged near 0, foff/rx_timing nonsensical -
// points at a structurally wrong signal) and "sync is close but not quite
// there" (sync_metric elevated, foff/timing sane - points at level/SNR).
// File-scope like the other debug state above: not per-channel, reset
// alongside it. ~140KB (mostly MODEM_STATS_NR_MAX x MODEM_STATS_NC_MAX
// symbol history) - too big to want on the stack per block, and only one
// FreeDV RX channel is realistically under test at a time.
static struct MODEM_STATS fdv_dbg_stats;

//ringbuffer (same scheme as rnnr.c)
static void fdv_rb_init(fdv_ring_buffer* rb, int capacity)
{
    rb->buf = malloc0(capacity * sizeof(float));
    rb->capacity = capacity;
    rb->head = 0;
    rb->tail = 0;
    rb->count = 0;
}

static void fdv_rb_free(fdv_ring_buffer* rb)
{
    _aligned_free(rb->buf);
    rb->buf = NULL;
    rb->capacity = 0;
    rb->head = rb->tail = rb->count = 0;
}

static void fdv_rb_clear(fdv_ring_buffer* rb)
{
    rb->head = rb->tail = rb->count = 0;
}

static void fdv_rb_put(fdv_ring_buffer* rb, float v)
{
    if (rb->count < rb->capacity)
    {
        rb->buf[rb->tail] = v;
        rb->tail = (rb->tail + 1) % rb->capacity;
        rb->count++;
    }
}

static int fdv_rb_get_bulk(fdv_ring_buffer* rb, float* dest, int n)
{
    int to_get = n < rb->count ? n : rb->count;
    for (int i = 0; i < to_get; i++)
    {
        dest[i] = rb->buf[rb->head];
        rb->head = (rb->head + 1) % rb->capacity;
    }
    rb->count -= to_get;
    return to_get;
}

// resamplers and rings depend on size/rate; built here so the rate and
// size setters can simply tear down and rebuild around the freedv handle
static void fdv_alloc_streams(FDV a)
{
    int down_out_cap = (int)((long long)a->size * a->modem_rate / a->rate) + 16;
    int up_out_cap = (int)((long long)a->n_max_speech * a->rate / a->speech_rate) + 16;

    a->rs_down_in = malloc0(a->size * sizeof(float));
    a->rs_down_out = malloc0(down_out_cap * sizeof(float));
    a->rs_up_in = malloc0(a->n_max_speech * sizeof(float));
    a->rs_up_out = malloc0(up_out_cap * sizeof(float));

    a->rs_down = create_resampleF(1, a->size, a->rs_down_in, a->rs_down_out, a->rate, a->modem_rate);
    a->rs_up = create_resampleF(1, a->n_max_speech, a->rs_up_in, a->rs_up_out, a->speech_rate, a->rate);

    fdv_rb_init(&a->demod_ring, a->modem_rate + a->n_max_modem);   // ~1 s + one modem frame
    fdv_rb_init(&a->out_ring, a->rate);                            // ~1 s at dsp rate

    a->primed = 0;
}

static void fdv_free_streams(FDV a)
{
    destroy_resampleF(a->rs_down);
    destroy_resampleF(a->rs_up);
    a->rs_down = NULL;
    a->rs_up = NULL;
    _aligned_free(a->rs_down_in);
    _aligned_free(a->rs_down_out);
    _aligned_free(a->rs_up_in);
    _aligned_free(a->rs_up_out);
    fdv_rb_free(&a->demod_ring);
    fdv_rb_free(&a->out_ring);
}

static void fdv_reset(FDV a)
{
    fdv_rb_clear(&a->demod_ring);
    fdv_rb_clear(&a->out_ring);
    flush_resampleF(a->rs_down);
    flush_resampleF(a->rs_up);
    a->primed = 0;
    a->sync = 0;
    a->snr = -100.0f;
    a->agc_seeded = 0;
}

FDV create_fdv(int run, int size, double* in, double* out, int rate)
{
    FDV a = malloc0(sizeof(fdv));
    InitializeCriticalSectionAndSpinCount(&a->cs, 2500);
    a->run = run;
    a->size = size;
    a->in = in;
    a->out = out;
    a->rate = rate;
    a->mode = FREEDV_MODE_700E;
    a->snr = -100.0f;
    a->agc_gain_db = 40.0f;
    a->agc_seeded = 0;

    a->f = freedv_open(a->mode);
    a->modem_rate = freedv_get_modem_sample_rate(a->f);
    a->speech_rate = freedv_get_speech_sample_rate(a->f);
    a->n_max_modem = freedv_get_n_max_modem_samples(a->f);
    a->n_max_speech = freedv_get_n_max_speech_samples(a->f);

    a->demod_in = malloc0(a->n_max_modem * sizeof(short));
    a->speech_out = malloc0(a->n_max_speech * sizeof(short));
    a->nin_buf = malloc0(a->n_max_modem * sizeof(float));

    fdv_alloc_streams(a);
    return a;
}

void destroy_fdv(FDV a)
{
    EnterCriticalSection(&a->cs);
    fdv_free_streams(a);
    freedv_close(a->f);
    a->f = NULL;
    LeaveCriticalSection(&a->cs);
    DeleteCriticalSection(&a->cs);
    _aligned_free(a->demod_in);
    _aligned_free(a->speech_out);
    _aligned_free(a->nin_buf);
    _aligned_free(a);
}

void flush_fdv(FDV a)
{
    EnterCriticalSection(&a->cs);
    fdv_reset(a);
    LeaveCriticalSection(&a->cs);
}

void setBuffers_fdv(FDV a, double* in, double* out)
{
    a->in = in;
    a->out = out;
}

void setSize_fdv(FDV a, int size)
{
    EnterCriticalSection(&a->cs);
    fdv_free_streams(a);
    a->size = size;
    fdv_alloc_streams(a);
    LeaveCriticalSection(&a->cs);
}

void setSamplerate_fdv(FDV a, int rate)
{
    EnterCriticalSection(&a->cs);
    fdv_free_streams(a);
    a->rate = rate;
    fdv_alloc_streams(a);
    LeaveCriticalSection(&a->cs);
}

static float fdv_block_rms(const float* x, int n)
{
    double s = 0.0;
    for (int i = 0; i < n; i++) { double v = (double)x[i]; s += v * v; }
    float r = (float)sqrt(s / (double)n);
    return (r < 1e-9f) ? 1e-9f : r;
}

/********************************************************************************************************
*                                                                                                       *
*                                    FreeDV 700E TX encode (see fdv.h)                                  *
*                                                                                                       *
********************************************************************************************************/

static void fdvtx_alloc_streams(FDVTX a)
{
    int down_out_cap = (int)((long long)a->size * a->speech_rate / a->rate) + 16;
    int up_out_cap = (int)((long long)a->n_nom_modem_samples * a->rate / a->modem_rate) + 16;

    a->rs_down_in = malloc0(a->size * sizeof(float));
    a->rs_down_out = malloc0(down_out_cap * sizeof(float));
    a->rs_up_in = malloc0(a->n_nom_modem_samples * sizeof(float));
    a->rs_up_out = malloc0(up_out_cap * sizeof(float));

    a->rs_down = create_resampleF(1, a->size, a->rs_down_in, a->rs_down_out, a->rate, a->speech_rate);
    a->rs_up = create_resampleF(1, a->n_nom_modem_samples, a->rs_up_in, a->rs_up_out, a->modem_rate, a->rate);

    fdv_rb_init(&a->speech_ring, a->speech_rate + a->n_speech_samples);  // ~1 s + one speech block
    fdv_rb_init(&a->out_ring, a->rate);                                  // ~1 s at dsp rate
}

static void fdvtx_free_streams(FDVTX a)
{
    destroy_resampleF(a->rs_down);
    destroy_resampleF(a->rs_up);
    a->rs_down = NULL;
    a->rs_up = NULL;
    _aligned_free(a->rs_down_in);
    _aligned_free(a->rs_down_out);
    _aligned_free(a->rs_up_in);
    _aligned_free(a->rs_up_out);
    fdv_rb_free(&a->speech_ring);
    fdv_rb_free(&a->out_ring);
}

static void fdvtx_reset(FDVTX a)
{
    fdv_rb_clear(&a->speech_ring);
    fdv_rb_clear(&a->out_ring);
    flush_resampleF(a->rs_down);
    flush_resampleF(a->rs_up);
    a->agc_seeded = 0;
}

FDVTX create_fdvtx(int run, int size, double* in, double* out, int rate)
{
    FDVTX a = malloc0(sizeof(fdvtx));
    InitializeCriticalSectionAndSpinCount(&a->cs, 2500);
    a->run = run;
    a->size = size;
    a->in = in;
    a->out = out;
    a->rate = rate;
    a->mode = FREEDV_MODE_700E;
    a->agc_gain_db = 40.0f;
    a->agc_seeded = 0;

    a->f = freedv_open(a->mode);
    a->modem_rate = freedv_get_modem_sample_rate(a->f);
    a->speech_rate = freedv_get_speech_sample_rate(a->f);
    a->n_speech_samples = freedv_get_n_speech_samples(a->f);
    a->n_nom_modem_samples = freedv_get_n_nom_modem_samples(a->f);

    a->speech_in = malloc0(a->n_speech_samples * sizeof(short));
    a->mod_out = malloc0(a->n_nom_modem_samples * sizeof(short));
    a->speech_scratch = malloc0(a->n_speech_samples * sizeof(float));

    fdvtx_alloc_streams(a);
    return a;
}

void destroy_fdvtx(FDVTX a)
{
    EnterCriticalSection(&a->cs);
    fdvtx_free_streams(a);
    freedv_close(a->f);
    a->f = NULL;
    LeaveCriticalSection(&a->cs);
    DeleteCriticalSection(&a->cs);
    _aligned_free(a->speech_in);
    _aligned_free(a->mod_out);
    _aligned_free(a->speech_scratch);
    _aligned_free(a);
}

void flush_fdvtx(FDVTX a)
{
    EnterCriticalSection(&a->cs);
    fdvtx_reset(a);
    LeaveCriticalSection(&a->cs);
}

void setBuffers_fdvtx(FDVTX a, double* in, double* out)
{
    a->in = in;
    a->out = out;
}

void setSize_fdvtx(FDVTX a, int size)
{
    EnterCriticalSection(&a->cs);
    fdvtx_free_streams(a);
    a->size = size;
    fdvtx_alloc_streams(a);
    LeaveCriticalSection(&a->cs);
}

void setSamplerate_fdvtx(FDVTX a, int rate)
{
    EnterCriticalSection(&a->cs);
    fdvtx_free_streams(a);
    a->rate = rate;
    fdvtx_alloc_streams(a);
    LeaveCriticalSection(&a->cs);
}

void xfdvtx(FDVTX a)
{
    if (a->run && a->f)
    {
        int i;
        EnterCriticalSection(&a->cs);

        // downsample this buffer's mono mic audio to the speech rate --
        // tap point is right after the input resampler, before mic gain/
        // panel/compressor/EQ/CESSB/ALC (see fdv.h's comment on this struct)
        for (i = 0; i < a->size; i++)
            a->rs_down_in[i] = (float)a->in[2 * i + 0];
        a->rs_down->size = a->size;
        int nsp = xresampleF(a->rs_down);

        for (i = 0; i < nsp; i++)
            fdv_rb_put(&a->speech_ring, a->rs_down_out[i]);

        // encode every complete speech block
        while (a->speech_ring.count >= a->n_speech_samples)
        {
            fdv_rb_get_bulk(&a->speech_ring, a->speech_scratch, a->n_speech_samples);

            // normalise the block into the short domain, same smoothed-AGC
            // shape as fdv's RX side (see xfdv's agc_gain_db handling)
            float rms = fdv_block_rms(a->speech_scratch, a->n_speech_samples) * 32767.0f;
            float cur_db = 20.0f * log10f(rms);
            float desired_db = FDV_TX_TARGET_RMS_DB - cur_db;
            if (!a->agc_seeded)
            {
                a->agc_gain_db = desired_db;
                a->agc_seeded = 1;
            }
            else
            {
                a->agc_gain_db += FDV_GAIN_SMOOTH * (desired_db - a->agc_gain_db);
            }
            if (a->agc_gain_db < FDV_GAIN_MIN_DB) a->agc_gain_db = FDV_GAIN_MIN_DB;
            if (a->agc_gain_db > FDV_GAIN_MAX_DB) a->agc_gain_db = FDV_GAIN_MAX_DB;
            float g = 32767.0f * powf(10.0f, a->agc_gain_db / 20.0f);

            for (i = 0; i < a->n_speech_samples; i++)
            {
                float v = a->speech_scratch[i] * g;
                if (v > FDV_SHORT_CEIL) v = FDV_SHORT_CEIL;
                if (v < -FDV_SHORT_CEIL) v = -FDV_SHORT_CEIL;
                a->speech_in[i] = (short)v;
            }

            freedv_tx(a->f, a->mod_out, a->speech_in);   // always produces n_nom_modem_samples

            // W5TSU: DEBUG - temporary raw dump of the encoder's modem output
            // (short PCM at a->modem_rate, i.e. freedv_get_modem_sample_rate())
            // so a first real ZZEF=1 arm can be verified off-box without any
            // MOX/PTT -- xtxa() runs continuously regardless of key state
            // (confirmed 2026-08-18 via the mic capture path), so this starts
            // filling as soon as the encoder is armed and power is on. Remove
            // once the encode path is confirmed sane.
            if (fdvtx_dbg_count < 2000)
            {
                const char* dir = getenv("TEMP");
                char path[512];
                if (dir) snprintf(path, sizeof(path), "%s\\fdvtx_debug_modem.raw", dir);
                else snprintf(path, sizeof(path), "C:\\fdvtx_debug_modem.raw");
                FILE* mf = fopen(path, "ab");
                if (mf)
                {
                    fwrite(a->mod_out, sizeof(short), a->n_nom_modem_samples, mf);
                    fclose(mf);
                }
                fdvtx_dbg_count++;
            }

            for (i = 0; i < a->n_nom_modem_samples; i++)
                a->rs_up_in[i] = a->mod_out[i] * (FDV_TX_MODEM_GAIN / 32768.0f);
            a->rs_up->size = a->n_nom_modem_samples;
            int nup = xresampleF(a->rs_up);
            for (i = 0; i < nup; i++)
                fdv_rb_put(&a->out_ring, a->rs_up_out[i]);
        }

        // drain modem audio into a->out this block. Unlike RX's "prime then
        // pass raw audio through while priming" strategy, an underrun here
        // outputs silence, never falls back to the raw mic -- with FDV TX
        // armed, live mic audio must never leak onto the air disguised as
        // (or alongside) a digital signal, same principle RADE V1 TX's own
        // silence-on-underrun already follows (ChannelMaster/radae.c step 6).
        if (a->out_ring.count >= a->size)
        {
            fdv_rb_get_bulk(&a->out_ring, a->rs_down_in, a->size); // reuse as scratch
            for (i = 0; i < a->size; i++)
            {
                a->out[2 * i + 0] = (double)a->rs_down_in[i];
                a->out[2 * i + 1] = 0.0;
            }
        }
        else
        {
            for (i = 0; i < a->size; i++)
            {
                a->out[2 * i + 0] = 0.0;
                a->out[2 * i + 1] = 0.0;
            }
        }

        LeaveCriticalSection(&a->cs);
    }
    // else: a->in and a->out are the same buffer (midbuff) in practice --
    // true passthrough by construction, nothing to copy
}

void xfdv(FDV a)
{
    if (a->run && a->f)
    {
        int i;
        EnterCriticalSection(&a->cs);

        // downsample this buffer's mono demod audio to the modem rate
        for (i = 0; i < a->size; i++)
            a->rs_down_in[i] = (float)a->in[2 * i + 0];
        a->rs_down->size = a->size;
        int n8k = xresampleF(a->rs_down);

        // W5TSU: DEBUG - temporary diagnostic dump, remove before merge.
        // Raw resampler output (a->rs_down_out): the 8 kHz signal exactly as
        // create_resampleF produces it, before fdv's own RMS/AGC normalizer
        // touches it and before it's chunked into nin-sized modem blocks.
        // Written as contiguous float32 so it can be diffed sample-for-sample
        // against a known-good 8 kHz modem reference (e.g. Tools/FreeDV's
        // ve9qrp raw/wav) to confirm or rule out the untested
        // create_resampleF decimate-by-6 path.
        {
            if (fdv_dbg_resamp_count < 4000) // W5TSU: DEBUG - was 150 (~0.2s of real 8kHz audio at n8k~10.7/call); raised to capture several real seconds for a meaningful sample-for-sample diff
            {
                const char* dir = getenv("TEMP");
                char path[512];
                if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_resamp.raw", dir);
                else snprintf(path, sizeof(path), "C:\\fdv_debug_resamp.raw");
                FILE* rsf = fopen(path, "ab");
                if (rsf)
                {
                    fwrite(a->rs_down_out, sizeof(float), n8k, rsf);
                    fclose(rsf);
                }
                fdv_dbg_resamp_count++;
            }
        }

        for (i = 0; i < n8k; i++)
            fdv_rb_put(&a->demod_ring, a->rs_down_out[i]);

        // run the modem for every complete nin block
        int nin = freedv_nin(a->f);

        // W5TSU: DEBUG - unconditional, every xfdv() call (capped), so we can
        // see nin/ring state even on calls where the while loop below never
        // runs at all - the thing the block-by-block text log can't show
        // since it only fires from inside that loop. Remove before merge.
        if (fdv_dbg_nin_count < 500)
        {
            const char* dir = getenv("TEMP");
            char path[512];
            if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_nin.txt", dir);
            else snprintf(path, sizeof(path), "C:\\fdv_debug_nin.txt");
            FILE* ninf = fopen(path, "a");
            if (ninf)
            {
                fprintf(ninf, "call=%d n8k=%d nin=%d ring_count=%d ring_capacity=%d f=%p loop_will_run=%d\n",
                    fdv_dbg_nin_count, n8k, nin, a->demod_ring.count, a->demod_ring.capacity, (void*)a->f,
                    a->demod_ring.count >= nin && nin > 0);
                fclose(ninf);
            }
            fdv_dbg_nin_count++;
        }

        // W5TSU: fix - codec2's own ofdm_demod() legitimately sets nin=0 when
        // its internal rxbuf already has enough samples buffered for the next
        // frame (ofdm.c: "use internal rxbuf samples if they are available"),
        // meaning "call freedv_rx() again right now with zero new samples to
        // drain me" - not "stop". The previous `nin > 0` guard here treated
        // that as a terminal condition instead, permanently stalling this
        // loop (and hence all decode/sync) the instant codec2 first reported
        // nin=0, which happens routinely once sync engages. Confirmed via
        // fdv_debug_nin.txt: nin dropped to 0 immediately after a sync gain/
        // loss and never recovered, while demod_ring kept filling
        // (unconsumed) toward its capacity for the rest of the session.
        // freedv-gui's own reference RX loop (FreeDVReceiveStep.cpp) has no
        // equivalent guard at all - it just checks a FIFO has >= nin bytes
        // (trivially true for nin=0) and keeps calling the modem.
        // W5TSU: safety bound on how many *consecutive* nin==0 "drain me"
        // iterations we'll honour in a row. The reference loop above has no
        // equivalent cap and assumes nin==0 always self-resolves within a
        // call or two - true in every case observed so far - but nothing in
        // the API contract guarantees that, and `count >= nin` is trivially
        // satisfied forever when nin stays 0, so an unbounded loop here risks
        // hanging the whole DSP thread if that assumption is ever wrong.
        // 16 is a generous multiple of what's actually been observed (1-2).
        int zero_nin_streak = 0;
        while (a->demod_ring.count >= nin && nin >= 0 && zero_nin_streak < 16)
        {
            if (nin == 0) zero_nin_streak++; else zero_nin_streak = 0;

            fdv_rb_get_bulk(&a->demod_ring, a->nin_buf, nin);

            // declared here (not inside the nin>0 block below) so the debug
            // logging further down - gated on the same nin>0 condition, but a
            // separate scope - can still see the values it needs
            float rms = 0.0f, cur_db = 0.0f;

            if (nin > 0)
            {
                // normalise the block into the short domain
                rms = fdv_block_rms(a->nin_buf, nin) * 32767.0f;
                cur_db = 20.0f * log10f(rms);          // block level at unity gain, in short counts
                float desired_db = FDV_TARGET_RMS_DB - cur_db;
                if (!a->agc_seeded)
                {
                    // lock onto the block's own level immediately instead of
                    // smoothing down from a fixed guess, which was overshooting
                    // by ~45dB and clipping demod_in for the first ~15 blocks -
                    // squarely inside the modem's OFDM sync acquisition window
                    a->agc_gain_db = desired_db;
                    a->agc_seeded = 1;
                }
                else
                {
                    // W5TSU: the frozen-gain experiment (skip this re-lock
                    // entirely once seeded) was tried and ruled out - "no sync"
                    // persisted with or without AGC, live-tested against a real
                    // instance (FreeDV-Plan.md). Restored to the smoothing
                    // update.
                    a->agc_gain_db += FDV_GAIN_SMOOTH * (desired_db - a->agc_gain_db);
                }
                if (a->agc_gain_db < FDV_GAIN_MIN_DB) a->agc_gain_db = FDV_GAIN_MIN_DB;
                if (a->agc_gain_db > FDV_GAIN_MAX_DB) a->agc_gain_db = FDV_GAIN_MAX_DB;
                float g = 32767.0f * powf(10.0f, a->agc_gain_db / 20.0f);

                for (i = 0; i < nin; i++)
                {
                    float v = a->nin_buf[i] * g;
                    if (v > FDV_SHORT_CEIL) v = FDV_SHORT_CEIL;
                    if (v < -FDV_SHORT_CEIL) v = -FDV_SHORT_CEIL;
                    a->demod_in[i] = (short)v;
                }
            }
            // nin == 0: nothing to normalise - demod_in is unused (freedv_rx
            // reads zero samples from it below); calling fdv_block_rms() with
            // n=0 would divide by zero and poison agc_gain_db with NaN, so
            // this path is skipped entirely rather than guarded internally.

            int nout = freedv_rx(a->f, a->speech_out, a->demod_in);
            freedv_get_modem_stats(a->f, &a->sync, &a->snr);

            // W5TSU: DEBUG - temporary diagnostic dump, remove before merge.
            // Guarded on nin>0: a nin=0 iteration has no real audio content
            // to record (demod_in wasn't touched this pass, rms/cur_db are
            // still their zero-initialised defaults).
            if (nin > 0)
            {
                if (fdv_dbg_audio_count < 150)
                {
                    const char* dir = getenv("TEMP");
                    char path[512];
                    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_audio.raw", dir);
                    else snprintf(path, sizeof(path), "C:\\fdv_debug_audio.raw");
                    FILE* rawf = fopen(path, "ab");
                    if (rawf)
                    {
                        fwrite(a->demod_in, sizeof(short), nin, rawf);
                        fclose(rawf);
                    }
                    fdv_dbg_audio_count++;
                }
            }
            if (nin > 0)
            {
                if (fdv_dbg_count < 40)
                {
                    // W5TSU: DEBUG - pull the OFDM demod's own internal
                    // sync/timing/freq-offset estimates. Only computed when
                    // actually about to be logged (see comment on
                    // fdv_dbg_stats) - freedv_get_modem_stats() above is
                    // still called unconditionally, it's cheap.
                    freedv_get_modem_extended_stats(a->f, &fdv_dbg_stats);

                    const char* dir = getenv("TEMP");
                    char path[512];
                    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug.txt", dir);
                    else snprintf(path, sizeof(path), "C:\\fdv_debug.txt");
                    FILE* dbgf = fopen(path, "a");
                    if (dbgf)
                    {
                        // W5TSU: DEBUG - self-report the rates this session
                        // is actually running with, once, instead of
                        // requiring a separate Setup -> DSP -> Options check
                        // to confirm fdv.c's 48 kHz assumption holds.
                        if (fdv_dbg_count == 0)
                            fprintf(dbgf, "rates: dsp_rate=%d modem_rate=%d speech_rate=%d dsp_size=%d\n",
                                a->rate, a->modem_rate, a->speech_rate, a->size);

                        fprintf(dbgf, "block=%d nin=%d rms=%.1f cur_db=%.1f agc_gain_db=%.1f in[0..3]=%.5f,%.5f,%.5f,%.5f demod_in[0..3]=%d,%d,%d,%d sync=%d snr=%.1f sync_metric=%.3f foff=%.1f rx_timing=%.2f clock_offset=%.1f\n",
                            fdv_dbg_count, nin, rms, cur_db, a->agc_gain_db,
                            a->in[0], a->in[2], a->in[4], a->in[6],
                            (int)a->demod_in[0], (int)a->demod_in[1], (int)a->demod_in[2], (int)a->demod_in[3],
                            a->sync, a->snr,
                            fdv_dbg_stats.sync_metric, fdv_dbg_stats.foff,
                            fdv_dbg_stats.rx_timing, fdv_dbg_stats.clock_offset);
                        fclose(dbgf);
                    }
                    fdv_dbg_count++;
                }
            }

            if (nout > 0)
            {
                // decoded speech: back to float, up to the dsp rate, into the output ring
                for (i = 0; i < nout; i++)
                    a->rs_up_in[i] = a->speech_out[i] * (FDV_SPEECH_GAIN / 32768.0f);
                a->rs_up->size = nout;
                int nup = xresampleF(a->rs_up);
                for (i = 0; i < nup; i++)
                    fdv_rb_put(&a->out_ring, a->rs_up_out[i]);
            }

            nin = freedv_nin(a->f);
        }

        // emit decoded speech once the ring has real depth; otherwise pass the
        // raw modem audio through so the operator can tune the signal in
        if (!a->primed && a->out_ring.count >= a->size + a->rate / 8)
            a->primed = 1;

        if (a->primed && a->out_ring.count >= a->size)
        {
            fdv_rb_get_bulk(&a->out_ring, a->rs_down_in, a->size); // reuse as scratch
            for (i = 0; i < a->size; i++)
            {
                a->out[2 * i + 0] = (double)a->rs_down_in[i];
                a->out[2 * i + 1] = 0.0;
            }
        }
        else
        {
            a->primed = 0;
            if (a->out != a->in)
                memcpy(a->out, a->in, a->size * sizeof(complex));
        }

        LeaveCriticalSection(&a->cs);
    }
    else if (a->out != a->in)
        memcpy(a->out, a->in, a->size * sizeof(complex));
}

/********************************************************************************************************
*                                                                                                       *
*                                           Public Properties                                           *
*                                                                                                       *
********************************************************************************************************/

PORT
void SetRXAFDVRun(int channel, int run)
{
    FDV a = rxa[channel].fdv.p;
    if (a->run != run)
    {
        EnterCriticalSection(&ch[channel].csDSP);
        flush_fdv(a);
        a->run = run;
        LeaveCriticalSection(&ch[channel].csDSP);
    }
}

PORT
int GetRXAFDVSync(int channel)
{
    return rxa[channel].fdv.p->sync;
}

PORT
double GetRXAFDVSnr(int channel)
{
    return (double)rxa[channel].fdv.p->snr;
}

// W5TSU: FreeDV 700E TX encode enable. Inert by default (run=0 at
// create_fdvtx) and nothing calls this yet -- console/CAT wiring, and the
// MOX/PTT arbiter RADE V1 TX needed (see console.cs's OnMoxPreChangeHandler_
// Radae), are both still open, matching this project's precedent of
// splitting encoder-wiring from PTT-wiring across separate sessions.
PORT
void SetTXAFDVRun(int channel, int run)
{
    FDVTX a = txa[channel].fdvtx.p;
    if (a->run != run)
    {
        EnterCriticalSection(&ch[channel].csDSP);
        flush_fdvtx(a);
        a->run = run;
        LeaveCriticalSection(&ch[channel].csDSP);
    }
}

PORT
int GetTXAFDVRun(int channel)
{
    return txa[channel].fdvtx.p->run;
}

// W5TSU: DEBUG - temporary diagnostic dump control, remove before merge.
// Call at the start of a Quick-Play test session so fdv_debug.txt/
// fdv_debug_audio.raw/fdv_debug_resamp.raw capture that run instead of
// staying silent because an earlier run in the same process already used
// up the counters.
PORT
void ResetRXAFDVDebug(void)
{
    fdv_dbg_count = 0;
    fdv_dbg_audio_count = 0;
    fdv_dbg_resamp_count = 0;
    fdv_dbg_nin_count = 0;

    const char* dir = getenv("TEMP");
    char path[512];

    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug.txt", dir);
    else snprintf(path, sizeof(path), "C:\\fdv_debug.txt");
    FILE* dbgf = fopen(path, "w");
    if (dbgf) fclose(dbgf);

    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_audio.raw", dir);
    else snprintf(path, sizeof(path), "C:\\fdv_debug_audio.raw");
    FILE* rawf = fopen(path, "wb");
    if (rawf) fclose(rawf);

    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_resamp.raw", dir);
    else snprintf(path, sizeof(path), "C:\\fdv_debug_resamp.raw");
    FILE* rsf = fopen(path, "wb");
    if (rsf) fclose(rsf);

    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug_nin.txt", dir);
    else snprintf(path, sizeof(path), "C:\\fdv_debug_nin.txt");
    FILE* ninf2 = fopen(path, "w");
    if (ninf2) fclose(ninf2);
}

// W5TSU: DEBUG - temporary, remove once the TX encode path is confirmed sane.
PORT
void ResetTXAFDVDebug(void)
{
    fdvtx_dbg_count = 0;

    const char* dir = getenv("TEMP");
    char path[512];
    if (dir) snprintf(path, sizeof(path), "%s\\fdvtx_debug_modem.raw", dir);
    else snprintf(path, sizeof(path), "C:\\fdvtx_debug_modem.raw");
    FILE* mf = fopen(path, "wb");
    if (mf) fclose(mf);
}
