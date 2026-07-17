import go.NativeSlice;

class Main {
	static function render(values:NativeSlice<Int>):String {
		var out = "";
		for (index in 0...values.length) {
			if (index > 0) {
				out += ",";
			}
			out += Std.string(values[index]);
		}
		return out;
	}

	static function main() {
		var source = [1, 2, 3];
		var native = NativeSlice.fromArray(source);
		native[0] = 7;
		var portable = native.toArray();
		portable.push(9);
		var dynamicSource:Array<Dynamic> = ["source"];
		var dynamicNative:NativeSlice<Dynamic> = NativeSlice.fromArray(dynamicSource);
		dynamicNative[0] = "native";
		var dynamicPortable = dynamicNative.toArray();
		dynamicPortable[0] = "portable";

		Sys.println("source=" + source.join(","));
		Sys.println("native=" + render(native));
		Sys.println("portable=" + portable.join(","));
		Sys.println("dynamic.source=" + Std.string(dynamicSource[0]));
		Sys.println("dynamic.native=" + Std.string(dynamicNative[0]));
		Sys.println("dynamic.portable=" + Std.string(dynamicPortable[0]));
	}
}
