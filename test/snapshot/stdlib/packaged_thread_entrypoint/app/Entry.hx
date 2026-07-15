package app;

import sys.thread.Lock;
import sys.thread.Thread;

class Entry {
	public static function main() {
		Sys.println("packaged-entry");
		var delay = new Lock();
		Thread.create(function() {
			delay.wait(0.02);
			Sys.println("packaged-worker");
		});
	}
}
