import haxe.Timer;

class Main {
	static function main() {
		var started = Timer.stamp();
		Sys.sleep(0.02);
		var elapsed = Timer.stamp() - started;
		Sys.println(elapsed >= 0.005);
		Sys.println(elapsed < 5.0);
	}
}
