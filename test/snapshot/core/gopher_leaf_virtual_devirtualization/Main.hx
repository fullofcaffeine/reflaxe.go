class Leaf {
  public function new() {}

  public function ping():Int {
    return pong();
  }

  public function pong():Int {
    return 7;
  }
}

class Main {
  static function main() {
    var leaf = new Leaf();
    Sys.println(leaf.ping());
  }
}
