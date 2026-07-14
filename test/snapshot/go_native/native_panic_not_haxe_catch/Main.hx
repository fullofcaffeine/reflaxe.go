@:go.import("log")
extern class GoLog {
	@:go.name("Panic")
	public static function panic(message:String):Void;
}

class Main {
	static function main():Void {
		Sys.println("native-start");
		try {
			GoLog.panic("native-failure");
		} catch (error:haxe.Exception) {
			Sys.println("incorrectly-caught=" + Std.string(error));
		}
	}
}
