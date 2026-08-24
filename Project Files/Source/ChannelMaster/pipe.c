/*  pipe.c

This file is part of a program that implements a Software-Defined Radio.

Copyright (C) 2014 Warren Pratt, NR0V

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

The author can be reached by email at  

warren@wpratt.com

*/

#include "cmcomm.h"

pipe pip  = {0};
PIPE ppip = &pip;

// W5TSU: FreeDV 700E RX1 loopback bridge (see pipe.h's comment on
// SetFDVLoopbackEnabled). Raw complex I/Q ring buffer -- bridges TX's
// processed output I/Q directly into RX1's raw antenna I/Q slot, both at
// their respective ch_outrate (checked equal at the RX-side pop, not
// assumed -- xmtr[tx].ch_outrate and rcvr[rx].ch_outrate are independently
// configurable fields with no compile-time guarantee they match, even
// though normal single-DSP-rate operation makes them equal in practice).
// On a rate mismatch, RX1 gets silence rather than either corrupted data
// or a misleadingly-normal antenna signal -- the operator explicitly armed
// a loopback test, so real antenna content leaking through would be a
// false "it worked."
#define FDVLOOP_BRIDGE_CAP 96000   /* 2 s at 48 kHz, matches RADAE_LOOP_BRIDGE_CAP's sizing convention */
static volatile long g_fdvloop_enabled = 0;
static complex g_fdvloop_bridge[FDVLOOP_BRIDGE_CAP];
static int     g_fdvloop_bridge_n = 0;
static long    g_fdvloop_rate_mismatch_warned = 0;

PORT void SetFDVLoopbackEnabled(int enable)
{
    long prev = _InterlockedExchange(&g_fdvloop_enabled, enable ? 1 : 0);
    if (!prev && enable)
    {
        g_fdvloop_bridge_n = 0;                       /* start clean */
        _InterlockedExchange(&g_fdvloop_rate_mismatch_warned, 0);
        OutputDebugStringA("[FDVLOOP] 700E loopback START\n");
    }
    else if (prev && !enable)
    {
        OutputDebugStringA("[FDVLOOP] 700E loopback STOP\n");
    }
}

PORT int GetFDVLoopbackEnabled(void)
{
    return (int)_InterlockedAnd(&g_fdvloop_enabled, 1);
}

void create_spc0()
{
	int i;
	for (i = 0; i < pcm->cmSPC[0]; i++)
	{
		int in_id = inid (2, i);
		// noise blanker
		ppip->spc0[i].panb = create_anb (
			0,									// run
			pcm->xcm_insize[in_id],				// buffsize
			pcm->in[in_id],						// input buffer
			pcm->in[in_id],						// output buffer
			pcm->xcm_inrate[in_id],				// sample rate
			0.0001,								// tau
			0.0001,								// hangtime
			0.0001,								// advtime
			0.05,								// backtau
			30.0);								// threshold

		// noise blanker II
		ppip->spc0[i].pnob = create_nob (
			0,									// run
			pcm->xcm_insize[in_id],				// buffsize
			pcm->in[in_id],						// input buffer
			pcm->in[in_id],						// output buffer
			pcm->xcm_inrate[in_id],				// sample rate
			0,									// mode
			0.0001,								// advslewtime
			0.0001,								// advtime
			0.0001,								// hangslewtime
			0.0001,								// hangtime
			0.025,								// max_imp_seq_time
			0.05,								// backtau
			30.0);								// threshold
	}
}

void destroy_spc0()
{
	int i;
	for (i = 0; i < pcm->cmSPC[0]; i++)
	{
		destroy_anb (ppip->spc0[i].panb);
		destroy_nob (ppip->spc0[i].pnob);
	}
}

