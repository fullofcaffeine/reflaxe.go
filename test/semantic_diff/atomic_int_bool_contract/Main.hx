#if go
import haxe.atomic.AtomicBool;
import haxe.atomic.AtomicInt;

typedef CompatAtomicBool = AtomicBool;
typedef CompatAtomicInt = AtomicInt;
#else
class CompatAtomicInt {
	var value:Int;

	public function new(value:Int) {
		this.value = value;
	}

	public function add(b:Int):Int {
		var previous = value;
		value += b;
		return previous;
	}

	public function sub(b:Int):Int {
		var previous = value;
		value -= b;
		return previous;
	}

	public function and(b:Int):Int {
		var previous = value;
		value &= b;
		return previous;
	}

	public function or(b:Int):Int {
		var previous = value;
		value |= b;
		return previous;
	}

	public function xor(b:Int):Int {
		var previous = value;
		value ^= b;
		return previous;
	}

	public function compareExchange(expected:Int, replacement:Int):Int {
		var previous = value;
		if (previous == expected) {
			value = replacement;
		}
		return previous;
	}

	public function exchange(next:Int):Int {
		var previous = value;
		value = next;
		return previous;
	}

	public function load():Int {
		return value;
	}

	public function store(next:Int):Int {
		value = next;
		return next;
	}
}

class CompatAtomicBool {
	static inline function toInt(value:Bool):Int {
		return value ? 1 : 0;
	}

	static inline function toBool(value:Int):Bool {
		return value == 1;
	}

	var inner:CompatAtomicInt;

	public function new(value:Bool) {
		inner = new CompatAtomicInt(toInt(value));
	}

	public function compareExchange(expected:Bool, replacement:Bool):Bool {
		return toBool(inner.compareExchange(toInt(expected), toInt(replacement)));
	}

	public function exchange(value:Bool):Bool {
		return toBool(inner.exchange(toInt(value)));
	}

	public function load():Bool {
		return toBool(inner.load());
	}

	public function store(value:Bool):Bool {
		return toBool(inner.store(toInt(value)));
	}
}
#end

class Main {
	static function out(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main() {
		var atom = new CompatAtomicInt(10);
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

		var flag = new CompatAtomicBool(false);
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
