class Holder<T> {
	var value:Iterator<T>;

	public function new(value:Iterator<T>) {
		this.value = value;
	}

	public function read():Iterator<T>
		return value;
}

class Main {
	static function same(left:Dynamic, right:Dynamic):Bool
		return left == right;

	static function main() {
		final source = {hasNext: () -> true, next: () -> 3};
		final holder = new Holder<Int>(source);
		final view = holder.read();
		final saved = view.next;
		Sys.println(same(view, source));
		Sys.println(saved == view.next);
		Sys.println(saved());
		source.next = () -> 8;
		Sys.println(view.next());
		Sys.println(saved());
		Sys.println(view.next == source.next);
		final missing:Iterator<Int> = null;
		Sys.println(new Holder<Int>(missing).read() == null);
		final nullable = {hasNext: () -> true, next: () -> (null : Null<Int>)};
		final nullableView = new Holder<Null<Int>>(nullable).read();
		Sys.println(nullableView.next() == null);
		nullable.next = () -> 5;
		Sys.println(nullableView.next());
	}
}
