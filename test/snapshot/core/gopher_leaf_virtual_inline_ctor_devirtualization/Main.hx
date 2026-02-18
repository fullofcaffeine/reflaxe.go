class Leaf {
  public function new() {}

  public function ping():Int {
    return pong();
  }

  public function pong():Int {
    return 13;
  }
}

class Main {
  static function main() {
    Sys.println((new Leaf()).ping());
  }
}
