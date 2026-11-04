#!/usr/bin/env python3
"""Compile guarded recovery declarations for every independent tag selection."""

import os
from pathlib import Path
import subprocess
import tempfile


ROOT = Path(__file__).resolve().parents[1]


def generate(destination: Path) -> dict[str, bytes]:
    subprocess.run(
        ["haxe", "-cp", "src", "-cp", "test/compiler", "--run",
         "GoRecoveryDispatchContract", str(destination)],
        cwd=ROOT, check=True, timeout=60,
    )
    return {path.name: path.read_bytes() for path in destination.glob("*.go")}


def main() -> None:
    with tempfile.TemporaryDirectory(prefix="go-recovery-dispatch-") as temporary:
        root = Path(temporary)
        first, second = root / "first", root / "second"
        first.mkdir()
        second.mkdir()
        if generate(first) != generate(second):
            raise AssertionError("Recovery output changed between identical compilations")
        (first / "go.mod").write_text("module recoverycontract\n\ngo 1.26\n")
        for tags in ["", "alpha", "beta", "alpha,beta"]:
            print(f"Recovery dispatch: tags={tags!r}", flush=True)
            subprocess.run(["go", "run", "-tags", tags, "."], cwd=first,
                           env=os.environ, check=True, timeout=180)


if __name__ == "__main__":
    main()
