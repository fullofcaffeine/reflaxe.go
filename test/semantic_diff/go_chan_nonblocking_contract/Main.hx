import go.Chan;
import go.Go;

class Main {
	static function main() {
		var ch:Chan<Int> = Go.newChan(1);

		Sys.println("trySend.1=" + ch.trySend(3));
		Sys.println("trySend.2=" + ch.trySend(4));

		Sys.println("recvOr.1=" + ch.recvOr(-1));
		Sys.println("recvOr.2=" + ch.recvOr(99));

		var empty = ch.tryRecv();
		Sys.println("tryRecv.empty.isErr=" + empty.isErr());

		ch.send(7);
		var got = ch.tryRecv();
		Sys.println("tryRecv.got.isOk=" + got.isOk());
		Sys.println("tryRecv.got.unwrap=" + got.unwrap());
	}
}
