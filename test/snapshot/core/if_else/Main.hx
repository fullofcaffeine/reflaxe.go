class Main {
  static function main() {
    var value:Int = 3;
    if (value > 2) {
      Sys.println("gt");
    } else {
      Sys.println("lte");
    }

    if (value == 3 && true) {
      Sys.println("yes");
    } else {
      Sys.println("no");
    }
  }
}
