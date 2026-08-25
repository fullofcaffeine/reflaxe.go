@:go.import("time")
@:go.name("Duration")
@:go.valueType
extern class GoDuration {}

@:go.import("time")
@:go.name("Time")
@:go.struct
extern class GoTime {
	@:go.valueArgs("0")
	@:go.valueReturn
	@:go.name("Add")
	public function add(duration:GoDuration):GoTime;
}

@:go.import("time")
extern class GoTimePkg {
	@:go.valueReturn
	@:go.name("Unix")
	public static function unix(seconds:Int, nanoseconds:Int):GoTime;

	@:go.tupleReturn
	@:go.tupleValueResults("0")
	@:go.name("ParseDuration")
	public static function parseDuration(value:String):ParseDurationResult;
}

@:go.import("fmt")
extern class GoFmt {
	@:go.name("Println")
	public static function println(value:String):Void;
}

final class ParseDurationResult {
	public var value(default, null):GoDuration;
	public var error(default, null):Null<go.Error>;

	public function new(value:GoDuration, error:Null<go.Error>) {
		this.value = value;
		this.error = error;
	}
}

class Main {
	static function main():Void {
		final duration = GoTimePkg.parseDuration("1s");
		final start = GoTimePkg.unix(41, 0);
		GoFmt.println("ok=" + (duration.error == null && start.add(duration.value) != null));
	}
}
