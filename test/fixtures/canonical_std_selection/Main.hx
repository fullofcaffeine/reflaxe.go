import go.Http;
import haxe.io.Bytes;
import haxe.io.BytesInput;
import hxrt.string.GoStringRuntime;

class Main {
	static function main():Void {
		Sys.println(Lambda.origin());
		Sys.println(GoStringRuntime.charCodeAt("A", 0));
		Sys.println(Http.statusText(200));
		var input = new BytesInput(Bytes.ofString("support-ok"));
		Sys.println(input.readAll().toString());
	}
}
