---
name: thetis-fork-merge
description: Consolidate this Hermes-Lite 2 fork against all four upstream/sibling Thetis repositories — ramdor/Thetis (official), mi0bot/OpenHPSDR-Thetis, ON7OFF/Thetis, and sv1eia/Thetis-RADE — into a chosen target branch. Generalizes thetis-merge (mi0bot+ramdor into master only) to all four sources and a configurable target. Use when asked to "merge upstream forks", "consolidate the Thetis repos", "sync trial-merge-upstream", or repeat this multi-fork merge process.
---

# Thetis multi-fork merge process

This fork doesn't sit downstream of one upstream — it sits at the convergence point of four
independent Thetis lineages, each related to this repo (and to each other) differently:

```
ramdor/Thetis (official, 2020 divergence point)
   │
   ├──ported by hand, not git-merged──► mi0bot/OpenHPSDR-Thetis ──git merge (shared history)──► this repo
   │
   ├──diverged 2020, same base, sibling fork──► ON7OFF/Thetis (HL2 port, inactive since 2025-08)
   │
   └──squash-imported 2026-05, no shared git history──► sv1eia/Thetis-RADE (RADE V1+V2 fork, active)
```

Only the mi0bot edge is a normal `git merge`. Everything else requires different tooling —
that's the whole reason this skill exists, generalizing `thetis-merge` (which only ever handled
mi0bot + cherry-picks from ramdor, targeting `master`) to all four sources and a target branch
you choose. `thetis-merge` still exists and is still correct for its narrower scope; this skill
supersedes it in capability but doesn't replace it as a file — ask before deleting the old one.

## Facts verified 2026-08-24 (recheck if stale — these are exactly the kind of claims that go
stale silently; the sv1eia entry below is itself a correction of an 6-week-old claim)

| Repo | Remote | Relationship to this repo | Status as of this check |
|------|--------|---------------------------|--------------------------|
| `mi0bot/OpenHPSDR-Thetis` | `upstream` | Direct git upstream — shared history, normal `git merge` works | 0 new commits beyond what we have (tip `0cef1c90`, 2026-06-12) — fully caught up |
| `ramdor/Thetis` (official) | `ramdor` | Shares **only** the 2020-10-29 merge-base `ed4c27c9` ("files cleanup") — mi0bot ported this by hand, not by merge. **Never `git merge` directly.** | Latest `852bf0ef` (2026-07-02). 1787 commits since the base not in ours — expected/by design (we only want portable fixes, not full parity with 6 years of independent development), not a backlog to clear |
| `ON7OFF/Thetis` | `on7off` | Also shares **only** the same 2020-10-29 base — a **sibling** fork off the old ramdor point, not downstream of mi0bot despite being an HL2 port like this one | Latest `81e6e093` (2025-08-28) — inactive ~1 year as of this check, matching the prior (now-superseded) `thetis-merge` skill's July-2026 claim. Two branches, 1 commit apart (`on7off-dev`, `feat/hl2-console-pass1`) — same lineage, treat as one |
| `sv1eia/Thetis-RADE` | `sv1eia` | **No shared git history at all.** Their repo is a single squash-imported root commit, `06c79f4f "Initial commit: Thetis-RADE fork from ramdor/Thetis v2.10.3.15-g2e"` (2026-05-06), 36 commits total. `git merge-base HEAD sv1eia/main` returns nothing | Latest `408f2b52` (2026-06-27) — **actively developed**, unlike the other two. Recently added RADE **V2** support and their own independent HL2 port |

