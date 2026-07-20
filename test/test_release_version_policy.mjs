#!/usr/bin/env node

import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { Writable } from "node:stream";
import { fileURLToPath, pathToFileURL } from "node:url";
import semanticRelease from "semantic-release";
import semver from "semver";
import {
  ReleasePolicyError,
  analyzeCommits,
  validateReleasePolicy,
} from "../scripts/release/analyze-commits.mjs";


const silentLogger = {
  log() {},
  error() {},
};

const silentStream = new Writable({
  write(_chunk, _encoding, callback) {
    callback();
  },
});

const sourcePolicy = JSON.parse(
  readFileSync(new URL("../release/policy.json", import.meta.url), "utf8"),
);
const currentSourceCommit = execFileSync(
  "git",
  ["rev-parse", "HEAD"],
  { encoding: "utf8" },
).trim();

function policyWithApprovedMajors(approved = []) {
  const policy = structuredClone(sourcePolicy);
  const largest = Math.max(1, ...approved);
  for (let major = 1; major <= largest; major += 1) {
    if (major > 1) {
      policy.releaseLines[String(major)] = {
        stage: "stable",
        requirements: [{
          id: "next-major-review",
          status: "pending",
          record: "review-major-" + major,
        }],
        approval: null,
      };
    }
    if (approved.includes(major)) {
      const line = policy.releaseLines[String(major)];
      for (const requirement of line.requirements) {
        requirement.status = "complete";
      }
      line.approval = {
        approvedBy: "Release maintainers",
        record: "approval-major-" + major,
        date: "2026-07-19",
        sourceCommit: currentSourceCommit,
      };
    }
  }
  return policy;
}

async function withPolicy(policy, action) {
  const directory = mkdtempSync(join(tmpdir(), "haxe-go-policy-fixture-"));
  const policyPath = join(directory, "policy.json");
  writeFileSync(policyPath, JSON.stringify(policy));
  try {
    return await action(policyPath);
  } finally {
    rmSync(directory, { recursive: true, force: true });
  }
}

async function analyzeWithPolicy(policy, releaseContext) {
  return withPolicy(
    policy,
    (policyPath) => analyzeCommits({ policyPath }, releaseContext),
  );
}

function context(version, commits, overrides = {}) {
  return {
    cwd: process.cwd(),
    branch: { name: "master", type: "release" },
    commits: commits.map((message, index) => ({
      message,
      hash: "commit-" + index,
    })),
    lastRelease: {
      version,
      gitTag: "v" + version,
      channels: [null],
    },
    logger: silentLogger,
    ...overrides,
  };
}

async function expectRelease({ name, version, commits, approved = [], type, next }) {
  const actual = await analyzeWithPolicy(
    policyWithApprovedMajors(approved),
    context(version, commits, {
      options: { packageVersion: "99.99.99" },
    }),
  );
  assert.equal(actual, type, name + ": release type");
  assert.equal(
    actual === null ? null : semver.inc(version, actual),
    next,
    name + ": next version",
  );
}

async function expectFailure(name, action, pattern) {
  await assert.rejects(action, (error) => {
    assert.ok(error instanceof ReleasePolicyError, name + ": policy error type");
    assert.equal(error.code, "EHAXEGORELEASEPOLICY", name + ": policy error code");
    assert.match(error.message, pattern, name + ": policy diagnostic");
    return true;
  });
}

