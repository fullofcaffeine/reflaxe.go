import reflaxe.go.ast.GoType;

/** Independent contracts for the type identities used by callable recovery. */
class GoTypeRecoveryContract {
	static function require(value:Bool, message:String):Void {
		if (!value)
			throw message;
	}

	static function main():Void {
		var concrete = GoType.parse("*hxrt.Function[func(int) int]");
		var erased = GoType.parse("*hxrt.Function[func(any) any]");
		require(!concrete.sameRepresentation(erased), "different callable payloads collapsed");
		require(concrete.sameRepresentation(GoType.parse("*hxrt.Function[func(int) int]")), "same callback failed to intern");
		require(GoType.parse("func(byte) rune").sameRepresentation(GoType.parse("func(uint8) int32")), "Go aliases split the same signature");
		require(GoType.parse("func(any)").sameRepresentation(GoType.parse("func(interface{})")), "empty interface differs from any");
		require(!GoType.parse("func(...any)").sameRepresentation(GoType.parse("func([]any)")), "variadic policy was erased");
		require(!GoType.parse("*a.Item").sameRepresentation(GoType.parse("*b.Item")), "nominal owners collapsed");
		var signature = concrete.functionCarrierSignature();
		require(signature != null && signature.params.length == 1 && signature.results.length == 1, "callback shape missing");
		require(signature.params[0].sameRepresentation(GoType.parse("int")), "callback argument changed");
		signature.params.resize(0);
		require(concrete.functionCarrierSignature().params.length == 1, "signature query leaked mutable type storage");
		require(GoType.parse("*other.Function[func(int) int]").functionCarrierSignature() == null, "foreign generic mistaken for Haxe carrier");
		Sys.println("ok");
	}
}
