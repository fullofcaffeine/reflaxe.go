import haxe.crypto.Sha256;
import haxe.io.Bytes;
import reflaxe.go.compiler.GoExistingModuleOwnership;

/**
	Contract for the SHA-256 digest behind existing-module ownership records.

	Why
	Ownership records store a SHA-256 hex digest for every compiler-owned file.
	A rerun decides whether a file is still compiler-owned, caller-owned, or
	modified by comparing the digest of the bytes on disk with the record. Any
	digest that differs from standard SHA-256 would misjudge every owned file
	and break records written by earlier compiler versions.

	What
	Pins `GoExistingModuleOwnership.digest` to standard SHA-256: the FIPS 180-2
	test messages, a one-million-byte input, multi-byte UTF-8 source text, and
	inputs of every byte value whose lengths sit on both sides of each padding
	boundary (55/56 and 63/64 bytes per block).

	How
	The pinned hex values come from Python's `hashlib.sha256`, so the contract
	does not trust any Haxe implementation. Every length from 0 to 300 bytes is
	also compared with the Haxe standard library's `haxe.crypto.Sha256`.
**/
class GoOwnershipDigestContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function pattern(length:Int):Bytes {
		final bytes = Bytes.alloc(length);
		for (index in 0...length) {
			bytes.set(index, (index * 167 + 13) & 0xFF);
		}
		return bytes;
	}

	static function expectDigest(bytes:Bytes, expected:String, label:String):Void {
		final actual = GoExistingModuleOwnership.digest(bytes);
		require(actual == expected, 'digest of $label is $actual, expected $expected');
	}

	static function main():Void {
		expectDigest(Bytes.alloc(0), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "the empty input");
		expectDigest(Bytes.ofString("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "abc");
		expectDigest(Bytes.ofString("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"),
			"248d6a61d20638b8e5c026930c3e6039a33ce45964ff2167f6ecedd419db06c1", "the 448-bit FIPS message");
		expectDigest(Bytes.ofString("abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu"),
			"cf5b16a778af8380036ce59e7b0492370b249b11e8f07a51afac45037afee9d1", "the 896-bit FIPS message");

		final million = Bytes.alloc(1000000);
		million.fill(0, million.length, "a".code);
		expectDigest(million, "cdc76e5c9914fb9281a1c7e284d73e67f1809a48a497200e046d39ccc7112cd0", "one million 'a' bytes");

		final source = Bytes.ofString("package main\n\n// Ünïcödé comment: 日本語 ✓\nfunc main() {\n\tprintln(\"héllo\")\n}\n");
		require(source.length == 87, "the UTF-8 source fixture has an unexpected length");
		expectDigest(source, "87f81e164777cd6dd1fc8a429f482c0235c3588db5028a15410dbfb6f4125e11", "UTF-8 Go source");

		final pinned:Array<{length:Int, digest:String}> = [
			{length: 1, digest: "9d1e0e2d9459d06523ad13e28a4093c2316baafe7aec5b25f30eba2e113599c4"},
			{length: 55, digest: "b2ca0adc388a66d3a5a7f6541239331d5a867eede554e11262241271faaf36c5"},
			{length: 56, digest: "3768eff44f1df02704a832cf708935fbed9bb74d0714fbe75454c266c4e12856"},
			{length: 57, digest: "a35b6298adfbc5ec16dc51a8d52145daa2a68f1d66e837d85bed6783d874c5cb"},
			{length: 63, digest: "dc06ea9e456da9f9fc2276564018231c36196ff7745bf505fd7f503aec5c67fd"},
			{length: 64, digest: "b68fe543b0b5a544e32eb08712e697bfcd3a3cb491563c3b1cba112d378f4bdb"},
			{length: 65, digest: "587cdb9ad9be371721f2f84a84eb09f2b698ec19bbe2d7c222179996547a04a6"},
			{length: 119, digest: "b27b61037be684924f5497ee79c784f09c971d8c8c3136f6270e0fd34775aa1b"},
			{length: 120, digest: "00c666e960132985e9c99dc98bf25782df791138d0f1252b69781694ed97d693"},
			{length: 127, digest: "805f2702e76ff5d578052c965ded05a689d9293fd86c5584b0db399d33e9745c"},
			{length: 128, digest: "cd77acfc16d0964a4c3ef06f0678a5c88d2d610de00f7b997676fea5b4d2ae9e"},
			{length: 129, digest: "2c0b373b4a62770395098399a7d08ca18fce440e3b6c35a22e5910eaf515b72c"},
			{length: 1000, digest: "fa10d4e76ab24b0ff9d6c596a00df178cd2a80a15c70770dee68d374ccb39bc6"},
			{length: 4099, digest: "4b5ecad58bcba9adefdb320c36c49a41d04eb303e40e2a80abf7a720bb4a8ac3"}
		];
		for (entry in pinned) {
			expectDigest(pattern(entry.length), entry.digest, 'the ${entry.length}-byte pattern');
		}

		for (length in 0...301) {
			final bytes = pattern(length);
			expectDigest(bytes, Sha256.make(bytes).toHex(), 'the ${length}-byte pattern against haxe.crypto.Sha256');
		}

		final offset = pattern(200).sub(37, 100);
		expectDigest(offset, Sha256.make(offset).toHex(), "a sub-range copy");

		Sys.println("ok");
	}
}
