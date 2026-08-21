class Main {
	static function main() {
		final value = "a😀bé😀";
		expect(value.indexOf("😀"), 1);
		expect(value.indexOf("😀", 2), 4);
		expect(value.indexOf("", 3), 3);
		expect(value.indexOf("", 99), 5);
		expect(value.indexOf("😀", -1), 4);
		expect(value.indexOf("", -1), 0);
		expect(value.indexOf("😀", -99), 1);
		expect(value.indexOf("missing"), -1);
	}

	static function expect(actual:Int, expected:Int):Void {
		if (actual != expected)
			throw 'expected ${expected}, got ${actual}';
	}
}
