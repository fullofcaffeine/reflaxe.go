package x;

/** Its default file name would end in `_js.go`, which Go builds only for GOOS=js. **/
class Js {
	public static function name():String {
		return "x.Js";
	}
}
