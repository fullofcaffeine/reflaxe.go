typedef View = {var n:Int; var ?label:String;}

class C {
	public var n:Int = 1;

	public function new() {}
}

class Present extends C {
	public var label:String = null;

	public function new() {
		super();
	}
}

class Main {
	static function main() {
		var erased:Dynamic = new C();
		var view:View = erased;
		Sys.println(view.n);
		Sys.println(view.label == null);
		Sys.println(Reflect.hasField(view, "label"));
		var present:View = new Present();
		Sys.println(present.label == null);
		Sys.println(Reflect.hasField(present, "label"));
	}
}
