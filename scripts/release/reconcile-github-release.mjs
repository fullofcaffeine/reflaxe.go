#!/usr/bin/env node

import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";
import {
  lstatSync,
  readFileSync,
} from "node:fs";
import { basename, dirname, isAbsolute, relative, resolve, sep } from "node:path";
import { fileURLToPath } from "node:url";
import semver from "semver";

const ERROR_CODE = "EHAXEGORELEASESTATE";
const FULL_SHA = /^[0-9a-f]{40}$/;
const SHA256_DIGEST = /^sha256:[0-9a-f]{64}$/;

/**
 * What: A stable diagnostic type for release-state policy failures.
 * Why: A retryable network error must not be confused with a contradictory
 * hosted release, because only the former may safely be retried.
 * How: Every policy rejection uses one code while preserving a direct,
 * operator-readable explanation in the message.
 */
export class ReleaseStateError extends Error {
  constructor(message, options = undefined) {
    super(message, options);
    this.name = "ReleaseStateError";
    this.code = ERROR_CODE;
  }
}

function fail(message, options = undefined) {
  throw new ReleaseStateError(message, options);
}

function normalizeSourceSha(value, label = "source SHA") {
  const sha = String(value ?? "").trim().toLowerCase();
  if (!FULL_SHA.test(sha)) fail(`${label} must be a full lowercase Git commit SHA`);
  return sha;
}

export function executeReleaseCommand(command, args, options = {}) {
  return execFileSync(command, args, {
    encoding: "utf8",
    stdio: ["pipe", "pipe", "inherit"],
    ...options,
  });
}

/**
 * What: Parse the one release-tag syntax accepted by haxe.go.
 * Why: Shell regular expressions and version-aware platform sorting implement
 * only approximations of SemVer and can select a different public identity.
 * How: The locked `semver` library parses the version, then this function
 * requires its exact canonical stable spelling and rejects the 0.0.0 sentinel.
 */
export function canonicalStableTag(value) {
  if (typeof value !== "string" || !value.startsWith("v")) {
    fail("release tag must use canonical vMAJOR.MINOR.PATCH form");
  }
  const rawVersion = value.slice(1);
  const parsed = semver.parse(rawVersion, { loose: false, includePrerelease: true });
  if (
    parsed === null
    || parsed.prerelease.length !== 0
    || parsed.build.length !== 0
    || rawVersion !== `${parsed.major}.${parsed.minor}.${parsed.patch}`
  ) {
    fail(`release tag is not canonical stable SemVer: ${value}`);
  }
  if (parsed.major === 0 && parsed.minor === 0 && parsed.patch === 0) {
    fail("the v0.0.0 development sentinel cannot become a release tag");
  }
  return `v${parsed.version}`;
}

/** Select the highest canonical stable release tag using SemVer precedence. */
export function selectLatestStableTag(values) {
  if (!Array.isArray(values)) fail("release tag candidates must be an array");
  const tags = [];
  for (const value of values) {
    try {
      tags.push(canonicalStableTag(value));
    } catch (error) {
      if (!(error instanceof ReleaseStateError)) throw error;
    }
  }
  if (tags.length === 0) fail("no canonical stable SemVer tag is available");
  tags.sort((left, right) => semver.rcompare(left.slice(1), right.slice(1)));
  return tags[0];
}

/** Return every canonical stable tag in descending SemVer order. */
export function canonicalStableTags(values) {
  if (!Array.isArray(values)) fail("release tag candidates must be an array");
  const tags = [];
  for (const value of values) {
    try {
      tags.push(canonicalStableTag(value));
    } catch (error) {
      if (!(error instanceof ReleaseStateError)) throw error;
    }
  }
  return [...new Set(tags)].sort((left, right) =>
    semver.rcompare(left.slice(1), right.slice(1))
  );
}

