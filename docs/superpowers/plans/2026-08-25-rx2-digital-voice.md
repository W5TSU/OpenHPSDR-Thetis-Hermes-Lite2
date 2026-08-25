# RX2 Digital Voice Support Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give RX2 the same Off/700E/RADE V1/RADE V2 Mode selector RX1 has (sub-project #2), scoped to RX2's own decode/loopback/level/status only — no backend changes, since `ChannelMaster/radae.c` and `wdsp/fdv.c` are already genuinely dual-RX by design.

**Architecture:** Pure UI/CAT-layer addition. A new compact "RX2 Core" group box on the existing Digital Voice tab, wired to `console.radio.GetDSPRX(1, 0)` (a separate `RadioDSPRX` instance from RX1's `GetDSPRX(0, 0)`, so the interlock sub-project #2 built into the property setters applies automatically with zero new code) and to the same already-`rx`-parameterized `radae.c`/`fdv.c` WDSP calls RX1 uses, just with `rx=1` instead of `rx=0`. Five new CAT commands mirror the RX1 family. One genuine new-code decision this task makes: RX2's Loopback checkbox is RADE V1/V2 only, not 700E — see Task 1's comments for why (700E's loopback bridge is structurally RX1-only in `pipe.c`).

**Tech Stack:** C# WinForms (.NET Framework 4.8), P/Invoke into `wdsp.dll`/`ChannelMaster.dll` — **zero new P/Invoke declarations needed**; every `WDSP.*` call this plan uses already exists in `dsp.cs` from sub-projects #1/#2, already `rx`-parameterized. No unit test framework exists in this codebase; verification is CI compile-check + live hardware CAT/UI checks (see Global Constraints).

**Spec:** [docs/superpowers/specs/2026-08-25-rx2-digital-voice-design.md](../specs/2026-08-25-rx2-digital-voice-design.md)

## Global Constraints

