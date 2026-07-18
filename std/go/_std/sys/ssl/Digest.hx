package sys.ssl;

import go.NativeSlice;
import haxe.io.Bytes;
import hxrt.ssl.NativeDigest;

/**
	What: Implements Haxe digest/signature helpers over typed native byte slices.
	Why: The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	because `sys.ssl.Digest` is extern. Hashing and asymmetric signing are native crypto operations, but Bytes
	conversion and the public API do not belong in raw injection or compiler shims.
	How: Copy Haxe Bytes to `NativeSlice<Int>`, call `NativeDigest`, and copy results
	back to ordinary Bytes.
**/
class Digest {
	public static function make(data:Bytes, algorithm:DigestAlgorithm):Bytes {
		var input = new Array<Int>();
		for (index in 0...data.length)
			input.push(data.get(index));
		var nativeResult = NativeDigest.make(NativeSlice.fromArray(input), cast algorithm);
		var result = Bytes.alloc(nativeResult.length);
		for (index in 0...nativeResult.length)
			result.set(index, nativeResult[index]);
		return result;
	}

	public static function sign(data:Bytes, privateKey:Key, algorithm:DigestAlgorithm):Bytes {
		var input = new Array<Int>();
		for (index in 0...data.length)
			input.push(data.get(index));
		var nativeResult = NativeDigest.sign(NativeSlice.fromArray(input), privateKey.handle, cast algorithm);
		var result = Bytes.alloc(nativeResult.length);
		for (index in 0...nativeResult.length)
			result.set(index, nativeResult[index]);
		return result;
	}

	public static function verify(data:Bytes, signature:Bytes, publicKey:Key, algorithm:DigestAlgorithm):Bool {
		var input = new Array<Int>();
		for (index in 0...data.length)
			input.push(data.get(index));
		var signatureValues = new Array<Int>();
		for (index in 0...signature.length)
			signatureValues.push(signature.get(index));
		return NativeDigest.verify(NativeSlice.fromArray(input), NativeSlice.fromArray(signatureValues), publicKey.handle, cast algorithm);
	}
}
