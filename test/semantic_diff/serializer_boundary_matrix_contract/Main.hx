interface LabelledValue {
	function label():String;
}

class GenericBox<T> implements LabelledValue {
	public var value:T;
	public var next:GenericBox<T>;

	public function new(value:T) {
		this.value = value;
	}

	public function label():String {
		return "box:" + Std.string(value);
	}
}

class InterfaceEnvelope {
	public var item:LabelledValue;

	public function new(item:LabelledValue) {
		this.item = item;
	}
}

class Main {
	static function throwsOnDecode(wire:String):Bool {
		var caught = false;
		try {
			haxe.Unserializer.run(wire);
		} catch (_:Dynamic) {
			caught = true;
		}
		return caught;
	}

	static function main() {
		var interfaceValue:LabelledValue = new GenericBox<String>("alpha");
		var envelope = new InterfaceEnvelope(interfaceValue);
		var decodedEnvelope:InterfaceEnvelope = cast haxe.Unserializer.run(haxe.Serializer.run(envelope));
		var decodedBox:GenericBox<String> = cast decodedEnvelope.item;
		Sys.println("boundary.interface.value=" + Std.string(decodedBox.value));
		Sys.println("boundary.interface.dispatch=" + decodedEnvelope.item.label());

		var propertyOk = true;
		for (index in 0...64) {
			var value = (index * index) - 37;
			var original = new GenericBox<Int>(value);
			var decoded:GenericBox<Int> = cast haxe.Unserializer.run(haxe.Serializer.run(original));
			if (decoded.value != value || decoded.next != null || decoded.label() != "box:" + value) {
				propertyOk = false;
			}
		}
		Sys.println("boundary.generic.matrix=" + propertyOk);

		var cycle = new GenericBox<Int>(7);
		cycle.next = cycle;
		var serializer = new haxe.Serializer();
		serializer.useCache = true;
		serializer.serialize(cycle);
		var decodedCycle:GenericBox<Int> = cast haxe.Unserializer.run(serializer.toString());
		Sys.println("boundary.generic.cycle=" + (decodedCycle.next == decodedCycle));

		Sys.println("boundary.error.invalidToken=" + throwsOnDecode("!"));
		Sys.println("boundary.error.truncatedString=" + throwsOnDecode("y3:ab"));
		Sys.println("boundary.error.unknownClass=" + throwsOnDecode("cy12:MissingClassg"));
	}
}
