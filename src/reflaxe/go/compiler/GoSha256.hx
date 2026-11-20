package reflaxe.go.compiler;

#if (macro || eval)
import haxe.ds.Vector;
import haxe.io.Bytes;

/**
	Standard SHA-256 written for the Haxe macro interpreter.

	Why
	Existing-module ownership records store a SHA-256 digest for every generated
	file, so a large compile hashes several megabytes at macro time. The macro
	interpreter has no native SHA-256: `sys.ssl.Digest.make` has no eval binding
	in Haxe 4.3. The portable `haxe.crypto.Sha256` is written so it also runs on
	JavaScript, where integers are not 32 bits wide. It splits every addition
	into 16-bit halves and first copies the input byte by byte into an
	`Array<Int>`. In the interpreter every integer operation allocates, so that
	extra work costs both time and garbage-collector load in a big compile.

	What
	`hex` returns the same lowercase hex digest as
	`haxe.crypto.Sha256.make(bytes).toHex()` for every input, about 2.5 times
	faster in the interpreter. It does not change or keep the input.

	How
	The interpreter's `Int` wraps at 32 bits, so additions need no carry fixes.
	Each 64-byte block is read directly from `bytes`; only the final one or two
	padded blocks go through a small copy. Rounds follow FIPS 180-4 section
	6.2.2, with `Ch` and `Maj` in their usual reduced forms. The contract test
	`GoOwnershipDigestContract` pins the result to Python's `hashlib`.
**/
class GoSha256 {
	static final ROUND_CONSTANTS:Vector<Int> = Vector.fromArrayCopy([
		0x428A2F98, 0x71374491, 0xB5C0FBCF, 0xE9B5DBA5, 0x3956C25B, 0x59F111F1, 0x923F82A4, 0xAB1C5ED5,
		0xD807AA98, 0x12835B01, 0x243185BE, 0x550C7DC3, 0x72BE5D74, 0x80DEB1FE, 0x9BDC06A7, 0xC19BF174,
		0xE49B69C1, 0xEFBE4786, 0x0FC19DC6, 0x240CA1CC, 0x2DE92C6F, 0x4A7484AA, 0x5CB0A9DC, 0x76F988DA,
		0x983E5152, 0xA831C66D, 0xB00327C8, 0xBF597FC7, 0xC6E00BF3, 0xD5A79147, 0x06CA6351, 0x14292967,
		0x27B70A85, 0x2E1B2138, 0x4D2C6DFC, 0x53380D13, 0x650A7354, 0x766A0ABB, 0x81C2C92E, 0x92722C85,
		0xA2BFE8A1, 0xA81A664B, 0xC24B8B70, 0xC76C51A3, 0xD192E819, 0xD6990624, 0xF40E3585, 0x106AA070,
		0x19A4C116, 0x1E376C08, 0x2748774C, 0x34B0BCB5, 0x391C0CB3, 0x4ED8AA4A, 0x5B9CCA4F, 0x682E6FF3,
		0x748F82EE, 0x78A5636F, 0x84C87814, 0x8CC70208, 0x90BEFFFA, 0xA4506CEB, 0xBEF9A3F7, 0xC67178F2
	]);

	/** Return the lowercase hex SHA-256 digest of `bytes`. */
	public static function hex(bytes:Bytes):String {
		final state = new Vector<Int>(8);
		state[0] = 0x6A09E667;
		state[1] = 0xBB67AE85;
		state[2] = 0x3C6EF372;
		state[3] = 0xA54FF53A;
		state[4] = 0x510E527F;
		state[5] = 0x9B05688C;
		state[6] = 0x1F83D9AB;
		state[7] = 0x5BE0CD19;
		final schedule = new Vector<Int>(64);

		final length = bytes.length;
		final fullBlocks = length >> 6;
		for (block in 0...fullBlocks) {
			compress(state, schedule, bytes, block << 6);
		}

		// Padding: the remaining bytes, one 0x80 byte, zeros, then the message
		// length in bits as a 64-bit big-endian integer.
		final remaining = length & 63;
		final tail = Bytes.alloc(remaining < 56 ? 64 : 128);
		tail.blit(0, bytes, fullBlocks << 6, remaining);
		tail.set(remaining, 0x80);
		writeBigEndian(tail, tail.length - 8, length >>> 29);
		writeBigEndian(tail, tail.length - 4, length << 3);
		var offset = 0;
		while (offset < tail.length) {
			compress(state, schedule, tail, offset);
			offset += 64;
		}

		final digest = Bytes.alloc(32);
		for (index in 0...8) {
			writeBigEndian(digest, index << 2, state[index]);
		}
		return digest.toHex();
	}

	static function compress(state:Vector<Int>, w:Vector<Int>, data:Bytes, offset:Int):Void {
		final k = ROUND_CONSTANTS;
		for (index in 0...16) {
			final position = offset + (index << 2);
			w[index] = (data.get(position) << 24) | (data.get(position + 1) << 16) | (data.get(position + 2) << 8) | data.get(position + 3);
		}
		for (index in 16...64) {
			final x = w[index - 15];
			final y = w[index - 2];
			final s0 = rotateRight(x, 7) ^ rotateRight(x, 18) ^ (x >>> 3);
			final s1 = rotateRight(y, 17) ^ rotateRight(y, 19) ^ (y >>> 10);
			w[index] = w[index - 16] + s0 + w[index - 7] + s1;
		}

		var a = state[0];
		var b = state[1];
		var c = state[2];
		var d = state[3];
		var e = state[4];
		var f = state[5];
		var g = state[6];
		var h = state[7];
		for (index in 0...64) {
			final t1 = h + (rotateRight(e, 6) ^ rotateRight(e, 11) ^ rotateRight(e, 25)) + (g ^ (e & (f ^ g))) + k[index] + w[index];
			final t2 = (rotateRight(a, 2) ^ rotateRight(a, 13) ^ rotateRight(a, 22)) + ((a & b) | (c & (a | b)));
			h = g;
			g = f;
			f = e;
			e = d + t1;
			d = c;
			c = b;
			b = a;
			a = t1 + t2;
		}
		state[0] = state[0] + a;
		state[1] = state[1] + b;
		state[2] = state[2] + c;
		state[3] = state[3] + d;
		state[4] = state[4] + e;
		state[5] = state[5] + f;
		state[6] = state[6] + g;
		state[7] = state[7] + h;
	}

	static inline function rotateRight(value:Int, bits:Int):Int {
		return (value >>> bits) | (value << (32 - bits));
	}

	static inline function writeBigEndian(target:Bytes, position:Int, value:Int):Void {
		target.set(position, (value >>> 24) & 0xFF);
		target.set(position + 1, (value >>> 16) & 0xFF);
		target.set(position + 2, (value >>> 8) & 0xFF);
		target.set(position + 3, value & 0xFF);
	}
}
#else
class GoSha256 {}
#end
