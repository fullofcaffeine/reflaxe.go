/** Local polymorphic callables retain their declared ABI at applied call sites. */
class Main {
	static function main():Void {
		function report<T>(expected:Array<T>, value:T):Void {
			Sys.println(Std.string(expected[0]) + ":" + Std.string(value));
		}
		report([2], 2);
		report(["word"], "word");
		function outer<A>(expected:Array<A>, value:A):Void {
			function inner<B>(items:Array<B>, item:B):Void {
				Sys.println(Std.string(items[0]) + ":" + Std.string(item));
			}
			inner(expected, value);
		}
		outer([42], 42);
	}
}
