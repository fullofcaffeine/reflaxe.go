#!/usr/bin/env bash

set -euo pipefail

haxe_version="${HAXE_VERSION:-4.3.7}"
neko_version="${NEKO_VERSION:-2.3.0}"
neko_release_tag="${NEKO_RELEASE_TAG:-v2-3-0}"
retry_attempts="${HAXE_SETUP_RETRY_ATTEMPTS:-4}"
retry_delay_sec="${HAXE_SETUP_RETRY_DELAY_SEC:-5}"
retry_backoff="${HAXE_SETUP_RETRY_BACKOFF:-linear}"
install_root="${HAXE_SETUP_INSTALL_ROOT:-${RUNNER_TEMP:-/tmp}/haxe-fallback}"

if [[ "$(uname -s)" != "Linux" ]]; then
  echo "[haxe-setup] error: setup-haxe-linux-fallback.sh is Linux-only" >&2
  exit 2
fi

if ! [[ "$retry_attempts" =~ ^[1-9][0-9]*$ ]]; then
  echo "[haxe-setup] error: HAXE_SETUP_RETRY_ATTEMPTS must be a positive integer" >&2
  exit 2
fi

if ! [[ "$retry_delay_sec" =~ ^[0-9]+$ ]]; then
  echo "[haxe-setup] error: HAXE_SETUP_RETRY_DELAY_SEC must be an integer >= 0" >&2
  exit 2
fi

case "$retry_backoff" in
  linear|exponential)
    ;;
  *)
    echo "[haxe-setup] error: HAXE_SETUP_RETRY_BACKOFF must be linear or exponential" >&2
    exit 2
    ;;
esac

compute_delay() {
  local attempt="$1"
  if [[ "$retry_backoff" == "exponential" ]]; then
    local multiplier=1
    local i=1
    while (( i < attempt )); do
      multiplier=$((multiplier * 2))
      i=$((i + 1))
    done
    echo $((retry_delay_sec * multiplier))
  else
    echo $((retry_delay_sec * attempt))
  fi
}

download_with_retry() {
  local url="$1"
  local destination="$2"
  local label="$3"
  local attempt=1

  while (( attempt <= retry_attempts )); do
    echo "[haxe-setup] downloading ${label}, attempt ${attempt}/${retry_attempts}"
    if curl -fL --retry 3 --retry-all-errors --connect-timeout 20 --max-time 240 -o "$destination" "$url"; then
      return 0
    fi
    if (( attempt < retry_attempts )); then
      local delay
      delay="$(compute_delay "$attempt")"
      echo "[haxe-setup] retrying ${label} in ${delay}s"
      sleep "$delay"
    fi
    attempt=$((attempt + 1))
  done

  echo "[haxe-setup] error: failed to download ${label} after ${retry_attempts} attempts" >&2
  return 1
}

extract_single_root() {
  local archive="$1"
  local destination="$2"
  local root

  rm -rf "$destination"
  mkdir -p "$destination"
  tar -xzf "$archive" -C "$destination"

  root="$(find "$destination" -mindepth 1 -maxdepth 1 -type d | sort | head -n 1 || true)"
  if [[ -z "$root" ]]; then
    echo "[haxe-setup] error: archive ${archive} did not extract a top-level directory" >&2
    return 1
  fi
  printf '%s\n' "$root"
}

ensure_executable() {
  local candidate="$1"
  if [[ ! -x "$candidate" ]]; then
    chmod +x "$candidate"
  fi
}

haxe_url="https://github.com/HaxeFoundation/haxe/releases/download/${haxe_version}/haxe-${haxe_version}-linux64.tar.gz"
neko_url="https://github.com/HaxeFoundation/neko/releases/download/${neko_release_tag}/neko-${neko_version}-linux64.tar.gz"

work_dir="$(mktemp -d "${TMPDIR:-/tmp}/haxe-setup.XXXXXX")"
trap 'rm -rf "$work_dir"' EXIT

haxe_archive="${work_dir}/haxe.tar.gz"
neko_archive="${work_dir}/neko.tar.gz"

download_with_retry "$haxe_url" "$haxe_archive" "haxe ${haxe_version} linux64 archive"
download_with_retry "$neko_url" "$neko_archive" "neko ${neko_version} linux64 archive"

tar -tzf "$haxe_archive" >/dev/null
tar -tzf "$neko_archive" >/dev/null

rm -rf "$install_root"
mkdir -p "$install_root"

haxe_root="$(extract_single_root "$haxe_archive" "${install_root}/haxe")"
neko_root="$(extract_single_root "$neko_archive" "${install_root}/neko")"

ensure_executable "${haxe_root}/haxe"
ensure_executable "${haxe_root}/haxelib"
ensure_executable "${neko_root}/neko"

export PATH="${haxe_root}:${neko_root}:${PATH}"
export HAXE_STD_PATH="${haxe_root}/std"
export NEKOPATH="${neko_root}"
export LD_LIBRARY_PATH="${neko_root}${LD_LIBRARY_PATH:+:${LD_LIBRARY_PATH}}"

if [[ -n "${GITHUB_PATH:-}" ]]; then
  printf '%s\n' "${haxe_root}" >> "${GITHUB_PATH}"
  printf '%s\n' "${neko_root}" >> "${GITHUB_PATH}"
fi

if [[ -n "${GITHUB_ENV:-}" ]]; then
  {
    printf 'HAXE_STD_PATH=%s\n' "${HAXE_STD_PATH}"
    printf 'NEKOPATH=%s\n' "${NEKOPATH}"
    printf 'LD_LIBRARY_PATH=%s\n' "${LD_LIBRARY_PATH}"
  } >> "${GITHUB_ENV}"
fi

echo "[haxe-setup] fallback install completed"
echo "[haxe-setup] haxe root: ${haxe_root}"
echo "[haxe-setup] neko root: ${neko_root}"

if ! haxe_version_out="$(haxe -version 2>/tmp/haxe-version.err)"; then
  cat /tmp/haxe-version.err >&2 || true
  echo "[haxe-setup] error: haxe executable did not run after fallback install" >&2
  exit 1
fi

if ! haxelib_version_out="$(haxelib version 2>/tmp/haxelib-version.err)"; then
  cat /tmp/haxelib-version.err >&2 || true
  echo "[haxe-setup] error: haxelib executable did not run after fallback install" >&2
  exit 1
fi

if ! neko_version_out="$(neko -version 2>/tmp/neko-version.err)"; then
  cat /tmp/neko-version.err >&2 || true
  echo "[haxe-setup] error: neko executable did not run after fallback install" >&2
  exit 1
fi

echo "[haxe-setup] haxe version: ${haxe_version_out}"
echo "[haxe-setup] haxelib version: ${haxelib_version_out}"
echo "[haxe-setup] neko version: ${neko_version_out}"
