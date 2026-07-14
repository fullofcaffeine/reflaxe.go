# GPT-5.6 Pro review evidence

## What it is

The review bundle is a reproducible evidence snapshot for an independent architecture and production-readiness review. It names the exact Git commit under review and carries both a direct Git archive and a line-numbered Repomix view. The archive is the source authority; Repomix is a reviewer convenience.

## Why it exists

A deep review cannot safely infer source identity, generated-output freshness, CI status, release state, or sibling precedent from an arbitrary working tree. The bundle keeps those authorities separate and explicit:

- source and committed generated portable and metal output come from one pushed Haxe.Go commit;
- GitHub Actions logs and job metadata must report that exact source SHA;
- release and tag state comes from the live GitHub API;
- compatibility and ownership inventories stay connected to their executable generators and fixtures;
- Rust-family portable specialization and release patterns are read from a pinned Haxe.Rust commit;
- Ruby release mechanics and Elixir's canonical `_std` layout are pinned as narrower sibling references.

The sibling material is precedent, not code imported into Haxe.Go and not proof that a Rust-, Ruby-, or Elixir-specific mechanism belongs in Go.

## How it works

The builder:

1. resolves every requested ref to a full commit and requires each commit to be present on an `origin/*` branch;
2. creates the primary source archive with `git archive`, independently of uncommitted files;
3. emits source and reference file inventories with Git blob IDs;
4. creates line-numbered views with pinned `repomix@1.14.0` without disabling its security scan;
5. requires every selected GitHub Actions run to be completed successfully at the exact primary SHA;
6. captures sanitized GitHub Actions logs, release/tag metadata, host-control responses, and a read-only roadmap snapshot;
7. rejects tracked secret-bearing path classes and unredacted machine-local paths;
8. runs Gitleaks, writes a per-file SHA-256 manifest, and creates a deterministic outer ZIP.

Run it only after the review commit and its CI are pushed:

```bash
python3 scripts/review/build_gpt56_evidence.py \
  --source-ref <full-haxe-go-sha> \
  --rust-ref <full-haxe-rust-sha> \
  --ruby-ref <full-haxe-ruby-sha> \
  --elixir-ref <full-haxe-elixir-sha> \
  --ci-run <exact-sha-run-id> \
  --release-tag <current-release-tag> \
  --output dist/review/haxe-go-gpt56-evidence-<short-sha>.zip
```

Repeat `--ci-run` for the quality, compiler harness, generated-example, static-analysis, and secret-scanning workflows. The committed review record must name the ZIP, size, outer SHA-256, source commit, sibling commits, run IDs, release tag, and intentional omissions.

## Intentional exclusions

- `.git`, `node_modules`, caches, ignored build output, `dist`, and all untracked files are outside `git archive`.
- `.beads/issues.jsonl` is the hashed legacy provenance archive. It is not current compiler or roadmap evidence and is intentionally omitted.
- `.beads/interactions.jsonl` is omitted; a read-only operational roadmap/Dolt snapshot is included separately.
- `*.pem`, `*.key`, and `infra/secrets/**` cause bundle construction to fail.
- Machine-local home/workspace paths in captured logs are replaced with stable placeholders. Repository-relative file and line evidence is preserved.
- Existing untracked Repomix files are never reused because their source commit and exclusion set are not independently provable.

If Repomix's heuristic scanner omits a tracked test fixture, the builder records the exact path and copies the raw file into a separate `repomix-security-exclusions` payload after the full bundle passes Gitleaks. Reviewers must treat that as an upload-view limitation, not infer that the source file is missing.

The resulting ZIP lives under ignored `dist/review/`; its small, auditable evidence record belongs in this directory. The review prompt and final adjudication will be committed alongside that record in later Beads slices.
