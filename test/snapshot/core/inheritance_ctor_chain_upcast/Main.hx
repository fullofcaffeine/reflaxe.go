class Base {
  public var value:Int;

  public function new(value:Int) {
    this.value = value;
  }

  public function read():Int {
    return value;
  }
}

class Child extends Base {
  public function new(value:Int) {
    super(value + 1);
  }
}

class Main {
  static function show(base:Base) {
    Sys.println(base.read());
  }

  static function main() {
    var child = new Child(4);
    var base:Base = child;
    show(child);
    Sys.println(base.read());
  }
}
