class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		var dateValue = Date.fromString("2020-01-02 03:04:05");
		safe("ser.date.value", function() return haxe.Serializer.run(dateValue));
		safe("ser.date.token", function() return haxe.Serializer.run(haxe.Unserializer.run("v1.577955845e+12")));

		var bytesValue = haxe.io.Bytes.alloc(3);
		bytesValue.set(0, 97);
		bytesValue.set(1, 0);
		bytesValue.set(2, 98);
		safe("ser.bytes.value", function() return haxe.Serializer.run(bytesValue));
		safe("ser.bytes.token", function() return haxe.Serializer.run(haxe.Unserializer.run("s4:YQBi")));
		safe("ser.bytes.tokenUtf8", function() return haxe.Serializer.run(haxe.Unserializer.run("s6:w6HOsg")));
	}
}
