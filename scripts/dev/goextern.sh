#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

cd "$ROOT/tools/goextern"
go run . --out "$ROOT/gen/goextern" "$@"
