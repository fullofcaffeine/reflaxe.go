#!/usr/bin/env python3
"""Build deterministic perf warning-history artifacts from harness output."""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any


GO_WARNING_RE = re.compile(
    r"^(?P<case>[^.]+)\.(?P<profile>[^.]+)\.(?P<metric>\S+) ratio "
    r"(?P<drift>[+-]?\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), budget=\+?(?P<threshold>\d+(?:\.\d+)?)%\)$"
)
GO_DELTA_WARNING_RE = re.compile(
    r"^delta\.(?P<case>[^.]+)\.(?P<metric>\S+) ratio "
    r"(?P<drift>[+-]?\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), budget=\+?(?P<threshold>\d+(?:\.\d+)?)%\)$"
)
APP_WARNING_RE = re.compile(
    r"^(?P<app>[^:]+)::(?P<variant>[^:]+)::(?P<profile>[^.]+)\.(?P<metric>\S+) "
    r"(?P<direction>rose|dropped) (?P<drift>\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), budget=(?P<threshold>\d+(?:\.\d+)?)%\)$"
)
APP_DELTA_WARNING_RE = re.compile(
    r"^delta\.(?P<app>[^:]+)::(?P<variant>[^.]+)\.(?P<metric>\S+) "
    r"(?P<direction>rose|dropped) (?P<drift>\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), budget=(?P<threshold>\d+(?:\.\d+)?)%\)$"
)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Build deterministic perf warning-history artifacts")
    parser.add_argument("--harness", required=True, help="Stable harness id, e.g. go-profile or app-profile")
    parser.add_argument("--comparison", required=True, type=Path, help="comparison.json produced by a perf harness")
    parser.add_argument("--out-json", required=True, type=Path, help="Output warning_history.json path")
    parser.add_argument("--out-md", required=True, type=Path, help="Output warning_history.md path")
    return parser.parse_args()


def parse_float(raw: str | None) -> float | None:
    if raw is None or raw == "":
        return None
    return float(raw)


def normalize_warning(harness: str, message: str) -> dict[str, Any]:
    for parser in (parse_go_delta_warning, parse_go_warning, parse_app_delta_warning, parse_app_warning):
        parsed = parser(harness, message)
        if parsed is not None:
            return parsed
    return {
        "harness": harness,
        "kind": "unparsed",
        "app": None,
        "case": None,
        "variant": None,
        "profile": None,
        "metric": "unparsed",
        "thresholdPct": None,
        "driftPct": None,
        "current": None,
        "baseline": None,
        "message": message,
        "groupKey": f"{harness}|unparsed|{message}",
    }


def parse_go_warning(harness: str, message: str) -> dict[str, Any] | None:
    match = GO_WARNING_RE.match(message)
    if match is None:
        return None
    data = match.groupdict()
    return warning_record(
        harness=harness,
        kind="profile",
        app=None,
        case=data["case"],
        variant=None,
        profile=data["profile"],
        metric=data["metric"],
        threshold_pct=parse_float(data["threshold"]),
        drift_pct=parse_float(data["drift"]),
        current=data["current"],
        baseline=data["baseline"],
        message=message,
    )


def parse_go_delta_warning(harness: str, message: str) -> dict[str, Any] | None:
    match = GO_DELTA_WARNING_RE.match(message)
    if match is None:
        return None
    data = match.groupdict()
    return warning_record(
        harness=harness,
        kind="delta",
        app=None,
        case=data["case"],
        variant=None,
        profile="portable_vs_metal",
        metric=data["metric"],
        threshold_pct=parse_float(data["threshold"]),
        drift_pct=parse_float(data["drift"]),
        current=data["current"],
        baseline=data["baseline"],
        message=message,
    )


def parse_app_warning(harness: str, message: str) -> dict[str, Any] | None:
    match = APP_WARNING_RE.match(message)
    if match is None:
        return None
    data = match.groupdict()
    return warning_record(
        harness=harness,
        kind="profile",
        app=data["app"],
        case=None,
        variant=data["variant"],
        profile=data["profile"],
        metric=data["metric"],
        threshold_pct=parse_float(data["threshold"]),
        drift_pct=signed_drift(data["direction"], data["drift"]),
        current=data["current"],
        baseline=data["baseline"],
        message=message,
    )


