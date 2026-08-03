#!/usr/bin/env node

import assert from "node:assert/strict";
import { createHash } from "node:crypto";
import {
  mkdtempSync,
  readFileSync,
  rmSync,
  symlinkSync,
  writeFileSync,
} from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import {
  GitHubReleaseAdapter,
  ReleaseStateError,
  canonicalStableTag,
  executeReleaseCommand,
  loadExpectedAssetManifest,
  reconcileHostedRelease,
  selectLatestStableTag,
} from "../scripts/release/reconcile-github-release.mjs";

const TAG = "v0.54.0";
const SOURCE_SHA = "a".repeat(40);

function sha256(bytes) {
  return createHash("sha256").update(bytes).digest("hex");
}

function releaseRecord({
  draft = true,
  immutable = false,
  assets = [],
  body = "Release notes fixture",
} = {}) {
  return {
    id: 42,
    tag_name: TAG,
    name: TAG,
    target_commitish: SOURCE_SHA,
    draft,
    prerelease: false,
    immutable,
    body,
    assets,
  };
}

function hostedAsset(expected, id, overrides = {}) {
  return {
    id,
    name: expected.name,
    state: "uploaded",
    size: expected.size,
    digest: expected.digest,
    ...overrides,
  };
}

class FakeGitHubAdapter {
  constructor({ release = null, tagSha = SOURCE_SHA } = {}) {
    this.release = release;
    this.releaseReads = null;
    this.tagSha = tagSha;
    this.tagShas = null;
    this.operations = [];
    this.nextAssetId = 100;
    this.loseCreateResponse = false;
    this.loseUploadResponseFor = null;
    this.losePublishResponse = false;
  }

  async getTagCommit(tag) {
    this.operations.push(`get-tag:${tag}`);
    if (Array.isArray(this.tagShas) && this.tagShas.length > 0) {
      return this.tagShas.shift();
    }
    return this.tagSha;
  }

  async getRelease(tag) {
    this.operations.push(`get-release:${tag}`);
    if (Array.isArray(this.releaseReads) && this.releaseReads.length > 0) {
      this.release = this.releaseReads.shift();
    }
    return this.release;
  }

  async createDraft({ tag, sourceSha, notes }) {
    this.operations.push("create-draft");
    this.release = releaseRecord();
    this.release.tag_name = tag;
    this.release.name = tag;
    this.release.target_commitish = sourceSha;
    this.release.body = notes;
    if (this.loseCreateResponse) throw new Error("lost create response");
    return this.release;
  }

  async uploadAsset(_release, expected) {
    this.operations.push(`upload:${expected.name}`);
    this.release.assets.push(
      hostedAsset(expected, this.nextAssetId++),
    );
    if (this.loseUploadResponseFor === expected.name) {
      throw new Error("lost upload response");
    }
  }

  async publishRelease() {
    this.operations.push("publish");
    this.release.draft = false;
    this.release.immutable = true;
    if (this.losePublishResponse) throw new Error("lost publish response");
    return this.release;
  }
}

function completeAdapter(expected, options = {}) {
  return new FakeGitHubAdapter({
    release: releaseRecord({
      draft: options.draft ?? true,
      immutable: options.immutable ?? false,
      assets: expected.map((asset, index) => hostedAsset(asset, index + 1)),
    }),
    tagSha: options.tagSha,
  });
}

async function expectStateFailure(label, action, pattern) {
  await assert.rejects(action, (error) => {
    assert.ok(error instanceof ReleaseStateError, `${label}: error type`);
    assert.equal(error.code, "EHAXEGORELEASESTATE", `${label}: error code`);
    assert.match(error.message, pattern, `${label}: diagnostic`);
    return true;
  });
}

