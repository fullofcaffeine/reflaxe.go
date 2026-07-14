# Vendored Reflaxe Provenance

This page identifies the exact Reflaxe bytes shipped in vendor/reflaxe and
explains how to verify or update them.

## Recorded origin chain

The framework source starts from the official
[SomeRanDev/reflaxe](https://github.com/SomeRanDev/reflaxe) commit
430b4187a6bf4813cf618fc3a73ccf494a2ab9f5. That commit has 59 files under
src/reflaxe.

The shipped snapshot comes from
[fullofcaffeine/reflaxe.rust](https://github.com/fullofcaffeine/reflaxe.rust)
commit f53bec2adae8ef000467e488b974f6514c1af98f, path vendor/reflaxe. Its Git
tree is a26137d0af5f297eb12e3750a62d0544a4755b76.

The repository copy contains the supplier tree's 62 files plus one explicit
license-restoration overlay: the exact `LICENSE` bytes from the pinned official
upstream commit. The resulting 63-file tree has canonical SHA-256 digest
29437d3bddaefab9cfb1213b2b7911e758794048d535eb01c68f7c88fdaca454.
Removing the declared overlay recreates the supplier Git tree object exactly.
The canonical digest is SHA-256 over path-sorted records of each file's SHA-256
and POSIX-relative path; the overlay declaration and every file digest live in
[provenance/reflaxe/vendor-manifest.json](../provenance/reflaxe/vendor-manifest.json).

The patch from official source to supplier source contains 14 modified framework source files.
It is committed as
[provenance/reflaxe/upstream-to-supplier.patch](../provenance/reflaxe/upstream-to-supplier.patch)
and has its own digest in the manifest. The OutputManager.hx portion uses
Git's binary-patch encoding because the inherited source contains
absolute-path-shaped examples; this keeps the provenance artifact free of
machine-local-looking paths without changing reconstructed bytes.

The official upstream LICENSE file digest is recorded in the manifest and is
the authority for the shipped `vendor/reflaxe/LICENSE` overlay.
This technical provenance record does not decide redistribution obligations;
license and release-distribution review remains a separate policy decision.

vendor/reflaxe/PATCHES.md and FUTURE_MODIFICATIONS.md are bytes inherited from
the supplier snapshot. They are not the authority for the origin identity
recorded here.

## Offline verification

Run:

~~~bash
npm run verify:vendor-reflaxe
~~~

The verifier:

1. rejects missing, added, changed, or symlinked vendor entries;
2. verifies all 63 file hashes, the canonical tree digest, the license overlay,
   and the supplier Git tree after removing only that declared overlay;
3. verifies the patch digest and its 14-file inventory;
4. reverse-applies the patch to the shipped source and checks the exact official
   source digest; and
5. reapplies the patch and checks that the exact supplier source is restored.

This path requires only committed files and Git. It does not access the
network or modify vendor/reflaxe.

## Pinned-network reconstruction

Run:

~~~bash
npm run verify:vendor-reflaxe:reconstruct
~~~

In addition to the offline checks, the verifier creates isolated temporary
repositories, fetches both immutable commits, verifies the supplier Git tree
and official source/license digests, applies the committed patch to official
source, and compares the result byte-for-byte with the supplier source. Branch
heads and mutable release tags are never reconstruction inputs.

## Vendor update procedure

1. Select and review an exact official upstream commit and an exact supplier
   commit. Record commit SHAs, never branch names.
2. Replace vendor/reflaxe only with the exact supplier subdirectory and record
   its Git tree SHA-1.
3. Regenerate the official-to-supplier source patch in a temporary Git
   repository. Use -c diff.mnemonicPrefix=false diff --binary --full-index so
   paths are normalized as a/src/reflaxe/... and b/src/reflaxe/.... Mark a
   file binary in the temporary repository's .git/info/attributes when its
   literal source would put a machine-local path shape into the patch.
4. Recompute the path-sorted SHA-256 inventory, source-tree digests, patch
   digest, exact modified-file list, upstream license digest, and both commit
   identities in vendor-manifest.json.
5. Run both verifier modes and the regression contract:

   ~~~bash
   npm run verify:vendor-reflaxe
   npm run verify:vendor-reflaxe:reconstruct
   python3 test/test_vendor_reflaxe_provenance.py
   ~~~

6. Review the vendor and patch diffs, then perform the separate license and
   distribution-policy review required for a release.

Any mismatch is a failed update, not a warning.
