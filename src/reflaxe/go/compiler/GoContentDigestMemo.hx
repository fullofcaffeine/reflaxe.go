package reflaxe.go.compiler;

#if (macro || eval)
import haxe.io.Bytes;

/** One exact byte sequence and the SHA-256 digest computed from it. */
private typedef GoContentDigestEntry = {
	final bytes:Bytes;
	final sha256:String;
}

/**
	Remembers SHA-256 digests of exact file contents for one output transaction.

	Why
	An existing-module commit checks ownership by hashing the same files several
	times: old owned files before staging, replaced destinations, backups, and the
	newly installed files. The compiler hashes in interpreted Haxe, so on a large
	program these repeated passes cost more time than formatting every generated
	file with `gofmt`.

	What
	`digest` always returns `GoExistingModuleOwnership.digest(bytes)`. A known
	digest is reused only when the bytes are byte-for-byte equal to the bytes that
	produced it, so every check still inspects the bytes it read at that moment.
	A changed, missing, or replaced file still fails exactly as before.

	How
	Entries are grouped by module-relative path. A lookup compares the candidate
	with each remembered entry of the same length through native `Bytes.compare`.
	A miss hashes the bytes and remembers a private copy, so later edits to a
	caller's buffer cannot change what an entry describes. `seed` records a digest
	that the output plan already computed from exactly those bytes.
**/
class GoContentDigestMemo {
	final entriesByPath:Map<String, Array<GoContentDigestEntry>> = [];

	public function new() {}

	/** Record a digest that was computed from exactly `bytes`. */
	public function seed(path:String, bytes:Bytes, sha256:String):Void {
		remember(path, bytes, sha256);
	}

	/** Return the SHA-256 hex digest of `bytes`, hashing only unseen content. */
	public function digest(path:String, bytes:Bytes):String {
		final entries = entriesByPath.get(path);
		if (entries != null) {
			for (entry in entries) {
				if (entry.bytes.length == bytes.length && entry.bytes.compare(bytes) == 0) {
					return entry.sha256;
				}
			}
		}
		final sha256 = GoExistingModuleOwnership.digest(bytes);
		remember(path, bytes, sha256);
		return sha256;
	}

	function remember(path:String, bytes:Bytes, sha256:String):Void {
		final entry:GoContentDigestEntry = {bytes: bytes.sub(0, bytes.length), sha256: sha256};
		final entries = entriesByPath.get(path);
		if (entries == null) {
			entriesByPath.set(path, [entry]);
		} else {
			entries.push(entry);
		}
	}
}
#else
class GoContentDigestMemo {}
#end
