package go;

/**
	What: One native Go `map[string]any` whose values are native Go values.

	Why: Go APIs such as `text/template` and `encoding/json` read data through
	reflection. A Haxe anonymous object reaches them as a runtime map whose
	String fields are `*string` carriers, and a Haxe class has no exported Go
	fields. Ports of Go code that pass `map[string]interface{}` data, such as a
	template data root, need the map that Go code would build.

	How: The compiler maps this type to `map[string]any`. `new` creates an empty
	map, and `set` stores one entry with the key as a Go `string` and the value
	as a `go.NativeAny`, so `map.set("Count", 3)` stores a Go `int`. Pass the map
	to a Go parameter typed `go.NativeAny`, `Dynamic` or `Any`, or nest it with
	`set`. Go `%T` prints `map[string]interface {}`, and a template reads
	`.Issue.Title` from the key "Title" of the nested map stored under "Issue".
	A null map is a nil Go map; `set` on it fails as in Go.
	Call `set` directly; method values are not supported.
**/
extern class NativeAnyMap {
	/** Creates an empty native map. */
	public function new();

	/** Stores `value` under the Go string `key`, replacing an earlier entry. */
	public function set(key:String, value:NativeAny):Void;
}
