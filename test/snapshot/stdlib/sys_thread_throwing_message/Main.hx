import sys.thread.Thread;

class BadMessage extends haxe.Exception {
	public function new() {
		super("original");
	}

	override function get_message():String {
		throw "message failed";
	}
}

class Main {
	static function main() {
		Thread.create(function() {
			throw new BadMessage();
		});
		Thread.createWithEventLoop(function() {
			throw new BadMessage();
		});
		Sys.println("main-survived");
	}
}