function normalizeExpectedAssets(expectedAssets) {
  if (!Array.isArray(expectedAssets) || expectedAssets.length === 0) {
    fail("release reconciliation requires at least one expected asset");
  }
  const names = new Set();
  return expectedAssets.map((asset, index) => {
    const label = `expected asset ${index + 1}`;
    if (!asset || typeof asset !== "object" || Array.isArray(asset)) {
      fail(`${label} must be an object`);
    }
    if (
      typeof asset.name !== "string"
      || asset.name.length === 0
      || basename(asset.name) !== asset.name
    ) {
      fail(`${label} name must be one safe file name`);
    }
    if (names.has(asset.name)) fail(`expected assets contain duplicate name ${asset.name}`);
    names.add(asset.name);
    if (!Number.isSafeInteger(asset.size) || asset.size < 0) {
      fail(`${asset.name} expected size must be a non-negative integer`);
    }
    if (typeof asset.digest !== "string" || !SHA256_DIGEST.test(asset.digest)) {
      fail(`${asset.name} expected digest must use lowercase sha256:<64 hex>`);
    }
    return Object.freeze({ ...asset });
  });
}

function assetMap(release, expectedAssets) {
  if (!Array.isArray(release.assets)) fail("GitHub Release assets must be an array");
  const names = release.assets.map((asset) => asset?.name);
  if (names.some((name) => typeof name !== "string" || name.length === 0)) {
    fail("GitHub Release contains an asset without a valid name");
  }
  if (new Set(names).size !== names.length) {
    fail("GitHub Release contains duplicate asset names");
  }
  const allowed = new Set(expectedAssets.map((asset) => asset.name));
  const unexpected = names.filter((name) => !allowed.has(name)).sort();
  if (unexpected.length > 0) {
    fail(`GitHub Release contains unexpected assets: ${unexpected.join(", ")}`);
  }
  return new Map(release.assets.map((asset) => [asset.name, asset]));
}

function verifyOneAsset(hosted, expected) {
  if (hosted.state !== "uploaded") {
    fail(`${expected.name} is not in uploaded state`);
  }
  if (hosted.size !== expected.size) {
    fail(`${expected.name} hosted size does not match the approved file`);
  }
  if (hosted.digest !== expected.digest) {
    fail(`${expected.name} hosted digest does not match the approved file`);
  }
}

function inspectAssets(release, expectedAssets) {
  const hosted = assetMap(release, expectedAssets);
  for (const expected of expectedAssets) {
    const actual = hosted.get(expected.name);
    if (actual) verifyOneAsset(actual, expected);
  }
  return {
    hosted,
    missing: expectedAssets.filter((asset) => !hosted.has(asset.name)),
  };
}

function verifyReleaseMetadata(release, tag, notes = "") {
  if (!release || typeof release !== "object") fail(`GitHub Release ${tag} does not exist`);
  if (release.tag_name !== tag) fail("GitHub Release tag does not match release identity");
  if (release.prerelease !== false) fail("stable GitHub Release must not be a prerelease");
  if (typeof release.draft !== "boolean") fail("GitHub Release draft state is missing");
  if (release.draft && release.immutable === true) {
    fail("draft GitHub Release unexpectedly reports immutable");
  }
  if (notes.length > 0 && release.body !== notes) {
    fail("GitHub Release notes do not match the approved bounded release notes");
  }
}

function requireCompleteAssets(release, expectedAssets) {
  const { missing } = inspectAssets(release, expectedAssets);
  if (missing.length > 0) {
    fail(`GitHub Release is missing assets: ${missing.map((asset) => asset.name).join(", ")}`);
  }
}

async function recoverLostMutation(adapter, tag, error) {
  try {
    return await adapter.getRelease(tag);
  } catch (queryError) {
    throw new ReleaseStateError(
      `GitHub mutation failed and authoritative state could not be re-read: ${error.message}`,
      { cause: queryError },
    );
  }
}

async function verifyRemoteTagIdentity(adapter, tag, sourceSha) {
  const remoteTagSha = await adapter.getTagCommit(tag);
  if (remoteTagSha === null || remoteTagSha === undefined) {
    fail(`remote tag ${tag} does not exist; reconciliation never creates tags`);
  }
  if (normalizeSourceSha(remoteTagSha, `remote tag ${tag}`) !== sourceSha) {
    fail(`remote tag ${tag} does not identify the tested source SHA ${sourceSha}`);
  }
}