def parse_app_delta_warning(harness: str, message: str) -> dict[str, Any] | None:
    match = APP_DELTA_WARNING_RE.match(message)
    if match is None:
        return None
    data = match.groupdict()
    return warning_record(
        harness=harness,
        kind="delta",
        app=data["app"],
        case=None,
        variant=data["variant"],
        profile="portable_vs_metal",
        metric=data["metric"],
        threshold_pct=parse_float(data["threshold"]),
        drift_pct=signed_drift(data["direction"], data["drift"]),
        current=data["current"],
        baseline=data["baseline"],
        message=message,
    )


def signed_drift(direction: str, drift: str) -> float:
    value = parse_float(drift) or 0.0
    if direction == "dropped":
        return -value
    return value


def warning_record(
    *,
    harness: str,
    kind: str,
    app: str | None,
    case: str | None,
    variant: str | None,
    profile: str | None,
    metric: str,
    threshold_pct: float | None,
    drift_pct: float | None,
    current: str | None,
    baseline: str | None,
    message: str,
) -> dict[str, Any]:
    group_key = "|".join(
        [
            harness,
            kind,
            app or "",
            case or "",
            variant or "",
            profile or "",
            metric,
            "" if threshold_pct is None else f"{threshold_pct:.6f}",
        ]
    )
    return {
        "harness": harness,
        "kind": kind,
        "app": app,
        "case": case,
        "variant": variant,
        "profile": profile,
        "metric": metric,
        "thresholdPct": threshold_pct,
        "driftPct": drift_pct,
        "current": current,
        "baseline": baseline,
        "message": message,
        "groupKey": group_key,
    }


def build_groups(warnings: list[dict[str, Any]]) -> list[dict[str, Any]]:
    grouped: dict[str, list[dict[str, Any]]] = {}
    for warning in warnings:
        grouped.setdefault(warning["groupKey"], []).append(warning)

    groups: list[dict[str, Any]] = []
    for key in sorted(grouped):
        entries = sorted(grouped[key], key=lambda item: item["message"])
        first = entries[0]
        groups.append(
            {
                "groupKey": key,
                "harness": first["harness"],
                "kind": first["kind"],
                "app": first["app"],
                "case": first["case"],
                "variant": first["variant"],
                "profile": first["profile"],
                "metric": first["metric"],
                "thresholdPct": first["thresholdPct"],
                "count": len(entries),
                "warnings": entries,
            }
        )
    return groups


def render_markdown(report: dict[str, Any]) -> str:
    lines = [
        "# Perf Warning History",
        "",
        f"- Harness: `{report['harness']}`",
        f"- Warning count: `{report['warningCount']}`",
        f"- Group count: `{report['groupCount']}`",
        "",
        "| Harness | Kind | App | Case | Variant | Profile | Metric | Threshold % | Count |",
        "| --- | --- | --- | --- | --- | --- | --- | ---: | ---: |",
    ]
    for group in report["groups"]:
        lines.append(
            "| "
            + " | ".join(
                [
                    md(group["harness"]),
                    md(group["kind"]),
                    md(group["app"]),
                    md(group["case"]),
                    md(group["variant"]),
                    md(group["profile"]),
                    md(group["metric"]),
                    "" if group["thresholdPct"] is None else f"{group['thresholdPct']:.2f}",
                    str(group["count"]),
                ]
            )
            + " |"
        )
    if not report["groups"]:
        lines.append("| - | - | - | - | - | - | - | - | 0 |")
    lines.append("")
    return "\n".join(lines) + "\n"


def md(value: Any) -> str:
    if value is None or value == "":
        return "-"
    return str(value).replace("|", "\\|")


def main() -> int:
    args = parse_args()
    comparison = json.loads(args.comparison.read_text(encoding="utf-8"))
    raw_warnings = comparison.get("warnings", [])
    if not isinstance(raw_warnings, list):
        raise TypeError("comparison warnings field must be a list")

    warnings = [normalize_warning(args.harness, str(message)) for message in raw_warnings]
    warnings.sort(key=lambda item: (item["groupKey"], item["message"]))
    groups = build_groups(warnings)
    report = {
        "schemaVersion": 1,
        "harness": args.harness,
        "comparisonPath": str(args.comparison),
        "warningCount": len(warnings),
        "groupCount": len(groups),
        "groups": groups,
    }

    args.out_json.parent.mkdir(parents=True, exist_ok=True)
    args.out_json.write_text(json.dumps(report, indent=2, sort_keys=True) + "\n", encoding="utf-8")
    args.out_md.parent.mkdir(parents=True, exist_ok=True)
    args.out_md.write_text(render_markdown(report), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
