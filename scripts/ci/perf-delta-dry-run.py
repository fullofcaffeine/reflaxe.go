#!/usr/bin/env python3
"""Extract non-blocking portable-vs-metal delta hard-gate candidates."""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any


GO_DELTA_HARD_RE = re.compile(
    r"^delta\.(?P<case>[^.]+)\.(?P<metric>\S+) ratio "
    r"(?P<drift>[+-]?\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), budget=\+?(?P<threshold>\d+(?:\.\d+)?)%\)$"
)
APP_DELTA_HARD_RE = re.compile(
    r"^delta\.(?P<app>[^:]+)::(?P<variant>[^.]+)\.(?P<metric>\S+) "
    r"(?P<direction>rose|dropped) (?P<drift>\d+(?:\.\d+)?)% "
    r"\(current=(?P<current>[^,]+), baseline=(?P<baseline>[^,]+), delta budget=(?P<threshold>\d+(?:\.\d+)?)%\)$"
)


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Build non-blocking delta hard-gate dry-run artifact")
    parser.add_argument("--harness", required=True, help="Stable harness id, e.g. go-profile or app-profile")
    parser.add_argument("--comparison", required=True, type=Path, help="comparison.json produced by a perf harness")
    parser.add_argument("--out-json", required=True, type=Path, help="Output delta_hard_gate_dry_run.json path")
    parser.add_argument("--out-md", required=True, type=Path, help="Output delta_hard_gate_dry_run.md path")
    return parser.parse_args()


def parse_candidate(harness: str, message: str) -> dict[str, Any] | None:
    if not message.startswith("delta."):
        return None

    go_match = GO_DELTA_HARD_RE.match(message)
    if go_match is not None:
        data = go_match.groupdict()
        return candidate_record(
            harness=harness,
            app=None,
            case=data["case"],
            variant=None,
            metric=data["metric"],
            drift_pct=float(data["drift"]),
            current=data["current"],
            baseline=data["baseline"],
            threshold_pct=float(data["threshold"]),
            message=message,
        )

    app_match = APP_DELTA_HARD_RE.match(message)
    if app_match is not None:
        data = app_match.groupdict()
        drift = float(data["drift"])
        if data["direction"] == "dropped":
            drift = -drift
        return candidate_record(
            harness=harness,
            app=data["app"],
            case=None,
            variant=data["variant"],
            metric=data["metric"],
            drift_pct=drift,
            current=data["current"],
            baseline=data["baseline"],
            threshold_pct=float(data["threshold"]),
            message=message,
        )

    return candidate_record(
        harness=harness,
        app=None,
        case=None,
        variant=None,
        metric="unparsed",
        drift_pct=None,
        current=None,
        baseline=None,
        threshold_pct=None,
        message=message,
    )


def candidate_record(
    *,
    harness: str,
    app: str | None,
    case: str | None,
    variant: str | None,
    metric: str,
    drift_pct: float | None,
    current: str | None,
    baseline: str | None,
    threshold_pct: float | None,
    message: str,
) -> dict[str, Any]:
    key = "|".join(
        [
            harness,
            app or "",
            case or "",
            variant or "",
            metric,
            "" if threshold_pct is None else f"{threshold_pct:.6f}",
        ]
    )
    return {
        "groupKey": key,
        "harness": harness,
        "app": app,
        "case": case,
        "variant": variant,
        "profile": "portable_vs_metal",
        "metric": metric,
        "driftPct": drift_pct,
        "current": current,
        "baseline": baseline,
        "thresholdPct": threshold_pct,
        "message": message,
    }


def render_markdown(report: dict[str, Any]) -> str:
    lines = [
        "# Delta Hard-Gate Dry Run",
        "",
        f"- Harness: `{report['harness']}`",
        f"- Would fail if enforced: `{report['wouldFailIfEnforced']}`",
        f"- Candidate count: `{report['candidateCount']}`",
        "",
        "| Harness | App | Case | Variant | Profile | Metric | Drift % | Threshold % |",
        "| --- | --- | --- | --- | --- | --- | ---: | ---: |",
    ]
    for candidate in report["candidates"]:
        lines.append(
            "| "
            + " | ".join(
                [
                    md(candidate["harness"]),
                    md(candidate["app"]),
                    md(candidate["case"]),
                    md(candidate["variant"]),
                    md(candidate["profile"]),
                    md(candidate["metric"]),
                    "" if candidate["driftPct"] is None else f"{candidate['driftPct']:.2f}",
                    "" if candidate["thresholdPct"] is None else f"{candidate['thresholdPct']:.2f}",
                ]
            )
            + " |"
        )
    if not report["candidates"]:
        lines.append("| - | - | - | - | - | - | - | - |")
    lines.append("")
    return "\n".join(lines) + "\n"


def md(value: Any) -> str:
    if value is None or value == "":
        return "-"
    return str(value).replace("|", "\\|")


def main() -> int:
    args = parse_args()
    comparison = json.loads(args.comparison.read_text(encoding="utf-8"))
    hard_failures = comparison.get("hardFailures", [])
    if not isinstance(hard_failures, list):
        raise TypeError("comparison hardFailures field must be a list")

    candidates = [
        candidate
        for message in hard_failures
        if (candidate := parse_candidate(args.harness, str(message))) is not None
    ]
    candidates.sort(key=lambda item: (item["groupKey"], item["message"]))
    report = {
        "schemaVersion": 1,
        "harness": args.harness,
        "mode": "non_blocking_delta_hard_gate_dry_run",
        "comparisonPath": str(args.comparison),
        "enforceDeltaBudget": bool(comparison.get("enforceDeltaBudget", False)),
        "wouldFailIfEnforced": len(candidates) > 0,
        "candidateCount": len(candidates),
        "candidates": candidates,
    }

    args.out_json.parent.mkdir(parents=True, exist_ok=True)
    args.out_json.write_text(json.dumps(report, indent=2, sort_keys=True) + "\n", encoding="utf-8")
    args.out_md.parent.mkdir(parents=True, exist_ok=True)
    args.out_md.write_text(render_markdown(report), encoding="utf-8")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
