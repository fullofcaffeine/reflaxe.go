#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
DIST_DIR="${DIST_DIR:-$ROOT/dist/examples}"
TARGETS=("linux/amd64" "linux/arm64" "darwin/arm64" "windows/amd64")
PROFILES=("portable" "gopher" "metal")

mkdir -p "$DIST_DIR"

if [ "$#" -gt 0 ]; then
  EXAMPLES=("$@")
else
  mapfile -t EXAMPLES < <(find "$ROOT/examples" -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort)
fi

built=0
for example in "${EXAMPLES[@]}"; do
  for profile in "${PROFILES[@]}"; do
    src="$ROOT/examples/$example/generated/$profile"
    if [ ! -f "$src/go.mod" ]; then
      continue
    fi

    for target in "${TARGETS[@]}"; do
      goos="${target%/*}"
      goarch="${target#*/}"
      outdir="$DIST_DIR/$example/$profile/${goos}_${goarch}"
      mkdir -p "$outdir"

      bin="$example"
      if [ "$goos" = "windows" ]; then
        bin="$bin.exe"
      fi

      echo "Building $example/$profile for $goos/$goarch"
      (
        cd "$src"
        GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 go build -o "$outdir/$bin" .
      )
      built=$((built + 1))
    done
  done
done

if [ "$built" -eq 0 ]; then
  echo "No binaries were built. Ensure examples/*/generated/<profile> exist."
  exit 1
fi

python3 - "$DIST_DIR" <<'PY'
from __future__ import annotations

import hashlib
import json
from pathlib import Path
import sys


def sha256(path: Path) -> str:
    hasher = hashlib.sha256()
    with path.open("rb") as handle:
        for chunk in iter(lambda: handle.read(1024 * 1024), b""):
            hasher.update(chunk)
    return hasher.hexdigest()


dist = Path(sys.argv[1]).resolve()
files: list[dict[str, object]] = []

for path in sorted(dist.rglob("*")):
    if path.is_dir():
        continue
    if path.name in {"manifest.json", "checksums.txt"}:
        continue
    digest = sha256(path)
    rel = path.relative_to(dist).as_posix()
    files.append({
        "path": rel,
        "size": path.stat().st_size,
        "sha256": digest,
    })

checksums = dist / "checksums.txt"
checksums.write_text("\n".join(f"{entry['sha256']}  {entry['path']}" for entry in files) + "\n", encoding="utf-8")

manifest = {
    "artifact_count": len(files),
    "files": files,
}
(dist / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")

print(f"Wrote {checksums}")
print(f"Wrote {dist / 'manifest.json'}")
PY
