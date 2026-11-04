typedef Slot<T> = {var next:Void->T;}

class Main {
	static function duplicate<T>(value:Slot<T>):Slot<T>
		return Reflect.copy(value);

	static function main() {
		var source:Slot<Int> = {next: () -> 3};
		var copied = duplicate(source);
		Sys.println(copied.next());
		source.next = () -> 9;
		Sys.println(copied.next());
	}
}
