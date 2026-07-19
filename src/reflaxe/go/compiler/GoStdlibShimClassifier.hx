package reflaxe.go.compiler;

#if macro
import haxe.macro.Type;

private typedef GoStdlibShimSurface = {
	final kind:String;
	final path:String;
	final groups:Array<String>;
}

/**
	What:
	- Selects the remaining compiler-emitted standard-library groups for exact
	  typed Haxe class and enum symbols.

	Why:
	- Broad package/name predicates hid which public symbols still caused compiler
	  stdlib ownership and made additions difficult to audit bidirectionally.

	How:
	- Keeps one flat, machine-auditable list of fully qualified symbols.
	- `test_compiler_stdlib_intrinsic_registry.py` compares every entry with the
	  canonical ownership decision in `docs/compiler-stdlib-intrinsics.json`.
**/
class GoStdlibShimClassifier {
	static final SURFACES:Array<GoStdlibShimSurface> = [{kind: "class", path: "Type", groups: ["type_metadata"]}];

	public static function requiredGroupsForClass(classType:ClassType):Array<String> {
		return requiredGroups("class", qualifiedPath(classType.pack, classType.name));
	}

	public static function requiredGroupsForEnum(enumType:EnumType):Array<String> {
		return requiredGroups("enum", qualifiedPath(enumType.pack, enumType.name));
	}

	static function requiredGroups(kind:String, path:String):Array<String> {
		for (surface in SURFACES) {
			if (surface.kind == kind && surface.path == path) {
				return surface.groups.copy();
			}
		}
		return [];
	}

	static function qualifiedPath(pack:Array<String>, name:String):String {
		var packageName = pack.join(".");
		return packageName == "" ? name : packageName + "." + name;
	}
}
#end
