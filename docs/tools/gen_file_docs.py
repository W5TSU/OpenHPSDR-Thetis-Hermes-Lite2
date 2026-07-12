"""Generate per-source-file documentation from CODE_OUTLINE.md + graphify graph.json.

For every source file referenced in docs/CODE_OUTLINE.md, writes
docs/files/<relpath>.md containing:
  - the file's role (from the outline table row) and its functional area
  - how it is used: cross-file callers/callees aggregated from the graph,
    plus the file's most externally-called symbols
  - an outline of classes/methods/functions with line numbers; each function
    carries its signature, a description (source doc-comment where present,
    naming-convention heuristic otherwise), and its callers from the graph
Also writes docs/files/README.md as an index grouped by functional area.

Run: python docs/tools/gen_file_docs.py   (after `graphify update "Project Files/Source"`)
"""
import json
import re
from collections import Counter, defaultdict
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
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
    for tok in tok_re.findall(fcell) + glob_re.findall(fcell):
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

parent = {}
children = defaultdict(list)
in_edges = defaultdict(list)   # target file -> (src node, tgt node, relation)
out_edges = defaultdict(list)  # source file -> same
callers = defaultdict(list)    # callable node id -> [caller node ids]
referrers = defaultdict(list)  # callable node id -> [referencing node ids]
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
    if rel == 'calls':
        callers[l['target']].append(l['source'])
    elif rel == 'references':
        referrers[l['target']].append(l['source'])
    # namespace nodes make cross-file "imports" noise (a shared namespace is
    # attributed to whichever file declared it first) — skip them
    if t.get('type') == 'namespace' or s.get('type') == 'namespace':
        continue
    sf, tf = s.get('source_file') or '', t.get('source_file') or ''
    if sf and tf and sf != tf:
        out_edges[sf].append((s, t, rel))
        in_edges[tf].append((s, t, rel))

# C# method labels by lowercased bare name -> files, used to spot the P/Invoke
# extern declaration that bridges a C function into the console (the graph
# does not link across the language boundary)
cs_decls = defaultdict(set)
for n in g['nodes']:
    sf = n.get('source_file') or ''
    lbl = n.get('label', '')
    if sf.endswith('.cs') and lbl.startswith('.') and lbl.endswith('()'):
        cs_decls[lbl[1:-2].lower()].add(sf)

def lineno(n):
    loc = n.get('source_location') or ''
    m = re.match(r'L(\d+)', loc)
    return int(m.group(1)) if m else 0

def is_callable(n):
    return n['label'].startswith('.') or n['label'].endswith('()')

# ------------------------------------------------- source-text extraction ---
_lines_cache = {}

def get_lines(rel):
    if rel not in _lines_cache:
        p = SRC / rel
        try:
            _lines_cache[rel] = p.read_text(encoding='utf-8', errors='replace').splitlines()
        except OSError:
            _lines_cache[rel] = []
    return _lines_cache[rel]

def clean_ws(s):
    return re.sub(r'\s+', ' ', s).strip()

def extract_signature(rel, ln):
    """Signature text starting at the node's line, cut at '{' or ';'."""
    lines = get_lines(rel)
    i = ln - 1
    if i < 0 or i >= len(lines):
        return ''
    parts = []
    for j in range(i, min(i + 5, len(lines))):
        parts.append(re.sub(r'//.*$', '', lines[j]).strip())
        joined = ' '.join(parts)
        for stop in ('{', ';'):
            if stop in joined:
                return clean_ws(joined.split(stop)[0])[:180]
        if joined.rstrip().endswith(')'):
            return clean_ws(joined)[:180]
    return clean_ws(' '.join(parts))[:180]

_noise_line = re.compile(r'^[\s/*=\-#]*$')

def extract_comment(rel, ln):
    """Doc comment immediately above the declaration, if any."""
    lines = get_lines(rel)
    i = ln - 2
    while i >= 0 and lines[i].strip().startswith('['):
        i -= 1  # skip C# attributes
    if i < 0 or i >= len(lines):
        return ''
    s = lines[i].strip()
    buf = []
    if s.startswith('///'):
        while i >= 0 and lines[i].strip().startswith('///'):
            buf.insert(0, lines[i].strip()[3:].strip())
            i -= 1
        text = ' '.join(buf)
        m = re.search(r'<summary>(.*?)</summary>', text, re.S)
        if m:
            text = m.group(1)
        text = re.sub(r'<[^>]+>', ' ', text)
    elif s.startswith('//'):
        while i >= 0 and lines[i].strip().startswith('//') and len(buf) < 8:
            t = lines[i].strip().lstrip('/').strip()
            if not _noise_line.match(t):
                buf.insert(0, t)
            i -= 1
        text = ' '.join(buf)
    elif s.endswith('*/'):
        j = i
        while j >= 0 and '/*' not in lines[j] and i - j < 40:
            j -= 1
        if j < 0 or '/*' not in lines[j]:
            return ''
        raw = [re.sub(r'^\s*/?\*+/?', '', l).strip() for l in lines[j:i + 1]]
        text = ' '.join(t for t in raw if not _noise_line.match(t))
    else:
        return ''
    text = clean_ws(text)
    if len(text) < 5 or 'copyright' in text.lower() or 'license' in text.lower():
        return ''
    if len(text) > 300:
        text = text[:297].rsplit(' ', 1)[0] + '…'
    return text

