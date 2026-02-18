enum Color {
  Red;
  RGB(r:Int, g:Int, b:Int);
}

class Main {
  static function isSome(value:Color):Int {
    if (value == null) {
      return 0;
    }
    return 1;
  }

  static function main() {
    var red = Color.Red;
    var rgb = Color.RGB(1, 2, 3);
    Sys.println(isSome(red));
    Sys.println(isSome(rgb));
  }
}
