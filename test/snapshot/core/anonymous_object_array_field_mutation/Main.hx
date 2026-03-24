class Main {
	static function main() {
		var box = {items: [1, 2]};
		box.items.push(3);
		Sys.println(box.items.length);
		Sys.println(box.items[2]);
		box.items.pop();
		Sys.println(box.items.length);

		var nested = {inner: {items: ["a"]}};
		nested.inner.items.push("b");
		Sys.println(nested.inner.items.length);
		Sys.println(nested.inner.items[1]);
		nested.inner.items.pop();
		Sys.println(nested.inner.items.length);
	}
}
