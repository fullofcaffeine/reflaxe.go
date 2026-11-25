import go.NativeAny;

/** A go-free module-level function next to one whose result type is `go.*`. */
@:goNative function plain():String {
	return "plain";
}

@:goNative function native():NativeAny {
	return NativeAny.ofString("native");
}
