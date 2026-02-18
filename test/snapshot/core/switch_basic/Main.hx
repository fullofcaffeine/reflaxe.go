class Main {
  static function pick(v:Int):Int {
    return switch (v) {
      case 0, 1: 10;
      case 2: 20;
      default: 30;
    }
  }

  static function main() {
    var v = 1;
    switch (v) {
      case 0:
        Sys.println(0);
      case 1:
        Sys.println(1);
      default:
        Sys.println(9);
    }
    Sys.println(pick(0));
    Sys.println(pick(2));
    Sys.println(pick(7));
  }
}
