class Counter {
  public var value:Int;

  public function new(start:Int) {
    this.value = start;
  }

  public function inc(step:Int):Int {
    value = value + step;
    return value;
  }
}

class Main {
  static function main() {
    var counter = new Counter(5);
    Sys.println(counter.value);
    Sys.println(counter.inc(2));
    Sys.println(counter.value);
  }
}
