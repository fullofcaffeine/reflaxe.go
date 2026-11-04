class Cell<T> {
	public var value:T;

	public function new(value:T) {
		this.value = value;
	}
}

class Main {
	static function main() {
		var source = new Cell<String>(null);
		var view:{value:String} = source;
		var observed:Dynamic = view.value;
		Sys.println(observed == null);
	}
}
