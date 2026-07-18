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
			case _:
				false;
		};
	}

	public static inline function isCompilerOwnedModule(moduleName:String):Bool {
		return isCompilerOwnedAuthority(moduleName);
	}

	/**
		What:
		Report whether a remaining compiler-owned class is safe to embed as a
		concrete superclass carrier in a generated Go struct.

		Why:
		The former socket and IO structs were the last such carriers. Remaining
		authorities do not currently materialize as source-embeddable classes.

		How:
		Return false until an individually reviewed compiler authority genuinely
		materializes as an embeddable struct again.
	**/
	public static function isEmbeddableCompilerOwnedSuper(_name:String):Bool {
		return false;
	}

	public static inline function canConstructEmptyTypeValue(goTypeName:String):Bool {
		return goTypeName != null && goTypeName != "" && !StringTools.startsWith(goTypeName, "*");
	}
}
