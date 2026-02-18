class Main {
  static function main() {
    var once = 0;
    do {
      once++;
    } while (false);

    var i = 0;
    var hit = 0;
    do {
      i++;
      if (i < 3) {
        continue;
      }
      hit = i;
    } while (i < 3);

    Sys.println(Std.string(once));
    Sys.println(Std.string(hit));
  }
}
