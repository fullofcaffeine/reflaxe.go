class Base {
  public function new() {}

  public function who():Int {
    return 1;
  }

  public function callWho():Int {
    return who();
  }
}

class Child extends Base {
  public function new() {
    super();
  }

  override public function who():Int {
    return 2;
  }
}

class Main {
  static function main() {
    var child = new Child();
    var base:Base = child;
    Sys.println(base.callWho());
  }
}
