import go.Go;

class Main {
	static function main():Void {
		var channel:go.Chan<Int> = Go.newChan();
		channel.close();
	}
}
