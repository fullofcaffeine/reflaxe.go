import go.Chan;
import go.Go;

class Main {
	static function main() {
		var ch:Chan<Int> = Go.newChan(1);

		var empty = ch.tryRecv();
		Sys.println(empty.value);

		ch.send(9);
		var got = ch.tryRecv();
		Sys.println(got.value);

		var emptyAgain = ch.tryRecv();
		Sys.println(emptyAgain.value);
	}
}
