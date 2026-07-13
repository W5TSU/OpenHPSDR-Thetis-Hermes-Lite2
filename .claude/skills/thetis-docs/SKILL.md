---
name: thetis-docs
description: Rebuild or update this repo's code documentation under code_documentation/ — refresh the graphify knowledge graph, regenerate the 281 per-file pages, keep CODE_OUTLINE.md and the doc index current. Use when asked to "update the docs", "regenerate documentation", "document new code", or after merges/refactors that change source files.
---

# Thetis code documentation process

The documentation lives in **`code_documentation/`** (renamed from `docs/` — never recreate a
`docs/` directory) and is generated from a [graphify](https://graphify.net) knowledge graph of
`Project Files/Source`. The graph itself lives in **`graphify-out/`** at the repo root and is
**gitignored** (rebuildable; graph.json is ~19 MB). Layout:

| Item | Nature |
|------|--------|
| `code_documentation/CODE_OUTLINE.md` | Hand-written. 18 functional areas → file tables. **The per-file role text in its tables is the source of truth** consumed by the generator. |
| `code_documentation/files/**.md` | Generated. One page per source file: role, graph-derived usage, symbol outline with signatures/descriptions/callers. Never edit by hand. |
| `code_documentation/tools/gen_file_docs.py` | The generator. Reads `graphify-out/graph.json` + CODE_OUTLINE.md tables. |
| `code_documentation/README.md` | Hand-written index — add a table row for any new document. |
| `NR3.md`, `History.md` | Hand-written deep dives. |

## Routine update (code changed, graph exists)

```bash
graphify update "Project Files/Source"                      # incremental AST re-extract, no LLM
$(cat graphify-out/.graphify_python) code_documentation/tools/gen_file_docs.py
```

Notes:
- The graphify interpreter path is cached in `graphify-out/.graphify_python`; if missing,
  resolve from the shebang of `$(which graphify)`.
- If files were **deleted**, graphify's shrink-guard refuses a smaller graph — rerun with
  `--force` / `GRAPHIFY_FORCE=1` as its error message says.
- If files were **added or repurposed**, first add/update their row in the right
  CODE_OUTLINE.md section table (backticked path + 1–2 sentence role) — otherwise the new file
  gets no page. Keep `*.Designer.cs`, vendored libs (`Project Files/lib/`, ASIO SDK), and
  `Source/packages/` out of the outline.

## Full rebuild (no graphify-out/, or graph corrupt)

Follow graphify's skill runbook (`graphify` CLI is installed; code-only corpus → AST extraction
only, **no API key, never wait for one**). Sequence: detect → AST extract → merge with empty
semantic file → build/cluster → health check → label communities → `graphify export html`.
Expectations from the first build (July 2026): ~479 code files, ~2.2M words (the >2M warning is
expected — proceed), ~15,500 nodes / 36,500 edges / ~430 communities, ~3% dangling edges
(normal), HTML auto-aggregates above 5,000 nodes. Community labels: name from each community's
dominant source file (hand map for major subsystems, derived names + numeric suffixes for the
rest — see prior labels in `graphify-out/.graphify_labels.json` if present).

## Generator behavior worth knowing (tools/gen_file_docs.py)

- Function descriptions: source doc-comment if present, else naming-convention heuristics
  (wdsp `create_`/`destroy_`/`flush_`/`x…` blocks, `Set`/`Get` accessors, WinForms `_Click`
  event handlers). No invented descriptions otherwise.
- "Called by" lines come from graph call edges; C functions with no static caller are checked
  against C# extern names to report **P/Invoke linkage** (the graph doesn't cross the
  language boundary itself).
- Namespace nodes are filtered from cross-file edges (they otherwise create bogus "imports"
  entries).

## Verification before committing

1. Link check — every relative link in generated pages must resolve:
   ```bash
   python3 -c "
   import re
   from pathlib import Path
   bad=0
   for p in Path('code_documentation').rglob('*.md'):
       for m in re.finditer(r'\]\(([^)#]+)(?:#[^)]*)?\)', p.read_text()):
           t=m.group(1)
           if t.startswith('http'): continue
           if not (p.parent/t).resolve().exists(): print('BROKEN:',p,'->',t); bad+=1
   print(bad,'broken')"
   ```
2. Spot-check one C page (e.g. `files/wdsp/RXA.c.md`) and one C# page
   (e.g. `files/Console/HPSDR/IoBoardHl2.cs.md`) for sane outlines and callers.
3. Every file named in CODE_OUTLINE.md must exist on disk (check by basename against
   `Project Files/Source`).
4. No stale `docs/` path references (`grep -rn "docs/" code_documentation | grep -v code_documentation`).

## Deep dives and the index

Hand-written documents (like `NR3.md`, `History.md`) follow this pattern: trace the feature
through every layer (console UI → radio.cs → dsp.cs P/Invoke → wdsp → native lib), cite files
with line numbers, cross-link to the generated per-file pages, and end with related-doc links.
After adding one: add a row to `code_documentation/README.md`'s table, and link from the
top-level `ReadMe.md` only if the user asks.

## Committing

Commit `code_documentation/` changes with a message describing what was regenerated and why
(e.g. "Regenerate code docs after <merge/feature>"). `graphify-out/` stays untracked. Push only
when the user's established workflow expects it.

## Exploring the graph (for answering questions, not just docs)

```bash
graphify query "How does X work?"        # BFS over the graph
graphify explain "IoBoardHl2"            # one node + neighbors
graphify path "CATCommands" "NetworkIO"  # shortest path
```
