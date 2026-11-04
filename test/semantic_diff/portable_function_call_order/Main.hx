class Main {
	static function main() {
		var callback:Int->Int = _ -> 1;
		function replace():Int {
			callback = _ -> 2;
			return 0;
		}
		Sys.println(callback(replace()));
		var earlier = 3;
		function mutate():Int {
			earlier = 9;
			return 4;
		}
		final pair = (left:Int, right:Int) -> left * 10 + right;
		Sys.println(pair(earlier, mutate()));
		var absent:Int->Int = null;
		function install():Int {
			absent = _ -> 8;
			return 0;
		}
		try {
			Sys.println(absent(install()));
		} catch (error:haxe.Exception) {
			Sys.println(error.message);
		}
		var present:Int->Int = _ -> 7;
		function clear():Int {
			present = null;
			return 0;
		}
		try {
			Sys.println(present(clear()));
		} catch (error:haxe.Exception) {
			Sys.println(error.message);
		}
	}
}
