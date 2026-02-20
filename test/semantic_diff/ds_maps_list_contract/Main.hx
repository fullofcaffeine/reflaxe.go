enum EKey {
	A;
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
		sm.set("b", 2);
		Sys.println("sm.a=" + Std.string(sm.get("a")));
		Sys.println("sm.exists.b0=" + sm.exists("b"));
		Sys.println("sm.remove.b=" + sm.remove("b"));
		Sys.println("sm.exists.b1=" + sm.exists("b"));

		var im = new haxe.ds.IntMap<String>();
		im.set(7, "seven");
		Sys.println("im.7=" + Std.string(im.get(7)));
		Sys.println("im.exists.7a=" + im.exists(7));
		Sys.println("im.remove.7=" + im.remove(7));
		Sys.println("im.exists.7b=" + im.exists(7));
		var imMissing = im.get(7);
		Sys.println("im.missing=" + Std.string(imMissing));
		Sys.println("im.missing.null=" + (imMissing == null));

		var om = new haxe.ds.ObjectMap<Box, String>();
		var b1 = new Box(1);
		var b2 = new Box(1);
		om.set(b1, "one");
		Sys.println("om.b1=" + Std.string(om.get(b1)));
		Sys.println("om.exists.b2=" + om.exists(b2));
		Sys.println("om.remove.b1=" + om.remove(b1));
		Sys.println("om.exists.b1=" + om.exists(b1));
		var omMissing = om.get(b2);
		Sys.println("om.missing=" + Std.string(omMissing));
		Sys.println("om.missing.null=" + (omMissing == null));

		var em = new haxe.ds.EnumValueMap<EKey, String>();
		em.set(EKey.A, "enumA");
		Sys.println("em.A=" + Std.string(em.get(EKey.A)));
		Sys.println("em.remove.A=" + em.remove(EKey.A));
		Sys.println("em.exists.A=" + em.exists(EKey.A));
		var emMissing = em.get(EKey.A);
		Sys.println("em.missing=" + Std.string(emMissing));
		Sys.println("em.missing.null=" + (emMissing == null));

		var list = new haxe.ds.List<Int>();
		list.add(4);
		list.push(5);
		list.push(6);
		Sys.println("list.len0=" + list.length);
		Sys.println("list.first=" + Std.string(list.first()));
		Sys.println("list.last=" + Std.string(list.last()));
		Sys.println("list.pop0=" + Std.string(list.pop()));
		Sys.println("list.pop1=" + Std.string(list.pop()));
		Sys.println("list.len1=" + list.length);

		var stringList = new haxe.ds.List<String>();
		var slFirst = stringList.first();
		var slLast = stringList.last();
		var slPop = stringList.pop();
		Sys.println("stringList.first=" + Std.string(slFirst));
		Sys.println("stringList.last=" + Std.string(slLast));
		Sys.println("stringList.pop=" + Std.string(slPop));
		Sys.println("stringList.first.null=" + (slFirst == null));
		Sys.println("stringList.last.null=" + (slLast == null));
		Sys.println("stringList.pop.null=" + (slPop == null));
	}
}
