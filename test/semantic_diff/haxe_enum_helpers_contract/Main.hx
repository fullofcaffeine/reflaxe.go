using haxe.EnumTools;

enum Access {
	Read;
	Write;
	Execute;
}

enum Token {
	Word(value:String);
	Stop;
	Pause;
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
		var flags:haxe.EnumFlags<Access> = Access.Read;
		flags.set(Access.Write);
		flags.setTo(Access.Execute, true);
		flags.setTo(Access.Read, false);
		Sys.println("flags.has.read=" + flags.has(Access.Read));
		Sys.println("flags.has.write=" + flags.has(Access.Write));
		Sys.println("flags.has.execute=" + flags.has(Access.Execute));
		Sys.println("flags.int=" + flags.toInt());
		var combined = flags | Access.Read;
		Sys.println("flags.combined.read=" + combined.has(Access.Read));

		var tokenEnum = Type.getEnum(Token.Stop);
		Sys.println("enum.name=" + tokenEnum.getName());
		Sys.println("enum.ctors=" + join(tokenEnum.getConstructors()));
		var stop:Token = tokenEnum.createByIndex(1);
		Sys.println("enum.stop.index=" + Type.enumIndex(stop));
		var word:Token = tokenEnum.createByName("Word", ["go"]);
		Sys.println("enum.word.name=" + Type.enumConstructor(word));
		Sys.println("enum.word.params=" + join([for (value in Type.enumParameters(word)) Std.string(value)]));
		Sys.println("enum.all=" + join([for (value in tokenEnum.createAll()) Type.enumConstructor(value)]));
		Sys.println("enum.equals=" + Type.enumEq(Token.Word("go"), word));
	}
}
