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
			case "EReg", "haxe.io.BufferInput", "haxe.io.Bytes", "haxe.io.BytesBuffer", "haxe.io.BytesInput", "haxe.io.BytesOutput", "haxe.io.Eof",
				"haxe.io.Error", "haxe.io.Input", "haxe.io.Output", "haxe.io.StringInput", "sys.Http", "sys.net.Host", "sys.net.Socket", "sys.net.UdpSocket":
				true;
			case _:
				false;
		};
	}

	public static inline function isCompilerOwnedModule(moduleName:String):Bool {
		return isCompilerOwnedAuthority(moduleName);
	}

	/**
		What:
		Identify compiler-owned classes that are safe to embed as concrete
		superclass carriers in generated Go structs.

		Why:
		Some compiler-owned stdlib authorities, such as `haxe.io.Input`, lower to
		Go interfaces rather than concrete structs. Treating those as embeddable
		superclasses produces invalid `*interface` fields.

		How:
		Keep the allowlist narrow and concrete. Only compiler-owned surfaces that
		really materialize as struct carriers should return `true` here.
	**/
	public static function isEmbeddableCompilerOwnedSuper(name:String):Bool {
		return switch (name) {
			case "sys.net.Socket", "sys.net.UdpSocket":
				true;
			case _:
				false;
		};
	}

	public static inline function canConstructEmptyTypeValue(goTypeName:String):Bool {
		return goTypeName != null && goTypeName != "" && !StringTools.startsWith(goTypeName, "*");
	}
}
