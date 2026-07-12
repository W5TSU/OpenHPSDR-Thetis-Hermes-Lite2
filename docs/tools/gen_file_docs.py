"""Generate per-source-file documentation from CODE_OUTLINE.md + graphify graph.json.

For every source file referenced in docs/CODE_OUTLINE.md, writes
docs/files/<relpath>.md containing:
  - the file's role (from the outline table row) and its functional area
  - how it is used: cross-file callers/callees aggregated from the graph,
    plus the file's most externally-called symbols
  - an outline of classes/methods/functions with line numbers
Also writes docs/files/README.md as an index grouped by functional area.
"""
import json
import os
import re
from collections import Counter, defaultdict
from pathlib import Path

ROOT = Path('/home/w5tsu/Development/OpenHPSDR-Thetis-Hermes-Lite2')
SRC = ROOT / 'Project Files/Source'
OUTLINE = ROOT / 'docs/CODE_OUTLINE.md'
GRAPH = ROOT / 'graphify-out/graph.json'
OUTDIR = ROOT / 'docs/files'

SECTION_DIR_HINTS = {
    7: ['wdsp/'],
    8: ['ChannelMaster/'],
    9: ['Console/', 'cmASIO/'],
    12: ['Midi2Cat/'],
    18: ['RawInput/'],
}

# ---------------------------------------------------------------- outline ---
all_src = [p for p in SRC.rglob('*') if p.suffix in ('.cs', '.c', '.cpp', '.h')]
by_basename = defaultdict(list)
for p in all_src:
    by_basename[p.name].append(p.relative_to(SRC).as_posix())

sections = {}          # num -> title
file_info = {}         # relpath -> {'section': num, 'desc': str}
cur_sec = None
sec_re = re.compile(r'^## (\d+)\. (.+)$')
row_re = re.compile(r'^\|(.+?)\|(.+)\|\s*$')
tok_re = re.compile(r'`([A-Za-z0-9_./*]+\.(?:cs|c|cpp|h))`')
glob_re = re.compile(r'`([A-Za-z0-9_./]*\*[A-Za-z0-9_.]*\.(?:cs|c|cpp|h))`')

def resolve(tok, sec):
    if '*' in tok:
        return sorted(q.relative_to(SRC).as_posix() for q in SRC.glob(tok))
    if '/' in tok:
        return [tok] if (SRC / tok).exists() else []
    cands = by_basename.get(tok, [])
    if len(cands) == 1:
        return cands
    for hint in SECTION_DIR_HINTS.get(sec, ['Console/']):
        hits = [c for c in cands if c.startswith(hint)]
        if len(hits) == 1:
            return hits
    return cands[:1]

for line in OUTLINE.read_text(encoding='utf-8').splitlines():
    m = sec_re.match(line)
    if m:
        cur_sec = int(m.group(1))
        sections[cur_sec] = m.group(2)
        continue
    m = row_re.match(line)
    if not m or cur_sec is None:
        continue
    fcell, desc = m.group(1), m.group(2).strip()
    if fcell.strip().startswith('File') or set(fcell.strip()) <= {'-', ' '}:
        continue
    toks = tok_re.findall(fcell) + glob_re.findall(fcell)
    # bare names without extension inside parens (e.g. `buttonts`) are covered
    # by the glob in the same cell, so ignore them
    for tok in toks:
        for rel in resolve(tok, cur_sec):
            if rel not in file_info:
                file_info[rel] = {'section': cur_sec, 'desc': desc}

print(f'outline: {len(file_info)} files across {len(sections)} sections')

# ------------------------------------------------------------------ graph ---
g = json.loads(GRAPH.read_text(encoding='utf-8'))
nodes = {n['id']: n for n in g['nodes']}
nodes_by_file = defaultdict(list)
for n in g['nodes']:
    sf = n.get('source_file') or ''
    if sf:
        nodes_by_file[sf].append(n)

STRUCT = ('contains', 'method', 'defines')
parent = {}
children = defaultdict(list)
in_edges = defaultdict(list)   # target file -> (src node, tgt node, relation)
out_edges = defaultdict(list)  # source file -> same
# structural pass: 'method' edges first so methods attach to their class even
# when a 'contains' edge from the file node exists too
for pass_rels in (('method',), ('contains', 'defines')):
    for l in g['links']:
        rel = l['relation']
        if rel not in pass_rels:
            continue
        s, t = nodes.get(l['source']), nodes.get(l['target'])
        if not s or not t:
            continue
        sf, tf = s.get('source_file') or '', t.get('source_file') or ''
        if sf and sf == tf and l['target'] not in parent and l['target'] != l['source']:
            parent[l['target']] = l['source']
            children[l['source']].append(l['target'])
for l in g['links']:
    rel = l['relation']
    if rel not in ('calls', 'references', 'imports', 'inherits', 'implements'):
        continue
    s, t = nodes.get(l['source']), nodes.get(l['target'])
    if not s or not t:
        continue
    # namespace nodes make cross-file "imports" noise (a shared namespace is
    # attributed to whichever file declared it first) — skip them
    if t.get('type') == 'namespace' or s.get('type') == 'namespace':
        continue
    sf, tf = s.get('source_file') or '', t.get('source_file') or ''
    if sf and tf and sf != tf:
        out_edges[sf].append((s, t, rel))
        in_edges[tf].append((s, t, rel))

def lineno(n):
    loc = n.get('source_location') or ''
    m = re.match(r'L(\d+)', loc)
    return int(m.group(1)) if m else 0

def is_callable(n):
    return n['label'].startswith('.') or n['label'].endswith('()')

