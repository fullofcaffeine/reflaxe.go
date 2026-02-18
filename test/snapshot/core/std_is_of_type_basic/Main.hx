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
    var child:Base = new Child();
    var base:Base = new Base();

    Sys.println(Std.isOfType(child, Child));
    Sys.println(Std.isOfType(child, Base));
    Sys.println(Std.isOfType(base, Child));
    Sys.println(Std.isOfType(null, Child));

    Sys.println(Std.isOfType(1, Int));
    Sys.println(Std.isOfType(1, Float));
    Sys.println(Std.isOfType(1.5, Int));
    Sys.println(Std.isOfType("x", String));
    Sys.println(Std.isOfType(true, Bool));
    Sys.println(Std.isOfType(null, Dynamic));
  }
}
