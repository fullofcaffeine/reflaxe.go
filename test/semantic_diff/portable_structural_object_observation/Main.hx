typedef NextOnly = {function next():Int;}

class Cursor {
	public var tag = "cursor";
	public var position = 0;

	public function new() {}

	public function hasNext():Bool
		return false;

	public function next():Int
		return position++;

	public function toString():String
		return tag;
}

class Main {
	static function same(left:Dynamic, right:Dynamic):Bool
		return left == right;

	static function main() {
		final source = new Cursor();
		final view:Iterator<Int> = source;
		Sys.println(Reflect.field(view, "tag"));
		Sys.println(Reflect.hasField(view, "tag"));
		final names = Reflect.fields(view);
		names.sort(Reflect.compare);
		Sys.println(names.join(","));
		Sys.println(Std.string(view));
		Sys.println(Type.getClassName(Type.getClass(view)) == Type.getClassName(Type.getClass(source)));
		Sys.println(Std.isOfType(view, Cursor));
		// Erased keys prove source identity across statically different views.
		final keys = new haxe.ds.ObjectMap<Dynamic, Int>();
		keys.set(source, 11);
		Sys.println(keys.get(view));
		Sys.println(Reflect.compareMethods(Reflect.field(view, "next"), source.next));
		Sys.println(haxe.Json.stringify(view) == haxe.Json.stringify(source));
		final copied = Reflect.copy(view);
		Sys.println(same(copied, source));
		Sys.println(Reflect.field(copied, "tag"));
		Reflect.setField(copied, "tag", "copy");
		Sys.println(source.tag);
		final callback = () -> 11;
		final rich = {tag: "extra", hasNext: () -> false, next: callback};
		final projected:NextOnly = rich;
		Sys.println(same(rich, projected));
		Sys.println(Reflect.field(projected, "tag"));
		final clone = Reflect.copy(projected);
		Sys.println(same(clone, projected));
		Sys.println(clone.next == callback);
		rich.next = () -> 22;
		Sys.println(projected.next());
		Sys.println(clone.next());
		final values = [1, 2];
		final cursor:Iterator<Int> = values.iterator();
		final a = cursor.next;
		final b = cursor.next;
		Sys.println(a == b);
		Sys.println(Reflect.compareMethods(a, b));
		Sys.println(cursor.next());
		final cursorFields = Reflect.fields(cursor);
		cursorFields.sort(Reflect.compare);
		Sys.println(cursorFields.join(","));
		values[1] = 9;
		Sys.println(cursor.next());
	}
}
