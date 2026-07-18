enum ReflectSample {
	Idle;
	Active(value:Int);
}

class PropertyBox {
	public var value(get, set):Int;
	public var plain:Int;

	var stored:Int;

	public function new(value:Int) {
		stored = value;
		plain = value + 10;
	}

	function get_value():Int {
		return stored + 1;
	}

	function set_value(value:Int):Int {
		stored = value * 2;
		return value;
	}
}

class FieldBase {
	public var baseField:Int;

	public function new(baseField:Int) {
		this.baseField = baseField;
	}
}

class FieldLeaf extends FieldBase {
	public var leafField:String;

	public function new(baseField:Int, leafField:String) {
		super(baseField);
		this.leafField = leafField;
	}
}

class Main {
	static function main() {
		var object:Dynamic = {name: "Ada", count: 2};
		var names = Reflect.fields(object);
		var foundCount = false;
		var foundName = false;
		for (name in names) {
			if (name == "count")
				foundCount = true;
			if (name == "name")
				foundName = true;
		}
		Sys.println(names.length);
		Sys.println(foundCount);
		Sys.println(foundName);

		var copied:Dynamic = Reflect.copy(object);
		Reflect.setField(copied, "name", "Bea");
		Sys.println(Reflect.field(object, "name"));
		Sys.println(Reflect.field(copied, "name"));
		Sys.println(Reflect.deleteField(copied, "count"));
		Sys.println(Reflect.hasField(copied, "count"));

		var box = new PropertyBox(4);
		var boxFields = Reflect.fields(box);
		var foundPlainField = false;
		var foundStoredField = false;
		var foundValueField = false;
		for (name in boxFields) {
			if (name == "plain")
				foundPlainField = true;
			if (name == "stored")
				foundStoredField = true;
			if (name == "value")
				foundValueField = true;
		}
		Sys.println(boxFields.length);
		Sys.println(foundPlainField);
		Sys.println(foundStoredField);
		Sys.println(foundValueField);
		Sys.println(Reflect.hasField(box, "value"));
		Sys.println(Reflect.field(box, "value") == null);
		Sys.println(Reflect.hasField(box, "plain"));
		Sys.println(Reflect.field(box, "plain"));
		Reflect.setField(box, "plain", 21);
		Sys.println(Reflect.field(box, "plain"));
		Sys.println(Reflect.getProperty(box, "value"));
		Reflect.setProperty(box, "value", 6);
		Sys.println(Reflect.getProperty(box, "value"));

		var leaf = new FieldLeaf(7, "leaf");
		var asBase:FieldBase = leaf;
		var inheritedFields = Reflect.fields(asBase);
		var foundBaseField = false;
		var foundLeafField = false;
		for (name in inheritedFields) {
			if (name == "baseField")
				foundBaseField = true;
			if (name == "leafField")
				foundLeafField = true;
		}
		Sys.println(inheritedFields.length);
		Sys.println(foundBaseField);
		Sys.println(foundLeafField);
		Sys.println(Reflect.field(asBase, "leafField"));
		Reflect.setField(asBase, "leafField", "changed");
		Sys.println(leaf.leafField);

		var add = function(left:Int, right:Int):Int return left + right;
		Sys.println(Reflect.isFunction(add));
		Sys.println(Reflect.callMethod(null, add, [3, 5]));
		Sys.println(Reflect.compareMethods(add, add));
		Sys.println(Reflect.compareMethods(add, function(left:Int, right:Int):Int return left + right));

		Sys.println(Reflect.isObject(object));
		Sys.println(Reflect.isObject("text"));
		Sys.println(Reflect.isEnumValue(Active(3)));
		Sys.println(Reflect.isEnumValue(object));

		var variadic:Dynamic = Reflect.makeVarArgs(function(arguments:Array<Dynamic>):Dynamic {
			var total = 0;
			for (argument in arguments) {
				var number:Int = argument;
				total += number;
			}
			return total;
		});
		Sys.println(Reflect.callMethod(null, variadic, [1, 2, 3, 4]));
	}
}
