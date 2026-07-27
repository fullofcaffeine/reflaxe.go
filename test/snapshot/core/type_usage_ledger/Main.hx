import go.Fmt;
import hxrt.string.GoStringRuntime;
import UsedBox.UsedSibling;

class Main {
	static inline final INCLUDE_DEAD_GENERIC = false;

	static function main() {
		if (INCLUDE_DEAD_GENERIC) {
			var dead = new DeadBox<String>("not emitted");
			Fmt.println(dead.value);
		}

		var box = new UsedBox<Int>(39);
		var nested = new UsedBox<UsedBox<Int>>(box);
		var sibling = new UsedSibling(1);
		var callback:Int->String = value -> Std.string(value);
		var details:{label:String, count:Int} = {
			label: callback(39),
			count: sibling.delta
		};
		Fmt.println(box.value);
		Fmt.println(38 + consumeDetails(details) + nestedDepth(nested) + GoStringRuntime.length("x"));
	}

	static function consumeDetails(details:{label:String, count:Int}):Int {
		return details.count;
	}

	static function nestedDepth(value:UsedBox<UsedBox<Int>>):Int {
		return value == null ? 0 : 1;
	}
}
