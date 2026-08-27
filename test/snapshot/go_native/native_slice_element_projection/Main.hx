import go.NativeStringSlice;

@:go.import("fmt")
extern class GoFmt {
	@:go.name("Println")
	static function println(value:String):Void;
}

class Main {
	static function renderStrings(values:NativeStringSlice):String {
		var rendered = "";
		for (index in 0...values.length) {
			if (index > 0)
				rendered += ",";
			rendered += values[index];
		}
		return rendered;
	}

	static function main():Void {
		final stringSource = ["alpha", "beta"];
		final nativeStrings:NativeStringSlice = NativeStringSlice.fromArray(stringSource);
		final inlineStrings:NativeStringSlice = NativeStringSlice.fromArray(["gamma"]);
		nativeStrings[0] = "delta";

		GoFmt.println("strings=" + renderStrings(nativeStrings));
		GoFmt.println("inline=" + renderStrings(inlineStrings));
	}
}
