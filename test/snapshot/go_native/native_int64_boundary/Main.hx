import go.NativeInt64;
import go.Result;

/** Standard-library boundary: ParseInt and FormatInt exchange exact Go int64 values. */
@:go.import("strconv")
extern class NativeStrconv {
	@:go.name("ParseInt")
	@:go.valueError
	public static function parseInt(value:String, base:Int, bitSize:Int):Result<NativeInt64>;

	@:go.name("FormatInt")
	public static function formatInt(value:NativeInt64, base:Int):String;

	@:go.name("ParseInt")
	@:go.tupleReturn
	public static function parseTuple(value:String, base:Int, bitSize:Int):ParsedInt64;
}

/** A precise two-result Go call without the generic Result storage carrier. */
final class ParsedInt64 {
	public var value(default, null):NativeInt64;
	public var error(default, null):Null<go.Error>;

	public function new(value:NativeInt64, error:Null<go.Error>) {
		this.value = value;
		this.error = error;
	}
}

/** Standard-library struct field whose Int64 member must retain its native width. */
@:go.import("database/sql")
@:go.name("NullInt64")
@:go.struct
extern class NativeNullInt64 {
	public function new();
	@:go.name("Int64") public var number:NativeInt64;
	@:go.name("Valid") public var valid:Bool;
}

@:goNative
class Main {
	static function main():Void {
		for (text in ["-9223372036854775808", "9223372036854775807", "9007199254740993"]) {
			final parsed = NativeStrconv.parseInt(text, 10, 64);
			if (!parsed.isOk())
				throw "ParseInt failed";
			final parsedTuple = NativeStrconv.parseTuple(text, 10, 64);
			if (parsedTuple.error != null)
				throw "ParseInt tuple failed";
			if (NativeStrconv.formatInt(parsedTuple.value, 10) != text)
				throw "tuple value changed";
			final row = new NativeNullInt64();
			row.number = parsed.unwrap();
			row.valid = true;
			final actual = NativeStrconv.formatInt(row.number, 10);
			if (!row.valid || actual != text)
				throw "int64 field round trip failed";
			Sys.println(actual);
		}
	}
}
