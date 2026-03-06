class Main {
	static function digestRest(...args:Int):String {
		var rest:haxe.Rest<Int> = args;
		var arrayView = rest.toArray();
		var appended = rest.append(8);
		var prepended = rest.prepend(-1);
		return arrayView.length
			+ ":"
			+ (arrayView.length > 0 ? arrayView[0] : -99)
			+ ":"
			+ (arrayView.length > 0 ? arrayView[arrayView.length - 1] : -99)
			+ "|append="
			+ appended.length
			+ ":"
			+ appended[appended.length - 1]
			+ "|prepend="
			+ prepended.length
			+ ":"
			+ prepended[0]
			+ "|len="
			+ rest.length;
	}

	static function main() {
		Sys.println(digestRest(3, 1, 4));
		Sys.println(digestRest());
	}
}
