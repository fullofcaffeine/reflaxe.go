class MathBox {
  public static var factor:Int = 3;

  public static function mul(value:Int):Int {
    return value * factor;
  }
}

class Main {
  static function main() {
    Sys.println(MathBox.mul(4));
    MathBox.factor = 5;
    Sys.println(MathBox.mul(4));
  }
}
