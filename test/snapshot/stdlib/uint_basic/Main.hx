class Main {
	static function main():Void {
		var base:UInt = cast 3;
		var scaled:UInt = base * 7;
		var shifted:UInt = scaled >> 1;
		Sys.println(Std.string(base));
		Sys.println(Std.string(scaled));
		Sys.println(Std.string(shifted));
	}
}
