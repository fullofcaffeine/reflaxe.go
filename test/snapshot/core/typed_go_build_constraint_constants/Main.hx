import AlwaysFeature;
import NeverFeature;

class Main {
	static function main():Void {
		if (Type.resolveClass("AlwaysFeature") == null)
			throw "unconditional selection lost common type metadata";
		if (Type.resolveClass("NeverFeature") != null)
			throw "impossible selection leaked type metadata";
		Sys.println("ok");
	}
}
