package reflaxe.go;

#if macro
import haxe.macro.Context;
#end

class RawNativeModeResolver {
	public static inline final DEFINE_NAME = "reflaxe_go_raw_native_mode";

	#if macro
	public static function resolve():RawNativeMode {
		var raw = Context.definedValue(DEFINE_NAME);
		if (raw == null) {
			return RawNativeMode.Interp;
		}

		var normalized = StringTools.trim(raw).toLowerCase();
		if (normalized == "" || normalized == "interp") {
			return RawNativeMode.Interp;
		}

		return switch (normalized) {
			case "utf16le": RawNativeMode.Utf16LE;
			case _:
				Context.fatalError('Invalid value "' + raw + '" for -D ' + DEFINE_NAME + ' (expected interp|utf16le)', Context.currentPos());
				RawNativeMode.Interp;
		}
	}
	#else
	public static function resolve():RawNativeMode {
		return RawNativeMode.Interp;
	}
	#end

	public static function label(mode:RawNativeMode):String {
		return switch (mode) {
			case RawNativeMode.Interp: "interp";
			case RawNativeMode.Utf16LE: "utf16le";
		}
	}
}
