# Digital Voice Mode Selector Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fold FreeDV 700E into the RADE Setup panel as one exclusive Mode selector (Off/700E/RADE V1/RADE V2), replacing three independent, uninterlocked enable paths with a single control, mutual exclusion enforced at the property-setter level.

**Architecture:** Four layers, built bottom-up: (1) a low-level RX/TX interlock in `radio.cs`'s existing `RXAFDVRun`/`RXRadaeEnabled`/`TXAFDVRun` property setters so arming one subsystem always disarms the other, from any entry point; (2) one new CAT command (`ZZEX`) for the unified mode index, plus a one-line change to the existing `ZZDK` command so it participates in the same interlock; (3) a WinForms UI layer — `cmbRadeProtocol` repurposed into a 4-item `cmbRadeMode` combo, moved into the RX1 Core group, with the old FreeDV tab and its controls removed entirely; (4) integration verification on real hardware (`hl2winbox`), since this project has no local Windows build/test environment and no automated test suite.

**Tech Stack:** C# WinForms (.NET Framework 4.8), P/Invoke into `wdsp.dll`/`ChannelMaster.dll` (both C/C++, unmodified by this plan — every backend function this plan calls already exists). No unit test framework exists in this codebase; verification is CI compile-check + live hardware CAT/UI checks (see Global Constraints).

**Spec:** [docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md](../specs/2026-08-25-digital-voice-mode-selector-design.md)

## Global Constraints

