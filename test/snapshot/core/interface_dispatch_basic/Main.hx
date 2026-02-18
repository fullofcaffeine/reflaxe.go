interface Named {
  public function name():String;
}

class Person implements Named {
  public function new() {}

  public function name():String {
    return "person";
  }
}

class Main {
  static function printNamed(value:Named):Void {
    Sys.println(value.name());
  }

  static function main() {
    var named:Named = new Person();
    printNamed(named);
    printNamed(new Person());
  }
}