**The sv1eia correction, spelled out**: `Documentation/FreeDV-Plan.md`'s 2026-08-10 sv1eia-eval
entry claims "both this fork and sv1eia's share real git history via ramdor/Thetis (common
ancestor `ed4c27c9`)... this is cherry-pick territory." A fresh check today finds no such
ancestor — `git merge-base` is empty. Either sv1eia's history was rewritten/re-imported since
that entry was written, or the original check was mistaken (possibly it found a merge-base
between different refs than intended). Either way, **treat sv1eia as vendor-drop-only** (§5
below) until this is re-verified, not as cherry-pick territory — and flag the discrepancy to
whoever owns `FreeDV-Plan.md` so that entry gets a correction note, the same way this doc's own
convention handles superseded claims (add a note, don't silently rewrite).

**A trap already found once**: `git diff --stat HEAD on7off/on7off-dev` reports **3300+ files,
millions of lines changed** — this is a tree-shape artifact (unrelated-looking paths read as
wholesale delete+add without rename detection), not a real measure of how different the code
is. Don't use raw diffstat against ON7OFF as a signal of anything; use the symbol-grep method in
§4 instead.

## Step 1 — Target branch and working branch

Unlike `thetis-merge` (always `master`), **the target branch is a parameter you choose**. State
it explicitly before starting — don't assume.

```bash
git checkout <target-branch> && git pull origin <target-branch>   # skip pull if it has no upstream tracking (e.g. a local-only trial branch)
git checkout -b fork_merge_$(date +%Y%m%d)
```

## Step 2 — Fetch all four

```bash
git fetch upstream   # mi0bot
git fetch ramdor
git fetch on7off
git fetch sv1eia
```

If any remote is missing (`git remote -v`), add it first:

```bash
git remote add upstream https://github.com/mi0bot/OpenHPSDR-Thetis.git
git remote add ramdor   https://github.com/ramdor/Thetis.git
git remote add on7off   https://github.com/ON7OFF/Thetis.git
git remote add sv1eia   https://github.com/sv1eia/Thetis-RADE.git
```

## Step 3 — mi0bot (direct merge, only if it has new commits)

```bash
git rev-list --count HEAD..upstream/master
```

If that's `0`, skip — nothing to merge, move to Step 4 (the value is usually there anyway,
same as `thetis-merge` found). If nonzero:

```bash
git merge upstream/master
```

Resolve conflicts per Step 6.

## Step 4 — ramdor + ON7OFF (cherry-pick territory — both connect only via the 2020 base)

Identical process for both, since both only share `ed4c27c9` with us:

1. Determine `<last-sync-date>` per repo — the newest "Updated from official codebase"/cherry-pick
   recorded in `ReleaseNotes.txt`, or this skill's own last-checked date above if never synced.
2. List candidates:
   ```bash
   git log --format='%h %ad %s' --date=short --since=<last-sync-date> ramdor/master
   git log --format='%h %ad %s' --date=short --since=<last-sync-date> on7off/on7off-dev
   ```
3. **Skip docs/status-only commits** (README, ReleaseNotes-only, project-status changes).
   ON7OFF commit messages are sometimes Dutch — translate before judging relevance, don't skip
   on language alone.
4. **Verify the change is actually missing** — pick a distinctive symbol, grep our tree. Don't
   use raw `git diff --stat` against ON7OFF (see the trap noted above).
   ```bash
   git grep "<distinctive symbol>" HEAD -- "Project Files/Source"
   ```
5. **Dry-run before touching the tree**:
   ```bash
   git merge-tree --write-tree --merge-base=<sha>^ HEAD <sha>
   ```
6. **Apply with provenance**:
   ```bash
   git cherry-pick -x <sha>
   ```

**ON7OFF as a conflict-resolution reference, not just a fix source**: their own history shows
they've already 3-way-merged mi0bot onto ramdor themselves (e.g. `3053df9b "3-way merge mi0bot
2.10.3.8 onto ramdor 2.10.3.11 (Console, clean pass)"`). If we hit the same conflict they did,
their resolution is worth checking as precedent — but don't cherry-pick their merge commits
directly, merge commits don't cherry-pick cleanly; look at what the merged result contains and
port the resolution by hand.

## Step 5 — sv1eia (vendor-drop only — no shared git history)

No common ancestor means no 3-way merge/cherry-pick tooling applies. Two cases:

- **New files unique to sv1eia** (their RADE additions — `ChannelMaster/radae*.{c,h}`,
  `Project Files/lib/{radae_c,opus_dnn,r8brain,freedv_text}`, etc.): copy directly.
  ```bash
  git checkout sv1eia/main -- <path>
  ```
  **Check for a path collision with this fork's own RADE work first**:
  ```bash
  git log --all --oneline -- <path>
  ```
  This repo's `FreeDV` branch already has its own, independently-built RADE V1 RX/TX
  implementation (`ChannelMaster/radae.c`, `radae_micdsp.c`, etc. — see
  `Documentation/FreeDV-Plan.md`). Don't blindly overwrite it with sv1eia's version — compare
  the two approaches first and bring the change to whoever owns that branch's RADE work before
  deciding which wins.
