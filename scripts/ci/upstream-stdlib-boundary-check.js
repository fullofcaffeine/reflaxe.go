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

const approvedStdOverrideLayout = [
  "std/*.hx",
  "std/haxe/**",
  "std/sys/**",
  "std/go/**",
  "std/hxrt/**",
  "std/reflaxe/go/internal/**",
];
const stdPathAllowlist = new Set(["std/AGENTS.md"]);

function hasSupportedStdExtension(path) {
  return path.endsWith(".hx") && !path.endsWith(".cross.hx");
}

function isApprovedStdOverridePath(path) {
  if (!path.startsWith("std/") || !hasSupportedStdExtension(path)) {
    return false;
  }

  if (
    path.startsWith("std/haxe/")
    || path.startsWith("std/sys/")
    || path.startsWith("std/go/")
    || path.startsWith("std/hxrt/")
    || path.startsWith("std/reflaxe/go/internal/")
  ) {
    return true;
  }

  return path.split("/").length === 2;
}

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

  if (!isApprovedStdOverridePath(path)) {
    disallowedStdFiles.push(path);
  }
}

if (disallowedStdFiles.length > 0) {
  fail(
    "stdlib override policy allows only ordinary .hx files (plus std/AGENTS.md) in the approved staged layout ("
      + approvedStdOverrideLayout.join(", ")
      + "). Found:\n"
      + summarize(disallowedStdFiles, 20)
  );
}

console.log(
  `[ci:guards] OK: upstream stdlib boundary (vendor/haxe untracked; approved staged layout: ${approvedStdOverrideLayout.join(
    ", "
  )})`
);

if (process.exitCode) {
  process.exit(process.exitCode);
}
