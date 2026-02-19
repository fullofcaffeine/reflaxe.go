#!/usr/bin/env python3

from __future__ import annotations

import argparse
from pathlib import Path
import shutil

ROOT = Path(__file__).resolve().parent.parent.parent
EXAMPLES_ROOT = ROOT / "examples"
PROFILES = ("portable", "gopher", "metal")
EXCLUDE_NAMES = {"go.sum", "_GeneratedFiles.json", ".DS_Store"}
EXCLUDE_DIRS = {".cache"}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Sync generated example Go trees from out_<profile>")
    parser.add_argument("--example", action="append", default=[], help="Example name filter (repeatable)")
    parser.add_argument("--profile", action="append", default=[], help="Profile filter (repeatable)")
    parser.add_argument("--bless", action="store_true", help="Copy out_<profile> to generated/<profile> before validation")
    return parser.parse_args()


def all_files(root: Path) -> list[Path]:
    if not root.exists():
        return []
    files: list[Path] = []
    for path in sorted(root.rglob("*")):
        if path.is_dir():
            continue
        if path.name in EXCLUDE_NAMES:
            continue
        if any(part in EXCLUDE_DIRS for part in path.parts):
            continue
        files.append(path)
    return files


def collect_tree_deltas(left: Path, right: Path) -> list[str]:
    left_files = {path.relative_to(left): path for path in all_files(left)} if left.exists() else {}
    right_files = {path.relative_to(right): path for path in all_files(right)} if right.exists() else {}

    rels = sorted(set(left_files) | set(right_files))
    deltas: list[str] = []
    for rel in rels:
        l = left_files.get(rel)
        r = right_files.get(rel)
        if l is None:
            deltas.append(f"Only in {right}: {rel.as_posix()}")
            continue
        if r is None:
            deltas.append(f"Only in {left}: {rel.as_posix()}")
            continue
        if l.read_text(encoding="utf-8", errors="replace") != r.read_text(encoding="utf-8", errors="replace"):
            deltas.append(f"Diff: {rel.as_posix()}")
    return deltas


def copy_tree(source: Path, target: Path) -> None:
    if target.exists():
        shutil.rmtree(target)
    target.mkdir(parents=True, exist_ok=True)
    for path in all_files(source):
        rel = path.relative_to(source)
        dest = target / rel
        dest.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(path, dest)


def main() -> int:
    args = parse_args()

    examples = [path for path in sorted(EXAMPLES_ROOT.iterdir()) if path.is_dir()]
    selected_examples = {item.strip() for item in args.example if item.strip()} if args.example else None
    selected_profiles = {item.strip() for item in args.profile if item.strip()} if args.profile else set(PROFILES)

    failures: list[str] = []

    for example_dir in examples:
        if selected_examples is not None and example_dir.name not in selected_examples:
            continue

        for profile in PROFILES:
            if profile not in selected_profiles:
                continue

            out_dir = example_dir / f"out_{profile}"
            generated_dir = example_dir / "generated" / profile

            compile_hxml = example_dir / f"compile.{profile}.hxml"
            if not compile_hxml.exists():
                continue

            if not out_dir.exists():
                failures.append(f"{example_dir.name}/{profile}: missing {out_dir}")
                continue

            if args.bless:
                copy_tree(out_dir, generated_dir)

            if not generated_dir.exists():
                failures.append(f"{example_dir.name}/{profile}: missing {generated_dir}")
                continue

            deltas = collect_tree_deltas(generated_dir, out_dir)
            if deltas:
                preview = "\n".join(deltas[:10])
                if len(deltas) > 10:
                    preview += f"\n... and {len(deltas) - 10} more"
                failures.append(f"{example_dir.name}/{profile}:\n{preview}")
            else:
                print(f"PASS {example_dir.name}/{profile}")

    if failures:
        print("\nGenerated trees are out of sync:")
        for item in failures:
            print(item)
        return 1

    print("\nAll generated trees are in sync")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
