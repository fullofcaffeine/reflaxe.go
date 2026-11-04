#!/usr/bin/env python3
"""Prove nested and native recovery across independent Go file selections."""

from itertools import combinations
from pathlib import Path
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]
FIXTURE = ROOT / "test/snapshot/core/typed_go_build_constraint_overlap"


def main() -> None:
    with tempfile.TemporaryDirectory(prefix="conditional-recovery-runtime-") as temporary:
        output = Path(temporary) / "out"
        subprocess.run([
            "haxe", "-cp", str(FIXTURE), "-lib", "reflaxe.go", "-main", "Main",
            "-D", f"go_output={output}", "-D", "go_no_build",
            "-D", "reflaxe_go_strict_examples", "-D", "reflaxe.dont_output_metadata_id",
        ], cwd=ROOT, check=True, timeout=120)
        for size in range(4):
            for selected in combinations(["alpha", "beta", "gamma"], size):
                tags = ",".join(selected)
                # Independent behavioral oracle: each admitted feature registers once.
                expected = int("alpha" in selected) + int("beta" in selected)
                expected += int(("alpha" in selected or "beta" in selected) and "gamma" not in selected)
                print(f"Conditional recovery: tags={tags!r}, expected={expected}", flush=True)
                result = subprocess.run(["go", "run", "-tags", tags, "."], cwd=output,
                                        text=True, capture_output=True, timeout=180)
                if result.returncode or result.stdout != f"{expected}\n":
                    raise AssertionError(f"tags={tags!r}: exit={result.returncode}\n{result.stdout}{result.stderr}")


    fixture = ROOT / "test/snapshot/core/conditional_nominal_type_test"
    reference = subprocess.run(["haxe", "-cp", str(fixture), "-main", "Main", "--interp"],
                               cwd=ROOT, text=True, capture_output=True, check=True, timeout=120)
    if reference.stdout != "1\n":
        raise AssertionError(f"Haxe nominal authority changed: {reference.stdout}{reference.stderr}")
    with tempfile.TemporaryDirectory(prefix="conditional-nominal-runtime-") as temporary:
        output = Path(temporary) / "out"
        subprocess.run([
            "haxe", "-cp", str(fixture), "-lib", "reflaxe.go", "-main", "Main",
            "-D", f"go_output={output}", "-D", "go_no_build",
            "-D", "reflaxe_go_strict_examples", "-D", "reflaxe.dont_output_metadata_id",
        ], cwd=ROOT, check=True, timeout=120)
        for tags, expected in [("", "0\n"), ("exception_probe", "1\n")]:
            print(f"Conditional nominal membership: tags={tags!r}", flush=True)
            result = subprocess.run(["go", "run", "-tags", tags, "."], cwd=output,
                                    text=True, capture_output=True, timeout=180)
            if result.returncode or result.stdout != expected:
                raise AssertionError(f"tags={tags!r}: exit={result.returncode}\n{result.stdout}{result.stderr}")


if __name__ == "__main__":
    main()