async function verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes }) {
  verifyReleaseMetadata(release, tag, notes);
  if (release.draft) fail(`GitHub Release ${tag} is still a draft`);
  if (release.immutable !== true) fail(`published GitHub Release ${tag} is not immutable`);
  requireCompleteAssets(release, expectedAssets);
  await verifyRemoteTagIdentity(adapter, tag, sourceSha);
  return { state: "verified-immutable", release };
}

async function waitForImmutable({ adapter, tag, expectedAssets, notes, attempts, delayMs, wait }) {
  let lastRelease = null;
  for (let attempt = 1; attempt <= attempts; attempt += 1) {
    lastRelease = await adapter.getRelease(tag);
    if (lastRelease && lastRelease.draft === false && lastRelease.immutable === true) {
      verifyReleaseMetadata(lastRelease, tag, notes);
      requireCompleteAssets(lastRelease, expectedAssets);
      return lastRelease;
    }
    if (attempt < attempts) await wait(delayMs);
  }
  if (lastRelease?.draft === false) fail(`published GitHub Release ${tag} is not immutable`);
  fail(`GitHub Release ${tag} is still a draft after publication`);
}

/**
 * What: Reconcile or verify the hosted state for one already-created tag.
 * Why: A rerun after a lost API response must finish the same release without
 * choosing another version, moving its tag, or overwriting contradictory bytes.
 * How: Verify GitHub's tag first and again at every mutation boundary, classify
 * the complete hosted asset set before mutation, upload only absent draft
 * assets, publish once, and finish by re-reading both the immutable Release and
 * tag identity.
 */
export async function reconcileHostedRelease({
  mode,
  tag: rawTag,
  sourceSha: rawSourceSha,
  expectedAssets: rawExpectedAssets,
  notes = "",
  adapter,
  immutableAttempts = 10,
  immutableDelayMs = 500,
  wait = (milliseconds) => new Promise((resolvePromise) => setTimeout(resolvePromise, milliseconds)),
}) {
  if (mode !== "reconcile" && mode !== "verify") {
    fail(`unsupported release-state mode: ${mode}`);
  }
  if (!adapter || typeof adapter !== "object") fail("GitHub release adapter is required");
  const tag = canonicalStableTag(rawTag);
  const sourceSha = normalizeSourceSha(rawSourceSha, "tested source SHA");
  const expectedAssets = normalizeExpectedAssets(rawExpectedAssets);

  await verifyRemoteTagIdentity(adapter, tag, sourceSha);

  let release = await adapter.getRelease(tag);
  if (!release) {
    if (mode === "verify") fail(`GitHub Release ${tag} does not exist`);
    try {
      release = await adapter.createDraft({ tag, sourceSha, notes });
    } catch (error) {
      release = await recoverLostMutation(adapter, tag, error);
      if (!release) throw error;
    }
    await verifyRemoteTagIdentity(adapter, tag, sourceSha);
  }

  verifyReleaseMetadata(release, tag, notes);
  if (!release.draft) {
    return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
  }
  if (mode === "verify") fail(`GitHub Release ${tag} is still a draft`);

  // This full inspection happens before the first upload. A conflicting,
  // duplicate, or unexpected hosted asset therefore cannot cause a partial
  // repair mutation.
  let { missing } = inspectAssets(release, expectedAssets);
  for (const expected of missing) {
    release = await adapter.getRelease(tag);
    verifyReleaseMetadata(release, tag, notes);
    if (!release.draft) {
      return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
    }
    let current = inspectAssets(release, expectedAssets);
    if (!current.missing.some((asset) => asset.name === expected.name)) continue;

    await verifyRemoteTagIdentity(adapter, tag, sourceSha);
    // Re-read after the tag check so a concurrently published Release is
    // observed before this command attempts another asset mutation.
    release = await adapter.getRelease(tag);
    verifyReleaseMetadata(release, tag, notes);
    if (!release.draft) {
      return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
    }
    current = inspectAssets(release, expectedAssets);
    if (!current.missing.some((asset) => asset.name === expected.name)) continue;
    try {
      await adapter.uploadAsset(release, expected);
    } catch (error) {
      release = await recoverLostMutation(adapter, tag, error);
      if (!release) throw error;
      verifyReleaseMetadata(release, tag, notes);
      if (!release.draft) {
        return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
      }
      const refreshed = inspectAssets(release, expectedAssets);
      if (refreshed.missing.some((asset) => asset.name === expected.name)) throw error;
    }
  }

  release = await adapter.getRelease(tag);
  verifyReleaseMetadata(release, tag, notes);
  if (!release.draft) {
    return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
  }
  ({ missing } = inspectAssets(release, expectedAssets));
  if (missing.length > 0) {
    fail(`draft GitHub Release is still missing assets: ${missing.map((asset) => asset.name).join(", ")}`);
  }

  await verifyRemoteTagIdentity(adapter, tag, sourceSha);
  release = await adapter.getRelease(tag);
  verifyReleaseMetadata(release, tag, notes);
  if (!release.draft) {
    return verifyPublishedRelease({ adapter, release, tag, sourceSha, expectedAssets, notes });
  }
  requireCompleteAssets(release, expectedAssets);
  try {
    await adapter.publishRelease(release);
  } catch (_error) {
    // A lost response is not success. The authoritative immutable-state query
    // below decides whether GitHub applied the mutation.
  }
  const published = await waitForImmutable({
    adapter,
    tag,
    expectedAssets,
    notes,
    attempts: immutableAttempts,
    delayMs: immutableDelayMs,
    wait,
  });
  await verifyRemoteTagIdentity(adapter, tag, sourceSha);
  return { state: "published-immutable", release: published };
}

