class Holder<T> {
	final value:T;

	public function new(value:T) {
		this.value = value;
	}

	public function reader():Void->T {
		return () -> value;
	}
}

class Main {
	static function main() {
		final read = new Holder<Int>(19).reader();
		final alias = read;
		Sys.println(read());
		Sys.println(Reflect.compareMethods(read, alias));
	}
}
