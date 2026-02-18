class Main {
  static function id(value:Int):Int {
    return value;
  }

  static function main() {
    var i = 1;
    var before = id(i++);
    Sys.println(Std.string(before));
    Sys.println(Std.string(i));
  }
}
