using haxe.EnumTools;
using haxe.EnumTools.EnumValueTools;

enum Mode {
	Alpha;
	Beta(code:Int);
	Gamma;
}

class Main {
	static function main() {
		var flags:haxe.EnumFlags<Mode> = Mode.Alpha;
		flags.set(Mode.Gamma);
		flags.unset(Mode.Alpha);
		Sys.println("flags.alpha=" + flags.has(Mode.Alpha));
		Sys.println("flags.gamma=" + flags.has(Mode.Gamma));
		Sys.println("flags.int=" + flags.toInt());

		var enumType = Type.getEnum(Mode.Alpha);
		var created:Mode = enumType.createByIndex(1, [7]);
		Sys.println("created.name=" + created.getName());
		Sys.println("created.index=" + created.getIndex());
		Sys.println("ctors=" + enumType.getConstructors().join("|"));
		Sys.println("all=" + [for (value in enumType.createAll()) value.getName()].join("|"));
	}
}
