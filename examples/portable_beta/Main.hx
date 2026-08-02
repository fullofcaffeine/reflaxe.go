import haxe.Json;
import haxe.ds.StringMap;
import haxe.io.Bytes;
import haxe.io.Path;
import sys.FileSystem;
import sys.io.File;

private enum DeliveryState {
	Queued;
	Delivered;
}

private class Delivery {
	public final name:String;
	public final state:DeliveryState;

	public function new(name:String, state:DeliveryState) {
		this.name = name;
		this.state = state;
	}
}

/**
	What: Runs a deliberately narrow set of portable-beta operations.
	Why: A green broad example must not silently widen the public beta claim.
	How: Each check uses an operation named by the compatibility manifest and
	throws when the observed result differs from its manually authored oracle.
**/
class Main {
	static inline final SCRATCH_FILE = ".portable_beta_contract.txt";

	static function require(condition:Bool, message:String):Void {
		if (!condition) {
			throw message;
		}
	}

	static function checkLanguageAndCollections():Void {
		var deliveries = new Array<Delivery>();
		deliveries.push(new Delivery("alpha", Delivered));
		deliveries.push(new Delivery("beta", Queued));
		deliveries.push(new Delivery("gamma", Delivered));

		var delivered = 0;
		for (index in 0...deliveries.length) {
			switch (deliveries[index].state) {
				case Delivered:
					delivered++;
				case Queued:
			}
		}
		require(delivered == 2, "enum and numeric-for contract");

		var removed = deliveries.pop();
		require(removed != null, "array pop contract");
		if (removed != null) {
			require(removed.name == "gamma", "array order contract");
		}
		require(deliveries.length == 2, "array length contract");

		var decorate = function(value:String):String {
			return "[" + value + "]";
		};
		require(decorate("portable") == "[portable]", "closure contract");
	}

	static function checkTextAndData():Void {
		var cleaned = StringTools.trim("  alpha-go  ");
		require(cleaned.length == 8, "string length contract");
		require(cleaned.charAt(5) == "-", "string charAt contract");
		require(cleaned.substring(0, 5) == "alpha", "string substring contract");
		require(StringTools.startsWith(cleaned, "alpha"), "startsWith contract");
		require(StringTools.endsWith(cleaned, "go"), "endsWith contract");
		require(StringTools.contains(cleaned, "ha-g"), "contains contract");
		require(StringTools.replace(cleaned, "-", " ") == "alpha go", "replace contract");

		var bytes = Bytes.ofString(cleaned);
		require(bytes.toString() == cleaned, "bytes round-trip contract");

		var names = new StringMap<String>();
		names.set("primary", "alpha");
		require(names.exists("primary"), "StringMap exists contract");
		require(names.get("primary") == "alpha", "StringMap get contract");
		require(names.remove("primary"), "StringMap remove contract");
		require(!names.exists("primary"), "StringMap removal contract");

		var parsed:Dynamic = Json.parse('{"name":"alpha","ready":true}');
		require(Reflect.hasField(parsed, "name"), "JSON field contract");
		var parsedName:String = cast Reflect.field(parsed, "name");
		require(parsedName == "alpha", "JSON value contract");
		var encoded = Json.stringify({name: "alpha"});
		require(StringTools.contains(encoded, '"name":"alpha"'), "JSON stringify contract");

		var path = new Path(Path.join(["reports", "result.json"]));
		require(path.dir == "reports", "Path directory contract");
		require(path.file == "result", "Path file contract");
		require(path.ext == "json", "Path extension contract");
	}

	static function checkFileLifecycle():Void {
		if (FileSystem.exists(SCRATCH_FILE)) {
			FileSystem.deleteFile(SCRATCH_FILE);
		}
		try {
			File.saveContent(SCRATCH_FILE, "portable-beta-ok");
			require(FileSystem.exists(SCRATCH_FILE), "file existence contract");
			require(File.getContent(SCRATCH_FILE) == "portable-beta-ok", "file content contract");
			FileSystem.deleteFile(SCRATCH_FILE);
		} catch (error:haxe.Exception) {
			if (FileSystem.exists(SCRATCH_FILE)) {
				FileSystem.deleteFile(SCRATCH_FILE);
			}
			throw error;
		}
		require(!FileSystem.exists(SCRATCH_FILE), "file cleanup contract");
	}

	static function checkExceptionContract():Void {
		var observed = "";
		try {
			throw "portable-beta-error";
		} catch (error:haxe.Exception) {
			observed = error.message;
		}
		require(observed == "portable-beta-error", "typed exception message contract");
	}

	static function main():Void {
		checkLanguageAndCollections();
		checkTextAndData();
		checkFileLifecycle();
		checkExceptionContract();
	}
}
