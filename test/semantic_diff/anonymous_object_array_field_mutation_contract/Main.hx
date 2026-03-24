class Main {
	static function mutate(box:{var items:Array<Int>;}):Void {
		box.items.push(7);
		box.items.push(9);
		box.items.pop();
	}

	static function main() {
		var box = {items: [1, 2]};
		mutate(box);
		Sys.println("box.len=" + box.items.length);
		Sys.println("box.last=" + box.items[box.items.length - 1]);

		var nested = {inner: {items: ["x"]}};
		nested.inner.items.push("y");
		var removed = nested.inner.items.pop();
		Sys.println("nested.removed=" + removed);
		Sys.println("nested.len=" + nested.inner.items.length);
		Sys.println("nested.only=" + nested.inner.items[0]);
	}
}
