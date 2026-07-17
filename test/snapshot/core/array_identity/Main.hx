class SnapshotArrayIdentityBox {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

class SnapshotArrayIdentityHolder {
	public var values:Array<Int>;

	public function new(values:Array<Int>) {
		this.values = values;
	}
}

/**
	What: Pins the shared portable carrier used for Haxe `Array` values in generated Go.
	Why: Raw Go slice headers lose length changes across assignments, calls, returns,
	callbacks, erased generics, and fields even though Haxe keeps one mutable identity.
	How: Exercise a compact cross-section of aliases, generic retention, sparse growth,
	identity comparison, and once-only receiver/argument evaluation.
**/
class Main {
	static var retainedGeneric:Array<Dynamic>;
	static var holderEvaluations:Int = 0;
	static var events:Array<String> = [];
	static var operationValues:Array<Int> = [];

	static function show<T>(values:Array<T>):String {
		return [for (value in values) Std.string(value)].join(",");
	}

	static function pushParameter(values:Array<Int>, value:Int):Void {
		values.push(value);
	}

	static function returnAlias(values:Array<Int>):Array<Int> {
		return values;
	}

	static function pushGeneric<T>(values:Array<T>, value:T):Void {
		values.push(value);
	}

	static function retainErased<T>(values:Array<T>):Void {
		retainedGeneric = cast values;
	}

	static function applyGenericCallback<T>(values:Array<T>, value:T, callback:(Array<T>, T) -> Void):Void {
		callback(values, value);
	}

	static function makeHolder():SnapshotArrayIdentityHolder {
		holderEvaluations++;
		events.push("target");
		return new SnapshotArrayIdentityHolder([20]);
	}

	static function markedIndex():Int {
		events.push("index");
		return 2;
	}

	static function markedValue():Int {
		events.push("value");
		return 23;
	}

	static function markedOperationValues():Array<Int> {
		events.push("target");
		return operationValues;
	}

	static function main():Void {
		var local = [1, 2];
		var localAlias = local;
		localAlias.push(3);
		localAlias.insert(1, 9);
		localAlias.remove(2);
		localAlias.pop();
		Sys.println("local=" + show(local));

		var parameter = [4];
		pushParameter(parameter, 5);
		Sys.println("parameter=" + show(parameter));

		var returnedSource = [6];
		var returnedAlias = returnAlias(returnedSource);
		returnedAlias.push(7);
		Sys.println("return=" + show(returnedSource));

		var genericInts = [8];
		pushGeneric(genericInts, 9);
		Sys.println("generic=" + show(genericInts));
		Sys.println("generic.has=" + Lambda.has(["a", "b"], "a"));
		var stringElements = "left,right".split(";");
		Sys.println("string.split=" + stringElements[0].split(",").length);

		var retainedSource = [10];
		retainErased(retainedSource);
		retainedGeneric.push(11);
		Sys.println("retained=" + show(retainedSource));

		var callbackSource = [12];
		applyGenericCallback(callbackSource, 13, function(values, value) values.push(value));
		Sys.println("callback=" + show(callbackSource));

		var sparse:Array<Null<Int>> = [];
		var sparseAlias = sparse;
		sparse[2] = 14;
		sparseAlias.push(15);
		Sys.println("sparse=" + show(sparse));

		var identity = [16];
		var identityAlias = identity;
		var identityCopy = identity.copy();
		Sys.println("identity=" + (identity == identityAlias) + ":" + (identity == identityCopy));

		holderEvaluations = 0;
		events = [];
		makeHolder().values[
			{
				events.push("index");
				2;
			}
		] = {
			events.push("value");
			23;
		};
		Sys.println("order=" + holderEvaluations + ":" + events.join(","));

		operationValues = [1, 2, 3];
		events = [];
		var compoundResult = markedOperationValues()[markedIndex()] += markedValue();
		Sys.println("compound=" + compoundResult + ":" + show(operationValues) + ":" + events.join(","));

		operationValues = [4, 5, 6];
		events = [];
		var postResult = markedOperationValues()[markedIndex()]++;
		Sys.println("post=" + postResult + ":" + show(operationValues) + ":" + events.join(","));
	}
}
