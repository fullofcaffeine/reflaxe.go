@:go.import("strconv")
extern class StrconvTuplePkg {
	@:go.tupleReturn
	@:go.name("Atoi")
	static function atoi(value:String):AtoiResult;
}

class AtoiResult {
	public var n(default, null):Int;
	public var err(default, null):Null<go.Error>;

	public function new(n:Int, err:Null<go.Error>) {
		this.n = n;
		this.err = err;
	}
}

@:go.import("time")
@:go.name("Time")
extern class TimeTuple {
	@:go.name("Now")
	static function now():TimeTuple;

	@:go.tupleReturn
	@:go.name("Zone")
	function zone():TimeZoneResult;
}

class TimeZoneResult {
	public var name(default, null):String;
	public var offset(default, null):Int;

	public function new(name:String, offset:Int) {
		this.name = name;
		this.offset = offset;
	}
}

class Main {
	static function main() {
		var ok = StrconvTuplePkg.atoi("12");
		var bad = StrconvTuplePkg.atoi("nope");
		var zone = TimeTuple.now().zone();
		Sys.println('atoi.ok=' + ok.n + ':' + (ok.err == null));
		Sys.println('atoi.err=' + (bad.err != null));
		Sys.println('zone.typed=' + (zone.name.length >= 0) + ':' + (zone.offset >= -86400));
	}
}
