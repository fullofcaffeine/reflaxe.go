class Main {
  static function main() {
    var m = new haxe.ds.IntMap<String>();
    m.set(1, "one");
    m.set(2, "two");
    var one:String = m.get(1);
    Sys.println(one);
    Sys.println(m.exists(3));
  }
}
