class Parent extends haxe.Exception {
	public function new() {
		super("parent");
	}

	override public function toString():String {
		return "parent-string";
	}
}

class Child extends Parent {
	public function new() {
		super();
	}

	override public function toString():String {
		return "child-string";
	}
}

class Main {
	static function main():Void {
		var child = new Child();
		var base:haxe.Exception = child;
		Sys.println("base-string=" + base.toString());
		try {
			throw base;
		} catch (parent:Parent) {
			Sys.println("parent-catch=" + Std.string(parent == child));
		} catch (other:Dynamic) {
			Sys.println("parent-catch=false");
		}
		var result = try {
			throw child;
			"missed";
		} catch (caught:haxe.Exception) {
			caught.toString();
		};
		Sys.println("expression=" + result);
		try {
			try {
				throw child;
			} catch (caught:haxe.Exception) {
				throw caught;
			}
		} catch (caught:Child) {
			Sys.println("rethrow=" + Std.string(caught == child));
		} catch (other:Dynamic) {
			Sys.println("rethrow=false");
		}
		Sys.println("base-not-value=" + Std.string(!Std.isOfType(base, haxe.ValueException)));
		var native = {marker: 42};
		var attached = new haxe.Exception("attached", base, native);
		Sys.println("native=" + Std.string(attached.native == native));
		Sys.println("previous=" + Std.string(attached.previous == base));
		try {
			throw 17;
		} catch (caught:haxe.Exception) {
			Sys.println("primitive=" + caught.message);
			Sys.println("primitive-value=" + Std.string(Std.isOfType(caught, haxe.ValueException)));
		}
	}
}
