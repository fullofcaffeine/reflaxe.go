#!/usr/bin/env node

const cp = require("child_process");
const fs = require("fs");
const nodePath = require("path");

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

function sameStringArray(left, right) {
  return JSON.stringify(left) === JSON.stringify(right);
}

function isRepoRelative(path) {
  return typeof path === "string"
    && path.length > 0
    && !nodePath.posix.isAbsolute(path)
    && !nodePath.win32.isAbsolute(path)
    && !path.split("/").includes("..");
}

const ledgerPath = "docs/stdlib-provenance-ledger.json";
const approvedGovernedStdLayout = [
  "std/*.hx",
  "std/*.cross.hx",
  "std/haxe/**",
  "std/sys/**",
  "std/go/**",
  "std/hxrt/**",
  "std/_std/**",
];
const stdPathAllowlist = new Set(["std/AGENTS.md"]);
const allowedOwnershipClasses = new Set([
  "upstream_std_override",
  "staged_support",
  "hxrt_binding",
  "public_go_facade",
  "obsolete",
  "intentional_boundary_fixture",
]);

function hasSupportedStdExtension(path) {
  return path.endsWith(".hx") || path.endsWith(".cross.hx");
}

function isGovernedStdSourcePath(path) {
  if (!path.startsWith("std/") || !hasSupportedStdExtension(path)) {
    return false;
  }

  if (
    path.startsWith("std/haxe/")
    || path.startsWith("std/sys/")
    || path.startsWith("std/go/")
    || path.startsWith("std/hxrt/")
    || path.startsWith("std/_std/")
  ) {
    return true;
  }

  return path.split("/").length === 2;
}

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

if (ledger.schemaVersion !== 2) {
  fail(`${ledgerPath} schemaVersion must be 2`);
}

if (!ledger.baselineUpstream || typeof ledger.baselineUpstream !== "object") {
  fail(`${ledgerPath} must define baselineUpstream metadata`);
}

const migrationContract = ledger.migrationContract;
if (migrationContract == null || typeof migrationContract !== "object") {
  fail(`${ledgerPath} must define migrationContract metadata`);
}

const canonicalOverrideRoot = migrationContract && migrationContract.canonicalOverrideRoot;
const targetSupportRoot = migrationContract && migrationContract.targetSupportRoot;
if (canonicalOverrideRoot !== "std/go/_std") {
  fail(`${ledgerPath} migrationContract.canonicalOverrideRoot must be std/go/_std`);
}
if (targetSupportRoot !== "std/hxrt") {
  fail(`${ledgerPath} migrationContract.targetSupportRoot must be std/hxrt`);
}

const ownershipDefinitions = migrationContract && migrationContract.ownershipClasses;
if (ownershipDefinitions == null || typeof ownershipDefinitions !== "object" || Array.isArray(ownershipDefinitions)) {
  fail(`${ledgerPath} migrationContract must define ownershipClasses`);
} else {
  const actualClasses = Object.keys(ownershipDefinitions).sort();
  const expectedClasses = [...allowedOwnershipClasses].sort();
  if (!sameStringArray(actualClasses, expectedClasses)) {
    fail(`${ledgerPath} ownershipClasses must be exactly ${expectedClasses.join(", ")}`);
  }
  for (const ownershipClass of expectedClasses) {
    if (typeof ownershipDefinitions[ownershipClass] !== "string" || ownershipDefinitions[ownershipClass].length === 0) {
      fail(`${ledgerPath} ownershipClasses.${ownershipClass} must explain the ownership boundary`);
    }
  }
}

const migrationBeads = migrationContract && migrationContract.migrationBeads;
if (migrationBeads == null || typeof migrationBeads !== "object" || Array.isArray(migrationBeads)) {
  fail(`${ledgerPath} migrationContract must define migrationBeads`);
} else {
  for (const ownershipClass of allowedOwnershipClasses) {
    const bead = migrationBeads[ownershipClass];
    if (typeof bead !== "string" || !/^haxe_go-[a-z0-9.-]+$/.test(bead)) {
      fail(`${ledgerPath} migrationBeads.${ownershipClass} must name a Bead`);
    }
  }
}

