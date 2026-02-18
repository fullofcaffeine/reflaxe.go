class Main {
  static function mayFail(flag:Bool):Int {
    if (flag) {
      throw "boom";
    }
    return 7;
  }

  static function main() {
    try {
      Sys.println(mayFail(false));
      Sys.println(mayFail(true));
    } catch (e:Dynamic) {
      Sys.println(e);
    }
    Sys.println(9);
  }
}
