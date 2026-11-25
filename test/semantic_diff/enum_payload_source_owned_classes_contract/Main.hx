/**
	Each payload type is a standard-library class that only this enum names.
	The compiler lowers enums before classes, so these classes are required
	while the enum is lowered. They must still be generated afterwards, or the
	enum's Go declaration names undefined types.

	The switch below has no payload cases on purpose: a payload case lowers a
	type test in `Main` that would require the classes again and hide the bug.
**/
enum Box {
	Empty;
	Stack(stack:haxe.ds.GenericStack<Int>);
	Tree(tree:haxe.ds.BalancedTree<String, Int>);
	Buffer(buffer:haxe.io.BytesBuffer);
}

class Main {
	static function label(box:Box):String {
		return switch (box) {
			case Empty: "empty";
			case _: "holder";
		}
	}

	static function main() {
		var box = Box.Empty;
		Sys.println("constructor=" + Type.enumConstructor(box));
		Sys.println("index=" + Type.enumIndex(box));
		Sys.println("label=" + label(box));
		Sys.println("is Empty=" + (box == Box.Empty));
		Sys.println("constructs=" + Type.getEnumConstructs(Box).join(","));
	}
}