- **Changes to files this repo already has** (most of `Console/`): read
  `git diff sv1eia/main -- <path>` as a patch — the same care as an emailed diff from someone
  with no shared history — and manually decide what to port. There is no dry-run merge-tree
  option here; the diff itself is the only tool.

sv1eia is the one actively-developed source of the four (latest commit 2026-06-27 as of this
check, added RADE V2 support and their own HL2 port since the file list above was last
characterized) — re-check their commit list each run rather than trusting a cached file list.

## Step 6 — Fork-specific conflict resolution rules

Properties of *this* fork's files, independent of which source triggered the conflict:

| File | Rule |
|------|------|
| `Console/titlebar.cs` | **Keep ours** — `BUILD_NAME` is this fork's identity (e.g. "HL2 (MI0BOT/W5TSU)"; must stay ≤ 22 chars so the splash's 32-char version line fits). |
| `Source/ReleaseNotes.txt` | **Keep ours**, then add a line for the picked change under the current HL2 section, citing the upstream repo + short SHA ("- [add] ... - from official Thetis <sha>", or "- [add] ... - from ON7OFF/Thetis <sha>"). |
| `Console/packages.config` + `Thetis.csproj` + `app.config` | Treat as one unit: take NuGet version bumps in **all three or none** — a mismatch breaks the build. Upstream's bumps are safe (CI restores from nuget.org). |
| `ReadMe.md` | **Keep the user's prose** (fork purpose, developer acknowledgments); take only newer release-headline/changelog facts. |
| `Console/AssemblyInfo.cs` | Keep ours; version changes happen only in the release step (and only when the target branch is `master` — see Step 10). |

After resolving, always check: `git grep -l "^<<<<<<<" -- "Project Files"` must be empty.

## Step 7 — Semantic verification (no local build possible — Linux box, Windows-only project)

- If a source **removed** code: grep for dangling references to every removed method/symbol.
- If a source **added UI**: confirm designer additions are self-contained (explicit properties,
  no `ApplyResources` needing missing `.resx` entries) and event handlers exist in the code-behind.
- New P/Invoke externs: confirm the matching `PORT` function exists in the C source.
- Anything touching `ChannelMaster/radae*`/`wdsp/fdv.c` specifically: cross-check against
  `Documentation/FreeDV-Plan.md` before accepting a foreign version of these files — this fork's
  own FreeDV/RADE work is independent of all four sources above and the most likely place a
  naive merge silently regresses something already fixed here.

## Step 8 — Build via CI

```bash
git push -u origin <branch>
gh workflow run build.yml --ref <branch> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
gh run list --workflow=build.yml --branch <branch> --limit 1 -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
gh run watch <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 --exit-status
```

(`gh` may default to the mi0bot repo — always pass `-R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`.)

## Step 9 — Merge to the target branch (after green CI and user approval)

```bash
git checkout <target-branch> && git merge --ff-only origin/<target-branch>
git merge <working-branch> --no-edit
git push origin <target-branch>
```

`ReadMe.md` is the usual conflict here if the target branch has moved meanwhile — apply the
Step 6 rule.

## Step 10 — Release (only when the target branch is `master` and the user asks)

Doesn't apply to a trial/consolidation branch. When it does apply, it's identical to
`thetis-merge`'s Step 8: bump `Console/AssemblyInfo.cs`, add a `ReleaseNotes.txt` section, update
`ReadMe.md`'s headline, commit as `Bump version to 2.10.3.NN HL2`, tag `v2.10.3.NN`, push — the
tag run's `release` job publishes the MSI.

## Reporting

Per source repo: how many new commits found, what was taken/skipped and why, how conflicts were
resolved (and whether ON7OFF's own prior resolution was used as precedent), CI result with run
link. For sv1eia specifically: which files were vendor-dropped vs. hand-ported, and any
collisions found with this fork's own RADE work. If a source's relationship-to-us claim (this
skill's Facts table) no longer matches what you found, say so explicitly and update the table —
don't let it go stale silently the way the sv1eia entry did. If nothing is mergeable anywhere,
that's a valid and common outcome — say so plainly.
