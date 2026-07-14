import sys.thread.Lock;
import sys.thread.Thread;

class Main {
	static function main():Void {
		var started = new Lock();
		Thread.create(() -> {
			started.release();
			new Lock().wait(0.02);
			Sys.println("child-before-throw");
			throw "child-failure";
		});
		started.wait();
		Sys.println("main-survived");
	}
}
