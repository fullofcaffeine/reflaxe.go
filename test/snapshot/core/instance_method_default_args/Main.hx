class Greeter {
	public function new() {}

	public function greet(name:String = "world", punct:String = "!"):String {
		return "hello " + name + punct;
	}

	public function wrap(prefix:String = "[", suffix:String = "]"):String {
		return prefix + "go" + suffix;
	}
}

class Main {
	static function main() {
		var greeter = new Greeter();
		Sys.println(greeter.greet());
		Sys.println(greeter.greet("Go"));
		Sys.println(greeter.greet("Go", "?"));
		Sys.println(greeter.wrap());
		Sys.println(greeter.wrap("<"));
		Sys.println(greeter.wrap("<", ">"));
	}
}
