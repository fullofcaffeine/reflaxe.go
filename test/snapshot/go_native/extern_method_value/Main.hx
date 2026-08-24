@:go.import("time")
@:go.name("Time")
extern class GoTime {
	@:go.name("Now")
	public static function now():GoTime;

	@:go.name("Unix")
	public function unix():Int;
}

@:go.import("log")
@:go.name("Logger")
extern class GoLogger {
	@:go.name("Default")
	public static function defaultLogger():GoLogger;

	@:go.name("SetPrefix")
	public function setPrefix(value:String):Void;

	@:go.name("Prefix")
	public function prefix():String;
}

class Main {
	static function main():Void {
		final now = GoTime.now();
		final unix:Void->Int = now.unix;
		Sys.println(unix() > 0);

		final logger = GoLogger.defaultLogger();
		final originalPrefix = logger.prefix();
		final setPrefix:String->Void = logger.setPrefix;
		final getPrefix:Void->String = logger.prefix;
		setPrefix("hx:");
		Sys.println(getPrefix() == "hx:");
		logger.setPrefix(originalPrefix);
	}
}
