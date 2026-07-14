import go.Go;

class Main {
	static function main():Void {
		var channel:go.Chan<Dynamic> = Go.newChan();
		channel.send("hello");
	}
}
