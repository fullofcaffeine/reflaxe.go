import go.Chan;
import go.Go;

class Main {
	static function main() {
		var buffered:Chan<Int> = Go.newChan(1);
		Sys.println(buffered.trySend(10));
		Sys.println(buffered.trySend(11));
		Sys.println(buffered.recvOr(-1));
		Sys.println(buffered.recvOr(-1));
		buffered.close();

		var direct = new Chan<Int>();
		Go.spawn(function() {
			direct.send(42);
		});
		Sys.println(direct.recv());
		direct.close();
	}
}
