package haxe.ds;

/**
	What
	- Implements the complete `haxe.ds.List<T>` API as ordinary staged Haxe source
	  backed by an array.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on
	  `haxe.go` yet because it uses private linked nodes, while the still-open
	  collection-adapter and serializer lowerings tracked by `haxe_go-vfp.8.7.17`
	  consume the compact `items` carrier. The public list behavior itself does not
	  require compiler ownership.

	How
	- Keep the compatibility carrier private, implement every public algorithm in
	  Haxe, return concrete source-owned iterators, and retain the class methods needed by the
	  transitional serializer. Once `haxe_go-vfp.8.7.17` removes representation
	  reads, this override can be reconsidered against the upstream linked list.
**/
@:keep
class List<T> {
	private var items:Array<T>;

	public var length(default, null):Int;

	public function new() {
		items = [];
		length = 0;
	}

	public function add(item:T):Void {
		items.push(item);
		length++;
	}

	public function push(item:T):Void {
		var next = [item];
		for (existing in items) {
			next.push(existing);
		}
		items = next;
		length++;
	}

	public function first():Null<T> {
		return length == 0 ? null : items[0];
	}

	public function last():Null<T> {
		return length == 0 ? null : items[length - 1];
	}

	public function pop():Null<T> {
		if (length == 0) {
			return null;
		}
		var first = items[0];
		var remaining = new Array<T>();
		for (index in 1...items.length) {
			remaining.push(items[index]);
		}
		items = remaining;
		length--;
		return first;
	}

	public inline function isEmpty():Bool {
		return length == 0;
	}

	public function clear():Void {
		items = [];
		length = 0;
	}

	public function remove(value:T):Bool {
		for (index in 0...items.length) {
			if (sameValue(items[index], value)) {
				var remaining = new Array<T>();
				for (copyIndex in 0...items.length) {
					if (copyIndex != index) {
						remaining.push(items[copyIndex]);
					}
				}
				items = remaining;
				length--;
				return true;
			}
		}
		return false;
	}

	/**
		What
		- Compares values read through the erased generic list carrier.

		Why
		- Generated Go stores `List<T>` elements as erased values, so direct equality
		  compares String pointers instead of preserving Haxe String value equality.

		How
		- Localize `Dynamic` to this generic boundary, preserve nulls, restore String
		  comparison by value, and use ordinary Haxe equality for other supported
		  element shapes. Generic target equality can replace this helper later.
	**/
	private static function sameValue(left:Dynamic, right:Dynamic):Bool {
		if (left == null || right == null) {
			return left == right;
		}
		if (Std.isOfType(left, String) || Std.isOfType(right, String)) {
			return Std.string(left) == Std.string(right);
		}
		return left == right;
	}

	public function iterator():GoListIterator<T> {
		return new GoListIterator(items);
	}

	@:pure @:runtime public function keyValueIterator():GoListKeyValueIterator<T> {
		return new GoListKeyValueIterator(items);
	}

	public function toString():String {
		var rendered = new Array<String>();
		for (index in 0...items.length) {
			rendered.push(Std.string(items[index]));
		}
		return "{" + rendered.join(", ") + "}";
	}

	public function join(separator:String):String {
		var rendered = new Array<String>();
		for (index in 0...items.length) {
			rendered.push(Std.string(items[index]));
		}
		return rendered.join(separator);
	}

	public function filter(predicate:T->Bool):List<T> {
		var filtered = new List<T>();
		var source = iterator();
		while (source.hasNext()) {
			var item = source.next();
			if (predicate(item)) {
				filtered.add(item);
			}
		}
		return filtered;
	}

	public function map<Result>(transform:T->Result):List<Result> {
		var mapped = new List<Result>();
		var source = iterator();
		while (source.hasNext()) {
			var item = source.next();
			mapped.add(transform(item));
		}
		return mapped;
	}
}

/**
	What
	- Provides the concrete iterator carrier for staged `List` values.

	Why
	- Anonymous structural iterators lower to dynamic field maps on `haxe.go`, but
	  transitional framework adapters call `hasNext` and `next` as Go methods.

	How
	- Snapshot the backing array reference and advance a typed source-level index.
**/
private class GoListIterator<T> {
	private final items:Array<T>;
	private var index:Int;

	public function new(items:Array<T>) {
		this.items = items;
		index = 0;
	}

	public function hasNext():Bool {
		return index < items.length;
	}

	public function next():T {
		return items[index++];
	}
}

/**
	What
	- Provides indexed iteration for staged `List` values.

	Why
	- A concrete iterator keeps the Go method boundary typed while preserving the
	  ordinary Haxe `{key, value}` result expected by `KeyValueIterator` consumers.

	How
	- Advance one index per call and construct the public key/value record in Haxe.
**/
private class GoListKeyValueIterator<T> {
	private final items:Array<T>;
	private var index:Int;

	public function new(items:Array<T>) {
		this.items = items;
		index = 0;
	}

	public function hasNext():Bool {
		return index < items.length;
	}

	public function next():{key:Int, value:T} {
		var key = index++;
		return {key: key, value: items[key]};
	}
}
