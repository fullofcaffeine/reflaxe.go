import haxe.PosInfos;

class Main {
	static function position(fileName:String, lineNumber:Int, customParams:Array<Dynamic>):PosInfos {
		return {
			fileName: fileName,
			lineNumber: lineNumber,
			className: "Main",
			methodName: "main",
			customParams: customParams
		};
	}

	static function main():Void {
		var manual = position("manual.hx", 7, ["tail", null]);
		Sys.println("format.plain=" + haxe.Log.formatOutput("value", null));
		Sys.println("format.pos=" + haxe.Log.formatOutput("value", manual));
		haxe.Log.trace("default", manual);

		var original = haxe.Log.trace;
		haxe.Log.trace = function(value:Dynamic, ?infos:PosInfos):Void {
			Sys.println("custom.value=" + Std.string(value));
			var count = infos == null || infos.customParams == null ? -1 : infos.customParams.length;
			Sys.println("custom.info="
				+ (infos != null)
				+ ":"
				+ infos.className
				+ ":"
				+ infos.methodName
				+ ":"
				+ (infos.fileName == "Main.hx")
				+ ":"
				+ count);
		};
		trace("rebound", "tail", null);

		haxe.Log.trace = original;
		var direct = haxe.Log.trace;
		direct("function.value", null);

		haxe.Log.trace = null;
		try {
			haxe.Log.trace("ignored", null);
			Sys.println("null=no-throw");
		} catch (_:Dynamic) {
			Sys.println("null=throws");
		}

		haxe.Log.trace = original;
		haxe.Log.trace("restored", null);
	}
}
