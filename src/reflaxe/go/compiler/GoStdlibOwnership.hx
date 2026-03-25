package reflaxe.go.compiler;

/**
	Canonical compiler-owned stdlib authority classifier.

	Why:
	- `GoCompiler` needs one source of truth for which stdlib surfaces are still
	  intentionally compiler-owned.
	- Source-owned std routing, module splitting, and ownership docs all depend on
	  the same answer.

	What:
	- Defines the current compiler-owned stdlib authorities.
	- Exposes a small utility for type-value materialization eligibility.

	How:
	- Keep the authority list centralized here instead of duplicating switch blocks
	  across `GoCompiler` and future compiler-owned emitters.
**/
class GoStdlibOwnership {
	public static function isCompilerOwnedAuthority(name:String):Bool {
		return switch (name) {
			case "EReg", "haxe.ds.EnumValueMap", "haxe.io.BufferInput", "haxe.io.Bytes", "haxe.io.BytesBuffer", "haxe.io.BytesInput", "haxe.io.BytesOutput",
				"haxe.io.Eof", "haxe.io.Error", "haxe.io.Input", "haxe.io.Output", "haxe.io.StringInput", "sys.Http", "sys.io.File", "sys.io.FileInput",
				"sys.io.FileOutput", "sys.io.FileSeek":
				true;
			case _:
				false;
		};
	}

	public static inline function isCompilerOwnedModule(moduleName:String):Bool {
		return isCompilerOwnedAuthority(moduleName);
	}

	public static inline function canConstructEmptyTypeValue(goTypeName:String):Bool {
		return goTypeName != null && goTypeName != "" && !StringTools.startsWith(goTypeName, "*");
	}
}
