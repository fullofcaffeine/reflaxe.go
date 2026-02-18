class Leaf {
  public function new() {}

  public function ping():Int {
    return pong();
  }

  public function pong():Int {
    return 11;
  }
}

class Main {
  static function main() {
    var leaf = new Leaf();
    var copy = leaf;
    Sys.println(copy.ping());
  }
}