void create_pipe()
{
	int i;
	create_siphonEXT(							// siphon for phase2 display for rcvr[0]
		0,										// id
		1,										// run
		pcm->xcm_insize[inid (0, 0)],			// buffer size
		512,									// sipsize
		512,									// fftsize
		0);										// specmode
	(*pip.create_Scope)(0);						// scope display for rcvr[0]
	ppip->rbuff = (double **) malloc0 (pcm->cmRCVR * sizeof (double *));
	for (i = 0; i < pcm->cmRCVR; i++)
	{
		ppip->rbuff[i] = (double *) malloc0 (pcm->rcvr[i].ch_outsize * sizeof (complex));
		(*pip.create_WavePlay)(i);
		(*pip.create_WaveRecord)(i);
		create_ivac(
			i,									// id
			0,									// run
			0,									// iq_type
			0,									// stereo
			pcm->xcm_inrate[inid(0, i)],		// rx i-q rate
			pcm->xcm_inrate[inid(1, 0)],		// mic rate
			pcm->rcvr[i].ch_outrate,			// receiver audio rate
			pcm->xmtr[0].ch_outrate,			// tx monitor rate
			48000,								// vac rate
			pcm->xcm_insize[inid(1, 0)],		// mic buffer size
			pcm->xcm_insize[inid(0, i)],		// iq buffer size
			pcm->rcvr[i].ch_outsize,			// receiver audio buffer size
			pcm->xmtr[0].ch_outsize,			// tx monitor buffer size
			1024);								// vac size
		ppip->rcvr[i].scope_run = 0;			// scope run
		ppip->rcvr[i].playwave_run = 0;			// playwave run
		ppip->rcvr[i].recordwave_run = 0;		// recordwave run
	}
	create_tci();
	create_spc0();
	create_radae();								// W5TSU: RADE V1 (experimental)
}

void destroy_pipe()
{
	int i;
	destroy_spc0();
	destroy_tci();
	destroy_radae();								// W5TSU: RADE V1 (experimental)
	for (i = 0; i < pcm->cmRCVR; i++)
	{
		_aligned_free (ppip->rbuff[i]);
		destroy_ivac (i);
	}
	_aligned_free (ppip->rbuff);
	destroy_siphonEXT (0);
}

void xplaywave(int rx, int state, double* data)
{
	// prevent 1000's of calls when not being used
	if (ppip->rcvr[rx].playwave_run)
	{
		(*pip.rcvr[rx].xplaywave)(state, data);
	}
}

void xrecordwave(int rx, int state, int pos, double* data)
{
	// prevent 1000's of calls when not being used
	if (ppip->rcvr[rx].recordwave_run)
	{
		(*pip.rcvr[rx].xrecordwave)(state, pos, data);
	}
}

void xscope(int rx, int state, double* data)
{
	// prevent 1000's of calls when not being used
	if (ppip->rcvr[rx].scope_run)
	{
		(*pip.rcvr[rx].xscope)(state, data);
	}
}

