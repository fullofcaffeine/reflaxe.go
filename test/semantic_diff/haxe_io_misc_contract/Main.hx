import haxe.Int64;
import haxe.io.BufferInput;
import haxe.io.Bytes;
import haxe.io.BytesData;
import haxe.io.Encoding;
import haxe.io.Eof;
import haxe.io.Error;
import haxe.io.FPHelper;
import haxe.io.Mime;
import haxe.io.Scheme;
import haxe.io.StringInput;

class Main {
	static function errorTag(err:Error):String {
		return switch (err) {
			case Blocked:
				"blocked";
			case Overflow:
				"overflow";
			case OutsideBounds:
				"outside";
			case Custom(v):
				"custom:" + Std.string(v);
		};
	}

	static function readTwo(input:haxe.io.Input):String {
		var out = Bytes.alloc(2);
		input.readBytes(out, 0, 2);
		return out.toString();
	}

	static function main() {
		var stringInput = new StringInput("abc");
		Sys.println("stringInput=" + stringInput.readByte() + "," + readTwo(stringInput));

		var buffered = new BufferInput(new StringInput("wxyz"), Bytes.alloc(2));
		Sys.println("bufferInput=" + buffered.readByte() + "," + readTwo(buffered));

		var source = Bytes.ofString("ab");
		var data:BytesData = source.getData();
		data[1] = "Z".code;
		Sys.println("bytesData=" + Bytes.ofData(data).toString());

		Sys.println("encoding=" + switch (Encoding.RawNative) {
			case RawNative:
				"raw";
			case UTF8:
				"utf8";
		});

		try {
			throw new Eof();
		} catch (e:Eof) {
			Sys.println("eof=" + Std.string(e));
		}

		try {
			throw Error.Custom("boom");
		} catch (e:Error) {
			Sys.println("error=" + errorTag(e));
		}

		var mime:Mime = Mime.ApplicationJson;
		var scheme:Scheme = Scheme.Https;
		Sys.println("mime=" + mime);
		Sys.println("scheme=" + scheme);

		Sys.println("fp.i32ToFloat=" + FPHelper.i32ToFloat(1065353216));
		Sys.println("fp.floatToI32=" + FPHelper.floatToI32(1.5));
		Sys.println("fp.i64ToDouble=" + FPHelper.i64ToDouble(0, 1072693248));
		Sys.println("fp.doubleToI64=" + Int64.toStr(FPHelper.doubleToI64(1.0)));
	}
}
