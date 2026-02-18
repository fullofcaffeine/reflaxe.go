class Keywords {
  public static function map():Int {
    return 30;
  }

  public static function range():Int {
    return 40;
  }
}

class Main {
  static function main() {
    Sys.println(a_b.Util.value());
    Sys.println(a.b.Util.value());
    Sys.println(Keywords.map());
    Sys.println(Keywords.range());
  }
}
