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

The dependency vulnerability audit copies both package.json and
package-lock.json into an isolated directory and performs npm ci
--ignore-scripts --omit=dev. It never resolves a new lock during an audit.

## Updating JavaScript dependencies

package.json records the lock generator as npm@11.16.0 in its packageManager
field. Use that exact npm version on the supported Node 24 line:

~~~bash
npm_version="$(node -p "require('./package.json').packageManager.split('@').pop()")"
npx --yes --package="npm@$npm_version" --call \
  'npm install --package-lock-only --ignore-scripts --no-audit --no-fund'
npm ci --ignore-scripts --no-audit --no-fund
npm run security:supply-chain
npm run security:deps
~~~

Change declared dependencies intentionally in package.json, inspect the full
lock diff, and commit both files together. Do not hand-edit the resolved
packages, and do not replace npm ci with npm install in CI.

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
