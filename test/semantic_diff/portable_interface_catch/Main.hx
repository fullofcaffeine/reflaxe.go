interface Tagged {
	public function tag():Int;
}

interface MoreTagged extends Tagged {}

class Problem implements MoreTagged {
	public function new() {}

	public function tag():Int {
		return 7;
	}
}

class SameShape {
	public function new() {}

	public function tag():Int {
		return 9;
	}
}

class Main {
	static function main() {
		try {
			throw new Problem();
		} catch (problem:Tagged) {
			Sys.println("tag=" + problem.tag());
		} catch (_:Dynamic) {
			Sys.println("miss");
		}
		try {
			throw new SameShape();
		} catch (_:Tagged) {
			Sys.println("incorrect-structural-match");
		} catch (_:Dynamic) {
			Sys.println("nominal-miss");
		}
	}
}
