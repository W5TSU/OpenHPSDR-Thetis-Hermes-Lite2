# RADE RX2 Native Loopback Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make RX2's already-built "Loopback" checkbox (sub-project 3) actually work, by removing `radae.c`'s RX1-only guard and fanning TX-encoded modem audio into every enabled RX's loop bridge — RX1 and RX2 simultaneously supported.

**Architecture:** Pure native-C fix in `ChannelMaster/radae.c`, the only file touched. Two localized changes: (1) `SetRadaeLoopbackEnabled`'s `if (rx != 0) return;` guard becomes the same `radae_rx_valid()` bounds check its getter already uses; (2) `xradae_tx` computes `lpb1` alongside `lpb0`, ORs them into `lpb_any` for `tx_rx` selection and MOX-gating, and its single-RX bridge-write branch becomes a `push_loop_bridge(loop_rx, scratch, have)` helper called once per enabled RX. The read side (`xradae_rx`) is already fully `rx`-parameterized and untouched. Zero C#/Go/CAT changes — the UI layer built in sub-project 3 already calls `WDSP.SetRadaeLoopbackEnabled(1, ...)` correctly and has just been silently no-op'd until now.

**Tech Stack:** C (MSVC, `ChannelMaster.dll`), built only by CI (GitHub Actions `build.yml` on `FreeDV` — full-solution msbuild, so `ChannelMaster.dll` is genuinely rebuilt). No unit test harness exists for ChannelMaster anywhere in this repo; verification is CI compile-check + a live hardware pass on the test box.

**Spec:** [docs/superpowers/specs/2026-08-26-rade-rx2-native-loopback-design.md](../specs/2026-08-26-rade-rx2-native-loopback-design.md)

## Global Constraints

- **Scope is `Project Files/Source/ChannelMaster/radae.c` only.** No `pipe.c` (700E's loopback stays RX1-only, per the user's explicit choice), no C#, no Go, no CAT changes.
- **RX1 and RX2 loopback must be independently toggleable and safe to run simultaneously** — never mutually exclusive.
- **`mic_io` must be silenced whenever *any* RX's loopback is active** — live mic audio must never reach the air/normal TX path during a loopback test, exactly as today for RX1 alone.
- **No new global state** beyond the static helper function `push_loop_bridge` — reuse the existing `[RADAE_NRX]` arrays (`g_loop_bridge`, `g_loop_bridge_n`, `g_loop_bridge_ovrun_count`, `g_radae_loopback_enabled`), no parallel bookkeeping.
- **CRLF landmine:** `radae.c` is 100% CRLF (verified 2026-08-26, 1667/1667 lines). The plain-text `Edit` tool has flattened whole files' line endings twice on this branch. All edits to `radae.c` use a Python byte-splicing script (`data.replace(old_bytes, new_bytes)` after `assert data.count(old_bytes) == 1`), then `git diff --stat` to confirm the changed-line count is on the order of tens, not hundreds.
- **No local Windows build.** The only compile check is CI (`gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`). **Push before triggering CI** — CI only sees the remote. A stale `ChannelMaster.dll` on the test box would silently mask this fix, so Task 2's deploy must confirm the deployed build's SHA before testing anything.
- **`<thetis-host>` is resolved at execution time, never hardcoded.** The test box has two SSH aliases for the same physical machine: `hl2winbox` (LAN, 192.168.2.12) and `hermes-pc` (VPN, 100.117.67.160). Try `ssh hl2winbox "echo ok"` first, then `ssh hermes-pc "echo ok"`; whichever responds is the alias for every ssh/scp step, and its HostName IP is `<thetis-host>` for every `thetisctl` step.
- Commit after every task; one commit per task minimum.

---

### Task 1: radae.c — remove the RX1-only guard, fan out the bridge write