function bytesDigest(bytes) {
  return `sha256:${createHash("sha256").update(bytes).digest("hex")}`;
}

/**
 * Load and locally verify the hand-off produced by the artifact build. Paths
 * are resolved relative to the manifest so the CI workspace location never
 * becomes part of release identity.
 */
export function loadExpectedAssetManifest(manifestPath) {
  const absoluteManifest = resolve(manifestPath);
  let document;
  try {
    document = JSON.parse(readFileSync(absoluteManifest, "utf8"));
  } catch (error) {
    fail(`cannot read release asset manifest: ${error.message}`, { cause: error });
  }
  if (document?.schemaVersion !== 1) fail("release asset manifest schemaVersion must be 1");
  const tag = canonicalStableTag(document.tag);
  const sourceSha = normalizeSourceSha(document.sourceSha, "asset manifest source SHA");
  if (!Array.isArray(document.assets)) fail("release asset manifest assets must be an array");
  const root = dirname(absoluteManifest);
  const candidates = document.assets.map((asset) => {
    if (!asset || typeof asset.path !== "string" || asset.path.length === 0) {
      fail("every release asset manifest entry requires a relative path");
    }
    if (isAbsolute(asset.path)) fail("release asset manifest paths must be relative");
    const path = resolve(root, asset.path);
    const relativePath = relative(root, path);
    if (
      relativePath === ".."
      || relativePath.startsWith(`..${sep}`)
      || isAbsolute(relativePath)
    ) {
      fail(`release asset path escapes its manifest: ${asset.path}`);
    }
    let size;
    let digest;
    try {
      const fileType = lstatSync(path);
      if (!fileType.isFile() || fileType.isSymbolicLink()) {
        fail(`${asset.path} must be a regular file, not a symlink`);
      }
      const bytes = readFileSync(path);
      size = bytes.length;
      digest = bytesDigest(bytes);
    } catch (error) {
      fail(`cannot read release asset ${asset.path}: ${error.message}`, { cause: error });
    }
    if (size !== asset.size) fail(`${asset.name} local size does not match the asset manifest`);
    if (digest !== asset.digest) fail(`${asset.name} local digest does not match the asset manifest`);
    return { name: asset.name, path, size, digest };
  });
  return { tag, sourceSha, assets: normalizeExpectedAssets(candidates) };
}

