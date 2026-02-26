import go.Chan;
import go.Go;

class Main {
	static function main():Void {
		var ch:Chan<Int> = Go.newChan(1);
		Sys.println(ch.trySend(3));
		Sys.println(ch.trySend(4));
		Sys.println(ch.recvOr(-1));
		Sys.println(ch.recvOr(-1));
		ch.close();
	}
}
