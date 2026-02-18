class Main {
  static function main() {
    var i = 1;
    var now = (i += 2);
    Sys.println(Std.string(now));
    Sys.println(Std.string(i));

    i *= 4;
    Sys.println(Std.string(i));
  }
}
