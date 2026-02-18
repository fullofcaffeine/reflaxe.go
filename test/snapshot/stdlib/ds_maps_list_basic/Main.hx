enum EKey {
  A;
  B(v:Int);
}

class Box {
  public var id:Int;

  public function new(id:Int) {
    this.id = id;
  }
}

class Main {
  static function main() {
    var sm = new haxe.ds.StringMap<Int>();
    sm.set("a", 1);
    var av:Int = sm.get("a");
    Sys.println(av);

    var om = new haxe.ds.ObjectMap<Box, String>();
    var box = new Box(7);
    om.set(box, "box");
    var ov:String = om.get(box);
    Sys.println(ov);

    var em = new haxe.ds.EnumValueMap<EKey, String>();
    em.set(EKey.A, "enum");
    var ev:String = em.get(EKey.A);
    Sys.println(ev);

    var list = new haxe.ds.List<Int>();
    list.add(4);
    list.add(5);
    Sys.println(list.length);
    Sys.println(list.first());
    Sys.println(list.last());
    Sys.println(list.pop());
    Sys.println(list.length);
  }
}
