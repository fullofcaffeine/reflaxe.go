import haxe.atomic.AtomicObject;

private class AtomicBox {
	public final value:Int;

	public function new(value:Int) {
		this.value = value;
	}
}

class Main {
	static function main() {
		var first = new AtomicBox(1);
		var second = new AtomicBox(2);
		var value = new AtomicObject(first);
		value.compareExchange(first, second);
	}
}