**Files:**
- Modify: `Project Files/Source/ChannelMaster/radae.c` (`SetRadaeLoopbackEnabled` at ~line 833; new static helper `push_loop_bridge` inserted before `xradae_tx` at ~line 1332; `xradae_tx`'s `lpb0`/`tx_rx` block at ~1345, MOX gate at ~1369, bridge-write block at ~1611)

**Interfaces:**
- Consumes: `radae_rx_valid(int rx)` (existing static, radae.c:692), `g_radae_loopback_enabled` / `g_loop_bridge` / `g_loop_bridge_n` / `g_loop_bridge_ovrun_count` (existing `[RADAE_NRX]` statics, radae.c:285-288), `g_radae_tx_rx`, `RADAE_LOOP_BRIDGE_CAP`.
- Produces: `SetRadaeLoopbackEnabled(int rx, int enable)` now accepts `rx` 0 and 1 (same exported signature, no header/P/Invoke change); new file-local `static void push_loop_bridge(int loop_rx, const float* scratch, int have)`. Task 2 relies only on the changed runtime behavior, not on any new symbol.

- [ ] **Step 1: Write the splicing script**

Create `scratchpad/splice_radae_rx2_loopback.py`. Every anchor below was verified byte-for-byte against the current file on 2026-08-26 — if any assert fires, re-read the file rather than loosening the assert.

```python
#!/usr/bin/env python3
"""Extend radae.c's loopback bridge to RX2 (sub-project 7, see
docs/superpowers/specs/2026-08-26-rade-rx2-native-loopback-design.md).
radae.c is 100% CRLF -- splice byte-exact."""

path = "Project Files/Source/ChannelMaster/radae.c"

def crlf(text):
    return text.replace("\r\n", "\n").replace("\n", "\r\n")

replacements = []

# 1. Setter: replace the hard RX1-only guard (and its stale comment) with
#    the same bounds check GetRadaeLoopbackEnabled already uses.
replacements.append(("setter guard", crlf('''PORT void SetRadaeLoopbackEnabled(int rx, int enable)
{
    /* Loopback is RX1-only -- no cross-protocol loopback can occur. */
    if (rx != 0) return;
'''), crlf('''PORT void SetRadaeLoopbackEnabled(int rx, int enable)
{
    if (!radae_rx_valid(rx)) return;
''')))

# 2. New static helper, inserted just above xradae_tx (after the
#    RADE_TX_SCALE_* defines).
replacements.append(("push_loop_bridge helper", crlf('''#define RADE_TX_SCALE_V1  0.5f
#define RADE_TX_SCALE_V2  0.66f

void xradae_tx(double* mic_io)
'''), crlf('''#define RADE_TX_SCALE_V1  0.5f
#define RADE_TX_SCALE_V2  0.66f

/* Copies up to `have` samples of TX-encoded modem audio (`scratch`) into
 * loop_rx's bridge, respecting its capacity and logging overruns the same
 * way for every RX.  Overflow is dropped and counted, matching the
 * original RX1-only behaviour.  Called once per enabled loopback RX from
 * xradae_tx -- each call touches only its own slot's fields, so the
 * single-writer-per-slot invariant that made the RX1-only code safe
 * carries over unchanged. */
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
                "[RADAE] RX%d loop_bridge OVRUN dropped=%d total=%ld\\n",
                loop_rx + 1, have - take, c);
            OutputDebugStringA(log);
        }
    }
}

void xradae_tx(double* mic_io)
''')))

# 3. lpb0/tx_rx selection: compute both flags, OR into lpb_any; update the
#    stale RX1-only comment.
replacements.append(("tx_rx selection", crlf('''    /* Loopback is RX1-only; otherwise the encoder follows the transmitting RX
     * (set by SetRadaeTxRx at the MOX edge).  tx_rx selects both the handle and
     * its per-RX geometry, so transmitting on an RX running V2 encodes V2. */
    const long lpb0 = _InterlockedAnd(&g_radae_loopback_enabled[0], 1);
    int tx_rx = lpb0 ? 0 : (int)_InterlockedAnd(&g_radae_tx_rx, 1);
'''), crlf('''    /* During any loopback (RX1 and/or RX2) the encoder is forced to RX1's
     * handle/protocol; otherwise it follows the transmitting RX (set by
     * SetRadaeTxRx at the MOX edge).  tx_rx selects both the handle and
     * its per-RX geometry, so transmitting on an RX running V2 encodes V2. */
    const long lpb0 = _InterlockedAnd(&g_radae_loopback_enabled[0], 1);
    const long lpb1 = _InterlockedAnd(&g_radae_loopback_enabled[1], 1);
    const long lpb_any = lpb0 || lpb1;
    int tx_rx = lpb_any ? 0 : (int)_InterlockedAnd(&g_radae_tx_rx, 1);
''')))

# 4. tx_scale comment: "during loopback tx_rx is forced to RX1" stays true,
#    but drop the "(RX1-only)" framing one comment up in the MOX gate.
replacements.append(("mox gate", crlf('''    /* MOX-state gating.  Mirror of the RX-side gate.  Loopback (RX1-only)
     * passes the gate so the encoder->bridge->RX1 path runs without MOX. */
'''), crlf('''    /* MOX-state gating.  Mirror of the RX-side gate.  Loopback (either RX)
     * passes the gate so the encoder->bridge->RX path runs without MOX. */
''')))
replacements.append(("mox gate condition", crlf('''        if (!mox && !lpb0 && !eoo && !hold)
'''), crlf('''        if (!mox && !lpb_any && !eoo && !hold)
''')))

# 5. Bridge-write block: single lpb0 branch -> per-RX fan-out via the
#    helper.  mic_io is silenced under the same condition as before (any
#    loopback active).
replacements.append(("bridge write", crlf('''        if (lpb0)
        {
            int avail = RADAE_LOOP_BRIDGE_CAP - g_loop_bridge_n[0];
            int take  = (have < avail) ? have : avail;
            if (take > 0)
            {
                memcpy(g_loop_bridge[0] + g_loop_bridge_n[0], scratch,
                       (size_t)take * sizeof(float));
                g_loop_bridge_n[0] += take;
            }
            if (take < have)
            {
                long c = ++g_loop_bridge_ovrun_count[0];
                if (c == 1 || (c % 50) == 0)
                {
                    char log[140];
                    sprintf_s(log, sizeof(log),
                        "[RADAE] RX1 loop_bridge OVRUN dropped=%d total=%ld\\n",
                        have - take, c);
                    OutputDebugStringA(log);
                }
            }
            for (i = 0; i < outsize; i++)
            {
                mic_io[2 * i]     = 0.0;
                mic_io[2 * i + 1] = 0.0;
            }
        }
'''), crlf('''        if (lpb_any)
        {
            if (lpb0) push_loop_bridge(0, scratch, have);
            if (lpb1) push_loop_bridge(1, scratch, have);
            for (i = 0; i < outsize; i++)
            {
                mic_io[2 * i]     = 0.0;
                mic_io[2 * i + 1] = 0.0;
            }
        }
''')))

data = open(path, "rb").read()
for name, old, new in replacements:
    old_b, new_b = old.encode("utf-8"), new.encode("utf-8")
    count = data.count(old_b)
    assert count == 1, f"{name}: anchor found {count} times, expected 1"
    data = data.replace(old_b, new_b)
    print(f"{name}: replaced")

open(path, "wb").write(data)
print("done")
```

Note the `\\n` inside the two `sprintf_s` format strings — that is Python escaping for a literal backslash-n in the C source, exactly matching the existing code's `\n` bytes. Do not "fix" it to a real newline.

- [ ] **Step 2: Run it and verify the diff**

```bash
python3 scratchpad/splice_radae_rx2_loopback.py
git diff --stat "Project Files/Source/ChannelMaster/radae.c"
```

Expected: 6 "replaced" lines then `done`, and a diff on the order of `+55/-45` lines. If it shows hundreds or thousands of changed lines, the line endings got flattened: `git checkout -- "Project Files/Source/ChannelMaster/radae.c"` and fix the script before retrying.

- [ ] **Step 3: Sweep for leftovers**

```bash
grep -n "lpb0\|lpb1\|lpb_any\|push_loop_bridge\|rx != 0" "Project Files/Source/ChannelMaster/radae.c"
```

Expected:
- `rx != 0`: **zero** hits (the guard is gone; no other line in this file uses that exact text).
- `lpb0`/`lpb1`: hits only inside `xradae_tx` — the two `_InterlockedAnd` computations, the `lpb_any` OR, and the two `push_loop_bridge` call-site conditions.
- `lpb_any`: three hits — the OR, the MOX gate condition, the bridge-write branch condition.
- `push_loop_bridge`: three hits — the definition, the two call sites.

Any hit outside those (in particular a surviving bare `if (lpb0)` branch, or `lpb0` still in the MOX gate) means a splice was missed.

- [ ] **Step 4: Commit and push**

```bash
git add "Project Files/Source/ChannelMaster/radae.c" && \
git commit -m "feat(radae): extend native loopback bridge to RX2

SetRadaeLoopbackEnabled's rx!=0 guard (which silently no-op'd the RX2
Loopback checkbox built in sub-project 3) becomes the same
radae_rx_valid() bounds check its getter already uses, and xradae_tx's
bridge-write fans TX-encoded modem audio into every enabled RX's slot
via a shared push_loop_bridge helper -- RX1 and RX2 loopback are
independently toggleable and safe to run simultaneously (one encoder,
forced to RX1's handle/protocol during any loopback, feeding both
bridges). mic_io is silenced whenever either loopback is active, same
no-RF guarantee as before. Read side (xradae_rx) was already fully
rx-parameterized and is untouched; zero C#/Go/CAT changes needed.

Part of sub-project 7, see
docs/superpowers/specs/2026-08-26-rade-rx2-native-loopback-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>" && \
git push
```

- [ ] **Step 5: CI compile check**

```bash
gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
# find the run id:
gh run list --workflow=build.yml -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 --limit 3
gh run watch <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
```

Expected: `conclusion: success`. The workflow builds the full solution, so `ChannelMaster.dll` is genuinely recompiled — a C error in the splice (typo'd identifier, unbalanced brace) fails here. If CI fails, read the msbuild error in the run log, fix via a follow-up splice (never the `Edit` tool), commit, push, and re-run.

---

### Task 2: Deploy and live-verify all five scenarios on the test box

**Files:** none (verification only).

**Interfaces:** none — consumes the finished state of Task 1 (CI-green commit on `FreeDV`).

This is the first time RX2's loopback bridge is ever exercised end-to-end. Any surprise (RX2 never syncs in loopback, simultaneous mode starves one RX, audio artifacts) is a genuine finding to report plainly, not a testing-process failure.

- [ ] **Step 1: Resolve `<thetis-host>` and deploy the CI build**

Resolve the host per Global Constraints (`ssh hl2winbox "echo ok"` first, then `ssh hermes-pc "echo ok"`; the responding alias is `<box>` below, its HostName IP — 192.168.2.12 LAN or 100.117.67.160 VPN — is `<thetis-host>`). Then:

```bash
gh run download <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 -n Thetis-HL2-installer -D scratchpad/deploy
scp scratchpad/deploy/*.msi <box>:Downloads/Thetis-Test-rx2loopback.msi
ssh <box> "msiexec /a %USERPROFILE%\\Downloads\\Thetis-Test-rx2loopback.msi /qn TARGETDIR=%USERPROFILE%\\Downloads\\thetis-extract"
ssh <box> "taskkill /IM Thetis.exe /F" # ok if it reports no such process
ssh <box> "robocopy %USERPROFILE%\\Downloads\\thetis-extract\\OpenHPSDR\\Thetis-Test \"C:\\Program Files\\OpenHPSDR\\Thetis-Test\" /MIR" # robocopy exit codes 0-7 are success
```

Relaunch via the scheduled-task interactive-session trick (a plain SSH-launched process lands in Session 0 and crashes on the app's early `MessageBox.Show`):

```bash
ssh <box> "schtasks /create /tn ThetisRelaunch /tr \"'C:\\Program Files\\OpenHPSDR\\Thetis-Test\\Thetis.exe'\" /sc onstart /ru mark /it /f"
ssh <box> "schtasks /run /tn ThetisRelaunch"
sleep 20
ssh <box> "schtasks /delete /tn ThetisRelaunch /f"
```

Confirm the deployed build is Task 1's commit — this matters more than usual because a stale `ChannelMaster.dll` would silently reproduce the exact pre-fix symptom (RX2 loopback checkbox does nothing):

```bash
cd Tools/thetis-ai-control && go run ./cmd/thetisctl cat --host <thetis-host> --timeout 8s version
```

Expected: the `git:` short SHA matches Task 1's commit. If not, stop and re-deploy — do not proceed to testing.

- [ ] **Step 2: Arm the radio and RX2**

Power the radio on and enable RX2 (`console.RX2Enabled` — main console's "RX2" button; RX2 decode is unobservable with it off). Both can be done over CAT:

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZPS1"   # power on (skip if already on)
# RX2 on: use the console's RX2 button via screenshot/click if no CAT path is
# confirmed for it this session; verify state visually with a screenshot.
```

If `scratchpad/cat_roundtrip.py` is missing (fresh session — the scratchpad is session-specific), recreate it exactly:

```python
#!/usr/bin/env python3
"""Raw CAT-over-TCP round-trip helper (same script sub-projects #2/#3 used).
Usage: cat_roundtrip.py [--host H] [--port P] SET:ZZxx... GET:ZZxx ...
SET:ZZFC2 sends "ZZFC2;" (fire-and-forget); GET:ZZFC sends "ZZFC;" and
prints the reply. Thetis's CAT-over-TCP server listens on port 13013."""
import socket, sys

args = sys.argv[1:]
host, port = None, 13013
while args and args[0].startswith("--"):
    if args[0] == "--host": host = args[1]; args = args[2:]
    elif args[0] == "--port": port = int(args[1]); args = args[2:]
    else: sys.exit(f"unknown flag {args[0]}")
if host is None:
    sys.exit("--host is required (use <thetis-host> resolved for this session)")

s = socket.create_connection((host, port), timeout=8)
s.settimeout(4)
for a in args:
    kind, _, code = a.partition(":")
    s.sendall((code + ";").encode("ascii"))
    if kind == "GET":
        buf = b""
        try:
            while not buf.endswith(b";"):
                chunk = s.recv(256)
                if not chunk: break
                buf += chunk
        except socket.timeout:
            pass
        print(f"{a} -> {buf.decode('ascii', 'replace') or '(no reply)'}")
    else:
        print(f"{a} sent")
s.close()
```

Pass `--host <thetis-host>` on every invocation below (the bare `python3 scratchpad/cat_roundtrip.py "..."` forms in Steps 3-8 all take it).

- [ ] **Step 3: Scenario 1 — RX1-only regression**

RX1 mode = RADE V1, RX1 loopback on, RX2 loopback off:

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZEX2" "SET:ZZDL1" "SET:ZZFE0" "GET:ZZDL" "GET:ZZFE"
```

Talk into the mic for ~15 s, then poll RX1 sync a few times over several seconds:

```bash
python3 scratchpad/cat_roundtrip.py "GET:ZZDZ" "GET:ZZDZ" "GET:ZZDZ"
```

Expected: `ZZDL -> 1`, `ZZFE -> 0`, and `ZZDZ` shows sync=1 with a plausible SNR while speaking — identical to pre-change behavior. **This is the regression gate: if RX1 loopback no longer works, stop here** — the refactor broke the existing path and that finding outranks everything below.

- [ ] **Step 4: Scenario 2 — RX2-only loopback (the headline new capability)**

RX1 loopback off; RX2 mode = RADE V1, RX2 loopback on:

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZDL0" "SET:ZZFC2" "SET:ZZFE1" "GET:ZZDL" "GET:ZZFE"
```

Talk into the mic for ~15 s, poll RX2 sync:

```bash
python3 scratchpad/cat_roundtrip.py "GET:ZZFG" "GET:ZZFG" "GET:ZZFG"
```

Expected: `ZZDL -> 0`, `ZZFE -> 1`, and `ZZFG` (RX2 RADE sync/SNR, sub-project 3's CAT command) flips to sync=1 with a plausible SNR. Decoded audio should be audible on RX2's output. Before this change, `SET:ZZFE1` round-tripped at the CAT layer but the native guard discarded it — sync staying 0 here forever was the old symptom.

- [ ] **Step 5: Scenario 3 — simultaneous RX1 + RX2 loopback**

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZDL1" "GET:ZZDL" "GET:ZZFE"
```

(RX2's loopback is still on from Step 4.) Talk into the mic for ~15 s, poll both:

```bash
python3 scratchpad/cat_roundtrip.py "GET:ZZDZ" "GET:ZZFG" "GET:ZZDZ" "GET:ZZFG"
```

Expected: both `ZZDZ` (RX1) and `ZZFG` (RX2) show sync=1 with plausible SNRs from the same single encode pass. If one RX syncs and the other doesn't, that's a real fan-out finding (e.g. one bridge starving), not a test artifact.

- [ ] **Step 6: Scenario 4 — mic_io silencing (no live-mic leak, no keying)**

Throughout Steps 3-5, confirm:
- The radio never keys (no MOX/TX indication — loopback passes the MOX gate without keying; confirm visually via screenshot of the console's TX indicator, or by the absence of any power output).
- No live (unencoded) mic audio is audible on either RX's output — only the decoded (encode→decode round-trip) audio, which has RADE's characteristic slight delay and vocoder timbre. Live-mic bleed would mean the `mic_io` silencing branch isn't covering some path.

- [ ] **Step 7: Scenario 5 — independent toggle-off and clean drain**

With both loopbacks running (state from Step 5), turn RX1's off while leaving RX2's on:

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZDL0" "GET:ZZDL" "GET:ZZFE"
```

Keep talking; poll both again:

```bash
python3 scratchpad/cat_roundtrip.py "GET:ZZDZ" "GET:ZZFG"
```

Expected: `ZZDL -> 0`, `ZZFE -> 1`; RX1 loses sync within a few seconds (its bridge was drained by `SetRadaeLoopbackEnabled`'s existing `if (!enable)` branch and no longer refills), while RX2 keeps sync=1 uninterrupted — no glitch, dropout, or resync on RX2 at the moment RX1 was toggled off. Then the mirror case: re-enable RX1 (`SET:ZZDL1`), confirm it re-syncs, then disable RX2 (`SET:ZZFE0`) and confirm RX1 stays synced while RX2 drops.

- [ ] **Step 8: Leave the box in a known state**

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZDL0" "SET:ZZFE0" "SET:ZZEX0" "SET:ZZFC0" "GET:ZZDL" "GET:ZZFE" "GET:ZZEX" "GET:ZZFC"
```

Expected: all four reads return `0`.

- [ ] **Step 9: Report results**

No commit for this task (verification-only). Report per-scenario pass/fail with the actual CAT reply strings, the Step 6 silencing observations, and any finding from Steps 4/5/7 with specifics — including "worked exactly as designed" stated plainly if that's what happened.
