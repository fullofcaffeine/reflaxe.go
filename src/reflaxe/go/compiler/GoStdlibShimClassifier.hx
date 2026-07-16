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
	static final SURFACES:Array<GoStdlibShimSurface> = [
		{kind: "class", path: "EReg", groups: ["regex_serializer"]},
		{kind: "class", path: "Reflect", groups: ["stdlib_symbols"]},
		{kind: "class", path: "Std", groups: ["stdlib_symbols"]},
		{kind: "class", path: "Type", groups: ["stdlib_symbols"]},
		{kind: "class", path: "UnicodeString", groups: ["stdlib_symbols"]},
		{kind: "class", path: "_UnicodeString.UnicodeString_Impl_", groups: ["stdlib_symbols"]},
		{kind: "class", path: "haxe.Serializer", groups: ["regex_serializer"]},
		{kind: "class", path: "haxe.Unserializer", groups: ["regex_serializer"]},
		{kind: "class", path: "haxe.ds.BalancedTree", groups: ["stdlib_symbols"]},
		{kind: "class", path: "haxe.io.BufferInput", groups: ["io"]},
		{kind: "class", path: "haxe.io.Bytes", groups: ["io"]},
		{kind: "class", path: "haxe.io.BytesBuffer", groups: ["io"]},
		{kind: "class", path: "haxe.io.BytesInput", groups: ["io"]},
		{kind: "class", path: "haxe.io.BytesOutput", groups: ["io"]},
		{kind: "class", path: "haxe.io.Encoding", groups: ["io"]},
		{kind: "class", path: "haxe.io.Eof", groups: ["io"]},
		{kind: "class", path: "haxe.io.Input", groups: ["io"]},
		{kind: "class", path: "haxe.io.Output", groups: ["io"]},
		{kind: "class", path: "haxe.io.Path", groups: ["stdlib_symbols"]},
		{kind: "class", path: "haxe.io.StringInput", groups: ["io"]},
		{kind: "class", path: "sys.Http", groups: ["http"]},
		{kind: "class", path: "sys.net.Host", groups: ["net_socket"]},
		{kind: "class", path: "sys.net.Socket", groups: ["net_socket"]},
		{kind: "class", path: "sys.net.UdpSocket", groups: ["net_socket"]},
		{kind: "class", path: "sys.ssl.Socket", groups: ["net_socket"]},
		{kind: "class", path: "sys.ssl._Socket.Socket_Impl_", groups: ["net_socket"]},
		{kind: "enum", path: "haxe.ds.Option", groups: ["stdlib_symbols"]},
		{kind: "enum", path: "haxe.io.Error", groups: ["io"]}
	];

	public static function needsIoHelperSurface(classType:ClassType, fieldName:String, isIoInputHelperMethodName:String->Bool,
			isIoOutputHelperMethodName:String->Bool):Bool {
		if (classType.pack.join(".") != "haxe.io") {
			return false;
		}
		return switch (classType.name) {
			case "Input", "BytesInput", "BufferInput", "StringInput":
				isIoInputHelperMethodName(fieldName);
			case "Output", "BytesOutput":
				isIoOutputHelperMethodName(fieldName);
			case _:
				false;
		};
	}

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
