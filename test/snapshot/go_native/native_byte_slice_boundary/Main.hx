import go.NativeByte;
import go.NativeError;
import go.NativeSlice;

/** Exact standard-library externs used to prove a stored Go []byte boundary. */
@:go.import("bytes")
@:go.package("bytes")
@:go.name("Buffer")
@:go.struct
extern class ByteBuffer {
	@:go.name("Bytes") public function bytes():NativeSlice<NativeByte>;
	@:go.name("String") public function string():String;
}

@:go.import("bytes")
@:go.package("bytes")
extern class ByteBuffers {
	@:go.name("NewBufferString") public static function newBufferString(text:String):ByteBuffer;
	@:go.name("Clone") public static function clone(value:NativeSlice<NativeByte>):NativeSlice<NativeByte>;
	@:go.name("IndexByte") public static function indexByte(value:NativeSlice<NativeByte>, needle:NativeByte):Int;
}

@:go.import("encoding/json")
@:go.package("json")
extern class NativeJSON {
	@:go.name("Compact") public static function compact(output:ByteBuffer, source:NativeSlice<NativeByte>):NativeError;
}

@:go.import("encoding/hex")
@:go.package("hex")
extern class NativeHex {
	@:go.tupleReturn
	@:go.name("DecodeString") public static function decodeString(text:String):DecodedHex;
	@:go.name("EncodeToString") public static function encodeToString(value:NativeSlice<NativeByte>):String;
}

/** The exact []byte,error result of encoding/hex.DecodeString. */
class DecodedHex {
	public final value1:NativeSlice<NativeByte>;
	public final value2:Null<NativeError>;

	public function new(value1:NativeSlice<NativeByte>, value2:Null<NativeError>) {
		this.value1 = value1;
		this.value2 = value2;
	}
}

@:goNative
class Main {
	static function check(ok:Bool, label:String):Void {
		if (!ok)
			throw label;
	}

	static function compactData(source:NativeSlice<NativeByte>):String {
		final output = ByteBuffers.newBufferString("");
		final error = NativeJSON.compact(output, source);
		if (error != null)
			throw error.text();
		return output.string();
	}

	static function keepInts(value:NativeSlice<Int>):Int {
		return value.length;
	}

	static function main():Void {
		check(compactData(ByteBuffers.newBufferString("{ \"ok\": true }").bytes()) == "{\"ok\":true}", "typed []byte parameter");
		final absent:NativeSlice<NativeByte> = null;
		check(absent == null && ByteBuffers.clone(absent) == null, "nil byte slice");

		final decoded = NativeHex.decodeString("00ff");
		check(decoded.value2 == null, "hex decode");
		final octets = decoded.value1;
		check(octets.length == 2 && NativeHex.encodeToString(octets) == "00ff", "byte 0 and 255");
		final high:NativeByte = octets[1];
		check(ByteBuffers.indexByte(octets, high) == 1, "byte scalar argument");

		final buffer = ByteBuffers.newBufferString("ab");
		final view = buffer.bytes();
		view[0] = high;
		check(NativeHex.encodeToString(buffer.bytes()) == "ff62", "Buffer.Bytes aliases backing storage");
		final copy = ByteBuffers.clone(view);
		copy[1] = octets[0];
		check(NativeHex.encodeToString(copy) == "ff00" && NativeHex.encodeToString(buffer.bytes()) == "ff62", "Clone owns a copy");
		final portable = view.toArray();
		final copiedBack = NativeSlice.fromArray(portable);
		copiedBack[0] = octets[0];
		check(NativeHex.encodeToString(copiedBack) == "0062"
			&& NativeHex.encodeToString(buffer.bytes()) == "ff62", "Array boundary copies");

		check(keepInts(NativeSlice.fromArray([1, 2])) == 2, "ordinary NativeSlice<Int>");
		Sys.println("native-byte:ok");
	}
}
