class BlockedInput extends haxe.io.Input {
	public function new() {}

	override public function readByte():Int {
		return throw new haxe.io.Eof();
	}

	override public function readBytes(buf:haxe.io.Bytes, pos:Int, len:Int):Int {
		return 0;
	}
}

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

	static function outsideInputTag():String {
		var input = new haxe.io.BytesInput(haxe.io.Bytes.ofString("a"));
		var result = "outsideInput=miss";
		try {
			input.readBytes(haxe.io.Bytes.alloc(1), 2, 1);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case OutsideBounds:
					"outsideInput=outside";
				case _:
					"outsideInput=other";
			};
		} catch (e:Dynamic) {
			result = "outsideInput=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function outsideOutputTag():String {
		var output = new haxe.io.BytesOutput();
		var result = "outsideOutput=miss";
		try {
			output.writeBytes(haxe.io.Bytes.ofString("a"), 2, 1);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case OutsideBounds:
					"outsideOutput=outside";
				case _:
					"outsideOutput=other";
			};
		} catch (e:Dynamic) {
			result = "outsideOutput=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function blockedTag():String {
		var output = new haxe.io.BytesOutput();
		var result = "blocked=miss";
		try {
			output.writeInput(new BlockedInput(), 1);
		} catch (e:haxe.io.Error) {
			result = switch (e) {
				case Blocked:
					"blocked=blocked";
				case _:
					"blocked=other";
			};
		} catch (e:Dynamic) {
			result = "blocked=dynamic:" + Std.string(e);
		}
		return result;
	}

	static function main() {
		Sys.println(overflowTag());
		Sys.println(customTag("boom"));
		Sys.println(customTag(123));
		Sys.println(writeThroughBaseOutputTag());
		Sys.println(outsideInputTag());
		Sys.println(outsideOutputTag());
		Sys.println(blockedTag());
	}
}
