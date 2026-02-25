import go.Go;

class Main {
	static function main() {
		var ch:go.Chan<Int> = Go.newChan();
		ch.close();
		trace("ok");
	}
}
