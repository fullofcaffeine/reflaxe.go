class Main {
	static function main() {
		var payload:Dynamic = {
			name: "delta",
			values: [1, 2, 3, 4],
			nested: {flag: true, note: "ready"}
		};

		var encoded = haxe.Serializer.run(payload);
		var decoded:Dynamic = haxe.Unserializer.run(encoded);
		var nested:Dynamic = Reflect.field(decoded, "nested");
		Sys.println("encoded.nonEmpty=" + (encoded != null && encoded != ""));
		Sys.println("decoded.name=" + (Reflect.field(decoded, "name") == "delta"));
		Sys.println("decoded.values.present=" + (Reflect.field(decoded, "values") != null));
		Sys.println("decoded.nested.flag=" + (Reflect.field(nested, "flag") == true));
		Sys.println("decoded.nested.note=" + (Reflect.field(nested, "note") == "ready"));

		var serializer = new haxe.Serializer();
		serializer.serialize({id: 7, title: "card"});
		var replay = haxe.Unserializer.run(serializer.toString());
		Sys.println("serializer.replay.id=" + Reflect.field(replay, "id"));
		Sys.println("serializer.replay.title=" + Reflect.field(replay, "title"));
	}
}
