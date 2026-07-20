import { analyzeCommits as analyzeConventionalCommits } from "@semantic-release/commit-analyzer";
import { generateNotes as generateConventionalNotes } from "@semantic-release/release-notes-generator";
import { execFileSync } from "node:child_process";
import fs from "node:fs";
import path from "node:path";
import { isDeepStrictEqual } from "node:util";
import semver from "semver";


const DEFAULT_POLICY_PATH = "release/policy.json";
const COMMIT_HEADER_PATTERN = /^(\w*)(?:\((.*)\))?(!)?: (.*)$/;

/**
 * A fail-closed release-policy error. A stable code lets CI distinguish a
 * version-policy rejection from parser, GitHub, or network failures.
 */
export class ReleasePolicyError extends Error {
  constructor(message) {
    super(message);
    this.name = "ReleasePolicyError";
    this.code = "EHAXEGORELEASEPOLICY";
  }
}


function fail(message) {
  throw new ReleasePolicyError(message);
}


function requireObject(value, label) {
  if (value === null || typeof value !== "object" || Array.isArray(value)) {
    fail(label + " must be an object");
  }
  return value;
}


function requireString(value, label) {
  if (typeof value !== "string" || value.trim().length === 0) {
    fail(label + " must be a non-empty string");
  }
  return value;
}


function requireExactKeys(value, expected, label) {
  const actual = Object.keys(value).sort();
  const wanted = [...expected].sort();
  if (actual.join(",") !== wanted.join(",")) {
    fail(label + " must contain exactly: " + wanted.join(", "));
  }
}


function validateApproval(approval, label) {
  if (approval === null) {
    return;
  }
  requireObject(approval, label);
  requireString(approval.approvedBy, label + ".approvedBy");
  requireString(approval.record, label + ".record");
  requireString(approval.date, label + ".date");
  if (!/^\d{4}-\d{2}-\d{2}$/.test(approval.date)) {
    fail(label + ".date must use YYYY-MM-DD");
  }
  const parsedDate = new Date(approval.date + "T00:00:00Z");
  if (
    Number.isNaN(parsedDate.getTime())
    || parsedDate.toISOString().slice(0, 10) !== approval.date
    || approval.date > new Date().toISOString().slice(0, 10)
  ) {
    fail(label + ".date must be a real, non-future date");
  }
  if (typeof approval.sourceCommit !== "string" || !/^[0-9a-f]{40}$/.test(approval.sourceCommit)) {
    fail(label + ".sourceCommit must be a full lowercase Git commit");
  }
}


/**
 * Validate the small release-line checklist. The file records why a stable
 * line is blocked or approved; Git tags remain the version authority.
 */