async function proveSemanticReleaseIntegration() {
  const cwd = mkdtempSync(join(tmpdir(), "haxe-go-version-policy-"));
  const policyPlugin = fileURLToPath(
    new URL("../scripts/release/analyze-commits.mjs", import.meta.url),
  );
  const git = (...args) => execFileSync("git", args, { cwd, stdio: "ignore" });

  try {
    git("init", "-q", "-b", "master");
    git("config", "user.name", "Release Policy");
    git("config", "user.email", "release@example.test");
    writeFileSync(
      join(cwd, "package.json"),
      '{"name":"lineage-fixture","version":"99.99.99"}\n',
    );
    writeFileSync(join(cwd, "release-policy.json"), JSON.stringify(sourcePolicy));
    writeFileSync(join(cwd, "fixture.txt"), "baseline\n");
    git("add", ".");
    git("commit", "-qm", "chore: baseline");
    git("tag", "v0.53.1");
    writeFileSync(join(cwd, "fixture.txt"), "baseline\nfixed\n");
    git("add", ".");
    git("commit", "-qm", "fix: exercise tag lineage");

    const result = await semanticRelease(
      {
        branches: ["master"],
        tagFormat: "v$" + "{version}",
        repositoryUrl: pathToFileURL(cwd).href,
        plugins: [[policyPlugin, { policyPath: "release-policy.json" }]],
        dryRun: true,
        ci: false,
      },
      {
        cwd,
        env: process.env,
        stdout: silentStream,
        stderr: silentStream,
      },
    );

    assert.equal(result?.nextRelease?.type, "patch");
    assert.equal(result?.nextRelease?.version, "0.53.2");
    assert.equal(result?.nextRelease?.gitTag, "v0.53.2");
    assert.match(
      result?.nextRelease?.notes ?? "",
      /compare\/v0\.53\.1\.\.\.v0\.53\.2/,
    );

    const published = await semanticRelease(
      {
        branches: ["master"],
        tagFormat: "v$" + "{version}",
        repositoryUrl: pathToFileURL(cwd).href,
        plugins: [[policyPlugin, { policyPath: "release-policy.json" }]],
        ci: false,
      },
      {
        cwd,
        env: process.env,
        stdout: silentStream,
        stderr: silentStream,
      },
    );
    assert.equal(published?.nextRelease?.gitTag, "v0.53.2");
    assert.equal(
      execFileSync("git", ["rev-parse", "v0.53.2^{commit}"], {
        cwd,
        encoding: "utf8",
      }).trim(),
      execFileSync("git", ["rev-parse", "HEAD"], {
        cwd,
        encoding: "utf8",
      }).trim(),
      "version-only semantic-release must create the tag at tested HEAD",
    );
  } finally {
    rmSync(cwd, { recursive: true, force: true });
  }
}

async function proveApprovalSourceBinding() {
  const cwd = mkdtempSync(join(tmpdir(), "haxe-go-approval-binding-"));
  const git = (...args) => execFileSync("git", args, { cwd, encoding: "utf8" }).trim();

  try {
    git("init", "-q", "-b", "master");
    git("config", "user.name", "Release Policy");
    git("config", "user.email", "release@example.test");
    mkdirSync(join(cwd, "release"));
    const reviewedPolicy = policyWithApprovedMajors([1]);
    reviewedPolicy.releaseLines["1"].approval = null;
    writeFileSync(join(cwd, "release", "policy.json"), JSON.stringify(reviewedPolicy));
    writeFileSync(join(cwd, "fixture.txt"), "reviewed source\n");
    git("add", ".");
    git("commit", "-qm", "chore: reviewed source");
    const reviewedSourceCommit = git("rev-parse", "HEAD");

    const approvedPolicy = structuredClone(reviewedPolicy);
    approvedPolicy.releaseLines["1"].approval = {
      approvedBy: "Release maintainers",
      record: "approval-major-1",
      date: "2026-07-19",
      sourceCommit: reviewedSourceCommit,
    };
    writeFileSync(join(cwd, "release", "policy.json"), JSON.stringify(approvedPolicy));
    git("add", "release/policy.json");
    git("commit", "-qm", "chore: record stable approval");

    const releaseType = await analyzeCommits(
      { policyPath: "release/policy.json" },
      context("0.99.4", ["feat!: graduate reviewed source"], { cwd }),
    );
    assert.equal(releaseType, "major", "approval-only metadata delta permits graduation");

    writeFileSync(join(cwd, "fixture.txt"), "reviewed source\nchanged afterward\n");
    git("add", "fixture.txt");
    git("commit", "-qm", "fix: mutate source after approval");
    await expectFailure(
      "stale approval after source change",
      () => analyzeCommits(
        { policyPath: "release/policy.json" },
        context("0.99.4", ["feat!: attempt stale graduation"], { cwd }),
      ),
      /approval for major 1 is stale.*fixture\.txt/,
    );

    const stablePatch = await analyzeCommits(
      { policyPath: "release/policy.json" },
      context("1.0.0", ["fix: ordinary stable maintenance"], { cwd }),
    );
    assert.equal(stablePatch, "patch", "past approval does not block stable maintenance");

    writeFileSync(join(cwd, "fixture.txt"), "reviewed source\n");
    git("add", "fixture.txt");
    git("commit", "-qm", "revert: restore reviewed source tree");
    approvedPolicy.releaseLines["2"] = {
      stage: "stable",
      requirements: [{
        id: "next-major-review",
        status: "pending",
        record: "review-major-2",
      }],
      approval: null,
    };
    writeFileSync(join(cwd, "release", "policy.json"), JSON.stringify(approvedPolicy));
    git("add", "release/policy.json");
    git("commit", "-qm", "chore: alter policy after approval");
    await expectFailure(
      "policy drift after approval",
      () => analyzeCommits(
        { policyPath: "release/policy.json" },
        context("0.99.4", ["feat!: attempt policy-drift graduation"], { cwd }),
      ),
      /only the target approval may differ/,
    );
  } finally {
    rmSync(cwd, { recursive: true, force: true });
  }
}

