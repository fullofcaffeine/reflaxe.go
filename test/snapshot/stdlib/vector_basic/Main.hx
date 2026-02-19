class Main {
	static function render(label:String, v:haxe.ds.Vector<Int>) {
		var out = "";
		var sum = 0;
		for (i in 0...v.length) {
			if (i > 0) {
				out += ",";
			}
			out += Std.string(v[i]);
			sum += v[i];
		}
		Sys.println(label + ":" + out);
		Sys.println(label + "_sum:" + sum);
	}

	static function main() {
		var v = new haxe.ds.Vector<Int>(4);
		v[0] = 3;
		v[1] = 1;
		v[2] = 4;
		v[3] = 1;

		Sys.println("len:" + v.length);
		render("base", v);

		v[1] = 9;
		Sys.println("len_after_set:" + v.length);
		render("mut", v);

		var w = new haxe.ds.Vector<Int>(6);
		for (i in 0...w.length) {
			w[i] = (i + 1) * 2;
		}
		Sys.println("w_len:" + w.length);
		render("w", w);
	}
}
