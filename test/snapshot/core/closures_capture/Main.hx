class Main {
  static function main() {
    var factor:Int = 3;
    var mul = function(v:Int):Int {
      return v * factor;
    }

    factor = 4;
    Sys.println(mul(2));
  }
}
