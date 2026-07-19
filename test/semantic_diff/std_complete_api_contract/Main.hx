class StdParent {
	public function new() {}
}

class StdChild extends StdParent {
	public function new() {
		super();
	}
}

class Main {
	static function emit(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main():Void {
		emit("parse.null", Std.parseInt(null));
		emit("parse.empty", Std.parseInt(""));
		emit("parse.space", Std.parseInt(" \t\n\r\x0b\x0c42tail"));
		emit("parse.plus", Std.parseInt("+17"));
		emit("parse.minus", Std.parseInt("-17"));
		emit("parse.hex.lower", Std.parseInt("0x2a!"));
		emit("parse.hex.upper", Std.parseInt("-0X2A!"));
		emit("parse.decimal.stop", Std.parseInt("12.9"));
		emit("parse.scientific.stop", Std.parseInt("10e2"));
		emit("parse.binary.unsupported", Std.parseInt("0b101"));
		emit("parse.octal.unsupported", Std.parseInt("010"));
		emit("parse.sign.space", Std.parseInt("+ 1"));
		emit("parse.invalid", Std.parseInt("word"));
		emit("parse.max", Std.parseInt("2147483647"));
		emit("parse.max.plus.one", Std.parseInt("2147483648"));
		emit("parse.min", Std.parseInt("-2147483648"));
		emit("parse.min.minus.one", Std.parseInt("-2147483649"));
		emit("parse.hex.max", Std.parseInt("0x7fffffff"));
		emit("parse.hex.high.bit", Std.parseInt("0x80000000"));
		emit("parse.hex.all.bits", Std.parseInt("0xffffffff"));

		emit("float.null.nan", Math.isNaN(Std.parseFloat(null)));
		emit("float.invalid.nan", Math.isNaN(Std.parseFloat("word")));
		emit("float.decimal", Std.parseFloat("  -1.5tail"));
		emit("float.exponent", Std.parseFloat("1.25e2tail"));
		emit("float.incomplete.exponent.nan", Math.isNaN(Std.parseFloat("1e+")));
		emit("int.positive", Std.int(1.9));
		emit("int.negative", Std.int(-1.9));

		var childAsParent:StdParent = new StdChild();
		var parent:StdParent = new StdParent();
		emit("downcast.hit", Std.downcast(childAsParent, StdChild) != null);
		emit("downcast.miss", Std.downcast(parent, StdChild) == null);
		emit("instance.hit", Std.instance(childAsParent, StdChild) != null);
		emit("is.alias", Std.is(childAsParent, StdChild));

		emit("random.zero", Std.random(0));
		emit("random.negative", Std.random(-3));
		emit("random.one", Std.random(1));
		var sample = Std.random(8);
		emit("random.range", sample >= 0 && sample < 8);
	}
}
