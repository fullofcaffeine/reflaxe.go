package haxe.ds;

/**
	What
	A staged `haxe.ds.GenericStack` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	The upstream implementation is the correct semantic contract, but its current
	shape still relies on array `join` lowering and inline generic field paths
	that `haxe.go` does not yet handle cleanly for direct source-owned runtime
	use. That left direct `GenericStack` usage as compile-only debt.

	How
	Keep the upstream public API and iterator behavior, but express `first`,
	`pop`, and `toString` in a form that lowers predictably on `haxe.go` today.
	This keeps the ownership in staged std code instead of pushing collection
	semantics into `GoCompiler`.
**/
class GenericCell<T> {
	public var elt:T;
	public var next:GenericCell<T>;

	public function new(elt:T, next:GenericCell<T>) {
		this.elt = elt;
		this.next = next;
	}
}

/**
	What
	A staged `haxe.ds.GenericStack` override for `haxe.go`.

	Why
	The upstream implementation is semantically correct, but direct source-owned
	compilation currently trips backend gaps around inline generic field access and
	array `join` lowering in `toString`.

	How
	Preserve the public API and iterator semantics while using explicit local
	control flow and `StringBuf` rendering that already lower cleanly on Go.
**/
class GenericStack<T> {
	public var head:GenericCell<T>;

	public function new() {}

	public inline function add(item:T) {
		head = new GenericCell<T>(item, head);
	}

	public function first():Null<T> {
		return head == null ? null : head.elt;
	}

	public function pop():Null<T> {
		var current = head;
		if (current == null) {
			return null;
		}
		head = current.next;
		return current.elt;
	}

	public inline function isEmpty():Bool {
		return head == null;
	}

	public function remove(v:T):Bool {
		var prev:GenericCell<T> = null;
		var current = head;
		while (current != null) {
			if (sameValue(current.elt, v)) {
				if (prev == null) {
					head = current.next;
				} else {
					prev.next = current.next;
				}
				return true;
			}
			prev = current;
			current = current.next;
		}
		return false;
	}

	/**
		Why
		`GenericStack<T>` erases element reads on `haxe.go` down to a runtime value
		boundary, so direct `==` here can degrade into Go interface identity for
		String-backed values instead of Haxe value equality.

		How
		Keep the `Dynamic` usage local to this staged std override and only use it to
		restore Haxe-style equality on erased generic values. Once generic equality
		semantics are handled centrally, this helper can be removed.
	**/
	static function sameValue(left:Dynamic, right:Dynamic):Bool {
		if (left == null || right == null) {
			return left == right;
		}
		if (Std.isOfType(left, String) || Std.isOfType(right, String)) {
			return Std.string(left) == Std.string(right);
		}
		return left == right;
	}

	public function iterator():Iterator<T> {
		var current = head;
		return {
			hasNext: function() {
				return current != null;
			},
			next: function() {
				var cell = current;
				current = cell.next;
				return cell.elt;
			}
		};
	}

	public function toString():String {
		var out = new StringBuf();
		out.add("{");
		var current = head;
		var firstItem = true;
		while (current != null) {
			if (!firstItem) {
				out.add(",");
			}
			out.add(Std.string(current.elt));
			firstItem = false;
			current = current.next;
		}
		out.add("}");
		return out.toString();
	}
}
