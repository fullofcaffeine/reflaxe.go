package go;

/**
	What: The source-level implementation behind the typed `go.Slice<T>` facade.

	Why: `go.Slice<T>` models a native Go slice, while portable Haxe `Array<T>`
	has shared object identity and sparse-growth behavior. Storing an `Array<T>`
	here would silently mix those two contracts whenever native specialization is
	not available.

	How: Keep the backing field as an explicit `NativeSlice<T>`. The compiler can
	lower its indexing and append operations to ordinary Go slice operations, and
	`toArray()` performs the documented shallow-copy conversion back to portable
	Array storage. This Go-native facade intentionally has no interpreter fallback;
	target-only behavior is verified by Go runtime contracts.
**/
class Slice<T> {
	public var data(default, null):NativeSlice<T>;

	public function new() {
		data = NativeSlice.fromArray([]);
	}

	public var length(get, never):Int;

	function get_length():Int {
		return data.length;
	}

	public function push(value:T):Void {
		data = NativeSlice.append(data, value);
	}

	public function get(index:Int):T {
		return data[index];
	}

	public function set(index:Int, value:T):Void {
		data[index] = value;
	}

	public function toArray():Array<T> {
		return data.toArray();
	}
}
