import haxe.atomic.AtomicBool;
import haxe.atomic.AtomicInt;

class Main {
	static function main() {
		var count = new AtomicInt(1);
		count.add(2);
		var flag = new AtomicBool(false);
		flag.exchange(true);
	}
}
