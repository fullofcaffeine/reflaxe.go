class AccessorRoot {
	var rootValue:Int;

	public function new(rootValue:Int) {
		this.rootValue = rootValue;
	}

	public function root():Int {
		return rootValue;
	}

	public function virtualLabel():String {
		return "root";
	}

	public function dispatchLabel():String {
		return virtualLabel();
	}
}

class AccessorBase extends AccessorRoot {
	var inheritedValue:Int;

	@:transient
	public var transientValue:String;

	public function new(rootValue:Int, inheritedValue:Int, transientValue:String) {
		super(rootValue);
		this.inheritedValue = inheritedValue;
		this.transientValue = transientValue;
	}

	public function inherited():Int {
		return inheritedValue;
	}

	public function transientField():String {
		return transientValue;
	}

	override public function virtualLabel():String {
		return "base";
	}
}

class AccessorChild extends AccessorBase {
	var ownValue:String;

	public function new(rootValue:Int, inheritedValue:Int, ownValue:String, transientValue:String) {
		super(rootValue, inheritedValue, transientValue);
		this.ownValue = ownValue;
	}

	public function own():String {
		return ownValue;
	}

	override public function virtualLabel():String {
		return "child:" + ownValue;
	}
}

class HookValue {
	var value:Int;

	public function new(value:Int) {
		this.value = value;
	}

	public function hxSerialize(serializer:haxe.Serializer):Void {
		serializer.serialize(value + 1);
	}

	public function hxUnserialize(unserializer:haxe.Unserializer):Void {
		var encodedValue:Int = unserializer.unserialize();
		value = encodedValue - 1;
	}

	public function current():Int {
		return value;
	}
}

class NumericWireValue {
	public var amount:Float;

	public function new(amount:Float) {
		this.amount = amount;
	}
}

class Main {
	static function main() {
		var original = new AccessorChild(3, 7, "leaf", "not-on-wire");
		var encoded = haxe.Serializer.run(original);
		var decoded:AccessorChild = cast haxe.Unserializer.run(encoded);
		Sys.println("access.encoded=" + encoded);
		Sys.println("access.root=" + decoded.root());
		Sys.println("access.inherited=" + decoded.inherited());
		Sys.println("access.own=" + decoded.own());
		Sys.println("access.transient=" + Std.string(decoded.transientField()));
		Sys.println("access.virtual=" + decoded.dispatchLabel());

		var hook:HookValue = cast haxe.Unserializer.run(haxe.Serializer.run(new HookValue(11)));
		Sys.println("access.hook=" + hook.current());

		// Integral Float values may be encoded with the compact `i` token by
		// another Haxe target. Assignment must restore the declared Float field.
		var numeric:NumericWireValue = cast haxe.Unserializer.run("cy16:NumericWireValuey6:amounti3g");
		Sys.println("access.numeric=" + numeric.amount);
	}
}
