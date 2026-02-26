class Main {
	static function intArrayToString(values:Array<Int>):String {
		var out = "";
		for (i in 0...values.length) {
			if (i > 0) {
				out += ",";
			}
			out += values[i];
		}
		return out;
	}

	static function main() {
		var values = [3, 1, 2];
		values.push(4);
		values.pop();
		values.push(8);
		Sys.println("array.len=" + values.length);
		Sys.println("array.values=" + intArrayToString(values));
		Sys.println("array.index0_2=" + values[0] + ":" + values[2]);

		var iterSum = 0;
		var iterTrace = [];
		for (i in 2...7) {
			iterSum += i;
			iterTrace.push(i);
		}
		Sys.println("iter.sum=" + iterSum);
		Sys.println("iter.trace=" + intArrayToString(iterTrace));

		var s = "Abc-Abc";
		Sys.println("string.len=" + s.length);
		Sys.println("string.charAt=" + s.charAt(0) + ":" + s.charAt(99));
		Sys.println("string.charCodeAt=" + s.charCodeAt(1));
		Sys.println("string.substring=" + s.substring(4, 7));
		Sys.println("string.fromCharCode=" + String.fromCharCode(65) + String.fromCharCode(122));
	}
}
