# GitHub Governance Policy

## What this protects

GitHub stores important security and release controls outside Git. A clean
workflow file therefore cannot prove that branch rules, tag rules, secret
scanning, or immutable releases are active. `github-governance-policy.json` is
the repository-owned declaration of those expected host settings, and
`scripts/security/verify-github-governance.py` compares the declaration with
both source files and the live GitHub API.

The policy requires:

- pull requests from ordinary contributors to pass the declared CI checks,
  receive one approval, use linear history, and resolve review conversations;
- deletion and force-push protection for `master` and `v*` release tags;
- [immutable releases](https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases),
  which lock a published release's tag to its commit while that release exists;
- [secret scanning and push protection](https://docs.github.com/en/code-security/concepts/secret-security/push-protection);
- vulnerability alerts and Dependabot security updates;
- read-only default `GITHUB_TOKEN` permissions, no workflow-authored pull
  request approvals, and host enforcement of full action commit SHAs.

These controls complement the repository's Gitleaks, dependency audits,
action-lock manifest, deterministic artifact checks, and same-tested-SHA
release workflow. No single control replaces the others.

`allowed_actions` remains `all` intentionally. The host requires every external
action reference to use a full commit SHA, while the source verifier and
`.github/actions-lock.json` require each referenced repository, human-readable
version, and reviewed SHA to agree. This permits a future reviewed action
addition through an ordinary policy change without maintaining a second
repository allowlist in mutable host settings.

GitHub reports non-provider secret patterns and secret validity checks as
disabled after enabling the repository's available security controls. Those
features belong to GitHub Secret Protection, which is not available for this
personal public repository. The policy records both controls under
`plan_limitations` with their observed status. The live verifier will fail if
that status changes, prompting a deliberate policy upgrade rather than silently
ignoring a newly available protection.

## Why administrators can bypass

This personal repository currently has one maintainer and its reviewed
automation lands work through direct administrator pushes to `master`.
Requiring a second human review with no bypass would make legitimate
maintenance and emergency recovery impossible. Both rulesets therefore name
the repository owner as an explicit, audited `always` bypass actor.

That bypass is a recovery and sole-maintainer compatibility path, not evidence
that a change passed CI. Release publication still runs only from the exact
`master` SHA after its workflow dependencies pass. If the project gains another
maintainer and adopts pull-request-only landing, replace the owner bypass with
the narrower governance model in the same reviewed policy change.

Immutable releases deliberately add a stronger boundary: after publication,
the associated tag cannot move or disappear while the release exists, even
though an administrator can bypass the general tag ruleset. Repair must follow
the documented release reconciliation process rather than silently retargeting
a published version.

## How to verify

Source-only verification is deterministic and needs no GitHub credentials:

~~~bash
python3 scripts/security/verify-github-governance.py --mode source
python3 test/test_github_governance_policy.py
~~~

Live verification requires `gh auth` with repository administration read
access:

~~~bash
python3 scripts/security/verify-github-governance.py --mode live
~~~

The live command fails if a declared setting is unavailable, disabled, renamed,
or weakened. It does not mutate GitHub. Change the JSON policy and its tests
first, review the effect, apply the matching host change through GitHub's API or
settings UI, and run the live verifier again.

GitHub's repository-rules API and Actions-permissions API are the canonical
references for the fields represented here:

- [Repository rulesets API](https://docs.github.com/en/rest/repos/rules)
- [GitHub Actions permissions API](https://docs.github.com/en/rest/actions/permissions)
- [Repository immutable-releases API](https://docs.github.com/en/rest/repos/repos#check-if-immutable-releases-are-enabled-for-a-repository)

Dependabot version-update scheduling remains in `.github/dependabot.yml`.
Dependabot security updates are a separate live repository setting; the policy
and live verifier require both rather than treating the YAML file as proof that
security updates are enabled.