await expectRelease({
  name: "major-zero fix",
  version: "0.53.1",
  commits: ["fix: preserve release identity"],
  type: "patch",
  next: "0.53.2",
});
await expectRelease({
  name: "major-zero feature",
  version: "0.53.1",
  commits: ["feat: stage Haxelib metadata"],
  type: "minor",
  next: "0.54.0",
});
await expectRelease({
  name: "experimental fix has a minor floor",
  version: "0.53.1",
  commits: ["fix(experimental): revise an excluded surface"],
  type: "minor",
  next: "0.54.0",
});
await expectRelease({
  name: "deprecation notice has a minor floor",
  version: "0.53.1",
  commits: ["fix(deprecation): announce a supported replacement"],
  type: "minor",
  next: "0.54.0",
});
await expectRelease({
  name: "major-zero breaking header",
  version: "0.53.1",
  commits: ["feat!: revise the release contract"],
  type: "minor",
  next: "0.54.0",
});
await expectRelease({
  name: "major-zero breaking footer",
  version: "0.53.1",
  commits: ["fix: revise contract\n\nBREAKING CHANGE: version staging changed"],
  type: "minor",
  next: "0.54.0",
});
await expectRelease({
  name: "no release",
  version: "0.53.1",
  commits: ["docs: explain release identity", "test: cover metadata staging"],
  type: null,
  next: null,
});

await expectFailure(
  "missing lineage",
  () => analyzeWithPolicy(
    policyWithApprovedMajors(),
    context("0.53.1", ["fix: test"], { lastRelease: {} }),
  ),
  /previous v<SemVer> Git tag is required/,
);
await expectFailure(
  "mismatched tag",
  () => analyzeWithPolicy(
    policyWithApprovedMajors(),
    context("0.53.1", ["fix: test"], {
      lastRelease: { version: "0.53.1", gitTag: "v9.9.9", channels: [null] },
    }),
  ),
  /must exactly match v0\.53\.1/,
);
await expectFailure(
  "development sentinel tag",
  () => analyzeWithPolicy(
    policyWithApprovedMajors(),
    context("0.0.0", ["fix: test"]),
  ),
  /development sentinel cannot be a release tag/,
);
await expectFailure(
  "invalid semver",
  () => analyzeWithPolicy(
    policyWithApprovedMajors(),
    context("0.53", ["fix: test"]),
  ),
  /not canonical SemVer/,
);
await expectFailure(
  "build metadata",
  () => analyzeWithPolicy(
    policyWithApprovedMajors(),
    context("0.53.1+rebuilt", ["fix: test"]),
  ),
  /unsupported build metadata/,
);
await expectFailure(
  "unknown stable major",
  () => analyzeWithPolicy(
    policyWithApprovedMajors([1]),
    context("2.1.0", ["fix: test"]),
  ),
  /stable major 2 is not present/,
);
await expectFailure(
  "unapproved stable breaking change",
  () => analyzeWithPolicy(
    policyWithApprovedMajors([1]),
    context("1.7.2", ["feat!: break stable API"]),
  ),
  /requires independent approval for major 2/,
);
await expectFailure(
  "stable experimental change without surface proof",
  () => analyzeWithPolicy(
    policyWithApprovedMajors([1]),
    context("1.7.2", ["fix(experimental): revise an excluded surface"]),
  ),
  /requires explicit surface proof/,
);
const nonContiguousPolicy = policyWithApprovedMajors([1, 3]);
await expectFailure(
  "non-contiguous approvals",
  () => Promise.resolve().then(() => validateReleasePolicy(nonContiguousPolicy)),
  /approvals must be contiguous/,
);

