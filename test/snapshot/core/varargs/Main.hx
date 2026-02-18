class Main {
  static function sum(...values:Int):Int {
    var total:Int = 0;
    var i:Int = 0;
    while (i < values.length) {
      total = total + values[i];
      i = i + 1;
    }
    return total;
  }

  static function main() {
    Sys.println(sum(1, 2, 3));
    Sys.println(sum(4));
  }
}
