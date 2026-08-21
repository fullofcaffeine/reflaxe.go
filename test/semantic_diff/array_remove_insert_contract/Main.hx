class ArrayMutationBox {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

typedef ArrayMutationHolder = {
	var values:Array<Int>;
}

/**
	What: Exercises portable Array removal and insertion semantics.
	Why: Go slices have no matching methods, and library code must not rebuild
	arrays manually to preserve Haxe behavior.
	How: Cover remove, insert, and shift across equality, identity, nullable,
	generic, alias, position, and anonymous-object field cases.
**/
class Main {
	static var events:Array<String> = [];
	static var holderEvaluations:Int = 0;

	static function makeSame():String {
		return String.fromCharCode(115) + "ame";
	}

	static function removeGeneric<T>(first:T, second:T, value:T):String {
		var values = [first, second];
		return values.remove(value) + ":" + values.length + ":" + Std.string(values[0]);
	}

	static function removeGenericCount<T>(first:T, second:T, value:T):String {
		var values = [first, second];
		return values.remove(value) + ":" + values.length;
	}

	static function removeGenericFour<T>(first:T, second:T, third:T, fourth:T, value:T):String {
		var values = [first, second, third, fourth];
		return values.remove(value) + ":" + [for (item in values) Std.string(item)].join(",");
	}

	static function insertGeneric<T>(first:T, second:T, pos:Int, value:T):String {
		var values = [first, second];
		values.insert(pos, value);
		return values.length + ":" + Std.string(values[1]);
	}

	static function shiftGeneric<T>(first:T, second:T):String {
		var values = [first, second];
		return Std.string(values.shift()) + ":" + values.length + ":" + Std.string(values[0]);
	}

	static function showNullableInts(values:Array<Null<Int>>):String {
		return [for (value in values) Std.string(value)].join(",");
	}

	static function makeHolder():ArrayMutationHolder {
		return {values: [1, 2, 1]};
	}

	static function makeCountedHolder():ArrayMutationHolder {
		holderEvaluations++;
		return makeHolder();
	}

	static function markedPosition():Int {
		events.push("position");
		return -1;
	}

	static function markedInsertValue():Int {
		events.push("value");
		return 2;
	}

	static function markedRemoveValue():Int {
		events.push("value");
		return 1;
	}

	static function main():Void {
		var duplicate = [1, 2, 1];
		Sys.println("remove.duplicate=" + duplicate.remove(1) + ":" + duplicate.join(","));
		Sys.println("remove.missing=" + duplicate.remove(9) + ":" + duplicate.join(","));

		var strings = [makeSame(), "tail"];
		Sys.println("remove.string=" + strings.remove(makeSame()) + ":" + strings.join(","));

		var nullableInts:Array<Null<Int>> = [null, 1, null];
		Sys.println("remove.null=" + nullableInts.remove(null) + ":" + showNullableInts(nullableInts));
		var nullableStrings:Array<Null<String>> = [null, "A", "null", "B"];
		Sys.println("remove.nullString.literal=" + nullableStrings.remove("null") + ":" + nullableStrings.join(","));
		Sys.println("remove.nullString.null=" + nullableStrings.remove(null) + ":" + nullableStrings.join(","));

		var first = new ArrayMutationBox(1);
		var second = new ArrayMutationBox(2);
		var boxes = [first, second];
		Sys.println("remove.object.other=" + boxes.remove(new ArrayMutationBox(1)) + ":" + boxes.length);
		Sys.println("remove.object.exact=" + boxes.remove(first) + ":" + boxes.length);

		var atStart = [1, 2];
		atStart.insert(0, 0);
		Sys.println("insert.start=" + atStart.join(","));

		var inMiddle = [1, 3];
		inMiddle.insert(1, 2);
		Sys.println("insert.middle=" + inMiddle.join(","));

		var atEnd = [1, 2];
		atEnd.insert(atEnd.length, 3);
		Sys.println("insert.end=" + atEnd.join(","));

		var oversized = [1, 2];
		oversized.insert(99, 3);
		Sys.println("insert.oversized=" + oversized.join(","));

		var negative = [1, 3];
		negative.insert(-1, 2);
		Sys.println("insert.negative=" + negative.join(","));

		var tooNegative = [2, 3];
		tooNegative.insert(-99, 1);
		Sys.println("insert.tooNegative=" + tooNegative.join(","));

		var empty = new Array<Int>();
		empty.insert(-1, 1);
		Sys.println("insert.empty=" + empty.join(","));

		events = [];
		var orderedInsert = [1, 3];
		orderedInsert.insert(markedPosition(), markedInsertValue());
		Sys.println("order.insert=" + events.join(",") + ":" + orderedInsert.join(","));
		events = [];
		var orderedRemove = [1, 2];
		var removedInOrder = orderedRemove.remove(markedRemoveValue());
		Sys.println("order.remove=" + events.join(",") + ":" + removedInOrder + ":" + orderedRemove.join(","));

		Sys.println("generic.remove.string=" + removeGeneric(makeSame(), "tail", makeSame()));
		Sys.println("generic.remove.object.other=" + removeGenericCount(first, second, new ArrayMutationBox(1)));
		Sys.println("generic.remove.object.exact=" + removeGenericCount(first, second, first));
		Sys.println("generic.insert.string=" + insertGeneric(makeSame(), "tail", -1, "middle"));
		Sys.println("generic.remove.null=" + removeGeneric(null, 2, null));
		Sys.println("generic.remove.nullString=" + removeGenericFour(null, "A", "null", "B", "null"));
		Sys.println("generic.insert.null=" + insertGeneric(null, 2, -99, null));
		Sys.println("generic.shift.string=" + shiftGeneric(makeSame(), "tail"));
		Sys.println("generic.shift.null=" + shiftGeneric(null, 2));

		var shifted = [1, 2];
		var shiftedAlias = shifted;
		Sys.println("shift.value=" + shifted.shift() + ":" + shiftedAlias.join(","));
		var emptyShift = new Array<Null<Int>>();
		Sys.println("shift.empty=" + Std.string(emptyShift.shift()) + ":" + emptyShift.length);
		var ignoredShift = [1, 2];
		ignoredShift.shift();
		Sys.println("shift.ignored=" + ignoredShift.join(","));

		var holder = makeHolder();
		Sys.println("field.remove=" + holder.values.remove(1) + ":" + holder.values.join(","));
		holder.values.insert(1, 1);
		Sys.println("field.insert=" + holder.values.join(","));

		holderEvaluations = 0;
		var removedFromTemporary = makeCountedHolder().values.remove(1);
		Sys.println("receiver.remove=" + removedFromTemporary + ":" + holderEvaluations);
		holderEvaluations = 0;
		makeCountedHolder().values.insert(1, 1);
		Sys.println("receiver.insert=" + holderEvaluations);
	}
}
