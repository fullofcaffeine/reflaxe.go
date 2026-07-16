package haxe.ds;

import hxrt.collections.NativeEnumValue;

/**
	What
	- Implements `EnumValueMap` as an ordinary staged AVL tree with recursive enum
	  value comparison.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` because its covariant `BalancedTree.copy()` override does not satisfy
	  Go's exact interface method rules, and its dynamic comparison paths currently
	  lower without the type assertions Go requires.

	How
	- Preserve the upstream balanced-tree algorithms and comparison contract in a
	  standalone source-owned class. Use one typed runtime predicate to recognize nested enum
	  parameters, explicit casts at dynamic array/enum boundaries, and kept `IMap`
	  bridge methods for Go's erased interface method set.
**/
class EnumValueMap<K:EnumValue, V> implements haxe.Constraints.IMap<K, V> {
	private var root:EnumValueTreeNode<K, V>;

	public function new() {}

	public function set(key:K, value:V):Void {
		root = setLoop(key, value, root);
	}

	public function get(key:K):Null<V> {
		var node = root;
		while (node != null) {
			var result = compare(key, node.key);
			if (result == 0) {
				return node.value;
			}
			node = result < 0 ? node.left : node.right;
		}
		return null;
	}

	public function exists(key:K):Bool {
		var node = root;
		while (node != null) {
			var result = compare(key, node.key);
			if (result == 0) {
				return true;
			}
			node = result < 0 ? node.left : node.right;
		}
		return false;
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

	public function keys():Iterator<K> {
		var values = keysLoop(root, []);
		var index = 0;
		return {
			hasNext: function() {
				return index < values.length;
			},
			next: function() {
				return values[index++];
			}
		};
	}

	public function iterator():Iterator<V> {
		var values = valuesLoop(root, []);
		var index = 0;
		return {
			hasNext: function() {
				return index < values.length;
			},
			next: function() {
				return values[index++];
			}
		};
	}

	public function keyValueIterator():KeyValueIterator<K, V> {
		var keys = keys();
		return {
			hasNext: function() {
				return keys.hasNext();
			},
			next: function() {
				var key = keys.next();
				return {key: key, value: cast get(key)};
			}
		};
	}

	public function copy():EnumValueMap<K, V> {
		var copied = new EnumValueMap<K, V>();
		copied.root = root;
		return copied;
	}

	public function toString():String {
		return root == null ? "[]" : "[" + root.toString() + "]";
	}

	public function clear():Void {
		root = null;
	}

	private function compare(left:EnumValue, right:EnumValue):Int {
		var result = Type.enumIndex(left) - Type.enumIndex(right);
		if (result != 0) {
			return result;
		}
		return compareArgs(Type.enumParameters(left), Type.enumParameters(right));
	}

	private function compareArgs(left:Array<Dynamic>, right:Array<Dynamic>):Int {
		var result = left.length - right.length;
		if (result != 0) {
			return result;
		}
		for (index in 0...left.length) {
			result = compareArg(left[index], right[index]);
			if (result != 0) {
				return result;
			}
		}
		return 0;
	}

	private function compareArg(left:Dynamic, right:Dynamic):Int {
		if (isEnumValue(left) && isEnumValue(right)) {
			return compare(cast left, cast right);
		}
		if (Std.isOfType(left, Array) && Std.isOfType(right, Array)) {
			return compareArgs(cast left, cast right);
		}
		return Reflect.compare(left, right);
	}

	private static function isEnumValue(value:Dynamic):Bool {
		return NativeEnumValue.isEnumValue(value);
	}

	private function setLoop(key:K, value:V, node:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		if (node == null) {
			return new EnumValueTreeNode<K, V>(null, key, value, null, -1);
		}
		var result = compare(key, node.key);
		if (result == 0) {
			return new EnumValueTreeNode<K, V>(node.left, key, value, node.right, node.getHeight());
		}
		if (result < 0) {
			return balance(setLoop(key, value, node.left), node.key, node.value, node.right);
		}
		return balance(node.left, node.key, node.value, setLoop(key, value, node.right));
	}

	private function removeLoop(key:K, node:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		if (node == null) {
			throw "Not_found";
		}
		var result = compare(key, node.key);
		if (result == 0) {
			return merge(node.left, node.right);
		}
		if (result < 0) {
			return balance(removeLoop(key, node.left), node.key, node.value, node.right);
		}
		return balance(node.left, node.key, node.value, removeLoop(key, node.right));
	}

	private static function keysLoop<K:EnumValue, V>(node:EnumValueTreeNode<K, V>, out:Array<K>):Array<K> {
		if (node != null) {
			keysLoop(node.left, out);
			out.push(node.key);
			keysLoop(node.right, out);
		}
		return out;
	}

	private static function valuesLoop<K:EnumValue, V>(node:EnumValueTreeNode<K, V>, out:Array<V>):Array<V> {
		if (node != null) {
			valuesLoop(node.left, out);
			out.push(node.value);
			valuesLoop(node.right, out);
		}
		return out;
	}

	private function merge(left:EnumValueTreeNode<K, V>, right:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		if (left == null) {
			return right;
		}
		if (right == null) {
			return left;
		}
		var minimum = minBinding(right);
		return balance(left, minimum.key, minimum.value, removeMinBinding(right));
	}

	private function minBinding(node:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		if (node == null) {
			throw "Not_found";
		}
		return node.left == null ? node : minBinding(node.left);
	}

	private function removeMinBinding(node:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		return node.left == null ? node.right : balance(removeMinBinding(node.left), node.key, node.value, node.right);
	}

	private function balance(left:EnumValueTreeNode<K, V>, key:K, value:V, right:EnumValueTreeNode<K, V>):EnumValueTreeNode<K, V> {
		var leftHeight = left == null ? 0 : left.getHeight();
		var rightHeight = right == null ? 0 : right.getHeight();
		if (leftHeight > rightHeight + 2) {
			var leftLeftHeight = left.left == null ? 0 : left.left.getHeight();
			var leftRightHeight = left.right == null ? 0 : left.right.getHeight();
			if (leftLeftHeight >= leftRightHeight) {
				return new EnumValueTreeNode<K, V>(left.left, left.key, left.value, new EnumValueTreeNode<K, V>(left.right, key, value, right, -1), -1);
			}
			return new EnumValueTreeNode<K, V>(new EnumValueTreeNode<K, V>(left.left, left.key, left.value, left.right.left, -1), left.right.key,
				left.right.value, new EnumValueTreeNode<K, V>(left.right.right, key, value, right, -1), -1);
		}
		if (rightHeight > leftHeight + 2) {
			var rightRightHeight = right.right == null ? 0 : right.right.getHeight();
			var rightLeftHeight = right.left == null ? 0 : right.left.getHeight();
			if (rightRightHeight > rightLeftHeight) {
				return new EnumValueTreeNode<K, V>(new EnumValueTreeNode<K, V>(left, key, value, right.left, -1), right.key, right.value, right.right, -1);
			}
			return new EnumValueTreeNode<K, V>(new EnumValueTreeNode<K, V>(left, key, value, right.left.left, -1), right.left.key, right.left.value,
				new EnumValueTreeNode<K, V>(right.left.right, right.key, right.value, right.right, -1), -1);
		}
		return new EnumValueTreeNode<K, V>(left, key, value, right, (leftHeight > rightHeight ? leftHeight : rightHeight) + 1);
	}

	@:keep private function getIMap(key:Dynamic):Dynamic {
		return get(cast key);
	}

	@:keep private function setIMap(key:Dynamic, value:Dynamic):Void {
		set(cast key, cast value);
	}

	@:keep private function existsIMap(key:Dynamic):Bool {
		return exists(cast key);
	}

	@:keep private function removeIMap(key:Dynamic):Bool {
		return remove(cast key);
	}

	@:keep private function copyIMap():haxe.Constraints.IMap<K, V> {
		return copy();
	}
}

/**
	What
	- Stores one node of the staged `EnumValueMap` AVL tree.

	Why
	- Keeping nodes as typed Haxe source preserves enum-key/value types without an
	  untyped compiler carrier.

	How
	- Cache subtree height at construction and expose only the fields used by the
	  owning map's balancing and ordered traversal algorithms.
**/
private class EnumValueTreeNode<K:EnumValue, V> {
	public var left:EnumValueTreeNode<K, V>;
	public var right:EnumValueTreeNode<K, V>;
	public var key:K;
	public var value:V;

	private var height:Int;

	public function new(left:EnumValueTreeNode<K, V>, key:K, value:V, right:EnumValueTreeNode<K, V>, height:Int) {
		this.left = left;
		this.key = key;
		this.value = value;
		this.right = right;
		if (height == -1) {
			var leftHeight = left == null ? 0 : left.getHeight();
			var rightHeight = right == null ? 0 : right.getHeight();
			this.height = (leftHeight > rightHeight ? leftHeight : rightHeight) + 1;
		} else {
			this.height = height;
		}
	}

	public inline function getHeight():Int {
		return height;
	}

	public function toString():String {
		return (left == null ? "" : left.toString() + ", ")
			+ Std.string(key)
			+ " => "
			+ Std.string(value)
			+ (right == null ? "" : ", " + right.toString());
	}
}
