package go;

class Slice<T> {
  public var data(default, null):Array<T>;

  public function new() {
    data = [];
  }

  public var length(get, never):Int;

  function get_length():Int {
    return data.length;
  }

  public function push(value:T):Void {
    data.push(value);
  }

  public function get(index:Int):T {
    return data[index];
  }

  public function set(index:Int, value:T):Void {
    data[index] = value;
  }

  public function toArray():Array<T> {
    return data;
  }
}
