class Main {
  static function id(value:Int):Int {
    return value;
  }

  static function main() {
    var i = 1;
    var now = id(++i);
    Sys.println(Std.string(now));
    Sys.println(Std.string(i));
  }
}
