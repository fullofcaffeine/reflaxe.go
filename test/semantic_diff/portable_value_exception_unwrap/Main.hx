class DifferentValue extends haxe.ValueException {
	public function new() {
		super("stored");
	}

	override function unwrap():Any {
		return "unwrapped";
	}
}

class Main {
	static function main() {
		try {
			throw new DifferentValue();
		} catch (value:Dynamic) {
			Sys.println(Std.string(value));
		}
	}
}
