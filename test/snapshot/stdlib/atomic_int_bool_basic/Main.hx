import haxe.atomic.AtomicBool;
import haxe.atomic.AtomicInt;

class Main {
	static function out(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main() {
		var atom = new AtomicInt(10);
		out("int.load.0", atom.load());
		out("int.add.old", atom.add(5));
		out("int.load.1", atom.load());
		out("int.sub.old", atom.sub(2));
		out("int.load.2", atom.load());
		out("int.and.old", atom.and(6));
		out("int.load.3", atom.load());
		out("int.or.old", atom.or(8));
		out("int.load.4", atom.load());
		out("int.xor.old", atom.xor(10));
		out("int.load.5", atom.load());
		out("int.cmp.miss.old", atom.compareExchange(7, 100));
		out("int.cmp.miss.now", atom.load());
		out("int.cmp.hit.old", atom.compareExchange(6, 11));
		out("int.cmp.hit.now", atom.load());
		out("int.xchg.old", atom.exchange(3));
		out("int.xchg.now", atom.load());
		out("int.store.ret", atom.store(42));
		out("int.store.now", atom.load());

		var flag = new AtomicBool(false);
		out("bool.load.0", flag.load());
		out("bool.cmp.miss.old", flag.compareExchange(true, false));
		out("bool.cmp.miss.now", flag.load());
		out("bool.cmp.hit.old", flag.compareExchange(false, true));
		out("bool.cmp.hit.now", flag.load());
		out("bool.xchg.old", flag.exchange(false));
		out("bool.xchg.now", flag.load());
		out("bool.store.ret", flag.store(true));
		out("bool.store.now", flag.load());
	}
}
