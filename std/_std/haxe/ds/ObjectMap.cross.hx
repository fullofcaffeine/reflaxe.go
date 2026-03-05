package haxe.ds;

extern class ObjectMap<K:{}, V> implements haxe.Constraints.IMap<K, V> {
	function new():Void;
	function set(key:K, value:V):Void;
	function get(key:K):Null<V>;
	function exists(key:K):Bool;
	function remove(key:K):Bool;
	function keys():Iterator<K>;
	function iterator():Iterator<V>;

	@:runtime inline function keyValueIterator():KeyValueIterator<K, V> {
		return new haxe.iterators.MapKeyValueIterator(this);
	}

	function copy():ObjectMap<K, V>;
	function toString():String;
	function clear():Void;
}
