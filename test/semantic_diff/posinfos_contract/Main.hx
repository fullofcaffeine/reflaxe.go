import haxe.PosInfos;

class Main {
	static function probe(label:String, ?pos:PosInfos):Void {
		Sys.println(label + ".class=" + pos.className);
		Sys.println(label + ".method=" + pos.methodName);
		Sys.println(label + ".line=" + pos.lineNumber);
		Sys.println(label + ".fileNonEmpty=" + (pos.fileName != ""));
	}

	static function main() {
		probe("main");
		var fn = function() {
			probe("closure");
		}
		fn();
	}
}
