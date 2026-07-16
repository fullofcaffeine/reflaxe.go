package haxe.crypto;

import haxe.io.Bytes;
import hxrt.crypto.NativeCrypto;

/**
	What:
	- Implements `Sha224.encode` and `Sha224.make` as staged Haxe source.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go` while generated
	  Go lacks Haxe array-growth semantics for indexed writes; the public API still
	  does not justify compiler-owned declarations.

	How:
	- Convert public `Bytes` through typed integer arrays and delegate only SHA-224
	  execution to Go's standard runtime implementation.
**/
class Sha224 {
	public static function encode(value:String):String {
		return NativeCrypto.sha224String(value);
	}

	public static function make(value:Bytes):Bytes {
		return fromValues(NativeCrypto.sha224Values(toValues(value)));
	}

	static function toValues(bytes:Bytes):Array<Int> {
		var values = new Array<Int>();
		for (index in 0...bytes.length)
			values.push(bytes.get(index));
		return values;
	}

	static function fromValues(values:Array<Int>):Bytes {
		var bytes = Bytes.alloc(values.length);
		for (index in 0...values.length)
			bytes.set(index, values[index]);
		return bytes;
	}
}
