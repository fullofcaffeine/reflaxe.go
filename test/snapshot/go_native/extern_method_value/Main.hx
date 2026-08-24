@:go.import("time")
@:go.name("Time")
extern class GoTime {
	@:go.name("Now")
	public static function now():GoTime;

	@:go.name("Unix")
	public function unix():Int;
}

class Main {
	static function main():Void {
		final now = GoTime.now();
		final unix:Void->Int = now.unix;
		Sys.println(unix() > 0);
	}
}
