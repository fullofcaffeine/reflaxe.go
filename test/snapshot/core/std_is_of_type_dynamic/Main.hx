class Base {
  public function new() {}
}

class Child extends Base {
  public function new() {
    super();
  }
}

class Main {
  static function main() {
    var d:Dynamic = new Child();
    Sys.println(Std.isOfType(d, Base));
    Sys.println(Std.isOfType(d, Child));

    d = new Base();
    Sys.println(Std.isOfType(d, Child));

    d = [1, 2];
    Sys.println(Std.isOfType(d, Array));

    d = 1;
    Sys.println(Std.isOfType(d, Array));

    d = null;
    Sys.println(Std.isOfType(d, Dynamic));
  }
}
