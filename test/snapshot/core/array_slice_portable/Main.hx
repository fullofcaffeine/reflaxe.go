class Main {
	static function printValues(values:Array<Int>):Void {
		Sys.println(values.join(","));
	}

	static function main() {
		final values = [0, 1, 2, 3];
		printValues(values.slice(1));
		printValues(values.slice(1, 3));
		printValues(values.slice(-2));
		printValues(values.slice(0, -1));
		Sys.println(values.slice(7).length);
		printValues(values.slice(1, null));
		final nullableEnd:Null<Int> = null;
		printValues(values.slice(1, nullableEnd));
		final concreteEnd:Null<Int> = 3;
		printValues(values.slice(1, concreteEnd));

		final copy = values.slice(0, 2);
		copy[0] = 9;
		Sys.println(values[0]);
		printValues(NativeSlicer.middle(values));
	}
}
