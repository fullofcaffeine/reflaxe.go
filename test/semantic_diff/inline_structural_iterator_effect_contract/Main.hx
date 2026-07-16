/**
	What: Exercises an effectful inline method that returns a structural iterator.
	Why: Replacing only the inline block's final ArrayIterator must not discard or
	duplicate the statements that run before it.
	How: Count the inline effect, mutate the captured array, then compare Haxe and Go
	output before and after consuming the iterator.
**/
class Main {
	static var effectCount = 0;
	static var argumentEffectCount = 0;

	/**
		What: Produces an iterator after one observable source effect.
		Why: This is the prefix shape used by upstream inline `Xml.iterator()`.
		How: Increment the shared counter before delegating to `Array.iterator()`.
	**/
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
