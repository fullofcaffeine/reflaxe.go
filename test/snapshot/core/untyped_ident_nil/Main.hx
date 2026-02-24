class Main {
	static function main() {
		var dynamicNil:Dynamic = untyped nil;
		if (dynamicNil == null) {
			Sys.println("ident:nil");
		} else {
			Sys.println("ident:non_nil");
		}
	}
}
