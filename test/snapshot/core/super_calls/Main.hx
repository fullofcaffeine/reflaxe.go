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

  public function callSuperWho():Int {
    return super.who();
  }

  public function callSuperCallWho():Int {
    return super.callWho();
  }
}

class Main {
  static function main() {
    var child = new Child();
    var base:Base = child;

    Sys.println(child.callWho());
    Sys.println(child.callSuperWho());
    Sys.println(child.callSuperCallWho());
    Sys.println(base.callWho());
  }
}
