import go.Go;

class Main {
	static function main() {
		var ch:go.Chan<Dynamic> = Go.newChan();
		ch.send("hello");
	}
}
