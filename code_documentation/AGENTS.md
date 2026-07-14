# code_documentation

## Purpose

Developer reference for the source tree: a hand-written outline and deep-dive documents plus
281 generated per-file pages built from a graphify knowledge graph.

## Ownership

- Hand-written (edit directly): `CODE_OUTLINE.md` (its file-role table rows are the
  generator's input), `README.md` (the index), `History.md`, `NR3.md`, `Protocols.md`,
  `Database.md`, `hermes-lite2+.md`
- Generated (never hand-edit; regenerate instead): everything under `files/`
- `tools/gen_file_docs.py` is the generator; the graph lives in `graphify-out/` at the repo
  root (gitignored, rebuildable)

## Local Contracts

- Full runbook: `.claude/skills/thetis-docs/SKILL.md`. Routine update:
  `graphify update "Project Files/Source"` then
  `$(cat graphify-out/.graphify_python) code_documentation/tools/gen_file_docs.py`
- New or repurposed source files need a `CODE_OUTLINE.md` table row (backticked path +
  1–2 sentence role) before regeneration, or they get no page
- Every new document added here gets a row in `README.md`'s table
- This directory is `code_documentation/` — never recreate a `docs/` directory

## Verification

- Relative-link check across all `*.md` (script in the thetis-docs skill) must report 0 broken
- Spot-check one C page (`files/wdsp/RXA.c.md`) and one C# page
  (`files/Console/HPSDR/IoBoardHl2.cs.md`) for sane outlines and callers

## Child DOX Index

(none)
