class DomainError extends haxe.Exception {
	public final code:String;

	public function new(code:String, message:String) {
		super(message);
		this.code = code;
	}
}

class Main {
	static function raiseDomain():Void {
		throw new DomainError("E42", "typed domain failure");
	}

	static function raisePlain():Void {
		throw "plain";
	}

	static function main() {
		try {
			raiseDomain();
		} catch (error:DomainError) {
			Sys.println("typed:" + error.code);
		} catch (error:haxe.Exception) {
			Sys.println("base:" + error.message);
		}

		try {
			raisePlain();
		} catch (error:DomainError) {
			Sys.println("typed:" + error.code);
		} catch (error:haxe.Exception) {
			Sys.println("base:" + error.message);
		}
	}
}
