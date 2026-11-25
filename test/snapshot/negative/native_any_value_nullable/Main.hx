import go.NativeAny;

@:go.import("time") @:go.name("Time") @:go.struct
extern class GoTime {
	public function new();
}

/**
	A `Null<T>` carrier of a value extern can hold null, and a null carrier has
	no Go value (Go has no nil `time.Time`), so `go.NativeAny.ofValue` rejects
	the nullable type.
**/
class Main {
	static function main():Void {
		final time:Null<GoTime> = Sys.args().length > 99 ? new GoTime() : null;
		final value = NativeAny.ofValue(time);
		Sys.println(value != null);
	}
}
