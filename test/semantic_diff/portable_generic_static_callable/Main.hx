typedef Record = {var value:Int;}

class Main {
	static function identity<T>(value:T):T
		return value;

	static function callback<T>(value:T->T):T->T
		return value;

	static function main() {
		final typed:Int->Int = identity;
		Sys.println(typed(7));
		final values = ["b", "a"];
		values.sort(Reflect.compare);
		Sys.println(values.join(","));
		final source:Record = {value: 9};
		final result:Record = identity(source);
		Sys.println(result.value);
		Sys.println(result == source);
		final cb = callback((value:Int) -> value + 1);
		Sys.println(cb(4));
	}
}
