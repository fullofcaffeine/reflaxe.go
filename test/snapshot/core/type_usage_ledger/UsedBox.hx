class UsedBox<T> {
	public var value:T;

	public function new(value:T) {
		this.value = value;
	}
}

class UsedSibling {
	public final delta:Int;

	public function new(delta:Int) {
		this.delta = delta;
	}
}
