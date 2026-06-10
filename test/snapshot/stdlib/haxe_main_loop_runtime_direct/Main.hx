class Main {
	static function main() {
		var ran = false;
		var event:haxe.MainLoop.MainEvent = null;
		event = haxe.MainLoop.add(function() {
			ran = true;
			Sys.println("mainloop.add=ran");
			event.stop();
		});
		haxe.EntryPoint.run();
		Sys.println("mainloop.after=" + ran);

		haxe.Timer.delay(function() {
			Sys.println("timer.delay=ran");
		}, 1);
		haxe.EntryPoint.run();
	}
}
