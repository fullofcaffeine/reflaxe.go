package go;

class Map<K, V> {
  final inner:haxe.ds.StringMap<V>;

  public function new() {
    inner = new haxe.ds.StringMap<V>();
  }

  public function set(key:K, value:V):Void {
    inner.set(Std.string(key), value);
  }

  public function get(key:K):Null<V> {
    return inner.get(Std.string(key));
  }

  public function exists(key:K):Bool {
    return inner.exists(Std.string(key));
  }
}