const tempRoot = mkdtempSync(join(tmpdir(), "haxe-go-release-state-"));
try {
  assert.equal(
    executeReleaseCommand(process.execPath, ["-e", "process.stdout.write('captured')"]),
    "captured",
    "child stdout must be captured so the CLI result stays machine-readable JSON",
  );
  const expected = ["archive", "checksum", "manifest", "provenance"].map(
    (kind) => {
      const bytes = Buffer.from(`${kind} fixture bytes\n`);
      const path = join(tempRoot, kind);
      writeFileSync(path, bytes);
      return {
        name: `reflaxe.go-0.54.0.${kind}`,
        path,
        size: bytes.length,
        digest: `sha256:${sha256(bytes)}`,
      };
    },
  );

  assert.equal(canonicalStableTag(TAG), TAG);
  for (const invalid of [
    "0.54.0",
    "v01.2.3",
    "v1.2.3-rc.1",
    "v1.2.3+build.4",
    "v0.0.0",
    "v1.2",
    "main",
  ]) {
    assert.throws(() => canonicalStableTag(invalid), ReleaseStateError);
  }
  assert.equal(
    selectLatestStableTag(["v2.0.0", "v10.0.0", "v1.99.0", "notes"]),
    "v10.0.0",
    "selection must use SemVer precedence, not lexical or platform sort",
  );

  // A missing Release is recovered inside the same command. A lost create
  // response is resolved by re-reading GitHub's authoritative state.
  const absent = new FakeGitHubAdapter();
  absent.loseCreateResponse = true;
  absent.losePublishResponse = true;
  const absentResult = await reconcileHostedRelease({
    mode: "reconcile",
    tag: TAG,
    sourceSha: SOURCE_SHA,
    expectedAssets: expected,
    notes: "Release notes fixture",
    adapter: absent,
    wait: async () => {},
  });
  assert.equal(absentResult.state, "published-immutable");
  assert.equal(
    absent.operations.filter((operation) => operation.startsWith("upload:"))
      .length,
    expected.length,
  );
  assert.equal(absent.operations.filter((operation) => operation === "create-draft").length, 1);
  assert.equal(absent.operations.filter((operation) => operation === "publish").length, 1);

  // A partial draft preserves matching assets and uploads only missing ones,
  // even if the upload succeeded but its response was lost.
  const partial = completeAdapter(expected.slice(0, 2));
  partial.loseUploadResponseFor = expected[2].name;
  const partialResult = await reconcileHostedRelease({
    mode: "reconcile",
    tag: TAG,
    sourceSha: SOURCE_SHA,
    expectedAssets: expected,
    notes: "Release notes fixture",
    adapter: partial,
    wait: async () => {},
  });
  assert.equal(partialResult.state, "published-immutable");
  assert.deepEqual(
    partial.operations.filter((operation) => operation.startsWith("upload:")),
    [`upload:${expected[2].name}`, `upload:${expected[3].name}`],
  );

  // A matching immutable rerun is read-only verification.
  const immutable = completeAdapter(expected, {
    draft: false,
    immutable: true,
  });
  const immutableResult = await reconcileHostedRelease({
    mode: "reconcile",
    tag: TAG,
    sourceSha: SOURCE_SHA,
    expectedAssets: expected,
    adapter: immutable,
    wait: async () => {},
  });
  assert.equal(immutableResult.state, "verified-immutable");
  assert.deepEqual(
    immutable.operations,
    [`get-tag:${TAG}`, `get-release:${TAG}`, `get-tag:${TAG}`],
  );

  // Conflicting or unexpected draft bytes fail before any hosted mutation.
  const conflict = completeAdapter(expected);
  conflict.release.assets[0].digest = `sha256:${"0".repeat(64)}`;
  await expectStateFailure(
    "conflicting draft digest",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: conflict,
      }),
    /digest does not match/,
  );
  assert.deepEqual(conflict.operations, [`get-tag:${TAG}`, `get-release:${TAG}`]);

  const unexpected = completeAdapter(expected);
  unexpected.release.assets.push({
    id: 99,
    name: "surprise.bin",
    state: "uploaded",
    size: 1,
    digest: `sha256:${"f".repeat(64)}`,
  });
  await expectStateFailure(
    "unexpected draft asset",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: unexpected,
      }),
    /unexpected assets: surprise\.bin/,
  );
  assert.deepEqual(unexpected.operations, [`get-tag:${TAG}`, `get-release:${TAG}`]);

  const conflictingNotes = completeAdapter(expected);
  conflictingNotes.release.body = "Generic generated notes";
  await expectStateFailure(
    "conflicting release notes",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        notes: "Release notes fixture",
        adapter: conflictingNotes,
      }),
    /release notes do not match/i,
  );
  assert.deepEqual(
    conflictingNotes.operations,
    [`get-tag:${TAG}`, `get-release:${TAG}`],
  );

  const duplicate = completeAdapter(expected);
  duplicate.release.assets.push({ ...duplicate.release.assets[0], id: 200 });
  await expectStateFailure(
    "duplicate draft asset",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: duplicate,
      }),
    /duplicate asset names/,
  );
  assert.deepEqual(duplicate.operations, [`get-tag:${TAG}`, `get-release:${TAG}`]);

  // Existing public state is never repaired in place.
  for (const [label, adapter, diagnostic] of [
    [
      "mutable published release",
      completeAdapter(expected, { draft: false, immutable: false }),
      /not immutable/,
    ],
    [
      "incomplete immutable release",
      completeAdapter(expected.slice(0, 3), {
        draft: false,
        immutable: true,
      }),
      /missing assets/,
    ],
    [
      "moved remote tag",
      completeAdapter(expected, { tagSha: "b".repeat(40) }),
      /remote tag.*tested source SHA/,
    ],
    [
      "missing remote tag",
      completeAdapter(expected, { tagSha: null }),
      /remote tag.*does not exist/,
    ],
  ]) {
    await expectStateFailure(
      label,
      () =>
        reconcileHostedRelease({
          mode: "reconcile",
          tag: TAG,
          sourceSha: SOURCE_SHA,
          expectedAssets: expected,
          adapter,
        }),
      diagnostic,
    );
    assert.equal(
      adapter.operations.some((operation) =>
        /^(create-draft|upload:|publish)/.test(operation),
      ),
      false,
      `${label}: no mutation`,
    );
  }
  const publishedConflict = completeAdapter(expected, {
    draft: false,
    immutable: true,
  });
  publishedConflict.release.assets[0].digest = `sha256:${"9".repeat(64)}`;
  await expectStateFailure(
    "conflicting immutable digest",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: publishedConflict,
      }),
    /digest does not match/,
  );
  assert.deepEqual(
    publishedConflict.operations,
    [`get-tag:${TAG}`, `get-release:${TAG}`],
  );

  const movedBeforePublish = completeAdapter(expected);
  movedBeforePublish.tagShas = [SOURCE_SHA, "b".repeat(40)];
  await expectStateFailure(
    "tag moved before publish",
    () =>
      reconcileHostedRelease({
        mode: "reconcile",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: movedBeforePublish,
      }),
    /remote tag.*tested source SHA/,
  );
  assert.equal(movedBeforePublish.operations.includes("publish"), false);

  const publishedByAnotherRunner = completeAdapter(expected.slice(0, 2));
  publishedByAnotherRunner.releaseReads = [
    publishedByAnotherRunner.release,
    releaseRecord({
      draft: false,
      immutable: true,
      assets: expected.map((asset, index) => hostedAsset(asset, index + 20)),
    }),
  ];
  const concurrentResult = await reconcileHostedRelease({
    mode: "reconcile",
    tag: TAG,
    sourceSha: SOURCE_SHA,
    expectedAssets: expected,
    adapter: publishedByAnotherRunner,
  });
  assert.equal(concurrentResult.state, "verified-immutable");
  assert.equal(
    publishedByAnotherRunner.operations.some((operation) =>
      /^(upload:|publish)/.test(operation),
    ),
    false,
    "a newly published release becomes verification-only before mutation",
  );

  // Verification mode is read-only and fails rather than creating state.
  const verifyAbsent = new FakeGitHubAdapter();
  await expectStateFailure(
    "verify absent",
    () =>
      reconcileHostedRelease({
        mode: "verify",
        tag: TAG,
        sourceSha: SOURCE_SHA,
        expectedAssets: expected,
        adapter: verifyAbsent,
      }),
    /does not exist/,
  );
  assert.equal(
    verifyAbsent.operations.some((operation) =>
      /^(create-draft|upload:|publish)/.test(operation),
    ),
    false,
  );

  assert.equal(
    expected.every((asset) => readFileSync(asset.path).length === asset.size),
    true,
  );

  const manifestPath = join(tempRoot, "release-assets.json");
  writeFileSync(
    manifestPath,
    `${JSON.stringify({
      schemaVersion: 1,
      tag: TAG,
      sourceSha: SOURCE_SHA,
      assets: expected.map(({ name, path, size, digest }) => ({
        name,
        path: basenameForFixture(path),
        size,
        digest,
      })),
    }, null, 2)}\n`,
  );
  const loaded = loadExpectedAssetManifest(manifestPath);
  assert.equal(loaded.tag, TAG);
  assert.equal(loaded.assets.length, expected.length);
  const escapingManifest = join(tempRoot, "escaping-assets.json");
  writeFileSync(
    escapingManifest,
    `${JSON.stringify({
      schemaVersion: 1,
      tag: TAG,
      sourceSha: SOURCE_SHA,
      assets: [{
        name: expected[0].name,
        path: "../outside.zip",
        size: expected[0].size,
        digest: expected[0].digest,
      }],
    })}\n`,
  );
  assert.throws(
    () => loadExpectedAssetManifest(escapingManifest),
    /path escapes its manifest/,
  );
  const linkedAsset = join(tempRoot, "linked-archive");
  symlinkSync(expected[0].path, linkedAsset);
  const linkedManifest = join(tempRoot, "linked-assets.json");
  writeFileSync(
    linkedManifest,
    `${JSON.stringify({
      schemaVersion: 1,
      tag: TAG,
      sourceSha: SOURCE_SHA,
      assets: [{
        name: expected[0].name,
        path: "linked-archive",
        size: expected[0].size,
        digest: expected[0].digest,
      }],
    })}\n`,
  );
  assert.throws(
    () => loadExpectedAssetManifest(linkedManifest),
    /must be a regular file, not a symlink/,
  );

  const requests = [];
  const apiAdapter = new GitHubReleaseAdapter({
    repository: "fullofcaffeine/reflaxe.go",
    token: "fixture-token",
    fetchImpl: async (url) => {
      requests.push(url);
      if (url.endsWith(`/releases/tags/${TAG}`)) {
        return new Response("not found", { status: 404 });
      }
      if (url.endsWith("/releases?per_page=100&page=1")) {
        return Response.json([releaseRecord()]);
      }
      return new Response("unexpected request", { status: 500 });
    },
  });
  const rediscoveredDraft = await apiAdapter.getRelease(TAG);
  assert.equal(rediscoveredDraft?.draft, true);
  assert.equal(requests.length, 2, "draft lookup must use GitHub API list truth after tag 404");

  const annotatedTagObject = "c".repeat(40);
  const tagRequests = [];
  const tagAdapter = new GitHubReleaseAdapter({
    repository: "fullofcaffeine/reflaxe.go",
    token: "fixture-token",
    fetchImpl: async (url, options) => {
      tagRequests.push({ method: options.method, url });
      if (url.endsWith(`/git/ref/tags/${TAG}`)) {
        return Response.json({ object: { type: "tag", sha: annotatedTagObject } });
      }
      if (url.endsWith(`/git/tags/${annotatedTagObject}`)) {
        return Response.json({ object: { type: "commit", sha: SOURCE_SHA } });
      }
      return new Response("unexpected request", { status: 500 });
    },
  });
  assert.equal(await tagAdapter.getTagCommit(TAG), SOURCE_SHA);
  assert.deepEqual(
    tagRequests.map(({ method }) => method),
    ["GET", "GET"],
    "annotated tag identity uses read-only GitHub API requests",
  );

  const createCommands = [];
  let releaseCreated = false;
  const guardedCreateAdapter = new GitHubReleaseAdapter({
    repository: "fullofcaffeine/reflaxe.go",
    token: "fixture-token",
    executeImpl(command, args, options) {
      createCommands.push({ args, command, input: options.input });
      releaseCreated = true;
    },
    fetchImpl: async (url, options) => {
      if (url.endsWith(`/releases/tags/${TAG}`)) {
        return releaseCreated
          ? Response.json(releaseRecord())
          : new Response("not found", { status: 404 });
      }
      if (url.endsWith("/releases?per_page=100&page=1")) {
        return Response.json([]);
      }
      return new Response(`unexpected ${options.method} request`, { status: 500 });
    },
  });
  const guardedDraft = await guardedCreateAdapter.createDraft({
    tag: TAG,
    sourceSha: SOURCE_SHA,
    notes: "fixture notes",
  });
  assert.equal(guardedDraft?.draft, true);
  assert.equal(createCommands.length, 1);
  assert.equal(createCommands[0].command, "gh");
  assert.equal(createCommands[0].args.includes("--verify-tag"), true);
  assert.equal(createCommands[0].args.includes("--draft"), true);
  assert.equal(createCommands[0].input, "fixture notes");

  const mutationRequests = [];
  const mutationAdapter = new GitHubReleaseAdapter({
    repository: "fullofcaffeine/reflaxe.go",
    token: "fixture-token",
    fetchImpl: async (url, options) => {
      mutationRequests.push({
        body: options.body,
        contentType: options.headers["Content-Type"],
        method: options.method,
        url,
      });
      return Response.json({ ok: true });
    },
  });
  await mutationAdapter.uploadAsset(releaseRecord(), expected[0]);
  await mutationAdapter.publishRelease(releaseRecord());
  assert.match(
    mutationRequests[0].url,
    /uploads\.github\.com\/repos\/fullofcaffeine\/reflaxe\.go\/releases\/42\/assets\?name=/,
  );
  assert.equal(mutationRequests[0].method, "POST");
  assert.equal(mutationRequests[0].contentType, "application/octet-stream");
  assert.deepEqual(mutationRequests[0].body, readFileSync(expected[0].path));
  assert.match(mutationRequests[1].url, /api\.github\.com\/repos\/fullofcaffeine\/reflaxe\.go\/releases\/42$/);
  assert.equal(mutationRequests[1].method, "PATCH");
  assert.equal(mutationRequests[1].contentType, "application/json");
  assert.deepEqual(JSON.parse(mutationRequests[1].body), { draft: false });

  console.log(
    "[release-reconciliation] OK: strict SemVer and fresh, retry, conflict, and immutable states",
  );
} finally {
  rmSync(tempRoot, { recursive: true, force: true });
}

function basenameForFixture(path) {
  return path.slice(path.lastIndexOf("/") + 1);
}
