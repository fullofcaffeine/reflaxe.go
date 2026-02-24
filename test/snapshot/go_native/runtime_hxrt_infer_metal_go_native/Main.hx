import go.Chan;
import go.Go;

class Main {
	static function main() {
		var channel:Chan<Int> = Go.newChan();
		Go.spawn(function() {
			channel.send(1);
		});
		var value = channel.recv();
		Sys.println(value);
	}
}
