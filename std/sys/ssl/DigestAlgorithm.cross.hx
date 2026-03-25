package sys.ssl;

/**
	What
	Direct `sys.ssl.DigestAlgorithm` support for `haxe.go`.

	Why
	- `sys.ssl.Digest` selects hashing/signing behavior through this enum
	  abstract.
	- Keeping the upstream names unchanged preserves callsite portability even
	  while the underlying digest implementation is target-specific.

	How
	- Match the upstream string-backed enum abstract values exactly so generated
	  Go and Haxe-level comparisons see the same public API.
**/
enum abstract DigestAlgorithm(String) to String {
	var MD5 = "MD5";
	var SHA1 = "SHA1";
	var SHA224 = "SHA224";
	var SHA256 = "SHA256";
	var SHA384 = "SHA384";
	var SHA512 = "SHA512";
	var RIPEMD160 = "RIPEMD160";
}
