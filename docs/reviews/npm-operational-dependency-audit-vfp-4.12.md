# Operational npm Dependency Audit: `haxe_go-vfp.4.12`

Date: 2026-07-14

Baseline source commit: `c61afc71`

## Decision

Audit the complete locked Node dependency tree, including packages declared as
`devDependencies`. In this repository, that npm category contains the tools
that analyze commits, choose release versions, generate notes, and publish
GitHub releases. They are operational supply-chain dependencies even though
they are not shipped in the Haxelib archive.

The release-blocking command is:

~~~bash
npm audit --include=dev --audit-level=high
~~~

The explicit `--include=dev` prevents `NODE_ENV=production` from silently
excluding release tooling. The audit reports every severity and exits nonzero
for high or critical findings. No advisory ignore, package override, or
reachability exception is configured.

## Baseline And Reachability

The baseline locked-tree audit under npm 11.16.0 reported 19 vulnerable package
findings: 1 critical, 10 high, 7 moderate, and 1 low. The following execution
map distinguishes the configured execution path from packages that npm
installed but the release configuration does not invoke:

1. `.releaserc.json` configures only the repository's custom
   `scripts/release/analyze-commits.mjs` plugin and
   `@semantic-release/github`.
2. The custom plugin imports `@semantic-release/commit-analyzer` and
   `@semantic-release/release-notes-generator`. Handlebars, lodash-es, and
   picomatch therefore occur on configured analyzer, notes, or GitHub paths.
3. `semantic-release` bundles `@semantic-release/npm` as a default plugin
   dependency, but the repository's explicit plugin list does not select it.
   Its nested npm CLI and packaging/signing subtree is installed but inactive
   for this release architecture.
4. `js-yaml` is installed through cosmiconfig, but `.releaserc.json` selects
   the JSON loader rather than the YAML adapter.

Installed-but-inactive paths were still upgraded. Their classification
explains exposure; it does not exempt bytes in the operational lock from the
gate.

## Complete 19-Finding Triage

The “path” column describes the baseline execution classification. “Fixed
lock” records the old vulnerable version or versions and the clean resolved
version or versions.

| Package | Severity | Baseline path | Fixed lock |
| --- | --- | --- | --- |
| `@sigstore/core` | `moderate` | Installed but inactive under bundled `@semantic-release/npm` signing support. | `2.0.0` -> `3.2.1` |
| `@sigstore/sign` | `moderate` | Installed but inactive under bundled `@semantic-release/npm` signing support. | `3.1.0` -> `4.1.1` |
| `@sigstore/verify` | `moderate` | Installed but inactive under bundled `@semantic-release/npm` signing support. | `2.1.1` -> `3.1.1` |
| `brace-expansion` | `moderate` | Installed but inactive under the bundled npm CLI glob stack. | `2.0.2` -> `5.0.7` |
| `diff` | `low` | Installed but inactive under the bundled npm CLI diff path. | `5.2.0` -> `8.0.4` |
| `glob` | `high` | Installed but inactive under the bundled npm CLI. | `10.4.5` -> `13.0.6` |
| `handlebars` | `critical` | Configured execution path through `conventional-changelog-writer` for analysis and notes. | `4.7.8` -> `4.7.9` |
| `ip-address` | `moderate` | Installed but inactive under the bundled npm CLI proxy stack. | `9.0.5` -> `10.2.0` |
| `js-yaml` | `moderate` | Installed config adapter; inactive because the repository uses `.releaserc.json`. | `4.1.1` -> `4.3.0` |
| `libnpmdiff` | `high` | Installed but inactive under the bundled npm CLI. | `7.0.1` -> `8.1.11` |
| `libnpmpublish` | `high` | Installed but inactive under bundled `@semantic-release/npm`. | `10.0.1` -> `11.2.0` |
| `lodash-es` | `high` | Configured execution path through semantic-release, analyzer, notes, and GitHub plugins. | `4.17.23` -> `4.18.1` |
| `minimatch` | `high` | Installed but inactive under the bundled npm CLI. | `9.0.5` -> `10.2.5` |
| `npm` | `high` | Bundled by inactive `@semantic-release/npm`; not the npm 11.16.0 process running the gate. | `10.9.4` -> `11.18.0` |
| `pacote` | `high` | Installed but inactive under the bundled npm CLI fetch/package paths. | `19.0.1`, `20.0.0` -> `21.5.1` |
| `picomatch` | `high` | Mixed: configured analyzer/GitHub paths plus the inactive bundled npm CLI path. | `2.3.1`, `4.0.2`, `4.0.3` -> `2.3.2`, `4.0.4`, `4.0.5` |
| `sigstore` | `high` | Installed but inactive under bundled npm publish and package-fetch support. | `3.1.0` -> `4.1.1` |
| `socks` | `moderate` | Installed but inactive under the bundled npm CLI proxy stack. | `2.8.5` -> `2.8.9` |
| `tar` | `high` | Installed but inactive under bundled npm packaging, cache, and node-gyp paths. | `6.2.1`, `7.4.3` -> `7.5.19` |

The critical Handlebars range ends at 4.7.8 and is patched in 4.7.9; the
lodash-es high-severity range ends at 4.17.23 and is patched in 4.18.0. The
lock resolves later fixed versions in both cases.

## Remediation

The supported Node 24 release toolchain now uses:

- `semantic-release` 25.0.7 instead of 24.2.9;
- `@semantic-release/github` 12.0.9 instead of 11.0.6;
- `@semantic-release/release-notes-generator` 14.1.1 instead of 14.1.0;
- the compatible patched transitive versions recorded above; and
- the existing lock generator pin, npm 11.16.0.

The semantic-release major upgrade moves its bundled inactive npm plugin from
12.0.2 to 13.1.5 and its nested npm CLI from 10.9.4 to 11.18.0. This removes
the inactive subtree findings without pretending they are reachable. It also
keeps the tree clean if release configuration changes later.

## Verification

Verification uses Node 24.18.0 and npm 11.16.0:

~~~text
npm ci --ignore-scripts --no-audit --no-fund
npm audit --include=dev --audit-level=high
found 0 vulnerabilities
~~~

Repeated npm 11.16.0 package-lock-only resolution leaves `package-lock.json`
byte-for-byte unchanged. The release-policy suite also exercises the real
semantic-release engine against temporary Git history, proving tag-derived
version selection and generated notes after the major upgrade.

## Residual Policy

- Any future high or critical finding anywhere in the operational lock blocks
  CI and release, whether the path is configured or merely installed.
- Moderate and low findings remain visible in the report and require triage;
  `--audit-level=high` changes the exit threshold, not report contents.
- A reachability disposition must name the exact configured plugin or loader
  boundary. It cannot replace an available compatible upgrade.
- Changes to `.releaserc.json` must re-evaluate paths currently classified as
  installed but inactive.
