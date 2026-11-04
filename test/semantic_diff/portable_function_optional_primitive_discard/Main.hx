class Main {
	static function make():(?Int) -> Int {
		return function(value:Int = 9):Int {
			Sys.println(value);
			return value;
		};
	}

	static function invoke(callback:(?Int) -> Void) {
		callback();
		callback(0);
		callback(4);
	}

	static function main() {
		invoke(make());
	}
}