const audits = migrationContract && migrationContract.compilerShimAudit;
const auditByGroup = new Map();
if (!Array.isArray(audits)) {
  fail(`${ledgerPath} migrationContract must contain compilerShimAudit`);
} else {
  for (const audit of audits) {
    if (audit == null || typeof audit !== "object") {
      fail(`${ledgerPath} compilerShimAudit contains a non-object entry`);
      continue;
    }
    const group = audit.group;
    if (typeof group !== "string" || group.length === 0) {
      fail(`${ledgerPath} compilerShimAudit entry is missing group`);
      continue;
    }
    if (auditByGroup.has(group)) {
      fail(`${ledgerPath} compilerShimAudit contains duplicate group: ${group}`);
    }
    auditByGroup.set(group, audit);

    if (typeof audit.decision !== "string" || audit.decision.length === 0) {
      fail(`${ledgerPath} compilerShimAudit.${group} is missing decision`);
    }
    if (!Array.isArray(audit.sourcePaths)) {
      fail(`${ledgerPath} compilerShimAudit.${group} must define sourcePaths`);
    } else if (new Set(audit.sourcePaths).size !== audit.sourcePaths.length) {
      fail(`${ledgerPath} compilerShimAudit.${group} contains duplicate sourcePaths`);
    }
    if (!Array.isArray(audit.references) || audit.references.length === 0) {
      fail(`${ledgerPath} compilerShimAudit.${group} must define references`);
    } else {
      for (const reference of audit.references) {
        if (!isRepoRelative(reference) || !fs.existsSync(reference) || !fs.statSync(reference).isFile()) {
          fail(`${ledgerPath} compilerShimAudit.${group} has missing or unsafe reference: ${reference}`);
        }
      }
    }
    if (typeof audit.notes !== "string" || audit.notes.length === 0) {
      fail(`${ledgerPath} compilerShimAudit.${group} must explain the retained compiler seam`);
    }
  }
}

if (!Array.isArray(ledger.entries)) {
  fail(`${ledgerPath} must contain an entries array`);
}

const trackedStdSources = gitTrackedUnder("std").filter((path) => {
  if (stdPathAllowlist.has(path)) {
    return false;
  }
  return isGovernedStdSourcePath(path);
});
const trackedSet = new Set(trackedStdSources);