# --------------------------------------------------- heuristic descriptions -
def spaced(name):
    words = re.sub(r'([a-z0-9])([A-Z])', r'\1 \2', name).replace('_', ' ')
    return clean_ws(words).lower()

CS_EVENTS = {
    'Click': 'is clicked', 'DoubleClick': 'is double-clicked',
    'CheckedChanged': 'checked state changes',
    'SelectedIndexChanged': 'selection changes', 'TextChanged': 'text changes',
    'ValueChanged': 'value changes', 'Scroll': 'is scrolled',
    'MouseDown': 'receives a mouse-down', 'MouseUp': 'receives a mouse-up',
    'MouseMove': 'receives mouse movement', 'MouseWheel': 'receives a mouse wheel event',
    'MouseEnter': 'is entered by the mouse', 'MouseLeave': 'is left by the mouse',
    'KeyDown': 'receives a key-down', 'KeyUp': 'receives a key-up',
    'KeyPress': 'receives a key press', 'Load': 'loads',
    'FormClosing': 'is closing', 'FormClosed': 'has closed',
    'Closing': 'is closing', 'Closed': 'has closed', 'Shown': 'is shown',
    'Paint': 'repaints', 'Resize': 'is resized', 'Tick': 'timer fires',
    'LostFocus': 'loses focus', 'GotFocus': 'gains focus',
    'Leave': 'is left', 'Enter': 'is entered',
    'DropDown': 'drops down', 'DropDownClosed': 'drop-down closes',
    'DoWork': 'background worker runs', 'RunWorkerCompleted': 'background work completes',
}
CS_VERBS = ('Update', 'Refresh', 'Initialize', 'Init', 'Setup', 'Add', 'Remove',
            'Create', 'Delete', 'Save', 'Load', 'Read', 'Write', 'Send',
            'Receive', 'Parse', 'Handle', 'Process', 'Start', 'Stop', 'Enable',
            'Disable', 'Toggle', 'Show', 'Hide', 'Open', 'Close', 'Apply',
            'Build', 'Clear', 'Reset', 'Check', 'Find', 'Select', 'Draw',
            'Render', 'Calculate', 'Compute', 'Convert', 'Format', 'Restore',
            'Register', 'Unregister', 'Connect', 'Disconnect')

def heuristic_desc(label, rel, class_name):
    name = label.lstrip('.').rstrip('()')
    is_c = rel.endswith(('.c', '.cpp', '.h')) and not label.startswith('.')
    if is_c:
        for pre, tmpl in (
            ('create_', 'Constructor for the {b} block: allocates its state/buffers and computes initial coefficients.'),
            ('destroy_', 'Destroys the {b} block, freeing its allocated buffers.'),
            ('flush_', 'Flushes (zeroes) the {b} block’s internal buffers/state.'),
            ('setBuffers_', 'Re-points the {b} block’s input/output buffers (called when the channel’s buffers change).'),
            ('setSamplerate_', 'Reconfigures the {b} block for a new sample rate.'),
            ('setSize_', 'Reconfigures the {b} block for a new buffer size.'),
        ):
            if name.startswith(pre):
                return tmpl.format(b='`' + name[len(pre):] + '`')
        if name.startswith('x') and len(name) > 2 and name[1].islower():
            return (f'Runs the `{name[1:]}` block on one buffer of samples — the '
                    'per-buffer processing entry called from the owning chain.')
        m = re.match(r'^(Set|Get)([A-Z]\w+)$', name)
        if m:
            verb = 'Sets' if m.group(1) == 'Set' else 'Returns'
            return (f'{verb} {spaced(m.group(2))} — API {m.group(1).lower()}ter, '
                    'typically called from the console via P/Invoke.')
        m = re.match(r'^(RXA|TXA)([A-Z]\w+)$', name)
        if m:
            return (f'{m.group(1)} chain operation — {spaced(m.group(2))}; '
                    'part of the receive/transmit chain API.')
        return ''
    # C# members
    m = re.match(r'^(.*)_([A-Za-z]+)$', name)
    if m and m.group(2) in CS_EVENTS:
        return f'WinForms event handler: runs when `{m.group(1)}` {CS_EVENTS[m.group(2)]}.'
    if class_name and name == class_name:
        return f'Constructor — creates and initializes a `{class_name}` instance.'
    if name == 'InitializeComponent':
        return 'Designer-generated UI construction (creates and lays out the form’s controls).'
    if name == 'Dispose':
        return 'Releases the object’s resources.'
    if name == 'ToString':
        return 'Returns the string representation.'
    m = re.match(r'^(get|Get)([A-Z_]\w*)$', name)
    if m:
        return f'Returns {spaced(m.group(2))}.'
    m = re.match(r'^(set|Set)([A-Z_]\w*)$', name)
    if m:
        return f'Sets {spaced(m.group(2))}.'
    m = re.match(r'^On([A-Z]\w+)$', name)
    if m:
        return f'Handles/raises the {spaced(m.group(1))} event.'
    for v in CS_VERBS:
        if name.startswith(v) and len(name) > len(v) and name[len(v)].isupper():
            verb = v + ('es' if v.endswith(('s', 'sh', 'ch')) else 's')
            return f'{verb} {spaced(name[len(v):])}.'
    return ''

