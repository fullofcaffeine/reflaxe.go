class Main {
	static function main() {
		Sys.println(Reflect.compare(2, 3));
		Sys.println(Reflect.compare(3, 3));
		Sys.println(Reflect.compare(4, 3));

		Sys.println(Reflect.compare("a", "b"));
		Sys.println(Reflect.compare("z", "a"));
	}
}
