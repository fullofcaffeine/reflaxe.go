package reflaxe.go.compiler;

class GoAutoLoweringModeTools {
	public static inline function label(mode:GoAutoLoweringMode):String {
		return switch (mode) {
			case Off:
				"off";
			case Auto:
				"auto";
			case AutoStrict:
				"auto_strict";
		};
	}
}
