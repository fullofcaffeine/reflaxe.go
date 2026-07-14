package reflaxe.go.analyze;

#if macro
import haxe.macro.Type;
import reflaxe.go.analyze.MetalLaneAnalyzer.MetalLaneDeclaration;
import reflaxe.go.analyze.MetalLaneAnalyzer.MetalLaneSnapshot;

/**
	Why
	New code should not expose the historical lane name in its signatures.

	What
	Canonical typed alias for one native-boundary declaration.

	How
	Aliases the compatibility parser record without changing its public shape.
**/
typedef GoNativeBoundaryDeclaration = MetalLaneDeclaration;

/**
	Why
	Boundary consumers need an explicit canonical return type.

	What
	Canonical typed alias for the deterministic module/declaration snapshot.

	How
	Preserves the compatibility parser representation while callers migrate to
	the positive native-boundary API.
**/
typedef GoNativeBoundarySnapshot = MetalLaneSnapshot;

/**
	Why
	Module-level native authority needs a profile-independent, positively named
	entry point while old macro imports remain compatible.

	What
	Collects modules declared with canonical `@:goNative` metadata or the
	compatibility alias `@:goMetal`.

	How
	Delegates parsing to the compatibility module and returns its deterministic
	module/declaration snapshot.
**/
class GoNativeBoundaryAnalyzer {
	public static function collect(moduleTypes:Array<ModuleType>):GoNativeBoundarySnapshot {
		return MetalLaneAnalyzer.collect(moduleTypes);
	}
}
#end
