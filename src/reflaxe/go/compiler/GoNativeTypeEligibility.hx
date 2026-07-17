package reflaxe.go.compiler;

#if macro
import haxe.macro.Type;

/**
	Why
	Eligibility rules differ for ordinary elements and comparable Go map keys.

	What
	Names the role a generic type plays in a typed native representation.

	How
	`GoNativeTypeEligibility.resolve` uses the role to apply the additional map-key
	comparability rule after central scalar mapping.
**/
enum abstract GoNativeEligibilityRole(String) {
	var ChanElement = "chan_element";
	var SliceElement = "slice_element";
	var MapKey = "map_key";
	var MapValue = "map_value";
	var ResultElement = "result_element";
}

/**
	Why
	Callers need stable failure evidence instead of a nullable type guess.

	What
	Returns eligibility, resolved Go type, and deterministic failure details.

	How
	Lowering and analyzers share this result for diagnostics, fallback policy, and
	report ledgers.
**/
typedef GoNativeTypeEligibilityResult = {
	final eligible:Bool;
	final goType:Null<String>;
	final reasonCode:Null<String>;
	final reason:Null<String>;
}

/**
	Why
	Typed representation eligibility is capability evidence shared by every
	policy preset; naming it after metal incorrectly makes it profile-owned.

	What
	Determines whether a Haxe generic role has a concrete, semantics-safe Go type.

	How
	Uses the central scalar mapper, rejects nullable/dynamic carriers, and applies
	Go comparability rules to map keys.
**/
class GoNativeTypeEligibility {
	public static function resolve(type:Type, role:GoNativeEligibilityRole, classTypeName:ClassType->String,
			enumTypeName:EnumType->String):GoNativeTypeEligibilityResult {
		if (type == null) {
			return ineligible("missing_type", "Could not resolve generic type for typed specialization.");
		}

		if (GoTypeMapper.isNullablePrimitiveType(type)) {
			return ineligible("nullable_primitive_dynamic_path",
				"Nullable primitive types currently lower through dynamic/null-boxed representation; typed specialization is disabled for semantic safety.");
		}

		// A portable Haxe Array is represented by a comparable pointer in generated
		// Go, but that backend detail is not authority for a Go-native map key. Keep
		// the shared Haxe carrier out of native map specialization so the boundary
		// cannot silently depend on compiler/runtime representation identity.
		if (role == GoNativeEligibilityRole.MapKey && GoTypeMapper.isHaxeArrayType(type)) {
			return ineligible("map_key_not_comparable",
				"Portable Haxe Array values are not eligible as Go-native map keys; choose an explicitly comparable Go key type.");
		}

		var goType = GoTypeMapper.scalarGoType(type, classTypeName, enumTypeName);
		if (goType == null || goType == "") {
			return ineligible("empty_go_type", "Could not resolve concrete Go type for typed specialization.");
		}
		if (goType == "any") {
			return ineligible("type_maps_to_any", "Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.");
		}
		if (role == GoNativeEligibilityRole.MapKey && !isComparableMapKeyType(goType)) {
			return ineligible("map_key_not_comparable", 'Go map keys must be comparable (key type resolved to ' + goType + ").");
		}

		return {
			eligible: true,
			goType: goType,
			reasonCode: null,
			reason: null
		};
	}

	static function isComparableMapKeyType(goType:String):Bool {
		if (goType == null || goType == "") {
			return false;
		}
		if (StringTools.startsWith(goType, "[]")) {
			return false;
		}
		if (StringTools.startsWith(goType, "map[")) {
			return false;
		}
		if (StringTools.startsWith(goType, "func(")) {
			return false;
		}
		if (StringTools.startsWith(goType, "[")) {
			var closing = goType.indexOf("]");
			if (closing <= 0 || closing + 1 >= goType.length) {
				return false;
			}
			var elementType = StringTools.trim(goType.substr(closing + 1));
			return isComparableMapKeyType(elementType);
		}
		return true;
	}

	static function ineligible(code:String, message:String):GoNativeTypeEligibilityResult {
		return {
			eligible: false,
			goType: null,
			reasonCode: code,
			reason: message
		};
	}
}
#end
