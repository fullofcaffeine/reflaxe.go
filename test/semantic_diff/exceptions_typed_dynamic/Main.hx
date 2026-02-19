class Main {
	static function raise(v:Int):Void {
		if (v == 0) {
			throw "text";
		}
		if (v == 1) {
			throw 11;
		}
		throw true;
	}

	static function handle(v:Int):Void {
		try {
			raise(v);
		} catch (s:String) {
			Sys.println("S:" + s);
		} catch (i:Int) {
			Sys.println(i);
		} catch (e:Dynamic) {
			Sys.println("D");
		}
	}

	static function main() {
		handle(0);
		handle(1);
		handle(2);
	}
}
