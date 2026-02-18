class Main {
  static function twice(value:Int):Int {
    return value * 2;
  }

  static function main() {
    function add(a:Int, b:Int):Int {
      return a + b;
    }

    var mul = function(v:Int):Int {
      return v * 3;
    }

    Sys.println(twice(5));
    Sys.println(add(2, 7));
    Sys.println(mul(4));
  }
}
