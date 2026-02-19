import haxe.ds.Option;

class Main {
	static function render(opt:Option<Int>):String {
		return switch (opt) {
			case Some(v): "some:" + v;
			case None: "none";
		}
	}

	static function main() {
		Sys.println(render(Some(7)));
		Sys.println(render(None));
	}
}
