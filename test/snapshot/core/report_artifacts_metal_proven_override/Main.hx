import go.Chan;
import go.Go;

class Main {
	static function main():Void {
		var channel:Chan<Int> = Go.newChan(1);
		channel.send(13);
		Sys.println(channel.recv());
		channel.close();
	}
}
