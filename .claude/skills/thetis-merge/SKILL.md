---
name: thetis-merge
description: Sync this Hermes-Lite 2 fork with its upstream Thetis repositories — merge new mi0bot commits, cherry-pick portable fixes from official ramdor/Thetis, resolve the fork's known conflict patterns, verify via CI, and optionally cut a release. Use when asked to "merge upstream", "sync with mi0bot/ramdor", "check for upstream changes", or repeat the Thetis merge process.
---

# Thetis upstream merge process

This fork (W5TSU/OpenHPSDR-Thetis-Hermes-Lite2) sits at the bottom of a chain:

```
ramdor/Thetis (official)  ──ports──►  mi0bot/OpenHPSDR-Thetis (HL2)  ──merges──►  this repo
```

Facts that shape the process (verified July 2026 — recheck if stale):

- **mi0bot** is the true git upstream (`upstream` remote). Its history is shared with ours;
  normal `git merge` works.
- **ramdor** (official Thetis) shares only a 2020-era merge base (`ed4c27c9`) — mi0bot ported
  its code by hand, not by merging. **Never `git merge` ramdor directly** (thousands of
  conflicts); take individual commits by `git cherry-pick -x` instead. Also: mi0bot's port
  commits ("<sha> - Updated from official codebase") cite ramdor *development-branch* SHAs that
  no longer resolve — compare by **content and date**, not by SHA ancestry.
- **ON7OFF/Thetis** was checked July 2026: inactive since Aug 2025, everything relevant already
  present here. Re-check only if the user asks.
- CI (`.github/workflows/build.yml`) auto-builds on push to `master` and on `v*` tags (tag =
  release job publishes the MSI). Other branches need `workflow_dispatch`.

## Step 1 — Create a working branch and fetch

```bash
git checkout master && git pull origin master
git checkout -b upstream_merge_$(date +%Y%m%d)
git fetch upstream                       # mi0bot
git rev-list --count HEAD..upstream/master
```

## Step 2 — Merge mi0bot (if it has new commits)

```bash
git merge upstream/master
```

Resolve conflicts with the fork rules in Step 4. If zero new commits, continue — the value is
usually in ramdor.

## Step 3 — Check official ramdor/Thetis for portable fixes

```bash
git fetch https://github.com/ramdor/Thetis.git master     # lands in FETCH_HEAD
git log --format='%h %ad %s' --date=short --since=<last-sync-date> FETCH_HEAD | head -30
```

Determine `<last-sync-date>` from the newest "Updated from official codebase" commit in
`git log upstream/master` or the last cherry-pick recorded in `ReleaseNotes.txt`.

For each candidate ramdor commit:

1. **Skip docs/status-only commits** (README, ReleaseNotes-only, project-status changes).
2. **Verify the change is actually missing** — pick a distinctive symbol from the commit and
   grep our tree (`git grep "<symbol>" HEAD -- "Project Files/Source"`). mi0bot often ported
   content already. Beware counting bugs: read the grep output, don't just sum fields.
3. **Dry-run the pick** before touching the tree:
   ```bash
   git merge-tree --write-tree --merge-base=<sha>^ HEAD <sha>   # lists CONFLICT lines
   ```
4. **Apply with provenance** (`-x` appends the upstream SHA, matching fork convention):
   ```bash
   git cherry-pick -x <sha>
   ```

## Step 4 — Fork-specific conflict resolution rules

These files conflict predictably; resolve them the same way every time:

| File | Rule |
|------|------|
| `Console/titlebar.cs` | **Keep ours** — `BUILD_NAME` is this fork's identity (e.g. "HL2 Beta 1 (MI0BOT)"). |
| `Source/ReleaseNotes.txt` | **Keep ours**, then add a line for the picked change under the current HL2 section, citing the upstream short SHA ("- [add] ... - from official Thetis <sha>"). |
| `Console/packages.config` + `Thetis.csproj` + `app.config` | Treat as one unit: take NuGet version bumps in **all three or none** — a mismatch breaks the build. Upstream's bumps are safe (CI restores from nuget.org). |
| `ReadMe.md` | **Keep the user's prose** (fork purpose, developer acknowledgments); take only newer release-headline/changelog facts. |
| `Console/AssemblyInfo.cs` | Keep ours; version changes happen only in the release step. |

After resolving, always check: `git grep -l "^<<<<<<<" -- "Project Files"` must be empty.

## Step 5 — Semantic verification (no local build possible — Linux box, Windows-only project)

- If upstream **removed** code: grep for dangling references to every removed method/symbol.
- If upstream **added UI**: confirm designer additions are self-contained (explicit properties,
  no `ApplyResources` needing missing `.resx` entries) and event handlers exist in the code-behind.
- New P/Invoke externs: confirm the matching `PORT` function exists in the C source.

## Step 6 — Build via CI

```bash
git push -u origin <branch>
gh workflow run build.yml --ref <branch> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
gh run list --workflow=build.yml --branch <branch> --limit 1 -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2
gh run watch <run-id> -R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2 --exit-status
```

(`gh` may default to the mi0bot repo — always pass `-R W5TSU/OpenHPSDR-Thetis-Hermes-Lite2`.)
Green build = NuGet restore + solution + installer all pass. Artifacts appear on the run.

## Step 7 — Merge to master (after green CI and user approval)

```bash
git checkout master && git merge --ff-only origin/master
git merge <branch> --no-edit        # true merge if master gained commits meanwhile
git push origin master              # auto-triggers CI
```

`ReadMe.md` is the usual conflict here (user edits it on GitHub) — apply the Step 4 rule.

## Step 8 — Release (only when the user asks)

The MSI version binds to Thetis.exe's file version (`Product.wxs` →
`bind.FileVersion.ThetisEXE` → `AssemblyInfo.cs`), so **bump before tagging**:

1. `Console/AssemblyInfo.cs`: `AssemblyVersion("2.10.3.NN")`.
2. `Source/ReleaseNotes.txt`: new `# 2.10.3.NN HL2 (date)` section listing the changes.
3. `ReadMe.md`: update the "Latest Release" headline and add a short changelog block.
4. Commit as `Bump version to 2.10.3.NN HL2`, push master, then:
   ```bash
   git tag v2.10.3.NN && git push origin v2.10.3.NN
   ```
5. The tag run's `release` job publishes the MSI at
   `https://github.com/W5TSU/OpenHPSDR-Thetis-Hermes-Lite2/releases/tag/v2.10.3.NN`.
   (The MSI filename shows only three version components — expected.)

## Reporting

Tell the user, for each upstream: how many new commits, what was taken/skipped and why, how
conflicts were resolved, CI result with run link, and (if released) the release URL. If nothing
is mergeable anywhere, say so plainly — that is a valid and common outcome.
