typedef Slot<T> = {var next:Void->T;}

class Main {
	static function replace<T>(value:Slot<T>, next:Void->T):Void {
		Reflect.setField(value, "next", next);
	}

	static function main() {
		var source:Slot<Int> = {next: () -> 3};
		replace(source, () -> 4);
		Sys.println(source.next());
	}
}
