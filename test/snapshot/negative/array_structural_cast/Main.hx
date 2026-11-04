class Main {
	static function main():Void {
		var source = [1, 2];
		var erased:Iterable<Int> = cast source;
		Sys.println(erased.iterator().next());
	}
}
