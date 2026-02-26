class Main {
	static function main() {
		var s:String = "h\u00E9z";

		Sys.println("len=" + s.length);
		Sys.println("char0=" + s.charAt(0));
		Sys.println("char1=" + s.charAt(1));
		Sys.println("char2=" + s.charAt(2));
		Sys.println("char_oob=" + s.charAt(99));

		Sys.println("code0=" + s.charCodeAt(0));
		Sys.println("code1=" + s.charCodeAt(1));
		Sys.println("code2=" + s.charCodeAt(2));
		Sys.println("code_neg=" + s.charCodeAt(-1));
		Sys.println("code_oob=" + s.charCodeAt(99));

		Sys.println("substring_0_2=" + s.substring(0, 2));
		Sys.println("substring_1_99=" + s.substring(1, 99));
		Sys.println("substr_1_2=" + s.substr(1, 2));
		Sys.println("substr_neg_1=" + s.substr(-1));
	}
}
