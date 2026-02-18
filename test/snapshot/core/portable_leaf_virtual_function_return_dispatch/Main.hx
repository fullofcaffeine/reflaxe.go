class Leaf {
  public function new() {}

  public function ping():Int {
    return pong();
  }

  public function pong():Int {
    return 17;
  }
}

class Factory {
  public static function makeLeaf():Leaf {
    return new Leaf();
  }
}

class Main {
  static function main() {
    Sys.println(Factory.makeLeaf().ping());
  }
}
