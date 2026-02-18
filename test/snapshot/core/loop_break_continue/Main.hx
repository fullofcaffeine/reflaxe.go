class Main {
  static function main() {
    var i = 0;
    var sum = 0;
    while (i < 10) {
      i++;
      if (i % 2 == 0) {
        continue;
      }
      if (i > 7) {
        break;
      }
      sum += i;
    }

    Sys.println(Std.string(sum));
    Sys.println(Std.string(i));
  }
}
