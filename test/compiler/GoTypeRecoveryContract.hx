import reflaxe.go.ast.GoType;
import reflaxe.go.compiler.GoRecoveryDispatch;

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

		// Helper names come from these keys, so equal representations must share one
		// spelling and different representations must never share one.
		require(GoType.parse("func(byte) rune").representationKey() == GoType.parse("func(uint8) int32").representationKey(),
			"alias spelling split a helper key");
		require(GoType.parse("func(interface{})").representationKey() == GoType.parse("func(any)").representationKey(), "empty interface split a helper key");
		require(GoType.parse("interface{B(); A() int}").representationKey() == GoType.parse("interface{A() int; B()}").representationKey(),
			"method order split a helper key");
		require(concrete.representationKey() != erased.representationKey(), "different callable payloads shared a helper key");
		require(GoType.parse("func(...any)").representationKey() != GoType.parse("func([]any)").representationKey(),
			"variadic policy was erased from a helper key");
		require(GoType.parse("func(int) int").representationKey() != GoType.parse("func(int, int)").representationKey(),
			"result and parameter lists shared a helper key");
		require(GoType.parse("func(byte)").render() == "func(byte)", "printer lost the source alias spelling");

		// Entry names are stable keys; a key shaped like a probe suffix stays distinct.
		require(GoRecoveryDispatch.entrySymbol("hx_recover_class_", "app__Shape") == "hx_recover_class_app__Shape", "helper name was not its type key");
		var probeShaped = GoRecoveryDispatch.entrySymbol("hx_recover_class_", "app__Shape__probe_0");
		require(probeShaped != "hx_recover_class_app__Shape" + "__probe_0", "helper name collided with a probe");
		require(probeShaped != GoRecoveryDispatch.entrySymbol("hx_recover_class_", "app__Shape__probe_0_"), "probe escape merged two keys");
		Sys.println("ok");
	}
}
