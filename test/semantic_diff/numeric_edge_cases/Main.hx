class Main {
	static function main() {
		var neg:Int = -3;
		Sys.println(neg >>> 1);

		var big:Int = 1 << 30;
		Sys.println(big >>> 2);

		Sys.println(-1 >>> 1);

		var a:Float = 0.1 + 0.2;
		Sys.println(a > 0.3);
	}
}
