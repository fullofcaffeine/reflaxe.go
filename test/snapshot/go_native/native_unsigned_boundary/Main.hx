import SyscallOpen;
import UnixSyscallOpen;
import WindowsSyscallOpen;
import go.NativeSlice;
import go.NativeUint;
import go.NativeUint16;
import go.NativeUint32;

/** `math/bits` declares exact `uint32`, `uint16` and `uint` APIs on every GOOS. */
@:go.import("math/bits")
@:go.package("bits")
extern class NativeBits {
	/** Go: `func OnesCount32(x uint32) int`. */
	@:go.name("OnesCount32") public static function onesCount32(x:NativeUint32):Int;

	/** Go: `func ReverseBytes32(x uint32) uint32`. */
	@:go.name("ReverseBytes32") public static function reverseBytes32(x:NativeUint32):NativeUint32;

	/** Go: `func Add32(x, y, carry uint32) (sum, carryOut uint32)`. */
	@:go.name("Add32")
	@:go.tupleReturn
	public static function add32(x:NativeUint32, y:NativeUint32, carry:NativeUint32):Add32Result;

	/** Go: `func ReverseBytes16(x uint16) uint16`. */
	@:go.name("ReverseBytes16") public static function reverseBytes16(x:NativeUint16):NativeUint16;

	/** Go: `func OnesCount(x uint) int`. */
	@:go.name("OnesCount") public static function onesCount(x:NativeUint):Int;

	/** Go: `func Len(x uint) int`. */
	@:go.name("Len") public static function len(x:NativeUint):Int;
}

/** Go untyped constants: each read takes the carrier's Go type, not the default `int`. */
@:go.import("math")
@:go.package("math")
extern class NativeMath {
	@:go.name("MaxUint32") public static var maxUint32(default, never):NativeUint32;
	@:go.name("MaxUint16") public static var maxUint16(default, never):NativeUint16;
	@:go.name("MaxUint") public static var maxUint(default, never):NativeUint;
}

/** `slices` is generic: Go infers the element type, `uint32`, `uint16` or `uint`, from the native slice. */
@:go.import("slices")
@:go.package("slices")
extern class NativeSlicesPkg {
	@:go.name("Max") public static function maxUint32(values:NativeSlice<NativeUint32>):NativeUint32;
	@:go.name("Index") public static function indexUint16(values:NativeSlice<NativeUint16>, value:NativeUint16):Int;
	@:go.name("Max") public static function maxUint(values:NativeSlice<NativeUint>):NativeUint;
}

/** `io/fs.FileMode` is the named Go type `type FileMode uint32`. It stays a separate value type. */
@:go.import("io/fs")
@:go.package("fs")
@:go.name("FileMode")
@:go.valueType
extern class NativeFileMode {
	/** Go: `func (m FileMode) Perm() FileMode`. */
	@:go.valueReturn
	@:go.name("Perm") public function perm():NativeFileMode;

	@:go.name("String") public function toString():String;
}

@:go.import("io/fs")
@:go.package("fs")
extern class NativeFs {
	/** The explicit Go conversion `fs.FileMode(bits)` from an exact `uint32`. */
	@:go.valueReturn
	@:go.name("FileMode") public static function fileMode(bits:NativeUint32):NativeFileMode;
}

/** `archive/zip.FileHeader` has exported `uint32` and `uint16` fields and a named `fs.FileMode` API. */
@:go.import("archive/zip")
@:go.package("zip")
@:go.name("FileHeader")
@:go.struct
extern class NativeZipHeader {
	public function new();
	@:go.name("ExternalAttrs") public var externalAttrs:NativeUint32;
	@:go.name("Method") public var method:NativeUint16;

	/** Go: `func (h *FileHeader) SetMode(mode fs.FileMode)`. */
	@:go.valueArgs("0")
	@:go.name("SetMode") public function setMode(mode:NativeFileMode):Void;

	/** Go: `func (h *FileHeader) Mode() (mode fs.FileMode)`. */
	@:go.valueReturn
	@:go.name("Mode") public function mode():NativeFileMode;
}

