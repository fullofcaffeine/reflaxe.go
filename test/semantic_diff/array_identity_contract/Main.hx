class ArrayIdentityBox {
	public final id:Int;

	public function new(id:Int) {
		this.id = id;
	}
}

class ArrayIdentityHolder {
	public var values:Array<Int>;

	public function new(values:Array<Int>) {
		this.values = values;
	}
}

typedef ArrayIdentityAnon = {
	var values:Array<Int>;
}

/**
	What: Exercises Haxe `Array` identity across every observable length-changing boundary.
	Why: A Go slice copies its length header, while Haxe assignments and calls alias one
	mutable Array object whose later length and contents must remain shared.
	How: Mutate through locals, fields, returns, retained values, callbacks, erased
	generics, dynamic storage, and sparse indexed writes while recording evaluation order.
**/
class Main {
	static var retained:Array<Int>;
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

	static function popParameter(values:Array<Int>):Null<Int> {
		return values.pop();
	}

	static function returnAlias(values:Array<Int>):Array<Int> {
		return values;
	}

	static function retain(values:Array<Int>):Void {
		retained = values;
	}

	static function applyCallback(values:Array<Int>, callback:Array<Int>->Void):Void {
		callback(values);
	}

	static function pushGeneric<T>(values:Array<T>, value:T):Void {
		values.push(value);
	}

	static function returnGeneric<T>(values:Array<T>):Array<T> {
		return values;
	}

	static function retainErased<T>(values:Array<T>):Void {
		retainedGeneric = cast values;
	}

	static function applyGenericCallback<T>(values:Array<T>, value:T, callback:(Array<T>, T) -> Void):Void {
		callback(values, value);
	}

	static function makeHolder():ArrayIdentityHolder {
		holderEvaluations++;
		events.push("target");
		return new ArrayIdentityHolder([20]);
	}

	static function makeAnon():ArrayIdentityAnon {
		holderEvaluations++;
		return {values: [30]};
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
		Sys.println("local.push=" + show(local) + ":" + local.length + ":" + localAlias.length);
		localAlias.insert(1, 9);
		Sys.println("local.insert=" + show(local));
		var removed = localAlias.remove(2);
		Sys.println("local.remove=" + removed + ":" + show(local));
		var popped = localAlias.pop();
		Sys.println("local.pop=" + popped + ":" + show(local));

		var parameter = [4];
		pushParameter(parameter, 5);
		Sys.println("parameter.push=" + show(parameter));
		var parameterAlias = parameter;
		var parameterPopped = popParameter(parameterAlias);
		Sys.println("parameter.pop=" + parameterPopped + ":" + show(parameter));

		var returnedSource = [6];
		var returned = returnAlias(returnedSource);
		returned.push(7);
		Sys.println("return.push=" + show(returnedSource));

		var holder = new ArrayIdentityHolder([8]);
		var fieldAlias = holder.values;
		fieldAlias.insert(1, 9);
		Sys.println("field.insert=" + show(holder.values));

		var anon:ArrayIdentityAnon = {values: [10]};
		var anonAlias = anon.values;
		anon.values.push(11);
		Sys.println("anon.push=" + show(anonAlias));

		var callbackSource = [12];
		applyCallback(callbackSource, function(values) values.push(13));
		Sys.println("callback.push=" + show(callbackSource));

		var retainedSource = [14];
		retain(retainedSource);
		retained.push(15);
		Sys.println("retained.push=" + show(retainedSource));

		var dynamicSource = [16];
		var dynamicValue:Dynamic = dynamicSource;
		var dynamicAlias:Array<Int> = cast dynamicValue;
		dynamicAlias.push(17);
		Sys.println("dynamic.push=" + show(dynamicSource));

		var genericInts = [18];
		pushGeneric(genericInts, 19);
		var genericReturned = returnGeneric(genericInts);
		genericReturned.insert(1, 181);
		Sys.println("generic.int=" + show(genericInts));

		var genericStrings = ["a"];
		pushGeneric(genericStrings, "b");
		Sys.println("generic.string=" + show(genericStrings));
		Sys.println("generic.has=" + Lambda.has(genericStrings, "a"));
		Sys.println("generic.stringSplit=" + genericStrings[0].split("").length);

		var genericNullable:Array<Null<Int>> = [null];
		pushGeneric(genericNullable, 2);
		Sys.println("generic.nullable=" + show(genericNullable));

		var firstBox = new ArrayIdentityBox(1);
		var genericBoxes = [firstBox];
		pushGeneric(genericBoxes, new ArrayIdentityBox(2));
		Sys.println("generic.reference=" + genericBoxes.length + ":" + genericBoxes[0].id + ":" + genericBoxes[1].id);

		var retainedErasedSource = [20];
		retainErased(retainedErasedSource);
		retainedGeneric.push(21);
		Sys.println("generic.retained=" + show(retainedErasedSource));

		var genericCallbackSource = [22];
		applyGenericCallback(genericCallbackSource, 23, function(values, value) values.push(value));
		Sys.println("generic.callback=" + show(genericCallbackSource));

		var sparse:Array<Null<Int>> = [];
		var sparseAlias = sparse;
		sparse[2] = 7;
		Sys.println("sparse.nullable=" + sparseAlias.length + ":" + show(sparseAlias));
		sparseAlias.push(8);
		Sys.println("sparse.followup=" + show(sparse));

		var sparseInts:Array<Int> = [];
		sparseInts[2] = 9;
		Sys.println("sparse.int=" + sparseInts.length + ":" + show(sparseInts));

		var identity = [24];
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
		Sys.println("order.index=" + holderEvaluations + ":" + events.join(","));

		holderEvaluations = 0;
		events = [];
		makeHolder().values.push(markedValue());
		Sys.println("order.fieldPush=" + holderEvaluations + ":" + events.join(","));

		holderEvaluations = 0;
		events = [];
		makeAnon().values.insert(markedIndex(), markedValue());
		Sys.println("order.anonInsert=" + holderEvaluations + ":" + events.join(","));

		operationValues = [1, 2, 3];
		events = [];
		var compoundResult = markedOperationValues()[markedIndex()] += markedValue();
		Sys.println("order.compound=" + compoundResult + ":" + show(operationValues) + ":" + events.join(","));

		operationValues = [4, 5, 6];
		events = [];
		markedOperationValues()[markedIndex()] += markedValue();
		Sys.println("order.compoundStmt=" + show(operationValues) + ":" + events.join(","));

		operationValues = [7, 8, 9];
		events = [];
		var postResult = markedOperationValues()[markedIndex()]++;
		Sys.println("order.post=" + postResult + ":" + show(operationValues) + ":" + events.join(","));

		operationValues = [10, 11, 12];
		events = [];
		var preResult = ++markedOperationValues()[markedIndex()];
		Sys.println("order.pre=" + preResult + ":" + show(operationValues) + ":" + events.join(","));

		var compoundStrings = ["a"];
		var compoundStringsAlias = compoundStrings;
		compoundStringsAlias[0] += "b";
		Sys.println("compound.string=" + compoundStrings[0]);
	}
}
