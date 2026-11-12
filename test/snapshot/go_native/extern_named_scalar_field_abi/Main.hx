/**
 * Isolates a standard-library named scalar stored by value in a native struct.
 * No socket is opened: the timeout is parsed, assigned, then read locally.
 */
@:goNative
class Main {
	static var reads:Int = 0;
	static var order:String = "";

	static function observed(dialer:Dialer):Dialer {
		reads++;
		order += "L";
		return dialer;
	}

	static function supplied(value:Duration):Duration {
		order += "R";
		return value;
	}

	static function main():Void {
		var parsed = TimePackage.parseDuration("750ms");
		if (parsed.value2 != null)
			throw "duration parse failed";
		var dialer = new Dialer();
		dialer.timeout = parsed.value1;
		var roundTrip = observed(dialer).timeout;
		if (roundTrip.string() != "750ms")
			throw "named scalar field changed value";
		if (reads != 1)
			throw "native field receiver evaluated more than once";
		reads = 0;
		order = "";
		var assigned = observed(dialer).timeout = supplied(parsed.value1);
		if (reads != 1 || order != "LR")
			throw "native field assignment must evaluate receiver once before right side";
		var second = TimePackage.parseDuration("1s");
		if (second.value2 != null)
			throw "second duration parse failed";
		dialer.timeout = second.value1;
		if (roundTrip.string() != "750ms" || assigned.string() != "750ms" || dialer.timeout.string() != "1s")
			throw "native value read or assignment result aliased field storage";

		// A pointer field has the same imported struct carrier but retains aliases.
		var request = new Request();
		var url = new URL();
		url.path = "before";
		request.url = url;
		var pointerRead = request.url;
		url.path = "after";
		if (pointerRead.path != "after" || request.url.path != "after")
			throw "unmarked pointer field lost native alias";
		Sys.println("named scalar field roundtrip: PASS");
	}
}

/** Exact opaque named scalar, matching the generator's time.Duration shape. */
@:go.import("time")
@:go.name("Duration")
@:go.valueType
extern class Duration {
	@:go.name("String")
	public function string():String;
}

/** Only the exported field under qualification is exposed. */
@:go.import("net")
@:go.name("Dialer")
@:go.struct
extern class Dialer {
	public function new();
	@:go.name("Timeout")
	@:go.valueField
	public var timeout:Duration;
}

/** Real standard-library pointer fields must keep their original native ABI. */
@:go.import("net/http")
@:go.name("Request")
@:go.struct
extern class Request {
	public function new();
	@:go.name("URL")
	public var url:URL;
}

@:go.import("net/url")
@:go.name("URL")
@:go.struct
extern class URL {
	public function new();
	@:go.name("Path")
	public var path:String;
}

@:go.import("time")
extern class TimePackage {
	@:go.name("ParseDuration")
	@:go.tupleReturn
	@:go.tupleValueResults("0")
	public static function parseDuration(value:String):DurationResult;
}

/** Lossless error carrier for the standard library multi-return signature. */
class DurationResult {
	public var value1(default, null):Duration;
	public var value2(default, null):Null<go.NativeError>;

	public function new(value1:Duration, value2:Null<go.NativeError>) {
		this.value1 = value1;
		this.value2 = value2;
	}
}
