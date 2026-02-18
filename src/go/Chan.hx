package go;

class Chan<T> {
  var queue:Array<T>;
  var readIndex:Int;

  public function new() {
    queue = [];
    readIndex = 0;
  }

  public function send(value:T):Void {
    queue.push(value);
  }

  public function recv():Null<T> {
    if (readIndex >= queue.length) {
      return null;
    }

    var value = queue[readIndex];
    readIndex++;
    return value;
  }
}