# ----------------------------------------------------------------- render ---
def render_tree(nid, depth, lines, rel_file):
    n = nodes[nid]
    kids = sorted(children.get(nid, []), key=lambda k: lineno(nodes[k]))
    label, ln = n['label'], lineno(n)
    if is_callable(n):
        lines.append('  ' * depth + f'- `{label}` — L{ln}')
        for k in kids:
            render_tree(k, depth + 1, lines, rel_file)
    else:
        kind = 'namespace' if n.get('type') == 'namespace' else 'type'
        lines.append('')
        lines.append('#### ' + f'`{label}` ({kind}, L{ln})')
        lines.append('')
        if not kids:
            lines.append('_No extracted members._')
        for k in kids:
            render_tree(k, depth, lines, rel_file)

def agg_by_file(pairs, key):
    c = defaultdict(Counter)
    for s, t, rel in pairs:
        c[(s if key == 'out' else s)['source_file'] if False else (t['source_file'] if key == 'out' else s['source_file'])][rel] += 1
    return c

generated = []
for rel, info in sorted(file_info.items()):
    fnodes = sorted(nodes_by_file.get(rel, []), key=lineno)
    depth = rel.count('/')
    up = '../' * (depth + 1)          # from docs/files/<dirs> back to docs/
    sec_num, sec_title = info['section'], sections[info['section']]
    anchor = re.sub(r'[^a-z0-9 -]', '', f'{sec_num} {sections[info["section"]]}'.lower()).replace(' ', '-')

    lines = [f'# `{rel}`', '']
    lines.append(f'**Functional area:** [{sec_num}. {sec_title}]({up}CODE_OUTLINE.md#{anchor})')
    lines.append('')
    lines.append(f'**Role:** {info["desc"]}')
    lines.append('')

    # --- how it is used ---
    lines.append('## How this file is used')
    lines.append('')
    inc = defaultdict(Counter)
    for s, t, r in in_edges.get(rel, []):
        inc[s['source_file']][r] += 1
    outc = defaultdict(Counter)
    for s, t, r in out_edges.get(rel, []):
        outc[t['source_file']][r] += 1

    def fmt_deps(c, verb):
        if not c:
            return [f'- {verb}: none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).']
        ranked = sorted(c.items(), key=lambda kv: -sum(kv[1].values()))
        out = []
        for f, rels in ranked[:12]:
            det = ', '.join(f'{r} ×{n}' for r, n in rels.most_common())
            out.append(f'  - `{f}` ({det})')
        if len(ranked) > 12:
            out.append(f'  - …and {len(ranked) - 12} more files')
        return [f'- {verb}:'] + out

    lines += fmt_deps(inc, 'Used by (incoming references from other files)')
    lines += fmt_deps(outc, 'Uses (outgoing references to other files)')

    # top externally-called symbols
    sym_calls = Counter()
    for s, t, r in in_edges.get(rel, []):
        if is_callable(t):
            sym_calls[t['label']] += 1
    if sym_calls:
        lines.append('- Most-referenced symbols from other files: '
                     + ', '.join(f'`{k}` (×{v})' for k, v in sym_calls.most_common(8)))
    lines.append('')

    # --- outline ---
    lines.append('## Outline')
    lines.append('')
    if not fnodes:
        lines.append('_No symbols extracted for this file (it may be data-only, '
                     'generated, or vendored)._')
    else:
        roots = [n for n in fnodes if n['id'] not in parent]
        placed = set()
        base = rel.rsplit('/', 1)[-1]
        file_roots = [n for n in roots if n['label'] == base]
        other_roots = [n for n in roots if n['label'] != base]
        top = []
        for fr in file_roots:
            top += sorted(children.get(fr['id'], []), key=lambda k: lineno(nodes[k]))
        top += [n['id'] for n in other_roots]
        # free functions first, then types
        funcs = [nid for nid in top if is_callable(nodes[nid])]
        types_ = [nid for nid in top if not is_callable(nodes[nid])]
        if funcs:
            lines.append('### Functions')
            lines.append('')
            for nid in funcs:
                render_tree(nid, 0, lines, rel)
            lines.append('')
        if types_:
            lines.append('### Types')
            for nid in types_:
                render_tree(nid, 0, lines, rel)
            lines.append('')

    lines.append('---')
    lines.append('_Generated from the graphify knowledge graph '
                 f'(`graphify-out/graph.json`); line numbers refer to '
                 f'`Project Files/Source/{rel}`. Regenerate after code changes '
                 'with `graphify update "Project Files/Source"` followed by '
                 '`python docs/tools/gen_file_docs.py`._')
    outp = OUTDIR / (rel + '.md')
    outp.parent.mkdir(parents=True, exist_ok=True)
    outp.write_text('\n'.join(lines) + '\n', encoding='utf-8')
    generated.append((rel, sec_num, sum(1 for n in fnodes if is_callable(n))))

# ------------------------------------------------------------------ index ---
idx = ['# Per-file source documentation', '',
       'One page per source file listed in [CODE_OUTLINE.md](../CODE_OUTLINE.md): '
       'the file\'s role, how it is used (graph-derived callers/callees), and an '
       'outline of its classes, methods, and functions with line numbers.', '']
by_sec = defaultdict(list)
for rel, sec, nfuncs in generated:
    by_sec[sec].append((rel, nfuncs))
for sec in sorted(by_sec):
    idx.append(f'## {sec}. {sections[sec]}')
    idx.append('')
    for rel, nfuncs in sorted(by_sec[sec]):
        idx.append(f'- [`{rel}`]({rel}.md) — {nfuncs} functions/methods')
    idx.append('')
(OUTDIR / 'README.md').write_text('\n'.join(idx) + '\n', encoding='utf-8')

print(f'generated {len(generated)} file docs + index at {OUTDIR}')
print('largest:', sorted(generated, key=lambda x: -x[2])[:5])
