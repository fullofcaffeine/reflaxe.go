class Main {
  static function main() {
    var a = -1;
    var b = (a >>>= 1);
    Sys.println(Std.string(a));
    Sys.println(Std.string(b));
  }
}
