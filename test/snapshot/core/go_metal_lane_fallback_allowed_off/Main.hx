@:goMetal
class Main {
	static function main() {
		// Intentional: unresolved go.Chan<T> remains allowed in lane modules when auto mode is off.
		var ch:go.Chan<Dynamic> = go.Go.newChan();
		ch.send("lane");
		trace("ok");
	}
}
