package reflaxe.go.macros;

/**
	Compatibility entry point for integrations using the historical gate name.
	New code should initialize `NativeAuthorityGate`.
**/
class PortableNativeImportGate {
	public static inline function init():Void {
		NativeAuthorityGate.init();
	}
}
