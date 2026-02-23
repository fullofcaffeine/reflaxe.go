import go.Chan;
import go.Go;

class Main {
	static function main() {
		var requests:Chan<Int> = Go.newChan();
		var responses:Chan<Int> = Go.newChan();

		Go.spawn(function() {
			var value = requests.recv();
			responses.send(value);
		});

		requests.send(41);
		Sys.println(responses.recv());

		var buffered:Chan<Int> = Go.newChan(1);
		Sys.println(buffered.trySend(7));
		Sys.println(buffered.trySend(8));
		Sys.println(buffered.recvOr(-1));
		Sys.println(buffered.recvOr(-1));
		buffered.close();
	}
}
