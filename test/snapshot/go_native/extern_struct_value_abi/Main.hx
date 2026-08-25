@:go.import("image")
@:go.name("Point")
@:go.struct
extern class GoPoint {
	public function new();

	@:go.name("X") public var x:Int;
	@:go.name("Y") public var y:Int;

	@:go.valueArgs("0")
	@:go.valueReturn
	@:go.name("Add")
	public function add(other:GoPoint):GoPoint;
}

@:go.import("time")
@:go.name("Time")
@:go.struct
extern class GoTime {
	@:go.valueArgs("0")
	@:go.name("Equal")
	public function equal(other:GoTime):Bool;
}

@:go.import("time")
extern class GoTimePkg {
	@:go.tupleReturn
	@:go.tupleValueResults("0")
	@:go.name("Parse")
	public static function parse(layout:String, value:String):ParseResult;
}

@:go.import("fmt")
extern class GoFmt {
	@:go.name("Println")
	public static function println(value:String):Void;
}

final class ParseResult {
	public var value(default, null):GoTime;
	public var error(default, null):Null<go.Error>;

	public function new(value:GoTime, error:Null<go.Error>) {
		this.value = value;
		this.error = error;
	}
}

class Main {
	static function main():Void {
		final left = new GoPoint();
		left.x = 20;
		final right = new GoPoint();
		right.y = 22;
		final sum = left.add(right);
		GoFmt.println("point=" + (sum.x + sum.y));

		final first = GoTimePkg.parse("2006-01-02", "2026-08-24");
		final second = GoTimePkg.parse("2006-01-02", "2026-08-24");
		GoFmt.println("time.equal=" + (first.error == null && second.error == null && first.value.equal(second.value)));
	}
}
