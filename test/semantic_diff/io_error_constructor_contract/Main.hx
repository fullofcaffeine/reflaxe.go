class Main {
	static function overflowTag():String {
		var out = new haxe.io.BytesOutput();
		var result = "overflow=miss";
		try {
			out.writeInt8(500);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case Overflow:
					"overflow=overflow";
				case Custom(v):
					"overflow=custom:" + Std.string(v);
				case Blocked:
					"overflow=blocked";
				case OutsideBounds:
					"overflow=outside";
			};
		} catch (e:Dynamic) {
			result = "overflow=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function customTag(value:Dynamic):String {
		var result = "custom=miss";
		try {
			throw haxe.io.Error.Custom(value);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case Custom(v):
					"custom=" + Std.string(v);
				case Overflow:
					"custom=overflow";
				case Blocked:
					"custom=blocked";
				case OutsideBounds:
					"custom=outside";
			};
		} catch (e:Dynamic) {
			result = "custom=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function writeThroughBaseOutputTag():String {
		var out:haxe.io.Output = new haxe.io.BytesOutput();
		var result = "base=miss";
		try {
			out.writeInt8(999);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case Overflow:
					"base=overflow";
				case _:
					"base=other";
			};
		} catch (e:Dynamic) {
			result = "base=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function main() {
		Sys.println(overflowTag());
		Sys.println(customTag("boom"));
		Sys.println(customTag(123));
		Sys.println(writeThroughBaseOutputTag());
	}
}