def called_by_line(node, rel):
    nid = node['id']
    cs = [nodes[c] for c in callers.get(nid, []) if c in nodes]
    if cs:
        seen, items = set(), []
        for c in sorted(cs, key=lambda n: (n.get('source_file') != rel,
                                           n.get('source_file') or '', lineno(n))):
            key = (c['label'], c.get('source_file'))
            if key in seen:
                continue
            seen.add(key)
            where = 'same file' if c.get('source_file') == rel else f"`{c.get('source_file')}`"
            items.append(f"`{c['label']}` ({where})")
        shown = items[:6]
        more = f' — and {len(items) - 6} more' if len(items) > 6 else ''
        return 'Called by: ' + ', '.join(shown) + more
    # no direct call edges — explain how it is reached
    name = node['label'].lstrip('.').rstrip('()')
    if rel.endswith(('.c', '.cpp')):
        decls = sorted(cs_decls.get(name.lower(), []))
        if decls:
            return ('Called from C# via P/Invoke — declared/wrapped in '
                    + ', '.join(f'`{d}`' for d in decls[:3]) + '.')
    refs = [nodes[r] for r in referrers.get(nid, []) if r in nodes]
    if refs:
        seen, items = set(), []
        for c in refs:
            key = (c['label'], c.get('source_file'))
            if key in seen:
                continue
            seen.add(key)
            where = 'same file' if c.get('source_file') == rel else f"`{c.get('source_file')}`"
            items.append(f"`{c['label']}` ({where})")
        return ('Referenced (not directly called) by: ' + ', '.join(items[:5])
                + ' — typically event wiring or a delegate/callback assignment.')
    mm = re.match(r'^(.*)_([A-Za-z]+)$', name)
    if mm and mm.group(2) in CS_EVENTS:
        return 'Called by: WinForms event wiring at runtime (no static call sites).'
    return ('No callers found in the graph — likely invoked via P/Invoke, '
            'UI/event wiring, a delegate, a thread start, or externally.')

# ----------------------------------------------------------------- render ---
def render_callable(nid, depth, lines, rel_file, class_name):
    n = nodes[nid]
    ind = '  ' * depth
    ln = lineno(n)
    sig = extract_signature(rel_file, ln)
    head = f'{ind}- **`{n["label"]}`** — L{ln}'
    if sig and sig != n['label']:
        head += f' — `{sig}`'
    lines.append(head)
    desc = extract_comment(rel_file, ln) or heuristic_desc(n['label'], rel_file, class_name)
    if desc:
        lines.append(f'{ind}  {desc}')
    lines.append(f'{ind}  {called_by_line(n, rel_file)}')

def render_tree(nid, depth, lines, rel_file, class_name=None):
    n = nodes[nid]
    kids = sorted(children.get(nid, []), key=lambda k: lineno(nodes[k]))
    if is_callable(n):
        render_callable(nid, depth, lines, rel_file, class_name)
        for k in kids:
            render_tree(k, depth + 1, lines, rel_file, class_name)
    else:
        kind = 'namespace' if n.get('type') == 'namespace' else 'type'
        lines.append('')
        lines.append('#### ' + f'`{n["label"]}` ({kind}, L{lineno(n)})')
        lines.append('')
        if not kids:
            lines.append('_No extracted members._')
        for k in kids:
            render_tree(k, depth, lines, rel_file, n['label'])

generated = []
for rel, info in sorted(file_info.items()):
    fnodes = sorted(nodes_by_file.get(rel, []), key=lineno)
    depth = rel.count('/')
    up = '../' * (depth + 1)          # from docs/files/<dirs> back to docs/
    sec_num, sec_title = info['section'], sections[info['section']]
    anchor = re.sub(r'[^a-z0-9 -]', '', f'{sec_num} {sec_title}'.lower()).replace(' ', '-')

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
    lines.append('_Each entry: symbol — line — signature, then a description '
                 '(from source comments where present, otherwise inferred from '
                 'naming conventions) and its callers as recorded in the graph._')
    lines.append('')
    if not fnodes:
        lines.append('_No symbols extracted for this file (it may be data-only, '
                     'generated, or vendored)._')
    else:
        roots = [n for n in fnodes if n['id'] not in parent]
        base = rel.rsplit('/', 1)[-1]
        file_roots = [n for n in roots if n['label'] == base]
        other_roots = [n for n in roots if n['label'] != base]
        top = []
        for fr in file_roots:
            top += sorted(children.get(fr['id'], []), key=lambda k: lineno(nodes[k]))
        top += [n['id'] for n in other_roots]
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
       'outline of its classes, methods, and functions — each with line number, '
       'signature, description, and callers.', '']
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
