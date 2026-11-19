import haxe.crypto.Sha256;
import haxe.io.Bytes;
import reflaxe.go.compiler.GoContentDigestMemo;

/**
	Contract for the existing-module transaction digest memo.

	Why
	The memo lets one output transaction skip repeated SHA-256 passes over the
	same generated files. A reused digest is only correct when the bytes being
	checked are exactly the bytes that produced it.

	What
	Proves that the memo returns real SHA-256 digests, reuses a known digest only
	for byte-identical content, and cannot be corrupted by later edits to a buffer
	that was handed to it.

	How
	A deliberately wrong seeded digest makes reuse observable without counters:
	it may come back only for identical bytes at the seeded path.
**/
class GoContentDigestMemoContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function sha256(bytes:Bytes):String {
		return Sha256.make(bytes).toHex();
	}

	static function copy(bytes:Bytes):Bytes {
		return bytes.sub(0, bytes.length);
	}

	static function main():Void {
		final path = "generated/go/app/main.go";
		final original = Bytes.ofString("package main\n\nfunc main() {}\n");

		final fresh = new GoContentDigestMemo();
		require(fresh.digest(path, original) == sha256(original), "a fresh digest is not the SHA-256 of the bytes");
		require(fresh.digest(path, copy(original)) == sha256(original), "identical bytes read again changed the digest");

		final seeded = new GoContentDigestMemo();
		final marker = "seeded-digest-for-reuse-check";
		seeded.seed(path, original, marker);
		require(seeded.digest(path, copy(original)) == marker, "byte-identical content did not reuse the known digest");

		final sameLength = Bytes.ofString("package main\n\nfunc mainX(){}\n");
		require(sameLength.length == original.length, "the same-length fixture has a different length");
		require(seeded.digest(path, sameLength) == sha256(sameLength), "different bytes of equal length reused a digest");

		final longer = Bytes.ofString("package main\n\nfunc main() {}\n// edited\n");
		require(seeded.digest(path, longer) == sha256(longer), "different bytes reused a digest");

		final buffer = copy(original);
		final guarded = new GoContentDigestMemo();
		guarded.seed(path, buffer, marker);
		buffer.set(0, "P".code);
		require(guarded.digest(path, buffer) == sha256(buffer), "editing a seeded buffer changed the remembered content");
		require(guarded.digest(path, copy(original)) == marker, "the remembered content did not survive a buffer edit");

		Sys.println("ok");
	}
}
