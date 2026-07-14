import go.Fmt;

class Main {
	static function main():Void {
		var normalized = StringTools.replace("metal_zip", "_", "-");
		Fmt.println(normalized == "metal-zip" ? 707 : -1);
	}
}
