import haxe.PosInfos;

class Main {
	static function read(?pos:PosInfos):Void {
		var custom = pos.customParams;
		var count = 0;
		if (custom != null) {
			count = custom.length;
		}
		Sys.println("count=" + count);
		Sys.println("line=" + pos.lineNumber);
	}

	static function main() {
		read();
	}
}