export function validateReleasePolicy(value) {
  const policy = requireObject(value, "release policy");
  requireExactKeys(
    policy,
    ["schemaVersion", "normalChannel", "deprecationPolicy", "experimentalPolicy", "releaseLines"],
    "release policy",
  );
  if (policy.schemaVersion !== 1) {
    fail("release/policy.json schemaVersion must be 1");
  }

  const channel = requireObject(policy.normalChannel, "normalChannel");
  requireExactKeys(
    channel,
    ["branch", "tagFormat", "prerelease", "productMaturityAuthority"],
    "normalChannel",
  );
  if (
    channel.branch !== "master"
    || channel.tagFormat !== "v${version}"
    || channel.prerelease !== false
    || channel.productMaturityAuthority !== "docs/compatibility-support-manifest.json"
  ) {
    fail("normalChannel must keep master, canonical v<SemVer> tags, no prerelease, and the compatibility manifest as its maturity authority");
  }

  const deprecation = requireObject(policy.deprecationPolicy, "deprecationPolicy");
  requireExactKeys(
    deprecation,
    ["majorZero", "stable", "emergencyExceptionRequires"],
    "deprecationPolicy",
  );
  const majorZeroDeprecation = requireObject(deprecation.majorZero, "deprecationPolicy.majorZero");
  requireExactKeys(
    majorZeroDeprecation,
    ["noticeRelease", "minimumFunctionalMinorReleasesAfterNotice", "earliestRemovalMinorOffset"],
    "deprecationPolicy.majorZero",
  );
  if (
    majorZeroDeprecation.noticeRelease !== "minor"
    || majorZeroDeprecation.minimumFunctionalMinorReleasesAfterNotice !== 1
    || majorZeroDeprecation.earliestRemovalMinorOffset !== 2
  ) {
    fail("major-zero deprecation must use notice minor N, functional N+1, and earliest removal N+2");
  }
  const stableDeprecation = requireObject(deprecation.stable, "deprecationPolicy.stable");
  requireExactKeys(
    stableDeprecation,
    ["noticeRelease", "removalRelease"],
    "deprecationPolicy.stable",
  );
  if (stableDeprecation.noticeRelease !== "minor" || stableDeprecation.removalRelease !== "next-major") {
    fail("stable deprecation must preserve functionality until the next major");
  }
  const emergency = deprecation.emergencyExceptionRequires;
  if (
    !Array.isArray(emergency)
    || emergency.join(",") !== "review-record,release-note,migration-guidance"
  ) {
    fail("emergency deprecation exceptions must require review, release notes, and migration guidance");
  }

  const experimental = requireObject(policy.experimentalPolicy, "experimentalPolicy");
  requireExactKeys(
    experimental,
    ["compatibilityPromise", "minimumChangeRelease", "stableMinorException"],
    "experimentalPolicy",
  );
  if (
    experimental.compatibilityPromise !== "excluded-unless-admitted"
    || experimental.minimumChangeRelease !== "minor"
    || experimental.stableMinorException !== "requires-explicit-surface-proof"
  ) {
    fail("experimental policy must stay excluded unless admitted and fail closed without surface proof");
  }

  const lines = requireObject(policy.releaseLines, "releaseLines");
  const majorKeys = Object.keys(lines);
  if (majorKeys.length < 2) {
    fail("releaseLines must include major 0 and the stable major 1 candidate");
  }
  for (let index = 0; index < majorKeys.length; index += 1) {
    if (majorKeys[index] !== String(index)) {
      fail("releaseLines must be the contiguous sequence 0..N");
    }
  }

  const initial = requireObject(lines["0"], "releaseLines.0");
  requireExactKeys(initial, ["stage", "breakingBump"], "releaseLines.0");
  if (initial.stage !== "initial-development" || initial.breakingBump !== "minor") {
    fail("releaseLines.0 must be initial-development with breakingBump minor");
  }
  if (Object.prototype.hasOwnProperty.call(initial, "approval")) {
    fail("releaseLines.0 must not define stable approval");
  }

  let foundPendingStableLine = false;
  for (let major = 1; major < majorKeys.length; major += 1) {
    const label = "releaseLines." + major;
    const line = requireObject(lines[String(major)], label);
    requireExactKeys(line, ["stage", "requirements", "approval"], label);
    if (line.stage !== "stable") {
      fail(label + ".stage must be stable");
    }
    if (!Array.isArray(line.requirements) || line.requirements.length === 0) {
      fail(label + ".requirements must be a non-empty array");
    }
    const ids = new Set();
    const pending = [];
    for (let index = 0; index < line.requirements.length; index += 1) {
      const requirementLabel = label + ".requirements[" + index + "]";
      const requirement = requireObject(line.requirements[index], requirementLabel);
      requireExactKeys(requirement, ["id", "status", "record"], requirementLabel);
      const id = requireString(requirement.id, requirementLabel + ".id");
      if (!/^[a-z0-9]+(?:-[a-z0-9]+)*$/.test(id) || ids.has(id)) {
        fail(requirementLabel + ".id must be unique lowercase kebab-case");
      }
      ids.add(id);
      if (requirement.status !== "pending" && requirement.status !== "complete") {
        fail(requirementLabel + ".status must be pending or complete");
      }
      if (requirement.status === "pending") {
        pending.push(id);
      }
      requireString(requirement.record, requirementLabel + ".record");
    }
    if (!Object.prototype.hasOwnProperty.call(line, "approval")) {
      fail(label + ".approval must be present and null until approved");
    }
    if (line.approval !== null) {
      requireExactKeys(
        requireObject(line.approval, label + ".approval"),
        ["approvedBy", "record", "date", "sourceCommit"],
        label + ".approval",
      );
    }
    validateApproval(line.approval, label + ".approval");
    if (line.approval !== null && pending.length > 0) {
      fail(label + " cannot be approved while requirements are pending: " + pending.join(", "));
    }
    if (line.approval === null) {
      foundPendingStableLine = true;
    } else if (foundPendingStableLine) {
      fail("stable-major approvals must be contiguous from major 1");
    }
  }

  return policy;
}


