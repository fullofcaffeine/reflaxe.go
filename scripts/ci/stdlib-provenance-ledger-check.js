#!/usr/bin/env node

const cp = require("child_process");
const fs = require("fs");

function fail(message) {
  console.error(`[ci:guards] ERROR: ${message}`);
  process.exitCode = 1;
}

function gitTrackedUnder(path) {
  try {
    const out = cp.execFileSync("git", ["ls-files", "-z", "--", path], {
      encoding: "utf8",
    });
    return out
      .split("\0")
      .filter(Boolean)
      .filter((trackedPath) => fs.existsSync(trackedPath));
  } catch (_) {
    return [];
  }
}

function summarize(paths, limit) {
  const slice = paths.slice(0, limit).map((path) => `- ${path}`);
  const suffix = paths.length > limit ? `\n- ... (${paths.length - limit} more)` : "";
  return `${slice.join("\n")}${suffix}`;
}

const ledgerPath = "docs/stdlib-provenance-ledger.json";
const approvedStdOverrideRoots = ["std/go/", "std/_std/"];
const stdPathAllowlist = new Set(["std/AGENTS.md"]);

if (!fs.existsSync(ledgerPath)) {
  fail(`missing stdlib provenance ledger: ${ledgerPath}`);
  process.exit(process.exitCode);
}

let ledger = null;
try {
  ledger = JSON.parse(fs.readFileSync(ledgerPath, "utf8"));
} catch (error) {
  fail(`invalid JSON in ${ledgerPath}: ${error}`);
  process.exit(process.exitCode);
}

if (ledger == null || typeof ledger !== "object") {
  fail(`${ledgerPath} must contain a JSON object`);
}

if (ledger.schemaVersion !== 1) {
  fail(`${ledgerPath} schemaVersion must be 1`);
}

if (!ledger.baselineUpstream || typeof ledger.baselineUpstream !== "object") {
  fail(`${ledgerPath} must define baselineUpstream metadata`);
}

if (!Array.isArray(ledger.entries)) {
  fail(`${ledgerPath} must contain an entries array`);
}

const trackedStdOverrideFiles = gitTrackedUnder("std").filter((path) => {
  if (stdPathAllowlist.has(path)) {
    return false;
  }
  const inApprovedRoot = approvedStdOverrideRoots.some((root) => path.startsWith(root));
  const supportedExtension = path.endsWith(".hx") || path.endsWith(".cross.hx");
  return inApprovedRoot && supportedExtension;
});
const trackedSet = new Set(trackedStdOverrideFiles);

const ledgerPaths = [];
for (const entry of ledger.entries) {
  if (entry == null || typeof entry !== "object") {
    fail(`${ledgerPath} contains a non-object entry`);
    continue;
  }

  const path = entry.path;
  const provenanceKind = entry.provenanceKind;
  const upstreamOraclePath = entry.upstreamOraclePath;

  if (typeof path !== "string" || path.length === 0) {
    fail(`${ledgerPath} entry is missing path`);
    continue;
  }

  const inApprovedRoot = approvedStdOverrideRoots.some((root) => path.startsWith(root));
  if (!inApprovedRoot) {
    fail(`${ledgerPath} entry path must stay under approved roots (${approvedStdOverrideRoots.join(", ")}): ${path}`);
  }

  if (!path.endsWith(".hx") && !path.endsWith(".cross.hx")) {
    fail(`${ledgerPath} entry path must target a .hx/.cross.hx file: ${path}`);
  }

  if (ledgerPaths.includes(path)) {
    fail(`${ledgerPath} contains duplicate path entry: ${path}`);
  }
  ledgerPaths.push(path);

  if (provenanceKind !== "upstream_std_sync" && provenanceKind !== "repo_authored_override") {
    fail(`${ledgerPath} entry provenanceKind must be upstream_std_sync or repo_authored_override for ${path}`);
  }

  if (provenanceKind === "upstream_std_sync") {
    if (typeof upstreamOraclePath !== "string" || upstreamOraclePath.length === 0) {
      fail(`${ledgerPath} entry is missing upstreamOraclePath for ${path}`);
    } else if (!upstreamOraclePath.startsWith("vendor/haxe/std/")) {
      fail(`${ledgerPath} entry upstreamOraclePath must point to vendor/haxe/std/** for ${path}: ${upstreamOraclePath}`);
    }
  }
}

const ledgerSet = new Set(ledgerPaths);
const missingCoverage = trackedStdOverrideFiles.filter((path) => !ledgerSet.has(path));
const staleCoverage = ledgerPaths.filter((path) => !trackedSet.has(path));

if (missingCoverage.length > 0) {
  fail(`stdlib provenance ledger missing tracked std override files:\n${summarize(missingCoverage, 20)}`);
}

if (staleCoverage.length > 0) {
  fail(`stdlib provenance ledger references non-tracked std override files:\n${summarize(staleCoverage, 20)}`);
}

if (process.exitCode) {
  process.exit(process.exitCode);
}

console.log(`[ci:guards] OK: stdlib provenance ledger covers ${trackedStdOverrideFiles.length} tracked std override files`);
