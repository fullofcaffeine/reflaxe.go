class Main {
  static function main() {
    var values:Array<Int> = [];
    values.push(4);
    values.push(9);
    values.pop();

    Sys.println(values.length);
    Sys.println(values[0]);
  }
}