/** GitHub API adapter used by the same-SHA release command. */
export class GitHubReleaseAdapter {
  constructor({
    repository,
    token,
    fetchImpl = globalThis.fetch,
    executeImpl = executeReleaseCommand,
    waitImpl = (delayMs) => new Promise((resolveWait) => setTimeout(resolveWait, delayMs)),
  }) {
    if (typeof repository !== "string" || !/^[A-Za-z0-9_.-]+\/[A-Za-z0-9_.-]+$/.test(repository)) {
      fail("repository must use OWNER/NAME form");
    }
    if (typeof token !== "string" || token.length === 0) fail("GitHub API token is required");
    if (typeof fetchImpl !== "function") fail("Fetch API is required");
    if (typeof executeImpl !== "function") fail("GitHub CLI executor is required");
    if (typeof waitImpl !== "function") fail("release visibility waiter is required");
    this.repository = repository;
    this.token = token;
    this.fetchImpl = fetchImpl;
    this.executeImpl = executeImpl;
    this.waitImpl = waitImpl;
    this.apiRoot = `https://api.github.com/repos/${repository}`;
  }

  async request(url, { method = "GET", body, headers = {}, allow404 = false } = {}) {
    const response = await this.fetchImpl(url, {
      method,
      body,
      headers: {
        Accept: "application/vnd.github+json",
        Authorization: `Bearer ${this.token}`,
        "X-GitHub-Api-Version": "2022-11-28",
        "User-Agent": "haxe-go-release-reconciliation",
        ...headers,
      },
    });
    if (allow404 && response.status === 404) return null;
    if (!response.ok) {
      const detail = (await response.text()).slice(0, 1000);
      throw new Error(`GitHub API ${method} ${url} failed (${response.status}): ${detail}`);
    }
    if (response.status === 204) return null;
    return response.json();
  }

  async getTagCommit(tag) {
    const encoded = encodeURIComponent(tag);
    const ref = await this.request(`${this.apiRoot}/git/ref/tags/${encoded}`, { allow404: true });
    if (!ref) return null;
    let object = ref.object;
    for (let depth = 0; depth < 8; depth += 1) {
      if (object?.type === "commit") return normalizeSourceSha(object.sha, `remote tag ${tag}`);
      if (object?.type !== "tag" || typeof object.sha !== "string") {
        fail(`remote tag ${tag} does not resolve to a Git commit`);
      }
      const tagObject = await this.request(`${this.apiRoot}/git/tags/${object.sha}`);
      object = tagObject.object;
    }
    fail(`remote tag ${tag} contains too many nested annotated tags`);
  }

  async getRelease(tag) {
    const encoded = encodeURIComponent(tag);
    const published = await this.request(`${this.apiRoot}/releases/tags/${encoded}`, { allow404: true });
    if (published) return published;
    // The tag endpoint can omit drafts. Authenticated bounded list queries are
    // therefore part of the authoritative lookup, as in the sibling compilers.
    for (let page = 1; page <= 10; page += 1) {
      const releases = await this.request(`${this.apiRoot}/releases?per_page=100&page=${page}`);
      if (!Array.isArray(releases)) fail("GitHub Releases list response must be an array");
      const match = releases.find((release) => release?.tag_name === tag);
      if (match) return match;
      if (releases.length < 100) return null;
    }
    fail(`GitHub Release lookup for ${tag} exceeded 1000 records`);
  }

