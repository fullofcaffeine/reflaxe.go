import go.Chan;
import go.Go;

class Main {
	static function main() {
		var ch:Chan<Int> = Go.newChan(2);
		ch.send(10);
		ch.send(20);
		Sys.println(ch.recv());
		Sys.println(ch.recv());
	}
}
