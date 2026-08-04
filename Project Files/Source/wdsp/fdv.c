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

// freedv_rx() wants 16-bit-range samples; normalise each nin block to
// this RMS (in short counts) with a smoothed gain so post-AGC level
// changes don't starve or clip the modem
#define FDV_TARGET_RMS_DB   (72.0f)     // ~4000 counts
#define FDV_GAIN_MIN_DB     (-20.0f)
#define FDV_GAIN_MAX_DB     (140.0f)
#define FDV_GAIN_SMOOTH     (0.3f)
#define FDV_SHORT_CEIL      (32000.0f)
#define FDV_SPEECH_GAIN     (0.30f)     // decoded speech level into midbuff

// W5TSU: DEBUG - temporary diagnostic dump state, remove before merge.
// Process-lifetime counters, not per-channel: reset via ResetRXAFDVDebug()
// so each Quick-Play session starts a fresh capture instead of exhausting
// the cap on the first run and silently going quiet on every run after.
static int fdv_dbg_count = 0;
static int fdv_dbg_audio_count = 0;
static int fdv_dbg_resamp_count = 0;

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
            if (fdv_dbg_resamp_count < 150)
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
        while (a->demod_ring.count >= nin && nin > 0)
        {
            fdv_rb_get_bulk(&a->demod_ring, a->nin_buf, nin);

            // normalise the block into the short domain
            float rms = fdv_block_rms(a->nin_buf, nin) * 32767.0f;
            float cur_db = 20.0f * log10f(rms);          // block level at unity gain, in short counts
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

            int nout = freedv_rx(a->f, a->speech_out, a->demod_in);
            freedv_get_modem_stats(a->f, &a->sync, &a->snr);

            // W5TSU: DEBUG - temporary diagnostic dump, remove before merge
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
            {
                if (fdv_dbg_count < 40)
                {
                    const char* dir = getenv("TEMP");
                    char path[512];
                    if (dir) snprintf(path, sizeof(path), "%s\\fdv_debug.txt", dir);
                    else snprintf(path, sizeof(path), "C:\\fdv_debug.txt");
                    FILE* dbgf = fopen(path, "a");
                    if (dbgf)
                    {
                        fprintf(dbgf, "block=%d nin=%d rms=%.1f cur_db=%.1f agc_gain_db=%.1f in[0..3]=%.5f,%.5f,%.5f,%.5f demod_in[0..3]=%d,%d,%d,%d sync=%d snr=%.1f\n",
                            fdv_dbg_count, nin, rms, cur_db, a->agc_gain_db,
                            a->in[0], a->in[2], a->in[4], a->in[6],
                            (int)a->demod_in[0], (int)a->demod_in[1], (int)a->demod_in[2], (int)a->demod_in[3],
                            a->sync, a->snr);
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
}