/** `image/color.RGBA64` has exported `uint16` fields; its `RGBA` method returns four `uint32` values. */
@:go.import("image/color")
@:go.package("color")
@:go.name("RGBA64")
@:go.struct
extern class NativeRgba64 {
	public function new();
	@:go.name("R") public var r:NativeUint16;
	@:go.name("A") public var a:NativeUint16;

	/** Go: `func (c RGBA64) RGBA() (r, g, b, a uint32)`. */
	@:go.name("RGBA")
	@:go.tupleReturn
	public function rgba():RgbaResult;
}

/** The two `uint32` results of `bits.Add32`, in Go order. */
final class Add32Result {
	public final sum:NativeUint32;
	public final carryOut:NativeUint32;

	public function new(sum:NativeUint32, carryOut:NativeUint32) {
		this.sum = sum;
		this.carryOut = carryOut;
	}
}

/** The four `uint32` results of `color.RGBA64.RGBA`, in Go order. */
final class RgbaResult {
	public final r:NativeUint32;
	public final g:NativeUint32;
	public final b:NativeUint32;
	public final a:NativeUint32;

	public function new(r:NativeUint32, g:NativeUint32, b:NativeUint32, a:NativeUint32) {
		this.r = r;
		this.g = g;
		this.b = b;
		this.a = a;
	}
}

/** Ordinary Haxe class fields that store each carrier. */
class Counters {
	public var wide:NativeUint32;
	public var narrow:NativeUint16;
	public var word:NativeUint;

	public function new(wide:NativeUint32, narrow:NativeUint16, word:NativeUint) {
		this.wide = wide;
		this.narrow = narrow;
		this.word = word;
	}
}

@:goNative
class Main {
	static function check(ok:Bool, label:String):Void {
		if (!ok)
			throw label;
	}

	/** Haxe parameters of the carrier type; `==` compares the native Go values. */
	static function sameUint32(left:NativeUint32, right:NativeUint32):Bool {
		return left == right;
	}

	static function sameUint16(left:NativeUint16, right:NativeUint16):Bool {
		return left == right;
	}

	static function sameUint(left:NativeUint, right:NativeUint):Bool {
		return left == right;
	}

