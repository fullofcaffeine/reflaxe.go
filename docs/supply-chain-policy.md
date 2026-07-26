# Supply-Chain Policy

This policy defines what the repository proves about JavaScript dependencies,
GitHub Actions, and the vendor Reflaxe snapshot before a release.

## What the gate protects

The release-blocking gate has three parts:

1. package-lock.json is the only JavaScript resolution input. CI installs it
   with npm ci, which fails when package.json and the lock disagree.
2. Every external GitHub Action is pinned to a 40-character commit SHA. The
   readable release tag and the same immutable commit are recorded in
   [.github/actions-lock.json](../.github/actions-lock.json).
3. The vendor Reflaxe tree is checked against a per-file manifest, an exact
   supplier Git tree, and a patch that round-trips to an exact official
   upstream commit.

Local composite actions such as ./.github/actions/setup-haxe-linux are
repository source and therefore do not use an external action pin.

## Why it exists

A mutable package resolution or action tag can change without a source commit.
Likewise, an unlabelled vendor directory cannot prove which upstream bytes a
release contains. These checks make those inputs reviewable, repeatable, and
fail closed when they drift.

## How it works

Run the complete offline gate from repository root:

~~~bash
npm run security:supply-chain
~~~

The command checks root metadata in package-lock.json, scans all workflow and
composite-action YAML for immutable pins and version comments, confirms weekly
Dependabot coverage, then runs the vendored-source verifier. The CI quality job
runs the command immediately after its clean install, and npm run
release:status runs it again before reporting healthy release wiring.

The dependency vulnerability audit copies package.json, package-lock.json, and
every declared repository-local `file:` dependency into an isolated directory.
It performs `npm ci --ignore-scripts --include=dev`, then requires `npm ls
--all` to prove the isolated install has no missing or invalid package links
before running the vulnerability scan. This is intentional: npm classifies the
repository's CI and release executables as development dependencies, but they
are operational supply-chain dependencies because they decide versions,
generate release notes, and publish releases. The audit never resolves a new
lock or treats a malformed install as clean.

## Updating JavaScript dependencies

`package.json` is the canonical npm version source through its exact
`packageManager` value, currently `npm@11.16.0`. That metadata tells people and
tools which version owns the lock, but it does not activate that npm version by
itself. Node's
[Corepack documentation](https://nodejs.org/docs/latest-v24.x/api/corepack.html)
explains that npm has no Corepack shim by default, so an ordinary `npm` command
otherwise uses the executable bundled with Node.

Every CI job that sets up Node immediately runs
`scripts/ci/setup-pinned-npm.sh`. The script reads `packageManager`, uses the
bundled npm only to install that exact global npm package with lifecycle scripts
and automatic audit disabled, clears the shell command cache, and fails unless
`npm --version` matches. The supply-chain gate rejects workflow npm invocations
that appear before this bootstrap. As a result, clean installs, dependency
audits, tests, performance jobs, and the release job all execute through the
same reviewed npm version.

For a local lock refresh, invoke the same exact version without changing the
developer's global npm:

~~~bash
npm_version="$(node -p "require('./package.json').packageManager.split('@').pop()")"
npx --yes --package="npm@$npm_version" --call \
  'npm install --package-lock-only --ignore-scripts --no-audit --no-fund'
npx --yes --package="npm@$npm_version" --call \
  'npm ci --ignore-scripts --no-audit --no-fund'
npm run security:supply-chain
npm run security:deps
~~~

Change declared dependencies intentionally in package.json, inspect the full
lock diff, and commit both files together. Do not hand-edit the resolved
packages, and do not replace npm ci with npm install in CI. The isolated
dependency audit additionally requires `npm ls --all`, so activating the exact
npm version cannot make a broken dependency tree appear clean.

## Updating GitHub Actions

Dependabot checks the github-actions ecosystem weekly. For each update:

1. Resolve the exact release tag in the action's official repository. For an
   annotated tag, use the dereferenced ^{} result; for a lightweight tag, use
   the tag result.
2. Update every workflow reference to the resolved 40-character commit and
   retain the exact release tag as its inline comment.
3. Update the matching entry in .github/actions-lock.json.
4. Run npm run security:supply-chain and
   python3 test/test_supply_chain_contract.py.

For example:

~~~bash
git ls-remote https://github.com/actions/checkout.git \
  'refs/tags/v6.0.3' 'refs/tags/v6.0.3^{}'
~~~

The lock manifest is intentionally independent of a moving tag. Dependabot is
update automation, not authority to bypass review of the resolved commit.

## Updating vendored Reflaxe

Follow the exact origin, digest, patch, and reconstruction procedure in
[Vendored Reflaxe Provenance](vendor-reflaxe-provenance.md). A vendor update is
not complete until both the offline verification and pinned-network
reconstruction pass.

## Failure policy

Missing locks, stale root metadata, mutable or unmanifested actions, unexpected
vendor files, digest drift, and failed patch reconstruction all return a
non-zero status. There is no release-mode skip.
