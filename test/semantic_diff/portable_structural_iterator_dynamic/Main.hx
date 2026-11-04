class Cursor {
	var index = 0;

	public function new() {}

	public dynamic function hasNext():Bool
		return index < 1;

	public dynamic function next():Int
		return index++;
}

class Main {
	static function main() {
		final source = new Cursor();
		final view:Iterator<Int> = source;
		Sys.println(view.hasNext());
		Sys.println(view.next());
		Sys.println(view.hasNext());
		final saved = view.next;
		Sys.println(saved == view.next);
		source.next = () -> 23;
		Sys.println(saved == view.next);
		Sys.println(saved());
		Sys.println(view.next());
	}
}
