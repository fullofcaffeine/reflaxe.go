import Main.ErrorMatch;
import Main.NativeErrors;
import Main.PathError;

/** AsType is a Go 1.26 API; the file constraint records its native minimum. */
@:goBuildConstraint("go1.26")
@:goNative
class NativeErrorProbe {
	public static function check():Void {
		final match:ErrorMatch<PathError> = NativeErrors.asType(null);
		Sys.println(match.matched);
		Sys.println(match.value == null);
	}
}
