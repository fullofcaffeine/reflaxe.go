#!/usr/bin/env python3
"""Prove that editing one module does not rename recovery helpers elsewhere.

What: compiles one fixture three times: unchanged, with the first use of two
helpers swapped in one module, and with new helpers demanded by that module.
Why: recovery helpers were numbered by first use across the whole compilation,
so an edit in one module renamed helper calls in unrelated generated files.
How: every generated file outside the edited module must stay byte-identical.
The file that holds helper declarations (main.go) may only gain lines when the
edit adds helpers, and must not change at all when only the use order changes.
"""

from difflib import SequenceMatcher
from pathlib import Path
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]
FIXTURE = ROOT / "test/snapshot/core/recovery_helper_names_stable"
EDITED_MODULE = "module_edited.go"
HELPER_DECLARATIONS = "main.go"


def compile_fixture(output: Path, *defines: str) -> dict[str, bytes]:
    command = [
        "haxe", "-cp", str(FIXTURE), "-lib", "reflaxe.go", "-main", "Main",
        "-D", f"go_output={output}", "-D", "go_no_build",
        "-D", "reflaxe_go_hxrt_features=core", "-D", "reflaxe_go_strict_examples",
        "-D", "reflaxe.dont_output_metadata_id", "-D", "no-traces", "-D", "no_traces",
    ]
    for define in defines:
        command += ["-D", define]
    subprocess.run(command, cwd=ROOT, check=True, timeout=120)
    return {
        path.relative_to(output).as_posix(): path.read_bytes()
        for path in sorted(output.rglob("*"))
        if path.is_file()
    }


def changed_files(before: dict[str, bytes], after: dict[str, bytes]) -> list[str]:
    return sorted(name for name in set(before) | set(after) if before.get(name) != after.get(name))


def removed_lines(before: bytes, after: bytes) -> list[str]:
    old = before.decode("utf-8").splitlines()
    new = after.decode("utf-8").splitlines()
    removed = []
    for tag, start, end, _, _ in SequenceMatcher(None, old, new, autojunk=False).get_opcodes():
        if tag in ("replace", "delete"):
            removed.extend(old[start:end])
    return removed


def main() -> None:
    with tempfile.TemporaryDirectory(prefix="recovery-helper-names-") as temporary:
        root = Path(temporary)
        baseline = compile_fixture(root / "baseline")
        reordered = compile_fixture(root / "reordered", "recovery_names_reorder")
        extended = compile_fixture(root / "extended", "recovery_names_edit")

    for name in (EDITED_MODULE, HELPER_DECLARATIONS, "module_stable.go"):
        if name not in baseline:
            raise AssertionError(f"Fixture no longer generates {name}: {sorted(baseline)}")

    changed = changed_files(baseline, reordered)
    if changed != [EDITED_MODULE]:
        raise AssertionError(
            "Reordering helper uses inside one module must change only that module, "
            f"but these files changed: {changed}"
        )

    changed = changed_files(baseline, extended)
    unexpected = [name for name in changed if name not in (EDITED_MODULE, HELPER_DECLARATIONS)]
    if unexpected or EDITED_MODULE not in changed:
        raise AssertionError(
            "Adding helper uses inside one module must change only that module and the "
            f"helper declarations, but these files changed: {changed}"
        )
    removed = removed_lines(baseline[HELPER_DECLARATIONS], extended[HELPER_DECLARATIONS])
    if removed:
        raise AssertionError(
            "Adding helpers must not rename or reorder existing helper declarations; "
            "removed lines:\n" + "\n".join(removed)
        )
    print("Recovery helper names are stable across unrelated module edits.")


if __name__ == "__main__":
    main()
