class Main {
	static function main():Void {
		var mutex = new sys.thread.Mutex();
		mutex.acquire();
		var message = StringTools.replace(StringTools.trim("  release-zip  "), "-", " ");
		mutex.release();
		Sys.println(message);
	}
}