	static function main():Void {
		// A uint32 variable, not a Go constant, reaches Go uint32 parameters.
		final perm:NativeUint32 = NativeUint32.fromInt(0x1A4);
		check(NativeBits.onesCount32(perm) == 4, "uint32 parameter");
		final reversed = NativeBits.reverseBytes32(perm);
		check(sameUint32(reversed, NativeUint32.fromInt(0xA4010000)), "uint32 result and negative constant");
		check(SyscallOpen.openReadOnly("go.mod", perm), "syscall.Open uint32 perm");

		// A negative Int wraps, as Go does: as a constant and as a variable.
		final minusOne = NativeBits.onesCount32(perm) - 5;
		final allBits = NativeUint32.fromInt(-1);
		check(sameUint32(allBits, NativeUint32.fromInt(minusOne)), "uint32 wrap");
		final sum = NativeBits.add32(allBits, NativeUint32.fromInt(1), NativeUint32.fromInt(0));
		check(NativeUint32.toInt(sum.sum) == 0 && NativeUint32.toInt(sum.carryOut) == 1, "uint32 tuple fields");

		// Named uint32: fs.FileMode is a separate value type reached by an explicit conversion.
		final header = new NativeZipHeader();
		header.setMode(NativeFs.fileMode(perm));
		final permText = header.mode().perm().toString();
		check(NativeUint32.toInt(header.externalAttrs) != 0, "uint32 struct field read");
		header.externalAttrs = perm;
		check(sameUint32(header.externalAttrs, perm), "uint32 struct field write");
		header.method = NativeUint16.fromInt(8);
		check(NativeUint16.toInt(header.method) == 8, "uint16 struct field");

		// uint16: parameters, results, constant masking and struct fields.
		check(sameUint16(NativeBits.reverseBytes16(NativeUint16.fromInt(0x1234)), NativeUint16.fromInt(0x3412)), "uint16 parameter");
		check(sameUint16(NativeUint16.fromInt(0x12345), NativeUint16.fromInt(0x2345)), "uint16 constant masking");
		check(sameUint16(NativeUint16.fromInt(-1), NativeUint16.fromInt(minusOne)), "uint16 wrap");
		final color = new NativeRgba64();
		color.r = NativeUint16.fromInt(0xFFFF);
		color.a = NativeUint16.fromInt(-1);
		final rgba = color.rgba();
		check(NativeUint32.toInt(rgba.r) == 0xFFFF
			&& NativeUint32.toInt(rgba.g) == 0
			&& sameUint32(rgba.a, rgba.r), "uint32 method tuple");

		// uint: the platform word, 32 or 64 bits.
		check(NativeBits.onesCount(NativeUint.fromInt(255)) == 8 && NativeBits.len(NativeUint.fromInt(256)) == 9, "uint parameter");
		final allWord = NativeUint.fromInt(-1);
		check(sameUint(allWord, NativeUint.fromInt(minusOne)) && NativeBits.onesCount(allWord) == NativeBits.len(allWord), "uint wrap");
		check(NativeUint.toInt(allWord) == -1, "uint to int");

		// Go untyped constants keep the carrier type in locals and in boxed array elements.
		final maxWide = NativeMath.maxUint32;
		check(sameUint32(maxWide, allBits) && NativeBits.onesCount32(maxWide) == 32, "uint32 constant local");
		final maxNarrow = NativeMath.maxUint16;
		check(sameUint16(maxNarrow, NativeUint16.fromInt(-1)), "uint16 constant local");
		final maxWords = NativeSlice.fromArray([for (_ in 0...2) NativeMath.maxUint]);
		check(sameUint(maxWords[1], allWord), "uint constant slice element");

		// Class fields.
		final counters = new Counters(perm, NativeUint16.fromInt(7), NativeUint.fromInt(9));
		counters.wide = reversed;
		counters.narrow = NativeBits.reverseBytes16(counters.narrow);
		check(sameUint32(counters.wide, reversed) && NativeUint16.toInt(counters.narrow) == 0x0700, "class fields");
		check(NativeUint.toInt(counters.word) == 9, "uint class field");

		// Native slice elements: []uint32, []uint16 and []uint.
		final wide = NativeSlice.fromArray([perm, NativeUint32.fromInt(7), reversed]);
		check(sameUint32(NativeSlicesPkg.maxUint32(wide), reversed), "uint32 slice");
		wide[1] = allBits;
		check(sameUint32(NativeSlicesPkg.maxUint32(wide), allBits), "uint32 slice element write");
		final narrow = NativeSlice.fromArray([NativeUint16.fromInt(1), NativeUint16.fromInt(2)]);
		check(NativeSlicesPkg.indexUint16(narrow, NativeUint16.fromInt(2)) == 1, "uint16 slice");
		final words = NativeSlice.fromArray([NativeUint.fromInt(3), allWord]);
		check(sameUint(NativeSlicesPkg.maxUint(words), allWord), "uint slice");

		Sys.println("perm: " + permText);
		Sys.println("ones: " + NativeBits.onesCount32(perm));
		Sys.println("round trip: " + NativeUint32.toInt(NativeUint32.fromInt(42)) + " " + NativeUint16.toInt(NativeUint16.fromInt(42)) + " "
			+ NativeUint.toInt(NativeUint.fromInt(42)));
		Sys.println("native-unsigned:ok");
	}
}