/**
 * Bind one pending major transition to real source history. The approval must be
 * committed after the reviewed source, so only policy/Beads metadata may
 * differ between that source and the release candidate.
 */
export function validateApprovalSourceCommit(policy, major, context = {}) {
  const cwd = context.cwd ?? process.cwd();
  const line = policy.releaseLines[String(major)];
  if (line?.approval === null || line?.approval === undefined) {
    fail("stable major " + major + " has no approval to validate");
  }
  const sourceCommit = line.approval.sourceCommit;
  try {
    execFileSync("git", ["cat-file", "-e", sourceCommit + "^{commit}"], {
      cwd,
      stdio: "ignore",
    });
  } catch {
    fail(
      "releaseLines."
        + major
        + ".approval.sourceCommit "
        + sourceCommit
        + " does not exist in the release repository",
    );
  }

  let headCommit;
  try {
    headCommit = execFileSync("git", ["rev-parse", "HEAD"], {
      cwd,
      encoding: "utf8",
      stdio: ["ignore", "pipe", "ignore"],
    }).trim();
  } catch {
    fail("cannot resolve the release repository HEAD for stable approval");
  }

  try {
    execFileSync("git", ["merge-base", "--is-ancestor", sourceCommit, headCommit], {
      cwd,
      stdio: "ignore",
    });
  } catch {
    fail(
      "releaseLines."
        + major
        + ".approval.sourceCommit must be an ancestor of the release candidate",
    );
  }

  let reviewedPolicy;
  try {
    reviewedPolicy = validateReleasePolicy(JSON.parse(execFileSync(
      "git",
      ["show", sourceCommit + ":release/policy.json"],
      {
        cwd,
        encoding: "utf8",
        stdio: ["ignore", "pipe", "ignore"],
      },
    )));
  } catch {
    fail("approved source commit does not contain a valid release/policy.json");
  }
  const expectedReviewedPolicy = structuredClone(policy);
  expectedReviewedPolicy.releaseLines[String(major)].approval = null;
  if (!isDeepStrictEqual(reviewedPolicy, expectedReviewedPolicy)) {
    fail(
      "stable approval for major "
        + major
        + " is stale; only the target approval may differ from the reviewed policy",
    );
  }

  let changedPaths;
  try {
    changedPaths = execFileSync(
      "git",
      ["diff", "--name-only", sourceCommit + ".." + headCommit, "--"],
      {
        cwd,
        encoding: "utf8",
        stdio: ["ignore", "pipe", "ignore"],
      },
    ).trim().split("\n").filter(Boolean);
  } catch {
    fail("cannot compare the approved source commit with the release candidate");
  }
  const forbidden = changedPaths.filter(
    (file) => file !== "release/policy.json" && file !== ".beads/interactions.jsonl",
  );
  if (forbidden.length > 0) {
    fail(
      "stable approval for major "
        + major
        + " is stale; source changed after review: "
        + forbidden.join(", "),
    );
  }

  return policy;
}


/** Load and validate the reviewed release-line policy from the source tree. */
export function loadReleasePolicy(pluginConfig = {}, context = {}) {
  const policyPath = pluginConfig.policyPath ?? DEFAULT_POLICY_PATH;
  requireString(policyPath, "policyPath");
  const absolutePath = path.resolve(context.cwd ?? process.cwd(), policyPath);
  let value;
  try {
    value = JSON.parse(fs.readFileSync(absolutePath, "utf8"));
  } catch (error) {
    fail("cannot read release policy " + absolutePath + ": " + error.message);
  }
  return validateReleasePolicy(value);
}


export function approvedStableMajors(policy) {
  return Object.entries(policy.releaseLines)
    .filter(([major, line]) => major !== "0" && line.approval !== null)
    .map(([major]) => Number(major));
}


