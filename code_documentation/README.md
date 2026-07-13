# Code Documentation

This directory documents the source code of **Thetis for the Hermes-Lite 2** — this repository's
fork of the [mi0bot/OpenHPSDR-Thetis](https://github.com/mi0bot/OpenHPSDR-Thetis) SDR console,
adapted for the Hermes-Lite 2 radio. It is reference documentation for developers: what each part
of the program does, which source files implement it, and how the files relate to each other.

## What is here

| Item | Contents |
|------|----------|
| [`CODE_OUTLINE.md`](CODE_OUTLINE.md) | The starting point. Maps the program's **18 functional areas** (main console, HPSDR network protocol, HL2 I/O board, spectrum display, the wdsp DSP engine, ChannelMaster audio routing, CAT control, CW keying, MIDI, and more) to the source files that implement each, with a 1–2 sentence role per file. |
| [`files/`](files/README.md) | **One page per source file** (281 pages, indexed in `files/README.md`). Each page gives the file's role, a graph-derived summary of how it is used — which files call into it, which files it calls, and its most externally-referenced symbols — and an outline of its classes, methods, and functions. Every function entry carries its source line number, signature, a short description, and how it is called (its callers from the graph, P/Invoke linkage into the C DLLs, or event wiring). |
| [`tools/gen_file_docs.py`](tools/gen_file_docs.py) | The generator that produces everything under `files/` from the knowledge graph and `CODE_OUTLINE.md`. |
| [`NR3.md`](NR3.md) | Deep dive: the RNNoise neural-network noise reduction (NR3) — architecture across all layers, how to use it, how models are loaded/validated, and how to train new models with the bundled RNNoise toolchain. |
| [`History.md`](History.md) | The lineage of this software, 2003–present: FlexRadio's SDR-1000/PowerSDR, the HPSDR group and TAPR, the OpenHPSDR hardware era, PowerSDR mRX, Thetis, and the fork map down to this Hermes-Lite 2 repository. |
| [`Database.md`](Database.md) | Deep dive: the settings database — what "Export Database" (Database Manager) writes, the `database.xml` DataSet-XML format, every table's schema (Options, State, TXProfile, Memory, BandStack2…), and how import/restore works. |
| [`Protocols.md`](Protocols.md) | Deep dive: openHPSDR network Protocols 1 & 2 — framing and port maps, how discovery picks the protocol, which boards run which protocol, the HL2's P1 extensions (I²C tunnelling, ACK), and how HL2 gateware is updated. |
| [`hermes-lite2+.md`](hermes-lite2+.md) | Deep dive: hardware (H/W) selection — where the radio-model data lives (`HPSDRModel`/`HPSDRHW`, `clsHardwareSpecific.cs`), how a selection is applied and persisted, the checklist for adding a new radio, and exactly how the Hermes-Lite 2 is defined. |

Documentation covers the six first-party subprojects under `Project Files/Source/` — Console,
wdsp, ChannelMaster, cmASIO, Midi2Cat, and RawInput (~473 source files). Vendored third-party
libraries (`Project Files/lib/`, the bundled ASIO SDK) and generated `*.Designer.cs` files are
intentionally excluded.

## How it was made

The documentation was produced in July 2026 by **Claude Code** (Anthropic's Claude, Fable 5
model), directed by Mark Grennan (W5TSU), using
**[graphify](https://graphify.net)** to build a knowledge graph of the source tree:

1. graphify parsed every first-party source file with tree-sitter (C, C#, and C++ grammars) —
   deterministic AST extraction, no LLM involved — producing a graph of **15,492 nodes and
   36,536 edges** (classes, methods, functions, and their `calls` / `references` / `contains` /
   `imports` relationships), clustered into 434 communities.
2. The functional areas in `CODE_OUTLINE.md` were derived from those graph communities plus the
   graph's "god node" and cross-community analysis, then written up by Claude with spot-checks
   against the actual source.
3. `tools/gen_file_docs.py` generated the per-file pages mechanically from the graph: symbol
   outlines and line numbers come straight from the AST extraction; each page's role text comes
   from its `CODE_OUTLINE.md` table row.

The graph itself lives in `graphify-out/` at the repository root. It is **not committed**
(gitignored — the 19 MB `graph.json` is fully rebuildable), but once built it can be queried
directly:

```bash
graphify query "How does the console send TX frequency to the HL2 I/O board?"
graphify explain "IoBoardHl2"
graphify path "CATCommands" "NetworkIO"
```

## Regenerating after code changes

```bash
graphify update "Project Files/Source"        # incremental graph rebuild (no LLM needed)
python code_documentation/tools/gen_file_docs.py            # regenerate code_documentation/files/ from the graph
```

`CODE_OUTLINE.md` is hand-written prose and only needs editing when files are added, removed, or
change purpose.

## Caveats

- **Line numbers drift.** They were captured from the source at generation time; regenerate after
  significant changes rather than trusting stale numbers.
- **Extraction is structural, not semantic.** A few symbols tree-sitter could not place
  structurally appear at file level instead of inside their class, and cross-file counts miss
  calls made through P/Invoke, delegates, events, or reflection.
- **Function descriptions are best-effort.** Where a doc comment exists in the source it is used
  verbatim; otherwise the description is inferred from the codebase's naming conventions
  (`create_`/`destroy_`/`x…` block functions in wdsp/ChannelMaster, `Set`/`Get` accessors,
  WinForms `_Click`-style event handlers). The code itself remains the authority on behavior.
