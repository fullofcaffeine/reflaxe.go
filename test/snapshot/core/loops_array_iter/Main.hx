class Main {
  static function main() {
    var values:Array<Int> = [2, 4, 6];
    var sum:Int = 0;

    for (value in values) {
      sum = sum + value;
    }

    Sys.println(sum);
  }
}
