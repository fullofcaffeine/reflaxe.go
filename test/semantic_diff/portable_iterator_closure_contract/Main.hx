/**
	What
	Proves iterator state/ordering plus the supported closure capture, callback,
	and identity subset that current haxe.go preserves.

	Why
	Replacing an iterator with a Go `range` loop can duplicate evaluation or
	snapshot state. Treating every callback as an untracked Go function can lose
	shared captures or confuse two method closures that have different receivers.

	How
	The same source runs through the Haxe interpreter and reflaxe.go. Typed and
	`Dynamic` values share the fixture so native-eligible and fallback signatures
	are checked against one source oracle. Identity covers aliases, distinct
	closure instances, and repeated access to one receiver. Different-receiver
	method identity remains the explicit `haxe_go-vfp.7.11` blocker, so this
	fixture does not grant `haxe.Function` admission.
**/
class PortableIterator {
	final values:Array<Int>;
	var index = 0;

	public function new(values:Array<Int>) {
		this.values = values;
	}

	public function hasNext():Bool {
		return index < values.length;
	}

	public function next():Int {
		return values[index++];
	}
}

class PortableCounter {
	public var value:Int;

	public function new(value:Int) {
		this.value = value;
	}

	public function add(delta:Int):Int {
		value += delta;
		return value;
	}
}

class Main {
	static function forward<T>(iterator:Iterator<T>):Iterator<T> {
		return iterator;
	}

	static function consume(iterator:Iterator<Int>):String {
		final values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}

	static function callTwice(callback:Int->Void):Void {
		callback(2);
		callback(3);
	}

	static function makeAccumulator(seed:Int):Int->Int {
		var total = seed;
		return function(delta:Int):Int {
			total += delta;
			return total;
		};
	}

	static function main() {
		var evaluations = 0;
		final values = [1, 2];
		function source():PortableIterator {
			evaluations++;
			return new PortableIterator(values);
		}

		final iterator = forward(source());
		final iteratorAlias = iterator;
		Sys.println("iterator.evaluations=" + evaluations);
		Sys.println("iterator.ready=" + iterator.hasNext() + ":" + iterator.hasNext());
		Sys.println("iterator.first=" + iteratorAlias.next());
		values.push(3);
		Sys.println("iterator.rest=" + consume(iterator));
		Sys.println("iterator.exhausted=" + iteratorAlias.hasNext() + ":" + iterator.hasNext());

		final dynamicSource:Array<Dynamic> = [1, "two", null];
		final dynamicIterator:Iterator<Dynamic> = dynamicSource.iterator();
		final dynamicValues = [];
		while (dynamicIterator.hasNext()) {
			dynamicValues.push(Std.string(dynamicIterator.next()));
		}
		Sys.println("iterator.dynamic=" + dynamicValues.join(","));

		var captured = 1;
		final add = delta -> captured += delta;
		final addAlias = add;
		callTwice(addAlias);
		Sys.println("closure.capture=" + captured + ":" + add(4));

		final firstAccumulator = makeAccumulator(10);
		final firstAlias = firstAccumulator;
		final secondAccumulator = makeAccumulator(10);
		Sys.println("closure.reuse=" + firstAccumulator(1) + ":" + firstAlias(2) + ":" + secondAccumulator(1));
		Sys.println("closure.identity=" + Reflect.compareMethods(firstAccumulator, firstAlias) + ":"
			+ Reflect.compareMethods(firstAccumulator, secondAccumulator));

		final firstCounter = new PortableCounter(1);
		final firstMethod = firstCounter.add;
		final firstMethodAlias = firstMethod;
		final sameMethodAgain = firstCounter.add;
		firstCounter.value = 10;
		Sys.println("closure.bound=" + firstMethod(2) + ":" + firstCounter.value);
		Sys.println("closure.method-identity=" + Reflect.compareMethods(firstMethod, firstMethodAlias) + ":"
			+ Reflect.compareMethods(firstMethod, sameMethodAgain));
		Sys.println("closure.null-identity=" + Reflect.compareMethods(null, null));

		final loopCallbacks:Array<Void->Int> = [];
		for (index in 0...3) {
			loopCallbacks.push(() -> index);
		}
		Sys.println("closure.loop=" + [for (callback in loopCallbacks) callback()].join(","));

		final dynamicCallback:Dynamic->Dynamic = value -> value == null ? "nil" : Std.string(value);
		Sys.println("closure.dynamic=" + dynamicCallback(null) + ":" + dynamicCallback(7));
	}
}
