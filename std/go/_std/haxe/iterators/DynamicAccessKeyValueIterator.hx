package haxe.iterators;

/**
	What:
	- Owns key/value iteration for the Go-target `DynamicAccess` representation.

	Why:
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`. because the generated dynamic-access carrier requires an explicit key snapshot and reflective field lookup.

	How:
	- Snapshot keys once, advance a typed index, and return the upstream `{key, value}` shape.
**/
class DynamicAccessKeyValueIterator<T> {
	final access:Dynamic;
	final keys:Array<String>;
	var index:Int;

	public inline function new(access:DynamicAccess<T>) {
		this.access = access;
		this.keys = (cast access : DynamicAccess<Dynamic>).keys();
		index = 0;
	}

	public inline function hasNext():Bool {
		return index < keys.length;
	}

	public inline function next():{key:String, value:Dynamic} {
		var key = keys[index++];
		return {value: Reflect.field(access, key), key: key};
	}
}
