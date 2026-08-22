class Main {
	static function compare(left:String, right:String):Int {
		return left < right ? -1 : left == right ? 0 : 1;
	}

	static function main() {
		final values = ["beta", "alpha", "gamma"];
		values.sort(compare);
		var score = 0;
		if (values[0] == "alpha")
			score += 1;
		if (values[1] == "beta")
			score += 2;
		if (values[2] == "gamma")
			score += 4;
		if ("alpha" < "beta")
			score += 8;
		if ("beta" <= "beta")
			score += 16;
		if ("gamma" > "beta")
			score += 32;
		if ("gamma" >= "gamma")
			score += 64;
		if (score != 127)
			throw "unexpected sort or string ordering result";
	}
}
