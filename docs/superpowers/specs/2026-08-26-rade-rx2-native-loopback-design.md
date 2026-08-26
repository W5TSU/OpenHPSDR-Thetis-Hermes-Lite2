# Sub-project 7: RADE RX2 Native Loopback Design

## Context

Sub-project 3 (RX2/dual-channel Digital Voice support, 2026-08-25) built full
UI and CAT parity for RX2's RADE loopback test — `chkRadeRX2Loopback` in
Setup, calling `WDSP.SetRadaeLoopbackEnabled(1, ...)` exactly like RX1's
control calls `SetRadaeLoopbackEnabled(0, ...)`. That sub-project's final
review confirmed the checkbox and CAT wiring are correct, but also confirmed
a real, permanent limitation one level down: `ChannelMaster/radae.c`'s
loopback bridge only ever runs for RX1. Checking RX2's loopback box arms
nothing — no encoder output ever reaches RX2's decoder, silently.

At the time, the user asked to hold off on fixing this and revisit later.
This document is that revisit: extend the native loopback bridge so RX2
(and, by construction, any other configured RX) can loop back independently
of RX1, with both able to run **simultaneously** (RX1 in loopback while RX2
also runs loopback, testing both decoders against the same TX encode pass
at once, per the user's explicit choice below).

## Scope

**In scope:** `ChannelMaster/radae.c` only — the two blockers identified by
direct inspection (2026-08-26):

1. `SetRadaeLoopbackEnabled(int rx, int enable)` (radae.c:833-855) hard-guards
   `if (rx != 0) return;` — RX2 (and any rx != 0) is refused outright, so the
   backing array `g_radae_loopback_enabled[RADAE_NRX]` (already sized for all
   RXes) never gets a `1` written for `rx != 0`.
2. `xradae_tx(double* mic_io)` (radae.c:1332+) computes only `lpb0` (line
   1348), uses it alone to pick `tx_rx` (line 1349) and to gate MOX-bypass
   (line 1369), and its bridge-write block (lines 1611-1638) only ever reads
   `lpb0` and only ever writes into `g_loop_bridge[0]` / silences `mic_io`
   for `lpb0`'s branch. RX2 loopback is invisible to the one function that
   actually produces the modem audio a loopback bridge would carry.

Everything downstream of the bridge — `xradae_rx`'s read side (radae.c:1099+)
— is **already** fully parameterized by `rx`: `g_radae_loopback_enabled[rx]`,
`g_loop_bridge[rx]`, `g_loop_bridge_n[rx]` are all read/drained per-RX
already (confirmed by direct inspection, lines 1116-1148). No read-side
change is needed.

**Out of scope:**
- `ChannelMaster/pipe.c`'s 700E loopback (`g_fdvloop_enabled`) — a single
  non-array global with no per-RX groundwork at all. Extending it to RX2
  would be new plumbing, not a two-line fix, and the user has explicitly
  chosen to leave 700E's loopback alone for now ("Yes, RADE only for now").
- Any C# or Go change. The Setup UI (`chkRadeRX2Loopback_CheckedChanged`)
  and CAT layer already call `WDSP.SetRadaeLoopbackEnabled(1, ...)`
  correctly (confirmed, sub-project 3's own final review) — once the native
  guard is removed and the write side fans out, RX2's existing UI control
  starts working with no changes on its side.

## Design

### 1. Remove the RX1-only guard in the setter

`SetRadaeLoopbackEnabled` currently refuses any `rx != 0` unconditionally.
Replace the hard guard with the same bounds check `GetRadaeLoopbackEnabled`
already uses (`radae_rx_valid(rx)`, defined at radae.c:692 as
`rx >= 0 && rx < RADAE_NRX`), so RX1 and RX2 are both accepted and anything
genuinely out of range is still safely ignored:

```c
PORT void SetRadaeLoopbackEnabled(int rx, int enable)
{
    if (!radae_rx_valid(rx)) return;
    long prev = _InterlockedExchange(&g_radae_loopback_enabled[rx], enable ? 1 : 0);
    if (!enable)
    {
        /* Drain that RX's bridge so the next loopback session starts clean. */
        g_loop_bridge_n[rx] = 0;
    }
    if (enable && !prev)
    {
        char log[80];
        sprintf_s(log, sizeof(log), "[RADAE] RX%d loopback START\n", rx + 1);
        OutputDebugStringA(log);
    }
    else if (!enable && prev)
    {
        char log[80];
        sprintf_s(log, sizeof(log), "[RADAE] RX%d loopback STOP\n", rx + 1);
        OutputDebugStringA(log);
    }
}
```

The stale comment ("Loopback is RX1-only -- no cross-protocol loopback can
occur") is removed along with the guard it justified — sub-project 3 already
established that RX1 and RX2 can run independent, non-interfering RADE
protocols (V1 on one, V2 on the other), so there is no cross-protocol hazard
here either; each RX's bridge is a separate array slot.

### 2. Fan the encoder's output into every enabled RX's bridge, simultaneously

Per the user's explicit choice, RX1 and RX2 loopback are **not mutually
exclusive** — both can be armed at once, and both then receive the same
single TX encode pass (there is only one active encoder/`tx_rx` at a time;
loopback taps its output, it doesn't run two encoders). This matches how a
real operator would use it: arm both decoders, key the mic once, confirm
both RX1 and RX2 lock and decode the same audio.

**`tx_rx` selection (line 1348-1350):** compute both flags and OR them:

```c
    const long lpb0 = _InterlockedAnd(&g_radae_loopback_enabled[0], 1);
    const long lpb1 = _InterlockedAnd(&g_radae_loopback_enabled[1], 1);
    const long lpb_any = lpb0 || lpb1;
    int tx_rx = lpb_any ? 0 : (int)_InterlockedAnd(&g_radae_tx_rx, 1);
    if (tx_rx < 0 || tx_rx >= RADAE_NRX) tx_rx = 0;
    if (g_rade[tx_rx] == NULL) return;
```

`tx_rx` forces to RX1 (0) whenever *either* loopback is active, same as
today's RX1-only behavior — the encoder always runs RX1's protocol/handle
during any loopback session. This preserves the existing "during loopback
tx_rx is forced to RX1" comment's intent (line 1355) and the tx_scale
comment just below it; both are updated to say "RX1 or RX2" instead of
"RX1-only" where they reference loopback forcing RX1.

**MOX-state gating (line 1369):** replace `!lpb0` with `!lpb_any` so the
gate passes (allowing the encoder to run without a real MOX event) whenever
either RX's loopback is armed:

```c
        if (!mox && !lpb_any && !eoo && !hold)
```

**Bridge-write block (lines 1611-1638):** replace the single `if (lpb0)`
branch with a small per-RX helper invoked for each enabled loopback, so the
existing overrun-logging code isn't duplicated. Add a static helper above
`xradae_tx` (near the other static helpers in this file):

```c
/* Copies up to `have` samples of TX-encoded modem audio (`scratch`) into
 * loop_rx's bridge, respecting its capacity and logging overruns the same
 * way for every RX. Returns nothing -- overflow is dropped and counted,
 * matching the existing RX1-only behavior. */
static void push_loop_bridge(int loop_rx, const float* scratch, int have)
{
    int avail = RADAE_LOOP_BRIDGE_CAP - g_loop_bridge_n[loop_rx];
    int take  = (have < avail) ? have : avail;
    if (take > 0)
    {
        memcpy(g_loop_bridge[loop_rx] + g_loop_bridge_n[loop_rx], scratch,
               (size_t)take * sizeof(float));
        g_loop_bridge_n[loop_rx] += take;
    }
    if (take < have)
    {
        long c = ++g_loop_bridge_ovrun_count[loop_rx];
        if (c == 1 || (c % 50) == 0)
        {
            char log[140];
            sprintf_s(log, sizeof(log),
                "[RADAE] RX%d loop_bridge OVRUN dropped=%d total=%ld\n",
                loop_rx + 1, have - take, c);
            OutputDebugStringA(log);
        }
    }
}
```

And the call site becomes:

```c
        if (lpb_any)
        {
            if (lpb0) push_loop_bridge(0, scratch, have);
            if (lpb1) push_loop_bridge(1, scratch, have);
            for (i = 0; i < outsize; i++)
            {
                mic_io[2 * i]     = 0.0;
                mic_io[2 * i + 1] = 0.0;
            }
        }
        else
        {
            for (i = 0; i < have; i++)
            {
                mic_io[2 * i]     = (double)scratch[i];
                mic_io[2 * i + 1] = 0.0;
            }
            for (; i < outsize; i++)
            {
                mic_io[2 * i]     = 0.0;
                mic_io[2 * i + 1] = 0.0;
            }
        }
```

`mic_io` is silenced under the same condition as today (any loopback
active) — live mic audio must never reach the air/normal RX path during a
loopback test regardless of which RX(es) are looping.

### Why this is safe / minimal

- `g_loop_bridge`, `g_loop_bridge_n`, `g_loop_bridge_ovrun_count`,
  `g_loop_bridge_underrun_count`, and `g_radae_loopback_enabled` are already
  declared as `[RADAE_NRX]` arrays (radae.c:285-288, confirmed) — no new
  storage, no resize.
- The read side (`xradae_rx`) is untouched — it already loops correctly per
  `rx` and needs no changes.
- No new synchronization is introduced: each RX's bridge fields are only
  ever touched by that RX's slot index (`push_loop_bridge(0, ...)` only
  touches index 0's fields, `push_loop_bridge(1, ...)` only index 1's) —
  the existing single-writer-per-slot invariant that made the RX1-only code
  safe carries over unchanged to two writers of two disjoint slots.
- No C#/Go/CAT change is required — this is confirmed by direct inspection
  of `chkRadeRX2Loopback_CheckedChanged` in `setup.cs` (built in sub-project
  3), which already calls `WDSP.SetRadaeLoopbackEnabled(1, ...)` — that call
  currently no-ops against the `if (rx != 0) return;` guard and will simply
  start working once the guard above is removed.

## Testing

Native-C-only change in a DLL with no existing unit test harness in this
repo (`ChannelMaster` has no test project) — verification is live, on
`hl2winbox`/`hermes-pc`, matching this project's established practice for
every prior `radae.c` change:

1. Build, deploy, arm RX1 loopback only (`chkRadeRX1Loopback`) with RX1
   RADE decode enabled — confirm unchanged behavior (regression check: RX1
   loopback must still work exactly as before this change).
2. Arm RX2 loopback only (`chkRadeRX2Loopback`) with RX2 RADE decode
   enabled, RX1 loopback off — key the mic, confirm RX2 syncs and decodes
   audible, intelligible speech. This is the headline new capability.
3. Arm **both** RX1 and RX2 loopback simultaneously, both decoders enabled
   — key the mic once, confirm both RX1 and RX2 independently sync and
   decode the same speech.
4. Confirm `mic_io` is silenced (no live-mic leak, no on-air keying) during
   all three scenarios above — same check sub-project 1's original loopback
   bug fix already established as required.
5. Toggle each loopback checkbox off mid-session and confirm that RX's
   bridge drains cleanly (`g_loop_bridge_n[rx]` reset via
   `SetRadaeLoopbackEnabled`'s existing `if (!enable)` branch) without
   affecting the other RX's still-running loopback.

## Global Constraints

- Scope is `ChannelMaster/radae.c` only. No `pipe.c` (700E), no C#, no Go.
- RX1 and RX2 loopback must be independently toggleable and safe to run
  simultaneously — never mutually exclusive, per the user's explicit choice.
- `mic_io` (and therefore any real hardware TX output) must be silenced
  whenever *any* RX's loopback is active, exactly as today for RX1 alone.
- No new global state beyond the helper function `push_loop_bridge` — reuse
  the existing per-RX arrays, do not introduce parallel bookkeeping.
