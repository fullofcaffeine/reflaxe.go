class Main {
	static function main() {
		var ch:go.Chan<Int> = go.Go.newChan();
		ch.close();
		trace("ok");
	}
}
