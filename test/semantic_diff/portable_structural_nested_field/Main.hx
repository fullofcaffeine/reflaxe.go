typedef Outer<T> = {var inner:{var next:Void->T;}; var ?link:Outer<T>;}

class Main {
	static function through<T>(value:Outer<T>):Outer<T>
		return value;

	static function read<T>(value:Outer<T>):T
		return value.inner.next();

	static function main() {
		var source:Outer<Int> = {inner: {next: () -> 3}};
		source.link = source;
		var view = through(source);
		Sys.println(view.inner.next());
		Sys.println(read(source));
		source.inner.next = () -> 9;
		Sys.println(view.inner.next());
		Sys.println(read(source));
		Sys.println(view.link == source);
		Sys.println(read(view.link));
	}
}
