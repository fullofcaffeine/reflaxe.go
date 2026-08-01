interface DefaultedValue {
	public function get(value:Int = 5):Int;
	public function label(value:String = "interface"):String;
}

class DefaultedValueBase {
	public function new() {}

	public function get(value:Int = 7):Int {
		return value;
	}
}

class DefaultedValueImpl extends DefaultedValueBase implements DefaultedValue {
	public function new() {
		super();
	}

	public function label(value:String = "implementation"):String {
		return value;
	}
}

class Main {
	static function main() {
		var value:DefaultedValue = new DefaultedValueImpl();
		Sys.println(value.get());
		Sys.println(value.get(9));
		Sys.println(value.label());
		Sys.println(value.label("explicit"));
	}
}
