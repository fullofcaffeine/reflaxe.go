class Base {
  public function new() {}

  public function ping():Int {
    return 1;
  }
}

class Child extends Base {
  override public function ping():Int {
    return 2;
  }
}

class Main {
  static function main() {
    var child = new Child();
    Sys.println(child.ping());
  }
}
