package reflaxe.go.compiler;

#if macro
import haxe.macro.Type;

enum abstract GoMetalEligibilityRole(String) {
	var ChanElement = "chan_element";
	var SliceElement = "slice_element";
	var MapKey = "map_key";
	var MapValue = "map_value";
	var ResultElement = "result_element";
}

typedef GoMetalTypeEligibilityResult = {
	final eligible:Bool;
	final goType:Null<String>;
	final reasonCode:Null<String>;
	final reason:Null<String>;
}

class GoMetalTypeEligibility {
	public static function resolve(type:Type, role:GoMetalEligibilityRole, classTypeName:ClassType->String,
			enumTypeName:EnumType->String):GoMetalTypeEligibilityResult {
		if (type == null) {
			return ineligible("missing_type", "Could not resolve generic type for typed specialization.");
		}

		if (GoTypeMapper.isNullablePrimitiveType(type)) {
			return ineligible("nullable_primitive_dynamic_path",
				"Nullable primitive types currently lower through dynamic/null-boxed representation; typed specialization is disabled for semantic safety.");
		}

		var goType = GoTypeMapper.scalarGoType(type, classTypeName, enumTypeName);
		if (goType == null || goType == "") {
			return ineligible("empty_go_type", "Could not resolve concrete Go type for typed specialization.");
		}
		if (goType == "any") {
			return ineligible("type_maps_to_any", "Generic type resolves to `any`; typed specialization requires a concrete non-`any` Go type.");
		}
		if (role == GoMetalEligibilityRole.MapKey && !isComparableMapKeyType(goType)) {
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
		// Arrays are comparable only when their element type is comparable.
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

	static function ineligible(code:String, message:String):GoMetalTypeEligibilityResult {
		return {
			eligible: false,
			goType: null,
			reasonCode: code,
			reason: message
		};
	}
}
#end
