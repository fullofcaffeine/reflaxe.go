class Main {
  static function fail():Void {
    throw "boom";
  }

  static function main() {
    try {
      fail();
    } catch (e:haxe.Exception) {
      Sys.println(e.message);
    }
  }
}
