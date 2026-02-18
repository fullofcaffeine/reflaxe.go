class Main {
  static function add(a:Int = 1, b:Int = 2):Int {
    return a + b;
  }

  static function main() {
    function local(v:Int = 10):Int {
      return v + 1;
    }

    Sys.println(add());
    Sys.println(add(5));
    Sys.println(add(5, 6));
    Sys.println(local());
    Sys.println(local(20));
  }
}
