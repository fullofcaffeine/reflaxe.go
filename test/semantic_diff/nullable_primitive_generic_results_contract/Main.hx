class Main {
	static function main():Void {
		var intMap = new Map<String, Int>();
		intMap.set("alpha", 2);
		var intPresent:Null<Int> = intMap.get("alpha");
		var intMissing:Null<Int> = intMap.get("missing");
		Sys.println("int.present.null=" + (intPresent == null));
		Sys.println("int.present.plus1=" + (intPresent == null ? -1 : intPresent + 1));
		Sys.println("int.missing.null=" + (intMissing == null));
		Sys.println("int.missing.plus1=" + (intMissing == null ? -1 : intMissing + 1));
		var intCast:Int = cast intMap.get("alpha");
		Sys.println("int.cast=" + intCast);

		var floatMap = new Map<String, Float>();
		floatMap.set("pi", 3.5);
		var floatPresent:Null<Float> = floatMap.get("pi");
		var floatMissing:Null<Float> = floatMap.get("missing");
		Sys.println("float.present.null=" + (floatPresent == null));
		Sys.println("float.present.string=" + Std.string(floatPresent));
		Sys.println("float.missing.null=" + (floatMissing == null));
		Sys.println("float.missing.string=" + Std.string(floatMissing));
		var floatCast:Float = cast floatMap.get("pi");
		Sys.println("float.cast=" + Std.string(floatCast));

		var boolMap = new Map<String, Bool>();
		boolMap.set("yes", true);
		var boolPresent:Null<Bool> = boolMap.get("yes");
		var boolMissing:Null<Bool> = boolMap.get("missing");
		Sys.println("bool.present.null=" + (boolPresent == null));
		Sys.println("bool.present.string=" + Std.string(boolPresent));
		Sys.println("bool.missing.null=" + (boolMissing == null));
		Sys.println("bool.missing.string=" + Std.string(boolMissing));
		var boolCast:Bool = cast boolMap.get("yes");
		Sys.println("bool.cast=" + Std.string(boolCast));

		var list = new List<Int>();
		var listFirstMissing:Null<Int> = list.first();
		var listLastMissing:Null<Int> = list.last();
		var listPopMissing:Null<Int> = list.pop();
		Sys.println("list.first.missing.null=" + (listFirstMissing == null));
		Sys.println("list.last.missing.null=" + (listLastMissing == null));
		Sys.println("list.pop.missing.null=" + (listPopMissing == null));

		list.add(5);
		list.add(8);
		var listFirstPresent:Null<Int> = list.first();
		var listLastPresent:Null<Int> = list.last();
		var listPopPresent:Null<Int> = list.pop();
		Sys.println("list.first.present.plus1=" + (listFirstPresent == null ? -1 : listFirstPresent + 1));
		Sys.println("list.last.present.plus1=" + (listLastPresent == null ? -1 : listLastPresent + 1));
		Sys.println("list.pop.present.plus1=" + (listPopPresent == null ? -1 : listPopPresent + 1));
	}
}