  async createDraft({ tag, sourceSha, notes }) {
    canonicalStableTag(tag);
    normalizeSourceSha(sourceSha, "draft source SHA");
    const args = [
      "release",
      "create",
      tag,
      "--repo",
      this.repository,
      "--verify-tag",
      "--draft",
      "--title",
      tag,
    ];
    if (notes.length > 0) args.push("--notes-file", "-");
    else args.push("--generate-notes");
    this.executeImpl("gh", args, {
      env: { ...process.env, GH_TOKEN: this.token },
      input: notes,
    });
    // GitHub may acknowledge `gh release create` before the draft appears in
    // either REST lookup. This is eventual consistency, not contradictory
    // state. Re-read the authoritative endpoints with a short bounded backoff;
    // a later workflow run can still resume the same draft if all reads fail.
    const visibilityDelaysMs = [0, 1000, 2000, 4000, 5000, 5000, 5000, 5000];
    for (const delayMs of visibilityDelaysMs) {
      if (delayMs > 0) await this.waitImpl(delayMs);
      const release = await this.getRelease(tag);
      if (release) return release;
    }
    fail(`GitHub did not make the newly created draft Release ${tag} visible after bounded retries`);
  }

  async uploadAsset(release, expected) {
    if (!Number.isSafeInteger(release?.id)) fail("draft GitHub Release is missing its numeric id");
    const bytes = readFileSync(expected.path);
    if (bytes.length !== expected.size || bytesDigest(bytes) !== expected.digest) {
      fail(`${expected.name} changed after the release asset manifest was verified`);
    }
    const query = new URLSearchParams({ name: expected.name });
    return this.request(
      `https://uploads.github.com/repos/${this.repository}/releases/${release.id}/assets?${query}`,
      {
        method: "POST",
        body: bytes,
        headers: { "Content-Type": "application/octet-stream" },
      },
    );
  }

  async publishRelease(release) {
    if (!Number.isSafeInteger(release?.id)) fail("draft GitHub Release is missing its numeric id");
    return this.request(`${this.apiRoot}/releases/${release.id}`, {
      method: "PATCH",
      body: JSON.stringify({ draft: false }),
      headers: { "Content-Type": "application/json" },
    });
  }
}

function parseArguments(argv) {
  const options = { mode: "reconcile" };
  for (let index = 0; index < argv.length; index += 1) {
    const name = argv[index];
    if (name === "--verify-only") {
      options.mode = "verify";
      continue;
    }
    if (!["--repository", "--assets", "--notes-file"].includes(name)) {
      fail(`unknown argument: ${name}`);
    }
    const value = argv[index + 1];
    if (!value) fail(`missing value for ${name}`);
    index += 1;
    options[name.slice(2).replace("-", "_")] = value;
  }
  if (!options.repository) options.repository = process.env.GITHUB_REPOSITORY;
  if (!options.assets) fail("--assets <release-assets.json> is required");
  return options;
}

async function main() {
  if (process.argv.slice(2).length === 1 && process.argv[2] === "--list-stable-tags") {
    const tags = readFileSync(0, "utf8").split(/\r?\n/).filter(Boolean);
    process.stdout.write(`${canonicalStableTags(tags).join("\n")}\n`);
    return;
  }
  const options = parseArguments(process.argv.slice(2));
  const manifest = loadExpectedAssetManifest(options.assets);
  const notes = options.notes_file ? readFileSync(resolve(options.notes_file), "utf8") : "";
  const token = process.env.GH_TOKEN || process.env.GITHUB_TOKEN || "";
  const adapter = new GitHubReleaseAdapter({ repository: options.repository, token });
  const result = await reconcileHostedRelease({
    mode: options.mode,
    tag: manifest.tag,
    sourceSha: manifest.sourceSha,
    expectedAssets: manifest.assets,
    notes,
    adapter,
  });
  process.stdout.write(`${JSON.stringify({
    tag: manifest.tag,
    sourceSha: manifest.sourceSha,
    state: result.state,
    assetCount: manifest.assets.length,
  })}\n`);
}

if (process.argv[1] && resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  main().catch((error) => {
    const prefix = error instanceof ReleaseStateError ? "policy" : "error";
    console.error(`[release-reconciliation] ${prefix}: ${error.message}`);
    process.exitCode = 1;
  });
}
