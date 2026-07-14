package haxe.ds;

/**
	What
	A staged `haxe.ds.BalancedTree` override for `haxe.go`.

	Why
	The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`.
	Direct `BalancedTree` runtime use currently needs source-owned std emission,
	but the upstream implementation still exposes a Go-lowering paper-cut in
	`remove`: the `try/catch` + direct `return` shape can leave the generated Go
	function without an explicit return path. The backend also had a legacy shim
	name collision that this module now replaces cleanly.

	How
	Keep the upstream tree structure, ordering semantics, and iterator behavior,
	but rewrite `remove` into an explicit success flag. This preserves ownership in
	staged std code and keeps the compiler focused on lowering rather than
	re-implementing collection logic.
**/
class BalancedTree<K, V> implements haxe.Constraints.IMap<K, V> {
	var root:TreeNode<K, V>;

	public function new() {}

	public function set(key:K, value:V) {
		root = setLoop(key, value, root);
	}

	public function get(key:K):Null<V> {
		var node = root;
		while (node != null) {
			var c = compare(key, node.key);
			if (c == 0)
				return node.value;
			if (c < 0)
				node = node.left;
			else
				node = node.right;
		}
		return null;
	}

	public function remove(key:K):Bool {
		var removed = true;
		try {
			root = removeLoop(key, root);
		} catch (_:String) {
			removed = false;
		}
		return removed;
	}

	public function exists(key:K):Bool {
		var node = root;
		while (node != null) {
			var c = compare(key, node.key);
			if (c == 0)
				return true;
			else if (c < 0)
				node = node.left;
			else
				node = node.right;
		}
		return false;
	}

	public function iterator():Iterator<V> {
		var ret = iteratorLoop(root, []);
		var index = 0;
		return {
			hasNext: function() {
				return index < ret.length;
			},
			next: function() {
				return ret[index++];
			}
		};
	}

	public function keys():Iterator<K> {
		var ret = keysLoop(root, []);
		var index = 0;
		return {
			hasNext: function() {
				return index < ret.length;
			},
			next: function() {
				return ret[index++];
			}
		};
	}

	public function keyValueIterator():KeyValueIterator<K, V> {
		var keyIterator = keys();
		return {
			hasNext: function() {
				return keyIterator.hasNext();
			},
			next: function() {
				var key = keyIterator.next();
				return {
					key: key,
					value: get(key)
				};
			}
		};
	}

	public function copy():BalancedTree<K, V> {
		var copied = new BalancedTree<K, V>();
		copied.root = root;
		return copied;
	}

	function setLoop(k:K, v:V, node:TreeNode<K, V>) {
		if (node == null)
			return new TreeNode<K, V>(null, k, v, null, -1);
		var c = compare(k, node.key);
		return if (c == 0) new TreeNode<K, V>(node.left, k, v, node.right, node.get_height()) else if (c < 0) {
			var nl = setLoop(k, v, node.left);
			balance(nl, node.key, node.value, node.right);
		} else {
			var nr = setLoop(k, v, node.right);
			balance(node.left, node.key, node.value, nr);
		}
	}

	function removeLoop(k:K, node:TreeNode<K, V>) {
		if (node == null)
			throw "Not_found";
		var c = compare(k, node.key);
		return if (c == 0) merge(node.left,
			node.right) else if (c < 0) balance(removeLoop(k, node.left), node.key, node.value,
			node.right) else balance(node.left, node.key, node.value, removeLoop(k, node.right));
	}

	static function iteratorLoop<K, V>(node:TreeNode<K, V>, acc:Array<V>):Array<V> {
		if (node != null) {
			acc = iteratorLoop(node.left, acc);
			acc.push(node.value);
			acc = iteratorLoop(node.right, acc);
		}
		return acc;
	}

	function keysLoop(node:TreeNode<K, V>, acc:Array<K>):Array<K> {
		if (node != null) {
			acc = keysLoop(node.left, acc);
			acc.push(node.key);
			acc = keysLoop(node.right, acc);
		}
		return acc;
	}

	function merge(t1, t2) {
		if (t1 == null)
			return t2;
		if (t2 == null)
			return t1;
		var t = minBinding(t2);
		return balance(t1, t.key, t.value, removeMinBinding(t2));
	}

	function minBinding(t:TreeNode<K, V>) {
		return if (t == null) throw "Not_found" else if (t.left == null) t else minBinding(t.left);
	}

	function removeMinBinding(t:TreeNode<K, V>) {
		return if (t.left == null) t.right else balance(removeMinBinding(t.left), t.key, t.value, t.right);
	}

	function balance(l:TreeNode<K, V>, k:K, v:V, r:TreeNode<K, V>):TreeNode<K, V> {
		var hl = l.get_height();
		var hr = r.get_height();
		return if (hl > hr + 2) {
			if (l.left.get_height() >= l.right.get_height())
				new TreeNode<K, V>(l.left, l.key, l.value, new TreeNode<K, V>(l.right, k, v, r, -1), -1);
			else
				new TreeNode<K, V>(new TreeNode<K, V>(l.left, l.key, l.value, l.right.left, -1), l.right.key, l.right.value,
					new TreeNode<K, V>(l.right.right, k, v, r, -1), -1);
		} else if (hr > hl + 2) {
			if (r.right.get_height() > r.left.get_height())
				new TreeNode<K, V>(new TreeNode<K, V>(l, k, v, r.left, -1), r.key, r.value, r.right, -1);
			else
				new TreeNode<K, V>(new TreeNode<K, V>(l, k, v, r.left.left, -1), r.left.key, r.left.value,
					new TreeNode<K, V>(r.left.right, r.key, r.value, r.right, -1), -1);
		} else {
			new TreeNode<K, V>(l, k, v, r, (hl > hr ? hl : hr) + 1);
		}
	}

	function compare(k1:K, k2:K) {
		return Reflect.compare(k1, k2);
	}

	public function toString():String {
		return root == null ? "[]" : "[" + root.toString() + "]";
	}

	public function clear():Void {
		root = null;
	}
}

class TreeNode<K, V> {
	public var left:TreeNode<K, V>;
	public var right:TreeNode<K, V>;
	public var key:K;
	public var value:V;

	var _height:Int;

	public function new(l, k, v, r, h:Int) {
		left = l;
		key = k;
		value = v;
		right = r;
		if (h == -1)
			_height = (left.get_height() > right.get_height() ? left.get_height() : right.get_height()) + 1;
		else
			_height = h;
	}

	extern public inline function get_height()
		return this == null ? 0 : _height;

	public function toString():String {
		return (left == null ? "" : left.toString() + ", ")
			+ Std.string(key)
			+ " => "
			+ Std.string(value)
			+ (right == null ? "" : ", " + right.toString());
	}
}
