# Examples

This repository ships two first-class multi-profile examples:

- `profile_storyboard`: compact profile-idiom walkthrough.
- `tui_todo`: canonical complex app with deterministic pseudo-TUI flow.

Both examples compile from one Haxe codebase into:

- `portable`
- `gopher`
- `metal`

## Quick commands

Compile and validate all examples:

```bash
python3 test/run-examples.py
```

Update committed generated Go trees from fresh outputs:

```bash
python3 test/run-examples.py --bless-generated
```

Build binary artifacts from committed generated Go trees:

```bash
bash scripts/examples/build-binaries.sh
```
