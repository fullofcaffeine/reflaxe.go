package haxe.ds;

extern class StringMap<T> implements haxe.Constraints.IMap<String, T> {
	function new():Void;
	function set(key:String, value:T):Void;
	function get(key:String):Null<T>;
	function exists(key:String):Bool;
	function remove(key:String):Bool;
	function keys():Iterator<String>;
	function iterator():Iterator<T>;

	@:runtime inline function keyValueIterator():KeyValueIterator<String, T> {
		return new haxe.iterators.MapKeyValueIterator(this);
	}

	function copy():StringMap<T>;
	function toString():String;
	function clear():Void;
}
