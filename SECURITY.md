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

You can run a full local secret scan with:

```bash
npm run security:gitleaks
```
