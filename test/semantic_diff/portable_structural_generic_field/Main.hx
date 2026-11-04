class Holder<T> {
	public var value:Iterator<T>;

	public function new(value:Iterator<T>) {
		this.value = value;
	}

	public function read():Iterator<T>
		return value;

	public function erase():Dynamic
		return value;
}

class Main {
	static function main() {
		var source = {hasNext: () -> true, next: () -> 3};
		var holder = new Holder<Int>(source);
		Sys.println(holder.value.next());
		source.next = () -> 9;
		Sys.println(holder.value.next());
	}
}
