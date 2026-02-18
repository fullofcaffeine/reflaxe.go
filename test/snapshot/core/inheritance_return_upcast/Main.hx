class Base {
  public function new() {}

  public function tag():Int {
    return 1;
  }
}

class Child extends Base {
  public function new() {
    super();
  }

  override public function tag():Int {
    return 2;
  }
}

class Factory {
  public static function makeBase(flag:Bool):Base {
    if (flag) {
      return new Child();
    }
    return new Base();
  }
}

class Main {
  static function main() {
    var asBase = Factory.makeBase(true);
    Sys.println(asBase.tag());
  }
}
