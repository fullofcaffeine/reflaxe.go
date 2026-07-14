import { analyzeCommits as analyzeConventionalCommits } from "@semantic-release/commit-analyzer";
import { generateNotes as generateConventionalNotes } from "@semantic-release/release-notes-generator";
import semver from "semver";


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


/**
 * Stable-major approvals must be exactly 1..N. This makes each future stable
 * major a separately reviewed decision and rejects skipped or duplicated
 * approvals.
 */
export function validateApprovedStableMajors(value) {
  if (!Array.isArray(value)) {
    fail("approvedStableMajors must be an array of explicitly approved positive integers");
  }

  for (let index = 0; index < value.length; index += 1) {
    const expected = index + 1;
    if (!Number.isSafeInteger(value[index]) || value[index] !== expected) {
      fail(
        "approvedStableMajors must be the contiguous sequence 1..N; expected "
          + expected
          + " at index "
          + index,
      );
    }
  }

  return value;
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


/**
 * Delegate Conventional Commit parsing to the installed official analyzer,
 * then apply only haxe.go's major-zero and stable-major approval rules.
 */
export async function analyzeCommits(pluginConfig, context) {
  const approvedStableMajors = validateApprovedStableMajors(
    pluginConfig?.approvedStableMajors ?? [],
  );
  const lineage = validateTagLineage(context?.lastRelease, approvedStableMajors);
  const conventionalType = await analyzeConventionalCommits(
    {
      preset: "conventionalcommits",
      presetConfig: { preMajor: true },
      parserOpts: {
        headerPattern: /^(\w*)(?:\((.*)\))?(!)?: (.*)$/,
        headerCorrespondence: ["type", "scope", "breakingMarker", "subject"],
      },
      releaseRules: [{ breakingMarker: "!", release: "major" }],
    },
    context,
  );

  if (conventionalType !== "major") {
    return conventionalType;
  }

  const nextStableMajor = lineage.major + 1;
  if (approvedStableMajors.includes(nextStableMajor)) {
    return "major";
  }
  if (lineage.major === 0) {
    context.logger.log(
      "haxe.go stable major 1 is not approved; treating the breaking 0.x change as a minor release",
    );
    return "minor";
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
