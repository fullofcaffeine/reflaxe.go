class Main {
  static function sum(...values:Int):Int {
    var total = 0;
    for (value in values) {
      total = total + value;
    }
    return total;
  }

  static function main() {
    Sys.println(sum(1, 2, 3));
    Sys.println(sum(4));
  }
}