const prematureApprovalPolicy = policyWithApprovedMajors();
prematureApprovalPolicy.releaseLines["1"].approval = {
  approvedBy: "Release maintainers",
  record: "premature-approval",
  date: "2026-07-19",
  sourceCommit: "b".repeat(40),
};
await expectFailure(
  "stable approval with pending requirements",
  () => Promise.resolve().then(() => validateReleasePolicy(prematureApprovalPolicy)),
  /cannot be approved while requirements are pending/,
);

const unknownTopLevelAuthorityPolicy = policyWithApprovedMajors();
unknownTopLevelAuthorityPolicy.approvedStableMajors = [1];
await expectFailure(
  "unknown top-level approval authority",
  () => Promise.resolve().then(() => validateReleasePolicy(unknownTopLevelAuthorityPolicy)),
  /release policy must contain exactly/,
);

const unknownReleaseLineAuthorityPolicy = policyWithApprovedMajors();
unknownReleaseLineAuthorityPolicy.releaseLines["1"].approved = true;
await expectFailure(
  "unknown release-line approval authority",
  () => Promise.resolve().then(() => validateReleasePolicy(unknownReleaseLineAuthorityPolicy)),
  /releaseLines\.1 must contain exactly/,
);

const nonexistentApprovalCommitPolicy = policyWithApprovedMajors([1]);
nonexistentApprovalCommitPolicy.releaseLines["1"].approval.sourceCommit = "f".repeat(40);
await expectFailure(
  "nonexistent approval source commit",
  () => analyzeWithPolicy(
    nonexistentApprovalCommitPolicy,
    context("0.99.4", ["feat!: attempt unbound stable graduation"]),
  ),
  /does not exist in the release repository/,
);

const prereleaseChannelPolicy = policyWithApprovedMajors();
prereleaseChannelPolicy.normalChannel.prerelease = true;
await expectFailure(
  "normal channel prerelease drift",
  () => Promise.resolve().then(() => validateReleasePolicy(prereleaseChannelPolicy)),
  /normalChannel must keep master/,
);

const deprecationDriftPolicy = policyWithApprovedMajors();
deprecationDriftPolicy.deprecationPolicy.majorZero.earliestRemovalMinorOffset = 1;
await expectFailure(
  "major-zero deprecation drift",
  () => Promise.resolve().then(() => validateReleasePolicy(deprecationDriftPolicy)),
  /earliest removal N\+2/,
);

const experimentalDriftPolicy = policyWithApprovedMajors();
experimentalDriftPolicy.experimentalPolicy.minimumChangeRelease = "patch";
await expectFailure(
  "experimental patch loophole",
  () => Promise.resolve().then(() => validateReleasePolicy(experimentalDriftPolicy)),
  /experimental policy must stay excluded/,
);

const malformedApprovalPolicy = policyWithApprovedMajors();
malformedApprovalPolicy.releaseLines["1"].approval = {
  approvedBy: "Release maintainers",
  record: "malformed-approval",
  date: "2026-07-19",
  sourceCommit: "short-sha",
};
await expectFailure(
  "malformed approval source commit",
  () => Promise.resolve().then(() => validateReleasePolicy(malformedApprovalPolicy)),
  /sourceCommit must be a full lowercase Git commit/,
);

await expectFailure(
  "missing release policy",
  () => analyzeCommits(
    { policyPath: "release/does-not-exist.json" },
    context("0.53.1", ["fix: test"]),
  ),
  /cannot read release policy/,
);

await proveSemanticReleaseIntegration();
await proveApprovalSourceBinding();
console.log("[release-version-policy] OK: tag-derived 0.x and commit-bound stable majors");
