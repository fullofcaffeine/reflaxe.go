using haxe.EnumTools;

enum Mode {
	Alpha;
	Beta(code:Int);
	Gamma;
}

class Main {
	static function join(values:Array<String>):String {
		var out = "";
		for (index in 0...values.length) {
			if (index > 0) {
				out += "|";
			}
			out += values[index];
		}
		return out;
	}

	static function main() {
		var flags:haxe.EnumFlags<Mode> = Mode.Alpha;
		flags.set(Mode.Gamma);
		flags.unset(Mode.Alpha);
		Sys.println("flags.alpha=" + flags.has(Mode.Alpha));
		Sys.println("flags.gamma=" + flags.has(Mode.Gamma));
		Sys.println("flags.int=" + flags.toInt());

		var enumType = Type.getEnum(Mode.Alpha);
		var created:Mode = enumType.createByIndex(1, [7]);
		Sys.println("created.name=" + Type.enumConstructor(created));
		Sys.println("created.index=" + Type.enumIndex(created));
		Sys.println("ctors=" + join(enumType.getConstructors()));
		Sys.println("all=" + join([for (value in enumType.createAll()) Type.enumConstructor(value)]));
	}
}
