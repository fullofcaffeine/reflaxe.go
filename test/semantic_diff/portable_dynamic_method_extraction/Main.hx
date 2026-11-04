class Main {
	public function new() {}

	public dynamic function next():Int
		return 1;

	static function main() {
		var source = new Main();
		var view:{function next():Int;} = source;
		var a = source.next;
		var b = source.next;
		Sys.println(a == b);
		Sys.println(Reflect.compareMethods(a, b));
		var c = view.next;
		Sys.println(a == c);
		Sys.println(Reflect.compareMethods(a, c));
		var replacement = () -> 2;
		source.next = replacement;
		var d = source.next;
		Sys.println(d == replacement);
		Sys.println(d == source.next);
		Sys.println(Reflect.compareMethods(d, replacement));
		Sys.println(a());
		Sys.println(d());
	}
}
