class Main {
  static function main() {
    var parsed = haxe.Json.parse("[1,true,\"x\"]");
    Sys.println(haxe.Json.stringify(parsed));
  }
}
