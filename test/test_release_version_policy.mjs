#!/usr/bin/env node

import assert from "node:assert/strict";
import { execFileSync } from "node:child_process";
import { mkdtempSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { Writable } from "node:stream";
import { fileURLToPath, pathToFileURL } from "node:url";
import semanticRelease from "semantic-release";
import semver from "semver";
import {
  ReleasePolicyError,
  analyzeCommits,
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
  const actual = await analyzeCommits(
    { approvedStableMajors: approved },
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
        plugins: [[policyPlugin, { approvedStableMajors: [] }]],
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
  name: "approved stable graduation",
  version: "0.99.4",
  commits: ["feat!: approve the stable contract"],
  approved: [1],
  type: "major",
  next: "1.0.0",
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
  () => analyzeCommits(
    { approvedStableMajors: [] },
    context("0.53.1", ["fix: test"], { lastRelease: {} }),
  ),
  /previous v<SemVer> Git tag is required/,
);
await expectFailure(
  "mismatched tag",
  () => analyzeCommits(
    { approvedStableMajors: [] },
    context("0.53.1", ["fix: test"], {
      lastRelease: { version: "0.53.1", gitTag: "v9.9.9", channels: [null] },
    }),
  ),
  /must exactly match v0\.53\.1/,
);
await expectFailure(
  "development sentinel tag",
  () => analyzeCommits(
    { approvedStableMajors: [] },
    context("0.0.0", ["fix: test"]),
  ),
  /development sentinel cannot be a release tag/,
);
await expectFailure(
  "invalid semver",
  () => analyzeCommits(
    { approvedStableMajors: [] },
    context("0.53", ["fix: test"]),
  ),
  /not canonical SemVer/,
);
await expectFailure(
  "build metadata",
  () => analyzeCommits(
    { approvedStableMajors: [] },
    context("0.53.1+rebuilt", ["fix: test"]),
  ),
  /unsupported build metadata/,
);
await expectFailure(
  "unknown stable major",
  () => analyzeCommits(
    { approvedStableMajors: [1] },
    context("2.1.0", ["fix: test"]),
  ),
  /stable major 2 is not present/,
);
await expectFailure(
  "unapproved stable breaking change",
  () => analyzeCommits(
    { approvedStableMajors: [1] },
    context("1.7.2", ["feat!: break stable API"]),
  ),
  /requires independent approval for major 2/,
);
await expectFailure(
  "non-contiguous approvals",
  () => analyzeCommits(
    { approvedStableMajors: [1, 3] },
    context("0.53.1", ["feat!: test"]),
  ),
  /contiguous sequence 1\.\.N/,
);

await proveSemanticReleaseIntegration();
console.log("[release-version-policy] OK: tag-derived 0.x and approved stable majors");
