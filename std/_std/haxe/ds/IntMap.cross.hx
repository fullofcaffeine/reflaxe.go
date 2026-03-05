package haxe.ds;

extern class IntMap<T> implements haxe.Constraints.IMap<Int, T> {
	function new():Void;
	function set(key:Int, value:T):Void;
	function get(key:Int):Null<T>;
	function exists(key:Int):Bool;
	function remove(key:Int):Bool;
	function keys():Iterator<Int>;
	function iterator():Iterator<T>;

	@:runtime inline function keyValueIterator():KeyValueIterator<Int, T> {
		return new haxe.iterators.MapKeyValueIterator(this);
	}

	function copy():IntMap<T>;
	function toString():String;
	function clear():Void;
}