void xpipe (int stream, int pos, double** buffs)
{
	double* buff = buffs[stream];
	int i, j;
	int rx, tx, sp0;
	int st = stype (stream);
	if      (st == 0) rx  = rxid (stream);
	else if (st == 1) tx  = txid (stream);
	else if (st == 2) sp0 = sp0id (stream);

	if (stream == 0)	// PowerSDR RX1
	{
		switch (pos)
		{
		case 0:	// IQ data
			if (_InterlockedAnd(&g_fdvloop_enabled, 1))										// W5TSU: FreeDV 700E loopback -- drain the bridge into RX1's raw antenna I/Q slot, ahead of xplaywave so an explicit Quick-Play action (if also armed) still wins
			{
				int need = pcm->xcm_insize[stream];
				if (pcm->xmtr[0].ch_outrate != pcm->rcvr[rx].ch_outrate)
				{
					if (!_InterlockedExchange(&g_fdvloop_rate_mismatch_warned, 1))
						OutputDebugStringA("[FDVLOOP] TX/RX1 ch_outrate mismatch -- feeding silence, not bridging\n");
					memset(buff, 0, need * sizeof(complex));
				}
				else
				{
					int have = (g_fdvloop_bridge_n < need) ? g_fdvloop_bridge_n : need;
					if (have > 0)
					{
						memcpy(buff, g_fdvloop_bridge, have * sizeof(complex));
						g_fdvloop_bridge_n -= have;
						if (g_fdvloop_bridge_n > 0)
							memmove(g_fdvloop_bridge, g_fdvloop_bridge + have, g_fdvloop_bridge_n * sizeof(complex));
					}
					if (have < need)
						memset(buff + 2 * have, 0, (need - have) * sizeof(complex));			// underrun -- pad with silence, never leak real antenna I/Q while armed
				}
			}
			if (_InterlockedAnd (&pcm->tci_rx_out_run, 1) && pcm->OutboundTCIRxIQ)
				(*pcm->OutboundTCIRxIQ)(rx, pcm->xcm_insize[stream], buff);						// to TCI
			xplaywave(rx, 0, buff);																// wav player
			xrecordwave(rx, 0, 0, buff);														// wav recorder
			xsiphonEXT(rx, buff);																// siphon for phase2 display
			Spectrum0(_InterlockedAnd(&pip.rcvr[0].top_pan3_run, 0xffffffff), rx, 1, 0, buff);	// stitched pan
			xvacOUT(rx, 0, buff);																// data to VAC
			break;
		case 1: // Audio data
			memcpy (ppip->rbuff[rx], buffs[0], pcm->rcvr[rx].ch_outsize * sizeof (complex));
			for (i = 1; i < pcm->cmSubRCVR; i++)
				for (j = 0; j < 2 * pcm->rcvr[rx].ch_outsize; j++)
					ppip->rbuff[rx][j] += buffs[i][j];
			xscope(rx, 0, ppip->rbuff[rx]);														// scope
			xradae_rx(rx, ppip->rbuff[rx]);											// W5TSU: RADE V1 (experimental)
			xvacOUT(rx, 1, ppip->rbuff[rx]);				// W5TSU: fix - was called before xradae_rx, so VAC got the same never-decoded original audio as the local speaker bug this fix's sibling addressed. Moved after decode -- VAC/TCI/local speaker/wav all now see identical (decoded-or-silence) content.
			if (GetRadaeRxEnabled(rx))												// W5TSU: decode must also reach local monitor audio -- xMixAudio (cmaster.c) reads buffs[], not rbuff[rx], so VAC/TCI/wav got the decode but the speakers didn't
			{
				memcpy (buffs[0], ppip->rbuff[rx], pcm->rcvr[rx].ch_outsize * sizeof (complex));
				for (i = 1; i < pcm->cmSubRCVR; i++)
					memset (buffs[i], 0, pcm->rcvr[rx].ch_outsize * sizeof (complex));
			}
			xtciOUT(rx, 1, ppip->rbuff[rx]);													// data to TCI rx audio
			xrecordwave(rx, 0, 1, ppip->rbuff[rx]);												// wav recorder
			break;
		}
	}
	else if (stream < pcm->cmRCVR)	// other PowerSDR receivers
	{
		switch (pos)
		{
		case 0: // IQ data
			if (_InterlockedAnd (&pcm->tci_rx_out_run, 1) && pcm->OutboundTCIRxIQ)
				(*pcm->OutboundTCIRxIQ)(rx, pcm->xcm_insize[stream], buff);						// to TCI
			xplaywave(rx, 0, buff);																// wav player
			xrecordwave(rx, 0, 0, buff);														// wav recorder
			xvacOUT(rx, 0, buff);																// data to VAC
			break;
		case 1: // Audio data
			memcpy (ppip->rbuff[rx], buffs[0], pcm->rcvr[rx].ch_outsize * sizeof (complex));
			for (i = 1; i < pcm->cmSubRCVR; i++)
				for (j = 0; j < 2 * pcm->rcvr[rx].ch_outsize; j++)
					ppip->rbuff[rx][j] += buffs[i][j];
			xradae_rx(rx, ppip->rbuff[rx]);											// W5TSU: RADE V1 (experimental)
			xvacOUT(rx, 1, ppip->rbuff[rx]);				// W5TSU: fix - was called before xradae_rx, so VAC got the same never-decoded original audio as the local speaker bug this fix's sibling addressed. Moved after decode -- VAC/TCI/local speaker/wav all now see identical (decoded-or-silence) content.
			if (GetRadaeRxEnabled(rx))												// W5TSU: decode must also reach local monitor audio -- xMixAudio (cmaster.c) reads buffs[], not rbuff[rx], so VAC/TCI/wav got the decode but the speakers didn't
			{
				memcpy (buffs[0], ppip->rbuff[rx], pcm->rcvr[rx].ch_outsize * sizeof (complex));
				for (i = 1; i < pcm->cmSubRCVR; i++)
					memset (buffs[i], 0, pcm->rcvr[rx].ch_outsize * sizeof (complex));
			}
			xtciOUT(rx, 1, ppip->rbuff[rx]);													// data to TCI rx audio
			xrecordwave(rx, 0, 1, ppip->rbuff[rx]);												// wav recorder
			break;
		}
	}
	else if (stream == inid (1, 0))	// PowerSDR single transmitter
	{
		switch (pos)
		{
		case 0: // MIC data
			if (!_InterlockedAnd (&pcm->xmtr[0].use_tci_audio, 1))								// stop vacs and playback if tci TX audio is active
			{
				xplaywave(0, 1, buff);															// wav player 0
				xplaywave(1, 1, buff);															// wav player 1
				if (pip.xmtr[0].txvac == 0)  { xvacIN(0, buff, 0);  xvacIN(1, buff, 1); }
				if (pip.xmtr[0].txvac == 1)  { xvacIN(1, buff, 0);  xvacIN(0, buff, 1); }
			}
			xradae_tx(buff);								// W5TSU: RADE V1 (experimental) -- encodes mic (or wav/VAC-sourced) audio in place, replacing it with the modem waveform before the normal SSB TX chain (fexchange0) runs on it; inert unless SetRadaeTxEnabled(1)
			xrecordwave(0, 1, 0, buff);															// wav recorder 0 //[2.10.3.6]MW0LGE moved after vac
			xrecordwave(1, 1, 0, buff);															// wav recorder 1
			break;
		case 1: // IQ data
			xscope(0, 1, buffs[2]);																// scope
			xvacOUT(0, 2, buffs[2]);															// data to VAC 0
			xvacOUT(1, 2, buffs[2]);															// data to VAC 1
			for (i = 0; i < pcm->cmRCVR; i++)
				xtciOUT(i, 2, buffs[2]);														// tx monitor into each TCI rx audio stream
			xrecordwave(0, 1, 1, buffs[2]);														// wav recorder 0
			xrecordwave(1, 1, 1, buffs[2]);														// wav recorder 1
			if (_InterlockedAnd(&g_fdvloop_enabled, 1))										// W5TSU: FreeDV 700E loopback -- push fully-processed TX I/Q into the bridge, RX1's IQ-data case (below) drains it
			{
				int n_tx = pcm->xmtr[0].ch_outsize;
				int take = (n_tx < FDVLOOP_BRIDGE_CAP - g_fdvloop_bridge_n) ? n_tx : (FDVLOOP_BRIDGE_CAP - g_fdvloop_bridge_n);
				if (take > 0)
				{
					memcpy(g_fdvloop_bridge[g_fdvloop_bridge_n], buffs[2], take * sizeof(complex));
					g_fdvloop_bridge_n += take;
				}
			}
			break;
		}
	}
	else if (stream == inid(2, 0))	// PowerSDR Stitched Rcvrs, Left side
	{
		xanb(ppip->spc0[sp0].panb);									// nb
		xnob(ppip->spc0[sp0].pnob);									// nb II
		Spectrum0 (_InterlockedAnd (&pip.rcvr[0].top_pan3_run, 0xffffffff), 0, 0, 0, buff);// stitched pan
	}
	else if (stream == inid(2, 1))	// PowerSDR Stitched Rcvrs, Right side
	{
		xanb (ppip->spc0[sp0].panb);								// nb
		xnob (ppip->spc0[sp0].pnob);								// nb2
		Spectrum0 (_InterlockedAnd (&pip.rcvr[0].top_pan3_run, 0xffffffff), 0, 2, 0, buff);// stitched pan
	}
}

