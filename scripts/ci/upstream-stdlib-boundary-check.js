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

const approvedStdOverrideRoots = ["std/go/", "std/_std/"];
const stdPathAllowlist = new Set(["std/AGENTS.md"]);

const trackedVendor = gitTrackedUnder("vendor/haxe");
if (trackedVendor.length > 0) {
  fail(
    "tracked files under vendor/haxe are not allowed. Keep upstream vendor roots untracked and sync required overrides into approved std roots. Found:\n"
      + summarize(trackedVendor, 20)
  );
}

const trackedStd = gitTrackedUnder("std");
const disallowedStdFiles = [];
for (const path of trackedStd) {
  if (stdPathAllowlist.has(path)) {
    continue;
  }

  const inApprovedRoot = approvedStdOverrideRoots.some((root) => path.startsWith(root));
  if (!inApprovedRoot) {
    disallowedStdFiles.push(path);
    continue;
  }

  if (!path.endsWith(".hx") && !path.endsWith(".cross.hx")) {
    disallowedStdFiles.push(path);
  }
}

if (disallowedStdFiles.length > 0) {
  fail(
    "stdlib override roots may only contain .hx/.cross.hx files (plus std/AGENTS.md) under approved roots ("
      + approvedStdOverrideRoots.join(", ")
      + "). Found:\n"
      + summarize(disallowedStdFiles, 20)
  );
}

console.log(
  `[ci:guards] OK: upstream stdlib boundary (vendor/haxe untracked; approved override roots: ${approvedStdOverrideRoots.join(
    ", "
  )})`
);

if (process.exitCode) {
  process.exit(process.exitCode);
}