const ledgerPaths = [];
const destinations = [];
const pathsByShimGroup = new Map([...auditByGroup.keys()].map((group) => [group, []]));
for (const entry of ledger.entries) {
  if (entry == null || typeof entry !== "object") {
    fail(`${ledgerPath} contains a non-object entry`);
    continue;
  }

  const path = entry.path;
  const provenanceKind = entry.provenanceKind;
  const upstreamOraclePath = entry.upstreamOraclePath;
  const ownershipClass = entry.ownershipClass;
  const destination = entry.destination;
  const migrationBead = entry.migrationBead;
  const compilerShimGroups = entry.compilerShimGroups;

  if (typeof path !== "string" || path.length === 0) {
    fail(`${ledgerPath} entry is missing path`);
    continue;
  }

  if (!isGovernedStdSourcePath(path)) {
    fail(`${ledgerPath} entry path must stay within the approved staged layout (${approvedGovernedStdLayout.join(", ")}): ${path}`);
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

  if (!allowedOwnershipClasses.has(ownershipClass)) {
    fail(`${ledgerPath} entry has unsupported ownershipClass for ${path}: ${ownershipClass}`);
  }

  if (!isRepoRelative(destination) || !destination.endsWith(".hx") || destination.endsWith(".cross.hx")) {
    fail(`${ledgerPath} entry destination must be a repo-relative ordinary .hx path for ${path}: ${destination}`);
  } else {
    if (destinations.includes(destination)) {
      fail(`${ledgerPath} contains duplicate migration destination: ${destination}`);
    }
    destinations.push(destination);

    if (ownershipClass === "upstream_std_override" && !destination.startsWith(`${canonicalOverrideRoot}/`)) {
      fail(`${ledgerPath} upstream_std_override must migrate under ${canonicalOverrideRoot}: ${path}`);
    }
    if (
      ownershipClass === "staged_support"
      && (!destination.startsWith("std/")
        || destination.startsWith(`${canonicalOverrideRoot}/`)
        || destination.startsWith(`${targetSupportRoot}/`))
    ) {
      fail(`${ledgerPath} staged_support must remain ordinary std support outside _std and hxrt: ${path}`);
    }
    if (ownershipClass === "hxrt_binding" && !destination.startsWith(`${targetSupportRoot}/`)) {
      fail(`${ledgerPath} hxrt_binding must migrate under ${targetSupportRoot}: ${path}`);
    }
    if (
      ownershipClass === "public_go_facade"
      && (!destination.startsWith("std/go/") || destination.startsWith(`${canonicalOverrideRoot}/`))
    ) {
      fail(`${ledgerPath} public_go_facade must remain under std/go outside _std: ${path}`);
    }
  }

  const expectedMigrationBead = migrationBeads && migrationBeads[ownershipClass];
  if (migrationBead !== expectedMigrationBead) {
    fail(`${ledgerPath} entry migrationBead must be ${expectedMigrationBead} for ${path}`);
  }

  if (!Array.isArray(compilerShimGroups)) {
    fail(`${ledgerPath} entry compilerShimGroups must be an array for ${path}`);
  } else {
    const normalizedGroups = [...new Set(compilerShimGroups)].sort();
    if (!sameStringArray(compilerShimGroups, normalizedGroups)) {
      fail(`${ledgerPath} entry compilerShimGroups must be sorted and unique for ${path}`);
    }
    for (const group of compilerShimGroups) {
      if (!auditByGroup.has(group)) {
        fail(`${ledgerPath} entry references unaudited compiler shim group ${group}: ${path}`);
        continue;
      }
      pathsByShimGroup.get(group).push(path);
    }
  }
}

const ledgerSet = new Set(ledgerPaths);
const missingCoverage = trackedStdSources.filter((path) => !ledgerSet.has(path));
const staleCoverage = ledgerPaths.filter((path) => !trackedSet.has(path));

if (missingCoverage.length > 0) {
  fail(`stdlib migration ledger missing tracked std sources:\n${summarize(missingCoverage, 20)}`);
}

if (staleCoverage.length > 0) {
  fail(`stdlib migration ledger references non-tracked std sources:\n${summarize(staleCoverage, 20)}`);
}

for (const [group, audit] of auditByGroup.entries()) {
  if (!Array.isArray(audit.sourcePaths)) {
    continue;
  }
  const declared = [...audit.sourcePaths].sort();
  const assigned = [...(pathsByShimGroup.get(group) || [])].sort();
  if (!sameStringArray(declared, assigned)) {
    fail(`${ledgerPath} compilerShimAudit.${group} sourcePaths do not match per-entry assignments`);
  }
  for (const sourcePath of declared) {
    if (!ledgerSet.has(sourcePath)) {
      fail(`${ledgerPath} compilerShimAudit.${group} references an unknown source: ${sourcePath}`);
    }
  }
}

const ambiguities = migrationContract && migrationContract.ambiguities;
if (!Array.isArray(ambiguities)) {
  fail(`${ledgerPath} migrationContract must define ambiguities`);
} else {
  const seenAmbiguities = new Set();
  for (const ambiguity of ambiguities) {
    const path = ambiguity && ambiguity.path;
    const followUpBead = ambiguity && ambiguity.followUpBead;
    if (!ledgerSet.has(path)) {
      fail(`${ledgerPath} ambiguity references an unknown source: ${path}`);
    }
    if (seenAmbiguities.has(path)) {
      fail(`${ledgerPath} contains duplicate ambiguity entry: ${path}`);
    }
    seenAmbiguities.add(path);
    if (typeof followUpBead !== "string" || !/^haxe_go-[a-z0-9.-]+$/.test(followUpBead)) {
      fail(`${ledgerPath} ambiguity must name a follow-up Bead for ${path}`);
    }
  }
}

if (process.exitCode) {
  process.exit(process.exitCode);
}

const ownershipCounts = new Map();
for (const entry of ledger.entries) {
  ownershipCounts.set(entry.ownershipClass, (ownershipCounts.get(entry.ownershipClass) || 0) + 1);
}
const summary = [...ownershipCounts.entries()]
  .sort(([left], [right]) => left.localeCompare(right))
  .map(([ownership, count]) => `${ownership}=${count}`)
  .join(", ");
console.log(`[ci:guards] OK: stdlib migration ledger covers ${trackedStdSources.length} tracked sources (${summary})`);
