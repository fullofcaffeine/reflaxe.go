class Main {
	static function makeBase():Base {
		return new Child();
	}

	static function main() {
		var base:Base = makeBase();
		Sys.println(base.callWho());
	}
}
