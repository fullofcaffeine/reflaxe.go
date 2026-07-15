import go.Chan;
import go.Go;
import sys.thread.Thread;

@:go.import("log")
extern class GoLog {
	@:go.name("Panic")
	public static function panic(message:String):Void;
}

class Main {
	static function main():Void {
		Sys.println("native-goroutine-start");
		var started:Chan<Bool> = Go.newChan();
		var never:Chan<Bool> = Go.newChan();
		Go.spawn(() -> {
			Thread.current();
			started.send(true);
			GoLog.panic("native-goroutine-failure");
		});
		started.recv();
		never.recv();
	}
}
