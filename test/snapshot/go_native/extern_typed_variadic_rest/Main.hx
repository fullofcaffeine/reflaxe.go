/** Portable standard-library boundary: unicode.In requires ...*RangeTable. */
@:go.import("unicode") @:go.name("RangeTable")
extern class NativeRangeTable {}

@:go.import("unicode")
extern class NativeUnicode {
	@:go.name("Letter") public static var letter:NativeRangeTable;
	@:go.name("Digit") public static var digit:NativeRangeTable;
	@:go.name("In") public static function inRange(r:Int, ranges:haxe.Rest<NativeRangeTable>):Bool;
}

@:go.import("os")
extern class NativeOS {
	@:go.name("Exit") public static function exit(code:Int):Void;
}

@:go.import("text/template") @:go.name("Template")
extern class NativeTemplate {
	@:go.name("New") public static function make(name:String):NativeTemplate;
	@:go.name("Option") public function option(options:haxe.Rest<String>):NativeTemplate;
}

@:go.import("fmt")
extern class NativeFormat {
	@:go.name("Sprint") public static function sprint(values:haxe.Rest<Dynamic>):String;
}

@:goNative
class Main {
	static function sum(...values:Int):Int {
		var total = 0;
		for (value in values)
			total += value;
		return total;
	}

	static function main():Void {
		final letter:NativeRangeTable = NativeUnicode.letter;
		if (!NativeUnicode.inRange(65, letter, NativeUnicode.digit))
			NativeOS.exit(1);
		if (!NativeUnicode.inRange(65, letter))
			NativeOS.exit(4);
		if (NativeUnicode.inRange(65, NativeUnicode.digit))
			NativeOS.exit(2);
		final template = NativeTemplate.make("t");
		final option = "missingkey=zero";
		template.option(option, "missingkey=default");
		if (NativeFormat.sprint("left", 7, "right") != "left7right")
			NativeOS.exit(3);
		if (sum(1, 2, 3) != 6 || sum() != 0)
			NativeOS.exit(5);
	}
}
