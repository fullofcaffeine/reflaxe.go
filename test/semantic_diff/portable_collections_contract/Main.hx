class PortableCollectionBox {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

/**
	What
	Exercises the portable collection behavior that a future typed Go carrier
	must preserve.

	Why
	A raw Go slice or map can look type-correct while losing Haxe aliasing,
	sparse-null, missing-value, nested-value, or callback-visible mutation
	semantics.

	How
	The same source runs through `--interp` and reflaxe.go. Fully typed and
	`Dynamic` shapes intentionally share the fixture so the registry's admitted
	and fallback cases keep one source-level oracle.
**/
class Main {
	static function mutateArray(values:Array<Int>, callback:Array<Int>->Void):Void {
		values.push(2);
		callback(values);
	}

	static function mutateStringMap(values:haxe.ds.StringMap<Int>, callback:haxe.ds.StringMap<Int>->Void):Void {
		values.set("second", 2);
		callback(values);
	}

	static function mutateIntMap(values:haxe.ds.IntMap<String>, callback:haxe.ds.IntMap<String>->Void):Void {
		values.set(2, "two");
		callback(values);
	}

	static function sortedStringKeys(values:haxe.ds.StringMap<Dynamic>):String {
		final keys = [for (key in values.keys()) key];
		haxe.ds.ArraySort.sort(keys, Reflect.compare);
		return keys.join(",");
	}

	static function sortedIntKeys(values:haxe.ds.IntMap<Dynamic>):String {
		final keys = [for (key in values.keys()) key];
		haxe.ds.ArraySort.sort(keys, Reflect.compare);
		return keys.join(",");
	}

	static function main() {
		final values = [1];
		final alias = values;
		mutateArray(values, callbackValues -> callbackValues.push(3));
		Sys.println("array.alias=" + alias.join(","));

		final sparse:Array<Int> = [];
		sparse[2] = 7;
		Sys.println("array.sparse=" + sparse.length + ":" + Std.string(sparse[0]) + ":" + sparse[2]);

		final nested = [[4]];
		final nestedAlias = nested[0];
		nested[0].push(5);
		Sys.println("array.nested=" + nestedAlias.join(","));

		final empty:Array<Int> = [];
		final absent:Null<Array<Int>> = null;
		Sys.println("array.empty-null=" + empty.length + ":" + (absent == null));

		final dynamicValues:Array<Dynamic> = [1, "two", null];
		final dynamicAlias = dynamicValues;
		dynamicValues.push(new PortableCollectionBox(3));
		Sys.println("array.dynamic=" + dynamicAlias.length + ":" + Std.string(dynamicAlias[2]));

		final strings = new haxe.ds.StringMap<Int>();
		final stringsAlias = strings;
		strings.set("first", 1);
		mutateStringMap(strings, callbackValues -> callbackValues.set("third", 3));
		Sys.println("string-map.alias=" + stringsAlias.get("second") + ":" + stringsAlias.get("third"));
		Sys.println("string-map.keys=" + sortedStringKeys(cast stringsAlias));

		final ints = new haxe.ds.IntMap<String>();
		final intsAlias = ints;
		ints.set(1, "one");
		mutateIntMap(ints, callbackValues -> callbackValues.set(3, "three"));
		Sys.println("int-map.alias=" + intsAlias.get(2) + ":" + intsAlias.get(3));
		Sys.println("int-map.keys=" + sortedIntKeys(cast intsAlias));

		final nestedMap = new haxe.ds.StringMap<Array<Int>>();
		final nestedMapValue = [8];
		nestedMap.set("items", nestedMapValue);
		nestedMap.get("items").push(9);
		Sys.println("string-map.nested=" + nestedMapValue.join(","));

		final nestedIntMap = new haxe.ds.IntMap<Array<Int>>();
		final nestedIntMapValue = [10];
		nestedIntMap.set(1, nestedIntMapValue);
		nestedIntMap.get(1).push(11);
		Sys.println("int-map.nested=" + nestedIntMapValue.join(","));
		final nestedIntMapCopy = nestedIntMap.copy();
		nestedIntMap.set(2, [20]);
		nestedIntMapValue.push(12);
		Sys.println("int-map.copy=" + nestedIntMapCopy.exists(2) + ":" + nestedIntMapCopy.get(1).join(","));

		final nullable = new haxe.ds.StringMap<Null<Int>>();
		nullable.set("present", null);
		Sys.println("string-map.null=" + nullable.exists("present") + ":" + (nullable.get("present") == null));

		final nullableIntMap = new haxe.ds.IntMap<Null<Int>>();
		nullableIntMap.set(7, null);
		Sys.println("int-map.null=" + nullableIntMap.exists(7) + ":" + (nullableIntMap.get(7) == null));
		final nullableIntMapCopy = nullableIntMap.copy();
		Sys.println("int-map.copy-null=" + nullableIntMapCopy.exists(7) + ":" + (nullableIntMapCopy.get(7) == null));

		final emptyMap = new haxe.ds.StringMap<Int>();
		Sys.println("string-map.empty-missing=" + emptyMap.keys().hasNext() + ":" + (emptyMap.get("missing") == null));

		final emptyIntMap = new haxe.ds.IntMap<Int>();
		Sys.println("int-map.empty-missing=" + emptyIntMap.keys().hasNext() + ":" + (emptyIntMap.get(7) == null));

		final dynamicMap = new haxe.ds.StringMap<Dynamic>();
		dynamicMap.set("number", 1);
		dynamicMap.set("nothing", null);
		Sys.println("string-map.dynamic=" + sortedStringKeys(dynamicMap) + ":" + dynamicMap.exists("nothing"));

		final dynamicIntMap = new haxe.ds.IntMap<Dynamic>();
		dynamicIntMap.set(1, "one");
		dynamicIntMap.set(2, null);
		Sys.println("int-map.dynamic=" + sortedIntKeys(dynamicIntMap) + ":" + dynamicIntMap.exists(2));
	}
}
