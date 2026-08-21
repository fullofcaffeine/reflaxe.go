class Main {
	static function emit(label:String, value:Dynamic):Void {
		Sys.println(label + "=" + Std.string(value));
	}

	static function main() {
		var parsedArrayDynamic:Dynamic = haxe.Json.parse("[1,true,\"x\",null,{\"a\":1,\"b\":2}]");
		emit("arr.stringify", haxe.Json.stringify(parsedArrayDynamic));

		var parsedObject:Dynamic = haxe.Json.parse("{\"a\":1,\"b\":[2,3],\"c\":{\"d\":\"ok\"},\"z\":null}");
		emit("obj.a", Reflect.field(parsedObject, "a"));
		emit("obj.b_json", haxe.Json.stringify(Reflect.field(parsedObject, "b")));
		emit("obj.c.d", Reflect.field(Reflect.field(parsedObject, "c"), "d"));
		emit("obj.z_is_null", Reflect.field(parsedObject, "z") == null);
		var typedName:String = cast Reflect.field(Reflect.field(parsedObject, "c"), "d");
		emit("obj.c.d_typed", typedName);

		var source:Dynamic = {
			a: 1,
			b: [2, 3],
			c: {
				d: "ok"
			}
		};
		var encoded = haxe.Json.stringify(source);
		var prettySource:Dynamic = {
			items: [1, 2]
		};
		emit("space.omitted", haxe.Json.stringify(prettySource));
		emit("space.null", haxe.Json.stringify(prettySource, null, null));
		emit("space.empty", haxe.Json.stringify(prettySource, null, ""));
		emit("space.two", haxe.Json.stringify(prettySource, null, "  "));
		emit("space.tab", haxe.Json.stringify(prettySource, null, "\t"));

		var decoded:Dynamic = haxe.Json.parse(encoded);
		emit("src.decoded.a", Reflect.field(decoded, "a"));
		emit("src.decoded.b_json", haxe.Json.stringify(Reflect.field(decoded, "b")));
		emit("src.decoded.cd", Reflect.field(Reflect.field(decoded, "c"), "d"));

		var printed = haxe.format.JsonPrinter.print(parsedArrayDynamic);
		emit("printer.matches_stringify", printed == haxe.Json.stringify(parsedArrayDynamic));

		var parsedByParser:Dynamic = haxe.format.JsonParser.parse(encoded);
		emit("parser.a", Reflect.field(parsedByParser, "a"));
		emit("parser.cd", Reflect.field(Reflect.field(parsedByParser, "c"), "d"));
	}
}
