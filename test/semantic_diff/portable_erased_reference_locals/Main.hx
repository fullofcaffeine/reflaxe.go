interface Named {
	public function name():String;
}

class Base {
	public var value:Int = 3;

	public function new() {}
}

class Child extends Base implements Named {
	public function new() {
		super();
	}

	public function name():String {
		return "child";
	}
}

enum Choice {
	One;
}

class Main {
	static function main():Void {
		var erased:Dynamic = new Child();
		var base:Base = erased;
		Sys.println(base.value);
		var named:Named = erased;
		Sys.println(named.name());
		base.value = 9;
		var child:Child = erased;
		Sys.println(child.value);
		var absent:Dynamic = null;
		var missing:Base = absent;
		Sys.println(missing == null);
		var missingNamed:Named = absent;
		Sys.println(missingNamed == null);
		var choiceValue:Dynamic = One;
		var choice:Choice = choiceValue;
		Sys.println(choice == One);
		var arrayValue:Dynamic = [4, 5];
		var values:Array<Int> = arrayValue;
		Sys.println(values[0]);
	}
}
