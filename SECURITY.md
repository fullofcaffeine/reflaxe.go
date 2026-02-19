# Security Policy

## Reporting

If you discover a security issue, avoid filing a public issue with exploit details.
Open a private security advisory on GitHub or contact the maintainers directly.

## Repository Protections

This repo enforces secret scanning in two places:

- Local pre-commit hook (`npm run hooks:install`)
  - staged local path leak guard
  - staged `gitleaks` scan
  - staged Haxe auto-format
- CI workflow (`.github/workflows/security-gitleaks.yml`)
  - runs `gitleaks` on PRs and pushes to `master`

It also enforces static security analysis in CI:

- CI workflow (`.github/workflows/security-static-analysis.yml`)
  - `CodeQL` analysis (`go`, `python`, `javascript-typescript`)
  - dependency review for PRs (`actions/dependency-review-action`)
  - dependency vulnerability audit (`npm audit` + `govulncheck`)

You can run a full local secret scan with:

```bash
npm run security:gitleaks
```

You can run the local dependency vulnerability audit with:

```bash
npm run security:deps
```

`security:deps` behavior knobs:

- `SKIP_GOVULNCHECK=1` skip Go vulnerability scanning entirely.
- `GOVULNCHECK_INSTALL_ATTEMPTS` retry count for installing `govulncheck` (default `3`).
- `GOVULNCHECK_INSTALL_RETRY_DELAY_SEC` base backoff seconds between retries (default `2`).
- `GOVULNCHECK_VERSION` pin install target version (default `latest`).
- `GOVULNCHECK_ALLOW_INSTALL_FAILURE=1` local soft-fail mode if install still fails after retries.

In CI (`CI=true`), govulncheck install failure remains a hard failure.
