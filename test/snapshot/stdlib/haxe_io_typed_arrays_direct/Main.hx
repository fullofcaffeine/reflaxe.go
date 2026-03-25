import haxe.io.ArrayBufferView;
import haxe.io.Bytes;
import haxe.io.Error;
import haxe.io.Float32Array;
import haxe.io.Float64Array;
import haxe.io.Int32Array;
import haxe.io.UInt16Array;
import haxe.io.UInt32Array;
import haxe.io.UInt8Array;

class Main {
	static function errTag(err:Error):String {
		return switch (err) {
			case OutsideBounds:
				"outside";
			case Blocked:
				"blocked";
			case Overflow:
				"overflow";
			case Custom(v):
				"custom:" + Std.string(v);
		};
	}

	static function fmtFloat(v:Float):String {
		return Std.string(Math.round(v * 100) / 100);
	}

	static function main() {
		var bytes = Bytes.alloc(12);
		for (i in 0...bytes.length) {
			bytes.set(i, i + 1);
		}

		var view = ArrayBufferView.fromBytes(bytes, 2, 6);
		Sys.println("view=" + view.byteOffset + "," + view.byteLength + "," + view.buffer.length);
		var sub = view.sub(1, 3);
		var subarray = view.subarray(2, 5);
		Sys.println("view.sub=" + sub.byteOffset + "," + sub.byteLength);
		Sys.println("view.subarray=" + subarray.byteOffset + "," + subarray.byteLength);

		var u8 = UInt8Array.fromBytes(bytes, 1, 4);
		u8[1] = 99;
		Sys.println("u8=" + u8.length + "," + u8[0] + "," + u8[1] + "," + bytes.get(2));

		var u16 = UInt16Array.fromArray([0x1234, 0x00FF, 0xCAFE], 0, 3);
		var u16sub = u16.sub(1, 2);
		Sys.println("u16=" + u16.length + "," + u16[0] + "," + u16sub.length + "," + u16sub[0]);

		var u32 = UInt32Array.fromArray([1, 2, 3], 0, 3);
		var u32view = u32.subarray(1, 3);
		Sys.println("u32=" + u32.length + "," + u32[2] + "," + u32view.length + "," + u32view[0]);

		var i32 = Int32Array.fromArray([-7, 42], 0, 2);
		i32[1] = -100;
		Sys.println("i32=" + i32.length + "," + i32[0] + "," + i32[1]);

		var f32 = new Float32Array(2);
		f32[0] = 1.25;
		f32[1] = 2.5;
		var f32sub = f32.subarray(1, 2);
		Sys.println("f32=" + f32.length + "," + fmtFloat(f32[0]) + "," + fmtFloat(f32sub[0]));

		var f64 = Float64Array.fromArray([3.5, -1.25], 0, 2);
		Sys.println("f64=" + f64.length + "," + fmtFloat(f64[0]) + "," + fmtFloat(f64[1]));

		try {
			ArrayBufferView.fromBytes(bytes, 9, 4);
			Sys.println("bounds=miss");
		} catch (err:Error) {
			Sys.println("bounds=" + errTag(err));
		}
	}
}
