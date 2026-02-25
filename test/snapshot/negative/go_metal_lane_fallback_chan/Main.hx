@:goMetal
class Main {
	static function main() {
		var ch:go.Chan<Dynamic> = go.Go.newChan();
		ch.send("lane");
	}
}
