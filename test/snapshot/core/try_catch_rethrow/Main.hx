class Main {
  static function main() {
    try {
      try {
        throw true;
      } catch (i:Int) {
        Sys.println(i);
      }
    } catch (e:Dynamic) {
      Sys.println("outer");
    }
  }
}