- **No local Windows build environment and no automated test suite.** This is a Windows-only WinForms + native-DLL project (see repo `CLAUDE.md`); there is no `dotnet`/MSBuild available locally and no unit test runner anywhere in the codebase. "Testing" in this plan means: careful manual review of each diff (brace matching, no orphaned references) for code-only tasks, then one real CI build + `hl2winbox` hardware verification pass at the end (Task 4) — not per-task automated tests.
- **CRLF/LF landmine — read before touching any of these five files:** `radio.cs`, `CATCommands.cs` (both genuinely mixed CRLF/LF), and `CATParser.cs`, `CATStructs.xml`, `setup.designer.cs` (all three 100% CRLF). The plain-text `Edit` tool has been proven to flatten an entire file's line endings when writing back a modified region — this happened once already in sub-project #1's CAT-command work, producing a spurious ~3600-line diff in `CATCommands.cs` that had to be `git checkout`-reverted. For all edits to these five files, use a Python byte-splicing script (`data.replace(old_bytes, new_bytes)` after an `assert data.count(old_bytes) == 1`), never the `Edit` tool. `setup.cs` is LF-dominant and tolerated plain `Edit`-tool changes cleanly in sub-project #1 — the `Edit` tool is fine there, but run `git diff --stat` after every edit regardless and confirm the changed-line count roughly matches what you intended (a few dozen lines, not thousands) before moving on.
- **Naming convention:** every WinForms control in this codebase uses the (incorrect but established) namespace `System.Windows.Forms.*` for its own thread-safe types (`CheckBoxTS`, `ComboBoxTS`, `GroupBoxTS`, `LabelTS`, `NumericUpDownTS`) — match this exactly, don't "fix" it.
- **Deploy pipeline** (established in sub-project #1, reused verbatim in Task 4): trigger `gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2` → poll `gh run list`/`gh run view --json status,conclusion` → download the `Thetis-HL2-installer` artifact via `gh run download <id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 -n Thetis-HL2-installer -D <dir>` → `scp` the MSI to `hl2winbox:Downloads/` → run `extract.ps1` (msiexec `/a` admin-extract) → run `stop_thetis.ps1` → run `copy_build.ps1` (robocopy `/MIR`) → run `relaunch.ps1` via `run_interactive.ps1 -ScriptName relaunch.ps1` (the scheduled-task/interactive-session trick — required because a plain SSH session runs in Session 0 and cannot launch a real interactive desktop app or take real screenshots) → verify via `./Tools/thetis-ai-control/thetisctl cat --host 100.117.67.160 --timeout 8s version`. All of these `.ps1` scripts already exist from sub-project #1's work in the session scratchpad; recreate them if starting a fresh session (their content is given again in Task 4).
- **Commit after every task**, referencing the sub-project spec in the commit body.

---

### Task 1: RX1/TX1 mode interlock (radio.cs)

**Files:**
- Modify: `Project Files/Source/Console/radio.cs:2280-2296` (class `RadioDSPRX`, property `RXAFDVRun`)
- Modify: `Project Files/Source/Console/radio.cs:2305-2320` (class `RadioDSPRX`, property `RXRadaeEnabled`)
- Modify: `Project Files/Source/Console/radio.cs:3152-3168` (class `RadioDSPTX`, property `TXAFDVRun`)

**Interfaces:**
- Consumes: existing private fields `rx_fdv_run`/`rx_fdv_run_dsp`, `rx_radae_enabled`/`rx_radae_enabled_dsp`, `tx_fdv_run`/`tx_fdv_run_dsp` (all already declared on these two classes), existing `update`/`force`/`thread`/`subrx` fields, `WDSP.SetRXAFDVRun`, `WDSP.SetRadaeRxEnabled`, `WDSP.SetTXAFDVRun` (all pre-existing P/Invoke declarations in `dsp.cs`), and `WDSP.SetRadaeTxEnabled(int enable)` (pre-existing, `dsp.cs:336-337`).
- Produces: `RXAFDVRun`'s setter now also clears `RXRadaeEnabled` (sets it to `0`) whenever armed (`value == 1`); `RXRadaeEnabled`'s setter symmetrically clears `RXAFDVRun`; `TXAFDVRun`'s setter now also calls `WDSP.SetRadaeTxEnabled(0)` whenever armed. **No signature changes** — all three remain `public int X { get; set; }`, so every existing caller (`ZZDV`, `ZZDW`, `ZZEF` in `CATCommands.cs`, and the setup.cs handlers Task 3 rewrites) needs no changes for this task alone.

- [ ] **Step 1: Write the interlock-splicing script**

Create `scratchpad/splice_radio_interlock.py` (path is illustrative — put it wherever your session's scratchpad lives):

```python
#!/usr/bin/env python3
"""Add the Digital Voice mode interlock to radio.cs's RXAFDVRun/
RXRadaeEnabled/TXAFDVRun setters. This region of radio.cs is LF-only
(confirmed via a CRLF/LF byte count before writing this script) --
splice with plain \\n, matching the surrounding file exactly."""

path = "Project Files/Source/Console/radio.cs"

old_rxafdvrun = '''        public int RXAFDVRun
        {
            get { return rx_fdv_run; }
            set
            {
                rx_fdv_run = value;
                if (update)
                {
                    if (value != rx_fdv_run_dsp || force)
                    {
                        WDSP.SetRXAFDVRun(WDSP.id(thread, subrx), value);
                        rx_fdv_run_dsp = value;
                    }
                }
            }
        }
'''

new_rxafdvrun = '''        public int RXAFDVRun
        {
            get { return rx_fdv_run; }
            set
            {
                rx_fdv_run = value;
                if (update)
                {
                    if (value != rx_fdv_run_dsp || force)
                    {
                        WDSP.SetRXAFDVRun(WDSP.id(thread, subrx), value);
                        rx_fdv_run_dsp = value;
                    }
                }
                // W5TSU: Digital Voice mode interlock (sub-project 2 of 5, see
                // docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md).
                // Arming 700E RX decode always disarms RADE RX decode on the
                // same receiver -- both tap the same RX1 antenna feed, only
                // one digital voice codec can meaningfully run at a time.
                // Guarded on value==1 so disarming (value=0) never recurses
                // into RXRadaeEnabled's own mirror-image check below.
                if (value == 1 && RXRadaeEnabled != 0)
                {
                    RXRadaeEnabled = 0;
                }
            }
        }
'''

old_rxradaeenabled = '''        public int RXRadaeEnabled
        {
            get { return rx_radae_enabled; }
            set
            {
                rx_radae_enabled = value;
                if (update)
                {
                    if (value != rx_radae_enabled_dsp || force)
                    {
                        WDSP.SetRadaeRxEnabled((int)thread, value);
                        rx_radae_enabled_dsp = value;
                    }
                }
            }
        }
'''

new_rxradaeenabled = '''        public int RXRadaeEnabled
        {
            get { return rx_radae_enabled; }
            set
            {
                rx_radae_enabled = value;
                if (update)
                {
                    if (value != rx_radae_enabled_dsp || force)
                    {
                        WDSP.SetRadaeRxEnabled((int)thread, value);
                        rx_radae_enabled_dsp = value;
                    }
                }
                // W5TSU: Digital Voice mode interlock -- mirror image of
                // RXAFDVRun's own check above. See that property's comment.
                if (value == 1 && RXAFDVRun != 0)
                {
                    RXAFDVRun = 0;
                }
            }
        }
'''

old_txafdvrun = '''        public int TXAFDVRun
        {
            get { return tx_fdv_run; }
            set
            {
                tx_fdv_run = value;

                if (update)
                {
                    if (value != tx_fdv_run_dsp || force)
                    {
                        WDSP.SetTXAFDVRun(WDSP.id(thread, 0), value);
                        tx_fdv_run_dsp = value;
                    }
                }
            }
        }
'''

new_txafdvrun = '''        public int TXAFDVRun
        {
            get { return tx_fdv_run; }
            set
            {
                tx_fdv_run = value;

                if (update)
                {
                    if (value != tx_fdv_run_dsp || force)
                    {
                        WDSP.SetTXAFDVRun(WDSP.id(thread, 0), value);
                        tx_fdv_run_dsp = value;
                    }
                }
                // W5TSU: Digital Voice mode interlock -- arming 700E TX
                // encode always disarms RADE TX encode. RADE TX has no
                // radio.cs-cached property (direct WDSP global, per
                // sub-project #1's wiring pattern), so this is a
                // one-directional call -- the matching disarm-the-other-way
                // call lives at RADE TX's own arm sites
                // (cmbRadeMode_SelectedIndexChanged in setup.cs, and CAT
                // ZZDK in CATCommands.cs), not here.
                if (value == 1)
                {
                    WDSP.SetRadaeTxEnabled(0);
                }
            }
        }
'''

data = open(path, "r", encoding="utf-8", newline="").read()

for name, old, new in [
    ("RXAFDVRun", old_rxafdvrun, new_rxafdvrun),
    ("RXRadaeEnabled", old_rxradaeenabled, new_rxradaeenabled),
    ("TXAFDVRun", old_txafdvrun, new_txafdvrun),
]:
    count = data.count(old)
    assert count == 1, f"{name}: anchor found {count} times, expected 1"
    data = data.replace(old, new)
    print(f"{name}: replaced")

open(path, "w", encoding="utf-8", newline="").write(data)
print("done")
```

- [ ] **Step 2: Run it and verify the diff**

```bash
python3 scratchpad/splice_radio_interlock.py
git diff --stat "Project Files/Source/Console/radio.cs"
```

Expected: the script prints `RXAFDVRun: replaced`, `RXRadaeEnabled: replaced`,
`TXAFDVRun: replaced`, `done`. `git diff --stat` shows roughly `+30, -0` for
this one file (three ~9-line insertions) — if it shows hundreds or thousands
of changed lines, the file's line endings got flattened; run
`git checkout -- "Project Files/Source/Console/radio.cs"` and retry with a
corrected script rather than committing.

- [ ] **Step 3: Review the diff by eye**

```bash
git diff "Project Files/Source/Console/radio.cs"
```

Confirm: exactly three additions, each inside its property's `set` block
after the existing `if (update) { ... }` block, each with matched braces,
and no other lines touched.

- [ ] **Step 4: Commit**

```bash
git add "Project Files/Source/Console/radio.cs"
git commit -m "feat(radae): interlock RX1/TX1 digital voice mode at the property level

RXAFDVRun (700E) and RXRadaeEnabled (RADE), and TXAFDVRun (700E TX)
against SetRadaeTxEnabled (RADE TX), now disarm each other whenever
either is armed -- enforced in the radio.cs property setters
themselves, not just in a UI control, so it holds regardless of entry
point (new Mode combo in a later task, or the existing per-subsystem
CAT commands).

Part of sub-project 2 of 5, see
docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
git push
```

---

### Task 2: CAT layer — unified mode command + ZZDK cross-disarm

**Files:**
- Modify: `Project Files/Source/Console/CAT/CATCommands.cs` (add method `ZZEX`, modify method `ZZDK`)
- Modify: `Project Files/Source/Console/CAT/CATParser.cs` (add one `case "ZZEX":` to the extended-command switch)
- Modify: `Project Files/Source/Console/CAT/CATStructs.xml` (add one `<catstruct code="ZZEX">` entry)

**Interfaces:**
- Consumes: `console.radio.GetDSPRX(0, 0).RXAFDVRun`/`.RXRadaeEnabled` and `console.radio.GetDSPTX(0).TXAFDVRun` (existing radio.cs properties, now interlocked per Task 1), `WDSP.SetRadaeTxEnabled`/`GetRadaeTxEnabled`, `WDSP.SetRadaeProtocolV2(int rx, int on)`/`GetRadaeProtocolV2(int rx)` (all pre-existing `dsp.cs` P/Invoke declarations), `parser.nSet`/`nGet`/`Error1` (existing `CATCommands` base members, same as every other command in the file).
- Produces: `public string ZZEX(string s)` — set: `"0"`/`"1"`/`"2"`/`"3"` (Off/700E/RADE V1/RADE V2), arms the matching subsystem (relying on Task 1's interlock to clear the other side); get: returns the currently-active mode as one of those four digits, derived the same way Task 3's `InitRadePanelFromBackend` derives `cmbRadeMode`'s initial index. `ZZDK`'s existing signature (`public string ZZDK(string s)`) is unchanged; only its set-branch body gains one line.

- [ ] **Step 1: Write the CAT-layer splicing script**

Create `scratchpad/splice_cat_zzex.py`:

```python
#!/usr/bin/env python3
"""Add ZZEX (unified Digital Voice mode command) and the ZZDK cross-disarm
fix, byte-exact, preserving each file's CRLF convention (the RADE cluster
in CATCommands.cs and all of CATParser.cs/CATStructs.xml are CRLF)."""

ROOT = "Project Files/Source/Console/CAT/"

def crlf(text):
    return text.replace("\r\n", "\n").replace("\n", "\r\n")

# --- 1. CATCommands.cs: modify ZZDK, add ZZEX right after ZZEW -----------

old_zzdk = crlf('''        public string ZZDK(string s)
        {
            if (s.Length == parser.nSet && (s == "0" || s == "1"))
            {
                WDSP.SetRadaeTxEnabled((s == "1") ? 1 : 0);
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                return (WDSP.GetRadaeTxEnabled() != 0) ? "1" : "0";
            }
            else
            {
                return parser.Error1;
            }
        }
''')

new_zzdk = crlf('''        public string ZZDK(string s)
        {
            if (s.Length == parser.nSet && (s == "0" || s == "1"))
            {
                if (s == "1")
                {
                    // W5TSU: Digital Voice mode interlock -- arming RADE TX
                    // always disarms 700E TX (see radio.cs TXAFDVRun's own
                    // interlock for the reverse direction).
                    console.radio.GetDSPTX(0).TXAFDVRun = 0;
                }
                WDSP.SetRadaeTxEnabled((s == "1") ? 1 : 0);
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                return (WDSP.GetRadaeTxEnabled() != 0) ? "1" : "0";
            }
            else
            {
                return parser.Error1;
            }
        }
''')

zzex_method = crlf('''        // Reads or sets the unified Digital Voice mode index (sub-project 2 of
        // 5, see docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md)
        // -- Setup DSP/Digital Voice panel's Mode combo (cmbRadeMode). 0=Off,
        // 1=700E, 2=RADE V1, 3=RADE V2. Arms the selected subsystem's RX1
        // decode + TX encode together; arming one side always clears the
        // other via the low-level interlock in radio.cs's RXAFDVRun/
        // RXRadaeEnabled/TXAFDVRun setters, so this only needs to arm the
        // target (or explicitly disarm everything for Off) -- same
        // simplification cmbRadeMode_SelectedIndexChanged uses. The existing
        // per-subsystem commands (ZZDV/ZZEF/ZZEG/ZZDW/ZZDK/ZZDL/ZZEP) are
        // untouched and still work directly. // W5TSU
        public string ZZEX(string s)
        {
            if (s.Length == parser.nSet && (s == "0" || s == "1" || s == "2" || s == "3"))
            {
                switch (s)
                {
                    case "0": // Off
                        console.radio.GetDSPRX(0, 0).RXAFDVRun = 0;
                        console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 0;
                        console.radio.GetDSPTX(0).TXAFDVRun = 0;
                        WDSP.SetRadaeTxEnabled(0);
                        break;
                    case "1": // 700E
                        console.radio.GetDSPRX(0, 0).RXAFDVRun = 1;
                        console.radio.GetDSPTX(0).TXAFDVRun = 1;
                        break;
                    case "2": // RADE V1
                        WDSP.SetRadaeProtocolV2(0, 0);
                        console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 1;
                        console.radio.GetDSPTX(0).TXAFDVRun = 0;
                        WDSP.SetRadaeTxEnabled(1);
                        break;
                    case "3": // RADE V2
                        WDSP.SetRadaeProtocolV2(0, 1);
                        console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 1;
                        console.radio.GetDSPTX(0).TXAFDVRun = 0;
                        WDSP.SetRadaeTxEnabled(1);
                        break;
                }
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                if (console.radio.GetDSPRX(0, 0).RXRadaeEnabled != 0)
                    return (WDSP.GetRadaeProtocolV2(0) != 0) ? "3" : "2";
                else if (console.radio.GetDSPRX(0, 0).RXAFDVRun != 0)
                    return "1";
                else
                    return "0";
            }
            else
            {
                return parser.Error1;
            }
        }
''')

anchor_zzew_tail = crlf('''            else
            {
                return parser.Error1;
            }
        }
''') + "        /// <summary>\n        /// Sets or reads the VAC Stereo checkbox\n"

path = ROOT + "CATCommands.cs"
data = open(path, "rb").read()

old_b = old_zzdk.encode("utf-8")
count = data.count(old_b)
assert count == 1, f"ZZDK anchor found {count} times, expected 1"
data = data.replace(old_b, new_zzdk.encode("utf-8"))

anchor_b = anchor_zzew_tail.encode("utf-8")
count = data.count(anchor_b)
assert count == 1, f"ZZEW-tail anchor found {count} times, expected 1"
replacement_b = (crlf('''            else
            {
                return parser.Error1;
            }
        }
''') + zzex_method + "        /// <summary>\n        /// Sets or reads the VAC Stereo checkbox\n").encode("utf-8")
data = data.replace(anchor_b, replacement_b)

open(path, "wb").write(data)
print("CATCommands.cs: ZZDK modified, ZZEX inserted")

# --- 2. CATParser.cs: one new switch case ---------------------------------

case_zzex = crlf('''                case "ZZEX":
                    rtncmd = cmdlist.ZZEX(suffix);
                    break;
''')

anchor_parser = crlf('''                case "ZZEW":
                    rtncmd = cmdlist.ZZEW(suffix);
                    break;
                case "ZZTC":
''')
replacement_parser = crlf('''                case "ZZEW":
                    rtncmd = cmdlist.ZZEW(suffix);
                    break;
''') + case_zzex + crlf('''                case "ZZTC":
''')

path = ROOT + "CATParser.cs"
data = open(path, "rb").read()
anchor_b = anchor_parser.encode("utf-8")
count = data.count(anchor_b)
assert count == 1, f"CATParser.cs anchor found {count} times, expected 1"
data = data.replace(anchor_b, replacement_parser.encode("utf-8"))
open(path, "wb").write(data)
print("CATParser.cs: ZZEX case inserted")

# --- 3. CATStructs.xml: one new catstruct entry ---------------------------

struct_zzex = crlf('''  <catstruct code="ZZEX">
    <desc>Digital Voice unified mode index: 0=Off 1=700E 2=RADE V1 3=RADE V2</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
''')

anchor_xml = crlf('''  <catstruct code="ZZEW">
    <desc>RADE diagnostics bypass ALL status</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
  <catstruct code="ZZTC">
''')
replacement_xml = crlf('''  <catstruct code="ZZEW">
    <desc>RADE diagnostics bypass ALL status</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
''') + struct_zzex + crlf('''  <catstruct code="ZZTC">
''')

path = ROOT + "CATStructs.xml"
data = open(path, "rb").read()
anchor_b = anchor_xml.encode("utf-8")
count = data.count(anchor_b)
assert count == 1, f"CATStructs.xml anchor found {count} times, expected 1"
data = data.replace(anchor_b, replacement_xml.encode("utf-8"))
open(path, "wb").write(data)
print("CATStructs.xml: ZZEX catstruct inserted")
```

- [ ] **Step 2: Run it and verify each diff**

```bash
python3 scratchpad/splice_cat_zzex.py
git diff --stat "Project Files/Source/Console/CAT/CATCommands.cs" \
                "Project Files/Source/Console/CAT/CATParser.cs" \
                "Project Files/Source/Console/CAT/CATStructs.xml"
```

Expected: all three "inserted"/"modified" print lines, and `git diff --stat`
shows roughly `CATCommands.cs +67/-1`, `CATParser.cs +3`, `CATStructs.xml
+7` — not hundreds/thousands of lines. If any file shows a huge diff, run
`git checkout -- <that file>` and re-run after fixing the script.

- [ ] **Step 3: Commit**

```bash
git add "Project Files/Source/Console/CAT/CATCommands.cs" \
        "Project Files/Source/Console/CAT/CATParser.cs" \
        "Project Files/Source/Console/CAT/CATStructs.xml"
git commit -m "feat(radae): ZZEX unified Digital Voice mode CAT command

One new command (ZZEX: 0=Off/1=700E/2=RADE V1/3=RADE V2) arms the
selected subsystem's RX1 decode + TX encode together, relying on
Task 1's radio.cs interlock to clear whichever subsystem was
previously active. ZZDK (RADE TX enable) gets one added line so
arming RADE TX also disarms 700E TX directly, closing the same gap
from the CAT side that Task 1 closed for RXAFDVRun/RXRadaeEnabled.
Existing per-subsystem commands (ZZDV/ZZEF/ZZEG/ZZDW/ZZDL/ZZEP) are
untouched.

Part of sub-project 2 of 5, see
docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
git push
```

---

### Task 3: Digital Voice tab — UI layout + handler rewiring

This is one task, not two, because `setup.designer.cs` (control definitions
and event-handler *names*) and `setup.cs` (the handler *bodies* those names
point to) must agree with each other to compile — splitting them would leave
an intermediate state that doesn't build.

**Files:**
- Modify: `Project Files/Source/Console/setup.designer.cs` (remove `tpDSPFreeDV`/`grpFreeDV`/`chkFreeDVDecode`/`lblFreeDVStatus`; remove `grpRadeProtocol`; repurpose `lblRadeProtocol`→`lblRadeMode`, `cmbRadeProtocol`→`cmbRadeMode` and relocate both into `grpRadeRX1Core`; remove `chkRadeRX1Enable`; retext `tpDSPRADE`'s tab label to "Digital Voice")
- Modify: `Project Files/Source/Console/setup.cs` (remove `chkFreeDVDecode_CheckedChanged`, `freedvStatusTimer_Tick`, `_freedv_status_timer`, `chkRadeRX1Enable_CheckedChanged`; replace `cmbRadeProtocol_SelectedIndexChanged` with `cmbRadeMode_SelectedIndexChanged`; rewrite `radeStatusTimer_Tick` and `chkRadeRX1Loopback_CheckedChanged` to be mode-aware; rewrite `InitRadePanelFromBackend` to derive the mode and gray-out state; update the tab-load call site)

**Interfaces:**
- Consumes: Task 1's interlocked `RXAFDVRun`/`RXRadaeEnabled`/`TXAFDVRun`, `WDSP.SetRadaeTxEnabled`/`GetRadaeTxEnabled`, `WDSP.SetRadaeProtocolV2`/`GetRadaeProtocolV2`, `WDSP.SetRadaeLoopbackEnabled(int rx, int enable)`/`GetRadaeLoopbackEnabled(int rx)`, `WDSP.SetFDVLoopbackEnabled(int enable)`/`GetFDVLoopbackEnabled()`, `WDSP.GetRadaeSync(int rx)`/`GetRadaeSnrDb(int rx)`, `WDSP.GetRXAFDVSync(int channel)`/`GetRXAFDVSnr(int channel)`, `WDSP.id(int thread, int subrx)` — all pre-existing.
- Produces: `cmbRadeMode` (public via designer field, `SelectedIndex` 0-3), `lblRadeMode`, and the rewritten handlers below — nothing outside this file references any of them, so no other file changes are needed for this task.

- [ ] **Step 1: Write the designer.cs splicing script**

Create `scratchpad/splice_designer_dvmode.py`:

```python
#!/usr/bin/env python3
"""Fold 700E into the RADE tab's Mode selector: remove the old FreeDV tab
and its controls, remove the standalone Protocol group, repurpose its
label+combo into cmbRadeMode inside RX1 Core, remove the old RX1 Enable
checkbox. setup.designer.cs is 100% CRLF -- splice byte-exact."""

path = "Project Files/Source/Console/setup.designer.cs"

def crlf(text):
    return text.replace("\r\n", "\n").replace("\n", "\r\n")

replacements = []  # list of (name, old, new), applied in order

# 1. Field declarations -- tpDSPFreeDV
replacements.append(("field: tpDSPFreeDV", crlf('''        private System.Windows.Forms.TabPage tpDSPFM;
        private System.Windows.Forms.TabPage tpDSPFreeDV;
        private System.Windows.Forms.GroupBoxTS grpFMTX;
'''), crlf('''        private System.Windows.Forms.TabPage tpDSPFM;
        private System.Windows.Forms.GroupBoxTS grpFMTX;
''')))

# 2. Field declarations -- grpFreeDV/chkFreeDVDecode/lblFreeDVStatus
replacements.append(("field: grpFreeDV cluster", crlf('''        private GroupBoxTS grpRNnoise;
        private GroupBoxTS grpFreeDV;
        private CheckBoxTS chkFreeDVDecode;
        private LabelTS lblFreeDVStatus;
        private System.Windows.Forms.TabPage tpDSPRADE;
'''), crlf('''        private GroupBoxTS grpRNnoise;
        private System.Windows.Forms.TabPage tpDSPRADE;
''')))

# 3. Field declarations -- RX1Core/Protocol cluster
replacements.append(("field: RX1Core/Protocol cluster", crlf('''        private GroupBoxTS grpRadeRX1Core;
        private CheckBoxTS chkRadeRX1Enable;
        private CheckBoxTS chkRadeRX1Loopback;
        private LabelTS lblRadeRxLevel;
        private NumericUpDownTS udRadeRxLevel;
        private LabelTS lblRadeRX1Status;
        private GroupBoxTS grpRadeProtocol;
        private LabelTS lblRadeProtocol;
        private ComboBoxTS cmbRadeProtocol;
'''), crlf('''        private GroupBoxTS grpRadeRX1Core;
        private LabelTS lblRadeMode;
        private ComboBoxTS cmbRadeMode;
        private CheckBoxTS chkRadeRX1Loopback;
        private LabelTS lblRadeRxLevel;
        private NumericUpDownTS udRadeRxLevel;
        private LabelTS lblRadeRX1Status;
''')))

# 4. Instantiation -- tpDSPFreeDV
replacements.append(("instantiate: tpDSPFreeDV", crlf('''            this.tpDSPFM = new System.Windows.Forms.TabPage();
            this.tpDSPFreeDV = new System.Windows.Forms.TabPage();
'''), crlf('''            this.tpDSPFM = new System.Windows.Forms.TabPage();
''')))

# 5. Instantiation -- grpFreeDV cluster
replacements.append(("instantiate: grpFreeDV cluster", crlf('''            this.grpRNnoise = new System.Windows.Forms.GroupBoxTS();
            this.grpFreeDV = new System.Windows.Forms.GroupBoxTS();
            this.chkFreeDVDecode = new System.Windows.Forms.CheckBoxTS();
            this.lblFreeDVStatus = new System.Windows.Forms.LabelTS();
            this.tpDSPRADE = new System.Windows.Forms.TabPage();
'''), crlf('''            this.grpRNnoise = new System.Windows.Forms.GroupBoxTS();
            this.tpDSPRADE = new System.Windows.Forms.TabPage();
''')))

# 6. Instantiation -- RX1Core/Protocol cluster
replacements.append(("instantiate: RX1Core/Protocol cluster", crlf('''            this.grpRadeRX1Core = new System.Windows.Forms.GroupBoxTS();
            this.chkRadeRX1Enable = new System.Windows.Forms.CheckBoxTS();
            this.chkRadeRX1Loopback = new System.Windows.Forms.CheckBoxTS();
            this.lblRadeRxLevel = new System.Windows.Forms.LabelTS();
            this.udRadeRxLevel = new System.Windows.Forms.NumericUpDownTS();
            this.lblRadeRX1Status = new System.Windows.Forms.LabelTS();
            this.grpRadeProtocol = new System.Windows.Forms.GroupBoxTS();
            this.lblRadeProtocol = new System.Windows.Forms.LabelTS();
            this.cmbRadeProtocol = new System.Windows.Forms.ComboBoxTS();
'''), crlf('''            this.grpRadeRX1Core = new System.Windows.Forms.GroupBoxTS();
            this.lblRadeMode = new System.Windows.Forms.LabelTS();
            this.cmbRadeMode = new System.Windows.Forms.ComboBoxTS();
            this.chkRadeRX1Loopback = new System.Windows.Forms.CheckBoxTS();
            this.lblRadeRxLevel = new System.Windows.Forms.LabelTS();
            this.udRadeRxLevel = new System.Windows.Forms.NumericUpDownTS();
            this.lblRadeRX1Status = new System.Windows.Forms.LabelTS();
''')))

# 7. SuspendLayout -- tpDSPFreeDV
replacements.append(("suspend: tpDSPFreeDV", crlf('''            this.tpDSPFM.SuspendLayout();
            this.tpDSPFreeDV.SuspendLayout();
            this.grpFMRX.SuspendLayout();
'''), crlf('''            this.tpDSPFM.SuspendLayout();
            this.grpFMRX.SuspendLayout();
''')))

# 8. SuspendLayout -- grpFreeDV/grpRadeProtocol
replacements.append(("suspend: grpFreeDV/grpRadeProtocol", crlf('''            this.grpRNnoise.SuspendLayout();
            this.grpFreeDV.SuspendLayout();
            this.tpDSPRADE.SuspendLayout();
            this.grpRadeMicCond.SuspendLayout();
            this.grpRadeRX1Core.SuspendLayout();
            this.grpRadeProtocol.SuspendLayout();
            this.grpRadeDiagnostics.SuspendLayout();
'''), crlf('''            this.grpRNnoise.SuspendLayout();
            this.tpDSPRADE.SuspendLayout();
            this.grpRadeMicCond.SuspendLayout();
            this.grpRadeRX1Core.SuspendLayout();
            this.grpRadeDiagnostics.SuspendLayout();
''')))

# 9. tcDSP.Controls.Add -- drop tpDSPFreeDV registration
replacements.append(("tcDSP.Controls.Add", crlf('''            this.tcDSP.Controls.Add(this.tpDSPFM);
            this.tcDSP.Controls.Add(this.tpDSPFreeDV);
            this.tcDSP.Controls.Add(this.tpDSPRADE);
'''), crlf('''            this.tcDSP.Controls.Add(this.tpDSPFM);
            this.tcDSP.Controls.Add(this.tpDSPRADE);
''')))

# 10. Properties block -- remove tpDSPFreeDV's own property block
replacements.append(("properties: tpDSPFreeDV block", crlf('''            this.chkEmphPos.CheckedChanged += new System.EventHandler(this.chkEmphPos_CheckedChanged);
            //
            // tpDSPFreeDV
            //
            this.tpDSPFreeDV.BackColor = System.Drawing.SystemColors.Control;
            this.tpDSPFreeDV.Controls.Add(this.grpFreeDV);
            this.tpDSPFreeDV.Location = new System.Drawing.Point(4, 22);
            this.tpDSPFreeDV.Name = "tpDSPFreeDV";
            this.tpDSPFreeDV.Padding = new System.Windows.Forms.Padding(3);
            this.tpDSPFreeDV.Size = new System.Drawing.Size(724, 414);
            this.tpDSPFreeDV.TabIndex = 12;
            this.tpDSPFreeDV.Text = "FreeDV";
            //
            // tpDSPAudio
'''), crlf('''            this.chkEmphPos.CheckedChanged += new System.EventHandler(this.chkEmphPos_CheckedChanged);
            //
            // tpDSPAudio
''')))

# 11. Properties block -- remove grpFreeDV/chkFreeDVDecode/lblFreeDVStatus
replacements.append(("properties: grpFreeDV cluster", crlf('''            this.grpRNnoise.Text = "NR3";
            //
            // grpFreeDV
            //
            this.grpFreeDV.Controls.Add(this.chkFreeDVDecode);
            this.grpFreeDV.Controls.Add(this.lblFreeDVStatus);
            this.grpFreeDV.Location = new System.Drawing.Point(16, 16);
            this.grpFreeDV.Name = "grpFreeDV";
            this.grpFreeDV.Size = new System.Drawing.Size(235, 76);
            this.grpFreeDV.TabIndex = 44;
            this.grpFreeDV.TabStop = false;
            this.grpFreeDV.Text = "FreeDV (prototype)";
            //
            // chkFreeDVDecode
            //
            this.chkFreeDVDecode.AutoSize = true;
            this.chkFreeDVDecode.Image = null;
            this.chkFreeDVDecode.Location = new System.Drawing.Point(16, 22);
            this.chkFreeDVDecode.Name = "chkFreeDVDecode";
            this.chkFreeDVDecode.Size = new System.Drawing.Size(160, 17);
            this.chkFreeDVDecode.TabIndex = 0;
            this.chkFreeDVDecode.Text = "Decode FreeDV 700E (RX1)";
            this.toolTip1.SetToolTip(this.chkFreeDVDecode, "Decode FreeDV 700E digital voice on RX1. Use USB/DIGU with a ~3kHz filter; modem audio passes through until sync.");
            this.chkFreeDVDecode.UseVisualStyleBackColor = true;
            this.chkFreeDVDecode.CheckedChanged += new System.EventHandler(this.chkFreeDVDecode_CheckedChanged);
            //
            // lblFreeDVStatus
            //
            this.lblFreeDVStatus.AutoSize = true;
            this.lblFreeDVStatus.Image = null;
            this.lblFreeDVStatus.Location = new System.Drawing.Point(16, 48);
            this.lblFreeDVStatus.Name = "lblFreeDVStatus";
            this.lblFreeDVStatus.Size = new System.Drawing.Size(21, 13);
            this.lblFreeDVStatus.TabIndex = 1;
            this.lblFreeDVStatus.Text = "off";
            //
'''), crlf('''            this.grpRNnoise.Text = "NR3";
            //
''')))

# 12. Properties block -- RX1Core header + chkRadeRX1Enable -> lblRadeMode/cmbRadeMode,
#     remove grpRadeProtocol's own header, keep+relocate+retext lblRadeProtocol/cmbRadeProtocol
replacements.append(("properties: RX1Core+Protocol big block", crlf('''            this.udRadeEQVol.ValueChanged += new System.EventHandler(this.udRadeEQVol_ValueChanged);
            //
            // grpRadeRX1Core
            //
            this.grpRadeRX1Core.Controls.Add(this.chkRadeRX1Enable);
            this.grpRadeRX1Core.Controls.Add(this.chkRadeRX1Loopback);
            this.grpRadeRX1Core.Controls.Add(this.lblRadeRxLevel);
            this.grpRadeRX1Core.Controls.Add(this.udRadeRxLevel);
            this.grpRadeRX1Core.Controls.Add(this.lblRadeRX1Status);
            this.grpRadeRX1Core.Location = new System.Drawing.Point(352, 16);
            this.grpRadeRX1Core.Name = "grpRadeRX1Core";
            this.grpRadeRX1Core.Size = new System.Drawing.Size(340, 136);
            this.grpRadeRX1Core.TabIndex = 47;
            this.grpRadeRX1Core.TabStop = false;
            this.grpRadeRX1Core.Text = "RX1 Core";
            //
            // chkRadeRX1Enable
            //
            this.chkRadeRX1Enable.AutoSize = true;
            this.chkRadeRX1Enable.Image = null;
            this.chkRadeRX1Enable.Location = new System.Drawing.Point(16, 22);
            this.chkRadeRX1Enable.Name = "chkRadeRX1Enable";
            this.chkRadeRX1Enable.Size = new System.Drawing.Size(90, 17);
            this.chkRadeRX1Enable.TabIndex = 0;
            this.chkRadeRX1Enable.Text = "RX1 RADE Enable";
            this.toolTip1.SetToolTip(this.chkRadeRX1Enable, "RX audio fed to the RADE decoder before it reaches the speakers; mic audio fed to the RADE encoder after WDSP audio enhancements (TXEQ/Compander/CFC/Phase/Leveler). Mode/filter selection is not changed automatically.");
            this.chkRadeRX1Enable.UseVisualStyleBackColor = true;
            this.chkRadeRX1Enable.CheckedChanged += new System.EventHandler(this.chkRadeRX1Enable_CheckedChanged);
            //
            // chkRadeRX1Loopback
            //
            this.chkRadeRX1Loopback.AutoSize = true;
            this.chkRadeRX1Loopback.Image = null;
            this.chkRadeRX1Loopback.Location = new System.Drawing.Point(16, 48);
            this.chkRadeRX1Loopback.Name = "chkRadeRX1Loopback";
            this.chkRadeRX1Loopback.Size = new System.Drawing.Size(132, 17);
            this.chkRadeRX1Loopback.TabIndex = 1;
            this.chkRadeRX1Loopback.Text = "RX1 RADE Loopback Test";
            this.toolTip1.SetToolTip(this.chkRadeRX1Loopback, "TX encoder's modem output is bridged directly into RX1's decoder input -- no RF, radio never keys. For verifying the encode/decode round trip before any real on-air attempt.");
            this.chkRadeRX1Loopback.UseVisualStyleBackColor = true;
            this.chkRadeRX1Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX1Loopback_CheckedChanged);
            //
            // lblRadeRxLevel
            //
            this.lblRadeRxLevel.AutoSize = true;
            this.lblRadeRxLevel.Image = null;
            this.lblRadeRxLevel.Location = new System.Drawing.Point(16, 76);
            this.lblRadeRxLevel.Name = "lblRadeRxLevel";
            this.lblRadeRxLevel.Size = new System.Drawing.Size(84, 13);
            this.lblRadeRxLevel.TabIndex = 0;
            this.lblRadeRxLevel.Text = "RX Level (dB):";
            //
            // udRadeRxLevel
            //
            this.udRadeRxLevel.Increment = new decimal(new int[] {
            1,
            0,
            0,
            0});
            this.udRadeRxLevel.Location = new System.Drawing.Point(140, 74);
            this.udRadeRxLevel.Maximum = new decimal(new int[] {
            40,
            0,
            0,
            0});
            this.udRadeRxLevel.Minimum = new decimal(new int[] {
            40,
            0,
            0,
            -2147483648});
            this.udRadeRxLevel.Name = "udRadeRxLevel";
            this.udRadeRxLevel.Size = new System.Drawing.Size(50, 20);
            this.udRadeRxLevel.TabIndex = 2;
            this.udRadeRxLevel.TinyStep = false;
            this.toolTip1.SetToolTip(this.udRadeRxLevel, "Decoder input gain. Default 0 dB.");
            this.udRadeRxLevel.Value = new decimal(new int[] {
            0,
            0,
            0,
            0});
            this.udRadeRxLevel.ValueChanged += new System.EventHandler(this.udRadeRxLevel_ValueChanged);
            //
            // lblRadeRX1Status
            //
            this.lblRadeRX1Status.AutoSize = true;
            this.lblRadeRX1Status.Image = null;
            this.lblRadeRX1Status.Location = new System.Drawing.Point(16, 104);
            this.lblRadeRX1Status.Name = "lblRadeRX1Status";
            this.lblRadeRX1Status.Size = new System.Drawing.Size(21, 13);
            this.lblRadeRX1Status.TabIndex = 0;
            this.lblRadeRX1Status.Text = "off";
            //
            // grpRadeProtocol
            //
            this.grpRadeProtocol.Controls.Add(this.lblRadeProtocol);
            this.grpRadeProtocol.Controls.Add(this.cmbRadeProtocol);
            this.grpRadeProtocol.Location = new System.Drawing.Point(352, 160);
            this.grpRadeProtocol.Name = "grpRadeProtocol";
            this.grpRadeProtocol.Size = new System.Drawing.Size(340, 60);
            this.grpRadeProtocol.TabIndex = 48;
            this.grpRadeProtocol.TabStop = false;
            this.grpRadeProtocol.Text = "Protocol";
            //
            // lblRadeProtocol
            //
            this.lblRadeProtocol.AutoSize = true;
            this.lblRadeProtocol.Image = null;
            this.lblRadeProtocol.Location = new System.Drawing.Point(16, 25);
            this.lblRadeProtocol.Name = "lblRadeProtocol";
            this.lblRadeProtocol.Size = new System.Drawing.Size(84, 13);
            this.lblRadeProtocol.TabIndex = 0;
            this.lblRadeProtocol.Text = "RADE Protocol:";
            //
            // cmbRadeProtocol
            //
            this.cmbRadeProtocol.DropDownStyle = System.Windows.Forms.ComboBoxStyle.DropDownList;
            this.cmbRadeProtocol.Items.AddRange(new object[] {
            "V1", "V2"});
            this.cmbRadeProtocol.Location = new System.Drawing.Point(140, 21);
            this.cmbRadeProtocol.Name = "cmbRadeProtocol";
            this.cmbRadeProtocol.Size = new System.Drawing.Size(80, 21);
            this.cmbRadeProtocol.TabIndex = 0;
            this.toolTip1.SetToolTip(this.cmbRadeProtocol, "Select the RADE protocol version this RX's decoder uses. Live-recycles only this RX's modem handle.");
            this.cmbRadeProtocol.SelectedIndexChanged += new System.EventHandler(this.cmbRadeProtocol_SelectedIndexChanged);
            //
'''), crlf('''            this.udRadeEQVol.ValueChanged += new System.EventHandler(this.udRadeEQVol_ValueChanged);
            //
            // grpRadeRX1Core
            //
            this.grpRadeRX1Core.Controls.Add(this.lblRadeMode);
            this.grpRadeRX1Core.Controls.Add(this.cmbRadeMode);
            this.grpRadeRX1Core.Controls.Add(this.chkRadeRX1Loopback);
            this.grpRadeRX1Core.Controls.Add(this.lblRadeRxLevel);
            this.grpRadeRX1Core.Controls.Add(this.udRadeRxLevel);
            this.grpRadeRX1Core.Controls.Add(this.lblRadeRX1Status);
            this.grpRadeRX1Core.Location = new System.Drawing.Point(352, 16);
            this.grpRadeRX1Core.Name = "grpRadeRX1Core";
            this.grpRadeRX1Core.Size = new System.Drawing.Size(340, 136);
            this.grpRadeRX1Core.TabIndex = 47;
            this.grpRadeRX1Core.TabStop = false;
            this.grpRadeRX1Core.Text = "RX1 Core";
            //
            // lblRadeMode
            //
            this.lblRadeMode.AutoSize = true;
            this.lblRadeMode.Image = null;
            this.lblRadeMode.Location = new System.Drawing.Point(16, 25);
            this.lblRadeMode.Name = "lblRadeMode";
            this.lblRadeMode.Size = new System.Drawing.Size(84, 13);
            this.lblRadeMode.TabIndex = 0;
            this.lblRadeMode.Text = "Mode:";
            //
            // cmbRadeMode
            //
            this.cmbRadeMode.DropDownStyle = System.Windows.Forms.ComboBoxStyle.DropDownList;
            this.cmbRadeMode.Items.AddRange(new object[] {
            "Off", "700E", "RADE V1", "RADE V2"});
            this.cmbRadeMode.Location = new System.Drawing.Point(140, 21);
            this.cmbRadeMode.Name = "cmbRadeMode";
            this.cmbRadeMode.Size = new System.Drawing.Size(140, 21);
            this.cmbRadeMode.TabIndex = 0;
            this.toolTip1.SetToolTip(this.cmbRadeMode, "Select the RX1 digital voice mode: Off, FreeDV 700E, or RADE V1/V2. Arming a mode starts its RX decode and TX encode together; picking a different mode disarms the previous one automatically.");
            this.cmbRadeMode.SelectedIndexChanged += new System.EventHandler(this.cmbRadeMode_SelectedIndexChanged);
            //
            // chkRadeRX1Loopback
            //
            this.chkRadeRX1Loopback.AutoSize = true;
            this.chkRadeRX1Loopback.Image = null;
            this.chkRadeRX1Loopback.Location = new System.Drawing.Point(16, 48);
            this.chkRadeRX1Loopback.Name = "chkRadeRX1Loopback";
            this.chkRadeRX1Loopback.Size = new System.Drawing.Size(132, 17);
            this.chkRadeRX1Loopback.TabIndex = 1;
            this.chkRadeRX1Loopback.Text = "RX1 RADE Loopback Test";
            this.toolTip1.SetToolTip(this.chkRadeRX1Loopback, "TX encoder's modem output is bridged directly into RX1's decoder input -- no RF, radio never keys. For verifying the encode/decode round trip before any real on-air attempt.");
            this.chkRadeRX1Loopback.UseVisualStyleBackColor = true;
            this.chkRadeRX1Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX1Loopback_CheckedChanged);
            //
            // lblRadeRxLevel
            //
            this.lblRadeRxLevel.AutoSize = true;
            this.lblRadeRxLevel.Image = null;
            this.lblRadeRxLevel.Location = new System.Drawing.Point(16, 76);
            this.lblRadeRxLevel.Name = "lblRadeRxLevel";
            this.lblRadeRxLevel.Size = new System.Drawing.Size(84, 13);
            this.lblRadeRxLevel.TabIndex = 0;
            this.lblRadeRxLevel.Text = "RX Level (dB):";
            //
            // udRadeRxLevel
            //
            this.udRadeRxLevel.Increment = new decimal(new int[] {
            1,
            0,
            0,
            0});
            this.udRadeRxLevel.Location = new System.Drawing.Point(140, 74);
            this.udRadeRxLevel.Maximum = new decimal(new int[] {
            40,
            0,
            0,
            0});
            this.udRadeRxLevel.Minimum = new decimal(new int[] {
            40,
            0,
            0,
            -2147483648});
            this.udRadeRxLevel.Name = "udRadeRxLevel";
            this.udRadeRxLevel.Size = new System.Drawing.Size(50, 20);
            this.udRadeRxLevel.TabIndex = 2;
            this.udRadeRxLevel.TinyStep = false;
            this.toolTip1.SetToolTip(this.udRadeRxLevel, "Decoder input gain. Default 0 dB.");
            this.udRadeRxLevel.Value = new decimal(new int[] {
            0,
            0,
            0,
            0});
            this.udRadeRxLevel.ValueChanged += new System.EventHandler(this.udRadeRxLevel_ValueChanged);
            //
            // lblRadeRX1Status
            //
            this.lblRadeRX1Status.AutoSize = true;
            this.lblRadeRX1Status.Image = null;
            this.lblRadeRX1Status.Location = new System.Drawing.Point(16, 104);
            this.lblRadeRX1Status.Name = "lblRadeRX1Status";
            this.lblRadeRX1Status.Size = new System.Drawing.Size(21, 13);
            this.lblRadeRX1Status.TabIndex = 0;
            this.lblRadeRX1Status.Text = "off";
            //
''')))

# 13. tpDSPRADE final registration -- drop grpRadeProtocol, retext tab
replacements.append(("tpDSPRADE registration", crlf('''            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeProtocol);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
            this.tpDSPRADE.Location = new System.Drawing.Point(4, 22);
            this.tpDSPRADE.Name = "tpDSPRADE";
            this.tpDSPRADE.Padding = new System.Windows.Forms.Padding(3);
            this.tpDSPRADE.Size = new System.Drawing.Size(724, 414);
            this.tpDSPRADE.TabIndex = 13;
            this.tpDSPRADE.Text = "RADE";
'''), crlf('''            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
            this.tpDSPRADE.Location = new System.Drawing.Point(4, 22);
            this.tpDSPRADE.Name = "tpDSPRADE";
            this.tpDSPRADE.Padding = new System.Windows.Forms.Padding(3);
            this.tpDSPRADE.Size = new System.Drawing.Size(724, 414);
            this.tpDSPRADE.TabIndex = 13;
            this.tpDSPRADE.Text = "Digital Voice";
''')))

# 14. ResumeLayout -- tpDSPFreeDV
replacements.append(("resume: tpDSPFreeDV", crlf('''            this.tpDSPFM.ResumeLayout(false);
            this.tpDSPFreeDV.ResumeLayout(false);
            this.grpFMRX.ResumeLayout(false);
'''), crlf('''            this.tpDSPFM.ResumeLayout(false);
            this.grpFMRX.ResumeLayout(false);
''')))

# 15. ResumeLayout -- grpFreeDV/grpRadeProtocol
replacements.append(("resume: grpFreeDV/grpRadeProtocol", crlf('''            this.grpRNnoise.ResumeLayout(false);
            this.grpFreeDV.ResumeLayout(false);
            this.grpFreeDV.PerformLayout();
            this.grpRadeMicCond.ResumeLayout(false);
            this.grpRadeMicCond.PerformLayout();
            this.grpRadeRX1Core.ResumeLayout(false);
            this.grpRadeRX1Core.PerformLayout();
            this.grpRadeProtocol.ResumeLayout(false);
            this.grpRadeProtocol.PerformLayout();
            this.grpRadeDiagnostics.ResumeLayout(false);
            this.grpRadeDiagnostics.PerformLayout();
            this.tpDSPRADE.ResumeLayout(false);
'''), crlf('''            this.grpRNnoise.ResumeLayout(false);
            this.grpRadeMicCond.ResumeLayout(false);
            this.grpRadeMicCond.PerformLayout();
            this.grpRadeRX1Core.ResumeLayout(false);
            this.grpRadeRX1Core.PerformLayout();
            this.grpRadeDiagnostics.ResumeLayout(false);
            this.grpRadeDiagnostics.PerformLayout();
            this.tpDSPRADE.ResumeLayout(false);
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

- [ ] **Step 2: Run it and verify the diff**

```bash
python3 scratchpad/splice_designer_dvmode.py
git diff --stat "Project Files/Source/Console/setup.designer.cs"
```

Expected: 15 "replaced" lines followed by `done`, and a diff on the order of
`-90/+55` lines (removals dominate — several whole controls are being
deleted). Sanity-check with a sweep for anything left behind:

```bash
grep -n "chkRADEDecode\|chkRadeRX1Enable\|chkFreeDVDecode\|lblFreeDVStatus\|grpFreeDV\b\|grpRadeProtocol\b\|tpDSPFreeDV\b\|cmbRadeProtocol\b\|lblRadeProtocol\b" \
  "Project Files/Source/Console/setup.designer.cs"
```

Expected: **no output** — every reference to all seven removed/renamed
identifiers should be gone (the old sub-project #1 rename already cleared
`chkRADEDecode`; this step clears everything else).

- [ ] **Step 3: Rewrite setup.cs's handlers**

Open `Project Files/Source/Console/setup.cs`. This file is LF-dominant and
tolerated the plain `Edit` tool cleanly in sub-project #1 — use `Edit`
here, but run `git diff --stat` after and confirm the changed-line count is
in the dozens, not thousands (same landmine check as Task 1/2, just via a
different tool this time).

**3a. Remove the two 700E-only members and the old RX1-Enable handler.**
Find and delete this entire block (currently at `setup.cs:36699-36718`):

```csharp
        // W5TSU: FreeDV RX decode prototype (wdsp fdv.c)
        private System.Windows.Forms.Timer _freedv_status_timer = null;
        private void chkFreeDVDecode_CheckedChanged(object sender, EventArgs e)
        {
            console.radio.GetDSPRX(0, 0).RXAFDVRun = chkFreeDVDecode.Checked ? 1 : 0;

            if (_freedv_status_timer == null)
            {
                _freedv_status_timer = new System.Windows.Forms.Timer();
                _freedv_status_timer.Interval = 500;
                _freedv_status_timer.Tick += freedvStatusTimer_Tick;
            }
            _freedv_status_timer.Enabled = chkFreeDVDecode.Checked;

            if (!chkFreeDVDecode.Checked)
            {
                lblFreeDVStatus.Text = "off";
                lblFreeDVStatus.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

        private void freedvStatusTimer_Tick(object sender, EventArgs e)
        {
            if (!console.PowerOn)
            {
                lblFreeDVStatus.Text = "radio off";
                lblFreeDVStatus.ForeColor = System.Drawing.SystemColors.ControlText;
                return;
            }

            bool sync = WDSP.GetRXAFDVSync(WDSP.id(0, 0)) != 0;
            if (sync)
            {
                double snr = WDSP.GetRXAFDVSnr(WDSP.id(0, 0));
                lblFreeDVStatus.Text = string.Format("SYNC   SNR {0:F1} dB", snr);
                lblFreeDVStatus.ForeColor = System.Drawing.Color.Green;
            }
            else
            {
                lblFreeDVStatus.Text = "no sync";
                lblFreeDVStatus.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

```

**3b. Replace `chkRadeRX1Enable_CheckedChanged`+`radeStatusTimer_Tick` with
the mode-aware version.** Find this block (currently `setup.cs:36743-36797`,
right after the block removed in 3a):

```csharp
        // W5TSU: RADE V1 RX decode prototype (ChannelMaster/radae.c). Same shape as
        // chkFreeDVDecode/freedvStatusTimer_Tick above, but RXRadaeEnabled isn't a
        // wdsp channel setting -- WDSP.GetRadaeSync/GetRadaeSnrDb take ChannelMaster's
        // plain rx index (0 = RX1), not WDSP.id(). No known-working decode confirmed
        // yet (Documentation/FreeDV-Plan.md, Stage C) -- this is a control surface for
        // testing it, not a claim it works.
        private System.Windows.Forms.Timer _rade_status_timer = null;
        // W5TSU: RADE core Setup panel (sub-project 1 of 5, see
        // docs/superpowers/specs/2026-08-24-rade-setup-panel-design.md).
        // chkRADEDecode/lblRADEStatus (the old small grpRADE box on the
        // FreeDV tab) are renamed chkRadeRX1Enable/lblRadeRX1Status and
        // moved into the new dedicated "RADE" tab -- same backend
        // (RXRadaeEnabled), same status-timer logic, just relocated.
        private void chkRadeRX1Enable_CheckedChanged(object sender, EventArgs e)
        {
            console.radio.GetDSPRX(0, 0).RXRadaeEnabled = chkRadeRX1Enable.Checked ? 1 : 0;

            if (_rade_status_timer == null)
            {
                _rade_status_timer = new System.Windows.Forms.Timer();
                _rade_status_timer.Interval = 500;
                _rade_status_timer.Tick += radeStatusTimer_Tick;
            }
            _rade_status_timer.Enabled = chkRadeRX1Enable.Checked;

            if (!chkRadeRX1Enable.Checked)
            {
                lblRadeRX1Status.Text = "off";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

        private void radeStatusTimer_Tick(object sender, EventArgs e)
        {
            if (!console.PowerOn)
            {
                lblRadeRX1Status.Text = "radio off";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
                return;
            }

            bool sync = WDSP.GetRadaeSync(0) != 0;
            if (sync)
            {
                int snr = WDSP.GetRadaeSnrDb(0);
                lblRadeRX1Status.Text = string.Format("SYNC   SNR {0} dB", snr);
                lblRadeRX1Status.ForeColor = System.Drawing.Color.Green;
            }
            else
            {
                lblRadeRX1Status.Text = "no sync";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

```

Replace it with:

```csharp
        // W5TSU: Digital Voice mode selector (sub-project 2 of 5, see
        // docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md).
        // Replaces chkRadeRX1Enable/chkFreeDVDecode (both removed) with one
        // exclusive Mode combo covering Off/700E/RADE V1/RADE V2. Arming a
        // subsystem's RX/TX here relies on the low-level interlock in
        // radio.cs's RXAFDVRun/RXRadaeEnabled/TXAFDVRun setters to disarm
        // whichever subsystem was previously active -- only "Off" needs to
        // disarm everything itself.
        private System.Windows.Forms.Timer _rade_status_timer = null;
        private void cmbRadeMode_SelectedIndexChanged(object sender, EventArgs e)
        {
            int mode = cmbRadeMode.SelectedIndex;

            switch (mode)
            {
                case 0: // Off
                    console.radio.GetDSPRX(0, 0).RXAFDVRun = 0;
                    console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 0;
                    console.radio.GetDSPTX(0).TXAFDVRun = 0;
                    WDSP.SetRadaeTxEnabled(0);
                    break;
                case 1: // 700E
                    console.radio.GetDSPRX(0, 0).RXAFDVRun = 1;
                    console.radio.GetDSPTX(0).TXAFDVRun = 1;
                    break;
                case 2: // RADE V1
                    WDSP.SetRadaeProtocolV2(0, 0);
                    console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 1;
                    console.radio.GetDSPTX(0).TXAFDVRun = 0;
                    WDSP.SetRadaeTxEnabled(1);
                    break;
                case 3: // RADE V2
                    WDSP.SetRadaeProtocolV2(0, 1);
                    console.radio.GetDSPRX(0, 0).RXRadaeEnabled = 1;
                    console.radio.GetDSPTX(0).TXAFDVRun = 0;
                    WDSP.SetRadaeTxEnabled(1);
                    break;
            }

            // RADE-only controls: gray out for Off/700E, enabled for RADE V1/V2.
            bool radeActive = (mode == 2 || mode == 3);
            udRadeRxLevel.Enabled = radeActive;
            grpRadeMicCond.Enabled = radeActive;
            grpRadeDiagnostics.Enabled = radeActive;

            // Loopback Test is meaningful for 700E and RADE alike, just not for Off.
            chkRadeRX1Loopback.Enabled = (mode != 0);
            if (mode == 0 && chkRadeRX1Loopback.Checked)
            {
                chkRadeRX1Loopback.CheckedChanged -= chkRadeRX1Loopback_CheckedChanged;
                chkRadeRX1Loopback.Checked = false;
                chkRadeRX1Loopback.CheckedChanged += chkRadeRX1Loopback_CheckedChanged;
            }

            if (_rade_status_timer == null)
            {
                _rade_status_timer = new System.Windows.Forms.Timer();
                _rade_status_timer.Interval = 500;
                _rade_status_timer.Tick += radeStatusTimer_Tick;
            }
            _rade_status_timer.Enabled = (mode != 0);

            if (mode == 0)
            {
                lblRadeRX1Status.Text = "off";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

        // W5TSU: unified status readout -- reads whichever backend matches
        // the currently selected mode. Same "SYNC   SNR X.X dB" / "no sync"
        // / "radio off" text and color convention chkFreeDVDecode's and
        // chkRadeRX1Enable's separate timers each used before this
        // sub-project merged them into one.
        private void radeStatusTimer_Tick(object sender, EventArgs e)
        {
            if (!console.PowerOn)
            {
                lblRadeRX1Status.Text = "radio off";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
                return;
            }

            int mode = cmbRadeMode.SelectedIndex;
            bool sync;
            string snrText;

            if (mode == 1) // 700E
            {
                sync = WDSP.GetRXAFDVSync(WDSP.id(0, 0)) != 0;
                snrText = sync ? string.Format("{0:F1}", WDSP.GetRXAFDVSnr(WDSP.id(0, 0))) : "";
            }
            else // RADE V1/V2 (mode 2/3) -- the timer is disabled for mode 0, never reaches here then
            {
                sync = WDSP.GetRadaeSync(0) != 0;
                snrText = sync ? string.Format("{0}", WDSP.GetRadaeSnrDb(0)) : "";
            }

            if (sync)
            {
                lblRadeRX1Status.Text = string.Format("SYNC   SNR {0} dB", snrText);
                lblRadeRX1Status.ForeColor = System.Drawing.Color.Green;
            }
            else
            {
                lblRadeRX1Status.Text = "no sync";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

```

**3c. Rewrite `chkRadeRX1Loopback_CheckedChanged` to be mode-aware.** Find
(currently `setup.cs:36798-36804`):

```csharp
        // W5TSU: RX1 RADE loopback bridge (ZZDL's UI equivalent, RX1-only
        // matching this prototype's scope) -- see radae.c's
        // SetRadaeLoopbackEnabled.
        private void chkRadeRX1Loopback_CheckedChanged(object sender, EventArgs e)
        {
            WDSP.SetRadaeLoopbackEnabled(0, chkRadeRX1Loopback.Checked ? 1 : 0);
        }

```

Replace with:

```csharp
        // W5TSU: RX1 loopback bridge -- mode-aware, calls whichever
        // subsystem's loopback matches the current Mode selection
        // (radae.c SetRadaeLoopbackEnabled for RADE V1/V2, ChannelMaster
        // SetFDVLoopbackEnabled for 700E). Disabled entirely for Off, see
        // cmbRadeMode_SelectedIndexChanged.
        private void chkRadeRX1Loopback_CheckedChanged(object sender, EventArgs e)
        {
            int mode = cmbRadeMode.SelectedIndex;
            bool on = chkRadeRX1Loopback.Checked;

            if (mode == 1) // 700E
            {
                WDSP.SetFDVLoopbackEnabled(on ? 1 : 0);
            }
            else if (mode == 2 || mode == 3) // RADE V1/V2
            {
                WDSP.SetRadaeLoopbackEnabled(0, on ? 1 : 0);
            }
        }

```

**3d. Remove `cmbRadeProtocol_SelectedIndexChanged`.** Find (currently
`setup.cs:36896-36900`):

```csharp
        // W5TSU: protocol version selector.
        private void cmbRadeProtocol_SelectedIndexChanged(object sender, EventArgs e)
        {
            WDSP.SetRadaeProtocolV2(0, cmbRadeProtocol.SelectedIndex == 1 ? 1 : 0);
        }

```

Delete it entirely — `cmbRadeMode_SelectedIndexChanged` (added in 3b)
replaces it.

**3e. Rewrite the top of `InitRadePanelFromBackend`.** Find (currently
`setup.cs:36936-36938`, the first three lines of the method body):

```csharp
        private void InitRadePanelFromBackend()
        {
            chkRadeRX1Loopback.CheckedChanged -= chkRadeRX1Loopback_CheckedChanged;
            chkRadeRX1Loopback.Checked = WDSP.GetRadaeLoopbackEnabled(0) != 0;
            chkRadeRX1Loopback.CheckedChanged += chkRadeRX1Loopback_CheckedChanged;

```

Replace with:

```csharp
        private void InitRadePanelFromBackend()
        {
            // W5TSU: derive cmbRadeMode's initial selection from the
            // interlocked backend state (sub-project 2 of 5) --
            // RXRadaeEnabled and RXAFDVRun are mutually exclusive by
            // construction (see radio.cs), so checking RXRadaeEnabled
            // first is unambiguous.
            cmbRadeMode.SelectedIndexChanged -= cmbRadeMode_SelectedIndexChanged;
            int mode;
            if (console.radio.GetDSPRX(0, 0).RXRadaeEnabled != 0)
                mode = (WDSP.GetRadaeProtocolV2(0) != 0) ? 3 : 2;
            else if (console.radio.GetDSPRX(0, 0).RXAFDVRun != 0)
                mode = 1;
            else
                mode = 0;
            cmbRadeMode.SelectedIndex = mode;
            cmbRadeMode.SelectedIndexChanged += cmbRadeMode_SelectedIndexChanged;

            bool radeActive = (mode == 2 || mode == 3);
            udRadeRxLevel.Enabled = radeActive;
            grpRadeMicCond.Enabled = radeActive;
            grpRadeDiagnostics.Enabled = radeActive;
            chkRadeRX1Loopback.Enabled = (mode != 0);

            if (_rade_status_timer == null)
            {
                _rade_status_timer = new System.Windows.Forms.Timer();
                _rade_status_timer.Interval = 500;
                _rade_status_timer.Tick += radeStatusTimer_Tick;
            }
            _rade_status_timer.Enabled = (mode != 0);
            if (mode == 0)
            {
                lblRadeRX1Status.Text = "off";
                lblRadeRX1Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }

            chkRadeRX1Loopback.CheckedChanged -= chkRadeRX1Loopback_CheckedChanged;
            if (mode == 1)
                chkRadeRX1Loopback.Checked = WDSP.GetFDVLoopbackEnabled() != 0;
            else if (mode == 2 || mode == 3)
                chkRadeRX1Loopback.Checked = WDSP.GetRadaeLoopbackEnabled(0) != 0;
            else
                chkRadeRX1Loopback.Checked = false;
            chkRadeRX1Loopback.CheckedChanged += chkRadeRX1Loopback_CheckedChanged;

```

**3f. Remove the now-redundant `cmbRadeProtocol` sync block further down
the same method.** Find (currently `setup.cs:37001-37003`):

```csharp
            cmbRadeProtocol.SelectedIndexChanged -= cmbRadeProtocol_SelectedIndexChanged;
            cmbRadeProtocol.SelectedIndex = WDSP.GetRadaeProtocolV2(0) != 0 ? 1 : 0;
            cmbRadeProtocol.SelectedIndexChanged += cmbRadeProtocol_SelectedIndexChanged;

```

Delete it — mode (which folds in the protocol bit) is now derived once at
the top of the method, per 3e.

**3g. Update the tab-load call site.** Find (currently `setup.cs:2467-2469`):

```csharp
            chkFreeDVDecode_CheckedChanged(this, e); // W5TSU: FreeDV RX decode
            chkRadeRX1Enable_CheckedChanged(this, e); // W5TSU: RADE V1 RX decode
            InitRadePanelFromBackend(); // W5TSU: RADE core Setup panel -- sync all other new controls to backend state
```

Replace with:

```csharp
            InitRadePanelFromBackend(); // W5TSU: Digital Voice Setup panel -- derives Mode + syncs every other control to backend state
```

(`chkFreeDVDecode`/`chkRadeRX1Enable` no longer exist; `InitRadePanelFromBackend`
now does everything both old calls used to do, plus the new mode derivation.)

- [ ] **Step 4: Verify the diff and sweep for leftovers**

```bash
git diff --stat "Project Files/Source/Console/setup.cs"
grep -n "chkFreeDVDecode\|freedvStatusTimer_Tick\|_freedv_status_timer\|chkRadeRX1Enable\b\|cmbRadeProtocol\b" \
  "Project Files/Source/Console/setup.cs"
```

Expected: a diff in the low hundreds of lines (several methods rewritten,
none of them huge), and the `grep` produces **no output**.

- [ ] **Step 5: Commit**

```bash
git add "Project Files/Source/Console/setup.designer.cs" "Project Files/Source/Console/setup.cs"
git commit -m "feat(radae): fold 700E into the Digital Voice Mode selector

cmbRadeProtocol (V1/V2 only) becomes cmbRadeMode (Off/700E/RADE V1/
RADE V2), moved into the RX1 Core group; the standalone Protocol
group box is removed. chkRadeRX1Enable is removed -- enable is now
implicit in the mode selection. The old FreeDV tab (grpFreeDV/
chkFreeDVDecode/lblFreeDVStatus/tpDSPFreeDV) is removed entirely,
same treatment sub-project #1 gave grpRADE. The two separate status
timers merge into one, reading whichever backend matches the active
mode. Loopback Test becomes mode-aware (SetFDVLoopbackEnabled for
700E, SetRadaeLoopbackEnabled for RADE). RADE-only controls (RX
Level, Mic/TX Conditioning, Diagnostics) gray out for Off/700E. Tab
renamed RADE -> Digital Voice.

Part of sub-project 2 of 5, see
docs/superpowers/specs/2026-08-25-digital-voice-mode-selector-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
git push
```

---

### Task 4: Build, deploy, and verify on hl2winbox

**Files:** none (verification only).

**Interfaces:** none — this task consumes the finished state of Tasks 1-3
and produces a verified, working deployment.

- [ ] **Step 1: Trigger the CI build**

```bash
gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
```

Poll until it completes:

```bash
gh run list --workflow=build.yml --branch=FreeDV --limit=1 -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
```

Wait (checking every 20-30s, or via `gh run watch <run-id> -R
W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`) until `status` is `completed`.
Expected: `conclusion` is `success`. If it fails, read the failing step's
log (`gh run view <run-id> --log-failed -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`)
— a compile error here almost always means Task 3's designer.cs/setup.cs
edits are out of sync (a control renamed in one file but not the other);
re-check Task 3's Step 4 grep sweep.

- [ ] **Step 2: Download and deploy the new build**

```bash
RUN_ID=<the run id from Step 1>
gh run download $RUN_ID -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 -n Thetis-HL2-installer -D scratchpad/deploy
scp scratchpad/deploy/Thetis-Test-v2.10.3.x64.msi hl2winbox:Downloads/Thetis-Test-v2.10.3.x64.msi
```

Recreate these four PowerShell scripts in `scratchpad/` if they don't
already exist from a previous sub-project's work, then `scp` each to
`hl2winbox:Downloads/` before running:

`extract.ps1`:
```powershell
$msi = Join-Path $env:USERPROFILE "Downloads\Thetis-Test-v2.10.3.x64.msi"
$dst = Join-Path $env:USERPROFILE "Downloads\Thetis-Extract"
Remove-Item -Recurse -Force $dst -ErrorAction SilentlyContinue
$p = Start-Process msiexec.exe -ArgumentList "/a `"$msi`" /qn TARGETDIR=`"$dst`"" -Wait -PassThru
Write-Output "msiexec exit code: $($p.ExitCode)"
Get-Item (Join-Path $dst "OpenHPSDR\Thetis-Test\Thetis.exe") | Select-Object LastWriteTime,Length
```

`stop_thetis.ps1`:
```powershell
Stop-Process -Name Thetis -Force -ErrorAction SilentlyContinue
Start-Sleep -Seconds 2
$p = Get-Process Thetis -ErrorAction SilentlyContinue
if ($p) { Write-Output "STILL RUNNING: $($p.Id)" } else { Write-Output "STOPPED" }
```

`copy_build.ps1`:
```powershell
$src = Join-Path $env:USERPROFILE "Downloads\Thetis-Extract\OpenHPSDR\Thetis-Test"
$dst = "C:\Program Files\OpenHPSDR\Thetis-Test"
robocopy $src $dst /MIR /R:2 /W:1
Write-Output "robocopy exit code: $LASTEXITCODE"
Get-Item (Join-Path $dst "Thetis.exe") | Select-Object LastWriteTime,Length
```

`relaunch.ps1`:
```powershell
Unregister-ScheduledTask -TaskName ThetisRelaunch -Confirm:$false -ErrorAction SilentlyContinue
$exe = "C:\Program Files\OpenHPSDR\Thetis-Test\Thetis.exe"
$wd  = "C:\Program Files\OpenHPSDR\Thetis-Test"
$action = New-ScheduledTaskAction -Execute $exe -WorkingDirectory $wd
$principal = New-ScheduledTaskPrincipal -UserId $env:USERNAME -LogonType Interactive
Register-ScheduledTask -TaskName ThetisRelaunch -Action $action -Principal $principal | Out-Null
Start-ScheduledTask -TaskName ThetisRelaunch
Start-Sleep -Seconds 6
Get-Process Thetis -ErrorAction SilentlyContinue | Select-Object Id,StartTime,SessionId
Unregister-ScheduledTask -TaskName ThetisRelaunch -Confirm:$false
```

`run_interactive.ps1` (generic launcher — runs any named script inside the
real interactive desktop session via a temporary scheduled task, hidden
window; needed because a plain SSH session runs in Session 0 and can't
launch a real interactive GUI app or take real screenshots):
```powershell
param([string]$ScriptName)
Unregister-ScheduledTask -TaskName ThetisInteractiveRun -Confirm:$false -ErrorAction SilentlyContinue
$scriptPath = Join-Path $env:USERPROFILE "Downloads\$ScriptName"
$action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-WindowStyle Hidden -ExecutionPolicy Bypass -File `"$scriptPath`""
$principal = New-ScheduledTaskPrincipal -UserId $env:USERNAME -LogonType Interactive
Register-ScheduledTask -TaskName ThetisInteractiveRun -Action $action -Principal $principal | Out-Null
Start-ScheduledTask -TaskName ThetisInteractiveRun
Start-Sleep -Seconds 4
Unregister-ScheduledTask -TaskName ThetisInteractiveRun -Confirm:$false
Write-Output "Task ran"
```

Then run the deploy sequence:

```bash
for f in extract.ps1 stop_thetis.ps1 copy_build.ps1 relaunch.ps1 run_interactive.ps1; do
    scp scratchpad/$f hl2winbox:Downloads/$f
done
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\extract.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\stop_thetis.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\copy_build.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\run_interactive.ps1 -ScriptName relaunch.ps1"
```

Expected at each step: `msiexec exit code: 0`, `STOPPED`, `robocopy exit
code: 1` or `3` (both mean success — robocopy's exit codes are a bitmask,
not 0/nonzero), `Task ran`.

- [ ] **Step 3: Verify the new build is live**

```bash
./Tools/thetis-ai-control/thetisctl cat --host 100.117.67.160 --timeout 8s version
```

Expected: the reported `git:` short SHA matches Task 3's commit (not an
older one).

- [ ] **Step 4: CAT round-trip — interlock and ZZEX correctness**

Reuse (or recreate) the raw-socket CAT tester from sub-project #1
(`thetisctl`'s `query` subcommand is get-only and can't send arbitrary set
commands, so this script exists specifically to exercise both directions):

`scratchpad/cat_roundtrip.py`:
```python
#!/usr/bin/env python3
"""Raw CAT-over-TCP round-trip tester for Thetis (Kenwood-style ASCII,
';'-terminated). Mirrors Tools/thetis-ai-control/internal/cat/client.go's
wire protocol."""
import socket
import sys
import time

HOST = "100.117.67.160"
PORT = 13013
TIMEOUT = 5.0

def read_reply(sock, buf):
    sock.settimeout(TIMEOUT)
    while b";" not in buf:
        chunk = sock.recv(4096)
        if not chunk:
            raise ConnectionError("connection closed")
        buf += chunk
    idx = buf.index(b";")
    reply, rest = buf[:idx], buf[idx+1:]
    return reply.decode("ascii", errors="replace"), rest

def main():
    s = socket.create_connection((HOST, PORT), timeout=TIMEOUT)
    buf = b""
    s.settimeout(0.5)
    try:
        while True:
            chunk = s.recv(4096)
            if not chunk:
                break
            buf += chunk
    except socket.timeout:
        pass

    ops = sys.argv[1:]
    i = 0
    while i < len(ops):
        op = ops[i]
        if op.startswith("SET:"):
            cmd = op[4:]
            s.sendall((cmd + ";").encode("ascii"))
            print(f"SET {cmd} -> (no reply expected)")
            time.sleep(0.15)
        elif op.startswith("GET:"):
            code = op[4:]
            s.sendall((code + ";").encode("ascii"))
            for _ in range(8):
                reply, buf = read_reply(s, buf)
                if reply.startswith(code):
                    print(f"GET {code} -> {reply[len(code):]}")
                    break
            else:
                print(f"GET {code} -> NO MATCHING REPLY")
        i += 1
    s.close()

if __name__ == "__main__":
    main()
```

Run the interlock/mode verification sequence:

```bash
python3 scratchpad/cat_roundtrip.py \
  "GET:ZZEX" \
  "SET:ZZEX1" "GET:ZZEX" "GET:ZZDV" "GET:ZZDW" \
  "SET:ZZEX2" "GET:ZZEX" "GET:ZZDW" "GET:ZZDV" "GET:ZZEP" \
  "SET:ZZEX3" "GET:ZZEX" "GET:ZZEP" \
  "SET:ZZDV1" "GET:ZZEX" "GET:ZZDW" \
  "SET:ZZDK1" "GET:ZZDK" \
  "SET:ZZEX0" "GET:ZZEX" "GET:ZZDV" "GET:ZZDW" "GET:ZZDK"
```

Expected, in order:
1. `GET:ZZEX -> 0` (assuming the box was left at Off from prior testing —
   if not, note whatever it reports and adjust the rest of this checklist's
   expectations accordingly, the logic is what's under test, not a specific
   starting value)
2. After `SET:ZZEX1`: `GET:ZZEX -> 1`, `GET:ZZDV -> 1` (700E RX armed),
   `GET:ZZDW -> 0` (RADE RX auto-disarmed by Task 1's interlock)
3. After `SET:ZZEX2`: `GET:ZZEX -> 2`, `GET:ZZDW -> 1` (RADE RX armed),
   `GET:ZZDV -> 0` (700E RX auto-disarmed), `GET:ZZEP -> 0` (protocol V1)
4. After `SET:ZZEX3`: `GET:ZZEX -> 3`, `GET:ZZEP -> 1` (protocol V2)
5. After `SET:ZZDV1` (raw 700E command, bypassing ZZEX entirely): `GET:ZZEX
   -> 1` (ZZEX's own GET correctly reflects the interlocked state even
   though it wasn't the one that changed it), `GET:ZZDW -> 0` (RADE RX
   auto-disarmed by the *same* Task 1 interlock, proving it's enforced low,
   not just inside ZZEX's own logic)
6. After `SET:ZZDK1` (raw RADE TX arm): `GET:ZZDK -> 1`
7. After `SET:ZZEX0`: `GET:ZZEX -> 0`, `GET:ZZDV -> 0`, `GET:ZZDW -> 0`,
   `GET:ZZDK -> 0` (everything disarmed, including the RADE TX armed via
   the raw `ZZDK1` in step 6 — proving Off's explicit disarm-everything
   path works)

If any of these disagree, the interlock or `ZZEX`/`ZZDK` logic has a bug —
stop and fix Task 1 or Task 2 rather than proceeding.

- [ ] **Step 5: Screenshot verification — layout, gray-out, tab removal**

Recreate `screenshot.ps1` if needed:
```powershell
Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing
$bounds = [System.Windows.Forms.SystemInformation]::VirtualScreen
$bmp = New-Object System.Drawing.Bitmap $bounds.Width, $bounds.Height
$g = [System.Drawing.Graphics]::FromImage($bmp)
$g.CopyFromScreen($bounds.Location, [System.Drawing.Point]::Empty, $bounds.Size)
$bmp.Save("$env:USERPROFILE\Downloads\thetis_screenshot.png")
```

And a click helper equivalent to sub-project #1's `close_stray_and_click.ps1`
(minimizes any non-Thetis windowed process first, then clicks the given
coordinates — needed because any script run via `run_interactive.ps1`
itself spawns a visible console window that can cover the target unless
cleared first):
```powershell
Add-Type @"
using System;
using System.Runtime.InteropServices;
public class Win32c {
    [DllImport("user32.dll")] public static extern bool SetCursorPos(int x, int y);
    [DllImport("user32.dll")] public static extern void mouse_event(uint dwFlags, uint dx, uint dy, uint dwData, int dwExtraInfo);
    [DllImport("user32.dll")] public static extern bool SetForegroundWindow(IntPtr hWnd);
    [DllImport("user32.dll")] public static extern bool ShowWindow(IntPtr hWnd, int nCmdShow);
    public const uint MOUSEEVENTF_LEFTDOWN = 0x0002;
    public const uint MOUSEEVENTF_LEFTUP = 0x0004;
    public static void Click(int x, int y) {
        SetCursorPos(x, y);
        mouse_event(MOUSEEVENTF_LEFTDOWN, 0, 0, 0, 0);
        System.Threading.Thread.Sleep(50);
        mouse_event(MOUSEEVENTF_LEFTUP, 0, 0, 0, 0);
    }
}
"@
Get-Process | Where-Object { $_.MainWindowTitle -ne "" -and $_.ProcessName -ne "Thetis" } | ForEach-Object {
    [Win32c]::ShowWindow($_.MainWindowHandle, 6)
}
Start-Sleep -Milliseconds 500
$thetisProc = Get-Process Thetis -ErrorAction SilentlyContinue | Select-Object -First 1
if ($thetisProc -and $thetisProc.MainWindowHandle -ne [IntPtr]::Zero) {
    [Win32c]::ShowWindow($thetisProc.MainWindowHandle, 9)
    [Win32c]::SetForegroundWindow($thetisProc.MainWindowHandle)
}
Start-Sleep -Milliseconds 800
[Win32c]::Click(30, 34)   # Setup button -- coordinates confirmed in sub-project #1's session
Start-Sleep -Seconds 1
```

Deploy and run (Setup → DSP → Digital Voice tab should now be one click
further than sub-project #1's flow, since the click above only opens
Setup — click the "DSP" and "Digital Voice" sub-tabs the same way once
Setup is visible, adjusting coordinates from a screenshot taken between
clicks):

```bash
scp scratchpad/screenshot.ps1 scratchpad/click_setup.ps1 hl2winbox:Downloads/
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\run_interactive.ps1 -ScriptName click_setup.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\run_interactive.ps1 -ScriptName screenshot.ps1"
scp hl2winbox:Downloads/thetis_screenshot.png scratchpad/dv_mode_screenshot.png
```

View `scratchpad/dv_mode_screenshot.png`. Confirm:
- The tab reads "Digital Voice", not "RADE".
- There is no separate "FreeDV" tab anywhere in the DSP tab strip.
- RX1 Core shows a "Mode:" label + 4-item dropdown where "RX1 RADE Enable"
  used to be, followed by the Loopback checkbox, RX Level spinner, and
  status label, in that order.
- With Mode showing "Off" or "700E" (whatever `ZZEX` was left at after
  Step 4's test — re-run `SET:ZZEX0` first if you want a clean Off state
  to screenshot), the RX Level spinner, Mic/TX Conditioning group, and
  Diagnostics group all render visibly grayed out (not hidden).
- Setup opened and rendered without any error dialog.

- [ ] **Step 6: Leave the box in a known state**

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZEX0"
```

This matches sub-project #1's practice of restoring test values afterward
rather than leaving the box mid-experiment.

- [ ] **Step 7: Report results**

No commit for this task (verification-only) — report the CAT round-trip
results from Step 4 and describe the screenshot from Step 5 back to
whoever requested the plan, same as sub-project #1's UI-verification
report.
