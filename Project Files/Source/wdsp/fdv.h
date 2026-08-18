/*  fdv.h

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

FreeDV digital voice RX decode using libcodec2 (LGPL 2.1, dynamically
linked - see lib/codec2_x64/COPYING). Sits post-AGC in the RXA chain:
demodulated SSB audio is resampled to the 8 kHz modem rate, run through
freedv_rx(), and the decoded speech is resampled back to the dsp rate.
While unsynced (or before the output ring primes) the raw modem audio
passes through so the operator can tune.
*/

#ifndef _fdv_h
#define _fdv_h

struct freedv; // opaque, from codec2/freedv_api.h

typedef struct _fdv_ring_buffer {
    float* buf;
    int capacity;
    int head;
    int tail;
    int count;
} fdv_ring_buffer;

typedef struct _fdv
{
    int run;
    int size;                   // dsp buffer size (complex samples per call)
    double* in;
    double* out;
    int rate;                   // dsp sample rate

    struct freedv* f;
    int mode;                   // FREEDV_MODE_* (700E for the prototype)
    int modem_rate;             // freedv_get_modem_sample_rate()
    int speech_rate;            // freedv_get_speech_sample_rate()
    int n_max_modem;            // freedv_get_n_max_modem_samples()
    int n_max_speech;           // freedv_get_n_max_speech_samples()

    // dsp rate <-> modem/speech rate converters (non-complex float)
    RESAMPLEF rs_down;          // rate -> modem_rate
    RESAMPLEF rs_up;            // speech_rate -> rate
    float* rs_down_in;
    float* rs_down_out;
    float* rs_up_in;
    float* rs_up_out;

    fdv_ring_buffer demod_ring;  // 8 kHz demod samples awaiting freedv_nin()
    fdv_ring_buffer out_ring;    // decoded speech at dsp rate awaiting output

    short* demod_in;            // freedv_rx() input, n_max_modem
    short* speech_out;          // freedv_rx() output, n_max_speech
    float* nin_buf;             // staging for one nin block

    // input level normalisation into the short domain for freedv_rx
    float agc_gain_db;
    int agc_seeded;              // 0 until the first block has set agc_gain_db directly

    int primed;                 // out_ring has built enough depth to stream
    int sync;
    float snr;

    CRITICAL_SECTION cs;
} fdv, *FDV;

extern FDV create_fdv (int run, int size, double* in, double* out, int rate);
extern void destroy_fdv (FDV a);
extern void flush_fdv (FDV a);
extern void xfdv (FDV a);
extern void setBuffers_fdv (FDV a, double* in, double* out);
extern void setSize_fdv (FDV a, int size);
extern void setSamplerate_fdv (FDV a, int rate);

// W5TSU: FreeDV 700E TX encode -- mirrors fdv's RX ring-buffer/resampler
// structure in the opposite direction. Sits early in the TXA chain, right
// after the input resampler and before mic gain/compressor/EQ/CESSB/ALC:
// when run, it replaces raw mic audio in place with the 700E modem
// waveform, which then flows through the rest of the normal TX chain
// unmodified (same choice already shipped for RADE V1 TX in ChannelMaster/
// radae.c -- whether passing a digital modem waveform through analog-voice
// dynamics processing, CESSB included, degrades it enough to matter is an
// open question shared by both features, not decided here).
typedef struct _fdvtx
{
    int run;
    int size;                   // dsp buffer size (complex samples per call)
    double* in;
    double* out;
    int rate;                   // dsp sample rate

    struct freedv* f;
    int mode;                   // FREEDV_MODE_* (700E for the prototype, matches fdv's RX side)
    int modem_rate;             // freedv_get_modem_sample_rate()
    int speech_rate;            // freedv_get_speech_sample_rate()
    int n_speech_samples;       // freedv_get_n_speech_samples() -- fixed freedv_tx() input block
    int n_nom_modem_samples;    // freedv_get_n_nom_modem_samples() -- fixed freedv_tx() output block

    // dsp rate <-> speech/modem rate converters (non-complex float)
    RESAMPLEF rs_down;          // rate -> speech_rate (mic audio down)
    RESAMPLEF rs_up;            // modem_rate -> rate (modem audio up)
    float* rs_down_in;
    float* rs_down_out;
    float* rs_up_in;
    float* rs_up_out;

    fdv_ring_buffer speech_ring; // speech-rate mic samples awaiting a full freedv_tx() block
    fdv_ring_buffer out_ring;    // modem audio at dsp rate awaiting output

    short* speech_in;           // freedv_tx() input, n_speech_samples
    short* mod_out;              // freedv_tx() output, n_nom_modem_samples
    float* speech_scratch;      // one popped speech_ring block, float domain, n_speech_samples

    // input level normalisation into the short domain, same smoothed-AGC
    // approach as fdv's RX side (see fdv.c) -- this tap is pre mic-gain/
    // panel, so the raw level here is whatever the audio device delivers
    float agc_gain_db;
    int agc_seeded;

    // W5TSU: MOX/PTT arbiter (console.cs's OnMoxPreChangeHandler_FDVTX).
    // While draining, xfdvtx() stops ingesting new mic audio into
    // speech_ring but keeps draining whatever's already buffered there and
    // in out_ring -- the console-side arbiter holds the real key-up
    // (SetChannelState, which is what actually stops xtxa()/xfdvtx() from
    // running at all) until GetTXAFDVDrained() reports empty, so the tail
    // of the operator's last words isn't silently discarded the instant
    // MOX drops. Same principle as RADE V1 TX's EOO-flush-before-drop
    // arbiter, adapted to 700E's very different architecture (see
    // ChannelMaster/radae.c's g_radae_tx_silence_hold).
    int draining;

    CRITICAL_SECTION cs;
} fdvtx, *FDVTX;

extern FDVTX create_fdvtx (int run, int size, double* in, double* out, int rate);
extern void destroy_fdvtx (FDVTX a);
extern void flush_fdvtx (FDVTX a);
extern void xfdvtx (FDVTX a);
extern void setBuffers_fdvtx (FDVTX a, double* in, double* out);
extern void setSize_fdvtx (FDVTX a, int size);
extern void setSamplerate_fdvtx (FDVTX a, int rate);

#endif //_fdv_h
