class Main {
	static function main() {
		var names = new Array<String>();
		Sys.println(names.length);
		names.push("go");
		names.push("haxe");
		Sys.println(names.length);
		Sys.println(names[0]);
		Sys.println(names[1]);

		var nums = new Array<Int>();
		nums.push(3);
		nums.push(5);
		var sum = 0;
		for (n in nums) {
			sum += n;
		}
		Sys.println(sum);
	}
}
