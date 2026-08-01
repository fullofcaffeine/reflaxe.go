#!/usr/bin/env python3
"""Run the watcher shipped by the installed reflaxe.go package."""

from __future__ import annotations

from pathlib import Path
import subprocess
import sys


def find_reflaxe_go_tools() -> tuple[Path, Path]:
	result = subprocess.run(
		["haxelib", "path", "reflaxe.go"], text=True, capture_output=True, check=False
	)
	if result.returncode == 0:
		for line in result.stdout.splitlines():
			candidate = Path(line.strip())
			for root in (candidate, candidate.parent, candidate.parent.parent):
				for tools in (root / "tools", root / "scripts" / "dev"):
					watcher = tools / "haxe_go_watch.py"
					wrapper = tools / "go-hx.sh"
					if watcher.is_file() and wrapper.is_file():
						return watcher, wrapper
	print("[hx-go-watch] error: installed reflaxe.go package does not contain its dev tools", file=sys.stderr)
	print("[hx-go-watch] run 'npm run setup' and try again", file=sys.stderr)
	raise SystemExit(2)


watcher, wrapper = find_reflaxe_go_tools()
project = Path(__file__).resolve().parents[2]
command = [
	sys.executable,
	str(watcher),
	"--wrapper",
	str(wrapper),
	"--project",
	str(project),
	*sys.argv[1:],
]
raise SystemExit(subprocess.call(command, cwd=project))
