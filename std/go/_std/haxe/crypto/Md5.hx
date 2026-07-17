package haxe.crypto;

import haxe.io.Bytes;
import go.NativeSlice;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements `Md5.encode` and `Md5.make` as staged Haxe source.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` while generated
	  Go lacks Haxe array-growth semantics for indexed writes; the public API still
	  does not justify compiler-owned declarations.

	How:
	- Convert public `Bytes` through typed integer arrays and delegate only MD5
	  execution to Go's standard runtime implementation.
**/
class Md5 {
	public static function encode(value:String):String {
		return NativeCrypto.md5String(value);
	}

	public static function make(value:Bytes):Bytes {
		return fromValues(NativeCrypto.md5Values(toValues(value)));
	}

	static function toValues(bytes:Bytes):NativeSlice<Int> {
		var values = new Array<Int>();
		for (index in 0...bytes.length)
			values.push(bytes.get(index));
		return NativeSlice.fromArray(values);
	}

	static function fromValues(values:NativeSlice<Int>):Bytes {
		var bytes = Bytes.alloc(values.length);
		for (index in 0...values.length)
			bytes.set(index, values[index]);
		return bytes;
	}
}
