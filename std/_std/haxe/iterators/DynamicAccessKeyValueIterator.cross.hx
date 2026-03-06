package haxe.iterators;

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
