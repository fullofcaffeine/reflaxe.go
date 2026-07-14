import go.Chan;
import go.Go;

class Main {
	static function main() {
		var ch:Chan<Int> = Go.newChan(1);

		var empty = ch.tryRecv();
		Sys.println(empty.isErr());

		ch.send(9);
		var got = ch.tryRecv();
		Sys.println(got.isOk());
		Sys.println(got.unwrap());

		var emptyAgain = ch.tryRecv();
		Sys.println(emptyAgain.isErr());

		ch.send(12);
		ch.close();
		var bufferedBeforeClose = ch.tryRecv();
		Sys.println(bufferedBeforeClose.isOk());
		Sys.println(bufferedBeforeClose.unwrap());
		var closed = ch.tryRecv();
		Sys.println(closed.isErr());
		Sys.println(closed.error());
		Sys.println(ch.recvOr(-1));
	}
}