- **No local Windows build environment and no automated test suite.** Same as sub-projects #1/#2: no `dotnet`/MSBuild available locally, no unit test runner anywhere in the codebase. "Testing" means careful manual diff review for code-only tasks, then one real CI build + `hl2winbox` hardware verification pass at the end.
- **CRLF/LF landmine — read before touching any of these four files:** `CATCommands.cs` (genuinely mixed CRLF/LF), and `CATParser.cs`, `CATStructs.xml`, `setup.designer.cs` (all three 100% CRLF). The plain-text `Edit` tool has been proven to flatten an entire file's line endings when writing back a modified region — this has happened twice already on this branch (once in sub-project #2's CAT-command work, once on this exact repo's `.gitignore` during #2's wrap-up), both times producing a spurious diff that had to be `git checkout`-reverted. For all edits to these four files, use a Python byte-splicing script (`data.replace(old_bytes, new_bytes)` after an `assert data.count(old_bytes) == 1`), never the `Edit` tool. `setup.cs` is LF-dominant and has tolerated plain `Edit`-tool changes cleanly on this branch twice now (sub-projects #1 and #2) — the `Edit` tool is fine there, but run `git diff --stat` after every edit regardless and confirm the changed-line count roughly matches what you intended.
- **Naming convention:** every WinForms control in this codebase uses the (incorrect but established) namespace `System.Windows.Forms.*` for its own thread-safe types (`CheckBoxTS`, `ComboBoxTS`, `GroupBoxTS`, `LabelTS`, `NumericUpDownTS`) — match this exactly, don't "fix" it.
- **The `initializing` guard is required on every new `ComboBoxTS`/`CheckBoxTS` change handler from the start.** This codebase's generic Setup persistence (`SaveOptions`/`getOptions`) restores every named control from disk and fires its change event for real during construction, inside an `initializing = true ... false` window. Sub-project #2 discovered this the hard way (a final-review finding, fixed after the fact) for `cmbRadeMode`. Every new handler this plan adds (`cmbRadeRX2Mode_SelectedIndexChanged`) must start with `if (initializing) return;` — do not rediscover this gap a second time.
- **Deploy pipeline** (established in #1/#2, reused verbatim in Task 3): trigger `gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2` → poll `gh run view --json status,conclusion` → download the `Thetis-HL2-installer` artifact → `scp` to `hl2winbox:Downloads/` → run `extract.ps1` → `stop_thetis.ps1` → `copy_build.ps1` → `relaunch.ps1` via `run_interactive.ps1` (scheduled-task/interactive-session trick — a plain SSH session runs in Session 0 and can't launch a real interactive GUI app or take real screenshots) → verify via `./Tools/thetis-ai-control/thetisctl cat --host 100.117.67.160 --timeout 8s version`. Recreate the `.ps1` scripts in `scratchpad/` if starting a fresh session — their content is given in Task 3.
- **RX2 itself must be powered on for any of this to do anything observable.** The main console has its own "RX2" toggle (`console.RX2Enabled`) independent of Setup; verification steps that expect RX2 decode/status to actually respond need RX2 turned on first.
- Commit after every task; one commit per task minimum.

---

### Task 1: RX2 Core UI — designer.cs + setup.cs

Combined into one task because `setup.designer.cs` (control definitions and
event-handler *names*) and `setup.cs` (the handler *bodies* those names
point to) must agree with each other to compile — splitting them would
leave a broken intermediate build, same reasoning sub-project #2's UI task
used.

**Files:**
- Modify: `Project Files/Source/Console/setup.designer.cs` (add `grpRadeRX2Core` and its 6 children: `lblRadeRX2Mode`, `cmbRadeRX2Mode`, `chkRadeRX2Loopback`, `lblRadeRX2Level`, `udRadeRX2Level`, `lblRadeRX2Status`; register the group in `tpDSPRADE`)
- Modify: `Project Files/Source/Console/setup.cs` (add `cmbRadeRX2Mode_SelectedIndexChanged`, `radeRX2StatusTimer_Tick`, `chkRadeRX2Loopback_CheckedChanged`, `udRadeRX2Level_ValueChanged`, `radeRX2AvailableTimer_Tick`; extend `InitRadePanelFromBackend`)

**Interfaces:**
- Consumes: `console.radio.GetDSPRX(1, 0).RXAFDVRun`/`.RXRadaeEnabled` (existing radio.cs properties — a *different* `RadioDSPRX` instance than RX1's `GetDSPRX(0, 0)`, already interlocked per-instance by sub-project #2, no new radio.cs code needed), `WDSP.SetRadaeProtocolV2(int rx, int on)`/`GetRadaeProtocolV2(int rx)`, `WDSP.SetRadaeLoopbackEnabled(int rx, int enable)`/`GetRadaeLoopbackEnabled(int rx)`, `WDSP.GetRadaeSync(int rx)`/`GetRadaeSnrDb(int rx)`, `WDSP.GetRXAFDVSync(int channel)`/`GetRXAFDVSnr(int channel)`, `WDSP.id(int thread, int subrx)`, `WDSP.SetRadaeRxScale(int rx, double scale)`/`GetRadaeRxScale(int rx)`, `console.RX2Enabled`, `console.PowerOn`, `initializing` (all pre-existing).
- Produces: `grpRadeRX2Core`, `cmbRadeRX2Mode` (`SelectedIndex` 0-3), `chkRadeRX2Loopback`, `udRadeRX2Level`, `lblRadeRX2Status` — Task 2 (CAT layer) does **not** reference any of these; it calls the same backend functions directly, independently, matching the established CAT/UI duplication pattern from #2.

- [ ] **Step 1: Write the designer.cs splicing script**

Create `scratchpad/splice_designer_rx2core.py`:

```python
#!/usr/bin/env python3
"""Add the RX2 Core group (Mode/Loopback/RX Level/Status) to the Digital
Voice tab, filling the existing 76px gap between RX1 Core and Diagnostics.
setup.designer.cs is 100% CRLF -- splice byte-exact."""

path = "Project Files/Source/Console/setup.designer.cs"

def crlf(text):
    return text.replace("\r\n", "\n").replace("\n", "\r\n")

replacements = []

# 1. Field declarations -- insert after chkRadeBypassAll's decl
replacements.append(("field decls", crlf('''        private CheckBoxTS chkRadeBypassAll;

'''), crlf('''        private CheckBoxTS chkRadeBypassAll;
        private GroupBoxTS grpRadeRX2Core;
        private LabelTS lblRadeRX2Mode;
        private ComboBoxTS cmbRadeRX2Mode;
        private CheckBoxTS chkRadeRX2Loopback;
        private LabelTS lblRadeRX2Level;
        private NumericUpDownTS udRadeRX2Level;
        private LabelTS lblRadeRX2Status;

''')))

# 2. Instantiation -- insert after chkRadeBypassAll's instantiation
replacements.append(("instantiation", crlf('''            this.chkRadeBypassAll = new System.Windows.Forms.CheckBoxTS();

            this.chkNR3_RNNoiseFixedGain = new System.Windows.Forms.CheckBoxTS();
'''), crlf('''            this.chkRadeBypassAll = new System.Windows.Forms.CheckBoxTS();
            this.grpRadeRX2Core = new System.Windows.Forms.GroupBoxTS();
            this.lblRadeRX2Mode = new System.Windows.Forms.LabelTS();
            this.cmbRadeRX2Mode = new System.Windows.Forms.ComboBoxTS();
            this.chkRadeRX2Loopback = new System.Windows.Forms.CheckBoxTS();
            this.lblRadeRX2Level = new System.Windows.Forms.LabelTS();
            this.udRadeRX2Level = new System.Windows.Forms.NumericUpDownTS();
            this.lblRadeRX2Status = new System.Windows.Forms.LabelTS();

            this.chkNR3_RNNoiseFixedGain = new System.Windows.Forms.CheckBoxTS();
''')))

# 3. SuspendLayout / BeginInit
replacements.append(("suspend/begininit", crlf('''            this.grpRadeDiagnostics.SuspendLayout();
'''), crlf('''            this.grpRadeDiagnostics.SuspendLayout();
            this.grpRadeRX2Core.SuspendLayout();
''')))
replacements.append(("begininit numericupdown", crlf('''            ((System.ComponentModel.ISupportInitialize)(this.udRadeRxLevel)).BeginInit();
'''), crlf('''            ((System.ComponentModel.ISupportInitialize)(this.udRadeRxLevel)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.udRadeRX2Level)).BeginInit();
''')))

# 4. Properties block -- insert the new group + children after
#    chkRadeBypassAll's property block, before "// tpDSPRADE"
replacements.append(("properties block", crlf('''            this.chkRadeBypassAll.CheckedChanged += new System.EventHandler(this.chkRadeBypassAll_CheckedChanged);
            //
            // tpDSPRADE
            //
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
'''), crlf('''            this.chkRadeBypassAll.CheckedChanged += new System.EventHandler(this.chkRadeBypassAll_CheckedChanged);
            //
            // grpRadeRX2Core
            //
            this.grpRadeRX2Core.Controls.Add(this.lblRadeRX2Mode);
            this.grpRadeRX2Core.Controls.Add(this.cmbRadeRX2Mode);
            this.grpRadeRX2Core.Controls.Add(this.chkRadeRX2Loopback);
            this.grpRadeRX2Core.Controls.Add(this.lblRadeRX2Level);
            this.grpRadeRX2Core.Controls.Add(this.udRadeRX2Level);
            this.grpRadeRX2Core.Controls.Add(this.lblRadeRX2Status);
            this.grpRadeRX2Core.Location = new System.Drawing.Point(352, 152);
            this.grpRadeRX2Core.Name = "grpRadeRX2Core";
            this.grpRadeRX2Core.Size = new System.Drawing.Size(340, 76);
            this.grpRadeRX2Core.TabIndex = 50;
            this.grpRadeRX2Core.TabStop = false;
            this.grpRadeRX2Core.Text = "RX2 Core";
            //
            // lblRadeRX2Mode
            //
            this.lblRadeRX2Mode.AutoSize = true;
            this.lblRadeRX2Mode.Image = null;
            this.lblRadeRX2Mode.Location = new System.Drawing.Point(16, 22);
            this.lblRadeRX2Mode.Name = "lblRadeRX2Mode";
            this.lblRadeRX2Mode.Size = new System.Drawing.Size(84, 13);
            this.lblRadeRX2Mode.TabIndex = 0;
            this.lblRadeRX2Mode.Text = "Mode:";
            //
            // cmbRadeRX2Mode
            //
            this.cmbRadeRX2Mode.DropDownStyle = System.Windows.Forms.ComboBoxStyle.DropDownList;
            this.cmbRadeRX2Mode.Items.AddRange(new object[] {
            "Off", "700E", "RADE V1", "RADE V2"});
            this.cmbRadeRX2Mode.Location = new System.Drawing.Point(60, 20);
            this.cmbRadeRX2Mode.Name = "cmbRadeRX2Mode";
            this.cmbRadeRX2Mode.Size = new System.Drawing.Size(85, 21);
            this.cmbRadeRX2Mode.TabIndex = 0;
            this.toolTip1.SetToolTip(this.cmbRadeRX2Mode, "Select the RX2 digital voice mode: Off, FreeDV 700E, or RADE V1/V2. Arms RX2\\'s own decode only -- TX encode is a single shared resource controlled via RX1\\'s Mode selector, not this one.");
            this.cmbRadeRX2Mode.SelectedIndexChanged += new System.EventHandler(this.cmbRadeRX2Mode_SelectedIndexChanged);
            //
            // chkRadeRX2Loopback
            //
            this.chkRadeRX2Loopback.AutoSize = true;
            this.chkRadeRX2Loopback.Image = null;
            this.chkRadeRX2Loopback.Location = new System.Drawing.Point(155, 22);
            this.chkRadeRX2Loopback.Name = "chkRadeRX2Loopback";
            this.chkRadeRX2Loopback.Size = new System.Drawing.Size(90, 17);
            this.chkRadeRX2Loopback.TabIndex = 1;
            this.chkRadeRX2Loopback.Text = "Loopback";
            this.toolTip1.SetToolTip(this.chkRadeRX2Loopback, "RADE V1/V2 only -- 700E\\'s loopback bridge is RX1-only at the ChannelMaster level, no RX2 equivalent exists. TX encoder\\'s modem output is bridged directly into RX2\\'s decoder input -- no RF, radio never keys.");
            this.chkRadeRX2Loopback.UseVisualStyleBackColor = true;
            this.chkRadeRX2Loopback.CheckedChanged += new System.EventHandler(this.chkRadeRX2Loopback_CheckedChanged);
            //
            // lblRadeRX2Level
            //
            this.lblRadeRX2Level.AutoSize = true;
            this.lblRadeRX2Level.Image = null;
            this.lblRadeRX2Level.Location = new System.Drawing.Point(16, 50);
            this.lblRadeRX2Level.Name = "lblRadeRX2Level";
            this.lblRadeRX2Level.Size = new System.Drawing.Size(84, 13);
            this.lblRadeRX2Level.TabIndex = 0;
            this.lblRadeRX2Level.Text = "RX Lvl (dB):";
            //
            // udRadeRX2Level
            //
            this.udRadeRX2Level.Increment = new decimal(new int[] {
            1,
            0,
            0,
            0});
            this.udRadeRX2Level.Location = new System.Drawing.Point(100, 48);
            this.udRadeRX2Level.Maximum = new decimal(new int[] {
            40,
            0,
            0,
            0});
            this.udRadeRX2Level.Minimum = new decimal(new int[] {
            40,
            0,
            0,
            -2147483648});
            this.udRadeRX2Level.Name = "udRadeRX2Level";
            this.udRadeRX2Level.Size = new System.Drawing.Size(45, 20);
            this.udRadeRX2Level.TabIndex = 2;
            this.udRadeRX2Level.TinyStep = false;
            this.toolTip1.SetToolTip(this.udRadeRX2Level, "RX2 decoder input gain. Default 0 dB.");
            this.udRadeRX2Level.Value = new decimal(new int[] {
            0,
            0,
            0,
            0});
            this.udRadeRX2Level.ValueChanged += new System.EventHandler(this.udRadeRX2Level_ValueChanged);
            //
            // lblRadeRX2Status
            //
            this.lblRadeRX2Status.AutoSize = true;
            this.lblRadeRX2Status.Image = null;
            this.lblRadeRX2Status.Location = new System.Drawing.Point(155, 50);
            this.lblRadeRX2Status.Name = "lblRadeRX2Status";
            this.lblRadeRX2Status.Size = new System.Drawing.Size(21, 13);
            this.lblRadeRX2Status.TabIndex = 0;
            this.lblRadeRX2Status.Text = "off";
            //
            // tpDSPRADE
            //
            this.tpDSPRADE.Controls.Add(this.grpRadeMicCond);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX1Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeRX2Core);
            this.tpDSPRADE.Controls.Add(this.grpRadeDiagnostics);
''')))

# 5. ResumeLayout / EndInit
replacements.append(("resume", crlf('''            this.grpRadeDiagnostics.ResumeLayout(false);
            this.grpRadeDiagnostics.PerformLayout();
            this.tpDSPRADE.ResumeLayout(false);
'''), crlf('''            this.grpRadeDiagnostics.ResumeLayout(false);
            this.grpRadeDiagnostics.PerformLayout();
            this.grpRadeRX2Core.ResumeLayout(false);
            this.grpRadeRX2Core.PerformLayout();
            this.tpDSPRADE.ResumeLayout(false);
''')))
replacements.append(("endinit numericupdown", crlf('''            ((System.ComponentModel.ISupportInitialize)(this.udRadeRxLevel)).EndInit();
'''), crlf('''            ((System.ComponentModel.ISupportInitialize)(this.udRadeRxLevel)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.udRadeRX2Level)).EndInit();
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

Note: the two escaped apostrophes (`\\'`) inside the tooltip strings are
Python's own triple-single-quote string-literal escaping, not a C# escape
(C# double-quoted strings don't need apostrophes escaped at all) — the byte
written to the file is a plain `'`.

- [ ] **Step 2: Run it and verify the diff**

```bash
python3 scratchpad/splice_designer_rx2core.py
git diff --stat "Project Files/Source/Console/setup.designer.cs"
```

Expected: 7 "replaced" lines followed by `done`, and a diff of roughly
`+130/-0` lines (pure addition, no removals — unlike sub-project #2's UI
task, this one only adds controls, it doesn't rename or remove any). If it
shows hundreds/thousands of changed lines instead, the file's line endings
got flattened; `git checkout -- "Project Files/Source/Console/setup.designer.cs"`
and retry after fixing the script.

- [ ] **Step 3: Add the setup.cs handlers**

Open `Project Files/Source/Console/setup.cs`. Find `chkRadeRX1Loopback_CheckedChanged`'s
closing brace (search for the method — it currently ends with a block
matching `WDSP.SetRadaeLoopbackEnabled(0, on ? 1 : 0);` followed by two
closing braces and a blank line). Insert this new block immediately after
that blank line, before the next method's comment (`// W5TSU: RX decoder
input gain...` / `udRadeRxLevel_ValueChanged`):

```csharp
        // W5TSU: RX2 Digital Voice mode selector (sub-project 3 of 5, see
        // docs/superpowers/specs/2026-08-25-rx2-digital-voice-design.md).
        // Mirrors cmbRadeMode_SelectedIndexChanged (RX1) exactly for the
        // decode-arm/disarm logic, but never touches TX -- TX encode is a
        // single shared resource that stays controlled only via RX1's Mode
        // selector (see the spec's "TX semantics" section). RX2's own
        // 700E/RADE interlock still applies automatically: GetDSPRX(1, 0)
        // is a separate RadioDSPRX instance with its own rx_fdv_run/
        // rx_radae_enabled fields, and radio.cs's interlock is already
        // per-instance, so it needs no new code here.
        private System.Windows.Forms.Timer _rade_rx2_status_timer = null;
        private System.Windows.Forms.Timer _rade_rx2_available_timer = null;
        private void cmbRadeRX2Mode_SelectedIndexChanged(object sender, EventArgs e)
        {
            if (initializing) return;

            int mode = cmbRadeRX2Mode.SelectedIndex;

            switch (mode)
            {
                case 0: // Off
                    console.radio.GetDSPRX(1, 0).RXAFDVRun = 0;
                    console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 0;
                    break;
                case 1: // 700E
                    console.radio.GetDSPRX(1, 0).RXAFDVRun = 1;
                    break;
                case 2: // RADE V1
                    WDSP.SetRadaeProtocolV2(1, 0);
                    console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 1;
                    break;
                case 3: // RADE V2
                    WDSP.SetRadaeProtocolV2(1, 1);
                    console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 1;
                    break;
            }

            // W5TSU: Loopback is RADE V1/V2 only for RX2 -- unlike RX1,
            // 700E's loopback bridge (ChannelMaster pipe.c's
            // g_fdvloop_enabled) only drains into RX1's raw antenna I/Q
            // slot; there is no RX2 equivalent. Confirmed by reading
            // pipe.c's xpipe() directly: the "other PowerSDR receivers"
            // branch RX2 runs through has no bridging logic like RX1's
            // "case 0" branch does.
            bool loopbackAvailable = (mode == 2 || mode == 3);
            chkRadeRX2Loopback.Enabled = loopbackAvailable;
            if (!loopbackAvailable && chkRadeRX2Loopback.Checked)
            {
                chkRadeRX2Loopback.CheckedChanged -= chkRadeRX2Loopback_CheckedChanged;
                chkRadeRX2Loopback.Checked = false;
                chkRadeRX2Loopback.CheckedChanged += chkRadeRX2Loopback_CheckedChanged;
            }
            WDSP.SetRadaeLoopbackEnabled(1, 0);
            if (loopbackAvailable && chkRadeRX2Loopback.Checked)
            {
                WDSP.SetRadaeLoopbackEnabled(1, 1);
            }

            if (_rade_rx2_status_timer == null)
            {
                _rade_rx2_status_timer = new System.Windows.Forms.Timer();
                _rade_rx2_status_timer.Interval = 500;
                _rade_rx2_status_timer.Tick += radeRX2StatusTimer_Tick;
            }
            _rade_rx2_status_timer.Enabled = (mode != 0);

            if (mode == 0)
            {
                lblRadeRX2Status.Text = "off";
                lblRadeRX2Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

        // W5TSU: RX2 status readout -- mirrors radeStatusTimer_Tick (RX1)
        // exactly, reading whichever backend matches RX2's own selected mode.
        private void radeRX2StatusTimer_Tick(object sender, EventArgs e)
        {
            if (!console.PowerOn)
            {
                lblRadeRX2Status.Text = "radio off";
                lblRadeRX2Status.ForeColor = System.Drawing.SystemColors.ControlText;
                return;
            }

            int mode = cmbRadeRX2Mode.SelectedIndex;
            bool sync;
            string snrText;

            if (mode == 1) // 700E
            {
                sync = WDSP.GetRXAFDVSync(WDSP.id(1, 0)) != 0;
                snrText = sync ? string.Format("{0:F1}", WDSP.GetRXAFDVSnr(WDSP.id(1, 0))) : "";
            }
            else // RADE V1/V2 (mode 2/3) -- the timer is disabled for mode 0, never reaches here then
            {
                sync = WDSP.GetRadaeSync(1) != 0;
                snrText = sync ? string.Format("{0}", WDSP.GetRadaeSnrDb(1)) : "";
            }

            if (sync)
            {
                lblRadeRX2Status.Text = string.Format("SYNC   SNR {0} dB", snrText);
                lblRadeRX2Status.ForeColor = System.Drawing.Color.Green;
            }
            else
            {
                lblRadeRX2Status.Text = "no sync";
                lblRadeRX2Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }
        }

        // W5TSU: RX2 loopback bridge -- RADE V1/V2 only, see
        // cmbRadeRX2Mode_SelectedIndexChanged's comment for why 700E has
        // no RX2 loopback path.
        private void chkRadeRX2Loopback_CheckedChanged(object sender, EventArgs e)
        {
            int mode = cmbRadeRX2Mode.SelectedIndex;
            if (mode == 2 || mode == 3)
            {
                WDSP.SetRadaeLoopbackEnabled(1, chkRadeRX2Loopback.Checked ? 1 : 0);
            }
        }

        // W5TSU: RX2 decoder input gain -- mirrors udRadeRxLevel_ValueChanged
        // (RX1) exactly, same linear<->dB conversion (see that handler's
        // comment for why).
        private void udRadeRX2Level_ValueChanged(object sender, EventArgs e)
        {
            double linear = Math.Pow(10.0, (double)udRadeRX2Level.Value / 20.0);
            WDSP.SetRadaeRxScale(1, linear);
        }

        // W5TSU: RX2 Core is only meaningful while RX2 itself is running --
        // console.RX2Enabled has no change event to hook, so this polls it
        // on the same 500ms cadence as the status timers. Runs
        // independently of RX2's own mode/status timer (that one only runs
        // while mode != 0), so toggling RX2 on/off in the main console
        // updates this group even while its own Mode is Off.
        private void radeRX2AvailableTimer_Tick(object sender, EventArgs e)
        {
            grpRadeRX2Core.Enabled = console.RX2Enabled;
        }

```

- [ ] **Step 4: Extend `InitRadePanelFromBackend`**

In the same file, find `InitRadePanelFromBackend`'s existing loopback-sync
block for RX1 (ends with `chkRadeRX1Loopback.CheckedChanged +=
chkRadeRX1Loopback_CheckedChanged;`), immediately followed by the comment
`// W5TSU: linear -> dB, inverse of udRadeRxLevel_ValueChanged's...`.
Insert this new block between those two — after RX1's loopback sync ends,
before RX1's mic-level sync begins:

```csharp
            // W5TSU: RX2 Core -- mirrors the RX1 derivation above exactly,
            // but for GetDSPRX(1, 0). See cmbRadeRX2Mode_SelectedIndexChanged's
            // own comment for why TX is never touched here.
            cmbRadeRX2Mode.SelectedIndexChanged -= cmbRadeRX2Mode_SelectedIndexChanged;
            int rx2Mode;
            if (console.radio.GetDSPRX(1, 0).RXRadaeEnabled != 0)
                rx2Mode = (WDSP.GetRadaeProtocolV2(1) != 0) ? 3 : 2;
            else if (console.radio.GetDSPRX(1, 0).RXAFDVRun != 0)
                rx2Mode = 1;
            else
                rx2Mode = 0;
            cmbRadeRX2Mode.SelectedIndex = rx2Mode;
            cmbRadeRX2Mode.SelectedIndexChanged += cmbRadeRX2Mode_SelectedIndexChanged;

            bool rx2LoopbackAvailable = (rx2Mode == 2 || rx2Mode == 3);
            chkRadeRX2Loopback.Enabled = rx2LoopbackAvailable;

            if (_rade_rx2_status_timer == null)
            {
                _rade_rx2_status_timer = new System.Windows.Forms.Timer();
                _rade_rx2_status_timer.Interval = 500;
                _rade_rx2_status_timer.Tick += radeRX2StatusTimer_Tick;
            }
            _rade_rx2_status_timer.Enabled = (rx2Mode != 0);
            if (rx2Mode == 0)
            {
                lblRadeRX2Status.Text = "off";
                lblRadeRX2Status.ForeColor = System.Drawing.SystemColors.ControlText;
            }

            chkRadeRX2Loopback.CheckedChanged -= chkRadeRX2Loopback_CheckedChanged;
            chkRadeRX2Loopback.Checked = rx2LoopbackAvailable && (WDSP.GetRadaeLoopbackEnabled(1) != 0);
            chkRadeRX2Loopback.CheckedChanged += chkRadeRX2Loopback_CheckedChanged;

            udRadeRX2Level.ValueChanged -= udRadeRX2Level_ValueChanged;
            double rx2Db = 20.0 * Math.Log10(Math.Max(0.001, WDSP.GetRadaeRxScale(1)));
            udRadeRX2Level.Value = (decimal)Math.Max(-40.0, Math.Min(40.0, rx2Db));
            udRadeRX2Level.ValueChanged += udRadeRX2Level_ValueChanged;

            if (_rade_rx2_available_timer == null)
            {
                _rade_rx2_available_timer = new System.Windows.Forms.Timer();
                _rade_rx2_available_timer.Interval = 500;
                _rade_rx2_available_timer.Tick += radeRX2AvailableTimer_Tick;
            }
            _rade_rx2_available_timer.Enabled = true;
            grpRadeRX2Core.Enabled = console.RX2Enabled;

```

- [ ] **Step 5: Verify the diff and sweep for consistency**

```bash
git diff --stat "Project Files/Source/Console/setup.cs"
grep -n "cmbRadeRX2Mode\|chkRadeRX2Loopback\|udRadeRX2Level\|lblRadeRX2Status\|grpRadeRX2Core" \
  "Project Files/Source/Console/setup.cs" "Project Files/Source/Console/setup.designer.cs" \
  | wc -l
```

Expected: `setup.cs` diff on the order of `+160/-0` lines. The `grep` count
sweep should show every one of the five new control names appearing in
BOTH files (designer.cs declares/instantiates them, setup.cs's new
handlers and `InitRadePanelFromBackend` reference them) — if any name is
missing from one file, the build will fail with an undefined-identifier
error.

- [ ] **Step 6: Commit**

```bash
git add "Project Files/Source/Console/setup.designer.cs" "Project Files/Source/Console/setup.cs"
git commit -m "feat(radae): add RX2 Core to the Digital Voice tab

New grpRadeRX2Core group (compact 2-row layout, filling the existing
76px gap between RX1 Core and Diagnostics) mirrors RX1's Mode/
Loopback/RX Level/Status controls for RX2, wired to a separate
RadioDSPRX instance (GetDSPRX(1, 0)) that inherits sub-project #2's
interlock automatically -- no radio.cs changes needed. RX2's Mode
selector never touches TX (single shared resource, stays controlled
via RX1's selector). RX2's Loopback checkbox is RADE V1/V2 only --
700E's loopback bridge is structurally RX1-only in pipe.c, confirmed
by reading xpipe() directly, so there's no RX2 700E loopback path to
wire. A dedicated poll timer grays the whole group out when RX2 itself
isn't running (console.RX2Enabled has no change event to hook).

Part of sub-project 3 of 5, see
docs/superpowers/specs/2026-08-25-rx2-digital-voice-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
git push
```

---

### Task 2: CAT layer — five new RX2 commands

**Files:**
- Modify: `Project Files/Source/Console/CAT/CATCommands.cs` (add methods `ZZFC`, `ZZFE`, `ZZFG`, `ZZFN`, `ZZFK`)
- Modify: `Project Files/Source/Console/CAT/CATParser.cs` (add five `case` entries to the extended-command switch)
- Modify: `Project Files/Source/Console/CAT/CATStructs.xml` (add five `<catstruct>` entries)

**Interfaces:**
- Consumes: `console.radio.GetDSPRX(1, 0).RXAFDVRun`/`.RXRadaeEnabled` (same properties Task 1 uses), `WDSP.SetRadaeProtocolV2`/`GetRadaeProtocolV2`, `WDSP.SetRadaeLoopbackEnabled`/`GetRadaeLoopbackEnabled`, `WDSP.GetRadaeSync`/`GetRadaeSnrDb`, `WDSP.GetRXAFDVSync`/`GetRXAFDVSnr`, `WDSP.id`, `WDSP.SetRadaeRxScale`/`GetRadaeRxScale`, `parser.nSet`/`nGet`/`Error1`, `AddLeadingZeros` (all pre-existing, all already used by the RX1 `ZZEX`/`ZZDL`/`ZZDZ`/`ZZDS`/`ZZEO` commands this task mirrors).
- Produces: five new public methods on `CATCommands` — no other task depends on them (this plan's UI task, Task 1, independently calls the same backend functions directly, matching the established CAT/UI duplication pattern from sub-project #2).

Five codes, next free letters in the `ZZFx` family (`ZZEx` is down to its
last two free letters after sub-project #2's 17 commands):

| Code | Setting | Mirrors | Format |
|---|---|---|---|
| `ZZFC` | RX2 mode (0=Off/1=700E/2=RADE V1/3=RADE V2) | `ZZEX`, but never touches TX | `0`/`1`/`2`/`3`, nSet=1 |
| `ZZFE` | RX2 RADE loopback enable | `ZZDL` | `0`/`1`, nSet=1 |
| `ZZFG` | RX2 RADE decode sync/SNR (get-only) | `ZZDZ` | nAns=5 |
| `ZZFN` | RX2 700E decode sync/SNR (get-only) | `ZZDS` | nAns=5 |
| `ZZFK` | RX2 decoder-input level (dB) | `ZZEO` | `<sign><2-digit>`, nSet=3 |

`ZZFG`/`ZZFN` split RX2 status by backend the same way RX1's `ZZDZ`/`ZZDS`
already do — there is no combined mode-aware status command at the CAT
layer for RX1 either, so this isn't a new pattern.

- [ ] **Step 1: Write the CAT-layer splicing script**

Create `scratchpad/splice_cat_rx2.py`:

```python
#!/usr/bin/env python3
"""Add ZZFC/ZZFE/ZZFG/ZZFN/ZZFK (RX2 Digital Voice CAT commands), byte-exact,
preserving each file's CRLF convention."""

ROOT = "Project Files/Source/Console/CAT/"

def crlf(text):
    return text.replace("\r\n", "\n").replace("\n", "\r\n")

# --- 1. CATCommands.cs: five new methods, right after ZZEX ---------------

methods = crlf('''        // Reads or sets RX2's unified Digital Voice mode index (sub-project 3
        // of 5, see docs/superpowers/specs/2026-08-25-rx2-digital-voice-design.md)
        // -- Setup DSP/Digital Voice panel's RX2 Core Mode combo
        // (cmbRadeRX2Mode). 0=Off, 1=700E, 2=RADE V1, 3=RADE V2. Mirrors
        // ZZEX exactly except it never touches TX -- TX encode is a single
        // shared resource controlled only via RX1's Mode selector (ZZEX). // W5TSU
        public string ZZFC(string s)
        {
            if (s.Length == parser.nSet && (s == "0" || s == "1" || s == "2" || s == "3"))
            {
                switch (s)
                {
                    case "0": // Off
                        console.radio.GetDSPRX(1, 0).RXAFDVRun = 0;
                        console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 0;
                        break;
                    case "1": // 700E
                        console.radio.GetDSPRX(1, 0).RXAFDVRun = 1;
                        break;
                    case "2": // RADE V1
                        WDSP.SetRadaeProtocolV2(1, 0);
                        console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 1;
                        break;
                    case "3": // RADE V2
                        WDSP.SetRadaeProtocolV2(1, 1);
                        console.radio.GetDSPRX(1, 0).RXRadaeEnabled = 1;
                        break;
                }
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                if (console.radio.GetDSPRX(1, 0).RXRadaeEnabled != 0)
                    return (WDSP.GetRadaeProtocolV2(1) != 0) ? "3" : "2";
                else if (console.radio.GetDSPRX(1, 0).RXAFDVRun != 0)
                    return "1";
                else
                    return "0";
            }
            else
            {
                return parser.Error1;
            }
        }
        // Reads or sets RX2's RADE loopback bridge (ChannelMaster/radae.c
        // SetRadaeLoopbackEnabled(1, ...)) -- mirrors ZZDL exactly for
        // rx=1. RADE V1/V2 only; 700E has no RX2 loopback path (see
        // cmbRadeRX2Mode_SelectedIndexChanged's comment in setup.cs). // W5TSU
        public string ZZFE(string s)
        {
            if (s.Length == parser.nSet && (s == "0" || s == "1"))
            {
                WDSP.SetRadaeLoopbackEnabled(1, (s == "1") ? 1 : 0);
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                return (WDSP.GetRadaeLoopbackEnabled(1) != 0) ? "1" : "0";
            }
            else
            {
                return parser.Error1;
            }
        }
        // Reads RX2 RADE decode sync/SNR status (get-only): mirrors ZZDZ
        // exactly for rx=1. "<sync 0|1><sign><snr dB, 3 digits>". // W5TSU
        public string ZZFG()
        {
            const int rx = 1;
            bool sync = WDSP.GetRadaeSync(rx) != 0;
            int snr = sync ? WDSP.GetRadaeSnrDb(rx) : 0;

            string sign = snr < 0 ? "-" : "+";
            snr = Math.Min(Math.Abs(snr), 999);

            return (sync ? "1" : "0") + sign + AddLeadingZeros(snr, 3);
        }
        // Reads RX2 700E decode sync/SNR status (get-only): mirrors ZZDS
        // exactly for WDSP.id(1, 0). "<sync 0|1><sign><snr*10, 3 digits>". // W5TSU
        public string ZZFN()
        {
            int ch = WDSP.id(1, 0);
            bool sync = WDSP.GetRXAFDVSync(ch) != 0;
            double snr = sync ? WDSP.GetRXAFDVSnr(ch) : 0.0;

            int snrTenths = (int)Math.Round(snr * 10.0);
            string sign = snrTenths < 0 ? "-" : "+";
            snrTenths = Math.Abs(snrTenths);
            snrTenths = Math.Min(snrTenths, 999);

            return (sync ? "1" : "0") + sign + AddLeadingZeros(snrTenths, 3);
        }
        // Reads or sets RX2's decoder-input level (ChannelMaster/radae.c
        // SetRadaeRxScale/GetRadaeRxScale(1, ...)) -- mirrors ZZEO exactly
        // for rx=1, same dB<->linear conversion (radae.c takes a linear
        // gain factor, not dB -- see ZZEO/ZZEC's own comments for the full
        // story). Signed 2-digit field, clamped to -40..+40. // W5TSU
        public string ZZFK(string s)
        {
            const int rx = 1;
            if (s.Length == parser.nSet)
            {
                int n = Convert.ToInt32(s);
                n = Math.Max(-40, Math.Min(40, n));
                WDSP.SetRadaeRxScale(rx, Math.Pow(10.0, n / 20.0));
                return "";
            }
            else if (s.Length == parser.nGet)
            {
                double db = 20.0 * Math.Log10(Math.Max(0.001, WDSP.GetRadaeRxScale(rx)));
                int n = (int)Math.Round(Math.Max(-40.0, Math.Min(40.0, db)));
                string sign = n < 0 ? "-" : "+";
                return sign + AddLeadingZeros(Math.Abs(n), 2);
            }
            else
            {
                return parser.Error1;
            }
        }
''')

# The tail comment ("/// <summary> ... VAC Stereo checkbox") is genuinely
# LF-only in the real file even though everything before it in this method
# is CRLF -- same mixed boundary sub-project #2's equivalent splice hit,
# confirmed the same way: read the raw bytes before trusting it.
code_before = crlf('''X(0, 0).RXAFDVRun != 0)
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
tail_comment = "        /// <summary>\n        /// Sets or reads the VAC Stereo checkbox\n"

anchor_b = (code_before + tail_comment).encode("utf-8")
replacement_b = (code_before + methods + tail_comment).encode("utf-8")

path = ROOT + "CATCommands.cs"
data = open(path, "rb").read()
count = data.count(anchor_b)
assert count == 1, f"CATCommands.cs anchor found {count} times, expected 1"
data = data.replace(anchor_b, replacement_b)
open(path, "wb").write(data)
print("CATCommands.cs: 5 methods inserted")

# --- 2. CATParser.cs: five new switch cases -------------------------------

cases = crlf('''                case "ZZFC":
                    rtncmd = cmdlist.ZZFC(suffix);
                    break;
                case "ZZFE":
                    rtncmd = cmdlist.ZZFE(suffix);
                    break;
                case "ZZFG":
                    rtncmd = cmdlist.ZZFG();
                    break;
                case "ZZFN":
                    rtncmd = cmdlist.ZZFN();
                    break;
                case "ZZFK":
                    rtncmd = cmdlist.ZZFK(suffix);
                    break;
''')

anchor_parser = crlf('''                case "ZZEX":
                    rtncmd = cmdlist.ZZEX(suffix);
                    break;
                case "ZZTC":
''')
replacement_parser = crlf('''                case "ZZEX":
                    rtncmd = cmdlist.ZZEX(suffix);
                    break;
''') + cases + crlf('''                case "ZZTC":
''')

path = ROOT + "CATParser.cs"
data = open(path, "rb").read()
anchor_b = anchor_parser.encode("utf-8")
count = data.count(anchor_b)
assert count == 1, f"CATParser.cs anchor found {count} times, expected 1"
data = data.replace(anchor_b, replacement_parser.encode("utf-8"))
open(path, "wb").write(data)
print("CATParser.cs: 5 cases inserted")

# --- 3. CATStructs.xml: five new catstruct entries ------------------------

structs = crlf('''  <catstruct code="ZZFC">
    <desc>RX2 unified mode index: 0=Off 1=700E 2=RADE V1 3=RADE V2</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
  <catstruct code="ZZFE">
    <desc>RX2 RADE loopback bridge enable status</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
  <catstruct code="ZZFG">
    <desc>RX2 RADE decode sync/SNR status</desc>
    <active>true</active>
    <nsetparms>-1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>5</nansparms>
  </catstruct>
  <catstruct code="ZZFN">
    <desc>RX2 700E decode sync/SNR status</desc>
    <active>true</active>
    <nsetparms>-1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>5</nansparms>
  </catstruct>
  <catstruct code="ZZFK">
    <desc>RX2 decoder-input level (dB)</desc>
    <active>true</active>
    <nsetparms>3</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>3</nansparms>
  </catstruct>
''')

anchor_xml = crlf('''  <catstruct code="ZZEX">
    <desc>Digital Voice unified mode index: 0=Off 1=700E 2=RADE V1 3=RADE V2</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
  <catstruct code="ZZTC">
''')
replacement_xml = crlf('''  <catstruct code="ZZEX">
    <desc>Digital Voice unified mode index: 0=Off 1=700E 2=RADE V1 3=RADE V2</desc>
    <active>true</active>
    <nsetparms>1</nsetparms>
    <ngetparms>0</ngetparms>
    <nansparms>1</nansparms>
  </catstruct>
''') + structs + crlf('''  <catstruct code="ZZTC">
''')

path = ROOT + "CATStructs.xml"
data = open(path, "rb").read()
anchor_b = anchor_xml.encode("utf-8")
count = data.count(anchor_b)
assert count == 1, f"CATStructs.xml anchor found {count} times, expected 1"
data = data.replace(anchor_b, replacement_xml.encode("utf-8"))
open(path, "wb").write(data)
print("CATStructs.xml: 5 catstruct entries inserted")
```

- [ ] **Step 2: Run it and verify each diff**

```bash
python3 scratchpad/splice_cat_rx2.py
git diff --stat "Project Files/Source/Console/CAT/CATCommands.cs" \
                "Project Files/Source/Console/CAT/CATParser.cs" \
                "Project Files/Source/Console/CAT/CATStructs.xml"
```

Expected: all three "inserted" print lines, and `git diff --stat` shows
roughly `CATCommands.cs +90`, `CATParser.cs +15`, `CATStructs.xml +35` —
not hundreds/thousands of lines. If any file shows a huge diff, `git
checkout -- <that file>` and re-run after fixing the script.

- [ ] **Step 3: Commit**

```bash
git add "Project Files/Source/Console/CAT/CATCommands.cs" \
        "Project Files/Source/Console/CAT/CATParser.cs" \
        "Project Files/Source/Console/CAT/CATStructs.xml"
git commit -m "feat(radae): five RX2 Digital Voice CAT commands

ZZFC (unified mode, mirrors ZZEX minus the TX side), ZZFE (RADE
loopback, mirrors ZZDL), ZZFG/ZZFN (RADE/700E sync-SNR status,
mirror ZZDZ/ZZDS -- same per-backend split RX1's status commands
already use, not a new pattern), ZZFK (decoder-input level, mirrors
ZZEO). All rx=1 against the same already-rx-parameterized backend
calls RX1's commands use -- zero new P/Invoke declarations.

Part of sub-project 3 of 5, see
docs/superpowers/specs/2026-08-25-rx2-digital-voice-design.md.

Co-Authored-By: Claude Sonnet 5 <noreply@anthropic.com>"
git push
```

---

### Task 3: Build, deploy, and verify on hl2winbox

**Files:** none (verification only).

**Interfaces:** none — consumes the finished state of Tasks 1-2.

This is the first time RX2's RADE decode path is ever exercised on real
hardware — treat any surprise here (decode never syncs, audio is silent,
etc.) as a genuine finding to report, not a testing-process failure. Turn
on RX2 in the main console (`console.RX2Enabled` / the "RX2" button) before
starting — none of this is observable with RX2 powered off.

- [ ] **Step 1: Trigger the CI build, download, and deploy**

Same pipeline as sub-projects #1/#2 (see Global Constraints for the exact
`.ps1` script contents if they need recreating):

```bash
gh workflow run build.yml --ref FreeDV -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
# poll: gh run view <run-id> --json status,conclusion -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
gh run download <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 -n Thetis-HL2-installer -D scratchpad/deploy
scp scratchpad/deploy/Thetis-Test-v2.10.3.x64.msi hl2winbox:Downloads/Thetis-Test-v2.10.3.x64.msi
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\extract.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\stop_thetis.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\copy_build.ps1"
ssh hl2winbox "powershell -ExecutionPolicy Bypass -File Downloads\\run_interactive.ps1 -ScriptName relaunch.ps1"
./Tools/thetis-ai-control/thetisctl cat --host 100.117.67.160 --timeout 8s version
```

Expected: CI `conclusion: success`; the version string's `git:` short SHA
matches Task 2's commit.

- [ ] **Step 2: Turn RX2 on**

RX2's own power state lives in the main console, not Setup — confirm it's
on before testing (either visually via a screenshot, per Step 4 below, or
by asking whoever is at the machine to click the "RX2" button first).

- [ ] **Step 3: CAT round-trip — mode, loopback, status, level**

Reuse `scratchpad/cat_roundtrip.py` from sub-project #2 (a raw CAT-over-TCP
socket script — `thetisctl`'s `query` subcommand is get-only and can't
send arbitrary sets):

```bash
python3 scratchpad/cat_roundtrip.py \
  "GET:ZZFC" \
  "SET:ZZFC1" "GET:ZZFC" "GET:ZZFN" \
  "SET:ZZFC2" "GET:ZZFC" "GET:ZZFG" \
  "SET:ZZFC3" "GET:ZZFC" "GET:ZZFG" \
  "SET:ZZFK-15" "GET:ZZFK" \
  "SET:ZZFE1" "GET:ZZFE" \
  "SET:ZZFC1" "GET:ZZFE" \
  "SET:ZZFC0" "GET:ZZFC" "GET:ZZFE"
```

Expected, in order:
1. `GET:ZZFC -> 0` (assuming a clean starting state — if not, the logic
   under test is what matters, not the exact starting digit)
2. After `SET:ZZFC1` (700E): `GET:ZZFC -> 1`; `GET:ZZFN` returns a
   `<sync><sign><3-digit>` string starting with `0` (not synced yet,
   nothing decoding) — confirms the 700E status path responds for RX2
3. After `SET:ZZFC2` (RADE V1): `GET:ZZFC -> 2`; `GET:ZZFG` returns a
   similar `<sync><sign><3-digit>` string
4. After `SET:ZZFC3` (RADE V2): `GET:ZZFC -> 3`
5. After `SET:ZZFK-15`: `GET:ZZFK -> -15`
6. After `SET:ZZFE1` (loopback on, mode is still 3/RADE V2 from step 4, so
   loopback is available): `GET:ZZFE -> 1`
7. After `SET:ZZFC1` (switch to 700E while loopback was on): `GET:ZZFE`
   should read back `0` — confirms `cmbRadeRX2Mode_SelectedIndexChanged`'s
   disarm-on-mode-change logic actually fires via the CAT path too (`ZZFC`
   and the UI combo both call the same interlocked properties, so this
   should hold even though `ZZFE` itself has no mode-awareness of its own)
8. After `SET:ZZFC0` (Off): `GET:ZZFC -> 0`, `GET:ZZFE -> 0`

If step 7 doesn't show `ZZFE` reading back to `0`, that's a real finding —
it would mean `ZZFC`'s mode-switch logic and `ZZFE`'s raw loopback flag can
diverge via CAT the same way sub-project #2's final review found for RX1's
UI checkbox, just reached from a different entry point. Stop and
investigate rather than proceeding to Step 4.

- [ ] **Step 4: First real hardware test — does RX2 RADE decode actually work?**

With RX2 on and `ZZFC2` (RADE V1) or `ZZFC3` (RADE V2) set, and `ZZFE1`
(loopback) set, talk into the mic (same zero-RF loopback test #1's spec
describes for RX1 — TX encoder's modem output bridges directly into RX2's
decoder input, no RF, radio never keys) and confirm audio comes back
through RX2. Poll `GET:ZZFG` a few times over a couple of seconds and
confirm `sync` flips to `1` with a plausible SNR. This is the first time
this exact path has ever been exercised — if it doesn't sync or produces
no audio, that's a genuine backend finding (something in `xradae_rx`'s
dual-RX handling doesn't work as the code comments claim), not a UI/CAT
bug — report it plainly rather than debugging blind; it may need its own
follow-up investigation outside this plan's scope.

- [ ] **Step 5: Screenshot verification — layout and gray-out**

Reuse the screenshot/click scripts from sub-project #2 (`screenshot.ps1`,
a click helper that minimizes stray windows and clicks Setup's known
coordinates — see #2's plan Task 4 for the full script contents if they
need recreating). Navigate to Setup → DSP → Digital Voice and confirm:
- "RX2 Core" group is visible in the gap between RX1 Core and Diagnostics,
  with both rows (Mode+Loopback, RX Level+Status) readable and not
  overlapping anything.
- With RX2 on: the group is enabled/interactive.
- With RX2 off (toggle it off in the main console, wait ~1s for the poll
  timer, re-screenshot): the group visibly grays out.
- With RX2 Core's Mode at "700E": confirm the Loopback checkbox is grayed
  (per this plan's design decision — 700E has no RX2 loopback path).

- [ ] **Step 6: Leave the box in a known state**

```bash
python3 scratchpad/cat_roundtrip.py "SET:ZZFC0"
```

- [ ] **Step 7: Report results**

No commit for this task (verification-only) — report the CAT round-trip
results, the Step 4 hardware finding (decode worked / didn't, with
specifics either way), and the screenshot observations back to whoever
requested the plan.