PORT
void SendCBCreateScope (void (__stdcall *create_Scope)(int id))
{
	pip.create_Scope = create_Scope;
}

PORT
void SendCBScope (int id, void (__stdcall *xscope)(int state, double* data))
{
	pip.rcvr[id].xscope = xscope;
}

PORT
void SetScopeRun(int id, int run)
{
	pip.rcvr[id].scope_run = run;
}

PORT
void SendCBCreateWRecord (void (__stdcall *create_WaveRecord)(int id))
{
	pip.create_WaveRecord = create_WaveRecord;
}

PORT
void SendCBWaveRecorder (int id, void (__stdcall *xrecordwave)(int state, int pos, double* data))
{
	pip.rcvr[id].xrecordwave = xrecordwave;
}

PORT
void SetWaveRecorderRun(int id, int run)
{
	pip.rcvr[id].recordwave_run = run;
}

PORT
void SendCBCreateWPlay (void (__stdcall *create_WavePlay)(int id))
{
	pip.create_WavePlay = create_WavePlay;
}

PORT
void SendCBWavePlayer (int id, void (__stdcall *xplaywave)(int state, double* data))
{
	pip.rcvr[id].xplaywave = xplaywave;
}

PORT
void SetWavePlayerRun(int id, int run)
{
	pip.rcvr[id].playwave_run = run;
}

PORT
void SetTopPan3Run (int run)
{
	_InterlockedExchange (&pip.rcvr[0].top_pan3_run, run);
}

PORT
void SetTXVAC (int txid, int txvac)
{
	_InterlockedExchange (&pip.xmtr[txid].txvac, txvac);
}
