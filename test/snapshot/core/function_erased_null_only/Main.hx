/** Null recovery must compile even when no live callable producer is reachable. */
class Main {
	// Dynamic is the erased storage boundary exercised by this regression.
	static function absent():Dynamic
		return null;

	static function main() {
		var callback:Void->Int = absent();
		Sys.println(callback == null);
		var erased:Dynamic = absent();
		try {
			erased();
			Sys.println("missed null call");
		} catch (error:Dynamic) {
			Sys.println(error);
		}
	}
}