/**
 * Validate semantic-release's Git-derived lineage. Development manifests are
 * deliberately sentinels, so a canonical prior tag is required and is the
 * only input allowed to select the next version.
 */
export function validateTagLineage(lastRelease, approvedStableMajors) {
  const version = lastRelease?.version;
  const gitTag = lastRelease?.gitTag;

  if (
    typeof version !== "string"
    || version.length === 0
    || typeof gitTag !== "string"
    || gitTag.length === 0
  ) {
    fail("a canonical previous v<SemVer> Git tag is required; package metadata cannot seed release lineage");
  }

  const parsed = semver.parse(version);
  if (parsed?.build.length > 0) {
    fail("last release version " + version + " uses unsupported build metadata");
  }
  if (semver.valid(version) !== version) {
    fail("last release version " + JSON.stringify(version) + " is not canonical SemVer");
  }
  if (version === "0.0.0") {
    fail("the 0.0.0 development sentinel cannot be a release tag");
  }
  if (parsed.prerelease.length > 0) {
    fail("prerelease lineage " + version + " is unsupported on the normal master release channel");
  }
  if (gitTag !== "v" + version) {
    fail("last release tag " + JSON.stringify(gitTag) + " must exactly match v" + version);
  }
  if (parsed.major > 0 && !approvedStableMajors.includes(parsed.major)) {
    fail("stable major " + parsed.major + " is not present in approvedStableMajors");
  }

  return parsed;
}


function commitsUseReleaseScope(context, expectedScope) {
  return (context?.commits ?? []).some((commit) => {
    const header = String(commit.message ?? "").split("\n", 1)[0];
    const match = COMMIT_HEADER_PATTERN.exec(header);
    return match !== null
      && match[2] === expectedScope
      && ["feat", "fix", "perf"].includes(match[1]);
  });
}


/**
 * Delegate Conventional Commit parsing to the installed official analyzer,
 * then apply only haxe.go's major-zero and stable-major approval rules.
 */
export async function analyzeCommits(pluginConfig, context) {
  const policy = loadReleasePolicy(pluginConfig, context);
  const approvedMajors = approvedStableMajors(policy);
  const lineage = validateTagLineage(context?.lastRelease, approvedMajors);
  const conventionalType = await analyzeConventionalCommits(
    {
      preset: "conventionalcommits",
      presetConfig: { preMajor: true },
      parserOpts: {
        headerPattern: COMMIT_HEADER_PATTERN,
        headerCorrespondence: ["type", "scope", "breakingMarker", "subject"],
      },
      releaseRules: [
        { breakingMarker: "!", release: "major" },
        { type: "fix", scope: "experimental", release: "minor" },
        { type: "perf", scope: "experimental", release: "minor" },
        { type: "fix", scope: "deprecation", release: "minor" },
      ],
    },
    context,
  );

  if (
    lineage.major > 0
    && conventionalType !== null
    && conventionalType !== "major"
    && commitsUseReleaseScope(context, "experimental")
  ) {
    fail(
      "a stable-line experimental change requires explicit surface proof; "
        + "no proof mechanism is configured",
    );
  }

  if (conventionalType !== "major") {
    return conventionalType;
  }

  const nextStableMajor = lineage.major + 1;
  if (approvedMajors.includes(nextStableMajor)) {
    validateApprovalSourceCommit(policy, nextStableMajor, context);
    return "major";
  }
  if (lineage.major === 0) {
    context.logger.log(
      "haxe.go stable major 1 is not approved; treating the breaking 0.x change as a minor release",
    );
    return policy.releaseLines["0"].breakingBump;
  }

  fail(
    "breaking stable major "
      + lineage.major
      + " requires independent approval for major "
      + nextStableMajor,
  );
}


/**
 * Generate hosted notes from the exact previous and next Git tags. This does
 * not write a tracked changelog or mutate the tested checkout.
 */
export async function generateNotes(_pluginConfig, context) {
  return generateConventionalNotes(
    {
      preset: "conventionalcommits",
      presetConfig: { preMajor: true },
    },
    context,
  );
}
