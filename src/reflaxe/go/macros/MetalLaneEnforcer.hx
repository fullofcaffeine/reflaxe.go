package reflaxe.go.macros;

/**
	Compatibility entry point for integrations that initialized the historical
	metal-lane macro directly. New code should use `NativeBoundaryEnforcer`.
**/
class MetalLaneEnforcer {
	public static inline function init():Void {
		NativeBoundaryEnforcer.init();
	}
}
