# GitHub Governance Evidence — `haxe_go-vfp.4.10`

Captured on 2026-07-26 at 2026-07-26T21:28:24Z for
`fullofcaffeine/reflaxe.go`. This record separates the observed live API state
from the policy and verifier committed with the implementation.

## Before

The authenticated GitHub REST API reported:

- no repository rulesets and no classic protection on `master`;
- immutable releases disabled;
- secret scanning, push protection, and Dependabot security updates disabled;
- Actions enabled with `allowed_actions: all` and
  `sha_pinning_required: false`;
- default workflow permissions already `read`, with workflow pull-request
  approval disabled.

The pre-change calls were:

~~~text
GET /repos/fullofcaffeine/reflaxe.go
GET /repos/fullofcaffeine/reflaxe.go/rulesets
GET /repos/fullofcaffeine/reflaxe.go/branches/master/protection
GET /repos/fullofcaffeine/reflaxe.go/immutable-releases
GET /repos/fullofcaffeine/reflaxe.go/actions/permissions
GET /repos/fullofcaffeine/reflaxe.go/actions/permissions/workflow
~~~

## Applied state

GitHub accepted and returned:

| Control | Verified state |
| --- | --- |
| `Protect master` | Active branch ruleset, ID `19778148` |
| `Protect release tags` | Active tag ruleset, ID `19778147` |
| Immutable releases | Enabled |
| Secret scanning | Enabled |
| Secret scanning push protection | Enabled |
| Dependabot vulnerability alerts and security updates | Enabled |
| Actions | Enabled; `allowed_actions: all`; `sha_pinning_required: true` |
| Default workflow token | Read-only; cannot approve pull requests |

The master ruleset requires the seven status checks declared in
`github-governance-policy.json`, linear history, one non-author approval,
resolved conversations, and blocks deletion and force pushes. The release-tag
ruleset blocks deletion and tag movement for `v*`. Both name the personal
repository owner as the explicit administrator recovery bypass.

The final fail-closed verification passed:

~~~text
$ python3 scripts/security/verify-github-governance.py --mode live
[github-governance] source: OK (5 workflows; 2 rulesets declared)
[github-governance] live: OK (fullofcaffeine/reflaxe.go; 2 active rulesets)
~~~

## Plan limitations

GitHub left `secret_scanning_non_provider_patterns` and
`secret_scanning_validity_checks` disabled after the available repository
controls were enabled. Both are GitHub Secret Protection features unavailable
to this personal public repository. The policy records their observed state and
the live verifier will fail if availability changes, prompting an intentional
upgrade.

## Recovery boundary

The owner bypass preserves the established direct-push and emergency-recovery
workflow. It does not count as CI or release evidence. Published releases have
the stronger immutable-release boundary: their tags remain locked while the
release exists, and repair follows `docs/release-reconciliation.md`.
