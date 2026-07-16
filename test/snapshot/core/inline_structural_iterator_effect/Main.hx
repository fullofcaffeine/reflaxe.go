/**
	What: Pins generated Go for an effectful inline structural iterator coercion.
	Why: The effect prefix and typed live-array cursor must coexist without retaining
	an erased ArrayIterator constructor.
	How: Count one inline effect, mutate the captured array, and consume the iterator.
**/
class Main {
	static var effectCount = 0;
	static var argumentEffectCount = 0;

	static inline function checkedIterator(values:Array<Int>):Iterator<Int> {
		effectCount++;
		return values.iterator();
	}

	static inline function checkedArgumentIterator(values:Array<Int>):Iterator<Int> {
		argumentEffectCount++;
		return values.iterator();
	}

	static function collect(iterator:Iterator<Int>):String {
		var values = [];
		while (iterator.hasNext()) {
			values.push(iterator.next());
		}
		return values.join(",");
	}

	static function main() {
		var values = [1, 2];
		var iterator:Iterator<Int> = checkedIterator(values);
		values[0] = 9;
		Sys.println("effect-before=" + effectCount);
		Sys.println("values=" + collect(iterator));
		Sys.println("effect-after=" + effectCount);
		Sys.println("argument-values=" + collect(checkedArgumentIterator([3, 4])));
		Sys.println("argument-effects=" + argumentEffectCount);
	}
}
