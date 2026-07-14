package reflaxe.go.compiler;

#if macro
import haxe.macro.Type;
import reflaxe.go.compiler.GoNativeTypeEligibility.GoNativeEligibilityRole;
import reflaxe.go.compiler.GoNativeTypeEligibility.GoNativeTypeEligibilityResult;

/** Compatibility alias for the historical profile-shaped role name. */
typedef GoMetalEligibilityRole = GoNativeEligibilityRole;

/** Compatibility alias for the historical profile-shaped result name. */
typedef GoMetalTypeEligibilityResult = GoNativeTypeEligibilityResult;

/**
	Compatibility entry point. New compiler code should call
	`GoNativeTypeEligibility`; eligibility is profile-independent.
**/
class GoMetalTypeEligibility {
	public static inline function resolve(type:Type, role:GoMetalEligibilityRole, classTypeName:ClassType->String,
			enumTypeName:EnumType->String):GoMetalTypeEligibilityResult {
		return GoNativeTypeEligibility.resolve(type, role, classTypeName, enumTypeName);
	}
}
#end
