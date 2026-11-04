typedef Slot<T> = {var next:Void->T;}

class Holder<T> {
	public var callback:T->T;

	public function new(callback:T->T)
		this.callback = callback;
}

class CallbackReader<T> {
	public var read:(T->T)->Int;

	public function new(read:(T->T)->Int)
		this.read = read;
}

class Main {
	static function replace<T>(value:Slot<T>, next:Void->T):Void {
		Reflect.setField(value, "next", next);
	}

	static function main() {
		var source:Slot<Int> = {next: () -> 3};
		replace(source, () -> 4);
		Sys.println(source.next());
		var holder = new Holder<Int>(value -> value + 1);
		Reflect.setField(holder, "callback", (value:Int) -> value + 2);
		Sys.println(holder.callback(3));
		var reader = new CallbackReader<Int>(callback -> callback(3));
		Sys.println(reader.read(value -> value + 1));
		var empty:Int->Int = null;
		Reflect.setField(holder, "callback", empty);
		Sys.println(holder.callback == null);
	}
}
