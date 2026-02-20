import haxe.Int32;

class Main {
	static function out(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main():Void {
		var max:Int32 = 0x7fffffff;
		var min:Int32 = -2147483648;
		out("overflow.add", max + 1);
		out("underflow.sub", min - 1);
		out("mul.wrap", ((cast 1073741824 : Int32) * 4));
		out("div.float", ((cast 7 : Int32) / 2));
		out("mod.neg", ((cast -7 : Int32) % 4));

		out("and", ((cast 252645135 : Int32) & -16711936));
		out("or", ((cast 252645135 : Int32) | -268435456));
		out("xor", ((cast 252645135 : Int32) ^ 16711935));
		out("complement", ~(cast 16711935 : Int32));

		out("shr", ((cast -16 : Int32) >> 2));
		out("ushr", ((cast -16 : Int32) >>> 2));
		out("shl", ((cast 1 : Int32) << 31));

		out("ucompare.neg.pos", Int32.ucompare((cast -1 : Int32), (cast 1 : Int32)));
		out("ucompare.max.neg", Int32.ucompare((cast 2147483647 : Int32), (cast -1 : Int32)));

		var x:Int32 = 2147483647;
		out("post.inc.ret", x++);
		out("post.inc.val", x);

		var y:Int32 = -2147483648;
		out("pre.dec.val", --y);

		var mix:Int32 = 5;
		out("add.int", mix + 3);
		out("int.add", 3 + mix);
		out("add.float", mix + 0.5);
		out("float.add", 0.5 + mix);
		out("eq.int", mix == 5);
		out("eq.float", mix == 5.0);
	}
}
